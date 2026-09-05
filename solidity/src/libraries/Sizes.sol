// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/// @dev Maximum DKG committee size. SINGLE source of truth for the on-chain
///      side — must equal `circuits/common.MaxN` on the Go side. Update both
///      and run `make circuits` to regenerate verifier wrappers and bindings.
uint256 constant MAX_N = 32;

/// @dev Number of independent keys dealt by every epoch's DKG (the "pool").
///      Must equal `circuits/common.MaxK` on the Go side. Each registered
///      application claims exactly one of them, so an epoch serves at most
///      `MAX_K` applications; unused keys cost only calldata at contribution
///      time. Bounded by 8 because the per-epoch activation bitmap is a
///      `uint8`.
uint256 constant MAX_K = 8;

/// @dev Depth of the per-pool-key Merkle tree over the committee's share
///      commitments `D_i`. Must equal `log2(MAX_N)` (and
///      `circuits/common.MerkleDepth` on the Go side): the tree always has
///      exactly `MAX_N` leaves. Leaf `i` (participant `i + 1`) is
///      `keccak256(0x00 ‖ D.x ‖ D.y)` for `i < committeeSize` and
///      `MERKLE_EMPTY_LEAF` otherwise; an internal node is
///      `keccak256(0x01 ‖ left ‖ right)`. The one-byte tags keep leaves and
///      nodes in separate hash domains.
uint256 constant MERKLE_DEPTH = 5;
/// @dev Leaf of a committee slot that does not exist (`i >= committeeSize`).
///      Must equal `circuits/common.MerkleEmptyLeaf` on the Go side.
bytes32 constant MERKLE_EMPTY_LEAF = keccak256("davinci-dkg:merkle-empty:v1");

// ─── Derived transcript sizes ────────────────────────────────────────────────
//
// A "word" is 32 bytes. Every transcript below is a fixed-width vector of
// field elements read straight out of calldata by `DKGManager`; the Go
// witness builders emit the identical layout, so a one-word divergence makes
// the BRLC commitment (and therefore the proof) fail.
//
// ── submitContribution ──────────────────────────────────────────────────────
//
//   words [0, 2KN)              commitments, key-major:
//                               for j in [0, K): for m in [0, N): A[j][m].x, A[j][m].y
//   words [2KN, 2KN+N)          recipientIndexes           (0 in inactive slots)
//   words [2KN+N, 2KN+3N)       recipientPubKeys (x, y)    (identity when inactive)
//   words [2KN+3N, 2KN+5N)      ephemerals (x, y)          (identity when inactive)
//   words [2KN+5N, 3KN+5N)      maskedShares, key-major:
//                               for j in [0, K): for i in [0, N): ms[j][i]
//
// Must equal `len(contribution.TranscriptScalars())` on the Go side.
uint256 constant CONTRIB_TRANSCRIPT_WORDS = 3 * MAX_K * MAX_N + 5 * MAX_N;
/// @dev Byte offset of the committee section (recipientIndexes ‖ recipientPubKeys)
///      inside the contribution transcript. `submitContribution` hashes
///      `[CONTRIB_COMMITTEE_BYTES_OFFSET, CONTRIB_COMMITTEE_BYTES_END)` in one
///      keccak and compares it against the snapshot taken when the lottery filled.
uint256 constant CONTRIB_COMMITTEE_BYTES_OFFSET = (2 * MAX_K * MAX_N) * 32;
uint256 constant CONTRIB_COMMITTEE_BYTES_END    = (2 * MAX_K * MAX_N + 3 * MAX_N) * 32;

// ── activatePoolKey ─────────────────────────────────────────────────────────
//
//   words [0, N)      participantIndexes           (0 for i >= acceptedCount)
//   words [N, 2N)     contributionHashes           (0 for i >= acceptedCount)
//   words [2N, 4N)    aggregateCommitments (x, y)  (identity for m >= t)
//   words [4N, 6N)    shareCommitments D_p (x, y)  (identity for i >= committeeSize)
//
// Rows `i < acceptedCount` of the first two regions name the accepted
// contributors, in any order. Slot `i` of `shareCommitments` is the share
// commitment of committee member `p = i + 1` — contributor or not — so every
// member's leaf exists in the tree.
// Must equal `len(poolkey.TranscriptScalars())` on the Go side.
uint256 constant POOLKEY_TRANSCRIPT_WORDS = 6 * MAX_N;
/// @dev Byte offset of `contributionHashes` (the per-contributor Poseidon
///      `commitmentsHash` the contract re-checks against storage).
uint256 constant POOLKEY_HASHES_BYTES_OFFSET = MAX_N * 32;
/// @dev Word offset of `aggregateCommitments`; `aggregateCommitments[0]` is the
///      pool key `P_j` itself.
uint256 constant POOLKEY_AGG_WORDS_OFFSET = 2 * MAX_N;
/// @dev Word offset of `shareCommitments`; leaf `i` of the share Merkle tree
///      is `keccak256(0x00 ‖ D_{i+1}.x ‖ D_{i+1}.y)`.
uint256 constant POOLKEY_SHARE_WORDS_OFFSET = 4 * MAX_N;

