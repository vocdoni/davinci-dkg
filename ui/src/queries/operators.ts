import { useQuery, useQueryClient } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useConfig } from '~providers/ConfigProvider'
import { QueryKeys } from './keys'
import { scanFromBlock } from '~lib/chain-scan'
import { aggregateOperatorStats, type OperatorStats } from '~lib/operator-stats'

/**
 * How often the operator leaderboard refetches. Much slower than the rest of
 * the explorer on purpose: this is five historical log scans plus one
 * `eth_getTransaction` per finalization and combine, so it is by far the most
 * expensive read on the page and its numbers move once per epoch at most.
 */
const REFETCH_MS = 60_000

/**
 * Per-operator activity since the manager's deployment block.
 *
 * Five chunked log scans (`SlotClaimed`, `ContributionSubmitted`,
 * `PartialDecryptionSubmitted`, `EpochLive`, `DecryptionCombined`) plus the
 * registry roster, folded into the leaderboard by `aggregateOperatorStats`.
 * `EpochLive` and `DecryptionCombined` do not name their submitter, so their
 * transactions are fetched (de-duplicated, one round trip each) to attribute
 * finalizations and combines.
 */
export function useOperatorStats() {
  const { dkg } = useDkgClient()
  const config = useConfig()
  const queryClient = useQueryClient()
  const fromBlock = scanFromBlock(config)

  return useQuery<OperatorStats[]>({
    queryKey: QueryKeys.operatorStats(fromBlock),
    queryFn: async () => {
      const [nodes, claims, contributions, partials, finalizations, combines] = await Promise.all([
        // Routed through the shared cache entry so the roster is scanned once
        // for both this leaderboard and the plain node list.
        queryClient.fetchQuery({
          queryKey: QueryKeys.registryNodes,
          queryFn: () => dkg.getRegistryNodes(fromBlock),
          staleTime: REFETCH_MS / 2,
        }),
        dkg.getSlotClaimedEvents({ fromBlock }),
        dkg.getContributionSubmittedEvents({ fromBlock }),
        dkg.getPartialDecryptionEvents({ fromBlock }),
        dkg.getAllEpochLiveEvents({ fromBlock }),
        dkg.getAllDecryptionCombinedEvents({ fromBlock }),
      ])
      const senders = await dkg.getTransactionSenders([
        ...finalizations.map((e) => e.transactionHash),
        ...combines.map((e) => e.transactionHash),
      ])
      return aggregateOperatorStats({
        nodes,
        claims,
        contributions,
        partials,
        finalizations,
        combines,
        senders,
      })
    },
    refetchInterval: REFETCH_MS,
    staleTime: REFETCH_MS / 2,
  })
}
