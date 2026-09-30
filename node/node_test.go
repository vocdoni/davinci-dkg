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
// however many cadences have passed; a Live epoch
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

// epochDead is abortEpoch's condition: past the selection deadline without a
// full committee, or past the assembly deadline short of
// minValidContributions. Both deadlines are inclusive on chain, and every
// other state can still progress and must never be aborted.
func TestEpochDeadMirrorsAbortEpoch(t *testing.T) {
	c := qt.New(t)
	const sel, asm = uint64(108), uint64(116)
	epoch := func(status uint8, claimed, contributions uint16) epochView {
		return epochView{
			Status: status, ClaimedCount: claimed, ContributionCount: contributions,
			Policy: web3.EpochPolicy{
				CommitteeSize: 4, MinValidContributions: 3,
				CommitteeSelectionDeadlineBlock: sel, KeyAssemblyDeadlineBlock: asm,
			},
		}
	}
	for _, tc := range []struct {
		name string
		e    epochView
		head uint64
		dead bool
	}{
		{"selection before the deadline", epoch(epochCommitteeSelection, 0, 0), sel - 1, false},
		{"selection at the deadline, a claim still lands", epoch(epochCommitteeSelection, 3, 0), sel, false},
		{"selection one block past, committee short", epoch(epochCommitteeSelection, 3, 0), sel + 1, true},
		{"selection one block past, nobody claimed", epoch(epochCommitteeSelection, 0, 0), sel + 1, true},
		{"selection long past both deadlines", epoch(epochCommitteeSelection, 1, 0), asm + 1000, true},
		{"selection with a full count (inconsistent record)", epoch(epochCommitteeSelection, 4, 0), sel + 1, false},
		{"assembly before the deadline", epoch(epochKeyAssembly, 4, 0), asm - 1, false},
		{"assembly past the selection deadline only", epoch(epochKeyAssembly, 4, 0), sel + 1, false},
		{"assembly at the deadline, a contribution still lands", epoch(epochKeyAssembly, 4, 2), asm, false},
		{"assembly one block past, one short", epoch(epochKeyAssembly, 4, 2), asm + 1, true},
		{"assembly one block past, none", epoch(epochKeyAssembly, 4, 0), asm + 1, true},
		{"assembly with enough contributions can still finalize", epoch(epochKeyAssembly, 4, 3), asm + 1000, false},
		{"assembly with more than enough", epoch(epochKeyAssembly, 4, 4), asm + 1, false},
		{"live", epoch(epochLive, 4, 3), asm + 1000, false},
		{"aborted", epoch(epochAborted, 0, 0), asm + 1000, false},
		{"completed", epoch(epochCompleted, 4, 4), asm + 1000, false},
		{"none", epoch(0, 0, 0), asm + 1000, false},
	} {
		c.Run(tc.name, func(c *qt.C) {
			c.Assert(epochDead(tc.e, tc.head), qt.Equals, tc.dead)
		})
	}
}

// A failed abort is classified from the epoch's status read afterwards and
// the decoded reason: Aborted means another abort won, InvalidPhase with the
// epoch still open means the endpoint does not see it dead yet, and anything
// else (an RPC fault, a funding error, a mined revert with the epoch still
// open) is a failure the gate retries.
func TestAbortOutcomeOfClassifiesRefusals(t *testing.T) {
	c := qt.New(t)
	c.Assert(abortOutcomeOf("transaction 0xabc reverted (status 0)", epochAborted), qt.Equals, abortLost)
	c.Assert(abortOutcomeOf("dial tcp: connection refused", epochAborted), qt.Equals, abortLost)
	c.Assert(abortOutcomeOf("execution reverted: InvalidPhase", epochAborted), qt.Equals, abortLost)
	c.Assert(abortOutcomeOf("execution reverted: InvalidPhase", epochCommitteeSelection), qt.Equals, abortRefused)
	c.Assert(abortOutcomeOf("execution reverted: InvalidPhase", epochKeyAssembly), qt.Equals, abortRefused)

	c.Assert(abortOutcomeOf("transaction 0xabc reverted (status 0)", epochCommitteeSelection), qt.Equals, abortFailed)
	c.Assert(abortOutcomeOf("gas required exceeds allowance (0)", epochKeyAssembly), qt.Equals, abortFailed)
	c.Assert(abortOutcomeOf("timeout waiting for transaction 0xabc", epochCommitteeSelection), qt.Equals, abortFailed)
	c.Assert(abortOutcomeOf("execution reverted", epochKeyAssembly), qt.Equals, abortFailed)
}

