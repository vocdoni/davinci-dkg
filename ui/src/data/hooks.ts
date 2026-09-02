// Hooks over the data source. Every one of them is a `useSyncExternalStore`
// on the same snapshot plus a memoised selector, so a page never holds
// derived state and never refetches: when the indexer publishes, the selectors
// that actually changed re-run and nothing else does.

import { useCallback, useEffect, useMemo, useRef, useSyncExternalStore } from 'react'
import {
  activity,
  applicationDetail,
  applicationRows,
  epochDetail,
  epochRows,
  eventFeed,
  networkStats,
  operatorDetail,
  operatorRows,
  partialMatrix,
  searchStore,
  type ActivityBucket,
  type ApplicationDetail,
  type ApplicationRow,
  type EpochDetail,
  type EpochFilter,
  type EpochRow,
  type FeedEntry,
  type NetworkStats,
  type OperatorDetail,
  type OperatorRow,
  type PartialMatrix,
  type SearchResult,
} from '../indexer/selectors'
import type {
  Address,
  Aid,
  EpochId,
  Hex,
  IndexerSnapshot,
  IndexerStatus,
  IndexerStore,
} from '../indexer/types'
import { useDataSource } from './context'
import type { DataSource, DataSourceKind } from './source'

/** The raw snapshot. Prefer the specific hooks below. */
export function useSnapshot(): IndexerSnapshot {
  const source = useDataSource()
  return useSyncExternalStore(source.subscribe, source.getSnapshot, source.getSnapshot)
}

export function useStore(): IndexerStore {
  return useSnapshot().store
}

export interface IndexerHandle {
  status: IndexerStatus
  kind: DataSourceKind
  /** Convenience: `status.phase === 'scanning'`. */
  scanning: boolean
  headBlock: number
  lastBlock: number
  progress: number
  refresh: () => Promise<void>
  source: DataSource
}

/** Scan progress, head block, errors, and a manual refresh. */
export function useIndexer(): IndexerHandle {
  const source = useDataSource()
  const { status } = useSnapshot()
  return useMemo(
    () => ({
      status,
      kind: source.kind,
      scanning: status.scanning,
      headBlock: status.headBlock,
      lastBlock: status.lastBlock,
      progress: status.progress,
      refresh: () => source.refresh(),
      source,
    }),
    [source, status],
  )
}

export function useNetworkStats(): NetworkStats {
  const store = useStore()
  return useMemo(() => networkStats(store), [store])
}

export function useEpochs(filter: EpochFilter = {}): EpochRow[] {
  const store = useStore()
  const { phase, query, limit } = filter
  return useMemo(() => epochRows(store, { phase, query, limit }), [store, phase, query, limit])
}

/**
 * Full epoch detail. Also asks the indexer to refresh the epoch's on-chain
 * struct and to resolve the transactions whose sender/gas the page shows.
 */
export function useEpoch(id: EpochId | string | undefined): EpochDetail | null {
  const source = useDataSource()
  const store = useStore()
  const detail = useMemo(() => (id ? epochDetail(store, id) : null), [store, id])

  useEffect(() => {
    if (id) source.ensureEpochState(id as EpochId)
  }, [source, id])

  const finalizationTx = detail?.finalization?.tx ?? null
  useEffect(() => {
    if (finalizationTx) source.ensureTxMeta([finalizationTx])
  }, [source, finalizationTx])

  const contributionCount = detail?.epoch.contributions.length ?? 0
  useEffect(() => {
    if (!id || contributionCount === 0) return
    const current = epochDetail(source.getSnapshot().store, id)
    if (!current) return
    source.ensureTxMeta(current.committee.map((row) => row.contributionTx))
  }, [source, id, contributionCount])

  return detail
}

export function useOperators(): OperatorRow[] {
  const store = useStore()
  return useMemo(() => operatorRows(store), [store])
}

