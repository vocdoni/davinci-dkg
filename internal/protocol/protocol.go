// Package protocol defines cross-layer constants shared between the Go
// node, the Solidity contracts, and the TypeScript SDK. These constants
// are the protocol's lingua franca: every code path that produces or
// consumes a mode flag, a role tag, or a Fiat-Shamir transcript reads
// them from this single source.
//
// Any change here MUST be propagated to:
//
//   - solidity/src/libraries/DKGProtocol.sol
//   - sdk/src/protocol.ts
//
// The cross-impl byte equality is covered by tests/vectors/protocol.json
// (PR-S1) which is generated from this file and asserted byte-for-byte
// by the SDK and a Foundry test.
//
// Implements PLAN.md §2 principles 2 and 3 ("two modes everywhere",
// shared role tags) and §5 opener (transcript domain prefixes).
package protocol

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// AppMode selects how the per-application correction term T is produced
// and consumed by the combine circuit. See solidity/src/libraries/DKGProtocol.sol
// for the full explanation.
type AppMode uint8

const (
	// ModePublicDerivation: PK_aid = PK_ep + S·G with
	// S = Hash(eid||PK_ep||aid). Combine circuit computes T = S·C_1
	// in-circuit (paper line 1088).
	ModePublicDerivation AppMode = 0
	// ModeOrganizerCoDec: PK_aid = PK_ep + PK_org. Decryption requires both
	// the committee threshold and the organizer's Δ_org = sk_org · C_1
	// with a Chaum-Pedersen DLEQ. Combine circuit consumes T = Δ_org
	// as a public-input curve point.
	ModeOrganizerCoDec AppMode = 1
)

// Role tags Chaum-Pedersen proofs as either a committee partial decryption
// or an organizer share. Bound into the Fiat-Shamir transcript (paper
// §4.4 lines 695–704) so the two cannot be replayed for one another.
type Role uint8

const (
	RoleCommittee Role = 1
	RoleOrganizer Role = 2
)

// Domain-prefix strings. These are the canonical UTF-8 inputs that hash
// to the transcript domains below. Exposed so test vectors can include
// the pre-image alongside its digest.
const (
	DomainOperatorRegisterV1Str  = "davinci-dkg:operator-register:v1"
	DomainOrganizerRegisterV1Str = "davinci-dkg:organizer-register:v1"
	DomainDLEQV1Str              = "davinci-dkg:dleq:v1"
	// DomainDerivationV1Str is intentionally absent — the per-application
	// `S = keccak256(eid || PK_ep || aid) mod q` derivation has NO domain
	// prefix (paper §4.3 / DEEPSEEK §2.3). A previous draft contemplated
	// one; removing the dead constant prevents readers from assuming the
	// transcript carries it.
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
)

// Hash exposes the canonical hash function used to derive the domain
// digests above; provided so callers don't have to import go-ethereum
// crypto directly when they want to derive their own keccak256 digests
// over protocol-shaped data.
func Hash(b []byte) common.Hash {
	return crypto.Keccak256Hash(b)
}
