// Cross-layer protocol constants shared with the Go node and Solidity
// contracts. Every Fiat-Shamir transcript in the protocol reads its domain
// prefix from this single source.
//
// Any change here MUST be propagated to:
//
//   - solidity/src/libraries/DKGProtocol.sol
//   - internal/protocol/protocol.go
//
// The cross-impl byte equality is covered by tests/vectors/protocol.json
// which is generated from `internal/protocol/protocol.go` and asserted
// byte-for-byte by the SDK and a Foundry test.

import { keccak256, toHex } from 'viem';

// ─── Fiat-Shamir / Schnorr / DLEQ transcript domain prefixes ─────────────────

/** Canonical UTF-8 strings hashed to derive the domain digests below. */
export const DomainOperatorRegisterV1Str = 'davinci-dkg:operator-register:v1';
export const DomainOrganizerRegisterV1Str = 'davinci-dkg:organizer-register:v1';
export const DomainDLEQV1Str = 'davinci-dkg:dleq:v1';
/**
 * Chaum-Pedersen transcript of the organizer's decryption share
 * `Δ = sk_org · C1`. The challenge is a keccak256 (not Poseidon) because the
 * organizer runs in a browser and the contract recomputes the same value to
 * bind the share into the committee's combine proof.
 */
export const DomainOrganizerShareV1Str = 'davinci-dkg:organizer-share:v1';

/**
 * Domain-prefix digests (keccak256 of the strings above). Bound into the
 * Fiat-Shamir transcript on chain and in the in-circuit Poseidon hash.
 *
 * The cross-impl byte equality is the basis for cross-protocol replay
 * safety: a v1 organizer Schnorr proof cannot be replayed as a v1
 * operator Schnorr proof because they bind a different domain.
 */
export const DomainOperatorRegisterV1 = keccak256(toHex(DomainOperatorRegisterV1Str));
export const DomainOrganizerRegisterV1 = keccak256(toHex(DomainOrganizerRegisterV1Str));
export const DomainDLEQV1 = keccak256(toHex(DomainDLEQV1Str));
export const DomainOrganizerShareV1 = keccak256(toHex(DomainOrganizerShareV1Str));
