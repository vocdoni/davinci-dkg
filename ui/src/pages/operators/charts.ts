// Chart data for the operators page.
//
// Pure on purpose: the "top N plus one grouped column" rule and the status
// split are the two things that decide whether a 300-operator registry reads
// as a distribution or as a texture, so they are unit-tested without a DOM.

import type { OperatorRow } from '~indexer/selectors'
import { CHART_COLORS, type BarDatum, type BarSeries, type DonutSlice } from '~kit/charts'
import { shortAddress } from '~lib/format'

/** The four things an operator is paid to do, in pipeline order. */
export const WORK_SERIES: BarSeries[] = [
  { key: 'contributions', label: 'contributions', color: CHART_COLORS.emerald },
  { key: 'partials', label: 'partials', color: CHART_COLORS.teal },
  { key: 'finalizations', label: 'finalizations', color: CHART_COLORS.slate },
  { key: 'combines', label: 'combines', color: CHART_COLORS.warmGray },
]

export interface WorkChart {
  data: BarDatum[]
  /** Address behind each bar; null for the trailing "others" column. */
  addresses: Array<string | null>
  /** How many operators the last column folds together. */
  grouped: number
}

function work(row: OperatorRow): number {
  return row.contributions + row.partials + row.finalizations + row.combines
}

/**
 * Work per operator: the busiest `top` operators as their own bars, everything
 * below the cut summed into one column so nothing is dropped from the totals.
 * Ordered by contributions (the scarce work — only a committee member can do
 * it), then partials, then address for a stable tie-break.
 */
export function operatorWorkChart(rows: readonly OperatorRow[], top = 32): WorkChart {
  const sorted = [...rows].sort(
    (a, b) =>
      b.contributions - a.contributions ||
      b.partials - a.partials ||
      work(b) - work(a) ||
      a.address.localeCompare(b.address)
  )
  const head = sorted.slice(0, top)
  const tail = sorted.slice(top)

  const data: BarDatum[] = head.map((row) => ({
    label: shortAddress(row.address),
    values: {
      contributions: row.contributions,
      partials: row.partials,
      finalizations: row.finalizations,
      combines: row.combines,
    },
    note: `${row.address} · ${row.epochsServed} epochs served`,
  }))
  const addresses: Array<string | null> = head.map((row) => row.address)

  if (tail.length > 0) {
    const sum = (pick: (row: OperatorRow) => number) => tail.reduce((n, row) => n + pick(row), 0)
    data.push({
      label: `+${tail.length}`,
      values: {
        contributions: sum((row) => row.contributions),
        partials: sum((row) => row.partials),
        finalizations: sum((row) => row.finalizations),
        combines: sum((row) => row.combines),
      },
      note: `${tail.length} operators outside the top ${top}, summed`,
    })
    addresses.push(null)
  }

  return { data, addresses, grouped: tail.length }
}

/**
 * Status donut. "Idle past the window" is carved out of the active set rather
 * than shown as a fourth state: on chain those operators are still active, but
 * anyone may reap them, which is the thing worth seeing.
 */
export function operatorStatusSlices(rows: readonly OperatorRow[]): DonutSlice[] {
  let active = 0
  let reapable = 0
  let inactive = 0
  for (const row of rows) {
    if (row.status === 'active') {
      if (row.reapable) reapable += 1
      else active += 1
    } else inactive += 1
  }
  return [
    { label: 'active', value: active, color: CHART_COLORS.emerald },
    { label: 'idle past window', value: reapable, color: CHART_COLORS.amber },
    { label: 'inactive', value: inactive, color: CHART_COLORS.warmGray },
  ]
}
