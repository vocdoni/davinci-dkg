import type { ReactNode } from 'react'
import { cn } from '~lib/cn'
import { DotIcon } from './icons'

export type BadgeTone = 'ok' | 'warn' | 'danger' | 'neutral' | 'accent'

const TONES: Record<BadgeTone, string> = {
  ok: 'text-emerald border-emerald/25 bg-emerald/5',
  accent: 'text-emerald border-emerald/40 bg-emerald/10',
  warn: 'text-amber border-amber/25 bg-amber/5',
  danger: 'text-red border-red/25 bg-red/5',
  neutral: 'text-pewter border-charcoal bg-transparent',
}

export interface BadgeProps {
  tone?: BadgeTone
  /** Prefix the label with a filled dot — for live/phase states. */
  dot?: boolean
  size?: 'sm' | 'md'
  title?: string
  className?: string
  children: ReactNode
}

/** Pill badge. Emerald for live/ok, amber for warnings, red for failures. */
export function Badge({ tone = 'neutral', dot = false, size = 'md', title, className, children }: BadgeProps) {
  return (
    <span
      title={title}
      className={cn(
        'inline-flex items-center gap-1.5 whitespace-nowrap rounded-pill border font-medium tracking-[0.02em]',
        size === 'sm' ? 'px-2 py-[1px] text-[10px]' : 'px-2.5 py-[3px] text-[11px]',
        TONES[tone],
        className
      )}
    >
      {dot ? <DotIcon size={6} className='shrink-0' /> : null}
      {children}
    </span>
  )
}
