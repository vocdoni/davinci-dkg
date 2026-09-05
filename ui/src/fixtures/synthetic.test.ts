import { describe, expect, it } from 'vitest'
import { buildFixture, DEFAULT_FIXTURE, GAS } from './synthetic'
import { encodeStore } from '../indexer/persist'
import { STORE_VERSION } from '../indexer/types'

const { store, meta } = buildFixture()

describe('synthetic fixture', () => {
  it('is deterministic', () => {
    const again = buildFixture()
    expect(encodeStore(again.store)).toEqual(encodeStore(store))
    expect(again.meta.epochIds).toEqual(meta.epochIds)
  })

  it('produces a store the indexer would produce', () => {
    expect(store.version).toBe(STORE_VERSION)
    expect(store.chain.chainId).toBe(11155111)
    expect(store.chain.managerAddress).toBe(DEFAULT_FIXTURE.managerAddress)
    expect(store.chain.headBlock).toBeGreaterThan(store.chain.deployBlock)
    expect(store.lastIndexedBlock).toBe(store.chain.headBlock)
    // Events are chronological and uniquely keyed by (block, logIndex).
    const keys = new Set<string>()
    let previous = -1
    for (const event of store.events) {
      expect(event.block).toBeGreaterThanOrEqual(previous)
      previous = event.block
      const key = `${event.block}:${event.logIndex}`
      expect(keys.has(key)).toBe(false)
      keys.add(key)
    }
  })

  it('registers 300 operators, reaps some and reactivates a few', () => {
    expect(meta.operators).toHaveLength(300)
    const registered = Object.values(store.operators).filter((o) => o.status !== 'none')
    expect(registered).toHaveLength(300)
    expect(registered.filter((o) => o.status === 'inactive')).toHaveLength(24)
    expect(registered.filter((o) => o.reactivations > 0)).toHaveLength(6)
    expect(store.chain.activeCount).toBe(276)
    expect(store.chain.nodeCount).toBe(300)
  })

  it('has 8 epochs: one aborted, one in KeyAssembly, six live', () => {
    expect(meta.epochIds).toHaveLength(8)
    const phases = meta.epochIds.map((id) => store.epochs[id].status)
    expect(phases.filter((p) => p === 'live')).toHaveLength(6)
    expect(phases.filter((p) => p === 'aborted')).toHaveLength(1)
    expect(phases.filter((p) => p === 'key-assembly')).toHaveLength(1)
    expect(store.epochs[meta.epochIds[2]].status).toBe('aborted')
    expect(store.epochs[meta.epochIds[7]].status).toBe('key-assembly')
  })

  it('fills 64-member committees at t = 33 and m_min = 40', () => {
    for (const id of meta.epochIds) {
      const epoch = store.epochs[id]
      expect(epoch.policy).toMatchObject({
        threshold: 33,
        committeeSize: 64,
        minValidContributions: 40,
        lotteryAlphaBps: 15_000,
      })
      const expected = epoch.status === 'aborted' ? 40 : 64
      expect(epoch.slots).toHaveLength(expected)
      expect(epoch.committee.filter(Boolean)).toHaveLength(expected)
    }
    // The aborted epoch never reached the committee size, and never finalized.
    const aborted = store.epochs[meta.epochIds[2]]
    expect(aborted.slots.length).toBeLessThan(64)
    expect(aborted.finalization).toBeNull()
    // KeyAssembly: full committee, not enough contributions yet.
    const assembling = store.epochs[meta.epochIds[7]]
    expect(assembling.slots).toHaveLength(64)
    expect(assembling.contributions.length).toBeLessThan(40)
    expect(assembling.finalization).toBeNull()
  })

  it('gives every Live epoch applications with 8 ciphertexts each', () => {
    const live = meta.epochIds.filter((id) => store.epochs[id].status === 'live')
    expect(live).toHaveLength(6)
    for (const id of live) {
      const epoch = store.epochs[id]
      expect(epoch.applications).toHaveLength(2)
      for (const key of epoch.applications) {
        expect(store.applications[key].ciphertexts).toHaveLength(8)
      }
    }
    expect(Object.keys(store.ciphertexts)).toHaveLength(96)
  })

  it('stores the whole pool at finalization and claims one key per application', () => {
    for (const id of meta.epochIds) {
      const epoch = store.epochs[id]
      expect(epoch.poolKeys).toHaveLength(16)
      if (epoch.status !== 'live') {
        expect(epoch.poolKeys.every((slot) => slot.key == null && slot.claimedBy == null)).toBe(true)
        continue
      }
      // Every key exists from the finalization block on, and each is distinct.
      expect(epoch.poolKeys.every((slot) => slot.key != null)).toBe(true)
      expect(new Set(epoch.poolKeys.map((slot) => `${slot.key!.x}:${slot.key!.y}`)).size).toBe(16)
      expect(epoch.poolKeys.filter((slot) => slot.claimedBy != null)).toHaveLength(2)
      expect(epoch.poolNext).toBe(2)
      epoch.applications.forEach((key, i) => {
        const app = store.applications[key]
        expect(app.poolIndex).toBe(i)
        expect(epoch.poolKeys[i].claimedBy).toBe(app.aid)
        expect(epoch.finalization!.block).toBeLessThanOrEqual(app.createdBlock)
      })
    }
  })

  it('delivers partials in waves of t, reveals most secrets, combines most', () => {
    const stagger = store.chain.staggerBlocks
    let awaitingReveal = 0
    let combined = 0
    for (const key of Object.keys(store.ciphertexts)) {
      const ct = store.ciphertexts[key]
      const app = store.applications[`${ct.epoch}:${ct.aid}`]
      if (app.mode === 'organizer-locked' && app.organizerSecret == null) {
        // The contract refuses every partial and combine before the reveal.
        expect(ct.partials).toHaveLength(0)
        expect(ct.combined).toBeNull()
        awaitingReveal += 1
        continue
      }
      expect(ct.partials.length).toBeGreaterThanOrEqual(33)
      const waves = new Set(ct.partials.map((p) => Math.floor((p.block - ct.block) / stagger)))
      expect(Math.min(...waves)).toBe(0)
      expect(Math.max(...waves)).toBeLessThanOrEqual(1)
      // The threshold is met inside the first wave.
      expect(ct.partials.filter((p) => p.block < ct.block + stagger).length).toBe(33)
      if (ct.combined) combined += 1
    }
    // One organizer (the newest Live epoch's) keeps its secret: 8 ciphertexts
    // parked; two automatic ciphertexts there are ready but not combined yet.
    expect(awaitingReveal).toBe(8)
    expect(combined).toBe(86)
    const locked = Object.values(store.applications).filter((app) => app.mode === 'organizer-locked')
    expect(locked).toHaveLength(6)
    expect(locked.filter((app) => app.organizerSecret != null)).toHaveLength(5)
    for (const app of locked) {
      if (app.organizerSecret == null) continue
      expect(app.organizerReveal?.block).toBeGreaterThan(app.createdBlock)
      // Nothing — no partial, no combine — lands before the reveal.
      for (const key of app.ciphertexts) {
        const ct = store.ciphertexts[key]
        for (const partial of ct.partials) expect(partial.block).toBeGreaterThan(app.organizerReveal!.block)
        expect(ct.combined!.block).toBeGreaterThan(app.organizerReveal!.block)
      }
    }
    // Automatic applications carry the identity as their organizer key.
    for (const app of Object.values(store.applications)) {
      if (app.mode === 'automatic') expect(app.organizerPK).toEqual({ x: 0n, y: 1n })
    }
  })

  it('uses the measured gas figures and a 12 s block cadence', () => {
    expect(store.chain.blockTimeSeconds).toBe(12)
    const finalization = store.epochs[meta.epochIds[0]].finalization
    expect(finalization?.tx).toBeTruthy()
    expect(store.txMeta[finalization?.tx as string].gasUsed).toBe(GAS.finalizeEpoch)
    const combine = Object.values(store.ciphertexts).find((ct) => ct.combined)?.combined
    expect(store.txMeta[combine?.tx as string].gasUsed).toBe(GAS.combineDecryption)
  })

  it('accepts overrides', () => {
    const small = buildFixture({ operators: 12, epochs: 2, committeeSize: 6, threshold: 3, seed: 7 })
    expect(small.meta.operators).toHaveLength(12)
    expect(small.meta.epochIds).toHaveLength(2)
    expect(small.store.epochs[small.meta.epochIds[0]].policy?.committeeSize).toBe(6)
  })
})
