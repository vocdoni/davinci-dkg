// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGAppManager} from "./interfaces/IDKGAppManager.sol";
import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";

/// @title  DKGAppManager
/// @notice Per-application surface (registration + organizer Schnorr verification +
///         organizer share submission) for the davinci-dkg system. Sibling to
///         `DKGManager`; communicates with it through the narrow `IDKGManager`
///         view surface (`getEpoch`, `getCollectivePublicKey`, `getCiphertextHash`)
///         and is consulted by `DKGManager` via `IDKGAppManager` at
///         `submitCiphertext` / `combineDecryption` time.
///
///         Was carved out of `DKGManager` to keep the latter's runtime size
///         under the EIP-170 24,576-byte limit. No protocol behaviour change.
contract DKGAppManager is IDKGAppManager {
    /// @dev Must mirror `DKGManager.MAX_CIPHERTEXT_INDEX`.
    uint16 internal constant MAX_CIPHERTEXT_INDEX = 256;

    /// @notice The sibling DKGManager whose epochs this contract registers
    ///         applications against. Immutable.
    address public immutable MANAGER;

    /// @notice The same partial-decrypt Groth16 verifier used by
    ///         DKGManager.submitPartialDecryption. Reused here for the
    ///         organizer-share DLEQ check (paper §6.3).
    address public immutable PARTIAL_DECRYPT_VERIFIER;

    /// @dev Per-application registrations, keyed by `(eid, aid)`. See
    ///      DKGTypes.Application for the record shape. Both
    ///      registerApplication and registerApplicationCoDec write here; the
    ///      mode flag distinguishes the two paths at decryption time.
    mapping(bytes12 epochId => mapping(bytes32 aid => DKGTypes.Application app)) internal applications;

    /// @dev List of registered aids per epoch (excluding bytes32(0)), exposed
    ///      via `getRegisteredAids` for explorers. Append-only; never reordered.
    mapping(bytes12 epochId => bytes32[] aids) internal epochAidsList;

    /// @dev Per-(eid, aid, ciphertextIndex) organizer share submissions for
    ///      mode-1 applications. Written by submitOrganizerShare and read by
    ///      DKGManager.combineDecryption (via `getCombineCorrection`) when the
    ///      application's mode is OrganizerCoDec.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => DKGTypes.OrganizerShareRecord)))
        internal organizerShares;

    constructor(address _manager, address _partialDecryptVerifier) {
        if (_manager == address(0)) revert InvalidAddress();
        if (_partialDecryptVerifier == address(0)) revert InvalidVerifier();
        MANAGER = _manager;
        PARTIAL_DECRYPT_VERIFIER = _partialDecryptVerifier;
    }

    // ─── Application lifecycle (paper §4.3, §6) ──────────────────────────────

    /// @notice Register an application against a finalized epoch in
    ///         public-derivation mode (paper §4.3). Computes the per-application
    ///         derivation tag `S = keccak256(eid || PK_ep || aid) mod q` and
    ///         stores the application record.
    function registerApplication(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy
    ) external {
        IDKGManager.Epoch memory epoch = IDKGManager(MANAGER).getEpoch(epochId);
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (aid == bytes32(0)) revert InvalidApplication();
        DKGTypes.Application storage existing = applications[epochId][aid];
        if (existing.exists) revert ApplicationAlreadyExists();

        DKGTypes.Point memory pkep = IDKGManager(MANAGER).getCollectivePublicKey(epochId);
        // Defense in depth: a finalized epoch always has the key written; a
        // y == 0 here would mean a corrupted finalize. Mirrors the original
        // DKGManager check on `_collectiveKey[epochId].y`.
        if (pkep.y == 0) revert InvalidEpoch();

        // S = keccak256(eid || PK_ep.x || PK_ep.y || aid) mod q
        uint256 s = uint256(
            keccak256(abi.encodePacked(epochId, pkep.x, pkep.y, aid))
        ) % BabyJubJub.SUBGROUP_ORDER;

        existing.creator = msg.sender;
        existing.mode = DKGTypes.AppMode.PublicDerivation;
        existing.derivationS = s;
        // organizerPK is only consumed in mode 1; leave the slot zero for mode 0.
        // The getter normalizes a zero slot to identity for consumers.
        existing.policy = policy;
        existing.createdAtBlock = uint64(block.number);
        existing.exists = true;
        epochAidsList[epochId].push(aid);

        emit ApplicationRegistered(
            epochId,
            aid,
            msg.sender,
            uint8(DKGProtocol.MODE_PUBLIC_DERIVATION),
            s,
            0,
            1
        );
    }

    /// @notice Register an application against a finalized epoch in
    ///         organizer co-decryption mode (paper §6). Verifies a Schnorr
    ///         proof of knowledge of `sk_org`:
    ///
    ///             c = Poseidon(domain || eid || aid || PK_org || A)
    ///             z·G == A + c·PK_org
    function registerApplicationCoDec(
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
        if (aid == bytes32(0)) revert InvalidApplication();
        DKGTypes.Application storage existing = applications[epochId][aid];
        if (existing.exists) revert ApplicationAlreadyExists();

        BabyJubJub.requireValidPoint(pkOrgX, pkOrgY);
        BabyJubJub.requireValidPoint(schnorrAx, schnorrAy);

        if (
            !_verifyOrganizerSchnorr(
                epochId, aid, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ
            )
        ) revert InvalidSchnorrProof();

        existing.creator = msg.sender;
        existing.mode = DKGTypes.AppMode.OrganizerCoDec;
        existing.derivationS = 0;
        existing.organizerPK = DKGTypes.Point({x: pkOrgX, y: pkOrgY});
        existing.policy = policy;
        existing.createdAtBlock = uint64(block.number);
        existing.exists = true;
        epochAidsList[epochId].push(aid);

        emit ApplicationRegistered(
            epochId,
            aid,
            msg.sender,
            uint8(DKGProtocol.MODE_ORGANIZER_CODEC),
            0,
            pkOrgX,
            pkOrgY
        );
    }

    /// @notice Submit an organizer's `Δ_org = sk_org · C_1` share for a
    ///         mode-1 application's ciphertext, together with its DLEQ proof.
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
    ) external {
        IDKGManager.Epoch memory epoch = IDKGManager(MANAGER).getEpoch(epochId);
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert InvalidCiphertext();
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        if (uint8(app.mode) != uint8(DKGProtocol.MODE_ORGANIZER_CODEC)) revert InvalidApplication();

        // Bind C1 to the on-chain ciphertext stored on the manager.
        bytes32 storedCt = IDKGManager(MANAGER).getCiphertextHash(epochId, aid, ciphertextIndex);
        if (storedCt == bytes32(0)) revert CiphertextNotSubmitted();
        if (_hash4(c1x, c1y, c2x, c2y) != storedCt) revert InvalidProofInput();

        if (organizerShares[epochId][aid][ciphertextIndex].accepted) revert AlreadyPartiallyDecrypted();

        BabyJubJub.requireValidPoint(deltaOrgX, deltaOrgY);

        IZKVerifier(PARTIAL_DECRYPT_VERIFIER).verifyProof(dleqProof, dleqInput);
        uint256[16] memory pi = abi.decode(dleqInput, (uint256[16]));
        if (
            pi[0] != _epochScalar(epochId)
                || pi[1] != uint256(aid)
                || pi[2] != uint256(ciphertextIndex)
                || pi[3] != uint256(DKGProtocol.ROLE_ORGANIZER)
                || pi[4] != 0
                || pi[5] != c1x
                || pi[6] != c1y
                || pi[7] != app.organizerPK.x
                || pi[8] != app.organizerPK.y
                || pi[9] != deltaOrgX
                || pi[10] != deltaOrgY
        ) revert InvalidProofInput();

        organizerShares[epochId][aid][ciphertextIndex] = DKGTypes.OrganizerShareRecord({
            deltaOrg: DKGTypes.Point({x: deltaOrgX, y: deltaOrgY}),
            accepted: true
        });

        emit OrganizerShareSubmitted(epochId, aid, ciphertextIndex, deltaOrgX, deltaOrgY);
    }

    /// @notice Read an application record.
    function getApplication(bytes12 epochId, bytes32 aid)
        external
        view
        returns (DKGTypes.Application memory app)
    {
        app = applications[epochId][aid];
        // Mode-0 apps don't store organizerPK (they leave it zeroed); normalize
        // to the BabyJubJub identity (0, 1) for off-chain consumers.
        if (app.organizerPK.y == 0) {
            app.organizerPK = DKGTypes.Point({x: 0, y: 1});
        }
    }

    // ─── Cross-contract APIs (called by DKGManager) ───────────────────────────

    function getCombineCorrection(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (uint8 mode, uint256 derivationS, uint256 deltaOrgX, uint256 deltaOrgY)
    {
        if (aid == bytes32(0)) {
            return (uint8(DKGProtocol.MODE_PUBLIC_DERIVATION), 0, 0, 1);
        }
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        mode = uint8(app.mode);
        if (mode == uint8(DKGProtocol.MODE_PUBLIC_DERIVATION)) {
            derivationS = app.derivationS;
            deltaOrgX = 0;
            deltaOrgY = 1;
        } else {
            DKGTypes.OrganizerShareRecord storage org = organizerShares[epochId][aid][ciphertextIndex];
            // Mirror the previous DKGManager error for a missing organizer share.
            if (!org.accepted) revert InsufficientPartialDecryptions();
            derivationS = 0;
            deltaOrgX = org.deltaOrg.x;
            deltaOrgY = org.deltaOrg.y;
        }
    }

    function requireCanSubmitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        address sender
    ) external view {
        if (aid == bytes32(0)) return;
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        DKGTypes.AppPolicy memory ap = app.policy;
        if (ap.authorizedSubmitter != address(0) && sender != ap.authorizedSubmitter) revert NotOwner();
        if (ap.notBeforeBlock != 0 && uint64(block.number) < ap.notBeforeBlock) revert DecryptionNotYetAllowed();
        if (ap.notAfterBlock  != 0 && uint64(block.number) > ap.notAfterBlock)  revert DecryptionExpired();
        if (ap.maxCiphertexts != 0 && ciphertextIndex > ap.maxCiphertexts) revert DecryptionLimitReached();
    }

    function getRegisteredAids(bytes12 epochId) external view returns (bytes32[] memory) {
        return epochAidsList[epochId];
    }

    // ─── Internals ────────────────────────────────────────────────────────────

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
    ) internal view returns (bool) {
        uint256 c = _organizerSchnorrChallenge(epochId, aid, pkX, pkY, ax, ay);
        return BabyJubJub.verifySchnorrEquation(z, c, ax, ay, pkX, pkY);
    }

    function _epochScalar(bytes12 epochId) internal pure returns (uint256) {
        return uint256(uint96(epochId));
    }

    function _hash4(uint256 a, uint256 b, uint256 c, uint256 d) internal pure returns (bytes32 h) {
        assembly ("memory-safe") {
            let p := mload(0x40)
            mstore(p, a)
            mstore(add(p, 0x20), b)
            mstore(add(p, 0x40), c)
            mstore(add(p, 0x60), d)
            h := keccak256(p, 0x80)
        }
    }
}
