// Browser-safe SDK entry point.
//
// Excludes the ElGamal crypto module and the high-level flow helpers
// (both depend on circomlibjs which uses Node.js built-ins unavailable in
// browsers).  All on-chain read/write and monitoring utilities are included.

export { DKGClient } from './client.js';
export { DKGWriter } from './writer.js';

export {
  RoundStatus,
  NodeStatus,
  roundStatusLabel,
  type RoundStatusValue,
  type NodeStatusValue,
  type RoundPolicy,
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

export { dkgManagerAbi, dkgRegistryAbi } from './abi.js';
export { buildRoundId, parseRoundId } from './utils.js';
export {
  waitForRoundStatus,
  waitForDecryption,
  watchNewRounds,
  watchRoundFinalized,
  watchDecryptionCombined,
  networkSummary,
} from './monitor.js';
