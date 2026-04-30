package decryptcombine

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestDecryptCombineRejectsInvalidModeBit (PLAN.md §9.2 partial /
// DEEPSEEK §1.4): the combine circuit constrains `mode·(mode-1) == 0`,
// so any value outside {0, 1} must be rejected. A future refactor that
// drops this constraint could let an attacker supply a fractional/
// blended correction term — this test catches that regression.
func TestDecryptCombineRejectsInvalidModeBit(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	// Try mode = 2 (out of {0, 1}): mode·(mode-1) = 2 ≠ 0, must fail.
	tampered := *witness
	tampered.Mode = big.NewInt(2)
	assert := test.NewAssert(t)
	assert.SolvingFailed(&DecryptCombineCircuit{}, &tampered, test.WithCurves(ecc.BN254))

	// Try mode = -1 (encoded as r-1 in BN254.Fr): also fails.
	tampered = *witness
	tampered.Mode = big.NewInt(-1)
	assert.SolvingFailed(&DecryptCombineCircuit{}, &tampered, test.WithCurves(ecc.BN254))
}

// TestDecryptCombineModeBitSelectsCorrection (PLAN.md §9.2 / DEEPSEEK
// §1.4): the circuit's `api.Select(mode, ΔOrg, S·C1)` must only bind
// the active branch into the verifier equation. With mode=0, swapping
// in a bogus DeltaOrg post-build must NOT break the proof (it's the
// inactive branch); same in reverse for mode=1 with a bogus S.
//
// This is the strongest possible regression check on the Select gate:
// if a future refactor accidentally adds DeltaOrg unconditionally to
// the combine equation, the mode=0 case below would suddenly fail.
func TestDecryptCombineModeBitSelectsCorrection(t *testing.T) {
	c := qt.New(t)

	// Use the existing testAssignment which has Mode=0 (default) and a
	// (0, 1) identity DeltaOrg. Swap in a bogus DeltaOrg post-build —
	// the proof must still succeed because mode=0 selects S·C1.
	bogusPoint := group.NewPoint()
	bogusPoint.ScalarBaseMult(big.NewInt(99))
	bogus := group.Encode(bogusPoint)

	t.Run("mode=0-ignores-DeltaOrg", func(t *testing.T) {
		witness, _, err := BuildWitness(testAssignment())
		c.Assert(err, qt.IsNil)
		// Confirm we're in mode 0.
		c.Assert(witness.Mode.(*big.Int).Sign(), qt.Equals, 0)
		// Tamper DeltaOrg post-build.
		from := bogus // local alias for the closure value
		witness.DeltaOrg.X = from.X
		witness.DeltaOrg.Y = from.Y
		// Re-derive CombineHash because DeltaOrg.X/Y are bound into it.
		// (The circuit's check is "user-supplied CombineHash equals
		// in-circuit re-hash" — that's a transcript-binding check, NOT
		// part of the correction-term selection. To isolate the Select
		// gate behaviour we update CombineHash to match the tampered
		// inputs.)
		witness2, _, err := BuildWitness(modeAssignment(asnWithDeltaOrg(testAssignment(), from, 0, big.NewInt(0))))
		c.Assert(err, qt.IsNil)
		witness.CombineHash = witness2.CombineHash
		assert := test.NewAssert(t)
		// Proof still solves: DeltaOrg is on the inactive (mode==1) branch.
		assert.SolvingSucceeded(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	t.Run("mode=1-ignores-S", func(t *testing.T) {
		// Build an honest mode=1 assignment: T = DeltaOrg, so we need
		// the partial decryption + DeltaOrg to satisfy
		// C2 = m·G + Σλ_k·δ_k + DeltaOrg. The test fixture already has
		// C2 = 3·G + 14·G; reuse it but encode the 14·G as DeltaOrg (the
		// committee partial drops to identity so Σλ_k·δ_k = 0).
		c1Point := group.Generator()
		c1Point.ScalarBaseMult(big.NewInt(5))
		messagePoint := group.Generator()
		messagePoint.ScalarBaseMult(big.NewInt(3))
		deltaOrgPoint := group.Generator()
		deltaOrgPoint.ScalarBaseMult(big.NewInt(14))
		c2Point := group.NewPoint()
		c2Point.Set(messagePoint)
		c2Point.Add(c2Point, deltaOrgPoint)

		identity := types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
		asn := Assignment{
			RoundHash:          big.NewInt(5555),
			Mode:               big.NewInt(1),
			S:                  big.NewInt(0),
			DeltaOrg:           group.Encode(deltaOrgPoint),
			Threshold:          1,
			CiphertextC1:       group.Encode(c1Point),
			CiphertextC2:       group.Encode(c2Point),
			ParticipantIndexes: []uint16{1},
			PartialDecryptions: []types.CurvePoint{identity},
			Plaintext:          big.NewInt(3),
		}
		witness, _, err := BuildWitness(asn)
		c.Assert(err, qt.IsNil)
		c.Assert(witness.Mode.(*big.Int).Sign(), qt.Equals, 1)

		// Tamper S to a non-zero garbage value and re-hash combine.
		witness.S = big.NewInt(0xdead)
		asn2 := asn
		asn2.S = big.NewInt(0xdead)
		witness2, _, err := BuildWitness(asn2)
		c.Assert(err, qt.IsNil)
		witness.CombineHash = witness2.CombineHash

		assert := test.NewAssert(t)
		// Proof still solves: S is on the inactive (mode==0) branch.
		assert.SolvingSucceeded(&DecryptCombineCircuit{}, witness, test.WithCurves(ecc.BN254))
	})
}

// modeAssignment / asnWithDeltaOrg are tiny helpers that let the mode=0
// test rebuild the assignment with a tampered DeltaOrg so we can pull
// the correctly-rebound CombineHash out of a fresh BuildWitness call.
// They exist only to keep the test body readable; the bound-but-not-
// active correction term is the actual subject under test.
func asnWithDeltaOrg(a Assignment, deltaOrg types.CurvePoint, mode uint8, s *big.Int) Assignment {
	a.DeltaOrg = deltaOrg
	a.Mode = new(big.Int).SetUint64(uint64(mode))
	a.S = s
	return a
}

func modeAssignment(a Assignment) Assignment { return a }
