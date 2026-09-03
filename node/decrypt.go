package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/log"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
)

const (
	// logRangeBlocks bounds each eth_getLogs call; most public RPC providers
	// cap the range at 10k blocks.
	logRangeBlocks = 10_000
	// reorgDepthBlocks is how many already-scanned blocks every tick re-scans
	// so a CiphertextSubmitted event that lands in a reorged block is still
	// picked up. The pending map dedupes the overlap.
	reorgDepthBlocks = 32
	// maxPendingCiphertexts bounds the pending set; beyond it the oldest
	// entries are dropped so a spammer cannot grow the node's memory and
	// per-tick RPC fan-out without bound.
	maxPendingCiphertexts = 1024
	// maxServiceBackoffTicks caps the exponential per-ciphertext backoff.
	maxServiceBackoffTicks = 64
)

// ctKey identifies one ciphertext slot: (epoch, application, index).
type ctKey struct {
	epoch [12]byte
	aid   [32]byte
	idx   uint16
}

func (k ctKey) String() string {
	return fmt.Sprintf("epoch=%x aid=%x idx=%d", k.epoch, k.aid[:4], k.idx)
}

// ciphertext is a decoded CiphertextSubmitted event payload.
type ciphertext struct {
	c1, c2 nodetypes.CurvePoint
	block  uint64 // block the event was emitted in
	seq    uint64 // discovery order, used to evict the oldest entry
}

// serviceBackoff skips a failing ciphertext for 1, 2, 4 … ticks (capped at
// maxServiceBackoffTicks) between attempts.
type serviceBackoff struct {
	wait uint32 // ticks to skip after the most recent failure
	skip uint32 // ticks still to skip
}

func (b *serviceBackoff) fail() {
	b.wait = min(max(2*b.wait, 1), maxServiceBackoffTicks)
	b.skip = b.wait
}

// due reports whether the entry should be serviced this tick and consumes
// one skipped tick otherwise.
func (b *serviceBackoff) due() bool {
	if b.skip > 0 {
		b.skip--
		return false
	}
	return true
}

// scanStart returns the first block the next CiphertextSubmitted scan must
// cover: head − lookback on the first run, otherwise the tail of the previous
// scan re-covered by reorgDepthBlocks.
func scanStart(lastScan, head, lookback uint64) uint64 {
	if lastScan == 0 {
		if head > lookback {
			return head - lookback
		}
		return 0
	}
	if lastScan+1 > reorgDepthBlocks {
		return lastScan + 1 - reorgDepthBlocks
	}
	return 0
}

// scanCiphertexts pulls every CiphertextSubmitted event since the last scan
// (or since head-lookback on first run) into the pending set.
func (n *Node) scanCiphertexts(ctx context.Context) error {
	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("head: %w", err)
	}
	from := scanStart(n.lastCtScan, head, n.lookback)
	for start := from; start <= head; start += logRangeBlocks {
		end := min(start+logRangeBlocks-1, head)
		it, err := n.manager.FilterCiphertextSubmitted(&bind.FilterOpts{Context: ctx, Start: start, End: &end}, nil, nil, nil)
		if err != nil {
			return fmt.Errorf("filter CiphertextSubmitted [%d,%d]: %w", start, end, err)
		}
		for it.Next() {
			ev := it.Event
			key := ctKey{ev.EpochId, ev.Aid, ev.CiphertextIndex}
			if _, seen := n.pending[key]; seen {
				continue
			}
			if _, parked := n.parked[key]; parked {
				continue
			}
			if _, done := n.served[key]; done {
				continue
			}
			n.trackCiphertext(key, &ciphertext{
				c1:    nodetypes.CurvePoint{X: new(big.Int).Set(ev.C1x), Y: new(big.Int).Set(ev.C1y)},
				c2:    nodetypes.CurvePoint{X: new(big.Int).Set(ev.C2x), Y: new(big.Int).Set(ev.C2y)},
				block: ev.Raw.BlockNumber,
			})
			log.Infow("ciphertext discovered", "ct", key.String(), "block", ev.Raw.BlockNumber)
		}
		err = it.Error()
		_ = it.Close()
		if err != nil {
			return fmt.Errorf("iterate CiphertextSubmitted: %w", err)
		}
		if err := n.wakeParked(ctx, start, end); err != nil {
			return err
		}
		n.lastCtScan = end
	}
	// Served slots older than the re-scan window can no longer be rediscovered.
	for key, block := range n.served {
		if block < from {
			delete(n.served, key)
		}
	}
	return nil
}

