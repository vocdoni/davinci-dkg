// `DataSourceProvider` — the single place a page's data comes from.
//
//   <DataSourceProvider source={createDataSource({ demo, client, config })}>
//     <App />
//   </DataSourceProvider>
//
// The provider owns the source's lifetime: it starts it on mount and stops it
// on unmount, so a page never has to think about the scan loop.

import { createContext, useContext, useEffect, type ReactNode } from 'react'
import type { DataSource } from './source'

const DataSourceContext = createContext<DataSource | null>(null)

export interface DataSourceProviderProps {
  source: DataSource
  /** Start the source with the provider (default true). */
  autoStart?: boolean
  children: ReactNode
}

export function DataSourceProvider({
  source,
  autoStart = true,
  children,
}: DataSourceProviderProps): JSX.Element {
  useEffect(() => {
    if (!autoStart) return
    source.start()
    return () => {
      source.stop()
    }
  }, [source, autoStart])

  return <DataSourceContext.Provider value={source}>{children}</DataSourceContext.Provider>
}

export function useDataSource(): DataSource {
  const source = useContext(DataSourceContext)
  if (!source) throw new Error('useDataSource must be used inside <DataSourceProvider>')
  return source
}

/** Null instead of throwing — for components that render outside the provider. */
export function useOptionalDataSource(): DataSource | null {
  return useContext(DataSourceContext)
}
