// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

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

    function setUp() public {
        registry = new DKGRegistry();
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
            disclosureAllowed
        );
        // Advance past seedBlock so blockhash is available.
        vm.roll(block.number + 2);
    }

    function _claimAllSlots(bytes12 roundId) internal {
        manager.claimSlot(roundId);
        vm.prank(address(0xBEEF));
        manager.claimSlot(roundId);
    }

    function createFinalizedRound() internal returns (bytes12 roundId) {
        roundId = createSelectedRound();

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );

        vm.prank(address(0xBEEF));
        manager.submitContribution(
            roundId,
            2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            contributionTranscript(2),
            contributionProof(),
            contributionInput(
                roundId,
                2,
                2,
                2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1)
            )
        );

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
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );

        vm.prank(address(0xBEEF));
        manager.submitContribution(
            roundId,
            2,
            bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
            bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1),
            contributionTranscript(2),
            contributionProof(),
            contributionInput(
                roundId,
                2,
                2,
                2,
                bytes32(uint256(CONTRIBUTION_COMMITMENTS_HASH) + 1),
                bytes32(uint256(CONTRIBUTION_ENCRYPTED_SHARES_HASH) + 1)
            )
        );

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
        bytes12 roundId = manager.createRound(2, 2, 2, 10000, 1, uint64(block.number + 5), uint64(block.number + 10), false);

        IDKGManager.Round memory round = manager.getRound(roundId);

        assertEq(round.organizer, address(this));
        assertEq(uint256(round.policy.threshold), 2);
        assertEq(uint256(round.policy.committeeSize), 2);
        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Registration));
    }

    function test_ClaimSlot_RejectsBeforeSeedReady() public {
        bytes12 roundId =
            manager.createRound(2, 2, 2, 10000, 1, uint64(block.number + 5), uint64(block.number + 10), false);
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
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
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
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );
    }

    function test_SubmitContribution_RejectsDuplicates() public {
        bytes12 roundId = createSelectedRound();

        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );

        vm.expectRevert(IDKGManager.AlreadyContributed.selector);
        manager.submitContribution(
            roundId,
            1,
            CONTRIBUTION_COMMITMENTS_HASH,
            CONTRIBUTION_ENCRYPTED_SHARES_HASH,
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
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
            contributionTranscript(2),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );

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

        manager.combineDecryption(
            roundId,
            1,
            COMBINED_DECRYPTION_HASH,
            COMBINED_PLAINTEXT_HASH,
            decryptCombineTranscript(2),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, COMBINED_PLAINTEXT_HASH)
        );

        DKGTypes.CombinedDecryptionRecord memory record = manager.getCombinedDecryption(roundId, 1);
        // Only `completed` is persisted; the hashes live in the DecryptionCombined event.
        assertEq(record.completed ? uint256(1) : uint256(0), 1);
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
            contributionTranscript(1),
            contributionProof(),
            contributionInput(roundId, 2, 2, 1, CONTRIBUTION_COMMITMENTS_HASH, CONTRIBUTION_ENCRYPTED_SHARES_HASH)
        );
    }

    function test_CombineDecryption_RejectsTamperedTranscript() public {
        bytes12 roundId = createFinalizedRound();

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

        vm.expectRevert(IDKGManager.InvalidProofInput.selector);
        manager.combineDecryption(
            roundId,
            1,
            COMBINED_DECRYPTION_HASH,
            COMBINED_PLAINTEXT_HASH,
            decryptCombineTranscript(1),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, COMBINED_PLAINTEXT_HASH)
        );
    }

    function test_CombineDecryption_RejectsMissingPartialsForRequestedCiphertext() public {
        bytes12 roundId = createFinalizedRound();

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

        vm.expectRevert(IDKGManager.InsufficientPartialDecryptions.selector);
        manager.combineDecryption(
            roundId,
            1,
            COMBINED_DECRYPTION_HASH,
            COMBINED_PLAINTEXT_HASH,
            decryptCombineTranscript(2),
            decryptCombineProof(),
            decryptCombineInput(roundId, 2, 2, COMBINED_DECRYPTION_HASH, COMBINED_PLAINTEXT_HASH)
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

    function test_AbortRound_PersistsAbortedStatus() public {
        bytes12 roundId = manager.createRound(2, 2, 2, 10000, 1, uint64(block.number + 5), uint64(block.number + 10), false);

        manager.abortRound(roundId);

        IDKGManager.Round memory round = manager.getRound(roundId);
        assertEq(uint256(round.status), uint256(DKGTypes.RoundStatus.Aborted));
    }
}