// trackCiphertext adds a slot to the pending set, evicting the oldest entry
// when the set is full.
//
// ponytail: an evicted slot is not remembered in `served`, so a flood of
// more than maxPendingCiphertexts events inside the reorg window is
// rediscovered and re-evicted every tick until it ages out of the window.
func (n *Node) trackCiphertext(key ctKey, ct *ciphertext) {
	n.ctSeq++
	ct.seq = n.ctSeq
	if len(n.pending) >= maxPendingCiphertexts {
		var oldest ctKey
		oldestSeq := ^uint64(0)
		for k, c := range n.pending {
			if c.seq < oldestSeq {
				oldest, oldestSeq = k, c.seq
			}
		}
		log.Warnw("pending ciphertext set is full — dropping the oldest entry",
			"dropped", oldest.String(), "cap", maxPendingCiphertexts)
		n.forget(oldest)
	}
	n.pending[key] = ct
}

// forget drops every piece of per-slot state.
// park moves a slot that only waits for an organizer share out of the
// per-tick service loop. Epochs stay Live on chain indefinitely, so without
// this every withheld share would cost every node a few RPC scans per tick
// forever; a parked slot costs nothing until OrganizerShareSubmitted wakes it.
func (n *Node) park(key ctKey) {
	ct, ok := n.pending[key]
	if !ok {
		return
	}
	delete(n.pending, key)
	delete(n.backoff, key)
	if len(n.parked) >= maxParkedCiphertexts {
		var oldest ctKey
		oldestSeq := ^uint64(0)
		for k, c := range n.parked {
			if c.seq < oldestSeq {
				oldest, oldestSeq = k, c.seq
			}
		}
		delete(n.parked, oldest)
	}
	n.parked[key] = ct
	log.Infow("parked until the organizer posts a share", "ct", key.String(), "parked", len(n.parked))
}

// wakeParked returns to the pending set every parked slot that received an
// organizer share in [start, end].
func (n *Node) wakeParked(ctx context.Context, start, end uint64) error {
	if len(n.parked) == 0 {
		return nil
	}
	it, err := n.appManager.FilterOrganizerShareSubmitted(&bind.FilterOpts{Context: ctx, Start: start, End: &end}, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("filter OrganizerShareSubmitted [%d,%d]: %w", start, end, err)
	}
	defer func() { _ = it.Close() }()
	for it.Next() {
		ev := it.Event
		key := ctKey{ev.EpochId, ev.Aid, ev.CiphertextIndex}
		if ct, ok := n.parked[key]; ok {
			delete(n.parked, key)
			n.pending[key] = ct
			log.Infow("organizer share seen — resuming the slot", "ct", key.String(), "block", ev.Raw.BlockNumber)
		}
	}
	if err := it.Error(); err != nil {
		return fmt.Errorf("iterate OrganizerShareSubmitted: %w", err)
	}
	return nil
}

func (n *Node) forget(key ctKey) {
	delete(n.pending, key)
	delete(n.partialDone, key)
	delete(n.backoff, key)
	delete(n.inflight, key)
	delete(n.parked, key)
	n.jobsMu.Lock()
	delete(n.combineJobs, key)
	n.jobsMu.Unlock()
}

// inflightTx is a partial or combine transaction this node has sent for a
// slot and not yet seen mined. Sending without blocking on the receipt is
// what lets one tick serve every pending ciphertext instead of one per
// block; the next tick settles it from the receipt.
type inflightTx struct {
	hash      common.Hash
	combine   bool
	plaintext *big.Int // combine only, for the log line
	sent      time.Time
}

// maxParkedCiphertexts bounds the slots kept aside for a share that may never
// come; beyond it the oldest is dropped (rediscovery re-parks it if it is
// still inside the scan window).
const maxParkedCiphertexts = 4096

// criticalYield bounds how long a combine job defers to an in-progress
// contribution or finalization before running anyway.
const criticalYield = 3 * time.Minute

// inflightTimeout is how long a sent tx may stay unmined before the slot is
// retried from scratch (the tx manager keeps rebroadcasting meanwhile).
const inflightTimeout = 2 * time.Minute