// A dead newest epoch after a healthy one is aborted as soon as it is dead;
// each Aborted epoch right before it doubles the wait, so a fleet that
// cannot fill a committee backs off towards the cadence instead of churning
// an epoch every few blocks. The streak walk stops at the first epoch that
// is not Aborted and at nonce 1, reads each Aborted record once, and reports
// a failed read instead of cutting the streak (and the backoff) short.
func TestAbortBackoffGrowsWithTheAbortedStreak(t *testing.T) {
	c := qt.New(t)
	c.Assert(abortBackoff(0), qt.Equals, uint64(0))
	c.Assert(abortBackoff(1), qt.Equals, uint64(autoRetryBlocks))
	c.Assert(abortBackoff(2), qt.Equals, uint64(2*autoRetryBlocks))
	c.Assert(abortBackoff(5), qt.Equals, uint64(16*autoRetryBlocks))
	c.Assert(abortBackoff(maxAbortStreak), qt.Equals, uint64(655_360), qt.Commentf("weeks of blocks: past any cadence"))

	ctx := context.Background()
	const prefix = uint32(0xdead)
	id := func(nonce uint64) [12]byte { return web3.EpochID(prefix, nonce) }
	chain := &countingEpochReader{status: map[[12]byte]uint8{
		id(1): epochAborted, id(2): epochLive, id(3): epochAborted, id(4): epochAborted, id(5): epochCommitteeSelection,
		id(8): epochAborted, id(9): epochCommitteeSelection, // 7 is missing: its read fails
	}}
	n := newTestNode()
	tick := &tickCtx{epochs: map[[12]byte]epochView{}}
	streak := func(nonce uint64) uint {
		got, err := n.abortStreak(ctx, tick, chain, prefix, nonce)
		c.Assert(err, qt.IsNil)
		return got
	}
	c.Assert(streak(5), qt.Equals, uint(2), qt.Commentf("4 and 3, then Live 2 ends it"))
	c.Assert(streak(3), qt.Equals, uint(0), qt.Commentf("2 is Live"))
	c.Assert(streak(2), qt.Equals, uint(1), qt.Commentf("the walk ends at nonce 1"))
	c.Assert(streak(1), qt.Equals, uint(0))
	_, err := n.abortStreak(ctx, tick, chain, prefix, 9)
	c.Assert(err, qt.IsNotNil, qt.Commentf("8 is Aborted, 7 cannot be read"))

	calls := chain.calls
	tick = &tickCtx{epochs: map[[12]byte]epochView{}}
	c.Assert(streak(5), qt.Equals, uint(2))
	c.Assert(chain.calls, qt.Equals, calls, qt.Commentf("closed epochs come from the cache on later ticks"))
}

