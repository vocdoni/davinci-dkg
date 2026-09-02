package node

import (
	"errors"
	"math/big"
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
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
