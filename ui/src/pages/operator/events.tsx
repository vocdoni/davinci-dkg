import { useMemo, useState } from 'react'
import { Link } from 'react-router-dom'
import { BlockCell, EmptyState, Pagination, Timeline, TimelineRow, TxCell } from '~kit'
import type { IndexedEvent } from '~indexer/types'
import { describeEvent, type NonceLookup } from './describe'

const PAGE_SIZE = 20

/**
 * Every event this operator produced, newest first and paginated — an operator
 * that has served fifty epochs has thousands of them, and the page must not
 * render a thousand rows to show twenty.
 */
export function OperatorEventLog({ events, nonceOf }: { events: IndexedEvent[]; nonceOf?: NonceLookup }) {
  const [page, setPage] = useState(0)
  const newestFirst = useMemo(() => [...events].reverse(), [events])
  const pageCount = Math.max(1, Math.ceil(newestFirst.length / PAGE_SIZE))
  const current = Math.min(page, pageCount - 1)
  const slice = newestFirst.slice(current * PAGE_SIZE, (current + 1) * PAGE_SIZE)

  if (events.length === 0) {
    return <EmptyState compact title='No events' description='Nothing this address did appears in the indexed range.' />
  }

  return (
    <div className='flex flex-col gap-4'>
      <Timeline>
        {slice.map((event, i) => {
          const described = describeEvent(event, nonceOf)
          return (
            <TimelineRow
              key={`${event.block}:${event.logIndex}`}
              tone={described.tone}
              last={i === slice.length - 1}
              title={
                described.href ? (
                  <Link to={described.href} className='transition-colors hover:text-emerald'>
                    {described.title}
                  </Link>
                ) : (
                  described.title
                )
              }
              description={described.detail}
              meta={<BlockCell block={event.block} />}
              right={event.tx ? <TxCell hash={event.tx} /> : null}
            />
          )
        })}
      </Timeline>
      {pageCount > 1 ? (
        <Pagination
          page={current}
          pageCount={pageCount}
          onPageChange={setPage}
          pageSize={PAGE_SIZE}
          total={newestFirst.length}
        />
      ) : null}
    </div>
  )
}
