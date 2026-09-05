// The indexer: one scan, one poll loop, one snapshot.
//
// Lifecycle
//   start()  → load the IndexedDB cache, resolve the sibling contract
//              addresses, backfill from the cursor (or `deployBlock`) in
//              adaptive chunks, then poll every `pollIntervalMs`.
//   tick()   → head block; if it moved, scan the new range, refresh the
//              state of whatever the new events touched, attribute the
//              transactions that need a sender, persist.
//
// It is an external store: `subscribe` / `getSnapshot` plug straight into
// `useSyncExternalStore`, and every publish produces a fresh top-level store
// object so memoised selectors invalidate exactly once per change.

import type { Address, PublicClient } from 'viem'
import { ALL_EVENT_ABIS } from './events'
import {
  applyApplicationState,
  applyChainState,
  applyEpochState,
  applyEvents,
  applyOperatorState,
  applyTxMeta,
  bumpStore,
  createEmptyStore,
} from './reduce'
import { createIdbStore, loadStore, saveStore, clearStore, type KVStore } from './persist'
import { DEFAULT_CHUNK, scanRange } from './scan'
import { StateReader } from './state'
import {
  applicationKey,
  epochKey,
  operatorKey,
  txKey,
  type Aid,
  type EpochId,
  type Hex,
  type IndexerError,
  type IndexerSnapshot,
  type IndexerStatus,
  type IndexerStore,
} from './types'

export interface IndexerConfig {
  client: PublicClient
  chainId: number
  chainName?: string
  managerAddress: Address
  deployBlock: number
  explorerUrl?: string
  registryAddress?: Address
  appManagerAddress?: Address
  /** Poll interval in ms; one chain block by default. */
  pollIntervalMs?: number
  /** Initial `getLogs` window; the scan adapts from here. */
  chunkSize?: number
  blockTimeSeconds?: number
  staggerBlocks?: number
  /** Persistence backend; `null` disables the cache. Defaults to IndexedDB. */
  kv?: KVStore | null
  /** Minimum ms between cache writes. */
  persistIntervalMs?: number
  /** Transactions resolved per tick, so a backfill cannot flood the RPC. */
  txPerTick?: number
  /**
   * Blocks after which an epoch struct / registry counter is re-read even
   * though no event touched it. Everything that changes an epoch's phase emits
   * an event, so this is only a safety net — keep it well above the poll
   * interval so an idle chain costs one `eth_blockNumber` per poll and nothing
   * else.
   */
  stateStaleBlocks?: number
}

const MAX_ERRORS = 20
const PUBLISH_THROTTLE_MS = 200

function emptyStatus(fromBlock: number): IndexerStatus {
  return {
    phase: 'idle',
    scanning: false,
    fromBlock,
    lastBlock: fromBlock,
    headBlock: fromBlock,
    progress: 0,
    eventCount: 0,
    requests: 0,
    lastPollAt: null,
    errors: [],
  }
}

export class Indexer {
  readonly config: Required<
    Pick<IndexerConfig, 'pollIntervalMs' | 'persistIntervalMs' | 'txPerTick' | 'stateStaleBlocks'>
  > &
    IndexerConfig
  /** Deploy block found by bisection when the config has none. */
  private discoveredDeployBlock: number | null = null

  private store: IndexerStore
  private status: IndexerStatus
  private snapshot: IndexerSnapshot
  private listeners = new Set<() => void>()

  private readonly kv: KVStore | null
  private readonly reader: StateReader
  private chunkSize: number

  private started = false
  private timer: ReturnType<typeof setTimeout> | null = null
  private abort: AbortController | null = null
  private running: Promise<void> | null = null
  private bootstrapped = false

  private dirtyEpochs = new Set<string>()
  private dirtyOperators = new Set<string>()
  private dirtyApplications = new Set<string>()
  private pendingTx = new Set<string>()

  private pendingPersist = false
  private lastPublish = 0
  private publishTimer: ReturnType<typeof setTimeout> | null = null
  private lastPersist = 0

