import { describe, expect, it } from 'vitest'
import type { IndexedEvent } from '~indexer/types'
import { EVENT_TONE, eventSummary } from './events'

const envelope = { block: 1_000, tx: null, logIndex: 0, epoch: null, aid: null, actor: null }

describe('eventSummary', () => {
  it('prints both numberings for a claim, because they differ by one', () => {
    const event: IndexedEvent = {
      ...envelope,
      name: 'SlotClaimed',
      data: { epochId: '0x01', claimer: '0x0000000000000000000000000000000000000001', slot: 0 },
    }
    expect(eventSummary(event)).toBe('slot 0 · participant index 1')
  })

  it('names the ciphertext and the 1-based participant index of a partial', () => {
    const event: IndexedEvent = {
      ...envelope,
      name: 'PartialDecryptionSubmitted',
      data: {
        epochId: '0x01',
        aid: '0x02',
        participant: '0x0000000000000000000000000000000000000001',
        participantIndex: 7,
        ciphertextIndex: 3,
        delta: { x: 1n, y: 2n },
      },
    }
    expect(eventSummary(event)).toBe('ciphertext 3 · participant index 7')
  })

  it('carries the recovered plaintext of a combine', () => {
    const event: IndexedEvent = {
      ...envelope,
      name: 'DecryptionCombined',
      data: { epochId: '0x01', aid: '0x02', ciphertextIndex: 1, combineHash: '0x03', plaintext: 42n },
    }
    expect(eventSummary(event)).toContain('m = 42')
  })

  it('has a line for every event the store can hold', () => {
    for (const name of Object.keys(EVENT_TONE)) {
      expect(EVENT_TONE[name as keyof typeof EVENT_TONE]).toBeDefined()
    }
  })

  it('marks the milestones and the failures, and nothing else', () => {
    expect(EVENT_TONE.EpochLive).toBe('ok')
    expect(EVENT_TONE.DecryptionCombined).toBe('ok')
    expect(EVENT_TONE.EpochAborted).toBe('danger')
    expect(EVENT_TONE.ContributionSubmitted).toBe('neutral')
  })
})
