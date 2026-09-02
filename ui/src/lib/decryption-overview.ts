// Per-application decryption progress for one epoch.
//
// The chain keeps applications in a mapping, so the only way to enumerate them
// is the `ApplicationRegistered` log; the rest of the pipeline is likewise
// reconstructed from events. Pure, so the shape of the "submitted → partials →
// organizer share → combined" visual is testable without a chain.

export interface DecryptionOverviewInput {
  applications: readonly { aid: string; creator: string; blockNumber: bigint }[]
  ciphertexts: readonly {
    aid: string
    ciphertextIndex: number
    submitter: string
    blockNumber: bigint
  }[]
  partials: readonly { aid: string; ciphertextIndex: number; participant: string }[]
  /** `OrganizerShareSubmitted`, with the aid the query fetched them for. */
  shares: readonly { aid: string; ciphertextIndex: number }[]
  combines: readonly { aid: string; ciphertextIndex: number; plaintext: bigint }[]
}

export interface CiphertextProgress {
  index: number
  submitter: string
  blockNumber: bigint
  /** Distinct committee members that published a partial for this ciphertext. */
  participants: string[]
  partials: number
  organizerShare: boolean
  combined: boolean
  plaintext: bigint | null
}

export interface ApplicationProgress {
  aid: string
  creator: string
  registeredAtBlock: bigint
  ciphertexts: CiphertextProgress[]
  submitted: number
  sharesReleased: number
  combined: number
}

/**
 * Fold the five event streams into one row per application, each carrying its
 * ciphertexts in index order. A ciphertext seen in the ciphertext log always
 * gets a row even if nothing else has happened to it yet; an application with
 * no ciphertexts still gets a row, because "registered but idle" is a state
 * worth showing.
 */
export function aggregateDecryptionProgress(
  input: DecryptionOverviewInput,
): ApplicationProgress[] {
  const apps = new Map<string, ApplicationProgress>()
  const cts = new Map<string, CiphertextProgress>()

  const ctKey = (aid: string, index: number) => `${aid.toLowerCase()}:${index}`

  const app = (aid: string): ApplicationProgress => {
    const key = aid.toLowerCase()
    let row = apps.get(key)
    if (!row) {
      row = {
        aid,
        creator: '',
        registeredAtBlock: 0n,
        ciphertexts: [],
        submitted: 0,
        sharesReleased: 0,
        combined: 0,
      }
      apps.set(key, row)
    }
    return row
  }

  for (const a of input.applications) {
    const row = app(a.aid)
    row.aid = a.aid
    row.creator = a.creator
    row.registeredAtBlock = a.blockNumber
  }

  for (const c of input.ciphertexts) {
    const row = app(c.aid)
    const key = ctKey(c.aid, c.ciphertextIndex)
    if (cts.has(key)) continue
    const ct: CiphertextProgress = {
      index: c.ciphertextIndex,
      submitter: c.submitter,
      blockNumber: c.blockNumber,
      participants: [],
      partials: 0,
      organizerShare: false,
      combined: false,
      plaintext: null,
    }
    cts.set(key, ct)
    row.ciphertexts.push(ct)
  }

  for (const p of input.partials) {
    const ct = cts.get(ctKey(p.aid, p.ciphertextIndex))
    if (!ct) continue
    // A member could in principle land the same partial twice across a reorg;
    // the threshold counts distinct members.
    if (ct.participants.some((a) => a.toLowerCase() === p.participant.toLowerCase())) continue
    ct.participants.push(p.participant)
    ct.partials = ct.participants.length
  }

  for (const s of input.shares) {
    const ct = cts.get(ctKey(s.aid, s.ciphertextIndex))
    // Re-submission overwrites on chain, so several events for one index are
    // normal and still mean exactly "the share is released".
    if (ct) ct.organizerShare = true
  }

  for (const c of input.combines) {
    const ct = cts.get(ctKey(c.aid, c.ciphertextIndex))
    if (!ct) continue
    ct.combined = true
    ct.plaintext = c.plaintext
  }

  const out = [...apps.values()]
  for (const row of out) {
    row.ciphertexts.sort((a, b) => a.index - b.index)
    row.submitted = row.ciphertexts.length
    row.sharesReleased = row.ciphertexts.filter((c) => c.organizerShare).length
    row.combined = row.ciphertexts.filter((c) => c.combined).length
  }
  // Newest application first — the one someone is most likely watching.
  out.sort((a, b) => Number(b.registeredAtBlock - a.registeredAtBlock))
  return out
}
