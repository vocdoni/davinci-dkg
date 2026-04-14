// DKGWriter tests — createRound, registerKey, claimSlot.
// Require a live testnet (RUN_INTEGRATION_TESTS=true).

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { DKGWriter, RoundStatus, buildElGamal, buildRoundId, parseRoundId } from '../src/index.js';
import { makePublicClient, makeWalletClient } from './helpers/accounts.js';
import { mineUntilSeedAvailable } from './helpers/chain.js';

function useHarness() {
  return {
    enabled:        inject('integrationEnabled') as boolean,
    rpcUrl:         inject('rpcUrl')          as string,
    managerAddress: inject('managerAddress')  as `0x${string}`,
  };
}

describe('DKGWriter', () => {
  let writer: DKGWriter;

  beforeAll(() => {
    const { enabled, rpcUrl, managerAddress } = useHarness();
    if (!enabled) return;
    writer = new DKGWriter({
      publicClient:    makePublicClient(rpcUrl),
      walletClient:    makeWalletClient(rpcUrl, 1), // use account #1 so we don't clash with fixture (account #0)
      managerAddress,
    });
  });

  // ── Registry ─────────────────────────────────────────────────────────────

  it('registerKey / updateKey registers a BabyJubJub key and marks the node as active', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    const eg         = await buildElGamal();
    const { pubKey } = eg.generateKeyPair();
    const account    = writer.walletClient.account!.address;

    // The Go fixture (flow.test.ts runs before writer.test.ts alphabetically)
    // may have already registered this account.  Use updateKey in that case.
    const existing = await writer.getNode(account);
    const hash = existing.status === 0
      ? await writer.registerKey(pubKey[0], pubKey[1])
      : await writer.updateKey(pubKey[0], pubKey[1]);

    const receipt = await writer.waitForTransaction(hash);
    expect(receipt.status).toBe('success');

    const isActive = await writer.isActive(account);
    expect(isActive).toBe(true);

    const node = await writer.getNode(account);
    expect(node.pubX).toBe(pubKey[0]);
    expect(node.pubY).toBe(pubKey[1]);
  });

  // ── Round creation ────────────────────────────────────────────────────────

  it('createRound creates a round in Registration status', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    const currentBlock = await writer.blockNumber();
    const nonceBefore  = await writer.roundNonce();

    const hash = await writer.createRound({
      threshold:                 1,
      committeeSize:             1,
      minValidContributions:     1,
      lotteryAlphaBps:           15000,
      seedDelay:                 1,
      registrationDeadlineBlock: currentBlock + 25n,
      contributionDeadlineBlock: currentBlock + 50n,
      disclosureAllowed:         false,
    });
    const receipt = await writer.waitForTransaction(hash);
    expect(receipt.status).toBe('success');

    // Round nonce incremented
    const nonceAfter = await writer.roundNonce();
    expect(nonceAfter).toBe(nonceBefore + 1n);

    // Derive round ID
    const prefix  = await writer._managerContract.read.ROUND_PREFIX();
    const roundId = buildRoundId(prefix, nonceBefore + 1n);

    const round = await writer.getRound(roundId);
    expect(round.status).toBe(RoundStatus.Registration);
    expect(round.policy.threshold).toBe(1);
    expect(round.policy.committeeSize).toBe(1);
  });

  it('buildRoundId and parseRoundId are inverses', () => {
    const prefix = 1337;
    const nonce  = 42n;
    const id     = buildRoundId(prefix, nonce);
    const parsed = parseRoundId(id);
    expect(parsed.prefix).toBe(prefix);
    expect(parsed.nonce).toBe(nonce);
  });

  // ── Slot claiming ─────────────────────────────────────────────────────────

  it('claimSlot claims a slot after seedDelay blocks and advances round to Contribution', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    // Re-register account #1's key if needed (idempotent: if already registered, update)
    const eg       = await buildElGamal();
    const { pubKey } = eg.generateKeyPair();
    const account  = writer.walletClient.account!.address;

    const node = await writer.getNode(account);
    if (node.status === 0) {
      const regHash = await writer.registerKey(pubKey[0], pubKey[1]);
      await writer.waitForTransaction(regHash);
    }

    const currentBlock = await writer.blockNumber();
    const seedDelay    = 1;

    const createHash = await writer.createRound({
      threshold:                 1,
      committeeSize:             1,
      minValidContributions:     1,
      lotteryAlphaBps:           15000,
      seedDelay,
      registrationDeadlineBlock: currentBlock + 30n,
      contributionDeadlineBlock: currentBlock + 60n,
      disclosureAllowed:         false,
    });
    await writer.waitForTransaction(createHash);

    const prefix  = await writer._managerContract.read.ROUND_PREFIX();
    const nonce   = await writer.roundNonce();
    const roundId = buildRoundId(prefix, nonce);

    // Mine past the seed block
    const round = await writer.getRound(roundId);
    await mineUntilSeedAvailable(writer.publicClient, round.seedBlock);

    // Claim slot
    const claimHash = await writer.claimSlot(roundId);
    const claimReceipt = await writer.waitForTransaction(claimHash);
    expect(claimReceipt.status).toBe('success');

    // Round should now be in Contribution phase
    const updated = await writer.getRound(roundId);
    expect(updated.status).toBe(RoundStatus.Contribution);

    // Our address should be in selectedParticipants
    const participants = await writer.selectedParticipants(roundId);
    expect(participants.map((a) => a.toLowerCase())).toContain(account.toLowerCase());
  });
});
