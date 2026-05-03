import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { QueryKeys } from './keys'
import { Polling } from '~constants/polling'

export function useRecentEpochs(limit = 20) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.epochsRecent(limit),
    queryFn: () => dkg.getRecentRounds(limit),
    refetchInterval: Polling.default,
  })
}

export function useEpoch(id: `0x${string}` | undefined) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: id ? QueryKeys.epoch(id) : ['epochs', 'idle'],
    queryFn: async () => {
      if (!id) throw new Error('epoch id required')
      const [epoch, participants] = await Promise.all([dkg.getEpoch(id), dkg.selectedParticipants(id)])
      return { epoch, participants }
    },
    enabled: Boolean(id),
    refetchInterval: Polling.default,
  })
}

export function useEpochEvents(id: `0x${string}` | undefined, fromBlock?: bigint) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: id ? QueryKeys.epochEvents(id, fromBlock) : ['epochs', 'events', 'idle'],
    queryFn: () => {
      if (!id) throw new Error('epoch id required')
      return dkg.getAllEpochEvents(id, fromBlock ?? 0n)
    },
    enabled: Boolean(id),
    refetchInterval: Polling.default,
  })
}
