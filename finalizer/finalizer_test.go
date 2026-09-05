package finalizer

import (
	"bytes"
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
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
)

// testLayout is the compact transcript shape the fixtures below use: a
// committee of five with threshold three.
func testLayout(t *testing.T) contribution.Layout {
	t.Helper()
	layout, err := contribution.NewLayout(3, 5)
	qt.Assert(t, err, qt.IsNil)
	return layout
}

// validSubmitContributionCalldata packs a well-formed submitContribution call
// whose transcript is the compact L_C(t, n)-word layout the contract expects
// for the test layout, with every recipient slot carrying its own index.
func validSubmitContributionCalldata(t *testing.T) (calldata, transcript []byte) {
	t.Helper()
	layout := testLayout(t)
	words := make([]*big.Int, 0, layout.Words())
	for i := 0; i < layout.Words(); i++ {
		words = append(words, big.NewInt(int64(1000+i)))
	}
	for i := range layout.CommitteeSize {
		words[layout.RecipientIndexOffset(i)] = big.NewInt(int64(i + 1))
	}
	transcript, err := encodeWords(words...)
	qt.Assert(t, err, qt.IsNil)
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	calldata, err = parsed.Pack("submitContribution",
		[12]byte{1}, uint16(1), [32]byte{2}, [32]byte{3}, transcript, []byte{0xaa}, []byte{0xbb})
	qt.Assert(t, err, qt.IsNil)
	return calldata, transcript
}

func TestContributionTranscriptRoundTripsRealCalldata(t *testing.T) {
	c := qt.New(t)
	layout := testLayout(t)
	calldata, transcript := validSubmitContributionCalldata(t)
	got, err := ContributionTranscript(calldata, layout)
	c.Assert(err, qt.IsNil)
	c.Assert(bytes.Equal(got, transcript), qt.IsTrue)

	// The same calldata read under another epoch policy has the wrong length
	// and is refused: the layout comes from the epoch, never from calldata.
	other, err := contribution.NewLayout(2, 5)
	c.Assert(err, qt.IsNil)
	_, err = ContributionTranscript(calldata, other)
	c.Assert(err, qt.ErrorMatches, "transcript is .* bytes, want .*")
}

// The decoder addresses the compact regions by the layout's offsets, so the
// node reads the ephemeral and masked share of its own slot and the finalizer
// the commitment coordinates of every key.
func TestDecodeContributionFollowsTheCompactLayout(t *testing.T) {
	c := qt.New(t)
	layout := testLayout(t)
	calldata, _ := validSubmitContributionCalldata(t)
	tr, err := DecodeContribution(calldata, layout)
	c.Assert(err, qt.IsNil)
	c.Assert(tr.Commitments, qt.HasLen, contribution.MaxKeys)
	c.Assert(tr.Commitments[0], qt.HasLen, layout.Threshold)
	c.Assert(tr.Commitments[1][2].X.Int64(), qt.Equals, int64(1000+layout.CommitmentOffset(1, 2)))
	c.Assert(tr.RecipientIndexes[3].Int64(), qt.Equals, int64(4))
	c.Assert(tr.Ephemerals[4].Y.Int64(), qt.Equals, int64(1000+layout.EphemeralOffset(4)+1))
	c.Assert(tr.MaskedShares[contribution.MaxKeys-1][0].Int64(),
		qt.Equals, int64(1000+layout.MaskedShareOffset(contribution.MaxKeys-1, 0)))
}

