// Network presets; no chain needed.

import { describe, it, expect } from 'vitest';
import { KNOWN_NETWORKS, NODE_DEFAULT_NETWORK, canonicalNetworkName, getNetwork, findNetwork } from '../src/index.js';

describe('network presets', () => {
  it('knows the node default, optimism, with its public endpoints', () => {
    expect(NODE_DEFAULT_NETWORK).toBe('optimism');
    const op = KNOWN_NETWORKS[NODE_DEFAULT_NETWORK];
    expect(op.chainId).toBe(10);
    expect(op.rpcUrls.length).toBeGreaterThan(1);
    expect(KNOWN_NETWORKS.gnosis).toBeUndefined();
  });

  it('refuses a preset with no DKGManager yet and serves a deployed one', () => {
    const zero = '0x0000000000000000000000000000000000000000';
    const op = KNOWN_NETWORKS[NODE_DEFAULT_NETWORK];
    if (op.managerAddress === zero) {
      expect(() => getNetwork(NODE_DEFAULT_NETWORK)).toThrow('network "optimism" has no DKGManager yet');
    } else {
      expect(getNetwork(NODE_DEFAULT_NETWORK)).toBe(op);
    }
    expect(getNetwork('sepolia')).toBe(KNOWN_NETWORKS.sepolia);
    // findNetwork stays a plain match.
    expect(findNetwork(op.chainId, op.managerAddress)).toBe('optimism');
  });

  it('resolves names and aliases case-insensitively', () => {
    expect(canonicalNetworkName(' OptimISM ')).toBe('optimism');
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
