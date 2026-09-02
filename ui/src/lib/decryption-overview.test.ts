import { describe, expect, it } from 'vitest'
import {
  aggregateDecryptionProgress,
  type DecryptionOverviewInput,
} from './decryption-overview'

const AID = '0x0a01'
const OTHER = '0x0b02'
const ORG = '0xorganizer'
const M1 = '0xMember1'
const M2 = '0xMember2'

function input(overrides: Partial<DecryptionOverviewInput> = {}): DecryptionOverviewInput {
  return {
    applications: [],
    ciphertexts: [],
    partials: [],
    shares: [],
    combines: [],
    ...overrides,
  }
}

describe('aggregateDecryptionProgress', () => {
  it('walks one ciphertext through the whole pipeline', () => {
    const [app] = aggregateDecryptionProgress(
      input({
        applications: [{ aid: AID, creator: ORG, blockNumber: 10n }],
        ciphertexts: [{ aid: AID, ciphertextIndex: 1, submitter: ORG, blockNumber: 11n }],
        partials: [
          { aid: AID, ciphertextIndex: 1, participant: M1 },
          { aid: AID, ciphertextIndex: 1, participant: M2 },
        ],
        shares: [{ aid: AID, ciphertextIndex: 1 }],
        combines: [{ aid: AID, ciphertextIndex: 1, plaintext: 42n }],
      }),
    )
    expect(app).toMatchObject({ submitted: 1, sharesReleased: 1, combined: 1 })
    expect(app.ciphertexts[0]).toMatchObject({
      index: 1,
      partials: 2,
      participants: [M1, M2],
      organizerShare: true,
      combined: true,
      plaintext: 42n,
    })
  })

  it('counts distinct members only, and tolerates repeated share events', () => {
    const [app] = aggregateDecryptionProgress(
      input({
        applications: [{ aid: AID, creator: ORG, blockNumber: 1n }],
        ciphertexts: [{ aid: AID, ciphertextIndex: 1, submitter: ORG, blockNumber: 2n }],
        partials: [
          { aid: AID, ciphertextIndex: 1, participant: M1 },
          { aid: AID, ciphertextIndex: 1, participant: M1.toLowerCase() },
        ],
        shares: [
          { aid: AID, ciphertextIndex: 1 },
          { aid: AID, ciphertextIndex: 1 },
        ],
      }),
    )
    expect(app.ciphertexts[0].partials).toBe(1)
    expect(app.sharesReleased).toBe(1)
  })

  it('keeps a registered-but-idle application, and never crosses aids', () => {
    const rows = aggregateDecryptionProgress(
      input({
        applications: [
          { aid: AID, creator: ORG, blockNumber: 5n },
          { aid: OTHER, creator: ORG, blockNumber: 9n },
        ],
        ciphertexts: [{ aid: AID, ciphertextIndex: 1, submitter: ORG, blockNumber: 6n }],
        partials: [{ aid: OTHER, ciphertextIndex: 1, participant: M1 }],
      }),
    )
    // Newest registration first.
    expect(rows.map((r) => r.aid)).toEqual([OTHER, AID])
    expect(rows[0].ciphertexts).toHaveLength(0)
    expect(rows[1].ciphertexts[0].partials).toBe(0)
  })

  it('orders ciphertexts by their on-chain index', () => {
    const [app] = aggregateDecryptionProgress(
      input({
        applications: [{ aid: AID, creator: ORG, blockNumber: 1n }],
        ciphertexts: [
          { aid: AID, ciphertextIndex: 3, submitter: ORG, blockNumber: 8n },
          { aid: AID, ciphertextIndex: 1, submitter: ORG, blockNumber: 2n },
          { aid: AID, ciphertextIndex: 2, submitter: ORG, blockNumber: 5n },
        ],
      }),
    )
    expect(app.ciphertexts.map((c) => c.index)).toEqual([1, 2, 3])
  })
})
