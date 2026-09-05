// Transcript codecs for the two proof-carrying calls whose calldata the SDK
// has to read back: the compact contribution (`submitContribution`) and the
// batched finalization (`finalizeEpoch`). Mirrors `circuits/common/brlc.go`,
// `circuits/contribution/layout.go`, `circuits/finalize/witness.go` and the
// `BRLC` library in Solidity; the cross-impl values are pinned by
// `sdk/tests/vectors.test.ts` against `tests/vectors/{contribution_compact,
// finalize_transcript}.json`.
//
// Every word is a 32-byte big-endian integer strictly below the BN254 scalar
// field modulus. Decoders reject — never reduce — a non-canonical word, so a
// transcript has exactly one encoding (docs/pool-keys-v4.md §2).

import { decodeFunctionData, keccak256, type Hex } from 'viem';
import {
  poseidon1, poseidon2, poseidon3, poseidon4, poseidon5, poseidon6, poseidon7, poseidon8,
  poseidon9, poseidon10, poseidon11, poseidon12, poseidon13, poseidon14, poseidon15, poseidon16,
} from 'poseidon-lite';
import { dkgManagerAbi } from './abi.js';
import {
  DomainContributionTranscriptV2,
  DomainFinalizeTranscriptV2,
} from './protocol.js';
import {
  ContributionLayout,
  FINALIZE_HASHES_START,
  FINALIZE_INDEXES_START,
  FINALIZE_KEY_WORDS,
  FINALIZE_KEYS_START,
  FINALIZE_TRANSCRIPT_WORDS,
  MAX_K,
  MAX_N,
  finalizeKeyOffset,
  finalizePoolKeyOffset,
  finalizeShareCommitmentOffset,
} from './sizes.js';

/** BN254 scalar-field modulus `p`: the field every transcript word lives in. */
export const FR_MODULUS =
  21888242871839275222246405745257275088548364400416034343698204186575808495617n;

/** A BabyJubJub point as the contract stores it: two field words (RTE form). */
export interface TranscriptPoint {
  x: bigint;
  y: bigint;
}

const IDENTITY: TranscriptPoint = { x: 0n, y: 1n };

// ─── word codec ──────────────────────────────────────────────────────────────

function stripHex(h: Hex): string {
  const s = h.startsWith('0x') ? h.slice(2) : h;
  if (s.length % 2 !== 0 || !/^[0-9a-fA-F]*$/.test(s)) {
    throw new Error('transcript: malformed hex');
  }
  return s;
}

/** Throws unless `word` is a canonical field element in `[0, p)`. */
export function requireCanonicalWord(word: bigint, position: number): void {
  if (word < 0n || word >= FR_MODULUS) {
    throw new Error(`transcript: word ${position} is not a canonical field element`);
  }
}

/**
 * Split a `bytes` transcript into its 32-byte words. Rejects a length that is
 * not a whole number of words and any word `≥ p`.
 */
export function wordsFromBytes(data: Hex): bigint[] {
  const s = stripHex(data);
  if (s.length % 64 !== 0) {
    throw new Error(`transcript: ${s.length / 2} bytes is not a whole number of 32-byte words`);
  }
  const out: bigint[] = [];
  for (let i = 0; i < s.length; i += 64) {
    const word = BigInt('0x' + s.slice(i, i + 64));
    requireCanonicalWord(word, i / 64);
    out.push(word);
  }
  return out;
}

/** Pack words as 32-byte big-endian values — the `bytes transcript` calldata. */
export function wordsToBytes(words: readonly bigint[]): Hex {
  let out = '0x';
  words.forEach((word, i) => {
    requireCanonicalWord(word, i);
    out += word.toString(16).padStart(64, '0');
  });
  return out as Hex;
}

/** `keccak256(abi.encodePacked(bytes32...))` over the words. */
export function keccakWords(words: readonly bigint[]): Hex {
  return keccak256(wordsToBytes(words));
}

/** `bytes32` big-endian encoding of a field element. */
export function bytes32(value: bigint): Hex {
  if (value < 0n || value >= 1n << 256n) throw new Error('bytes32: value out of range');
  return ('0x' + value.toString(16).padStart(64, '0')) as Hex;
}

// ─── Fiat–Shamir: anchor, challenge, BRLC ────────────────────────────────────

