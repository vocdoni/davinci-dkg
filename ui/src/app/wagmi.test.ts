import { describe, expect, it } from 'vitest'
import type { RuntimeConfig } from '~config/runtime-config'
import { blockTimeSeconds, chainFromConfig } from './wagmi'

const GNOSIS: RuntimeConfig = {
  rpcUrl: 'https://rpc.example/gnosis',
  managerAddress: '0x9999F38Ff8Bf959E98Ddd5D4551f82775219c01B',
  chainId: 100,
  chainName: 'gnosis',
  explorerUrl: 'https://gnosisscan.io',
  deployBlock: 48_483_860,
  demo: false,
}

describe('chainFromConfig', () => {
  it('gives Gnosis Chain its currency and block time, and the configured endpoints', () => {
    const chain = chainFromConfig(GNOSIS)
    expect(chain.id).toBe(100)
    expect(chain.nativeCurrency.symbol).toBe('XDAI')
    expect(chain.testnet).toBe(false)
    expect(chain.rpcUrls.default.http).toEqual(['https://rpc.example/gnosis'])
    expect(chain.blockExplorers?.default.url).toBe('https://gnosisscan.io')
    expect(blockTimeSeconds(100)).toBe(5)
  })

  it('falls back to an Ether testnet with the default block time for other chains', () => {
    const chain = chainFromConfig({ ...GNOSIS, chainId: 11155111, chainName: 'sepolia', explorerUrl: undefined })
    expect(chain.nativeCurrency.symbol).toBe('ETH')
    expect(chain.testnet).toBe(true)
    expect(chain.blockExplorers).toBeUndefined()
    expect(blockTimeSeconds(11155111)).toBeUndefined()
    expect(blockTimeSeconds(31337)).toBeUndefined()
  })
})
