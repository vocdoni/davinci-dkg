// Package dleq holds the Chaum-Pedersen DLEQ proof bundle and the
// organizer-share prover/verifier.
//
// Two transcripts live in this protocol and they must not be confused:
//
//   - Committee partial decryptions use a Poseidon transcript derived
//     in-circuit. `circuits/partialdecrypt/witness.go` builds the
//     challenge over the full
//     (eid, aid, ctIdx, i, D_i, C_1, δ_i, A_i, B_i) tuple that the
//     in-circuit verifier expects; the `Proof` struct here is the bundle
//     the witness builder hands to the prover.
//
//   - The organizer share (Δ = sk_org·C_1) uses a keccak transcript so a
//     browser-only organizer needs nothing but keccak and BabyJubJub
//     arithmetic. `OrganizerShareChallenge` is the single source of truth
//     for that encoding, shared by the prover, the verifier, the
//     decrypt-combine witness builder and the cross-impl vectors.
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