/**
 * The BRLC challenge anchor: `keccak(digest_1 ‖ … ‖ digest_k ‖ keccak(words))`.
 * Every value the prover chooses — the calldata transcript and the Poseidon
 * digests that fix the witness — is hashed in before `ρ` exists. Mirrors
 * `ccommon.ChallengeAnchor`.
 */
export function challengeAnchor(words: readonly bigint[], ...digests: bigint[]): Hex {
  const transcriptHash = keccakWords(words);
  let packed = '0x';
  for (const digest of digests) packed += bytes32(digest).slice(2);
  packed += transcriptHash.slice(2);
  return keccak256(packed as Hex);
}

/**
 * `ρ = keccak256(eid12 ‖ domain32 ‖ anchor32) mod p` — `BRLC.deriveChallenge`
 * on chain and `ccommon.DeriveChallengeNative` in Go. `epochId` is the raw
 * 12-byte epoch identifier.
 */
export function deriveChallenge(epochId: Hex, domain: Hex, anchor: Hex): bigint {
  const eid = stripHex(epochId);
  if (eid.length !== 24) throw new Error(`deriveChallenge: epochId must be 12 bytes, got ${eid.length / 2}`);
  const dom = stripHex(domain);
  if (dom.length !== 64) throw new Error('deriveChallenge: domain must be 32 bytes');
  const anc = stripHex(anchor);
  if (anc.length !== 64) throw new Error('deriveChallenge: anchor must be 32 bytes');
  return BigInt(keccak256(('0x' + eid + dom + anc) as Hex)) % FR_MODULUS;
}

/**
 * Binding random linear combination `Σ ρ^(q+1)·w[q] mod p` over the words in
 * order — what `BRLC.commitCalldata` streams from calldata and what the
 * circuits' (gated) fold must equal.
 */
export function brlcCommit(challenge: bigint, words: readonly bigint[]): bigint {
  const rho = ((challenge % FR_MODULUS) + FR_MODULUS) % FR_MODULUS;
  let acc = 0n;
  let power = rho;
  for (const word of words) {
    acc = (acc + power * word) % FR_MODULUS;
    power = (power * rho) % FR_MODULUS;
  }
  return acc;
}

// ─── Poseidon MultiHash ──────────────────────────────────────────────────────

const POSEIDON = [
  poseidon1, poseidon2, poseidon3, poseidon4, poseidon5, poseidon6, poseidon7, poseidon8,
  poseidon9, poseidon10, poseidon11, poseidon12, poseidon13, poseidon14, poseidon15, poseidon16,
] as const;

function poseidonChunk(inputs: readonly bigint[]): bigint {
  const hash = POSEIDON[inputs.length - 1];
  if (!hash) throw new Error(`poseidon: unsupported arity ${inputs.length}`);
  return hash(inputs as bigint[]);
}

/**
 * The Poseidon sponge every circuit digest goes through (davinci-node's
 * `MultiPoseidon`, gnark-crypto-primitives `MultiHash`): up to 16 inputs are
 * hashed directly; longer vectors are hashed in 16-input chunks and the chunk
 * hashes are folded again the same way. Inputs are plain field elements.
 */
export function multiPoseidon(inputs: readonly bigint[]): bigint {
  if (inputs.length === 0) throw new Error('multiPoseidon: no inputs');
  if (inputs.length <= 16) return poseidonChunk(inputs);
  const hashes: bigint[] = [];
  for (let i = 0; i < inputs.length; i += 16) hashes.push(poseidonChunk(inputs.slice(i, i + 16)));
  if (hashes.length === 1) return hashes[0];
  if (hashes.length <= 16) return poseidonChunk(hashes);
  return multiPoseidon(hashes);
}

// ─── Compact contribution transcript ─────────────────────────────────────────

/** The structured content of one compact contribution transcript, no padding. */
export interface ContributionTranscript {
  /** `commitments[j][m]`: coefficient `m < t` of pool key `j`'s polynomial, as a point. */
  commitments: TranscriptPoint[][];
  /** `n` committee indexes; slot `i` holds `i + 1`. */
  recipientIndexes: bigint[];
  /** One registry key per committee slot. */
  recipientKeys: TranscriptPoint[];
  /** One ephemeral per committee slot, shared by every key's share to that member. */
  ephemerals: TranscriptPoint[];
  /** `maskedShares[j][i]`: pool key `j`'s masked share for committee slot `i`. */
  maskedShares: bigint[][];
}

