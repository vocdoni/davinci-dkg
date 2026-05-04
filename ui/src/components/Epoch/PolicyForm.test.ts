import { describe, it, expect } from 'vitest'
import { defaultPolicyForm, MAX_COMMITTEE_SIZE, validatePolicyForm, type PolicyFormState } from './PolicyForm'

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
})
