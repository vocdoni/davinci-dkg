// Filtering and header totals for the applications table. Pure, so the
// "secrets kept" definition is stated once and tested.

import type { ApplicationRow } from '~indexer/selectors'
import type { AppModeName } from '~indexer/types'
import { hexPrefix } from '../operators/filter'

export interface ApplicationFilter {
  /** Matches the application id, the organizer or any allow-listed submitter by prefix. */
  query?: string
  /** Epoch id, or `all`. */
  epoch?: string
  /** Application mode, or `all`. */
  mode?: AppModeName | 'all'
}

export function filterApplications(rows: readonly ApplicationRow[], filter: ApplicationFilter = {}): ApplicationRow[] {
  const { query = '', epoch = 'all', mode = 'all' } = filter
  const prefix = hexPrefix(query)
  return rows.filter((row) => {
    if (epoch !== 'all' && row.epoch.toLowerCase() !== epoch.toLowerCase()) return false
    if (mode !== 'all' && row.mode !== mode) return false
    if (!prefix) return true
    const bare = (value: string) => value.toLowerCase().replace(/^0x/, '')
    return (
      bare(row.aid).startsWith(prefix) ||
      bare(row.creator).startsWith(prefix) ||
      (row.submitters ?? []).some((submitter) => bare(submitter).startsWith(prefix))
    )
  })
}

export interface ApplicationSummary {
  applications: number
  ciphertexts: number
  decrypted: number
  /**
   * Organizer-locked applications whose `sk_org` is not on chain yet. Every
   * ciphertext under them is stuck whatever the committee does: the combine
   * proof needs the secret, and only the organizer can reveal it.
   */
  locked: number
}

export function summarizeApplications(rows: readonly ApplicationRow[]): ApplicationSummary {
  let ciphertexts = 0
  let decrypted = 0
  let locked = 0
  for (const row of rows) {
    ciphertexts += row.ciphertexts
    decrypted += row.decrypted
    if (!row.unlocked) locked += 1
  }
  return { applications: rows.length, ciphertexts, decrypted, locked }
}
