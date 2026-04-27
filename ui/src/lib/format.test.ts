import { describe, it, expect } from 'vitest'
import { bigIntToHex, blocksRemaining, blocksToDuration, shortAddress, shortHash } from './format'

describe('shortHash', () => {
  it('returns empty for null/undefined', () => {
    expect(shortHash(null)).toBe('')
    expect(shortHash(undefined)).toBe('')
  })

  it('returns the value unchanged for non-hex strings', () => {
    expect(shortHash('hello')).toBe('hello')
  })

  it('returns the value unchanged when shorter than head+tail+2', () => {
    expect(shortHash('0xabcd', 6, 4)).toBe('0xabcd')
  })

  it('truncates with default head/tail', () => {
    const long = '0x' + 'a'.repeat(64)
    const out = shortHash(long)
    expect(out.startsWith('0xaaaaaa')).toBe(true)
    expect(out.endsWith('aaaa')).toBe(true)
    expect(out).toContain('…')
  })
})

describe('shortAddress', () => {
  it('truncates an Ethereum address to 4+4', () => {
    const addr = '0x1234567890abcdef1234567890abcdef12345678'
    expect(shortAddress(addr)).toBe('0x1234…5678')
  })
})

describe('bigIntToHex', () => {
  it('zero-pads to 64 hex chars', () => {
    expect(bigIntToHex(0n)).toBe('0x' + '0'.repeat(64))
    expect(bigIntToHex(255n)).toBe('0x' + '0'.repeat(62) + 'ff')
  })
})

describe('blocksToDuration', () => {
  it('returns "now" for zero/negative', () => {
    expect(blocksToDuration(0)).toBe('now')
    expect(blocksToDuration(-5)).toBe('now')
  })
  it('formats sub-minute', () => {
    expect(blocksToDuration(2)).toBe('~24s')
  })
  it('formats minutes', () => {
    expect(blocksToDuration(10)).toBe('~2 min')
  })
  it('formats hours', () => {
    expect(blocksToDuration(450)).toMatch(/h$/)
  })
  it('formats days for very large block counts', () => {
    expect(blocksToDuration(10_000)).toMatch(/d$|h$/)
  })
})

describe('blocksRemaining', () => {
  it('returns null when current is missing', () => {
    expect(blocksRemaining(null, 10)).toBeNull()
    expect(blocksRemaining(undefined, 10)).toBeNull()
  })
  it('returns null when target is missing', () => {
    expect(blocksRemaining(5n, undefined)).toBeNull()
  })
  it('clamps to zero when target is in the past', () => {
    expect(blocksRemaining(20n, 10n)).toBe(0)
  })
  it('returns the positive delta', () => {
    expect(blocksRemaining(10n, 25n)).toBe(15)
    expect(blocksRemaining(10n, 25)).toBe(15) // mixed bigint/number target
  })
})
