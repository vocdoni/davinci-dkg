import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useConfig } from '~providers/ConfigProvider'
import { QueryKeys } from './keys'
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
    // ~10k blocks). Falls back to 0 when startBlock isn't configured.
    queryFn: () => dkg.getRegistryNodes(config.startBlock ? BigInt(config.startBlock) : 0n),
    refetchInterval: Polling.default,
    staleTime: Polling.default / 2,
  })
}

export function useRoundCount() {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['chain', 'roundNonce'],
    queryFn: () => dkg.roundNonce(),
    refetchInterval: Polling.default,
  })
}
