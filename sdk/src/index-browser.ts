// Browser-safe SDK entry point.
//
// Excludes the ElGamal crypto module and the high-level flow helpers; all
// on-chain read/write, monitoring and proof utilities are included — the
// organizer's registration Schnorr prover (`proveOrganizer`) among them,
// since the organizer is expected to run in a browser.

export { DKGClient, type FinalizeRecord, type ShareProofRecord } from './client.js';
export { DKGWriter, normalizeAppPolicy, type SubmitCiphertextResult } from './writer.js';

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
  type PollOptions,
  type EpochEvent,
  type EpochEntry,
  AppMode,
  appModeLabel,
  type AppModeValue,
  type AppPolicy,
  type AppPolicyInput,
  type ApplicationRecord,
  type ActivityScanOptions,
  type SlotClaimedEvent,
  type ContributionSubmittedEvent,
  type PartialDecryptionEvent,
  type EpochLiveEvent,
  type DecryptionCombinedEvent,
  type ApplicationRegisteredEvent,
  type PoolStatus,
  type PoolKeyClaimedEvent,
  type OrganizerSecretRevealedEvent,
} from './types.js';

export { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';
export {
  proveOrganizer,
  organizerPublicKey,
  verifyOrganizerSchnorr,
  organizerSchnorrChallenge,
} from './schnorr.js';
export { DomainOrganizerRegisterV1 } from './protocol.js';
export { buildEpochId, parseEpochId } from './utils.js';
export {
  MAX_N,
  MAX_K,
  MERKLE_DEPTH,
  contributionTranscriptWords,
  ContributionLayout,
  FINALIZE_KEY_WORDS,
  FINALIZE_TRANSCRIPT_WORDS,
  finalizeKeyOffset,
  finalizePoolKeyOffset,
  finalizeShareCommitmentOffset,
} from './sizes.js';
export {
  wordsFromBytes,
  decodeContributionTranscript,
  decodeContributionCalldata,
  decodeFinalizeTranscript,
  decodeFinalizeCalldata,
  type TranscriptPoint,
  type ContributionTranscript,
  type FinalizeTranscript,
  type FinalizeCall,
} from './transcript.js';
export {
  MERKLE_EMPTY_LEAF,
  shareCommitmentLeaf,
  merkleNode,
  shareCommitmentLeaves,
  merkleRoot,
  merklePath,
  merkleRootFromPath,
  verifyMerklePath,
  shareProof,
} from './merkle.js';
export {
  waitForEpochPhase,
  waitForDecryption,
  decryptionProgress,
  watchNewRounds,
  watchNewEpochs,
  watchEpochLive,
  watchDecryptionCombined,
  watchCiphertextSubmitted,
  networkSummary,
} from './monitor.js';
