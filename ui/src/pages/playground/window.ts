// The decryption window as the register step edits it: two `datetime-local`
// values, each empty for "unbounded", turned into the unix seconds the
// contract stores and checked against the rules `registerApplication`
// enforces before any transaction is built.

export interface WindowSeconds {
  notBefore: number
  notAfter: number
}

/** `datetime-local` value → unix seconds; empty → 0 (unbounded); null when not a date. */
export function windowSeconds(value: string): number | null {
  if (value.trim() === '') return 0
  const ms = Date.parse(value)
  if (!Number.isFinite(ms)) return null
  return Math.floor(ms / 1000)
}

/** Unix seconds → a `datetime-local` value in local time; 0 → empty. */
export function windowInput(seconds: number): string {
  if (!seconds) return ''
  const d = new Date(seconds * 1000)
  const pad = (n: number) => String(n).padStart(2, '0')
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())}T${pad(d.getHours())}:${pad(d.getMinutes())}`
}

/**
 * Validate a window the way the contract will: an end must be in the future
 * (`InvalidPolicy()` otherwise) and, when both sides are set, the window must
 * open before it closes.
 */
export function validateWindow(notBefore: string, notAfter: string, nowMs = Date.now()): WindowSeconds | { error: string } {
  const from = windowSeconds(notBefore)
  const until = windowSeconds(notAfter)
  if (from == null) return { error: 'The window start is not a date' }
  if (until == null) return { error: 'The window end is not a date' }
  if (until !== 0 && until * 1000 <= nowMs) {
    return { error: 'decryptNotAfter must be in the future — the contract rejects a window that is already closed' }
  }
  if (from !== 0 && until !== 0 && from >= until) return { error: 'The decryption window must open before it closes' }
  return { notBefore: from, notAfter: until }
}
