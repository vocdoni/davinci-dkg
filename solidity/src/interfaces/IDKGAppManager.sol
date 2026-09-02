// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "../libraries/DKGTypes.sol";

interface IDKGAppManager {
    // ─── Events ────────────────────────────────────────────────────────────────
    event ApplicationRegistered(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        address indexed creator,
        uint8 mode,
        uint256 derivationS,
        uint256 organizerPKx,
        uint256 organizerPKy
    );
    event OrganizerShareSubmitted(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        uint16 indexed ciphertextIndex,
        uint256 deltaOrgX,
        uint256 deltaOrgY
    );

    // ─── Errors ────────────────────────────────────────────────────────────────
    error InvalidApplication();
    error ApplicationAlreadyExists();
    error InvalidSchnorrProof();
    error InvalidEpoch();
    error InvalidPhase();
    error InvalidAddress();
    error InvalidVerifier();
    error InvalidCiphertext();
    error CiphertextNotSubmitted();
    error AlreadyPartiallyDecrypted();
    error InvalidProofInput();
    error NotOwner();
    error DecryptionNotYetAllowed();
    error DecryptionExpired();
    error DecryptionLimitReached();
    error InsufficientPartialDecryptions();

    // ─── Application lifecycle ────────────────────────────────────────────────
    function registerApplication(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy
    ) external;

    function registerApplicationCoDec(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy,
        uint256 pkOrgX,
        uint256 pkOrgY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external;

    function submitOrganizerShare(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y,
        uint256 deltaOrgX,
        uint256 deltaOrgY,
        bytes calldata dleqProof,
        bytes calldata dleqInput
    ) external;

    function getApplication(bytes12 epochId, bytes32 aid)
        external
        view
        returns (DKGTypes.Application memory);

    // ─── Cross-contract APIs (called by DKGManager) ───────────────────────────

    /// @notice Returns the per-app correction parameters DKGManager.combineDecryption
    ///         needs. Reverts if `aid != 0` and the app does not exist.
    ///         For aid == 0 (legacy per-epoch path) returns (mode=0, S=0, identity).
    function getCombineCorrection(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (uint8 mode, uint256 derivationS, uint256 deltaOrgX, uint256 deltaOrgY);

    /// @notice Enforce the per-app submitCiphertext access policy. Reverts on policy fail.
    ///         For aid == 0 the call is a no-op; the caller (DKGManager.submitCiphertext)
    ///         treats aid == 0 (the epoch key itself) as an open application.
    function requireCanSubmitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        address sender
    ) external view;

    /// @notice Returns the list of registered aids for an epoch (excluding bytes32(0)).
    function getRegisteredAids(bytes12 epochId) external view returns (bytes32[] memory);
}
