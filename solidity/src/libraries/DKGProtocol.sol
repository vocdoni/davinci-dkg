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
/// @dev    Two modes everywhere, shared role tags, and per-transcript
///         domain prefixes for cross-protocol replay safety.
library DKGProtocol {
    // ─── Application registration modes ──────────────────────────────────────
    //
    // Selects how the per-application correction term `T` is produced and
    // consumed by the combine circuit. The flag is stored on each
    // Application record at registration time and supplied to
    // `DecryptCombineVerifier` as a public input.
    //
    // mode = 0 (public derivation, paper §4.3):
    //          PK_aid = PK_ep + S·G with S = Hash(eid || PK_ep || aid).
    //          The contract stores S and the combine circuit computes
    //          T = S·C_1 in-circuit (paper line 1088).
    //
    // mode = 1 (organizer co-decryption, paper §6):
    //          PK_aid = PK_ep + PK_org. Decryption requires both the
    //          committee threshold partial decryptions and the organizer's
    //          Δ_org = sk_org · C_1 with a Chaum-Pedersen DLEQ. The
    //          combine circuit consumes T = Δ_org as a public-input
    //          curve point.
    uint8 internal constant MODE_PUBLIC_DERIVATION = 0;
    uint8 internal constant MODE_ORGANIZER_CODEC = 1;

    // ─── DLEQ role tags ─────────────────────────────────────────────────────
    //
    // Public-input tag distinguishing committee and organizer Chaum-Pedersen
    // proofs (paper §4.4 lines 695–704). Bound into the Fiat-Shamir transcript
    // so a committee proof cannot be replayed as an organizer share or vice
    // versa.
    uint8 internal constant ROLE_COMMITTEE = 1;
    uint8 internal constant ROLE_ORGANIZER = 2;

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
    bytes32 internal constant DOMAIN_DLEQ_V1 =
        keccak256("davinci-dkg:dleq:v1");
}
