import type { ReactNode } from 'react'
import { cn } from '~lib/cn'

export type TimelineTone = 'ok' | 'warn' | 'danger' | 'neutral' | 'muted'

const DOTS: Record<TimelineTone, string> = {
  ok: 'bg-emerald border-emerald',
  warn: 'bg-amber border-amber',
  danger: 'bg-red border-red',
  neutral: 'bg-silver border-silver',
  muted: 'bg-transparent border-charcoal',
}

export interface TimelineRowProps {
  title: ReactNode
  /** Left gutter: block number, time, index. Rendered mono. */
  meta?: ReactNode
  description?: ReactNode
  /** Right-aligned slot: tx link, gas, badge. */
  right?: ReactNode
  tone?: TimelineTone
  /** Hides the connector below the dot (use on the final row). */
  last?: boolean
  className?: string
}

/** One event in a vertical rail. Composed by the epoch and operator pages. */
export function TimelineRow({ title, meta, description, right, tone = 'neutral', last, className }: TimelineRowProps) {
  return (
    <li className={cn('relative flex gap-3 pb-4 last:pb-0', className)}>
      <div className='relative flex w-3 shrink-0 justify-center pt-1.5'>
        <span className={cn('z-10 h-2 w-2 rounded-full border', DOTS[tone])} />
        {!last ? <span aria-hidden='true' className='absolute top-4 bottom-[-16px] w-px bg-charcoal' /> : null}
      </div>
      <div className='flex min-w-0 flex-1 flex-wrap items-baseline justify-between gap-x-4 gap-y-1'>
        <div className='min-w-0'>
          <div className='truncate text-[13px] text-silver'>{title}</div>
          {description ? <div className='mt-0.5 text-xs text-ash'>{description}</div> : null}
        </div>
        <div className='flex shrink-0 items-center gap-3 font-mono text-[11px] tnum text-ash'>
          {meta}
          {right}
        </div>
      </div>
    </li>
  )
}

export function Timeline({ children, className }: { children: ReactNode; className?: string }) {
  return <ol className={cn('m-0 list-none p-0', className)}>{children}</ol>
}