// laterWaveDue says whether a member of wave > 0 may post its partial: only
// once its wave's delay since the ciphertext has passed *and* the earlier
// waves have stopped landing partials for staggerBlocks. Under load the
// first wave keeps landing one partial per block for many blocks, and
// counting from the ciphertext alone would make every later wave pile on.
func laterWaveDue(head, ctBlock, lastPartialBlock, wave uint64) bool {
	return head >= ctBlock+wave*staggerBlocks && head >= lastPartialBlock+staggerBlocks
}

// settleInflight looks at the receipt of the slot's in-flight tx. It returns
// pending=true while the tx is unmined (nothing else to do this tick) and
// done=true once a combine of ours is mined. A mined-but-reverted tx or a
// timeout clears the record so the normal path retries.
func (n *Node) settleInflight(ctx context.Context, key ctKey, fl inflightTx) (pending, done bool, err error) {
	receipt, err := n.contracts.Client().TransactionReceipt(ctx, fl.hash)
	switch {
	case errors.Is(err, ethereum.NotFound):
		if time.Since(fl.sent) < inflightTimeout {
			return true, false, nil
		}
		delete(n.inflight, key)
		return false, false, fmt.Errorf("tx %s unmined after %s, retrying", fl.hash.Hex(), inflightTimeout)
	case err != nil:
		return true, false, fmt.Errorf("receipt %s: %w", fl.hash.Hex(), err)
	}
	delete(n.inflight, key)
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		if fl.combine {
			// Almost always a lost race; the on-chain record settles it.
			if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
				return false, true, nil
			}
		}
		return false, false, fmt.Errorf("tx %s reverted, retrying", fl.hash.Hex())
	}
	if fl.combine {
		log.Infow("decryption combined", "ct", key.String(), "plaintext", fl.plaintext.String(), "tx", fl.hash.Hex())
		return false, true, nil
	}
	n.partialDone[key] = true
	return false, false, nil
}

// tickCache memoises the chain reads shared by every pending ciphertext
// within one tick: one head and one getEpoch per epoch instead of one per
// slot.
type tickCache struct {
	head    uint64
	headErr error
	epochs  map[[12]byte]epochView
	apps    map[appKey]nodetypes.CurvePoint // (epoch, aid) → PK_org
}

// appKey identifies one application: (epoch, aid).
type appKey struct {
	epoch [12]byte
	aid   [32]byte
}

func (n *Node) newTickCache(ctx context.Context) *tickCache {
	head, err := n.contracts.Client().BlockNumber(ctx)
	return &tickCache{
		head:    head,
		headErr: err,
		epochs:  make(map[[12]byte]epochView),
		apps:    make(map[appKey]nodetypes.CurvePoint),
	}
}

func (n *Node) cachedEpoch(ctx context.Context, tc *tickCache, epochID [12]byte) (epochView, error) {
	if e, ok := tc.epochs[epochID]; ok {
		return e, nil
	}
	e, err := n.contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return e, fmt.Errorf("get epoch: %w", err)
	}
	tc.epochs[epochID] = e
	return e, nil
}

// serviceCiphertexts advances every pending ciphertext and drops the ones
// that need no further work. Slots whose last attempt failed are skipped
// for an exponentially growing number of ticks.
func (n *Node) serviceCiphertexts(ctx context.Context) {
	if len(n.pending) == 0 {
		return
	}
	tc := n.newTickCache(ctx)
	for key, ct := range n.pending {
		if b := n.backoff[key]; b != nil && !b.due() {
			continue
		}
		done, err := n.serviceCiphertext(ctx, tc, key, ct)
		if err != nil {
			b := n.backoff[key]
			if b == nil {
				b = &serviceBackoff{}
				n.backoff[key] = b
			}
			b.fail()
			log.Warnw("ciphertext service failed", "ct", key.String(), "retryInTicks", b.wait, "err", err)
		} else {
			delete(n.backoff, key)
		}
		if done {
			n.forget(key)
			n.served[key] = ct.block
		}
	}
}

