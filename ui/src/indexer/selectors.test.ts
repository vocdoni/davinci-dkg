import { describe, expect, it } from 'vitest'
import { buildFixture } from '../fixtures/synthetic'
import {
  activity,
  applicationDetail,
  applicationRows,
  epochDetail,
  epochRows,
  eventFeed,
  formatParticipation,
  fractionOfHashSpace,
  networkStats,
  operatorDetail,
  operatorRows,
  partialMatrix,
  recoverRegistrySnapshot,
  searchStore,
} from './selectors'
import { bumpStore } from './reduce'

const { store, meta } = buildFixture()
const liveEpochId = meta.epochIds.find((id) => store.epochs[id].status === 'live')!
const abortedEpochId = meta.epochIds[2]
const assemblingEpochId = meta.epochIds[7]

describe('networkStats', () => {
  it('summarises the whole deployment', () => {
    const stats = networkStats(store)
    expect(stats).toMatchObject({
      chainId: 11155111,
      operatorsRegistered: 300,
      operatorsActive: 276,
      operatorsInactive: 24,
      epochs: 8,
      epochsLive: 6,
      epochsAborted: 1,
      thresholdInForce: 33,
      committeeSizeInForce: 64,
      applications: 12,
      ciphertexts: 96,
      ciphertextsDecrypted: 68,
    })
    // 7 full committees plus the aborted epoch's partial one.
    expect(stats.claims).toBe(7 * 64 + 40)
    expect(stats.partials).toBeGreaterThan(3_000)
    expect(stats.blocksToNextEpoch).toBeGreaterThan(0)
  })

  it('memoises per store identity', () => {
    expect(networkStats(store)).toBe(networkStats(store))
    expect(networkStats(bumpStore(store))).not.toBe(networkStats(store))
  })
})

describe('epochRows', () => {
  it('lists epochs newest first with progress fractions', () => {
    const rows = epochRows(store)
    expect(rows).toHaveLength(8)
    expect(rows[0].nonce).toBe(8)
    expect(rows[rows.length - 1].nonce).toBe(1)
    const live = rows.find((row) => row.id === liveEpochId)!
    expect(live.claimProgress).toBe(1)
    expect(live.contributionProgress).toBeGreaterThan(0.6)
    expect(live.finalizer).toBeTruthy()
    expect(live.finalizationGas).toBe(1_112_337)
    expect(live.ciphertexts).toBe(16)
  })

  it('filters by phase and query', () => {
    expect(epochRows(store, { phase: 'aborted' })).toHaveLength(1)
    expect(epochRows(store, { phase: 'live' })).toHaveLength(6)
    expect(epochRows(store, { query: '3' })).toHaveLength(1)
    expect(epochRows(store, { query: '3' })[0].nonce).toBe(3)
    expect(epochRows(store, { query: store.epochs[liveEpochId].creator })).toHaveLength(1)
    expect(epochRows(store, { query: liveEpochId })).toHaveLength(1)
    expect(epochRows(store, { limit: 3 })).toHaveLength(3)
  })
})

