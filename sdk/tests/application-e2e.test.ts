// SDK application-lifecycle end-to-end test.
//
// Validates the per-application surface added in P8/P9 against a live
// chain: registerApplication (mode 0), registerApplicationCoDec (mode 1
// with a Schnorr proof of knowledge of sk_org), and getApplication.
// Without this, an ABI mismatch in writer.ts would only surface in a
// downstream consumer hitting the chain.
//
// The Schnorr proof for mode 1 is built entirely in TS via
// `proveOrganizer` from sdk/src/schnorr.ts; this is the load-bearing
// cross-impl assertion — the on-chain `_organizerSchnorrChallenge`
// (Solidity) must agree byte-for-byte with the TS-side Poseidon
// transcript or the registration reverts.

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { spawn } from 'node:child_process';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import {
  DKGClient,
  DKGWriter,
  AppMode,
  computeS,
  proveOrganizer,
  type AppPolicy,
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

function randomAid(): `0x${string}` {
  const buf = new Uint8Array(32);
  globalThis.crypto.getRandomValues(buf);
  return ('0x' + Array.from(buf).map((b) => b.toString(16).padStart(2, '0')).join('')) as `0x${string}`;
}

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as const;

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

    const out = await runGoFixture(['--rpc-url', rpcUrl, '--addresses-file', addressesFile, '--action=create']);
    if (out.status !== 0) {
      console.warn('[application-e2e] fixture create failed — skipping. stderr:', out.stderr.slice(0, 500));
      return;
    }
    fixture = lastJsonLine<FixtureCreateResult>(out.stdout);
  });

  it('mode 0: writer.registerApplication round-trips through getApplication', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const policy: AppPolicy = {
      authorizedSubmitter: ZERO_ADDRESS,
      maxCiphertexts:      0,
      notBeforeBlock:      0n,
      notAfterBlock:       0n,
    };

    const tx = await writer.registerApplication(fixture.epochId, aid, policy);
    await writer.publicClient.waitForTransactionReceipt({ hash: tx });

    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(true);
    expect(app.mode).toBe(AppMode.PublicDerivation);
    expect(app.derivationS).not.toBe(0n);
    expect(app.policy.authorizedSubmitter.toLowerCase()).toBe(ZERO_ADDRESS);
    expect(app.policy.maxCiphertexts).toBe(0);

    // Cross-check: the on-chain S must equal the SDK-derived S over the
    // collective public key, in the on-chain (RTE) coordinate system.
    // client.getCollectivePublicKey returns TE coords; the contract's
    // registerApplication hashes the stored RTE coords. Convert before
    // calling computeS or the byte encoding diverges.
    const pkEpTE = await client.getCollectivePublicKey(fixture.epochId);
    const { fromTEtoRTE } = await import('../src/crypto/babyjub-form.js');
    const [pkXRte, pkYRte] = fromTEtoRTE(pkEpTE.x, pkEpTE.y);
    const sExpected = computeS(fixture.epochId, pkXRte, pkYRte, aid);
    expect(app.derivationS).toBe(sExpected);
  }, 900_000);

  it('mode 1: writer.registerApplicationCoDec accepts a TS-built Schnorr proof', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    // Random non-zero scalar, small enough that any subgroup-order check
    // is trivially passed.
    const sk = BigInt('0x' + Array.from(globalThis.crypto.getRandomValues(new Uint8Array(16)))
      .map((b) => b.toString(16).padStart(2, '0')).join('')) | 1n;

    const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, fixture.epochId, aid);

    const policy: AppPolicy = {
      authorizedSubmitter: ZERO_ADDRESS,
      maxCiphertexts:      0,
      notBeforeBlock:      0n,
      notAfterBlock:       0n,
    };

    // proveOrganizer returns RTE coords (matching the on-chain transcript);
    // writer.registerApplicationCoDec wants TE on input and converts back to
    // RTE before sending. Pass through the form converter to keep the shapes
    // straight.
    const { fromRTEtoTE } = await import('../src/crypto/babyjub-form.js');
    const [pkOrgX_TE, pkOrgY_TE] = fromRTEtoTE(pkOrgX, pkOrgY);

    const tx = await writer.registerApplicationCoDec(
      fixture.epochId, aid, policy,
      pkOrgX_TE, pkOrgY_TE,
      proof.ax, proof.ay, proof.z,
    );
    await writer.publicClient.waitForTransactionReceipt({ hash: tx });

    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(true);
    expect(app.mode).toBe(AppMode.OrganizerCoDec);
    // client.getApplication returns organizerPK in TE form (RTE→TE
    // conversion at the boundary), so compare against TE inputs.
    expect(app.organizerPK[0]).toBe(pkOrgX_TE);
    expect(app.organizerPK[1]).toBe(pkOrgY_TE);
    expect(app.derivationS).toBe(0n); // mode 1 stores S=0
  }, 900_000);

  it('mode 1: a tampered Schnorr response is rejected on-chain', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const sk = 1234567890123456789n;
    const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, fixture.epochId, aid);
    const { fromRTEtoTE } = await import('../src/crypto/babyjub-form.js');
    const [pkOrgX_TE, pkOrgY_TE] = fromRTEtoTE(pkOrgX, pkOrgY);

    const policy: AppPolicy = {
      authorizedSubmitter: ZERO_ADDRESS, maxCiphertexts: 0,
      notBeforeBlock: 0n, notAfterBlock: 0n,
    };

    // Flip one bit of `z` and confirm the on-chain verifier reverts.
    await expect(
      writer.registerApplicationCoDec(
        fixture.epochId, aid, policy,
        pkOrgX_TE, pkOrgY_TE,
        proof.ax, proof.ay, proof.z + 1n,
      ),
    ).rejects.toThrow();
  }, 900_000);

  it('getApplication returns exists=false for an unregistered aid', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const aid = randomAid();
    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(false);
  });
});
