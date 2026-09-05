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
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/finalizer"
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
	c1, c2    nodetypes.CurvePoint
	submitter common.Address
	block     uint64 // block the event was emitted in
	seq       uint64 // discovery order, used to evict the oldest entry
	// wakeBlock is the block at which a parked slot became serviceable again
	// (the reveal's block, or the head at which the decryption window was
	// seen open). Partial waves and the combine rotation count from it, not
	// from the ciphertext's own block, so a slot that waited weeks does not
	// have every member post and combine at once.
	wakeBlock uint64
}

// anchor is the block the slot's stagger schedules count from.
func (ct *ciphertext) anchor() uint64 { return max(ct.block, ct.wakeBlock) }

// parkedSlot is a ciphertext taken out of the per-tick service loop until
// something it waits for happens: the organizer's reveal (wakeAt == 0, woken
// by the OrganizerSecretRevealed event) or the decryption window opening
// (wakeAt = decryptNotBefore, unix seconds, re-checked against the head time
// once per tick without any RPC call).
type parkedSlot struct {
	ct     *ciphertext
	wakeAt uint64
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
		if err := n.discoverCiphertexts(ctx, nil, nil, start, end); err != nil {
			return err
		}
		if err := n.wakeParked(ctx, start, end, head); err != nil {
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

// discoverCiphertexts pulls the CiphertextSubmitted events of [start, end]
// (optionally narrowed to one epoch / application) into the pending set,
// skipping slots the node already tracks or has finished.
func (n *Node) discoverCiphertexts(ctx context.Context, epochs [][12]byte, aids [][32]byte, start, end uint64) error {
	it, err := n.manager.FilterCiphertextSubmitted(&bind.FilterOpts{Context: ctx, Start: start, End: &end}, epochs, aids, nil)
	if err != nil {
		return fmt.Errorf("filter CiphertextSubmitted [%d,%d]: %w", start, end, err)
	}
	defer func() { _ = it.Close() }()
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
			c1:        nodetypes.CurvePoint{X: new(big.Int).Set(ev.C1x), Y: new(big.Int).Set(ev.C1y)},
			c2:        nodetypes.CurvePoint{X: new(big.Int).Set(ev.C2x), Y: new(big.Int).Set(ev.C2y)},
			submitter: ev.Submitter,
			block:     ev.Raw.BlockNumber,
		})
		log.Infow("ciphertext discovered", "ct", key.String(), "block", ev.Raw.BlockNumber)
	}
	if err := it.Error(); err != nil {
		return fmt.Errorf("iterate CiphertextSubmitted: %w", err)
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

// park moves a slot that cannot be served yet out of the per-tick service
// loop: a locked application whose organizer has not revealed (wakeAt == 0)
// or a decryption window that has not opened (wakeAt = decryptNotBefore).
// Epochs stay Live on chain indefinitely, so without this every withheld
// secret would cost every node a few RPC scans per tick forever; a parked
// slot costs nothing until the reveal event or the head time wakes it.
func (n *Node) park(key ctKey, wakeAt uint64) {
	ct, ok := n.pending[key]
	if !ok {
		return
	}
	delete(n.pending, key)
	delete(n.backoff, key)
	if len(n.parked) >= maxParkedCiphertexts {
		var oldest ctKey
		oldestSeq := ^uint64(0)
		for k, p := range n.parked {
			if p.ct.seq < oldestSeq {
				oldest, oldestSeq = k, p.ct.seq
			}
		}
		delete(n.parked, oldest)
	}
	n.parked[key] = &parkedSlot{ct: ct, wakeAt: wakeAt}
	if wakeAt == 0 {
		log.Infow("parked until the organizer reveals its secret", "ct", key.String(), "parked", len(n.parked))
	} else {
		log.Infow("parked until the decryption window opens", "ct", key.String(), "notBefore", wakeAt, "parked", len(n.parked))
	}
}

// wake returns a parked slot to the pending set, anchoring its stagger
// schedules on wakeBlock.
func (n *Node) wake(key ctKey, p *parkedSlot, wakeBlock uint64) {
	delete(n.parked, key)
	p.ct.wakeBlock = wakeBlock
	n.pending[key] = p.ct
}

// wakeParked returns to the pending set every parked slot of an application
// whose organizer revealed its secret in [start, end], and rescans that
// application's ciphertexts from its registration block: the reveal is
// per-application and happens once, and ciphertexts submitted while the
// application was locked may be older than the lookback window of a node
// that restarted meanwhile.
func (n *Node) wakeParked(ctx context.Context, start, end, head uint64) error {
	if len(n.parked) == 0 {
		return nil
	}
	it, err := n.appManager.FilterOrganizerSecretRevealed(&bind.FilterOpts{Context: ctx, Start: start, End: &end}, nil, nil)
	if err != nil {
		return fmt.Errorf("filter OrganizerSecretRevealed [%d,%d]: %w", start, end, err)
	}
	defer func() { _ = it.Close() }()
	for it.Next() {
		ev := it.Event
		woken := 0
		for key, p := range n.parked {
			if key.epoch != ev.EpochId || key.aid != ev.Aid {
				continue
			}
			n.wake(key, p, ev.Raw.BlockNumber)
			woken++
			log.Infow("organizer secret revealed — resuming the slot", "ct", key.String(), "block", ev.Raw.BlockNumber)
		}
		if woken == 0 {
			continue
		}
		if err := n.rescanApplication(ctx, ev.EpochId, ev.Aid, ev.Raw.BlockNumber, head); err != nil {
			// The woken slots are already pending; the rescan only adds
			// older siblings and the next reveal-free tick cannot redo it,
			// so log rather than fail the whole scan.
			log.Warnw("rescan of the revealed application failed", "aid", fmt.Sprintf("%x", ev.Aid[:4]), "err", err)
		}
	}
	if err := it.Error(); err != nil {
		return fmt.Errorf("iterate OrganizerSecretRevealed: %w", err)
	}
	return nil
}

// rescanApplication pulls every ciphertext of (epoch, aid) since the
// application's registration block into the pending set, anchored on
// wakeBlock, so nothing submitted while the slot was parked is missed.
func (n *Node) rescanApplication(ctx context.Context, epochID [12]byte, aid [32]byte, wakeBlock, head uint64) error {
	rec, err := n.appManager.GetApplication(&bind.CallOpts{Context: ctx}, epochID, aid)
	if err != nil {
		return fmt.Errorf("get application: %w", err)
	}
	seqBefore := n.ctSeq
	for start := rec.CreatedAtBlock; start <= head; start += logRangeBlocks {
		end := min(start+logRangeBlocks-1, head)
		if err := n.discoverCiphertexts(ctx, [][12]byte{epochID}, [][32]byte{aid}, start, end); err != nil {
			return err
		}
	}
	for _, ct := range n.pending {
		if ct.seq > seqBefore {
			ct.wakeBlock = wakeBlock
		}
	}
	return nil
}

// wakeDue returns to the pending set every slot whose decryption window has
// opened by the head time. No RPC call is involved: the tick's head is all it
// reads.
func (n *Node) wakeDue(head, headTime uint64) {
	for key, p := range n.parked {
		if p.wakeAt != 0 && headTime >= p.wakeAt {
			n.wake(key, p, head)
			log.Infow("decryption window open — resuming the slot", "ct", key.String(), "notBefore", p.wakeAt)
		}
	}
}

// timeParked reports whether any parked slot waits on the decryption window.
func (n *Node) timeParked() bool {
	for _, p := range n.parked {
		if p.wakeAt != 0 {
			return true
		}
	}
	return false
}

// forget drops every piece of per-slot state.
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
	head     uint64
	headTime uint64 // timestamp of head, for decryption deadlines
	headErr  error
	epochs   map[[12]byte]epochView
	apps     map[appKey]appView
}

