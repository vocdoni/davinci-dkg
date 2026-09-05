package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	qt "github.com/frankban/quicktest"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
)

// A ciphertext event can be reorged out after we saw it; re-scanning the
// last reorgDepthBlocks every tick (the pending map dedupes) keeps the view
// consistent without a full rescan.
func TestScanStartOverlapsTheReorgWindow(t *testing.T) {
	c := qt.New(t)
	// First scan: head − lookback, saturating at genesis.
	c.Assert(scanStart(0, 1_000, 200), qt.Equals, uint64(800))
	c.Assert(scanStart(0, 100, 200), qt.Equals, uint64(0))
	// Subsequent scans re-cover the last reorgDepthBlocks blocks.
	c.Assert(scanStart(5_000, 5_003, 200), qt.Equals, uint64(5_000-reorgDepthBlocks+1))
	c.Assert(scanStart(10, 12, 200), qt.Equals, uint64(0))
}

// Failed service attempts back off exponentially per ciphertext (1, 2, 4 …
// ticks, capped at 64) so one broken slot cannot burn RPC and prover time
// every tick.
func TestServiceBackoffDoublesAndCaps(t *testing.T) {
	c := qt.New(t)
	var b serviceBackoff
	c.Assert(b.due(), qt.IsTrue, qt.Commentf("fresh entries are serviced immediately"))

	waits := []uint32{}
	for range 8 {
		b.fail()
		skipped := uint32(0)
		for !b.due() {
			skipped++
		}
		waits = append(waits, skipped)
	}
	c.Assert(waits, qt.DeepEquals, []uint32{1, 2, 4, 8, 16, 32, 64, 64})
}

// The pending set is bounded; when a spammer floods CiphertextSubmitted the
// oldest entries are dropped (and logged) rather than growing without bound.
func TestTrackCiphertextCapsPendingAndDropsOldest(t *testing.T) {
	c := qt.New(t)
	n := &Node{pending: map[ctKey]*ciphertext{}, partialDone: map[ctKey]bool{}}
	keyAt := func(i int) ctKey { return ctKey{idx: uint16(i % 65536), aid: [32]byte{byte(i >> 16)}} }
	for i := range maxPendingCiphertexts + 3 {
		n.trackCiphertext(keyAt(i), &ciphertext{block: uint64(i)})
	}
	c.Assert(n.pending, qt.HasLen, maxPendingCiphertexts)
	for i := range 3 {
		_, ok := n.pending[keyAt(i)]
		c.Assert(ok, qt.IsFalse, qt.Commentf("oldest entry %d should have been evicted", i))
	}
	_, ok := n.pending[keyAt(3)]
	c.Assert(ok, qt.IsTrue)
	_, ok = n.pending[keyAt(maxPendingCiphertexts+2)]
	c.Assert(ok, qt.IsTrue)
}

// fakeLogChain is a bind.ContractBackend whose FilterLogs enforces the
// provider-style range cap and records every query (and range) it was asked
// for.
type fakeLogChain struct {
	logs    []ethtypes.Log
	ranges  [][2]uint64
	queries []ethereum.FilterQuery
}

var errFakeUnsupported = errors.New("fakeLogChain: unsupported")

func (f *fakeLogChain) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]ethtypes.Log, error) {
	from, to := q.FromBlock.Uint64(), q.ToBlock.Uint64()
	f.ranges = append(f.ranges, [2]uint64{from, to})
	f.queries = append(f.queries, q)
	if to < from || to-from+1 > logRangeBlocks {
		return nil, fmt.Errorf("fakeLogChain: log range [%d,%d] too wide", from, to)
	}
	var out []ethtypes.Log
	for _, l := range f.logs {
		if l.BlockNumber >= from && l.BlockNumber <= to && topicsMatch(q.Topics, l.Topics) {
			out = append(out, l)
		}
	}
	return out, nil
}

func topicsMatch(want [][]common.Hash, got []common.Hash) bool {
	for i, set := range want {
		if len(set) == 0 {
			continue
		}
		if i >= len(got) || !slices.Contains(set, got[i]) {
			return false
		}
	}
	return true
}

func (*fakeLogChain) CodeAt(context.Context, common.Address, *big.Int) ([]byte, error) {
	return nil, errFakeUnsupported
}

func (*fakeLogChain) CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error) {
	return nil, errFakeUnsupported
}

