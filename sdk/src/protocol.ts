// Cross-layer protocol constants shared with the Go node and Solidity
// contracts. Every code path that produces or consumes a mode flag, a
// role tag, or a Fiat-Shamir transcript reads them from this single
// source.
//
// Any change here MUST be propagated to:
//
//   - solidity/src/libraries/DKGProtocol.sol
//   - internal/protocol/protocol.go
//
// The cross-impl byte equality is covered by tests/vectors/protocol.json
// which is generated from `internal/protocol/protocol.go` and asserted
// byte-for-byte by the SDK and a Foundry test.
//
// Two modes everywhere, shared role tags, and per-transcript domain
// prefixes for cross-protocol replay safety.

import { keccak256, toHex } from 'viem';

// ─── Application registration modes ──────────────────────────────────────────

/**
 * Selects how the per-application correction term T is produced and
 * consumed by the combine circuit. See `solidity/src/libraries/DKGProtocol.sol`
 * for the full explanation.
 */
export const AppMode = {
  /**
   * Public derivation: PK_aid = PK_ep + S·G with S = Hash(eid||PK_ep||aid).
   * Combine circuit computes T = S·C_1 in-circuit (paper line 1088).
   */
  PublicDerivation: 0,
  /**
   * Organizer co-decryption: PK_aid = PK_ep + PK_org. Decryption requires
   * both the committee threshold and the organizer's Δ_org = sk_org · C_1
   * with a Chaum-Pedersen DLEQ. Combine circuit consumes T = Δ_org as a
   * public-input curve point.
   */
  OrganizerCoDec: 1,
} as const;

export type AppModeValue = (typeof AppMode)[keyof typeof AppMode];

// ─── DLEQ role tags ──────────────────────────────────────────────────────────

/**
 * Tags a Chaum-Pedersen proof as either a committee partial decryption
 * or an organizer share. Bound into the Fiat-Shamir transcript (paper
 * §4.4 lines 695–704) so the two cannot be replayed for one another.
 */
export const Role = {
  Committee: 1,
  Organizer: 2,
} as const;

export type RoleValue = (typeof Role)[keyof typeof Role];

// ─── Fiat-Shamir / Schnorr / DLEQ transcript domain prefixes ─────────────────

/** Canonical UTF-8 strings hashed to derive the domain digests below.
 * Note: the per-application `S = keccak256(eid || PK_ep || aid) mod q`
 * derivation deliberately has NO domain prefix in the on-chain or paper
 * definitions. */
export const DomainOperatorRegisterV1Str = 'davinci-dkg:operator-register:v1';
export const DomainOrganizerRegisterV1Str = 'davinci-dkg:organizer-register:v1';
export const DomainDLEQV1Str = 'davinci-dkg:dleq:v1';
/**
 * Schnorr proof of knowledge of a ciphertext's randomness r (C1 = r·G),
 * checked off-chain by every committee node before it releases a partial
 * decryption; stops a ciphertext's C1 from being re-submitted under another
 * application as a decryption oracle.
 */
export const DomainCiphertextPoKV1Str = 'davinci-dkg:ciphertext-pok:v1';

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
export const DomainCiphertextPoKV1 = keccak256(toHex(DomainCiphertextPoKV1Str));
