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
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
)

// validSubmitContributionCalldata packs a well-formed submitContribution call
// whose transcript is the 8·MaxN-word layout the contract expects.
func validSubmitContributionCalldata(t *testing.T) (calldata, transcript []byte) {
	t.Helper()
	words := make([]*big.Int, 0, 8*ccommon.MaxN)
	for i := 0; i < 8*ccommon.MaxN; i++ {
		words = append(words, big.NewInt(int64(1000+i)))
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
	calldata, transcript := validSubmitContributionCalldata(t)
	got, err := ContributionTranscript(calldata)
	c.Assert(err, qt.IsNil)
	c.Assert(bytes.Equal(got, transcript), qt.IsTrue)
}

// Calldata comes from arbitrary transactions located through an event log; a
// hostile or merely unexpected payload must yield an error, never a panic.
func TestContributionTranscriptRejectsHostileCalldata(t *testing.T) {
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
		"transcript wrong size": func() []byte { out := bytes.Clone(valid); out[4+7*32+31] = 0x20; return out }(),
	}
	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("panicked on hostile calldata: %v", r)
				}
			}()
			tr, err := ContributionTranscript(data)
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

func (*fakeChain) CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error) {
	return nil, errFakeUnsupported
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
