// ── Core clients ──────────────────────────────────────────────────────────────
export { DKGClient } from './client.js';
export { DKGWriter, type SubmitCiphertextResult } from './writer.js';

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
  type AppPolicy,
  type ApplicationRecord,
} from './types.js';

// ── ABI ───────────────────────────────────────────────────────────────────────
export { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';

// ── Utilities ─────────────────────────────────────────────────────────────────
export { buildEpochId, parseEpochId } from './utils.js';

// ── Monitor / polling ─────────────────────────────────────────────────────────
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

// ── High-level flow helpers ───────────────────────────────────────────────────
export {
  encrypt,
  encryptForApplication,
  decrypt,
  waitForCollectivePublicKeyHash,
  waitForCombinedDecryption,
  demonstrateEncryptDecryptFlow,
  type CollectivePublicKey,
} from './flow.js';

// ── Crypto ────────────────────────────────────────────────────────────────────
export { buildElGamal, applicationKey, randomOrganizerSecret } from './crypto/index.js';
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
  DomainDLEQV1,
  DomainOrganizerShareV1,
  DomainOperatorRegisterV1Str,
  DomainOrganizerRegisterV1Str,
  DomainDLEQV1Str,
  DomainOrganizerShareV1Str,
} from './protocol.js';

// ── Schnorr / DLEQ provers and verifiers ─────────────────────────────────────
export {
  verifyOperatorSchnorr,
  operatorSchnorrChallenge,
  proveOperator,
  verifyOrganizerSchnorr,
  organizerSchnorrChallenge,
  proveOrganizer,
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

// ── Organizer decryption share ───────────────────────────────────────────────
export {
  organizerShareChallenge,
  proveOrganizerShare,
  verifyOrganizerShare,
  type OrganizerShare,
  type OrganizerShareProof,
} from './dleq.js';
