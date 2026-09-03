package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	qt "github.com/frankban/quicktest"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
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
// provider-style range cap and records every range it was asked for.
type fakeLogChain struct {
	logs   []ethtypes.Log
	ranges [][2]uint64
}

var errFakeUnsupported = errors.New("fakeLogChain: unsupported")

func (f *fakeLogChain) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]ethtypes.Log, error) {
	from, to := q.FromBlock.Uint64(), q.ToBlock.Uint64()
	f.ranges = append(f.ranges, [2]uint64{from, to})
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

// Partial decryptions are read from the event log between the epoch's seed
// block and head; that span grows with the epoch's age, so it has to be
// chunked to stay under provider limits, and the scan stops once threshold
// distinct partials are in hand.
func TestAcceptedPartialsScansLogsInBoundedChunks(t *testing.T) {
	c := qt.New(t)
	key := ctKey{epoch: [12]byte{1}, aid: [32]byte{2}, idx: 3}
	other := ctKey{epoch: [12]byte{1}, aid: [32]byte{2}, idx: 4}
	const seedBlock, head = uint64(1_000), uint64(36_000)
	chain := &fakeLogChain{logs: []ethtypes.Log{
		partialLog(t, other, 1, 12_000), // different ciphertext slot
		partialLog(t, key, 2, 12_500),
		partialLog(t, key, 2, 12_600), // duplicate participant
		partialLog(t, key, 5, 23_000),
		partialLog(t, key, 7, 34_000), // beyond threshold, never needed
	}}
	m, err := gtypes.NewDKGManager(common.Address{}, chain)
	c.Assert(err, qt.IsNil)
	n := &Node{manager: m}

	idxs, deltas, readyBlock, err := n.acceptedPartials(context.Background(), key, seedBlock, head, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(idxs, qt.DeepEquals, []uint16{2, 5})
	c.Assert(deltas, qt.HasLen, 2)
	c.Assert(deltas[1].X.Int64(), qt.Equals, int64(50))
	c.Assert(readyBlock, qt.Equals, uint64(23_000))

	c.Assert(chain.ranges[0][0], qt.Equals, seedBlock-1)
	for i, r := range chain.ranges {
		c.Assert(r[1]-r[0]+1 <= logRangeBlocks, qt.IsTrue, qt.Commentf("range %d = %v", i, r))
		if i > 0 {
			c.Assert(r[0], qt.Equals, chain.ranges[i-1][1]+1)
		}
	}
	last := chain.ranges[len(chain.ranges)-1]
	c.Assert(last[1] < 34_000, qt.IsTrue, qt.Commentf("scan should stop once threshold partials are found, got %v", chain.ranges))

	chain.ranges = nil
	idxs, _, _, err = n.acceptedPartials(context.Background(), key, seedBlock, head, 3)
	c.Assert(err, qt.IsNil)
	c.Assert(idxs, qt.DeepEquals, []uint16{2, 5, 7})
	c.Assert(chain.ranges[len(chain.ranges)-1][1], qt.Equals, head)
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

func TestParkMovesASlotOutOfPendingAndForgetClearsIt(t *testing.T) {
	n := &Node{
		pending: map[ctKey]*ciphertext{}, parked: map[ctKey]*ciphertext{},
		partialDone: map[ctKey]bool{}, backoff: map[ctKey]*serviceBackoff{}, inflight: map[ctKey]inflightTx{},
		combineJobs: map[ctKey]*combineResult{},
	}
	key := ctKey{idx: 1}
	n.trackCiphertext(key, &ciphertext{block: 10})
	n.backoff[key] = &serviceBackoff{}
	n.park(key)
	if _, ok := n.pending[key]; ok {
		t.Fatal("parked slot still pending")
	}
	if _, ok := n.parked[key]; !ok {
		t.Fatal("slot not parked")
	}
	if _, ok := n.backoff[key]; ok {
		t.Fatal("park kept the backoff record")
	}
	n.park(key) // idempotent: not pending any more
	n.forget(key)
	if _, ok := n.parked[key]; ok {
		t.Fatal("forget left the slot parked")
	}
}

func TestTaintsSurviveARestart(t *testing.T) {
	dir := t.TempDir()
	n := &Node{taintedApps: map[appKey]bool{}, taintFile: taintPath(dir)}
	key := appKey{epoch: [12]byte{1, 2, 3}, aid: [32]byte{9}}
	n.taintedApps[key] = true
	n.saveTaints()
	again := &Node{taintedApps: map[appKey]bool{}, taintFile: taintPath(dir)}
	again.loadTaints()
	if !again.taintedApps[key] {
		t.Fatal("taint not reloaded from disk")
	}
	empty := &Node{taintedApps: map[appKey]bool{}, taintFile: taintPath(t.TempDir())}
	empty.loadTaints() // no file: nothing to load, no error
	if len(empty.taintedApps) != 0 {
		t.Fatal("unexpected taints")
	}
}
