import { useRef, type ReactNode } from 'react'
import {
  flexRender,
  getCoreRowModel,
  getSortedRowModel,
  useReactTable,
  type Header,
  type OnChangeFn,
  type Row,
  type SortingState,
} from '@tanstack/react-table'
import { useVirtualizer } from '@tanstack/react-virtual'
import { cn } from '~lib/cn'
import { ArrowDownIcon, ArrowUpIcon } from './icons'
import { EmptyState } from './EmptyState'
import { Skeleton } from './Skeleton'
import type { AnyColumnDef, CellAlign } from './table-types'

export type { AnyColumnDef, CellAlign } from './table-types'

const ALIGN: Record<CellAlign, string> = { left: 'text-left', right: 'text-right', center: 'text-center' }
const JUSTIFY: Record<CellAlign, string> = { left: 'justify-start', right: 'justify-end', center: 'justify-center' }

export interface DataTableProps<T> {
  data: T[]
  columns: AnyColumnDef<T>[]
  /** Skeleton rows instead of data. */
  loading?: boolean
  loadingRows?: number
  /** Rendered when `data` is empty and not loading. */
  empty?: ReactNode
  getRowId?: (row: T, index: number) => string
  onRowClick?: (row: T) => void
  /** Controlled sorting; omit both for uncontrolled with `initialSorting`. */
  sorting?: SortingState
  onSortingChange?: OnChangeFn<SortingState>
  initialSorting?: SortingState
  /**
   * Windowed body. Required above ~200 rows; needs `maxHeight` so the
   * container, not the page, is the scroller.
   */
  virtualized?: boolean
  rowHeight?: number
  maxHeight?: number
  stickyHeader?: boolean
  className?: string
}

/**
 * The explorer's one table. Dense rows, hairline separators, sticky header,
 * click-to-sort headers, and an optional virtualised body so a 300-operator
 * registry costs the same as a 30-row one.
 */
