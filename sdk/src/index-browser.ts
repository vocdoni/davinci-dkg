// Browser-safe SDK entry point.
//
// Excludes the ElGamal crypto module and the high-level flow helpers
// (both depend on circomlibjs which uses Node.js built-ins unavailable in
// browsers).  All on-chain read/write and monitoring utilities are included.

export { DKGClient } from './client.js';
export { DKGWriter } from './writer.js';

export {
  EpochPhase,
  NodeStatus,
  OpenDecryptionPolicy,
  roundStatusLabel,
  type EpochPhaseValue,
  type NodeStatusValue,
  type EpochPolicy,
  type DecryptionPolicy,
  type Epoch,
  type ContributionRecord,
  type PartialDecryptionRecord,
  type CombinedDecryptionRecord,
  type NodeKey,
  type DKGConfig,
  type DKGWriterConfig,
  type BabyJubPoint,
  type ElGamalCiphertext,
  type PollOptions,
  type EpochEvent,
  type EpochEntry,
} from './types.js';

export { dkgManagerAbi, dkgRegistryAbi } from './abi.js';
export { buildEpochId, parseEpochId } from './utils.js';
export {
  waitForEpochPhase,
  waitForDecryption,
  watchNewRounds,
  watchEpochLive,
  watchDecryptionCombined,
  watchCiphertextSubmitted,
  networkSummary,
} from './monitor.js';
