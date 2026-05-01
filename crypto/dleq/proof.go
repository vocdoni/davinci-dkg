// Package dleq holds the Chaum-Pedersen DLEQ proof bundle.
//
// The active partial-decrypt path in `circuits/partialdecrypt/witness.go`
// derives the Fiat-Shamir challenge directly over the full
//
//	(eid, aid, ctIdx, role, i, G, C_1, D_i, δ_i, A_i, B_i)
//
// transcript that the in-circuit verifier expects. The `Proof` struct
// here is the bundle the witness builder hands to the prover.
package dleq

import (
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Proof is a non-interactive Chaum-Pedersen proof of equal discrete
// logarithms. The fields are:
//
//	A1, A2   — the prover's commitments (w·G and w·base).
//	Response — the prover's z = w + c·secret (mod L).
//
// The Fiat-Shamir challenge `c` is implicit; both prover and verifier
// recompute it over the agreed transcript.
type Proof struct {
	A1       types.CurvePoint
	A2       types.CurvePoint
	Response *big.Int
}
