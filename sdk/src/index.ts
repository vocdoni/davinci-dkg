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
  type CiphertextPoK,
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
  watchNewRounds,
  watchEpochLive,
  watchDecryptionCombined,
  watchCiphertextSubmitted,
  networkSummary,
} from './monitor.js';

// ── High-level flow helpers ───────────────────────────────────────────────────
export {
  encrypt,
  encryptWithProof,
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
  DomainCiphertextPoKV1,
  DomainOperatorRegisterV1Str,
  DomainOrganizerRegisterV1Str,
  DomainDLEQV1Str,
  DomainCiphertextPoKV1Str,
  type AppModeValue,
  type RoleValue,
} from './protocol.js';

export { computeS, validateDerivePKAppInput, SUBGROUP_ORDER } from './derive.js';
export type { DerivePKAppInput } from './derive.js';

// ── Schnorr / DLEQ provers and verifiers ─────────────────────────────────────
export {
  verifyOperatorSchnorr,
  operatorSchnorrChallenge,
  proveOperator,
  verifyOrganizerSchnorr,
  organizerSchnorrChallenge,
  proveOrganizer,
  proveCiphertext,
  verifyCiphertextPoK,
  ciphertextPoKChallenge,
  verifyDleq,
  dleqChallenge,
  DOMAIN_PARTIAL_DECRYPT,
  BN254_Q,
  type OperatorSchnorrProof,
  type OrganizerSchnorrProof,
  type DleqPoints,
  type DleqTranscriptInputs,
} from './schnorr.js';
