import type { ReactNode } from 'react'
import { cn } from '~lib/cn'

export interface KeyValueItem {
  label: ReactNode
  value: ReactNode
  /** Render the value in JetBrains Mono with tabular numbers. */
  mono?: boolean
  /** Small explanatory line under the value. */
  hint?: ReactNode
}

export interface KeyValueProps {
  items: KeyValueItem[]
  /** 1 (stacked rows, default) or 2 columns. */
  columns?: 1 | 2
  /** Hairline between rows. */
  divided?: boolean
  className?: string
}

/** Label → value list. The raw-record view on every detail page. */
export function KeyValue({ items, columns = 1, divided = true, className }: KeyValueProps) {
  return (
    <dl className={cn('grid gap-x-8', columns === 2 ? 'sm:grid-cols-2' : 'grid-cols-1', className)}>
      {items.map((item, i) => (
        <div
          key={i}
          className={cn(
            'flex flex-wrap items-baseline justify-between gap-x-6 gap-y-1 py-2.5',
            divided && 'border-b border-charcoal last:border-b-0'
          )}
        >
          <dt className='label-caps text-[11px] text-pewter'>{item.label}</dt>
          <dd className='min-w-0 text-right'>
            <div className={cn('truncate text-[13px] text-silver', item.mono && 'font-mono tnum text-ghost')}>
              {item.value}
            </div>
            {item.hint ? <div className='text-[11px] text-ash'>{item.hint}</div> : null}
          </dd>
        </div>
      ))}
    </dl>
  )
}
