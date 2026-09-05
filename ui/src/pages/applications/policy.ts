// Pure wording for the policy facts every application surface shows: the
// mode, the submission policy and the decryption window. Kept apart from the
// cells so the tables, the application page and the tests share one
// definition.

import type { BadgeTone } from '~kit'
import type { ApplicationRow } from '~indexer/selectors'
import type { AppModeName, Address } from '~indexer/types'

export const MODE_TONE: Record<AppModeName, BadgeTone> = {
  'organizer-locked': 'neutral',
  automatic: 'accent',
}

export const MODE_HELP: Record<AppModeName, string> = {
  'organizer-locked':
    'PK_aid = P_j + PK_org: the contract refuses partials and combines until the organizer reveals sk_org, once, on chain',
  automatic: 'PK_aid = P_j, no organizer key: t partials inside the decryption window are all it takes',
}

/** Who may call `submitCiphertext`, as the contract resolves the policy. */
export type SubmissionPolicy =
  | { kind: 'unknown' }
  | { kind: 'open' }
  | { kind: 'registrant' }
  | { kind: 'allow-list'; submitters: Address[] }

export function submissionPolicy(row: Pick<ApplicationRow, 'openSubmission' | 'submitters'>): SubmissionPolicy {
  if (row.openSubmission == null || row.submitters == null) return { kind: 'unknown' }
  if (row.openSubmission) return { kind: 'open' }
  if (row.submitters.length === 0) return { kind: 'registrant' }
  return { kind: 'allow-list', submitters: row.submitters }
}

export function submissionPolicyLabel(policy: SubmissionPolicy): string {
  switch (policy.kind) {
    case 'unknown':
      return '—'
    case 'open':
      return 'open'
    case 'registrant':
      return 'registrant only'
    case 'allow-list':
      return policy.submitters.length === 1 ? 'allow-list' : `allow-list (${policy.submitters.length})`
  }
}

/** Where the wall clock sits relative to a decryption window. */
export type WindowState = 'unbounded' | 'not-yet-open' | 'open' | 'closed'

export interface DecryptionWindow {
  /** Local date (or date-time, see `precision`), or null for an unbounded side. */
  from: string | null
  until: string | null
  /** Always the local date-time, for titles when `from`/`until` are dates only. */
  fromFull: string | null
  untilFull: string | null
  state: WindowState
}

export type WindowPrecision = 'datetime' | 'date'

function localDate(seconds: number, precision: WindowPrecision = 'datetime'): string {
  const date = new Date(seconds * 1000)
  return precision === 'date'
    ? date.toLocaleDateString(undefined, { dateStyle: 'medium' })
    : date.toLocaleString(undefined, { dateStyle: 'medium', timeStyle: 'short' })
}

/**
 * The decryption window (`policy.decryptNotBefore` / `decryptNotAfter`, unix
 * seconds, 0 = unbounded) as text plus whether the chain would accept a
 * partial or a combine right now. Null while the policy is unread.
 */
export function describeDecryptionWindow(
  notBefore: number | null,
  notAfter: number | null,
  nowMs = Date.now(),
  precision: WindowPrecision = 'datetime'
): DecryptionWindow | null {
  if (notBefore == null || notAfter == null) return null
  const from = notBefore === 0 ? null : localDate(notBefore, precision)
  const until = notAfter === 0 ? null : localDate(notAfter, precision)
  const fromFull = notBefore === 0 ? null : localDate(notBefore)
  const untilFull = notAfter === 0 ? null : localDate(notAfter)
  let state: WindowState = 'open'
  if (notBefore === 0 && notAfter === 0) state = 'unbounded'
  else if (notAfter !== 0 && notAfter * 1000 <= nowMs) state = 'closed'
  else if (notBefore !== 0 && notBefore * 1000 > nowMs) state = 'not-yet-open'
  return { from, until, fromFull, untilFull, state }
}

export const WINDOW_LABEL: Record<WindowState, string> = {
  unbounded: 'any time',
  'not-yet-open': 'not yet open',
  open: 'open',
  closed: 'closed',
}

export const WINDOW_TONE: Record<WindowState, BadgeTone> = {
  unbounded: 'neutral',
  'not-yet-open': 'warn',
  open: 'ok',
  closed: 'danger',
}
