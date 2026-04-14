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
  disclosureAllowed: boolean;
}

export interface Round {
  organizer: Address;
  policy: RoundPolicy;
  status: RoundStatusValue;
  nonce: bigint;
  seedBlock: bigint;
  seed: Hex;
  lotteryThreshold: bigint;
  claimedCount: number;
  contributionCount: number;
  partialDecryptionCount: number;
  revealedShareCount: number;
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
  combineHash: Hex;
  plaintextHash: Hex;
  completed: boolean;
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
  /** Address of the deployed DKGRegistry contract */
  registryAddress: Address;
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
