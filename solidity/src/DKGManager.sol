// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGIdLib} from "./libraries/DKGIdLib.sol";
import {BRLC} from "./libraries/BRLC.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";
import {PhaseLib} from "./libraries/PhaseLib.sol";
import {PoseidonT5} from "poseidon-solidity/PoseidonT5.sol";

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
    // ──────────────────────────────────────────────────────────────────────
    // Single source of truth for the per-circuit array bound.
    //
    // MAX_N is the only number to edit when changing the maximum committee
    // size. It must agree with `circuits/common.MaxN` (Go side); the test
    // `TestSolidityMaxNMatchesGoMaxN` enforces the equality at CI time.
    // Changing this requires recompiling every circuit, regenerating the
    // proving keys, and redeploying the verifier wrappers.
    // ──────────────────────────────────────────────────────────────────────
    uint256 internal constant MAX_N            = 32;
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
    uint64 public epochNonce;

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
    // CIRCUITS_AUDIT2 #2: ciphertexts, committee partials, partial counts, and
    // combined plaintexts are all keyed by (epochId, aid, ctIdx). The aid
    // namespace prevents two applications under the same epoch from colliding
    // on ctIdx (one app blocking another, partials shared/rejected across apps,
    // a combine for one aid marking the slot completed for all aids). The
    // legacy per-epoch path (no application registered) uses `aid = bytes32(0)`.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => mapping(address participant => DKGTypes.PartialDecryptionRecord partialDecryption)))) internal epochPartialDecryptions;
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => uint16 count))) internal epochPartialDecryptionCounts;
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => DKGTypes.CombinedDecryptionRecord combined))) internal
        epochCombinedDecryptions;
    /// @dev Stores keccak256(abi.encode(scX, scY)) for each share commitment, packing
    /// the original (x,y) pair into a single 32-byte slot. Saves one cold SSTORE per
    /// committee member at finalize time. The pre-image (x,y) is exposed in the
    /// EpochFinalized event for off-chain consumers.
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

    /// @dev keccak256(abi.encode(c1x, c1y, c2x, c2y)) for each ciphertext submitted to a
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

    /// @dev Per-application registrations, keyed by `(eid, aid)`. See PLAN §4.3
    ///      and DKGTypes.Application for the record shape. Both
    ///      registerApplication and registerApplicationCoDec write here; the
    ///      mode flag distinguishes the two paths at decryption time.
    mapping(bytes12 epochId => mapping(bytes32 aid => DKGTypes.Application app)) internal applications;
    /// @dev List of registered aids per epoch (excluding bytes32(0)). Used by
    ///      _evictRound to enumerate the per-aid storage that needs zeroing
    ///      out for SSTORE refunds. Append-only; never reordered.
    mapping(bytes12 epochId => bytes32[] aids) internal epochAidsList;

    /// @dev Per-(eid, aid, ciphertextIndex) organizer share submissions for
    ///      mode-1 applications. Written by submitOrganizerShare and read by
    ///      combineDecryption when the application's mode is OrganizerCoDec.
    mapping(bytes12 epochId => mapping(bytes32 aid => mapping(uint16 ciphertextIndex => DKGTypes.OrganizerShareRecord)))
        internal organizerShares;

    constructor(
        uint32 _chainId,
        address _registry,
        address _contributionVerifier,
        address _partialDecryptVerifier,
        address _finalizeVerifier,
        address _decryptCombineVerifier
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
    }

    /// @notice Create a new DKG epoch.
    /// @dev    Snapshots `REGISTRY.nodeCount()` to derive the per-epoch
    ///         lottery threshold and pins the seed block at
    ///         `block.number + seedDelay`. The caller becomes the epoch
    ///         organizer but does not select committee members — every
    ///         registered node that passes the lottery can self-claim a slot.
    /// @param  threshold                  Shamir reconstruction threshold `t`.
    /// @param  committeeSize              Target committee size `n`. Must
    ///                                    be in [1, MAX_N]; values above
    ///                                    MAX_N revert with `InvalidPolicy`.
    /// @param  minValidContributions      Minimum accepted contributions
    ///                                    required to allow `finalizeEpoch`.
    ///                                    Must be ≥ `threshold` (otherwise
    ///                                    the epoch could finalize with
    ///                                    fewer share holders than needed
    ///                                    to decrypt — `InvalidPolicy`).
    /// @param  lotteryAlphaBps            Oversubscription factor α encoded as
    ///                                    basis points (10000 = α=1.0). The
    ///                                    expected eligible set size is
    ///                                    `α · committeeSize`.
    /// @param  seedDelay                  Number of blocks after `createEpoch`
    ///                                    that must elapse before the seed
    ///                                    block is valid. Must be ≥ 1.
    /// @param  registrationDeadlineBlock  Block height after which the
    ///                                    registration window is considered
    ///                                    stalled and `extendRegistration`
    ///                                    may reroll the seed.
    /// @param  contributionDeadlineBlock  Block height after which the epoch
    ///                                    may be aborted for inactivity.
    /// @param  finalizeNotBeforeBlock     Earliest block at which finalizeEpoch
    ///                                    can succeed. Must be strictly greater
    ///                                    than contributionDeadlineBlock so all
    ///                                    selected participants have a window
    ///                                    to submit contributions before the
    ///                                    finalize set is closed.
    /// @return                            The 12-byte epoch identifier
    ///                                    `uint32 prefix || uint64 nonce`.
    function createEpoch(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps,
        uint16 seedDelay,
        uint64 registrationDeadlineBlock,
        uint64 contributionDeadlineBlock,
        uint64 finalizeNotBeforeBlock,
        DKGTypes.DecryptionPolicy calldata decryptionPolicy
    ) external returns (bytes12) {
        if (
            threshold == 0 || committeeSize == 0 || threshold > committeeSize
                || committeeSize > MAX_N
                || minValidContributions == 0 || minValidContributions > committeeSize
                || minValidContributions < threshold
                || lotteryAlphaBps < 10000 || seedDelay == 0 || seedDelay > 256
                || registrationDeadlineBlock <= uint64(block.number) + uint64(seedDelay)
                || contributionDeadlineBlock <= registrationDeadlineBlock
                || finalizeNotBeforeBlock <= contributionDeadlineBlock
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

        epochs[epochId] = Epoch({
            organizer: msg.sender,
            policy: DKGTypes.EpochPolicy({
                threshold: threshold,
                committeeSize: committeeSize,
                minValidContributions: minValidContributions,
                lotteryAlphaBps: lotteryAlphaBps,
                seedDelay: seedDelay,
                registrationDeadlineBlock: registrationDeadlineBlock,
                contributionDeadlineBlock: contributionDeadlineBlock,
                finalizeNotBeforeBlock: finalizeNotBeforeBlock
            }),
            decryptionPolicy: decryptionPolicy,
            status: DKGTypes.EpochPhase.Registration,
            nonce: epochNonce,
            seedBlock: uint64(block.number) + uint64(seedDelay),
            seed: bytes32(0),
            lotteryThreshold: lotteryThreshold,
            claimedCount: 0,
            contributionCount: 0,
            partialDecryptionCount: 0,
            ciphertextCount: 0
        });

        emit EpochCreated(epochId, msg.sender, uint64(block.number) + uint64(seedDelay), lotteryThreshold);
        return epochId;
    }

    /// @notice Eligible registered nodes call this to claim a slot in the epoch's
    /// committee. The first `committeeSize` callers that pass the lottery and arrive
    /// before `registrationDeadlineBlock` form the committee.
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
        if (!PhaseLib.inRegistration(epoch.status, epoch.policy.registrationDeadlineBlock)) revert InvalidPhase();
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
            epoch.status = DKGTypes.EpochPhase.Contribution;
            emit RegistrationClosed(epochId);
        }
    }

    /// @notice Re-roll the lottery seed if the epoch failed to fill within the
    /// registration window. Anyone may call once the original deadline has passed; the
    /// new seed is derived from the current block.
    /// @notice Reroll the lottery seed for a stalled registration window.
    /// @dev    Callable by anyone after `registrationDeadlineBlock` if the
    ///         committee has not filled. Captures a fresh `blockhash` as the
    ///         new seed, resets the claimed count, and pushes the deadline
    ///         forward by one `seedDelay` window.
    /// @param  epochId The epoch identifier.
    function extendRegistration(bytes12 epochId) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Registration) revert InvalidPhase();
        if (epoch.claimedCount == epoch.policy.committeeSize) revert InvalidPhase();
        if (block.number <= uint256(epoch.policy.registrationDeadlineBlock)) revert InvalidPhase();

        // Capture the original window length BEFORE we mutate seedBlock.
        uint64 oldDeadline = epoch.policy.registrationDeadlineBlock;
        uint64 oldSeedBlock = epoch.seedBlock;
        uint64 window = oldDeadline - (oldSeedBlock - uint64(epoch.policy.seedDelay));

        uint64 newSeedBlock = uint64(block.number) + uint64(epoch.policy.seedDelay);
        uint64 newRegistrationDeadline = uint64(block.number) + window;

        // Guard: the extended registration window must close before the contribution
        // deadline; otherwise the epoch would become permanently stuck with no way to
        // advance to the Contribution phase.
        if (newRegistrationDeadline >= epoch.policy.contributionDeadlineBlock) revert InvalidPolicy();

        epoch.seed = bytes32(0);
        epoch.seedBlock = newSeedBlock;
        epoch.policy.registrationDeadlineBlock = newRegistrationDeadline;
        emit RegistrationExtended(epochId, newSeedBlock, newRegistrationDeadline);
    }

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
        // every registered application aid.
        bytes32[] storage regAids = epochAidsList[oldRoundId];
        uint256 aidCount = regAids.length + 1;
        for (uint256 i = 0; i < n; i++) {
            address participant = parts[i];
            delete selectedOperators[oldRoundId][participant];
            delete epochShareCommitmentHashes[oldRoundId][uint16(i + 1)];
            delete epochContributions[oldRoundId][participant];
            // Clear per-ciphertext partial decryption records across all aids.
            for (uint256 a = 0; a < aidCount; a++) {
                bytes32 aid = a == 0 ? bytes32(0) : regAids[a - 1];
                for (uint16 ci = 1; ci <= MAX_CIPHERTEXT_INDEX; ci++) {
                    if (epochPartialDecryptions[oldRoundId][aid][ci][participant].accepted) {
                        delete epochPartialDecryptions[oldRoundId][aid][ci][participant];
                    }
                }
            }
        }
        // Clear per-ciphertext combined decryption records, counts, and ciphertext hashes.
        for (uint256 a = 0; a < aidCount; a++) {
            bytes32 aid = a == 0 ? bytes32(0) : regAids[a - 1];
            for (uint16 ci = 1; ci <= MAX_CIPHERTEXT_INDEX; ci++) {
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
            if (a > 0) delete applications[oldRoundId][aid];
        }
        delete epochAidsList[oldRoundId];
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
        uint256 commitment0X,
        uint256 commitment0Y,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (!PhaseLib.inContribution(epoch.status, epoch.policy.contributionDeadlineBlock)) revert InvalidPhase();
        if (!selectedOperators[epochId][msg.sender]) revert NotSelectedParticipant();
        if (contributorIndex == 0 || contributorIndex > epoch.policy.committeeSize) revert InvalidContribution();
        if (epochParticipants[epochId][contributorIndex - 1] != msg.sender) revert InvalidProofInput();

        DKGTypes.ContributionRecord storage record = epochContributions[epochId][msg.sender];
        if (record.accepted) revert AlreadyContributed();

        IZKVerifier(CONTRIBUTION_VERIFIER).verifyProof(proof, input);
        uint256[10] memory publicInputs = abi.decode(input, (uint256[10]));
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
        // publicInputs[8] = CommitmentX0 (contributor's individual public key share x)
        // publicInputs[9] = CommitmentY0 (contributor's individual public key share y)
        if (publicInputs[8] != commitment0X || publicInputs[9] != commitment0Y) revert InvalidProofInput();

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

        // Accumulate the collective public key: add commitment[0] = a_{i,0}·G
        // to the running sum. The ZK proof guarantees commitment0X/Y is the
        // correct zeroth Feldman commitment point of this contributor's polynomial.
        // Identity is (0, 1); the initial mapping value (0, 0) is treated as (0, 1).
        DKGTypes.Point storage cpk = _collectiveKey[epochId];
        uint256 accX = cpk.x;
        uint256 accY = cpk.y == 0 ? 1 : cpk.y; // treat uninitialized (0,0) as identity (0,1)
        (uint256 newX, uint256 newY) = BabyJubJub.pointAdd(accX, accY, commitment0X, commitment0Y);
        cpk.x = newX;
        cpk.y = newY;

        // Refresh the contributor's liveness timestamp on the registry.
        // A successful proof-gated contribution is the strongest possible
        // signal that the operator is alive and well-configured.
        IDKGRegistry(REGISTRY).markActive(msg.sender);

        emit ContributionSubmitted(epochId, msg.sender, contributorIndex, commitmentsHash, encryptedSharesHash);
    }

    /// @notice Returns the accumulated collective public key for a epoch.
    ///         This is the running sum of all accepted contributors' commitment[0]
    ///         points (a_{i,0}·G). Once the epoch is finalized it equals the
    ///         full collective public key. The y-coordinate of an uninitialized
    ///         (no contributions yet) key is returned as 1 (the identity element).
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
    ///         slot per entry; the pre-image is emitted in `EpochFinalized`.
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
        if (epoch.status == DKGTypes.EpochPhase.Finalized) revert AlreadyFinalized();
        if (epoch.status != DKGTypes.EpochPhase.Contribution) revert InvalidPhase();
        // finalizeNotBeforeBlock gate — semantically a "phase not yet open"
        // condition, so we reuse InvalidPhase to keep the contract small.
        if (block.number < uint256(epoch.policy.finalizeNotBeforeBlock)) revert InvalidPhase();
        if (epoch.contributionCount < epoch.policy.minValidContributions) revert InsufficientContributions();
        if (
            aggregateCommitmentsHash == bytes32(0) || collectivePublicKeyHash == bytes32(0)
                || shareCommitmentHash == bytes32(0)
        ) revert InvalidFinalization();

        IZKVerifier(FINALIZE_VERIFIER).verifyProof(proof, input);
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

        // Transcript layout (2N² + 5N words):
        //   words [0..N)              participantIndexes
        //   words [N..N+2N²)          contributionCommitments  (N contributors × N points × 2 coords)
        //   words [N+2N²..N+2N²+2N)   aggregateCommitments     (N points × 2 coords)
        //   words [N+2N²+2N..2N²+5N)  shareCommitments         (N points × 2 coords)
        if (transcript.length != FINALIZE_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        uint256 dOff;
        assembly { dOff := transcript.offset }

        _verifyFinalizeTranscript(epochId, epoch, challenge, publicInputs[8], transcript);

        // CIRCUITS_AUDIT2 #1: defence-in-depth — the proof's aggregateCommitments[0]
        // must equal the on-chain accumulated `_collectiveKey`. Without this a
        // valid finalize proof over a duplicated/omitted contributor subset
        // (which the duplicate-row bitmap already rejects) would still be
        // distinguishable from any future bug that lets an inconsistent set
        // through. Reads aggregate[0].x/y straight from the transcript calldata.
        {
            uint256 agg0Base = dOff + FINALIZE_AGG0_WORDS_OFFSET * 32;
            uint256 agg0X;
            uint256 agg0Y;
            assembly ("memory-safe") {
                agg0X := calldataload(agg0Base)
                agg0Y := calldataload(add(agg0Base, 0x20))
            }
            DKGTypes.Point storage cpkRef = _collectiveKey[epochId];
            uint256 cpkX = cpkRef.x;
            uint256 cpkY = cpkRef.y == 0 ? 1 : cpkRef.y;
            if (agg0X != cpkX || agg0Y != cpkY) revert InvalidProofInput();
        }

        epoch.status = DKGTypes.EpochPhase.Finalized;
        // The three commitment hashes are not persisted to storage; they are emitted
        // in EpochFinalized below and reconstructed off-chain from the event log.

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
            epochShareCommitmentHashes[epochId][uint16(pIdx)] = keccak256(abi.encode(scX, scY));
        }

        emit EpochFinalized(epochId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash);
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
        // CIRCUITS_AUDIT #4: track which participant indexes have been
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
            address participant = epochParticipants[epochId][participantIndex - 1];
            DKGTypes.PartialDecryptionRecord storage partialRecord =
                epochPartialDecryptions[epochId][aid][ciphertextIndex][participant];
            if (!partialRecord.accepted || partialRecord.participantIndex != participantIndex) revert InvalidProofInput();
            if (partialRecord.ciphertextIndex != ciphertextIndex) revert InvalidProofInput();
            if (pdX != partialRecord.delta.x || pdY != partialRecord.delta.y) revert InvalidProofInput();
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
        // CIRCUITS_AUDIT2 #1: reject duplicate participant indexes in the
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
    /// @dev `aid` binds the proof transcript to a specific application
    ///      (P9). Pass `bytes32(0)` for the legacy per-epoch path that
    ///      pre-dates the application surface; the circuit witness builder
    ///      defaults Aid/CtIdx/Role to zero/zero/COMMITTEE in that mode.
    /// @dev `c1x/c1y/c2x/c2y` are the raw ciphertext coordinates as
    ///      submitted via submitCiphertext. The contract verifies
    ///      `keccak256(abi.encode(...))` matches the stored ciphertext
    ///      hash and then binds the proof's public-input C1 (pi[5..6])
    ///      to the authoritative on-chain ciphertext (CIRCUITS_AUDIT #2).
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
        if (epoch.status != DKGTypes.EpochPhase.Finalized) revert InvalidPhase();
        if (!selectedOperators[epochId][msg.sender]) revert NotSelectedParticipant();
        if (
            participantIndex == 0 || participantIndex > epoch.policy.committeeSize || ciphertextIndex == 0
                || ciphertextIndex > MAX_CIPHERTEXT_INDEX || deltaHash == bytes32(0)
        ) revert InvalidPartialDecryption();
        if (epochParticipants[epochId][participantIndex - 1] != msg.sender) revert InvalidProofInput();

        // CIRCUITS_AUDIT #2: bind to the authoritative on-chain ciphertext.
        // Without this the prover can supply Δ_i = sk_i · B for an
        // arbitrary B, and the stored partial decryption is only meaningful
        // relative to that B — combine then aggregates points that aren't
        // decryptions of the submitted ciphertext.
        // CIRCUITS_AUDIT2 #2: ciphertext + partial storage are keyed by aid
        // so two applications under the same epoch don't collide on ctIdx.
        bytes32 storedCt = _ciphertexts[epochId][aid][ciphertextIndex];
        if (storedCt == bytes32(0)) revert CiphertextNotSubmitted();
        if (keccak256(abi.encode(c1x, c1y, c2x, c2y)) != storedCt) revert InvalidProofInput();

        DKGTypes.PartialDecryptionRecord storage record = epochPartialDecryptions[epochId][aid][ciphertextIndex][msg.sender];
        if (record.accepted) revert AlreadyPartiallyDecrypted();

        IZKVerifier(PARTIAL_DECRYPT_VERIFIER).verifyProof(proof, input);
        // P5 layout: [eid, aid, ctIdx, role, i, C1.x, C1.y, D_i.x, D_i.y,
        // delta.x, delta.y, A1.x, A1.y, A2.x, A2.y, response].
        // Committee partial decryptions always use role = COMMITTEE = 1
        // (organizer shares go through submitOrganizerShare instead).
        uint256[16] memory publicInputs = abi.decode(input, (uint256[16]));
        bytes32 storedScHash = epochShareCommitmentHashes[epochId][participantIndex];
        if (
            publicInputs[0] != _epochScalar(epochId)
                || publicInputs[1] != uint256(aid)
                || publicInputs[2] != ciphertextIndex
                || publicInputs[3] != uint256(DKGProtocol.ROLE_COMMITTEE)
                || publicInputs[4] != participantIndex
                // pi[5..6] = base point (C_1) — bind to the just-verified
                // on-chain ciphertext (CIRCUITS_AUDIT #2).
                || publicInputs[5] != c1x
                || publicInputs[6] != c1y
                || storedScHash == bytes32(0)
                || keccak256(abi.encode(publicInputs[7], publicInputs[8])) != storedScHash
        ) revert InvalidProofInput();
        if (deltaHash != keccak256(abi.encodePacked(publicInputs[9], publicInputs[10]))) revert InvalidProofInput();

        // Persist only what combineDecryption actually reads:
        //   - participantIndex + accepted: identity gate
        //   - delta.x/.y: BRLC verification
        DKGTypes.PartialDecryptionRecord storage prec =
            epochPartialDecryptions[epochId][aid][ciphertextIndex][msg.sender];
        prec.participantIndex = participantIndex;
        prec.ciphertextIndex = ciphertextIndex; // packed in slot 0 anyway
        prec.accepted = true;
        prec.delta.x = publicInputs[9];
        prec.delta.y = publicInputs[10];
        epoch.partialDecryptionCount++;
        epochPartialDecryptionCounts[epochId][aid][ciphertextIndex]++;

        emit PartialDecryptionSubmitted(epochId, msg.sender, participantIndex, ciphertextIndex, deltaHash);
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
        if (epoch.status != DKGTypes.EpochPhase.Finalized) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert InvalidCiphertext();

        // Well-formedness: coords must be canonical (< Q), on-curve, non-identity,
        // and in the prime-order subgroup. The first three are cheap (4 mulmods per
        // point); the subgroup check is one full BJJ scalar mul (~60k gas per point)
        // but is required to honor the paper's group-validation policy at every
        // entry point (paper §4.1, addressing DEEPSEEK §2.2). Without subgroup
        // membership a griefing submitter could pre-claim every index with a small-
        // order point the combine circuit can never accept.
        _requireValidEncryptionPoint(c1x, c1y);
        _requireValidEncryptionPoint(c2x, c2y);

        // Per-epoch DecryptionPolicy gates the legacy aid=0 path; per-application
        // AppPolicy (CIRCUITS_AUDIT2 #2) gates aid != 0.
        if (aid == bytes32(0)) {
            DKGTypes.DecryptionPolicy memory p = epoch.decryptionPolicy;
            if (p.ownerOnly && msg.sender != epoch.organizer) revert NotOwner();
            if (p.notBeforeBlock     != 0 && uint64(block.number)    < p.notBeforeBlock)     revert DecryptionNotYetAllowed();
            if (p.notBeforeTimestamp != 0 && uint64(block.timestamp) < p.notBeforeTimestamp) revert DecryptionNotYetAllowed();
            if (p.notAfterBlock      != 0 && uint64(block.number)    > p.notAfterBlock)      revert DecryptionExpired();
            if (p.notAfterTimestamp  != 0 && uint64(block.timestamp) > p.notAfterTimestamp)  revert DecryptionExpired();
            if (p.maxDecryptions     != 0 && epoch.ciphertextCount   >= p.maxDecryptions)    revert DecryptionLimitReached();
        } else {
            DKGTypes.Application storage app = applications[epochId][aid];
            if (!app.exists) revert InvalidApplication();
            DKGTypes.AppPolicy memory ap = app.policy;
            if (ap.authorizedSubmitter != address(0) && msg.sender != ap.authorizedSubmitter) revert NotOwner();
            if (ap.notBeforeBlock != 0 && uint64(block.number) < ap.notBeforeBlock) revert DecryptionNotYetAllowed();
            if (ap.notAfterBlock  != 0 && uint64(block.number) > ap.notAfterBlock)  revert DecryptionExpired();
            // maxCiphertexts is enforced per-app via ciphertextIndex bound:
            // index in [1, maxCiphertexts] when maxCiphertexts > 0.
            if (ap.maxCiphertexts != 0 && ciphertextIndex > ap.maxCiphertexts) revert DecryptionLimitReached();
        }

        if (_ciphertexts[epochId][aid][ciphertextIndex] != bytes32(0)) revert CiphertextAlreadySubmitted();

        _ciphertexts[epochId][aid][ciphertextIndex] = keccak256(abi.encode(c1x, c1y, c2x, c2y));
        unchecked { epoch.ciphertextCount += 1; }

        emit CiphertextSubmitted(epochId, aid, ciphertextIndex, msg.sender, c1x, c1y, c2x, c2y);
    }

    /// @dev Validate that (x, y) is a canonical, on-curve, non-identity point
    ///      in the prime-order subgroup. Reverts with InvalidCiphertext()
    ///      (rather than the BabyJubJub library's specific errors) so callers
    ///      observe a single failure mode at submitCiphertext time.
    ///      DEEPSEEK §2.2 hardening: matches the registry's
    ///      `_requireValidEncryptionPoint` naming + policy.
    function _requireValidEncryptionPoint(uint256 x, uint256 y) internal view {
        if (!BabyJubJub.isCanonical(x, y)) revert InvalidCiphertext();
        if (BabyJubJub.isIdentity(x, y)) revert InvalidCiphertext();
        if (!BabyJubJub.isOnCurve(x, y)) revert InvalidCiphertext();
        if (!BabyJubJub.isInPrimeSubgroup(x, y)) revert InvalidCiphertext();
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
        if (epoch.status != DKGTypes.EpochPhase.Finalized) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX || combineHash == bytes32(0)) revert InvalidCombinedDecryption();
        bytes32 storedCtHash = _ciphertexts[epochId][aid][ciphertextIndex];
        if (storedCtHash == bytes32(0)) revert CiphertextNotSubmitted();
        if (epochPartialDecryptionCounts[epochId][aid][ciphertextIndex] < epoch.policy.threshold) revert InsufficientPartialDecryptions();

        DKGTypes.CombinedDecryptionRecord storage record = epochCombinedDecryptions[epochId][aid][ciphertextIndex];
        if (record.completed) revert AlreadyCombined();

        IZKVerifier(DECRYPT_COMBINE_VERIFIER).verifyProof(proof, input);
        uint256 shareCount = _validateAndPostCombine(
            epochId, aid, ciphertextIndex, combineHash, plaintext, epoch, input, transcript, storedCtHash
        );
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
            DKGTypes.Application storage app = applications[epochId][aid];
            if (!app.exists) revert InvalidApplication();
            mode = uint8(app.mode);
            if (mode == uint8(DKGProtocol.MODE_PUBLIC_DERIVATION)) {
                derivationS = app.derivationS;
            } else {
                DKGTypes.OrganizerShareRecord storage org = organizerShares[epochId][aid][ciphertextIndex];
                if (!org.accepted) revert InsufficientPartialDecryptions();
                deltaOrg = org.deltaOrg;
            }
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

    /// @notice Submit an organizer's Δ_org = sk_org · C_1 share for a
    ///         mode-1 application's ciphertext, together with its DLEQ proof.
    ///         The proof reuses `PartialDecryptVerifier` with role=ORGANIZER
    ///         (paper §6.3 line 1161).
    /// @param  epochId         Anchor epoch identifier.
    /// @param  aid             Application identifier (must exist in mode 1).
    /// @param  ciphertextIndex Per-application ciphertext index.
    /// @param  deltaOrgX       X coordinate of Δ_org.
    /// @param  deltaOrgY       Y coordinate of Δ_org.
    /// @param  dleqProof       Encoded DLEQ proof bytes (ABI-encoded uint[8] / uint[4]).
    /// @param  dleqInput       Encoded DLEQ public input (uint[16]).
    /// @dev `c1x/c1y/c2x/c2y` are the raw ciphertext coordinates as
    ///      submitted via submitCiphertext. Verified against the stored
    ///      ciphertext hash and then bound to the proof's pi[5..6]
    ///      (CIRCUITS_AUDIT #1).
    function submitOrganizerShare(
        bytes12 epochId,
        bytes32 aid,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y,
        uint256 deltaOrgX,
        uint256 deltaOrgY,
        bytes calldata dleqProof,
        bytes calldata dleqInput
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Finalized) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert InvalidCiphertext();
        DKGTypes.Application storage app = applications[epochId][aid];
        if (!app.exists) revert InvalidApplication();
        if (uint8(app.mode) != uint8(DKGProtocol.MODE_ORGANIZER_CODEC)) revert InvalidApplication();

        // CIRCUITS_AUDIT #1: bind C1 to the on-chain ciphertext before
        // accepting any (PK_org, Δ_org) commitment. Without this the
        // organizer can produce a valid DLEQ for an unrelated base, and
        // the stored Δ_org corrects the wrong discrete log at combine.
        bytes32 storedCt = _ciphertexts[epochId][aid][ciphertextIndex];
        if (storedCt == bytes32(0)) revert CiphertextNotSubmitted();
        if (keccak256(abi.encode(c1x, c1y, c2x, c2y)) != storedCt) revert InvalidProofInput();

        if (organizerShares[epochId][aid][ciphertextIndex].accepted) revert AlreadyPartiallyDecrypted();

        BabyJubJub.requireValidPoint(deltaOrgX, deltaOrgY);

        // Verify the Chaum-Pedersen DLEQ proof via PartialDecryptVerifier.
        // The verifier checks that (PK_org, Δ_org) have the same discrete log
        // wrt (G, C_1), with role=ORGANIZER bound into the Fiat-Shamir
        // transcript. The contract is responsible for binding the public
        // inputs (eid, aid, ctIdx, role, C_1, PK_org, Δ_org) to authoritative
        // on-chain state before invoking the verifier.
        IZKVerifier(PARTIAL_DECRYPT_VERIFIER).verifyProof(dleqProof, dleqInput);
        uint256[16] memory pi = abi.decode(dleqInput, (uint256[16]));
        // CIRCUITS_AUDIT #1: the public-input layout per
        // circuits/partialdecrypt/circuit.go is
        //   [0]=eid, [1]=aid, [2]=ctIdx, [3]=role, [4]=i,
        //   [5..6]=Base/C1, [7..8]=PublicKey/D_i (= PK_org for organizer),
        //   [9..10]=Delta (= Δ_org for organizer),
        //   [11..12]=A1, [13..14]=A2, [15]=Response.
        // The previous code checked PK_org against pi[9..10] and
        // Δ_org against pi[11..12]; both were off by two slots, so a
        // malicious caller could supply an arbitrary point as Δ_org.
        if (
            pi[0] != _epochScalar(epochId)
                || pi[1] != uint256(aid)
                || pi[2] != uint256(ciphertextIndex)
                || pi[3] != uint256(DKGProtocol.ROLE_ORGANIZER)
                // pi[4] = participant index — must be 0 for organizer (per paper §6.3)
                || pi[4] != 0
                // pi[5..6] = base point — bound to the stored ciphertext.
                || pi[5] != c1x
                || pi[6] != c1y
                // pi[7..8] = D_i — for organizer, this is PK_org.
                || pi[7] != app.organizerPK.x
                || pi[8] != app.organizerPK.y
                // pi[9..10] = Δ_org.
                || pi[9] != deltaOrgX
                || pi[10] != deltaOrgY
        ) revert InvalidProofInput();

        organizerShares[epochId][aid][ciphertextIndex] = DKGTypes.OrganizerShareRecord({
            deltaOrg: DKGTypes.Point({x: deltaOrgX, y: deltaOrgY}),
            dleqHash: keccak256(dleqInput),
            accepted: true
        });

        emit OrganizerShareSubmitted(epochId, aid, ciphertextIndex, deltaOrgX, deltaOrgY);
    }

    // ─── Application lifecycle (paper §4.3, §6, PLAN.md §4.3) ────────────────

    /// @notice Register an application against a finalized epoch in
    ///         public-derivation mode (paper §4.3). Computes the per-application
    ///         derivation tag `S = keccak256(eid || PK_ep || aid) mod q` and
    ///         stores the application record. The implicit per-application
    ///         encryption key is `PK_aid = PK_ep + S·G`, recomputable on-chain
    ///         or off-chain by any reader from the stored `S`.
    /// @param  epochId  Anchor epoch identifier (must be Finalized).
    /// @param  aid      Caller-supplied non-zero application identifier.
    /// @param  policy   Per-application access policy (gates submitCiphertext).
    function registerApplication(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Finalized) revert InvalidPhase();
        if (aid == bytes32(0)) revert InvalidApplication();
        DKGTypes.Application storage existing = applications[epochId][aid];
        if (existing.exists) revert ApplicationAlreadyExists();

        DKGTypes.Point storage pkep = _collectiveKey[epochId];
        if (pkep.y == 0) revert InvalidEpoch();

        // S = keccak256(eid || PK_ep.x || PK_ep.y || aid) mod q
        uint256 s = uint256(
            keccak256(abi.encodePacked(epochId, pkep.x, pkep.y, aid))
        ) % BabyJubJub.SUBGROUP_ORDER;

        existing.creator = msg.sender;
        existing.mode = DKGTypes.AppMode.PublicDerivation;
        existing.derivationS = s;
        existing.organizerPK = DKGTypes.Point({x: 0, y: 1}); // identity
        existing.policy = policy;
        existing.createdAtBlock = uint64(block.number);
        existing.exists = true;
        epochAidsList[epochId].push(aid);

        emit ApplicationRegistered(
            epochId,
            aid,
            msg.sender,
            uint8(DKGProtocol.MODE_PUBLIC_DERIVATION),
            s,
            0,
            1
        );
    }

    /// @notice Register an application against a finalized epoch in
    ///         organizer co-decryption mode (paper §6). Verifies a Schnorr
    ///         proof of knowledge of `sk_org` per paper §6.2 / paper line 1138:
    ///
    ///             c = Poseidon(domain || eid || aid || PK_org || A)
    ///             z·G == A + c·PK_org
    ///
    ///         The implicit per-application key is `PK_aid = PK_ep + PK_org`.
    /// @param  epochId    Anchor epoch identifier (must be Finalized).
    /// @param  aid        Caller-supplied non-zero application identifier.
    /// @param  policy     Per-application access policy.
    /// @param  pkOrgX     Organizer public key X coordinate.
    /// @param  pkOrgY     Organizer public key Y coordinate.
    /// @param  schnorrAx  Schnorr nonce point A = w·G — X.
    /// @param  schnorrAy  Schnorr nonce point Y.
    /// @param  schnorrZ   Schnorr response z = w + c·sk_org (mod L).
    function registerApplicationCoDec(
        bytes12 epochId,
        bytes32 aid,
        DKGTypes.AppPolicy calldata policy,
        uint256 pkOrgX,
        uint256 pkOrgY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external {
        Epoch storage epoch = epochs[epochId];
        if (epoch.organizer == address(0)) revert InvalidEpoch();
        if (epoch.status != DKGTypes.EpochPhase.Finalized) revert InvalidPhase();
        if (aid == bytes32(0)) revert InvalidApplication();
        DKGTypes.Application storage existing = applications[epochId][aid];
        if (existing.exists) revert ApplicationAlreadyExists();

        BabyJubJub.requireValidPoint(pkOrgX, pkOrgY);
        BabyJubJub.requireValidPoint(schnorrAx, schnorrAy);

        if (
            !_verifyOrganizerSchnorr(
                epochId, aid, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ
            )
        ) revert InvalidSchnorrProof();

        existing.creator = msg.sender;
        existing.mode = DKGTypes.AppMode.OrganizerCoDec;
        existing.derivationS = 0;
        existing.organizerPK = DKGTypes.Point({x: pkOrgX, y: pkOrgY});
        existing.policy = policy;
        existing.createdAtBlock = uint64(block.number);
        existing.exists = true;
        epochAidsList[epochId].push(aid);

        emit ApplicationRegistered(
            epochId,
            aid,
            msg.sender,
            uint8(DKGProtocol.MODE_ORGANIZER_CODEC),
            0,
            pkOrgX,
            pkOrgY
        );
    }

    /// @notice Read an application record.
    function getApplication(bytes12 epochId, bytes32 aid)
        external
        view
        returns (DKGTypes.Application memory)
    {
        return applications[epochId][aid];
    }

    /// @dev Two-pass Poseidon (T5+T5) Fiat-Shamir transcript for the organizer
    ///      Schnorr proof. Layout (paper §6.2 line 1138):
    ///
    ///         inner = T5(domain, eid, PK_org.x, PK_org.y)
    ///         c     = T5(inner, aid_field, A.x, A.y)
    ///
    ///      `aid_field = uint256(aid) % BN254.Q` to keep the value in the
    ///      Poseidon input field. The shared `DOMAIN_ORGANIZER_REGISTER_V1`
    ///      digest namespaces the proof so it cannot be replayed against the
    ///      operator-registry Schnorr verification.
    function _organizerSchnorrChallenge(
        bytes12 epochId,
        bytes32 aid,
        uint256 pkX,
        uint256 pkY,
        uint256 ax,
        uint256 ay
    ) internal pure returns (uint256) {
        uint256 domainField = uint256(DKGProtocol.DOMAIN_ORGANIZER_REGISTER_V1) % BabyJubJub.Q;
        uint256[4] memory in1;
        in1[0] = domainField;
        in1[1] = uint256(uint96(epochId));
        in1[2] = pkX;
        in1[3] = pkY;
        uint256 inner = PoseidonT5.hash(in1);
        uint256[4] memory in2;
        in2[0] = inner;
        in2[1] = uint256(aid) % BabyJubJub.Q;
        in2[2] = ax;
        in2[3] = ay;
        return PoseidonT5.hash(in2);
    }

    /// @dev Verify the organizer Schnorr PoK: `z·G == A + c·PK_org`.
    function _verifyOrganizerSchnorr(
        bytes12 epochId,
        bytes32 aid,
        uint256 pkX,
        uint256 pkY,
        uint256 ax,
        uint256 ay,
        uint256 z
    ) internal view returns (bool) {
        uint256 c = _organizerSchnorrChallenge(epochId, aid, pkX, pkY, ax, ay);
        (uint256 zGx, uint256 zGy) = BabyJubJub.scalarMulBase(z);
        (uint256 cPKx, uint256 cPKy) = BabyJubJub.scalarMul(c, pkX, pkY);
        (uint256 rhsX, uint256 rhsY) = BabyJubJub.pointAdd(ax, ay, cPKx, cPKy);
        return zGx == rhsX && zGy == rhsY;
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
            epoch.status == DKGTypes.EpochPhase.Finalized
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

    function getPartialDecryption(bytes12 epochId, bytes32 aid, address participant, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.PartialDecryptionRecord memory)
    {
        return epochPartialDecryptions[epochId][aid][ciphertextIndex][participant];
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
    /// the EpochFinalized event log.
    function getShareCommitmentHash(bytes12 epochId, uint16 participantIndex)
        external
        view
        returns (bytes32)
    {
        return epochShareCommitmentHashes[epochId][participantIndex];
    }

    /// @notice keccak256(abi.encode(c1x, c1y, c2x, c2y)) of the ciphertext stored
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
}