// appView is what the decrypt loop needs from an application record: which
// pool key the application encrypts under, its organizer key and — once the
// organizer has revealed it, or trivially for an Automatic application — the
// secret the combine circuit needs.
type appView struct {
	poolIndex        uint8
	pkOrg            nodetypes.CurvePoint
	secret           *big.Int // nil while an organizer-locked secret is still withheld
	openSubmission   bool     // anyone may submit: taint per submitter, not per application
	decryptNotBefore uint64   // unix seconds, 0 = none
	decryptNotAfter  uint64   // unix seconds, 0 = never
}

// appKey identifies one application: (epoch, aid).
type appKey struct {
	epoch [12]byte
	aid   [32]byte
}

func (n *Node) newTickCache(ctx context.Context) *tickCache {
	tc := &tickCache{epochs: make(map[[12]byte]epochView), apps: make(map[appKey]appView)}
	hdr, err := n.contracts.Client().HeaderByNumber(ctx, nil)
	if err != nil {
		tc.headErr = err
		return tc
	}
	tc.head, tc.headTime = hdr.Number.Uint64(), hdr.Time
	return tc
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
	if len(n.pending) == 0 && !n.timeParked() {
		return
	}
	tc := n.newTickCache(ctx)
	if tc.headErr == nil {
		n.wakeDue(tc.head, tc.headTime)
	}
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
	if n.tainted(key, ct.submitter) {
		return true, nil
	}
	idx := myIndex(selected, n.address)
	if idx == 0 {
		// Only committee members hold a share of the pool key, and only they take
		// part in the combine rotation; nothing for this node to do.
		log.Debugw("ciphertext belongs to an epoch this node is not a member of — ignoring", "ct", key.String())
		return true, nil
	}
	app, err := n.application(ctx, tc, key)
	if err != nil {
		return false, err
	}
	if app.decryptNotBefore != 0 || app.decryptNotAfter != 0 {
		if tc.headErr != nil {
			return false, fmt.Errorf("read head: %w", tc.headErr)
		}
		if app.decryptNotAfter != 0 && tc.headTime > app.decryptNotAfter {
			// The contract rejects partials and combines from here on.
			log.Infow("decryption deadline passed — dropping the slot", "ct", key.String(), "deadline", app.decryptNotAfter)
			return true, nil
		}
		if tc.headTime < app.decryptNotBefore {
			// The contract refuses partials and combines until the window
			// opens; wakeDue brings the slot back once the head time passes it.
			n.park(key, app.decryptNotBefore)
			return false, nil
		}
	}
	if app.secret == nil {
		// Organizer-locked and not revealed: the contract refuses every
		// partial and combine (OrganizerSecretNotRevealed), so nothing is
		// posted — not even a partial — until the reveal event wakes the slot.
		n.park(key, 0)
		return false, nil
	}
	if !n.partialDone[key] {
		// Only t partials are ever needed. The first t members of a
		// seed-derived rotation respond at once; each later wave of t steps
		// in staggerBlocks later and only if partials are still missing, so
		// an honest ciphertext costs t partials rather than n and a hostile
		// submitter cannot make the whole committee spend. Waves count from
		// the slot's anchor: the ciphertext block, or the block it was woken
		// at if it sat parked (no partial can exist from before that).
		threshold := uint64(epoch.Policy.Threshold)
		wave := staggerSlot(epoch.Seed, uint64(key.idx), idx, uint16(len(selected))) / max(threshold, 1)
		if wave > 0 {
			if tc.headErr != nil {
				return false, fmt.Errorf("read head: %w", tc.headErr)
			}
			if tc.head < ct.anchor()+wave*staggerBlocks {
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
			} else if !laterWaveDue(tc.head, ct.anchor(), lastBlock, wave) {
				return false, nil
			}
		}
	}
	if !n.partialDone[key] {
		toxic, err := n.submitPartial(ctx, key, ct, idx, epoch, selected, app)
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
	return n.tryCombine(ctx, key, ct, epoch, app, idx, uint16(len(selected)), tc.head)
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

// submitPartial posts δ_i = e_{j,i}·C1 — the share of the application's pool
// key j — with its proof and the Merkle path that pins the share commitment
// to the root activatePoolKey stored. Returns toxic=true when the ciphertext
// is malformed and must never be decrypted.
func (n *Node) submitPartial(
	ctx context.Context,
	key ctKey,
	ct *ciphertext,
	idx uint16,
	epoch epochView,
	selected []common.Address,
	app appView,
) (bool, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	if rec, err := n.manager.GetPartialDecryption(callOpts, key.epoch, key.aid, idx, key.idx); err == nil && rec.Accepted {
		n.partialDone[key] = true
		return false, nil
	}

	// Refuse small-order / off-curve ciphertexts before touching the share:
	// δ_i = e_{j,i}·C1 for a cofactor point would leak the share mod 8
	// on-chain. The contract deliberately skips the prime-subgroup check (it
	// costs ~0.17 M gas per submission), so this is the only place it
	// happens — it is load-bearing, not belt-and-braces.
	if err := group.ValidateCiphertext(ct.c1, ct.c2); err != nil {
		log.Warnw("rejecting toxic ciphertext — refusing partial decryption", "ct", key.String(), "err", err)
		return true, nil
	}

	dShare, err := n.buildPrivateShare(ctx, key.epoch, app.poolIndex, idx, selected, epoch, callOpts)
	if err != nil {
		return false, fmt.Errorf("build private share: %w", err)
	}
	shareProof, err := n.shareProof(ctx, key.epoch, app.poolIndex, idx, epoch, selected)
	if err != nil {
		return false, fmt.Errorf("share inclusion proof: %w", err)
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
	proof, err := n.runtimes.partialDecrypt.ProveAndVerify(witness)
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
		ct.c1.X, ct.c1.Y, ct.c2.X, ct.c2.Y, dHash, proofBytes, inputBytes, shareProof)
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

// application reads (and memoises for the tick) what the decrypt loop needs
// from the application record: its pool key, PK_org, the organizer secret
// once it is knowable and the decryption window. Only registered
// applications can own a ciphertext — the contract rejects submitCiphertext
// for an unknown aid — so a missing record here means the chain is lying or
// we are looking at the wrong contract.
func (n *Node) application(ctx context.Context, tc *tickCache, key ctKey) (appView, error) {
	ak := appKey{epoch: key.epoch, aid: key.aid}
	if app, ok := tc.apps[ak]; ok {
		return app, nil
	}
	rec, err := n.appManager.GetApplication(&bind.CallOpts{Context: ctx}, key.epoch, key.aid)
	if err != nil {
		return appView{}, fmt.Errorf("get application: %w", err)
	}
	if !rec.Exists {
		return appView{}, fmt.Errorf("application %x is not registered", key.aid)
	}
	app := appView{
		poolIndex:        rec.PoolIndex,
		pkOrg:            nodetypes.CurvePoint{X: new(big.Int).Set(rec.OrganizerPK.X), Y: new(big.Int).Set(rec.OrganizerPK.Y)},
		openSubmission:   rec.Policy.OpenSubmission,
		decryptNotBefore: rec.Policy.DecryptNotBefore,
		decryptNotAfter:  rec.Policy.DecryptNotAfter,
	}
	// An automatic application stores the identity key and a zero secret, a
	// locked one a zero secret until revealOrganizerSecret lands. Either way
	// the combine needs sk_org with PK_org = sk_org·G, so one scalar
	// multiplication settles both cases and keeps a lying RPC from making us
	// burn a proof that can never verify.
	if rec.Policy.Mode == uint8(nodetypes.AppModeAutomatic) || rec.OrganizerSecret.Sign() != 0 {
		pk := group.NewPoint()
		pk.ScalarBaseMult(rec.OrganizerSecret)
		if enc := group.Encode(pk); enc.X.Cmp(app.pkOrg.X) == 0 && enc.Y.Cmp(app.pkOrg.Y) == 0 {
			app.secret = new(big.Int).Set(rec.OrganizerSecret)
		} else {
			log.Warnw("organizer secret does not match PK_org — waiting for a correct reveal",
				"aid", fmt.Sprintf("%x", key.aid[:4]))
		}
	}
	tc.apps[ak] = app
	return app, nil
}

// shareProof returns this node's Merkle path into the share-commitment root
// activatePoolKey stored for (epoch, pool key). The leaves are derived from
// the accepted contributions' calldata — the same reconstruction the
// activator proved — and checked against the stored root, so the node never
// depends on the activation transaction's outer calldata (an activation
// relayed through a contract carries a different selector and would strand
// every member). Reading that calldata is only a fast path, taken when it
// decodes and matches the root. The path is fixed for the life of the epoch,
// so it is built once per key and cached.
func (n *Node) shareProof(
	ctx context.Context,
	epochID [12]byte,
	keyIndex uint8,
	myIdx uint16,
	epoch epochView,
	selected []common.Address,
) ([][32]byte, error) {
	slot := poolSlot{epoch: epochID, key: keyIndex}
	if path, ok := n.shareProofs[slot]; ok {
		return path, nil
	}
	root, err := n.manager.GetPoolShareRoot(&bind.CallOpts{Context: ctx}, epochID, keyIndex)
	if err != nil {
		return nil, fmt.Errorf("get pool share root: %w", err)
	}
	if root == ([32]byte{}) {
		return nil, fmt.Errorf("pool key %d is not activated yet", keyIndex)
	}
	leaves, err := n.shareLeavesFromCalldata(ctx, epochID, keyIndex, epoch.SeedBlock, root)
	if err != nil {
		log.Debugw("share commitments: activation calldata unusable, rebuilding from the contributions",
			"epoch", roundHex(epochID), "key", keyIndex, "err", err)
		leaves, err = n.shareLeavesFromContributions(ctx, epochID, keyIndex, epoch, selected, root)
		if err != nil {
			return nil, err
		}
	}
	siblings, err := ccommon.MerklePath(leaves, int(myIdx)-1)
	if err != nil {
		return nil, fmt.Errorf("merkle path for index %d: %w", myIdx, err)
	}
	path := make([][32]byte, len(siblings))
	copy(path, siblings[:])
	n.shareProofs[slot] = path
	return path, nil
}

// shareLeavesFromCalldata is the fast path: decode the share commitments from
// the activatePoolKey transaction found through the PoolKeyActivated event.
// Anything that does not decode as a direct call or does not fold into the
// stored root is an error, and the caller falls back to the contributions.
func (n *Node) shareLeavesFromCalldata(
	ctx context.Context,
	epochID [12]byte,
	keyIndex uint8,
	seedBlock uint64,
	root [32]byte,
) ([ccommon.MaxN][32]byte, error) {
	start := uint64(0)
	if seedBlock > 0 {
		start = seedBlock - 1
	}
	data, err := finalizer.ActivationCalldata(ctx, n.contracts.Client(), n.manager, epochID, keyIndex, start)
	if err != nil {
		return [ccommon.MaxN][32]byte{}, err
	}
	idxs, commitments, err := finalizer.PoolKeyShareCommitments(data)
	if err != nil {
		return [ccommon.MaxN][32]byte{}, fmt.Errorf("decode activation transcript: %w", err)
	}
	return shareLeaves(idxs, commitments, root)
}

// shareLeavesFromContributions rebuilds D_p for every committee member from
// the accepted contributions' calldata, exactly as the activation proof did.
func (n *Node) shareLeavesFromContributions(
	ctx context.Context,
	epochID [12]byte,
	keyIndex uint8,
	epoch epochView,
	selected []common.Address,
	root [32]byte,
) ([ccommon.MaxN][32]byte, error) {
	pi, err := finalizer.PoolKeyStatement(ctx, n.contracts, n.manager, epochID,
		epoch.Policy.Threshold, epoch.Policy.CommitteeSize, selected, keyIndex)
	if err != nil {
		return [ccommon.MaxN][32]byte{}, fmt.Errorf("reconstruct pool key %d: %w", keyIndex, err)
	}
	size := int(epoch.Policy.CommitteeSize)
	if len(pi.ShareCommitments) < size {
		return [ccommon.MaxN][32]byte{}, fmt.Errorf("pool key statement has %d share commitments for a committee of %d",
			len(pi.ShareCommitments), size)
	}
	idxs := make([]uint16, size)
	for i := range idxs {
		idxs[i] = uint16(i + 1)
	}
	return shareLeaves(idxs, pi.ShareCommitments[:size], root)
}

// shareLeaves lays the member-indexed commitments out as tree leaves and
// refuses a set that does not fold into the root the contract stored: a
// path built from it would only buy a reverted partial.
func shareLeaves(idxs []uint16, commitments []nodetypes.CurvePoint, root [32]byte) ([ccommon.MaxN][32]byte, error) {
	leaves, err := ccommon.ShareCommitmentLeaves(idxs, commitments)
	if err != nil {
		return leaves, fmt.Errorf("build share commitment leaves: %w", err)
	}
	if got := ccommon.MerkleRoot(leaves); got != root {
		return leaves, fmt.Errorf("share commitments fold into root %x, contract stored %x", got[:4], root[:4])
	}
	return leaves, nil
}

// tryCombine interpolates threshold partials, recovers the plaintext by BSGS
// and posts the combine proof. A slot becomes combinable once `t` partial
// decryptions are on chain; the organizer secret is already known here
// (serviceCiphertext parks a locked application's slots until the reveal).
// Committee members take turns in a seed-derived rotation (like
// auto-finalize) starting at the block the last partial landed in — or the
// block the slot was woken at, whichever is later, so a slot that waited for
// the window or the reveal does not have every member combine at once;
// normally a single member pays for the combine and later slots only step in
// if the earlier ones did not.
func (n *Node) tryCombine(
	ctx context.Context,
	key ctKey,
	ct *ciphertext,
	epoch epochView,
	app appView,
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
	if app.secret == nil {
		n.park(key, 0)
		return false, nil
	}
	slot := staggerSlot(epoch.Seed, uint64(key.idx), myIdx, committeeSize)
	if waitUntil := max(readyBlock, ct.wakeBlock) + slot*staggerBlocks; head < waitUntil {
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
		go n.runCombineJob(key, ct, idxs, deltas, app, threshold)
		return false, nil
	}
	if res == nil {
		return false, nil // still computing
	}
	n.jobsMu.Lock()
	delete(n.combineJobs, key)
	n.jobsMu.Unlock()
	if res.taint {
		// With a closed submitter set an undecryptable ciphertext is the
		// registrant's doing: serve nothing else of this application in this
		// epoch, which caps what one registration can make the committee
		// compute. With open submission anyone could have sent it, so only
		// that submitter is cut off — one search per (application, address).
		tk := taintKey{epoch: key.epoch, aid: key.aid}
		if app.openSubmission {
			tk.submitter = ct.submitter
		}
		n.taints[tk] = true
		n.saveTaints()
		log.Warnw("tainted: ignoring the remaining ciphertexts for this epoch",
			"ct", key.String(), "openSubmission", app.openSubmission, "submitter", tk.submitter)
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
	app appView,
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
	res := n.combine(key, ct, idxs, deltas, app, threshold)
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
	app appView,
	threshold uint16,
) *combineResult {
	ctx := context.Background()
	// M·G = C2 − Σ λ_k·δ_k − sk_org·C1
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
	c1, err := group.Decode(ct.c1)
	if err != nil {
		return &combineResult{err: err}
	}
	correction := group.NewPoint()
	correction.ScalarMult(c1, app.secret)
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
		OrganizerPK:        app.pkOrg,
		OrganizerSecret:    app.secret,
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
	proof, err := n.runtimes.combine.ProveAndVerify(witness)
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
