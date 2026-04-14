import { createWalletClient, custom, defineChain, type WalletClient, type Address } from 'viem';
import { loadConfig, getClient, type RuntimeConfig } from './client';

declare global {
  interface Window {
    ethereum?: {
      request: (args: { method: string; params?: unknown[] }) => Promise<unknown>;
      on: (event: string, handler: (...args: unknown[]) => void) => void;
      removeListener: (event: string, handler: (...args: unknown[]) => void) => void;
    };
  }
}

export function hasWallet(): boolean {
  return typeof window !== 'undefined' && !!window.ethereum;
}

export async function connectWallet(): Promise<{ address: Address; walletClient: WalletClient }> {
  if (!window.ethereum) {
    throw new Error('No browser wallet detected. Install MetaMask or a compatible wallet.');
  }

  const cfg = await loadConfig();
  const chainIdHex = `0x${cfg.chainId.toString(16)}`;

  // Request account access
  const accounts = (await window.ethereum.request({ method: 'eth_requestAccounts' })) as Address[];
  if (!accounts.length) throw new Error('No accounts returned by wallet');
  const address = accounts[0];

  // Switch to the DKG chain, adding it if needed
  await switchToChain(cfg, chainIdHex);

  const chain = defineChain({
    id: cfg.chainId,
    name: cfg.chainName || `chain-${cfg.chainId}`,
    nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 },
    rpcUrls: { default: { http: [cfg.rpcUrl] } },
  });

  const walletClient = createWalletClient({
    account: address,
    chain,
    transport: custom(window.ethereum),
  });

  return { address, walletClient };
}

async function switchToChain(cfg: RuntimeConfig, chainIdHex: string): Promise<void> {
  try {
    await window.ethereum!.request({
      method: 'wallet_switchEthereumChain',
      params: [{ chainId: chainIdHex }],
    });
  } catch (err: unknown) {
    const code = (err as { code?: number }).code;
    if (code === 4902 || code === -32603) {
      // Chain not known to the wallet — add it
      await window.ethereum!.request({
        method: 'wallet_addEthereumChain',
        params: [{
          chainId: chainIdHex,
          chainName: cfg.chainName || `Chain ${cfg.chainId}`,
          nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 },
          rpcUrls: [cfg.rpcUrl],
        }],
      });
    } else {
      throw err;
    }
  }
}

export { getClient, loadConfig };
