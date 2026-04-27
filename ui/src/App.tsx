import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { WagmiProvider } from 'wagmi'
import { RainbowKitProvider, darkTheme } from '@rainbow-me/rainbowkit'
import '@rainbow-me/rainbowkit/styles.css'
import { Theme } from '~theme/Theme'
import { ConfigProvider } from '~providers/ConfigProvider'
import { DebugModeProvider } from '~providers/DebugModeProvider'
import { Router } from '~router/Router'
import { wagmiConfig } from '~lib/wagmi'
import { Polling } from '~constants/polling'

// Provider order is load-bearing:
//   Theme           — must wrap everything else so Chakra components in any
//                     descendant (including the config bootstrap screen)
//                     have access to the system tokens.
//   DebugMode       — pure React state; cheap to put outside config so the
//                     debug toggle works even on the config-failed screen.
//   Config          — gates on /config.json before rendering anything that
//                     touches chain/manager. Until config loads, children
//                     don't mount.
//   Wagmi           — depends on config indirectly only; we pass a static
//                     wagmiConfig so order with Config doesn't matter.
//   QueryClient     — wagmi v2 requires it.
//   RainbowKit      — must be inside Wagmi + QueryClient.
//   Router          — last, so route elements can use any of the above hooks.
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchInterval: Polling.default,
      staleTime: Polling.default / 2,
      retry: 1,
    },
  },
})

export function App() {
  return (
    <Theme>
      <DebugModeProvider>
        <ConfigProvider>
          <WagmiProvider config={wagmiConfig}>
            <QueryClientProvider client={queryClient}>
              <RainbowKitProvider theme={darkTheme({ accentColor: '#22d3ee', borderRadius: 'medium' })}>
                <Router />
              </RainbowKitProvider>
            </QueryClientProvider>
          </WagmiProvider>
        </ConfigProvider>
      </DebugModeProvider>
    </Theme>
  )
}
