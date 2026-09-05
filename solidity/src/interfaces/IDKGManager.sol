// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {DKGTypes} from "../libraries/DKGTypes.sol";

interface IDKGManager {
    struct Epoch {
        address organizer;
        DKGTypes.EpochPolicy policy;
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
    event CommitteeFilled(bytes12 indexed epochId);
    event ContributionSubmitted(
        bytes12 indexed epochId,
        address indexed contributor,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash
    );
    /// @notice The epoch froze its accepted contributor set and entered
    ///         Service. No key is published here — each of the epoch's
    ///         `MAX_K` pool keys is proven separately by `activatePoolKey`.
    event EpochLive(bytes12 indexed epochId, uint16 contributionCount);
    /// @notice Pool key `keyIndex` of `epochId` was proven and stored.
    ///         `(x, y)` is `P_j = Σ_d A_{d,j,0}` over the accepted
    ///         contributors; the matching share-commitment Merkle root is
    ///         readable through `getPoolShareRoot`.
    event PoolKeyActivated(bytes12 indexed epochId, uint8 indexed keyIndex, uint256 x, uint256 y);
    /// @notice Application `aid` claimed pool key `keyIndex` of `epochId`.
    event PoolKeyClaimed(bytes12 indexed epochId, bytes32 indexed aid, uint8 keyIndex);
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
    /// @dev The operator registered after the epoch was created; only the
    ///      snapshotted registry enters the lottery.
    error NotInSnapshot();
    error NotSelectedParticipant();
    error AlreadyContributed();
    error AlreadyLive();
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
    error DecryptionLimitReached();
    error CiphertextAlreadySubmitted();
    error CiphertextNotSubmitted();
    error InvalidCiphertext();
    /// @dev All `MAX_K` pool keys of the epoch have already been claimed by
    ///      applications. Register against a newer epoch.
    error PoolExhausted();
    /// @dev The pool key has not been activated yet (`activatePoolKey` never
    ///      ran for it), so it cannot be claimed or read.
    error PoolKeyNotActive();
    /// @dev A pool key may be activated exactly once per epoch.
    error PoolKeyAlreadyActive();

    /// @notice Create a new epoch. All phase deadlines are derived from
    ///         `EPOCH_DURATION_BLOCKS` (immutable, set at deploy) and the
    ///         per-phase constants in `libraries/Sizes.sol`. Permissionless:
    ///         any caller can fire it once `block.number >= nextEpochStartBlock()`,
    ///         and earlier when the newest epoch is `Live` with at most one
    ///         unclaimed pool key left, or `Aborted`.
    function createEpoch(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps
    ) external returns (bytes12);

    /// @notice Earliest block at which the next `createEpoch` may succeed on
    ///         the cadence alone. Equals `lastEpochStartBlock +
    ///         EPOCH_DURATION_BLOCKS` (or the current block before the first
    ///         epoch). A nearly spent pool or an aborted epoch short-circuits it.
    function nextEpochStartBlock() external view returns (uint64);

    /// @notice The deploy-time epoch length in blocks.
    function epochDurationBlocks() external view returns (uint256);

    function claimSlot(bytes12 epochId) external;
    /// @notice Submit a ciphertext for threshold decryption under the
    ///         application key `PK_aid`. The index is assigned on chain
    ///         (1, 2, … per application) and returned. `aid` must name a
    ///         registered application whose policy admits `msg.sender`.
    ///         Coordinates must be canonical, on-curve and non-identity;
    ///         prime-subgroup membership of C1 is checked off chain by the
    ///         committee nodes before they release a partial decryption.
    function submitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y
    ) external returns (uint16 ciphertextIndex);
    function submitContribution(
        bytes12 epochId,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    /// @notice Freeze the accepted contributor set and open the Service
    ///         window. Proof-less: the epoch key material is proven per pool
    ///         key by `activatePoolKey`.
    function finalizeEpoch(bytes12 epochId) external;
    /// @notice Prove and store one of the epoch's `MAX_K` pool keys.
    ///         Permissionless, one call per key, any order, only while Live.
    ///         `transcriptDigest` is the prover's digest of the witness
    ///         transcript (public input 5): the BRLC challenge is derived
    ///         from it and from the calldata, so the transcript is fixed
    ///         before the challenge exists.
    function activatePoolKey(
        bytes12 epochId,
        uint8 keyIndex,
        bytes32 transcriptDigest,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external;
    /// @notice Assign the next unclaimed (and already activated) pool key to
    ///         `aid`. Callable only by the sibling app manager, from
    ///         `registerApplication`.
    function claimPoolKey(bytes12 epochId, bytes32 aid) external returns (uint8 keyIndex);
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
        bytes calldata input,
        bytes32[] calldata shareProof
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
    /// @notice The activated pool key `P_j`. Reverts `PoolKeyNotActive` when
    ///         key `keyIndex` has not been proven yet.
    function getPoolKey(bytes12 epochId, uint8 keyIndex) external view returns (uint256 x, uint256 y);
    /// @notice `nextIndex` is the pool key the next registration claims;
    ///         `activated` is a `MAX_K`-wide bitmap of the proven keys.
    function getPoolStatus(bytes12 epochId) external view returns (uint8 nextIndex, uint8 activated);
    /// @notice keccak Merkle root over the `MAX_N` share commitments of pool
    ///         key `keyIndex`: leaf `i` (participant `i + 1`) is
    ///         `keccak256(0x00 ‖ D.x ‖ D.y)` for every committee member and
    ///         `MERKLE_EMPTY_LEAF` beyond `committeeSize`; internal nodes are
    ///         `keccak256(0x01 ‖ left ‖ right)`. `bytes32(0)` when not activated.
    function getPoolShareRoot(bytes12 epochId, uint8 keyIndex) external view returns (bytes32);
    /// @notice The pool key claimed by `aid`. Reverts `PoolKeyNotActive` when
    ///         the application never claimed one.
    function getAppPoolIndex(bytes12 epochId, bytes32 aid) external view returns (uint8);
    function getCiphertextHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) external view returns (bytes32);
    function ciphertextCount(bytes12 epochId, bytes32 aid) external view returns (uint16);
    function getPlaintext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) external view returns (uint256);
    function getContributionVerifierVKeyHash() external view returns (bytes32);
    function getPartialDecryptVerifierVKeyHash() external view returns (bytes32);
    function getPoolKeyVerifierVKeyHash() external view returns (bytes32);
    function getDecryptCombineVerifierVKeyHash() external view returns (bytes32);
}