// serviceCiphertext submits this node's partial decryption (if it sits on
// the committee) and then tries to combine. Returns done=true once the slot
// is combined on-chain, can never be, or is none of this node's business.
func (n *Node) serviceCiphertext(ctx context.Context, tc *tickCache, key ctKey, ct *ciphertext) (bool, error) {
	if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
		return true, nil
	}
	if fl, ok := n.inflight[key]; ok {
		pending, done, err := n.settleInflight(ctx, key, fl)
		if pending || done || err != nil {
			return done, err
		}
	}
	epoch, err := n.cachedEpoch(ctx, tc, key.epoch)
	if err != nil {
		return false, err
	}
	if epoch.Status != epochLive {
		// Partials and combines are only accepted while the epoch is Live;
		// an aborted or completed epoch can never serve this slot.
		return epoch.Status == epochAborted || epoch.Status == epochCompleted, nil
	}
	selected, err := n.selected(ctx, key.epoch)
	if err != nil {
		return false, err
	}
	if n.taintedApps[appKey{epoch: key.epoch, aid: key.aid}] {
		return true, nil
	}
	idx := myIndex(selected, n.address)
	if idx == 0 {
		// Only committee members hold a share of sk_ep, and only they take
		// part in the combine rotation; nothing for this node to do.
		log.Debugw("ciphertext belongs to an epoch this node is not a member of — ignoring", "ct", key.String())
		return true, nil
	}
	if !n.partialDone[key] {
		// Only t partials are ever needed. The first t members of a
		// seed-derived rotation respond at once; each later wave of t steps
		// in staggerBlocks later and only if partials are still missing, so
		// an honest ciphertext costs t partials rather than n and a hostile
		// submitter cannot make the whole committee spend.
		threshold := uint64(epoch.Policy.Threshold)
		wave := staggerSlot(epoch.Seed, uint64(key.idx), idx, uint16(len(selected))) / max(threshold, 1)
		if wave > 0 {
			if tc.headErr != nil {
				return false, fmt.Errorf("read head: %w", tc.headErr)
			}
			if tc.head < ct.block+wave*staggerBlocks {
				return false, nil
			}
			// With fewer than t partials the scan sees all of them, so
			// lastBlock is the newest partial for this slot.
			idxs, _, lastBlock, err := n.acceptedPartials(ctx, key, epoch.SeedBlock, tc.head, epoch.Policy.Threshold)
			if err != nil {
				return false, err
			}
			if uint64(len(idxs)) >= threshold {
				n.partialDone[key] = true
			} else if !laterWaveDue(tc.head, ct.block, lastBlock, wave) {
				return false, nil
			}
		}
	}
	if !n.partialDone[key] {
		toxic, err := n.submitPartial(ctx, key, ct, idx, epoch, selected)
		if err != nil {
			return false, err
		}
		if toxic {
			return true, nil
		}
	}
	if tc.headErr != nil {
		return false, fmt.Errorf("read head: %w", tc.headErr)
	}
	return n.tryCombine(ctx, tc, key, ct, epoch, idx, uint16(len(selected)), tc.head)
}

// selected caches the committee of a Live epoch (it never changes).
func (n *Node) selected(ctx context.Context, epochID [12]byte) ([]common.Address, error) {
	if s, ok := n.selectedCache[epochID]; ok {
		return s, nil
	}
	s, err := n.contracts.SelectedParticipants(ctx, epochID)
	if err != nil {
		return nil, fmt.Errorf("selected participants: %w", err)
	}
	n.selectedCache[epochID] = s
	return s, nil
}

