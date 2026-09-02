import type { HTMLAttributes, ReactNode } from 'react'
import { cn } from '~lib/cn'

export type CardLevel = 'carbon' | 'onyx'

export interface CardProps extends HTMLAttributes<HTMLDivElement> {
  /** Surface level: `carbon` (level 1, default) or `onyx` (level 2). */
  level?: CardLevel
  /** Lift the border to warm-gray and the surface to onyx on hover. */
  hover?: boolean
  /** Drop the default padding (tables and charts manage their own). */
  flush?: boolean
}

export function Card({ level = 'carbon', hover = false, flush = false, className, children, ...rest }: CardProps) {
  return (
    <div
      {...rest}
      className={cn(
        'rounded-md border border-charcoal transition-colors duration-200',
        level === 'carbon' ? 'bg-carbon' : 'bg-onyx',
        hover && 'hover:border-warm-gray hover:bg-onyx',
        !flush && 'p-5',
        className
      )}
    >
      {children}
    </div>
  )
}

export interface CardHeaderProps {
  title: ReactNode
  /** 13 px uppercase eyebrow above the title. */
  label?: ReactNode
  description?: ReactNode
  actions?: ReactNode
  className?: string
}

export function CardHeader({ title, label, description, actions, className }: CardHeaderProps) {
  return (
    <div className={cn('flex items-start justify-between gap-4 border-b border-charcoal px-5 py-4', className)}>
      <div className='min-w-0'>
        {label ? <div className='label-caps mb-1.5 text-emerald'>{label}</div> : null}
        <h2 className='truncate text-[15px] font-semibold text-ghost'>{title}</h2>
        {description ? <p className='mt-1 text-[13px] leading-relaxed text-ash'>{description}</p> : null}
      </div>
      {actions ? <div className='flex shrink-0 items-center gap-2'>{actions}</div> : null}
    </div>
  )
}

export function CardBody({ className, children, ...rest }: HTMLAttributes<HTMLDivElement>) {
  return (
    <div {...rest} className={cn('p-5', className)}>
      {children}
    </div>
  )
}

export function CardFooter({ className, children, ...rest }: HTMLAttributes<HTMLDivElement>) {
  return (
    <div {...rest} className={cn('flex items-center gap-2 border-t border-charcoal px-5 py-3', className)}>
      {children}
    </div>
  )
}

/** A `Card` with a header bar and flush body — the standard explorer panel. */
export function Panel({
  title,
  label,
  description,
  actions,
  children,
  className,
  bodyClassName,
}: CardHeaderProps & { children: ReactNode; bodyClassName?: string }) {
  return (
    <Card flush className={cn('overflow-hidden', className)}>
      <CardHeader title={title} label={label} description={description} actions={actions} />
      <div className={cn('p-5', bodyClassName)}>{children}</div>
    </Card>
  )
}
