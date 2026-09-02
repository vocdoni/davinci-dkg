import { useState, type ReactNode } from 'react'
import { bandScale, formatCompact, linearScale, niceTicks, stackMax, stackSeries } from './scale'
import { seriesColor } from './colors'
import { ChartFrame, type LegendItem } from './ChartFrame'
import { useChartTooltip } from './chart-tooltip'

export interface BarSeries {
  key: string
  label: string
  /** Defaults to the palette order (emerald, teal, slate, …). */
  color?: string
}

export interface BarDatum {
  /** X-axis label. Every nth is drawn, so it can be long-ish. */
  label: string
  values: Record<string, number>
  /** Extra lines in the tooltip — a block range, an epoch id. */
  note?: ReactNode
}

export interface StackedBarsProps {
  data: BarDatum[]
  series: BarSeries[]
  height?: number
  loading?: boolean
  /** Approximate y-axis tick count. */
  yTicks?: number
  onBarClick?: (datum: BarDatum, index: number) => void
  className?: string
}

const PAD = { top: 8, right: 8, bottom: 22, left: 40 }

/**
 * Categorical stacked bars — the "activity per epoch" chart. Hovering a column
 * highlights it and breaks the stack down by series in the tooltip.
 */
export function StackedBars({
  data,
  series,
  height = 220,
  loading = false,
  yTicks = 4,
  onBarClick,
  className,
}: StackedBarsProps) {
  const { tooltip, show, hide } = useChartTooltip()
  const [hovered, setHovered] = useState<number | null>(null)

  const keys = series.map((s) => s.key)
  const colors = series.map((s, i) => s.color ?? seriesColor(i))
  const legend: LegendItem[] = series.map((s, i) => ({ label: s.label, color: colors[i] as string }))
  const values = data.map((d) => d.values)
  const max = stackMax(values, keys)
  const ticks = niceTicks(0, max === 0 ? 1 : max, yTicks)

  return (
    <ChartFrame
      height={height}
      legend={legend}
      loading={loading}
      empty={data.length === 0}
      emptyLabel='No activity in this range'
      tooltip={tooltip}
      className={className}
    >
      {(width) => {
        const plotW = Math.max(width - PAD.left - PAD.right, 10)
        const plotH = Math.max(height - PAD.top - PAD.bottom, 10)
        const band = bandScale(data.length, plotW, 0.28)
        const y = linearScale([0, max === 0 ? 1 : max], [plotH, 0])
        const stacks = stackSeries(values, keys)
        // Only every nth label fits; pick the stride from the band width.
        const labelPx = Math.max(...data.map((d) => d.label.length), 4) * 6.2 + 10
        const stride = Math.max(1, Math.ceil((data.length * labelPx) / plotW))

        return (
          <svg width={width} height={height} role='img' aria-label='Stacked activity chart'>
            <g transform={`translate(${PAD.left},${PAD.top})`}>
              {ticks.map((tick) => (
                <g key={tick}>
                  <line x1={0} x2={plotW} y1={y(tick)} y2={y(tick)} stroke='#3d3a39' strokeWidth={1} opacity={0.5} />
                  <text
                    x={-8}
                    y={y(tick)}
                    dy='0.32em'
                    textAnchor='end'
                    fontSize={10}
                    fill='#8a8380'
                    fontFamily='var(--font-mono)'
                  >
                    {formatCompact(tick)}
                  </text>
                </g>
              ))}

              {data.map((datum, i) => {
                const segments = stacks[i] ?? []
                const total = segments.reduce((sum, s) => sum + s.value, 0)
                return (
                  <g
                    key={i}
                    onMouseMove={(e) => {
                      setHovered(i)
                      show(e, <BarTooltip datum={datum} series={series} colors={colors} total={total} />)
                    }}
                    onMouseLeave={() => {
                      setHovered(null)
                      hide()
                    }}
                    onClick={onBarClick ? () => onBarClick(datum, i) : undefined}
                    className={onBarClick ? 'cursor-pointer' : undefined}
                  >
                    <rect
                      x={band.at(i) - (band.step - band.bandWidth) / 2}
                      y={0}
                      width={band.step}
                      height={plotH}
                      fill={hovered === i ? '#1a1a1a' : 'transparent'}
                    />
                    {segments.map((segment, s) =>
                      segment.value > 0 ? (
                        <rect
                          key={segment.key}
                          x={band.at(i)}
                          y={y(segment.y1)}
                          width={band.bandWidth}
                          height={Math.max(y(segment.y0) - y(segment.y1), 1)}
                          fill={colors[s]}
                          opacity={hovered == null || hovered === i ? 1 : 0.45}
                        />
                      ) : null
                    )}
                  </g>
                )
              })}

              <line x1={0} x2={plotW} y1={plotH} y2={plotH} stroke='#3d3a39' />
              {data.map((datum, i) =>
                i % stride === 0 ? (
                  <text
                    key={i}
                    x={band.center(i)}
                    y={plotH + 14}
                    textAnchor='middle'
                    fontSize={10}
                    fill='#8a8380'
                    fontFamily='var(--font-mono)'
                  >
                    {datum.label}
                  </text>
                ) : null
              )}
            </g>
          </svg>
        )
      }}
    </ChartFrame>
  )
}

function BarTooltip({
  datum,
  series,
  colors,
  total,
}: {
  datum: BarDatum
  series: BarSeries[]
  colors: string[]
  total: number
}) {
  return (
    <div>
      <div className='font-mono text-[11px] text-ghost'>{datum.label}</div>
      {datum.note ? <div className='mt-0.5 text-[10px] text-ash'>{datum.note}</div> : null}
      <div className='mt-1.5 space-y-0.5'>
        {series.map((s, i) => (
          <div key={s.key} className='flex items-center justify-between gap-3'>
            <span className='flex items-center gap-1.5'>
              <span className='h-2 w-2 rounded-[2px]' style={{ background: colors[i] }} />
              {s.label}
            </span>
            <span className='font-mono tnum text-ghost'>{datum.values[s.key] ?? 0}</span>
          </div>
        ))}
        <div className='flex items-center justify-between gap-3 border-t border-charcoal pt-0.5 text-ash'>
          <span>total</span>
          <span className='font-mono tnum'>{total}</span>
        </div>
      </div>
    </div>
  )
}
