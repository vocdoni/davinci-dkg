// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGAppManager} from "./interfaces/IDKGAppManager.sol";
import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";

/// @title  DKGAppManager
/// @notice Per-application surface (organizer registration, pool-key claim,
///         submission policy and decryption window) for the davinci-dkg
///         system. Sibling to `DKGManager`; drives it through the narrow
///         `IDKGManager` surface (`getEpoch`, `claimPoolKey`) and is consulted
///         by it via `IDKGAppManager` at `submitCiphertext` /
///         `submitPartialDecryption` / `combineDecryption` time.
///
///         Was carved out of `DKGManager` to keep the latter's runtime size
///         under the EIP-170 24,576-byte limit. No protocol behaviour change.
contract DKGAppManager is IDKGAppManager {
    /// @dev Bounds the allow-list scan in `requireCanSubmitCiphertext`.
    uint256 internal constant MAX_SUBMITTERS = 32;

    /// @notice The sibling DKGManager whose epochs this contract registers
    ///         applications against. Immutable.
    address public immutable MANAGER;

    /// @dev Per-application registrations, keyed by `(eid, aid)`. See
    ///      DKGTypes.Application for the record shape.
    mapping(bytes12 epochId => mapping(bytes32 aid => DKGTypes.Application app)) internal applications;

    /// @dev List of registered aids per epoch, exposed via `getRegisteredAids`
    ///      for explorers. Append-only; never reordered.
    mapping(bytes12 epochId => bytes32[] aids) internal epochAidsList;

    constructor(address _manager) {
        if (_manager == address(0)) revert InvalidAddress();
        MANAGER = _manager;
    }

    // ─── Application lifecycle ───────────────────────────────────────────────

    /// @notice Register an application against a Live epoch and claim the
    ///         epoch's next activated pool key `P_j`.
    ///
    ///         `OrganizerLocked`: verifies a Schnorr proof of knowledge of
    ///         `sk_org`,
    ///
    ///             c = keccak256(domain || eid || aid || PK_org || A) mod L
    ///             z·G == A + c·PK_org
    ///
    ///         and the application key is `PK_aid = P_j + PK_org`. Losing
    ///         `sk_org` leaves the ciphertexts sealed until (and unless) the
    ///         organizer calls `revealOrganizerSecret`.
    ///
    ///         `Automatic`: there is no organizer key. The `pkOrg*` and
    ///         `schnorr*` arguments are ignored, `organizerPK` is stored as
    ///         the identity `(0, 1)` and `PK_aid = P_j` — the committee
    ///         decrypts on its own, and confidentiality rests on the threshold
    ///         plus the fact that `P_j` is this application's own key.
    /// @dev    Submission is open to anyone (`policy.openSubmission`), to the
    ///         `policy.submitters` allow-list, or — when both are empty — to
    ///         `msg.sender` only. The pool claim is the last step so a
    ///         rejected registration never burns a key.
    function registerApplication(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy,
        uint256 pkOrgX,
        uint256 pkOrgY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external {
        IDKGManager.Epoch memory epoch = IDKGManager(MANAGER).getEpoch(epochId);
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        _requireValidAid(aid);
        DKGTypes.Application storage app = applications[epochId][aid];
        if (app.exists) revert ApplicationAlreadyExists();
        _requireValidPolicy(policy);

        uint256 keyX;
        uint256 keyY;
        if (policy.mode == DKGTypes.AppMode.Automatic) {
            // No organizer half at all: store the identity so the combine
            // transcript's PK_org binding needs no mode-specific branch.
            keyY = 1;
        } else {
            // The subgroup check is load-bearing, not hygiene: a small-order
            // organizer key satisfies the Schnorr PoP with a ground
            // challenge, but `sk·G` is then never a multiple of `PK_org`,
            // so the application can neither be revealed nor combined, and
            // its "locked" ciphertexts are decryptable by the committee
            // alone. The Schnorr nonce A needs no such check.
            BabyJubJub.requireValidPoint(pkOrgX, pkOrgY);
            if (!BabyJubJub.isInPrimeSubgroup(pkOrgX, pkOrgY)) revert PointNotInSubgroup();
            BabyJubJub.requireValidPoint(schnorrAx, schnorrAy);
            if (
                !_verifyOrganizerSchnorr(epochId, aid, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
            ) revert InvalidSchnorrProof();
            keyX = pkOrgX;
            keyY = pkOrgY;
        }

        // Reverts PoolExhausted / PoolKeyNotActive; only the app manager may
        // call it, and it moves the epoch's pool cursor forward by one.
        uint8 poolIndex = IDKGManager(MANAGER).claimPoolKey(epochId, aid);

        app.creator = msg.sender;
        app.organizerPK = DKGTypes.Point({x: keyX, y: keyY});
        app.poolIndex = poolIndex;
        // Field by field: a calldata struct holding a dynamic array cannot be
        // assigned to storage in one go.
        DKGTypes.AppPolicy storage stored = app.policy;
        stored.mode = policy.mode;
        stored.openSubmission = policy.openSubmission;
        stored.submitters = policy.submitters;
        stored.maxCiphertexts = policy.maxCiphertexts;
        stored.notBeforeBlock = policy.notBeforeBlock;
        stored.notAfterBlock = policy.notAfterBlock;
        stored.decryptNotBefore = policy.decryptNotBefore;
        stored.decryptNotAfter = policy.decryptNotAfter;
        app.createdAtBlock = uint64(block.number);
        app.exists = true;
        epochAidsList[epochId].push(aid);

        emit ApplicationRegistered(epochId, aid, msg.sender, keyX, keyY, policy.mode, poolIndex);
    }

    /// @notice Publish `sk_org` for an `OrganizerLocked` application.
    /// @dev    Permissionless and one-shot: the contract checks
    ///         `sk_org·G == PK_org`, so only the real secret is accepted and
    ///         whoever holds it may publish it. Afterwards the committee can
    ///         combine every ciphertext of the application by itself — the
    ///         combine circuit takes `sk_org` as a private witness, so there
    ///         is no per-ciphertext organizer artefact either way. There is no
    ///         un-reveal.
    function revealOrganizerSecret(bytes12 epochId, bytes32 aid, uint256 organizerSecret) external {
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        // Automatic applications have no secret to reveal: their key is the
        // identity and `organizerSecret` is structurally 0.
        if (app.policy.mode != DKGTypes.AppMode.OrganizerLocked) revert AlreadyRevealed();
        if (app.organizerSecret != 0) revert AlreadyRevealed();
        if (organizerSecret == 0 || organizerSecret >= BabyJubJub.SUBGROUP_ORDER) revert InvalidOrganizerSecret();
        (uint256 gx, uint256 gy) = BabyJubJub.scalarMulBase(organizerSecret);
        DKGTypes.Point storage pk = app.organizerPK;
        if (gx != pk.x || gy != pk.y) revert InvalidOrganizerSecret();

        app.organizerSecret = organizerSecret;
        emit OrganizerSecretRevealed(epochId, aid, organizerSecret);
    }

    /// @notice Read an application record.
    function getApplication(bytes12 epochId, bytes32 aid)
        external
        view
        returns (DKGTypes.Application memory)
    {
        return applications[epochId][aid];
    }

    // ─── Cross-contract APIs (called by DKGManager) ───────────────────────────

    function getOrganizerPK(bytes12 epochId, bytes32 aid) external view returns (uint256, uint256) {
        DKGTypes.Point storage pk = applications[epochId][aid].organizerPK;
        return (pk.x, pk.y);
    }

    function requireDecryptionOpen(bytes12 epochId, bytes32 aid) external view {
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        DKGTypes.AppPolicy storage ap = app.policy;
        // A sealed locked application has nothing the committee may act on:
        // `t` partials alone would already fix `P_j`'s half of the
        // decryption, so neither partials nor combines exist before the
        // organizer reveals `sk_org`.
        if (ap.mode == DKGTypes.AppMode.OrganizerLocked && app.organizerSecret == 0) {
            revert OrganizerSecretNotRevealed();
        }
        uint64 opensAt = ap.decryptNotBefore;
        if (opensAt != 0 && block.timestamp < opensAt) revert DecryptionNotOpen();
        uint64 deadline = ap.decryptNotAfter;
        if (deadline != 0 && block.timestamp > deadline) revert DecryptionClosed();
    }

    function requireCanSubmitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        address sender
    ) external view {
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        DKGTypes.AppPolicy storage ap = app.policy;
        if (!ap.openSubmission && !_isSubmitter(app, sender)) revert NotOwner();
        if (ap.notBeforeBlock != 0 && uint64(block.number) < ap.notBeforeBlock) revert DecryptionNotYetAllowed();
        if (ap.notAfterBlock  != 0 && uint64(block.number) > ap.notAfterBlock)  revert DecryptionExpired();
        if (ap.maxCiphertexts != 0 && ciphertextIndex > ap.maxCiphertexts) revert DecryptionLimitReached();
        // A ciphertext nobody may ever decrypt is not worth storing.
        if (ap.decryptNotAfter != 0 && block.timestamp > ap.decryptNotAfter) revert DecryptionClosed();
    }

    function getRegisteredAids(bytes12 epochId) external view returns (bytes32[] memory) {
        return epochAidsList[epochId];
    }

    // ─── Internals ────────────────────────────────────────────────────────────

    /// @dev Allow-list semantics: empty list → the registrant only.
    function _isSubmitter(DKGTypes.Application storage app, address sender) internal view returns (bool) {
        address[] storage list = app.policy.submitters;
        uint256 n = list.length;
        if (n == 0) return sender == app.creator;
        for (uint256 i; i < n; i++) {
            if (list[i] == sender) return true;
        }
        return false;
    }

    /// @dev Reject contradictory or useless policies up front.
    function _requireValidPolicy(DKGTypes.AppPolicy calldata p) internal view {
        uint256 n = p.submitters.length;
        if (n > MAX_SUBMITTERS || (p.openSubmission && n != 0)) revert InvalidPolicy();
        for (uint256 i; i < n; i++) {
            if (p.submitters[i] == address(0)) revert InvalidPolicy();
        }
        // A decryption window that is already over, or that closes before it
        // opens, would make every ciphertext of the application dead on
        // arrival.
        if (p.decryptNotAfter != 0 && p.decryptNotAfter <= block.timestamp) revert InvalidPolicy();
        if (p.decryptNotBefore != 0 && p.decryptNotAfter != 0 && p.decryptNotBefore > p.decryptNotAfter) {
            revert InvalidPolicy();
        }
        // Same discipline for the submission block window: a window that
        // closes before it opens accepts no ciphertext at all.
        if (p.notBeforeBlock != 0 && p.notAfterBlock != 0 && p.notBeforeBlock > p.notAfterBlock) {
            revert InvalidPolicy();
        }
    }

    /// @dev `aid` is bound into the partial-decrypt and combine proofs as a
    ///      BN254 scalar-field public input, so an id at or above the field
    ///      modulus can never be proven against: reject it up front.
    function _requireValidAid(bytes32 aid) internal pure {
        if (aid == bytes32(0) || uint256(aid) >= BabyJubJub.Q) revert InvalidApplication();
    }

    /// @dev Fiat-Shamir transcript for the organizer Schnorr proof:
    ///        challenge = keccak256(domain || eid || aid || PK_org || A) mod L
    ///      keccak256 instead of Poseidon — see the matching note on
    ///      DKGRegistry._operatorSchnorrChallenge.
    function _organizerSchnorrChallenge(
        bytes12 epochId,
        bytes32 aid,
        uint256 pkX,
        uint256 pkY,
        uint256 ax,
        uint256 ay
    ) internal pure returns (uint256) {
        return uint256(keccak256(abi.encodePacked(
            DKGProtocol.DOMAIN_ORGANIZER_REGISTER_V1,
            epochId,
            aid,
            pkX,
            pkY,
            ax,
            ay
        ))) % BabyJubJub.SUBGROUP_ORDER;
    }

    /// @dev Verify the organizer Schnorr PoK: `z·G == A + c·PK_org`.
    function _verifyOrganizerSchnorr(
        bytes12 epochId,
        bytes32 aid,
        uint256 pkX,
        uint256 pkY,
        uint256 ax,
        uint256 ay,
        uint256 z
    ) internal pure returns (bool) {
        uint256 c = _organizerSchnorrChallenge(epochId, aid, pkX, pkY, ax, ay);
        return BabyJubJub.verifySchnorrEquation(z, c, ax, ay, pkX, pkY);
    }

}
