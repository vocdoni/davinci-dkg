// Filtering and header totals for the applications table. Pure, so the
// "pending shares" definition is stated once and tested.

import type { ApplicationRow } from '~indexer/selectors'
import { hexPrefix } from '../operators/filter'

export interface ApplicationFilter {
  /** Matches the application id, the organizer or the authorized submitter by prefix. */
  query?: string
  /** Epoch id, or `all`. */
  epoch?: string
}

export function filterApplications(rows: readonly ApplicationRow[], filter: ApplicationFilter = {}): ApplicationRow[] {
  const { query = '', epoch = 'all' } = filter
  const prefix = hexPrefix(query)
  return rows.filter((row) => {
    if (epoch !== 'all' && row.epoch.toLowerCase() !== epoch.toLowerCase()) return false
    if (!prefix) return true
    const bare = (value: string) => value.toLowerCase().replace(/^0x/, '')
    return (
      bare(row.aid).startsWith(prefix) ||
      bare(row.creator).startsWith(prefix) ||
      (row.authorizedSubmitter != null && bare(row.authorizedSubmitter).startsWith(prefix))
    )
  })
}

export interface ApplicationSummary {
  applications: number
  ciphertexts: number
  decrypted: number
  /**
   * Ciphertexts still without an organizer share. Every one of them is stuck
   * whatever the committee does: `combineDecryption` needs `t` partials *and*
   * `Δ = sk_org·C1`.
   */
  pendingShares: number
}

export function summarizeApplications(rows: readonly ApplicationRow[]): ApplicationSummary {
  let ciphertexts = 0
  let decrypted = 0
  let shares = 0
  for (const row of rows) {
    ciphertexts += row.ciphertexts
    decrypted += row.decrypted
    shares += row.sharesPublished
  }
  return {
    applications: rows.length,
    ciphertexts,
    decrypted,
    pendingShares: Math.max(0, ciphertexts - shares),
  }
}