// submitPartial posts δ_i = d_i·C1 with its DLEQ proof. Returns toxic=true
// when the ciphertext is malformed and must never be decrypted.
func (n *Node) submitPartial(
	ctx context.Context,
	key ctKey,
	ct *ciphertext,
	idx uint16,
	epoch epochView,
	selected []common.Address,
) (bool, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	if rec, err := n.manager.GetPartialDecryption(callOpts, key.epoch, key.aid, idx, key.idx); err == nil && rec.Accepted {
		n.partialDone[key] = true
		return false, nil
	}

	// Refuse small-order / off-curve ciphertexts before touching the share:
	// δ_i = d_i·C1 for a cofactor point would leak d_i mod 8 on-chain. The
	// contract deliberately skips the prime-subgroup check (it costs ~2 M
	// gas), so this is the only place it happens — it is load-bearing, not
	// belt-and-braces.
	if err := group.ValidateCiphertext(ct.c1, ct.c2); err != nil {
		log.Warnw("rejecting toxic ciphertext — refusing partial decryption", "ct", key.String(), "err", err)
		return true, nil
	}

	dShare, err := n.buildPrivateShare(ctx, key.epoch, idx, selected, epoch, callOpts)
	if err != nil {
		return false, fmt.Errorf("build private share: %w", err)
	}
	nonce, err := randomScalars(1)
	if err != nil {
		return false, err
	}
	witness, pi, err := partialdecrypt.BuildWitness(partialdecrypt.Assignment{
		RoundHash:        roundScalar(key.epoch),
		Aid:              new(big.Int).SetBytes(key.aid[:]),
		CtIdx:            new(big.Int).SetUint64(uint64(key.idx)),
		ParticipantIndex: idx,
		Base:             ct.c1,
		Secret:           dShare,
		Nonce:            nonce[0],
	})
	if err != nil {
		return false, fmt.Errorf("build partial decrypt witness: %w", err)
	}
	runtime, err := partialdecrypt.Artifacts.LoadOrSetupForCircuit(ctx, &partialdecrypt.PartialDecryptCircuit{})
	if err != nil {
		return false, fmt.Errorf("load partial decrypt circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return false, fmt.Errorf("prove partial decrypt: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return false, err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return false, err
	}
	dHash := ethcrypto.Keccak256Hash(
		common.LeftPadBytes(pi.Delta.X.Bytes(), 32),
		common.LeftPadBytes(pi.Delta.Y.Bytes(), 32),
	)

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return false, err
	}
	tx, err := n.manager.SubmitPartialDecryption(auth, key.epoch, key.aid, idx, key.idx,
		ct.c1.X, ct.c1.Y, ct.c2.X, ct.c2.Y, dHash, proofBytes, inputBytes)
	if err != nil {
		reason := decodeContractError(err)
		if strings.Contains(reason, "AlreadyPartiallyDecrypted") {
			n.partialDone[key] = true
			return false, nil
		}
		if isPermanentRevert(err) {
			// Retrying the same proof would fail the same way.
			n.partialDone[key] = true
		}
		return false, fmt.Errorf("submit partial decryption: %s", reason)
	}
	n.txm.RecordPending(tx)
	n.inflight[key] = inflightTx{hash: tx.Hash(), sent: time.Now()}
	log.Infow("partial decryption submitted", "ct", key.String(), "index", idx, "tx", tx.Hash().Hex())
	return false, nil
}

// acceptedPartials reads up to `threshold` distinct partials for the slot
// from the event log (the contract only stores their hashes), scanning
// [seedBlock−1, head] in logRangeBlocks chunks and stopping as soon as
// enough are in hand. readyBlock is the block the last of them landed in.
func (n *Node) acceptedPartials(
	ctx context.Context,
	key ctKey,
	seedBlock, head uint64,
	threshold uint16,
) (idxs []uint16, deltas []nodetypes.CurvePoint, readyBlock uint64, err error) {
	start := uint64(0)
	if seedBlock > 0 {
		start = seedBlock - 1
	}
	seen := map[uint16]bool{}
	for from := start; from <= head && len(idxs) < int(threshold); from += logRangeBlocks {
		end := min(from+logRangeBlocks-1, head)
		it, err := n.manager.FilterPartialDecryptionSubmitted(
			&bind.FilterOpts{Context: ctx, Start: from, End: &end},
			[][12]byte{key.epoch}, [][32]byte{key.aid}, nil,
		)
		if err != nil {
			return nil, nil, 0, fmt.Errorf("filter PartialDecryptionSubmitted [%d,%d]: %w", from, end, err)
		}
		for it.Next() && len(idxs) < int(threshold) {
			e := it.Event
			if e.CiphertextIndex != key.idx || seen[e.ParticipantIndex] {
				continue
			}
			seen[e.ParticipantIndex] = true
			idxs = append(idxs, e.ParticipantIndex)
			deltas = append(deltas, nodetypes.CurvePoint{X: new(big.Int).Set(e.DeltaX), Y: new(big.Int).Set(e.DeltaY)})
			readyBlock = max(readyBlock, e.Raw.BlockNumber)
		}
		err = it.Error()
		_ = it.Close()
		if err != nil {
			return nil, nil, 0, fmt.Errorf("iterate PartialDecryptionSubmitted: %w", err)
		}
	}
	return idxs, deltas, readyBlock, nil
}

// organizerShare is the organizer's Δ = sk_org·C1 for one ciphertext slot,
// with the Chaum-Pedersen DLEQ that ties it to the application's registered
// PK_org and the block the share landed in.
type organizerShare struct {
	pkOrg nodetypes.CurvePoint
	delta nodetypes.CurvePoint
	proof dleq.Proof
	block uint64
}

