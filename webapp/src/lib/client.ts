import { createPublicClient, http, type PublicClient } from 'viem';
import { defineChain } from 'viem';

export interface RuntimeConfig {
  rpcUrl: string;
  managerAddress: `0x${string}`;
  registryAddress: `0x${string}`;
  chainId: number;
  chainName: string;
  startBlock?: number;
}

let cachedConfig: RuntimeConfig | null = null;
let cachedBaseConfig: RuntimeConfig | null = null;
let cachedClient: PublicClient | null = null;

const RPC_OVERRIDE_KEY = 'dkg-explorer:rpc-url';

export function getRpcOverride(): string | null {
  try {
    return localStorage.getItem(RPC_OVERRIDE_KEY);
  } catch {
    return null;
  }
}

export function setRpcOverride(url: string | null) {
  try {
    if (url && url.trim() !== '') {
      localStorage.setItem(RPC_OVERRIDE_KEY, url.trim());
    } else {
      localStorage.removeItem(RPC_OVERRIDE_KEY);
    }
  } catch {
    // localStorage can throw in private-mode browsers; the override just
    // won't persist across reloads, which is acceptable.
  }
  cachedConfig = null;
  cachedClient = null;
  // cachedBaseConfig is intentionally kept — it reflects /config.json and
  // is invariant across override changes.
}

export async function loadConfig(): Promise<RuntimeConfig> {
  if (cachedConfig) return cachedConfig;
  if (!cachedBaseConfig) {
    const res = await fetch('/config.json', { cache: 'no-store' });
    if (!res.ok) {
      throw new Error(`Failed to load /config.json: ${res.status}`);
    }
    cachedBaseConfig = (await res.json()) as RuntimeConfig;
  }
  const override = getRpcOverride();
  cachedConfig = override ? { ...cachedBaseConfig, rpcUrl: override } : { ...cachedBaseConfig };
  return cachedConfig;
}

export async function loadBaseConfig(): Promise<RuntimeConfig> {
  if (cachedBaseConfig) return cachedBaseConfig;
  await loadConfig();
  return cachedBaseConfig!;
}

export async function getClient(): Promise<PublicClient> {
  if (cachedClient) return cachedClient;
  const cfg = await loadConfig();
  const chain = defineChain({
    id: cfg.chainId,
    name: cfg.chainName || `chain-${cfg.chainId}`,
    nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 },
    rpcUrls: { default: { http: [cfg.rpcUrl] } },
  });
  cachedClient = createPublicClient({
    chain,
    transport: http(cfg.rpcUrl),
  }) as PublicClient;
  return cachedClient;
}

export function resetClient() {
  cachedClient = null;
  cachedConfig = null;
}

export function resetAll() {
  cachedClient = null;
  cachedConfig = null;
  cachedBaseConfig = null;
}
