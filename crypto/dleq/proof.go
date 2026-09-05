// Package dleq holds the Chaum-Pedersen DLEQ proof bundle used by the
// committee's partial decryptions.
//
// The transcript is Poseidon and is derived in-circuit:
// `circuits/partialdecrypt/witness.go` builds the challenge over the full
// (eid, aid, ctIdx, i, D_i, C_1, δ_i, A_i, B_i) tuple that the in-circuit
// verifier expects; the `Proof` struct here is the bundle the witness
// builder hands to the prover.
//
// There is no organizer-share DLEQ any more: the combine circuit proves
// knowledge of sk_org directly (see docs/pool-keys.md, "Combine").
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
