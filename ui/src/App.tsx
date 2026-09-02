import { useMemo, type ReactNode } from 'react'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { RainbowKitProvider } from '@rainbow-me/rainbowkit'
import { RouterProvider } from 'react-router-dom'
import { WagmiProvider, usePublicClient } from 'wagmi'
import '@rainbow-me/rainbowkit/styles.css'
import { ConfigProvider } from '~config/ConfigProvider'
import { useRuntimeConfig } from '~config/config-context'
import { createWagmiConfig } from '~app/wagmi'
import { walletTheme } from '~app/rainbowkit-theme'
import { router } from '~routes/router'
import { DataSourceProvider } from '~data/context'
import { createDataSource } from '~data/create'

// Provider order is load-bearing:
//   Config      — gates on /config.json; nothing chain-aware mounts before it.
//   Wagmi       — built *from* that config, so one bundle serves any chain.
//   QueryClient — wagmi v2 requires it, and the indexer stream shares it.
//   RainbowKit  — must sit inside Wagmi + QueryClient.
//   Router      — last, so route elements can use all of the above. The Radix
//                 tooltip provider lives in the Shell, one level further in.
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 6_000,
      retry: 1,
      refetchOnWindowFocus: false,
    },
  },
})

// The indexer (or the demo fixture) is the single data source every page reads.
function DataProviders({ children }: { children: ReactNode }) {
  const config = useRuntimeConfig()
  const client = usePublicClient()
  const source = useMemo(() => createDataSource({ client, config }), [client, config])
  return <DataSourceProvider source={source}>{children}</DataSourceProvider>
}

function ChainProviders({ children }: { children: ReactNode }) {
  const config = useRuntimeConfig()
  const wagmiConfig = useMemo(() => createWagmiConfig(config), [config])
  return (
    <WagmiProvider config={wagmiConfig}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider theme={walletTheme} modalSize='compact'>
          <DataProviders>{children}</DataProviders>
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  )
}

export function App() {
  return (
    <ConfigProvider>
      <ChainProviders>
        <RouterProvider router={router} />
      </ChainProviders>
    </ConfigProvider>
  )
}
