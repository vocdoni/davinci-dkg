// Network presets; no chain needed.

import { describe, it, expect } from 'vitest';
import { KNOWN_NETWORKS, NODE_DEFAULT_NETWORK, canonicalNetworkName, getNetwork, findNetwork } from '../src/index.js';

describe('network presets', () => {
  it('knows the node default, gnosis, with its public endpoints', () => {
    expect(NODE_DEFAULT_NETWORK).toBe('gnosis');
    const gnosis = getNetwork(NODE_DEFAULT_NETWORK);
    expect(gnosis.chainId).toBe(100);
    expect(gnosis.managerAddress).toBe('0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF');
    expect(gnosis.startBlock).toBe(48_632_905n);
    expect(gnosis.rpcUrls.length).toBeGreaterThan(1);
    expect(findNetwork(100, '0xc6fb38c42ed3fb35d363a702218746d5c7da36bf')).toBe('gnosis');
  });

  it('resolves names and aliases case-insensitively', () => {
    expect(canonicalNetworkName(' Gnosis ')).toBe('gnosis');
    expect(canonicalNetworkName('SEP')).toBe('sepolia');
    expect(canonicalNetworkName('constructor')).toBeUndefined();
    expect(() => getNetwork('mainnet')).toThrow(/unknown network "mainnet"/);
  });

  it('finds the preset of a deployment by chain and manager', () => {
    const sepolia = KNOWN_NETWORKS.sepolia;
    expect(findNetwork(sepolia.chainId, sepolia.managerAddress.toUpperCase().replace('0X', '0x'))).toBe('sepolia');
    expect(findNetwork(1, sepolia.managerAddress)).toBeUndefined();
    expect(findNetwork(sepolia.chainId, '0x00000000000000000000000000000000000000c0')).toBeUndefined();
  });
});