func (*fakeLogChain) HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error) {
	return nil, errFakeUnsupported
}

func (*fakeLogChain) PendingCodeAt(context.Context, common.Address) ([]byte, error) {
	return nil, errFakeUnsupported
}

func (*fakeLogChain) PendingNonceAt(context.Context, common.Address) (uint64, error) {
	return 0, errFakeUnsupported
}

func (*fakeLogChain) SuggestGasPrice(context.Context) (*big.Int, error) {
	return nil, errFakeUnsupported
}

func (*fakeLogChain) SuggestGasTipCap(context.Context) (*big.Int, error) {
	return nil, errFakeUnsupported
}

func (*fakeLogChain) EstimateGas(context.Context, ethereum.CallMsg) (uint64, error) {
	return 0, errFakeUnsupported
}

func (*fakeLogChain) SendTransaction(context.Context, *ethtypes.Transaction) error {
	return errFakeUnsupported
}

func (*fakeLogChain) SubscribeFilterLogs(context.Context, ethereum.FilterQuery, chan<- ethtypes.Log) (ethereum.Subscription, error) {
	return nil, errFakeUnsupported
}

func partialLog(t *testing.T, key ctKey, participant uint16, block uint64) ethtypes.Log {
	t.Helper()
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := parsed.Events["PartialDecryptionSubmitted"]
	addr := common.BigToAddress(big.NewInt(int64(participant)))
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{key.epoch}, []any{key.aid}, []any{addr})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack(participant, key.idx, big.NewInt(int64(participant)*10), big.NewInt(int64(participant)*10+1))
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Topics:      []common.Hash{topics[0][0], topics[1][0], topics[2][0], topics[3][0]},
		Data:        data,
		BlockNumber: block,
	}
}

// bindTestChain wires a test node to a fake backend: the abigen bindings the
// event parsers and the getApplication read go through, and the log filterer
// the scan uses.
func bindTestChain(t *testing.T, n *Node, chain interface {
	logFilterer
	bind.ContractBackend
},
) {
	t.Helper()
	m, err := gtypes.NewDKGManager(common.Address{}, chain)
	qt.Assert(t, err, qt.IsNil)
	am, err := gtypes.NewDKGAppManager(common.Address{}, chain)
	qt.Assert(t, err, qt.IsNil)
	n.manager, n.appManager, n.logs = m, am, chain
}

// Partial decryptions are no longer read with a log query per slot per tick:
// the scan records every PartialDecryptionSubmitted of a tracked slot as it
// passes, in chain order, and acceptedPartials answers from that record —
// the first t distinct participants, exactly as the event log would give
// them, with the block the last of them landed in. Partials of slots the
// node does not track are not kept, a re-delivered event is not a second
// partial, and the re-scanned reorg window is rebuilt from fresh logs.
func TestAcceptedPartialsComeFromTheScanInChainOrder(t *testing.T) {
	c := qt.New(t)
	key := ctKey{epoch: [12]byte{1}, aid: [32]byte{2}, idx: 3}
	other := ctKey{epoch: [12]byte{1}, aid: [32]byte{2}, idx: 4}
	const head = uint64(36_000)
	chain := &fakeLogChain{logs: []ethtypes.Log{
		partialLog(t, other, 1, 12_000), // an untracked slot
		partialLog(t, key, 5, 23_000),
		partialLog(t, key, 2, 12_500), // out of order in the response: sorted by block
		partialLog(t, key, 2, 12_600), // duplicate participant
		partialLog(t, key, 7, 35_990), // inside the reorg window of the next scan
	}}
	n := newTestNode()
	bindTestChain(t, n, chain)
	n.lookback = head // the first scan covers the whole history here
	n.trackCiphertext(key, &ciphertext{block: 12_000})

	c.Assert(n.scanCiphertexts(context.Background(), head), qt.IsNil)
	idxs, deltas, readyBlock := n.acceptedPartials(key, 2)
	c.Assert(idxs, qt.DeepEquals, []uint16{2, 5})
	c.Assert(deltas, qt.HasLen, 2)
	c.Assert(deltas[1].X.Int64(), qt.Equals, int64(50))
	c.Assert(readyBlock, qt.Equals, uint64(23_000))
	idxs, _, readyBlock = n.acceptedPartials(key, 3)
	c.Assert(idxs, qt.DeepEquals, []uint16{2, 5, 7})
	c.Assert(readyBlock, qt.Equals, uint64(35_990))
	_, untracked := n.partials[other]
	c.Assert(untracked, qt.IsFalse, qt.Commentf("partials of untracked slots must not be kept"))

	// The next tick re-scans the reorg window: a partial whose block was
	// reorged out disappears, the rest (older than the window, or still in
	// the logs) survive once and only once.
	chain.logs = chain.logs[:4] // participant 7's block is gone
	c.Assert(n.scanCiphertexts(context.Background(), head+1), qt.IsNil)
	c.Assert(chain.ranges[len(chain.ranges)-1][0], qt.Equals, scanStart(head, head+1, n.lookback))
	idxs, _, _ = n.acceptedPartials(key, 3)
	c.Assert(idxs, qt.DeepEquals, []uint16{2, 5}, qt.Commentf("the window rebuild must drop the reorged partial"))
	c.Assert(n.partials[key], qt.HasLen, 2)

	// forget drops the record with the slot.
	n.forget(key)
	_, kept := n.partials[key]
	c.Assert(kept, qt.IsFalse)
}

