import { cn } from '~lib/cn'
import { Button } from './Button'
import { ChevronLeftIcon, ChevronRightIcon } from './icons'

export interface PaginationProps {
  /** Zero-based page index. */
  page: number
  pageCount: number
  onPageChange: (page: number) => void
  /** Shown as "x–y of total" when both are given. */
  pageSize?: number
  total?: number
  className?: string
}

/** Compact pager for lists that are paginated rather than virtualised. */
export function Pagination({ page, pageCount, onPageChange, pageSize, total, className }: PaginationProps) {
  const clamped = Math.min(Math.max(page, 0), Math.max(pageCount - 1, 0))
  const from = pageSize != null ? clamped * pageSize + 1 : null
  const to = pageSize != null && total != null ? Math.min((clamped + 1) * pageSize, total) : null
  return (
    <div className={cn('flex items-center justify-between gap-4 text-xs text-ash', className)}>
      <span className='font-mono tnum'>
        {from != null && to != null && total != null
          ? `${from}–${to} of ${total}`
          : `page ${clamped + 1} / ${pageCount || 1}`}
      </span>
      <div className='flex items-center gap-1'>
        <Button
          size='icon'
          variant='subtle'
          aria-label='Previous page'
          disabled={clamped <= 0}
          onClick={() => onPageChange(clamped - 1)}
        >
          <ChevronLeftIcon />
        </Button>
        <span className='px-2 font-mono tnum text-silver'>
          {clamped + 1} / {pageCount || 1}
        </span>
        <Button
          size='icon'
          variant='subtle'
          aria-label='Next page'
          disabled={clamped >= pageCount - 1}
          onClick={() => onPageChange(clamped + 1)}
        >
          <ChevronRightIcon />
        </Button>
      </div>
    </div>
  )
}
