import { describe, expect, it } from 'vitest'
import type { ApplicationRow } from '~indexer/selectors'
import type { Address, Aid, EpochId } from '~indexer/types'
import { describeDecryptionWindow, submissionPolicy, submissionPolicyLabel } from './policy'
import { filterApplications, summarizeApplications } from './filter'

const EPOCH_A = '0x0000000000000000000000a1' as EpochId
const EPOCH_B = '0x0000000000000000000000b2' as EpochId

function row(overrides: Partial<ApplicationRow> & { aid: string; epoch: EpochId }): ApplicationRow {
  return {
    key: `${overrides.epoch}:${overrides.aid}`,
    creator: '0x1111111111111111111111111111111111111111' as Address,
    organizerPK: { x: 1n, y: 2n },
    mode: 'organizer-locked',
    poolIndex: 0,
    poolKey: { x: 3n, y: 4n },
    openSubmission: false,
    submitters: ['0x2222222222222222222222222222222222222222' as Address],
    maxCiphertexts: 8,
    notBeforeBlock: 0,
    notAfterBlock: 0,
    decryptNotBefore: 0,
    decryptNotAfter: 0,
    unlocked: false,
    organizerSecret: null,
    revealBlock: null,
    revealTx: null,
    createdBlock: 10,
    createdTx: null,
    ciphertexts: 4,
    decrypted: 2,
    ...overrides,
    aid: overrides.aid as Aid,
  }
}

const rows: ApplicationRow[] = [
  row({ epoch: EPOCH_A, aid: '0xabcd01' }),
  row({
    epoch: EPOCH_A,
    aid: '0xbeef02',
    creator: '0x3333333333333333333333333333333333333333' as Address,
    unlocked: true,
    organizerSecret: 7n,
    revealBlock: 12,
  }),
  row({
    epoch: EPOCH_B,
    aid: '0xabcd03',
    mode: 'automatic',
    poolIndex: 1,
    organizerPK: { x: 0n, y: 1n },
    openSubmission: true,
    submitters: [],
    unlocked: true,
    ciphertexts: 2,
    decrypted: 0,
  }),
]

describe('filterApplications', () => {
  it('filters by epoch', () => {
    expect(filterApplications(rows, { epoch: EPOCH_B })).toHaveLength(1)
    expect(filterApplications(rows, { epoch: 'all' })).toHaveLength(3)
  })

  it('matches the application id by prefix', () => {
    expect(filterApplications(rows, { query: '0xabcd' })).toHaveLength(2)
    expect(filterApplications(rows, { query: 'beef' })).toHaveLength(1)
  })

  it('matches the organizer and any allow-listed submitter', () => {
    expect(filterApplications(rows, { query: '0x3333' })).toHaveLength(1)
    // The open-submission application has no allow-list to match.
    expect(filterApplications(rows, { query: '0x2222' })).toHaveLength(2)
  })

  it('filters by mode', () => {
    expect(filterApplications(rows, { mode: 'automatic' })).toHaveLength(1)
    expect(filterApplications(rows, { mode: 'organizer-locked' })).toHaveLength(2)
    expect(filterApplications(rows, { mode: 'all' })).toHaveLength(3)
  })

  it('combines the filters', () => {
    expect(filterApplications(rows, { epoch: EPOCH_A, query: '0xabcd' })).toHaveLength(1)
    expect(filterApplications(rows, { epoch: EPOCH_B, mode: 'organizer-locked' })).toHaveLength(0)
  })
})

describe('submission policy and decryption window wording', () => {
  it('names the three submission policies', () => {
    expect(submissionPolicyLabel(submissionPolicy({ openSubmission: true, submitters: [] }))).toBe('open')
    expect(submissionPolicyLabel(submissionPolicy({ openSubmission: false, submitters: [] }))).toBe('registrant only')
    expect(submissionPolicyLabel(submissionPolicy(rows[0]))).toBe('allow-list')
    expect(submissionPolicyLabel(submissionPolicy({ openSubmission: null, submitters: null }))).toBe('—')
  })

  it('reads the window as dates and knows where the clock sits', () => {
    const now = 1_700_000_000_000
    expect(describeDecryptionWindow(null, 0, now)).toBeNull()
    expect(describeDecryptionWindow(0, null, now)).toBeNull()
    expect(describeDecryptionWindow(0, 0, now)).toEqual({
      from: null,
      until: null,
      fromFull: null,
      untilFull: null,
      state: 'unbounded',
    })
    // The compact form drops the time of day but keeps it for the title.
    const compact = describeDecryptionWindow(1_700_000_000 - 60, 0, now, 'date')!
    expect(compact.from).not.toContain(':')
    expect(compact.fromFull).toContain(':')
    expect(describeDecryptionWindow(0, 1_700_000_000 + 60, now)?.state).toBe('open')
    expect(describeDecryptionWindow(0, 1_700_000_000 - 60, now)?.state).toBe('closed')
    expect(describeDecryptionWindow(1_700_000_000 + 60, 0, now)?.state).toBe('not-yet-open')
    expect(describeDecryptionWindow(1_700_000_000 - 60, 1_700_000_000 + 60, now)?.state).toBe('open')
    // A closed window is closed even if it also has not "opened" yet on paper.
    expect(describeDecryptionWindow(1_700_000_000 + 60, 1_700_000_000 - 60, now)?.state).toBe('closed')
    const window = describeDecryptionWindow(1_700_000_000, 1_700_000_000 + 60, now)!
    expect(window.from).not.toBeNull()
    expect(window.until).not.toBeNull()
    expect(describeDecryptionWindow(0, 1_700_000_000, now)?.from).toBeNull()
  })
})

describe('summarizeApplications', () => {
  it('totals the pipeline, counting every organizer that has not revealed', () => {
    expect(summarizeApplications(rows)).toEqual({
      applications: 3,
      ciphertexts: 10,
      decrypted: 4,
      locked: 1,
    })
  })

  it('is zero for an empty list', () => {
    expect(summarizeApplications([])).toEqual({ applications: 0, ciphertexts: 0, decrypted: 0, locked: 0 })
  })
})