// Calldata comes from arbitrary transactions located through an event log; a
// hostile or merely unexpected payload must yield an error, never a panic.
func TestContributionTranscriptRejectsHostileCalldata(t *testing.T) {
	layout := testLayout(t)
	valid, _ := validSubmitContributionCalldata(t)
	const transcriptOffsetWord = 4 + 4*32 // selector + fifth head word

	withOffset := func(word []byte) []byte {
		out := bytes.Clone(valid)
		copy(out[transcriptOffsetWord:transcriptOffsetWord+32], word)
		return out
	}
	allOnes := bytes.Repeat([]byte{0xff}, 32)
	int64Min := make([]byte, 32)
	int64Min[24] = 0x80 // 2^63: negative once cast to int64
	uint64Max := make([]byte, 32)
	copy(uint64Max[24:], bytes.Repeat([]byte{0xff}, 8))
	nearEnd := make([]byte, 32)
	// An offset that is in range for the length word but whose 32-byte length
	// read overflows the payload.
	nearEnd[30] = byte((len(valid) - 4 - 16) >> 8)
	nearEnd[31] = byte(len(valid) - 4 - 16)

	wrongSelector := bytes.Clone(valid)
	wrongSelector[0] ^= 0xff

	cases := map[string][]byte{
		"empty":                 {},
		"selector only":         valid[:4],
		"short head":            valid[:4+6*32],
		"offset 2^256-1":        withOffset(allOnes),
		"offset 2^63":           withOffset(int64Min),
		"offset 2^64-1":         withOffset(uint64Max),
		"offset near end":       withOffset(nearEnd),
		"truncated transcript":  valid[:4+8*32+100],
		"wrong selector":        wrongSelector,
		"random garbage":        bytes.Repeat([]byte{0x5a}, 4+7*32+64),
		"length word all ones":  func() []byte { out := bytes.Clone(valid); copy(out[4+7*32:4+8*32], allOnes); return out }(),
		"transcript wrong size": func() []byte { out := bytes.Clone(valid); out[4+7*32+31] ^= 0x01; return out }(),
	}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("panicked on hostile calldata: %v", r)
				}
			}()
			tr, err := ContributionTranscript(data, layout)
			qt.Assert(t, err, qt.Not(qt.IsNil))
			qt.Assert(t, tr, qt.IsNil)
		})
	}
}

// fakeChain is a minimal bind.ContractBackend plus the two ethclient calls
// ContributionCalldata needs. It refuses eth_getLogs ranges wider than
// LogRangeBlocks, like most public RPC providers do.
type fakeChain struct {
	head   uint64
	logs   []ethtypes.Log
	txs    map[common.Hash]*ethtypes.Transaction
	ranges [][2]uint64
	// eth_call answers for the pre-send recheck: the epoch and the
	// contribution records (a missing address is a member that never
	// contributed).
	epoch   *gtypes.IDKGManagerEpoch
	records map[common.Address]gtypes.DKGTypesContributionRecord
}

var errFakeUnsupported = errors.New("fakeChain: unsupported")

func (f *fakeChain) BlockNumber(context.Context) (uint64, error) { return f.head, nil }

func (f *fakeChain) TransactionByHash(_ context.Context, h common.Hash) (*ethtypes.Transaction, bool, error) {
	tx, ok := f.txs[h]
	if !ok {
		return nil, false, ethereum.NotFound
	}
	return tx, false, nil
}

func (f *fakeChain) FilterLogs(_ context.Context, q ethereum.FilterQuery) ([]ethtypes.Log, error) {
	from, to := q.FromBlock.Uint64(), q.ToBlock.Uint64()
	f.ranges = append(f.ranges, [2]uint64{from, to})
	if to < from || to-from+1 > LogRangeBlocks {
		return nil, fmt.Errorf("fakeChain: log range [%d,%d] too wide", from, to)
	}
	var out []ethtypes.Log
	for _, l := range f.logs {
		if l.BlockNumber >= from && l.BlockNumber <= to && topicsMatch(q.Topics, l.Topics) {
			out = append(out, l)
		}
	}
	return out, nil
}

// topicsMatch applies eth_getLogs topic semantics: each position is a set of
// acceptable values, an empty set matches anything.
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

func (*fakeChain) CodeAt(context.Context, common.Address, *big.Int) ([]byte, error) {
	return nil, errFakeUnsupported
}

func (f *fakeChain) CallContract(_ context.Context, call ethereum.CallMsg, _ *big.Int) ([]byte, error) {
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	if len(call.Data) < 4 {
		return nil, errFakeUnsupported
	}
	method, err := parsed.MethodById(call.Data[:4])
	if err != nil {
		return nil, err
	}
	switch method.Name {
	case "getEpoch":
		if f.epoch == nil {
			return nil, errFakeUnsupported
		}
		return method.Outputs.Pack(*f.epoch)
	case "getContribution":
		args, err := method.Inputs.Unpack(call.Data[4:])
		if err != nil {
			return nil, err
		}
		contributor, _ := args[1].(common.Address)
		return method.Outputs.Pack(f.records[contributor])
	default:
		return nil, errFakeUnsupported
	}
}

func (*fakeChain) HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error) {
	return nil, errFakeUnsupported
}

func (*fakeChain) PendingCodeAt(context.Context, common.Address) ([]byte, error) {
	return nil, errFakeUnsupported
}

func (*fakeChain) PendingNonceAt(context.Context, common.Address) (uint64, error) {
	return 0, errFakeUnsupported
}

