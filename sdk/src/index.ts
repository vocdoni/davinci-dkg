// ── Core clients ──────────────────────────────────────────────────────────────
export { DKGClient } from './client.js';
export { DKGWriter } from './writer.js';

// ── Types ─────────────────────────────────────────────────────────────────────
export {
  RoundStatus,
  NodeStatus,
  OpenDecryptionPolicy,
  roundStatusLabel,
  type RoundStatusValue,
  type NodeStatusValue,
  type RoundPolicy,
  type DecryptionPolicy,
  type Round,
  type ContributionRecord,
  type PartialDecryptionRecord,
  type CombinedDecryptionRecord,
  type RevealedShareRecord,
  type NodeKey,
  type DKGConfig,
  type DKGWriterConfig,
  type BabyJubPoint,
  type ElGamalCiphertext,
  type PollOptions,
  type RoundEvent,
  type RoundEntry,
} from './types.js';

// ── ABI ───────────────────────────────────────────────────────────────────────
export { dkgManagerAbi, dkgRegistryAbi } from './abi.js';

// ── Utilities ─────────────────────────────────────────────────────────────────
export { buildRoundId, parseRoundId } from './utils.js';

// ── Monitor / polling ─────────────────────────────────────────────────────────
export {
  waitForRoundStatus,
  waitForDecryption,
  watchNewRounds,
  watchRoundFinalized,
  watchDecryptionCombined,
  watchCiphertextSubmitted,
  networkSummary,
} from './monitor.js';

// ── High-level flow helpers ───────────────────────────────────────────────────
export {
  encrypt,
  decrypt,
  waitForCollectivePublicKeyHash,
  waitForCombinedDecryption,
  demonstrateEncryptDecryptFlow,
  type CollectivePublicKey,
} from './flow.js';

// ── Crypto ────────────────────────────────────────────────────────────────────
export { buildElGamal } from './crypto/index.js';
export type { ElGamal } from './crypto/index.js';
