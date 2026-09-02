import { describe, expect, it } from 'vitest'
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { createMemoryRouter, RouterProvider } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { WagmiProvider } from 'wagmi'
import { ConfigContext } from '~config/config-context'
import { DEMO_CONFIG } from '~config/runtime-config'
import { createWagmiConfig } from './wagmi'
import { Shell } from './Shell'

// Demo mode on purpose: it is the only configuration in which the shell needs
// no wallet stack, so this smoke test covers the top bar, chain pill, search
// box and footer without RainbowKit's portals.
function renderShell(initialPath = '/') {
  const router = createMemoryRouter(
    [
      {
        path: '/',
        element: <Shell />,
        children: [
          { index: true, element: <p>page body</p> },
          { path: 'operators/:address', element: <p>operator page</p> },
        ],
      },
    ],
    { initialEntries: [initialPath] }
  )
  const wagmiConfig = createWagmiConfig(DEMO_CONFIG)
  return render(
    <ConfigContext.Provider value={DEMO_CONFIG}>
      <WagmiProvider config={wagmiConfig}>
        <QueryClientProvider client={new QueryClient()}>
          <RouterProvider router={router} />
        </QueryClientProvider>
      </WagmiProvider>
    </ConfigContext.Provider>
  )
}

describe('Shell', () => {
  it('renders the brand, the primary nav and the outlet', () => {
    renderShell()
    expect(screen.getByText('davinci-dkg')).toBeInTheDocument()
    for (const label of ['Overview', 'Epochs', 'Operators', 'Applications', 'Playground', 'Docs']) {
      expect(screen.getAllByRole('link', { name: label }).length).toBeGreaterThan(0)
    }
    expect(screen.getByText('page body')).toBeInTheDocument()
  })

  it('shows the chain, the demo badge and a link to the kit', () => {
    renderShell()
    expect(screen.getByText(DEMO_CONFIG.chainName)).toBeInTheDocument()
    expect(screen.getByText('demo')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: 'Design kit' })).toHaveAttribute('href', '/kit')
  })

  it('routes a searched address to the operator page', async () => {
    const user = userEvent.setup()
    renderShell()
    const [box] = screen.getAllByRole('textbox', { name: /Search epochs/ })
    await user.type(box as HTMLElement, '0x3f9b338706a31f26d49159478015c8aaeab908ad{Enter}')
    expect(await screen.findByText('operator page')).toBeInTheDocument()
    expect(screen.queryByText('page body')).toBeNull()
  })
})
