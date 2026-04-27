import { describe, it, expect } from 'vitest'
import { RoundStatus, type Round } from '@vocdoni/davinci-dkg-sdk'
import { phaseSequence, roundPhase, roundPhaseColor, roundPhaseLabel, roundSummary } from './round-utils'

function mkRound(overrides: Partial<Round> = {}): Round {
  return {
    organizer: '0x0000000000000000000000000000000000000001' as `0x${string}`,
    nonce: 1n,
    seed: ('0x' + '0'.repeat(64)) as `0x${string}`,
    seedBlock: 100n,
    lotteryThreshold: 0n,
    status: RoundStatus.Registration,
    claimedCount: 0,
    contributionCount: 0,
    partialDecryptionCount: 0,
    revealedShareCount: 0,
    ciphertextCount: 0,
    policy: {
      threshold: 2,
      committeeSize: 3,
      minValidContributions: 2,
      lotteryAlphaBps: 15000,
      seedDelay: 2,
      registrationDeadlineBlock: 200n,
      contributionDeadlineBlock: 300n,
      finalizeNotBeforeBlock: 305n,
      disclosureAllowed: false,
    },
    decryptionPolicy: {
      ownerOnly: false,
      maxDecryptions: 0,
      notBeforeBlock: 0n,
      notBeforeTimestamp: 0n,
      notAfterBlock: 0n,
      notAfterTimestamp: 0n,
    },
    ...overrides,
  }
}

describe('roundPhase + label + color', () => {
  it('maps each known status to its phase', () => {
    expect(roundPhase(mkRound({ status: RoundStatus.Registration }))).toBe('registration')
    expect(roundPhase(mkRound({ status: RoundStatus.Contribution }))).toBe('contribution')
    expect(roundPhase(mkRound({ status: RoundStatus.Finalized }))).toBe('finalized')
    expect(roundPhase(mkRound({ status: RoundStatus.Completed }))).toBe('completed')
    expect(roundPhase(mkRound({ status: RoundStatus.Aborted }))).toBe('aborted')
  })

  it('every phase has a label and a color', () => {
    for (const phase of [...phaseSequence, 'aborted', 'unknown'] as const) {
      expect(roundPhaseLabel(phase)).toBeTruthy()
      expect(roundPhaseColor(phase)).toBeTruthy()
    }
  })

  it('phase sequence is the canonical four-step timeline', () => {
    expect(phaseSequence).toEqual(['registration', 'contribution', 'finalized', 'completed'])
  })
})

describe('roundSummary', () => {
  it('mentions the registration deadline when in Registration', () => {
    const r = mkRound({ status: RoundStatus.Registration, claimedCount: 1 })
    const out = roundSummary(r, 100n)
    expect(out).toMatch(/1\/3/)
    expect(out).toMatch(/Registration|claim/i)
  })

  it('reports threshold-met during Contribution', () => {
    const r = mkRound({ status: RoundStatus.Contribution, contributionCount: 2 })
    expect(roundSummary(r, 100n)).toMatch(/Threshold met/i)
  })

  it('reports awaiting contributions when below threshold', () => {
    const r = mkRound({ status: RoundStatus.Contribution, contributionCount: 0 })
    expect(roundSummary(r, 100n)).toMatch(/Awaiting contributions/i)
  })

  it('says finalized when in Finalized', () => {
    expect(roundSummary(mkRound({ status: RoundStatus.Finalized }), 100n)).toMatch(/finalized/i)
  })

  it('says aborted when in Aborted', () => {
    expect(roundSummary(mkRound({ status: RoundStatus.Aborted }), 100n)).toMatch(/aborted/i)
  })
})
