// Circuit bounds and transcript layouts shared with the Go circuits and the
// Solidity contracts. Mirrors `circuits/common/sizes.go`,
// `circuits/contribution/layout.go`, `circuits/finalize/witness.go` and
// `solidity/src/libraries/Sizes.sol` — a one-word divergence here makes a
// decoded transcript, a Merkle path or a BRLC commitment silently wrong, so
// every offset below is pinned by `sdk/tests/vectors.test.ts` against the
// generated `tests/vectors/*.json`.
//
// A "word" is a 32-byte big-endian field element; every offset is in words.

/** `MaxN`: the committee cap every circuit is compiled for. */
export const MAX_N = 32;
/** `MaxK`: pool keys dealt per epoch, one per application. */
export const MAX_K = 16;
/** `log2(MAX_N)`: depth of the share-commitment Merkle tree of one pool key. */
export const MERKLE_DEPTH = 5;

// ─── Compact contribution transcript (docs/pool-keys-v4.md §3, §5) ───────────
//
// `L_C = MaxK·(2t+n) + 5n` words, no padding, regions in this order:
//
//   [0, 2Kt)          for j, for m < t: A[j][m].x, A[j][m].y
//   [2Kt, 2Kt+n)      recipient indexes, committee order (slot i holds i+1)
//   [2Kt+n, 2Kt+3n)   recipient public keys (x, y)
//   [2Kt+3n, 2Kt+5n)  ephemerals (x, y)
//   [2Kt+5n, L_C)     for j, for i < n: masked share ms[j][i]
//
// The offsets are functions of the epoch's public `threshold` and
// `committeeSize`, never of the calldata itself.

/** `L_C(t, n) = MaxK·(2t+n) + 5n`, the compact contribution transcript length in words. */
export function contributionTranscriptWords(threshold: number, committeeSize: number): number {
  return MAX_K * (2 * threshold + committeeSize) + 5 * committeeSize;
}

/**
 * Word offsets of one epoch policy's compact contribution transcript. The
 * constructor validates `1 ≤ t ≤ n ≤ MaxN`, as the Go `NewLayout` does.
 */
export class ContributionLayout {
  readonly threshold: number;
  readonly committeeSize: number;

  constructor(threshold: number, committeeSize: number) {
    if (!Number.isInteger(threshold) || threshold < 1) {
      throw new Error(`contribution layout: threshold must be at least 1, got ${threshold}`);
    }
    if (!Number.isInteger(committeeSize) || committeeSize > MAX_N) {
      throw new Error(`contribution layout: committee size ${committeeSize} exceeds max ${MAX_N}`);
    }
    if (threshold > committeeSize) {
      throw new Error(`contribution layout: threshold ${threshold} exceeds committee size ${committeeSize}`);
    }
    this.threshold = threshold;
    this.committeeSize = committeeSize;
  }

  /** `L_C`, the transcript length in words. */
  get words(): number {
    return contributionTranscriptWords(this.threshold, this.committeeSize);
  }

  /** The exact calldata length the contract requires. */
  get bytes(): number {
    return 32 * this.words;
  }

  /** Word of `A[key][coefficient].x`; `.y` follows. */
  commitmentOffset(key: number, coefficient: number): number {
    return 2 * (key * this.threshold + coefficient);
  }

  /** `2Kt`, the first recipient-index word. */
  get recipientIndexesStart(): number {
    return 2 * MAX_K * this.threshold;
  }

  /** Word holding recipient slot `i`'s index (which must equal `i + 1`). */
  recipientIndexOffset(recipient: number): number {
    return this.recipientIndexesStart + recipient;
  }

  /** `2Kt + n`, the first recipient public-key word. */
  get recipientKeysStart(): number {
    return this.recipientIndexesStart + this.committeeSize;
  }

  /** Word of recipient slot `i`'s public key `x`; `y` follows. */
  recipientKeyOffset(recipient: number): number {
    return this.recipientKeysStart + 2 * recipient;
  }

  /** `2Kt + 3n`, the first ephemeral word. */
  get ephemeralsStart(): number {
    return this.recipientIndexesStart + 3 * this.committeeSize;
  }

  /** Word of recipient slot `i`'s ephemeral `x`; `y` follows. */
  ephemeralOffset(recipient: number): number {
    return this.ephemeralsStart + 2 * recipient;
  }

  /** `2Kt + 5n`, the first masked-share word. */
  get maskedSharesStart(): number {
    return this.recipientIndexesStart + 5 * this.committeeSize;
  }

  /** Word of `ms[key][recipient]`. */
  maskedShareOffset(key: number, recipient: number): number {
    return this.maskedSharesStart + key * this.committeeSize + recipient;
  }

  /**
   * The half-open word interval `[2Kt, 2Kt+3n)` — indexes followed by public
   * keys, exactly `3n` words — that `DKGManager` keccaks and compares against
   * its committee snapshot.
   */
  get committeeRegion(): { start: number; end: number } {
    return { start: this.recipientIndexesStart, end: this.ephemeralsStart };
  }
}

// ─── Finalization transcript (docs/pool-keys-v4.md §7) ───────────────────────
//
// Fixed `L_F = 2·MaxN + MaxK·(2 + 2·MaxN)` words (1,120 at the current bounds):
//
//   [0, N)      participant indexes I[d]   (0 for rows d ≥ acceptedCount)
//   [N, 2N)     contribution hashes h[d]   (0 for rows d ≥ acceptedCount)
//   then per key j, a (2 + 2N)-word block at 2N + j·(2 + 2N):
//     P[j].x, P[j].y, D[j][0].x, D[j][0].y, …, D[j][N−1].x, D[j][N−1].y
//                                (D[j][i] = identity (0, 1) for i ≥ committeeSize)
//
// Slot `i` of each key's share commitments is committee member `i + 1` —
// contributor or not — and is that member's Merkle leaf.

/** Words per key block: `P_j` plus `MaxN` share commitments. */
export const FINALIZE_KEY_WORDS = 2 + 2 * MAX_N;
/** `L_F`: the finalization transcript length in words. */
export const FINALIZE_TRANSCRIPT_WORDS = 2 * MAX_N + MAX_K * FINALIZE_KEY_WORDS;
/** First participant-index word. */
export const FINALIZE_INDEXES_START = 0;
/** First contribution-hash word. */
export const FINALIZE_HASHES_START = MAX_N;
/** First word of key 0's block. */
export const FINALIZE_KEYS_START = 2 * MAX_N;

/** First word of key `j`'s block, which holds `P[j].x`. */
export function finalizeKeyOffset(key: number): number {
  return FINALIZE_KEYS_START + key * FINALIZE_KEY_WORDS;
}

/** Word of `P[key].x`; `.y` follows. */
export function finalizePoolKeyOffset(key: number): number {
  return finalizeKeyOffset(key);
}

/**
 * Word of `D[key][member].x` (`.y` follows), `member` zero-based: the share
 * commitment of committee position `member + 1`.
 */
export function finalizeShareCommitmentOffset(key: number, member: number): number {
  return finalizeKeyOffset(key) + 2 + 2 * member;
}

// ─── Decrypt-combine transcript (unchanged by v4) ────────────────────────────

/** `6 + 3·MaxN`: ciphertext, `PK_org`, participant indexes, partial decryptions. */
export const COMBINE_TRANSCRIPT_WORDS = 6 + 3 * MAX_N;