// organizerPK reads (and memoises for the tick) the application's registered
// organizer key. Only registered applications can own a ciphertext — the
// contract rejects submitCiphertext for an unknown aid — so a missing record
// here means the chain is lying or we are looking at the wrong contract.
func (n *Node) organizerPK(ctx context.Context, tc *tickCache, key ctKey) (nodetypes.CurvePoint, error) {
	ak := appKey{epoch: key.epoch, aid: key.aid}
	if pk, ok := tc.apps[ak]; ok {
		return pk, nil
	}
	rec, err := n.appManager.GetApplication(&bind.CallOpts{Context: ctx}, key.epoch, key.aid)
	if err != nil {
		return nodetypes.CurvePoint{}, fmt.Errorf("get application: %w", err)
	}
	if !rec.Exists {
		return nodetypes.CurvePoint{}, fmt.Errorf("application %x is not registered", key.aid)
	}
	pk := nodetypes.CurvePoint{X: new(big.Int).Set(rec.OrganizerPK.X), Y: new(big.Int).Set(rec.OrganizerPK.Y)}
	tc.apps[ak] = pk
	return pk, nil
}

// latestOrganizerShare reads the newest OrganizerShareSubmitted event for the
// slot and verifies its DLEQ against the application's registered PK_org.
//
// ok=false means "not combinable yet": either nobody has posted a share, or
// the newest one does not verify. The contract stores the share words without
// checking them, precisely so a malformed submission cannot brick a
// ciphertext — the organizer can overwrite it, so we simply retry on the next
// tick instead of giving up on the slot.
func (n *Node) latestOrganizerShare(
	ctx context.Context,
	tc *tickCache,
	key ctKey,
	c1 nodetypes.CurvePoint,
	seedBlock, head uint64,
) (organizerShare, bool, error) {
	pkOrg, err := n.organizerPK(ctx, tc, key)
	if err != nil {
		return organizerShare{}, false, err
	}

	start := uint64(0)
	if seedBlock > 0 {
		start = seedBlock - 1
	}
	var share organizerShare
	found := false
	for from := start; from <= head; from += logRangeBlocks {
		end := min(from+logRangeBlocks-1, head)
		it, err := n.appManager.FilterOrganizerShareSubmitted(
			&bind.FilterOpts{Context: ctx, Start: from, End: &end},
			[][12]byte{key.epoch}, [][32]byte{key.aid}, []uint16{key.idx},
		)
		if err != nil {
			return organizerShare{}, false, fmt.Errorf("filter OrganizerShareSubmitted [%d,%d]: %w", from, end, err)
		}
		for it.Next() {
			e := it.Event
			share = organizerShare{
				pkOrg: pkOrg,
				delta: nodetypes.CurvePoint{X: new(big.Int).Set(e.DeltaX), Y: new(big.Int).Set(e.DeltaY)},
				proof: dleq.Proof{
					A1:       nodetypes.CurvePoint{X: new(big.Int).Set(e.A1x), Y: new(big.Int).Set(e.A1y)},
					A2:       nodetypes.CurvePoint{X: new(big.Int).Set(e.A2x), Y: new(big.Int).Set(e.A2y)},
					Response: new(big.Int).Set(e.Z),
				},
				block: e.Raw.BlockNumber,
			}
			found = true
		}
		err = it.Error()
		_ = it.Close()
		if err != nil {
			return organizerShare{}, false, fmt.Errorf("iterate OrganizerShareSubmitted: %w", err)
		}
	}
	if !found {
		log.Debugw("combine: waiting for the organizer share", "ct", key.String())
		return organizerShare{}, false, nil
	}
	if !dleq.VerifyOrganizerShare(key.epoch, key.aid, key.idx, pkOrg, c1, share.delta, share.proof) {
		fingerprint := ethcrypto.Keccak256Hash(
			share.delta.X.Bytes(), share.delta.Y.Bytes(), share.proof.Response.Bytes(),
		)
		if n.badShares[key] != fingerprint { // warn once per distinct bad share, not every tick
			n.badShares[key] = fingerprint
			log.Warnw("organizer share does not verify — waiting for the organizer to resubmit",
				"ct", key.String(), "block", share.block)
		}
		return organizerShare{}, false, nil
	}
	delete(n.badShares, key)
	return share, true, nil
}

