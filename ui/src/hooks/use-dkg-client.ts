import { useMemo } from 'react'
import { createPublicClient, defineChain, http, type PublicClient } from 'viem'
import { DKGClient } from '@vocdoni/davinci-dkg-sdk'
import { useConfig } from '~providers/ConfigProvider'

// Returns a memoised DKGClient + viem PublicClient for the active runtime
// config. Re-created when the config object identity changes (i.e. only on
// boot for now; an RPC-override mechanism in Settings will swap the config
// in place to retrigger).
//
// The `publicClient as never` cast on the DKGClient call is load-bearing:
// `link:../sdk` makes sdk a sibling pnpm package with its OWN viem in
// sdk/node_modules, so the SDK's `PublicClient` type lives in a different
// .pnpm hash from the UI's. Both shapes are structurally identical at
// runtime (same viem major), but tsc reports them as distinct types. The
// cast tells tsc to trust us; it can be removed once we move sdk + ui
// into a shared pnpm workspace and viem is hoisted to a single instance.
export function useDkgClient() {
  const config = useConfig()
  return useMemo(() => {
    const chain = defineChain({
      id: config.chainId,
      name: config.chainName || `chain-${config.chainId}`,
      nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 },
      rpcUrls: { default: { http: [config.rpcUrl] } },
    })
    const publicClient = createPublicClient({ chain, transport: http(config.rpcUrl) }) as PublicClient
    const dkg = new DKGClient({
      publicClient: publicClient as never,
      managerAddress: config.managerAddress,
      ...(config.registryAddress ? { registryAddress: config.registryAddress } : {}),
    })
    return { dkg, publicClient, chain }
  }, [config])
}
