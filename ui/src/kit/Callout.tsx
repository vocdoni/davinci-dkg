import type { ReactNode } from 'react'
import { cn } from '~lib/cn'
import { InfoIcon, WarningIcon } from './icons'

export type CalloutTone = 'info' | 'ok' | 'warn' | 'danger'

const TONES: Record<CalloutTone, { box: string; icon: string }> = {
  info: { box: 'border-charcoal bg-onyx/40', icon: 'text-pewter' },
  ok: { box: 'border-emerald/20 bg-emerald/[0.04]', icon: 'text-emerald' },
  warn: { box: 'border-amber/20 bg-amber/[0.04]', icon: 'text-amber' },
  danger: { box: 'border-red/20 bg-red/[0.04]', icon: 'text-red' },
}

export interface CalloutProps {
  tone?: CalloutTone
  title?: ReactNode
  actions?: ReactNode
  className?: string
  children?: ReactNode
}

/** Inline notice. Borders and a tinted wash — never a shadowed toast. */
export function Callout({ tone = 'info', title, actions, className, children }: CalloutProps) {
  const style = TONES[tone]
  const Glyph = tone === 'warn' || tone === 'danger' ? WarningIcon : InfoIcon
  return (
    <div className={cn('flex gap-3 rounded-md border p-4', style.box, className)}>
      <Glyph className={cn('mt-0.5 shrink-0', style.icon)} />
      <div className='min-w-0 flex-1'>
        {title ? <div className='text-[13px] font-semibold text-ghost'>{title}</div> : null}
        {children ? (
          <div className={cn('text-[13px] leading-relaxed text-ash', title && 'mt-1')}>{children}</div>
        ) : null}
      </div>
      {actions ? <div className='flex shrink-0 items-start gap-2'>{actions}</div> : null}
    </div>
  )
}
