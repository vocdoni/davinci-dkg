package contribution

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
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
	// witness.Shares is [MaxKeys][MaxRecipients]frontend.Variable (assigned
	// to *big.Int by BuildWitness); type-assert before big-int arithmetic.
	honest, ok := witness.Shares[0][0].(*big.Int)
	c.Assert(ok, qt.IsTrue, qt.Commentf("expected witness.Shares[0][0] to be *big.Int"))
	rbjj := group.ScalarField()
	tampered := new(big.Int).Mul(big.NewInt(7), rbjj)
	tampered.Add(tampered, honest)
	witness.Shares[0][0] = tampered

	// Use SolvingFailed (not ProveAndVerify) so we don't pay the cost of
	// running through the prover — the constraint solver alone is enough
	// to confirm the witness violates an in-circuit assertion.
	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// TestContributionCircuitRejectsOversizedConstantTerms exercises the range
// check on the one coefficient that is still a witness: the constant term
// whose commitment pins the pool-key contribution. Higher coefficients are
// implied by the shares and carry cofactor preimages instead.
func TestContributionCircuitRejectsOversizedConstantTerms(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	rbjj := group.ScalarField()
	// Make key 0's constant term equal to r_bjj exactly (out of range by 1).
	asn.Coefficients[0][0] = new(big.Int).Set(rbjj)

	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// TestContributionCircuitRejectsTorsionCommitments is the regression for the
// cofactor preimages: adding the order-2 point T = (0, p−1) to C_1 and C_2
// leaves every Feldman check intact (x·T + x²·T = x(x+1)·T = O for every
// integer x) but the perturbed commitments are outside the prime subgroup
// and have no preimage under multiplication by 8, so the proof must fail.
func TestContributionCircuitRejectsTorsionCommitments(t *testing.T) {
	c := qt.New(t)
	witness, _, err := BuildWitness(testAssignment())
	c.Assert(err, qt.IsNil)

	torsion := group.NewPoint().SetPoint(big.NewInt(0), new(big.Int).Sub(group.BaseField(), big.NewInt(1)))
	for _, m := range []int{1, 2} {
		x, ok := witness.Commitments[0][m].X.(*big.Int)
		c.Assert(ok, qt.IsTrue)
		y, ok := witness.Commitments[0][m].Y.(*big.Int)
		c.Assert(ok, qt.IsTrue)
		point, err := group.Decode(types.CurvePoint{X: x, Y: y})
		c.Assert(err, qt.IsNil)
		point.Add(point, torsion)
		witness.Commitments[0][m] = ccommon.CircuitPoint(group.Encode(point))
	}
	assert := test.NewAssert(t)
	assert.SolvingFailed(&ContributionCircuit{}, witness, test.WithCurves(ecc.BN254))
}
