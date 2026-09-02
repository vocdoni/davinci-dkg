import type { ReactNode } from 'react'
import { linearScale, num } from './scale'
import { CHART_COLORS, PHASE_COLORS, type PhaseColorKey } from './colors'
import { ChartFrame, type LegendItem } from './ChartFrame'
import { useChartTooltip } from './chart-tooltip'

export interface CadenceEpoch {
  id: string
  label?: string
  start: number | bigint
  end: number | bigint
  phase: PhaseColorKey
  /** Extra tooltip lines: committee size, contributions, ciphertexts. */
  detail?: ReactNode
}

export interface CadenceStripProps {
  epochs: CadenceEpoch[]
  current?: number | bigint | null
  height?: number
  loading?: boolean
  onEpochClick?: (epoch: CadenceEpoch) => void
  className?: string
}

const PAD = { top: 10, right: 10, bottom: 20, left: 10 }

/**
 * Epochs laid on a shared block axis, coloured by phase. Overlap is the point:
 * the strip shows the cadence — how epochs stagger, where a gap opened, which
 * one aborted.
 */
export function CadenceStrip({
  epochs,
  current,
  height = 84,
  loading = false,
  onEpochClick,
  className,
}: CadenceStripProps) {
  const { tooltip, show, hide } = useChartTooltip()
  const head = current == null ? null : num(current)
  const lo = Math.min(...epochs.map((e) => num(e.start)), head ?? Number.POSITIVE_INFINITY)
  const hi = Math.max(...epochs.map((e) => num(e.end)), head ?? Number.NEGATIVE_INFINITY)
  const phases = Array.from(new Set(epochs.map((e) => e.phase)))
  const legend: LegendItem[] = phases.map((phase) => ({ label: phase, color: PHASE_COLORS[phase] }))

  return (
    <ChartFrame
      height={height}
      legend={legend}
      loading={loading}
      empty={epochs.length === 0}
      emptyLabel='No epochs yet'
      tooltip={tooltip}
      className={className}
    >
      {(width) => {
        const plotW = Math.max(width - PAD.left - PAD.right, 10)
        const x = linearScale([lo, hi === lo ? lo + 1 : hi], [0, plotW])
        // Stagger overlapping epochs into lanes so none is hidden behind another.
        const lanes: number[] = []
        const laneOf = epochs.map((epoch) => {
          const from = num(epoch.start)
          const index = lanes.findIndex((end) => end <= from)
          const lane = index === -1 ? lanes.length : index
          lanes[lane] = num(epoch.end)
          return lane
        })
        const laneCount = Math.max(lanes.length, 1)
        const laneH = Math.max((height - PAD.top - PAD.bottom) / laneCount - 2, 6)

        return (
          <svg width={width} height={height} role='img' aria-label='Epoch cadence'>
            <g transform={`translate(${PAD.left},0)`}>
              {epochs.map((epoch, i) => {
                const x0 = x(num(epoch.start))
                const x1 = x(num(epoch.end))
                const y0 = PAD.top + (laneOf[i] as number) * (laneH + 2)
                return (
                  <rect
                    key={epoch.id}
                    x={x0}
                    y={y0}
                    width={Math.max(x1 - x0, 2)}
                    height={laneH}
                    rx={2}
                    fill={PHASE_COLORS[epoch.phase]}
                    opacity={0.85}
                    className={onEpochClick ? 'cursor-pointer' : undefined}
                    onClick={onEpochClick ? () => onEpochClick(epoch) : undefined}
                    onMouseMove={(e) =>
                      show(
                        e,
                        <div>
                          <div className='font-mono text-[11px] text-ghost'>{epoch.label ?? epoch.id}</div>
                          <div className='font-mono tnum text-[10px] text-ash'>
                            {num(epoch.start).toLocaleString()} → {num(epoch.end).toLocaleString()} · {epoch.phase}
                          </div>
                          {epoch.detail ? <div className='mt-0.5 text-[10px] text-ash'>{epoch.detail}</div> : null}
                        </div>
                      )
                    }
                    onMouseLeave={hide}
                  />
                )
              })}
              {head != null ? (
                <line
                  x1={x(head)}
                  x2={x(head)}
                  y1={PAD.top - 6}
                  y2={height - PAD.bottom + 2}
                  stroke={CHART_COLORS.emerald}
                  strokeWidth={1.5}
                  strokeDasharray='3 2'
                />
              ) : null}
              <text x={0} y={height - 6} fontSize={10} fill='#8a8380' fontFamily='var(--font-mono)'>
                {lo.toLocaleString()}
              </text>
              <text
                x={plotW}
                y={height - 6}
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
