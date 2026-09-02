// Browser-safe SDK entry point.
//
// Excludes the ElGamal crypto module and the high-level flow helpers; all
// on-chain read/write, monitoring and proof utilities are included. Note
// that `DKGWriter.submitCiphertext` needs a `CiphertextPoK`: build it with
// `proveCiphertext` from the ciphertext's randomness.

export { DKGClient } from './client.js';
export { DKGWriter, type SubmitCiphertextResult } from './writer.js';

export {
  EpochPhase,
  NodeStatus,
  roundStatusLabel,
  type EpochPhaseValue,
  type NodeStatusValue,
  type EpochPolicy,
  type CreateEpochParams,
  type EpochBounds,
  type Epoch,
  type ContributionRecord,
  type PartialDecryptionRecord,
  type CombinedDecryptionRecord,
  type NodeKey,
  type DKGConfig,
  type DKGWriterConfig,
  type BabyJubPoint,
  type ElGamalCiphertext,
  type CiphertextPoK,
  type PollOptions,
  type EpochEvent,
  type EpochEntry,
  type AppPolicy,
  type ApplicationRecord,
} from './types.js';

export { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';
export { proveCiphertext, verifyCiphertextPoK, ciphertextPoKChallenge } from './schnorr.js';
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
