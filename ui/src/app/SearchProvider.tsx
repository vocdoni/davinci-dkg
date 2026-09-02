import { useCallback, useMemo, useRef, useState, type ReactNode } from 'react'
import { useNavigate } from 'react-router-dom'
import { useRuntimeConfig } from '~config/config-context'
import { resolveSearch, type SearchResolver, type SearchTarget } from './search'
import { SearchContext, type SearchApi } from './search-context'

/** Owns the global search box's state and the resolver registry. */
export function SearchProvider({ children }: { children: ReactNode }) {
  const navigate = useNavigate()
  const { explorerUrl } = useRuntimeConfig()
  const [query, setQuery] = useState('')
  const [error, setError] = useState<string | null>(null)
  const resolvers = useRef<SearchResolver[]>([])

  const registerResolver = useCallback((resolver: SearchResolver) => {
    resolvers.current = [...resolvers.current, resolver]
    return () => {
      resolvers.current = resolvers.current.filter((r) => r !== resolver)
    }
  }, [])

  const resolve = useCallback(
    (value: string) => resolveSearch(value, { explorerUrl }, resolvers.current),
    [explorerUrl]
  )

  const submit = useCallback(
    (value: string): SearchTarget => {
      const target = resolve(value)
      if (target.kind === 'route') {
        setError(null)
        setQuery('')
        navigate(target.path)
      } else if (target.kind === 'external') {
        setError(null)
        window.open(target.url, '_blank', 'noopener,noreferrer')
      } else {
        setError(target.label)
      }
      return target
    },
    [navigate, resolve]
  )

  const api = useMemo<SearchApi>(
    () => ({
      query,
      setQuery: (value: string) => {
        setQuery(value)
        if (error) setError(null)
      },
      resolve,
      submit,
      error,
      registerResolver,
    }),
    [query, error, resolve, submit, registerResolver]
  )

  return <SearchContext.Provider value={api}>{children}</SearchContext.Provider>
}
