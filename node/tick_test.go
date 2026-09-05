package node

import (
	"context"
	"errors"
	"math/big"
	"testing"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/web3"
)

// fakeHeadChain answers HeaderByNumber with a scripted head and counts the reads.
type fakeHeadChain struct {
	head  uint64
	time  uint64
	err   error
	reads int
}

func (f *fakeHeadChain) HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error) {
	f.reads++
	if f.err != nil {
		return nil, f.err
	}
	return &ethtypes.Header{Number: new(big.Int).SetUint64(f.head), Time: f.time}, nil
}

// A tick reads the head once, and when it is the head the previous complete
// tick already processed nothing on chain can have changed: the tick is
// skipped. A tick that failed half way is not "complete", so the same head
// is processed again; a lower head (another endpoint behind the last one)
// is a change like any other.
func TestBeginTickSkipsAnUnchangedHead(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	n := newTestNode()
	chain := &fakeHeadChain{head: 100, time: 1_700_000_000}

	tc, err := n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNil)
	c.Assert(tc, qt.IsNotNil, qt.Commentf("the first tick always runs"))
	c.Assert(tc.head, qt.Equals, uint64(100))
	c.Assert(tc.headTime, qt.Equals, uint64(1_700_000_000))
	n.endTick(tc)

	tc, err = n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNil)
	c.Assert(tc, qt.IsNil, qt.Commentf("no new block: the tick must be skipped"))
	c.Assert(chain.reads, qt.Equals, 2, qt.Commentf("the skip decision costs exactly the one head read"))

	chain.head = 101
	tc, err = n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNil)
	c.Assert(tc, qt.IsNotNil)
	tc.failed = true
	n.endTick(tc)
	tc, err = n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNil)
	c.Assert(tc, qt.IsNotNil, qt.Commentf("a failed tick must be repeated at the same head"))
	n.endTick(tc)
	tc, err = n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNil)
	c.Assert(tc, qt.IsNil)

	chain.head = 99 // an endpoint behind the previous one
	tc, err = n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNil)
	c.Assert(tc, qt.IsNotNil, qt.Commentf("a different head, even a lower one, is processed"))

	chain.err = errors.New("rpc down")
	_, err = n.beginTick(ctx, chain)
	c.Assert(err, qt.IsNotNil)
}

// countingEpochReader answers getEpoch from a status map and counts reads.
type countingEpochReader struct {
	status map[[12]byte]uint8
	calls  int
}

func (f *countingEpochReader) GetEpoch(_ context.Context, id [12]byte) (web3.EpochView, error) {
	f.calls++
	if _, ok := f.status[id]; !ok {
		return web3.EpochView{}, errors.New("no such epoch")
	}
	return web3.EpochView{Status: f.status[id], Policy: web3.EpochPolicy{Threshold: 2}}, nil
}

// A Live (or Aborted, or Completed) epoch's record never changes on chain,
// so it is read once for the life of the process; an epoch still in
// CommitteeSelection or KeyAssembly is read once per tick, however many
// steps of the tick look at it. The cache is bounded and evicts the oldest
// nonce first.
func TestEpochCacheReadsClosedEpochsOnce(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	const prefix = uint32(0xdead)
	id := func(nonce uint64) [12]byte { return web3.EpochID(prefix, nonce) }
	chain := &countingEpochReader{status: map[[12]byte]uint8{
		id(1): epochLive, id(2): epochAborted, id(3): epochKeyAssembly, id(4): epochCommitteeSelection,
	}}
	n := newTestNode()

	tick1 := &tickCtx{epochs: map[[12]byte]epochView{}}
	for _, nonce := range []uint64{1, 2, 3, 4, 1, 2, 3, 4} {
		e, err := n.epoch(ctx, tick1, chain, id(nonce))
		c.Assert(err, qt.IsNil)
		c.Assert(e.Status, qt.Equals, chain.status[id(nonce)])
	}
	c.Assert(chain.calls, qt.Equals, 4, qt.Commentf("every epoch is read once within a tick"))

	tick2 := &tickCtx{epochs: map[[12]byte]epochView{}}
	for _, nonce := range []uint64{1, 2, 3, 4} {
		_, err := n.epoch(ctx, tick2, chain, id(nonce))
		c.Assert(err, qt.IsNil)
	}
	c.Assert(chain.calls, qt.Equals, 6, qt.Commentf("only the two open epochs are read again on the next tick"))
	_, live := n.epochCache[id(1)]
	_, open := n.epochCache[id(3)]
	c.Assert(live, qt.IsTrue)
	c.Assert(open, qt.IsFalse)

	// The lifecycle walk reads through the same cache: a Live epoch just
	// outside the window costs nothing after the first tick.
	reader := cachingEpochReader{n: n, tc: tick2, chain: chain}
	_, err := reader.GetEpoch(ctx, id(1))
	c.Assert(err, qt.IsNil)
	c.Assert(chain.calls, qt.Equals, 6)

	// Without a tick context (a read outside the loop) nothing is cached for
	// an open epoch and a read error is returned as is.
	_, err = n.epoch(ctx, nil, chain, id(3))
	c.Assert(err, qt.IsNil)
	_, err = n.epoch(ctx, nil, chain, id(9))
	c.Assert(err, qt.IsNotNil)

	// Bounded: past maxCachedEpochs the lowest nonce goes first.
	for nonce := uint64(10); nonce < 10+maxCachedEpochs; nonce++ {
		n.cacheEpoch(id(nonce), epochView{Status: epochLive})
	}
	c.Assert(n.epochCache, qt.HasLen, maxCachedEpochs)
	_, kept := n.epochCache[id(1)]
	c.Assert(kept, qt.IsFalse, qt.Commentf("the oldest epoch is evicted first"))
	_, kept = n.epochCache[id(10+maxCachedEpochs-1)]
	c.Assert(kept, qt.IsTrue)
}

// The histogram is the per-tick average of the requests issued between two
// snapshots, most frequent first, and says so when nothing was sent.
func TestRPCHistogramAveragesTheDeltaPerTick(t *testing.T) {
	c := qt.New(t)
	prev := map[string]uint64{"eth_call": 100, "eth_getLogs": 10, "eth_getBlockByNumber": 5}
	cur := map[string]uint64{"eth_call": 250, "eth_getLogs": 30, "eth_getBlockByNumber": 25, "batch": 2}
	c.Assert(rpcHistogram(prev, cur, 20), qt.Equals, "eth_call=7.5 eth_getBlockByNumber=1.0 eth_getLogs=1.0 batch=0.1")
	c.Assert(rpcHistogram(cur, cur, 20), qt.Equals, "none")
}
