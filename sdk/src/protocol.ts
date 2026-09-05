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

// ─── Schnorr registration and BRLC transcript domain prefixes ────────────────

/** Canonical UTF-8 strings hashed to derive the domain digests below. */
export const DomainOperatorRegisterV1Str = 'davinci-dkg:operator-register:v1';
export const DomainOrganizerRegisterV1Str = 'davinci-dkg:organizer-register:v1';
/**
 * BRLC transcript domains: bound into every proof-carrying call's challenge
 * (`keccak(eid ‖ domain ‖ anchor) mod p`, see `deriveChallenge`). One per
 * circuit whose transcript the contract streams: the compact contribution
 * (v2, docs/pool-keys-v4.md §3), the batched finalization (v2, §7 — it
 * replaces the former per-key `davinci-dkg:poolkey:v1` of `activatePoolKey`)
 * and decrypt-combine (unchanged).
 */
export const DomainContributionTranscriptV2Str = 'davinci-dkg:contribution:v2';
export const DomainFinalizeTranscriptV2Str = 'davinci-dkg:finalize:v2';
export const DomainDecryptCombineTranscriptV1Str = 'davinci-dkg:decrypt-combine:v1';

/**
 * Domain-prefix digests (keccak256 of the strings above). Bound into the
 * Fiat-Shamir transcript on chain and in the in-circuit Poseidon hash.
 *
 * The cross-impl byte equality is the basis for cross-protocol replay
 * safety: a v1 organizer Schnorr proof cannot be replayed as a v1
 * operator Schnorr proof because they bind a different domain, and a v4
 * contribution transcript cannot be replayed as a finalization one.
 */
export const DomainOperatorRegisterV1 = keccak256(toHex(DomainOperatorRegisterV1Str));
export const DomainOrganizerRegisterV1 = keccak256(toHex(DomainOrganizerRegisterV1Str));
export const DomainContributionTranscriptV2 = keccak256(toHex(DomainContributionTranscriptV2Str));
export const DomainFinalizeTranscriptV2 = keccak256(toHex(DomainFinalizeTranscriptV2Str));
export const DomainDecryptCombineTranscriptV1 = keccak256(toHex(DomainDecryptCombineTranscriptV1Str));
