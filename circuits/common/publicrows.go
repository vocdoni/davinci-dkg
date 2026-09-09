package common

import (
	"fmt"

	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
)

// CertifyPublicInputs adds one multiplication row per public input in which
// that input is the only variable on the left-hand side of the constraint.
//
// The weak simulation-extractability theorem for Groth16 (Baghery, Kohlweiss,
// Siim, Volkhov, FC 2021, Theorem 1) requires the left QAP polynomials of the
// public wires to be linearly independent and span-disjoint from those of the
// witness wires. gnark encodes `AssertIsEqual(public, computed)` as
// 1 · public = computed, which puts the public wire on the right-hand side
// and leaves its left polynomial zero, so a circuit whose public inputs are
// only ever asserted equal to derived values violates the hypothesis. A row
// with the public wire alone on the left certifies it: evaluating any linear
// combination of left polynomials at these rows isolates its public
// coefficients. The rows cost one constraint per public input and change
// nothing else about the relation.
func CertifyPublicInputs(api frontend.API, inputs ...frontend.Variable) {
	for _, in := range inputs {
		api.Mul(in, in)
	}
}

// MissingDedicatedPublicRows returns the public wires of a compiled R1CS,
// wire 0 being the constant one, that have no constraint in which they are
// the only left-hand variable; the certificate of CertifyPublicInputs holds
// iff the result is empty.
func MissingDedicatedPublicRows(ccs constraint.ConstraintSystem) ([]int, error) {
	cs, ok := ccs.(interface {
		GetR1Cs() []constraint.R1C
		GetNbPublicVariables() int
	})
	if !ok {
		return nil, fmt.Errorf("constraint system %T is not an R1CS", ccs)
	}
	nbPublic := cs.GetNbPublicVariables()
	dedicated := make([]bool, nbPublic)
	for _, r1c := range cs.GetR1Cs() {
		if len(r1c.L) != 1 {
			continue
		}
		term := r1c.L[0]
		if wire := term.WireID(); wire < nbPublic && term.CoeffID() != constraint.CoeffIdZero {
			dedicated[wire] = true
		}
	}
	var missing []int
	for wire, ok := range dedicated {
		if !ok {
			missing = append(missing, wire)
		}
	}
	return missing, nil
}
