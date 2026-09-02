// Demo data source: the synthetic fixture behind the `DataSource` interface,
// with a head block that keeps moving so countdowns, "live since" strips and
// the block ticker behave like they do on a real chain.

import { bumpStore } from '../indexer/reduce'
import type { DataSource } from '../data/source'
import type { IndexerSnapshot, IndexerStatus, IndexerStore } from '../indexer/types'
import { buildFixture, type FixtureOptions } from './synthetic'

export interface DemoSourceOptions extends FixtureOptions {
  /** Wall-clock ms between fake blocks; 0 freezes the chain. */
  blockIntervalMs?: number
}

export function createDemoDataSource(options: DemoSourceOptions = {}): DataSource {
  const { blockIntervalMs = 12_000, ...fixtureOptions } = options
  const { store: initial } = buildFixture(fixtureOptions)
  let store: IndexerStore = initial
  let status: IndexerStatus = {
    phase: 'live',
    scanning: false,
    fromBlock: store.chain.deployBlock,
    lastBlock: store.lastIndexedBlock,
    headBlock: store.chain.headBlock,
    progress: 1,
    eventCount: store.events.length,
    requests: 0,
    lastPollAt: Date.now(),
    errors: [],
  }
  let snapshot: IndexerSnapshot = { store, status }
  const listeners = new Set<() => void>()
  let timer: ReturnType<typeof setInterval> | null = null

  const publish = (): void => {
    store = bumpStore(store)
    snapshot = { store, status }
    for (const listener of listeners) listener()
  }

  const advance = (): void => {
    const head = store.chain.headBlock + 1
    store.chain = { ...store.chain, headBlock: head }
    store.lastIndexedBlock = head
    status = { ...status, headBlock: head, lastBlock: head, lastPollAt: Date.now() }
    publish()
  }

  return {
    kind: 'demo',
    subscribe(listener) {
      listeners.add(listener)
      return () => {
        listeners.delete(listener)
      }
    },
    getSnapshot: () => snapshot,
    start() {
      if (timer || blockIntervalMs <= 0) return
      timer = setInterval(advance, blockIntervalMs)
    },
    stop() {
      if (timer) clearInterval(timer)
      timer = null
    },
    async refresh() {
      advance()
    },
    ensureTxMeta() {
      // The fixture ships every transaction it references.
    },
    ensureEpochState() {},
    ensureOperatorState() {},
    ensureApplicationState() {},
  }
}
