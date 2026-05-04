// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {IDKGAppManager} from "./interfaces/IDKGAppManager.sol";
import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGIdLib} from "./libraries/DKGIdLib.sol";
import {BRLC} from "./libraries/BRLC.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";
import {PhaseLib} from "./libraries/PhaseLib.sol";
import {
    MAX_N,
    DEFAULT_EPOCH_DURATION_BLOCKS,
    DEFAULT_COMMITTEE_SELECTION_BLOCKS,
    DEFAULT_KEY_ASSEMBLY_BLOCKS,
    DEFAULT_FINALIZE_GAP_BLOCKS,
    SEED_DELAY_BLOCKS
} from "./libraries/Sizes.sol";

/// @title  DKGManager
/// @notice On-chain orchestrator for every phase of a davinci-dkg epoch.
/// @dev    Lifecycle: Registration (trustless lottery) → Contribution →
///         Finalized → Completed (or Aborted). Every state-mutating entry
///         point that makes a cryptographic claim is gated by a Groth16
///         verifier — no dispute phase, no complaint flow. Historic epoch
///         storage is bounded by a ring buffer of EPOCH_HISTORY_SIZE (64)
///         entries; evicted epochs remain reconstructible from event logs.
///         The share-commitment list is stored as `keccak256(x, y)` per
///         participant (1 SSTORE instead of 2) and the transcripts used by
///         finalize/combine are read straight out of calldata via assembly
///         to avoid per-element bounds checks.
contract DKGManager is IDKGManager {
    /// @dev Sibling app manager has not been wired yet — applications cannot
    ///      yet be registered, organizer shares cannot be submitted, and any
    ///      cross-contract path (submitCiphertext with aid != 0,
    ///      combineDecryption with aid != 0) reverts. The legacy aid == 0 path
    ///      remains available before wiring.
    error AppManagerNotSet();
    error AppManagerAlreadySet();

    // ──────────────────────────────────────────────────────────────────────
    // Single source of truth for the per-circuit array bound.
    //
    // MAX_N is imported from `libraries/Sizes.sol` — the single source of
    // truth on the on-chain side. It must agree with `circuits/common.MaxN`
    // on the Go side. Changing it requires recompiling every circuit,
    // regenerating the proving keys, and redeploying the verifier wrappers.
    // The Foundry test helpers in `test/TestHelpers.t.sol` also import the
    // same constant so the per-N gas tables can be regenerated without code
    // edits beyond the single line in `Sizes.sol`.
    // ──────────────────────────────────────────────────────────────────────
    uint256 internal constant MAX_COEFFICIENTS = MAX_N;
    uint256 internal constant MAX_RECIPIENTS   = MAX_N;
    uint256 internal constant MAX_PARTICIPANTS = MAX_N;
    uint256 internal constant MAX_SHARES       = MAX_N;

    // Derived transcript word counts (1 word = 32 bytes).
    //
    //   submitContribution: commitmentPoints (2N) ‖ recipientIndexes (N) ‖
    //                       recipientPubKeys (2N) ‖ ephemerals (2N) ‖
    //                       maskedShares (N)                              = 8N
    //   finalizeEpoch:      participantIndexes (N) ‖
    //                       contributionCommitments (2N²) ‖
    //                       aggregateCommitments (2N) ‖
    //                       shareCommitments (2N)                         = 2N² + 5N
    //   combineDecryption:  ciphertext (4) ‖ participantIndexes (N) ‖
    //                       partialDecryptions (2N)                       = 4 + 3N
    uint256 internal constant CONTRIB_TRANSCRIPT_WORDS     = 8 * MAX_N;
    uint256 internal constant FINALIZE_TRANSCRIPT_WORDS    = 2 * MAX_N * MAX_N + 5 * MAX_N;
    uint256 internal constant COMBINE_TRANSCRIPT_WORDS     = 4 + 3 * MAX_N;
    // contribution-time per-section byte offsets:
    uint256 internal constant CONTRIB_PUBKEYS_BYTES_OFFSET = (2 * MAX_N + MAX_N) * 32;          // start of recipientPubKeys
    uint256 internal constant CONTRIB_PUBKEYS_BYTES_END    = (2 * MAX_N + MAX_N + 2 * MAX_N) * 32; // end of recipientPubKeys
    uint256 internal constant CONTRIB_DIGEST_BYTES_LEN     = 2 * MAX_N * 32;                    // first 2N words = commitmentPoints
    // finalize-time per-section byte offsets:
    uint256 internal constant FINALIZE_CONTRIB_BYTES_OFFSET = MAX_N * 32;                       // participantIndexes end
    uint256 internal constant FINALIZE_CONTRIB_BYTES_LEN    = 2 * MAX_N * MAX_N * 32;           // contributionCommitments length in bytes
    uint256 internal constant FINALIZE_PER_CONTRIB_BYTES    = 2 * MAX_N * 32;                   // bytes per contributor's commitments slice
    uint256 internal constant FINALIZE_SHARE_WORDS_OFFSET   = MAX_N + 2 * MAX_N * MAX_N + 2 * MAX_N; // shareCommitments start, in words
    uint256 internal constant FINALIZE_AGG0_WORDS_OFFSET    = MAX_N + 2 * MAX_N * MAX_N;             // aggregateCommitments[0] start, in words
    // combine-time per-section byte offsets:
    uint256 internal constant COMBINE_PARTIALS_BYTES_OFFSET = (4 + MAX_N) * 32;                 // partialDecryptions start, in bytes

    /// @dev Number of recent epoch IDs retained on-chain. After this many `createEpoch`
    /// calls, the oldest live epoch's storage is evicted (its data is wiped) and only
    /// the event log retains it. Tunable; 64 is large enough to cover several days of
    /// epochs at typical cadences.
    uint256 internal constant EPOCH_HISTORY_SIZE = 64;

    /// @dev Upper bound on ciphertext indices accepted by `submitPartialDecryption`
    /// and `combineDecryption`. Prevents unbounded storage spam by a committee member
    /// who submits decryptions for arbitrarily large ciphertext indices.
    uint16 internal constant MAX_CIPHERTEXT_INDEX = 256;

    uint32 public immutable CHAIN_ID;
    address public immutable REGISTRY;
    uint32 public immutable EPOCH_PREFIX;
    address public immutable CONTRIBUTION_VERIFIER;
    address public immutable PARTIAL_DECRYPT_VERIFIER;
    address public immutable FINALIZE_VERIFIER;
    address public immutable DECRYPT_COMBINE_VERIFIER;
    /// @notice Total epoch length in blocks. Set per-deploy by the constructor.
    ///         Defaults to `DEFAULT_EPOCH_DURATION_BLOCKS` (Sizes.sol) when the
    ///         constructor argument is 0. Wall-clock duration depends on the
    ///         deployment chain's block time and is estimated off-chain.
    uint256 public immutable EPOCH_DURATION_BLOCKS;
    /// @dev Phase deadline offsets, derived from EPOCH_DURATION_BLOCKS at
    ///      construction so we don't pay the BPS division on every createEpoch.
    uint64 internal immutable COMMITTEE_SELECTION_DEADLINE_OFFSET;
    uint64 internal immutable KEY_ASSEMBLY_DEADLINE_OFFSET;
    uint64 internal immutable LIVE_NOT_BEFORE_OFFSET;
    uint64 public epochNonce;
    /// @notice Block at which the most recent epoch was created. Anchor for
    ///         `nextEpochStartBlock()`.
    uint64 public lastEpochStartBlock;

    /// @notice The sibling DKGAppManager that owns the per-application
    ///         registration / organizer-share storage and verification logic.
    ///         Set exactly once via `setAppManager` by the deployer (the
    ///         constructor cannot take it because the app manager's
    ///         constructor itself takes this contract's address — cyclic
    ///         dependency resolved with a one-shot setter).
    address public appManager;
    bool private _appManagerSet;
    /// @dev Address that deployed this contract; the only address allowed to
    ///      call `setAppManager`. Immutable after construction.
    address private immutable _deployer;

    /// @dev Fixed-size ring buffer of recent epoch IDs. New epochs push here at
    /// createEpoch; once the buffer is full, the displaced entry tells us which epoch
    /// to evict. `recentEpochsCount` counts total pushes; current write index is
    /// `recentEpochsCount % EPOCH_HISTORY_SIZE`.
    bytes12[EPOCH_HISTORY_SIZE] internal recentEpochs;
    uint64 internal recentEpochsCount;

    mapping(bytes12 epochId => Epoch epoch) internal epochs;
    mapping(bytes12 epochId => mapping(address operator => bool selected)) internal selectedOperators;
    mapping(bytes12 epochId => address[] participants) internal epochParticipants;
    mapping(bytes12 epochId => mapping(address contributor => DKGTypes.ContributionRecord contribution)) internal
        epochContributions;
    // Ciphertexts, committee partials, partial counts, and combined plaintexts
    // are all keyed by (epochId, aid, ctIdx). The aid namespace prevents two
    // applications under the same epoch from colliding on ctIdx (one app
    // blocking another, partials shared/rejected across apps, a combine for
    // one aid marking the slot completed for all aids). The legacy per-epoch
    // path (no application registered) uses `aid = bytes32(0)`.
    //
    // Partial decryptions are stored as `keccak256(δ.x, δ.y)` only — the raw
    // δ point is read back from the combine transcript at combine time and
    // re-hashed against the stored value. A per-(epochId, aid, ctIdx) bitmap
    // tracks which participantIndexes have submitted: bit (1 << pIdx) is set
    // iff participant pIdx has submitted. Saves two cold SSTOREs per partial
    // decryption versus storing the full δ point + a bool, and lets the
    // combine path key by participantIndex directly without an address
    // lookup.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => mapping(uint16 participantIndex => bytes32 deltaHash)))) internal epochPartialDeltaHash;
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => uint256 bitmap))) internal epochPartialBitmap;
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => uint16 count))) internal epochPartialDecryptionCounts;
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => DKGTypes.CombinedDecryptionRecord combined))) internal
        epochCombinedDecryptions;
    /// @dev Stores _hash2(scX, scY) for each share commitment, packing
    /// the original (x,y) pair into a single 32-byte slot. Saves one cold SSTORE per
    /// committee member at finalize time. The pre-image (x,y) is exposed in the
    /// EpochLive event for off-chain consumers.
    mapping(bytes12 epochId => mapping(uint16 participantIndex => bytes32 shareCommitmentHash)) internal epochShareCommitmentHashes;

    /// @dev Stores keccak256 over the canonical (recipientIndexes ‖ recipientPubKeys)
    /// section of any valid submitContribution transcript for this epoch. Set once at
    /// selectParticipants. Lets submitContribution verify the entire 96-word committee
    /// section in one keccak instead of 32 storage reads + 32 external registry calls.
    mapping(bytes12 epochId => bytes32 prefixHash) internal epochContribPrefixHash;

    /// @dev Accumulates the collective public key on-chain as contributions are submitted.
    ///      Each accepted contribution adds its commitment[0] point (a_{i,0}·G) to the
    ///      running sum. The identity element (0,1) is the initial value. Once the epoch
    ///      is finalized the value equals sum_i(a_{i,0}·G) = the collective public key.
    mapping(bytes12 epochId => DKGTypes.Point) internal _collectiveKey;

    /// @dev _hash4(c1x, c1y, c2x, c2y) for each ciphertext submitted to a
    ///      epoch. Written once per (epochId, aid, ciphertextIndex) by submitCiphertext and
    ///      verified by combineDecryption to bind the combine proof to the authoritative
    ///      on-chain ciphertext (preventing a combiner from swapping in a different ct).
    ///      The raw coordinates are available via the CiphertextSubmitted event log.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => bytes32 ciphertextHash))) internal _ciphertexts;

    // BabyJubJub curve constants moved to libraries/BabyJubJub.sol — point
    // validation flows through `_requireValidEncryptionPoint` which calls
    // `BabyJubJub.{isCanonical,isOnCurve,isIdentity,isInPrimeSubgroup}`.

    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:decrypt-combine:v1");
    bytes32 internal constant FINALIZE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:finalize:v1");

    constructor(
        uint32 _chainId,
        address _registry,
        address _contributionVerifier,
        address _partialDecryptVerifier,
        address _finalizeVerifier,
        address _decryptCombineVerifier,
        uint256 _epochDurationBlocks,
        uint256 _committeeSelectionBlocks,
        uint256 _keyAssemblyBlocks,
        uint256 _finalizeGapBlocks
    ) {
        if (uint32(block.chainid) != _chainId) revert InvalidChainId();
        if (_registry == address(0)) revert InvalidAddress();
        if (
            _contributionVerifier == address(0) || _partialDecryptVerifier == address(0) || _finalizeVerifier == address(0)
                || _decryptCombineVerifier == address(0)
        ) revert InvalidVerifier();
        CHAIN_ID = _chainId;
        REGISTRY = _registry;
        EPOCH_PREFIX = DKGIdLib.getPrefix(_chainId, address(this));
        CONTRIBUTION_VERIFIER = _contributionVerifier;
        PARTIAL_DECRYPT_VERIFIER = _partialDecryptVerifier;
        FINALIZE_VERIFIER = _finalizeVerifier;
        DECRYPT_COMBINE_VERIFIER = _decryptCombineVerifier;

        uint256 dur     = _epochDurationBlocks      == 0 ? DEFAULT_EPOCH_DURATION_BLOCKS      : _epochDurationBlocks;
        uint256 csl     = _committeeSelectionBlocks == 0 ? DEFAULT_COMMITTEE_SELECTION_BLOCKS : _committeeSelectionBlocks;
        uint256 keyAsm  = _keyAssemblyBlocks        == 0 ? DEFAULT_KEY_ASSEMBLY_BLOCKS        : _keyAssemblyBlocks;
        uint256 finGap  = _finalizeGapBlocks        == 0 ? DEFAULT_FINALIZE_GAP_BLOCKS        : _finalizeGapBlocks;

        // Sanity:
        //   - the seed-revelation block (startBlock + SEED_DELAY_BLOCKS) must
        //     land strictly inside CommitteeSelection so claimers get at
        //     least one block after the seed resolves;
        //   - every phase must fit inside EPOCH_DURATION with the Service
        //     window left non-empty;
        //   - all four constants must fit in uint64 since the per-epoch
        //     deadline blocks are stored as uint64.
        uint256 prep = csl + keyAsm + finGap;
        if (
            csl <= SEED_DELAY_BLOCKS
                || keyAsm == 0
                || finGap == 0
                || prep >= dur
                || dur > type(uint64).max
        ) {
            revert InvalidPolicy();
        }

        EPOCH_DURATION_BLOCKS               = dur;
        COMMITTEE_SELECTION_DEADLINE_OFFSET = uint64(csl);
        KEY_ASSEMBLY_DEADLINE_OFFSET        = uint64(csl + keyAsm);
        LIVE_NOT_BEFORE_OFFSET              = uint64(prep);

        _deployer = msg.sender;
    }

    /// @inheritdoc IDKGManager
    function epochDurationBlocks() external view returns (uint256) {
        return EPOCH_DURATION_BLOCKS;
    }

    /// @inheritdoc IDKGManager
    /// @dev Returns the earliest block at which `createEpoch` will succeed.
    ///      For the very first epoch (nothing scheduled yet) this is the
    ///      current block, so the call goes through immediately. Otherwise
    ///      `lastEpochStartBlock + EPOCH_DURATION_BLOCKS` — the cadence
    ///      anchor.
    function nextEpochStartBlock() public view returns (uint64) {
        if (lastEpochStartBlock == 0) return uint64(block.number);
        return lastEpochStartBlock + uint64(EPOCH_DURATION_BLOCKS);
    }

    /// @notice Pin the sibling DKGAppManager address. May only be called once
    ///         by the deployer; subsequent calls revert. Mirrors
    ///         DKGRegistry.setManager. Resolves the cyclic constructor
    ///         dependency between the two contracts.
    function setAppManager(address a) external {
        if (msg.sender != _deployer) revert Unauthorized();
        if (_appManagerSet) revert AppManagerAlreadySet();
        if (a == address(0)) revert InvalidAddress();
        appManager = a;
        _appManagerSet = true;
    }

    /// @notice Create a new DKG epoch. Permissionless: any caller may fire it
    ///         once `block.number >= nextEpochStartBlock()`. All phase
    ///         deadlines are derived from `EPOCH_DURATION_BLOCKS` and the BPS
    ///         constants in `Sizes.sol`.
    /// @dev    The caller becomes the epoch organizer (recorded for
    ///         organizational accounting; the protocol does not grant the
    ///         organizer any special privileges over committee selection,
    ///         which is trustlessly drawn by the lottery).
    /// @param  threshold              Shamir reconstruction threshold `t`.
    /// @param  committeeSize          Target committee size `n` ∈ [1, MAX_N].
    /// @param  minValidContributions  Minimum accepted contributions required
    ///                                for `finalizeEpoch`. Must be ≥ threshold.
    /// @param  lotteryAlphaBps        Oversubscription factor α in basis points
    ///                                (10000 = 1.0). Expected eligible set
    ///                                size is `α · committeeSize`.
    /// @param  decryptionPolicy       Per-epoch ciphertext-submission policy
    ///                                (gates the legacy aid=0 path).
    /// @return                        The 12-byte epoch identifier
    ///                                `uint32 prefix || uint64 nonce`.
    function createEpoch(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps,
        DKGTypes.DecryptionPolicy calldata decryptionPolicy
    ) external returns (bytes12) {
        // Permissionless cadence guard. The first epoch (lastEpochStartBlock
        // == 0) goes through immediately; subsequent epochs require one full
        // EPOCH_DURATION_BLOCKS to have elapsed since the previous start.
        if (block.number < nextEpochStartBlock()) revert InvalidPhase();

        if (
            threshold == 0 || committeeSize == 0 || threshold > committeeSize
                || committeeSize > MAX_N
                || minValidContributions == 0 || minValidContributions > committeeSize
                || minValidContributions < threshold
                || lotteryAlphaBps < 10000
        ) revert InvalidPolicy();
        // Two invariants worth calling out:
        //   - minValidContributions < threshold would let the epoch finalize
        //     with fewer share holders than the Shamir threshold — the
        //     resulting key would be unrecoverable.
        //   - committeeSize > MAX_N would pass createEpoch but silently
        //     break later: every per-epoch proof (contribution, finalize,
        //     combine, reveal) assumes fixed-width transcripts of size
        //     `MAX_N`. The contract's committee storage and event payloads
        //     would also overflow their hashed-prefix layout. Cap upfront
        //     so the failure is the obvious revert here, not a confusing
        //     ProofInvalid() three transactions deep.

        // DecryptionPolicy sanity: if both directions of the same clock are set,
        // the window must be non-empty. maxDecryptions is capped at MAX_CIPHERTEXT_INDEX.
        if (
            (decryptionPolicy.notBeforeBlock != 0 && decryptionPolicy.notAfterBlock != 0
                && decryptionPolicy.notAfterBlock <= decryptionPolicy.notBeforeBlock)
                || (decryptionPolicy.notBeforeTimestamp != 0 && decryptionPolicy.notAfterTimestamp != 0
                    && decryptionPolicy.notAfterTimestamp <= decryptionPolicy.notBeforeTimestamp)
                || decryptionPolicy.maxDecryptions > MAX_CIPHERTEXT_INDEX
        ) revert InvalidDecryptionPolicy();

        // Snapshot the currently ACTIVE node count and derive the per-node lottery
        // threshold so that on average `lotteryAlpha × committeeSize` live nodes pass.
        // Using activeCount (rather than nodeCount) keeps the lottery denominator
        // aligned with the set of nodes that can actually produce a contribution —
        // reaped stragglers are automatically excluded.
        uint256 registered = uint256(IDKGRegistry(REGISTRY).activeCount());
        if (registered == 0) revert InvalidPolicy();
        // numerator = α × committeeSize (in basis points domain); 10000 = α × 1.0
        // expectedPass = registered × (numerator / 10000)
        // threshold = floor(2^256 × expectedPass / registered)
        //           = floor(2^256 × numerator / 10000)         when registered > expectedPass
        // We cap the threshold at type(uint256).max - 1 so the comparison is strict.
        uint256 numerator = uint256(lotteryAlphaBps) * uint256(committeeSize);
        // expected = registered × numerator / 10000; if expected ≥ registered,
        // every node passes (threshold = max). Otherwise compute proportional.
        uint256 lotteryThreshold;
        if (numerator >= 10000) {
            // α × committeeSize ≥ registered: everyone passes
            lotteryThreshold = type(uint256).max;
        } else {
            // threshold = (2^256 × numerator) / (10000 × registered) ; use mulDiv-style
            // safe expansion. Since numerator/10000 ≤ committeeSize, and we're scaling
            // to 2^256, a simple shift suffices: shift by 256 then divide.
            // Equivalent: (uint256.max / registered) × (numerator / 10000), avoiding overflow.
            uint256 fraction = (type(uint256).max / 10000) * numerator; // ≤ uint256.max
            lotteryThreshold = fraction / registered;
        }

        epochNonce++;
        bytes12 epochId = DKGIdLib.computeEpochId(EPOCH_PREFIX, epochNonce);

        // Evict the oldest live epoch if the history buffer is full.
        uint256 writeSlot = uint256(recentEpochsCount) % EPOCH_HISTORY_SIZE;
        if (recentEpochsCount >= EPOCH_HISTORY_SIZE) {
            bytes12 evictedKey = recentEpochs[writeSlot];
            if (evictedKey != bytes12(0)) {
                _evictRound(evictedKey);
            }
        }
        recentEpochs[writeSlot] = epochId;
        unchecked { recentEpochsCount += 1; }

        uint64 startBlock = uint64(block.number);
        uint64 seedBlock  = startBlock + uint64(SEED_DELAY_BLOCKS);
        epochs[epochId] = Epoch({
            organizer: msg.sender,
            policy: DKGTypes.EpochPolicy({
                threshold: threshold,
                committeeSize: committeeSize,
                minValidContributions: minValidContributions,
                lotteryAlphaBps: lotteryAlphaBps,
                committeeSelectionDeadlineBlock: startBlock + COMMITTEE_SELECTION_DEADLINE_OFFSET,
                keyAssemblyDeadlineBlock: startBlock + KEY_ASSEMBLY_DEADLINE_OFFSET,
                liveNotBeforeBlock:    startBlock + LIVE_NOT_BEFORE_OFFSET
            }),
            decryptionPolicy: decryptionPolicy,
            status: DKGTypes.EpochPhase.CommitteeSelection,
            nonce: epochNonce,
            startBlock: startBlock,
            seedBlock: seedBlock,
            seed: bytes32(0),
            lotteryThreshold: lotteryThreshold,
            claimedCount: 0,
            contributionCount: 0,
            partialDecryptionCount: 0,
            ciphertextCount: 0
        });

        lastEpochStartBlock = startBlock;
        emit EpochCreated(epochId, msg.sender, startBlock, seedBlock, lotteryThreshold);
        return epochId;
    }

    /// @notice Eligible registered nodes call this to claim a slot in the epoch's
    /// committee. The first `committeeSize` callers that pass the lottery and arrive
    /// before `committeeSelectionDeadlineBlock` form the committee.
    /// @notice Claim a committee slot in the trustless lottery.
    /// @dev    Admissible iff `keccak256(seed ‖ msg.sender) < lotteryThreshold`.
    ///         The first call after `block.number ≥ seedBlock` lazily resolves
    ///         `seed = blockhash(seedBlock)`; the contract emits
    ///         `SeedResolved` on that call. Further calls are served
    ///         first-come-first-served until `committeeSize` slots are filled,
    ///         at which point the committee snapshot is taken and the epoch
    ///         advances to Contribution.
    /// @param  epochId The epoch identifier returned by `createEpoch`.
    function claimSlot(bytes12 epochId) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (!PhaseLib.inCommitteeSelection(epoch.status, epoch.policy.committeeSelectionDeadlineBlock)) revert InvalidPhase();
        if (epoch.claimedCount >= epoch.policy.committeeSize) revert SlotsFull();
        if (selectedOperators[epochId][msg.sender]) revert AlreadyClaimed();

        // Lazy seed resolution: capture blockhash(seedBlock) on the first claimer
        // that arrives at or after seedBlock.
        bytes32 seed = epoch.seed;
        if (seed == bytes32(0)) {
            uint256 sb = uint256(epoch.seedBlock);
            if (block.number <= sb) revert SeedNotReady();
            // blockhash(b) returns 0 for any b ≥ block.number or b + 256 < block.number
            bytes32 fresh = blockhash(sb);
            if (fresh == bytes32(0)) revert SeedExpired();
            epoch.seed = fresh;
            seed = fresh;
            emit SeedResolved(epochId, fresh);
        }

        // Caller must be an active registered node.
        IDKGRegistry.NodeKey memory node = IDKGRegistry(REGISTRY).getNode(msg.sender);
        if (node.status != IDKGRegistry.NodeStatus.ACTIVE) revert NotRegistered();

        // Lottery check: hash(seed ‖ caller) must fall under the epoch threshold.
        if (uint256(keccak256(abi.encodePacked(seed, msg.sender))) >= epoch.lotteryThreshold) {
            revert NotEligible();
        }

        // First-come-first-served slot allocation.
        uint16 slot = epoch.claimedCount;
        epochParticipants[epochId].push(msg.sender);
        selectedOperators[epochId][msg.sender] = true;
        epoch.claimedCount = slot + 1;
        emit SlotClaimed(epochId, msg.sender, slot);

        // When the last slot is filled, snapshot the committee key hash and transition
        // to Contribution. The snapshot is needed so submitContribution can verify the
        // entire (recipientIndexes ‖ recipientPubKeys) calldata block in one keccak.
        if (slot + 1 == epoch.policy.committeeSize) {
            _snapshotCommittee(epochId, epoch.policy.committeeSize);
            epoch.status = DKGTypes.EpochPhase.KeyAssembly;
            emit CommitteeFilled(epochId);
        }
    }

    // `extendRegistration` was removed in the auto-cadence refactor. With a
    // fixed `EPOCH_DURATION_BLOCKS`-derived schedule, a stalled registration
    // simply gets aborted (`abortEpoch`) and the next scheduled epoch picks
    // up automatically once the cadence threshold passes.

