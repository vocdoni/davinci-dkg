// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {DKGManager} from "../src/DKGManager.sol";
import {DKGAppManager} from "../src/DKGAppManager.sol";
import {IDKGManager} from "../src/interfaces/IDKGManager.sol";
import {IDKGAppManager} from "../src/interfaces/IDKGAppManager.sol";
import {DKGTypes} from "../src/libraries/DKGTypes.sol";
import {BabyJubJub} from "../src/libraries/BabyJubJub.sol";
import {
    MockContributionVerifier,
    MockPartialDecryptVerifier,
    MockDecryptCombineVerifier,
    TestHelpers
} from "./TestHelpers.t.sol";
import {MockVerifier} from "./MockVerifier.sol";

/// @title DKGManagerAppTest
/// @notice Tests for the per-application registration surface. Registration
///         claims one of the epoch's `MAX_K` committee-held pool keys and,
///         in `OrganizerLocked` mode, proves knowledge of `sk_org` with a
///         Schnorr PoP so the application key becomes `P_j + PK_org`.
///         `Automatic` applications have no organizer half at all.
///
///         Setup drives the epoch lifecycle against the mock verifiers up to
///         Live (one finalize proves and stores the whole pool); these tests
///         then exercise the application entry points only. The end-to-end
///         ciphertext / partial / combine flow lives in the broader
///         DKGManagerTest suite.
contract DKGManagerAppTest is Test, TestHelpers {
    DKGRegistry public registry;
    DKGManager public manager;
    DKGAppManager public appManager;

    // Schnorr operator vectors (ALICE / BEEF), copied from DKGManager.t.sol.
    // Committee members must be codeless addresses: submitContribution /
    // finalizeEpoch reject contract callers (DirectCallRequired), and the
    // test contract itself is a contract.
    address internal alice = address(0xA11CE);
    uint256 internal constant ALICE_PUBX =
        14666979294172776374634275498241310759674509575452256743546893482427808967539;
    uint256 internal constant ALICE_PUBY =
        16568773859060023308888034681483224034881825787861296311803627274237556869649;
    uint256 internal constant ALICE_AX =
        206825796546323573150801242605652465781823711295602898923500400375891572685;
    uint256 internal constant ALICE_AY =
        14223109950820609858208501628131627333362274908883166606890781114024603385194;
    uint256 internal constant ALICE_Z =
        1415349030302839010594375480381656799266590814245522079555192493283473059241;
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
        vm.prank(alice);
        registry.registerKey(ALICE_PUBX, ALICE_PUBY, ALICE_AX, ALICE_AY, ALICE_Z);
        vm.prank(address(0xBEEF));
        registry.registerKey(BEEF_PUBX, BEEF_PUBY, BEEF_AX, BEEF_AY, BEEF_Z);
        manager = new DKGManager(
            31337,
            address(registry),
            address(new MockContributionVerifier()),
            address(new MockPartialDecryptVerifier()),
            address(new MockVerifier()),
            address(new MockDecryptCombineVerifier()),
            0, 0, 0, 0, 0, 0, 0
        );
        registry.setManager(address(manager));
        appManager = new DKGAppManager(address(manager));
        manager.setAppManager(address(appManager));
        // Operators only enter lotteries of epochs created after they registered.
        vm.roll(block.number + 1);
    }

    /// @dev Drive an epoch through its whole lifecycle (createEpoch →
    ///      claimSlot × 2 → submitContribution × 2 → proof-carrying
    ///      finalizeEpoch) so registration can claim any key. The mocks
    ///      accept anything proof-shaped, so this shortcuts neatly.
    ///      Contributions and the finalize are pranked with matching
    ///      msg.sender/tx.origin — they must be direct EOA calls.
    function _liveEpoch() internal returns (bytes12 epochId) {
        uint64 next = manager.nextEpochStartBlock();
        if (block.number < uint256(next)) vm.roll(uint256(next));
        epochId = manager.createEpoch(2, 2, 2, 10000);
        vm.roll(block.number + 2);
        vm.prank(alice);
        manager.claimSlot(epochId);
        vm.prank(address(0xBEEF));
        manager.claimSlot(epochId);

        vm.prank(alice, alice);
        manager.submitContribution(
            epochId, 1, contributorCommitmentsHash(1), contributorSharesHash(1),
            contributionTranscript(2, 2), contributionProof(),
            contributionInput(epochId, 2, 2, 1, contributorCommitmentsHash(1), contributorSharesHash(1))
        );
        vm.prank(address(0xBEEF), address(0xBEEF));
        manager.submitContribution(
            epochId, 2, contributorCommitmentsHash(2), contributorSharesHash(2),
            contributionTranscript(2, 2), contributionProof(),
            contributionInput(epochId, 2, 2, 2, contributorCommitmentsHash(2), contributorSharesHash(2))
        );

        IDKGManager.Epoch memory r = manager.getEpoch(epochId);
        if (block.number < uint256(r.policy.liveNotBeforeBlock)) {
            vm.roll(uint256(r.policy.liveNotBeforeBlock));
        }
        vm.prank(alice, alice);
        manager.finalizeEpoch(
            epochId,
            testFinalizeDigest(),
            finalizeTranscript(2, 2),
            finalizeProof(),
            finalizeInput(epochId, 2, 2, 2)
        );
    }

    /// @dev Register `aid` with a freshly computed organizer PoP.
    function _register(bytes12 epochId, bytes32 aid, DKGTypes.AppPolicy memory policy) internal {
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);
    }

    /// @dev Register `aid` in Automatic mode: no organizer key at all.
    function _registerAutomatic(bytes12 epochId, bytes32 aid, DKGTypes.AppPolicy memory policy) internal {
        policy.mode = DKGTypes.AppMode.Automatic;
        appManager.registerApplication(epochId, aid, policy, 0, 0, 0, 0, 0);
    }

    function _emptyAppPolicy() internal pure returns (DKGTypes.AppPolicy memory) {
        return DKGTypes.AppPolicy({
            mode: DKGTypes.AppMode.OrganizerLocked,
            openSubmission: false,
            submitters: new address[](0),
            maxCiphertexts: 0,
            notBeforeBlock: 0,
            notAfterBlock: 0,
            decryptNotBefore: 0,
            decryptNotAfter: 0
        });
    }

    function _submitters(address a) internal pure returns (address[] memory list) {
        list = new address[](1);
        list[0] = a;
    }

    function _submitCiphertextAs(address who, bytes12 epochId, bytes32 aid) internal {
        vm.prank(who);
        manager.submitCiphertext(epochId, aid, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);
    }

    function _submitPartialAs(address who, bytes12 epochId, bytes32 aid, uint16 pIdx) internal {
        uint8 keyIndex = manager.getAppPoolIndex(epochId, aid);
        vm.prank(who);
        manager.submitPartialDecryption(
            epochId, aid, pIdx, 1,
            TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y,
            partialDecryptionHash(pIdx),
            partialDecryptionProof(),
            partialDecryptionInput(epochId, aid, pIdx, 1, keyIndex),
            shareProofFor(keyIndex, 2, pIdx)
        );
    }

    // ─── Registration ─────────────────────────────────────────────────────────

    function test_RegisterApplication_PersistsRecord() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(42));
        _register(epochId, aid, _emptyAppPolicy());

        (uint256 pkx, uint256 pky) = testOrganizerPK();
        DKGTypes.Application memory app = appManager.getApplication(epochId, aid);
        assertEq(app.creator, address(this));
        assertEq(app.organizerPK.x, pkx);
        assertEq(app.organizerPK.y, pky);
        // An empty allow-list means the registrant only.
        assertEq(app.policy.submitters.length, 0);
        assertFalse(app.policy.openSubmission);
        assertEq(uint8(app.policy.mode), uint8(DKGTypes.AppMode.OrganizerLocked));
        assertEq(app.organizerSecret, 0);
        // The first application of the epoch takes pool key 0.
        assertEq(uint256(app.poolIndex), 0);
        assertEq(uint256(manager.getAppPoolIndex(epochId, aid)), 0);
        assertEq(uint256(app.createdAtBlock), block.number);
        assertTrue(app.exists);

        bytes32[] memory aids = appManager.getRegisteredAids(epochId);
        assertEq(aids.length, 1);
        assertEq(uint256(aids[0]), uint256(aid));
    }

    /// @dev An explicit allow-list is kept as given and is exclusive: the
    ///      registrant is not implicitly on it.
    function test_RegisterApplication_KeepsExplicitSubmitters() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(43));
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();
        policy.submitters = _submitters(address(0xCAFE));
        _register(epochId, aid, policy);
        address[] memory stored = appManager.getApplication(epochId, aid).policy.submitters;
        assertEq(stored.length, 1);
        assertEq(stored[0], address(0xCAFE));

        vm.expectRevert(IDKGAppManager.NotOwner.selector);
        _submitCiphertextAs(address(this), epochId, aid);
        _submitCiphertextAs(address(0xCAFE), epochId, aid);
    }

    function test_SubmitCiphertext_RegistrantOnlyByDefault() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(46));
        _register(epochId, aid, _emptyAppPolicy());
        vm.expectRevert(IDKGAppManager.NotOwner.selector);
        _submitCiphertextAs(address(0xCAFE), epochId, aid);
        _submitCiphertextAs(address(this), epochId, aid);
    }

    function test_SubmitCiphertext_OpenSubmission() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(47));
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();
        policy.openSubmission = true;
        _register(epochId, aid, policy);
        _submitCiphertextAs(address(0xCAFE), epochId, aid);
        _submitCiphertextAs(address(0xD00D), epochId, aid);
    }

    // ─── Automatic mode ────────────────────────────────────────────────────

    /// @dev Automatic registration ignores the key and Schnorr arguments and
    ///      stores the identity: there is no organizer half to lose.
    function test_RegisterApplication_Automatic_StoresIdentityKey() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(48));
        (uint256 pkx, uint256 pky) = testOrganizerPK();
        // Pass a perfectly good organizer key and PoP: both must be ignored.
        (,, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();
        policy.mode = DKGTypes.AppMode.Automatic;
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        DKGTypes.Application memory app = appManager.getApplication(epochId, aid);
        assertEq(uint8(app.policy.mode), uint8(DKGTypes.AppMode.Automatic));
        assertEq(app.organizerSecret, 0);
        assertEq(app.organizerPK.x, 0);
        assertEq(app.organizerPK.y, 1);
        (uint256 gx, uint256 gy) = appManager.getOrganizerPK(epochId, aid);
        assertEq(gx, 0);
        assertEq(gy, 1);
    }

    /// @dev Automatic registration also accepts an all-zero key argument.
    function test_RegisterApplication_Automatic_AcceptsZeroArguments() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(49));
        _registerAutomatic(epochId, aid, _emptyAppPolicy());
        assertTrue(appManager.getApplication(epochId, aid).exists);
    }

    // ─── revealOrganizerSecret ────────────────────────────────────────────

    function test_RevealOrganizerSecret_StoresSecretOnce() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(60));
        _register(epochId, aid, _emptyAppPolicy());
        assertEq(appManager.getApplication(epochId, aid).organizerSecret, 0);

        // Permissionless: anyone holding the secret may publish it.
        vm.prank(address(0xCAFE));
        appManager.revealOrganizerSecret(epochId, aid, TEST_ORG_SK);
        assertEq(appManager.getApplication(epochId, aid).organizerSecret, TEST_ORG_SK);

        vm.expectRevert(IDKGAppManager.AlreadyRevealed.selector);
        appManager.revealOrganizerSecret(epochId, aid, TEST_ORG_SK);
    }

    function test_RevealOrganizerSecret_RejectsWrongSecret() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(61));
        _register(epochId, aid, _emptyAppPolicy());

        vm.expectRevert(IDKGAppManager.InvalidOrganizerSecret.selector);
        appManager.revealOrganizerSecret(epochId, aid, TEST_ORG_SK + 1);
        vm.expectRevert(IDKGAppManager.InvalidOrganizerSecret.selector);
        appManager.revealOrganizerSecret(epochId, aid, 0);
        vm.expectRevert(IDKGAppManager.InvalidOrganizerSecret.selector);
        appManager.revealOrganizerSecret(epochId, aid, BabyJubJub.SUBGROUP_ORDER);
        // A congruent-but-out-of-range scalar is rejected too.
        vm.expectRevert(IDKGAppManager.InvalidOrganizerSecret.selector);
        appManager.revealOrganizerSecret(epochId, aid, TEST_ORG_SK + BabyJubJub.SUBGROUP_ORDER);
        // …and the honest one still lands.
        appManager.revealOrganizerSecret(epochId, aid, TEST_ORG_SK);
    }

    /// @dev An Automatic application has no secret to reveal.
    function test_RevealOrganizerSecret_RejectsAutomatic() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(62));
        _registerAutomatic(epochId, aid, _emptyAppPolicy());
        vm.expectRevert(IDKGAppManager.AlreadyRevealed.selector);
        appManager.revealOrganizerSecret(epochId, aid, TEST_ORG_SK);
    }

    function test_RevealOrganizerSecret_RejectsUnknownApplication() public {
        bytes12 epochId = _liveEpoch();
        vm.expectRevert(IDKGAppManager.InvalidApplication.selector);
        appManager.revealOrganizerSecret(epochId, bytes32(uint256(0xDEAD)), TEST_ORG_SK);
    }

    // ─── Policy validation ─────────────────────────────────────────────────

    function test_RegisterApplication_RejectsBadPolicy() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(51));
        // The PoP is computed up front: its precompile call would otherwise
        // be the call `expectRevert` watches.
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();

        policy.openSubmission = true;
        policy.submitters = _submitters(address(0xCAFE));
        vm.expectRevert(IDKGAppManager.InvalidPolicy.selector);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        policy = _emptyAppPolicy();
        policy.submitters = _submitters(address(0));
        vm.expectRevert(IDKGAppManager.InvalidPolicy.selector);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        policy = _emptyAppPolicy();
        policy.submitters = new address[](33);
        for (uint256 i; i < 33; i++) policy.submitters[i] = address(uint160(i + 1));
        vm.expectRevert(IDKGAppManager.InvalidPolicy.selector);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        // A decryption window that already closed.
        policy = _emptyAppPolicy();
        policy.decryptNotAfter = uint64(block.timestamp);
        vm.expectRevert(IDKGAppManager.InvalidPolicy.selector);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        // …or one that closes before it opens.
        policy = _emptyAppPolicy();
        policy.decryptNotBefore = uint64(block.timestamp + 100);
        policy.decryptNotAfter = uint64(block.timestamp + 50);
        vm.expectRevert(IDKGAppManager.InvalidPolicy.selector);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        // A submission block window that closes before it opens.
        policy = _emptyAppPolicy();
        policy.notBeforeBlock = 100;
        policy.notAfterBlock = 50;
        vm.expectRevert(IDKGAppManager.InvalidPolicy.selector);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);

        // The same policy with 32 submitters and a sane window is fine.
        policy = _emptyAppPolicy();
        policy.submitters = new address[](32);
        for (uint256 i; i < 32; i++) policy.submitters[i] = address(uint160(i + 1));
        policy.decryptNotBefore = uint64(block.timestamp + 1);
        policy.decryptNotAfter = uint64(block.timestamp + 2);
        appManager.registerApplication(epochId, aid, policy, pkx, pky, ax, ay, z);
        assertTrue(appManager.getApplication(epochId, aid).exists);
    }

    /// @dev A small-order organizer key passes the Schnorr equation with a
    ///      ground challenge, but the application could then never be
    ///      revealed nor combined: registration refuses it outright.
    function test_RegisterApplication_RejectsSmallOrderOrganizerKey() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(46));
        (,, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGAppManager.PointNotInSubgroup.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), 0, BabyJubJub.Q - 1, ax, ay, z);
    }

    // ─── Decryption window ─────────────────────────────────────────────────

    /// @dev `decryptNotBefore` gates partials and combines but NOT ciphertext
    ///      submission: an organizer collects ballots first and only opens the
    ///      tally later.
    function test_DecryptionWindow_OpensAtNotBefore() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(70));
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();
        policy.decryptNotBefore = uint64(block.timestamp + 100);
        _registerAutomatic(epochId, aid, policy);
        // Submission stays open before the decryption window.
        _submitCiphertextAs(address(this), epochId, aid);

        vm.expectRevert(IDKGAppManager.DecryptionNotOpen.selector);
        appManager.requireDecryptionOpen(epochId, aid);

        uint8 keyIndex = manager.getAppPoolIndex(epochId, aid);
        bytes memory input = partialDecryptionInput(epochId, aid, 1, 1, keyIndex);
        bytes32[] memory path = shareProofFor(keyIndex, 2, 1);
        vm.expectRevert(IDKGAppManager.DecryptionNotOpen.selector);
        vm.prank(alice);
        manager.submitPartialDecryption(
            epochId, aid, 1, 1,
            TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y,
            partialDecryptionHash(1), partialDecryptionProof(), input, path
        );
        vm.expectRevert(IDKGAppManager.DecryptionNotOpen.selector);
        manager.combineDecryption(epochId, aid, 1, bytes32(uint256(1)), 0, "", "", "");

        // Open exactly at the boundary.
        vm.warp(block.timestamp + 100);
        appManager.requireDecryptionOpen(epochId, aid);
        _submitPartialAs(alice, epochId, aid, 1);
    }

    function test_DecryptionWindow_ClosesSubmissionPartialsAndCombine() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(52));
        DKGTypes.AppPolicy memory policy = _emptyAppPolicy();
        policy.decryptNotAfter = uint64(block.timestamp + 100);
        _registerAutomatic(epochId, aid, policy);
        _submitCiphertextAs(address(this), epochId, aid);
        appManager.requireDecryptionOpen(epochId, aid);

        // Still open at the deadline itself.
        vm.warp(block.timestamp + 100);
        appManager.requireDecryptionOpen(epochId, aid);
        _submitPartialAs(alice, epochId, aid, 1);

        vm.warp(block.timestamp + 1);
        vm.expectRevert(IDKGAppManager.DecryptionClosed.selector);
        appManager.requireDecryptionOpen(epochId, aid);
        vm.expectRevert(IDKGAppManager.DecryptionClosed.selector);
        _submitCiphertextAs(address(this), epochId, aid);

        uint8 keyIndex = manager.getAppPoolIndex(epochId, aid);
        bytes memory input = partialDecryptionInput(epochId, aid, 2, 1, keyIndex);
        bytes32[] memory path = shareProofFor(keyIndex, 2, 2);
        vm.prank(address(0xBEEF));
        vm.expectRevert(IDKGAppManager.DecryptionClosed.selector);
        manager.submitPartialDecryption(
            epochId, aid, 2, 1,
            TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y,
            partialDecryptionHash(2), partialDecryptionProof(), input, path
        );
        vm.expectRevert(IDKGAppManager.DecryptionClosed.selector);
        manager.combineDecryption(epochId, aid, 1, bytes32(uint256(1)), 0, "", "", "");
    }

    // ─── Registration guards ───────────────────────────────────────────────

    function test_RegisterApplication_RejectsBadSchnorrProof() public {
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(44));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);

        // wrong response
        vm.expectRevert(IDKGAppManager.InvalidSchnorrProof.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z + 1);

        // proof bound to a different aid
        (,, uint256 bx, uint256 by, uint256 bz) = organizerPoP(epochId, bytes32(uint256(45)));
        vm.expectRevert(IDKGAppManager.InvalidSchnorrProof.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, bx, by, bz);

        // the honest proof still registers — and only now is a key claimed.
        assertEq(uint256(manager.getPoolStatus(epochId)), 0);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
        assertTrue(appManager.getApplication(epochId, aid).exists);
        assertEq(uint256(manager.getPoolStatus(epochId)), 1);
    }

    function test_RegisterApplication_RejectsZeroAid() public {
        bytes12 epochId = _liveEpoch();
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, bytes32(0));
        vm.expectRevert(IDKGAppManager.InvalidApplication.selector);
        appManager.registerApplication(epochId, bytes32(0), _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    /// @dev `aid` is bound into every partial-decrypt and combine proof as
    ///      a BN254 scalar-field public input, so ids at or above the field
    ///      modulus could never be decrypted. Reject them at registration.
    function test_RegisterApplication_RejectsAidOutsideScalarField() public {
        bytes12 epochId = _liveEpoch();
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
        bytes12 epochId = _liveEpoch();
        bytes32 aid = bytes32(uint256(42));
        _register(epochId, aid, _emptyAppPolicy());
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGAppManager.ApplicationAlreadyExists.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
        // The failed duplicate must not have burned a second pool key.
        assertEq(uint256(manager.getPoolStatus(epochId)), 1);
    }

    function test_RegisterApplication_RejectsUnknownEpoch() public {
        bytes12 epochId = bytes12(uint96(0xdead));
        bytes32 aid = bytes32(uint256(1));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGAppManager.InvalidEpoch.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    function test_RegisterApplication_RejectsBeforeFinalization() public {
        uint64 next = manager.nextEpochStartBlock();
        if (block.number < uint256(next)) vm.roll(uint256(next));
        bytes12 epochId = manager.createEpoch(2, 2, 2, 10000);
        bytes32 aid = bytes32(uint256(1));
        (uint256 pkx, uint256 pky, uint256 ax, uint256 ay, uint256 z) = organizerPoP(epochId, aid);
        vm.expectRevert(IDKGAppManager.InvalidPhase.selector);
        appManager.registerApplication(epochId, aid, _emptyAppPolicy(), pkx, pky, ax, ay, z);
    }

    /// @dev Every application of an epoch gets its own committee key.
    function test_RegisterApplication_ClaimsDistinctPoolKeys() public {
        bytes12 epochId = _liveEpoch();
        _registerAutomatic(epochId, bytes32(uint256(80)), _emptyAppPolicy());
        _registerAutomatic(epochId, bytes32(uint256(81)), _emptyAppPolicy());
        assertEq(uint256(appManager.getApplication(epochId, bytes32(uint256(80))).poolIndex), 0);
        assertEq(uint256(appManager.getApplication(epochId, bytes32(uint256(81))).poolIndex), 1);

        (uint256 x0, uint256 y0) = manager.getPoolKey(epochId, 0);
        (uint256 x1, uint256 y1) = manager.getPoolKey(epochId, 1);
        assertTrue(x0 != x1 || y0 != y1);
    }

    function test_RequireDecryptionOpen_RejectsUnknownApplication() public {
        bytes12 epochId = _liveEpoch();
        vm.expectRevert(IDKGAppManager.InvalidApplication.selector);
        appManager.requireDecryptionOpen(epochId, bytes32(uint256(0xDEAD)));
    }
}
