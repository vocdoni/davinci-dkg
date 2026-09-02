// Session-scoped store for organizer secrets.
//
// `sk_org` is an application's ONLY decryption capability: it never leaves the
// browser, the contract stores only `PK_org = sk_org·G`, and nothing on chain
// can reconstruct it. Losing it makes every ciphertext under that `aid`
// permanently undecryptable.
//
// We therefore keep it in `sessionStorage`, deliberately:
//
//   • the playground needs it again a few steps later (to release the share),
//     so holding it only in React state would lose it on a page reload;
//   • `sessionStorage` dies with the tab, so a demo secret does not linger on
//     a shared machine the way `localStorage` would.
//
// It is NOT a place to keep a production secret. The UI says so next to every
// generated value, and a real organizer is expected to copy it into whatever
// they actually use for key custody.

import type { Hex } from 'viem'

const PREFIX = 'davinci-dkg:organizer-secret:'

function key(epochId: Hex, aid: Hex): string {
  return `${PREFIX}${epochId.toLowerCase()}:${aid.toLowerCase()}`
}

/** Persist `sk_org` for `(epochId, aid)` for the lifetime of the tab. */
export function saveOrganizerSecret(epochId: Hex, aid: Hex, skOrg: bigint): void {
  try {
    globalThis.sessionStorage?.setItem(key(epochId, aid), skOrg.toString())
  } catch {
    // Private mode / storage disabled — the caller still holds the value in
    // React state for this page load, and the UI has already shown it.
  }
}

/** Read back a previously stored `sk_org`, or null when there is none. */
export function loadOrganizerSecret(epochId: Hex, aid: Hex): bigint | null {
  try {
    const raw = globalThis.sessionStorage?.getItem(key(epochId, aid))
    if (!raw) return null
    const v = BigInt(raw)
    return v === 0n ? null : v
  } catch {
    return null
  }
}

/** Forget a stored `sk_org` (used when the user re-rolls a secret). */
export function clearOrganizerSecret(epochId: Hex, aid: Hex): void {
  try {
    globalThis.sessionStorage?.removeItem(key(epochId, aid))
  } catch {
    // nothing to do
  }
}

/**
 * Parse a user-supplied secret. Accepts decimal or `0x`-hex; returns null for
 * anything that is not a non-zero integer. Range (`< q`) is enforced by the
 * SDK's prover, which throws with a precise message.
 */
export function parseOrganizerSecret(input: string): bigint | null {
  const trimmed = input.trim()
  if (trimmed === '') return null
  if (!/^(0x[0-9a-fA-F]+|\d+)$/.test(trimmed)) return null
  try {
    const v = BigInt(trimmed)
    return v === 0n ? null : v
  } catch {
    return null
  }
}
