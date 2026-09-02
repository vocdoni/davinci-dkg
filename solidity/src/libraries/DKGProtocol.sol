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
    // ─── Fiat-Shamir / Schnorr / DLEQ transcript domain prefixes ─────────────
    //
    // These are versioned `keccak256` digests of the canonical UTF-8 strings.
    // The cross-impl byte equality is the basis for cross-protocol replay
    // safety: a v1 organizer Schnorr proof cannot be replayed as a v1
    // operator Schnorr proof because they bind a different domain.
    bytes32 internal constant DOMAIN_OPERATOR_REGISTER_V1 =
        keccak256("davinci-dkg:operator-register:v1");
    bytes32 internal constant DOMAIN_ORGANIZER_REGISTER_V1 =
        keccak256("davinci-dkg:organizer-register:v1");
    /// @dev Chaum–Pedersen challenge domain of the organizer's decryption
    ///      share `Δ = sk_org · C_1`. The challenge is keccak (not Poseidon)
    ///      so a browser-only organizer needs nothing but keccak and
    ///      BabyJubJub arithmetic to produce it; `DKGManager.combineDecryption`
    ///      recomputes it from calldata and the combine circuit consumes it
    ///      from the transcript.
    bytes32 internal constant DOMAIN_ORGANIZER_SHARE_V1 =
        keccak256("davinci-dkg:organizer-share:v1");
    /// @dev In-circuit Poseidon domain of the committee's partial-decryption
    ///      DLEQ proofs.
    bytes32 internal constant DOMAIN_DLEQ_V1 =
        keccak256("davinci-dkg:dleq:v1");
}
