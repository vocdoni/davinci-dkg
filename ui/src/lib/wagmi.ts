import { http, createConfig } from 'wagmi'
import { sepolia, foundry } from 'wagmi/chains'
import { connectorsForWallets } from '@rainbow-me/rainbowkit'
import { metaMaskWallet, walletConnectWallet, injectedWallet } from '@rainbow-me/rainbowkit/wallets'

// WalletConnect projectId is read from build-time env. If it's missing the
// WalletConnect option silently disappears from the wallet picker — better
// than crashing the boot, and zero risk in dev where most contributors
// connect via MetaMask anyway.
const projectId = (import.meta.env.VITE_WALLETCONNECT_PROJECT_ID as string | undefined) ?? ''

const wallets = [
  {
    groupName: 'Popular',
    wallets: projectId ? [metaMaskWallet, walletConnectWallet, injectedWallet] : [metaMaskWallet, injectedWallet],
  },
]

const connectors = connectorsForWallets(wallets, {
  appName: 'davinci-dkg explorer',
  projectId: projectId || 'davinci-dkg-no-walletconnect',
})

export const wagmiConfig = createConfig({
  chains: [sepolia, foundry],
  connectors,
  transports: {
    [sepolia.id]: http(),
    [foundry.id]: http('http://127.0.0.1:8545'),
  },
})
