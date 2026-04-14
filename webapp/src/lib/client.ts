import { createPublicClient, http, defineChain, type PublicClient } from 'viem';
import { DKGClient } from '@vocdoni/davinci-dkg-sdk';

export interface RuntimeConfig {
  rpcUrl: string;
  managerAddress: `0x${string}`;
  /** Registry address served by the node for display purposes. May be absent in idle mode. */
  registryAddress?: `0x${string}`;
  chainId: number;
  chainName: string;
  startBlock?: number;
}

let cachedConfig: RuntimeConfig | null = null;
let cachedBaseConfig: RuntimeConfig | null = null;
let cachedClient: PublicClient | null = null;
let cachedDKGClient: DKGClient | null = null;

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
  cachedDKGClient = null;
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

/**
 * Return (and cache) a DKGClient configured with the current runtime config.
 *
 * The same instance is reused across React Query calls within a single
 * browser session.  Call `resetClient()` after an RPC override change to
 * force a fresh client on the next request.
 */
export async function getDKGClient(): Promise<DKGClient> {
  if (cachedDKGClient) return cachedDKGClient;
  const [cfg, publicClient] = await Promise.all([loadConfig(), getClient()]);
  cachedDKGClient = new DKGClient({
    publicClient,
    managerAddress: cfg.managerAddress,
    // registryAddress is optional — the SDK derives it from the manager on first use.
    ...(cfg.registryAddress ? { registryAddress: cfg.registryAddress } : {}),
  });
  return cachedDKGClient;
}

export function resetClient() {
  cachedClient = null;
  cachedConfig = null;
  cachedDKGClient = null;
}

export function resetAll() {
  cachedClient = null;
  cachedConfig = null;
  cachedBaseConfig = null;
  cachedDKGClient = null;
}
