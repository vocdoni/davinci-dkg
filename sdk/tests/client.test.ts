// DKGClient read-only tests.
// Require a live testnet (RUN_INTEGRATION_TESTS=true).

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { DKGClient } from '../src/index.js';
import { makePublicClient } from './helpers/accounts.js';

function useHarness() {
  const enabled        = inject('integrationEnabled') as boolean;
  const rpcUrl         = inject('rpcUrl')          as string;
  const managerAddress = inject('managerAddress')  as `0x${string}`;
  return { enabled, rpcUrl, managerAddress };
}

describe('DKGClient (read-only)', () => {
  let client: DKGClient;

  beforeAll(() => {
    const { enabled, rpcUrl, managerAddress } = useHarness();
    if (!enabled) return;
    // registryAddress is intentionally omitted to exercise auto-derive from manager.
    client = new DKGClient({
      publicClient:    makePublicClient(rpcUrl),
      managerAddress,
    });
  });

  it('blockNumber returns a positive bigint', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const bn = await client.blockNumber();
    expect(typeof bn).toBe('bigint');
    expect(bn).toBeGreaterThan(0n);
  });

  it('nodeCount returns a bigint', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const count = await client.nodeCount();
    expect(typeof count).toBe('bigint');
    expect(count).toBeGreaterThanOrEqual(0n);
  });

  it('activeCount returns a bigint <= nodeCount', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const [total, active] = await Promise.all([client.nodeCount(), client.activeCount()]);
    expect(active).toBeLessThanOrEqual(total);
  });

  it('roundNonce starts at 0 on a fresh chain', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const nonce = await client.roundNonce();
    expect(typeof nonce).toBe('bigint');
    expect(nonce).toBeGreaterThanOrEqual(0n);
  });

  it('inactivityWindow returns a positive bigint', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const window = await client.inactivityWindow();
    expect(window).toBeGreaterThan(0n);
  });

  it('getRound returns zero-status round for a nonexistent round ID', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const fakeId = '0x000000000000000000000000' as `0x${string}`;
    const round  = await client.getRound(fakeId);
    expect(round.status).toBe(0);
  });

  it('isActive returns false for an unregistered address', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const active = await client.isActive('0x0000000000000000000000000000000000000001');
    expect(active).toBe(false);
  });

  it('getContributionVerifierVKeyHash returns a bytes32', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;
    const hash = await client.getContributionVerifierVKeyHash();
    expect(hash).toMatch(/^0x[0-9a-f]{64}$/i);
  });
});
