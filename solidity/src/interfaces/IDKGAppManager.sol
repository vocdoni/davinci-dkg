// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "../libraries/DKGTypes.sol";

interface IDKGAppManager {
    // ─── Events ────────────────────────────────────────────────────────────────
    /// @notice An application was registered and claimed pool key `poolIndex`
    ///         of the epoch. `organizerPKx/y` is the identity `(0, 1)` in
    ///         `Automatic` mode.
    event ApplicationRegistered(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        address indexed creator,
        uint256 organizerPKx,
        uint256 organizerPKy,
        DKGTypes.AppMode mode,
        uint8 poolIndex
    );
    /// @notice The organizer published `sk_org` for an `OrganizerLocked`
    ///         application. From here on the committee combines by itself:
    ///         every ciphertext of the application becomes decryptable by the
    ///         threshold alone. Emitted at most once per application.
    event OrganizerSecretRevealed(bytes12 indexed epochId, bytes32 indexed aid, uint256 organizerSecret);

    // ─── Errors ────────────────────────────────────────────────────────────────
    error InvalidApplication();
    error ApplicationAlreadyExists();
    error InvalidSchnorrProof();
    /// @dev The locked-mode organizer key is not in the curve's prime-order
    ///      subgroup. The Schnorr PoP does not imply membership — a
    ///      small-order key passes it with a ground challenge — and such an
    ///      application could never be revealed or combined.
    error PointNotInSubgroup();
    error InvalidEpoch();
    error InvalidPhase();
    error InvalidAddress();
    error NotOwner();
    error DecryptionNotYetAllowed();
    error DecryptionExpired();
    error DecryptionLimitReached();
    /// @dev `block.timestamp` is past the application's `decryptNotAfter`.
    error DecryptionClosed();
    /// @dev `block.timestamp` has not reached the application's
    ///      `decryptNotBefore` yet.
    error DecryptionNotOpen();
    error InvalidOrganizerSecret();
    error InvalidPolicy();
    /// @dev The application is `OrganizerLocked` and `revealOrganizerSecret`
    ///      has not run yet: no partial decryption and no combine may exist
    ///      before the reveal. Ciphertext submission is unaffected.
    error OrganizerSecretNotRevealed();
    /// @dev `revealOrganizerSecret` was called twice, or on an application
    ///      whose secret is structurally absent (`Automatic`).
    error AlreadyRevealed();
    /// @dev Re-declared from `IDKGManager` (identical selectors): the
    ///      manager raises them from `claimPoolKey`, which registration
    ///      calls, so they surface out of `registerApplication`.
    error PoolExhausted();
    error PoolKeyNotActive();

    // ─── Application lifecycle ────────────────────────────────────────────────

    /// @notice Register an application against a Live epoch and claim the
    ///         epoch's next activated pool key `P_j`.
    ///
    ///         In `OrganizerLocked` mode the caller proves knowledge of
    ///         `sk_org` with a Schnorr PoP over `DOMAIN_ORGANIZER_REGISTER_V1`
    ///         and the application key is `PK_aid = P_j + PK_org`.
    ///
    ///         In `Automatic` mode the key and Schnorr arguments are ignored,
    ///         `organizerPK` is stored as the identity `(0, 1)` and
    ///         `PK_aid = P_j`.
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

    /// @notice Publish `sk_org` for an `OrganizerLocked` application, once.
    ///         Permissionless — the contract checks `sk·G == PK_org`, so only
    ///         the real secret is accepted, and whoever holds it may publish
    ///         it. There is no way back: the committee can decrypt every
    ///         ciphertext of the application from then on.
    function revealOrganizerSecret(bytes12 epochId, bytes32 aid, uint256 organizerSecret) external;

    function getApplication(bytes12 epochId, bytes32 aid)
        external
        view
        returns (DKGTypes.Application memory);

    // ─── Cross-contract APIs (called by DKGManager) ───────────────────────────

    /// @notice `PK_org` of a registered application — the identity `(0, 1)`
    ///         for `Automatic` ones, zero for an unknown aid.
    function getOrganizerPK(bytes12 epochId, bytes32 aid) external view returns (uint256, uint256);

    /// @notice Reverts with `InvalidApplication` for an unknown aid, with
    ///         `OrganizerSecretNotRevealed` for an `OrganizerLocked`
    ///         application whose secret is still sealed, with
    ///         `DecryptionNotOpen` before the application's `decryptNotBefore`
    ///         and with `DecryptionClosed` after its `decryptNotAfter`
    ///         (neither when the bound is 0).
    function requireDecryptionOpen(bytes12 epochId, bytes32 aid) external view;

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
