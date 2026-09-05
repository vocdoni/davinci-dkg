// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/// @title DKGProtocol
/// @notice Cross-layer protocol constants. Solidity, Go and TypeScript all
///         hold copies of the same values; the cross-impl byte-equality is
///         covered by `tests/vectors/protocol.json` (driven from the Go side
///         and asserted by the SDK + Solidity).
///
///         Any change here MUST be propagated to:
///
///           - `internal/protocol/protocol.go`        (Go node)
///           - `sdk/src/protocol.ts`                  (TypeScript SDK)
///
/// @dev    Per-transcript domain prefixes for cross-protocol replay safety.
library DKGProtocol {
    // ─── Schnorr registration transcript domain prefixes ─────────────────────
    //
    // These are versioned `keccak256` digests of the canonical UTF-8 strings.
    // The cross-impl byte equality is the basis for cross-protocol replay
    // safety: a v1 organizer Schnorr proof cannot be replayed as a v1
    // operator Schnorr proof because they bind a different domain.
    bytes32 internal constant DOMAIN_OPERATOR_REGISTER_V1 =
        keccak256("davinci-dkg:operator-register:v1");
    bytes32 internal constant DOMAIN_ORGANIZER_REGISTER_V1 =
        keccak256("davinci-dkg:organizer-register:v1");

    // ─── BRLC transcript domains ─────────────────────────────────────────────
    //
    // The Fiat-Shamir domain every proof-carrying call binds into its
    // challenge: `keccak(eid || domain || anchor) mod p` (see BRLC.sol). One
    // per circuit whose transcript the contract streams.
    bytes32 internal constant DOMAIN_CONTRIBUTION_TRANSCRIPT_V1 =
        keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DOMAIN_POOLKEY_TRANSCRIPT_V1 =
        keccak256("davinci-dkg:poolkey:v1");
    bytes32 internal constant DOMAIN_DECRYPT_COMBINE_TRANSCRIPT_V1 =
        keccak256("davinci-dkg:decrypt-combine:v1");
}
