import { useMemo, useState } from 'react'
import { Pagination, Panel } from '~kit'
import type { IndexedEvent } from '~indexer/types'
import { EventTable } from '~pages/epochs/EventTable'

const PAGE_SIZE = 25

/**
 * Every log this epoch produced, newest first and paginated — a busy epoch
 * emits one event per claim, per contribution, per ciphertext and per partial,
 * so this is the longest list on the page.
 */
export function EventLogPanel({ events }: { events: IndexedEvent[] }) {
  const [page, setPage] = useState(0)
  const ordered = useMemo(() => [...events].reverse(), [events])
  const pageCount = Math.max(1, Math.ceil(ordered.length / PAGE_SIZE))
  const clamped = Math.min(page, pageCount - 1)
  const slice = useMemo(
    () => ordered.slice(clamped * PAGE_SIZE, clamped * PAGE_SIZE + PAGE_SIZE),
    [ordered, clamped]
  )

  return (
    <Panel
      label='Log'
      title={`${events.length.toLocaleString()} events`}
      description='Claims, contributions, finalization, ciphertexts, partials, organizer shares and combines, in block order.'
      bodyClassName='p-0'
    >
      <EventTable
        events={slice}
        maxHeight={PAGE_SIZE * 42 + 40}
        emptyTitle='No events for this epoch'
        emptyDescription='Nothing has been logged against this epoch id yet.'
      />
      <div className='border-t border-charcoal px-4 py-3'>
        <Pagination
          page={clamped}
          pageCount={pageCount}
          pageSize={PAGE_SIZE}
          total={ordered.length}
          onPageChange={setPage}
        />
      </div>
    </Panel>
  )
}