// ── combineDecryption ───────────────────────────────────────────────────────
//
//   words [0, 4)      C1.x C1.y C2.x C2.y
//   words [4, 6)      PK_org.x PK_org.y            (identity for Automatic apps)
//   words [6, 6+N)    participant indexes x_k      (0 in inactive slots)
//   words [6+N, 6+3N) partial decryptions δ_k      (identity when inactive)
//
// Must equal `len(decryptcombine.TranscriptScalars())` on the Go side.
uint256 constant COMBINE_TRANSCRIPT_WORDS = 6 + 3 * MAX_N;
/// @dev Ciphertext (4 words) + organizer key (2 words) — the fixed head that
///      precedes the per-participant section.
uint256 constant COMBINE_HEAD_WORDS = 6;
uint256 constant COMBINE_INDEXES_BYTES_OFFSET  = COMBINE_HEAD_WORDS * 32;
uint256 constant COMBINE_PARTIALS_BYTES_OFFSET = (COMBINE_HEAD_WORDS + MAX_N) * 32;

// ─── Epoch scheduling ────────────────────────────────────────────────────────
//
// Every duration is expressed in BLOCKS (chain-agnostic). The wall-clock time
// depends on the deployment chain's block time and is estimated off-chain by
// the SDK / UI from sampled block timestamps.
//
// All four block constants below are set per-deploy as `DKGManager`
// constructor immutables; the defaults match the values in
// `script/DeployAll.s.sol` when no env override is provided. They use ~12 s
// blocks (Sepolia / mainnet) — for chains with different block times scale
// the *_BLOCKS constants accordingly so wall-time matches.
//
// Lifecycle (Preparation = first three; Service = the rest):
//
//   [0,  CSL)               CommitteeSelection : claimSlot accepted
//   [CSL, CSL+KA)           KeyAssembly        : submitContribution accepted
//   [CSL+KA, +GAP)          finalize gap       : finalizeEpoch may run
//   [CSL+KA+GAP, END)       Live               : pool keys activate, apps
//                                                 register, ciphertexts decrypt
//
// CommitteeSelection and KeyAssembly are absolute, NOT proportional. The
// lottery is one keccak per claimer and the contribution proof is one tx
// per committee member, so they need a fixed budget — not a fraction of the
// epoch. Long epochs (multi-day) keep the same short Preparation; the extra
// time falls into Service.
//
// After END blocks anyone may call `createEpoch` again to mint the next
// epoch (permissionless), and also earlier when the newest epoch is `Live`
// with at most one unclaimed pool key left (an epoch only serves `MAX_K`
// applications) or has been `Aborted`. The previous epoch stays `Live` for
// the duration of its Service window, so its keys remain usable while the
// next epoch bootstraps.

uint256 constant DEFAULT_EPOCH_DURATION_BLOCKS       = 100;  // ~20 min @ 12 s
uint256 constant DEFAULT_COMMITTEE_SELECTION_BLOCKS  = 25;   // ~5 min  @ 12 s
uint256 constant DEFAULT_KEY_ASSEMBLY_BLOCKS         = 25;   // ~5 min  @ 12 s
uint256 constant DEFAULT_FINALIZE_GAP_BLOCKS         = 5;    // ~1 min  @ 12 s

uint256 constant SEED_DELAY_BLOCKS = 1;
// `blockhash(startBlock + SEED_DELAY_BLOCKS)` is the lottery seed. The
// constant must be strictly less than `COMMITTEE_SELECTION_BLOCKS` so
// claimers get at least one block to call after the seed resolves; the
// constructor enforces this.