// abortReady is epochDead plus the backoff, counted from the deadline that
// killed the epoch: the selection deadline in CommitteeSelection, the
// assembly deadline in KeyAssembly. The node aborts one block after
// deadline+backoff, never at it.
func TestAbortReadyWaitsOutTheBackoff(t *testing.T) {
	c := qt.New(t)
	const sel, asm = uint64(100), uint64(125)
	policy := web3.EpochPolicy{
		CommitteeSize: 2, MinValidContributions: 2,
		CommitteeSelectionDeadlineBlock: sel, KeyAssemblyDeadlineBlock: asm,
	}
	selecting := epochView{Status: epochCommitteeSelection, ClaimedCount: 1, Policy: policy}
	assembling := epochView{Status: epochKeyAssembly, ClaimedCount: 2, ContributionCount: 1, Policy: policy}
	finalizable := epochView{Status: epochKeyAssembly, ClaimedCount: 2, ContributionCount: 2, Policy: policy}

	c.Assert(abortReady(selecting, sel, 0), qt.IsFalse)
	c.Assert(abortReady(selecting, sel+1, 0), qt.IsTrue, qt.Commentf("no streak: as soon as it is dead"))
	c.Assert(abortReady(selecting, sel+abortBackoff(2), 2), qt.IsFalse)
	c.Assert(abortReady(selecting, sel+abortBackoff(2)+1, 2), qt.IsTrue)

	c.Assert(abortReady(assembling, asm, 0), qt.IsFalse)
	c.Assert(abortReady(assembling, asm+1, 0), qt.IsTrue)
	c.Assert(abortReady(assembling, sel+abortBackoff(1)+1, 1), qt.IsFalse,
		qt.Commentf("KeyAssembly counts from the assembly deadline"))
	c.Assert(abortReady(assembling, asm+abortBackoff(1), 1), qt.IsFalse)
	c.Assert(abortReady(assembling, asm+abortBackoff(1)+1, 1), qt.IsTrue)

	c.Assert(abortReady(finalizable, asm+1_000_000, 0), qt.IsFalse, qt.Commentf("it can still be finalized"))
	c.Assert(abortReady(epochView{Status: epochLive, Policy: policy}, asm+1_000_000, 0), qt.IsFalse)
}

// The gate makes one attempt per trigger: the same trigger again only
// autoRetryBlocks after the last attempt and only once that attempt's
// transaction has left the txmanager, while a new trigger goes at once.
func TestAutoGateSpacesAttemptsPerTrigger(t *testing.T) {
	c := qt.New(t)
	id := func(nonce uint64) [12]byte { return web3.EpochID(0xdead, nonce) }
	dead := autoTrigger{action: autoAbort, slot: poolSlot{epoch: id(4)}}
	const h = uint64(1_000)
	var g autoGate

	c.Assert(g.due(dead, h, false), qt.IsTrue, qt.Commentf("first attempt"))
	c.Assert(g.due(dead, h, false), qt.IsFalse, qt.Commentf("same head"))
	c.Assert(g.due(dead, h+1, false), qt.IsFalse, qt.Commentf("the attempt is recorded"))
	c.Assert(g.due(dead, h+autoRetryBlocks-1, false), qt.IsFalse)
	c.Assert(g.due(dead, h+autoRetryBlocks, true), qt.IsFalse, qt.Commentf("its transaction is still pending"))
	c.Assert(g.due(dead, h+autoRetryBlocks, false), qt.IsTrue, qt.Commentf("retry, exactly autoRetryBlocks later"))
	c.Assert(g.due(dead, h+autoRetryBlocks+1, false), qt.IsFalse, qt.Commentf("the retry is recorded too"))
	c.Assert(g.due(dead, h+2*autoRetryBlocks+5, true), qt.IsFalse)
	c.Assert(g.due(dead, h+2*autoRetryBlocks+5, false), qt.IsTrue, qt.Commentf("once the transaction is gone"))

	// The abort landed: the Aborted epoch is a new (early-create) trigger,
	// due at once even with the abort's transaction not yet pruned.
	aborted := autoTrigger{action: autoEarly, slot: poolSlot{epoch: id(4), key: abortedEpochSlot}}
	c.Assert(g.due(aborted, h+2*autoRetryBlocks+6, true), qt.IsTrue)
	c.Assert(g.due(aborted, h+2*autoRetryBlocks+7, false), qt.IsFalse)

	// A registration moving the pool cursor, or a new cadence threshold, is
	// a new trigger as well.
	drained := autoTrigger{action: autoEarly, slot: poolSlot{epoch: id(5), key: 15}}
	c.Assert(g.due(drained, h+2*autoRetryBlocks+8, false), qt.IsTrue)
	c.Assert(g.due(autoTrigger{action: autoEarly, slot: poolSlot{epoch: id(5), key: 16}}, h+2*autoRetryBlocks+8, false), qt.IsTrue)
	cadence := autoTrigger{action: autoCadence, next: 2_000}
	c.Assert(g.due(cadence, 2_000, false), qt.IsTrue)
	c.Assert(g.due(cadence, 2_001, false), qt.IsFalse)
	c.Assert(g.due(autoTrigger{action: autoCadence, next: 2_100}, 2_100, false), qt.IsTrue)
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
