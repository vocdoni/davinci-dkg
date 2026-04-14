// Viem client factories for tests.

import {
  createPublicClient,
  createWalletClient,
  http,
  defineChain,
} from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import { ANVIL_PRIVATE_KEYS } from './harness.js';

export const CHAIN_ID = 1337;

/** Build the viem chain object for the local Anvil instance. */
export function makeChain(rpcUrl: string) {
  return defineChain({
    id: CHAIN_ID,
    name: 'Anvil',
    nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 },
    rpcUrls: { default: { http: [rpcUrl] } },
  });
}

/** Create a read-only viem PublicClient connected to the testnet. */
export function makePublicClient(rpcUrl: string) {
  return createPublicClient({
    chain: makeChain(rpcUrl),
    transport: http(rpcUrl),
  });
}

/**
 * Create a WalletClient for the given Anvil account index (0–4).
 * Account 0 is the default organizer / deployer key.
 */
export function makeWalletClient(rpcUrl: string, accountIndex: 0 | 1 | 2 | 3 | 4 = 0) {
  const account = privateKeyToAccount(ANVIL_PRIVATE_KEYS[accountIndex]);
  return createWalletClient({
    chain:     makeChain(rpcUrl),
    transport: http(rpcUrl),
    account,
  });
}

/** Convenience: create both clients for the same account. */
export function makeClients(rpcUrl: string, accountIndex: 0 | 1 | 2 | 3 | 4 = 0) {
  return {
    publicClient: makePublicClient(rpcUrl),
    walletClient: makeWalletClient(rpcUrl, accountIndex),
  };
}
