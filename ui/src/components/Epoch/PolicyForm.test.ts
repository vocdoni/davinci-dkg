import { describe, it, expect } from 'vitest'
import {
  defaultPolicyForm,
  MAX_COMMITTEE_SIZE,
  MIN_LOTTERY_ALPHA_BPS,
  validatePolicyForm,
  type PolicyFormState,
} from './PolicyForm'

function form(overrides: Partial<PolicyFormState>): PolicyFormState {
  return { ...defaultPolicyForm, ...overrides }
}

describe('validatePolicyForm', () => {
  it('returns null for the default form', () => {
    expect(validatePolicyForm(defaultPolicyForm)).toBeNull()
  })

  it('rejects threshold < 1', () => {
    expect(validatePolicyForm(form({ threshold: '0' }))).toMatch(/Threshold must be at least 1/i)
  })

  it('rejects committee size < 1', () => {
    expect(validatePolicyForm(form({ committeeSize: '0', threshold: '0', minValidContributions: '0' }))).toBeTruthy()
  })

  it('rejects threshold > committee size', () => {
    expect(validatePolicyForm(form({ threshold: '5', committeeSize: '3', minValidContributions: '5' }))).toMatch(
      /Threshold cannot exceed committee size/i
    )
  })

  it('rejects min > committee size', () => {
    expect(validatePolicyForm(form({ threshold: '2', committeeSize: '3', minValidContributions: '5' }))).toMatch(
      /Min valid contributions cannot exceed committee size/i
    )
  })

  it('rejects min < threshold (the load-bearing footgun)', () => {
    const err = validatePolicyForm(form({ threshold: '5', committeeSize: '10', minValidContributions: '3' }))
    expect(err).toMatch(/Min valid contributions \(3\) must be ≥ threshold \(5\)/)
    // The message should explain the consequence so users understand why we
    // refuse to submit, not just that we did.
    expect(err).toMatch(/finalize but no one will be able to decrypt/i)
  })

  it('accepts min > threshold (extra redundancy)', () => {
    expect(validatePolicyForm(form({ threshold: '3', committeeSize: '10', minValidContributions: '7' }))).toBeNull()
  })

  it('accepts min == threshold (the auto-linked default case)', () => {
    expect(validatePolicyForm(form({ threshold: '4', committeeSize: '10', minValidContributions: '4' }))).toBeNull()
  })

  it('accepts committee size exactly at MAX_COMMITTEE_SIZE', () => {
    const n = String(MAX_COMMITTEE_SIZE)
    expect(validatePolicyForm(form({ threshold: '2', committeeSize: n, minValidContributions: '2' }))).toBeNull()
  })

  it('rejects committee size > MAX_COMMITTEE_SIZE (circuit MaxN cap)', () => {
    const n = String(MAX_COMMITTEE_SIZE + 1)
    const err = validatePolicyForm(form({ threshold: '2', committeeSize: n, minValidContributions: '2' }))
    expect(err).toMatch(new RegExp(`cannot exceed ${MAX_COMMITTEE_SIZE}`))
    expect(err).toMatch(/MaxN/)
  })

  it('rejects lottery α below 1.0× (contract floor)', () => {
    expect(validatePolicyForm(form({ lotteryAlphaBps: String(MIN_LOTTERY_ALPHA_BPS - 1) }))).toMatch(/Lottery α must be at least/)
    expect(validatePolicyForm(form({ lotteryAlphaBps: String(MIN_LOTTERY_ALPHA_BPS) }))).toBeNull()
  })

  describe('deployment bounds', () => {
    const bounds = { minThreshold: 2, minCommitteeSize: 3, maxLotteryAlphaBps: 20000 }

    it('accepts the default form under the bounds', () => {
      expect(validatePolicyForm(defaultPolicyForm, bounds)).toBeNull()
    })

    it('is a no-op when bounds are unknown', () => {
      expect(validatePolicyForm(form({ threshold: '1', minValidContributions: '1' }), null)).toBeNull()
      expect(validatePolicyForm(form({ threshold: '1', minValidContributions: '1' }), undefined)).toBeNull()
    })

    it('rejects threshold below MIN_THRESHOLD', () => {
      expect(validatePolicyForm(form({ threshold: '1', minValidContributions: '1' }), bounds)).toMatch(/MIN_THRESHOLD/)
    })

    it('rejects committee size below MIN_COMMITTEE_SIZE', () => {
      expect(validatePolicyForm(form({ threshold: '2', committeeSize: '2', minValidContributions: '2' }), bounds)).toMatch(
        /MIN_COMMITTEE_SIZE/
      )
    })

    it('rejects lottery α above MAX_LOTTERY_ALPHA_BPS', () => {
      expect(validatePolicyForm(form({ lotteryAlphaBps: '20001' }), bounds)).toMatch(/MAX_LOTTERY_ALPHA_BPS/)
      expect(validatePolicyForm(form({ lotteryAlphaBps: '20000' }), bounds)).toBeNull()
    })
  })
})
