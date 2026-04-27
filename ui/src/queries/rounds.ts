import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { QueryKeys } from './keys'
import { Polling } from '~constants/polling'

export function useRecentRounds(limit = 20) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.roundsRecent(limit),
    queryFn: () => dkg.getRecentRounds(limit),
    refetchInterval: Polling.default,
  })
}

export function useRound(id: `0x${string}` | undefined) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: id ? QueryKeys.round(id) : ['rounds', 'idle'],
    queryFn: async () => {
      if (!id) throw new Error('round id required')
      const [round, participants] = await Promise.all([dkg.getRound(id), dkg.selectedParticipants(id)])
      return { round, participants }
    },
    enabled: Boolean(id),
    refetchInterval: Polling.default,
  })
}

export function useRoundEvents(id: `0x${string}` | undefined, fromBlock?: bigint) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: id ? QueryKeys.roundEvents(id, fromBlock) : ['rounds', 'events', 'idle'],
    queryFn: () => {
      if (!id) throw new Error('round id required')
      return dkg.getAllRoundEvents(id, fromBlock ?? 0n)
    },
    enabled: Boolean(id),
    refetchInterval: Polling.default,
  })
}