func (*fakeChain) SuggestGasPrice(context.Context) (*big.Int, error) { return nil, errFakeUnsupported }

func (*fakeChain) SuggestGasTipCap(context.Context) (*big.Int, error) { return nil, errFakeUnsupported }

func (*fakeChain) EstimateGas(context.Context, ethereum.CallMsg) (uint64, error) {
	return 0, errFakeUnsupported
}

func (*fakeChain) SendTransaction(context.Context, *ethtypes.Transaction) error {
	return errFakeUnsupported
}

func (*fakeChain) SubscribeFilterLogs(context.Context, ethereum.FilterQuery, chan<- ethtypes.Log) (ethereum.Subscription, error) {
	return nil, errFakeUnsupported
}

// contributionLog builds a ContributionSubmitted log exactly as the contract
// would emit it, so the abigen iterator can unpack it.
func contributionLog(t *testing.T, epochID [12]byte, contributor common.Address, block uint64, tx *ethtypes.Transaction) ethtypes.Log {
	t.Helper()
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := parsed.Events["ContributionSubmitted"]
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{epochID}, []any{contributor})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack(uint16(1), [32]byte{}, [32]byte{})
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Topics:      []common.Hash{topics[0][0], topics[1][0], topics[2][0]},
		Data:        data,
		BlockNumber: block,
		TxHash:      tx.Hash(),
	}
}

// epochLiveLog builds an EpochLive log as the contract emits it at the end of
// finalizeEpoch.
func epochLiveLog(t *testing.T, epochID [12]byte, block uint64, tx *ethtypes.Transaction) ethtypes.Log {
	t.Helper()
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	ev := parsed.Events["EpochLive"]
	topics, err := abi.MakeTopics([]any{ev.ID}, []any{epochID})
	qt.Assert(t, err, qt.IsNil)
	data, err := ev.Inputs.NonIndexed().Pack(uint16(3))
	qt.Assert(t, err, qt.IsNil)
	return ethtypes.Log{
		Topics:      []common.Hash{topics[0][0], topics[1][0]},
		Data:        data,
		BlockNumber: block,
		TxHash:      tx.Hash(),
	}
}

// Public RPCs cap eth_getLogs at ~10k blocks, so the scan from the epoch's
// seed block to head must be chunked instead of issued as one query.
func TestContributionCalldataScansLogsInBoundedChunks(t *testing.T) {
	c := qt.New(t)
	calldata, _ := validSubmitContributionCalldata(t)
	tx := ethtypes.NewTx(&ethtypes.LegacyTx{Data: calldata})
	epochID := [12]byte{9}
	contributor := common.HexToAddress("0x1234")
	const seedBlock, eventBlock, head = uint64(100), uint64(31_500), uint64(45_000)

	chain := &fakeChain{
		head: head,
		logs: []ethtypes.Log{contributionLog(t, epochID, contributor, eventBlock, tx)},
		txs:  map[common.Hash]*ethtypes.Transaction{tx.Hash(): tx},
	}
	m, err := gtypes.NewDKGManager(common.Address{}, chain)
	c.Assert(err, qt.IsNil)

	got, err := ContributionCalldata(context.Background(), chain, m, epochID, contributor, seedBlock-1)
	c.Assert(err, qt.IsNil)
	c.Assert(bytes.Equal(got, calldata), qt.IsTrue)

	c.Assert(len(chain.ranges) > 1, qt.IsTrue, qt.Commentf("ranges: %v", chain.ranges))
	c.Assert(chain.ranges[0][0], qt.Equals, seedBlock-1)
	for i, r := range chain.ranges {
		c.Assert(r[1]-r[0]+1 <= LogRangeBlocks, qt.IsTrue, qt.Commentf("range %d = %v", i, r))
		if i > 0 {
			c.Assert(r[0], qt.Equals, chain.ranges[i-1][1]+1)
		}
	}
	last := chain.ranges[len(chain.ranges)-1]
	c.Assert(last[0] <= eventBlock && eventBlock <= last[1], qt.IsTrue, qt.Commentf("stopped at %v", last))

	chain.ranges = nil
	_, err = ContributionCalldata(context.Background(), chain, m, epochID, common.HexToAddress("0x5678"), seedBlock-1)
	c.Assert(err, qt.ErrorMatches, "no ContributionSubmitted event.*")
	c.Assert(chain.ranges[len(chain.ranges)-1][1], qt.Equals, head)
}

