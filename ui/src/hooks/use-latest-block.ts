import { useEffect, useState, useSyncExternalStore } from 'react'
import { useBlockNumber } from 'wagmi'
import { useRuntimeConfig } from '~config/config-context'
import { useOptionalDataSource } from '~data/context'

/** Chain head, polled. In demo mode it ticks off a synthetic clock instead. */
export interface LatestBlock {
  blockNumber: bigint | null
  isLoading: boolean
  isError: boolean
}

/** Nominal L1 block time; also the demo clock's tick. */
export const BLOCK_TIME_MS = 12_000

/** Where the demo clock starts, so screenshots have plausible block numbers. */
const DEMO_ORIGIN_BLOCK = 11_900_000n

const noSubscribe = () => () => {}

function useDemoBlock(): bigint {
  // Anchor the synthetic clock on the fixture's head block so the top bar agrees with every page.
  const source = useOptionalDataSource()
  const head = useSyncExternalStore(source?.subscribe ?? noSubscribe, () => source?.getSnapshot().status.headBlock ?? null)
  const [tick, setTick] = useState(0)
  useEffect(() => {
    const id = setInterval(() => setTick((t) => t + 1), BLOCK_TIME_MS)
    return () => clearInterval(id)
  }, [])
  return (head && Number.isFinite(head) ? BigInt(head) : DEMO_ORIGIN_BLOCK) + BigInt(tick)
}

/**
 * The chain head. Every countdown, phase window and "live since" in the app
 * reads it from here so they all agree on the same block within a render.
 */
export function useLatestBlock(): LatestBlock {
  const { demo } = useRuntimeConfig()
  const demoBlock = useDemoBlock()
  const query = useBlockNumber({ watch: !demo, query: { enabled: !demo, refetchInterval: BLOCK_TIME_MS } })

  if (demo) return { blockNumber: demoBlock, isLoading: false, isError: false }
  return { blockNumber: query.data ?? null, isLoading: query.isLoading, isError: query.isError }
}
