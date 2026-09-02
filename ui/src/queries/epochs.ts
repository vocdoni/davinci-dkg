import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useConfig } from '~providers/ConfigProvider'
import { QueryKeys } from './keys'
import { Polling } from '~constants/polling'
import { scanFromBlock } from '~lib/chain-scan'
import { aggregateDecryptionProgress } from '~lib/decryption-overview'

export function useRecentEpochs(limit = 20) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.epochsRecent(limit),
    queryFn: () => dkg.getRecentEpochs(limit),
    refetchInterval: Polling.default,
  })
}

// Deploy-time createEpoch bounds (MIN_THRESHOLD / MIN_COMMITTEE_SIZE /
// MAX_LOTTERY_ALPHA_BPS). Immutables on the manager, so never refetched.
export function useEpochBounds() {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.epochBounds,
    queryFn: () => dkg.getEpochBounds(),
    staleTime: Infinity,
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

/**
 * The manager's `EPOCH_DURATION_BLOCKS` immutable. Needed to place the "Live"
 * window on the block-axis timeline: the epoch's last block is
 * `startBlock + duration - 1`.
 */
export function useEpochDuration() {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: QueryKeys.epochDuration,
    queryFn: () => dkg.getEpochDurationBlocks(),
    staleTime: Infinity,
  })
}

/**
 * Per-application decryption progress for one epoch, reconstructed from the
 * five event streams the pipeline leaves behind. Enabled only for Live epochs
 * — before that there are no applications to show and the scans would be pure
 * cost.
 */
export function useEpochDecryptionProgress(id: `0x${string}` | undefined, enabled = true) {
  const { dkg } = useDkgClient()
  const config = useConfig()
  const fromBlock = scanFromBlock(config)

  return useQuery({
    queryKey: id ? QueryKeys.epochDecryption(id) : ['epochs', 'decryptionProgress', 'idle'],
    queryFn: async () => {
      if (!id) throw new Error('epoch id required')
      const [applications, ciphertexts, partials, combines] = await Promise.all([
        dkg.getApplicationRegisteredEvents({ epochId: id, fromBlock }),
        dkg.getCiphertextSubmittedEvents(id),
        dkg.getPartialDecryptionEvents({ epochId: id, fromBlock }),
        dkg.getAllDecryptionCombinedEvents({ epochId: id, fromBlock }),
      ])
      // Organizer shares live on the app manager and are only indexed by
      // (epochId, aid), so they need one scan per application. There are a
      // handful of applications per epoch in practice.
      const shares = (
        await Promise.all(
          applications.map(async (a) =>
            (await dkg.getOrganizerShareEvents(id, a.aid)).map((s) => ({
              aid: a.aid,
              ciphertextIndex: s.ciphertextIndex,
            })),
          ),
        )
      ).flat()
      return aggregateDecryptionProgress({ applications, ciphertexts, partials, shares, combines })
    },
    enabled: Boolean(id) && enabled,
    refetchInterval: Polling.decryption * 2,
  })
}

/**
 * The committee members that have published a partial decryption for one
 * ciphertext. `PartialDecryptionSubmitted` names its submitter, so this is the
 * one place the UI can show *who* is answering — the epoch-wide counter on the
 * epoch record cannot be split per ciphertext.
 */
export function useCiphertextPartials(
  id: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
  ciphertextIndex: number | null | undefined,
) {
  const { dkg } = useDkgClient()
  const config = useConfig()
  const fromBlock = scanFromBlock(config)

  return useQuery({
    queryKey:
      id && aid && ciphertextIndex != null
        ? QueryKeys.ciphertextPartials(id, aid, ciphertextIndex)
        : ['epochs', 'partials', 'idle'],
    queryFn: () => {
      if (!id || !aid || ciphertextIndex == null) throw new Error('epochId + aid + index required')
      return dkg.getPartialDecryptionEvents({ epochId: id, aid, ciphertextIndex, fromBlock })
    },
    enabled: Boolean(id && aid && ciphertextIndex != null),
    refetchInterval: Polling.decryption,
  })
}
