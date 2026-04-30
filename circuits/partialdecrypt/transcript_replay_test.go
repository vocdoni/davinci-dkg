package partialdecrypt

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// Mirror of the constants in internal/protocol/protocol.go. Kept inline
// here so the circuit test doesn't pull in the cross-layer protocol
// package (which carries an extra dep on ethcrypto).
const (
	roleCommittee = 1
	roleOrganizer = 2
)

// TestPartialDecryptRejectsCrossEpochReplay (PLAN.md §9.2 / DEEPSEEK §1.3):
// A valid (δ_i, A_i, B_i, z_i) for one (eid, aid, ctIdx) must NOT satisfy
// the circuit when re-bound to a different (eid', aid', ctIdx'). This is
// the load-bearing regression test for the H-1 vulnerability (the
// pre-rewrite circuit didn't bind RoundHash into the Fiat-Shamir
// transcript, allowing cross-round replay).
//
// Strategy: build a valid witness for `(eid_1, aid_1, ctIdx_1)` so
// Response = w + c1·secret. Then mutate any one of the bound fields
// (eid, aid, ctIdx, role, participantIndex) — Response stays fixed but
// the circuit recomputes c2 ≠ c1, so z·G ≠ A + c2·PK and the constraint
// solver fails.
func TestPartialDecryptRejectsCrossEpochReplay(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	asn.RoundHash = big.NewInt(1111)
	asn.Aid = big.NewInt(0xAA)
	asn.CtIdx = big.NewInt(7)
	asn.Role = big.NewInt(roleCommittee)
	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	cases := []struct {
		name  string
		mutate func(w *PartialDecryptCircuit)
	}{
		{"different-eid", func(w *PartialDecryptCircuit) { w.RoundHash = big.NewInt(2222) }},
		{"different-aid", func(w *PartialDecryptCircuit) { w.Aid = big.NewInt(0xBB) }},
		{"different-ctIdx", func(w *PartialDecryptCircuit) { w.CtIdx = big.NewInt(8) }},
		{"different-participant", func(w *PartialDecryptCircuit) { w.ParticipantIndex = big.NewInt(99) }},
		{"committee→organizer", func(w *PartialDecryptCircuit) { w.Role = big.NewInt(roleOrganizer) }},
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

// TestPartialDecryptOrganizerRoleConsistency (PLAN.md §9.2 / DEEPSEEK §1.5):
// the organizer DLEQ shares a circuit with the committee DLEQ; the role
// tag is the only thing distinguishing them in-circuit. A witness built
// for role=ORGANIZER must verify with role=ORGANIZER (sanity check) AND
// must be rejected if role is silently flipped to COMMITTEE post-build —
// the contract relies on this to enforce that organizer shares can only
// be submitted via submitOrganizerShare with role=ORGANIZER.
func TestPartialDecryptOrganizerRoleConsistency(t *testing.T) {
	c := qt.New(t)

	asn := testAssignment()
	asn.Role = big.NewInt(roleOrganizer)
	asn.ParticipantIndex = 1 // organizer slot — circuit doesn't care about value, only that it's bound
	witness, _, err := BuildWitness(asn)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)

	// Sanity: the honest organizer witness solves.
	t.Run("organizer-role-honest", func(t *testing.T) {
		assert.SolvingSucceeded(&PartialDecryptCircuit{}, witness, test.WithCurves(ecc.BN254))
	})

	// Replay: same proof material relabeled as committee → fails.
	t.Run("organizer→committee", func(t *testing.T) {
		tampered := *witness
		tampered.Role = big.NewInt(roleCommittee)
		assert.SolvingFailed(&PartialDecryptCircuit{}, &tampered, test.WithCurves(ecc.BN254))
	})
}

// touch group import in case future test additions need it.
var _ = group.Generator