// The finalization transaction is located through the EpochLive event the
// same way, so a member can read its Merkle leaves back from calldata.
func TestFinalizeCalldataIsLocatedThroughEpochLive(t *testing.T) {
	c := qt.New(t)
	calldata := validFinalizeEpochCalldata(t, 5)
	tx := ethtypes.NewTx(&ethtypes.LegacyTx{Data: calldata})
	epochID := [12]byte{1}
	chain := &fakeChain{
		head: 2_000,
		logs: []ethtypes.Log{epochLiveLog(t, epochID, 1_500, tx)},
		txs:  map[common.Hash]*ethtypes.Transaction{tx.Hash(): tx},
	}
	m, err := gtypes.NewDKGManager(common.Address{}, chain)
	c.Assert(err, qt.IsNil)

	got, err := FinalizeCalldata(context.Background(), chain, m, epochID, 100)
	c.Assert(err, qt.IsNil)
	c.Assert(bytes.Equal(got, calldata), qt.IsTrue)

	_, err = FinalizeCalldata(context.Background(), chain, m, [12]byte{2}, 100)
	c.Assert(err, qt.ErrorMatches, "no EpochLive event.*")
}

// validFinalizeEpochCalldata packs a well-formed finalizeEpoch call whose
// transcript words are their own offsets (plus 5000) and whose public inputs
// name a committee of `committeeSize`. Arguments are supplied by ABI name so
// the fixture follows the contract signature.
func validFinalizeEpochCalldata(t *testing.T, committeeSize uint64) []byte {
	t.Helper()
	words := make([]*big.Int, 0, finalize.TranscriptWords)
	for i := 0; i < finalize.TranscriptWords; i++ {
		words = append(words, big.NewInt(int64(5000+i)))
	}
	transcript, err := encodeWords(words...)
	qt.Assert(t, err, qt.IsNil)
	inputs := make([]*big.Int, finalize.PublicInputWords)
	for i := range inputs {
		inputs[i] = big.NewInt(int64(7000 + i))
	}
	inputs[2] = new(big.Int).SetUint64(committeeSize)
	input, err := encodeWords(inputs...)
	qt.Assert(t, err, qt.IsNil)

	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	qt.Assert(t, err, qt.IsNil)
	method, ok := parsed.Methods["finalizeEpoch"]
	qt.Assert(t, ok, qt.IsTrue)
	byName := map[string]any{
		"epochId":          [12]byte{1},
		"transcriptDigest": [32]byte{3},
		"transcript":       transcript,
		"proof":            []byte{0xaa},
		"input":            input,
	}
	args := make([]any, 0, len(method.Inputs))
	for _, in := range method.Inputs {
		value, ok := byName[in.Name]
		qt.Assert(t, ok, qt.IsTrue, qt.Commentf("unexpected finalizeEpoch argument %q", in.Name))
		args = append(args, value)
	}
	calldata, err := parsed.Pack("finalizeEpoch", args...)
	qt.Assert(t, err, qt.IsNil)
	return calldata
}

// Key j's block of the finalize transcript holds P_j then one share
// commitment per committee slot: member p's commitment is at slot p−1 whether
// or not p contributed, and the decoder returns exactly the committee the
// public inputs name.
func TestFinalizeShareCommitmentsAreMemberIndexed(t *testing.T) {
	c := qt.New(t)
	data := validFinalizeEpochCalldata(t, 5)

	for _, key := range []uint8{0, 7, finalize.MaxKeys - 1} {
		idxs, points, err := FinalizeShareCommitments(data, key)
		c.Assert(err, qt.IsNil)
		c.Assert(idxs, qt.DeepEquals, []uint16{1, 2, 3, 4, 5})
		c.Assert(points, qt.HasLen, 5)
		for i, p := range points {
			q := finalize.ShareCommitmentOffset(int(key), i)
			c.Assert(p.X.Int64(), qt.Equals, int64(5000+q), qt.Commentf("key %d member %d", key, i+1))
			c.Assert(p.Y.Int64(), qt.Equals, int64(5000+q+1))
		}
		pk, err := FinalizePoolKey(data, key)
		c.Assert(err, qt.IsNil)
		c.Assert(pk.X.Int64(), qt.Equals, int64(5000+finalize.PoolKeyOffset(int(key))))
		c.Assert(pk.Y.Int64(), qt.Equals, int64(5000+finalize.PoolKeyOffset(int(key))+1))
	}

	for _, size := range []uint64{0, finalize.MaxParticipants + 1} {
		_, _, err := FinalizeShareCommitments(validFinalizeEpochCalldata(t, size), 0)
		c.Assert(err, qt.ErrorMatches, "committee size .* out of range.*", qt.Commentf("size %d", size))
	}
	_, _, err := FinalizeShareCommitments(data, finalize.MaxKeys)
	c.Assert(err, qt.ErrorMatches, "key .* outside the pool.*")
	_, _, err = FinalizeShareCommitments([]byte{1, 2, 3}, 0)
	c.Assert(err, qt.Not(qt.IsNil))
}

