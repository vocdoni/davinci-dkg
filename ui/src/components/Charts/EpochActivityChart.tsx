import type { EpochEntry } from '@vocdoni/davinci-dkg-sdk'
import { BarChart, type BarGroup } from './BarChart'
import { ChartPanel } from './ChartPanel'
import {
  SeriesCiphertexts,
  SeriesClaims,
  SeriesContributions,
  SeriesPartials,
} from './palette'

const SERIES = [SeriesClaims, SeriesContributions, SeriesCiphertexts, SeriesPartials]

interface Props {
  /** Newest-first, as `getRecentEpochs` returns them. */
  epochs: EpochEntry[]
  loading?: boolean
  /** How many epochs to plot, oldest on the left. */
  limit?: number
}

/**
 * Activity across the most recent epochs: how many slots the lottery handed
 * out, how many contributions came back, how many ciphertexts the epoch was
 * asked to hold and how many partial decryptions that produced.
 *
 * All four numbers are counters on the epoch record itself, so this chart costs
 * no extra RPC — it re-reads the same query the epoch list uses.
 */
export function EpochActivityChart({ epochs, loading, limit = 10 }: Props) {
  const groups: BarGroup[] = epochs
    .slice(0, limit)
    .reverse()
    .map(({ id, epoch }) => ({
      key: id,
      label: `#${epoch.nonce.toString()}`,
      sublabel: `block #${epoch.startBlock.toString()}`,
      values: {
        claims: epoch.claimedCount,
        contributions: epoch.contributionCount,
        ciphertexts: epoch.ciphertextCount,
        partials: epoch.partialDecryptionCount,
      },
    }))

  return (
    <ChartPanel
      title={`Activity over the last ${Math.min(limit, Math.max(groups.length, 1))} epochs`}
      caption='Per epoch, oldest on the left. Claims and contributions are the key generation; ciphertexts and partial decryptions are the work that key then did.'
      loading={loading}
      height={180}
    >
      <BarChart
        series={SERIES}
        groups={groups}
        height={180}
        ariaLabel='Slots claimed, contributions, ciphertexts and partial decryptions per epoch'
        emptyMessage='No epochs on this deployment yet.'
      />
    </ChartPanel>
  )
}