export function DataTable<T>({
  data,
  columns,
  loading = false,
  loadingRows = 8,
  empty,
  getRowId,
  onRowClick,
  sorting,
  onSortingChange,
  initialSorting,
  virtualized = false,
  rowHeight = 40,
  maxHeight = 560,
  stickyHeader = true,
  className,
}: DataTableProps<T>) {
  const table = useReactTable({
    data,
    columns,
    getRowId,
    state: sorting ? { sorting } : undefined,
    initialState: initialSorting ? { sorting: initialSorting } : undefined,
    onSortingChange,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
  })

  const rows = table.getRowModel().rows
  const headers = table.getHeaderGroups()[0]?.headers ?? []

  if (loading) {
    return (
      <div className={cn('w-full', className)}>
        <HeaderRow headers={headers} sticky={false} onSort={undefined} />
        <div>
          {Array.from({ length: loadingRows }, (_, i) => (
            <div
              key={i}
              className='flex items-center gap-4 border-b border-charcoal px-4'
              style={{ height: rowHeight }}
            >
              {headers.map((h) => (
                <div key={h.id} className='flex-1' style={trackStyle(h.column.columnDef.meta?.width)}>
                  <Skeleton className='h-3 w-3/4' />
                </div>
              ))}
            </div>
          ))}
        </div>
      </div>
    )
  }

  if (rows.length === 0) {
    return (
      <div className={cn('w-full', className)}>
        <HeaderRow headers={headers} sticky={false} onSort={undefined} />
        {empty ?? <EmptyState title='Nothing here yet' description='No rows matched this view.' />}
      </div>
    )
  }

  if (virtualized) {
    return (
      <VirtualBody
        rows={rows}
        headers={headers}
        rowHeight={rowHeight}
        maxHeight={maxHeight}
        onRowClick={onRowClick}
        className={className}
      />
    )
  }

  return (
    <div className={cn('w-full overflow-x-auto scroll-slim', className)} style={{ maxHeight }}>
      <table className='w-full border-collapse text-[13px]'>
        <thead className={cn(stickyHeader && 'sticky top-0 z-10')}>
          <tr className='bg-carbon'>
            {headers.map((header) => {
              const meta = header.column.columnDef.meta
              const align = meta?.align ?? (meta?.numeric ? 'right' : 'left')
              const sortable = header.column.getCanSort()
              return (
                <th
                  key={header.id}
                  scope='col'
                  title={meta?.headerTooltip}
                  style={{ width: meta?.width }}
                  className={cn(
                    'label-caps border-b border-charcoal px-4 py-2.5 text-[10px] whitespace-nowrap text-pewter',
                    ALIGN[align],
                    sortable && 'cursor-pointer select-none hover:text-ghost'
                  )}
                  onClick={sortable ? header.column.getToggleSortingHandler() : undefined}
                >
                  <span className={cn('inline-flex items-center gap-1', align === 'right' && 'flex-row-reverse')}>
                    {flexRender(header.column.columnDef.header, header.getContext())}
                    <SortGlyph direction={header.column.getIsSorted()} visible={sortable} />
                  </span>
                </th>
              )
            })}
          </tr>
        </thead>
        <tbody>
          {rows.map((row) => (
            <tr
              key={row.id}
              onClick={onRowClick ? () => onRowClick(row.original) : undefined}
              className={cn(
                'border-b border-charcoal/60 last:border-b-0',
                onRowClick && 'cursor-pointer transition-colors hover:bg-onyx'
              )}
            >
              {row.getVisibleCells().map((cell) => {
                const meta = cell.column.columnDef.meta
                const align = meta?.align ?? (meta?.numeric ? 'right' : 'left')
                return (
                  <td
                    key={cell.id}
                    className={cn(
                      'px-4 py-2 text-silver',
                      ALIGN[align],
                      meta?.numeric && 'font-mono tnum whitespace-nowrap'
                    )}
                  >
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </td>
                )
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

type Headers<T> = Header<T, unknown>[]

function trackStyle(width?: string) {
  return width ? { flex: `0 0 ${width}`, width } : { flex: '1 1 0', minWidth: 0 }
}

function SortGlyph({ direction, visible }: { direction: false | 'asc' | 'desc'; visible: boolean }) {
  if (!visible) return null
  if (direction === 'asc') return <ArrowUpIcon size={11} className='text-emerald' />
  if (direction === 'desc') return <ArrowDownIcon size={11} className='text-emerald' />
  return <ArrowDownIcon size={11} className='text-charcoal' />
}

function HeaderRow<T>({
  headers,
  sticky,
  onSort,
}: {
  headers: Headers<T>
  sticky: boolean
  onSort: undefined | (() => void)
}) {
  return (
    <div
      className={cn(
        'flex items-center gap-4 border-b border-charcoal bg-carbon px-4 py-2.5',
        sticky && 'sticky top-0 z-10'
      )}
      onClick={onSort}
    >
      {headers.map((header) => {
        const meta = header.column.columnDef.meta
        const align = meta?.align ?? (meta?.numeric ? 'right' : 'left')
        const sortable = header.column.getCanSort()
        return (
          <div
            key={header.id}
            style={trackStyle(meta?.width)}
            title={meta?.headerTooltip}
            onClick={sortable ? header.column.getToggleSortingHandler() : undefined}
            className={cn(
              'label-caps flex min-w-0 items-center gap-1 text-[10px] text-pewter',
              JUSTIFY[align],
              sortable && 'cursor-pointer select-none hover:text-ghost'
            )}
          >
            <span className='truncate'>{flexRender(header.column.columnDef.header, header.getContext())}</span>
            <SortGlyph direction={header.column.getIsSorted()} visible={sortable} />
          </div>
        )
      })}
    </div>
  )
}

function VirtualBody<T>({
  rows,
  headers,
  rowHeight,
  maxHeight,
  onRowClick,
  className,
}: {
  rows: Row<T>[]
  headers: Headers<T>
  rowHeight: number
  maxHeight: number
  onRowClick?: (row: T) => void
  className?: string
}) {
  const scrollRef = useRef<HTMLDivElement>(null)
  const virtualizer = useVirtualizer({
    count: rows.length,
    getScrollElement: () => scrollRef.current,
    estimateSize: () => rowHeight,
    overscan: 12,
  })

  return (
    <div className={cn('w-full', className)}>
      <HeaderRow headers={headers} sticky={false} onSort={undefined} />
      <div ref={scrollRef} className='overflow-y-auto scroll-slim' style={{ maxHeight }}>
        <div style={{ height: virtualizer.getTotalSize(), position: 'relative' }}>
          {virtualizer.getVirtualItems().map((virtualRow) => {
            const row = rows[virtualRow.index]
            if (!row) return null
            return (
              <div
                key={row.id}
                data-index={virtualRow.index}
                ref={virtualizer.measureElement}
                onClick={onRowClick ? () => onRowClick(row.original) : undefined}
                style={{ transform: `translateY(${virtualRow.start}px)` }}
                className={cn(
                  'absolute top-0 left-0 flex w-full items-center gap-4 border-b border-charcoal/60 px-4 text-[13px]',
                  onRowClick && 'cursor-pointer transition-colors hover:bg-onyx'
                )}
              >
                {row.getVisibleCells().map((cell) => {
                  const meta = cell.column.columnDef.meta
                  const align = meta?.align ?? (meta?.numeric ? 'right' : 'left')
                  return (
                    <div
                      key={cell.id}
                      style={{ ...trackStyle(meta?.width), height: rowHeight }}
                      className={cn(
                        'flex min-w-0 items-center text-silver',
                        JUSTIFY[align],
                        meta?.numeric && 'font-mono tnum whitespace-nowrap'
                      )}
                    >
                      <span className='truncate'>{flexRender(cell.column.columnDef.cell, cell.getContext())}</span>
                    </div>
                  )
                })}
              </div>
            )
          })}
        </div>
      </div>
    </div>
  )
}
