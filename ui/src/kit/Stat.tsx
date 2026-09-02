import type { ReactNode } from 'react'
import { cn } from '~lib/cn'
import { Skeleton } from './Skeleton'

export interface StatProps {
  label: ReactNode
  value: ReactNode
  /** Small line under the value: a ratio, a delta, a block number. */
  hint?: ReactNode
  /** Badge or icon rendered on the right of the label row. */
  aside?: ReactNode
  tone?: 'default' | 'accent' | 'warn' | 'danger'
  loading?: boolean
  mono?: boolean
  className?: string
}

const TONES = {
  default: 'text-ghost',
  accent: 'text-emerald',
  warn: 'text-amber',
  danger: 'text-red',
} as const

/** Label / value / hint triple. The unit every header strip is built from. */
export function Stat({ label, value, hint, aside, tone = 'default', loading, mono, className }: StatProps) {
  return (
    <div className={cn('min-w-0', className)}>
      <div className='flex items-center justify-between gap-2'>
        <span className='label-caps truncate text-[11px] text-pewter'>{label}</span>
        {aside}
      </div>
      {loading ? (
        <Skeleton className='mt-2 h-7 w-24' />
      ) : (
        <div
          className={cn('mt-1.5 truncate text-2xl font-semibold tracking-tight', mono && 'font-mono tnum', TONES[tone])}
        >
          {value}
        </div>
      )}
      {hint ? <div className='mt-1 truncate text-xs text-ash'>{hint}</div> : null}
    </div>
  )
}

/** Row of stats separated by hairlines. Collapses to two columns under 1024 px. */
export function StatRow({ children, className }: { children: ReactNode; className?: string }) {
  return (
    <div
      className={cn(
        'grid grid-cols-2 gap-px overflow-hidden rounded-md border border-charcoal bg-charcoal lg:grid-cols-4',
        className
      )}
    >
      {children}
    </div>
  )
}

/** A `Stat` sized for `StatRow` (carbon cell, no border of its own). */
export function StatCell(props: StatProps) {
  return <Stat {...props} className={cn('bg-carbon p-4', props.className)} />
}