// tryCombine interpolates threshold partials, recovers the plaintext by BSGS
// and posts the combine proof. A slot becomes combinable only once BOTH
// `t` partial decryptions and a verifying organizer share are on chain.
// Committee members take turns in a seed-derived rotation (like
// auto-finalize) starting at the block the last of those two landed in, so
// normally a single member pays for the combine; later slots only step in if
// the earlier ones did not.
func (n *Node) tryCombine(
	ctx context.Context,
	tc *tickCache,
	key ctKey,
	ct *ciphertext,
	epoch epochView,
	myIdx, committeeSize uint16,
	head uint64,
) (bool, error) {
	threshold := epoch.Policy.Threshold
	idxs, deltas, readyBlock, err := n.acceptedPartials(ctx, key, epoch.SeedBlock, head, threshold)
	if err != nil {
		return false, err
	}
	if len(idxs) < int(threshold) {
		return false, nil
	}
	share, ready, err := n.latestOrganizerShare(ctx, tc, key, ct.c1, epoch.SeedBlock, head)
	if err != nil {
		return false, err
	}
	if !ready {
		n.park(key)
		return false, nil
	}
	// The rotation is anchored on whichever of the two prerequisites landed
	// last, so an organizer share posted long after the partials does not
	// make every slot eligible at once.
	slot := staggerSlot(epoch.Seed, uint64(key.idx), myIdx, committeeSize)
	if waitUntil := max(readyBlock, share.block) + slot*staggerBlocks; head < waitUntil {
		log.Debugw("combine: waiting for our slot", "ct", key.String(), "slot", slot, "head", head, "waitUntil", waitUntil)
		return false, nil
	}

	// The dlog search and the Groth16 proof take seconds (and a poisoned
	// slot burns the full 2^50 search); they run off the tick loop so one
	// slot never stalls the partials of every other pending ciphertext.
	n.jobsMu.Lock()
	res, running := n.combineJobs[key]
	n.jobsMu.Unlock()
	if !running {
		n.jobsMu.Lock()
		n.combineJobs[key] = nil
		n.jobsMu.Unlock()
		go n.runCombineJob(key, ct, idxs, deltas, share, threshold)
		return false, nil
	}
	if res == nil {
		return false, nil // still computing
	}
	n.jobsMu.Lock()
	delete(n.combineJobs, key)
	n.jobsMu.Unlock()
	if res.taint {
		// Only the application's authorised submitter can produce its
		// ciphertexts, so an undecryptable one is the organizer's doing:
		// serve nothing else of this application in this epoch, which caps
		// what one registration can make the committee compute.
		n.taintedApps[appKey{epoch: key.epoch, aid: key.aid}] = true
		n.saveTaints()
		log.Warnw("application tainted: ignoring its remaining ciphertexts for this epoch", "ct", key.String())
	}
	switch {
	case res.err != nil:
		return res.permanent, res.err
	case res.done:
		return true, nil
	}
	n.inflight[key] = inflightTx{hash: res.tx, combine: true, plaintext: res.plaintext, sent: time.Now()}
	return false, nil
}

// combineResult is what a combine job leaves behind for the tick loop.
type combineResult struct {
	tx        common.Hash
	plaintext *big.Int
	done      bool // the slot was combined by someone else meanwhile
	err       error
	permanent bool // the slot can never be combined by us; stop trying
	taint     bool // the plaintext was out of range: stop serving this application
}

// runCombineJob recovers the plaintext, proves the combine and sends the
// transaction. It only reads its arguments and the chain, and reports back
// through combineJobs; the tick loop owns every other piece of node state.
func (n *Node) runCombineJob(
	key ctKey,
	ct *ciphertext,
	idxs []uint16,
	deltas []nodetypes.CurvePoint,
	share organizerShare,
	threshold uint16,
) {
	// One combine at a time per node: the dlog search saturates every core,
	// and several at once (e.g. re-discovered poisoned slots after a restart)
	// would starve the partials and the RPC client.
	// Contributions and finalizations have a deadline; a search has none.
	// Let the epoch-critical work of the tick loop finish first.
	for wait := time.Duration(0); n.critical.Load() > 0 && wait < criticalYield; wait += 500 * time.Millisecond {
		time.Sleep(500 * time.Millisecond)
	}
	n.combineSem <- struct{}{}
	res := n.combine(key, ct, idxs, deltas, share, threshold)
	<-n.combineSem
	n.jobsMu.Lock()
	n.combineJobs[key] = res
	n.jobsMu.Unlock()
}