/**
 * Parse the `L_C` words of a compact contribution transcript for an epoch
 * with `threshold` and `committeeSize` (taken from the epoch, never from the
 * calldata). Rejects a wrong length, a non-canonical word and a recipient slot
 * whose index is not its committee position. Mirrors `Layout.Decode`.
 */
export function decodeContributionTranscript(
  words: readonly bigint[],
  threshold: number,
  committeeSize: number,
): ContributionTranscript {
  const layout = new ContributionLayout(threshold, committeeSize);
  if (words.length !== layout.words) {
    throw new Error(
      `compact transcript: got ${words.length} words, expected ${layout.words} for t=${threshold} n=${committeeSize}`,
    );
  }
  words.forEach((word, q) => requireCanonicalWord(word, q));
  const point = (q: number): TranscriptPoint => ({ x: words[q], y: words[q + 1] });
  const commitments: TranscriptPoint[][] = [];
  for (let j = 0; j < MAX_K; j++) {
    const row: TranscriptPoint[] = [];
    for (let m = 0; m < threshold; m++) row.push(point(layout.commitmentOffset(j, m)));
    commitments.push(row);
  }
  const recipientIndexes: bigint[] = [];
  const recipientKeys: TranscriptPoint[] = [];
  const ephemerals: TranscriptPoint[] = [];
  for (let i = 0; i < committeeSize; i++) {
    const index = words[layout.recipientIndexOffset(i)];
    if (index !== BigInt(i + 1)) {
      throw new Error(`compact transcript: recipient slot ${i} carries index ${index}, expected ${i + 1}`);
    }
    recipientIndexes.push(index);
    recipientKeys.push(point(layout.recipientKeyOffset(i)));
    ephemerals.push(point(layout.ephemeralOffset(i)));
  }
  const maskedShares: bigint[][] = [];
  for (let j = 0; j < MAX_K; j++) {
    const row: bigint[] = [];
    for (let i = 0; i < committeeSize; i++) row.push(words[layout.maskedShareOffset(j, i)]);
    maskedShares.push(row);
  }
  return { commitments, recipientIndexes, recipientKeys, ephemerals, maskedShares };
}

/** Inverse of `decodeContributionTranscript`: lay a transcript out as `L_C` words. */
export function encodeContributionTranscript(
  transcript: ContributionTranscript,
  threshold: number,
  committeeSize: number,
): bigint[] {
  const layout = new ContributionLayout(threshold, committeeSize);
  const n = committeeSize;
  if (transcript.commitments.length !== MAX_K || transcript.maskedShares.length !== MAX_K) {
    throw new Error(`compact transcript: expected ${MAX_K} commitment and masked-share sets`);
  }
  if (
    transcript.recipientIndexes.length !== n ||
    transcript.recipientKeys.length !== n ||
    transcript.ephemerals.length !== n
  ) {
    throw new Error(`compact transcript: expected ${n} recipient indexes, keys and ephemerals`);
  }
  const words: bigint[] = [];
  for (let j = 0; j < MAX_K; j++) {
    if (transcript.commitments[j].length !== threshold) {
      throw new Error(`compact transcript: key ${j} has ${transcript.commitments[j].length} commitments, expected ${threshold}`);
    }
    for (const c of transcript.commitments[j]) words.push(c.x, c.y);
  }
  words.push(...transcript.recipientIndexes);
  for (const k of transcript.recipientKeys) words.push(k.x, k.y);
  for (const e of transcript.ephemerals) words.push(e.x, e.y);
  for (let j = 0; j < MAX_K; j++) {
    if (transcript.maskedShares[j].length !== n) {
      throw new Error(`compact transcript: key ${j} has ${transcript.maskedShares[j].length} masked shares, expected ${n}`);
    }
    words.push(...transcript.maskedShares[j]);
  }
  if (words.length !== layout.words) throw new Error('compact transcript: internal length mismatch');
  return words;
}

/**
 * The contribution's `commitmentsHash` public input: the two-level Poseidon
 * digest `H(eid, contributorIndex, threshold, keyDigest_0 … keyDigest_15)`
 * with `keyDigest_j = H(A[j][0].x, A[j][0].y, …)` over the commitment vector
 * padded to `MaxN` coefficients with the identity `(0, 1)`. Digests absorb
 * the padded vectors even though the calldata is compact (§4).
 */
