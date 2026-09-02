import type { ReactNode } from 'react'
import { cn } from '~lib/cn'
import { useMeasuredWidth } from '~hooks/use-measured-width'
import { CHART_COLORS, heatColor } from './colors'
import { ChartLegend, ChartTooltipLayer, type LegendItem } from './ChartFrame'
import { useChartTooltip } from './chart-tooltip'
import { Skeleton } from '../Skeleton'
import { EmptyState } from '../EmptyState'

export interface MatrixCell {
  row: number
  col: number
  /** Drives the default heat ramp when no explicit colour is given. */
  value?: number
  /** Explicit fill — used for wave colouring in the decryption matrix. */
  color?: string
  /** Tooltip body for this cell. */
  detail?: ReactNode
}

export interface MatrixProps {
  /** Row labels, top to bottom (committee members). */
  rows: string[]
  /** Column labels, left to right (ciphertexts). */
  columns: string[]
  cells: MatrixCell[]
  /** Square cell edge in px. Defaults to whatever fits the panel width, clamped to 12–24. */
  cellSize?: number
  gap?: number
  /** Value that maps to full emerald; defaults to the largest cell value. */
  max?: number
  rowLabelWidth?: number
  legend?: LegendItem[]
  loading?: boolean
  onCellClick?: (cell: MatrixCell) => void
  className?: string
}

/**
 * Heatmap grid with sticky row labels. The label column stays put while the
 * cells scroll horizontally *inside the panel* — the page never scrolls
 * sideways, which is the rule for the 64 × 32 decryption matrix.
 */
export function Matrix({
  rows,
  columns,
  cells,
  cellSize: cellSizeProp,
  gap = 2,
  max,
  rowLabelWidth = 116,
  legend,
  loading = false,
  onCellClick,
  className,
}: MatrixProps) {
  const { tooltip, show, hide } = useChartTooltip()
  const [gridRef, gridWidth] = useMeasuredWidth<HTMLDivElement>()
  const fitted = gridWidth ? Math.floor(gridWidth / Math.max(columns.length, 1)) - gap : 12
  const cellSize = cellSizeProp ?? Math.min(24, Math.max(12, fitted))
  const step = cellSize + gap
  const gridW = Math.max(columns.length * step - gap, 0)
  const gridH = Math.max(rows.length * step - gap, 0)
  const headerH = 18
  const ceiling = max ?? cells.reduce((m, c) => Math.max(m, c.value ?? 0), 0)
  // Every nth column label, so headers never collide at cellSize 10-12.
  const labelStride = Math.max(1, Math.ceil(26 / step))

  if (loading)
    return <Skeleton className={cn('w-full', className)} height={Math.min(gridH + headerH, 420)} rounded='md' />
  if (rows.length === 0 || columns.length === 0) {
    return <EmptyState compact title='Nothing to plot' description='This epoch has no partial decryptions yet.' />
  }

  return (
    <div className={cn('w-full', className)}>
      <div data-chart-root className='relative flex w-full'>
        <div className='shrink-0 pr-2' style={{ width: rowLabelWidth, paddingTop: headerH }}>
          {rows.map((label, r) => (
            <div
              key={r}
              title={label}
              className='truncate text-right font-mono text-[10px] leading-none text-ash'
              style={{ height: cellSize, marginBottom: gap, lineHeight: `${cellSize}px` }}
            >
              {label}
            </div>
          ))}
        </div>

        <div ref={gridRef} className='min-w-0 flex-1 overflow-x-auto scroll-slim'>
          <svg width={Math.max(gridW, 1)} height={gridH + headerH} role='img' aria-label='Partial decryption matrix'>
            {columns.map((label, c) =>
              c % labelStride === 0 ? (
                <text
                  key={c}
                  x={c * step + cellSize / 2}
                  y={headerH - 6}
                  textAnchor='middle'
                  fontSize={9}
                  fill='#8a8380'
                  fontFamily='var(--font-mono)'
                >
                  {label}
                </text>
              ) : null
            )}
            <g transform={`translate(0,${headerH})`}>
              {rows.map((_, r) =>
                columns.map((__, c) => (
                  <rect
                    key={`bg-${r}-${c}`}
                    x={c * step}
                    y={r * step}
                    width={cellSize}
                    height={cellSize}
                    rx={2}
                    fill={CHART_COLORS.onyx}
                  />
                ))
              )}
              {cells.map((cell, i) => (
                <rect
                  key={i}
                  x={cell.col * step}
                  y={cell.row * step}
                  width={cellSize}
                  height={cellSize}
                  rx={2}
                  fill={cell.color ?? heatColor(cell.value ?? 0, ceiling)}
                  className={onCellClick ? 'cursor-pointer' : undefined}
                  onMouseMove={(e) =>
                    show(
                      e,
                      cell.detail ?? (
                        <div className='font-mono text-[10px]'>
                          <div className='text-ghost'>{rows[cell.row]}</div>
                          <div className='text-ash'>
                            {columns[cell.col]} · {cell.value ?? 1}
                          </div>
                        </div>
                      )
                    )
                  }
                  onMouseLeave={hide}
                  onClick={onCellClick ? () => onCellClick(cell) : undefined}
                />
              ))}
            </g>
          </svg>
        </div>
        <ChartTooltipLayer tooltip={tooltip} width={rowLabelWidth + gridW} />
      </div>
      {legend && legend.length > 0 ? <ChartLegend items={legend} className='mt-3' /> : null}
    </div>
  )
}
