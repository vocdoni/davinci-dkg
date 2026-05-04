// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "../libraries/DKGTypes.sol";

interface IDKGManager {
    struct Epoch {
        address organizer;
        DKGTypes.EpochPolicy policy;
        DKGTypes.DecryptionPolicy decryptionPolicy;
        DKGTypes.EpochPhase status;
        uint64 nonce;
        uint64 startBlock;           // block.number at createEpoch — anchor for nextEpochStartBlock
        uint64 seedBlock;            // block whose blockhash becomes the lottery seed
        bytes32 seed;                // captured lazily on the first claimSlot
        uint256 lotteryThreshold;    // node-eligibility threshold (snapshotted at createEpoch)
        uint16 claimedCount;
        uint16 contributionCount;
        uint16 partialDecryptionCount;
        uint16 ciphertextCount;      // number of submitCiphertext calls accepted for this epoch
    }

    event EpochCreated(bytes12 indexed epochId, address indexed organizer, uint64 startBlock, uint64 seedBlock, uint256 lotteryThreshold);
    event SeedResolved(bytes12 indexed epochId, bytes32 seed);
    event SlotClaimed(bytes12 indexed epochId, address indexed claimer, uint16 slot);
    event RegistrationClosed(bytes12 indexed epochId);
    event EpochEvicted(bytes12 indexed epochId);
    event ContributionSubmitted(
        bytes12 indexed epochId,
        address indexed contributor,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    );
    event EpochFinalized(
        bytes12 indexed epochId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash
    );
    event PartialDecryptionSubmitted(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        address indexed participant,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        uint256 deltaX,
        uint256 deltaY
    );
    event CiphertextSubmitted(
        bytes12 indexed epochId,
        bytes32 indexed aid,
        uint16 indexed ciphertextIndex,
        address submitter,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y
    );
    event DecryptionCombined(
        bytes12 indexed epochId, bytes32 indexed aid, uint16 indexed ciphertextIndex, bytes32 combineHash, uint256 plaintext
    );
    event EpochAborted(bytes12 indexed epochId);

    error InvalidPolicy();
    error InvalidChainId();
    error InvalidAddress();
    error InvalidEpoch();
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
    error InvalidCombinedDecryption();
    error InsufficientPartialDecryptions();
    error InvalidProofInput();
    error NotOwner();
    error InvalidDecryptionPolicy();
    error DecryptionNotYetAllowed();
    error DecryptionExpired();
    error DecryptionLimitReached();
    error CiphertextAlreadySubmitted();
    error CiphertextNotSubmitted();
    error InvalidCiphertext();

    /// @notice Create a new epoch. All phase deadlines are derived from
    ///         `EPOCH_DURATION_BLOCKS` (immutable, set at deploy) and the
    ///         per-phase BPS constants in `libraries/Sizes.sol`. Permissionless:
    ///         any caller can fire it once `block.number >= nextEpochStartBlock()`.
    function createEpoch(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps,
        DKGTypes.DecryptionPolicy calldata decryptionPolicy
    ) external returns (bytes12);

    /// @notice Earliest block at which the next `createEpoch` may succeed.
    ///         Equals `lastEpochStartBlock + EPOCH_DURATION_BLOCKS` (or 0
    ///         before the first epoch, meaning "any block").
    function nextEpochStartBlock() external view returns (uint64);

    /// @notice The deploy-time epoch length in blocks.
    function epochDurationBlocks() external view returns (uint256);

    function claimSlot(bytes12 epochId) external;
    function submitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y
    ) external;
    function submitContribution(
        bytes12 epochId,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function finalizeEpoch(
        bytes12 epochId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function submitPartialDecryption(
        bytes12 epochId,
        bytes32 aid,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y,
        bytes32 deltaHash,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function combineDecryption(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        bytes32 combineHash,
        uint256 plaintext,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    function abortEpoch(bytes12 epochId) external;
    function getEpoch(bytes12 epochId) external view returns (Epoch memory);
    function selectedParticipants(bytes12 epochId) external view returns (address[] memory);
    function getContribution(bytes12 epochId, address contributor) external view returns (DKGTypes.ContributionRecord memory);
    function getPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.PartialDecryptionRecord memory);
    function getCombinedDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.CombinedDecryptionRecord memory);
    function getShareCommitmentHash(bytes12 epochId, uint16 participantIndex) external view returns (bytes32);
    function getCollectivePublicKey(bytes12 epochId) external view returns (DKGTypes.Point memory);
    function getCiphertextHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) external view returns (bytes32);
    function getPlaintext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) external view returns (uint256);
    function getDecryptionPolicy(bytes12 epochId) external view returns (DKGTypes.DecryptionPolicy memory);
    function getContributionVerifierVKeyHash() external view returns (bytes32);
    function getPartialDecryptVerifierVKeyHash() external view returns (bytes32);
    function getFinalizeVerifierVKeyHash() external view returns (bytes32);
    function getDecryptCombineVerifierVKeyHash() external view returns (bytes32);
}
