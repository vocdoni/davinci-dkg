// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGManager} from "./interfaces/IDKGManager.sol";
import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";
import {IZKVerifier} from "./interfaces/IZKVerifier.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGIdLib} from "./libraries/DKGIdLib.sol";
import {BRLC} from "./libraries/BRLC.sol";
import {Lagrange} from "./libraries/Lagrange.sol";
import {DKGTypes} from "./libraries/DKGTypes.sol";
import {PhaseLib} from "./libraries/PhaseLib.sol";

/// @title  DKGManager
/// @notice On-chain orchestrator for every phase of a davinci-dkg round.
/// @dev    Lifecycle: Registration (trustless lottery) → Contribution →
///         Finalized → Completed (or Aborted). Every state-mutating entry
///         point that makes a cryptographic claim is gated by a Groth16
///         verifier — no dispute phase, no complaint flow. Historic round
///         storage is bounded by a ring buffer of ROUND_HISTORY_SIZE (64)
///         entries; evicted rounds remain reconstructible from event logs.
///         The share-commitment list is stored as `keccak256(x, y)` per
///         participant (1 SSTORE instead of 2) and the transcripts used by
///         finalize/combine/reconstruct are read straight out of calldata
///         via assembly to avoid per-element bounds checks.
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
    //   finalizeRound:      participantIndexes (N) ‖
    //                       contributionCommitments (2N²) ‖
    //                       aggregateCommitments (2N) ‖
    //                       shareCommitments (2N)                         = 2N² + 5N
    //   combineDecryption:  ciphertext (4) ‖ participantIndexes (N) ‖
    //                       partialDecryptions (2N)                       = 4 + 3N
    //   reconstructSecret:  participantIndexes (N) ‖ revealedShares (N)   = 2N
    uint256 internal constant CONTRIB_TRANSCRIPT_WORDS     = 8 * MAX_N;
    uint256 internal constant FINALIZE_TRANSCRIPT_WORDS    = 2 * MAX_N * MAX_N + 5 * MAX_N;
    uint256 internal constant COMBINE_TRANSCRIPT_WORDS     = 4 + 3 * MAX_N;
    uint256 internal constant RECONSTRUCT_TRANSCRIPT_WORDS = 2 * MAX_N;
    // contribution-time per-section byte offsets:
    uint256 internal constant CONTRIB_PUBKEYS_BYTES_OFFSET = (2 * MAX_N + MAX_N) * 32;          // start of recipientPubKeys
    uint256 internal constant CONTRIB_PUBKEYS_BYTES_END    = (2 * MAX_N + MAX_N + 2 * MAX_N) * 32; // end of recipientPubKeys
    uint256 internal constant CONTRIB_DIGEST_BYTES_LEN     = 2 * MAX_N * 32;                    // first 2N words = commitmentPoints
    // finalize-time per-section byte offsets:
    uint256 internal constant FINALIZE_CONTRIB_BYTES_OFFSET = MAX_N * 32;                       // participantIndexes end
    uint256 internal constant FINALIZE_CONTRIB_BYTES_LEN    = 2 * MAX_N * MAX_N * 32;           // contributionCommitments length in bytes
    uint256 internal constant FINALIZE_PER_CONTRIB_BYTES    = 2 * MAX_N * 32;                   // bytes per contributor's commitments slice
    uint256 internal constant FINALIZE_SHARE_WORDS_OFFSET   = MAX_N + 2 * MAX_N * MAX_N + 2 * MAX_N; // shareCommitments start, in words
    // combine-time per-section byte offsets:
    uint256 internal constant COMBINE_PARTIALS_BYTES_OFFSET = (4 + MAX_N) * 32;                 // partialDecryptions start, in bytes
    uint256 internal constant RECONSTRUCT_VALUES_BYTES_OFFSET = MAX_N * 32;                     // revealedShares start, in bytes

    /// @dev Number of recent round IDs retained on-chain. After this many `createRound`
    /// calls, the oldest live round's storage is evicted (its data is wiped) and only
    /// the event log retains it. Tunable; 64 is large enough to cover several days of
    /// rounds at typical cadences.
    uint256 internal constant ROUND_HISTORY_SIZE = 64;

    /// @dev Upper bound on ciphertext indices accepted by `submitPartialDecryption`
    /// and `combineDecryption`. Prevents unbounded storage spam by a committee member
    /// who submits decryptions for arbitrarily large ciphertext indices.
    uint16 internal constant MAX_CIPHERTEXT_INDEX = 256;

    uint32 public immutable CHAIN_ID;
    address public immutable REGISTRY;
    uint32 public immutable ROUND_PREFIX;
    address public immutable CONTRIBUTION_VERIFIER;
    address public immutable PARTIAL_DECRYPT_VERIFIER;
    address public immutable FINALIZE_VERIFIER;
    address public immutable DECRYPT_COMBINE_VERIFIER;
    address public immutable REVEAL_SUBMIT_VERIFIER;
    address public immutable REVEAL_SHARE_VERIFIER;
    uint64 public roundNonce;

    /// @dev Fixed-size ring buffer of recent round IDs. New rounds push here at
    /// createRound; once the buffer is full, the displaced entry tells us which round
    /// to evict. `recentRoundsCount` counts total pushes; current write index is
    /// `recentRoundsCount % ROUND_HISTORY_SIZE`.
    bytes12[ROUND_HISTORY_SIZE] internal recentRounds;
    uint64 internal recentRoundsCount;

    mapping(bytes12 roundId => Round round) internal rounds;
    mapping(bytes12 roundId => mapping(address operator => bool selected)) internal selectedOperators;
    mapping(bytes12 roundId => address[] participants) internal roundParticipants;
    mapping(bytes12 roundId => mapping(address contributor => DKGTypes.ContributionRecord contribution)) internal
        roundContributions;
    mapping(bytes12 roundId => mapping(uint16 ciphertextIndex => mapping(address participant => DKGTypes.PartialDecryptionRecord partialDecryption))) internal roundPartialDecryptions;
    mapping(bytes12 roundId => mapping(uint16 ciphertextIndex => uint16 count)) internal roundPartialDecryptionCounts;
    mapping(bytes12 roundId => mapping(uint16 ciphertextIndex => DKGTypes.CombinedDecryptionRecord combined)) internal
        roundCombinedDecryptions;
    mapping(bytes12 roundId => mapping(address participant => DKGTypes.RevealedShareRecord share)) internal
        roundRevealedShares;
    /// @dev Stores keccak256(abi.encode(scX, scY)) for each share commitment, packing
    /// the original (x,y) pair into a single 32-byte slot. Saves one cold SSTORE per
    /// committee member at finalize time. The pre-image (x,y) is exposed in the
    /// RoundFinalized event for off-chain consumers.
    mapping(bytes12 roundId => mapping(uint16 participantIndex => bytes32 shareCommitmentHash)) internal roundShareCommitmentHashes;

    /// @dev Stores keccak256 over the canonical (recipientIndexes ‖ recipientPubKeys)
    /// section of any valid submitContribution transcript for this round. Set once at
    /// selectParticipants. Lets submitContribution verify the entire 96-word committee
    /// section in one keccak instead of 32 storage reads + 32 external registry calls.
    mapping(bytes12 roundId => bytes32 prefixHash) internal roundContribPrefixHash;

    /// @dev Accumulates the collective public key on-chain as contributions are submitted.
    ///      Each accepted contribution adds its commitment[0] point (a_{i,0}·G) to the
    ///      running sum. The identity element (0,1) is the initial value. Once the round
    ///      is finalized the value equals sum_i(a_{i,0}·G) = the collective public key.
    mapping(bytes12 roundId => DKGTypes.Point) internal _collectiveKey;

    /// @dev keccak256(abi.encode(c1x, c1y, c2x, c2y)) for each ciphertext submitted to a
    ///      round. Written once per (roundId, ciphertextIndex) by submitCiphertext and
    ///      verified by combineDecryption to bind the combine proof to the authoritative
    ///      on-chain ciphertext (preventing a combiner from swapping in a different ct).
    ///      The raw coordinates are available via the CiphertextSubmitted event log.
    mapping(bytes12 roundId => mapping(uint16 ciphertextIndex => bytes32 ciphertextHash)) internal _ciphertexts;

    // BabyJubJub curve parameters in REDUCED twisted-Edwards form over BN254.Fr.
    //   A·x² + y² = 1 + D·x²·y²  (mod Q)   with A = -1
    // This matches the form used by gnark's bn254/twistededwards curve and by the
    // davinci-dkg ZK circuits (and thus by the ciphertexts nodes emit via
    // `group.Encode`). Do NOT confuse with the iden3/circomlib-style form used
    // internally by `libraries/BabyJubJub.sol` (A=168700, D=168696) — those are
    // a different isomorphic affine chart.
    uint256 private constant BabyJubJub_Q =
        21888242871839275222246405745257275088548364400416034343698204186575808495617;
    // A = -1 mod Q
    uint256 private constant BabyJubJub_A_NEG =
        21888242871839275222246405745257275088548364400416034343698204186575808495616;
    uint256 private constant BabyJubJub_D =
        12181644023421730124874158521699555681764249180949974110617291017600649128846;

    bytes32 internal constant CONTRIBUTION_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:contribution:v1");
    bytes32 internal constant DECRYPT_COMBINE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:decrypt-combine:v1");
    bytes32 internal constant FINALIZE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:finalize:v1");
    bytes32 internal constant REVEAL_SHARE_TRANSCRIPT_DOMAIN = keccak256("davinci-dkg:reveal-share:v1");

    constructor(
        uint32 _chainId,
        address _registry,
        address _contributionVerifier,
        address _partialDecryptVerifier,
        address _finalizeVerifier,
        address _decryptCombineVerifier,
        address _revealSubmitVerifier,
        address _revealShareVerifier
    ) {
        if (uint32(block.chainid) != _chainId) revert InvalidChainId();
        if (_registry == address(0)) revert InvalidAddress();
        if (
            _contributionVerifier == address(0) || _partialDecryptVerifier == address(0) || _finalizeVerifier == address(0)
                || _decryptCombineVerifier == address(0) || _revealSubmitVerifier == address(0)
                || _revealShareVerifier == address(0)
        ) revert InvalidVerifier();
        CHAIN_ID = _chainId;
        REGISTRY = _registry;
        ROUND_PREFIX = DKGIdLib.getPrefix(_chainId, address(this));
        CONTRIBUTION_VERIFIER = _contributionVerifier;
        PARTIAL_DECRYPT_VERIFIER = _partialDecryptVerifier;
        FINALIZE_VERIFIER = _finalizeVerifier;
        DECRYPT_COMBINE_VERIFIER = _decryptCombineVerifier;
        REVEAL_SUBMIT_VERIFIER = _revealSubmitVerifier;
        REVEAL_SHARE_VERIFIER = _revealShareVerifier;
    }

    /// @notice Create a new DKG round.
    /// @dev    Snapshots `REGISTRY.nodeCount()` to derive the per-round
    ///         lottery threshold and pins the seed block at
    ///         `block.number + seedDelay`. The caller becomes the round
    ///         organizer but does not select committee members — every
    ///         registered node that passes the lottery can self-claim a slot.
    /// @param  threshold                  Shamir reconstruction threshold `t`.
    /// @param  committeeSize              Target committee size `n`.
    /// @param  minValidContributions      Minimum accepted contributions
    ///                                    required to allow `finalizeRound`.
    /// @param  lotteryAlphaBps            Oversubscription factor α encoded as
    ///                                    basis points (10000 = α=1.0). The
    ///                                    expected eligible set size is
    ///                                    `α · committeeSize`.
    /// @param  seedDelay                  Number of blocks after `createRound`
    ///                                    that must elapse before the seed
    ///                                    block is valid. Must be ≥ 1.
    /// @param  registrationDeadlineBlock  Block height after which the
    ///                                    registration window is considered
    ///                                    stalled and `extendRegistration`
    ///                                    may reroll the seed.
    /// @param  contributionDeadlineBlock  Block height after which the round
    ///                                    may be aborted for inactivity.
    /// @param  disclosureAllowed          When true, enables the reveal-share
    ///                                    reconstruction phase on this round.
    /// @return                            The 12-byte round identifier
    ///                                    `uint32 prefix || uint64 nonce`.
    function createRound(
        uint16 threshold,
        uint16 committeeSize,
        uint16 minValidContributions,
        uint16 lotteryAlphaBps,
        uint16 seedDelay,
        uint64 registrationDeadlineBlock,
        uint64 contributionDeadlineBlock,
        bool disclosureAllowed,
        DKGTypes.DecryptionPolicy calldata decryptionPolicy
    ) external returns (bytes12) {
        if (
            threshold == 0 || committeeSize == 0 || threshold > committeeSize
                || minValidContributions == 0 || minValidContributions > committeeSize
                || lotteryAlphaBps < 10000 || seedDelay == 0 || seedDelay > 256
                || registrationDeadlineBlock <= uint64(block.number) + uint64(seedDelay)
                || contributionDeadlineBlock <= registrationDeadlineBlock
        ) revert InvalidPolicy();

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

        roundNonce++;
        bytes12 roundId = DKGIdLib.computeRoundId(ROUND_PREFIX, roundNonce);

        // Evict the oldest live round if the history buffer is full.
        uint256 writeSlot = uint256(recentRoundsCount) % ROUND_HISTORY_SIZE;
        if (recentRoundsCount >= ROUND_HISTORY_SIZE) {
            bytes12 evictedKey = recentRounds[writeSlot];
            if (evictedKey != bytes12(0)) {
                _evictRound(evictedKey);
            }
        }
        recentRounds[writeSlot] = roundId;
        unchecked { recentRoundsCount += 1; }

        rounds[roundId] = Round({
            organizer: msg.sender,
            policy: DKGTypes.RoundPolicy({
                threshold: threshold,
                committeeSize: committeeSize,
                minValidContributions: minValidContributions,
                lotteryAlphaBps: lotteryAlphaBps,
                seedDelay: seedDelay,
                registrationDeadlineBlock: registrationDeadlineBlock,
                contributionDeadlineBlock: contributionDeadlineBlock,
                disclosureAllowed: disclosureAllowed
            }),
            decryptionPolicy: decryptionPolicy,
            status: DKGTypes.RoundStatus.Registration,
            nonce: roundNonce,
            seedBlock: uint64(block.number) + uint64(seedDelay),
            seed: bytes32(0),
            lotteryThreshold: lotteryThreshold,
            claimedCount: 0,
            contributionCount: 0,
            partialDecryptionCount: 0,
            revealedShareCount: 0,
            ciphertextCount: 0
        });

        emit RoundCreated(roundId, msg.sender, uint64(block.number) + uint64(seedDelay), lotteryThreshold);
        return roundId;
    }

    /// @notice Eligible registered nodes call this to claim a slot in the round's
    /// committee. The first `committeeSize` callers that pass the lottery and arrive
    /// before `registrationDeadlineBlock` form the committee.
    /// @notice Claim a committee slot in the trustless lottery.
    /// @dev    Admissible iff `keccak256(seed ‖ msg.sender) < lotteryThreshold`.
    ///         The first call after `block.number ≥ seedBlock` lazily resolves
    ///         `seed = blockhash(seedBlock)`; the contract emits
    ///         `SeedResolved` on that call. Further calls are served
    ///         first-come-first-served until `committeeSize` slots are filled,
    ///         at which point the committee snapshot is taken and the round
    ///         advances to Contribution.
    /// @param  roundId The round identifier returned by `createRound`.
    function claimSlot(bytes12 roundId) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (!PhaseLib.inRegistration(round.status, round.policy.registrationDeadlineBlock)) revert InvalidPhase();
        if (round.claimedCount >= round.policy.committeeSize) revert SlotsFull();
        if (selectedOperators[roundId][msg.sender]) revert AlreadyClaimed();

        // Lazy seed resolution: capture blockhash(seedBlock) on the first claimer
        // that arrives at or after seedBlock.
        bytes32 seed = round.seed;
        if (seed == bytes32(0)) {
            uint256 sb = uint256(round.seedBlock);
            if (block.number <= sb) revert SeedNotReady();
            // blockhash(b) returns 0 for any b ≥ block.number or b + 256 < block.number
            bytes32 fresh = blockhash(sb);
            if (fresh == bytes32(0)) revert SeedExpired();
            round.seed = fresh;
            seed = fresh;
            emit SeedResolved(roundId, fresh);
        }

        // Caller must be an active registered node.
        IDKGRegistry.NodeKey memory node = IDKGRegistry(REGISTRY).getNode(msg.sender);
        if (node.status != IDKGRegistry.NodeStatus.ACTIVE) revert NotRegistered();

        // Lottery check: hash(seed ‖ caller) must fall under the round threshold.
        if (uint256(keccak256(abi.encodePacked(seed, msg.sender))) >= round.lotteryThreshold) {
            revert NotEligible();
        }

        // First-come-first-served slot allocation.
        uint16 slot = round.claimedCount;
        roundParticipants[roundId].push(msg.sender);
        selectedOperators[roundId][msg.sender] = true;
        round.claimedCount = slot + 1;
        emit SlotClaimed(roundId, msg.sender, slot);

        // When the last slot is filled, snapshot the committee key hash and transition
        // to Contribution. The snapshot is needed so submitContribution can verify the
        // entire (recipientIndexes ‖ recipientPubKeys) calldata block in one keccak.
        if (slot + 1 == round.policy.committeeSize) {
            _snapshotCommittee(roundId, round.policy.committeeSize);
            round.status = DKGTypes.RoundStatus.Contribution;
            emit RegistrationClosed(roundId);
        }
    }

    /// @notice Re-roll the lottery seed if the round failed to fill within the
    /// registration window. Anyone may call once the original deadline has passed; the
    /// new seed is derived from the current block.
    /// @notice Reroll the lottery seed for a stalled registration window.
    /// @dev    Callable by anyone after `registrationDeadlineBlock` if the
    ///         committee has not filled. Captures a fresh `blockhash` as the
    ///         new seed, resets the claimed count, and pushes the deadline
    ///         forward by one `seedDelay` window.
    /// @param  roundId The round identifier.
    function extendRegistration(bytes12 roundId) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (round.status != DKGTypes.RoundStatus.Registration) revert InvalidPhase();
        if (round.claimedCount == round.policy.committeeSize) revert InvalidPhase();
        if (block.number <= uint256(round.policy.registrationDeadlineBlock)) revert InvalidPhase();

        // Capture the original window length BEFORE we mutate seedBlock.
        uint64 oldDeadline = round.policy.registrationDeadlineBlock;
        uint64 oldSeedBlock = round.seedBlock;
        uint64 window = oldDeadline - (oldSeedBlock - uint64(round.policy.seedDelay));

        uint64 newSeedBlock = uint64(block.number) + uint64(round.policy.seedDelay);
        uint64 newRegistrationDeadline = uint64(block.number) + window;

        // Guard: the extended registration window must close before the contribution
        // deadline; otherwise the round would become permanently stuck with no way to
        // advance to the Contribution phase.
        if (newRegistrationDeadline >= round.policy.contributionDeadlineBlock) revert InvalidPolicy();

        round.seed = bytes32(0);
        round.seedBlock = newSeedBlock;
        round.policy.registrationDeadlineBlock = newRegistrationDeadline;
        emit RegistrationExtended(roundId, newSeedBlock, newRegistrationDeadline);
    }

    /// @dev Internal helper: build the canonical (recipientIndexes ‖ pubKeys) layout
    /// for the committee that's just been filled and store its keccak256. Drives the
    /// O(1) committee verification in submitContribution.
    function _snapshotCommittee(bytes12 roundId, uint16 committeeSize) internal {
        uint256[MAX_N] memory canonicalIdxs;
        uint256[2 * MAX_N] memory canonicalKeys;
        address[] storage participants = roundParticipants[roundId];
        for (uint256 i = 0; i < committeeSize; i++) {
            canonicalIdxs[i] = i + 1;
            IDKGRegistry.NodeKey memory node = IDKGRegistry(REGISTRY).getNode(participants[i]);
            canonicalKeys[i * 2] = node.pubX;
            canonicalKeys[i * 2 + 1] = node.pubY;
        }
        for (uint256 i = committeeSize; i < MAX_N; i++) {
            canonicalKeys[i * 2 + 1] = 1; // identity-pad unused slots
        }
        roundContribPrefixHash[roundId] = keccak256(abi.encodePacked(canonicalIdxs, canonicalKeys));
    }

    /// @dev Wipes all storage tied to an old round when it falls out of the recent
    /// rounds ring buffer. Refunds gas via SSTORE-zero on the storage slots being
    /// cleared. Off-chain consumers must rely on the historical event log.
    ///
    /// Cleans up all five previously-leaking mappings in addition to the core
    /// round data: contributions, partial decryptions (per-ciphertext counts and
    /// per-participant records), combined decryptions, and revealed shares.
    function _evictRound(bytes12 oldRoundId) internal {
        Round storage r = rounds[oldRoundId];
        if (r.organizer == address(0)) return;
        address[] storage parts = roundParticipants[oldRoundId];
        uint256 n = parts.length;
        for (uint256 i = 0; i < n; i++) {
            address participant = parts[i];
            delete selectedOperators[oldRoundId][participant];
            delete roundShareCommitmentHashes[oldRoundId][uint16(i + 1)];
            delete roundContributions[oldRoundId][participant];
            delete roundRevealedShares[oldRoundId][participant];
            // Clear per-ciphertext partial decryption records and counts.
            for (uint16 ci = 1; ci <= MAX_CIPHERTEXT_INDEX; ci++) {
                if (roundPartialDecryptions[oldRoundId][ci][participant].accepted) {
                    delete roundPartialDecryptions[oldRoundId][ci][participant];
                }
            }
        }
        // Clear per-ciphertext combined decryption records, counts, and ciphertext hashes.
        for (uint16 ci = 1; ci <= MAX_CIPHERTEXT_INDEX; ci++) {
            if (roundPartialDecryptionCounts[oldRoundId][ci] > 0) {
                delete roundPartialDecryptionCounts[oldRoundId][ci];
            }
            if (roundCombinedDecryptions[oldRoundId][ci].completed) {
                delete roundCombinedDecryptions[oldRoundId][ci];
            }
            if (_ciphertexts[oldRoundId][ci] != bytes32(0)) {
                delete _ciphertexts[oldRoundId][ci];
            }
        }
        delete roundParticipants[oldRoundId];
        delete roundContribPrefixHash[oldRoundId];
        delete _collectiveKey[oldRoundId];
        delete rounds[oldRoundId];
        emit RoundEvicted(oldRoundId);
    }

    /// @notice Submit a contributor's polynomial commitments, encrypted
    ///         shares and Groth16 proof of correctness.
    /// @dev    The committee membership + BabyJubJub public keys are verified
    ///         against a single `keccak256` snapshot taken when the lottery
    ///         filled; the transcript is read straight from calldata via the
    ///         BRLC helper. The transaction reverts if the proof fails.
    function submitContribution(
        bytes12 roundId,
        uint16 contributorIndex,
        bytes32 commitmentsHash,
        bytes32 encryptedSharesHash,
        uint256 commitment0X,
        uint256 commitment0Y,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (!PhaseLib.inContribution(round.status, round.policy.contributionDeadlineBlock)) revert InvalidPhase();
        if (!selectedOperators[roundId][msg.sender]) revert NotSelectedParticipant();
        if (contributorIndex == 0 || contributorIndex > round.policy.committeeSize) revert InvalidContribution();
        if (roundParticipants[roundId][contributorIndex - 1] != msg.sender) revert InvalidProofInput();

        DKGTypes.ContributionRecord storage record = roundContributions[roundId][msg.sender];
        if (record.accepted) revert AlreadyContributed();

        IZKVerifier(CONTRIBUTION_VERIFIER).verifyProof(proof, input);
        uint256[10] memory publicInputs = abi.decode(input, (uint256[10]));
        if (
            publicInputs[0] != _roundScalar(roundId) || publicInputs[1] != round.policy.threshold
                || publicInputs[2] != round.policy.committeeSize || publicInputs[3] != contributorIndex
                || bytes32(publicInputs[4]) != commitmentsHash || bytes32(publicInputs[5]) != encryptedSharesHash
        ) revert InvalidProofInput();
        uint256 challenge = BRLC.deriveChallenge(
            roundId,
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
        if (keccak256(transcript[CONTRIB_DIGEST_BYTES_LEN:CONTRIB_PUBKEYS_BYTES_END]) != roundContribPrefixHash[roundId]) {
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
        DKGTypes.ContributionRecord storage rec = roundContributions[roundId][msg.sender];
        rec.contributorIndex = contributorIndex;
        rec.commitmentVectorDigest = commitmentDigest;
        rec.accepted = true;
        round.contributionCount++;

        // Accumulate the collective public key: add commitment[0] = a_{i,0}·G
        // to the running sum. The ZK proof guarantees commitment0X/Y is the
        // correct zeroth Feldman commitment point of this contributor's polynomial.
        // Identity is (0, 1); the initial mapping value (0, 0) is treated as (0, 1).
        DKGTypes.Point storage cpk = _collectiveKey[roundId];
        uint256 accX = cpk.x;
        uint256 accY = cpk.y == 0 ? 1 : cpk.y; // treat uninitialized (0,0) as identity (0,1)
        (uint256 newX, uint256 newY) = BabyJubJub.pointAdd(accX, accY, commitment0X, commitment0Y);
        cpk.x = newX;
        cpk.y = newY;

        // Refresh the contributor's liveness timestamp on the registry.
        // A successful proof-gated contribution is the strongest possible
        // signal that the operator is alive and well-configured.
        IDKGRegistry(REGISTRY).markActive(msg.sender);

        emit ContributionSubmitted(roundId, msg.sender, contributorIndex, commitmentsHash, encryptedSharesHash);
    }

    /// @notice Returns the accumulated collective public key for a round.
    ///         This is the running sum of all accepted contributors' commitment[0]
    ///         points (a_{i,0}·G). Once the round is finalized it equals the
    ///         full collective public key. The y-coordinate of an uninitialized
    ///         (no contributions yet) key is returned as 1 (the identity element).
    function getCollectivePublicKey(bytes12 roundId) external view returns (DKGTypes.Point memory) {
        DKGTypes.Point storage cpk = _collectiveKey[roundId];
        if (cpk.y == 0) {
            return DKGTypes.Point({x: 0, y: 1});
        }
        return cpk;
    }

    /// @notice Aggregate accepted contributions, publish the collective
    ///         public key, and transition the round to Finalized.
    /// @dev    Callable by anyone once `contributionCount ≥
    ///         policy.minValidContributions`. Stores share commitments as
    ///         `keccak256(x, y)` per participant to keep storage to a single
    ///         slot per entry; the pre-image is emitted in `RoundFinalized`.
    function finalizeRound(
        bytes12 roundId,
        bytes32 aggregateCommitmentsHash,
        bytes32 collectivePublicKeyHash,
        bytes32 shareCommitmentHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (round.status == DKGTypes.RoundStatus.Finalized) revert AlreadyFinalized();
        if (round.status != DKGTypes.RoundStatus.Contribution) revert InvalidPhase();
        if (round.contributionCount < round.policy.minValidContributions) revert InsufficientContributions();
        if (
            aggregateCommitmentsHash == bytes32(0) || collectivePublicKeyHash == bytes32(0)
                || shareCommitmentHash == bytes32(0)
        ) revert InvalidFinalization();

        IZKVerifier(FINALIZE_VERIFIER).verifyProof(proof, input);
        uint256[9] memory publicInputs = abi.decode(input, (uint256[9]));
        if (
            publicInputs[0] != _roundScalar(roundId) || publicInputs[1] != round.policy.threshold
                || publicInputs[2] != round.policy.committeeSize || publicInputs[3] != round.contributionCount
                || bytes32(publicInputs[4]) != aggregateCommitmentsHash
                || bytes32(publicInputs[5]) != collectivePublicKeyHash
                || bytes32(publicInputs[6]) != shareCommitmentHash
        ) revert InvalidProofInput();

        uint256 challenge = BRLC.deriveChallenge(
            roundId,
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

        _verifyFinalizeTranscript(roundId, round, challenge, publicInputs[8], transcript);

        round.status = DKGTypes.RoundStatus.Finalized;
        // The three commitment hashes are not persisted to storage; they are emitted
        // in RoundFinalized below and reconstructed off-chain from the event log.

        // Persist share commitments directly from calldata, in the same loop as the
        // already-validated participantIndexes pass.
        uint256 ccount = round.contributionCount;
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
            roundShareCommitmentHashes[roundId][uint16(pIdx)] = keccak256(abi.encode(scX, scY));
        }

        emit RoundFinalized(roundId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash);
    }

    /// @dev SECURITY (C-2). Recompute sk = Σ λ_i d_i mod r from the transcript
    /// and compare it to the claimed ReconstructedSecretHash public input. The
    /// transcript layout is `participantIndexes (N words) ‖ revealedShares (N words)`;
    /// only the first `shareCount` of each are meaningful (validated by
    /// `_verifyReconstructTranscript`).
    function _verifyLagrangeReconstruction(
        uint256 shareCount,
        bytes32 expectedSecret,
        bytes calldata transcript
    ) internal view {
        uint256[] memory xs = new uint256[](shareCount);
        uint256[] memory shares = new uint256[](shareCount);
        uint256 dOff;
        assembly { dOff := transcript.offset }
        uint256 piBase = dOff;
        uint256 svBase = dOff + RECONSTRUCT_VALUES_BYTES_OFFSET;
        for (uint256 i = 0; i < shareCount; i++) {
            uint256 xi;
            uint256 si;
            assembly ("memory-safe") {
                xi := calldataload(add(piBase, mul(i, 0x20)))
                si := calldataload(add(svBase, mul(i, 0x20)))
            }
            xs[i] = xi;
            shares[i] = si;
        }
        Lagrange.verifyReconstruction(xs, shares, shareCount, uint256(expectedSecret));
    }

    /// @dev Verifies the reconstructSecret transcript directly from calldata.
    function _verifyReconstructTranscript(
        bytes12 roundId,
        uint16 committeeSize,
        uint256 shareCount,
        bytes calldata transcript
    ) internal view {
        uint256 dOff;
        assembly { dOff := transcript.offset }
        uint256 piBase = dOff;                                          // participantIndexes
        uint256 svBase = dOff + RECONSTRUCT_VALUES_BYTES_OFFSET;        // revealedShares

        for (uint256 i = 0; i < shareCount; i++) {
            uint256 pIdxRaw;
            uint256 sVal;
            assembly ("memory-safe") {
                pIdxRaw := calldataload(add(piBase, mul(i, 0x20)))
                sVal := calldataload(add(svBase, mul(i, 0x20)))
            }
            uint16 participantIndex = uint16(pIdxRaw);
            if (participantIndex == 0 || participantIndex > committeeSize) revert InvalidProofInput();
            address participant = roundParticipants[roundId][participantIndex - 1];
            DKGTypes.RevealedShareRecord storage record = roundRevealedShares[roundId][participant];
            if (!record.accepted || record.participantIndex != participantIndex) revert InvalidProofInput();
            if (record.shareValue != sVal) revert InvalidProofInput();
        }
        for (uint256 i = shareCount; i < MAX_N; i++) {
            uint256 pIdxRaw;
            uint256 sVal;
            assembly ("memory-safe") {
                pIdxRaw := calldataload(add(piBase, mul(i, 0x20)))
                sVal := calldataload(add(svBase, mul(i, 0x20)))
            }
            if (pIdxRaw != 0 || sVal != 0) revert InvalidProofInput();
        }
    }

    /// @dev Verifies the combineDecryption transcript directly from calldata.
    function _verifyCombineTranscript(
        bytes12 roundId,
        uint16 ciphertextIndex,
        Round storage round,
        uint256 shareCount,
        bytes calldata transcript
    ) internal view {
        uint256 dOff;
        assembly { dOff := transcript.offset }
        uint256 piBase = dOff + 4 * 32;                               // participantIndexes start
        uint256 pdBase = dOff + COMBINE_PARTIALS_BYTES_OFFSET;        // partialDecryptions start

        uint256 cs = round.policy.committeeSize;
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
            address participant = roundParticipants[roundId][participantIndex - 1];
            DKGTypes.PartialDecryptionRecord storage partialRecord =
                roundPartialDecryptions[roundId][ciphertextIndex][participant];
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
        bytes12 roundId,
        Round storage round,
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

        uint256 ccount = round.contributionCount;
        uint256 cSize = round.policy.committeeSize;
        for (uint256 i = 0; i < ccount; i++) {
            uint256 pIdxRaw;
            assembly ("memory-safe") {
                pIdxRaw := calldataload(add(piBase, mul(i, 0x20)))
            }
            uint16 participantIndex = uint16(pIdxRaw);
            if (participantIndex == 0 || participantIndex > cSize) revert InvalidProofInput();
            address participant = roundParticipants[roundId][participantIndex - 1];
            DKGTypes.ContributionRecord storage contribution = roundContributions[roundId][participant];
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
    /// @dev    Keyed by `(roundId, participant, ciphertextIndex)` to support
    ///         multiple ciphertexts per round. The Groth16 proof is a
    ///         Chaum–Pedersen DLEQ establishing that `δ_i` and the committed
    ///         share `D_i` share a discrete log with respect to `C_1` and `G`.
    function submitPartialDecryption(
        bytes12 roundId,
        uint16 participantIndex,
        uint16 ciphertextIndex,
        bytes32 deltaHash,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (round.status != DKGTypes.RoundStatus.Finalized) revert InvalidPhase();
        if (!selectedOperators[roundId][msg.sender]) revert NotSelectedParticipant();
        if (
            participantIndex == 0 || participantIndex > round.policy.committeeSize || ciphertextIndex == 0
                || ciphertextIndex > MAX_CIPHERTEXT_INDEX || deltaHash == bytes32(0)
        ) revert InvalidPartialDecryption();
        if (roundParticipants[roundId][participantIndex - 1] != msg.sender) revert InvalidProofInput();

        DKGTypes.PartialDecryptionRecord storage record = roundPartialDecryptions[roundId][ciphertextIndex][msg.sender];
        if (record.accepted) revert AlreadyPartiallyDecrypted();

        IZKVerifier(PARTIAL_DECRYPT_VERIFIER).verifyProof(proof, input);
        uint256[13] memory publicInputs = abi.decode(input, (uint256[13]));
        bytes32 storedScHash = roundShareCommitmentHashes[roundId][participantIndex];
        if (
            publicInputs[0] != _roundScalar(roundId) || publicInputs[1] != participantIndex
                || storedScHash == bytes32(0)
                || keccak256(abi.encode(publicInputs[4], publicInputs[5])) != storedScHash
        ) revert InvalidProofInput();
        if (deltaHash != keccak256(abi.encodePacked(publicInputs[6], publicInputs[7]))) revert InvalidProofInput();

        // Persist only what combineDecryption actually reads:
        //   - participantIndex + accepted: identity gate
        //   - delta.x/.y: BRLC verification
        // Drop the redundant `participant`, `ciphertextIndex`, and `deltaHash` slots.
        DKGTypes.PartialDecryptionRecord storage prec =
            roundPartialDecryptions[roundId][ciphertextIndex][msg.sender];
        prec.participantIndex = participantIndex;
        prec.ciphertextIndex = ciphertextIndex; // packed in slot 0 anyway
        prec.accepted = true;
        prec.delta.x = publicInputs[6];
        prec.delta.y = publicInputs[7];
        round.partialDecryptionCount++;
        roundPartialDecryptionCounts[roundId][ciphertextIndex]++;

        emit PartialDecryptionSubmitted(roundId, msg.sender, participantIndex, ciphertextIndex, deltaHash);
    }

    /// @notice Submit a ciphertext to be threshold-decrypted by the committee.
    /// @dev    Enforces the round's DecryptionPolicy: owner-only, block/timestamp
    ///         windows, and a cap on accepted ciphertexts per round. Write-once
    ///         per `ciphertextIndex`. Stores `keccak256(c1x, c1y, c2x, c2y)` so
    ///         `combineDecryption` can bind its proof's ciphertext public inputs
    ///         to the authoritative on-chain value. The raw coordinates are only
    ///         exposed via the `CiphertextSubmitted` event (nodes watch it).
    function submitCiphertext(
        bytes12 roundId,
        uint16 ciphertextIndex,
        uint256 c1x,
        uint256 c1y,
        uint256 c2x,
        uint256 c2y
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (round.status != DKGTypes.RoundStatus.Finalized) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX) revert InvalidCiphertext();

        // Well-formedness: coords must be canonical (< Q) and both points must lie on
        // BabyJubJub. Without this, a griefing submitter could pre-claim every index
        // with garbage that the ZK combine circuit can never accept, permanently
        // bricking the round's decryption in open-policy mode. Costs ≈ 4 mulmods per
        // point (~2k gas total) — negligible vs. the rest of the call.
        if (!_isOnBabyJubJub(c1x, c1y) || !_isOnBabyJubJub(c2x, c2y)) revert InvalidCiphertext();

        DKGTypes.DecryptionPolicy memory p = round.decryptionPolicy;
        if (p.ownerOnly && msg.sender != round.organizer) revert NotOwner();
        if (p.notBeforeBlock     != 0 && uint64(block.number)    < p.notBeforeBlock)     revert DecryptionNotYetAllowed();
        if (p.notBeforeTimestamp != 0 && uint64(block.timestamp) < p.notBeforeTimestamp) revert DecryptionNotYetAllowed();
        if (p.notAfterBlock      != 0 && uint64(block.number)    > p.notAfterBlock)      revert DecryptionExpired();
        if (p.notAfterTimestamp  != 0 && uint64(block.timestamp) > p.notAfterTimestamp)  revert DecryptionExpired();
        if (p.maxDecryptions     != 0 && round.ciphertextCount   >= p.maxDecryptions)    revert DecryptionLimitReached();

        if (_ciphertexts[roundId][ciphertextIndex] != bytes32(0)) revert CiphertextAlreadySubmitted();

        _ciphertexts[roundId][ciphertextIndex] = keccak256(abi.encode(c1x, c1y, c2x, c2y));
        unchecked { round.ciphertextCount += 1; }

        emit CiphertextSubmitted(roundId, ciphertextIndex, msg.sender, c1x, c1y, c2x, c2y);
    }

    /// @dev Returns true iff (x, y) are canonical field elements (< Q) and satisfy the
    ///      reduced-form BabyJubJub equation  −x² + y² ≡ 1 + D·x²·y² (mod Q).
    ///      Requiring canonical coords ensures `keccak256(abi.encode(x, y))` in
    ///      `submitCiphertext` matches the combine transcript's canonical `(x, y)` —
    ///      any non-canonical form would bind to a different hash and fail combine.
    function _isOnBabyJubJub(uint256 x, uint256 y) internal pure returns (bool) {
        uint256 q = BabyJubJub_Q;
        if (x >= q || y >= q) return false;
        uint256 xx = mulmod(x, x, q);
        uint256 yy = mulmod(y, y, q);
        // lhs = (-1)·x² + y² = (Q - xx) + yy  (all mod Q)
        uint256 lhs = addmod(mulmod(BabyJubJub_A_NEG, xx, q), yy, q);
        // rhs = 1 + D·x²·y²
        uint256 rhs = addmod(1, mulmod(BabyJubJub_D, mulmod(xx, yy, q), q), q);
        return lhs == rhs;
    }

    /// @notice Combine `t` partial decryptions via Lagrange interpolation and
    ///         persist the recovered plaintext on-chain.
    /// @dev    Callable by anyone once at least `threshold` partial
    ///         decryptions with matching `ciphertextIndex` are on-chain and the
    ///         ciphertext itself has been submitted via `submitCiphertext`.
    ///         The proof's ciphertext public inputs are bound to the stored
    ///         ciphertext hash; a combiner cannot substitute a different ct.
    function combineDecryption(
        bytes12 roundId,
        uint16 ciphertextIndex,
        bytes32 combineHash,
        uint256 plaintext,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (round.status != DKGTypes.RoundStatus.Finalized) revert InvalidPhase();
        if (ciphertextIndex == 0 || ciphertextIndex > MAX_CIPHERTEXT_INDEX || combineHash == bytes32(0)) revert InvalidCombinedDecryption();
        bytes32 storedCtHash = _ciphertexts[roundId][ciphertextIndex];
        if (storedCtHash == bytes32(0)) revert CiphertextNotSubmitted();
        if (roundPartialDecryptionCounts[roundId][ciphertextIndex] < round.policy.threshold) revert InsufficientPartialDecryptions();

        DKGTypes.CombinedDecryptionRecord storage record = roundCombinedDecryptions[roundId][ciphertextIndex];
        if (record.completed) revert AlreadyCombined();

        IZKVerifier(DECRYPT_COMBINE_VERIFIER).verifyProof(proof, input);
        uint256[7] memory publicInputs = abi.decode(input, (uint256[7]));
        if (
            publicInputs[0] != _roundScalar(roundId) || publicInputs[1] != round.policy.threshold
                || bytes32(publicInputs[3]) != combineHash || publicInputs[4] != plaintext
        ) revert InvalidProofInput();
        if (publicInputs[2] < round.policy.threshold) revert InvalidProofInput();
        uint256 challenge = BRLC.deriveChallenge(
            roundId,
            DECRYPT_COMBINE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(combineHash, bytes32(plaintext)))
        );
        if (publicInputs[5] != challenge) revert InvalidProofInput();

        // Transcript layout (4 + 3N words):
        //   words [0..4)         ciphertext
        //   words [4..4+N)       participantIndexes
        //   words [4+N..4+3N)    partialDecryptions  (N points × 2 coords)
        if (transcript.length != COMBINE_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        // Bind the combine proof's ciphertext public inputs (first 4 transcript words)
        // to the on-chain ciphertext hash written at submitCiphertext time.
        if (keccak256(transcript[0:128]) != storedCtHash) revert InvalidProofInput();
        _verifyCombineTranscript(roundId, ciphertextIndex, round, publicInputs[2], transcript);

        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, COMBINE_TRANSCRIPT_WORDS) != publicInputs[6]) revert InvalidProofInput();

        record.completed = true;
        record.plaintext = plaintext;

        emit DecryptionCombined(roundId, ciphertextIndex, combineHash, plaintext);
    }

    /// @notice Publish a committee member's secret share `d_i` under the
    ///         disclosure path.
    /// @dev    Only callable when the round was created with
    ///         `disclosureAllowed = true`. The Groth16 proof establishes
    ///         `d_i · G = D_i`, binding the revealed scalar to the on-chain
    ///         share commitment.
    function submitRevealedShare(
        bytes12 roundId,
        uint16 participantIndex,
        uint256 shareValue,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (!round.policy.disclosureAllowed) revert DisclosureDisabled();
        if (round.status != DKGTypes.RoundStatus.Finalized) revert InvalidPhase();
        if (!selectedOperators[roundId][msg.sender]) revert NotSelectedParticipant();
        if (participantIndex == 0 || participantIndex > round.policy.committeeSize || shareValue == 0) {
            revert InvalidRevealedShare();
        }
        if (roundParticipants[roundId][participantIndex - 1] != msg.sender) revert InvalidProofInput();

        DKGTypes.RevealedShareRecord storage record = roundRevealedShares[roundId][msg.sender];
        if (record.accepted) revert AlreadyRevealed();

        IZKVerifier(REVEAL_SUBMIT_VERIFIER).verifyProof(proof, input);
        uint256[5] memory publicInputs = abi.decode(input, (uint256[5]));
        bytes32 storedScHash = roundShareCommitmentHashes[roundId][participantIndex];
        if (
            publicInputs[0] != _roundScalar(roundId) || publicInputs[1] != participantIndex || publicInputs[2] != shareValue
                || storedScHash == bytes32(0)
                || keccak256(abi.encode(publicInputs[3], publicInputs[4])) != storedScHash
        ) revert InvalidProofInput();

        bytes32 shareHash = bytes32(shareValue);

        // Persist only what reconstructSecret reads:
        //   - shareValue (used in the BRLC verify)
        //   - participantIndex + accepted (identity gate)
        // Drop redundant `participant` and `shareHash`.
        DKGTypes.RevealedShareRecord storage rrec = roundRevealedShares[roundId][msg.sender];
        rrec.participantIndex = participantIndex;
        rrec.accepted = true;
        rrec.shareValue = shareValue;
        round.revealedShareCount++;

        emit RevealedShareSubmitted(roundId, msg.sender, participantIndex, shareHash);
    }

    /// @notice Reconstruct the round secret `sk = F(0)` from `≥ t` revealed
    ///         shares via Lagrange interpolation and transition the round to
    ///         Completed.
    /// @dev    Only callable when `disclosureAllowed = true`.
    function reconstructSecret(
        bytes12 roundId,
        bytes32 disclosureHash,
        bytes32 reconstructedSecretHash,
        bytes calldata transcript,
        bytes calldata proof,
        bytes calldata input
    ) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (!round.policy.disclosureAllowed) revert DisclosureDisabled();
        if (round.status != DKGTypes.RoundStatus.Finalized) revert InvalidPhase();
        if (disclosureHash == bytes32(0) || reconstructedSecretHash == bytes32(0)) revert InvalidReconstruction();
        if (round.revealedShareCount < round.policy.threshold) revert InsufficientRevealedShares();

        IZKVerifier(REVEAL_SHARE_VERIFIER).verifyProof(proof, input);
        uint256[7] memory publicInputs = abi.decode(input, (uint256[7]));
        if (
            publicInputs[0] != _roundScalar(roundId) || publicInputs[1] != round.policy.threshold
                || bytes32(publicInputs[3]) != disclosureHash || bytes32(publicInputs[4]) != reconstructedSecretHash
        ) revert InvalidProofInput();
        if (publicInputs[2] < round.policy.threshold) revert InvalidProofInput();
        uint256 challenge = BRLC.deriveChallenge(
            roundId,
            REVEAL_SHARE_TRANSCRIPT_DOMAIN,
            keccak256(abi.encodePacked(disclosureHash, reconstructedSecretHash))
        );
        if (publicInputs[5] != challenge) revert InvalidProofInput();

        // Transcript layout (2N words):
        //   words [0..N)   participantIndexes
        //   words [N..2N)  revealedShares
        if (transcript.length != RECONSTRUCT_TRANSCRIPT_WORDS * 32) revert InvalidProofInput();
        _verifyReconstructTranscript(roundId, round.policy.committeeSize, publicInputs[2], transcript);

        uint256 dOff;
        assembly { dOff := transcript.offset }
        if (BRLC.commitCalldata(challenge, dOff, RECONSTRUCT_TRANSCRIPT_WORDS) != publicInputs[6]) revert InvalidProofInput();

        // SECURITY (C-2): the reveal-share circuit deliberately does NOT
        // prove the Lagrange identity sk = Σ λ_i d_i. Without this on-chain
        // recomputation an attacker can publish any value as the round's
        // reconstructed secret. We recompute it here from the same transcript
        // calldata that BRLC just committed to and require it to match the
        // claimed ReconstructedSecretHash public input.
        _verifyLagrangeReconstruction(publicInputs[2], reconstructedSecretHash, transcript);

        round.status = DKGTypes.RoundStatus.Completed;
        // disclosureHash + reconstructedSecretHash are emitted in the event.

        emit SecretReconstructed(roundId, disclosureHash, reconstructedSecretHash);
    }

    /// @notice Abort a non-terminal round. Organizer only.
    /// @dev    Finalized rounds may NOT be aborted: the collective public key has
    ///         already been published and messages may already be encrypted to it.
    ///         Aborting after finalization would permanently block decryption for
    ///         those messages. Only Registration and Contribution phases are
    ///         abortable.
    /// @param  roundId The round identifier.
    function abortRound(bytes12 roundId) external {
        Round storage round = rounds[roundId];
        if (round.organizer == address(0)) revert InvalidRound();
        if (msg.sender != round.organizer) revert Unauthorized();
        if (
            round.status == DKGTypes.RoundStatus.Finalized
                || round.status == DKGTypes.RoundStatus.Completed
                || round.status == DKGTypes.RoundStatus.Aborted
        ) {
            revert InvalidPhase();
        }

        round.status = DKGTypes.RoundStatus.Aborted;
        emit RoundAborted(roundId);
    }

    function getRound(bytes12 roundId) external view returns (Round memory) {
        return rounds[roundId];
    }

    function selectedParticipants(bytes12 roundId) external view returns (address[] memory) {
        return roundParticipants[roundId];
    }

    function getContribution(bytes12 roundId, address contributor)
        external
        view
        returns (DKGTypes.ContributionRecord memory)
    {
        return roundContributions[roundId][contributor];
    }

    function getPartialDecryption(bytes12 roundId, address participant, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.PartialDecryptionRecord memory)
    {
        return roundPartialDecryptions[roundId][ciphertextIndex][participant];
    }

    function getCombinedDecryption(bytes12 roundId, uint16 ciphertextIndex)
        external
        view
        returns (DKGTypes.CombinedDecryptionRecord memory)
    {
        return roundCombinedDecryptions[roundId][ciphertextIndex];
    }

    function getRevealedShare(bytes12 roundId, address participant)
        external
        view
        returns (DKGTypes.RevealedShareRecord memory)
    {
        return roundRevealedShares[roundId][participant];
    }

    /// @notice Returns the keccak256(abi.encode(x, y)) commitment hash for a
    /// participant's share commitment. The pre-image (x,y) is exposed off-chain via
    /// the RoundFinalized event log.
    function getShareCommitmentHash(bytes12 roundId, uint16 participantIndex)
        external
        view
        returns (bytes32)
    {
        return roundShareCommitmentHashes[roundId][participantIndex];
    }

    /// @notice keccak256(abi.encode(c1x, c1y, c2x, c2y)) of the ciphertext stored
    ///         at `ciphertextIndex` for `roundId`. Returns bytes32(0) if no
    ///         ciphertext has been submitted at this slot.
    function getCiphertextHash(bytes12 roundId, uint16 ciphertextIndex) external view returns (bytes32) {
        return _ciphertexts[roundId][ciphertextIndex];
    }

    /// @notice Recovered plaintext for (roundId, ciphertextIndex). Returns 0 if
    ///         the decryption has not been combined yet; callers should also
    ///         consult `getCombinedDecryption(...)` / `DecryptionCombined`
    ///         events to disambiguate "not yet combined" from "plaintext is 0".
    function getPlaintext(bytes12 roundId, uint16 ciphertextIndex) external view returns (uint256) {
        return roundCombinedDecryptions[roundId][ciphertextIndex].plaintext;
    }

    function getDecryptionPolicy(bytes12 roundId) external view returns (DKGTypes.DecryptionPolicy memory) {
        return rounds[roundId].decryptionPolicy;
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

    function getRevealSubmitVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(REVEAL_SUBMIT_VERIFIER).provingKeyHash();
    }

    function getRevealShareVerifierVKeyHash() external view returns (bytes32) {
        return IZKVerifier(REVEAL_SHARE_VERIFIER).provingKeyHash();
    }

    function _roundScalar(bytes12 roundId) internal pure returns (uint256) {
        return uint256(uint96(roundId));
    }
}
