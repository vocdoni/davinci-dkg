import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { QueryKeys } from './keys'
import { Polling } from '~constants/polling'

export function useBlockNumber() {
  const { publicClient } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.blockNumber,
    queryFn: () => publicClient.getBlockNumber(),
    refetchInterval: Polling.blockNumber,
    staleTime: Polling.blockNumber / 2,
  })
}
