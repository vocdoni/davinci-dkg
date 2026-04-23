// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {DKGManager} from "../src/DKGManager.sol";
import {IDKGManager} from "../src/interfaces/IDKGManager.sol";
import {DKGTypes} from "../src/libraries/DKGTypes.sol";
import {
    MockContributionVerifier,
    MockPartialDecryptVerifier,
    MockFinalizeVerifier,
    MockDecryptCombineVerifier,
    MockRevealSubmitVerifier,
    MockRevealShareVerifier,
    TestHelpers
} from "./TestHelpers.t.sol";

contract DKGManagerTest is Test, TestHelpers {
    DKGRegistry public registry;
    MockContributionVerifier public verifier;
    MockPartialDecryptVerifier public partialVerifier;
    MockFinalizeVerifier public finalizeVerifier;
    MockDecryptCombineVerifier public decryptCombineVerifier;
    MockRevealSubmitVerifier public revealSubmitVerifier;
    MockRevealShareVerifier public revealShareVerifier;
    DKGManager public manager;

    uint64 internal constant INACTIVITY_WINDOW = 1_000;

    function setUp() public {
        registry = new DKGRegistry(INACTIVITY_WINDOW);
        registry.registerKey(101, 201);
        vm.prank(address(0xBEEF));
        registry.registerKey(102, 202);
        verifier = new MockContributionVerifier();
        partialVerifier = new MockPartialDecryptVerifier();
        finalizeVerifier = new MockFinalizeVerifier();
        decryptCombineVerifier = new MockDecryptCombineVerifier();
        revealSubmitVerifier = new MockRevealSubmitVerifier();
        revealShareVerifier = new MockRevealShareVerifier();
        manager = new DKGManager(
            31337,
            address(registry),
            address(verifier),
            address(partialVerifier),
            address(finalizeVerifier),
            address(decryptCombineVerifier),
            address(revealSubmitVerifier),
            address(revealShareVerifier)
        );
        // Wire the liveness callback so submitContribution can refresh
        // the contributor's lastActiveBlock on the registry.
        registry.setManager(address(manager));
    }

    function createSelectedRound() internal returns (bytes12 roundId) {
        roundId = _createLotteryRound(false);
        _claimAllSlots(roundId);
    }

    /// @dev Creates a round with a lottery threshold of "everyone passes" so the two
    /// registered test nodes (address(this) and 0xBEEF) are both eligible.
    function _createLotteryRound(bool disclosureAllowed) internal returns (bytes12 roundId) {
        // α = 1.0 (10000 bps), committeeSize = 2, registered = 2 → all pass.
        // seedDelay = 1, registrationDeadline = +5, contributionDeadline = +20.
        roundId = manager.createRound(
            2,                         // threshold
            2,                         // committeeSize
            2,                         // minValidContributions
            10000,                     // lotteryAlphaBps (α = 1.0)
            1,                          // seedDelay
            uint64(block.number + 5),  // registrationDeadlineBlock
            uint64(block.number + 20), // contributionDeadlineBlock
            uint64(block.number + 21), // finalizeNotBeforeBlock
            disclosureAllowed,
            _emptyDecryptionPolicy()
        );
        // Advance past seedBlock so blockhash is available.
        vm.roll(block.number + 2);
    }

    function _claimAllSlots(bytes12 roundId) internal {
        manager.claimSlot(roundId);
        vm.prank(address(0xBEEF));
        manager.claimSlot(roundId);
    }

    /// @dev Advance to a block at or after `finalizeNotBeforeBlock` for the round.
    function _advanceToFinalize(bytes12 roundId) internal {
        IDKGManager.Round memory r = manager.getRound(roundId);
        if (block.number < uint256(r.policy.finalizeNotBeforeBlock)) {
            vm.roll(uint256(r.policy.finalizeNotBeforeBlock));
        }
    }

    function createFinalizedRound() internal returns (bytes12 roundId) {
        roundId = createSelectedRound();

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );

        vm.prank(address(0xBEEF));
        manager.submitContribution(
            roundId,
            2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(
                roundId,
                2,
                2,
                2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
                0,
                1
            )
        );

        _advanceToFinalize(roundId);
        manager.finalizeRound(
            roundId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(2),
            finalizeProof(),
            finalizeInput(
                roundId,
                2,
                2,
                2,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );
    }

    function createDisclosureRound() internal returns (bytes12 roundId) {
        roundId = _createLotteryRound(true);
        _claimAllSlots(roundId);

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );

        vm.prank(address(0xBEEF));
        manager.submitContribution(
            roundId,
            2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(
                roundId,
                2,
                2,
                2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
                0,
                1
            )
        );

        _advanceToFinalize(roundId);
        manager.finalizeRound(
            roundId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(2),
            finalizeProof(),
            finalizeInput(
                roundId,
                2,
                2,
                2,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );
    }

    function test_CreateRound_PersistsPolicy() public {
        // committeeSize=2, registered=2, α=1.0 → all eligible
        bytes12 roundId = manager.createRound(2, 2, 2, 10000, 1, uint64(block.number + 5), uint64(block.number + 10), uint64(block.number + 11), false, _emptyDecryptionPolicy());

        IDKGManager.Round memory round = manager.getRound(roundId);

        assertEq(round.organizer, address(this));
        assertEq(uint256(round.policy.threshold), 2);
        assertEq(uint256(round.policy.committeeSize), 2);
        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Registration));
    }

    function test_ClaimSlot_RejectsBeforeSeedReady() public {
        bytes12 roundId =
            manager.createRound(2, 2, 2, 10000, 1, uint64(block.number + 5), uint64(block.number + 10), uint64(block.number + 11), false, _emptyDecryptionPolicy());
        // Don't roll forward — seed block has not been mined yet.
        vm.expectRevert(IDKGManager.SeedNotReady.selector);
        manager.claimSlot(roundId);
    }

    function test_ClaimSlot_RejectsDuplicates() public {
        bytes12 roundId = _createLotteryRound(false);
        manager.claimSlot(roundId);
        vm.expectRevert(IDKGManager.AlreadyClaimed.selector);
        manager.claimSlot(roundId);
    }

    function test_ClaimSlot_FillsCommitteeAndAdvancesPhase() public {
        bytes12 roundId = _createLotteryRound(false);
        _claimAllSlots(roundId);

        IDKGManager.Round memory round = manager.getRound(roundId);
        address[] memory selected = manager.selectedParticipants(roundId);

        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Contribution));
        assertEq(selected.length, 2);
        assertEq(selected[0], address(this));
        assertEq(selected[1], address(0xBEEF));
    }

    function test_ClaimSlot_RejectsAfterSlotsFull() public {
        bytes12 roundId = _createLotteryRound(false);
        _claimAllSlots(roundId);
        // round.status is now Contribution, so further claimSlot calls should revert
        // on phase, not on slots-full (we never reach the slots-full check).
        vm.prank(address(0xCAFE));
        vm.expectRevert(IDKGManager.InvalidPhase.selector);
        manager.claimSlot(roundId);
    }

    function test_GetContributionVerifierVKeyHash() public view {
        assertEq(uint256(manager.getContributionVerifierVKeyHash()), uint256(CONTRIBUTION_PROVING_KEY_HASH));
    }

    function test_SubmitContribution_PersistsRecord() public {
        bytes12 roundId = createSelectedRound();

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );

        DKGTypes.ContributionRecord memory record = manager.getContribution(roundId, address(this));
        IDKGManager.Round memory round = manager.getRound(roundId);

        // contributor / commitmentsHash / encryptedSharesHash are no longer persisted
        // (they live in the ContributionSubmitted event); only the fields the contract
        // itself needs at finalize time remain in storage.
        assertEq(uint256(record.contributorIndex), 1);
        assertEq(record.accepted ? uint256(1) : uint256(0), 1);
        assertEq(uint256(round.contributionCount), 1);
    }

    function test_SubmitContribution_RejectsUnselectedOperator() public {
        bytes12 roundId = createSelectedRound();

        vm.prank(address(0xCAFE));
        vm.expectRevert(IDKGManager.NotSelectedParticipant.selector);
        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );
    }

    function test_SubmitContribution_RejectsDuplicates() public {
        bytes12 roundId = createSelectedRound();

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );

        vm.expectRevert(IDKGManager.AlreadyContributed.selector);
        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );
    }

    function test_SubmitContribution_RejectsBadProofInput() public {
        bytes12 roundId = createSelectedRound();

        vm.expectRevert(MockContributionVerifier.InvalidProofInput.selector);
        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            CONTRIBUTION_INPUT_BAD
        );
    }

    function test_GetPartialDecryptVerifierVKeyHash() public view {
        assertEq(uint256(manager.getPartialDecryptVerifierVKeyHash()), uint256(PARTIAL_DECRYPTION_PROVING_KEY_HASH));
    }

    function test_GetRevealSubmitVerifierVKeyHash() public view {
        assertEq(uint256(manager.getRevealSubmitVerifierVKeyHash()), uint256(REVEAL_SUBMIT_PROVING_KEY_HASH));
    }

    function test_FinalizeRound_PersistsHashes() public {
        bytes12 roundId = createFinalizedRound();

        IDKGManager.Round memory round = manager.getRound(roundId);

        // The three commitment hashes are no longer stored; consumers read them from
        // the RoundFinalized event.
        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Finalized));
    }

    function test_FinalizeRound_RejectsInsufficientContributions() public {
        bytes12 roundId = createSelectedRound();

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );

        _advanceToFinalize(roundId);
        vm.expectRevert(IDKGManager.InsufficientContributions.selector);
        manager.finalizeRound(
            roundId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(1),
            finalizeProof(),
            finalizeInput(
                roundId,
                2,
                2,
                1,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );
    }

    function test_FinalizeRound_RejectsBeforeFinalizeNotBeforeBlock() public {
        bytes12 roundId = createSelectedRound();

        manager.submitContribution(
            roundId, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, 1, contributionTranscript(2), contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );
        vm.prank(address(0xBEEF));
        manager.submitContribution(
            roundId, 2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            0, 1, contributionTranscript(2), contributionProof(),
            contributionInput(
                roundId, 2, 2, 2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
                0, 1
            )
        );

        // Both contributions in, threshold met — but finalizeNotBeforeBlock
        // not yet reached. Must revert with InvalidPhase (the gate reuses this
        // selector to keep the contract under the EIP-170 size limit).
        IDKGManager.Round memory r = manager.getRound(roundId);
        assertTrue(block.number < uint256(r.policy.finalizeNotBeforeBlock));

        vm.expectRevert(IDKGManager.InvalidPhase.selector);
        manager.finalizeRound(
            roundId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(2), finalizeProof(),
            finalizeInput(
                roundId, 2, 2, 2,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );

        // Roll exactly to finalizeNotBeforeBlock — should succeed.
        vm.roll(uint256(r.policy.finalizeNotBeforeBlock));
        manager.finalizeRound(
            roundId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(2), finalizeProof(),
            finalizeInput(
                roundId, 2, 2, 2,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );
        IDKGManager.Round memory after_ = manager.getRound(roundId);
        assertEq(uint256(after_.status), uint256(DKGTypes.RoundStatus.Finalized));
    }

    function test_CreateRound_RejectsFinalizeNotBeforeAtOrBelowContribution() public {
        // finalizeNotBeforeBlock == contributionDeadlineBlock → revert
        vm.expectRevert(IDKGManager.InvalidPolicy.selector);
        manager.createRound(
            2, 2, 2, 10000, 1,
            uint64(block.number + 5),
            uint64(block.number + 10),
            uint64(block.number + 10),
            false,
            _emptyDecryptionPolicy()
        );
        // finalizeNotBeforeBlock < contributionDeadlineBlock → revert
        vm.expectRevert(IDKGManager.InvalidPolicy.selector);
        manager.createRound(
            2, 2, 2, 10000, 1,
            uint64(block.number + 5),
            uint64(block.number + 10),
            uint64(block.number + 9),
            false,
            _emptyDecryptionPolicy()
        );
    }

    function test_SubmitPartialDecryption_PersistsRecord() public {
        bytes12 roundId = createFinalizedRound();

        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );

        DKGTypes.PartialDecryptionRecord memory record = manager.getPartialDecryption(roundId, address(this), 1);
        IDKGManager.Round memory round = manager.getRound(roundId);

        // `participant` and `deltaHash` are no longer persisted; off-chain consumers
        // read those from the PartialDecryptionSubmitted event.
        assertEq(uint256(record.participantIndex), 1);
        assertEq(uint256(record.ciphertextIndex), 1);
        assertEq(record.accepted ? uint256(1) : uint256(0), 1);
        assertEq(uint256(round.partialDecryptionCount), 1);
    }

    function test_SubmitPartialDecryption_RejectsBadProofInput() public {
        bytes12 roundId = createFinalizedRound();

        vm.expectRevert(MockPartialDecryptVerifier.InvalidProofInput.selector);
        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            PARTIAL_DECRYPTION_INPUT_BAD
        );
    }

    function test_SubmitPartialDecryption_RejectsBeforeFinalization() public {
        bytes12 roundId = createSelectedRound();

        vm.expectRevert(IDKGManager.InvalidPhase.selector);
        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
    }

    function test_SubmitPartialDecryption_RejectsDuplicates() public {
        bytes12 roundId = createFinalizedRound();

        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );

        vm.expectRevert(IDKGManager.AlreadyPartiallyDecrypted.selector);
        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
    }

    function test_SubmitPartialDecryption_AllowsDistinctCiphertexts() public {
        bytes12 roundId = createFinalizedRound();

        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
        manager.submitPartialDecryption(
            roundId,
            1,
            2,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );

        DKGTypes.PartialDecryptionRecord memory first = manager.getPartialDecryption(roundId, address(this), 1);
        DKGTypes.PartialDecryptionRecord memory second = manager.getPartialDecryption(roundId, address(this), 2);

        assertEq(uint256(first.ciphertextIndex), 1);
        assertEq(uint256(second.ciphertextIndex), 2);
    }

    function test_CombineDecryption_PersistsRecord() public {
        bytes12 roundId = createFinalizedRound();
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
        vm.prank(address(0xBEEF));
        manager.submitPartialDecryption(
            roundId,
            2,
            1,
            partialDecryptionHash(2),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 2, partialDecryptionHash(2))
        );

        uint256 plaintext = uint256(COMBINED_PLAINTEXT_HASH);
        manager.combineDecryption(
            roundId,
            1,
            COMBINED_DECRYPTION_HASH,
            plaintext,
            decryptCombineTranscript(2),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, plaintext)
        );

        DKGTypes.CombinedDecryptionRecord memory record = manager.getCombinedDecryption(roundId, 1);
        assertEq(record.completed ? uint256(1) : uint256(0), 1);
        assertEq(record.plaintext, plaintext);
        assertEq(manager.getPlaintext(roundId, 1), plaintext);
    }

    function test_SubmitRevealedShare_PersistsRecord() public {
        bytes12 roundId = createDisclosureRound();

        manager.submitRevealedShare(
            roundId, 1, uint256(REVEALED_SHARE_HASH), revealShareProof(), revealSubmitInput(roundId, 1, uint256(REVEALED_SHARE_HASH))
        );

        DKGTypes.RevealedShareRecord memory record = manager.getRevealedShare(roundId, address(this));
        IDKGManager.Round memory round = manager.getRound(roundId);

        // `participant` and `shareHash` are no longer persisted; the event carries them.
        assertEq(uint256(record.participantIndex), 1);
        assertEq(record.shareValue, uint256(REVEALED_SHARE_HASH));
        assertEq(record.accepted ? uint256(1) : uint256(0), 1);
        assertEq(uint256(round.revealedShareCount), 1);
    }

    function test_SubmitRevealedShare_RejectsWhenDisclosureDisabled() public {
        bytes12 roundId = createFinalizedRound();

        vm.expectRevert(IDKGManager.DisclosureDisabled.selector);
        manager.submitRevealedShare(
            roundId, 1, uint256(REVEALED_SHARE_HASH), revealShareProof(), revealSubmitInput(roundId, 1, uint256(REVEALED_SHARE_HASH))
        );
    }

    function test_ReconstructSecret_PersistsHashAndCompletesRound() public {
        bytes12 roundId = createDisclosureRound();

        manager.submitRevealedShare(
            roundId, 1, uint256(REVEALED_SHARE_HASH), revealShareProof(), revealSubmitInput(roundId, 1, uint256(REVEALED_SHARE_HASH))
        );
        vm.prank(address(0xBEEF));
        manager.submitRevealedShare(
            roundId,
            2,
            uint256(REVEALED_SHARE_HASH) + 1,
            revealShareProof(),
            revealSubmitInput(roundId, 2, uint256(REVEALED_SHARE_HASH) + 1)
        );

        manager.reconstructSecret(
            roundId,
            DISCLOSURE_HASH,
            RECONSTRUCTED_SECRET_HASH,
            revealShareTranscript(2),
            revealShareProof(),
            revealShareInput(roundId, 2, 2, DISCLOSURE_HASH, RECONSTRUCTED_SECRET_HASH)
        );
 
        IDKGManager.Round memory round = manager.getRound(roundId);
        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Completed));
        // disclosureHash + reconstructedSecretHash live in the SecretReconstructed event now.
    }

    function test_SubmitContribution_RejectsTamperedTranscript() public {
        bytes12 roundId = createSelectedRound();

        vm.expectRevert(IDKGManager.InvalidProofInput.selector);
        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(1),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );
    }

    function test_CombineDecryption_RejectsTamperedTranscript() public {
        bytes12 roundId = createFinalizedRound();
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        manager.submitPartialDecryption(
            roundId,
            1,
            1,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
        vm.prank(address(0xBEEF));
        manager.submitPartialDecryption(
            roundId,
            2,
            1,
            partialDecryptionHash(2),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 2, partialDecryptionHash(2))
        );

        uint256 plaintext = uint256(COMBINED_PLAINTEXT_HASH);
        vm.expectRevert(IDKGManager.InvalidProofInput.selector);
        manager.combineDecryption(
            roundId,
            1,
            COMBINED_DECRYPTION_HASH,
            plaintext,
            decryptCombineTranscript(1),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, plaintext)
        );
    }

    function test_CombineDecryption_RejectsMissingPartialsForRequestedCiphertext() public {
        bytes12 roundId = createFinalizedRound();
        // A ciphertext at index 1 must exist for combine to reach the partial-count check.
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        manager.submitPartialDecryption(
            roundId,
            1,
            2,
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
        vm.prank(address(0xBEEF));
        manager.submitPartialDecryption(
            roundId,
            2,
            2,
            partialDecryptionHash(2),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 2, partialDecryptionHash(2))
        );

        uint256 plaintext = uint256(COMBINED_PLAINTEXT_HASH);
        vm.expectRevert(IDKGManager.InsufficientPartialDecryptions.selector);
        manager.combineDecryption(
            roundId,
            1,
            COMBINED_DECRYPTION_HASH,
            plaintext,
            decryptCombineTranscript(2),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, plaintext)
        );
    }

    function test_ReconstructSecret_RejectsTamperedTranscript() public {
        bytes12 roundId = createDisclosureRound();

        manager.submitRevealedShare(
            roundId, 1, uint256(REVEALED_SHARE_HASH), revealShareProof(), revealSubmitInput(roundId, 1, uint256(REVEALED_SHARE_HASH))
        );
        vm.prank(address(0xBEEF));
        manager.submitRevealedShare(
            roundId,
            2,
            uint256(REVEALED_SHARE_HASH) + 1,
            revealShareProof(),
            revealSubmitInput(roundId, 2, uint256(REVEALED_SHARE_HASH) + 1)
        );

        vm.expectRevert(IDKGManager.InvalidProofInput.selector);
        manager.reconstructSecret(
            roundId,
            DISCLOSURE_HASH,
            RECONSTRUCTED_SECRET_HASH,
            revealShareTranscript(1),
            revealShareProof(),
            revealShareInput(roundId, 2, 2, DISCLOSURE_HASH, RECONSTRUCTED_SECRET_HASH)
        );
    }

    // ── liveness integration ───────────────────────────────────────────────

    function test_CreateRound_UsesActiveCountAsDenominator() public {
        // Register a third node that we'll reap before the round is created.
        address ghost = address(0xDEAD);
        vm.prank(ghost);
        registry.registerKey(33, 44);
        assertEq(registry.activeCount(), 3);
        assertEq(registry.nodeCount(), 3);

        // Age the ghost out and reap.
        vm.roll(block.number + INACTIVITY_WINDOW + 1);
        registry.reap(ghost);
        assertEq(registry.activeCount(), 2);
        assertEq(registry.nodeCount(), 3);

        // Round created now sees registered = activeCount = 2, not 3.
        bytes12 roundId = _createLotteryRound(false);

        // With α=1.0 and active=2, the threshold should be uint256.max
        // (numerator 20000 ≥ 10000) — same as the vanilla 2-node setup.
        IDKGManager.Round memory round = manager.getRound(roundId);
        assertEq(round.lotteryThreshold, type(uint256).max);
    }

    function test_ClaimSlot_RejectsReapedNode() public {
        // Reap address(0xBEEF) (one of the two registered test nodes).
        vm.roll(block.number + INACTIVITY_WINDOW + 1);
        registry.reap(address(0xBEEF));

        // Start a fresh round with activeCount = 1.
        bytes12 roundId = manager.createRound(
            1,
            1,
            1,
            10000,
            1,
            uint64(block.number + 5),
            uint64(block.number + 20),
            uint64(block.number + 21),
            false,
            _emptyDecryptionPolicy()
        );
        vm.roll(block.number + 2);

        // Reaped node cannot claim — the existing status check in claimSlot
        // now triggers because the registry flipped them to INACTIVE.
        vm.prank(address(0xBEEF));
        vm.expectRevert(IDKGManager.NotRegistered.selector);
        manager.claimSlot(roundId);

        // The still-active node can.
        manager.claimSlot(roundId);
    }

    function test_SubmitContribution_RefreshesLastActiveBlock() public {
        bytes12 roundId = createSelectedRound();

        // Record the pre-contribution lastActiveBlock for address(this).
        uint64 before = registry.getNode(address(this)).lastActiveBlock;

        // Advance a few blocks so the SSTORE actually changes value.
        vm.roll(block.number + 10);

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, // commitment0X
            1, // commitment0Y
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );

        uint64 afterBlock = registry.getNode(address(this)).lastActiveBlock;
        assertEq(uint256(afterBlock), block.number);
        assertTrue(afterBlock > before);
    }

    function test_AbortRound_PersistsAbortedStatus() public {
        bytes12 roundId = manager.createRound(2, 2, 2, 10000, 1, uint64(block.number + 5), uint64(block.number + 10), uint64(block.number + 11), false, _emptyDecryptionPolicy());

        manager.abortRound(roundId);

        IDKGManager.Round memory round = manager.getRound(roundId);
        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Aborted));
    }

    // ── M-01: abortRound cannot abort a Finalized round ───────────────────

    function test_AbortRound_RejectsFinalized() public {
        bytes12 roundId = createFinalizedRound();

        vm.expectRevert(IDKGManager.InvalidPhase.selector);
        manager.abortRound(roundId);
    }

    // ── M-02: extendRegistration deadline validation ───────────────────────

    function test_ExtendRegistration_UpdatesSeedBlockAndDeadline() public {
        bytes12 roundId = manager.createRound(
            2, 2, 2, 10000, 1,
            uint64(block.number + 3),   // registrationDeadline
            uint64(block.number + 100), // contributionDeadline
            uint64(block.number + 101), // finalizeNotBeforeBlock
            false,
            _emptyDecryptionPolicy()
        );
        IDKGManager.Round memory before = manager.getRound(roundId);

        // Advance past the registration deadline without filling all slots.
        vm.roll(block.number + 4);
        manager.extendRegistration(roundId);

        IDKGManager.Round memory afterRound = manager.getRound(roundId);
        // Seed block and registration deadline must have advanced.
        assertTrue(afterRound.seedBlock > before.seedBlock);
        assertTrue(afterRound.policy.registrationDeadlineBlock > before.policy.registrationDeadlineBlock);
        // Seed must be reset.
        assertEq(uint256(afterRound.seed), 0);
    }

    function test_ExtendRegistration_RejectsWhenNewDeadlineExceedsContribution() public {
        // registrationDeadlineBlock = block.number + 3
        // contributionDeadlineBlock = block.number + 4  (very tight)
        // window = (block.number+3) - ((block.number+1) - 1) = 3
        // After rolling to block.number+10: newDeadline = 10+3 = 13 > 4 → should revert.
        uint64 base = uint64(block.number);
        bytes12 roundId = manager.createRound(
            2, 2, 2, 10000, 1,
            base + 3,  // registrationDeadline
            base + 4,  // contributionDeadline — very tight
            base + 5,  // finalizeNotBeforeBlock
            false,
            _emptyDecryptionPolicy()
        );

        // Advance well past the registration deadline.
        vm.roll(block.number + 10);

        vm.expectRevert(IDKGManager.InvalidPolicy.selector);
        manager.extendRegistration(roundId);
    }

    // ── M-03/M-04: constructor validation ─────────────────────────────────

    function test_Constructor_RejectsZeroRegistry() public {
        vm.expectRevert(IDKGManager.InvalidAddress.selector);
        new DKGManager(
            31337,
            address(0), // zero registry
            address(verifier),
            address(partialVerifier),
            address(finalizeVerifier),
            address(decryptCombineVerifier),
            address(revealSubmitVerifier),
            address(revealShareVerifier)
        );
    }

    function test_Constructor_RejectsWrongChainId() public {
        vm.expectRevert(IDKGManager.InvalidChainId.selector);
        new DKGManager(
            1, // mainnet — not the test chain (31337)
            address(registry),
            address(verifier),
            address(partialVerifier),
            address(finalizeVerifier),
            address(decryptCombineVerifier),
            address(revealSubmitVerifier),
            address(revealShareVerifier)
        );
    }

    // ── L-02: ciphertextIndex upper bound ─────────────────────────────────

    function test_SubmitPartialDecryption_RejectsCiphertextIndexTooLarge() public {
        bytes12 roundId = createFinalizedRound();

        vm.expectRevert(IDKGManager.InvalidPartialDecryption.selector);
        manager.submitPartialDecryption(
            roundId,
            1,
            257, // > MAX_CIPHERTEXT_INDEX (256)
            partialDecryptionHash(1),
            partialDecryptionProof(),
            partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
    }

    function test_CombineDecryption_RejectsCiphertextIndexTooLarge() public {
        bytes12 roundId = createFinalizedRound();

        uint256 plaintext = uint256(COMBINED_PLAINTEXT_HASH);
        vm.expectRevert(IDKGManager.InvalidCombinedDecryption.selector);
        manager.combineDecryption(
            roundId,
            257, // > MAX_CIPHERTEXT_INDEX (256)
            COMBINED_DECRYPTION_HASH,
            plaintext,
            decryptCombineTranscript(2),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, plaintext)
        );
    }

    // ── submitCiphertext + DecryptionPolicy ──────────────────────────────────

    function test_SubmitCiphertext_StoresHashAndIncrementsCount() public {
        bytes12 roundId = createFinalizedRound();
        bytes32 expected = keccak256(abi.encode(TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y));

        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        assertEq(uint256(manager.getCiphertextHash(roundId, 1)), uint256(expected));
        assertEq(uint256(manager.getRound(roundId).ciphertextCount), 1);
    }

    function test_SubmitCiphertext_RejectsDuplicateIndex() public {
        bytes12 roundId = createFinalizedRound();
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        // Second submission uses on-curve coords (identity) so it passes the
        // well-formedness check and reaches the write-once guard.
        vm.expectRevert(IDKGManager.CiphertextAlreadySubmitted.selector);
        manager.submitCiphertext(roundId, 1, 0, 1, 0, 1);
    }

    function test_SubmitCiphertext_RejectsBeforeFinalized() public {
        bytes12 roundId = _createLotteryRound(false);

        vm.expectRevert(IDKGManager.InvalidPhase.selector);
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);
    }

    function test_SubmitCiphertext_OwnerOnly_BlocksOthers() public {
        bytes12 roundId = _createRoundWithDecryptionPolicy(
            DKGTypes.DecryptionPolicy({
                ownerOnly: true,
                maxDecryptions: 0,
                notBeforeBlock: 0, notBeforeTimestamp: 0,
                notAfterBlock: 0, notAfterTimestamp: 0
            })
        );

        vm.prank(address(0xCAFE));
        vm.expectRevert(IDKGManager.NotOwner.selector);
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        // Owner can submit.
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);
    }

    function test_SubmitCiphertext_NotBeforeBlock_Blocks() public {
        uint64 unlockBlock = uint64(block.number + 1000);
        bytes12 roundId = _createRoundWithDecryptionPolicy(
            DKGTypes.DecryptionPolicy({
                ownerOnly: false,
                maxDecryptions: 0,
                notBeforeBlock: unlockBlock,
                notBeforeTimestamp: 0,
                notAfterBlock: 0,
                notAfterTimestamp: 0
            })
        );

        vm.expectRevert(IDKGManager.DecryptionNotYetAllowed.selector);
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);

        vm.roll(uint256(unlockBlock));
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);
    }

    // Note: timestamp-based gates (notBeforeTimestamp / notAfterTimestamp) share the
    // same revert paths as the block-based gates tested here; the minimal forge-std
    // shim in this repo does not expose vm.warp, so timestamp cases are exercised via
    // the Go integration suite (cmd/davinci-dkg-node tests) rather than here.

    function test_SubmitCiphertext_NotAfterBlock_Blocks() public {
        uint64 cutoff = uint64(block.number + 50);
        bytes12 roundId = _createRoundWithDecryptionPolicy(
            DKGTypes.DecryptionPolicy({
                ownerOnly: false,
                maxDecryptions: 0,
                notBeforeBlock: 0,
                notBeforeTimestamp: 0,
                notAfterBlock: cutoff,
                notAfterTimestamp: 0
            })
        );

        vm.roll(uint256(cutoff) + 1);
        vm.expectRevert(IDKGManager.DecryptionExpired.selector);
        manager.submitCiphertext(roundId, 1, TEST_CT_C1X, TEST_CT_C1Y, TEST_CT_C2X, TEST_CT_C2Y);
    }

    function test_SubmitCiphertext_MaxDecryptions_Caps() public {
        bytes12 roundId = _createRoundWithDecryptionPolicy(
            DKGTypes.DecryptionPolicy({
                ownerOnly: false,
                maxDecryptions: 2,
                notBeforeBlock: 0, notBeforeTimestamp: 0,
                notAfterBlock: 0, notAfterTimestamp: 0
            })
        );

        // Both c1 and c2 = identity for simplicity; the cap check doesn't care about
        // distinct coordinates, only the submission count.
        manager.submitCiphertext(roundId, 1, 0, 1, 0, 1);
        manager.submitCiphertext(roundId, 2, 0, 1, 0, 1);

        vm.expectRevert(IDKGManager.DecryptionLimitReached.selector);
        manager.submitCiphertext(roundId, 3, 0, 1, 0, 1);
    }

    function test_SubmitCiphertext_RejectsOffCurvePoint() public {
        bytes12 roundId = createFinalizedRound();

        // (7001, 8001) does NOT satisfy a·x² + y² = 1 + d·x²·y² (mod Q).
        vm.expectRevert(IDKGManager.InvalidCiphertext.selector);
        manager.submitCiphertext(roundId, 1, 7001, 8001, 0, 1);

        // Canonical-range but off-curve: e.g. (1, 1) — 168700 + 1 = 168701; 1 + 168696 = 168697.
        vm.expectRevert(IDKGManager.InvalidCiphertext.selector);
        manager.submitCiphertext(roundId, 1, 1, 1, 0, 1);

        // Bad c2 with valid c1 also reverts.
        vm.expectRevert(IDKGManager.InvalidCiphertext.selector);
        manager.submitCiphertext(roundId, 1, 0, 1, 1, 1);

        // Coordinates ≥ Q (non-canonical) rejected even if they'd be on-curve post-reduction.
        uint256 Q = 21888242871839275222246405745257275088548364400416034343698204186575808495617;
        vm.expectRevert(IDKGManager.InvalidCiphertext.selector);
        manager.submitCiphertext(roundId, 1, Q, 1, 0, 1);
    }

    function test_CreateRound_RejectsInvalidDecryptionPolicyWindow() public {
        // notAfterBlock <= notBeforeBlock is a degenerate window.
        DKGTypes.DecryptionPolicy memory bad = DKGTypes.DecryptionPolicy({
            ownerOnly: false,
            maxDecryptions: 0,
            notBeforeBlock: 100,
            notBeforeTimestamp: 0,
            notAfterBlock: 100, // equal → empty window
            notAfterTimestamp: 0
        });

        vm.expectRevert(IDKGManager.InvalidDecryptionPolicy.selector);
        manager.createRound(
            2, 2, 2, 10000, 1,
            uint64(block.number + 5),
            uint64(block.number + 20),
            uint64(block.number + 21),
            false,
            bad
        );
    }

    function test_CombineDecryption_RejectsWhenCiphertextNotSubmitted() public {
        bytes12 roundId = createFinalizedRound();

        manager.submitPartialDecryption(
            roundId, 1, 1, partialDecryptionHash(1),
            partialDecryptionProof(), partialDecryptionInput(roundId, 1, partialDecryptionHash(1))
        );
        vm.prank(address(0xBEEF));
        manager.submitPartialDecryption(
            roundId, 2, 1, partialDecryptionHash(2),
            partialDecryptionProof(), partialDecryptionInput(roundId, 2, partialDecryptionHash(2))
        );

        uint256 plaintext = uint256(COMBINED_PLAINTEXT_HASH);
        vm.expectRevert(IDKGManager.CiphertextNotSubmitted.selector);
        manager.combineDecryption(
            roundId, 1, COMBINED_DECRYPTION_HASH, plaintext,
            decryptCombineTranscript(2), decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, plaintext)
        );
    }

    /// @dev Build a finalized round with a custom DecryptionPolicy.
    function _createRoundWithDecryptionPolicy(DKGTypes.DecryptionPolicy memory dp)
        internal
        returns (bytes12 roundId)
    {
        roundId = manager.createRound(
            2, 2, 2, 10000, 1,
            uint64(block.number + 5),
            uint64(block.number + 20),
            uint64(block.number + 21),
            false,
            dp
        );
        vm.roll(block.number + 2);
        _claimAllSlots(roundId);

        manager.submitContribution(
            roundId, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            0, 1, contributionTranscript(2), contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH, 0, 1)
        );
        vm.prank(address(0xBEEF));
        manager.submitContribution(
            roundId, 2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            0, 1, contributionTranscript(2), contributionProof(),
            contributionInput(
                roundId, 2, 2, 2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
                0, 1
            )
        );
        _advanceToFinalize(roundId);
        manager.finalizeRound(
            roundId,
            FINALIZED_AGGREGATE_COMMITMENTS_HASH,
            FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
            FINALIZED_SHARE_COMMITMENT_HASH,
            finalizeTranscript(2), finalizeProof(),
            finalizeInput(
                roundId, 2, 2, 2,
                FINALIZED_AGGREGATE_COMMITMENTS_HASH,
                FINALIZED_COLLECTIVE_PUBLIC_KEY_HASH,
                FINALIZED_SHARE_COMMITMENT_HASH
            )
        );
    }
}
