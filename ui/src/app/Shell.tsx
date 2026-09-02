import { useIndexerSearchResolver } from '~data/hooks'
import { useOptionalDataSource } from '~data/context'
import { useRegisterSearchResolver } from '~app/search-context'
import { Outlet, ScrollRestoration } from 'react-router-dom'
import { PageContainer, TooltipProvider } from '~kit'
import { SearchProvider } from './SearchProvider'
import { TopBar } from './TopBar'
import { Footer } from './Footer'

/**
 * App frame every route renders inside. It owns the Radix tooltip provider, so
 * any kit component with a tooltip works on any page without extra wiring.
 *
 * Pages receive a 1400 px container and are responsible only for their own
 * content; a page that needs full bleed can break out with a negative margin.
 */
export function Shell() {
  return (
    <SearchProvider>
      <IndexerSearchBridge />
      <TooltipProvider>
        <div className='flex min-h-screen flex-col bg-obsidian'>
          <TopBar />
          <main className='flex-1 py-8'>
            <PageContainer>
              <Outlet />
            </PageContainer>
          </main>
          <Footer />
        </div>
      </TooltipProvider>
      <ScrollRestoration />
    </SearchProvider>
  )
}

// Routes epoch ids, application ids, addresses and tx hashes through the
// indexer's store before the shape-based fallback of the search box.
function IndexerSearchBridge() {
  // Tests and isolated renders may mount the shell without a data source.
  const source = useOptionalDataSource()
  return source ? <IndexerSearchBridgeInner /> : null
}

function IndexerSearchBridgeInner() {
  const resolver = useIndexerSearchResolver()
  useRegisterSearchResolver(resolver)
  return null
}
