import { BarChart, type BarGroup } from './BarChart'
import { ChartPanel } from './ChartPanel'
import { SeriesContributions, SeriesPartials } from './palette'
import { shortHash } from '~lib/format'
import type { OperatorStats } from '~lib/operator-stats'

const SERIES = [SeriesContributions, SeriesPartials]

interface Props {
  rows: OperatorStats[]
  loading?: boolean
  /** Cap the number of bars; the leaderboard below carries the full list. */
  limit?: number
}

/**
 * Work done per operator since the deployment block: accepted contributions
 * (the key-generation half) next to published partial decryptions (the
 * decryption half). Both are self-attributing events, so no transaction
 * lookups are involved in these two numbers.
 */
export function OperatorActivityChart({ rows, loading, limit = 12 }: Props) {
  const groups: BarGroup[] = rows.slice(0, limit).map((r) => ({
    key: r.operator,
    label: shortHash(r.operator, 4, 4),
    sublabel: r.operator,
    values: { contributions: r.contributions, partials: r.partials },
  }))

  return (
    <ChartPanel
      title='Work per operator'
      caption='Accepted contributions and published partial decryptions per operator, over the whole history of this deployment.'
      loading={loading}
      height={190}
    >
      <BarChart
        series={SERIES}
        groups={groups}
        height={190}
        ariaLabel='Contributions and partial decryptions per operator'
        emptyMessage='No operator activity recorded yet.'
      />
    </ChartPanel>
  )
}
