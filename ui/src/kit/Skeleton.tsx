import { cn } from '~lib/cn'

export interface SkeletonProps {
  className?: string
  /** Convenience for inline sizing when a utility class would be noise. */
  width?: number | string
  height?: number | string
  rounded?: 'sm' | 'md' | 'pill'
}

/** Loading placeholder: onyx block, slow opacity pulse. Never a spinner in a table. */
export function Skeleton({ className, width, height, rounded = 'sm' }: SkeletonProps) {
  return (
    <div
      aria-hidden='true'
      style={{ width, height }}
      className={cn(
        'animate-skeleton bg-onyx',
        rounded === 'pill' ? 'rounded-pill' : rounded === 'md' ? 'rounded-md' : 'rounded-sm',
        className
      )}
    />
  )
}

export function SkeletonText({ lines = 3, className }: { lines?: number; className?: string }) {
  return (
    <div className={cn('space-y-2', className)}>
      {Array.from({ length: lines }, (_, i) => (
        <Skeleton key={i} className={cn('h-3', i === lines - 1 ? 'w-2/3' : 'w-full')} />
      ))}
    </div>
  )
}
