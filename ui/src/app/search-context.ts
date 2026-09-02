import { createContext, useContext, useEffect } from 'react'
import type { SearchResolver, SearchTarget } from './search'

export interface SearchApi {
  /** Current text in the global search box. */
  query: string
  setQuery: (query: string) => void
  /** Resolve without navigating — for a preview row under the input. */
  resolve: (query: string) => SearchTarget
  /** Resolve and navigate (or open the explorer in a new tab). */
  submit: (query: string) => SearchTarget
  /** The last unresolved query, so the box can explain itself. */
  error: string | null
  registerResolver: (resolver: SearchResolver) => () => void
}

export const SearchContext = createContext<SearchApi | null>(null)

export function useSearch(): SearchApi {
  const api = useContext(SearchContext)
  if (!api) throw new Error('useSearch must be used inside <SearchProvider>')
  return api
}

/**
 * Contribute a resolver that runs *before* the built-in shape matching.
 *
 * This is the extension point for the indexer stream: once the store knows
 * which epoch an application id belongs to, register a resolver that returns
 * the `/applications/:epoch/:aid` route for it, and the global search box picks
 * it up with no change to the shell.
 */
export function useRegisterSearchResolver(resolver: SearchResolver): void {
  const { registerResolver } = useSearch()
  useEffect(() => registerResolver(resolver), [registerResolver, resolver])
}
