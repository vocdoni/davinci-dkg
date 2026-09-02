import type { ReactNode } from 'react'
import { cn } from '~lib/cn'

/** Page gutter. 1400 px max, 24 px side padding, 16 px under 1024 px. */
export function PageContainer({
  children,
  className,
  width = 'default',
}: {
  children: ReactNode
  className?: string
  /** `wide` removes the max width for full-bleed matrices. */
  width?: 'default' | 'wide' | 'prose'
}) {
  return (
    <div
      className={cn(
        'mx-auto w-full px-4 lg:px-6',
        width === 'default' && 'max-w-[1400px]',
        width === 'prose' && 'max-w-[820px]',
        className
      )}
    >
      {children}
    </div>
  )
}

export interface SectionHeaderProps {
  /** 13 px uppercase emerald eyebrow. */
  label?: ReactNode
  title: ReactNode
  description?: ReactNode
  actions?: ReactNode
  className?: string
  size?: 'page' | 'section'
}

/** Uppercase label over a title — the one heading pattern the design uses. */
export function SectionHeader({ label, title, description, actions, className, size = 'section' }: SectionHeaderProps) {
  return (
    <div className={cn('flex flex-wrap items-end justify-between gap-4', className)}>
      <div className='min-w-0'>
        {label ? <div className='label-caps mb-2 text-emerald'>{label}</div> : null}
        <h1
          className={cn(
            'font-semibold tracking-tight text-ghost',
            size === 'page' ? 'text-[28px] leading-tight' : 'text-lg'
          )}
        >
          {title}
        </h1>
        {description ? <p className='mt-2 max-w-2xl text-sm leading-relaxed text-ash'>{description}</p> : null}
      </div>
      {actions ? <div className='flex shrink-0 flex-wrap items-center gap-2'>{actions}</div> : null}
    </div>
  )
}

/** Vertical rhythm between the blocks of a page. */
export function Stack({ children, className }: { children: ReactNode; className?: string }) {
  return <div className={cn('flex flex-col gap-6', className)}>{children}</div>
}
