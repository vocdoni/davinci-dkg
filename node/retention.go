package node

import (
	"bytes"
	"cmp"
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/finalizer"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// A member rebuilds its share of a pool key the first time an application
// claims the key, often days after the epoch went Live, from every accepted
// dealer's submitContribution calldata. That calldata is located through the
// ContributionSubmitted log, and many endpoints refuse log ranges that old
// (publicnode answers 403 beyond about 10,000 blocks). So once an epoch is
// Live a member keeps every accepted dealer's calldata in the on-disk cache
// while the logs are recent; the share builds then read the cache.

// keepFetchesPerTick bounds how many dealers' calldata retention fetches in
// one tick. A fetch can cost an event-log scan and, on an endpoint slow to
// serve old ranges, its HTTP timeout; a backlog (an upgrade with several
// uncached Live epochs) is worked off over later ticks instead of holding
// up the next one.
const keepFetchesPerTick = 1

// keepPassBudget bounds the elapsed time of one retention pass, its record
// reads and dealer fetches included. Each pooled read shares it among the
// endpoints (web3.WithFairAttempts); what it cuts short is a failed
// attempt, retried after the usual backoff.
const keepPassBudget = 10 * time.Second

// retention is a Live epoch whose contributions this member keeps.
type retention struct {
	epoch    epochView
	selected []common.Address
	myIdx    uint16
}

// queueRetention queues a Live epoch this node is a member of for
// retainContributions. Without a cache there is nothing to keep it in. An
// incomplete committee is an error (see selected), so the epoch is visited
// again rather than finished as one this node has no share in.
func (n *Node) queueRetention(ctx context.Context, chain committeeReader, epochID [12]byte, epoch epochView) error {
	if !n.contribCache.enabled() {
		return nil
	}
	if _, queued := n.retain[epochID]; queued {
		return nil
	}
	selected, err := n.selected(ctx, chain, epochID, epoch.Policy.CommitteeSize)
	if err != nil {
		return err
	}
	if idx := myIndex(selected, n.address); idx != 0 {
		n.retain[epochID] = retention{epoch: epoch, selected: selected, myIdx: idx}
	}
	return nil
}

// retainContributions makes one attempt per tick on the oldest queued
// epoch that is not backing off, at the end of the tick: its records are
// read, valid cache entries are checked and at most keepFetchesPerTick
// dealers are fetched, all within keepPassBudget. The epoch leaves the queue
// once every accepted dealer's calldata is kept; a failed attempt backs it
// off. Retention does not run at all in a tick whose epoch discovery failed
// (the next tick retries that first), nor while an epoch is still selecting
// its committee or assembling its key, so a slow history read never delays a
// slot claim, a contribution or a finalization.
func (n *Node) retainContributions(ctx context.Context, tc *tickCtx, client shareChain) {
	if len(n.retain) == 0 || tc.failed || n.keyGenerationActive(tc) {
		return
	}
	ctx, cancel := context.WithTimeout(ctx, cmp.Or(n.keepBudget, keepPassBudget))
	defer cancel()
	// Each endpoint gets a share of the budget: one that stalls on old log
	// ranges is skipped for its cooldown instead of spending every pass.
	ctx = web3.WithFairAttempts(ctx)
	ids := slices.SortedFunc(maps.Keys(n.retain), func(a, b [12]byte) int { return bytes.Compare(a[:], b[:]) })
	for _, epochID := range ids {
		if b := n.keepRetry[epochID]; b != nil && !b.due() {
			continue
		}
		_, done, err := n.keepContributions(ctx, client, epochID, n.retain[epochID], keepFetchesPerTick)
		if err != nil {
			b := n.keepRetry[epochID]
			if b == nil {
				b = &serviceBackoff{}
				n.keepRetry[epochID] = b
			}
			b.fail()
			log.Warnw("keeping contribution calldata failed", "epoch", roundHex(epochID), "retryInTicks", b.wait, "err", err)
			return
		}
		if done {
			delete(n.retain, epochID)
			delete(n.keepRetry, epochID)
			log.Infow("contribution calldata kept", "epoch", roundHex(epochID))
		}
		return
	}
}

// keyGenerationActive reports whether an epoch read this tick and not
// finished for this node can still need a slot claim, a contribution or a
// finalization: in CommitteeSelection before its selection deadline, or in
// KeyAssembly before the last slot of its finalize rotation.
func (n *Node) keyGenerationActive(tc *tickCtx) bool {
	for id, e := range tc.epochs {
		if n.terminal[id] {
			continue
		}
		switch e.Status {
		case epochCommitteeSelection:
			if tc.head < e.Policy.CommitteeSelectionDeadlineBlock {
				return true
			}
		case epochKeyAssembly:
			if tc.head <= n.staggerWaitUntil(e.Policy.LiveNotBeforeBlock, uint64(e.Policy.CommitteeSize)) {
				return true
			}
		}
	}
	return false
}

// keepContributions makes sure the on-disk cache holds every accepted
// committee member's submitContribution calldata for the epoch, this node's
// own included, each checked as a share rebuild checks it, for every pool
// key: the commitments reproduce the dealer's stored commitmentsHash and the
// share dealt to myIdx opens them. Valid cache entries cost no request; at
// most maxFetches dealers are fetched, and done is false while some remain.
func (n *Node) keepContributions(
	ctx context.Context, client shareChain, epochID [12]byte, r retention, maxFetches int,
) (fetched int, done bool, err error) {
	layout, err := contribution.NewLayout(int(r.epoch.Policy.Threshold), int(r.epoch.Policy.CommitteeSize))
	if err != nil {
		return 0, false, fmt.Errorf("epoch policy: %w", err)
	}
	if r.myIdx == 0 || int(r.myIdx) > layout.CommitteeSize {
		return 0, false, fmt.Errorf("member index %d outside the committee of %d", r.myIdx, layout.CommitteeSize)
	}
	records, err := n.acceptedRecords(ctx, client, epochID, r.epoch, r.selected)
	if err != nil {
		return 0, false, err
	}
	roundHash := roundScalar(epochID)
	for i, addr := range r.selected {
		rec := records[i]
		if !rec.Accepted {
			continue
		}
		contribIdx := uint16(i + 1)
		check := func(data []byte) error {
			tr, err := decodeContribution(data, layout)
			if err != nil {
				return err
			}
			if err := checkCommitments(tr, layout, contribIdx, roundHash, rec.CommitmentsHash); err != nil {
				return err
			}
			for key := range uint8(contribution.MaxKeys) {
				if _, err := openShare(tr, contribIdx, r.myIdx, key, roundHash, n.bjjSecret); err != nil {
					return err
				}
			}
			return nil
		}
		if n.cachedContribution(epochID, addr, check) {
			continue
		}
		if fetched == maxFetches {
			return fetched, false, nil
		}
		fetched++
		if err := n.fetchContribution(ctx, client, epochID, addr, r.epoch.SeedBlock, check); err != nil {
			return fetched, false, fmt.Errorf("contribution of %s (idx %d): %w", addr.Hex(), contribIdx, err)
		}
	}
	return fetched, true, nil
}

// acceptedRecords reads the committee's contribution records (entry i is
// member i+1's) and returns them only when they are the Live epoch's whole
// accepted set: a full committee, every accepted record under its member's
// index, and exactly the epoch's ContributionCount accepted, the set
// finalizeEpoch froze (the checks finalizer.reconstruct makes). An endpoint
// behind the chain shows fewer; that is an error to retry, never a set to
// build a share or finish retention on.
func (n *Node) acceptedRecords(
	ctx context.Context, client bind.ContractCaller, epochID [12]byte, epoch epochView, selected []common.Address,
) ([]gtypes.DKGTypesContributionRecord, error) {
	if len(selected) != int(epoch.Policy.CommitteeSize) {
		return nil, fmt.Errorf("committee has %d of its %d members: the endpoint is behind the chain, retrying",
			len(selected), epoch.Policy.CommitteeSize)
	}
	records, err := finalizer.ContributionRecords(ctx, client, n.managerAddr, nil, epochID, selected)
	if err != nil {
		return nil, err
	}
	accepted := 0
	for i, rec := range records {
		if !rec.Accepted {
			continue
		}
		if rec.ContributorIndex != uint16(i+1) {
			return nil, fmt.Errorf("contribution of %s is stored under index %d, committee position is %d",
				selected[i].Hex(), rec.ContributorIndex, i+1)
		}
		accepted++
	}
	if accepted != int(epoch.ContributionCount) {
		return nil, fmt.Errorf("%d accepted contributions, the epoch has %d: the endpoint is behind the chain, retrying",
			accepted, epoch.ContributionCount)
	}
	return records, nil
}

// cachedContribution reports whether the cache holds contributor's calldata
// and it passes check; an entry that fails is logged and treated as a miss.
func (n *Node) cachedContribution(epochID [12]byte, contributor common.Address, check func(data []byte) error) bool {
	data, ok := n.contribCache.Get(epochID, contributor)
	if !ok {
		return false
	}
	if err := check(data); err != nil {
		log.Warnw("cached contribution calldata unusable, refetching",
			"epoch", roundHex(epochID), "contributor", contributor, "err", err)
		return false
	}
	return true
}

// errCacheWrite marks a fetched, valid calldata the cache could not store:
// a share build goes on without the entry, retention does not count it kept.
var errCacheWrite = errors.New("contribution cache write failed")

// fetchContribution locates contributor's submitContribution transaction
// through the ContributionSubmitted log from fromBlock and stores its
// calldata once it passes check; a failed write is errCacheWrite.
func (n *Node) fetchContribution(
	ctx context.Context,
	client calldataReader,
	epochID [12]byte,
	contributor common.Address,
	fromBlock uint64,
	check func(data []byte) error,
) error {
	data, err := finalizer.ContributionCalldata(ctx, client, n.manager, epochID, contributor, fromBlock)
	if err != nil {
		return err
	}
	if err := check(data); err != nil {
		return err
	}
	if err := n.contribCache.Store(epochID, contributor, data); err != nil {
		return fmt.Errorf("%w: %w", errCacheWrite, err)
	}
	return nil
}
