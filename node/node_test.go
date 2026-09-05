package node

import (
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

// The activation reserve is the contiguous run of activated keys from the
// pool's next index: registration claims in order and reverts on the first
// gap, so an activated key past a gap counts for nothing. The key to activate
// is always the first gap at or after nextIndex.
func TestNextKeyToActivateUsesContiguityNotPopcount(t *testing.T) {
	c := qt.New(t)
	pick := func(next, activated, ahead uint8) (int, bool) {
		k, ok := nextKeyToActivate(next, activated, ahead)
		return int(k), ok
	}
	// Fresh Live epoch: key 0 first.
	k, ok := pick(0, 0b0000_0000, 2)
	c.Assert(ok, qt.IsTrue)
	c.Assert(k, qt.Equals, 0)
	// Keys 0 and 1 up, nothing claimed, ahead=2: reserve satisfied.
	_, ok = pick(0, 0b0000_0011, 2)
	c.Assert(ok, qt.IsFalse)
	// Keys 0 and 2 up but 1 missing: popcount says two, contiguity says one.
	k, ok = pick(0, 0b0000_0101, 2)
	c.Assert(ok, qt.IsTrue)
	c.Assert(k, qt.Equals, 1)
	// Key 0 claimed (next=1), keys 0,1 up: need key 2.
	k, ok = pick(1, 0b0000_0011, 2)
	c.Assert(ok, qt.IsTrue)
	c.Assert(k, qt.Equals, 2)
	// Near the end of the pool the window clips at MaxK.
	k, ok = pick(7, 0b0111_1111, 2)
	c.Assert(ok, qt.IsTrue)
	c.Assert(k, qt.Equals, 7)
	_, ok = pick(7, 0b1111_1111, 2)
	c.Assert(ok, qt.IsFalse)
	_, ok = pick(8, 0b1111_1111, 2)
	c.Assert(ok, qt.IsFalse)
}

// Older Live epochs whose pools still need activations stay on the visiting
// list next to the newest two, oldest first, until they turn terminal; the
// list is bounded by dropping the oldest tracked epoch.
func TestEpochsToVisitCoversTrackedLiveEpochsAndIsBounded(t *testing.T) {
	c := qt.New(t)
	n := &Node{terminal: map[[12]byte]bool{}, liveEpochs: map[[12]byte]uint64{}}
	const prefix = uint32(0xdead)
	id := func(nonce uint64) [12]byte { return web3.EpochID(prefix, nonce) }

	c.Assert(n.epochsToVisit(prefix, 1), qt.DeepEquals, [][12]byte{id(1)})
	c.Assert(n.epochsToVisit(prefix, 5), qt.DeepEquals, [][12]byte{id(4), id(5)})

	n.trackLive(id(2), 2)
	n.trackLive(id(4), 4) // already among the newest two: no duplicate
	c.Assert(n.epochsToVisit(prefix, 5), qt.DeepEquals, [][12]byte{id(2), id(4), id(5)})

	n.finish(id(2))
	c.Assert(n.epochsToVisit(prefix, 5), qt.DeepEquals, [][12]byte{id(4), id(5)})
	_, tracked := n.liveEpochs[id(2)]
	c.Assert(tracked, qt.IsFalse)

	for nonce := uint64(10); nonce < 10+maxTrackedLiveEpochs+3; nonce++ {
		n.trackLive(id(nonce), nonce)
	}
	c.Assert(n.liveEpochs, qt.HasLen, maxTrackedLiveEpochs)
	_, tracked = n.liveEpochs[id(10)]
	c.Assert(tracked, qt.IsFalse, qt.Commentf("the oldest tracked epoch is dropped first"))
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
