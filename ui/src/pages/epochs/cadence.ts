// Block-clock arithmetic: "what is this epoch waiting for, and how far away is
// it?". Pure functions over the store's own types so the maths is unit-tested
// against the fixture rather than eyeballed in a screenshot.
//
// Everything here tolerates a non-finite head or target: the indexer publishes
// a snapshot before the chain head is known, and a countdown must degrade to
// "—" instead of rendering NaN.

import type { NetworkStats } from '~indexer/selectors'
import type { EpochEntity } from '~indexer/types'
import { blocksToDuration } from '~lib/format'

export interface Countdown {
  /** What the clock is counting down to. */
  label: string
  targetBlock: number
  /** Blocks still to go; 0 once the target is reached. */
  blocks: number
  /** True when the target block is already behind the head. */
  passed: boolean
}

export interface NextEpochCountdown extends Countdown {
  /**
   * `chain` when `nextEpochStartBlock()` was read from the manager, `cadence`
   * when it was derived from the newest epoch's start plus the epoch duration.
   */
  source: 'chain' | 'cadence'
}

function finite(value: number | null | undefined): value is number {
  return typeof value === 'number' && Number.isFinite(value)
}

/**
 * `null` for anything that is not a real block height. Every block a page
 * prints goes through this: a snapshot published before the head is known must
 * render an em dash, never `NaN`.
 */
export function blockOrNull(value: number | null | undefined): number | null {
  return finite(value) ? value : null
}

/** Blocks between `head` and `targetBlock`, or null when either is unknown. */
export function countdownTo(label: string, targetBlock: number | null | undefined, head: number): Countdown | null {
  if (!finite(targetBlock) || !finite(head)) return null
  const remaining = targetBlock - head
  return { label, targetBlock, blocks: Math.max(0, remaining), passed: remaining <= 0 }
}

/**
 * The deadline the epoch's current phase is racing: the selection window, the
 * assembly window, or the end of the service window. Terminal phases (aborted,
 * closed) have nothing to wait for.
 */
export function phaseCountdown(
  epoch: EpochEntity | null | undefined,
  head: number,
  epochDurationBlocks: number | null | undefined
): Countdown | null {
  if (!epoch) return null
  switch (epoch.status) {
    case 'committee-selection':
      return countdownTo('selection closes', epoch.policy?.committeeSelectionDeadlineBlock, head)
    case 'key-assembly':
      return countdownTo('assembly closes', epoch.policy?.keyAssemblyDeadlineBlock, head)
    case 'live':
      return countdownTo('epoch ends', finite(epochDurationBlocks) ? epoch.startBlock + epochDurationBlocks : null, head)
    default:
      return null
  }
}

/**
 * When `createEpoch` may fire again. Prefers the manager's own
 * `nextEpochStartBlock`; falls back to the newest epoch's cadence anchor
 * (`startBlock + EPOCH_DURATION_BLOCKS`), which is how the contract computes it.
 */
export function nextEpochCountdown(stats: NetworkStats): NextEpochCountdown | null {
  const chain = countdownTo('next epoch', stats.nextEpochStartBlock, stats.headBlock)
  if (chain) return { ...chain, source: 'chain' }
  const newest = stats.newestEpoch
  if (!newest || !finite(stats.epochDurationBlocks)) return null
  const cadence = countdownTo('next epoch', newest.startBlock + stats.epochDurationBlocks, stats.headBlock)
  return cadence ? { ...cadence, source: 'cadence' } : null
}

/**
 * `"128 blocks · ~26 min"`. Once the target block is behind the head there is
 * nothing to count, so the caller picks the tense: a deadline has `passed`, a
 * cadence anchor is `due`.
 */
export function formatCountdown(countdown: Countdown | null, blockTimeSeconds = 12, passedText = 'due'): string {
  if (!countdown) return '—'
  if (countdown.passed) return passedText
  return `${countdown.blocks.toLocaleString()} blocks · ${blocksToDuration(countdown.blocks, blockTimeSeconds)}`
}

/** How long an epoch has been live, as a block delta and a duration. */
export function elapsedSince(
  block: number | null | undefined,
  head: number,
  blockTimeSeconds = 12
): { blocks: number; text: string } | null {
  if (!finite(block) || !finite(head)) return null
  const blocks = Math.max(0, head - block)
  return { blocks, text: blocks === 0 ? 'just now' : `${blocksToDuration(blocks, blockTimeSeconds)} ago` }
}

/**
 * Words in the finalize transcript the BRLC challenge is taken over:
 * `N` participant indexes + `2·N²` contribution-commitment coordinates +
 * `2·N` aggregate-commitment coordinates + `2·N` share-commitment coordinates
 * = `2·N² + 5·N`. Each word is a 32-byte field element.
 */
export function transcriptWords(committeeSize: number): number {
  if (!finite(committeeSize) || committeeSize <= 0) return 0
  return 2 * committeeSize * committeeSize + 5 * committeeSize
}
