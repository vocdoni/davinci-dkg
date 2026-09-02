import type { ReactNode } from 'react'
import { arcPath, formatPercent } from './scale'
import { seriesColor } from './colors'
import { ChartFrame, type LegendItem } from './ChartFrame'
import { useChartTooltip } from './chart-tooltip'

export interface DonutSlice {
  label: string
  value: number
  color?: string
}

export interface DonutProps {
  slices: DonutSlice[]
  size?: number
  thickness?: number
  /** Big number in the hole; defaults to the total. */
  centerValue?: ReactNode
  centerLabel?: ReactNode
  loading?: boolean
  className?: string
}

/** Composition of a small set — operator status, phase mix. */
export function Donut({
  slices,
  size = 168,
  thickness = 18,
  centerValue,
  centerLabel,
  loading = false,
  className,
}: DonutProps) {
  const { tooltip, show, hide } = useChartTooltip()
  const total = slices.reduce((sum, s) => sum + Math.max(s.value, 0), 0)
  const colors = slices.map((s, i) => s.color ?? seriesColor(i))
  const legend: LegendItem[] = slices.map((s, i) => ({
    label: s.label,
    color: colors[i] as string,
    value: total > 0 ? formatPercent(s.value / total, 0) : '0%',
  }))

  return (
    <ChartFrame
      height={size}
      legend={legend}
      loading={loading}
      empty={total === 0}
      emptyLabel='Nothing to break down'
      tooltip={tooltip}
      className={className}
    >
      {(width) => {
        const cx = Math.min(width, size) / 2
        const cy = size / 2
        const outer = Math.min(width, size) / 2 - 2
        const inner = Math.max(outer - thickness, 4)
        let angle = 0
        return (
          <svg width={width} height={size} role='img' aria-label='Distribution'>
            {slices.map((slice, i) => {
              const sweep = (Math.max(slice.value, 0) / total) * 360
              const d = arcPath(cx, cy, outer, inner, angle, angle + sweep)
              angle += sweep
              if (!d) return null
              return (
                <path
                  key={i}
                  d={d}
                  fill={colors[i]}
                  onMouseMove={(e) =>
                    show(
                      e,
                      <div>
                        <div className='text-ghost'>{slice.label}</div>
                        <div className='font-mono tnum text-[10px] text-ash'>
                          {slice.value} · {formatPercent(slice.value / total, 1)}
                        </div>
                      </div>
                    )
                  }
                  onMouseLeave={hide}
                />
              )
            })}
            <text
              x={cx}
              y={cy - 2}
              textAnchor='middle'
              fontSize={22}
              fontWeight={600}
              fill='#ffffff'
              fontFamily='var(--font-mono)'
            >
              {centerValue ?? total}
            </text>
            {centerLabel ? (
              <text x={cx} y={cy + 16} textAnchor='middle' fontSize={10} fill='#8a8380'>
                {centerLabel}
              </text>
            ) : null}
          </svg>
        )
      }}
    </ChartFrame>
  )
}