  constructor(config: IndexerConfig) {
    this.config = {
      pollIntervalMs: 12_000,
      persistIntervalMs: 5_000,
      txPerTick: 25,
      stateStaleBlocks: 50,
      ...config,
    }
    this.kv = config.kv === null ? null : (config.kv ?? createIdbStore())
    this.chunkSize = config.chunkSize ?? DEFAULT_CHUNK
    this.store = createEmptyStore({
      chainId: config.chainId,
      chainName: config.chainName,
      managerAddress: config.managerAddress,
      registryAddress: config.registryAddress ?? null,
      appManagerAddress: config.appManagerAddress ?? null,
      explorerUrl: config.explorerUrl,
      deployBlock: config.deployBlock,
      blockTimeSeconds: config.blockTimeSeconds,
      staggerBlocks: config.staggerBlocks,
    })
    this.status = emptyStatus(config.deployBlock)
    this.snapshot = { store: this.store, status: this.status }
    this.reader = new StateReader({
      client: config.client,
      managerAddress: config.managerAddress,
      registryAddress: config.registryAddress ?? null,
      appManagerAddress: config.appManagerAddress ?? null,
    })
  }

  // ── external store ─────────────────────────────────────────────────────────

  subscribe = (listener: () => void): (() => void) => {
    this.listeners.add(listener)
    return () => {
      this.listeners.delete(listener)
    }
  }

  getSnapshot = (): IndexerSnapshot => this.snapshot

  // ── lifecycle ──────────────────────────────────────────────────────────────

  start(): void {
    if (this.started) return
    this.started = true
    this.abort = new AbortController()
    void this.loop()
  }

  stop(): void {
    this.started = false
    this.abort?.abort()
    this.abort = null
    if (this.timer) clearTimeout(this.timer)
    this.timer = null
    if (this.publishTimer) clearTimeout(this.publishTimer)
    this.publishTimer = null
  }

  /** Run one tick now (also used by a manual "refresh" button). */
  async refresh(): Promise<void> {
    if (this.running) return this.running
    this.running = this.tick().finally(() => {
      this.running = null
    })
    return this.running
  }

  async clearCache(): Promise<void> {
    if (this.kv) await clearStore(this.kv, this.config.chainId, this.config.managerAddress)
  }

  /** Ask for `from`/`gasUsed` of transactions the UI is about to show. */
  ensureTxMeta(hashes: Array<Hex | null | undefined>): void {
    let added = false
    for (const hash of hashes) {
      if (!hash) continue
      const key = txKey(hash)
      if (this.store.txMeta[key] || this.pendingTx.has(key)) continue
      this.pendingTx.add(key)
      added = true
    }
    if (added && this.started) void this.refresh()
  }

  /** Ask for a fresh `getEpoch` on the next tick. */
  ensureEpochState(id: EpochId): void {
    this.dirtyEpochs.add(epochKey(id))
  }

  ensureOperatorState(address: Address): void {
    this.dirtyOperators.add(operatorKey(address))
  }

  ensureApplicationState(epoch: EpochId, aid: Aid): void {
    this.dirtyApplications.add(applicationKey(epoch, aid))
  }

  // ── internals ──────────────────────────────────────────────────────────────

  private async loop(): Promise<void> {
    while (this.started) {
      await this.refresh()
      if (!this.started) break
      await new Promise<void>((resolve) => {
        this.timer = setTimeout(resolve, this.config.pollIntervalMs)
      })
    }
  }

  private async bootstrap(): Promise<void> {
    if (this.bootstrapped) return
    this.bootstrapped = true
    this.status = { ...this.status, phase: 'loading' }
    this.publish(true)

    if (this.kv) {
      const cached = await loadStore(this.kv, this.config.chainId, this.config.managerAddress)
      if (cached) {
        // Config wins over whatever the cache remembers about the deployment.
        cached.chain = {
          ...cached.chain,
          explorerUrl: this.config.explorerUrl ?? cached.chain.explorerUrl,
          chainName: this.config.chainName ?? cached.chain.chainName,
          blockTimeSeconds: this.config.blockTimeSeconds ?? cached.chain.blockTimeSeconds,
          staggerBlocks: this.config.staggerBlocks ?? cached.chain.staggerBlocks,
        }
        this.store = cached
        this.reader.registryAddress = cached.chain.registryAddress
        this.reader.appManagerAddress = cached.chain.appManagerAddress
        this.status = {
          ...this.status,
          fromBlock: cached.lastIndexedBlock + 1,
          lastBlock: cached.lastIndexedBlock,
          eventCount: cached.events.length,
        }
        this.publish(true)
      }
    }

    if (!this.store.chain.registryAddress || !this.store.chain.appManagerAddress) {
      try {
        const meta = await this.reader.readChainMeta(this.store.lastIndexedBlock)
        applyChainState(this.store, meta)
      } catch (err) {
        this.pushError('state', err)
      }
    }
  }

