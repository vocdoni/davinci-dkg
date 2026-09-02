import { describe, expect, it } from 'vitest'
import {
  aggregateOperatorStats,
  formatParticipation,
  participationScore,
  summarizeOperatorStats,
  type OperatorStatsInput,
} from './operator-stats'

const A = '0xAAaaAAaaAAaaAAaaAAaaAAaaAAaaAAaaAAaaAAaa'
const B = '0xBbbbBBbbBBbbBBbbBBbbBBbbBBbbBBbbBBbbBBbb'
const C = '0xCcccCCccCCccCCccCCccCCccCCccCCccCCccCCcc'

function node(operator: string, status = 1, last = 100n, reg = 10n) {
  return { operator, status, lastActiveBlock: last, registeredAtBlock: reg }
}

function input(overrides: Partial<OperatorStatsInput> = {}): OperatorStatsInput {
  return {
    nodes: [],
    claims: [],
    contributions: [],
    partials: [],
    finalizations: [],
    combines: [],
    senders: new Map(),
    ...overrides,
  }
}

describe('participationScore', () => {
  it('is contributions over claims in percent', () => {
    expect(participationScore(3, 4)).toBe(75)
    expect(participationScore(4, 4)).toBe(100)
    expect(participationScore(0, 3)).toBe(0)
  })

  it('is null — not 100 — when the operator never claimed a slot', () => {
    expect(participationScore(0, 0)).toBeNull()
    expect(formatParticipation(participationScore(0, 0))).toBe('—')
  })

  it('rounds to whole percent and does not clamp a truncated history', () => {
    expect(participationScore(1, 3)).toBe(33)
    expect(participationScore(2, 3)).toBe(67)
    // More contributions than claims can only happen when the scan window
    // starts after some claims; the anomaly is surfaced, not hidden.
    expect(participationScore(3, 2)).toBe(150)
  })
})

describe('aggregateOperatorStats', () => {
  it('counts each self-attributing event against its operator', () => {
    const rows = aggregateOperatorStats(
      input({
        nodes: [node(A), node(B)],
        claims: [{ claimer: A }, { claimer: A }, { claimer: B }],
        contributions: [{ contributor: A }, { contributor: A }],
        partials: [{ participant: B }, { participant: B }, { participant: B }],
      }),
    )
    const a = rows.find((r) => r.operator === A)!
    const b = rows.find((r) => r.operator === B)!
    expect(a).toMatchObject({ claims: 2, contributions: 2, partials: 0, participation: 100 })
    expect(b).toMatchObject({ claims: 1, contributions: 0, partials: 3, participation: 0 })
  })

  it('attributes finalizations and combines through the transaction sender', () => {
    const rows = aggregateOperatorStats(
      input({
        nodes: [node(A), node(B)],
        finalizations: [{ transactionHash: '0xF1' }, { transactionHash: '0xF2' }],
        combines: [{ transactionHash: '0xC1' }],
        // The map is keyed by lower-cased hash, as the SDK returns it.
        senders: new Map([
          ['0xf1', A],
          ['0xf2', B],
          ['0xc1', A],
        ]),
      }),
    )
    expect(rows.find((r) => r.operator === A)).toMatchObject({ finalizations: 1, combines: 1 })
    expect(rows.find((r) => r.operator === B)).toMatchObject({ finalizations: 1, combines: 0 })
  })

  it('drops events whose transaction sender could not be resolved', () => {
    const rows = aggregateOperatorStats(
      input({
        nodes: [node(A)],
        finalizations: [{ transactionHash: null }, { transactionHash: '0xdead' }],
        senders: new Map(),
      }),
    )
    expect(rows).toHaveLength(1)
    expect(rows[0].finalizations).toBe(0)
  })

  it('matches addresses case-insensitively but displays the registry casing', () => {
    const rows = aggregateOperatorStats(
      input({
        nodes: [node(A)],
        claims: [{ claimer: A.toLowerCase() }],
        contributions: [{ contributor: A.toUpperCase().replace('0X', '0x') }],
      }),
    )
    expect(rows).toHaveLength(1)
    expect(rows[0].operator).toBe(A)
    expect(rows[0]).toMatchObject({ claims: 1, contributions: 1 })
  })

  it('keeps registry liveness on the row and gives idle operators a row', () => {
    const rows = aggregateOperatorStats(
      input({ nodes: [node(A, 2, 42n, 7n)] }),
    )
    expect(rows[0]).toMatchObject({
      status: 2,
      lastActiveBlock: 42n,
      registeredAtBlock: 7n,
      claims: 0,
      participation: null,
    })
  })

  it('includes an operator seen only in events, flagged as unregistered', () => {
    const rows = aggregateOperatorStats(
      input({ nodes: [node(A)], contributions: [{ contributor: C }] }),
    )
    const c = rows.find((r) => r.operator === C)!
    expect(c.status).toBe(0)
    expect(c.registeredAtBlock).toBeNull()
    expect(c.contributions).toBe(1)
  })

  it('sorts by contributions, then partials, then claims, then address', () => {
    const rows = aggregateOperatorStats(
      input({
        nodes: [node(A), node(B), node(C)],
        contributions: [{ contributor: B }, { contributor: B }, { contributor: C }],
        partials: [{ participant: C }, { participant: A }],
      }),
    )
    expect(rows.map((r) => r.operator)).toEqual([B, C, A])
  })
})

describe('summarizeOperatorStats', () => {
  it('totals the columns and scores the deployment as a whole', () => {
    const rows = aggregateOperatorStats(
      input({
        nodes: [node(A), node(B, 2)],
        claims: [{ claimer: A }, { claimer: A }, { claimer: B }, { claimer: B }],
        contributions: [{ contributor: A }, { contributor: A }, { contributor: B }],
        partials: [{ participant: A }],
        finalizations: [{ transactionHash: '0xf1' }],
        senders: new Map([['0xf1', A]]),
      }),
    )
    expect(summarizeOperatorStats(rows)).toEqual({
      operators: 2,
      activeOperators: 1,
      claims: 4,
      contributions: 3,
      partials: 1,
      finalizations: 1,
      combines: 0,
      participation: 75,
    })
  })

  it('reports a null score for a deployment with no claims yet', () => {
    expect(summarizeOperatorStats([]).participation).toBeNull()
  })
})
