// Members × ciphertexts, coloured by decryption wave.
//
// A node schedules its i-th decryption attempt `i · staggerBlocks` after the
// ciphertext, so the wave a partial landed in is exactly how far down the
// stagger the committee had to go before `t` answers were in: an all-emerald
// column answered on the first pass, a cold one had to wait for stragglers.

import { EmptyState } from '~kit'
import { Matrix, waveColor, CHART_COLORS, type LegendItem, type MatrixCell as KitMatrixCell } from '~kit/charts'
import type { PartialMatrix } from '~indexer/selectors'
import { shortAddress } from '~lib/format'

export function ApplicationPartialMatrix({ matrix }: { matrix: PartialMatrix | null }) {
  if (!matrix || matrix.columns.length === 0 || matrix.rows.length === 0) {
    return (
      <EmptyState
        compact
        title='Nothing to plot'
        description='This application has no ciphertexts with partial decryptions yet.'
      />
    )
  }

  let maxWave = 0
  for (const row of matrix.cells) {
    for (const cell of row) if (cell) maxWave = Math.max(maxWave, cell.wave ?? 0)
  }
  const waves = maxWave + 1

  const cells: KitMatrixCell[] = []
  matrix.cells.forEach((row, r) => {
    row.forEach((cell, c) => {
      if (!cell) return
      const operator = matrix.rows[r]?.operator ?? ''
      cells.push({
        row: r,
        col: c,
        color: waveColor(cell.wave ?? 0, waves),
        detail: (
          <div className='font-mono text-[10px]'>
            <div className='text-ghost'>
              #{cell.participantIndex} {shortAddress(operator)}
            </div>
            <div className='text-ash'>
              ciphertext {cell.ciphertextIndex} · block {cell.block} · wave {cell.wave}
            </div>
          </div>
        ),
      })
    })
  })

  const legend: LegendItem[] = [
    ...Array.from({ length: Math.min(waves, 6) }, (_, i) => ({
      label: `wave ${i}`,
      color: waveColor(i, waves),
    })),
    { label: 'no partial', color: CHART_COLORS.onyx },
  ]

  return (
    <Matrix
      rows={matrix.rows.map((row) => `#${row.participantIndex} ${shortAddress(row.operator)}`)}
      columns={matrix.columns.map((column) => String(column.ciphertextIndex))}
      cells={cells}
      rowLabelWidth={132}
      legend={legend}
    />
  )
}
