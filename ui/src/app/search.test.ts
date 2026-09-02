import { describe, expect, it } from 'vitest'
import { resolveSearch, type SearchResolver } from './search'

const ctx = { explorerUrl: 'https://sepolia.etherscan.io/' }

describe('resolveSearch', () => {
  it('routes an address to the operator page', () => {
    const target = resolveSearch('0x3f9b338706a31f26d49159478015c8aaeab908ad', ctx)
    expect(target).toMatchObject({ kind: 'route' })
    expect(target.kind === 'route' && target.path).toBe('/operators/0x3f9b338706a31f26d49159478015c8aaeab908ad')
  })

  it('routes a bytes12 epoch id to the epoch page', () => {
    const target = resolveSearch('0x0102030405060708090a0b0c', ctx)
    expect(target.kind === 'route' && target.path).toBe('/epochs/0x0102030405060708090a0b0c')
  })

  it('sends a 32-byte value to the explorer as a transaction', () => {
    const target = resolveSearch(`0x${'ab'.repeat(32)}`, ctx)
    expect(target.kind).toBe('external')
    expect(target.kind === 'external' && target.url).toBe(`https://sepolia.etherscan.io/tx/0x${'ab'.repeat(32)}`)
  })

  it('sends a bare number to the explorer as a block', () => {
    const target = resolveSearch('11619019', ctx)
    expect(target.kind === 'external' && target.url).toBe('https://sepolia.etherscan.io/block/11619019')
  })

  it('degrades to unknown without an explorer', () => {
    expect(resolveSearch(`0x${'ab'.repeat(32)}`, {}).kind).toBe('unknown')
    expect(resolveSearch('12345', {}).kind).toBe('unknown')
  })

  it('is empty-safe and rejects nonsense', () => {
    expect(resolveSearch('   ', ctx).kind).toBe('unknown')
    expect(resolveSearch('not a thing', ctx).kind).toBe('unknown')
  })

  it('gives registered resolvers the first look', () => {
    const resolver: SearchResolver = (query) =>
      query.startsWith('0xab') ? { kind: 'route', path: '/applications/0xepoch/' + query, label: 'app' } : null
    const target = resolveSearch(`0x${'ab'.repeat(32)}`, ctx, [resolver])
    expect(target.kind).toBe('route')
  })

  it('trims whitespace around the query', () => {
    expect(resolveSearch('  0x0102030405060708090a0b0c  ', ctx).kind).toBe('route')
  })
})
