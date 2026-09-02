import type { ReactNode } from 'react'
import { cn } from '~lib/cn'

export interface ProgressBarProps {
  /** How many of `total` are in. */
  value: number
  total: number
  /** Draws a tick at this count — the threshold t the epoch needs. */
  threshold?: number
  /** `t of n` caption above the bar; pass `false` to hide it. */
  label?: ReactNode | false
  tone?: 'accent' | 'warn' | 'danger' | 'neutral'
  size?: 'sm' | 'md'
  className?: string
}

const FILLS = {
  accent: 'bg-emerald',
  warn: 'bg-amber',
  danger: 'bg-red',
  neutral: 'bg-warm-gray',
} as const

const pct = (n: number, d: number) => (d <= 0 ? 0 : Math.max(0, Math.min(100, (n / d) * 100)))

/**
 * `t of n` bar with an optional threshold marker. The marker is what makes it
 * readable at a glance: everything left of the tick is "not yet decryptable".
 */
export function ProgressBar({
  value,
  total,
  threshold,
  label,
  tone = 'accent',
  size = 'md',
  className,
}: ProgressBarProps) {
  const filled = pct(value, total)
  const mark = threshold != null ? pct(threshold, total) : null
  const reached = threshold == null || value >= threshold
  return (
    <div className={cn('min-w-0', className)}>
      {label !== false ? (
        <div className='mb-1.5 flex items-baseline justify-between gap-2 text-[11px]'>
          <span className='text-pewter'>{label}</span>
          <span className={cn('font-mono tnum', reached ? 'text-emerald' : 'text-silver')}>
            {value}
            <span className='text-ash'> / {total}</span>
            {threshold != null ? <span className='text-ash'> · t={threshold}</span> : null}
          </span>
        </div>
      ) : null}
      <div
        role='progressbar'
        aria-valuenow={value}
        aria-valuemin={0}
        aria-valuemax={total}
        className={cn('relative w-full overflow-hidden rounded-pill bg-onyx', size === 'sm' ? 'h-1' : 'h-1.5')}
      >
        <div
          className={cn('h-full rounded-pill transition-[width] duration-500', FILLS[reached ? tone : 'neutral'])}
          style={{ width: `${filled}%` }}
        />
        {mark != null ? (
          <div
            aria-hidden='true'
            title={`threshold ${threshold}`}
            className='absolute inset-y-[-2px] w-px bg-pewter'
            style={{ left: `${mark}%` }}
          />
        ) : null}
      </div>
    </div>
  )
}
