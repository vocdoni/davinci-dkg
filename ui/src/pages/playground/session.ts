// The resumable half of the playground that is *not* the secret.
//
// `~lib/organizer-secret` owns `sk_org`; this owns everything else the
// walkthrough needs to pick up where it left off after a reload: the value
// that was encrypted, the ciphertext built from it, the index the contract
// assigned, the share decision and the transaction records. None of it is
// sensitive — all of it is either public on chain or reconstructible — but
// losing it would strand the "verify locally" step, which compares what the
// chain stores against what this browser built.
//
// Same lifetime as the secret (`sessionStorage`, dies with the tab) so a
// resumed walkthrough never has half its state.

import type { PlaygroundState, SerialCiphertext, ShareDecision, TxSlot, TxState } from './machine'

const PREFIX = 'davinci-dkg:playground:'

export interface PlaygroundSession {
  value: string
  ciphertext: SerialCiphertext | null
  ciphertextIndex: number | null
  share: ShareDecision
  cap: number
  submitter: string
  txs: Partial<Record<TxSlot, TxState>>
}

function key(epochId: string, aid: string): string {
  return `${PREFIX}${epochId.toLowerCase()}:${aid.toLowerCase()}`
}

export function saveSession(epochId: string, aid: string, state: PlaygroundState): void {
  const session: PlaygroundSession = {
    value: state.value,
    ciphertext: state.ciphertext,
    ciphertextIndex: state.ciphertextIndex,
    share: state.share,
    cap: state.cap,
    submitter: state.submitter,
    txs: state.txs,
  }
  try {
    globalThis.sessionStorage?.setItem(key(epochId, aid), JSON.stringify(session))
  } catch {
    // Private mode or storage disabled: the walkthrough still works for this
    // page load, it just cannot be resumed after a reload.
  }
}

export function loadSession(epochId: string, aid: string): PlaygroundSession | null {
  try {
    const raw = globalThis.sessionStorage?.getItem(key(epochId, aid))
    if (!raw) return null
    const parsed = JSON.parse(raw) as Partial<PlaygroundSession>
    if (typeof parsed !== 'object' || parsed === null) return null
    return {
      value: typeof parsed.value === 'string' ? parsed.value : '',
      ciphertext: parsed.ciphertext ?? null,
      ciphertextIndex: typeof parsed.ciphertextIndex === 'number' ? parsed.ciphertextIndex : null,
      share: parsed.share === 'released' || parsed.share === 'withheld' ? parsed.share : 'undecided',
      cap: typeof parsed.cap === 'number' ? parsed.cap : 4,
      submitter: typeof parsed.submitter === 'string' ? parsed.submitter : '',
      txs: parsed.txs ?? {},
    }
  } catch {
    return null
  }
}

export function clearSession(epochId: string, aid: string): void {
  try {
    globalThis.sessionStorage?.removeItem(key(epochId, aid))
  } catch {
    // nothing to do
  }
}
