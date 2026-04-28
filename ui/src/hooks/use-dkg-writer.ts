import { useMemo } from 'react'
import { useWalletClient } from 'wagmi'
import { DKGWriter } from '@vocdoni/davinci-dkg-sdk'
import { useDkgClient } from './use-dkg-client'
import { useConfig } from '~providers/ConfigProvider'

// Returns a memoised DKGWriter once the user has connected a wallet via
// wagmi/RainbowKit. Returns null until then so callers can render
// connect-wallet placeholders cleanly.
export function useDkgWriter(): DKGWriter | null {
  const { publicClient } = useDkgClient()
  const config = useConfig()
  const { data: walletClient } = useWalletClient()

  return useMemo(() => {
    if (!walletClient) return null
    // See use-dkg-client.ts for why these `as never` casts are necessary
    // — duplicate viem packages from the link:../sdk dependency.
    return new DKGWriter({
      publicClient: publicClient as never,
      walletClient: walletClient as never,
      managerAddress: config.managerAddress,
      ...(config.registryAddress ? { registryAddress: config.registryAddress } : {}),
    })
  }, [publicClient, walletClient, config.managerAddress, config.registryAddress])
}
