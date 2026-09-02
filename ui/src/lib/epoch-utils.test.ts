import { describe, it, expect } from 'vitest'
import { EpochPhase, type Epoch } from '@vocdoni/davinci-dkg-sdk'
import {
  phaseSequence,
  roundFailure,
  roundPhase,
  roundPhaseColor,
  roundPhaseLabel,
  roundSummary,
} from './epoch-utils'

function mkRound(overrides: Partial<Epoch> = {}): Epoch {
  return {
    organizer: '0x0000000000000000000000000000000000000001' as `0x${string}`,
    nonce: 1n,
    seed: ('0x' + '0'.repeat(64)) as `0x${string}`,
    startBlock: 99n,
    seedBlock: 100n,
    lotteryThreshold: 0n,
    status: EpochPhase.CommitteeSelection,
    claimedCount: 0,
    contributionCount: 0,
    partialDecryptionCount: 0,
    ciphertextCount: 0,
    policy: {
      threshold: 2,
      committeeSize: 3,
      minValidContributions: 2,
      lotteryAlphaBps: 15000,
      committeeSelectionDeadlineBlock: 200n,
      keyAssemblyDeadlineBlock: 300n,
      liveNotBeforeBlock: 305n,
    },
    ...overrides,
  }
}

describe('roundPhase + label + color', () => {
  it('maps each known status to its phase', () => {
    expect(roundPhase(mkRound({ status: EpochPhase.CommitteeSelection }))).toBe('committee-selection')
    expect(roundPhase(mkRound({ status: EpochPhase.KeyAssembly }))).toBe('key-assembly')
    expect(roundPhase(mkRound({ status: EpochPhase.Live }))).toBe('live')
    expect(roundPhase(mkRound({ status: EpochPhase.Completed }))).toBe('completed')
    expect(roundPhase(mkRound({ status: EpochPhase.Aborted }))).toBe('aborted')
  })

  it('every phase has a label and a color', () => {
    for (const phase of [...phaseSequence, 'aborted', 'unknown'] as const) {
      expect(roundPhaseLabel(phase)).toBeTruthy()
      expect(roundPhaseColor(phase)).toBeTruthy()
    }
  })

  it('phase sequence is the canonical four-step timeline', () => {
    expect(phaseSequence).toEqual(['committee-selection', 'key-assembly', 'live', 'completed'])
  })
})

describe('roundSummary', () => {
  it('mentions claiming during CommitteeSelection', () => {
    const r = mkRound({ status: EpochPhase.CommitteeSelection, claimedCount: 1 })
    const out = roundSummary(r, 100n)
    expect(out).toMatch(/1\/3/)
    expect(out).toMatch(/committee|claim/i)
  })

  it('reports threshold-met during KeyAssembly', () => {
    const r = mkRound({ status: EpochPhase.KeyAssembly, contributionCount: 2 })
    expect(roundSummary(r, 100n)).toMatch(/Threshold met/i)
  })

  it('reports awaiting contributions when below threshold', () => {
    const r = mkRound({ status: EpochPhase.KeyAssembly, contributionCount: 0 })
    expect(roundSummary(r, 100n)).toMatch(/Awaiting contributions/i)
  })

  it('says live when in Live', () => {
    expect(roundSummary(mkRound({ status: EpochPhase.Live }), 100n)).toMatch(/live/i)
  })

  it('says aborted when in Aborted', () => {
    expect(roundSummary(mkRound({ status: EpochPhase.Aborted }), 100n)).toMatch(/aborted/i)
  })
})

// Mirrors DKGManager.abortEpoch: the epoch is dead (and anyone may abort it)
// exactly when roundFailure is non-null.
describe('roundFailure', () => {
  it('is null without a current block', () => {
    expect(roundFailure(mkRound(), null)).toBeNull()
  })

  it('is null while the committee-selection deadline has not passed', () => {
    expect(roundFailure(mkRound({ status: EpochPhase.CommitteeSelection }), 200n)).toBeNull()
  })

  it('flags CommitteeSelection past its deadline regardless of how many slots were claimed', () => {
    // A full committee would have moved the epoch to KeyAssembly, so still
    // being in CommitteeSelection past the deadline means it never filled —
    // even when claimed >= minValidContributions.
    const r = mkRound({ status: EpochPhase.CommitteeSelection, claimedCount: 2 })
    expect(roundFailure(r, 201n)).toEqual({ kind: 'committee-selection', have: 2, need: 3, total: 3 })
  })

  it('flags KeyAssembly past its deadline only when contributions are below minValidContributions', () => {
    const short = mkRound({ status: EpochPhase.KeyAssembly, contributionCount: 1 })
    expect(roundFailure(short, 301n)).toEqual({ kind: 'key-assembly', have: 1, need: 2, total: 3 })
    const enough = mkRound({ status: EpochPhase.KeyAssembly, contributionCount: 2 })
    expect(roundFailure(enough, 301n)).toBeNull()
    expect(roundFailure(short, 300n)).toBeNull()
  })

  it('is null for Live, Completed and Aborted epochs', () => {
    for (const status of [EpochPhase.Live, EpochPhase.Completed, EpochPhase.Aborted] as const) {
      expect(roundFailure(mkRound({ status }), 10_000n)).toBeNull()
    }
  })
})
