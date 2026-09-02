// Public surface of the indexer. Pages import from `~/indexer` (types and
// selectors) and from `~/data` (hooks); nothing else reaches inside.

export * from './types'
export * from './selectors'
export { Indexer, createIndexer, type IndexerConfig } from './indexer'
export {
  createEmptyStore,
  bumpStore,
  applyEvents,
  applyEpochState,
  applyOperatorState,
  applyApplicationState,
  applyChainState,
  applyTxMeta,
  ensureEpoch,
  ensureOperator,
  phaseFromStatus,
  statusFromPhase,
  nodeStatusFromCode,
  type StoreSeed,
  type EpochStateUpdate,
  type OperatorStateUpdate,
  type ApplicationStateUpdate,
} from './reduce'
export {
  normalizeLog,
  compareEvents,
  ALL_EVENT_ABIS,
  MANAGER_EVENT_ABIS,
  REGISTRY_EVENT_ABIS,
  APP_MANAGER_EVENT_ABIS,
  type RawLog,
} from './events'
export {
  scanRange,
  estimateRequests,
  isRangeError,
  DEFAULT_CHUNK,
  MIN_CHUNK,
  MAX_CHUNK,
  type ScanOptions,
  type ScanResult,
  type ChunkReport,
} from './scan'
export {
  StateReader,
  createStateReader,
  MULTICALL3_ADDRESS,
  type StateReaderOptions,
} from './state'
export {
  createIdbStore,
  memoryStore,
  loadStore,
  saveStore,
  clearStore,
  cacheKey,
  encodeStore,
  decodeStore,
  type KVStore,
  type PersistedEnvelope,
} from './persist'