export function useOperator(address: Address | string | undefined): OperatorDetail | null {
  const source = useDataSource()
  const store = useStore()
  useEffect(() => {
    if (address) source.ensureOperatorState(address as Address)
  }, [source, address])
  return useMemo(() => (address ? operatorDetail(store, address) : null), [store, address])
}

export function useApplications(): ApplicationRow[] {
  const store = useStore()
  return useMemo(() => applicationRows(store), [store])
}

export function useApplication(
  epoch: EpochId | string | undefined,
  aid: Aid | string | undefined,
): ApplicationDetail | null {
  const source = useDataSource()
  const store = useStore()
  const detail = useMemo(
    () => (epoch && aid ? applicationDetail(store, epoch, aid) : null),
    [store, epoch, aid],
  )

  useEffect(() => {
    if (epoch && aid) source.ensureApplicationState(epoch as EpochId, aid as Aid)
  }, [source, epoch, aid])

  const combinedCount = detail?.summary.combined ?? 0
  useEffect(() => {
    if (!epoch || !aid || combinedCount === 0) return
    const current = applicationDetail(source.getSnapshot().store, epoch, aid)
    if (!current) return
    const hashes: Array<Hex | null> = current.ciphertexts.map((row) => row.combined.tx)
    source.ensureTxMeta(hashes)
  }, [source, epoch, aid, combinedCount])

  return detail
}

/** Per-epoch activity for the newest `count` epochs, oldest first. */
export function useActivity(count = 30): ActivityBucket[] {
  const store = useStore()
  return useMemo(() => activity(store, count), [store, count])
}

export function useEventFeed(count = 20): FeedEntry[] {
  const store = useStore()
  return useMemo(() => eventFeed(store, count), [store, count])
}

export function usePartialMatrix(
  epoch: EpochId | string | undefined,
  aid?: Aid | string,
): PartialMatrix | null {
  const store = useStore()
  return useMemo(() => (epoch ? partialMatrix(store, epoch, aid) : null), [store, epoch, aid])
}

export function useSearch(query: string, limit = 10): SearchResult[] {
  const store = useStore()
  return useMemo(() => searchStore(store, query, limit), [store, query, limit])
}

/**
 * Alias of {@link useSearch}, for pages that also import the shell's
 * `useSearch()` (the global search box API) and need both in one file.
 */
export const useStoreSearch = useSearch

/**
 * A resolver for the shell's global search box, backed by the store.
 *
 * The shell only knows what a query *looks* like; this knows what exists. Wire
 * it up once, near the search provider:
 *
 *   useRegisterSearchResolver(useIndexerSearchResolver())
 *
 * The return type is structurally the shell's `SearchTarget`, so no import
 * crosses between the two layers.
 */
export type StoreSearchTarget =
  | { kind: 'route'; path: string; label: string }
  | { kind: 'external'; url: string; label: string }

export function useIndexerSearchResolver(): (
  query: string,
  ctx?: { explorerUrl?: string },
) => StoreSearchTarget | null {
  const source = useDataSource()
  const store = useStore()
  // The store object changes on every publish; keeping it in a ref lets the
  // resolver identity stay stable, so the shell does not re-register it every
  // poll.
  const latest = useRef(store)
  latest.current = store
  return useCallback(
    (query: string) => {
      const [best] = searchStore(latest.current, query, 1)
      if (!best) return null
      return best.external
        ? { kind: 'external' as const, url: best.href, label: best.label }
        : { kind: 'route' as const, path: best.href, label: best.label }
    },
    [source],
  )
}

/** Ask for `from` / `gasUsed` of arbitrary transactions a page is rendering. */
export function useTxMeta(hashes: Array<Hex | null | undefined>): void {
  const source = useDataSource()
  const key = hashes.filter(Boolean).join(',')
  useEffect(() => {
    if (key.length === 0) return
    source.ensureTxMeta(key.split(',') as Hex[])
  }, [source, key])
}
