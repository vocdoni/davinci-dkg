// Search and status filtering for the operators table.
//
// Addresses are matched by *prefix*, never by substring: a substring match on
// a hex string turns a single typed digit into "everything", which is worse
// than no filter at all.

import type { OperatorRow } from '~indexer/selectors'

export type OperatorStatusFilter = 'all' | 'active' | 'inactive' | 'idle'

export const STATUS_OPTIONS: Array<{ value: OperatorStatusFilter; label: string }> = [
  { value: 'all', label: 'All operators' },
  { value: 'active', label: 'Active' },
  { value: 'inactive', label: 'Inactive' },
  { value: 'idle', label: 'Idle past the window' },
]

export interface OperatorFilter {
  query?: string
  status?: OperatorStatusFilter
}

/**
 * Normalised hex prefix: accepts `0xabc`, `abc` and any casing, and rejects
 * anything that is not hex. Shared with the applications filter (aids and
 * addresses are both matched by prefix).
 */
export function hexPrefix(query: string): string | null {
  const q = query.trim().toLowerCase()
  if (q === '') return null
  const bare = q.startsWith('0x') ? q.slice(2) : q
  if (bare === '' || !/^[0-9a-f]+$/.test(bare)) return null
  return bare
}

export function filterOperators(rows: readonly OperatorRow[], filter: OperatorFilter = {}): OperatorRow[] {
  const { query = '', status = 'all' } = filter
  const prefix = hexPrefix(query)
  return rows.filter((row) => {
    if (status === 'active' && row.status !== 'active') return false
    if (status === 'inactive' && row.status !== 'inactive') return false
    if (status === 'idle' && !row.reapable) return false
    if (prefix && !row.address.toLowerCase().slice(2).startsWith(prefix)) return false
    return true
  })
}
