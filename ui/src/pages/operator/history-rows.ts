// Per-epoch history rows for one operator.
//
// The selector gives the counters; the transactions behind them only exist in
// the event log, so this joins the two. Pure, so the join and the trend series
// are tested without a DOM.

import type { OperatorEpochHistory } from '~indexer/selectors'
import { epochKey, type Hex, type IndexedEvent } from '~indexer/types'

export interface OperatorHistoryRow extends OperatorEpochHistory {
  claimTx: Hex | null
  claimBlock: number | null
  contributionTx: Hex | null
  contributionGas: number | null
}

export type GasLookup = (tx: Hex | null) => number | null

/** History plus the claim and contribution transactions, newest epoch first. */
export function operatorHistoryRows(
  history: readonly OperatorEpochHistory[],
  events: readonly IndexedEvent[],
  gasOf: GasLookup = () => null
): OperatorHistoryRow[] {
  const claims = new Map<string, { tx: Hex | null; block: number }>()
  const contributions = new Map<string, Hex | null>()
  for (const event of events) {
    if (!event.epoch) continue
    const key = epochKey(event.epoch)
    if (event.name === 'SlotClaimed' && !claims.has(key)) claims.set(key, { tx: event.tx, block: event.block })
    if (event.name === 'ContributionSubmitted' && !contributions.has(key)) contributions.set(key, event.tx)
  }

  return [...history]
    .sort((a, b) => b.nonce - a.nonce)
    .map((entry) => {
      const key = epochKey(entry.epoch)
      const claim = claims.get(key) ?? null
      const contributionTx = contributions.get(key) ?? null
      return {
        ...entry,
        claimTx: claim?.tx ?? null,
        claimBlock: claim?.block ?? null,
        contributionTx,
        contributionGas: gasOf(contributionTx),
      }
    })
}

/**
 * Participation as it evolved: cumulative contributions over cumulative claims
 * per epoch, oldest first. Cumulative rather than per-epoch because per-epoch
 * participation is a boolean and a spark of zeros and hundreds says nothing.
 */
export function participationTrend(rows: readonly OperatorHistoryRow[]): number[] {
  let claims = 0
  let contributions = 0
  return [...rows]
    .sort((a, b) => a.nonce - b.nonce)
    .map((row) => {
      if (row.claimed) claims += 1
      if (row.contributed) contributions += 1
      return claims === 0 ? 0 : Math.round((contributions / claims) * 100)
    })
}

/** Partials published per epoch, oldest first. */
export function partialsTrend(rows: readonly OperatorHistoryRow[]): number[] {
  return [...rows].sort((a, b) => a.nonce - b.nonce).map((row) => row.partials)
}
