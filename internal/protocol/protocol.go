// Package protocol defines cross-layer constants shared between the Go
// node, the Solidity contracts, and the TypeScript SDK. These constants
// are the protocol's lingua franca: every code path that produces or
// consumes a Fiat-Shamir transcript reads its domain prefix from this
// single source.
//
// Any change here MUST be propagated to:
//
//   - solidity/src/libraries/DKGProtocol.sol
//   - sdk/src/protocol.ts
//
// The cross-impl byte equality is covered by tests/vectors/protocol.json
// which is generated from this file and asserted byte-for-byte
// by the SDK and a Foundry test.
package protocol

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Domain-prefix strings. These are the canonical UTF-8 inputs that hash
// to the transcript domains below. Exposed so test vectors can include
// the pre-image alongside its digest.
const (
	DomainOperatorRegisterV1Str  = "davinci-dkg:operator-register:v1"
	DomainOrganizerRegisterV1Str = "davinci-dkg:organizer-register:v1"
	DomainDLEQV1Str              = "davinci-dkg:dleq:v1"
	// DomainOrganizerShareV1Str separates the organizer's Chaum-Pedersen
	// DLEQ over Δ = sk_org·C1. The challenge is keccak (not Poseidon) so a
	// browser-only organizer needs nothing but keccak and BabyJubJub
	// arithmetic; the combine SNARK consumes the challenge as a transcript
	// word and the contract recomputes it from calldata.
	DomainOrganizerShareV1Str = "davinci-dkg:organizer-share:v1"
)

// Domain-prefix digests (keccak256 of the strings above). These are the
// values bound into the Fiat-Shamir transcript on the chain and in the
// in-circuit Poseidon hash.
//
// The cross-impl byte equality is the basis for cross-protocol replay
// safety: a v1 organizer Schnorr proof cannot be replayed as a v1
// operator Schnorr proof because they bind a different domain.
var (
	DomainOperatorRegisterV1  = crypto.Keccak256Hash([]byte(DomainOperatorRegisterV1Str))
	DomainOrganizerRegisterV1 = crypto.Keccak256Hash([]byte(DomainOrganizerRegisterV1Str))
	DomainDLEQV1              = crypto.Keccak256Hash([]byte(DomainDLEQV1Str))
	DomainOrganizerShareV1    = crypto.Keccak256Hash([]byte(DomainOrganizerShareV1Str))
)

// Hash exposes the canonical hash function used to derive the domain
// digests above; provided so callers don't have to import go-ethereum
// crypto directly when they want to derive their own keccak256 digests
// over protocol-shaped data.
func Hash(b []byte) common.Hash {
	return crypto.Keccak256Hash(b)
}
