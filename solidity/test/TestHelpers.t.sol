// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../src/interfaces/IZKVerifier.sol";
import {BRLC} from "../src/libraries/BRLC.sol";
import {DKGTypes} from "../src/libraries/DKGTypes.sol";
import {MAX_N} from "../src/libraries/Sizes.sol";
import {TestInputs} from "./TestInputs.t.sol";

abstract contract TestHelpers is TestInputs {
    /// @dev All-zero decryption policy: owner-only disabled, no time locks,
    /// no submission cap. Used by tests that don't care about submission gating.

    /// @dev Canonical on-curve ciphertext used by the ZK-mock tests. Both c1
    /// and c2 must be in the prime-order subgroup of BabyJubJub (gnark RTE
    /// form) — the DKGManager's `_requireValidEncryptionPoint` enforces this
    /// at submitCiphertext time. Identity (0, 1) is not acceptable.
    ///   c1 = 1·G  (the gnark RTE generator)
    ///   c2 = 4096·G (the SCHNORR_THIS vector's pubkey, secret = 0x1000)
    uint256 internal constant TEST_CT_C1X = 9671717474070082183213120605117400219616337014328744928644933853176787189663;
    uint256 internal constant TEST_CT_C1Y = 16950150798460657717958625567821834550301663161624707787222815936182638968203;
    uint256 internal constant TEST_CT_C2X = 17765672829315743641357949553430354448961270408100494783209553303687184365803;
    uint256 internal constant TEST_CT_C2Y = 13591243454297365848719372676992908085762757043204242277513940025707896351954;

    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:decrypt-combine:v1");
    bytes32 internal constant FINALIZE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:finalize:v1");

    function contributionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(1), 2, 3, 4, 5, 6, 7, 8]);
    }

    function contributionInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            CONTRIBUTION_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(commitmentsHash, encryptedSharesHash, keccak256(contributionTranscript(committeeSize))))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(threshold),
                uint256(committeeSize),
                uint256(contributorIndex),
                uint256(commitmentsHash),
                uint256(encryptedSharesHash),
                challenge,
                contributionTranscriptCommitment(challenge, committeeSize)
            ]
        );
    }

    // The test fixtures below mirror the on-chain transcript layouts,
    // parameterized in `MAX_N` (imported from src/libraries/Sizes.sol):
    //   contribution: 8N words = (2N + N + 2N + 2N + N) layout.
    //   finalize:     2N²+5N words.
    //   combine:      4 + 3N words.
    //   reconstruct:  2N words.
    /// @dev The first two committee slots use real on-curve BabyJubJub keys
    ///      (THIS and BEEF Schnorr vectors from cmd/operator-schnorr-vectors).
    ///      Slots 3+ stay at zero and are padded with the identity (0,1) below.
    ///      This matches what the DKGManager test fixtures register at setUp,
    ///      so the committee snapshot hash and the contribution transcript
    ///      stay byte-identical. All tests fix committeeSize = 2.
    function _slotPubKey(uint256 i) internal pure returns (uint256 x, uint256 y) {
        if (i == 0) {
            return (
                17765672829315743641357949553430354448961270408100494783209553303687184365803,
                13591243454297365848719372676992908085762757043204242277513940025707896351954
            );
        }
        if (i == 1) {
            return (
                10228722604559478181013548940833210623190136968531440936190496170400150013980,
                13886497050333420293068628977630539070604271411621054562122682889313139677221
            );
        }
        revert("test fixture: only 2 committee members supported");
    }

    function contributionTranscript(uint16 committeeSize) internal pure returns (bytes memory) {
        uint256[2 * MAX_N] memory commitments;
        uint256[MAX_N] memory recipientIndexes;
        uint256[2 * MAX_N] memory recipientPubKeys;
        uint256[2 * MAX_N] memory ephemerals;
        uint256[MAX_N] memory maskedShares;
        for (uint256 i = 0; i < MAX_N; i++) {
            commitments[i * 2 + 1] = 1;
            recipientPubKeys[i * 2 + 1] = 1;
            ephemerals[i * 2 + 1] = 1;
        }
        for (uint256 i = 0; i < committeeSize; i++) {
            commitments[i * 2 + 1] = 0;
            recipientIndexes[i] = i + 1;
            (uint256 px, uint256 py) = _slotPubKey(i);
            recipientPubKeys[i * 2] = px;
            recipientPubKeys[i * 2 + 1] = py;
            ephemerals[i * 2] = 300 + i + 1;
            ephemerals[i * 2 + 1] = 400 + i + 1;
            maskedShares[i] = 500 + i + 1;
        }
        return abi.encode(commitments, recipientIndexes, recipientPubKeys, ephemerals, maskedShares);
    }

    function contributionTranscriptCommitment(uint256 challenge, uint16 committeeSize) internal pure returns (uint256) {
        uint256[] memory values = new uint256[](8 * MAX_N);
        for (uint256 i = 0; i < MAX_N; i++) {
            values[i * 2 + 1] = 1;                       // commitments y pad
            values[3 * MAX_N + i * 2 + 1] = 1;           // recipientPubKeys y pad (offset 2N+N)
            values[5 * MAX_N + i * 2 + 1] = 1;           // ephemerals y pad (offset 2N+N+2N)
        }
        uint256 cursor = 2 * MAX_N; // recipientIndexes start (after 2N commitments)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[i * 2 + 1] = 0;
            values[cursor++] = i + 1;
        }
        cursor = 3 * MAX_N; // recipientPubKeys start (2N+N)
        for (uint256 i = 0; i < committeeSize; i++) {
            (uint256 px, uint256 py) = _slotPubKey(i);
            values[cursor++] = px;
            values[cursor++] = py;
        }
        cursor = 5 * MAX_N; // ephemerals start (2N+N+2N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 300 + i + 1;
            values[cursor++] = 400 + i + 1;
        }
        cursor = 7 * MAX_N; // maskedShares start (2N+N+2N+2N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 500 + i + 1;
        }
        return BRLC.commit(challenge, values);
    }

    function partialDecryptionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(11), 12, 13, 14, 15, 16, 17, 18]);
    }

    /// partialdecrypt layout (16 public inputs):
    ///   [0] eid, [1] aid, [2] ctIdx, [3] role, [4] participantIndex,
    ///   [5..6] C1.x/y, [7..8] D_i.x/y, [9..10] delta.x/y,
    ///   [11..12] A1.x/y, [13..14] A2.x/y, [15] response.
    /// Tests pass aid=0 (legacy path) and role=COMMITTEE=1.
    /// 3-arg overload defaults ciphertextIndex to 1 for the most common case.
    function partialDecryptionInput(bytes12 epochId, uint16 participantIndex, bytes32 unused)
        internal
        pure
        returns (bytes memory)
    {
        return partialDecryptionInputCt(epochId, participantIndex, 1, unused);
    }

    function partialDecryptionInputCt(
        bytes12 epochId,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        bytes32
    ) internal pure returns (bytes memory) {
        uint256[16] memory inputs;
        inputs[0] = uint256(uint96(epochId));
        inputs[1] = 0;                          // aid (legacy)
        inputs[2] = ciphertextIndex;
        inputs[3] = 1;                          // role = COMMITTEE
        inputs[4] = participantIndex;
        // pi[5..6] = C1, must match the test ciphertext fixture so
        // submitPartialDecryption's ciphertext binding accepts.
        inputs[5] = TEST_CT_C1X;
        inputs[6] = TEST_CT_C1Y;
        inputs[7] = 1000 + participantIndex;    // D_i.x
        inputs[8] = 2000 + participantIndex;    // D_i.y
        inputs[9] = 7000 + participantIndex;    // delta.x
        inputs[10] = 8000 + participantIndex;   // delta.y
        return abi.encode(inputs);
    }

    function partialDecryptionHash(uint16 participantIndex) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint256(7000 + participantIndex), uint256(8000 + participantIndex)));
    }

    function finalizeProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(21), 22, 23, 24, 25, 26, 27, 28]);
    }

    function finalizeInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 acceptedCount,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            FINALIZE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(
                aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash,
                FINALIZED_ROWS_HASH, keccak256(finalizeTranscript(acceptedCount))
            ))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(threshold),
                uint256(committeeSize),
                uint256(acceptedCount),
                uint256(aggregateCommitmentsHash),
                uint256(collectivePublicKeyHash),
                uint256(shareCommitmentHash),
                uint256(FINALIZED_ROWS_HASH),
                challenge,
                finalizeTranscriptCommitment(challenge, acceptedCount)
            ]
        );
    }

    function finalizeTranscript(uint16 acceptedCount) internal pure returns (bytes memory) {
        uint256[MAX_N] memory participantIndexes;
        uint256[2 * MAX_N * MAX_N] memory contributionCommitments;
        uint256[2 * MAX_N] memory aggregateCommitments;
        uint256[2 * MAX_N] memory shareCommitments;
        // Per participant, mirror contributionTranscript(committeeSize=2):
        // pt[0]=pt[1]=(0,0); pt[2..N-1]=(0,1) → odd indices 5,7,...,2N-1 = 1
        // (each participant's commitments occupy 2N words).
        for (uint256 i = 0; i < acceptedCount; i++) {
            participantIndexes[i] = i + 1;
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                contributionCommitments[i * (2 * MAX_N) + k] = 1;
            }
            shareCommitments[i * 2] = 1000 + i + 1;
            shareCommitments[i * 2 + 1] = 2000 + i + 1;
        }
        // The contract requires aggregate[0] to equal
        // the accumulated _collectiveKey. The fixtures all use commitment0 =
        // (0, 1) (identity), so the running sum is identity (0, 1).
        aggregateCommitments[1] = 1;
        return abi.encode(participantIndexes, contributionCommitments, aggregateCommitments, shareCommitments);
    }

    function finalizeTranscriptCommitment(uint256 challenge, uint16 acceptedCount) internal pure returns (uint256) {
        // Layout (2N²+5N words): [0..N) participantIndexes,
        //                        [N..N+2N²) contributionCommitments,
        //                        [N+2N²..N+2N²+2N) aggregateCommitments,
        //                        [N+2N²+2N..2N²+5N) shareCommitments.
        uint256[] memory values = new uint256[](2 * MAX_N * MAX_N + 5 * MAX_N);
        for (uint256 i = 0; i < acceptedCount; i++) {
            values[i] = i + 1;
            uint256 offset = MAX_N + i * (2 * MAX_N); // N + i*2N
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                values[offset + k] = 1;
            }
        }
        // shareCommitments start: N + 2N² + 2N
        uint256 shareOffsetBase = MAX_N + 2 * MAX_N * MAX_N + 2 * MAX_N;
        for (uint256 i = 0; i < acceptedCount; i++) {
            uint256 offset = shareOffsetBase + i * 2;
            values[offset] = 1000 + i + 1;
            values[offset + 1] = 2000 + i + 1;
        }
        // aggregateCommitments[0] = identity (0, 1) — see finalizeTranscript.
        // Slot offset is N + 2N²; aggregate[0].y is at offset+1.
        values[MAX_N + 2 * MAX_N * MAX_N + 1] = 1;
        return BRLC.commit(challenge, values);
    }

    /// @dev Variant of finalizeTranscript used by the duplicate-row regression
    /// test. Sets participantIndexes = [1, 1] for an
    /// acceptedCount of 2 instead of [1, 2].
    function finalizeTranscriptWithDuplicateRows() internal pure returns (bytes memory) {
        uint256[MAX_N] memory participantIndexes;
        uint256[2 * MAX_N * MAX_N] memory contributionCommitments;
        uint256[2 * MAX_N] memory aggregateCommitments;
        uint256[2 * MAX_N] memory shareCommitments;
        participantIndexes[0] = 1;
        participantIndexes[1] = 1;
        for (uint256 i = 0; i < 2; i++) {
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                contributionCommitments[i * (2 * MAX_N) + k] = 1;
            }
            shareCommitments[i * 2] = 1000 + i + 1;
            shareCommitments[i * 2 + 1] = 2000 + i + 1;
        }
        aggregateCommitments[1] = 1;
        return abi.encode(participantIndexes, contributionCommitments, aggregateCommitments, shareCommitments);
    }

    function finalizeInputWithDuplicateRows(
        bytes12 epochId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            FINALIZE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(
                aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash,
                FINALIZED_ROWS_HASH, keccak256(finalizeTranscriptWithDuplicateRows())
            ))
        );
        return abi.encode(
            [
                uint256(uint96(epochId)),
                uint256(2),
                uint256(2),
                uint256(2),
                uint256(aggregateCommitmentsHash),
                uint256(collectivePublicKeyHash),
                uint256(shareCommitmentHash),
                uint256(FINALIZED_ROWS_HASH),
                challenge,
                _finalizeTranscriptCommitmentWithDuplicateRows(challenge)
            ]
        );
    }

    function _finalizeTranscriptCommitmentWithDuplicateRows(uint256 challenge) private pure returns (uint256) {
        uint256[] memory values = new uint256[](2 * MAX_N * MAX_N + 5 * MAX_N);
        values[0] = 1;
        values[1] = 1;
        for (uint256 i = 0; i < 2; i++) {
            uint256 offset = MAX_N + i * (2 * MAX_N);
            for (uint256 k = 5; k < 2 * MAX_N; k += 2) {
                values[offset + k] = 1;
            }
        }
        uint256 shareOffsetBase = MAX_N + 2 * MAX_N * MAX_N + 2 * MAX_N;
        for (uint256 i = 0; i < 2; i++) {
            uint256 offset = shareOffsetBase + i * 2;
            values[offset] = 1000 + i + 1;
            values[offset + 1] = 2000 + i + 1;
        }
        values[MAX_N + 2 * MAX_N * MAX_N + 1] = 1;
        return BRLC.commit(challenge, values);
    }

    function decryptCombineProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(31), 32, 33, 34, 35, 36, 37, 38]);
    }

    /// @dev 13-element layout matching the P5/P6 combine circuit and verifier:
    ///      eid, aid, ctIdx, mode, S, deltaOrgX, deltaOrgY, threshold,
    ///      shareCount, combineHash, plaintextHash, challenge, transcriptCommitment.
    ///      For the legacy per-epoch tests we pass aid=0, ctIdx=0, mode=0, S=0,
    ///      deltaOrg=identity (matching the contract's legacy combine path).
    function decryptCombineInput(
        bytes12 epochId,
        uint16 threshold,
        uint16 shareCount,
        bytes32 combineHash,
        uint256 plaintext
    ) internal pure returns (bytes memory) {
        // Test fixture default: legacy combine path with aid=0, ctIdx=1 (the
        // only ciphertext used in DKGManagerTest). Matches the contract's
        // expectations for `combineDecryption(epoch, bytes32(0), 1, ...)`.
        return decryptCombineInputFull(
            epochId, bytes32(0), 1, 0, 0,
            0, 1,
            threshold, shareCount, combineHash, plaintext
        );
    }

    function decryptCombineInputFull(
        bytes12 epochId,
        bytes32 aid,
        uint16 ctIdx,
        uint8 mode,
        uint256 derivationS,
        uint256 deltaOrgX,
        uint256 deltaOrgY,
        uint16 threshold,
        uint16 shareCount,
        bytes32 combineHash,
        uint256 plaintext
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            DECRYPT_COMBINE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(combineHash, bytes32(plaintext), keccak256(decryptCombineTranscript(shareCount))))
        );
        uint256[13] memory v;
        v[0] = uint256(uint96(epochId));
        v[1] = uint256(aid);
        v[2] = uint256(ctIdx);
        v[3] = uint256(mode);
        v[4] = derivationS;
        v[5] = deltaOrgX;
        v[6] = deltaOrgY;
        v[7] = uint256(threshold);
        v[8] = uint256(shareCount);
        v[9] = uint256(combineHash);
        v[10] = plaintext;
        v[11] = challenge;
        v[12] = decryptCombineTranscriptCommitment(challenge, shareCount);
        return abi.encode(v);
    }

    function decryptCombineTranscript(uint16 shareCount) internal pure returns (bytes memory) {
        uint256[4] memory ciphertext;
        uint256[MAX_N] memory participantIndexes;
        uint256[2 * MAX_N] memory partialDecryptions;
        ciphertext[0] = TEST_CT_C1X;
        ciphertext[1] = TEST_CT_C1Y;
        ciphertext[2] = TEST_CT_C2X;
        ciphertext[3] = TEST_CT_C2Y;
        for (uint256 i = 0; i < MAX_N; i++) {
            partialDecryptions[i * 2 + 1] = 1;
        }
        for (uint256 i = 0; i < shareCount; i++) {
            participantIndexes[i] = i + 1;
            partialDecryptions[i * 2] = 7000 + i + 1;
            partialDecryptions[i * 2 + 1] = 8000 + i + 1;
        }
        return abi.encode(ciphertext, participantIndexes, partialDecryptions);
    }

    function decryptCombineTranscriptCommitment(uint256 challenge, uint16 shareCount) internal pure returns (uint256) {
        // Layout (4+3N words): [0..4) ciphertext, [4..4+N) participantIndexes,
        //                      [4+N..4+3N) partialDecryptions.
        uint256[] memory values = new uint256[](4 + 3 * MAX_N);
        values[0] = TEST_CT_C1X;
        values[1] = TEST_CT_C1Y;
        values[2] = TEST_CT_C2X;
        values[3] = TEST_CT_C2Y;
        uint256 partialBase = 4 + MAX_N; // partialDecryptions start
        for (uint256 i = 0; i < MAX_N; i++) {
            values[partialBase + i * 2 + 1] = 1; // y pad
        }
        uint256 cursor = 4;
        for (uint256 i = 0; i < shareCount; i++) {
            values[cursor++] = i + 1;
        }
        cursor = partialBase;
        for (uint256 i = 0; i < shareCount; i++) {
            values[cursor++] = 7000 + i + 1;
            values[cursor++] = 8000 + i + 1;
        }
        return BRLC.commit(challenge, values);
    }

}

contract MockContributionVerifier is IZKVerifier, TestInputs {
    error InvalidProofInput();

    function verifyProof(bytes calldata proof, bytes calldata input) external pure override {
        if (proof.length == 0 || keccak256(input) == keccak256(CONTRIBUTION_INPUT_BAD)) {
            revert InvalidProofInput();
        }
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return CONTRIBUTION_PROVING_KEY_HASH;
    }
}

contract MockPartialDecryptVerifier is IZKVerifier, TestInputs {
    error InvalidProofInput();

    function verifyProof(bytes calldata proof, bytes calldata input) external pure override {
        if (proof.length == 0 || keccak256(input) == keccak256(PARTIAL_DECRYPTION_INPUT_BAD)) revert InvalidProofInput();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return PARTIAL_DECRYPTION_PROVING_KEY_HASH;
    }
}

contract MockFinalizeVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return FINALIZE_PROVING_KEY_HASH;
    }
}

contract MockDecryptCombineVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return DECRYPT_COMBINE_PROVING_KEY_HASH;
    }
}
