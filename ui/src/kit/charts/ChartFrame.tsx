import type { ReactNode } from 'react'
import { useMeasuredWidth } from '~hooks/use-measured-width'
import { cn } from '~lib/cn'
import { Skeleton } from '../Skeleton'
import { EmptyState } from '../EmptyState'
import type { ChartTooltipState } from './chart-tooltip'

export interface LegendItem {
  label: ReactNode
  color: string
  /** Optional value shown after the label — a count, a percentage. */
  value?: ReactNode
  /** Draw the swatch as an outline (for "expected"/reference series). */
  outline?: boolean
}

export function ChartLegend({ items, className }: { items: LegendItem[]; className?: string }) {
  if (items.length === 0) return null
  return (
    <ul className={cn('flex flex-wrap items-center gap-x-4 gap-y-1.5 text-[11px] text-pewter', className)}>
      {items.map((item, i) => (
        <li key={i} className='flex items-center gap-1.5'>
          <span
            aria-hidden='true'
            className='h-2 w-2 shrink-0 rounded-[2px]'
            style={
              item.outline
                ? { border: `1px solid ${item.color}`, background: 'transparent' }
                : { background: item.color }
            }
          />
          <span>{item.label}</span>
          {item.value != null ? <span className='font-mono tnum text-ash'>{item.value}</span> : null}
        </li>
      ))}
    </ul>
  )
}

export function ChartTooltipLayer({ tooltip, width }: { tooltip: ChartTooltipState | null; width: number }) {
  if (!tooltip) return null
  // Flip to the left of the cursor near the right edge so the box never leaves
  // the panel; the caller's container is `overflow-hidden`.
  const flip = tooltip.x > width - 160
  return (
    <div
      role='tooltip'
      style={{ left: tooltip.x, top: tooltip.y, transform: `translate(${flip ? '-100%' : '0'}, -100%)` }}
      className='pointer-events-none absolute z-20 -mt-2 max-w-[220px] rounded-sm border border-charcoal bg-onyx px-2.5 py-1.5 text-[11px] leading-snug text-silver shadow-[0_8px_24px_rgba(0,0,0,0.55)]'
    >
      {tooltip.content}
    </div>
  )
}

export interface ChartFrameProps {
  /** Drawing height in CSS pixels (excludes the legend). */
  height: number
  legend?: LegendItem[]
  loading?: boolean
  /** Show the empty state instead of the drawing. */
  empty?: boolean
  emptyLabel?: string
  emptyDescription?: string
  tooltip?: ChartTooltipState | null
  className?: string
  /** Receives the measured pixel width; render the SVG with it. */
  children: (width: number) => ReactNode
}

/**
 * Shell every chart shares: measures its own width (the SVG draws in CSS
 * pixels, so a viewBox scale would blow the type up with the geometry),
 * renders a skeleton until the first measurement, and owns the legend and the
 * tooltip layer so no chart re-implements them.
 */
export function ChartFrame({
  height,
  legend,
  loading = false,
  empty = false,
  emptyLabel = 'No data yet',
  emptyDescription,
  tooltip,
  className,
  children,
}: ChartFrameProps) {
  const [ref, width] = useMeasuredWidth<HTMLDivElement>()
  const ready = width != null && width > 0 && !loading

  return (
    <div className={cn('w-full', className)}>
      <div ref={ref} data-chart-root className='relative w-full overflow-hidden' style={{ minHeight: height }}>
        {empty ? (
          <EmptyState compact title={emptyLabel} description={emptyDescription} />
        ) : ready ? (
          children(width)
        ) : (
          <Skeleton className='w-full' height={height} rounded='md' />
        )}
        {ready && !empty ? <ChartTooltipLayer tooltip={tooltip ?? null} width={width} /> : null}
      </div>
      {legend && legend.length > 0 ? <ChartLegend items={legend} className='mt-3' /> : null}
    </div>
  )
}
