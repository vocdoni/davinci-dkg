import { describe, expect, it } from 'vitest'
import { isDemoRequested, parseRuntimeConfig } from './runtime-config'

const valid = {
  rpcUrl: 'https://rpc.example/sepolia',
  managerAddress: '0x3f9b338706a31f26d49159478015c8aaeab908ad',
  chainId: 11155111,
  chainName: 'sepolia',
  deployBlock: 11619019,
  explorerUrl: 'https://sepolia.etherscan.io/',
}

describe('parseRuntimeConfig', () => {
  it('accepts the shape render-ui-config.sh writes', () => {
    const cfg = parseRuntimeConfig(valid)
    expect(cfg.chainId).toBe(11155111)
    expect(cfg.deployBlock).toBe(11619019)
  })

  it('strips the trailing slash off the explorer URL', () => {
    expect(parseRuntimeConfig(valid).explorerUrl).toBe('https://sepolia.etherscan.io')
  })

  it('defaults deployBlock to 0 when absent', () => {
    const { deployBlock: _omit, ...rest } = valid
    expect(parseRuntimeConfig(rest).deployBlock).toBe(0)
  })

  it('rejects a bad manager address', () => {
    expect(() => parseRuntimeConfig({ ...valid, managerAddress: '0xnope' })).toThrow(/managerAddress/)
  })

  it('rejects a missing rpcUrl', () => {
    expect(() => parseRuntimeConfig({ ...valid, rpcUrl: '' })).toThrow(/rpcUrl/)
  })
})

describe('isDemoRequested', () => {
  it.each(['?demo=1', '?demo', '?demo=true', '?foo=1&demo=1'])('is true for %s', (search) => {
    expect(isDemoRequested(search)).toBe(true)
  })

  it.each(['', '?demo=0', '?demo=false', '?other=1'])('is false for "%s"', (search) => {
    expect(isDemoRequested(search)).toBe(false)
  })
})
