import { render, type RenderOptions, type RenderResult } from '@testing-library/react'
import { MemoryRouter } from 'react-router-dom'
import type { ReactElement, ReactNode } from 'react'
import { ConfigContext } from '~config/config-context'
import { DEMO_CONFIG, type RuntimeConfig } from '~config/runtime-config'
import { TooltipProvider } from '~kit'
import { DataSourceProvider } from '~data/context'
import { createDataSource } from '~data/create'

/**
 * Renders a kit component with the two contexts every one of them may reach
 * for: the runtime config (explorer URLs, demo flag) and a router (Address and
 * Hash render `<Link>`s). No wallet or query client — nothing in `~kit` needs
 * a chain.
 */
export function renderWithProviders(
  ui: ReactElement,
  options: RenderOptions & { config?: Partial<RuntimeConfig>; route?: string } = {}
): RenderResult {
  const { config, route = '/', ...rest } = options
  const value: RuntimeConfig = { ...DEMO_CONFIG, demo: false, ...config }
  const Wrapper = ({ children }: { children: ReactNode }) => (
    <ConfigContext.Provider value={value}>
      <MemoryRouter initialEntries={[route]}>
        <DataSourceProvider source={createDataSource({ client: null, config: { ...value, demo: true } })}>
          <TooltipProvider>{children}</TooltipProvider>
        </DataSourceProvider>
      </MemoryRouter>
    </ConfigContext.Provider>
  )
  return render(ui, { wrapper: Wrapper, ...rest })
}
