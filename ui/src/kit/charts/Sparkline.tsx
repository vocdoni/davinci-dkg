import { areaPath, extent, linePath, linearScale } from './scale'
import { CHART_COLORS } from './colors'
import { cn } from '~lib/cn'

export interface SparklineProps {
  values: number[]
  width?: number
  height?: number
  color?: string
  /** Fill under the line. */
  area?: boolean
  /** Dot on the final point. */
  showLast?: boolean
  /** Baseline at zero rather than at the series minimum. */
  zeroBaseline?: boolean
  ariaLabel?: string
  className?: string
}

/**
 * Inline trend line. Fixed-size and unmeasured on purpose: sparklines live in
 * table cells, where a ResizeObserver per row would be absurd.
 */
export function Sparkline({
  values,
  width = 96,
  height = 24,
  color = CHART_COLORS.emerald,
  area = true,
  showLast = true,
  zeroBaseline = true,
  ariaLabel = 'trend',
  className,
}: SparklineProps) {
  if (values.length === 0) return <span className={cn('text-[11px] text-ash', className)}>—</span>

  const pad = 2
  const { min, max } = extent(values)
  const lo = zeroBaseline ? Math.min(min, 0) : min
  const hi = max === lo ? lo + 1 : max
  const x = linearScale([0, Math.max(values.length - 1, 1)], [pad, width - pad])
  const y = linearScale([lo, hi], [height - pad, pad])
  const points = values.map((value, i) => ({ x: x(i), y: y(value) }))
  const last = points[points.length - 1] as { x: number; y: number }

  return (
    <svg width={width} height={height} role='img' aria-label={ariaLabel} className={cn('overflow-visible', className)}>
      {area ? <path d={areaPath(points, height - pad)} fill={color} opacity={0.12} /> : null}
      <path
        d={linePath(points)}
        fill='none'
        stroke={color}
        strokeWidth={1.5}
        strokeLinejoin='round'
        strokeLinecap='round'
      />
      {showLast ? <circle cx={last.x} cy={last.y} r={2} fill={color} /> : null}
    </svg>
  )
}
