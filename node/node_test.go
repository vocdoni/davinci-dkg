package node

import (
	"context"
	"errors"
	"math/big"
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/web3"
)

// d_i = Σ_j f_j(i) over every accepted contribution. A partial sum is a wrong
// share, and a wrong share produces partial decryptions that look valid but
// interpolate to garbage, so the aggregation must refuse anything less than
// the full set instead of caching it.
func TestSumRecoveredSharesRequiresEveryAcceptedContribution(t *testing.T) {
	c := qt.New(t)
	q := group.ScalarField()
	a, b := big.NewInt(5), new(big.Int).Sub(q, big.NewInt(2))

	sum, err := sumRecoveredShares([]*big.Int{a, b}, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(sum.String(), qt.Equals, "3") // (5 + q − 2) mod q

	_, err = sumRecoveredShares([]*big.Int{a}, 2)
	c.Assert(err, qt.ErrorMatches, ".*recovered 1/2.*")

	_, err = sumRecoveredShares(nil, 0)
	c.Assert(err, qt.Not(qt.IsNil))

	_, err = sumRecoveredShares([]*big.Int{a, new(big.Int).Sub(q, a)}, 2)
	c.Assert(err, qt.ErrorMatches, ".*zero.*")
}

// Every committee member gets a distinct slot in the finalize/combine
// rotation, so with a live population exactly one node normally pays for the
// transaction; the rotation start moves with the seed and the salt.
func TestStaggerSlotIsAPermutationOfTheCommittee(t *testing.T) {
	c := qt.New(t)
	seed := common.HexToHash("0x00000000000000000000000000000000000000000000000000000000000000f3") // 243 = 3 mod 5
	const n = uint16(5)

	var slots []int
	for idx := uint16(1); idx <= n; idx++ {
		slots = append(slots, int(staggerSlot(seed, 0, idx, n)))
	}
	sort.Ints(slots)
	c.Assert(slots, qt.DeepEquals, []int{0, 1, 2, 3, 4})

	// seed mod n == 3 → member 4 (index 3) opens the rotation.
	c.Assert(staggerSlot(seed, 0, 4, n), qt.Equals, uint64(0))
	c.Assert(staggerSlot(seed, 0, 5, n), qt.Equals, uint64(1))
	c.Assert(staggerSlot(seed, 0, 1, n), qt.Equals, uint64(2))
	// A salt of 1 rotates the start by one member.
	c.Assert(staggerSlot(seed, 1, 5, n), qt.Equals, uint64(0))
	// Degenerate committee: everyone is slot 0.
	c.Assert(staggerSlot(seed, 7, 1, 1), qt.Equals, uint64(0))
}

// fakeEpochReader answers getEpoch from an epoch → status map (missing =
// None) and counts the reads the lifecycle scan spends past its window.
type fakeEpochReader struct {
	status map[[12]byte]uint8
	err    error
	calls  int
}

func (f *fakeEpochReader) GetEpoch(_ context.Context, id [12]byte) (web3.EpochView, error) {
	f.calls++
	if f.err != nil {
		return web3.EpochView{}, f.err
	}
	return web3.EpochView{Status: f.status[id]}, nil
}

// The lifecycle scan covers the newest epochLookback epochs minus the ones
// already seen terminal, oldest first, so an epoch that qualified but was
// never finalized stays discoverable across cadences and restarts. Past the
// window it spends exactly one getEpoch per tick when the epoch just outside
// it is closed.
func TestEpochsToVisitCoversTheLookbackMinusTerminal(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	n := &Node{terminal: map[[12]byte]bool{}, finalizeRetry: map[[12]byte]*serviceBackoff{}}
	const prefix = uint32(0xdead)
	id := func(nonce uint64) [12]byte { return web3.EpochID(prefix, nonce) }
	chain := &fakeEpochReader{status: map[[12]byte]uint8{id(20 - epochLookback): epochLive}}

	c.Assert(n.epochsToVisit(ctx, chain, prefix, 1), qt.DeepEquals, [][12]byte{id(1)})
	c.Assert(n.epochsToVisit(ctx, chain, prefix, 3), qt.DeepEquals, [][12]byte{id(1), id(2), id(3)})
	c.Assert(chain.calls, qt.Equals, 0, qt.Commentf("nothing older than nonce 1 to look at"))

	visited := n.epochsToVisit(ctx, chain, prefix, 20)
	c.Assert(visited, qt.HasLen, epochLookback)
	c.Assert(visited[0], qt.Equals, id(20-epochLookback+1))
	c.Assert(visited[len(visited)-1], qt.Equals, id(20))
	c.Assert(chain.calls, qt.Equals, 1, qt.Commentf("the Live epoch just outside the window ends the walk"))

	n.finalizeRetry[id(19)] = &serviceBackoff{}
	n.finish(id(19))
	n.finish(id(13))
	visited = n.epochsToVisit(ctx, chain, prefix, 20)
	c.Assert(visited, qt.HasLen, epochLookback-2)
	for _, v := range visited {
		c.Assert(v, qt.Not(qt.Equals), id(19))
		c.Assert(v, qt.Not(qt.Equals), id(13))
	}
	_, retrying := n.finalizeRetry[id(19)]
	c.Assert(retrying, qt.IsFalse, qt.Commentf("a terminal epoch drops its finalize backoff"))
}

// Past the fixed window the scan keeps stepping back while the chain still
// reports unfinished epochs and stops at the first closed one (or nonce 1),
// so an epoch that qualified but was never finalized stays discoverable
// however many cadences have passed (docs/pool-keys-v4.md §10); a Live epoch
// right outside the window ends the walk at once.
func TestEpochsToVisitWalksPastTheWindowWhileEpochsAreUnfinished(t *testing.T) {
	c := qt.New(t)
	ctx := context.Background()
	const prefix = uint32(0xdead)
	const nonce = uint64(20)
	id := func(n uint64) [12]byte { return web3.EpochID(prefix, n) }
	// The window is [first, nonce]; out1 is the epoch right outside it.
	const first = nonce - epochLookback + 1
	const out1, out2, out3, out4 = first - 1, first - 2, first - 3, first - 4
	window := make([][12]byte, 0, epochLookback)
	for k := first; k <= nonce; k++ {
		window = append(window, id(k))
	}
	withOlder := func(older ...[12]byte) [][12]byte { return append(older, window...) }
	fresh := func() *Node {
		return &Node{terminal: map[[12]byte]bool{}, finalizeRetry: map[[12]byte]*serviceBackoff{}}
	}

	// Two unfinished epochs outside the window are visited, oldest first;
	// the Live epoch behind them ends the walk and hides everything older.
	chain := &fakeEpochReader{status: map[[12]byte]uint8{
		id(out1): epochKeyAssembly, id(out2): epochKeyAssembly, id(out3): epochLive, id(out4): epochKeyAssembly,
	}}
	n := fresh()
	c.Assert(n.epochsToVisit(ctx, chain, prefix, nonce), qt.DeepEquals, withOlder(id(out2), id(out1)))
	c.Assert(chain.calls, qt.Equals, 3)

	// Live right outside the window: nothing older is visited or even read.
	chain = &fakeEpochReader{status: map[[12]byte]uint8{id(out1): epochLive, id(out2): epochKeyAssembly}}
	c.Assert(fresh().epochsToVisit(ctx, chain, prefix, nonce), qt.DeepEquals, window)
	c.Assert(chain.calls, qt.Equals, 1)

	// An unfinished epoch this node already finished with (not selected) is
	// skipped but does not end the walk, and neither does a dead
	// CommitteeSelection epoch: the KeyAssembly epoch behind them is found.
	chain = &fakeEpochReader{status: map[[12]byte]uint8{
		id(out1): epochKeyAssembly, id(out2): epochCommitteeSelection,
		id(out3): epochKeyAssembly, id(out4): epochAborted,
	}}
	n = fresh()
	n.finish(id(out1))
	c.Assert(n.epochsToVisit(ctx, chain, prefix, nonce), qt.DeepEquals, withOlder(id(out3), id(out2)))
	c.Assert(chain.calls, qt.Equals, 4)

	// The walk ends at nonce 1 when every older epoch is unfinished.
	chain = &fakeEpochReader{status: map[[12]byte]uint8{id(1): epochKeyAssembly}}
	c.Assert(fresh().epochsToVisit(ctx, chain, prefix, epochLookback+1), qt.HasLen, epochLookback+1)
	c.Assert(chain.calls, qt.Equals, 1)

	// A read failure ends the walk for this tick without dropping the window.
	chain = &fakeEpochReader{err: errors.New("rpc down")}
	c.Assert(fresh().epochsToVisit(ctx, chain, prefix, nonce), qt.DeepEquals, window)
	c.Assert(chain.calls, qt.Equals, 1)
}

// A finalize attempt only stops when the race is really lost: AlreadyLive
// from the contract, or the epoch no longer in KeyAssembly. Any other revert
// (InvalidFinalization, a stale proof, a mined revert) must be retried.
func TestFinalizeRaceLostOnlyOnAlreadyLiveOrPhaseChange(t *testing.T) {
	c := qt.New(t)
	c.Assert(finalizeRaceLost("execution reverted: AlreadyLive", epochKeyAssembly), qt.IsTrue)
	c.Assert(finalizeRaceLost("execution reverted", epochLive), qt.IsTrue)
	c.Assert(finalizeRaceLost("transaction 0xabc reverted (status 0)", epochAborted), qt.IsTrue)

	c.Assert(finalizeRaceLost("execution reverted", epochKeyAssembly), qt.IsFalse)
	c.Assert(finalizeRaceLost("execution reverted: InvalidFinalization", epochKeyAssembly), qt.IsFalse)
	c.Assert(finalizeRaceLost("execution reverted: InvalidPhase", epochKeyAssembly), qt.IsFalse)
	c.Assert(finalizeRaceLost("transaction 0xabc reverted (status 0)", epochKeyAssembly), qt.IsFalse)
	c.Assert(finalizeRaceLost("dial tcp: connection refused", epochKeyAssembly), qt.IsFalse)
}

// The tx manager reports a mined-but-reverted transaction as
// "reverted (status 0)"; that is as final as an eth_call revert.
func TestIsPermanentRevertRecognisesMinedReverts(t *testing.T) {
	c := qt.New(t)
	c.Assert(isPermanentRevert(errors.New("transaction 0xabc reverted (status 0)")), qt.IsTrue)
	c.Assert(isPermanentRevert(errors.New("execution reverted")), qt.IsTrue)
	c.Assert(isPermanentRevert(errors.New("timeout waiting for transaction 0xabc")), qt.IsFalse)
	c.Assert(isPermanentRevert(errors.New("rpc: the node reverted to a snapshot")), qt.IsFalse)
	c.Assert(isPermanentRevert(nil), qt.IsFalse)
}

// The startup banner must not leak RPC credentials (API keys live in the
// URL path or userinfo on most providers).
func TestRPCHostHidesCredentials(t *testing.T) {
	c := qt.New(t)
	c.Assert(rpcHost("https://user:secret@eth-sepolia.example.com/v3/apikey123"), qt.Equals, "eth-sepolia.example.com")
	c.Assert(rpcHost("http://127.0.0.1:8545"), qt.Equals, "127.0.0.1:8545")
	c.Assert(rpcHost("ws://[::1]:8546/ws?key=1"), qt.Equals, "[::1]:8546")
	c.Assert(rpcHost("not a url"), qt.Equals, "<unparseable rpc url>")
}
