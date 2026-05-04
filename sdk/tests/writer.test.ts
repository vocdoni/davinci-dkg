// DKGWriter tests — createEpoch, registerKey, claimSlot.
// Require a live testnet (RUN_INTEGRATION_TESTS=true).

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { DKGWriter, EpochPhase, buildElGamal, buildEpochId, parseEpochId } from '../src/index.js';
import { makePublicClient, makeWalletClient } from './helpers/accounts.js';
import { mineUntilEpochAllowed, mineUntilSeedAvailable } from './helpers/chain.js';

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
    const { privKey, pubKey } = eg.generateKeyPair();
    const account    = writer.walletClient.account!.address;

    // The Go fixture (flow.test.ts runs before writer.test.ts alphabetically)
    // may have already registered this account.  Use updateKey in that case.
    const existing = await writer.getNode(account);
    const hash = existing.status === 0
      ? await writer.registerKey(privKey)
      : await writer.updateKey(privKey);

    const receipt = await writer.waitForTransaction(hash);
    expect(receipt.status).toBe('success');

    const isActive = await writer.isActive(account);
    expect(isActive).toBe(true);

    // The contract stores the pubkey in RTE form; SDK ElGamal works in TE.
    // Convert before comparing.
    const node = await writer.getNode(account);
    const { fromTEtoRTE } = await import('../src/crypto/babyjub-form.js');
    const [pkX_RTE, pkY_RTE] = fromTEtoRTE(pubKey[0], pubKey[1]);
    expect(node.pubX).toBe(pkX_RTE);
    expect(node.pubY).toBe(pkY_RTE);
  });

  // ── Epoch creation ────────────────────────────────────────────────────────

  it('createEpoch creates a epoch in Registration status', async () => {
    const { enabled } = useHarness();
    if (!enabled) return;

    await mineUntilEpochAllowed(writer.publicClient, writer.managerAddress);
    const currentBlock = await writer.blockNumber();
    const nonceBefore  = await writer.epochNonce();

    const hash = await writer.createEpoch({
      threshold:                 1,
      committeeSize:             1,
      minValidContributions:     1,
      lotteryAlphaBps:           15000,
      seedDelay:                 1,
      committeeSelectionDeadlineBlock: currentBlock + 25n,
      keyAssemblyDeadlineBlock: currentBlock + 50n,
      liveNotBeforeBlock:    currentBlock + 51n,
    });
    const receipt = await writer.waitForTransaction(hash);
    expect(receipt.status).toBe('success');

    // Epoch nonce incremented
    const nonceAfter = await writer.epochNonce();
    expect(nonceAfter).toBe(nonceBefore + 1n);

    // Derive epoch ID
    const prefix  = await writer._managerContract.read.EPOCH_PREFIX();
    const epochId = buildEpochId(prefix, nonceBefore + 1n);

    const epoch = await writer.getEpoch(epochId);
    expect(epoch.status).toBe(EpochPhase.CommitteeSelection);
    expect(epoch.policy.threshold).toBe(1);
    expect(epoch.policy.committeeSize).toBe(1);
  });

  it('buildEpochId and parseEpochId are inverses', () => {
    const prefix = 1337;
    const nonce  = 42n;
    const id     = buildEpochId(prefix, nonce);
    const parsed = parseEpochId(id);
    expect(parsed.prefix).toBe(prefix);
    expect(parsed.nonce).toBe(nonce);
  });

  // ── Slot claiming ─────────────────────────────────────────────────────────

  it('claimSlot claims a slot after seedDelay blocks and advances epoch to Contribution', async () => {
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

    await mineUntilEpochAllowed(writer.publicClient, writer.managerAddress);
    const currentBlock = await writer.blockNumber();
    const seedDelay    = 1;

    const createHash = await writer.createEpoch({
      threshold:                 1,
      committeeSize:             1,
      minValidContributions:     1,
      lotteryAlphaBps:           15000,
      seedDelay,
      committeeSelectionDeadlineBlock: currentBlock + 30n,
      keyAssemblyDeadlineBlock: currentBlock + 60n,
      liveNotBeforeBlock:    currentBlock + 61n,
    });
    await writer.waitForTransaction(createHash);

    const prefix  = await writer._managerContract.read.EPOCH_PREFIX();
    const nonce   = await writer.epochNonce();
    const epochId = buildEpochId(prefix, nonce);

    // Mine past the seed block
    const epoch = await writer.getEpoch(epochId);
    await mineUntilSeedAvailable(writer.publicClient, epoch.seedBlock);

    // Claim slot
    const claimHash = await writer.claimSlot(epochId);
    const claimReceipt = await writer.waitForTransaction(claimHash);
    expect(claimReceipt.status).toBe('success');

    // Epoch should now be in Contribution phase
    const updated = await writer.getEpoch(epochId);
    expect(updated.status).toBe(EpochPhase.KeyAssembly);

    // Our address should be in selectedParticipants
    const participants = await writer.selectedParticipants(epochId);
    expect(participants.map((a) => a.toLowerCase())).toContain(account.toLowerCase());
  });
});
