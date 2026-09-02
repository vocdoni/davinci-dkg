// Per-operator aggregation of the DKG event log.
//
// Everything here is pure: the query layer does the RPC work and hands the raw
// events in, this file turns them into the leaderboard rows. That split is what
// makes the numbers testable — see operator-stats.test.ts, which feeds it
// synthetic events rather than a chain.
//
// Attribution rules, and why they differ per event:
//
//   SlotClaimed                  → `claimer`      (indexed, self-attributing)
//   ContributionSubmitted        → `contributor`  (indexed, self-attributing)
//   PartialDecryptionSubmitted   → `participant`  (indexed, self-attributing)
//   EpochLive                    → transaction sender. The event says an epoch
//                                  was finalized, not by whom; exactly one is
//                                  emitted per epoch, and whoever paid for the
//                                  finalize proof is its `tx.from`.
//   DecryptionCombined           → transaction sender, same reasoning (one per
//                                  ciphertext).

/** Minimal shape of a registry record; matches the SDK's `NodeKey`. */
export interface OperatorRegistryRecord {
  operator: string
  status: number
  lastActiveBlock: bigint
  registeredAtBlock: bigint
}

export interface OperatorStatsInput {
  /** Registry roster. Operators with no events still get a row. */
  nodes: readonly OperatorRegistryRecord[]
  claims: readonly { claimer: string }[]
  contributions: readonly { contributor: string }[]
  partials: readonly { participant: string }[]
  /** `EpochLive` events — attributed through `senders`. */
  finalizations: readonly { transactionHash: string | null }[]
  /** `DecryptionCombined` events — attributed through `senders`. */
  combines: readonly { transactionHash: string | null }[]
  /** Lower-cased transaction hash → sender address. */
  senders: ReadonlyMap<string, string>
}

export interface OperatorStats {
  /** Display form of the address (checksummed when the registry knows it). */
  operator: string
  /** `NodeStatus` value; 0 when the operator is not (or no longer) registered. */
  status: number
  /** Null for an address seen in events but absent from the registry. */
  lastActiveBlock: bigint | null
  registeredAtBlock: bigint | null
  claims: number
  contributions: number
  partials: number
  finalizations: number
  combines: number
  /**
   * Participation score — see {@link participationScore}. Percent of claimed
   * slots that turned into an accepted contribution, or null when the operator
   * has never claimed a slot.
   */
  participation: number | null
}

export interface OperatorStatsSummary {
  operators: number
  activeOperators: number
  claims: number
  contributions: number
  partials: number
  finalizations: number
  combines: number
  /** Aggregate contributions/claims across every operator, or null. */
  participation: number | null
}

/**
 * The score, defined once so the table header, the tooltip and the tests all
 * mean the same thing:
 *
 *   participation = contributions / slots claimed, as a percentage
 *
 * A slot claimed in the lottery is a commitment to publish a contribution in
 * the key-assembly window; the ratio is how often the operator kept it. It is
 * deliberately NOT "100% when both are zero" — an operator that never won the
 * lottery has no track record, and rendering that as a perfect score would put
 * it above operators that actually did the work. Callers show `—` for null.
 *
 * The rounding is to whole percent; values above 100 are possible only when
 * the scan window starts after some claims (a truncated history) and are left
 * as they are rather than clamped, so the anomaly stays visible.
 */
export function participationScore(contributions: number, claims: number): number | null {
  if (claims <= 0) return null
  return Math.round((contributions / claims) * 100)
}

/** Formats a score for display; `—` when there is no track record. */
export function formatParticipation(score: number | null): string {
  return score == null ? '—' : `${score}%`
}

export function aggregateOperatorStats(input: OperatorStatsInput): OperatorStats[] {
  const rows = new Map<string, OperatorStats>()

  const touch = (address: string | null | undefined): OperatorStats | null => {
    if (!address) return null
    const key = address.toLowerCase()
    let row = rows.get(key)
    if (!row) {
      row = {
        operator: address,
        status: 0,
        lastActiveBlock: null,
        registeredAtBlock: null,
        claims: 0,
        contributions: 0,
        partials: 0,
        finalizations: 0,
        combines: 0,
        participation: null,
      }
      rows.set(key, row)
    }
    return row
  }

  // Registry first, so its checksummed address and liveness win over whatever
  // casing the event logs happen to carry.
  for (const node of input.nodes) {
    const row = touch(node.operator)
    if (!row) continue
    row.operator = node.operator
    row.status = node.status
    row.lastActiveBlock = node.lastActiveBlock
    row.registeredAtBlock = node.registeredAtBlock
  }

  for (const e of input.claims) {
    const row = touch(e.claimer)
    if (row) row.claims++
  }
  for (const e of input.contributions) {
    const row = touch(e.contributor)
    if (row) row.contributions++
  }
  for (const e of input.partials) {
    const row = touch(e.participant)
    if (row) row.partials++
  }
  for (const e of input.finalizations) {
    const row = touch(senderOf(input.senders, e.transactionHash))
    if (row) row.finalizations++
  }
  for (const e of input.combines) {
    const row = touch(senderOf(input.senders, e.transactionHash))
    if (row) row.combines++
  }

  const out = [...rows.values()]
  for (const row of out) row.participation = participationScore(row.contributions, row.claims)

  // Most useful first: the operators actually carrying the key generation, then
  // the ones only decrypting, then the idle ones. Address breaks ties so the
  // order is stable between refetches.
  out.sort(
    (a, b) =>
      b.contributions - a.contributions ||
      b.partials - a.partials ||
      b.claims - a.claims ||
      a.operator.toLowerCase().localeCompare(b.operator.toLowerCase()),
  )
  return out
}

/** Column totals for the stat cards above the leaderboard. */
export function summarizeOperatorStats(rows: readonly OperatorStats[]): OperatorStatsSummary {
  const sum = (pick: (r: OperatorStats) => number) => rows.reduce((n, r) => n + pick(r), 0)
  const claims = sum((r) => r.claims)
  const contributions = sum((r) => r.contributions)
  return {
    operators: rows.length,
    activeOperators: rows.filter((r) => r.status === 1).length,
    claims,
    contributions,
    partials: sum((r) => r.partials),
    finalizations: sum((r) => r.finalizations),
    combines: sum((r) => r.combines),
    participation: participationScore(contributions, claims),
  }
}

function senderOf(senders: ReadonlyMap<string, string>, hash: string | null): string | null {
  if (!hash) return null
  return senders.get(hash.toLowerCase()) ?? null
}