export function contributionCommitmentsHash(
  epochId: Hex,
  contributorIndex: number,
  threshold: number,
  commitments: readonly (readonly TranscriptPoint[])[],
): bigint {
  const keyDigests = commitments.map((row) => {
    const inputs: bigint[] = [];
    for (let m = 0; m < MAX_N; m++) {
      const point = row[m] ?? IDENTITY;
      inputs.push(point.x, point.y);
    }
    return multiPoseidon(inputs);
  });
  return multiPoseidon([BigInt(epochId), BigInt(contributorIndex), BigInt(threshold), ...keyDigests]);
}

/**
 * The contribution's `encryptedSharesHash` public input:
 * `H(eid, contributorIndex, committeeSize, rowDigest_0 … rowDigest_31)` with
 * `rowDigest_i = H(index_i, key_i.x, key_i.y, eph_i.x, eph_i.y, ms[0][i] … ms[15][i])`;
 * rows beyond the committee carry index 0, identity points and zero shares.
 */
export function contributionEncryptedSharesHash(
  epochId: Hex,
  contributorIndex: number,
  committeeSize: number,
  transcript: Pick<ContributionTranscript, 'recipientIndexes' | 'recipientKeys' | 'ephemerals' | 'maskedShares'>,
): bigint {
  const rowDigests: bigint[] = [];
  for (let i = 0; i < MAX_N; i++) {
    const key = transcript.recipientKeys[i] ?? IDENTITY;
    const eph = transcript.ephemerals[i] ?? IDENTITY;
    const inputs: bigint[] = [transcript.recipientIndexes[i] ?? 0n, key.x, key.y, eph.x, eph.y];
    for (let j = 0; j < MAX_K; j++) inputs.push(transcript.maskedShares[j]?.[i] ?? 0n);
    rowDigests.push(multiPoseidon(inputs));
  }
  return multiPoseidon([BigInt(epochId), BigInt(contributorIndex), BigInt(committeeSize), ...rowDigests]);
}

/**
 * The BRLC challenge of a contribution: anchor
 * `keccak(commitmentsHash ‖ encryptedSharesHash ‖ keccak(transcript))` under
 * the `davinci-dkg:contribution:v2` domain.
 */
export function contributionChallenge(
  epochId: Hex,
  words: readonly bigint[],
  commitmentsHash: bigint,
  encryptedSharesHash: bigint,
): { anchor: Hex; challenge: bigint; transcriptCommitment: bigint } {
  const anchor = challengeAnchor(words, commitmentsHash, encryptedSharesHash);
  const challenge = deriveChallenge(epochId, DomainContributionTranscriptV2, anchor);
  return { anchor, challenge, transcriptCommitment: brlcCommit(challenge, words) };
}

/** `submitContribution` calldata, decoded and validated against the epoch policy. */
export interface ContributionCall {
  epochId: Hex;
  contributorIndex: number;
  commitmentsHash: Hex;
  encryptedSharesHash: Hex;
  /** The raw `L_C` words. */
  words: bigint[];
  transcript: ContributionTranscript;
  proof: Hex;
  /** The eight public inputs in verifier order. */
  publicInputs: bigint[];
}

/**
 * Decode a `submitContribution` transaction's calldata. `threshold` and
 * `committeeSize` come from the epoch (authoritative state), which is what
 * fixes the layout; the call reverts on chain otherwise, so a mismatch here
 * means the transaction is not what the caller thinks it is.
 */
export function decodeContributionCalldata(
  data: Hex,
  threshold: number,
  committeeSize: number,
): ContributionCall {
  const decoded = decodeFunctionData({ abi: dkgManagerAbi, data });
  if (decoded.functionName !== 'submitContribution') {
    throw new Error(`decodeContributionCalldata: expected submitContribution, got ${decoded.functionName}`);
  }
  const [epochId, contributorIndex, commitmentsHash, encryptedSharesHash, transcriptBytes, proof, input] =
    decoded.args as unknown as [Hex, number, Hex, Hex, Hex, Hex, Hex];
  const words = wordsFromBytes(transcriptBytes);
  const transcript = decodeContributionTranscript(words, threshold, committeeSize);
  const publicInputs = wordsFromBytes(input);
  if (publicInputs.length !== 8) {
    throw new Error(`decodeContributionCalldata: expected 8 public inputs, got ${publicInputs.length}`);
  }
  return {
    epochId,
    contributorIndex: Number(contributorIndex),
    commitmentsHash,
    encryptedSharesHash,
    words,
    transcript,
    proof,
    publicInputs,
  };
}