  private async tick(): Promise<void> {
    try {
      await this.bootstrap()
      const head = Number(await this.config.client.getBlockNumber())
      this.status = {
        ...this.status,
        headBlock: head,
        lastPollAt: Date.now(),
        requests: this.status.requests + 1,
      }
      applyChainState(this.store, { headBlock: head })

      if (this.config.deployBlock <= 0 && this.discoveredDeployBlock == null && this.store.lastIndexedBlock < 0) {
        // No deploy block configured: find the manager's creation block once
        // (about 25 eth_getCode calls) instead of scanning from genesis.
        this.discoveredDeployBlock = await discoverDeployBlock(this.config.client, this.config.managerAddress, head)
        this.status = { ...this.status, fromBlock: this.discoveredDeployBlock, requests: this.status.requests + 25 }
      }
      const deployBlock = this.discoveredDeployBlock ?? this.config.deployBlock
      const from = Math.max(deployBlock, this.store.lastIndexedBlock + 1)
      if (head >= from) {
        await this.scan(from, head)
      }
      this.store.lastIndexedBlock = Math.max(this.store.lastIndexedBlock, head)
      await this.refreshState(head)
      await this.resolveTransactions()
      this.status = {
        ...this.status,
        phase: 'live',
        scanning: false,
        progress: 1,
        lastBlock: this.store.lastIndexedBlock,
        eventCount: this.store.events.length,
        requests: this.status.requests + this.reader.requests,
      }
      this.reader.requests = 0
      this.publish(true)
      // Only rewrite the cache when something actually changed: encoding the
      // store is the most expensive thing an idle poll could possibly do.
      if (this.pendingPersist) {
        this.pendingPersist = false
        await this.persist()
      }
    } catch (err) {
      this.pushError('poll', err)
      this.status = { ...this.status, phase: 'error', scanning: false }
      this.publish(true)
    }
  }

  private async scan(from: number, to: number): Promise<void> {
    const span = to - from + 1
    const backfill = span > this.chunkSize
    if (backfill) {
      this.status = { ...this.status, phase: 'scanning', scanning: true, fromBlock: from, progress: 0 }
      this.publish(true)
    }
    const result = await scanRange({
      client: this.config.client,
      addresses: this.contractAddresses(),
      events: ALL_EVENT_ABIS,
      fromBlock: from,
      toBlock: to,
      chunkSize: this.chunkSize,
      signal: this.abort?.signal,
      onChunk: async (chunk) => {
        const applied = applyEvents(this.store, chunk.events)
        if (applied > 0) this.pendingPersist = true
        this.markDirty(applied > 0)
        this.store.lastIndexedBlock = Math.max(this.store.lastIndexedBlock, chunk.to)
        this.status = {
          ...this.status,
          lastBlock: this.store.lastIndexedBlock,
          eventCount: this.store.events.length,
          progress: span <= 0 ? 1 : Math.min(1, (chunk.to - from + 1) / span),
          requests: this.status.requests + 1,
        }
        this.publish()
        if (backfill) await this.persistThrottled()
      },
    })
    this.chunkSize = result.chunkSize
  }

  private contractAddresses(): Address[] {
    const addresses = [this.store.chain.managerAddress]
    if (this.store.chain.registryAddress) addresses.push(this.store.chain.registryAddress)
    if (this.store.chain.appManagerAddress) addresses.push(this.store.chain.appManagerAddress)
    return addresses
  }

  /** Mark whatever the newest events touched as needing a state read. */
  private markDirty(hasNew: boolean): void {
    if (!hasNew) return
    // Only the tail matters: everything before the previous cursor was already
    // marked when it arrived.
    const events = this.store.events
    for (let i = events.length - 1; i >= 0 && i >= events.length - 500; i--) {
      const ev = events[i]
      if (ev.epoch) this.dirtyEpochs.add(epochKey(ev.epoch))
      if (ev.actor) this.dirtyOperators.add(operatorKey(ev.actor))
      if (ev.epoch && ev.aid) this.dirtyApplications.add(applicationKey(ev.epoch, ev.aid))
      if (ev.name === 'EpochLive' && ev.tx) this.pendingTx.add(txKey(ev.tx))
      if (ev.name === 'DecryptionCombined' && ev.tx) this.pendingTx.add(txKey(ev.tx))
    }
  }

