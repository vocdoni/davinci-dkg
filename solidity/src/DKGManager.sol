// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {IDKGAppManager} from "./interfaces/IDKGAppManager.sol";
import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGIdLib} from "./libraries/DKGIdLib.sol";
import {BRLC} from "./libraries/BRLC.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {PhaseLib} from "./libraries/PhaseLib.sol";
import {
    MAX_N,
    MAX_K,
    MERKLE_DEPTH,
    MERKLE_EMPTY_LEAF,
    CONTRIB_TRANSCRIPT_WORDS,
    CONTRIB_COMMITTEE_BYTES_OFFSET,
    CONTRIB_COMMITTEE_BYTES_END,
    POOLKEY_TRANSCRIPT_WORDS,
    POOLKEY_HASHES_BYTES_OFFSET,
    POOLKEY_AGG_WORDS_OFFSET,
    POOLKEY_SHARE_WORDS_OFFSET,
    COMBINE_TRANSCRIPT_WORDS,
    COMBINE_INDEXES_BYTES_OFFSET,
    COMBINE_PARTIALS_BYTES_OFFSET,
    DEFAULT_EPOCH_DURATION_BLOCKS,
    DEFAULT_COMMITTEE_SELECTION_BLOCKS,
    DEFAULT_KEY_ASSEMBLY_BLOCKS,
    DEFAULT_FINALIZE_GAP_BLOCKS,
    SEED_DELAY_BLOCKS
} from "./libraries/Sizes.sol";

