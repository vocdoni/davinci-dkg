import type { ReactNode } from 'react'
import { linearScale, num } from './scale'
import { CHART_COLORS, PHASE_COLORS, type PhaseColorKey } from './colors'
import { ChartFrame, type LegendItem } from './ChartFrame'
import { useChartTooltip } from './chart-tooltip'

export interface BlockWindow {
  label: string
  from: number | bigint
  to: number | bigint
  /** A phase key from the palette, or any explicit colour. */
  tone?: PhaseColorKey | string
  /** Extra tooltip lines. */
  detail?: ReactNode
}

export interface BlockTimelineProps {
  windows: BlockWindow[]
  /** Chain head. Drawn as a labelled emerald marker. */
  current?: number | bigint | null
  /** Axis bounds; default to the span of the windows (plus the marker). */
  start?: number | bigint
  end?: number | bigint
  height?: number
  loading?: boolean
  className?: string
}

const PAD = { top: 20, right: 12, bottom: 26, left: 12 }

function toneColor(tone?: PhaseColorKey | string): string {
  if (!tone) return CHART_COLORS.teal
  return (PHASE_COLORS as Record<string, string>)[tone] ?? tone
}

/**
 * The epoch lifecycle drawn on a block axis: committee selection, key
 * assembly, live, finalize — with the current block marked. This is the one
 * picture that answers "what is this epoch waiting for?".
 */
export function BlockTimeline({
  windows,
  current,
  start,
  end,
  height = 96,
  loading = false,
  className,
}: BlockTimelineProps) {
  const { tooltip, show, hide } = useChartTooltip()
  const froms = windows.map((w) => num(w.from))
  const tos = windows.map((w) => num(w.to))
  const head = current == null ? null : num(current)
  const lo = start != null ? num(start) : Math.min(...froms, head ?? Number.POSITIVE_INFINITY)
  const hi = end != null ? num(end) : Math.max(...tos, head ?? Number.NEGATIVE_INFINITY)
  const legend: LegendItem[] = windows.map((w) => ({ label: w.label, color: toneColor(w.tone) }))

  return (
    <ChartFrame
      height={height}
      legend={legend}
      loading={loading}
      empty={windows.length === 0}
      emptyLabel='No windows to draw'
      tooltip={tooltip}
      className={className}
    >
      {(width) => {
        const plotW = Math.max(width - PAD.left - PAD.right, 10)
        const x = linearScale([lo, hi === lo ? lo + 1 : hi], [0, plotW])
        const laneY = PAD.top + 6
        const laneH = Math.max(height - PAD.top - PAD.bottom - 12, 14)
        return (
          <svg width={width} height={height} role='img' aria-label='Epoch lifecycle on the block axis'>
            <g transform={`translate(${PAD.left},0)`}>
              <rect x={0} y={laneY} width={plotW} height={laneH} rx={4} fill={CHART_COLORS.onyx} />
              {windows.map((w, i) => {
                const x0 = x(num(w.from))
                const x1 = x(num(w.to))
                const color = toneColor(w.tone)
                const w0 = Math.max(x1 - x0, 2)
                return (
                  <g
                    key={i}
                    onMouseMove={(e) =>
                      show(
                        e,
                        <div>
                          <div className='text-ghost'>{w.label}</div>
                          <div className='font-mono tnum text-[10px] text-ash'>
                            {num(w.from).toLocaleString()} → {num(w.to).toLocaleString()}
                          </div>
                          {w.detail ? <div className='mt-0.5 text-[10px] text-ash'>{w.detail}</div> : null}
                        </div>
                      )
                    }
                    onMouseLeave={hide}
                  >
                    <rect x={x0} y={laneY} width={w0} height={laneH} rx={3} fill={color} opacity={0.28} />
                    <rect x={x0} y={laneY} width={2} height={laneH} fill={color} />
                    {w0 > 54 ? (
                      <text x={x0 + 6} y={laneY + laneH / 2} dy='0.32em' fontSize={10} fill={color}>
                        {w.label}
                      </text>
                    ) : null}
                  </g>
                )
              })}

              {head != null ? (
                <g>
                  <line
                    x1={x(head)}
                    x2={x(head)}
                    y1={laneY - 8}
                    y2={laneY + laneH + 6}
                    stroke={CHART_COLORS.emerald}
                    strokeWidth={1.5}
                  />
                  <circle cx={x(head)} cy={laneY - 9} r={3} fill={CHART_COLORS.emerald} />
                  <text
                    x={Math.min(Math.max(x(head), 24), plotW - 24)}
                    y={laneY - 14}
                    textAnchor='middle'
                    fontSize={10}
                    fill={CHART_COLORS.emerald}
                    fontFamily='var(--font-mono)'
                  >
                    {head.toLocaleString()}
                  </text>
                </g>
              ) : null}

              <text x={0} y={height - 8} fontSize={10} fill='#8a8380' fontFamily='var(--font-mono)'>
                {lo.toLocaleString()}
              </text>
              <text
                x={plotW}
                y={height - 8}
                textAnchor='end'
                fontSize={10}
                fill='#8a8380'
                fontFamily='var(--font-mono)'
              >
                {hi.toLocaleString()}
              </text>
            </g>
          </svg>
        )
      }}
    </ChartFrame>
  )
}
