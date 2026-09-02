import { Button, Callout, ProgressBar } from '~kit'
import type { IndexerHandle } from '~data/hooks'
import { blockOrNull } from '~pages/epochs/cadence'

/**
 * The scan is visible or it is not there: while the indexer backfills, every
 * page shows partial data, and the user is owed the reason. Silent once the
 * indexer is live and caught up.
 */
export function IndexerStatus({ indexer }: { indexer: IndexerHandle }) {
  const { status } = indexer
  const errors = status.errors.slice(-3)
  const from = blockOrNull(status.fromBlock)
  const last = blockOrNull(status.lastBlock)
  const head = blockOrNull(status.headBlock)

  if (status.phase === 'error' && errors.length > 0) {
    return (
      <Callout
        tone='danger'
        title='The indexer stopped on an error'
        actions={
          <Button size='sm' variant='ghost' onClick={() => void indexer.refresh()}>
            Retry
          </Button>
        }
      >
        <ul className='mt-1 space-y-1 font-mono text-[11px]'>
          {errors.map((error) => (
            <li key={`${error.at}:${error.scope}`}>
              {error.scope}: {error.message}
            </li>
          ))}
        </ul>
      </Callout>
    )
  }

  if (status.phase === 'live' || status.phase === 'idle') return null

  const scanned = from != null && last != null ? Math.max(0, last - from) : 0
  const total = from != null && head != null ? Math.max(1, head - from) : 1

  return (
    <Callout
      tone='info'
      title={status.phase === 'loading' ? 'Starting the indexer' : 'Indexing history'}
      actions={
        <span className='font-mono text-[11px] text-ash'>{status.requests.toLocaleString()} getLogs requests</span>
      }
    >
      <div className='mt-2 max-w-2xl'>
        <ProgressBar
          value={scanned}
          total={total}
          label={`block ${last?.toLocaleString() ?? '—'} of ${head?.toLocaleString() ?? '—'}`}
        />
        <p className='mt-2 text-[12px] text-ash'>
          {status.eventCount.toLocaleString()} events indexed so far. Pages fill in as the scan advances.
        </p>
      </div>
    </Callout>
  )
}
