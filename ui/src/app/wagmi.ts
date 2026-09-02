import { createConfig, http } from 'wagmi'
import { defineChain, type Chain } from 'viem'
import { connectorsForWallets } from '@rainbow-me/rainbowkit'
import { injectedWallet, metaMaskWallet, walletConnectWallet } from '@rainbow-me/rainbowkit/wallets'
import type { RuntimeConfig } from '~config/runtime-config'

// WalletConnect projectId comes from build-time env. Without it the
// WalletConnect option silently disappears from the picker — better than
// crashing the boot, and irrelevant in dev where MetaMask is the norm.
const projectId = (import.meta.env.VITE_WALLETCONNECT_PROJECT_ID as string | undefined) ?? ''

/**
 * The chain is built from `/config.json` rather than picked from a hard-coded
 * list: the same bundle has to serve Sepolia, a local Anvil testnet and any
 * future deployment without a rebuild.
 */
export function chainFromConfig(config: RuntimeConfig): Chain {
  return defineChain({
    id: config.chainId,
    name: config.chainName,
    nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 },
    rpcUrls: { default: { http: [config.rpcUrl] } },
    blockExplorers: config.explorerUrl ? { default: { name: 'Explorer', url: config.explorerUrl } } : undefined,
    testnet: config.chainId !== 1,
  })
}

export function createWagmiConfig(config: RuntimeConfig) {
  const chain = chainFromConfig(config)
  const connectors = connectorsForWallets(
    [
      {
        groupName: 'Popular',
        wallets: projectId ? [metaMaskWallet, walletConnectWallet, injectedWallet] : [metaMaskWallet, injectedWallet],
      },
    ],
    { appName: 'davinci-dkg explorer', projectId: projectId || 'davinci-dkg-no-walletconnect' }
  )

  return createConfig({
    chains: [chain],
    connectors,
    transports: { [chain.id]: http(config.rpcUrl) },
    // One poll per block time; every "live" number in the UI derives from the
    // head, so a tighter interval buys nothing but RPC bill.
    pollingInterval: 12_000,
  })
}
