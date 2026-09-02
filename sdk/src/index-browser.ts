// Browser-safe SDK entry point.
//
// Excludes the ElGamal crypto module and the high-level flow helpers; all
// on-chain read/write, monitoring and proof utilities are included — the
// organizer's share prover (`proveOrganizerShare`) among them, since the
// organizer is expected to run in a browser.

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
  type PollOptions,
  type EpochEvent,
  type EpochEntry,
  type AppPolicy,
  type ApplicationRecord,
  type ActivityScanOptions,
  type SlotClaimedEvent,
  type ContributionSubmittedEvent,
  type PartialDecryptionEvent,
  type EpochLiveEvent,
  type DecryptionCombinedEvent,
  type ApplicationRegisteredEvent,
} from './types.js';

export { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';
export { proveOrganizer, verifyOrganizerSchnorr, organizerSchnorrChallenge } from './schnorr.js';
export {
  organizerShareChallenge,
  proveOrganizerShare,
  verifyOrganizerShare,
  type OrganizerShare,
  type OrganizerShareProof,
} from './dleq.js';
export { DomainOrganizerRegisterV1, DomainOrganizerShareV1 } from './protocol.js';
export { buildEpochId, parseEpochId } from './utils.js';
export {
  waitForEpochPhase,
  waitForDecryption,
  waitForOrganizerShare,
  decryptionProgress,
  watchNewRounds,
  watchNewEpochs,
  watchEpochLive,
  watchDecryptionCombined,
  watchCiphertextSubmitted,
  networkSummary,
} from './monitor.js';
