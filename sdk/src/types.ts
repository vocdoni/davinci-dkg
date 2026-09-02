import type { Address, Hex, PublicClient, WalletClient } from 'viem';

export type { Address, Hex } from 'viem';

// ── Epoch status ──────────────────────────────────────────────────────────────

/**
 * Epoch state machine. The first three states group into the "Preparation"
 * phase (committee assembly + key generation); `Live` is the "Service" phase
 * in which apps register their derived keys and ciphertexts get decrypted.
 */
export const EpochPhase = {
  None: 0,
  CommitteeSelection: 1,
  KeyAssembly: 2,
  Live: 3,
  Aborted: 4,
  Completed: 5,
} as const;

export type EpochPhaseValue = (typeof EpochPhase)[keyof typeof EpochPhase];

export function roundStatusLabel(status: number): string {
  switch (status) {
    case EpochPhase.None: return 'None';
    case EpochPhase.CommitteeSelection: return 'CommitteeSelection';
    case EpochPhase.KeyAssembly: return 'KeyAssembly';
    case EpochPhase.Live: return 'Live';
    case EpochPhase.Aborted: return 'Aborted';
    case EpochPhase.Completed: return 'Completed';
    default: return `Unknown(${status})`;
  }
}

// ── Node status ───────────────────────────────────────────────────────────────

export const NodeStatus = {
  None: 0,
  Active: 1,
  Inactive: 2,
} as const;

export type NodeStatusValue = (typeof NodeStatus)[keyof typeof NodeStatus];

// ── Contract types ────────────────────────────────────────────────────────────

/**
 * Per-epoch DKG policy. Phase deadline blocks are derived ON-CHAIN at
 * `createEpoch` time from the contract's immutable `EPOCH_DURATION_BLOCKS`
 * plus the per-phase BPS constants (registration / contribution /
 * finalize gap). Callers no longer supply them: `writer.createEpoch` only
 * takes the policy fields below. The on-chain Epoch struct continues to
 * surface the resolved deadline blocks (populated by createEpoch from the
 * derived offsets) for downstream phase-check reads.
 */
export interface EpochPolicy {
  threshold: number;
  committeeSize: number;
  minValidContributions: number;
  /** Over-subscription factor in basis points (min 10000 = 1.0×). Default 15000 = 1.5×. */
  lotteryAlphaBps: number;
  /**
   * On-chain-derived deadline blocks. Populated by the contract from
   * `EPOCH_DURATION_BLOCKS`, surfaced via `getEpoch` for phase checks.
   * Callers constructing a fresh policy for `createEpoch` may leave them
   * at 0 — they are not transmitted by the writer.
   */
  committeeSelectionDeadlineBlock: bigint;
  keyAssemblyDeadlineBlock: bigint;
  liveNotBeforeBlock: bigint;
}

/**
 * The caller-supplied part of an epoch policy — exactly the four arguments
 * of `DKGManager.createEpoch`. Phase deadlines are derived on-chain, and the
 * former per-epoch `DecryptionPolicy` no longer exists: ciphertext submission
 * is gated per application (`AppPolicy`) instead.
 */
export type CreateEpochParams = Pick<
  EpochPolicy,
  'threshold' | 'committeeSize' | 'minValidContributions' | 'lotteryAlphaBps'
>;

/**
 * Deploy-time bounds `createEpoch` enforces on `CreateEpochParams`
 * (`DKGManager` immutables `MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`,
 * `MAX_LOTTERY_ALPHA_BPS`). On top of these the contract always requires
 * `1 ≤ threshold ≤ minValidContributions ≤ committeeSize ≤ MaxN` and
 * `lotteryAlphaBps ≥ 10000`. Read them with `client.getEpochBounds()`.
 */
export interface EpochBounds {
  minThreshold: number;
  minCommitteeSize: number;
  maxLotteryAlphaBps: number;
}

export interface Epoch {
  organizer: Address;
  policy: EpochPolicy;
  status: EpochPhaseValue;
  nonce: bigint;
  /** Block in which this epoch was created (anchor for nextEpochStartBlock). */
  startBlock: bigint;
  seedBlock: bigint;
  seed: Hex;
  lotteryThreshold: bigint;
  claimedCount: number;
  contributionCount: number;
  partialDecryptionCount: number;
  ciphertextCount: number;
}

export interface ContributionRecord {
  contributor: Address;
  contributorIndex: number;
  commitmentsHash: Hex;
  encryptedSharesHash: Hex;
  commitmentVectorDigest: Hex;
  accepted: boolean;
}