// ─── Finalization transcript ─────────────────────────────────────────────────

/** The structured content of the fixed finalization transcript, padded to the circuit bounds. */
export interface FinalizeTranscript {
  /** `MaxN` dealer rows: the accepted contributors' indexes, 0 beyond `acceptedCount`. */
  participantIndexes: bigint[];
  /** `MaxN` dealer rows: the stored `commitmentsHash` of each accepted contributor, 0 beyond. */
  contributionHashes: bigint[];
  /** `MaxK` pool keys `P_j` (RTE form, as stored on chain). */
  poolKeys: TranscriptPoint[];
  /**
   * `shareCommitments[j][i]`: `D_{j,i}`, the share commitment of committee
   * member `i + 1` under key `j`; the identity `(0, 1)` for `i ≥ committeeSize`.
   * `MaxK × MaxN`.
   */
  shareCommitments: TranscriptPoint[][];
}

/** Parse the `L_F` words of a finalization transcript. Mirrors the layout of `finalize.PublicInputs`. */
export function decodeFinalizeTranscript(words: readonly bigint[]): FinalizeTranscript {
  if (words.length !== FINALIZE_TRANSCRIPT_WORDS) {
    throw new Error(`finalize transcript: got ${words.length} words, expected ${FINALIZE_TRANSCRIPT_WORDS}`);
  }
  words.forEach((word, q) => requireCanonicalWord(word, q));
  const point = (q: number): TranscriptPoint => ({ x: words[q], y: words[q + 1] });
  const participantIndexes = words.slice(FINALIZE_INDEXES_START, FINALIZE_INDEXES_START + MAX_N);
  const contributionHashes = words.slice(FINALIZE_HASHES_START, FINALIZE_HASHES_START + MAX_N);
  const poolKeys: TranscriptPoint[] = [];
  const shareCommitments: TranscriptPoint[][] = [];
  for (let j = 0; j < MAX_K; j++) {
    poolKeys.push(point(finalizePoolKeyOffset(j)));
    const row: TranscriptPoint[] = [];
    for (let i = 0; i < MAX_N; i++) row.push(point(finalizeShareCommitmentOffset(j, i)));
    shareCommitments.push(row);
  }
  return { participantIndexes, contributionHashes, poolKeys, shareCommitments };
}

/** Inverse of `decodeFinalizeTranscript`. */
export function encodeFinalizeTranscript(transcript: FinalizeTranscript): bigint[] {
  if (transcript.participantIndexes.length !== MAX_N || transcript.contributionHashes.length !== MAX_N) {
    throw new Error(`finalize transcript: expected ${MAX_N} dealer rows`);
  }
  if (transcript.poolKeys.length !== MAX_K || transcript.shareCommitments.length !== MAX_K) {
    throw new Error(`finalize transcript: expected ${MAX_K} keys`);
  }
  const words: bigint[] = [...transcript.participantIndexes, ...transcript.contributionHashes];
  for (let j = 0; j < MAX_K; j++) {
    words.push(transcript.poolKeys[j].x, transcript.poolKeys[j].y);
    if (transcript.shareCommitments[j].length !== MAX_N) {
      throw new Error(`finalize transcript: key ${j} has ${transcript.shareCommitments[j].length} share commitments`);
    }
    for (const d of transcript.shareCommitments[j]) words.push(d.x, d.y);
  }
  if (words.length !== FINALIZE_TRANSCRIPT_WORDS) throw new Error('finalize transcript: internal length mismatch');
  return words;
}

/** The three Poseidon levels of the finalization digest (docs/pool-keys-v4.md §7). */
export interface FinalizeDigestParts {
  /** `R = H(0, I[0..N), h[0..N))`. */
  rows: bigint;
  /** `B_j = H(1, j, P[j].x, P[j].y, D[j][0].x, …)`. */
  keys: bigint[];
  /** `T = H(2, eid, t, n, a, K, L_F, R, B_0 … B_(K−1))` — public input 4. */
  digest: bigint;
}

