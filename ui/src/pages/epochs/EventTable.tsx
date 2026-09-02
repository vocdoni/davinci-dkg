import { useMemo } from 'react'
import { Address, BlockCell, DataTable, EmptyState, Hash, TxCell, type AnyColumnDef } from '~kit'
import type { IndexedEvent } from '~indexer/types'
import { paths } from '~routes/paths'
import { cn } from '~lib/cn'
import { EVENT_TONE, EVENT_TONE_CLASS, eventSummary } from './events'

export interface EventTableProps {
  events: IndexedEvent[]
  /** Show the epoch column — off on the epoch page, where it is constant. */
  showEpoch?: boolean
  loading?: boolean
  loadingRows?: number
  maxHeight?: number
  emptyTitle?: string
  emptyDescription?: string
}

const eventKey = (event: IndexedEvent): string => `${event.block}:${event.logIndex}:${event.name}`

/**
 * One row per log: what happened, in which block, to whom, and the transaction
 * that carried it. Nothing is folded away — the feed on the overview and the
 * epoch's full log are the same table with one column swapped.
 */
export function EventTable({
  events,
  showEpoch = false,
  loading = false,
  loadingRows = 8,
  maxHeight = 520,
  emptyTitle = 'No events yet',
  emptyDescription = 'Nothing has been logged in this range.',
}: EventTableProps) {
  const columns = useMemo<AnyColumnDef<IndexedEvent>[]>(() => {
    const cols: AnyColumnDef<IndexedEvent>[] = [
      {
        id: 'name',
        header: 'Event',
        accessorKey: 'name',
        enableSorting: false,
        meta: { width: '208px' },
        cell: ({ row }) => (
          <span className={cn('font-mono text-[12px]', EVENT_TONE_CLASS[EVENT_TONE[row.original.name]])}>
            {row.original.name}
          </span>
        ),
      },
      {
        id: 'block',
        header: 'Block',
        accessorKey: 'block',
        enableSorting: false,
        meta: { numeric: true, width: '104px' },
        cell: ({ row }) => <BlockCell block={row.original.block} />,
      },
    ]
    if (showEpoch) {
      cols.push({
        id: 'epoch',
        header: 'Epoch',
        enableSorting: false,
        meta: { width: '132px' },
        cell: ({ row }) =>
          row.original.epoch ? (
            <Hash value={row.original.epoch} chars={6} copy={false} href={paths.epoch(row.original.epoch)} />
          ) : (
            <span className='text-ash'>—</span>
          ),
      })
    }
    cols.push(
      {
        id: 'actor',
        header: 'Actor',
        enableSorting: false,
        meta: { width: '148px' },
        cell: ({ row }) =>
          row.original.actor ? (
            <Address value={row.original.actor} copy={false} explorer={false} to={paths.operator(row.original.actor)} />
          ) : (
            <span className='text-ash'>—</span>
          ),
      },
      {
        id: 'detail',
        header: 'Detail',
        enableSorting: false,
        cell: ({ row }) => <span className='font-mono text-[11px] text-ash'>{eventSummary(row.original)}</span>,
      },
      {
        id: 'tx',
        header: 'Transaction',
        enableSorting: false,
        meta: { width: '124px' },
        cell: ({ row }) =>
          row.original.tx ? <TxCell hash={row.original.tx} /> : <span className='text-ash'>—</span>,
      }
    )
    return cols
  }, [showEpoch])

  return (
    <div className='overflow-x-auto scroll-slim'>
      <div className='min-w-[860px]'>
        <DataTable
          data={events}
          columns={columns}
          loading={loading}
          loadingRows={loadingRows}
          maxHeight={maxHeight}
          getRowId={eventKey}
          empty={<EmptyState compact title={emptyTitle} description={emptyDescription} />}
        />
      </div>
    </div>
  )
}
