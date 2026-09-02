import { useMemo } from 'react'
import { EmptyState, Panel } from '~kit'
import { CHART_COLORS, Matrix, waveColor, type LegendItem, type MatrixCell } from '~kit/charts'
import type { ApplicationRow, PartialMatrix } from '~indexer/selectors'
import { shortAddress, shortHash } from '~lib/format'

/** Marker colours for the two rows that are not committee members. */
const SHARE_COLOR = CHART_COLORS.teal
const COMBINED_COLOR = CHART_COLORS.emerald

export interface DecryptionMatrixProps {
  matrix: PartialMatrix | null
  applications: ApplicationRow[]
}

/**
 * Members × ciphertexts. A cell is a published partial decryption, coloured by
 * the wave it landed in — the node schedules its i-th attempt `i·stagger`
 * blocks after the ciphertext, so emerald answered first and slate trailed.
 * The last two rows are the organizer share and the combined plaintext, which
 * is what turns t partials into a decryption.
 */
export function DecryptionMatrix({ matrix, applications }: DecryptionMatrixProps) {
  const aidLetters = useMemo(() => {
    const letters = new Map<string, string>()
    applications.forEach((app, i) => letters.set(app.aid.toLowerCase(), String.fromCharCode(65 + (i % 26))))
    return letters
  }, [applications])

  const model = useMemo(() => {
    if (!matrix || matrix.columns.length === 0 || matrix.rows.length === 0) return null

    const label = (aid: string, index: number) => `${aidLetters.get(aid.toLowerCase()) ?? '?'}#${index}`
    let waves = 1
    for (const line of matrix.cells) {
      for (const cell of line) {
        if (cell && Number.isFinite(cell.wave) && cell.wave != null) waves = Math.max(waves, cell.wave + 1)
      }
    }

    const cells: MatrixCell[] = []
    matrix.cells.forEach((line, slot) => {
      line.forEach((cell, column) => {
        if (!cell) return
        const wave = Number.isFinite(cell.wave) && cell.wave != null ? cell.wave : 0
        const member = matrix.rows[slot]
        cells.push({
          row: slot,
          col: column,
          color: waveColor(wave, waves),
          detail: (
            <div className='font-mono text-[10px]'>
              <div className='text-ghost'>
                index {cell.participantIndex} · slot {cell.slot}
              </div>
              <div className='text-ash'>{member ? shortAddress(member.operator) : ''}</div>
              <div className='text-ash'>
                {label(cell.aid, cell.ciphertextIndex)} · wave {wave + 1}
              </div>
              <div className='text-ash'>
                block {Number.isFinite(cell.block) && cell.block != null ? cell.block.toLocaleString() : '—'}
              </div>
              <div className='text-ash'>{cell.tx ? shortHash(cell.tx, 8, 6) : 'no transaction'}</div>
            </div>
          ),
        })
      })
    })

    const shareRow = matrix.rows.length
    const combinedRow = shareRow + 1
    matrix.columns.forEach((column) => {
      const name = label(column.aid, column.ciphertextIndex)
      if (column.share) {
        cells.push({
          row: shareRow,
          col: column.column,
          color: SHARE_COLOR,
          detail: (
            <div className='font-mono text-[10px]'>
              <div className='text-ghost'>{name}</div>
              <div className='text-ash'>organizer share published</div>
            </div>
          ),
        })
      }
      if (column.combined) {
        cells.push({
          row: combinedRow,
          col: column.column,
          color: COMBINED_COLOR,
          detail: (
            <div className='font-mono text-[10px]'>
              <div className='text-ghost'>{name}</div>
              <div className='text-ash'>
                combined · {column.partials} partials of t = {column.threshold}
              </div>
            </div>
          ),
        })
      }
    })

    const rows = [
      ...matrix.rows.map(
        (member) => `${String(member.participantIndex).padStart(2, '0')} ${shortAddress(member.operator)}`
      ),
      'organizer share',
      'combined',
    ]

    const legend: LegendItem[] = [
      ...Array.from({ length: waves }, (_, w) => ({ label: `wave ${w + 1}`, color: waveColor(w, waves) })),
      { label: 'no partial', color: CHART_COLORS.onyx },
      { label: 'organizer share', color: SHARE_COLOR },
      { label: 'combined', color: COMBINED_COLOR },
    ]

    return {
      rows,
      columns: matrix.columns.map((column) => label(column.aid, column.ciphertextIndex)),
      cells,
      legend,
      waves,
      partials: cells.length,
      combined: matrix.columns.filter((column) => column.combined).length,
      shares: matrix.columns.filter((column) => column.share).length,
    }
  }, [matrix, aidLetters])

  return (
    <Panel
      label='Decryption'
      title='Partial decryptions'
      description='Every δ_i published against every ciphertext of this epoch. t partials plus the organizer share make a ciphertext combinable.'
      actions={
        matrix ? (
          <span className='font-mono text-[11px] text-ash'>
            t = {matrix.threshold} · stagger {matrix.staggerBlocks} blocks
          </span>
        ) : null
      }
    >
      {!model ? (
        <EmptyState
          compact
          title='Nothing to plot'
          description='This epoch has no ciphertexts yet, so no partial decryption has been published.'
        />
      ) : (
        <>
          <p className='mb-4 text-[12px] text-ash'>
            {model.columns.length} ciphertexts · {model.partials} partials in {model.waves} wave
            {model.waves === 1 ? '' : 's'} · {model.shares} organizer shares · {model.combined} combined. Columns are
            labelled by application letter and ciphertext index.
          </p>
          <Matrix
            rows={model.rows}
            columns={model.columns}
            cells={model.cells}
            rowLabelWidth={132}
            legend={model.legend}
          />
        </>
      )}
    </Panel>
  )
}