func (n *Node) combine(
	key ctKey,
	ct *ciphertext,
	idxs []uint16,
	deltas []nodetypes.CurvePoint,
	share organizerShare,
	threshold uint16,
) *combineResult {
	ctx := context.Background()
	// M·G = C2 − Σ λ_k·δ_k − Δ_org
	combinedEnc, err := ccommon.InterpolatePointsAtZeroNative(ccommon.Uint16sToBigInts(idxs), deltas)
	if err != nil {
		return &combineResult{err: fmt.Errorf("interpolate partials: %w", err)}
	}
	combined, err := group.Decode(combinedEnc)
	if err != nil {
		return &combineResult{err: err}
	}
	c2, err := group.Decode(ct.c2)
	if err != nil {
		return &combineResult{err: err}
	}
	correction, err := group.Decode(share.delta)
	if err != nil {
		return &combineResult{err: fmt.Errorf("decode organizer share: %w", err)}
	}
	negCombined := group.NewPoint()
	negCombined.Neg(combined)
	negCorrection := group.NewPoint()
	negCorrection.Neg(correction)
	mG := group.NewPoint()
	mG.Add(c2, negCombined)
	mG.Add(mG, negCorrection)

	plaintext, err := dlogBSGS(mG)
	if err != nil {
		// Plaintext ≥ MaxDLogPlaintext: retrying can never succeed.
		// An attacker-made ciphertext, not a node fault: warn, don't error.
		log.Warnw("combine: dlog failed, the plaintext is out of range (must be < 2^50)", "ct", key.String(), "err", err)
		return &combineResult{done: true, taint: true}
	}

	witness, pi, err := decryptcombine.BuildWitness(decryptcombine.Assignment{
		RoundHash:          roundScalar(key.epoch),
		Aid:                new(big.Int).SetBytes(key.aid[:]),
		CtIdx:              new(big.Int).SetUint64(uint64(key.idx)),
		DeltaOrg:           share.delta,
		OrganizerPK:        share.pkOrg,
		OrganizerProof:     share.proof,
		Threshold:          threshold,
		CiphertextC1:       ct.c1,
		CiphertextC2:       ct.c2,
		ParticipantIndexes: idxs,
		PartialDecryptions: deltas,
		Plaintext:          plaintext,
	})
	if err != nil {
		return &combineResult{err: fmt.Errorf("build combine witness: %w", err)}
	}
	runtime, err := decryptcombine.Artifacts.LoadOrSetupForCircuit(ctx, &decryptcombine.DecryptCombineCircuit{})
	if err != nil {
		return &combineResult{err: fmt.Errorf("load combine circuit: %w", err)}
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return &combineResult{err: fmt.Errorf("prove combine: %w", err)}
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return &combineResult{err: err}
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return &combineResult{err: err}
	}
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return &combineResult{err: err}
	}

	// A later slot may still race an earlier one that was merely slow;
	// re-check right before paying for the tx so a winner that landed while
	// we were proving saves us gas.
	if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
		return &combineResult{done: true}
	}
	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return &combineResult{err: err}
	}
	tx, err := n.manager.CombineDecryption(auth, key.epoch, key.aid, key.idx,
		common.BigToHash(pi.CombineHash), pi.PlaintextHash, transcriptBytes, proofBytes, inputBytes)
	if err != nil {
		reason := decodeContractError(err)
		if strings.Contains(reason, "AlreadyCombined") {
			return &combineResult{done: true}
		}
		// A mined-but-reverted tx is almost always a lost race; the
		// on-chain re-check at the top of the next attempt settles it.
		if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
			return &combineResult{done: true}
		}
		if isPermanentRevert(err) {
			return &combineResult{err: errors.New("combine rejected on-chain, giving up: " + reason), permanent: true}
		}
		return &combineResult{err: fmt.Errorf("submit combine: %w", err)}
	}
	n.txm.RecordPending(tx)
	log.Infow("combine submitted", "ct", key.String(), "plaintext", plaintext.String(), "tx", tx.Hash().Hex())
	return &combineResult{tx: tx.Hash(), plaintext: plaintext}
}
