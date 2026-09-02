import { describe, expect, it } from 'vitest'
import type { OperatorRow } from '~indexer/selectors'
import { operatorStatusSlices, operatorWorkChart } from './charts'
import { filterOperators, hexPrefix } from './filter'

function row(overrides: Partial<OperatorRow> & { address: string }): OperatorRow {
  return {
    pubKey: null,
    status: 'active',
    registeredAtBlock: 100,
    lastActiveBlock: 200,
    epochsServed: 0,
    claims: 0,
    contributions: 0,
    partials: 0,
    finalizations: 0,
    combines: 0,
    participation: null,
    idleBlocks: 10,
    reapable: false,
    ...overrides,
    address: overrides.address as OperatorRow['address'],
  } as OperatorRow
}

const rows: OperatorRow[] = [
  row({ address: '0xaa00000000000000000000000000000000000001', contributions: 9, partials: 4 }),
  row({ address: '0xbb00000000000000000000000000000000000002', contributions: 5, partials: 20, combines: 3 }),
  row({ address: '0xcc00000000000000000000000000000000000003', contributions: 1, finalizations: 2 }),
  row({ address: '0xdd00000000000000000000000000000000000004', status: 'inactive', idleBlocks: null }),
  row({ address: '0xee00000000000000000000000000000000000005', reapable: true, partials: 1 }),
]

describe('operatorWorkChart', () => {
  it('orders by contributions and groups the tail into one column', () => {
    const chart = operatorWorkChart(rows, 2)
    expect(chart.data).toHaveLength(3)
    expect(chart.addresses.slice(0, 2)).toEqual([rows[0].address, rows[1].address])
    expect(chart.grouped).toBe(3)
    // The grouped column is not clickable and sums everything below the cut.
    expect(chart.addresses[2]).toBeNull()
    expect(chart.data[2].values).toEqual({ contributions: 1, partials: 1, finalizations: 2, combines: 0 })
    expect(chart.data[2].label).toBe('+3')
  })

  it('omits the others column when everything fits', () => {
    const chart = operatorWorkChart(rows, 10)
    expect(chart.grouped).toBe(0)
    expect(chart.data).toHaveLength(rows.length)
    expect(chart.addresses.every((address) => address !== null)).toBe(true)
  })

  it('keeps the totals intact across the cut', () => {
    const total = (key: string, chart = operatorWorkChart(rows, 2)) =>
      chart.data.reduce((sum, datum) => sum + (datum.values[key] ?? 0), 0)
    expect(total('contributions')).toBe(15)
    expect(total('partials')).toBe(25)
    expect(total('combines')).toBe(3)
  })
})

describe('operatorStatusSlices', () => {
  it('carves the reapable operators out of the active set', () => {
    const slices = operatorStatusSlices(rows)
    expect(slices.map((slice) => [slice.label, slice.value])).toEqual([
      ['active', 3],
      ['idle past window', 1],
      ['inactive', 1],
    ])
    expect(slices.reduce((sum, slice) => sum + slice.value, 0)).toBe(rows.length)
  })
})

describe('hexPrefix', () => {
  it('accepts a bare or prefixed hex string in any casing', () => {
    expect(hexPrefix('0xAB')).toBe('ab')
    expect(hexPrefix('  ab ')).toBe('ab')
  })

  it('rejects anything that is not hex', () => {
    expect(hexPrefix('')).toBeNull()
    expect(hexPrefix('0x')).toBeNull()
    expect(hexPrefix('zz')).toBeNull()
  })
})

describe('filterOperators', () => {
  it('matches addresses by prefix, never by substring', () => {
    expect(filterOperators(rows, { query: 'bb' })).toHaveLength(1)
    // "00" appears inside every address but starts none of them.
    expect(filterOperators(rows, { query: '00' })).toHaveLength(0)
  })

  it('filters by status, with idle carved out of active', () => {
    expect(filterOperators(rows, { status: 'active' })).toHaveLength(4)
    expect(filterOperators(rows, { status: 'inactive' })).toHaveLength(1)
    expect(filterOperators(rows, { status: 'idle' }).map((r) => r.address)).toEqual([rows[4].address])
  })

  it('returns everything with an empty filter', () => {
    expect(filterOperators(rows)).toHaveLength(rows.length)
  })
})
