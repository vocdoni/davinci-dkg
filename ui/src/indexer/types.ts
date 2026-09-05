// Entity store for the explorer's in-browser indexer.
//
// Everything a page renders comes from this store: the indexer fills it from
// contract events plus the state those events do not carry, the synthetic
// fixture builds the identical shape without a chain, and `selectors.ts` is
// the only thing that reads it. Blocks are plain numbers (an Ethereum block
// height fits in a double with room to spare); curve coordinates, plaintexts
// and the lottery threshold stay `bigint` because they are field elements.
//
// The store is persisted verbatim (see `persist.ts`), so keep it free of
// class instances, functions and cycles.

import type { Address, Hex } from 'viem'

export type { Address, Hex }

/** Bumped whenever the shape below changes; a mismatch drops the cache. */
export const STORE_VERSION = 5

/** `bytes12` epoch id (4-byte prefix ‖ 8-byte nonce). */
export type EpochId = Hex
/** `bytes32` application id. */
export type Aid = Hex

export interface Point {
  x: bigint
  y: bigint
}

export type EpochPhaseName =
  | 'none'
  | 'committee-selection'
  | 'key-assembly'
  | 'live'
  | 'aborted'
  | 'completed'

export type NodeStatusName = 'none' | 'active' | 'inactive'

/** `MaxK`: pool keys dealt per epoch (`circuits/common/sizes.go`, `Sizes.sol`). */
export const POOL_SIZE = 8
/** `MaxN`: the committee cap the circuits are compiled for. */
export const MAX_COMMITTEE = 32

/**
 * Whether an organizer key sits on top of the pool key (`DKGTypes.AppMode`).
 * Organizer-locked: `PK_aid = P_j + PK_org`, and the organizer reveals
 * `sk_org` once before the committee can combine. Automatic: `PK_aid = P_j`,
 * no organizer key at all (`organizerPK` is the identity `(0, 1)`).
 */
export type AppModeName = 'organizer-locked' | 'automatic'

export function appModeName(mode: number | bigint): AppModeName {
  return Number(mode) === 1 ? 'automatic' : 'organizer-locked'
}

// ── Normalised events ────────────────────────────────────────────────────────

/** Payload per event, keyed by the contract's event name. */
export interface EventDataMap {
  // DKGRegistry
  NodeRegistered: { operator: Address; pubX: bigint; pubY: bigint }
  NodeUpdated: { operator: Address; pubX: bigint; pubY: bigint }
  NodeMarkedActive: { operator: Address; atBlock: number }
  NodeReaped: { operator: Address; lastActiveBlock: number }
  NodeReactivated: { operator: Address }
  ManagerSet: { manager: Address }
  // DKGManager
  EpochCreated: {
    epochId: EpochId
    organizer: Address
    startBlock: number
    seedBlock: number
    lotteryThreshold: bigint
  }
  SeedResolved: { epochId: EpochId; seed: Hex }
  SlotClaimed: { epochId: EpochId; claimer: Address; slot: number }
  CommitteeFilled: { epochId: EpochId }
  ContributionSubmitted: {
    epochId: EpochId
    contributor: Address
    contributorIndex: number
    commitmentsHash: Hex
    encryptedSharesHash: Hex
  }
  /** Proof-less finalize: freezes the accepted contributor set, no key yet. */
  EpochLive: { epochId: EpochId; contributionCount: number }
  /** Pool key `P_j` proven and stored; `key` is in TE form. */
  PoolKeyActivated: { epochId: EpochId; keyIndex: number; key: Point }
  PoolKeyClaimed: { epochId: EpochId; aid: Aid; keyIndex: number }
  CiphertextSubmitted: {
    epochId: EpochId
    aid: Aid
    ciphertextIndex: number
    submitter: Address
    c1: Point
    c2: Point
  }
  PartialDecryptionSubmitted: {
    epochId: EpochId
    aid: Aid
    participant: Address
    participantIndex: number
    ciphertextIndex: number
    delta: Point
  }
  DecryptionCombined: {
    epochId: EpochId
    aid: Aid
    ciphertextIndex: number
    combineHash: Hex
    plaintext: bigint
  }
  EpochAborted: { epochId: EpochId }
  // DKGAppManager
  ApplicationRegistered: {
    epochId: EpochId
    aid: Aid
    creator: Address
    /** TE form, matching `getApplication`; the identity `(0, 1)` when automatic. */
    organizerPK: Point
    mode: AppModeName
    /** Pool key the registration claimed. */
    poolIndex: number
  }
  OrganizerSecretRevealed: { epochId: EpochId; aid: Aid; organizerSecret: bigint }
}

export type IndexedEventName = keyof EventDataMap