/// @dev Internal helper: build the canonical (recipientIndexes ‖ pubKeys) layout
    /// for the committee that's just been filled and store its keccak256. Drives the
    /// O(1) committee verification in submitContribution.
    function _snapshotCommittee(bytes12 epochId, uint16 committeeSize) internal {
        uint256[MAX_N] memory canonicalIdxs;
        uint256[2 * MAX_N] memory canonicalKeys;
        address[] storage participants = epochParticipants[epochId];
        for (uint256 i = 0; i < committeeSize; i++) {
            canonicalIdxs[i] = i + 1;
            IDKGRegistry.NodeKey memory node = IDKGRegistry(REGISTRY).getNode(participants[i]);
            canonicalKeys[i * 2] = node.pubX;
            canonicalKeys[i * 2 + 1] = node.pubY;
        }
        for (uint256 i = committeeSize; i < MAX_N; i++) {
            canonicalKeys[i * 2 + 1] = 1; // identity-pad unused slots
        }
        epochContribPrefixHash[epochId] = keccak256(abi.encodePacked(canonicalIdxs, canonicalKeys));
    }

    /// @dev Wipes all storage tied to an old epoch when it falls out of the recent
    /// epochs ring buffer. Refunds gas via SSTORE-zero on the storage slots being
    /// cleared. Off-chain consumers must rely on the historical event log.
    ///
    /// Cleans up all four previously-leaking mappings in addition to the core
    /// epoch data: contributions, partial decryptions (per-ciphertext counts
    /// and per-participant records), and combined decryptions.
    function _evictRound(bytes12 oldRoundId) internal {
        Epoch storage r = epochs[oldRoundId];
        if (r.organizer == address(0)) return;
        address[] storage parts = epochParticipants[oldRoundId];
        uint256 n = parts.length;
        // Build the cleanup aid set: bytes32(0) (legacy per-epoch path) plus
        // every registered application aid (sourced from the sibling app
        // manager). When the app manager is not yet wired we only clean up
        // the legacy path.
        bytes32[] memory regAids;
        if (_appManagerSet) {
            regAids = IDKGAppManager(appManager).getRegisteredAids(oldRoundId);
        }
        uint256 aidCount = regAids.length + 1;
        for (uint256 i = 0; i < n; i++) {
            address participant = parts[i];
            delete selectedOperators[oldRoundId][participant];
            delete epochShareCommitmentHashes[oldRoundId][uint16(i + 1)];
            delete epochContributions[oldRoundId][participant];
        }
        // Clear per-ciphertext partial-decryption hashes and bitmaps + combined
        // decryption records, counts, and ciphertext hashes.
        for (uint256 a = 0; a < aidCount; a++) {
            bytes32 aid = a == 0 ? bytes32(0) : regAids[a - 1];
            for (uint16 ci = 1; ci <= MAX_CIPHERTEXT_INDEX; ci++) {
                if (epochPartialBitmap[oldRoundId][aid][ci] != 0) {
                    uint256 bm = epochPartialBitmap[oldRoundId][aid][ci];
                    for (uint16 pIdx = 1; pIdx <= MAX_N; pIdx++) {
                        if ((bm >> pIdx) & 1 == 1) {
                            delete epochPartialDeltaHash[oldRoundId][aid][ci][pIdx];
                        }
                    }
                    delete epochPartialBitmap[oldRoundId][aid][ci];
                }
                if (epochPartialDecryptionCounts[oldRoundId][aid][ci] > 0) {
                    delete epochPartialDecryptionCounts[oldRoundId][aid][ci];
                }
                if (epochCombinedDecryptions[oldRoundId][aid][ci].completed) {
                    delete epochCombinedDecryptions[oldRoundId][aid][ci];
                }
                if (_ciphertexts[oldRoundId][aid][ci] != bytes32(0)) {
                    delete _ciphertexts[oldRoundId][aid][ci];
                }
            }
            // applications[oldRoundId][aid] now lives on DKGAppManager.
            // Its lazy-cleanup is acceptable per the audit notes.
        }
        delete epochParticipants[oldRoundId];
        delete epochContribPrefixHash[oldRoundId];
        delete _collectiveKey[oldRoundId];
        delete epochs[oldRoundId];
        emit EpochEvicted(oldRoundId);
    }

    /// @notice Submit a contributor's polynomial commitments, encrypted
    ///         shares and Groth16 proof of correctness.
    /// @dev    The committee membership + BabyJubJub public keys are verified
    ///         against a single `keccak256` snapshot taken when the lottery
    ///         filled; the transcript is read straight from calldata via the
    ///         BRLC helper. The transaction reverts if the proof fails.
    function submitContribution(
        bytes12 epochId,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (!PhaseLib.inKeyAssembly(epoch.status, epoch.policy.keyAssemblyDeadlineBlock)) revert InvalidPhase();
        if (!selectedOperators[epochId][msg.sender]) revert NotSelectedParticipant();
        if (contributorIndex == 0 || contributorIndex > epoch.policy.committeeSize) revert InvalidContribution();
        if (epochParticipants[epochId][contributorIndex - 1] != msg.sender) revert InvalidProofInput();

        DKGTypes.ContributionRecord storage record = epochContributions[epochId][msg.sender];
        if (record.accepted) revert AlreadyContributed();

        // Cheap public-input checks first; only invoke the expensive verifier
        // once we've confirmed the proof targets the right epoch + contributor.
        uint256[8] memory publicInputs = abi.decode(input, (uint256[8]));
        if (
            publicInputs[0] != _epochScalar(epochId) || publicInputs[1] != epoch.policy.threshold
                || publicInputs[2] != epoch.policy.committeeSize || publicInputs[3] != contributorIndex
                || bytes32(publicInputs[4]) != commitmentsHash || bytes32(publicInputs[5]) != encryptedSharesHash
        ) revert InvalidProofInput();
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            CONTRIBUTION_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(commitmentsHash, encryptedSharesHash))
        );
        if (publicInputs[6] != challenge) revert InvalidProofInput();
        // publicInputs[7] = TranscriptCommitment (verified below via BRLC)
        IZKVerifier(CONTRIBUTION_VERIFIER).verifyProof(proof, input);

        // Transcript layout (8N words = 256 N=32, 128 N=16):
        //   words [0..2N)     commitmentPoints  (N points × 2 coords)
        //   words [2N..3N)    recipientIndexes
        //   words [3N..5N)    recipientPubKeys  (N points × 2 coords)
        //   words [5N..7N)    ephemerals
        //   words [7N..8N)    maskedShares
        if (transcript.length != CONTRIB_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        bytes32 commitmentDigest = keccak256(transcript[0:CONTRIB_DIGEST_BYTES_LEN]);

        // Single-shot committee verification: bytes [recipientIndexes..recipientPubKeys-end)
        // of the transcript hold the canonical recipientIndexes ‖ recipientPubKeys section.
        // Compare against the hash snapshotted when the lottery filled. This replaces the
        // previous per-recipient loop with N storage reads + N external registry calls.
        if (keccak256(transcript[CONTRIB_DIGEST_BYTES_LEN:CONTRIB_PUBKEYS_BYTES_END]) != epochContribPrefixHash[epochId]) {
            revert InvalidProofInput();
        }

        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, CONTRIB_TRANSCRIPT_WORDS) != publicInputs[7]) revert InvalidProofInput();

        // Only persist the fields the contract itself actually needs:
        //   - commitmentVectorDigest: re-checked at finalize time
        //   - contributorIndex + accepted: identity & dup-prevention gates
        // commitmentsHash, encryptedSharesHash, and the redundant `contributor` are
        // only emitted in the event below; off-chain consumers read them from logs.
        DKGTypes.ContributionRecord storage rec = epochContributions[epochId][msg.sender];
        rec.contributorIndex = contributorIndex;
        rec.commitmentVectorDigest = commitmentDigest;
        rec.accepted = true;
        epoch.contributionCount++;

        // Refresh the contributor's liveness timestamp on the registry.
        // A successful proof-gated contribution is the strongest possible
        // signal that the operator is alive and well-configured.
        IDKGRegistry(REGISTRY).markActive(msg.sender);

        emit ContributionSubmitted(epochId, msg.sender, contributorIndex, commitmentsHash, encryptedSharesHash);
    }

    /// @notice Returns the collective public key for an epoch.
    ///         Written exactly once at `finalizeEpoch` from
    ///         `aggregateCommitments[0] = Σᵢ aᵢ,₀·G`. Returns the identity
    ///         element `(0, 1)` for epochs that have not yet been finalized.
    function getCollectivePublicKey(bytes12 epochId) external view returns (DKGTypes.Point memory) {
        DKGTypes.Point storage cpk = _collectiveKey[epochId];
        if (cpk.y == 0) {
            return DKGTypes.Point({x: 0, y: 1});
        }
        return cpk;
    }

    /// @notice Aggregate accepted contributions, publish the collective
    ///         public key, and transition the epoch to Finalized.
    /// @dev    Callable by anyone once `contributionCount ≥
    ///         policy.minValidContributions`. Stores share commitments as
    ///         `keccak256(x, y)` per participant to keep storage to a single
    ///         slot per entry; the pre-image is emitted in `EpochLive`.
    function finalizeEpoch(
        bytes12 epochId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status == DKGTypes.EpochPhase.Live) revert AlreadyLive();
        if (epoch.status != DKGTypes.EpochPhase.KeyAssembly) revert InvalidPhase();
        // liveNotBeforeBlock gate — semantically a "phase not yet open"
        // condition, so we reuse InvalidPhase to keep the contract small.
        if (block.number < uint256(epoch.policy.liveNotBeforeBlock)) revert InvalidPhase();
        if (epoch.contributionCount < epoch.policy.minValidContributions) revert InsufficientContributions();
        if (
            aggregateCommitmentsHash == bytes32(0) || collectivePublicKeyHash == bytes32(0)
                || shareCommitmentHash == bytes32(0)
        ) revert InvalidFinalization();

        // Cheap public-input checks first; only invoke the verifier when the
        // proof targets the right epoch / aggregate.
        uint256[9] memory publicInputs = abi.decode(input, (uint256[9]));
        if (
            publicInputs[0] != _epochScalar(epochId) || publicInputs[1] != epoch.policy.threshold
                || publicInputs[2] != epoch.policy.committeeSize || publicInputs[3] != epoch.contributionCount
                || bytes32(publicInputs[4]) != aggregateCommitmentsHash
                || bytes32(publicInputs[5]) != collectivePublicKeyHash
                || bytes32(publicInputs[6]) != shareCommitmentHash
        ) revert InvalidProofInput();

        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            FINALIZE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash))
        );
        if (publicInputs[7] != challenge) revert InvalidProofInput();
        IZKVerifier(FINALIZE_VERIFIER).verifyProof(proof, input);

        // Transcript layout (2N² + 5N words):
        //   words [0..N)              participantIndexes
        //   words [N..N+2N²)          contributionCommitments  (N contributors × N points × 2 coords)
        //   words [N+2N²..N+2N²+2N)   aggregateCommitments     (N points × 2 coords)
        //   words [N+2N²+2N..2N²+5N)  shareCommitments         (N points × 2 coords)
        if (transcript.length != FINALIZE_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        uint256 dOff;
        assembly { dOff := transcript.offset }

        _verifyFinalizeTranscript(epochId, epoch, challenge, publicInputs[8], transcript);

        // The collective public key is `aggregateCommitments[0]`. The proof
        // attests that the aggregate is the correctly-summed Σᵢ Cᵢ over the
        // accepted contributor set, and `_verifyFinalizeTranscript` rejects
        // duplicate participant rows and binds each row to the on-chain
        // accepted contribution digest. Read aggregate[0] straight from the
        // transcript calldata and persist it once.
        {
            uint256 agg0Base = dOff + FINALIZE_AGG0_WORDS_OFFSET * 32;
            uint256 agg0X;
            uint256 agg0Y;
            assembly ("memory-safe") {
                agg0X := calldataload(agg0Base)
                agg0Y := calldataload(add(agg0Base, 0x20))
            }
            DKGTypes.Point storage cpkRef = _collectiveKey[epochId];
            cpkRef.x = agg0X;
            cpkRef.y = agg0Y;
        }

        epoch.status = DKGTypes.EpochPhase.Live;
        // The three commitment hashes are not persisted to storage; they are emitted
        // in EpochLive below and reconstructed off-chain from the event log.

        // Persist share commitments directly from calldata, in the same loop as the
        // already-validated participantIndexes pass.
        uint256 ccount = epoch.contributionCount;
        uint256 piBase = dOff;                                       // participantIndexes
        uint256 scBase = dOff + FINALIZE_SHARE_WORDS_OFFSET * 32;    // shareCommitments
        for (uint256 i = 0; i < ccount; i++) {
            uint256 pIdx;
            uint256 scX;
            uint256 scY;
            assembly ("memory-safe") {
                pIdx := calldataload(add(piBase, mul(i, 0x20)))
                scX := calldataload(add(scBase, mul(i, 0x40)))
                scY := calldataload(add(scBase, add(mul(i, 0x40), 0x20)))
            }
            epochShareCommitmentHashes[epochId][uint16(pIdx)] = _hash2(scX, scY);
        }

        emit EpochLive(epochId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash);
    }

    /// @dev Verifies the combineDecryption transcript directly from calldata.
    function _verifyCombineTranscript(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        Epoch storage epoch,
        uint256 shareCount,
        bytes calldata transcript
    ) internal view {
        uint256 dOff;
        assembly { dOff := transcript.offset }
        uint256 piBase = dOff + 4 * 32;                               // participantIndexes start
        uint256 pdBase = dOff + COMBINE_PARTIALS_BYTES_OFFSET;        // partialDecryptions start

        uint256 cs = epoch.policy.committeeSize;
        // Track which participant indexes have been
        // seen so duplicates can't pad the qualifying set with copies of
        // the same partial. Bitmap fits because participantIndex ≤ MAX_N
        // ≤ 32. The contract-side check + the in-circuit `mask =
        // PrefixMask(ShareCount, MaxShares)` together enforce a proper
        // qualifying set without needing in-circuit pairwise inequality.
        uint256 seenIndexes;
        for (uint256 i = 0; i < shareCount; i++) {
            uint256 pIdxRaw;
            uint256 pdX;
            uint256 pdY;
            assembly ("memory-safe") {
                pIdxRaw := calldataload(add(piBase, mul(i, 0x20)))
                pdX := calldataload(add(pdBase, mul(i, 0x40)))
                pdY := calldataload(add(pdBase, add(mul(i, 0x40), 0x20)))
            }
            uint16 participantIndex = uint16(pIdxRaw);
            if (participantIndex == 0 || participantIndex > cs) revert InvalidProofInput();
            uint256 indexBit = uint256(1) << participantIndex;
            if (seenIndexes & indexBit != 0) revert InvalidProofInput();
            seenIndexes |= indexBit;
            // Bitmap bit must be set for this participant — i.e. they
            // submitted a partial under this (epochId, aid, ctIdx).
            if (epochPartialBitmap[epochId][aid][ciphertextIndex] & indexBit == 0) {
                revert InvalidProofInput();
            }
            // The transcript-supplied δ must hash to the value we stored at
            // submitPartialDecryption time.
            if (
                keccak256(abi.encodePacked(pdX, pdY))
                    != epochPartialDeltaHash[epochId][aid][ciphertextIndex][participantIndex]
            ) revert InvalidProofInput();
        }
        for (uint256 i = shareCount; i < MAX_N; i++) {
            uint256 pIdxRaw;
            uint256 pdX;
            uint256 pdY;
            assembly ("memory-safe") {
                pIdxRaw := calldataload(add(piBase, mul(i, 0x20)))
                pdX := calldataload(add(pdBase, mul(i, 0x40)))
                pdY := calldataload(add(pdBase, add(mul(i, 0x40), 0x20)))
            }
            if (pIdxRaw != 0) revert InvalidProofInput();
            if (pdX != 0 || pdY != 1) revert InvalidProofInput();
        }
    }

    /// @dev Verifies per-contributor commitment digests and the BRLC commitment over the
    /// finalize transcript directly out of calldata (no abi.decode, no memory copies).
    function _verifyFinalizeTranscript(
        bytes12 epochId,
        Epoch storage epoch,
        uint256 challenge,
        uint256 expectedBrlc,
        bytes calldata transcript
    ) internal view {
        uint256 dOff;
        assembly { dOff := transcript.offset }
        uint256 piBase = dOff;                                         // participantIndexes

        // contributionCommitments occupies the next 2N² words.
        bytes calldata contribCommBytes =
            transcript[FINALIZE_CONTRIB_BYTES_OFFSET:FINALIZE_CONTRIB_BYTES_OFFSET + FINALIZE_CONTRIB_BYTES_LEN];

        uint256 ccount = epoch.contributionCount;
        uint256 cSize = epoch.policy.committeeSize;
        // Reject duplicate participant indexes in the
        // active prefix. Without this an attacker could repeat one accepted
        // contributor's row and omit another, finalising an aggregate that
        // disagrees with the on-chain accumulated `_collectiveKey`. Bitmap
        // fits because participantIndex ≤ MAX_N ≤ 32.
        uint256 seenIndexes;
        for (uint256 i = 0; i < ccount; i++) {
            uint256 pIdxRaw;
            assembly ("memory-safe") {
                pIdxRaw := calldataload(add(piBase, mul(i, 0x20)))
            }
            uint16 participantIndex = uint16(pIdxRaw);
            if (participantIndex == 0 || participantIndex > cSize) revert InvalidProofInput();
            uint256 indexBit = uint256(1) << participantIndex;
            if (seenIndexes & indexBit != 0) revert InvalidProofInput();
            seenIndexes |= indexBit;
            address participant = epochParticipants[epochId][participantIndex - 1];
            DKGTypes.ContributionRecord storage contribution = epochContributions[epochId][participant];
            if (!contribution.accepted || contribution.contributorIndex != participantIndex) revert InvalidProofInput();

            // Each contributor's commitments occupy 2N words.
            bytes32 digest = keccak256(
                contribCommBytes[i * FINALIZE_PER_CONTRIB_BYTES:(i + 1) * FINALIZE_PER_CONTRIB_BYTES]
            );
            if (digest != contribution.commitmentVectorDigest) revert InvalidProofInput();
        }

        // Stream BRLC over the entire 2N² + 5N word transcript region.
        if (BRLC.commitCalldata(challenge, dOff, FINALIZE_TRANSCRIPT_WORDS) != expectedBrlc) revert InvalidProofInput();
    }

    /// @notice Submit a committee member's partial decryption `δ_i = d_i · C_1`.
    /// @dev    Keyed by `(epochId, participant, ciphertextIndex)` to support
    ///         multiple ciphertexts per epoch. The Groth16 proof is a
    ///         Chaum–Pedersen DLEQ establishing that `δ_i` and the committed
    ///         share `D_i` share a discrete log with respect to `C_1` and `G`.
    /// @dev `aid` binds the proof transcript to a specific application.
    ///      Pass `bytes32(0)` for the legacy per-epoch path that
    ///      pre-dates the application surface; the circuit witness builder
    ///      defaults Aid/CtIdx/Role to zero/zero/COMMITTEE in that mode.
    /// @dev `c1x/c1y/c2x/c2y` are the raw ciphertext coordinates as
    ///      submitted via submitCiphertext. The contract verifies
    ///      `keccak256(abi.encode(...))` matches the stored ciphertext
    ///      hash and then binds the proof's public-input C1 (pi[5..6])
    ///      to the authoritative on-chain ciphertext.
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
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (!selectedOperators[epochId][msg.sender]) revert NotSelectedParticipant();
        if (
            participantIndex == 0 || participantIndex > epoch.policy.committeeSize || ciphertextIndex == 0
                || ciphertextIndex > MAX_CIPHERTEXT_INDEX || deltaHash == bytes32(0)
        ) revert InvalidPartialDecryption();
        if (epochParticipants[epochId][participantIndex - 1] != msg.sender) revert InvalidProofInput();

        // Bind to the authoritative on-chain ciphertext. Without this the
        // prover can supply Δ_i = sk_i · B for an arbitrary B, and the stored
        // partial decryption is only meaningful relative to that B — combine
        // then aggregates points that aren't decryptions of the submitted
        // ciphertext. Ciphertext + partial storage are keyed by aid so two
        // applications under the same epoch don't collide on ctIdx.
        bytes32 storedCt = _ciphertexts[epochId][aid][ciphertextIndex];
        if (storedCt == bytes32(0)) revert CiphertextNotSubmitted();
        if (_hash4(c1x, c1y, c2x, c2y) != storedCt) revert InvalidProofInput();

        uint256 indexBit = uint256(1) << participantIndex;
        if (epochPartialBitmap[epochId][aid][ciphertextIndex] & indexBit != 0) {
            revert AlreadyPartiallyDecrypted();
        }

        // Layout: [eid, aid, ctIdx, role, i, C1.x, C1.y, D_i.x, D_i.y,
        // delta.x, delta.y, A1.x, A1.y, A2.x, A2.y, response].
        // Committee partial decryptions always use role = COMMITTEE = 1
        // (organizer shares go through submitOrganizerShare instead).
        // Cheap public-input checks fail before the expensive verifier call.
        uint256[16] memory publicInputs = abi.decode(input, (uint256[16]));
        bytes32 storedScHash = epochShareCommitmentHashes[epochId][participantIndex];
        if (
            publicInputs[0] != _epochScalar(epochId)
                || publicInputs[1] != uint256(aid)
                || publicInputs[2] != ciphertextIndex
                || publicInputs[3] != uint256(DKGProtocol.ROLE_COMMITTEE)
                || publicInputs[4] != participantIndex
                // pi[5..6] = base point (C_1) — bind to the just-verified
                // on-chain ciphertext.
                || publicInputs[5] != c1x
                || publicInputs[6] != c1y
                || storedScHash == bytes32(0)
                || _hash2(publicInputs[7], publicInputs[8]) != storedScHash
        ) revert InvalidProofInput();
        if (deltaHash != keccak256(abi.encodePacked(publicInputs[9], publicInputs[10]))) revert InvalidProofInput();
        IZKVerifier(PARTIAL_DECRYPT_VERIFIER).verifyProof(proof, input);

        // Persist the δ commitment as a single 32-byte hash plus a bitmap bit.
        // The combine path reads δ.x/δ.y back from the proof transcript and
        // re-hashes against the stored value, so we don't store the raw point.
        epochPartialDeltaHash[epochId][aid][ciphertextIndex][participantIndex] = deltaHash;
        epochPartialBitmap[epochId][aid][ciphertextIndex] |= indexBit;
        epoch.partialDecryptionCount++;
        epochPartialDecryptionCounts[epochId][aid][ciphertextIndex]++;

        emit PartialDecryptionSubmitted(
            epochId, aid, msg.sender, participantIndex, ciphertextIndex,
            publicInputs[9], publicInputs[10]
        );
    }

    /// @notice Submit a ciphertext to be threshold-decrypted by the committee.
    /// @dev    Enforces the epoch's DecryptionPolicy: owner-only, block/timestamp
    ///         windows, and a cap on accepted ciphertexts per epoch. Write-once
    ///         per `ciphertextIndex`. Stores `keccak256(c1x, c1y, c2x, c2y)` so
    ///         `combineDecryption` can bind its proof's ciphertext public inputs
    ///         to the authoritative on-chain value. The raw coordinates are only
    ///         exposed via the `CiphertextSubmitted` event (nodes watch it).
    function submitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert InvalidCiphertext();

        // Well-formedness: coords must be canonical (< Q), on-curve, non-identity,
        // and in the prime-order subgroup. The first three are cheap (4 mulmods per
        // point); the subgroup check is one full BJJ scalar mul (~60k gas per point)
        // but is required to honor the paper's group-validation policy at every
        // entry point (paper §4.1). Without subgroup
        // membership a griefing submitter could pre-claim every index with a small-
        // order point the combine circuit can never accept.
        _requireValidEncryptionPoint(c1x, c1y);
        _requireValidEncryptionPoint(c2x, c2y);

        // Per-epoch DecryptionPolicy gates the legacy aid=0 path; per-application
        // AppPolicy gates aid != 0.
        if (aid == bytes32(0)) {
            DKGTypes.DecryptionPolicy memory p = epoch.decryptionPolicy;
            if (p.ownerOnly && msg.sender != epoch.organizer) revert NotOwner();
            if (p.notBeforeBlock     != 0 && uint64(block.number)    < p.notBeforeBlock)     revert DecryptionNotYetAllowed();
            if (p.notBeforeTimestamp != 0 && uint64(block.timestamp) < p.notBeforeTimestamp) revert DecryptionNotYetAllowed();
            if (p.notAfterBlock      != 0 && uint64(block.number)    > p.notAfterBlock)      revert DecryptionExpired();
            if (p.notAfterTimestamp  != 0 && uint64(block.timestamp) > p.notAfterTimestamp)  revert DecryptionExpired();
            if (p.maxDecryptions     != 0 && epoch.ciphertextCount   >= p.maxDecryptions)    revert DecryptionLimitReached();
        } else {
            if (!_appManagerSet) revert AppManagerNotSet();
            IDKGAppManager(appManager).requireCanSubmitCiphertext(epochId, aid, ciphertextIndex, msg.sender);
        }

        if (_ciphertexts[epochId][aid][ciphertextIndex] != bytes32(0)) revert CiphertextAlreadySubmitted();

        _ciphertexts[epochId][aid][ciphertextIndex] = _hash4(c1x, c1y, c2x, c2y);
        unchecked { epoch.ciphertextCount += 1; }

        emit CiphertextSubmitted(epochId, aid, ciphertextIndex, msg.sender, c1x, c1y, c2x, c2y);
    }

    /// @dev Validate that (x, y) is a canonical, on-curve, non-identity point
    ///      on BabyJubJub. Reverts with InvalidCiphertext() so callers observe
    ///      a single failure mode at submitCiphertext time.
    ///
    ///      The prime-subgroup membership check (`[L]·P == identity`) is
    ///      intentionally NOT performed here. A small-order ciphertext point
    ///      cannot be combined into a valid plaintext (the combine SNARK's
    ///      `M = m·G` constraint would have no solution), so the worst an
    ///      attacker can do is occupy a ciphertext slot that no committee
    ///      member will ever decrypt — and committee node software is
    ///      required to subgroup-check before computing a partial decryption,
    ///      so no `δ_i = sk_i · c1` ever lands on-chain to leak `sk_i mod h`.
    ///      The off-chain enforcement layer is the canonical defence against
    ///      small-order ciphertext submissions.
    ///
    ///      Skipping the on-chain check saves ~2 M gas per submission (one
    ///      full BJJ scalar multiplication per coordinate); the residual
    ///      defence relies on (i) the off-chain committee node policy and
    ///      (ii) the application's `authorizedSubmitter` / `maxCiphertexts`
    ///      policy gating the slot count.
    function _requireValidEncryptionPoint(uint256 x, uint256 y) internal pure {
        if (!BabyJubJub.isCanonical(x, y)) revert InvalidCiphertext();
        if (BabyJubJub.isIdentity(x, y)) revert InvalidCiphertext();
        if (!BabyJubJub.isOnCurve(x, y)) revert InvalidCiphertext();
    }

    /// @notice Combine `t` partial decryptions via Lagrange interpolation and
    ///         persist the recovered plaintext on-chain.
    /// @dev    Callable by anyone once at least `threshold` partial
    ///         decryptions with matching `ciphertextIndex` are on-chain and the
    ///         ciphertext itself has been submitted via `submitCiphertext`.
    ///         The proof's ciphertext public inputs are bound to the stored
    ///         ciphertext hash; a combiner cannot substitute a different ct.
    /// @notice Per-application combine. The `aid` parameter selects the
    ///         application registered against `epochId`; if the app exists
    ///         in mode 1 (organizer co-decryption), the previously stored
    ///         `Δ_org` is supplied to the verifier as the correction point;
    ///         in mode 0, the stored derivation tag `S` is supplied.
    /// @dev    Public-input layout (13 fields): eid, aid, ctIdx, mode, S,
    ///         deltaOrgX, deltaOrgY, threshold, shareCount, combineHash,
    ///         plaintextHash, challenge, transcriptCommitment. Matches the
    ///         circuit definition in `circuits/decryptcombine/circuit.go`
    ///         per paper §5.5 lines 1051–1077.
    function combineDecryption(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        bytes32 combineHash,
        uint256 plaintext,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX || combineHash == bytes32(0)) revert InvalidCombinedDecryption();
        bytes32 storedCtHash = _ciphertexts[epochId][aid][ciphertextIndex];
        if (storedCtHash == bytes32(0)) revert CiphertextNotSubmitted();
        if (epochPartialDecryptionCounts[epochId][aid][ciphertextIndex] < epoch.policy.threshold) revert InsufficientPartialDecryptions();

        DKGTypes.CombinedDecryptionRecord storage record = epochCombinedDecryptions[epochId][aid][ciphertextIndex];
        if (record.completed) revert AlreadyCombined();

        // Validate cheap public-input bindings before invoking the verifier
        // so a mismatched (eid, aid, ctIdx, ...) submission fails before the
        // ~280 k pairing check.
        uint256 shareCount = _validateAndPostCombine(
            epochId, aid, ciphertextIndex, combineHash, plaintext, epoch, input, transcript, storedCtHash
        );
        IZKVerifier(DECRYPT_COMBINE_VERIFIER).verifyProof(proof, input);
        _verifyCombineTranscript(epochId, aid, ciphertextIndex, epoch, shareCount, transcript);

        record.completed = true;
        record.plaintext = plaintext;

        emit DecryptionCombined(epochId, aid, ciphertextIndex, combineHash, plaintext);
    }

    /// @dev Resolves application correction (mode 0 / mode 1), validates the
    /// public-input layout against eid/aid/ctIdx/mode/S/Δ_org/threshold/
    /// combineHash/plaintext/challenge, range-checks shareCount, binds the
    /// transcript's first 128 bytes to the stored ciphertext hash, and verifies
    /// the BRLC commitment over the full transcript region. Split out of
    /// combineDecryption to keep the parent's stack within Yul's depth limit.
    function _validateAndPostCombine(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        bytes32 combineHash,
        uint256 plaintext,
        Epoch storage epoch,
        bytes calldata input,
        bytes calldata transcript,
        bytes32 storedCtHash
    ) internal view returns (uint256 shareCount) {
        // Resolve the application's per-app correction. aid == 0 is the
        // legacy per-epoch combine path: mode = 0, S = 0, Δ_org = identity.
        uint8 mode = uint8(DKGProtocol.MODE_PUBLIC_DERIVATION);
        uint256 derivationS;
        DKGTypes.Point memory deltaOrg = DKGTypes.Point({x: 0, y: 1});
        if (aid != bytes32(0)) {
            if (!_appManagerSet) revert AppManagerNotSet();
            (uint8 m, uint256 s, uint256 dx, uint256 dy) =
                IDKGAppManager(appManager).getCombineCorrection(epochId, aid, ciphertextIndex);
            mode = m;
            derivationS = s;
            deltaOrg = DKGTypes.Point({x: dx, y: dy});
        }

        uint256[13] memory publicInputs = abi.decode(input, (uint256[13]));
        if (
            publicInputs[0] != _epochScalar(epochId)
                || publicInputs[1] != uint256(aid)
                || publicInputs[2] != uint256(ciphertextIndex)
                || publicInputs[3] != uint256(mode)
                || publicInputs[4] != derivationS
                || publicInputs[5] != deltaOrg.x
                || publicInputs[6] != deltaOrg.y
                || publicInputs[7] != epoch.policy.threshold
                || bytes32(publicInputs[9]) != combineHash
                || publicInputs[10] != plaintext
        ) revert InvalidProofInput();
        if (publicInputs[8] < epoch.policy.threshold) revert InvalidProofInput();
        if (publicInputs[8] > MAX_N) revert InvalidProofInput();
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            DECRYPT_COMBINE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(combineHash, bytes32(plaintext)))
        );
        if (publicInputs[11] != challenge) revert InvalidProofInput();
        if (transcript.length != COMBINE_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        if (keccak256(transcript[0:128]) != storedCtHash) revert InvalidProofInput();
        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, COMBINE_TRANSCRIPT_WORDS) != publicInputs[12]) revert InvalidProofInput();
        return publicInputs[8];
    }


    /// @notice Abort a non-terminal epoch. Organizer only.
    /// @dev    Finalized epochs may NOT be aborted: the collective public key has
    ///         already been published and messages may already be encrypted to it.
    ///         Aborting after finalization would permanently block decryption for
    ///         those messages. Only Registration and Contribution phases are
    ///         abortable.
    /// @param  epochId The epoch identifier.
    function abortEpoch(bytes12 epochId) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (msg.sender != epoch.organizer) revert Unauthorized();
        if (
            epoch.status == DKGTypes.EpochPhase.Live
                || epoch.status == DKGTypes.EpochPhase.Completed
                || epoch.status == DKGTypes.EpochPhase.Aborted
        ) {
            revert InvalidPhase();
        }

        epoch.status = DKGTypes.EpochPhase.Aborted;
        emit EpochAborted(epochId);
    }

    function getEpoch(bytes12 epochId) external view returns (Epoch memory) {
        return epochs[epochId];
    }

    function selectedParticipants(bytes12 epochId) external view returns (address[] memory) {
        return epochParticipants[epochId];
    }

    function getContribution(bytes12 epochId, address contributor)
        external
        view
        returns (DKGTypes.ContributionRecord memory)
    {
        return epochContributions[epochId][contributor];
    }

    /// @notice Returns the stored partial-decryption record for a single
    ///         committee member's submission, keyed by participant index.
    ///         Returns `accepted = false` if no submission was made.
    function getPartialDecryption(bytes12 epochId, bytes32 aid, uint16 participantIndex, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.PartialDecryptionRecord memory)
    {
        bool accepted = (epochPartialBitmap[epochId][aid][ciphertextIndex] >> participantIndex) & 1 == 1;
        return DKGTypes.PartialDecryptionRecord({
            participantIndex: accepted ? participantIndex : 0,
            ciphertextIndex: accepted ? ciphertextIndex : 0,
            deltaHash: epochPartialDeltaHash[epochId][aid][ciphertextIndex][participantIndex],
            accepted: accepted
        });
    }

    function getCombinedDecryption(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.CombinedDecryptionRecord memory)
    {
        return epochCombinedDecryptions[epochId][aid][ciphertextIndex];
    }

    /// @notice Returns the keccak256(abi.encode(x, y)) commitment hash for a
    /// participant's share commitment. The pre-image (x,y) is exposed off-chain via
    /// the EpochLive event log.
    function getShareCommitmentHash(bytes12 epochId, uint16 participantIndex)
        external
        view
        returns (bytes32)
    {
        return epochShareCommitmentHashes[epochId][participantIndex];
    }

    /// @notice _hash4(c1x, c1y, c2x, c2y) of the ciphertext stored
    ///         at `ciphertextIndex` for `epochId`. Returns bytes32(0) if no
    ///         ciphertext has been submitted at this slot.
    function getCiphertextHash(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) external view returns (bytes32) {
        return _ciphertexts[epochId][aid][ciphertextIndex];
    }

    /// @notice Recovered plaintext for (epochId, aid, ciphertextIndex). Returns 0 if
    ///         the decryption has not been combined yet; callers should also
    ///         consult `getCombinedDecryption(...)` / `DecryptionCombined`
    ///         events to disambiguate "not yet combined" from "plaintext is 0".
    function getPlaintext(bytes12 epochId, bytes32 aid, uint16 ciphertextIndex) external view returns (uint256) {
        return epochCombinedDecryptions[epochId][aid][ciphertextIndex].plaintext;
    }

    function getDecryptionPolicy(bytes12 epochId) external view returns (DKGTypes.DecryptionPolicy memory) {
        return epochs[epochId].decryptionPolicy;
    }

    function getContributionVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(CONTRIBUTION_VERIFIER).provingKeyHash();
    }

    function getPartialDecryptVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(PARTIAL_DECRYPT_VERIFIER).provingKeyHash();
    }

    function getFinalizeVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(FINALIZE_VERIFIER).provingKeyHash();
    }

    function getDecryptCombineVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(DECRYPT_COMBINE_VERIFIER).provingKeyHash();
    }

    function _epochScalar(bytes12 epochId) internal pure returns (uint256) {
        return uint256(uint96(epochId));
    }

    /// @dev keccak256 of two 32-byte words written into scratch memory.
    ///      Skips the abi.encode allocation/length-prefix overhead.
    function _hash2(uint256 a, uint256 b) internal pure returns (bytes32 h) {
        assembly ("memory-safe") {
            let p := mload(0x40)
            mstore(p, a)
            mstore(add(p, 0x20), b)
            h := keccak256(p, 0x40)
        }
    }

    /// @dev keccak256 of four 32-byte words written into scratch memory.
    function _hash4(uint256 a, uint256 b, uint256 c, uint256 d) internal pure returns (bytes32 h) {
        assembly ("memory-safe") {
            let p := mload(0x40)
            mstore(p, a)
            mstore(add(p, 0x20), b)
            mstore(add(p, 0x40), c)
            mstore(add(p, 0x60), d)
            h := keccak256(p, 0x80)
        }
    }
}
