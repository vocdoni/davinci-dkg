// Store-backed suggestions under a page's own search box.
//
// The shell's global search box routes by shape; this shows what the *store*
// actually knows, so an epoch id pasted into the operators filter (or an
// address pasted into the applications filter) is one click from its page
// instead of a dead end. Shared by both stream D list pages; it lives in a
// page folder because that is what this stream owns.

import { Link } from 'react-router-dom'
import { useStoreSearch } from '~data/hooks'
import type { SearchKind } from '~indexer/selectors'
import { ExternalIcon } from '~kit'
import { cn } from '~lib/cn'

export interface StoreSearchHitsProps {
  query: string
  /** Kinds the page's own table already filters on, so nothing is offered twice. */
  skip?: SearchKind[]
  limit?: number
  className?: string
}

export function StoreSearchHits({ query, skip = [], limit = 5, className }: StoreSearchHitsProps) {
  const results = useStoreSearch(query, limit + skip.length)
  const hits = results.filter((hit) => !skip.includes(hit.kind)).slice(0, limit)
  if (query.trim() === '' || hits.length === 0) return null

  return (
    <div className={cn('rounded-md border border-charcoal bg-onyx/40 p-3', className)}>
      <div className='label-caps mb-2 text-[10px] text-pewter'>elsewhere in the explorer</div>
      <ul className='m-0 flex flex-col gap-1 p-0'>
        {hits.map((hit) => (
          <li key={`${hit.kind}:${hit.id}`} className='min-w-0'>
            {hit.external ? (
              <a
                href={hit.href}
                target='_blank'
                rel='noreferrer noopener'
                className='flex min-w-0 items-baseline gap-2 text-[13px] text-silver transition-colors hover:text-emerald'
              >
                <span className='truncate font-mono'>{hit.label}</span>
                <span className='shrink-0 text-[11px] text-ash'>{hit.detail}</span>
                <ExternalIcon size={11} className='shrink-0 text-ash' />
              </a>
            ) : (
              <Link
                to={hit.href}
                className='flex min-w-0 items-baseline gap-2 text-[13px] text-silver transition-colors hover:text-emerald'
              >
                <span className='truncate font-mono'>{hit.label}</span>
                <span className='shrink-0 text-[11px] text-ash'>{hit.detail}</span>
              </Link>
            )}
          </li>
        ))}
      </ul>
    </div>
  )
}
