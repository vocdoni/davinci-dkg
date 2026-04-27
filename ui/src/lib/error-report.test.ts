import { describe, it, expect } from 'vitest'
import { buildErrorReport } from './error-report'

describe('buildErrorReport', () => {
  it('includes the error message and stack', () => {
    const err = new Error('boom')
    const report = buildErrorReport(err, { route: '/x', chainId: 1 })
    expect(report).toContain('boom')
    expect(report).toContain('### Stack')
    expect(report).toContain('/x')
  })

  it('handles non-Error throwables', () => {
    expect(buildErrorReport('string failure')).toContain('string failure')
    expect(buildErrorReport({ foo: 1 })).toContain('[object Object]')
  })

  it('omits absent context lines', () => {
    const out = buildErrorReport(new Error('x'))
    expect(out).not.toContain('- route:')
    expect(out).not.toContain('- wallet:')
  })

  it('formats a chain line when chainId or chainName is present', () => {
    expect(buildErrorReport(new Error('x'), { chainId: 11155111 })).toMatch(/- chain: \(unknown\) \(id 11155111\)/)
    expect(buildErrorReport(new Error('x'), { chainName: 'sepolia' })).toMatch(/- chain: sepolia \(id \?\)/)
  })
})