func TestLaterWaveDueNeedsBothDelayAndStalledProgress(t *testing.T) {
	const ct = 100
	// Wave 1 may not fire before its own delay has passed…
	if laterWaveDue(ct+staggerBlocks-1, ct, 0, 1) {
		t.Fatal("wave 1 fired before ctBlock+staggerBlocks")
	}
	// …nor while wave 0 keeps landing partials.
	if laterWaveDue(ct+10, ct, ct+9, 1) {
		t.Fatal("wave 1 fired while earlier waves were still landing partials")
	}
	// It fires once both the delay and staggerBlocks of silence have passed.
	if !laterWaveDue(ct+10, ct, ct+10-staggerBlocks, 1) {
		t.Fatal("wave 1 did not fire after the delay and a stalled wave 0")
	}
	// Wave 2 waits twice as long from the ciphertext.
	if laterWaveDue(ct+2*staggerBlocks-1, ct, 0, 2) {
		t.Fatal("wave 2 fired before ctBlock+2*staggerBlocks")
	}
}

func newTestNode() *Node {
	return &Node{
		pending: map[ctKey]*ciphertext{}, parked: map[ctKey]*parkedSlot{}, served: map[ctKey]uint64{},
		partialDone: map[ctKey]bool{}, backoff: map[ctKey]*serviceBackoff{}, inflight: map[ctKey]inflightTx{},
		combineJobs: map[ctKey]*combineResult{}, taints: map[taintKey]bool{},
		epochCache: map[[12]byte]epochView{}, apps: map[appKey]appView{}, partials: map[ctKey][]partialRecord{},
	}
}

func TestParkMovesASlotOutOfPendingAndForgetClearsIt(t *testing.T) {
	n := newTestNode()
	key := ctKey{idx: 1}
	n.trackCiphertext(key, &ciphertext{block: 10})
	n.backoff[key] = &serviceBackoff{}
	n.park(key, 0)
	if _, ok := n.pending[key]; ok {
		t.Fatal("parked slot still pending")
	}
	if _, ok := n.parked[key]; !ok {
		t.Fatal("slot not parked")
	}
	if _, ok := n.backoff[key]; ok {
		t.Fatal("park kept the backoff record")
	}
	n.park(key, 0) // idempotent: not pending any more
	n.forget(key)
	if _, ok := n.parked[key]; ok {
		t.Fatal("forget left the slot parked")
	}
}

// A slot whose decryption window has not opened is parked with the window's
// start; every tick compares the parked slots against the head time and
// wakes the due ones, anchoring their stagger schedules on the wake block so
// the committee does not pile on at once.
func TestWakeDueResumesWindowedSlotsAndAnchorsThem(t *testing.T) {
	c := qt.New(t)
	n := newTestNode()
	windowed, locked := ctKey{idx: 1}, ctKey{idx: 2}
	n.trackCiphertext(windowed, &ciphertext{block: 10})
	n.trackCiphertext(locked, &ciphertext{block: 11})
	n.park(windowed, 1_700_000_000)
	n.park(locked, 0)
	c.Assert(n.timeParked(), qt.IsTrue)

	n.wakeDue(500, 1_699_999_999)
	c.Assert(n.pending, qt.HasLen, 0, qt.Commentf("woke before the window opened"))

	n.wakeDue(520, 1_700_000_000)
	ct, ok := n.pending[windowed]
	c.Assert(ok, qt.IsTrue)
	c.Assert(ct.wakeBlock, qt.Equals, uint64(520))
	c.Assert(ct.anchor(), qt.Equals, uint64(520))
	_, stillParked := n.parked[locked]
	c.Assert(stillParked, qt.IsTrue, qt.Commentf("a reveal-parked slot is not woken by time"))
	c.Assert(n.timeParked(), qt.IsFalse)
}

