import type { ReactNode } from 'react'
import { arcPath, clamp, formatPercent } from './scale'
import { CHART_COLORS } from './colors'
import { ChartFrame } from './ChartFrame'

export interface GaugeProps {
  /** Fraction in [0, 1]. */
  value: number
  label?: ReactNode
  /** Line under the percentage — e.g. the raw τ value. */
  caption?: ReactNode
  size?: number
  thickness?: number
  color?: string
  /** Second, fainter arc: the reference fraction (e.g. α·N/2^256). */
  reference?: number
  loading?: boolean
  className?: string
}

/**
 * Half-circle gauge. Built for "τ as a fraction of the hash space": the number
 * is tiny and the shape carries the meaning, so the arc is the headline and
 * the percentage sits inside it.
 */
export function Gauge({
  value,
  label,
  caption,
  size = 180,
  thickness = 14,
  color = CHART_COLORS.emerald,
  reference,
  loading = false,
  className,
}: GaugeProps) {
  const fraction = clamp(Number.isFinite(value) ? value : 0, 0, 1)
  const height = Math.round(size / 2) + 34

  return (
    <ChartFrame height={height} loading={loading} className={className}>
      {(width) => {
        const cx = width / 2
        const outer = Math.min(size, width) / 2 - 2
        const inner = Math.max(outer - thickness, 4)
        const cy = outer + 4
        // -90° … +90° drawn as a 180° sweep starting at 9 o'clock.
        const track = arcPath(cx, cy, outer, inner, 270, 450)
        const filled = arcPath(cx, cy, outer, inner, 270, 270 + 180 * fraction)
        const ref = reference == null ? null : 270 + 180 * clamp(reference, 0, 1)
        return (
          <svg width={width} height={height} role='img' aria-label={typeof label === 'string' ? label : 'gauge'}>
            <path d={track} fill={CHART_COLORS.onyx} />
            {filled ? <path d={filled} fill={color} /> : null}
            {ref != null ? (
              <path d={arcPath(cx, cy, outer, inner, ref - 0.6, ref + 0.6)} fill={CHART_COLORS.pewter} />
            ) : null}
            <text
              x={cx}
              y={cy - 6}
              textAnchor='middle'
              fontSize={20}
              fontWeight={600}
              fill='#ffffff'
              fontFamily='var(--font-mono)'
            >
              {formatPercent(fraction, fraction < 0.01 ? 3 : 1)}
            </text>
            {label ? (
              <text x={cx} y={cy + 14} textAnchor='middle' fontSize={10} fill='#b8b3b0'>
                {label}
              </text>
            ) : null}
            {caption ? (
              <text
                x={cx}
                y={height - 8}
                textAnchor='middle'
                fontSize={10}
                fill='#8a8380'
                fontFamily='var(--font-mono)'
              >
                {caption}
              </text>
            ) : null}
          </svg>
        )
      }}
    </ChartFrame>
  )
}
