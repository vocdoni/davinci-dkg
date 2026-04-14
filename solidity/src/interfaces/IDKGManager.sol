// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "../libraries/DKGTypes.sol";

interface IDKGManager {
    struct Round {
        address organizer;
        DKGTypes.RoundPolicy policy;
        DKGTypes.RoundStatus status;
        uint64 nonce;
        uint64 seedBlock;            // block whose blockhash becomes the lottery seed
        bytes32 seed;                // captured lazily on the first claimSlot
        uint256 lotteryThreshold;    // node-eligibility threshold (snapshotted at createRound)
        uint16 claimedCount;
        uint16 contributionCount;
        uint16 partialDecryptionCount;
        uint16 revealedShareCount;
    }

    event RoundCreated(bytes12 indexed roundId, address indexed organizer, uint64 seedBlock, uint256 lotteryThreshold);
    event RegistrationExtended(bytes12 indexed roundId, uint64 newSeedBlock, uint64 newRegistrationDeadline);
    event SeedResolved(bytes12 indexed roundId, bytes32 seed);
    event SlotClaimed(bytes12 indexed roundId, address indexed claimer, uint16 slot);
    event RegistrationClosed(bytes12 indexed roundId);
    event RoundEvicted(bytes12 indexed roundId);
    event ContributionSubmitted(
        bytes12 indexed roundId,
        address indexed contributor,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    );
    event RoundFinalized(
        bytes12 indexed roundId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    );
    event PartialDecryptionSubmitted(
        bytes12 indexed roundId,
        address indexed participant,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        bytes32 deltaHash
    );
    event DecryptionCombined(
        bytes12 indexed roundId, uint16 indexed ciphertextIndex, bytes32 combineHash, bytes32 plaintextHash
    );
    event RevealedShareSubmitted(bytes12 indexed roundId, address indexed participant, uint16 participantIndex, bytes32 shareHash);
    event SecretReconstructed(bytes12 indexed roundId, bytes32 disclosureHash, bytes32 reconstructedSecretHash);
    event RoundAborted(bytes12 indexed roundId);

    error InvalidPolicy();
    error InvalidChainId();
    error InvalidAddress();
    error InvalidRound();
    error InvalidPhase();
    error NotEligible();
    error AlreadyClaimed();
    error SlotsFull();
    error SeedNotReady();
    error SeedExpired();
    error NotRegistered();
    error NotSelectedParticipant();
    error AlreadyContributed();
    error AlreadyFinalized();
    error AlreadyPartiallyDecrypted();
    error InvalidCommitteeSize();
    error InvalidContribution();
    error InvalidFinalization();
    error InvalidPartialDecryption();
    error InsufficientContributions();
    error InvalidVerifier();
    error Unauthorized();
    error AlreadyCombined();
    error AlreadyRevealed();
    error InvalidCombinedDecryption();
    error InvalidRevealedShare();
    error InvalidReconstruction();
    error InsufficientPartialDecryptions();
    error InsufficientRevealedShares();
    error DisclosureDisabled();
    error InvalidProofInput();

    function createRound(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps,
        uint16 seedDelay,
        uint64 registrationDeadlineBlock,
        uint64 contributionDeadlineBlock,
        bool disclosureAllowed
    ) external returns (bytes12);

    function claimSlot(bytes12 roundId) external;
    function extendRegistration(bytes12 roundId) external;
    function submitContribution(
        bytes12 roundId,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function finalizeRound(
        bytes12 roundId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function submitPartialDecryption(
        bytes12 roundId,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        bytes32 deltaHash,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function combineDecryption(
        bytes12 roundId,
        uint16 ciphertextIndex,
        bytes32 combineHash,
        bytes32 plaintextHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function submitRevealedShare(
        bytes12 roundId,
        uint16 participantIndex,
        uint256 shareValue,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function reconstructSecret(
        bytes12 roundId,
        bytes32 disclosureHash,
        bytes32 reconstructedSecretHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function abortRound(bytes12 roundId) external;
    function getRound(bytes12 roundId) external view returns (Round memory);
    function selectedParticipants(bytes12 roundId) external view returns (address[] memory);
    function getContribution(bytes12 roundId, address contributor) external view returns (DKGTypes.ContributionRecord memory);
    function getPartialDecryption(bytes12 roundId, address participant, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.PartialDecryptionRecord memory);
    function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.CombinedDecryptionRecord memory);
    function getRevealedShare(bytes12 roundId, address participant)
        external
        view
        returns (DKGTypes.RevealedShareRecord memory);
    function getShareCommitmentHash(bytes12 roundId, uint16 participantIndex) external view returns (bytes32);
    function getContributionVerifierVKeyHash() external view returns (bytes32);
    function getPartialDecryptVerifierVKeyHash() external view returns (bytes32);
    function getFinalizeVerifierVKeyHash() external view returns (bytes32);
    function getDecryptCombineVerifierVKeyHash() external view returns (bytes32);
    function getRevealSubmitVerifierVKeyHash() external view returns (bytes32);
    function getRevealShareVerifierVKeyHash() external view returns (bytes32);
}
