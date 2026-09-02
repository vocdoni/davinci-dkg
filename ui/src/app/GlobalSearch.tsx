import { useState, type FormEvent } from 'react'
import { Input, SearchIcon } from '~kit'
import { cn } from '~lib/cn'
import { useSearch } from './search-context'

/**
 * One box for every identifier in the protocol. It routes by shape — see
 * `search.ts` — and says why when it can't.
 */
export function GlobalSearch({ className }: { className?: string }) {
  const { query, setQuery, submit, error } = useSearch()
  const [focused, setFocused] = useState(false)

  const onSubmit = (event: FormEvent) => {
    event.preventDefault()
    if (query.trim() !== '') submit(query)
  }

  return (
    <form onSubmit={onSubmit} className={cn('relative', className)} role='search'>
      <Input
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        onFocus={() => setFocused(true)}
        onBlur={() => setFocused(false)}
        size='sm'
        mono
        spellCheck={false}
        autoComplete='off'
        aria-label='Search epochs, applications, operators and transactions'
        placeholder='Epoch, application, address, block, tx…'
        iconLeft={<SearchIcon size={14} />}
      />
      {error && focused ? (
        <p className='absolute top-full right-0 left-0 z-30 mt-1 rounded-sm border border-red/30 bg-carbon px-3 py-2 text-[11px] text-red'>
          {error}
        </p>
      ) : null}
    </form>
  )
}
