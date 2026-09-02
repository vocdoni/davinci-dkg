// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {DKGManager} from "../src/DKGManager.sol";
import {DKGAppManager} from "../src/DKGAppManager.sol";
import {IDKGManager} from "../src/interfaces/IDKGManager.sol";
import {IDKGAppManager} from "../src/interfaces/IDKGAppManager.sol";
import {DKGTypes} from "../src/libraries/DKGTypes.sol";
import {
    MockContributionVerifier,
    MockPartialDecryptVerifier,
    MockFinalizeVerifier,
    MockDecryptCombineVerifier,
    TestHelpers
} from "./TestHelpers.t.sol";

/// @title DKGManagerAppTest
/// @notice Tests for the per-application registration surface. Registration is
///         organizer-only: the caller proves knowledge of `sk_org` with a
///         Schnorr PoP and the application key becomes `PK_ep + PK_org`.
///
///         Setup drives the epoch lifecycle against the mock verifiers up to
///         Live; these tests then exercise the application entry points only.
///         The end-to-end ciphertext / organizer-share / combine flow lives in
///         the broader DKGManagerTest suite.
contract DKGManagerAppTest is Test, TestHelpers {
    DKGRegistry public registry;
    DKGManager public manager;
    DKGAppManager public appManager;

    // Schnorr operator vectors (THIS / BEEF), copied from DKGManager.t.sol.
    uint256 internal constant THIS_PUBX =
        17765672829315743641357949553430354448961270408100494783209553303687184365803;
    uint256 internal constant THIS_PUBY =
        13591243454297365848719372676992908085762757043204242277513940025707896351954;
    uint256 internal constant THIS_AX =
        8735331066154608227876753674818247062971894554643955333501578253334124637538;
    uint256 internal constant THIS_AY =
        21272836954776476886917169960736847641993451090581680169106828110848349153636;
    uint256 internal constant THIS_Z =
        1369885591156151396853044744701991591603236518469391786750891091614769051717;
    uint256 internal constant BEEF_PUBX =
        10228722604559478181013548940833210623190136968531440936190496170400150013980;
    uint256 internal constant BEEF_PUBY =
        13886497050333420293068628977630539070604271411621054562122682889313139677221;
    uint256 internal constant BEEF_AX =
        14566234973743316386655481200503006158732416518867749191804063760238069794878;
    uint256 internal constant BEEF_AY =
        13078057343376780948266496972118389875598288245980415388506507934563371992806;
    uint256 internal constant BEEF_Z =
        855437853746059451716869189734643730464853812739723681978635195783451059776;

    bytes32 internal constant TEST_AID = bytes32(uint256(7));

    function setUp() public {
        registry = new DKGRegistry(1_000);
        registry.registerKey(THIS_PUBX, THIS_PUBY, THIS_AX, THIS_AY, THIS_Z);
        vm.prank(address(0xBEEF));
        registry.registerKey(BEEF_PUBX, BEEF_PUBY, BEEF_AX, BEEF_AY, BEEF_Z);
        manager = new DKGManager(
            31337,
            address(registry),
            address(new MockContributionVerifier()),
            address(new MockPartialDecryptVerifier()),
            address(new MockFinalizeVerifier()),
            address(new MockDecryptCombineVerifier()),
            0, 0, 0, 0, 0, 0, 0
        );
        registry.setManager(address(manager));
        appManager = new DKGAppManager(address(manager));
        manager.setAppManager(address(appManager));
        // Operators only enter lotteries of epochs created after they registered.
        vm.roll(block.number + 1);
    }

    // Helper: build a finalized-enough epoch by going through the full
    // lifecycle (createEpoch → claimSlot → submitContribution × 2 →
    // finalizeEpoch). For application-registration tests we only need the
    // resulting `_collectiveKey[eid]` to be non-identity. The mocks accept
    // anything proof-shaped, so this shortcuts neatly.
    function _finalizedEpoch() internal returns (bytes12 epochId) {
        uint64 next = manager.nextEpochStartBlock();
        if (block.number < uint256(next)) vm.roll(uint256(next));
        epochId = manager.createEpoch(2, 2, 2, 10000);
        vm.roll(block.number + 2);
        manager.claimSlot(epochId);
        vm.prank(address(0xBEEF));
        manager.claimSlot(epochId);

        manager.submitContribution(
            epochId, 1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            contributionTranscript(2),
            contributionProof(),
            contributionInput(epochId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );
        vm.prank(address(0xBEEF));
        manager.submitContribution(
            epochId, 2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            contributionTranscript(2),
            contributionProof(),
            contributionInput(
                epochId,
                2,
                2,
                2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1)
            )
        );
        IDKGManager.Epoch memory r = manager.getEpoch(epochId);
        if (block.number < uint256(r.policy.liveNotBeforeBlock)) {
            vm.roll(uint256(r.policy.liveNotBeforeBlock));
        }
        manager.finalizeEpoch(
            epochId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(2),
            finalizeProof(),
            finalizeInput(
                epochId, 2, 2, 2,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );
    }

    /// @dev Register `aid` with a freshly computed organizer PoP.
    function _register(bytes12 epochId, bytes32 aid, DKGTypes.AppPolicy memory policy) internal {
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);
    }

    function _emptyAppPolicy() internal pure returns (DKGTypes.AppPolicy memory) {
        return DKGTypes.AppPolicy({
            authorizedSubmitter: address(0),
            maxCiphertexts: 0,
            notBeforeBlock: 0,
            notAfterBlock: 0
        });
    }

    // ─── Registration (organizer-only) ─────────────────────────────────────

    function test_RegisterApplication_PersistsRecord() public {
        bytes12 epochId = _finalizedEpoch();
        bytes32 aid = bytes32(uint256(42));
        _register(epochId, aid, _emptyAppPolicy());

        (uint256 pkx, uint256 pky) = testOrganizerPK();
        DKGTypes.Application memory app = appManager.getApplication(epochId, aid);
        assertEq(app.creator, address(this));
        assertEq(app.organizerPK.x, pkx);
        assertEq(app.organizerPK.y, pky);
        // A zero authorizedSubmitter resolves to the registrant: there is no
        // open submission.
        assertEq(app.policy.authorizedSubmitter, address(this));
        assertEq(uint256(app.createdAtBlock), block.number);
        assertTrue(app.exists);

        bytes32[] memory aids = appManager.getRegisteredAids(epochId);
        assertEq(aids.length, 1);
        assertEq(uint256(aids[0]), uint256(aid));
    }

    /// @dev An explicit submitter is kept as given.
    function test_RegisterApplication_KeepsExplicitSubmitter() public {
        bytes12 epochId = _finalizedEpoch();
        bytes32 aid = bytes32(uint256(43));
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();
        policy.authorizedSubmitter = address(0xCAFE);
        _register(epochId, aid, policy);
        assertEq(appManager.getApplication(epochId, aid).policy.authorizedSubmitter, address(0xCAFE));
    }

    function test_RegisterApplication_RejectsBadSchnorrProof() public {
        bytes12 epochId = _finalizedEpoch();
        bytes32 aid = bytes32(uint256(44));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);

        // wrong response
        vm.expectRevert(IDKGAppManager.InvalidSchnorrProof.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z + 1);

        // proof bound to a different aid
        (,, uint256 bx, uint256 by, uint256 bz) = organizerPoP(epochId, bytes32(uint256(45)));
        vm.expectRevert(IDKGAppManager.InvalidSchnorrProof.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, bx, by, bz);

        // the honest proof still registers
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
        assertTrue(appManager.getApplication(epochId, aid).exists);
    }

    function test_RegisterApplication_RejectsZeroAid() public {
        bytes12 epochId = _finalizedEpoch();
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, bytes32(0));
        vm.expectRevert(IDKGAppManager.InvalidApplication.selector);
        appManager.registerApplication(epochId, bytes32(0), _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    /// @dev `aid` is bound into every partial-decrypt and combine proof as
    ///      a BN254 scalar-field public input, so ids at or above the field
    ///      modulus could never be decrypted. Reject them at registration.
    function test_RegisterApplication_RejectsAidOutsideScalarField() public {
        bytes12 epochId = _finalizedEpoch();
        bytes32 aid = bytes32(uint256(21888242871839275222246405745257275088548364400416034343698204186575808495617));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGAppManager.InvalidApplication.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
        // One below the modulus is the largest valid id.
        _register(
            epochId,
            bytes32(uint256(21888242871839275222246405745257275088548364400416034343698204186575808495616)),
            _emptyAppPolicy()
        );
    }

    function test_RegisterApplication_RejectsDuplicate() public {
        bytes12 epochId = _finalizedEpoch();
        bytes32 aid = bytes32(uint256(42));
        _register(epochId, aid, _emptyAppPolicy());
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGAppManager.ApplicationAlreadyExists.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    function test_RegisterApplication_RejectsUnknownEpoch() public {
        bytes12 epochId = bytes12(uint96(0xdead));
        bytes32 aid = bytes32(uint256(1));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGManager.InvalidEpoch.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    function test_RegisterApplication_RejectsBeforeFinalization() public {
        uint64 next = manager.nextEpochStartBlock();
        if (block.number < uint256(next)) vm.roll(uint256(next));
        bytes12 epochId = manager.createEpoch(2, 2, 2, 10000);
        bytes32 aid = bytes32(uint256(1));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGManager.InvalidPhase.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    // ─── Organizer share phase gating ──────────────────────────────────────

    function test_SubmitOrganizerShare_RejectsUnknownEpoch() public {
        OrgShare memory sh = testOrganizerShare(bytes12(uint96(0xfeed)), TEST_AID, 1);
        vm.expectRevert(IDKGManager.InvalidEpoch.selector);
        appManager.submitOrganizerShare(
            bytes12(uint96(0xfeed)), TEST_AID, 1,
            TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y,
            sh.deltaX, sh.deltaY, sh.a1x, sh.a1y, sh.a2x, sh.a2y, sh.z
        );
    }

    function test_SubmitOrganizerShare_RejectsCiphertextIndexZero() public {
        bytes12 epochId = _finalizedEpoch();
        OrgShare memory sh = testOrganizerShare(epochId, TEST_AID, 1);
        vm.expectRevert(IDKGAppManager.InvalidCiphertext.selector);
        appManager.submitOrganizerShare(
            epochId, TEST_AID, 0,
            TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y,
            sh.deltaX, sh.deltaY, sh.a1x, sh.a1y, sh.a2x, sh.a2y, sh.z
        );
    }
}