func TestTaintsSurviveARestart(t *testing.T) {
	dir := t.TempDir()
	n := &Node{taints: map[taintKey]bool{}, taintFile: taintPath(dir)}
	app := taintKey{epoch: [12]byte{1, 2, 3}, aid: [32]byte{9}}
	sub := taintKey{epoch: [12]byte{1, 2, 3}, aid: [32]byte{8}, submitter: common.HexToAddress("0xabc")}
	n.taints[app] = true
	n.taints[sub] = true
	n.saveTaints()
	again := &Node{taints: map[taintKey]bool{}, taintFile: taintPath(dir)}
	again.loadTaints()
	if !again.taints[app] || !again.taints[sub] {
		t.Fatal("taints not reloaded from disk")
	}
	empty := &Node{taints: map[taintKey]bool{}, taintFile: taintPath(t.TempDir())}
	empty.loadTaints() // no file: nothing to load, no error
	if len(empty.taints) != 0 {
		t.Fatal("unexpected taints")
	}
}

// Files written before per-submitter taints existed hold epoch:aid entries;
// they load as whole-application taints, and the lookup honours both kinds.
func TestTaintFileLoadsLegacyEntriesAsWholeApplicationTaints(t *testing.T) {
	c := qt.New(t)
	dir := t.TempDir()
	path := taintPath(dir)
	legacy := "010203000000000000000000:0900000000000000000000000000000000000000000000000000000000000000"
	c.Assert(os.WriteFile(path, []byte(`["`+legacy+`","garbage"]`), 0o600), qt.IsNil)
	n := newTestNode()
	n.taintFile = path
	n.loadTaints()
	c.Assert(n.taints, qt.HasLen, 1)

	key := ctKey{epoch: [12]byte{1, 2, 3}, aid: [32]byte{9}}
	c.Assert(n.tainted(key, common.HexToAddress("0x1")), qt.IsTrue, qt.Commentf("whole-app taint covers every submitter"))
	other := ctKey{epoch: [12]byte{1, 2, 3}, aid: [32]byte{8}}
	c.Assert(n.tainted(other, common.HexToAddress("0x1")), qt.IsFalse)
	n.taints[taintKey{epoch: other.epoch, aid: other.aid, submitter: common.HexToAddress("0x1")}] = true
	c.Assert(n.tainted(other, common.HexToAddress("0x1")), qt.IsTrue)
	c.Assert(n.tainted(other, common.HexToAddress("0x2")), qt.IsFalse, qt.Commentf("per-submitter taint spares the others"))
}

// fakeAppChain answers DKGAppManager getApplication calls on top of
// fakeLogChain's log serving.
type fakeAppChain struct {
	fakeLogChain
	app       gtypes.DKGTypesApplication
	failCalls bool
}

func (f *fakeAppChain) CallContract(_ context.Context, msg ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	if f.failCalls {
		return nil, errFakeUnsupported
	}
	abiJSON, err := gtypes.DKGAppManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	m := abiJSON.Methods["getApplication"]
	if len(msg.Data) < 4 || string(msg.Data[:4]) != string(m.ID) {
		return nil, errFakeUnsupported
	}
	return m.Outputs.Pack(f.app)
}

func revealLog(t *testing.T, epoch [12]byte, aid [32]byte, block uint64) ethtypes.Log {
	t.Helper()
	abiJSON, err := gtypes.DKGAppManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := abiJSON.Events["OrganizerSecretRevealed"]
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{epoch}, []any{aid})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack(big.NewInt(12345))
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Address:     common.Address{0xaa},
		Topics:      []common.Hash{topics[0][0], topics[1][0], topics[2][0]},
		Data:        data,
		BlockNumber: block,
	}
}

