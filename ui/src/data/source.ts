// The contract between the pages and whatever is feeding them.
//
// Two implementations exist: the live indexer (viem + the SDK against a real
// chain) and the synthetic fixture (`?demo=1`, no RPC at all). Both are
// external stores, so every hook in this folder is a `useSyncExternalStore`
// over the same snapshot shape and no page knows which one it is talking to.

import type { Address, Aid, EpochId, Hex, IndexerSnapshot } from '../indexer/types'
import { Indexer, type IndexerConfig } from '../indexer/indexer'

export type DataSourceKind = 'live' | 'demo'

export interface DataSource {
  readonly kind: DataSourceKind
  /** `useSyncExternalStore` pair. */
  subscribe(listener: () => void): () => void
  getSnapshot(): IndexerSnapshot
  /** Begin scanning / advancing. Idempotent. */
  start(): void
  stop(): void
  /** Force one poll (the "refresh" button, or after a write lands). */
  refresh(): Promise<void>
  /** Lazily resolve `from` / `gasUsed` for transactions a page is showing. */
  ensureTxMeta(hashes: Array<Hex | null | undefined>): void
  ensureEpochState(id: EpochId): void
  ensureOperatorState(address: Address): void
  ensureApplicationState(epoch: EpochId, aid: Aid): void
}

/** Wrap a live `Indexer` as a `DataSource`. */
export function createLiveDataSource(config: IndexerConfig): DataSource & { indexer: Indexer } {
  const indexer = new Indexer(config)
  return {
    kind: 'live',
    indexer,
    subscribe: indexer.subscribe,
    getSnapshot: indexer.getSnapshot,
    start: () => indexer.start(),
    stop: () => indexer.stop(),
    refresh: () => indexer.refresh(),
    ensureTxMeta: (hashes) => indexer.ensureTxMeta(hashes),
    ensureEpochState: (id) => indexer.ensureEpochState(id),
    ensureOperatorState: (address) => indexer.ensureOperatorState(address),
    ensureApplicationState: (epoch, aid) => indexer.ensureApplicationState(epoch, aid),
  }
}
