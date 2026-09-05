package node

import (
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/davinci-dkg/log"
	"github.com/vocdoni/davinci-dkg/web3"
)

// A poll cycle used to read the chain head in every subsystem (liveness,
// auto-create, each lifecycle step, the ciphertext scan, the service loop)
// and re-read records that cannot change (a Live epoch, an application's
// policy, its pool key) for every slot on every tick. Against a public RPC
// endpoint that is what exhausted the quota, so the tick now reads the head
// once, shares it, skips the chain reads altogether when no block arrived
// since the last complete tick and keeps every immutable record in memory.
// This file holds the per-tick context, the immutable-record caches and the
// per-method RPC accounting the node logs.

// tickCtx is what one poll cycle shares between its steps: the head (number
// and timestamp) read once at the start, the epochs read this tick that can
// still change (CommitteeSelection / KeyAssembly, which the lifecycle and the
// early-create check both look at) and whether any read failed, in which case
// the head is not recorded as fully processed and the next tick repeats the
// work even if no block arrived.
type tickCtx struct {
	head     uint64
	headTime uint64 // timestamp of head, for decryption windows
	epochs   map[[12]byte]epochView
	failed   bool
}

// headReader is the slice of the chain the tick needs to start;
// *web3.PooledBackend implements it.
type headReader interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*ethtypes.Header, error)
}

// beginTick reads the head once (number and timestamp, one
// eth_getBlockByNumber) and returns the tick's context — or nil when the head
// is the one the last complete tick already processed: nothing on chain can
// have changed since, so the whole tick is skipped except for settling the
// transactions in flight.
func (n *Node) beginTick(ctx context.Context, chain headReader) (*tickCtx, error) {
	hdr, err := chain.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	head := hdr.Number.Uint64()
	if n.headSeen && head == n.lastHead {
		return nil, nil
	}
	return &tickCtx{head: head, headTime: hdr.Time, epochs: make(map[[12]byte]epochView)}, nil
}

// endTick records the tick's head as processed unless a read failed along
// the way, so a tick that hit an RPC error is repeated at the same head.
func (n *Node) endTick(tc *tickCtx) {
	if tc.failed {
		return
	}
	n.lastHead, n.headSeen = tc.head, true
}

// refreshHead re-reads the head after a step that waited for a transaction
// (a contribution, a finalization): the rest of the tick then compares
// deadlines against the current block rather than the one the tick began at.
func (n *Node) refreshHead(ctx context.Context, tc *tickCtx) {
	hdr, err := n.contracts.Client().HeaderByNumber(ctx, nil)
	if err != nil {
		log.Warnw("tick: cannot refresh the chain head", "err", err)
		return
	}
	tc.head, tc.headTime = hdr.Number.Uint64(), hdr.Time
}

// epochTerminal reports whether an epoch record can never change again: Live
// (every pool key and share root stored by finalizeEpoch; there is no
// transition out of Live), Aborted and Completed.
func epochTerminal(status uint8) bool {
	return status == epochLive || status == epochAborted || status == epochCompleted
}

// maxCachedEpochs bounds the immutable-epoch cache; past it the epoch with
// the lowest nonce is dropped (and costs one read if it is ever needed
// again). The lookback window, the closed epoch just outside it and every
// epoch with a tracked ciphertext fit comfortably.
const maxCachedEpochs = 256

// epoch returns an epoch's record: from the immutable cache when the epoch is
// Live, Aborted or Completed, from this tick's reads when it was already read
// this tick, else from the chain (one eth_call), storing it in whichever of
// the two the status allows.
func (n *Node) epoch(ctx context.Context, tc *tickCtx, chain epochReader, epochID [12]byte) (epochView, error) {
	if e, ok := n.epochCache[epochID]; ok {
		return e, nil
	}
	if tc != nil {
		if e, ok := tc.epochs[epochID]; ok {
			return e, nil
		}
	}
	e, err := chain.GetEpoch(ctx, epochID)
	if err != nil {
		return e, err
	}
	if epochTerminal(e.Status) {
		n.cacheEpoch(epochID, e)
	} else if tc != nil {
		tc.epochs[epochID] = e
	}
	return e, nil
}

func (n *Node) cacheEpoch(epochID [12]byte, e epochView) {
	if len(n.epochCache) >= maxCachedEpochs {
		var oldest [12]byte
		oldestNonce := ^uint64(0)
		for id := range n.epochCache {
			if nonce := binary.BigEndian.Uint64(id[4:]); nonce < oldestNonce {
				oldest, oldestNonce = id, nonce
			}
		}
		delete(n.epochCache, oldest)
	}
	n.epochCache[epochID] = e
}

// cachingEpochReader is the epochReader the lifecycle scan walks with: every
// read goes through Node.epoch, so a closed epoch is read once for the life
// of the process and an open one once per tick.
type cachingEpochReader struct {
	n     *Node
	tc    *tickCtx
	chain epochReader
}

func (r cachingEpochReader) GetEpoch(ctx context.Context, epochID [12]byte) (web3.EpochView, error) {
	return r.n.epoch(ctx, r.tc, r.chain, epochID)
}

// epochPrefixValue returns the manager's EPOCH_PREFIX, an immutable read once.
func (n *Node) epochPrefixValue(ctx context.Context) (uint32, error) {
	if n.prefixKnown {
		return n.epochPrefix, nil
	}
	prefix, err := n.manager.EPOCHPREFIX(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, fmt.Errorf("epoch prefix: %w", err)
	}
	n.epochPrefix, n.prefixKnown = prefix, true
	return prefix, nil
}

// choresEveryTicks is how often the informational chores run: the wallet
// balance line and the RPC histogram, both once every this many polls.
const choresEveryTicks = 20

// logRPCStats logs the JSON-RPC requests sent since the previous histogram,
// per method, averaged over the ticks in between.
func (n *Node) logRPCStats() {
	cur := n.contracts.Pool().Snapshot()
	ticks := n.ticks - n.rpcSeenAt
	if ticks == 0 {
		return
	}
	log.Infow("rpc calls per tick: " + rpcHistogram(n.rpcSeen, cur, ticks))
	n.rpcSeen, n.rpcSeenAt = cur, n.ticks
}

// rpcHistogram formats the per-method request counts accumulated between two
// snapshots as per-tick averages, most frequent first:
// "eth_call=7.5 eth_getLogs=1.0 eth_getBlockByNumber=1.0".
func rpcHistogram(prev, cur map[string]uint64, ticks uint64) string {
	type entry struct {
		method string
		count  uint64
	}
	entries := make([]entry, 0, len(cur))
	for method, count := range cur {
		if delta := count - prev[method]; delta > 0 {
			entries = append(entries, entry{method, delta})
		}
	}
	if len(entries) == 0 {
		return "none"
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].count != entries[j].count {
			return entries[i].count > entries[j].count
		}
		return entries[i].method < entries[j].method
	})
	parts := make([]string, len(entries))
	for i, e := range entries {
		parts[i] = fmt.Sprintf("%s=%.1f", e.method, float64(e.count)/float64(ticks))
	}
	return strings.Join(parts, " ")
}