func ciphertextLog(t *testing.T, key ctKey, block uint64) ethtypes.Log {
	t.Helper()
	abiJSON, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := abiJSON.Events["CiphertextSubmitted"]
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{key.epoch}, []any{key.aid}, []any{key.idx})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack(common.HexToAddress("0xfeed"), big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(4))
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Address:     common.Address{0xbb},
		Topics:      []common.Hash{topics[0][0], topics[1][0], topics[2][0], topics[3][0]},
		Data:        data,
		BlockNumber: block,
	}
}

func combinedLog(t *testing.T, key ctKey, block uint64) ethtypes.Log {
	t.Helper()
	abiJSON, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := abiJSON.Events["DecryptionCombined"]
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{key.epoch}, []any{key.aid}, []any{key.idx})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack([32]byte{0xc0}, big.NewInt(42))
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Address:     common.Address{0xbb},
		Topics:      []common.Hash{topics[0][0], topics[1][0], topics[2][0], topics[3][0]},
		Data:        data,
		BlockNumber: block,
		TxHash:      common.HexToHash("0xc0ffee"),
	}
}

// A node that restarted after a locked application's ciphertexts were parked
// holds none of them: the reveal event alone — no parked slot — must still
// rescan the application from its registration block, and a failed rescan
// must surface so the scan cursor does not advance past the event.
func TestRevealRescansApplicationsWithoutParkedSlots(t *testing.T) {
	c := qt.New(t)
	key := ctKey{epoch: [12]byte{7}, aid: [32]byte{9}, idx: 1}
	chain := &fakeAppChain{app: gtypes.DKGTypesApplication{
		Creator:         common.HexToAddress("0x1"),
		OrganizerPK:     gtypes.DKGTypesPoint{X: big.NewInt(1), Y: big.NewInt(2)},
		OrganizerSecret: big.NewInt(0),
		Policy:          gtypes.DKGTypesAppPolicy{Mode: 1, Submitters: []common.Address{}},
		CreatedAtBlock:  500,
		Exists:          true,
	}}
	chain.logs = []ethtypes.Log{
		revealLog(t, key.epoch, key.aid, 1_500),
		ciphertextLog(t, key, 800),              // submitted while locked, before this node's scan window
		revealLog(t, key.epoch, key.aid, 3_500), // a second application's reveal, for the failure case
	}
	n := newTestNode()
	bindTestChain(t, n, chain)
	c.Assert(n.parked, qt.HasLen, 0)

	c.Assert(n.scanRange(context.Background(), 1_000, 2_000, 2_000), qt.IsNil)

	ct, ok := n.pending[key]
	c.Assert(ok, qt.IsTrue, qt.Commentf("the reveal must rescan the application even with nothing parked"))
	c.Assert(ct.wakeBlock, qt.Equals, uint64(1_500), qt.Commentf("rescanned slots anchor on the reveal block"))
	c.Assert(ct.submitter, qt.Equals, common.HexToAddress("0xfeed"))
	// The rescan's own filter starts at the registration block, is narrowed
	// to the application and asks for its ciphertexts, partials and combines
	// at once.
	rescan := chain.queries[len(chain.queries)-1]
	c.Assert(rescan.Topics, qt.HasLen, 3)
	c.Assert(rescan.Topics[0], qt.HasLen, 3)
	c.Assert(rescan.FromBlock.Uint64(), qt.Equals, uint64(500))

	chain.failCalls = true
	err := n.scanRange(context.Background(), 3_000, 4_000, 4_000)
	c.Assert(err, qt.Not(qt.IsNil), qt.Commentf("a failed rescan must fail the range so the cursor does not advance"))
}

