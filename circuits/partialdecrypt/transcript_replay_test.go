package partialdecrypt

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
)

// TestPartialDecryptRejectsCrossEpochReplay:
// A valid (δ_i, A_i, B_i, z_i) for one (eid, aid, ctIdx) must NOT satisfy
// the circuit when re-bound to a different (eid', aid', ctIdx'). The
// circuit binds RoundHash, Aid, CtIdx and ParticipantIndex into the
// Fiat-Shamir transcript to prevent cross-round replay.
//
// Strategy: build a valid witness for `(eid_1, aid_1, ctIdx_1)` so
// Response = w + c1·secret. Then mutate any one of the bound fields —
// Response stays fixed but the circuit recomputes c2 ≠ c1, so
// z·G ≠ A + c2·PK and the constraint solver fails.
func TestPartialDecryptRejectsCrossEpochReplay(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	asn.RoundHash = big.NewInt(1111)
	asn.Aid = big.NewInt(0xAA)
	asn.CtIdx = big.NewInt(7)
	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	cases := []struct {
		name   string
		mutate func(w *PartialDecryptCircuit)
	}{
		{"different-eid", func(w *PartialDecryptCircuit) { w.RoundHash = big.NewInt(2222) }},
		{"different-aid", func(w *PartialDecryptCircuit) { w.Aid = big.NewInt(0xBB) }},
		{"different-ctIdx", func(w *PartialDecryptCircuit) { w.CtIdx = big.NewInt(8) }},
		{"different-participant", func(w *PartialDecryptCircuit) { w.ParticipantIndex = big.NewInt(99) }},
	}

	assert := test.NewAssert(t)
	for _, ca := range cases {
		t.Run(ca.name, func(t *testing.T) {
			tampered := *witness
			ca.mutate(&tampered)
			// Mutating a public input only works if the in-circuit
			// challenge derivation is actually sensitive to it. If a
			// future refactor accidentally drops one of these from the
			// transcript, the constraint solver will succeed and this
			// assertion will catch the regression.
			assert.SolvingFailed(&PartialDecryptCircuit{}, &tampered, test.WithCurves(ecc.BN254))
		})
	}
}

// TestPartialDecryptRejectsZeroParticipantIndex: committee slots are
// one-based, so the witness builder must refuse index 0 rather than emit a
// proof the contract's index checks would reject anyway.
func TestPartialDecryptRejectsZeroParticipantIndex(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	asn.ParticipantIndex = 0
	_, _, err := BuildWitness(asn)
	c.Assert(err, qt.Not(qt.IsNil))
}
