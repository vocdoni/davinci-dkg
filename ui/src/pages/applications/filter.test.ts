import { describe, expect, it } from 'vitest'
import type { ApplicationRow } from '~indexer/selectors'
import type { Address, Aid, EpochId } from '~indexer/types'
import { filterApplications, summarizeApplications } from './filter'

const EPOCH_A = '0x0000000000000000000000a1' as EpochId
const EPOCH_B = '0x0000000000000000000000b2' as EpochId

function row(overrides: Partial<ApplicationRow> & { aid: string; epoch: EpochId }): ApplicationRow {
  return {
    key: `${overrides.epoch}:${overrides.aid}`,
    creator: '0x1111111111111111111111111111111111111111' as Address,
    organizerPK: { x: 1n, y: 2n },
    authorizedSubmitter: '0x2222222222222222222222222222222222222222' as Address,
    maxCiphertexts: 8,
    notBeforeBlock: 0,
    notAfterBlock: 0,
    createdBlock: 10,
    createdTx: null,
    ciphertexts: 4,
    decrypted: 2,
    sharesPublished: 3,
    ...overrides,
    aid: overrides.aid as Aid,
  }
}

const rows: ApplicationRow[] = [
  row({ epoch: EPOCH_A, aid: '0xabcd01' }),
  row({ epoch: EPOCH_A, aid: '0xbeef02', creator: '0x3333333333333333333333333333333333333333' as Address }),
  row({ epoch: EPOCH_B, aid: '0xabcd03', ciphertexts: 2, decrypted: 0, sharesPublished: 0 }),
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

  it('matches the organizer and the authorized submitter', () => {
    expect(filterApplications(rows, { query: '0x3333' })).toHaveLength(1)
    expect(filterApplications(rows, { query: '0x2222' })).toHaveLength(3)
  })

  it('combines both filters', () => {
    expect(filterApplications(rows, { epoch: EPOCH_A, query: '0xabcd' })).toHaveLength(1)
  })
})

describe('summarizeApplications', () => {
  it('totals the pipeline, counting every share-less ciphertext as pending', () => {
    expect(summarizeApplications(rows)).toEqual({
      applications: 3,
      ciphertexts: 10,
      decrypted: 4,
      pendingShares: 4,
    })
  })

  it('is zero for an empty list', () => {
    expect(summarizeApplications([])).toEqual({
      applications: 0,
      ciphertexts: 0,
      decrypted: 0,
      pendingShares: 0,
    })
  })
})
