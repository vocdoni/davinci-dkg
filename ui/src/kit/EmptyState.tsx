import type { ReactNode } from 'react'
import { cn } from '~lib/cn'

export interface EmptyStateProps {
  title: ReactNode
  description?: ReactNode
  action?: ReactNode
  compact?: boolean
  className?: string
}

/** What a table or panel shows when the chain genuinely has nothing to show. */
export function EmptyState({ title, description, action, compact = false, className }: EmptyStateProps) {
  return (
    <div className={cn('flex flex-col items-center justify-center text-center', compact ? 'py-8' : 'py-16', className)}>
      <div className='mb-3 h-8 w-8 rounded-md border border-dashed border-charcoal' />
      <p className='text-[13px] font-medium text-silver'>{title}</p>
      {description ? <p className='mt-1 max-w-sm text-xs leading-relaxed text-ash'>{description}</p> : null}
      {action ? <div className='mt-4'>{action}</div> : null}
    </div>
  )
}
