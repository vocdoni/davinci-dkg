import type { Address, Hex, PublicClient, WalletClient } from 'viem';

export type { Address, Hex } from 'viem';

// ── Round status ──────────────────────────────────────────────────────────────

export const RoundStatus = {
  None: 0,
  Registration: 1,
  Contribution: 2,
  Finalized: 3,
  Aborted: 4,
  Completed: 5,
} as const;

export type RoundStatusValue = (typeof RoundStatus)[keyof typeof RoundStatus];

export function roundStatusLabel(status: number): string {
  switch (status) {
    case RoundStatus.None: return 'None';
    case RoundStatus.Registration: return 'Registration';
    case RoundStatus.Contribution: return 'Contribution';
    case RoundStatus.Finalized: return 'Finalized';
    case RoundStatus.Aborted: return 'Aborted';
    case RoundStatus.Completed: return 'Completed';
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

export interface RoundPolicy {
  threshold: number;
  committeeSize: number;
  minValidContributions: number;
  /** Over-subscription factor in basis points (min 10000 = 1.0×). Default 15000 = 1.5×. */
  lotteryAlphaBps: number;
  /** Number of blocks between createRound and seed availability (1–256). */
  seedDelay: number;
  registrationDeadlineBlock: bigint;
  contributionDeadlineBlock: bigint;
  /**
   * Earliest block at which `finalizeRound` can succeed. Must be strictly
   * greater than `contributionDeadlineBlock`; gives selected participants a
   * window to submit before the contribution set is frozen. The contract
   * reverts with `FinalizeTooEarly` if `block.number < finalizeNotBeforeBlock`.
   */
  finalizeNotBeforeBlock: bigint;
  disclosureAllowed: boolean;
}

/**
 * Gates `submitCiphertext` for a round. All checks AND together; a zero-valued
 * field is a no-op for that check. Policy gates SUBMISSION only — once a
 * ciphertext is on-chain, committee decryption proceeds regardless of these
 * fields.
 */
export interface DecryptionPolicy {
  /** If true, only the round organizer may call `submitCiphertext`. */
  ownerOnly: boolean;
  /** Maximum accepted ciphertexts per round; 0 = unlimited (bounded by MAX_CIPHERTEXT_INDEX). */
  maxDecryptions: number;
  /** submitCiphertext reverts if `block.number < notBeforeBlock`; 0 = no lock. */
  notBeforeBlock: bigint;
  /** submitCiphertext reverts if `block.timestamp < notBeforeTimestamp`; 0 = no lock. */
  notBeforeTimestamp: bigint;
  /** submitCiphertext reverts if `block.number > notAfterBlock`; 0 = no deadline. */
  notAfterBlock: bigint;
  /** submitCiphertext reverts if `block.timestamp > notAfterTimestamp`; 0 = no deadline. */
  notAfterTimestamp: bigint;
}

/** Convenience: an all-zero DecryptionPolicy = no submission gating. */
export const OpenDecryptionPolicy: DecryptionPolicy = {
  ownerOnly: false,
  maxDecryptions: 0,
  notBeforeBlock: 0n,
  notBeforeTimestamp: 0n,
  notAfterBlock: 0n,
  notAfterTimestamp: 0n,
};

export interface Round {
  organizer: Address;
  policy: RoundPolicy;
  decryptionPolicy: DecryptionPolicy;
  status: RoundStatusValue;
  nonce: bigint;
  seedBlock: bigint;
  seed: Hex;
  lotteryThreshold: bigint;
  claimedCount: number;
  contributionCount: number;
  partialDecryptionCount: number;
  revealedShareCount: number;
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
  participant: Address;
  participantIndex: number;
  ciphertextIndex: number;
  deltaHash: Hex;
  delta: { x: bigint; y: bigint };
  accepted: boolean;
}

export interface CombinedDecryptionRecord {
  ciphertextIndex: number;
  completed: boolean;
  /** Recovered plaintext scalar; zero if `completed` is false. */
  plaintext: bigint;
}

export interface RevealedShareRecord {
  participant: Address;
  participantIndex: number;
  shareValue: bigint;
  shareHash: Hex;
  accepted: boolean;
}

export interface NodeKey {
  operator: Address;
  pubX: bigint;
  pubY: bigint;
  status: NodeStatusValue;
  lastActiveBlock: bigint;
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

// ── Monitor types ─────────────────────────────────────────────────────────────

export interface PollOptions {
  /** Interval between polls in ms (default: 2000) */
  intervalMs?: number;
  /** Maximum wait time in ms. Throws if exceeded. (default: 120_000) */
  timeoutMs?: number;
}

// ── Event query types ─────────────────────────────────────────────────────────

/** A single parsed contract event returned by getAllRoundEvents / getRoundCreatedEvents. */
export interface RoundEvent {
  eventName: string;
  args: Record<string, unknown>;
  blockNumber: bigint;
  transactionHash: `0x${string}`;
}

/** A round entry returned by getRecentRounds. */
export interface RoundEntry {
  id: `0x${string}`;
  round: Round;
}
