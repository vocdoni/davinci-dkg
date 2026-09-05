// ── Core clients ──────────────────────────────────────────────────────────────
export { DKGClient } from './client.js';
export { DKGWriter, normalizeAppPolicy, type SubmitCiphertextResult } from './writer.js';

// ── Types ─────────────────────────────────────────────────────────────────────
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
  type PoolKeyActivatedEvent,
  type PoolKeyClaimedEvent,
  type OrganizerSecretRevealedEvent,
} from './types.js';

// ── ABI ───────────────────────────────────────────────────────────────────────
export { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';

// ── Utilities ─────────────────────────────────────────────────────────────────
export { buildEpochId, parseEpochId } from './utils.js';

// ── Monitor / polling ─────────────────────────────────────────────────────────
export {
  waitForEpochPhase,
  waitForDecryption,
  waitForPoolKeyActivated,
  decryptionProgress,
  watchNewRounds,
  watchNewEpochs,
  watchEpochLive,
  watchDecryptionCombined,
  watchCiphertextSubmitted,
  networkSummary,
} from './monitor.js';

// ── High-level flow helpers ───────────────────────────────────────────────────
export {
  encrypt,
  encryptForApplication,
  decrypt,
  waitForPoolKey,
  waitForCombinedDecryption,
  demonstrateEncryptDecryptFlow,
} from './flow.js';

// ── Crypto ────────────────────────────────────────────────────────────────────
export { buildElGamal, applicationKey, randomOrganizerSecret, randomAid } from './crypto/index.js';
export type { ElGamal } from './crypto/index.js';
export {
  fromRTEtoTE,
  fromTEtoRTE,
  pointFromRTEtoTE,
  pointFromTEtoRTE,
} from './crypto/babyjub-form.js';

// ── Protocol constants ────────────────────────────────────────────────────────
export {
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  DomainOperatorRegisterV1Str,
  DomainOrganizerRegisterV1Str,
  DomainContributionTranscriptV1,
  DomainPoolKeyTranscriptV1,
  DomainDecryptCombineTranscriptV1,
  DomainContributionTranscriptV1Str,
  DomainPoolKeyTranscriptV1Str,
  DomainDecryptCombineTranscriptV1Str,
} from './protocol.js';

// ── Schnorr / DLEQ provers and verifiers ─────────────────────────────────────
export {
  verifyOperatorSchnorr,
  operatorSchnorrChallenge,
  proveOperator,
  verifyOrganizerSchnorr,
  organizerSchnorrChallenge,
  proveOrganizer,
  organizerPublicKey,
  verifyDleq,
  dleqChallenge,
  DOMAIN_PARTIAL_DECRYPT,
  BN254_Q,
  SUBGROUP_ORDER,
  type OperatorSchnorrProof,
  type OrganizerSchnorrProof,
  type DleqPoints,
  type DleqTranscriptInputs,
} from './schnorr.js';
