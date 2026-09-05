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
///      time. All of them are proven and stored by the single proof-carrying
///      `finalizeEpoch`, and the pool cursor plus key indexes fit a `uint8`.
uint256 constant MAX_K = 16;

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
// ── submitContribution (compact, v4) ────────────────────────────────────────
//
// With `t` = threshold, `n` = committeeSize and `L_C = K·(2t+n)+5n` words
// (`contribTranscriptWords` below): no inactive padding travels in calldata,
// the fixed-width regions are simply truncated at the live counts.
//
//   words [0, 2Kt)              commitments, key-major:
//                               for j in [0, K): for m in [0, t): A[j][m].x, A[j][m].y
//   words [2Kt, 2Kt+n)          recipientIndexes   (word i MUST be i+1)
//   words [2Kt+n, 2Kt+3n)       recipientPubKeys (x, y)
//   words [2Kt+3n, 2Kt+5n)      ephemerals (x, y)
//   words [2Kt+5n, L_C)         maskedShares, key-major:
//                               for j in [0, K): for i in [0, n): ms[j][i]
//
// Must equal `len(contribution.TranscriptScalars())` on the Go side.

/// @dev `L_C = MAX_K·(2t+n) + 5n` — the compact contribution transcript
///      length in words, derived from the epoch policy (t = threshold,
///      n = committeeSize). Mirrors `ContributionTranscriptWords(t, n)` in
///      `circuits/common/sizes.go` on the Go side.
function contribTranscriptWords(uint256 t, uint256 n) pure returns (uint256) {
    return MAX_K * (2 * t + n) + 5 * n;
}

/// @dev Byte offset of the committee section (recipientIndexes ‖
///      recipientPubKeys) inside the compact contribution transcript.
///      `submitContribution` hashes
///      `[contribCommitteeBytesOffset(t), contribCommitteeBytesEnd(t, n))`
///      in one keccak and compares it against the unpadded `3n`-word
///      snapshot taken when the lottery filled.
function contribCommitteeBytesOffset(uint256 t) pure returns (uint256) {
    return 2 * MAX_K * t * 32;
}

function contribCommitteeBytesEnd(uint256 t, uint256 n) pure returns (uint256) {
    return (2 * MAX_K * t + 3 * n) * 32;
}

// ── finalizeEpoch (proof-carrying, v4) ──────────────────────────────────────
//
//   words [0, N)        participantIndexes           (0 for rows >= acceptedCount)
//   words [N, 2N)       contributionHashes           (0 for rows >= acceptedCount)
//   then per key j in [0, K), a (2 + 2N)-word row:
//     P[j].x, P[j].y, D[j][0].x, D[j][0].y, …, D[j][N-1].x, D[j][N-1].y
//                       (D[j][i] = identity (0, 1) for i >= committeeSize)
//
// Rows `i < acceptedCount` of the first two regions name the accepted
// contributors, in any order (builders SHOULD emit ascending indexes). Slot
// `i` of each key's share commitments is the share commitment of committee
// member `p = i + 1` — contributor or not — so every member's leaf exists in
// the tree. `finalizeEpoch` verifies one Groth16 proof over the whole vector,
// stores all MAX_K keys and Merkle roots, and flips the epoch Live.
// Must equal `len(finalize.TranscriptScalars())` on the Go side.
uint256 constant FINALIZE_TRANSCRIPT_WORDS = 2 * MAX_N + MAX_K * (2 + 2 * MAX_N);
/// @dev Byte offset of the per-key rows: key `j`'s `P[j]` sits at word
///      `2·MAX_N + j·(2 + 2·MAX_N)` of the finalize transcript.
uint256 constant FINALIZE_KEY_WORDS_STRIDE = 2 + 2 * MAX_N;

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