// One eth_getLogs per range covers every event kind the node acts on:
// CiphertextSubmitted, OrganizerSecretRevealed, PartialDecryptionSubmitted
// and DecryptionCombined, over the manager and the app manager together,
// dispatched by topic 0 to the abigen parsers. Each kind must land where its
// dedicated filter used to put it: a new slot in pending, a reveal waking
// the parked slots of its application (and refreshing its cached view), a
// partial in the slot's record (this node's own marking the partial done),
// a combine retiring the slot.
func TestScanRangeDispatchesEveryEventKindFromOneFilter(t *testing.T) {
	c := qt.New(t)
	epoch := [12]byte{7}
	fresh := ctKey{epoch: epoch, aid: [32]byte{1}, idx: 1}
	locked := ctKey{epoch: epoch, aid: [32]byte{2}, idx: 1}
	combined := ctKey{epoch: epoch, aid: [32]byte{3}, idx: 1}
	me := common.HexToAddress("0x0000000000000000000000000000000000000004")
	chain := &fakeAppChain{app: gtypes.DKGTypesApplication{ // organizer-locked, secret still withheld
		Creator:         common.HexToAddress("0x1"),
		OrganizerPK:     gtypes.DKGTypesPoint{X: big.NewInt(1), Y: big.NewInt(2)},
		OrganizerSecret: big.NewInt(0),
		Policy:          gtypes.DKGTypesAppPolicy{Mode: uint8(nodetypes.AppModeOrganizerLocked), Submitters: []common.Address{}},
		CreatedAtBlock:  9_000,
		Exists:          true,
	}}
	chain.logs = []ethtypes.Log{
		ciphertextLog(t, fresh, 10_010),
		partialLog(t, fresh, 4, 10_020), // participant 4 is this node
		partialLog(t, fresh, 6, 10_030),
		revealLog(t, locked.epoch, locked.aid, 10_040),
		combinedLog(t, combined, 10_050),
	}
	n := newTestNode()
	n.address = me
	n.lookback = 1_000
	bindTestChain(t, n, chain)
	n.trackCiphertext(locked, &ciphertext{block: 9_500})
	n.park(locked, 0)
	n.apps[appKey{epoch: locked.epoch, aid: locked.aid}] = appView{poolIndex: 1}
	n.trackCiphertext(combined, &ciphertext{block: 9_600})
	n.inflight[combined] = inflightTx{combine: true}

	c.Assert(n.scanCiphertexts(context.Background(), 10_100), qt.IsNil)

	// One filter for the range, over both contracts, all four signatures.
	ids, err := eventIDs()
	c.Assert(err, qt.IsNil)
	first := chain.queries[0]
	c.Assert(first.Addresses, qt.HasLen, 2)
	c.Assert(first.Topics, qt.HasLen, 1)
	c.Assert(first.Topics[0], qt.DeepEquals, []common.Hash{ids.ciphertext, ids.reveal, ids.partial, ids.combined})
	c.Assert(first.FromBlock.Uint64(), qt.Equals, uint64(10_100-1_000))
	c.Assert(first.ToBlock.Uint64(), qt.Equals, uint64(10_100))

	// CiphertextSubmitted → pending, with its partials recorded behind it.
	ct, ok := n.pending[fresh]
	c.Assert(ok, qt.IsTrue)
	c.Assert(ct.block, qt.Equals, uint64(10_010))
	idxs, _, ready := n.acceptedPartials(fresh, 2)
	c.Assert(idxs, qt.DeepEquals, []uint16{4, 6})
	c.Assert(ready, qt.Equals, uint64(10_030))
	c.Assert(n.partialDone[fresh], qt.IsTrue, qt.Commentf("this node's own partial marks the slot's partial done"))

	// OrganizerSecretRevealed → the parked slot is back in pending, anchored
	// on the reveal, and the application view was re-read.
	woken, ok := n.pending[locked]
	c.Assert(ok, qt.IsTrue, qt.Commentf("the reveal must wake the parked slot"))
	c.Assert(woken.wakeBlock, qt.Equals, uint64(10_040))
	c.Assert(n.apps[appKey{epoch: locked.epoch, aid: locked.aid}].createdAt, qt.Equals, uint64(9_000))

	// DecryptionCombined → the slot is retired, in-flight record and all,
	// and remembered as served.
	_, stillPending := n.pending[combined]
	c.Assert(stillPending, qt.IsFalse, qt.Commentf("a combined slot must be retired"))
	_, stillInflight := n.inflight[combined]
	c.Assert(stillInflight, qt.IsFalse)
	c.Assert(n.served[combined], qt.Equals, uint64(10_050))

	// The scan cursor advanced; the next tick re-covers only the window.
	c.Assert(n.lastCtScan, qt.Equals, uint64(10_100))
	c.Assert(scanStart(n.lastCtScan, 10_101, n.lookback), qt.Equals, uint64(10_101-reorgDepthBlocks))
}
