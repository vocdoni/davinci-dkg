import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { Polling } from '~constants/polling'

export function useCollectivePublicKey(epochId: `0x${string}` | undefined) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['collectivePublicKey', epochId],
    queryFn: () => {
      if (!epochId) throw new Error('epochId required')
      return dkg.getCollectivePublicKey(epochId)
    },
    enabled: Boolean(epochId),
    refetchInterval: Polling.default,
  })
}

// Per-application reads. The cache key includes (epochId, aid) so that
// multiple AppRegistrationForms / DecryptionPipelines on the same page
// don't collide.

export function useApplication(
  epochId: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['application', epochId, aid],
    queryFn: () => {
      if (!epochId || !aid) throw new Error('epochId + aid required')
      return dkg.getApplication(epochId, aid)
    },
    enabled: Boolean(epochId && aid),
    refetchInterval: Polling.default,
  })
}
