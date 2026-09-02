// IndexedDB persistence for the entity store.
//
// The store is keyed by `chainId:managerAddress` and tagged with
// `STORE_VERSION`; anything written by an older build (or by another
// deployment) is ignored and re-scanned rather than migrated. The value is a
// JSON string with `bigint`s wrapped, so the same code path works against
// IndexedDB, `idb-keyval`, `localStorage` or an in-memory mock.

import {
  createStore as createIdbKeyvalStore,
  del as idbDel,
  get as idbGet,
  set as idbSet,
} from 'idb-keyval'
import { STORE_VERSION, type IndexerStore } from './types'

/**
 * The three calls the indexer needs from a key/value store. `idb-keyval`'s
 * `get` / `set` / `del` satisfy this as-is:
 *
 *   import * as idb from 'idb-keyval'
 *   new Indexer({ ..., kv: idb })
 */
export interface KVStore {
  get(key: string): Promise<unknown>
  set(key: string, value: unknown): Promise<void>
  del(key: string): Promise<void>
}

export interface PersistedEnvelope {
  version: number
  chainId: number
  manager: string
  savedAt: number
  lastIndexedBlock: number
  /** `encodeStore` output. */
  data: string
}

export function cacheKey(chainId: number, managerAddress: string): string {
  return `dkg-explorer:v${STORE_VERSION}:${chainId}:${managerAddress.toLowerCase()}`
}

interface BigIntBox {
  $bigint: string
}

function isBigIntBox(v: unknown): v is BigIntBox {
  return typeof v === 'object' && v !== null && typeof (v as BigIntBox).$bigint === 'string'
}

/** JSON with `bigint` boxed as `{ $bigint: "…" }`. */
export function encodeStore(store: IndexerStore): string {
  return JSON.stringify(store, (_key, value) =>
    typeof value === 'bigint' ? { $bigint: value.toString() } : value,
  )
}

export function decodeStore(text: string): IndexerStore {
  return JSON.parse(text, (_key, value) => (isBigIntBox(value) ? BigInt(value.$bigint) : value)) as IndexerStore
}

/**
 * Read the cached store for this deployment. Returns null when nothing is
 * cached, when the cache was written by another version/chain/manager, or
 * when it cannot be parsed — every one of which just means "scan again".
 */
export async function loadStore(
  kv: KVStore,
  chainId: number,
  managerAddress: string,
): Promise<IndexerStore | null> {
  try {
    const raw = (await kv.get(cacheKey(chainId, managerAddress))) as PersistedEnvelope | undefined
    if (!raw || typeof raw !== 'object') return null
    if (raw.version !== STORE_VERSION) return null
    if (raw.chainId !== chainId) return null
    if (raw.manager?.toLowerCase() !== managerAddress.toLowerCase()) return null
    const store = decodeStore(raw.data)
    if (store.version !== STORE_VERSION) return null
    return store
  } catch {
    return null
  }
}

export async function saveStore(kv: KVStore, store: IndexerStore): Promise<void> {
  const envelope: PersistedEnvelope = {
    version: STORE_VERSION,
    chainId: store.chain.chainId,
    manager: store.chain.managerAddress,
    savedAt: Date.now(),
    lastIndexedBlock: store.lastIndexedBlock,
    data: encodeStore(store),
  }
  await kv.set(cacheKey(store.chain.chainId, store.chain.managerAddress), envelope)
}

export async function clearStore(kv: KVStore, chainId: number, managerAddress: string): Promise<void> {
  await kv.del(cacheKey(chainId, managerAddress))
}

// ── stores ───────────────────────────────────────────────────────────────────

/** Non-persistent fallback: used in tests, SSR and private-mode browsers. */
export function memoryStore(): KVStore {
  const map = new Map<string, unknown>()
  return {
    async get(key) {
      return map.get(key)
    },
    async set(key, value) {
      map.set(key, value)
    },
    async del(key) {
      map.delete(key)
    },
  }
}

/**
 * IndexedDB via `idb-keyval`, in its own database so it never collides with
 * anything else the page stores. Falls back to memory wherever IndexedDB is
 * unavailable (SSR, private mode, jsdom), which just means "scan every time".
 */
export function createIdbStore(dbName = 'davinci-dkg-explorer', storeName = 'indexer'): KVStore {
  if (typeof indexedDB === 'undefined') return memoryStore()
  try {
    const store = createIdbKeyvalStore(dbName, storeName)
    return {
      get: (key) => idbGet(key, store),
      set: (key, value) => idbSet(key, value, store),
      del: (key) => idbDel(key, store),
    }
  } catch {
    return memoryStore()
  }
}
