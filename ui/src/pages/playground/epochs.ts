// Which epochs the walkthrough may run against, straight off the indexer.
//
// The epoch step needs three things per Live epoch — the committee's `t of n`,
// the block it went Live, and how much of its pool is left — and one fallback for the
// case that matters most on a quiet deployment: no epoch is Live at all, in
// which case the step shows the newest epoch's phase and how far away the next
// one is instead of an empty list.

import { useMemo } from 'react'
import { useEpochs, useNetworkStats, useStore } from '~data/hooks'
import { resolveHeadBlock } from './head-block'
import type { EpochOption } from './controller'
import type { EpochStepData } from './steps'

export interface EpochStepInputs {
  data: EpochStepData
  options: EpochOption[]
}

export function useEpochStepData(): EpochStepInputs {
  const store = useStore()
  const live = useEpochs({ phase: 'live' })
  const newest = useEpochs({ limit: 1 })
  const stats = useNetworkStats()

  return useMemo(() => {
    const options: EpochOption[] = live.map((row) => ({
      id: row.id,
      nonce: row.nonce,
      threshold: row.threshold,
      committeeSize: row.committeeSize,
      liveSinceBlock: row.liveSinceBlock,
      poolClaimed: row.poolClaimed,
      poolFree: row.poolFree,
    }))
    const head = newest[0] ?? null
    const headBlock = resolveHeadBlock(store)
    const next = stats.nextEpochStartBlock
    return {
      options,
      data: {
        live: options,
        newest: head
          ? {
              id: head.id,
              nonce: head.nonce,
              phase: head.phase,
              startBlock: head.startBlock,
              endBlock: head.endBlock,
            }
          : null,
        headBlock,
        blockTimeSeconds: stats.blockTimeSeconds,
        nextEpochStartBlock: next,
        // `networkStats` derives this from `chain.headBlock`, which the
        // fixture leaves as NaN; recompute it off the head we trust.
        blocksToNextEpoch: next != null ? Math.max(0, next - headBlock) : null,
      },
    }
  }, [live, newest, store, stats.blockTimeSeconds, stats.nextEpochStartBlock])
}
