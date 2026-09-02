// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGAppManager} from "./interfaces/IDKGAppManager.sol";
import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";

/// @title  DKGAppManager
/// @notice Per-application surface (organizer registration + organizer share
///         publication) for the davinci-dkg system. Sibling to `DKGManager`;
///         communicates with it through the narrow `IDKGManager` view surface
///         (`getEpoch`, `getCiphertextHash`, `getCombinedDecryption`) and is
///         consulted by `DKGManager` via `IDKGAppManager` at
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

    /// @dev Per-application registrations, keyed by `(eid, aid)`. See
    ///      DKGTypes.Application for the record shape.
    mapping(bytes12 epochId => mapping(bytes32 aid => DKGTypes.Application app)) internal applications;

    /// @dev List of registered aids per epoch, exposed via `getRegisteredAids`
    ///      for explorers. Append-only; never reordered.
    mapping(bytes12 epochId => bytes32[] aids) internal epochAidsList;

    /// @dev Per-(eid, aid, ciphertextIndex) organizer share commitments:
    ///      `keccak256(abi.encodePacked(Δ.x, Δ.y, A1.x, A1.y, A2.x, A2.y, z))`.
    ///      Written by `submitOrganizerShare`, read by
    ///      `DKGManager.combineDecryption` through `getOrganizerShareHash`,
    ///      which binds the combine transcript's organizer words to it. The
    ///      DLEQ itself is verified inside the combine SNARK, never here.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => bytes32 shareHash)))
        internal organizerShares;

    constructor(address _manager) {
        if (_manager == address(0)) revert InvalidAddress();
        MANAGER = _manager;
    }

    // ─── Application lifecycle ───────────────────────────────────────────────

    /// @notice Register an application against a Live epoch. Verifies a Schnorr
    ///         proof of knowledge of `sk_org`:
    ///
    ///             c = keccak256(domain || eid || aid || PK_org || A) mod L
    ///             z·G == A + c·PK_org
    ///
    ///         The application key is `PK_aid = PK_ep + PK_org`; losing
    ///         `sk_org` makes every ciphertext of the application permanently
    ///         undecryptable.
    /// @dev    `policy.authorizedSubmitter == address(0)` resolves to
    ///         `msg.sender` — there is no open submission.
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
        DKGTypes.Application storage existing = applications[epochId][aid];
        if (existing.exists) revert ApplicationAlreadyExists();

        BabyJubJub.requireValidPoint(pkOrgX, pkOrgY);
        BabyJubJub.requireValidPoint(schnorrAx, schnorrAy);

        if (
            !_verifyOrganizerSchnorr(
                epochId, aid, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ
            )
        ) revert InvalidSchnorrProof();

        DKGTypes.AppPolicy memory resolved = policy;
        if (resolved.authorizedSubmitter == address(0)) resolved.authorizedSubmitter = msg.sender;

        existing.creator = msg.sender;
        existing.organizerPK = DKGTypes.Point({x: pkOrgX, y: pkOrgY});
        existing.policy = resolved;
        existing.createdAtBlock = uint64(block.number);
        existing.exists = true;
        epochAidsList[epochId].push(aid);

        emit ApplicationRegistered(epochId, aid, msg.sender, pkOrgX, pkOrgY);
    }

    /// @notice Publish the organizer's `Δ = sk_org · C_1` for one ciphertext
    ///         together with the Chaum–Pedersen DLEQ `(A1, A2, z)` proving
    ///         `log_G(PK_org) == log_{C_1}(Δ)`.
    /// @dev    Permissionless and unverified on chain: the contract only binds
    ///         the words to `(eid, aid, ctIdx)` by storing their hash. The
    ///         combine SNARK verifies the DLEQ against the registered
    ///         `PK_org` and the challenge `e` that `DKGManager` recomputes
    ///         from the same words, so a bogus share can never produce a
    ///         combine proof — it can only be overwritten by a correct one.
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
    ) external {
        IDKGManager.Epoch memory epoch = IDKGManager(MANAGER).getEpoch(epochId);
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert InvalidCiphertext();
        if (!applications[epochId][aid].exists) revert InvalidApplication();

        // Bind (C1, C2) to the on-chain ciphertext stored on the manager: the
        // DLEQ is only meaningful relative to the authoritative C1.
        bytes32 storedCt = IDKGManager(MANAGER).getCiphertextHash(epochId, aid, ciphertextIndex);
        if (storedCt == bytes32(0)) revert CiphertextNotSubmitted();
        if (_hash4(c1x, c1y, c2x, c2y) != storedCt) revert InvalidProofInput();

        // A malformed share must never brick a ciphertext, so re-submission
        // overwrites — until the plaintext is on chain and the share is moot.
        if (IDKGManager(MANAGER).getCombinedDecryption(epochId, aid, ciphertextIndex).completed) {
            revert AlreadyCombined();
        }

        // Cheap well-formedness. Δ must be a real point (a small-order or
        // identity Δ can never satisfy the in-circuit DLEQ anyway, but
        // rejecting it here keeps the stored words meaningful); A1/A2 must be
        // canonical and on-curve, `z` a canonical scalar.
        BabyJubJub.requireValidPoint(deltaX, deltaY);
        if (!BabyJubJub.isOnCurve(a1x, a1y) || !BabyJubJub.isOnCurve(a2x, a2y)) revert InvalidProofInput();
        if (z >= BabyJubJub.SUBGROUP_ORDER) revert InvalidProofInput();

        organizerShares[epochId][aid][ciphertextIndex] =
            keccak256(abi.encodePacked(deltaX, deltaY, a1x, a1y, a2x, a2y, z));

        emit OrganizerShareSubmitted(epochId, aid, ciphertextIndex, deltaX, deltaY, a1x, a1y, a2x, a2y, z);
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

    function getOrganizerShareHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (bytes32)
    {
        return organizerShares[epochId][aid][ciphertextIndex];
    }

    function requireCanSubmitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        address sender
    ) external view {
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        DKGTypes.AppPolicy memory ap = app.policy;
        // Always set: registration resolves a zero submitter to the registrant.
        if (sender != ap.authorizedSubmitter) revert NotOwner();
        if (ap.notBeforeBlock != 0 && uint64(block.number) < ap.notBeforeBlock) revert DecryptionNotYetAllowed();
        if (ap.notAfterBlock  != 0 && uint64(block.number) > ap.notAfterBlock)  revert DecryptionExpired();
        if (ap.maxCiphertexts != 0 && ciphertextIndex > ap.maxCiphertexts) revert DecryptionLimitReached();
    }

    function getRegisteredAids(bytes12 epochId) external view returns (bytes32[] memory) {
        return epochAidsList[epochId];
    }

    // ─── Internals ────────────────────────────────────────────────────────────

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