describe('epochDetail', () => {
  it('exposes the lottery with τ as a fraction of the hash space', () => {
    const detail = epochDetail(store, liveEpochId)!
    expect(detail.lottery.seed).toMatch(/^0x[0-9a-f]{64}$/)
    // τ = α·n·2²⁵⁶/R with α = 1.5, n = 64, R = 276 ⇒ 0.3478…
    expect(detail.lottery.thresholdFraction).toBeCloseTo((1.5 * 64) / 276, 6)
    expect(detail.lottery.alpha).toBe(1.5)
    expect(detail.lottery.registrySnapshot).toBe(276)
    expect(detail.lottery.admissibleProbability).toBeCloseTo(0.3478, 4)
    expect(detail.lottery.claims).toHaveLength(64)
    expect(detail.lottery.claims[0].slot).toBe(0)
  })

  it('builds one committee row per slot, joined to its 1-based contribution', () => {
    const detail = epochDetail(store, liveEpochId)!
    expect(detail.committee).toHaveLength(64)
    // Slots count from 0, participant indices from 1.
    expect(detail.committee[0]).toMatchObject({ slot: 0, participantIndex: 1 })
    expect(detail.committee[63]).toMatchObject({ slot: 63, participantIndex: 64 })
    for (const row of detail.committee) {
      const contribution = store.contributions[`${detail.epoch.id}:${row.participantIndex}`]
      expect(row.contributed).toBe(contribution != null)
      if (contribution) expect(contribution.contributor).toBe(row.operator)
    }
    const contributed = detail.committee.filter((row) => row.contributed)
    expect(contributed.length).toBe(detail.epoch.contributions.length)
    for (const row of contributed) {
      expect(row.contributionBlock).toBeGreaterThan(row.claimBlock!)
      expect(row.commitmentsHash).toMatch(/^0x[0-9a-f]{64}$/)
      expect(row.shareCommitmentHash).toMatch(/^0x[0-9a-f]{64}$/)
    }
    expect(detail.committee.every((row) => row.operator !== undefined)).toBe(true)
    expect(detail.contributions).toHaveLength(detail.epoch.contributions.length)
  })

  it('carries windows, applications, finalization and the epoch event log', () => {
    const detail = epochDetail(store, liveEpochId)!
    expect(detail.windows.committeeSelectionDeadline).toBe(detail.epoch.startBlock + 25)
    expect(detail.windows.keyAssemblyDeadline).toBe(detail.epoch.startBlock + 50)
    expect(detail.windows.liveNotBefore).toBe(detail.epoch.startBlock + 55)
    expect(detail.windows.endBlock).toBe(detail.epoch.startBlock + 300)
    expect(detail.applications).toHaveLength(2)
    expect(detail.applications[0].ciphertexts).toBe(8)
    expect(detail.finalization?.gasUsed).toBe(1_112_337)
    expect(detail.collectivePublicKey).not.toBeNull()
    expect(detail.events.length).toBeGreaterThan(64)
    expect(detail.events.every((event) => event.epoch === detail.epoch.id)).toBe(true)
  })

  it('reports the aborted and assembling epochs honestly', () => {
    const aborted = epochDetail(store, abortedEpochId)!
    expect(aborted.epoch.status).toBe('aborted')
    expect(aborted.finalization).toBeNull()
    expect(aborted.lottery.claims.length).toBeLessThan(64)
    expect(aborted.applications).toHaveLength(0)

    const assembling = epochDetail(store, assemblingEpochId)!
    expect(assembling.epoch.status).toBe('key-assembly')
    expect(assembling.committee).toHaveLength(64)
    expect(assembling.contributions.length).toBeLessThan(40)
    expect(assembling.collectivePublicKey).toBeNull()
  })

  it('returns null for an unknown epoch', () => {
    expect(epochDetail(store, '0x2f1105e900000000000000ff')).toBeNull()
  })
})

describe('lottery maths', () => {
  it('converts τ to a fraction and back to the registry size', () => {
    expect(fractionOfHashSpace(0n)).toBe(0)
    expect(fractionOfHashSpace(1n << 255n)).toBeCloseTo(0.5, 6)
    expect(fractionOfHashSpace(1n << 256n)).toBe(1)
    const epoch = store.epochs[liveEpochId]
    expect(recoverRegistrySnapshot(epoch, 64)).toBe(276)
  })
})

describe('operators', () => {
  it('lists only registry members, with per-operator work counters', () => {
    const rows = operatorRows(store)
    expect(rows).toHaveLength(300)
    const totalClaims = rows.reduce((sum, row) => sum + row.claims, 0)
    expect(totalClaims).toBe(7 * 64 + 40)
    const totalContributions = rows.reduce((sum, row) => sum + row.contributions, 0)
    expect(totalContributions).toBe(networkStats(store).contributions)
    expect(rows.reduce((sum, row) => sum + row.finalizations, 0)).toBe(6)
    expect(rows.reduce((sum, row) => sum + row.combines, 0)).toBe(68)
  })

  it('shows participation as contributions/claims, and "—" without claims', () => {
    const rows = operatorRows(store)
    const served = rows.filter((row) => row.claims > 0)
    expect(served.length).toBeGreaterThan(100)
    for (const row of served) {
      expect(row.participation).toBeGreaterThanOrEqual(0)
      expect(row.participation).toBeLessThanOrEqual(1)
      expect(row.participation).toBeCloseTo(row.contributions / row.claims, 9)
    }
    // The aborted and still-assembling epochs leave claims without a
    // contribution, so both ends of the range are represented.
    expect(served.some((row) => row.participation === 1)).toBe(true)
    expect(served.some((row) => (row.participation ?? 1) < 1)).toBe(true)
    const idle = rows.filter((row) => row.claims === 0)
    expect(idle.length).toBeGreaterThan(0)
    for (const row of idle) expect(row.participation).toBeNull()
    expect(formatParticipation(null)).toBe('—')
    expect(formatParticipation(0.9375, 1)).toBe('93.8%')
  })

  it('gives an operator its per-epoch history', () => {
    const busiest = [...operatorRows(store)].sort((a, b) => b.claims - a.claims)[0]
    const detail = operatorDetail(store, busiest.address)!
    expect(detail.row.address).toBe(busiest.address)
    expect(detail.history.length).toBeGreaterThan(0)
    // Newest epoch first.
    expect(detail.history[0].nonce).toBeGreaterThanOrEqual(detail.history[detail.history.length - 1].nonce)
    const claimed = detail.history.filter((entry) => entry.claimed)
    expect(claimed).toHaveLength(busiest.claims)
    expect(detail.events.length).toBeGreaterThan(0)
    expect(detail.events.every((event) => event.actor === busiest.address)).toBe(true)
    expect(operatorDetail(store, '0x0000000000000000000000000000000000009999')).toBeNull()
  })
})