const DIGEST_TAG_ROWS = 0n;
const DIGEST_TAG_KEY = 1n;
const DIGEST_TAG_OUTER = 2n;

/**
 * Recompute the finalization `transcriptDigest` from the transcript words and
 * the public counts — what the circuit asserts against public input 4 and
 * what the contract checks against the `transcriptDigest` argument.
 */
export function finalizeTranscriptDigest(
  epochId: Hex,
  threshold: number,
  committeeSize: number,
  acceptedCount: number,
  words: readonly bigint[],
): FinalizeDigestParts {
  if (words.length !== FINALIZE_TRANSCRIPT_WORDS) {
    throw new Error(`finalize transcript digest: got ${words.length} words, want ${FINALIZE_TRANSCRIPT_WORDS}`);
  }
  const rows = multiPoseidon([DIGEST_TAG_ROWS, ...words.slice(0, FINALIZE_KEYS_START)]);
  const keys: bigint[] = [];
  for (let j = 0; j < MAX_K; j++) {
    keys.push(
      multiPoseidon([DIGEST_TAG_KEY, BigInt(j), ...words.slice(finalizeKeyOffset(j), finalizeKeyOffset(j) + FINALIZE_KEY_WORDS)]),
    );
  }
  const digest = multiPoseidon([
    DIGEST_TAG_OUTER,
    BigInt(epochId),
    BigInt(threshold),
    BigInt(committeeSize),
    BigInt(acceptedCount),
    BigInt(MAX_K),
    BigInt(FINALIZE_TRANSCRIPT_WORDS),
    rows,
    ...keys,
  ]);
  return { rows, keys, digest };
}

/**
 * The BRLC challenge of a finalization: anchor
 * `keccak(transcriptDigest ‖ keccak(transcript))` under the
 * `davinci-dkg:finalize:v2` domain.
 */
export function finalizeChallenge(
  epochId: Hex,
  words: readonly bigint[],
  transcriptDigest: bigint,
): { anchor: Hex; challenge: bigint; transcriptCommitment: bigint } {
  const anchor = challengeAnchor(words, transcriptDigest);
  const challenge = deriveChallenge(epochId, DomainFinalizeTranscriptV2, anchor);
  return { anchor, challenge, transcriptCommitment: brlcCommit(challenge, words) };
}

/** `finalizeEpoch` calldata, decoded. */
export interface FinalizeCall {
  epochId: Hex;
  /** Public input 4, as passed in the `transcriptDigest` argument. */
  transcriptDigest: bigint;
  /** The raw `L_F` words. */
  words: bigint[];
  transcript: FinalizeTranscript;
  proof: Hex;
  /** The seven public inputs in verifier order. */
  publicInputs: bigint[];
  /** `acceptedCount` (public input 3): rows `[0, acceptedCount)` name the accepted dealers. */
  acceptedCount: number;
}

/**
 * Decode a `finalizeEpoch` transaction's calldata: the epoch's whole key pool
 * and every member's share commitment, straight from the words the contract
 * verified. Rejects anything that is not a canonical, full-length transcript
 * and a public-input vector whose digest position disagrees with the argument.
 */
export function decodeFinalizeCalldata(data: Hex): FinalizeCall {
  const decoded = decodeFunctionData({ abi: dkgManagerAbi, data });
  if (decoded.functionName !== 'finalizeEpoch') {
    throw new Error(`decodeFinalizeCalldata: expected finalizeEpoch, got ${decoded.functionName}`);
  }
  const [epochId, digestHex, transcriptBytes, proof, input] = decoded.args as unknown as [Hex, Hex, Hex, Hex, Hex];
  const words = wordsFromBytes(transcriptBytes);
  const transcript = decodeFinalizeTranscript(words);
  const publicInputs = wordsFromBytes(input);
  if (publicInputs.length !== 7) {
    throw new Error(`decodeFinalizeCalldata: expected 7 public inputs, got ${publicInputs.length}`);
  }
  const transcriptDigest = BigInt(digestHex);
  if (publicInputs[4] !== transcriptDigest) {
    throw new Error('decodeFinalizeCalldata: public input 4 does not match the transcriptDigest argument');
  }
  return {
    epochId,
    transcriptDigest,
    words,
    transcript,
    proof,
    publicInputs,
    acceptedCount: Number(publicInputs[3]),
  };
}