/** Common envelope every normalised event carries. */
interface EventEnvelope {
  block: number
  tx: Hex | null
  logIndex: number
  /** Epoch this event belongs to, when it names one. */
  epoch: EpochId | null
  /** Application this event belongs to, when it names one. */
  aid: Aid | null
  /** The operator/organizer the event attributes itself to, when it names one. */
  actor: Address | null
}

export type IndexedEvent = {
  [K in IndexedEventName]: EventEnvelope & { name: K; data: EventDataMap[K] }
}[IndexedEventName]

// ── Entities ─────────────────────────────────────────────────────────────────

export interface OperatorEntity {
  address: Address
  /** BabyJubJub encryption key, latest value after any `updateKey`. */
  pubKey: Point | null
  status: NodeStatusName
  /** Block at which the operator (re-)entered the active set. */
  registeredAtBlock: number
  lastActiveBlock: number
  firstSeenBlock: number
  keyUpdates: number
  reaps: number
  reactivations: number
  /** Indices into `IndexerStore.events`, ascending. */
  events: number[]
  /** Block at which the registry record was last read; 0 = never. */
  stateBlock: number
}

export interface EpochPolicy {
  threshold: number
  committeeSize: number
  minValidContributions: number
  lotteryAlphaBps: number
  committeeSelectionDeadlineBlock: number
  keyAssemblyDeadlineBlock: number
  liveNotBeforeBlock: number
}

export interface SlotEntity {
  key: string
  epoch: EpochId
  slot: number
  operator: Address
  block: number
  tx: Hex | null
}

export interface ContributionEntity {
  key: string
  epoch: EpochId
  index: number
  contributor: Address
  commitmentsHash: Hex
  encryptedSharesHash: Hex
  block: number
  tx: Hex | null
}

export interface FinalizationEntity {
  /** Resolved from the transaction sender (lazy); `EpochLive` names nobody. */
  by: Address | null
  block: number
  tx: Hex | null
  /** Accepted contributions frozen by `finalizeEpoch`. */
  contributionCount: number
}

/**
 * One of the `POOL_SIZE` keys an epoch deals. A slot exists from the moment
 * the epoch does; `key` is null until `activatePoolKey` proves it, and
 * `claimedBy` until a registration claims it.
 */
export interface PoolKeyEntity {
  index: number
  /** `P_j`, TE form; null until activated. */
  key: Point | null
  activatedBlock: number | null
  activatedTx: Hex | null
  claimedBy: Aid | null
  claimedBlock: number | null
  claimedTx: Hex | null
}

export interface EpochCounts {
  claims: number
  contributions: number
  ciphertexts: number
  partials: number
  combines: number
  applications: number
}

export interface EpochEntity {
  id: EpochId
  nonce: number
  creator: Address
  createdBlock: number
  createdTx: Hex | null
  /** `block.number` at `createEpoch` — the cadence anchor. */
  startBlock: number
  seedBlock: number
  seed: Hex | null
  seedResolvedBlock: number | null
  /** τ: `keccak(seed ‖ operator) < τ` makes an operator eligible. */
  lotteryThreshold: bigint
  status: EpochPhaseName
  policy: EpochPolicy | null
  /** Registry `activeCount()` snapshotted at `createEpoch`, when known. */
  registrySnapshot: number | null
  /** Committee in slot order; index i is participant index i. */
  committee: Address[]
  committeeFilledBlock: number | null
  abortedBlock: number | null
  /** Keys into `slots`, slot order. */
  slots: string[]
  /** Keys into `contributions`, submission order. */
  contributions: string[]
  finalization: FinalizationEntity | null
  /** The pool, always `POOL_SIZE` long, by key index. */
  poolKeys: PoolKeyEntity[]
  /** Index of the next key a registration would claim (`getPoolStatus.nextIndex`). */
  poolNext: number
  /** Keys into `applications`, registration order. */
  applications: string[]
  counts: EpochCounts
  events: number[]
  stateBlock: number
}

export interface AppPolicyEntity {
  /** Anyone may call submitCiphertext. */
  openSubmission: boolean
  /** Allow-list; empty means the registrant only. */
  submitters: Address[]
  maxCiphertexts: number
  notBeforeBlock: number
  notAfterBlock: number
  /** Decryption window as unix seconds; 0 = unbounded on that side. */
  decryptNotBefore: number
  decryptNotAfter: number
}