/// @title  DKGManager
/// @notice On-chain orchestrator for every phase of a davinci-dkg epoch.
/// @dev    Lifecycle: CommitteeSelection (trustless lottery) → KeyAssembly →
///         Live → Completed (or Aborted). Every state-mutating entry point
///         that makes a cryptographic claim is gated by a Groth16 verifier —
///         no dispute phase, no complaint flow. Epoch storage is never
///         evicted: an application may outlive many epochs and `createEpoch`
///         must stay O(1) gas regardless of history.
///
///         Each epoch's DKG deals `MAX_K` independent keys. `finalizeEpoch`
///         only freezes the accepted contributor set; each pool key is proven
///         separately by `activatePoolKey`, which stores `P_j` and a keccak
///         Merkle root over the committee's share commitments for that key
///         (one SSTORE instead of `MAX_N`). Partial decryptions carry the
///         matching Merkle path. Every transcript is read straight out of
///         calldata via assembly to avoid per-element bounds checks.
contract DKGManager is IDKGManager {
    /// @dev Sibling app manager has not been wired yet — every application
    ///      path (submitCiphertext / combineDecryption) reverts until it is.
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

    /// @dev Upper bound on ciphertext indices accepted by `submitPartialDecryption`
    /// and `combineDecryption`. Prevents unbounded storage spam by a committee member
    /// who submits decryptions for arbitrarily large ciphertext indices.
    uint16 internal constant MAX_CIPHERTEXT_INDEX = 256;

    uint32 public immutable CHAIN_ID;
    address public immutable REGISTRY;
    uint32 public immutable EPOCH_PREFIX;
    address public immutable CONTRIBUTION_VERIFIER;
    address public immutable PARTIAL_DECRYPT_VERIFIER;
    address public immutable POOL_KEY_VERIFIER;
    address public immutable DECRYPT_COMBINE_VERIFIER;
    /// @notice Total epoch length in blocks. Set per-deploy by the constructor.
    ///         Defaults to `DEFAULT_EPOCH_DURATION_BLOCKS` (Sizes.sol) when the
    ///         constructor argument is 0. Wall-clock duration depends on the
    ///         deployment chain's block time and is estimated off-chain.
    uint256 public immutable EPOCH_DURATION_BLOCKS;
    /// @notice Deployment floors/ceilings for the per-epoch policy. createEpoch
    ///         is permissionless, so whoever wins the cadence race must not be
    ///         able to shrink the committee or widen the lottery below what
    ///         the deployment intends.
    uint16 public immutable MIN_THRESHOLD;
    uint16 public immutable MIN_COMMITTEE_SIZE;
    uint16 public immutable MAX_LOTTERY_ALPHA_BPS;
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
    /// @dev Number of ciphertexts submitted per (epoch, aid); the next index.
    mapping(bytes12 epochId => mapping(bytes32 aid => uint16 count)) internal ciphertextCounts;
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => DKGTypes.CombinedDecryptionRecord combined))) internal
        epochCombinedDecryptions;
    /// @dev Stores keccak256 over the canonical (recipientIndexes ‖ recipientPubKeys)
    /// section of any valid submitContribution transcript for this epoch. Set once at
    /// selectParticipants. Lets submitContribution verify the entire 96-word committee
    /// section in one keccak instead of 32 storage reads + 32 external registry calls.
    mapping(bytes12 epochId => bytes32 prefixHash) internal epochContribPrefixHash;

    // ─── Pool keys ──────────────────────────────────────────────────────────
    //
    // Every epoch deals `MAX_K` keys. A key becomes usable in two steps:
    // `activatePoolKey` proves it (writing `poolKeys` + `poolShareRoots` and
    // setting its bit in `poolActivated`), and `claimPoolKey` — driven by the
    // app manager at registration — hands the next activated one to an
    // application. `poolNext` only ever moves forward, so an epoch serves at
    // most `MAX_K` applications and `createEpoch` opens early once it is
    // nearly spent.

    /// @dev `P_j = Σ_d A_{d,j,0}`, written exactly once per key by
    ///      `activatePoolKey` from the proof-verified `aggregateCommitments[0]`.
    mapping(bytes12 epochId => mapping(uint8 keyIndex => DKGTypes.Point key)) internal poolKeys;
    /// @dev keccak Merkle root (depth `MERKLE_DEPTH`, `MAX_N` leaves) over the
    ///      committee's share commitments `D_i` for one pool key. Leaf
    ///      `participantIndex - 1` is `keccak256(0x00 ‖ D.x ‖ D.y)` for every
    ///      committee member, `MERKLE_EMPTY_LEAF` beyond `committeeSize`, and
    ///      internal nodes are `keccak256(0x01 ‖ left ‖ right)`. One SSTORE
    ///      replaces the per-participant list; the pre-images travel in the
    ///      activation calldata.
    mapping(bytes12 epochId => mapping(uint8 keyIndex => bytes32 root)) internal poolShareRoots;
    /// @dev Bit `j` is set once pool key `j` of the epoch has been activated.
    mapping(bytes12 epochId => uint8 bitmap) internal poolActivated;
    /// @dev Index of the pool key the next registration claims. Reaching
    ///      `MAX_K` exhausts the epoch.
    mapping(bytes12 epochId => uint8 nextIndex) internal poolNext;
    /// @dev `keyIndex + 1` of the pool key an application claimed; 0 means it
    ///      never claimed one (the +1 keeps key 0 distinguishable from unset).
    mapping(bytes12 epochId => mapping(bytes32 aid => uint8 claimedIndexPlusOne)) internal appPoolIndexes;

    /// @dev _hash4(c1x, c1y, c2x, c2y) for each ciphertext submitted to a
    ///      epoch. Written once per (epochId, aid, ciphertextIndex) by submitCiphertext and
    ///      verified by combineDecryption to bind the combine proof to the authoritative
    ///      on-chain ciphertext (preventing a combiner from swapping in a different ct).
    ///      The raw coordinates are available via the CiphertextSubmitted event log.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => bytes32 ciphertextHash))) internal _ciphertexts;

    // BabyJubJub curve constants moved to libraries/BabyJubJub.sol — point
    // validation flows through `_requireValidEncryptionPoint` which calls
    // `BabyJubJub.{isCanonical,isOnCurve,isIdentity,isInPrimeSubgroup}`.

    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = DKGProtocol.DOMAIN_CONTRIBUTION_TRANSCRIPT_V1;
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = DKGProtocol.DOMAIN_DECRYPT_COMBINE_TRANSCRIPT_V1;
    bytes32 internal constant POOLKEY_TRANSCRIPT_DOMAIN = DKGProtocol.DOMAIN_POOLKEY_TRANSCRIPT_V1;

    constructor(
        uint32 _chainId,
        address _registry,
        address _contributionVerifier,
        address _partialDecryptVerifier,
        address _poolKeyVerifier,
        address _decryptCombineVerifier,
        uint256 _epochDurationBlocks,
        uint256 _committeeSelectionBlocks,
        uint256 _keyAssemblyBlocks,
        uint256 _finalizeGapBlocks,
        uint16 _minThreshold,
        uint16 _minCommitteeSize,
        uint16 _maxLotteryAlphaBps
    ) {
        if (uint32(block.chainid) != _chainId) revert InvalidChainId();
        if (_registry == address(0)) revert InvalidAddress();
        if (
            _contributionVerifier == address(0) || _partialDecryptVerifier == address(0) || _poolKeyVerifier == address(0)
                || _decryptCombineVerifier == address(0)
        ) revert InvalidVerifier();
        CHAIN_ID = _chainId;
        REGISTRY = _registry;
        EPOCH_PREFIX = DKGIdLib.getPrefix(_chainId, address(this));
        CONTRIBUTION_VERIFIER = _contributionVerifier;
        PARTIAL_DECRYPT_VERIFIER = _partialDecryptVerifier;
        POOL_KEY_VERIFIER = _poolKeyVerifier;
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

        MIN_THRESHOLD         = _minThreshold == 0 ? 1 : _minThreshold;
        MIN_COMMITTEE_SIZE    = _minCommitteeSize == 0 ? 1 : _minCommitteeSize;
        MAX_LOTTERY_ALPHA_BPS = _maxLotteryAlphaBps == 0 ? type(uint16).max : _maxLotteryAlphaBps;
        if (MIN_THRESHOLD > MIN_COMMITTEE_SIZE || MIN_COMMITTEE_SIZE > MAX_N || MAX_LOTTERY_ALPHA_BPS < 10000) {
            revert InvalidPolicy();
        }

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
    /// @return                        The 12-byte epoch identifier
    ///                                `uint32 prefix || uint64 nonce`.
    function createEpoch(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps
    ) external returns (bytes12) {
        // Permissionless cadence guard. The first epoch (lastEpochStartBlock
        // == 0) goes through immediately; subsequent epochs require one full
        // EPOCH_DURATION_BLOCKS to have elapsed since the previous start —
        // unless the newest epoch is Live and nearly out of pool keys, or
        // Aborted. An epoch serves at most MAX_K applications, so waiting for
        // the cadence with an exhausted pool would make registrations revert
        // PoolExhausted with nowhere else to go; an aborted epoch serves
        // nobody at all. An epoch still in preparation keeps the cadence.
        if (block.number < nextEpochStartBlock()) {
            bytes12 newest = DKGIdLib.computeEpochId(EPOCH_PREFIX, epochNonce);
            DKGTypes.EpochPhase phase = epochs[newest].status;
            bool spent = phase == DKGTypes.EpochPhase.Live && poolNext[newest] >= MAX_K - 1;
            if (!spent && phase != DKGTypes.EpochPhase.Aborted) revert InvalidPhase();
        }

        if (
            threshold == 0 || committeeSize == 0 || threshold > committeeSize
                || committeeSize > MAX_N
                || minValidContributions == 0 || minValidContributions > committeeSize
                || minValidContributions < threshold
                || lotteryAlphaBps < 10000
                || threshold < MIN_THRESHOLD || committeeSize < MIN_COMMITTEE_SIZE
                || lotteryAlphaBps > MAX_LOTTERY_ALPHA_BPS
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

        // Snapshot the currently ACTIVE node count and derive the per-node lottery
        // threshold so that on average `lotteryAlpha × committeeSize` live nodes pass.
        // Using activeCount (rather than nodeCount) keeps the lottery denominator
        // aligned with the set of nodes that can actually produce a contribution —
        // reaped stragglers are automatically excluded.
        uint256 registered = uint256(IDKGRegistry(REGISTRY).activeCount());
        if (registered == 0) revert InvalidPolicy();
        // numerator = α·n in basis points; the admissible fraction of the
        // hash space is numerator / (10000 · registered).
        uint256 numerator = uint256(lotteryAlphaBps) * uint256(committeeSize);
        uint256 denominator = 10000 * registered;
        uint256 lotteryThreshold;
        if (numerator >= denominator) {
            // α·n ≥ registered: everyone passes.
            lotteryThreshold = type(uint256).max;
        } else {
            // threshold = 2^256 · numerator / denominator, computed as
            // (max / denominator) · numerator which cannot overflow because
            // numerator < denominator.
            lotteryThreshold = (type(uint256).max / denominator) * numerator;
        }

        epochNonce++;
        bytes12 epochId = DKGIdLib.computeEpochId(EPOCH_PREFIX, epochNonce);

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
        // Only identities that existed when the epoch was created may claim:
        // otherwise fresh addresses could be ground against the revealed seed.
        if (node.registeredAtBlock >= epoch.startBlock) revert NotInSnapshot();

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
    /// O(1) committee verification in submitContribution. The same keys are
    /// published (unpadded, slot order) in the `CommitteeSnapshot` event so
    /// off-chain consumers read the frozen view instead of the live registry.
    function _snapshotCommittee(bytes12 epochId, uint16 committeeSize) internal {
        uint256[MAX_N] memory canonicalIdxs;
        uint256[2 * MAX_N] memory canonicalKeys;
        uint256[] memory snapshotKeys = new uint256[](2 * committeeSize);
        address[] storage participants = epochParticipants[epochId];
        for (uint256 i = 0; i < committeeSize; i++) {
            canonicalIdxs[i] = i + 1;
            IDKGRegistry.NodeKey memory node = IDKGRegistry(REGISTRY).getNode(participants[i]);
            canonicalKeys[i * 2] = node.pubX;
            canonicalKeys[i * 2 + 1] = node.pubY;
            snapshotKeys[i * 2] = node.pubX;
            snapshotKeys[i * 2 + 1] = node.pubY;
        }
        for (uint256 i = committeeSize; i < MAX_N; i++) {
            canonicalKeys[i * 2 + 1] = 1; // identity-pad unused slots
        }
        epochContribPrefixHash[epochId] = keccak256(abi.encodePacked(canonicalIdxs, canonicalKeys));
        emit CommitteeSnapshot(epochId, committeeSize, snapshotKeys);
    }

    /// @notice Submit a contributor's polynomial commitments, encrypted
    ///         shares and Groth16 proof of correctness.
    /// @dev    The committee membership + BabyJubJub public keys are verified
    ///         against a single `keccak256` snapshot taken when the lottery
    ///         filled; the transcript is read straight from calldata via the
    ///         BRLC helper. The transaction reverts if the proof fails.
    ///
    ///         Direct EOA calls only: nodes and the SDK rebuild every
    ///         contribution from the outer transaction calldata, which an
    ///         internal (contract-relayed) call never carries — such a
    ///         contribution would be accepted on chain yet unrecoverable,
    ///         and one relaying committee member blocks every key
    ///         activation of the epoch. The `code.length` clause also
    ///         excludes EIP-7702 delegated accounts, whose outer calldata is
    ///         a batch, not the DKG call.
    function submitContribution(
        bytes12 epochId,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        if (msg.sender != tx.origin || msg.sender.code.length != 0) revert DirectCallRequired();
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
        // Transcript layout (3KN + 5N words; see Sizes.sol):
        //   words [0..2KN)          commitments, key-major (K keys × N points × 2)
        //   words [2KN..2KN+N)      recipientIndexes
        //   words [2KN+N..2KN+3N)   recipientPubKeys  (N points × 2 coords)
        //   words [2KN+3N..2KN+5N)  ephemerals
        //   words [2KN+5N..3KN+5N)  maskedShares, key-major
        if (transcript.length != CONTRIB_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        // The BRLC challenge is Fiat–Shamir over the calldata itself. Deriving
        // it from the prover's own digests would let the prover know ρ before
        // choosing the calldata and publish a transcript that differs from the
        // proven witness while preserving Σ ρ^i·v_i.
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            CONTRIBUTION_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(commitmentsHash, encryptedSharesHash, keccak256(transcript)))
        );
        if (publicInputs[6] != challenge) revert InvalidProofInput();
        // publicInputs[7] = TranscriptCommitment (verified below via BRLC)
        IZKVerifier(CONTRIBUTION_VERIFIER).verifyProof(proof, input);

        // Single-shot committee verification: the bytes spanning
        // recipientIndexes ‖ recipientPubKeys hold the canonical committee
        // section. Compare against the hash snapshotted when the lottery
        // filled. This replaces a per-recipient loop with N storage reads + N
        // external registry calls.
        if (
            keccak256(transcript[CONTRIB_COMMITTEE_BYTES_OFFSET:CONTRIB_COMMITTEE_BYTES_END])
                != epochContribPrefixHash[epochId]
        ) revert InvalidProofInput();

        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, CONTRIB_TRANSCRIPT_WORDS) != publicInputs[7]) revert InvalidProofInput();

        // Only persist the fields the contract itself actually needs:
        //   - commitmentsHash: the Poseidon digest over all MAX_K commitment
        //     vectors, re-checked per participant row by `activatePoolKey`;
        //   - contributorIndex + accepted: identity & dup-prevention gates.
        // encryptedSharesHash and the redundant `contributor` are only emitted
        // in the event below; off-chain consumers read them from logs.
        DKGTypes.ContributionRecord storage rec = epochContributions[epochId][msg.sender];
        rec.contributorIndex = contributorIndex;
        rec.commitmentsHash = commitmentsHash;
        rec.accepted = true;
        epoch.contributionCount++;

        // Refresh the contributor's liveness timestamp on the registry.
        // A successful proof-gated contribution is the strongest possible
        // signal that the operator is alive and well-configured.
        IDKGRegistry(REGISTRY).markActive(msg.sender);

        emit ContributionSubmitted(epochId, msg.sender, contributorIndex, commitmentsHash, encryptedSharesHash);
    }

    /// @notice Freeze the epoch's accepted contributor set and open the
    ///         Service window.
    /// @dev    Proof-less by design: with `MAX_K` keys per epoch there is no
    ///         single "collective key" to publish here. Each key is proven
    ///         independently by `activatePoolKey`, which re-checks every
    ///         participant row against the `commitmentsHash` stored at
    ///         contribution time — so the contributor set this call freezes is
    ///         exactly the set every activation must aggregate over.
    ///         Callable by anyone once the epoch reached `liveNotBeforeBlock`
    ///         with at least `policy.minValidContributions` contributions.
    function finalizeEpoch(bytes12 epochId) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status == DKGTypes.EpochPhase.Live) revert AlreadyLive();
        if (epoch.status != DKGTypes.EpochPhase.KeyAssembly) revert InvalidPhase();
        // liveNotBeforeBlock gate — semantically a "phase not yet open"
        // condition, so we reuse InvalidPhase to keep the contract small.
        if (block.number < uint256(epoch.policy.liveNotBeforeBlock)) revert InvalidPhase();
        uint16 accepted = epoch.contributionCount;
        if (accepted < epoch.policy.minValidContributions) revert InsufficientContributions();

        epoch.status = DKGTypes.EpochPhase.Live;
        emit EpochLive(epochId, accepted);
    }

    /// @notice Prove and store one of the epoch's `MAX_K` pool keys.
    /// @dev    Permissionless, one call per key, in any order, only while the
    ///         epoch is Live. The proof attests that, over the accepted
    ///         contributors listed in the transcript, each contributor's
    ///         on-chain `commitmentsHash` is reproduced from its commitments
    ///         for this key plus the digests of its other keys, that
    ///         `aggregate[m] = Σ_i A_i[j][m]`, and that
    ///         `D_p = Σ_m p^m · aggregate[m]` for every committee member `p`.
    ///         The contract independently binds every row to an accepted
    ///         on-chain contribution, rejects duplicate participant indexes,
    ///         requires the unused rows and share slots to be blank, and
    ///         re-derives the BRLC commitment from calldata so the published
    ///         transcript is the one that was proven.
    /// @param  keyIndex         Which of the `MAX_K` keys this activates.
    /// @param  transcriptDigest The prover's digest of the witness transcript
    ///                          (public input 5). It enters the challenge
    ///                          anchor so the witness is fixed before ρ exists.
    /// @param  transcript       `POOLKEY_TRANSCRIPT_WORDS` words; layout in
    ///                          Sizes.sol.
    /// @param  input            `[eid, t, n, acceptedCount, keyIndex,
    ///                          transcriptDigest, challenge,
    ///                          transcriptCommitment]`.
    /// @dev    Direct EOA calls only, exactly like submitContribution: the
    ///         committee reconstructs every activation (pool key, share
    ///         commitments, Merkle root pre-images) from the outer
    ///         transaction calldata, so a contract-relayed activation would
    ///         be valid on chain yet unrecoverable. `code.length` also
    ///         excludes EIP-7702 delegated accounts (batch outer calldata).
    function activatePoolKey(
        bytes12 epochId,
        uint8 keyIndex,
        bytes32 transcriptDigest,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        if (msg.sender != tx.origin || msg.sender.code.length != 0) revert DirectCallRequired();
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();
        if (keyIndex >= MAX_K) revert InvalidProofInput();
        uint8 keyBit = uint8(1) << keyIndex;
        if (poolActivated[epochId] & keyBit != 0) revert PoolKeyAlreadyActive();
        if (transcript.length != POOLKEY_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();

        // Cheap public-input checks first; only invoke the expensive verifier
        // once the proof provably targets this epoch and key.
        uint256[8] memory publicInputs = abi.decode(input, (uint256[8]));
        if (
            publicInputs[0] != _epochScalar(epochId) || publicInputs[1] != epoch.policy.threshold
                || publicInputs[2] != epoch.policy.committeeSize || publicInputs[3] != epoch.contributionCount
                || publicInputs[4] != uint256(keyIndex) || bytes32(publicInputs[5]) != transcriptDigest
        ) revert InvalidProofInput();
        // Fiat–Shamir over the prover's transcript digest AND the calldata
        // (see submitContribution): the digest fixes the witness before ρ
        // exists, the calldata hash pins the published words to it.
        uint256 challenge = BRLC.deriveChallenge(
            epochId, POOLKEY_TRANSCRIPT_DOMAIN, keccak256(abi.encodePacked(transcriptDigest, keccak256(transcript)))
        );
        if (publicInputs[6] != challenge) revert InvalidProofInput();
        IZKVerifier(POOL_KEY_VERIFIER).verifyProof(proof, input);

        bytes32 root = _verifyPoolKeyRows(epochId, epoch, transcript);

        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, POOLKEY_TRANSCRIPT_WORDS) != publicInputs[7]) {
            revert InvalidProofInput();
        }

        // The pool key is `aggregateCommitments[0]`. Read it straight from the
        // (now fully bound) transcript calldata and persist it once.
        uint256 pkX;
        uint256 pkY;
        {
            uint256 aggBase = dOff + POOLKEY_AGG_WORDS_OFFSET * 32;
            assembly ("memory-safe") {
                pkX := calldataload(aggBase)
                pkY := calldataload(add(aggBase, 0x20))
            }
        }
        DKGTypes.Point storage stored = poolKeys[epochId][keyIndex];
        stored.x = pkX;
        stored.y = pkY;
        poolShareRoots[epochId][keyIndex] = root;
        poolActivated[epochId] = poolActivated[epochId] | keyBit;

        emit PoolKeyActivated(epochId, keyIndex, pkX, pkY);
    }

    /// @dev Validates the rows of a pool-key activation transcript against
    ///      on-chain contribution state and returns the Merkle root over the
    ///      share commitments.
    ///
    ///      For every active row `i < contributionCount`: the participant
    ///      index is in `[1, committeeSize]` and unique, names an accepted
    ///      contributor registered under that very index, and carries that
    ///      contributor's stored Poseidon `commitmentsHash`. Without the
    ///      uniqueness check a prover could repeat one contributor's row and
    ///      omit another, aggregating over the wrong set. Every row
    ///      `i >= contributionCount` must be blank (index and hash both 0).
    ///
    ///      Share slot `i` belongs to committee member `i + 1`, whether it
    ///      contributed or not: leaf `i` is `keccak256(0x00 ‖ D.x ‖ D.y)` for
    ///      `i < committeeSize`, and the slot must hold the identity `(0, 1)`
    ///      with leaf `MERKLE_EMPTY_LEAF` beyond that. Internal nodes are
    ///      `keccak256(0x01 ‖ left ‖ right)` — the tags keep leaves and
    ///      nodes in separate domains. The tree is a fixed
    ///      `MERKLE_DEPTH`-level fold over `MAX_N` leaves, so a partial
    ///      decryption later proves its `D_i` with a `MERKLE_DEPTH`-long path.
    function _verifyPoolKeyRows(
        bytes12 epochId,
        Epoch storage epoch,
        bytes calldata transcript
    ) internal view returns (bytes32) {
        uint256 dOff;
        assembly { dOff := transcript.offset }
        uint256 piBase = dOff;                                     // participantIndexes
        uint256 chBase = dOff + POOLKEY_HASHES_BYTES_OFFSET;       // contributionHashes
        uint256 scBase = dOff + POOLKEY_SHARE_WORDS_OFFSET * 32;   // shareCommitments

        uint256 ccount = epoch.contributionCount;
        uint256 cSize = epoch.policy.committeeSize;
        uint256 seenIndexes;
        for (uint256 i = 0; i < MAX_N; i++) {
            uint256 pIdx;
            uint256 rowHash;
            assembly ("memory-safe") {
                pIdx := calldataload(add(piBase, mul(i, 0x20)))
                rowHash := calldataload(add(chBase, mul(i, 0x20)))
            }
            if (i >= ccount) {
                if (pIdx != 0 || rowHash != 0) revert InvalidProofInput();
                continue;
            }
            if (pIdx == 0 || pIdx > cSize) revert InvalidProofInput();
            uint256 indexBit = uint256(1) << pIdx;
            if (seenIndexes & indexBit != 0) revert InvalidProofInput();
            seenIndexes |= indexBit;

            address participant = epochParticipants[epochId][pIdx - 1];
            DKGTypes.ContributionRecord storage contribution = epochContributions[epochId][participant];
            if (!contribution.accepted || contribution.contributorIndex != uint16(pIdx)) revert InvalidProofInput();
            if (bytes32(rowHash) != contribution.commitmentsHash) revert InvalidProofInput();
        }

        bytes32[MAX_N] memory nodes;
        for (uint256 i = 0; i < MAX_N; i++) {
            uint256 scX;
            uint256 scY;
            assembly ("memory-safe") {
                scX := calldataload(add(scBase, mul(i, 0x40)))
                scY := calldataload(add(scBase, add(mul(i, 0x40), 0x20)))
            }
            if (i < cSize) {
                nodes[i] = _hashLeaf(scX, scY);
            } else {
                if (scX != 0 || scY != 1) revert InvalidProofInput();
                nodes[i] = MERKLE_EMPTY_LEAF;
            }
        }
        uint256 width = MAX_N;
        for (uint256 level = 0; level < MERKLE_DEPTH; level++) {
            width >>= 1;
            for (uint256 i = 0; i < width; i++) {
                nodes[i] = _hashNode(nodes[i * 2], nodes[i * 2 + 1]);
            }
        }
        return nodes[0];
    }

    /// @notice Assign the next unclaimed pool key of `epochId` to `aid`.
    /// @dev    Callable only by the sibling app manager, which drives it from
    ///         `registerApplication` after it has validated the epoch phase,
    ///         the aid and the policy. Reverts `PoolExhausted` when all
    ///         `MAX_K` keys are taken and `PoolKeyNotActive` when the next one
    ///         has not been proven yet — an application must never end up
    ///         holding a key nobody can decrypt under.
    function claimPoolKey(bytes12 epochId, bytes32 aid) external returns (uint8 keyIndex) {
        if (msg.sender != appManager) revert Unauthorized();
        keyIndex = poolNext[epochId];
        if (keyIndex >= MAX_K) revert PoolExhausted();
        if (poolActivated[epochId] & (uint8(1) << keyIndex) == 0) revert PoolKeyNotActive();
        unchecked {
            poolNext[epochId] = keyIndex + 1;
            appPoolIndexes[epochId][aid] = keyIndex + 1;
        }
        emit PoolKeyClaimed(epochId, aid, keyIndex);
    }

    /// @notice The activated pool key `P_j` of an epoch.
    function getPoolKey(bytes12 epochId, uint8 keyIndex) external view returns (uint256, uint256) {
        if (keyIndex >= MAX_K || poolActivated[epochId] & (uint8(1) << keyIndex) == 0) revert PoolKeyNotActive();
        DKGTypes.Point storage key = poolKeys[epochId][keyIndex];
        return (key.x, key.y);
    }

    /// @notice `nextIndex` is the key the next registration claims (`MAX_K`
    ///         once the pool is spent); `activated` is the bitmap of proven
    ///         keys, bit `j` for key `j`.
    function getPoolStatus(bytes12 epochId) external view returns (uint8 nextIndex, uint8 activated) {
        return (poolNext[epochId], poolActivated[epochId]);
    }

    /// @notice Merkle root over the share commitments of one pool key;
    ///         `bytes32(0)` while the key is not activated.
    function getPoolShareRoot(bytes12 epochId, uint8 keyIndex) external view returns (bytes32) {
        return poolShareRoots[epochId][keyIndex];
    }

    /// @notice The pool key claimed by an application.
    function getAppPoolIndex(bytes12 epochId, bytes32 aid) external view returns (uint8) {
        return _appPoolIndex(epochId, aid);
    }

    /// @dev Unwraps the `+1` marker; reverts when the aid never claimed a key.
    function _appPoolIndex(bytes12 epochId, bytes32 aid) internal view returns (uint8) {
        uint8 claimed = appPoolIndexes[epochId][aid];
        if (claimed == 0) revert PoolKeyNotActive();
        unchecked { return claimed - 1; }
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
        uint256 piBase = dOff + COMBINE_INDEXES_BYTES_OFFSET;         // participantIndexes start
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
            if (pIdxRaw == 0 || pIdxRaw > cs) revert InvalidProofInput();
            uint16 participantIndex = uint16(pIdxRaw);
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

    /// @notice Submit a committee member's partial decryption
    ///         `δ_i = e_{j,i} · C_1` under the application's pool key.
    /// @dev    Keyed by `(epochId, aid, ciphertextIndex, participantIndex)` to
    ///         support multiple ciphertexts and applications per epoch. The
    ///         Groth16 proof is a Chaum–Pedersen DLEQ establishing that `δ_i`
    ///         and the committed share `D_i` share a discrete log with respect
    ///         to `C_1` and `G`.
    /// @dev `aid` binds the proof transcript to a specific application; only
    ///      registered applications can own a ciphertext, so an unknown aid
    ///      fails the ciphertext binding below.
    /// @dev `c1x/c1y/c2x/c2y` are the raw ciphertext coordinates as
    ///      submitted via submitCiphertext. The contract verifies
    ///      `keccak256(abi.encode(...))` matches the stored ciphertext
    ///      hash and then binds the proof's public-input C1 (pi[4..5])
    ///      to the authoritative on-chain ciphertext.
    /// @param shareProof `MERKLE_DEPTH` siblings, bottom-up, proving that
    ///        `keccak256(0x00 ‖ D_i.x ‖ D_i.y)` (public inputs 6/7) sits at
    ///        leaf `participantIndex - 1` of the share root of the
    ///        application's pool key. Every committee member has a leaf,
    ///        contributor or not. This is what pins the member to the share
    ///        it was dealt FOR THIS APPLICATION — a δ computed under another
    ///        key can never produce a path into this root.
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
        // Outside the application's decryption window nothing may leak: in
        // Automatic mode `t` partials alone already decrypt off chain.
        IDKGAppManager(appManager).requireDecryptionOpen(epochId, aid);

        uint256 indexBit = uint256(1) << participantIndex;
        if (epochPartialBitmap[epochId][aid][ciphertextIndex] & indexBit != 0) {
            revert AlreadyPartiallyDecrypted();
        }

        // Layout (15 words): [eid, aid, ctIdx, i, C1.x, C1.y, D_i.x, D_i.y,
        // delta.x, delta.y, A1.x, A1.y, A2.x, A2.y, response].
        // Cheap public-input checks fail before the expensive verifier call.
        uint256[15] memory publicInputs = abi.decode(input, (uint256[15]));
        if (
            publicInputs[0] != _epochScalar(epochId)
                || publicInputs[1] != uint256(aid)
                || publicInputs[2] != ciphertextIndex
                || publicInputs[3] != participantIndex
                // pi[4..5] = base point (C_1) — bind to the just-verified
                // on-chain ciphertext.
                || publicInputs[4] != c1x
                || publicInputs[5] != c1y
        ) revert InvalidProofInput();
        if (deltaHash != keccak256(abi.encodePacked(publicInputs[8], publicInputs[9]))) revert InvalidProofInput();
        // pi[6..7] = D_i, the member's share commitment for THIS application's
        // pool key. Proven against the root activatePoolKey stored.
        _verifyShareProof(
            poolShareRoots[epochId][_appPoolIndex(epochId, aid)],
            _hashLeaf(publicInputs[6], publicInputs[7]),
            participantIndex,
            shareProof
        );
        IZKVerifier(PARTIAL_DECRYPT_VERIFIER).verifyProof(proof, input);

        // Persist the δ commitment as a single 32-byte hash plus a bitmap bit.
        // The combine path reads δ.x/δ.y back from the proof transcript and
        // re-hashes against the stored value, so we don't store the raw point.
        epochPartialDeltaHash[epochId][aid][ciphertextIndex][participantIndex] = deltaHash;
        epochPartialBitmap[epochId][aid][ciphertextIndex] |= indexBit;
        unchecked { epoch.partialDecryptionCount++; } // informational only

        emit PartialDecryptionSubmitted(
            epochId, aid, msg.sender, participantIndex, ciphertextIndex,
            publicInputs[8], publicInputs[9]
        );
    }

    /// @dev Verify a `MERKLE_DEPTH`-long inclusion path for `leaf` at leaf
    ///      index `participantIndex - 1` against `root`. Siblings are ordered
    ///      bottom-up and combined as `keccak256(0x01 ‖ left ‖ right)`,
    ///      matching the fold in `_verifyPoolKeyRows`.
    function _verifyShareProof(
        bytes32 root,
        bytes32 leaf,
        uint16 participantIndex,
        bytes32[] calldata shareProof
    ) internal pure {
        if (shareProof.length != MERKLE_DEPTH) revert InvalidProofInput();
        uint256 index;
        unchecked { index = uint256(participantIndex) - 1; }
        bytes32 node = leaf;
        for (uint256 level = 0; level < MERKLE_DEPTH; level++) {
            bytes32 sibling = shareProof[level];
            node = index & 1 == 0 ? _hashNode(node, sibling) : _hashNode(sibling, node);
            index >>= 1;
        }
        if (node != root) revert InvalidProofInput();
    }

    /// @notice Submit a ciphertext to be threshold-decrypted by the committee
    ///         under the application key `PK_aid`.
    /// @dev    The index is assigned on chain per (epoch, aid), so a caller
    ///         cannot squat on or front-run a specific slot. The stored value is
    ///         `keccak256(c1x, c1y, c2x, c2y)`, which binds the partial and
    ///         combine proofs to this exact ciphertext.
    ///
    ///         There is no proof of knowledge of the randomness `r`: the
    ///         submitter of a homomorphically aggregated ciphertext cannot
    ///         know it. Cross-application replay is instead prevented by the
    ///         per-application pool key — every application's partials are
    ///         computed under its own `P_j`, so decrypting a copied `C_1`
    ///         under another application yields a value under a different key
    ///         and is useless.
    ///
    ///         `aid` must name an application registered on the sibling app
    ///         manager, which also enforces the per-application submission
    ///         policy (authorized submitter, block window, ciphertext cap).
    function submitCiphertext(
        bytes12 epochId,
        bytes32 aid,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y
    ) external returns (uint16 ciphertextIndex) {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Live) revert InvalidPhase();

        // Well-formedness: coords must be canonical (< Q), on-curve and
        // non-identity. Prime-subgroup membership of C_1 is NOT checked here
        // — see `_requireValidEncryptionPoint`.
        _requireValidEncryptionPoint(c1x, c1y);
        _requireValidEncryptionPoint(c2x, c2y);

        ciphertextIndex = ciphertextCounts[epochId][aid] + 1;
        if (ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert DecryptionLimitReached();
        if (!_appManagerSet) revert AppManagerNotSet();
        // Reverts InvalidApplication for an unregistered (or zero) aid.
        IDKGAppManager(appManager).requireCanSubmitCiphertext(epochId, aid, ciphertextIndex, msg.sender);
        ciphertextCounts[epochId][aid] = ciphertextIndex;
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
    ///      (ii) the application's submission policy / `maxCiphertexts`
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
    ///         application registered against `epochId`. There is no
    ///         per-ciphertext organizer artefact any more: the circuit proves
    ///         knowledge of `sk_org` with `PK_org == sk_org·G` and
    ///         `Δ = sk_org·C_1` internally, and this contract only pins the
    ///         transcript's `PK_org` to the application's registered key.
    ///         `Automatic` applications carry the identity `(0, 1)` there, so
    ///         the same statement covers both modes.
    /// @dev    Public-input layout (9 fields): eid, aid, ctIdx, threshold,
    ///         shareCount, combineHash, plaintext, challenge,
    ///         transcriptCommitment. Matches the circuit definition in
    ///         `circuits/decryptcombine/circuit.go`.
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
        IDKGAppManager(appManager).requireDecryptionOpen(epochId, aid);
        if (_popcount(epochPartialBitmap[epochId][aid][ciphertextIndex]) < epoch.policy.threshold) {
            revert InsufficientPartialDecryptions();
        }

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

    /// @dev Validates the 9-word public-input vector against
    /// eid/aid/ctIdx/threshold/combineHash/plaintext/challenge, range-checks
    /// shareCount, binds the transcript's first 128 bytes to the stored
    /// ciphertext hash and its organizer words to the application's registered
    /// `PK_org`, and verifies the BRLC commitment over the full transcript
    /// region. Split out of combineDecryption to keep the parent's stack
    /// within Yul's depth limit.
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
        // Layout (9 words): [eid, aid, ctIdx, threshold, shareCount,
        // combineHash, plaintext, challenge, transcriptCommitment].
        uint256[9] memory publicInputs = abi.decode(input, (uint256[9]));
        if (
            publicInputs[0] != _epochScalar(epochId)
                || publicInputs[1] != uint256(aid)
                || publicInputs[2] != uint256(ciphertextIndex)
                || publicInputs[3] != epoch.policy.threshold
                || bytes32(publicInputs[5]) != combineHash
                || publicInputs[6] != plaintext
        ) revert InvalidProofInput();
        if (publicInputs[4] < epoch.policy.threshold) revert InvalidProofInput();
        if (publicInputs[4] > MAX_N) revert InvalidProofInput();
        if (transcript.length != COMBINE_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        if (keccak256(transcript[0:128]) != storedCtHash) revert InvalidProofInput();
        _verifyOrganizerKey(epochId, aid, transcript);
        // Fiat–Shamir over the calldata (see submitContribution).
        uint256 challenge = BRLC.deriveChallenge(
            epochId,
            DECRYPT_COMBINE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(combineHash, bytes32(plaintext), keccak256(transcript)))
        );
        if (publicInputs[7] != challenge) revert InvalidProofInput();
        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, COMBINE_TRANSCRIPT_WORDS) != publicInputs[8]) revert InvalidProofInput();
        return publicInputs[4];
    }

    /// @dev Binds `w[4..5]` of the combine transcript to the application's
    ///      registered `PK_org`, so a combiner cannot swap in an organizer key
    ///      it controls and prove the statement against that instead. For
    ///      `Automatic` applications the registered key is the identity
    ///      `(0, 1)` and the circuit's `OrganizerSecret` is 0, so the same
    ///      binding covers both modes with no special case.
    function _verifyOrganizerKey(bytes12 epochId, bytes32 aid, bytes calldata transcript) internal view {
        if (!_appManagerSet) revert AppManagerNotSet();
        uint256 wx;
        uint256 wy;
        assembly ("memory-safe") {
            wx := calldataload(add(transcript.offset, 0x80))
            wy := calldataload(add(transcript.offset, 0xa0))
        }
        (uint256 pkOrgX, uint256 pkOrgY) = IDKGAppManager(appManager).getOrganizerPK(epochId, aid);
        if (wx != pkOrgX || wy != pkOrgY) revert InvalidProofInput();
    }

    /// @notice Record that an epoch is dead. Permissionless, and only possible
    ///         once the epoch provably cannot progress: the committee did not
    ///         fill before the selection deadline, or fewer than
    ///         `minValidContributions` arrived before the key-assembly
    ///         deadline. Nobody — not even the creator — can abort an epoch
    ///         that can still be finalized.
    /// @param  epochId The epoch identifier.
    function abortEpoch(bytes12 epochId) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        bool dead = (
            epoch.status == DKGTypes.EpochPhase.CommitteeSelection
                && block.number > uint256(epoch.policy.committeeSelectionDeadlineBlock)
        ) || (
            epoch.status == DKGTypes.EpochPhase.KeyAssembly
                && block.number > uint256(epoch.policy.keyAssemblyDeadlineBlock)
                && epoch.contributionCount < epoch.policy.minValidContributions
        );
        if (!dead) revert InvalidPhase();

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

    /// @notice Number of ciphertexts submitted so far for (epochId, aid).
    function ciphertextCount(bytes12 epochId, bytes32 aid) external view returns (uint16) {
        return ciphertextCounts[epochId][aid];
    }

    function _popcount(uint256 x) internal pure returns (uint256 n) {
        while (x != 0) {
            x &= x - 1;
            n++;
        }
    }

    function getContributionVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(CONTRIBUTION_VERIFIER).provingKeyHash();
    }

    function getPartialDecryptVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(PARTIAL_DECRYPT_VERIFIER).provingKeyHash();
    }

    function getPoolKeyVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(POOL_KEY_VERIFIER).provingKeyHash();
    }

    function getDecryptCombineVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(DECRYPT_COMBINE_VERIFIER).provingKeyHash();
    }

    function _epochScalar(bytes12 epochId) internal pure returns (uint256) {
        return uint256(uint96(epochId));
    }

    /// @dev Share-commitment leaf `keccak256(0x00 ‖ x ‖ y)` (65 bytes), written
    ///      into scratch memory to skip the abi.encodePacked allocation.
    function _hashLeaf(uint256 x, uint256 y) internal pure returns (bytes32 h) {
        assembly ("memory-safe") {
            let p := mload(0x40)
            mstore8(p, 0x00)
            mstore(add(p, 0x01), x)
            mstore(add(p, 0x21), y)
            h := keccak256(p, 0x41)
        }
    }

    /// @dev Merkle internal node `keccak256(0x01 ‖ left ‖ right)` (65 bytes).
    function _hashNode(bytes32 left, bytes32 right) internal pure returns (bytes32 h) {
        assembly ("memory-safe") {
            let p := mload(0x40)
            mstore8(p, 0x01)
            mstore(add(p, 0x01), left)
            mstore(add(p, 0x21), right)
            h := keccak256(p, 0x41)
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
