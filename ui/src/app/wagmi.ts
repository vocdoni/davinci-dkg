import { createConfig, http } from 'wagmi'
import { defineChain, type Chain } from 'viem'
import { gnosis } from 'viem/chains'
import { connectorsForWallets } from '@rainbow-me/rainbowkit'
import { injectedWallet, metaMaskWallet, walletConnectWallet } from '@rainbow-me/rainbowkit/wallets'
import type { RuntimeConfig } from '~config/runtime-config'

// WalletConnect projectId comes from build-time env. Without it the
// WalletConnect option silently disappears from the picker — better than
// crashing the boot, and irrelevant in dev where MetaMask is the norm.
const projectId = (import.meta.env.VITE_WALLETCONNECT_PROJECT_ID as string | undefined) ?? ''

// Chains whose currency, block time or testnet flag differ from the default
// (Ether, 12 s blocks, testnet unless mainnet).
const KNOWN_CHAINS: readonly Chain[] = [gnosis]

function knownChain(chainId: number): Chain | undefined {
  return KNOWN_CHAINS.find((chain) => chain.id === chainId)
}

/**
 * The chain is built from `/config.json` rather than picked from a hard-coded
 * list: the same bundle has to serve Gnosis Chain, Sepolia, a local Anvil
 * testnet and any future deployment without a rebuild. A known chain only
 * lends its native currency, block time and testnet flag.
 */
export function chainFromConfig(config: RuntimeConfig): Chain {
  const known = knownChain(config.chainId)
  return defineChain({
    id: config.chainId,
    name: config.chainName,
    nativeCurrency: known?.nativeCurrency ?? { name: 'Ether', symbol: 'ETH', decimals: 18 },
    blockTime: known?.blockTime,
    rpcUrls: { default: { http: [config.rpcUrl] } },
    blockExplorers: config.explorerUrl ? { default: { name: 'Explorer', url: config.explorerUrl } } : undefined,
    testnet: known ? known.testnet === true : config.chainId !== 1,
  })
}

/** Seconds per block of a known chain; undefined keeps the indexer's 12 s. */
export function blockTimeSeconds(chainId: number): number | undefined {
  const ms = knownChain(chainId)?.blockTime
  return ms === undefined ? undefined : ms / 1000
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