describe('applications', () => {
  it('lists applications newest first with pipeline counters', () => {
    const rows = applicationRows(store)
    expect(rows).toHaveLength(12)
    for (const row of rows) {
      expect(row.ciphertexts).toBe(8)
      expect(row.decrypted).toBeLessThanOrEqual(row.ciphertexts)
      expect(row.sharesPublished).toBe(6)
      expect(row.maxCiphertexts).toBe(8)
      expect(row.authorizedSubmitter).toMatch(/^0x[0-9a-f]{40}$/)
    }
    expect(rows[0].createdBlock).toBeGreaterThanOrEqual(rows[rows.length - 1].createdBlock)
  })

  it('describes every ciphertext of an application', () => {
    const { epoch, aid } = meta.applications[0]
    const detail = applicationDetail(store, epoch, aid)!
    expect(detail.ciphertexts).toHaveLength(8)
    expect(detail.summary.total).toBe(8)
    expect(detail.summary.withShare).toBe(6)

    const combined = detail.ciphertexts.find((row) => row.state === 'combined')!
    expect(combined.combined.plaintext).toBeGreaterThan(0n)
    expect(combined.combined.gasUsed).toBe(430_432)
    expect(combined.combined.by).toMatch(/^0x[0-9a-f]{40}$/)
    expect(combined.partialCount).toBeGreaterThanOrEqual(combined.threshold)

    // Every fourth ciphertext has no organizer share and cannot combine.
    const stuck = detail.ciphertexts.filter((row) => !row.share.present)
    expect(stuck).toHaveLength(2)
    for (const row of stuck) {
      expect(row.state).toBe('awaiting-share')
      expect(row.combined.done).toBe(false)
    }
    // Partials carry the wave they landed in and their 0-based slot.
    expect(new Set(combined.partials.map((partial) => partial.wave))).toEqual(new Set([0]))
    for (const partial of combined.partials) {
      expect(partial.slot).toBe(partial.participantIndex - 1)
      expect(store.epochs[combined.epoch].committee[partial.slot]).toBe(partial.participant)
    }
    const late = detail.ciphertexts.find((row) => row.partials.some((partial) => partial.wave === 1))
    expect(late).toBeDefined()
  })

  it('returns null for an unknown application', () => {
    expect(applicationDetail(store, meta.epochIds[0], '0x' + '11'.repeat(32))).toBeNull()
  })
})

describe('partialMatrix', () => {
  it('is committee × ciphertexts for one application (64 × 8)', () => {
    const { epoch, aid } = meta.applications.find((a) => a.epoch === liveEpochId)!
    const matrix = partialMatrix(store, epoch, aid)!
    expect(matrix.rows).toHaveLength(64)
    expect(matrix.columns).toHaveLength(8)
    expect(matrix.cells).toHaveLength(64)
    expect(matrix.cells[0]).toHaveLength(8)
    expect(matrix.threshold).toBe(33)
    expect(matrix.staggerBlocks).toBe(3)

    // Every column holds at least t partials, and each cell knows its wave.
    for (const column of matrix.columns) {
      const filled = matrix.cells.map((row) => row[column.column]).filter(Boolean)
      expect(filled.length).toBe(column.partials)
      expect(filled.length).toBeGreaterThanOrEqual(33)
      for (const cell of filled) {
        expect(cell!.participantIndex).toBe(cell!.slot + 1)
        expect(cell!.wave).toBeGreaterThanOrEqual(0)
        expect(cell!.block).toBeGreaterThanOrEqual(column.submitBlock)
        expect(cell!.wave).toBe(Math.floor((cell!.block! - column.submitBlock) / 3))
      }
    }
    // Row totals agree with the cells, and rows are addressed by slot.
    matrix.rows.forEach((row, i) => {
      expect(row.slot).toBe(i)
      expect(row.participantIndex).toBe(i + 1)
      expect(row.operator).toBe(store.epochs[epoch].committee[i])
      expect(row.partials).toBe(matrix.cells[i].filter(Boolean).length)
    })
  })

  it('spans every application of the epoch when no aid is given', () => {
    const matrix = partialMatrix(store, liveEpochId)!
    expect(matrix.columns).toHaveLength(16)
    expect(new Set(matrix.columns.map((column) => column.aid)).size).toBe(2)
  })

  it('is empty but well-formed for an epoch with no ciphertexts', () => {
    const matrix = partialMatrix(store, assemblingEpochId)!
    expect(matrix.rows).toHaveLength(64)
    expect(matrix.columns).toHaveLength(0)
    expect(partialMatrix(store, '0x2f1105e900000000000000ff')).toBeNull()
  })
})

