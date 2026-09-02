import { useMemo } from 'react'
import { useNavigate } from 'react-router-dom'
import { Panel } from '~kit'
import { StackedBars, type BarDatum, type BarSeries } from '~kit/charts'
import type { ActivityBucket } from '~indexer/selectors'
import { paths } from '~routes/paths'

const SERIES: BarSeries[] = [
  { key: 'claims', label: 'claims' },
  { key: 'contributions', label: 'contributions' },
  { key: 'ciphertexts', label: 'ciphertexts' },
  { key: 'partials', label: 'partials' },
]

/**
 * What the network actually did, epoch by epoch. The four series are the four
 * transactions the protocol asks for, so a short bar is a phase that nobody
 * showed up for.
 */
export function ActivityPanel({ buckets, loading }: { buckets: ActivityBucket[]; loading: boolean }) {
  const navigate = useNavigate()
  const data = useMemo<BarDatum[]>(
    () =>
      buckets.map((bucket) => ({
        label: `#${bucket.nonce}`,
        values: {
          claims: bucket.claims,
          contributions: bucket.contributions,
          ciphertexts: bucket.ciphertexts,
          partials: bucket.partials,
        },
        note: `${bucket.phase} · start block ${bucket.startBlock.toLocaleString()}`,
      })),
    [buckets]
  )

  return (
    <Panel
      label='Network'
      title='Activity per epoch'
      description={`Claims, contributions, ciphertexts and partial decryptions over the last ${buckets.length || 30} epochs.`}
    >
      <StackedBars
        data={data}
        series={SERIES}
        height={220}
        loading={loading}
        onBarClick={(_, index) => {
          const bucket = buckets[index]
          if (bucket) navigate(paths.epoch(bucket.epoch))
        }}
      />
    </Panel>
  )
}