  private async refreshState(head: number): Promise<void> {
    const stale = (block: number): boolean =>
      block === 0 || head - block >= this.config.stateStaleBlocks
    // The newest epochs are re-read when their struct has gone stale. Their
    // phase changes all emit an event (CommitteeFilled / EpochLive /
    // EpochAborted), which already marks them dirty, so this only catches
    // counters drifting on a chain nobody is touching.
    for (const key of this.store.epochOrder.slice(-3)) {
      const epoch = this.store.epochs[key]
      if (epoch.status !== 'aborted' && stale(epoch.stateBlock)) this.dirtyEpochs.add(key)
    }

    const epochIds = [...this.dirtyEpochs].filter((key) => this.store.epochs[key])
    this.dirtyEpochs.clear()
    if (epochIds.length > 0) {
      this.pendingPersist = true
      try {
        const updates = await this.reader.readEpochs(epochIds as EpochId[], head)
        for (const [id, update] of updates) applyEpochState(this.store, id as EpochId, update)
      } catch (err) {
        this.pushError('state', err)
      }
    }

    const operators = [...this.dirtyOperators].filter((key) => this.store.operators[key])
    this.dirtyOperators.clear()
    if (operators.length > 0) {
      try {
        const updates = await this.reader.readOperators(operators as Address[], head)
        for (const [address, update] of updates) applyOperatorState(this.store, address as Address, update)
      } catch (err) {
        this.pushError('state', err)
      }
    }

    const apps = [...this.dirtyApplications]
      .map((key) => this.store.applications[key])
      .filter((app) => app != null && app.policy == null)
    this.dirtyApplications.clear()
    if (apps.length > 0) {
      try {
        const updates = await this.reader.readApplications(
          apps.map((app) => ({ epoch: app.epoch, aid: app.aid })),
          head,
        )
        for (const [key, update] of updates) {
          const app = this.store.applications[key]
          if (app) applyApplicationState(this.store, app.epoch, app.aid, update)
        }
      } catch (err) {
        this.pushError('state', err)
      }
    }

    if (stale(this.store.chain.stateBlock)) {
      this.pendingPersist = true
      try {
        applyChainState(this.store, await this.reader.readRegistryCounters(head))
      } catch (err) {
        this.pushError('state', err)
      }
    }
  }

  private async resolveTransactions(): Promise<void> {
    if (this.pendingTx.size === 0) return
    const batch = [...this.pendingTx].slice(0, this.config.txPerTick)
    for (const hash of batch) this.pendingTx.delete(hash)
    try {
      const metas = await this.reader.readTxMeta(batch as Hex[])
      if (metas.length > 0) this.pendingPersist = true
      applyTxMeta(this.store, metas)
    } catch (err) {
      this.pushError('tx', err)
    }
  }

  private pushError(scope: IndexerError['scope'], err: unknown): void {
    const message = err instanceof Error ? err.message : String(err)
    const errors = [...this.status.errors, { at: Date.now(), scope, message }].slice(-MAX_ERRORS)
    this.status = { ...this.status, errors }
  }

  private async persist(): Promise<void> {
    if (!this.kv) return
    try {
      this.lastPersist = Date.now()
      await saveStore(this.kv, this.store)
    } catch (err) {
      this.pushError('persist', err)
    }
  }

  private async persistThrottled(): Promise<void> {
    if (Date.now() - this.lastPersist < this.config.persistIntervalMs) return
    await this.persist()
  }

  private publish(immediate = false): void {
    this.store = bumpStore(this.store)
    this.snapshot = { store: this.store, status: this.status }
    const now = Date.now()
    if (!immediate && now - this.lastPublish < PUBLISH_THROTTLE_MS) {
      if (!this.publishTimer) {
        this.publishTimer = setTimeout(() => {
          this.publishTimer = null
          this.publish(true)
        }, PUBLISH_THROTTLE_MS)
      }
      return
    }
    if (this.publishTimer) {
      clearTimeout(this.publishTimer)
      this.publishTimer = null
    }
    this.lastPublish = now
    for (const listener of this.listeners) listener()
  }
}

export function createIndexer(config: IndexerConfig): Indexer {
  return new Indexer(config)
}

/**
 * Binary-search the first block at which `address` has code. Contracts are
 * never destroyed here, so "has code" is monotone in the block number.
 */
export async function discoverDeployBlock(
  client: Pick<PublicClient, 'getCode'>,
  address: Address,
  head: number,
): Promise<number> {
  let lo = 0
  let hi = head
  while (lo < hi) {
    const mid = Math.floor((lo + hi) / 2)
    const code = await client.getCode({ address, blockNumber: BigInt(mid) })
    if (code && code !== '0x') hi = mid
    else lo = mid + 1
  }
  return lo
}
