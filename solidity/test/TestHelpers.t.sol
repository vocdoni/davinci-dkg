// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../src/interfaces/IZKVerifier.sol";
import {BRLC} from "../src/libraries/BRLC.sol";
import {TestInputs} from "./TestInputs.t.sol";

abstract contract TestHelpers is TestInputs {
    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:decrypt-combine:v1");
    bytes32 internal constant FINALIZE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:finalize:v1");
    bytes32 internal constant REVEAL_SHARE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:reveal-share:v1");

    function contributionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(1), 2, 3, 4, 5, 6, 7, 8]);
    }

    function contributionInput(
        bytes12 roundId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            roundId,
            CONTRIBUTION_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(commitmentsHash, encryptedSharesHash))
        );
        return abi.encode(
            [
                uint256(uint96(roundId)),
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

    // The test fixtures below mirror the on-chain transcript layouts at N=16:
    //   contribution: 8N = 128 words = (2N + N + 2N + 2N + N) layout.
    //   finalize:     2N²+5N = 592 words.
    //   combine:      4 + 3N = 52 words.
    //   reconstruct:  2N = 32 words.
    function contributionTranscript(uint16 committeeSize) internal pure returns (bytes memory) {
        uint256[32] memory commitments;       // 2N
        uint256[16] memory recipientIndexes;  // N
        uint256[32] memory recipientPubKeys;  // 2N
        uint256[32] memory ephemerals;        // 2N
        uint256[16] memory maskedShares;      // N
        for (uint256 i = 0; i < 16; i++) {
            commitments[i * 2 + 1] = 1;
            recipientPubKeys[i * 2 + 1] = 1;
            ephemerals[i * 2 + 1] = 1;
        }
        for (uint256 i = 0; i < committeeSize; i++) {
            commitments[i * 2 + 1] = 0;
            recipientIndexes[i] = i + 1;
            recipientPubKeys[i * 2] = 100 + i + 1;
            recipientPubKeys[i * 2 + 1] = 200 + i + 1;
            ephemerals[i * 2] = 300 + i + 1;
            ephemerals[i * 2 + 1] = 400 + i + 1;
            maskedShares[i] = 500 + i + 1;
        }
        return abi.encode(commitments, recipientIndexes, recipientPubKeys, ephemerals, maskedShares);
    }

    function contributionTranscriptCommitment(uint256 challenge, uint16 committeeSize) internal pure returns (uint256) {
        uint256[] memory values = new uint256[](128); // 8N
        for (uint256 i = 0; i < 16; i++) {
            values[i * 2 + 1] = 1;        // commitments y pad
            values[48 + i * 2 + 1] = 1;   // recipientPubKeys y pad (offset 2N+N=48)
            values[80 + i * 2 + 1] = 1;   // ephemerals y pad (offset 2N+N+2N=80)
        }
        uint256 cursor = 32; // recipientIndexes start (after 2N commitments)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[i * 2 + 1] = 0;
            values[cursor++] = i + 1;
        }
        cursor = 48; // recipientPubKeys start (2N+N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 100 + i + 1;
            values[cursor++] = 200 + i + 1;
        }
        cursor = 80; // ephemerals start (2N+N+2N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 300 + i + 1;
            values[cursor++] = 400 + i + 1;
        }
        cursor = 112; // maskedShares start (2N+N+2N+2N)
        for (uint256 i = 0; i < committeeSize; i++) {
            values[cursor++] = 500 + i + 1;
        }
        return BRLC.commit(challenge, values);
    }

    function partialDecryptionProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(11), 12, 13, 14, 15, 16, 17, 18]);
    }

    function partialDecryptionInput(bytes12 roundId, uint16 participantIndex, bytes32)
        internal
        pure
        returns (bytes memory)
    {
        uint256[13] memory inputs;
        inputs[0] = uint256(uint96(roundId));
        inputs[1] = participantIndex;
        inputs[4] = 1000 + participantIndex;
        inputs[5] = 2000 + participantIndex;
        inputs[6] = 7000 + participantIndex;
        inputs[7] = 8000 + participantIndex;
        return abi.encode(inputs);
    }

    function partialDecryptionHash(uint16 participantIndex) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(uint256(7000 + participantIndex), uint256(8000 + participantIndex)));
    }

    function finalizeProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(21), 22, 23, 24, 25, 26, 27, 28]);
    }

    function finalizeInput(
        bytes12 roundId,
        uint16 threshold,
        uint16 committeeSize,
        uint16 acceptedCount,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            roundId, FINALIZE_TRANSCRIPT_DOMAIN, keccak256(abi.encodePacked(aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash))
        );
        return abi.encode(
            [
                uint256(uint96(roundId)),
                uint256(threshold),
                uint256(committeeSize),
                uint256(acceptedCount),
                uint256(aggregateCommitmentsHash),
                uint256(collectivePublicKeyHash),
                uint256(shareCommitmentHash),
                challenge,
                finalizeTranscriptCommitment(challenge, acceptedCount)
            ]
        );
    }

    function finalizeTranscript(uint16 acceptedCount) internal pure returns (bytes memory) {
        uint256[16] memory participantIndexes;        // N
        uint256[512] memory contributionCommitments;  // 2N²
        uint256[32] memory aggregateCommitments;      // 2N
        uint256[32] memory shareCommitments;          // 2N
        // Per participant, mirror contributionTranscript(committeeSize=2):
        // pt[0]=pt[1]=(0,0); pt[2..15]=(0,1) → odd indices 5,7,...,31 = 1
        // (each participant's commitments occupy 2N=32 words).
        for (uint256 i = 0; i < acceptedCount; i++) {
            participantIndexes[i] = i + 1;
            for (uint256 k = 5; k < 32; k += 2) {
                contributionCommitments[i * 32 + k] = 1;
            }
            shareCommitments[i * 2] = 1000 + i + 1;
            shareCommitments[i * 2 + 1] = 2000 + i + 1;
        }
        return abi.encode(participantIndexes, contributionCommitments, aggregateCommitments, shareCommitments);
    }

    function finalizeTranscriptCommitment(uint256 challenge, uint16 acceptedCount) internal pure returns (uint256) {
        // Layout (2N²+5N words): [0..N) participantIndexes,
        //                        [N..N+2N²) contributionCommitments,
        //                        [N+2N²..N+2N²+2N) aggregateCommitments,
        //                        [N+2N²+2N..2N²+5N) shareCommitments.
        uint256[] memory values = new uint256[](592); // 2*16*16 + 5*16 = 592
        for (uint256 i = 0; i < acceptedCount; i++) {
            values[i] = i + 1;
            uint256 offset = 16 + i * 32; // N + i*2N
            for (uint256 k = 5; k < 32; k += 2) {
                values[offset + k] = 1;
            }
        }
        for (uint256 i = 0; i < acceptedCount; i++) {
            uint256 offset = 560 + i * 2; // N + 2N² + 2N = 560 (shareCommitments start)
            values[offset] = 1000 + i + 1;
            values[offset + 1] = 2000 + i + 1;
        }
        return BRLC.commit(challenge, values);
    }

    function decryptCombineProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(31), 32, 33, 34, 35, 36, 37, 38]);
    }

    function decryptCombineInput(
        bytes12 roundId,
        uint16 threshold,
        uint16 shareCount,
        bytes32 combineHash,
        bytes32 plaintextHash
    ) internal pure returns (bytes memory) {
        uint256 challenge = BRLC.deriveChallenge(
            roundId,
            DECRYPT_COMBINE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(combineHash, plaintextHash))
        );
        return abi.encode(
            [
                uint256(uint96(roundId)),
                uint256(threshold),
                uint256(shareCount),
                uint256(combineHash),
                uint256(plaintextHash),
                challenge,
                decryptCombineTranscriptCommitment(challenge, shareCount)
            ]
        );
    }

    function decryptCombineTranscript(uint16 shareCount) internal pure returns (bytes memory) {
        uint256[4] memory ciphertext;
        uint256[16] memory participantIndexes;     // N
        uint256[32] memory partialDecryptions;     // 2N
        ciphertext[0] = 7001;
        ciphertext[1] = 8001;
        ciphertext[2] = 9001;
        ciphertext[3] = 10001;
        for (uint256 i = 0; i < 16; i++) {
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
        uint256[] memory values = new uint256[](52); // 4 + 3*16
        values[0] = 7001;
        values[1] = 8001;
        values[2] = 9001;
        values[3] = 10001;
        for (uint256 i = 0; i < 16; i++) {
            values[20 + i * 2 + 1] = 1; // partialDecryptions y pad (offset 4+N=20)
        }
        uint256 cursor = 4;
        for (uint256 i = 0; i < shareCount; i++) {
            values[cursor++] = i + 1;
        }
        cursor = 20; // partialDecryptions start (4+N)
        for (uint256 i = 0; i < shareCount; i++) {
            values[cursor++] = 7000 + i + 1;
            values[cursor++] = 8000 + i + 1;
        }
        return BRLC.commit(challenge, values);
    }

    function revealShareProof() internal pure returns (bytes memory) {
        return abi.encode([uint256(41), 42, 43, 44, 45, 46, 47, 48]);
    }

    function revealSubmitInput(bytes12 roundId, uint16 participantIndex, uint256 shareValue)
        internal
        pure
        returns (bytes memory)
    {
        return abi.encode([uint256(uint96(roundId)), uint256(participantIndex), shareValue, 1000 + participantIndex, 2000 + participantIndex]);
    }

    function revealShareInput(
        bytes12 roundId,
        uint16 threshold,
        uint16 shareCount,
        bytes32 disclosureHash,
        bytes32 reconstructedSecretHash
    )
        internal
        pure
        returns (bytes memory)
    {
        uint256 challenge = BRLC.deriveChallenge(
            roundId,
            REVEAL_SHARE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(disclosureHash, reconstructedSecretHash))
        );
        return abi.encode(
            [
                uint256(uint96(roundId)),
                uint256(threshold),
                uint256(shareCount),
                uint256(disclosureHash),
                uint256(reconstructedSecretHash),
                challenge,
                revealShareTranscriptCommitment(challenge, shareCount)
            ]
        );
    }

    function revealShareTranscript(uint16 shareCount) internal pure returns (bytes memory) {
        uint256[16] memory participantIndexes; // N
        uint256[16] memory shares;             // N
        for (uint256 i = 0; i < shareCount; i++) {
            participantIndexes[i] = i + 1;
            shares[i] = uint256(REVEALED_SHARE_HASH) + i;
        }
        return abi.encode(participantIndexes, shares);
    }

    function revealShareTranscriptCommitment(uint256 challenge, uint16 shareCount) internal pure returns (uint256) {
        uint256[] memory values = new uint256[](32); // 2N
        for (uint256 i = 0; i < shareCount; i++) {
            values[i] = i + 1;
            values[16 + i] = uint256(REVEALED_SHARE_HASH) + i; // shares start at N
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

contract MockRevealSubmitVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return REVEAL_SUBMIT_PROVING_KEY_HASH;
    }
}

contract MockRevealShareVerifier is IZKVerifier, TestInputs {
    function verifyProof(bytes calldata proof, bytes calldata) external pure override {
        if (proof.length == 0) revert();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return REVEAL_SHARE_PROVING_KEY_HASH;
    }
}