export interface PartialDecryptionRecord {
  participantIndex: number;
  ciphertextIndex: number;
  deltaHash: Hex;
  accepted: boolean;
}

export interface CombinedDecryptionRecord {
  ciphertextIndex: number;
  completed: boolean;
  /** Recovered plaintext scalar; zero if `completed` is false. */
  plaintext: bigint;
}

export interface NodeKey {
  operator: Address;
  pubX: bigint;
  pubY: bigint;
  status: NodeStatusValue;
  lastActiveBlock: bigint;
  /**
   * Block of the first `registerKey`. A node only enters the lottery of
   * epochs created after this block (`NotInSnapshot` otherwise), so a fresh
   * identity cannot be ground against an already-revealed seed.
   */
  registeredAtBlock: bigint;
}

// ── Application (P8/P9) ──────────────────────────────────────────────────────

/** Per-application policy gating submitCiphertext, mirrors `DKGTypes.AppPolicy`. */
export interface AppPolicy {
  /** Address authorized to submit ciphertexts; zero address means open. */
  authorizedSubmitter: Address;
  /** Maximum ciphertexts under this aid; 0 means unlimited. */
  maxCiphertexts: number;
  /** Earliest block at which submitCiphertext is valid; 0 means no floor. */
  notBeforeBlock: bigint;
  /** Latest block at which submitCiphertext is valid; 0 means no ceiling. */
  notAfterBlock: bigint;
}

/**
 * Cached on-chain `Application` record. `mode === 0` means public derivation
 * (PK_aid = PK_ep + S·G); `mode === 1` means organizer co-decryption
 * (PK_aid = PK_ep + PK_org). `derivationS` is meaningful only in mode 0;
 * `organizerPK` is meaningful only in mode 1.
 */
export interface ApplicationRecord {
  creator: Address;
  mode: 0 | 1;
  derivationS: bigint;
  organizerPK: BabyJubPoint;
  policy: AppPolicy;
  createdAtBlock: bigint;
  exists: boolean;
}

// ── SDK config ────────────────────────────────────────────────────────────────

export interface DKGConfig {
  /** viem PublicClient connected to the target chain */
  publicClient: PublicClient;
  /** Address of the deployed DKGManager contract */
  managerAddress: Address;
  /**
   * Address of the deployed DKGRegistry contract.
   * When omitted, the client reads it from DKGManager.REGISTRY() on first use.
   */
  registryAddress?: Address;
  /**
   * Address of the deployed DKGAppManager contract (sibling to DKGManager
   * that owns the per-application registration surface).
   * When omitted, the client reads it from DKGManager.appManager() on first use.
   */
  appManagerAddress?: Address;
}

export interface DKGWriterConfig extends DKGConfig {
  /** viem WalletClient for signing transactions */
  walletClient: WalletClient;
}

// ── ElGamal types ─────────────────────────────────────────────────────────────

/** A BabyJubJub curve point as [x, y] bigints. */
export type BabyJubPoint = [bigint, bigint];

export interface ElGamalCiphertext {
  /** Ephemeral key: c1 = k * G */
  c1: BabyJubPoint;
  /** Encrypted message: c2 = m*G + k*PubKey */
  c2: BabyJubPoint;
}

/**
 * Schnorr proof of knowledge of the ElGamal randomness `r` behind
 * `c1 = r·G`, bound to `(epochId, aid, c1, c2)`. Every committee node
 * verifies it before releasing a partial decryption, so a ciphertext
 * submitted without a valid proof is never decrypted. Build it with
 * `encryptWithProof` (or `proveCiphertext` if you already hold `r`).
 * Coordinates are in on-chain (RTE) form, exactly as `submitCiphertext`
 * sends them and `CiphertextSubmitted` reports them.
 */
export interface CiphertextPoK {
  /** Witness point A = w·G. */
  ax: bigint;
  ay: bigint;
  /** Response z = w + c·r mod L. */
  z: bigint;
}

// ── Monitor types ─────────────────────────────────────────────────────────────

export interface PollOptions {
  /** Interval between polls in ms (default: 2000) */
  intervalMs?: number;
  /** Maximum wait time in ms. Throws if exceeded. (default: 120_000) */
  timeoutMs?: number;
}

// ── Event query types ─────────────────────────────────────────────────────────

/** A single parsed contract event returned by getAllEpochEvents / getEpochCreatedEvents. */
export interface EpochEvent {
  eventName: string;
  args: Record<string, unknown>;
  blockNumber: bigint;
  transactionHash: `0x${string}`;
}

/** A epoch entry returned by getRecentRounds. */
export interface EpochEntry {
  id: `0x${string}`;
  epoch: Epoch;
}
