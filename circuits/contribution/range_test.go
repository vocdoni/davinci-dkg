package contribution

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// TestContributionCircuitRejectsOversizedShares is the load-bearing
// regression test for paper §5.1 ("range-check attack: craft a malicious
// s_i(j) = honest + k·q and confirm the contribution circuit rejects
// it"). Without the range check an attacker could submit a polynomial
// share that wraps around the BabyJubJub subgroup order; the in-circuit
// `AssertIsLessOrEqual(c.Shares[i], subgroupOrderMinusOne)` constraint
// is the fix.
//
// We craft `tampered = honest + 7·r_bjj` (which still fits in the BN254
// scalar field but exceeds the subgroup order) and assert the circuit
// rejects it. A future refactor that removes the range check would let
// SolvingSucceeded return true here and fail the test.
func TestContributionCircuitRejectsOversizedShares(t *testing.T) {
	c := qt.New(t)

	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	// Replace Shares[0] with `honest + 7·r_bjj`. The witness is in the
	// BN254 scalar field (~7r); 7r < p so this stays representable but
	// is well outside [0, r). The circuit's `AssertIsLessOrEqual(...,
	// subgroupOrderMinusOne)` constraint must reject it.
	//
	// witness.Shares is []frontend.Variable (assigned to *big.Int by
	// BuildWitness); type-assert before doing big-int arithmetic.
	honest, ok := witness.Shares[0].(*big.Int)
	c.Assert(ok, qt.IsTrue, qt.Commentf("expected witness.Shares[0] to be *big.Int"))
	rbjj := group.ScalarField()
	tampered := new(big.Int).Mul(big.NewInt(7), rbjj)
	tampered.Add(tampered, honest)
	witness.Shares[0] = tampered

	// Use SolvingFailed (not ProveAndVerify) so we don't pay the cost of
	// running through the prover — the constraint solver alone is enough
	// to confirm the witness violates an in-circuit assertion.
	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// TestContributionCircuitRejectsOversizedCoefficients exercises the
// companion range check on Coefficients (paper §5.1 line 879). A
// coefficient ≥ r_bjj would make Feldman commitments inconsistent under
// modular reduction; the circuit asserts each coefficient < r_bjj.
func TestContributionCircuitRejectsOversizedCoefficients(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	rbjj := group.ScalarField()
	// Make Coefficients[0] equal to r_bjj exactly (out of range by 1).
	asn.Coefficients[0] = new(big.Int).Set(rbjj)

	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}
