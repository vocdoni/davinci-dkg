import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from './use-dkg-client'

// Number of recent blocks sampled to estimate block time. Larger sample =
// smoother estimate but more RPC traffic. Six is a reasonable default for
// a chain with stable cadence.
const SAMPLE_BLOCKS = 6n
// Treat the estimate as fresh for 5 minutes — block time on a given chain
// is effectively constant so we don't need to re-sample frequently. We only
// set staleTime (no refetchInterval), so the value is recomputed on the
// next mount or query invalidation after the window elapses.
const STALE_TIME_MS = 5 * 60_000

const FALLBACK_BLOCK_TIME_S = 12

/**
 * Returns the estimated wall-clock seconds per block for the connected
 * chain, sampled by reading two recent block headers and dividing the
 * timestamp delta by the block-number delta. Falls back to 12s (Ethereum
 * mainnet cadence) if the sample fails or the chain has fewer than
 * `SAMPLE_BLOCKS` blocks.
 *
 * Cached for 5 minutes via React Query, so the cost amortises across every
 * countdown component on the page.
 */
export function useBlockTimeSeconds(): number {
  const client = useDkgClient()
  const { data } = useQuery({
    queryKey: ['blockTimeSeconds', client?.chain?.id],
    enabled: !!client,
    staleTime: STALE_TIME_MS,
    queryFn: async () => {
      if (!client) return FALLBACK_BLOCK_TIME_S
      const head = await client.publicClient.getBlockNumber()
      if (head <= SAMPLE_BLOCKS) return FALLBACK_BLOCK_TIME_S
      const [latest, earlier] = await Promise.all([
        client.publicClient.getBlock({ blockNumber: head }),
        client.publicClient.getBlock({ blockNumber: head - SAMPLE_BLOCKS }),
      ])
      const dt = Number(latest.timestamp - earlier.timestamp)
      const dn = Number(head - (head - SAMPLE_BLOCKS))
      if (dt <= 0 || dn <= 0) return FALLBACK_BLOCK_TIME_S
      return dt / dn
    },
  })
  return data ?? FALLBACK_BLOCK_TIME_S
}
