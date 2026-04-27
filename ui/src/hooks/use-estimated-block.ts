import { useEffect, useRef, useState } from 'react'
import { useBlockNumber } from '~queries/chain'

// Returns a 1Hz-ticking estimate of the current block height. Between RPC
// fetches it extrapolates: estimate = lastFetched + ⌊elapsed/blockTime⌋.
//
// The block-time default of 12s matches Ethereum mainnet; on Anvil/devnets
// blocks come faster but for a *countdown* the resulting display ticks
// down monotonically anyway — it just resyncs to the truth on the next
// 4s React-Query refetch.
//
// Returns null until the first real fetch lands; callers should treat
// that as "loading".
export function useEstimatedBlock(blockTimeSeconds = 12): bigint | null {
  const { data: fetched, dataUpdatedAt } = useBlockNumber()
  const [now, setNow] = useState<number>(() => Date.now())
  const tick = useRef<ReturnType<typeof setInterval> | null>(null)

  useEffect(() => {
    tick.current = setInterval(() => setNow(Date.now()), 1000)
    return () => {
      if (tick.current) clearInterval(tick.current)
    }
  }, [])

  if (fetched == null || dataUpdatedAt === 0) return null
  const elapsedSec = Math.max(0, Math.floor((now - dataUpdatedAt) / 1000))
  const blocksGuess = BigInt(Math.floor(elapsedSec / blockTimeSeconds))
  return fetched + blocksGuess
}