export interface ApplicationEntity {
  key: string
  epoch: EpochId
  aid: Aid
  creator: Address
  /** PK_org, TE form; the identity `(0, 1)` for an automatic application. */
  organizerPK: Point
  /** From the `ApplicationRegistered` event, so known before the record is read. */
  mode: AppModeName
  /** Pool key claimed at registration; null until the event or record is seen. */
  poolIndex: number | null
  /** `sk_org` once revealed; null until then, and always for an automatic application. */
  organizerSecret: bigint | null
  /** The `OrganizerSecretRevealed` event behind `organizerSecret`, when indexed. */
  organizerReveal: { block: number; tx: Hex | null } | null
  policy: AppPolicyEntity | null
  createdBlock: number
  createdTx: Hex | null
  /** Keys into `ciphertexts`, index order. */
  ciphertexts: string[]
  events: number[]
  stateBlock: number
}

export interface PartialEntity {
  participant: Address
  participantIndex: number
  block: number
  tx: Hex | null
  delta: Point
}

export interface CombineEntity {
  /** Resolved from the transaction sender (lazy). */
  by: Address | null
  block: number
  tx: Hex | null
  combineHash: Hex
  plaintext: bigint
}

export interface CiphertextEntity {
  key: string
  epoch: EpochId
  aid: Aid
  index: number
  submitter: Address
  c1: Point
  c2: Point
  block: number
  tx: Hex | null
  partials: PartialEntity[]
  combined: CombineEntity | null
}

export interface TxMeta {
  hash: Hex
  from: Address | null
  gasUsed: number | null
  blockNumber: number
  status: 'success' | 'reverted' | null
}

/** Chain-level facts, partly from config and partly read on chain. */
export interface ChainMeta {
  chainId: number
  chainName: string
  managerAddress: Address
  registryAddress: Address | null
  appManagerAddress: Address | null
  explorerUrl: string
  deployBlock: number
  headBlock: number
  epochPrefix: number | null
  epochDurationBlocks: number | null
  committeeSelectionBlocks: number | null
  keyAssemblyBlocks: number | null
  nextEpochStartBlock: number | null
  inactivityWindow: number | null
  activeCount: number | null
  nodeCount: number | null
  /** Seconds per block, used to turn block deltas into wall-clock estimates. */
  blockTimeSeconds: number
  /** Node's per-slot delay between decryption waves (`node.staggerBlocks`). */
  staggerBlocks: number
  stateBlock: number
}

export interface IndexerStore {
  version: number
  chain: ChainMeta
  /** Highest block whose logs are in the store. */
  lastIndexedBlock: number
  operators: Record<string, OperatorEntity>
  /** Operator keys in first-seen order. */
  operatorOrder: string[]
  epochs: Record<string, EpochEntity>
  /** Epoch keys in creation order (ascending nonce). */
  epochOrder: string[]
  slots: Record<string, SlotEntity>
  contributions: Record<string, ContributionEntity>
  applications: Record<string, ApplicationEntity>
  applicationOrder: string[]
  ciphertexts: Record<string, CiphertextEntity>
  txMeta: Record<string, TxMeta>
  /** Every event ever seen, ascending by (block, logIndex). */
  events: IndexedEvent[]
}

// ── Status ───────────────────────────────────────────────────────────────────

export interface IndexerError {
  at: number
  scope: 'scan' | 'state' | 'tx' | 'persist' | 'poll'
  message: string
}

export interface IndexerStatus {
  phase: 'idle' | 'loading' | 'scanning' | 'live' | 'error'
  /** True while a backfill (not an incremental poll) is running. */
  scanning: boolean
  /** Where the scan started — the deployment block, or the cached cursor. */
  fromBlock: number
  /** Last block whose logs have been indexed. */
  lastBlock: number
  /** Chain head as of the last poll. */
  headBlock: number
  /** 0…1 over the backfill range; 1 when caught up. */
  progress: number
  eventCount: number
  /** Number of `eth_getLogs` requests issued since start (for cost display). */
  requests: number
  lastPollAt: number | null
  errors: IndexerError[]
}

export interface IndexerSnapshot {
  store: IndexerStore
  status: IndexerStatus
}

// ── Keys ─────────────────────────────────────────────────────────────────────

export function operatorKey(address: Address | string): string {
  return address.toLowerCase()
}

export function epochKey(id: EpochId | string): string {
  return id.toLowerCase()
}

export function slotKey(epoch: EpochId, slot: number): string {
  return `${epoch.toLowerCase()}:${slot}`
}

export function contributionKey(epoch: EpochId, index: number): string {
  return `${epoch.toLowerCase()}:${index}`
}

export function applicationKey(epoch: EpochId, aid: Aid): string {
  return `${epoch.toLowerCase()}:${aid.toLowerCase()}`
}

export function ciphertextKey(epoch: EpochId, aid: Aid, index: number): string {
  return `${epoch.toLowerCase()}:${aid.toLowerCase()}:${index}`
}

export function txKey(hash: Hex | string): string {
  return hash.toLowerCase()
}
