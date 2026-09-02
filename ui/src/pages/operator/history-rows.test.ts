import { describe, expect, it } from 'vitest'
import type { OperatorEpochHistory } from '~indexer/selectors'
import type { EpochId, Hex, IndexedEvent } from '~indexer/types'
import { operatorHistoryRows, participationTrend, partialsTrend } from './history-rows'

const EPOCH_A = '0x0000000000000000000000a1' as EpochId
const EPOCH_B = '0x0000000000000000000000b2' as EpochId
const OPERATOR = '0x00000000000000000000000000000000000000aa' as const

function history(overrides: Partial<OperatorEpochHistory> & { epoch: EpochId; nonce: number }): OperatorEpochHistory {
  return {
    claimed: false,
    slot: null,
    contributed: false,
    contributionBlock: null,
    finalized: false,
    partials: 0,
    combines: 0,
    ...overrides,
  }
}

const events: IndexedEvent[] = [
  {
    name: 'SlotClaimed',
    block: 10,
    tx: '0xaa' as Hex,
    logIndex: 0,
    epoch: EPOCH_A,
    aid: null,
    actor: OPERATOR,
    data: { epochId: EPOCH_A, claimer: OPERATOR, slot: 2 },
  },
  {
    name: 'ContributionSubmitted',
    block: 14,
    tx: '0xbb' as Hex,
    logIndex: 0,
    epoch: EPOCH_A,
    aid: null,
    actor: OPERATOR,
    data: {
      epochId: EPOCH_A,
      contributor: OPERATOR,
      contributorIndex: 3,
      commitmentsHash: '0xcc' as Hex,
      encryptedSharesHash: '0xdd' as Hex,
    },
  },
  {
    name: 'SlotClaimed',
    block: 40,
    tx: '0xee' as Hex,
    logIndex: 0,
    epoch: EPOCH_B,
    aid: null,
    actor: OPERATOR,
    data: { epochId: EPOCH_B, claimer: OPERATOR, slot: 5 },
  },
]

const rows = operatorHistoryRows(
  [
    history({
      epoch: EPOCH_A,
      nonce: 1,
      claimed: true,
      slot: 2,
      contributed: true,
      contributionBlock: 14,
      partials: 4,
    }),
    history({ epoch: EPOCH_B, nonce: 2, claimed: true, slot: 5, partials: 1 }),
  ],
  events,
  (tx) => (tx === '0xbb' ? 462_523 : null)
)

describe('operatorHistoryRows', () => {
  it('is newest epoch first', () => {
    expect(rows.map((row) => row.nonce)).toEqual([2, 1])
  })

  it('joins the claim and contribution transactions onto each epoch', () => {
    const first = rows[1]
    expect(first.claimTx).toBe('0xaa')
    expect(first.claimBlock).toBe(10)
    expect(first.contributionTx).toBe('0xbb')
    expect(first.contributionGas).toBe(462_523)
  })

  it('leaves an epoch without a contribution empty rather than guessing', () => {
    expect(rows[0].contributionTx).toBeNull()
    expect(rows[0].contributionGas).toBeNull()
  })
})

describe('trends', () => {
  it('reports cumulative participation, oldest epoch first', () => {
    // Epoch 1: claimed and contributed → 100%. Epoch 2: claimed, no
    // contribution → 1 of 2.
    expect(participationTrend(rows)).toEqual([100, 50])
  })

  it('reports partials per epoch, oldest first', () => {
    expect(partialsTrend(rows)).toEqual([4, 1])
  })
})
