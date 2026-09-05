// SDK application-lifecycle end-to-end test.
//
// Validates the per-application surface against a live chain:
// registerApplication in both modes (organizer-locked with a Schnorr proof of
// possession of sk_org, revealed later via revealOrganizerSecret; automatic
// with no organizer key at all — the contract always stores the fixed
// identity and a zero secret) and getApplication. Without this, an ABI
// mismatch in writer.ts would only surface in a downstream consumer hitting
// the chain.
//
// The proof of possession is built entirely in TS via `proveOrganizer` from
// sdk/src/schnorr.ts; this is the load-bearing cross-impl assertion — the
// on-chain `_organizerSchnorrChallenge` (Solidity) must agree byte-for-byte
// with the TS-side keccak transcript or the registration reverts.

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { spawn } from 'node:child_process';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import {
  AppMode,
  DKGClient,
  DKGWriter,
  normalizeAppPolicy,
  proveOrganizer,
  randomOrganizerSecret,
  type AppPolicyInput,
  randomAid,
} from '../src/index.js';
import { makePublicClient, makeWalletClient } from './helpers/accounts.js';

const __filename = fileURLToPath(import.meta.url);
const __dirname  = path.dirname(__filename);

function useHarness() {
  return {
    enabled:        inject('integrationEnabled') as boolean,
    rpcUrl:         inject('rpcUrl')           as string,
    managerAddress: inject('managerAddress')   as `0x${string}`,
    addressesFile:  inject('addressesFile')    as string,
  };
}

interface FixtureCreateResult {
  epochId: `0x${string}`;
  share:   string;
}

async function runGoFixture(args: string[]) {
  const repoRoot    = path.resolve(__dirname, '..', '..');
  const fixtureMain = path.join(repoRoot, 'cmd', 'sdk-test-fixture');
  return new Promise<{ status: number | null; stdout: string; stderr: string }>((resolve, reject) => {
    const child = spawn('go', ['run', fixtureMain, ...args], { cwd: repoRoot });
    let stdout = '', stderr = '';
    child.stdout.on('data', (d) => { stdout += d.toString('utf8'); });
    child.stderr.on('data', (d) => { stderr += d.toString('utf8'); });
    const t = setTimeout(() => { child.kill('SIGKILL'); reject(new Error('Go fixture timed out')); }, 600_000);
    child.on('error', (e) => { clearTimeout(t); reject(e); });
    child.on('close', (code) => { clearTimeout(t); resolve({ status: code, stdout, stderr }); });
  });
}

function lastJsonLine<T>(stdout: string): T | null {
  const line = stdout.split('\n').map((l) => l.trim()).filter((l) => l.startsWith('{')).at(-1);
  if (!line) return null;
  try { return JSON.parse(line) as T; } catch { return null; }
}


