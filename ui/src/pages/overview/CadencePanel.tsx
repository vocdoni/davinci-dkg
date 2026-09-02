import { useMemo } from 'react'
import { useNavigate } from 'react-router-dom'
import { Panel } from '~kit'
import { CadenceStrip, type CadenceEpoch } from '~kit/charts'
import type { EpochRow } from '~indexer/selectors'
import { paths } from '~routes/paths'
import { PHASE_COLOR_KEY } from '~pages/epochs/phase'
import { blockOrNull } from '~pages/epochs/cadence'

/**
 * Epochs on the block axis. Overlap is the point: a Live epoch keeps serving
 * while the next one runs its lottery, so the strip is where the cadence — and
 * any gap in it — becomes visible.
 */
export function CadencePanel({ rows, head, loading }: { rows: EpochRow[]; head: number; loading: boolean }) {
  const navigate = useNavigate()
  const current = blockOrNull(head)

  const epochs = useMemo<CadenceEpoch[]>(() => {
    // Oldest first: the strip packs epochs into lanes by ascending start block.
    const ordered = [...rows].sort((a, b) => a.startBlock - b.startBlock)
    return ordered.map((row) => {
      const start = row.startBlock
      const end = row.endBlock ?? Math.max(current ?? start, start + 1)
      return {
        id: row.id,
        label: `#${row.nonce}`,
        start,
        end: Math.max(end, start + 1),
        phase: PHASE_COLOR_KEY[row.phase],
        detail: `${row.contributions}/${row.committeeSize} contributions · ${row.ciphertexts} ciphertexts`,
      }
    })
  }, [rows, current])

  return (
    <Panel
      label='Network'
      title='Epoch cadence'
      description='Every epoch on a shared block axis, coloured by phase. The dashed line is the chain head.'
    >
      <CadenceStrip
        epochs={epochs}
        current={current}
        height={132}
        loading={loading}
        onEpochClick={(epoch) => navigate(paths.epoch(epoch.id))}
      />
    </Panel>
  )
}
