// Monitor / event-watching tests.
// Require a live testnet (RUN_INTEGRATION_TESTS=true).

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import {
  DKGClient,
  DKGWriter,
  EpochPhase,
  buildElGamal,
  buildEpochId,
  waitForEpochPhase,
  watchNewRounds,
  networkSummary,
} from '../src/index.js';
import { makePublicClient, makeWalletClient } from './helpers/accounts.js';
import { mineUntilSeedAvailable } from './helpers/chain.js';

function useHarness() {
  return {
    enabled:        inject('integrationEnabled') as boolean,
    rpcUrl:         inject('rpcUrl')          as string,
    managerAddress: inject('managerAddress')  as `0x${string}`,
  };
}

describe('Monitor utilities', () => {
  let client: DKGClient;
  let writer: DKGWriter; // account #2 to avoid nonce clashes

  beforeAll(() => {
    const { enabled, rpcUrl, managerAddress } = useHarness();
    if (!enabled) return;
    const publicClient = makePublicClient(rpcUrl);
    client = new DKGClient({ publicClient, managerAddress });
    writer = new DKGWriter({
      publicClient,
      walletClient:  makeWalletClient(rpcUrl, 2),
      managerAddress,
    });
  });

  it('networkSummary returns sensible values', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    const summary = await networkSummary(client);
    expect(summary.blockNumber).toBeGreaterThan(0n);
    expect(typeof summary.totalNodes).toBe('bigint');
    expect(typeof summary.activeNodes).toBe('bigint');
    expect(typeof summary.epochNonce).toBe('bigint');
  });

  it('watchNewRounds fires the callback when a epoch is created', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    const seen: `0x${string}`[] = [];
    const unsub = watchNewRounds(client, (epochId) => {
      seen.push(epochId);
    });

    // Register + create a epoch
    const eg = await buildElGamal();
    const { pubKey } = eg.generateKeyPair();
    const account    = writer.walletClient.account!.address;

    const node = await client.getNode(account);
    if (node.status === 0) {
      const regHash = await writer.registerKey(pubKey[0], pubKey[1]);
      await writer.waitForTransaction(regHash);
    }

    const currentBlock = await client.blockNumber();
    const hash = await writer.createEpoch({
      threshold:                 1,
      committeeSize:             1,
      minValidContributions:     1,
      lotteryAlphaBps:           15000,
      seedDelay:                 1,
      registrationDeadlineBlock: currentBlock + 30n,
      contributionDeadlineBlock: currentBlock + 60n,
      finalizeNotBeforeBlock:    currentBlock + 61n,
    });
    await writer.waitForTransaction(hash);

    // Give the subscription a moment to receive the event
    await new Promise((r) => setTimeout(r, 3_000));
    unsub();

    expect(seen.length).toBeGreaterThan(0);
  });

  it('waitForEpochPhase resolves once the epoch reaches the target status', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    // Register account #2's key
    const eg = await buildElGamal();
    const { pubKey } = eg.generateKeyPair();
    const account    = writer.walletClient.account!.address;

    const node = await client.getNode(account);
    if (node.status === 0) {
      const regHash = await writer.registerKey(pubKey[0], pubKey[1]);
      await writer.waitForTransaction(regHash);
    }

    const currentBlock = await client.blockNumber();
    const createHash   = await writer.createEpoch({
      threshold:                 1,
      committeeSize:             1,
      minValidContributions:     1,
      lotteryAlphaBps:           15000,
      seedDelay:                 1,
      registrationDeadlineBlock: currentBlock + 30n,
      contributionDeadlineBlock: currentBlock + 60n,
      finalizeNotBeforeBlock:    currentBlock + 61n,
    });
    await writer.waitForTransaction(createHash);

    const prefix  = await writer._managerContract.read.EPOCH_PREFIX();
    const nonce   = await writer.epochNonce();
    const epochId = buildEpochId(prefix, nonce);

    // Should already be in Registration — resolves immediately
    await waitForEpochPhase(client, epochId, EpochPhase.Registration, {
      intervalMs: 500,
      timeoutMs: 15_000,
    });

    // Mine past seed and claim slot → triggers Contribution
    const epoch = await client.getEpoch(epochId);
    await mineUntilSeedAvailable(client.publicClient, epoch.seedBlock);

    const claimHash = await writer.claimSlot(epochId);
    await writer.waitForTransaction(claimHash);

    await waitForEpochPhase(client, epochId, EpochPhase.Contribution, {
      intervalMs: 500,
      timeoutMs:  30_000,
    });

    const updated = await client.getEpoch(epochId);
    expect(updated.status).toBe(EpochPhase.Contribution);
  });

  it('getEpochCreatedEvents returns historical events', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    const events = await client.getEpochCreatedEvents({ fromBlock: 0n });
    expect(events.length).toBeGreaterThan(0);
    for (const e of events) {
      expect(e.epochId).toMatch(/^0x[0-9a-f]{24}$/i);
      expect(e.blockNumber).toBeGreaterThan(0n);
    }
  });
});