describe('SDK application lifecycle end-to-end (live chain)', () => {
  let client:  DKGClient;
  let writer:  DKGWriter;
  let fixture: FixtureCreateResult | null = null;

  beforeAll(async () => {
    const { enabled, rpcUrl, managerAddress, addressesFile } = useHarness();
    if (!enabled) return;

    client = new DKGClient({ publicClient: makePublicClient(rpcUrl), managerAddress });
    writer = new DKGWriter({
      publicClient: makePublicClient(rpcUrl),
      walletClient: makeWalletClient(rpcUrl, 1),
      managerAddress,
    });

    const out = await runGoFixture(['--rpc-url', rpcUrl, '--addresses-file', addressesFile, '--action=create', '--keys=3']);
    if (out.status !== 0) {
      console.warn('[application-e2e] fixture create failed — skipping. stderr:', out.stderr.slice(0, 500));
      return;
    }
    fixture = lastJsonLine<FixtureCreateResult>(out.stdout);
  });

  it('writer.registerApplication round-trips through getApplication', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const skOrg = randomOrganizerSecret();
    // Every field optional: the defaults are organizer-locked, registrant-only.
    const policy: AppPolicyInput = { maxCiphertexts: 0 };

    const tx = await writer.registerApplication(fixture.epochId, aid, policy, skOrg);
    await writer.publicClient.waitForTransactionReceipt({ hash: tx });

    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(true);
    expect(app.policy.mode).toBe(AppMode.OrganizerLocked);
    expect(app.policy.maxCiphertexts).toBe(0);
    // An empty list means "the registering address only"; nothing is resolved.
    expect(app.policy.submitters).toEqual([]);
    expect(app.policy.openSubmission).toBe(false);
    expect(app.policy.decryptNotAfter).toBe(0n);
    // Organizer-locked: nothing about sk_org is on chain.
    expect(app.organizerSecret).toBe(0n);

    // The stored PK_org must be sk_org·G. proveOrganizer returns RTE coords
    // (the on-chain transcript form); client.getApplication converts the
    // stored record back to TE, so compare through the form converter.
    const { pkOrgX, pkOrgY } = proveOrganizer(skOrg, fixture.epochId, aid, 1n);
    const { fromRTEtoTE } = await import('../src/crypto/babyjub-form.js');
    const [pkOrgX_TE, pkOrgY_TE] = fromRTEtoTE(pkOrgX, pkOrgY);
    expect(app.organizerPK[0]).toBe(pkOrgX_TE);
    expect(app.organizerPK[1]).toBe(pkOrgY_TE);
  }, 900_000);

  it('an automatic application has no organizer key and skips the Schnorr proof', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const submitter = writer.walletClient.account!.address;
    const deadline = BigInt(Math.floor(Date.now() / 1000) + 24 * 3600);
    // Automatic mode takes no organizer secret at all — pass none.
    const tx = await writer.registerApplication(
      fixture.epochId,
      aid,
      { mode: AppMode.Automatic, submitters: [submitter], maxCiphertexts: 4, decryptNotAfter: deadline },
    );
    await writer.publicClient.waitForTransactionReceipt({ hash: tx });

    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(true);
    expect(app.policy.mode).toBe(AppMode.Automatic);
    expect(app.policy.submitters.map((a) => a.toLowerCase())).toEqual([submitter.toLowerCase()]);
    expect(app.policy.maxCiphertexts).toBe(4);
    expect(app.policy.decryptNotAfter).toBe(deadline);
    // No organizer key at all: the contract stores the fixed identity and a
    // zero secret, unconditionally.
    expect(app.organizerSecret).toBe(0n);
    expect(app.organizerPK).toEqual([0n, 1n]);
    expect(await client.getOrganizerPK(fixture.epochId, aid)).toEqual(app.organizerPK);
    expect(await client.isDecryptionOpen(fixture.epochId, aid)).toBe(true);

    const events = await client.getApplicationRegisteredEvents({ epochId: fixture.epochId });
    const mine = events.find((e) => e.aid.toLowerCase() === aid.toLowerCase());
    expect(mine?.mode).toBe(AppMode.Automatic);
  }, 900_000);

  it('a tampered Schnorr response is rejected on-chain', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const sk = 1234567890123456789n;
    const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, fixture.epochId, aid);

    // Flip one bit of `z` and confirm the on-chain verifier reverts. The
    // writer builds the proof itself, so go through the raw ABI here.
    await expect(
      writer.publicClient.simulateContract({
        address: await client.getAppManagerAddress(),
        abi: (await import('../src/abi.js')).dkgAppManagerAbi,
        functionName: 'registerApplication',
        args: [
          fixture.epochId, aid, normalizeAppPolicy(),
          pkOrgX, pkOrgY, proof.ax, proof.ay, proof.z + 1n,
        ],
        account: writer.walletClient.account!.address,
      }),
    ).rejects.toThrow();
  }, 900_000);

  it('revealOrganizerSecret publishes sk_org once for a locked application', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const skOrg = randomOrganizerSecret();
    const policy: AppPolicyInput = { maxCiphertexts: 0 };
    const regTx = await writer.registerApplication(fixture.epochId, aid, policy, skOrg);
    await writer.publicClient.waitForTransactionReceipt({ hash: regTx });
    expect((await client.getApplication(fixture.epochId, aid)).organizerSecret).toBe(0n);

    const revealTx = await writer.revealOrganizerSecret(fixture.epochId, aid, skOrg);
    await writer.publicClient.waitForTransactionReceipt({ hash: revealTx });
    expect((await client.getApplication(fixture.epochId, aid)).organizerSecret).toBe(skOrg);

    // AlreadyRevealed() on a second call.
    await expect(writer.revealOrganizerSecret(fixture.epochId, aid, skOrg)).rejects.toThrow();
  }, 900_000);

  it('getApplication returns exists=false for an unregistered aid', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(false);
  });
});