// Proving takes a while, so right before sending the finalizer re-reads the
// epoch and every committee member's record and compares them with the
// snapshot the statement was built from. A reorg that swapped an accepted
// dealer for another or changed a stored commitmentsHash leaves the count
// unchanged, so the count alone is blind to it; every such divergence must
// surface as ErrStale before any transaction, a finalization that landed as
// ErrAlreadyLive.
func TestRecheckRefusesAStaleAcceptedSet(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	epochID := [12]byte{4}
	a, b, d := common.HexToAddress("0xa"), common.HexToAddress("0xb"), common.HexToAddress("0xd")
	record := func(addr common.Address, index uint16, hash byte) gtypes.DKGTypesContributionRecord {
		return gtypes.DKGTypesContributionRecord{
			Contributor: addr, ContributorIndex: index, CommitmentsHash: [32]byte{hash}, Accepted: true,
		}
	}
	hash := func(hash byte) *big.Int {
		word := [32]byte{hash}
		return new(big.Int).SetBytes(word[:])
	}
	epoch := gtypes.IDKGManagerEpoch{
		Status:            phaseKeyAssembly,
		ContributionCount: 2,
		LotteryThreshold:  big.NewInt(0),
		Policy:            gtypes.DKGTypesEpochPolicy{Threshold: 2, CommitteeSize: 3},
	}
	snap := &snapshot{
		epoch:     epoch,
		committee: []common.Address{a, b, d},
		assignment: finalize.Assignment{
			ParticipantIndexes: []uint16{1, 3},
			ContributionHashes: []*big.Int{hash(0xa1), hash(0xd3)},
		},
	}
	chain := &fakeChain{
		epoch:   &epoch,
		records: map[common.Address]gtypes.DKGTypesContributionRecord{a: record(a, 1, 0xa1), d: record(d, 3, 0xd3)},
	}
	c.Assert(snap.recheck(ctx, chain, epochID), qt.IsNil)

	// Same count, one dealer swapped: b took d's place.
	chain.records = map[common.Address]gtypes.DKGTypesContributionRecord{a: record(a, 1, 0xa1), b: record(b, 2, 0xb2)}
	err := snap.recheck(ctx, chain, epochID)
	c.Assert(errors.Is(err, ErrStale), qt.IsTrue, qt.Commentf("%v", err))
	c.Assert(err, qt.ErrorMatches, ".*contribution of 0x.*[bB].* \\(member 2\\) changed while proving")

	// Same dealers, d's stored commitmentsHash changed.
	chain.records = map[common.Address]gtypes.DKGTypesContributionRecord{a: record(a, 1, 0xa1), d: record(d, 3, 0xd4)}
	c.Assert(errors.Is(snap.recheck(ctx, chain, epochID), ErrStale), qt.IsTrue)

	// Fewer accepted records than the statement, count notwithstanding.
	chain.records = map[common.Address]gtypes.DKGTypesContributionRecord{a: record(a, 1, 0xa1)}
	c.Assert(errors.Is(snap.recheck(ctx, chain, epochID), ErrStale), qt.IsTrue)

	// The accepted count moved.
	chain.records = map[common.Address]gtypes.DKGTypesContributionRecord{a: record(a, 1, 0xa1), d: record(d, 3, 0xd3)}
	moved := epoch
	moved.ContributionCount = 3
	chain.epoch = &moved
	c.Assert(errors.Is(snap.recheck(ctx, chain, epochID), ErrStale), qt.IsTrue)

	// Someone else finalized meanwhile: a benign race, not staleness.
	live := epoch
	live.Status = phaseLive
	chain.epoch = &live
	c.Assert(errors.Is(snap.recheck(ctx, chain, epochID), ErrAlreadyLive), qt.IsTrue)

	chain.epoch = &epoch
	c.Assert(snap.recheck(ctx, chain, epochID), qt.IsNil, qt.Commentf("the unchanged set still passes"))
}
