import { describe, expect, it } from 'vitest'
import { applyEvents, createEmptyStore } from './reduce'
import {
  cacheKey,
  clearStore,
  decodeStore,
  encodeStore,
  loadStore,
  saveStore,
  type KVStore,
} from './persist'
import { STORE_VERSION, type Address, type EpochId, type Hex } from './types'

const MANAGER = '0x3f9b338706a31f26d49159478015c8aaeab908ad' as Address
const EPOCH = '0x2f1105e90000000000000001' as EpochId
const ALICE = '0x1111111111111111111111111111111111111111' as Address

/**
 * Stands in for `idb-keyval`: same `get`/`set`/`del` signature, and it clones
 * through JSON so a value that survives here survives structured clone too.
 */
function mockIdbKeyval(): KVStore & { raw: Map<string, string> } {
  const raw = new Map<string, string>()
  return {
    raw,
    async get(key) {
      const value = raw.get(key)
      return value === undefined ? undefined : JSON.parse(value)
    },
    async set(key, value) {
      raw.set(key, JSON.stringify(value))
    },
    async del(key) {
      raw.delete(key)
    },
  }
}

function seededStore() {
  const store = createEmptyStore({ chainId: 11155111, managerAddress: MANAGER, deployBlock: 100 })
  applyEvents(store, [
    {
      name: 'EpochCreated',
      block: 300,
      tx: ('0x' + 'ab'.repeat(32)) as Hex,
      logIndex: 0,
      epoch: EPOCH,
      aid: null,
      actor: ALICE,
      data: {
        epochId: EPOCH,
        organizer: ALICE,
        startBlock: 300,
        seedBlock: 301,
        // A value no JSON number can hold — the codec has to keep it exact.
        lotteryThreshold: (1n << 255n) + 12345n,
      },
    },
  ])
  return store
}

describe('store codec', () => {
  it('round-trips bigints exactly', () => {
    const store = seededStore()
    const decoded = decodeStore(encodeStore(store))
    expect(decoded.epochs[EPOCH].lotteryThreshold).toBe((1n << 255n) + 12345n)
    expect(typeof decoded.epochs[EPOCH].lotteryThreshold).toBe('bigint')
    expect(decoded.events).toHaveLength(1)
    expect(decoded.chain.managerAddress).toBe(MANAGER)
  })
})

describe('persistence', () => {
  it('saves and reloads the store for one deployment', async () => {
    const kv = mockIdbKeyval()
    const store = seededStore()
    store.lastIndexedBlock = 512
    await saveStore(kv, store)

    expect([...kv.raw.keys()]).toEqual([cacheKey(11155111, MANAGER)])

    const loaded = await loadStore(kv, 11155111, MANAGER)
    expect(loaded).not.toBeNull()
    expect(loaded?.lastIndexedBlock).toBe(512)
    expect(loaded?.epochs[EPOCH].nonce).toBe(1)
    expect(loaded?.epochs[EPOCH].lotteryThreshold).toBe((1n << 255n) + 12345n)
  })

  it('ignores a cache from another chain, manager or schema version', async () => {
    const kv = mockIdbKeyval()
    const store = seededStore()
    await saveStore(kv, store)

    expect(await loadStore(kv, 1, MANAGER)).toBeNull()
    expect(await loadStore(kv, 11155111, '0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef')).toBeNull()

    const key = cacheKey(11155111, MANAGER)
    const envelope = JSON.parse(kv.raw.get(key) as string)
    envelope.version = STORE_VERSION + 1
    kv.raw.set(key, JSON.stringify(envelope))
    expect(await loadStore(kv, 11155111, MANAGER)).toBeNull()
  })

  it('survives a corrupted cache entry', async () => {
    const kv = mockIdbKeyval()
    await kv.set(cacheKey(11155111, MANAGER), {
      version: STORE_VERSION,
      chainId: 11155111,
      manager: MANAGER,
      savedAt: 0,
      lastIndexedBlock: 0,
      data: '{not json',
    })
    expect(await loadStore(kv, 11155111, MANAGER)).toBeNull()
  })

  it('clears the cache', async () => {
    const kv = mockIdbKeyval()
    await saveStore(kv, seededStore())
    await clearStore(kv, 11155111, MANAGER)
    expect(await loadStore(kv, 11155111, MANAGER)).toBeNull()
  })

  it('keys the cache by chain and manager', () => {
    expect(cacheKey(11155111, MANAGER)).toBe(`dkg-explorer:v${STORE_VERSION}:11155111:${MANAGER}`)
    expect(cacheKey(31337, MANAGER)).not.toBe(cacheKey(11155111, MANAGER))
  })
})
