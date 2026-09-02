import { describe, expect, it } from 'vitest'
import type { Aid, EpochId, Hex, IndexedEvent } from '~indexer/types'
import { describeEvent, epochLabel } from './describe'

const EPOCH = '0x0000000000000000000000a1' as EpochId
const AID = `0x${'2c'.repeat(32)}` as Aid
const OPERATOR = '0x00000000000000000000000000000000000000aa' as const

function event<T extends IndexedEvent>(partial: T): T {
  return partial
}

const nonceOf = (id: EpochId) => (id === EPOCH ? 7 : null)

describe('epochLabel', () => {
  it('prefers the nonce and falls back to a short id', () => {
    expect(epochLabel(EPOCH, nonceOf)).toBe('epoch #7')
    expect(epochLabel(EPOCH)).toBe('epoch 0x000000…00a1')
    expect(epochLabel(null)).toBe('the registry')
  })
})

describe('describeEvent', () => {
  it('names the epoch and links to it for a slot claim', () => {
    const described = describeEvent(
      event({
        name: 'SlotClaimed',
        block: 10,
        tx: '0xaa' as Hex,
        logIndex: 0,
        epoch: EPOCH,
        aid: null,
        actor: OPERATOR,
        data: { epochId: EPOCH, claimer: OPERATOR, slot: 4 },
      }),
      nonceOf
    )
    expect(described.title).toBe('Claimed slot 4 in epoch #7')
    // Slots are 0-based, participant indices are 1-based; both are shown.
    expect(described.detail).toBe('participant index 5')
    expect(described.href).toBe(`/epochs/${EPOCH}`)
  })

  it('links an application event to the application, not the epoch', () => {
    const described = describeEvent(
      event({
        name: 'OrganizerShareSubmitted',
        block: 20,
        tx: '0xbb' as Hex,
        logIndex: 1,
        epoch: EPOCH,
        aid: AID,
        actor: null,
        data: {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 3,
          delta: { x: 1n, y: 2n },
          a1: { x: 3n, y: 4n },
          a2: { x: 5n, y: 6n },
          z: 7n,
        },
      }),
      nonceOf
    )
    expect(described.title).toBe('Released the organizer share for ciphertext 3')
    expect(described.href).toBe(`/applications/${EPOCH}/${AID}`)
  })

  it('marks the destructive events', () => {
    const reaped = describeEvent(
      event({
        name: 'NodeReaped',
        block: 30,
        tx: null,
        logIndex: 0,
        epoch: null,
        aid: null,
        actor: OPERATOR,
        data: { operator: OPERATOR, lastActiveBlock: 12 },
      })
    )
    expect(reaped.tone).toBe('danger')
    expect(reaped.detail).toBe('last active at block 12')
    expect(reaped.href).toBeUndefined()
  })
})
