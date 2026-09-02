// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "../libraries/DKGTypes.sol";

interface IDKGAppManager {
    // ─── Events ────────────────────────────────────────────────────────────────
    event ApplicationRegistered(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        address indexed creator,
        uint256 organizerPKx,
        uint256 organizerPKy
    );
    /// @notice The organizer's decryption share `Δ = sk_org·C_1` for one
    ///         ciphertext, together with its Chaum–Pedersen DLEQ `(A1, A2, z)`.
    ///         The contract stores only `keccak256(Δ ‖ A1 ‖ A2 ‖ z)`; the words
    ///         themselves are only available here, and the DLEQ is verified
    ///         inside the combine SNARK (and off chain by the committee).
    event OrganizerShareSubmitted(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        uint16 indexed ciphertextIndex,
        uint256 deltaX,
        uint256 deltaY,
        uint256 a1x,
        uint256 a1y,
        uint256 a2x,
        uint256 a2y,
        uint256 z
    );

    // ─── Errors ────────────────────────────────────────────────────────────────
    error InvalidApplication();
    error ApplicationAlreadyExists();
    error InvalidSchnorrProof();
    error InvalidEpoch();
    error InvalidPhase();
    error InvalidAddress();
    error InvalidCiphertext();
    error CiphertextNotSubmitted();
    error AlreadyCombined();
    error InvalidProofInput();
    error NotOwner();
    error DecryptionNotYetAllowed();
    error DecryptionExpired();
    error DecryptionLimitReached();

    // ─── Application lifecycle ────────────────────────────────────────────────

    /// @notice Register an application against a Live epoch. The caller proves
    ///         knowledge of `sk_org` with a Schnorr PoP over
    ///         `DOMAIN_ORGANIZER_REGISTER_V1`; the application key is
    ///         `PK_aid = PK_ep + PK_org`.
    function registerApplication(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy,
        uint256 pkOrgX,
        uint256 pkOrgY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external;

    /// @notice Publish the organizer's decryption share for one ciphertext.
    ///         Permissionless: the share is self-authenticating because the
    ///         combine SNARK verifies its DLEQ against the registered
    ///         `PK_org`. Overwrites any previous share until the ciphertext
    ///         has been combined.
    function submitOrganizerShare(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y,
        uint256 deltaX,
        uint256 deltaY,
        uint256 a1x,
        uint256 a1y,
        uint256 a2x,
        uint256 a2y,
        uint256 z
    ) external;

    function getApplication(bytes12 epochId, bytes32 aid)
        external
        view
        returns (DKGTypes.Application memory);

    // ─── Cross-contract APIs (called by DKGManager) ───────────────────────────

    /// @notice `keccak256(abi.encodePacked(deltaX, deltaY, a1x, a1y, a2x, a2y, z))`
    ///         of the organizer share stored for `(epochId, aid, ciphertextIndex)`,
    ///         or `bytes32(0)` if none has been submitted.
    function getOrganizerShareHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (bytes32);

    /// @notice Enforce the per-app submitCiphertext access policy. Reverts on
    ///         policy failure, and on an aid that was never registered.
    function requireCanSubmitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        address sender
    ) external view;

    /// @notice Returns the list of registered aids for an epoch.
    function getRegisteredAids(bytes12 epochId) external view returns (bytes32[] memory);
}
