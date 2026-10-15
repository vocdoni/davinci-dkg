// Well-known deployments, mirroring `KnownNetworks` in config/networks.go.
// Only the DKGManager is listed: the registry and the app manager are read
// from it on chain (see DKGConfig).

import type { Address } from 'viem';

export interface NetworkDeployment {
  /** EIP-155 chain id. */
  chainId: number;
  /** The deployed DKGManager. */
  managerAddress: Address;
  /** Block the DKGManager was deployed at: the floor for event log scans. */
  startBlock: bigint;
  /** Public JSON-RPC endpoints; empty when the network has none built in. */
  rpcUrls: readonly string[];
}

export const KNOWN_NETWORKS: Readonly<Record<string, NetworkDeployment>> = {
  optimism: {
    chainId: 10,
    managerAddress: '0x9E799A1D22710056E62c849e3FFE19Dd01db2D41',
    startBlock: 158_245_366n,
    rpcUrls: ['https://mainnet.optimism.io', 'https://optimism-rpc.publicnode.com'],
  },
  sepolia: {
    chainId: 11155111,
    managerAddress: '0xc73b7a868eca6ac7e3e647e2665aa16a793cf551',
    startBlock: 11_668_198n,
    rpcUrls: [],
  },
};

/** The network `davinci-dkg-node` joins when started without `--network` or `--manager`. */
export const NODE_DEFAULT_NETWORK = 'optimism';

const ALIASES: Readonly<Record<string, string>> = { sep: 'sepolia' };

/** Resolve a network name or alias (case-insensitive) to its canonical name, or undefined. */
export function canonicalNetworkName(name: string): string | undefined {
  const lower = name.trim().toLowerCase();
  const own = (obj: object, key: string) => Object.prototype.hasOwnProperty.call(obj, key);
  const canonical = own(ALIASES, lower) ? ALIASES[lower] : lower;
  return own(KNOWN_NETWORKS, canonical) ? canonical : undefined;
}

/** The deployment of a known network; throws on an unknown name or a preset with no DKGManager yet. */
export function getNetwork(name: string): NetworkDeployment {
  const canonical = canonicalNetworkName(name);
  if (canonical === undefined) {
    throw new Error(`unknown network "${name}"; supported: ${Object.keys(KNOWN_NETWORKS).join(', ')}`);
  }
  const dep = KNOWN_NETWORKS[canonical];
  if (/^0x0{40}$/i.test(dep.managerAddress)) {
    throw new Error(`network "${canonical}" has no DKGManager yet`);
  }
  return dep;
}

/** The known network a deployment belongs to (same chain and manager), or undefined. */
export function findNetwork(chainId: number, managerAddress: string): string | undefined {
  const manager = managerAddress.toLowerCase();
  return Object.keys(KNOWN_NETWORKS).find((name) => {
    const dep = KNOWN_NETWORKS[name];
    return dep.chainId === chainId && dep.managerAddress.toLowerCase() === manager;
  });
}