describe('activity and the event feed', () => {
  it('buckets work per epoch, oldest first', () => {
    const buckets = activity(store, 30)
    expect(buckets).toHaveLength(8)
    expect(buckets[0].nonce).toBe(1)
    expect(buckets[buckets.length - 1].nonce).toBe(8)
    const first = buckets[0]
    expect(first.claims).toBe(64)
    expect(first.ciphertexts).toBe(16)
    expect(first.partials).toBeGreaterThan(500)
    expect(activity(store, 3)).toHaveLength(3)
  })

  it('returns the newest events with their transactions', () => {
    const feed = eventFeed(store, 20)
    expect(feed).toHaveLength(20)
    for (let i = 1; i < feed.length; i++) {
      expect(feed[i - 1].block).toBeGreaterThanOrEqual(feed[i].block)
    }
    expect(feed[0].event).toBe(store.events[store.events.length - 1])
    expect(eventFeed(store, 5)).toHaveLength(5)
  })
})

describe('searchStore', () => {
  it('resolves epoch nonces, ids, addresses, aids and transactions', () => {
    const byNonce = searchStore(store, '3')
    expect(byNonce[0]).toMatchObject({ kind: 'epoch', href: `/epochs/${meta.epochIds[2]}` })

    const byId = searchStore(store, liveEpochId)
    expect(byId[0]).toMatchObject({ kind: 'epoch', id: liveEpochId })

    const operator = operatorRows(store)[0].address
    const byAddress = searchStore(store, operator)
    expect(byAddress[0]).toMatchObject({ kind: 'operator', href: `/operators/${operator}` })

    const { aid, epoch } = meta.applications[0]
    const byAid = searchStore(store, aid)
    expect(byAid[0]).toMatchObject({ kind: 'application', href: `/applications/${epoch}/${aid}` })

    const tx = store.epochs[liveEpochId].finalization!.tx!
    const byTx = searchStore(store, tx)
    // No in-app transaction page: a transaction routes to the block explorer.
    expect(byTx[0]).toMatchObject({
      kind: 'transaction',
      id: tx,
      external: true,
      href: `https://sepolia.etherscan.io/tx/${tx}`,
    })
    expect(byId[0].external).toBe(false)

    expect(searchStore(store, '')).toHaveLength(0)
    expect(searchStore(store, 'not-hex')).toHaveLength(0)
    expect(searchStore(store, liveEpochId.slice(0, 8), 5).length).toBeGreaterThan(0)
  })
})

describe('performance at fixture scale', () => {
  it('runs every selector well inside a frame budget', () => {
    const fresh = bumpStore(store) // defeat the memo cache
    const timings: Record<string, number> = {}
    const time = (name: string, run: () => unknown): void => {
      const started = performance.now()
      run()
      timings[name] = performance.now() - started
    }
    time('networkStats', () => networkStats(fresh))
    time('epochRows', () => epochRows(fresh))
    time('operatorRows', () => operatorRows(fresh))
    time('epochDetail', () => epochDetail(fresh, liveEpochId))
    time('applicationRows', () => applicationRows(fresh))
    time('applicationDetail', () =>
      applicationDetail(fresh, meta.applications[0].epoch, meta.applications[0].aid),
    )
    time('partialMatrix', () => partialMatrix(fresh, liveEpochId))
    time('activity', () => activity(fresh, 30))
    time('eventFeed', () => eventFeed(fresh, 20))
    time('searchStore', () => searchStore(fresh, '0x2f'))
    for (const [name, ms] of Object.entries(timings)) {
      expect(ms, `${name} took ${ms.toFixed(1)}ms`).toBeLessThan(50)
    }
  })
})
