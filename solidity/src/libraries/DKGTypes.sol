// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

library DKGTypes {
    struct Point {
        uint256 x;
        uint256 y;
    }

    enum RoundStatus {
        None,
        Registration,    // accepting claimSlot calls (replaces Readiness)
        Contribution,
        Finalized,
        Aborted,
        Completed
    }

    struct RoundPolicy {
        uint16 threshold;
        uint16 committeeSize;
        uint16 minValidContributions;
        uint16 lotteryAlphaBps;            // candidate-pool size = α × committeeSize, α expressed in basis points (10000 = 1.0)
        uint16 seedDelay;                  // blocks between createRound and the block whose hash becomes the seed
        uint64 registrationDeadlineBlock;  // last block in which claimSlot is accepted
        uint64 contributionDeadlineBlock;
        bool disclosureAllowed;
    }

    struct ContributionRecord {
        address contributor;
        uint16 contributorIndex;
        bytes32 commitmentsHash;
        bytes32 encryptedSharesHash;
        bytes32 commitmentVectorDigest;
        bool accepted;
    }

    struct PartialDecryptionRecord {
        address participant;
        uint16 participantIndex;
        uint16 ciphertextIndex;
        bytes32 deltaHash;
        Point delta;
        bool accepted;
    }

    struct CombinedDecryptionRecord {
        uint16 ciphertextIndex;
        bytes32 combineHash;
        bytes32 plaintextHash;
        bool completed;
    }

    struct RevealedShareRecord {
        address participant;
        uint16 participantIndex;
        uint256 shareValue;
        bytes32 shareHash;
        bool accepted;
    }
}
