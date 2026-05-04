// ── Core clients ──────────────────────────────────────────────────────────────
export { DKGClient } from './client.js';
export { DKGWriter } from './writer.js';

// ── Types ─────────────────────────────────────────────────────────────────────
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
  type AppPolicy,
  type ApplicationRecord,
} from './types.js';

// ── ABI ───────────────────────────────────────────────────────────────────────
export { dkgManagerAbi, dkgRegistryAbi } from './abi.js';

// ── Utilities ─────────────────────────────────────────────────────────────────
export { buildEpochId, parseEpochId } from './utils.js';

// ── Monitor / polling ─────────────────────────────────────────────────────────
export {
  waitForEpochPhase,
  waitForDecryption,
  watchNewRounds,
  watchEpochFinalized,
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

// ── Protocol constants + per-application derivation ──────────────────────────
export {
  AppMode,
  Role,
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  DomainDLEQV1,
  DomainOperatorRegisterV1Str,
  DomainOrganizerRegisterV1Str,
  DomainDLEQV1Str,
  type AppModeValue,
  type RoleValue,
} from './protocol.js';

export { computeS, validateDerivePKAppInput, SUBGROUP_ORDER } from './derive.js';
export type { DerivePKAppInput } from './derive.js';

// ── Schnorr / DLEQ verifiers ─────────────────────────────────────────────────
export {
  verifyOperatorSchnorr,
  operatorSchnorrChallenge,
  verifyOrganizerSchnorr,
  organizerSchnorrChallenge,
  proveOrganizer,
  verifyDleq,
  dleqChallenge,
  DOMAIN_PARTIAL_DECRYPT,
  BN254_Q,
  type OperatorSchnorrProof,
  type OrganizerSchnorrProof,
  type DleqPoints,
  type DleqTranscriptInputs,
} from './schnorr.js';
