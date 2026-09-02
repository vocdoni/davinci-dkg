import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useConfig } from '~providers/ConfigProvider'
import { QueryKeys } from './keys'
import { scanFromBlock } from '~lib/chain-scan'
import { Polling } from '~constants/polling'

export function useRegistryStats() {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.registryStats,
    queryFn: async () => {
      const [active, total, inactivity] = await Promise.all([
        dkg.activeCount(),
        dkg.nodeCount(),
        dkg.inactivityWindow(),
      ])
      return { active, total, inactivity }
    },
    refetchInterval: Polling.default,
  })
}

export function useRegistryNodes() {
  const { dkg } = useDkgClient()
  const config = useConfig()
  return useQuery({
    queryKey: QueryKeys.registryNodes,
    // The SDK takes a fromBlock for the event scan; using the manager's
    // deployment block keeps free-tier RPCs happy (most cap getLogs at
    // ~10k blocks). `useOperatorStats` fetches through this same key, so the
    // roster is scanned once no matter how many views want it.
    queryFn: () => dkg.getRegistryNodes(scanFromBlock(config)),
    refetchInterval: Polling.default,
    staleTime: Polling.default / 2,
  })
}

export function useEpochCount() {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['chain', 'epochNonce'],
    queryFn: () => dkg.epochNonce(),
    refetchInterval: Polling.default,
  })
}
