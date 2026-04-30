// SDK combine end-to-end test.
//
// Validates that `writer.combineDecryption(...)` — including the new `aid`
// parameter wired in P9 — actually composes correctly with the on-chain
// `combineDecryption` ABI. The Go fixture builds the Groth16 proof + combine
// payload (since gnark proving in TS is not feasible), but the on-chain
// combine call itself goes through the SDK writer. Without this test, an
// ABI signature drift in writer.ts would only be caught by a downstream
// consumer hitting the chain.
//
// Flow:
//   1. Go fixture (`--action=create`) → finalized single-participant epoch.
//   2. SDK encrypts a plaintext, calls writer.submitCiphertext.
//   3. Go fixture (`--action=prepare-combine`) → submits the partial
//      decryption on-chain and emits the combine payload bytes.
//   4. SDK calls writer.combineDecryption(epochId, 0x00…00, idx, …) with
//      those bytes. The 32-byte zero `aid` is the legacy per-epoch path
//      (no application registered).
//   5. SDK reads getPlaintext and asserts the recovered value matches.

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { spawn } from 'node:child_process';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import { DKGClient, DKGWriter, buildElGamal } from '../src/index.js';
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
  collectivePublicKeyHash: `0x${string}`;
  share: string;
}

interface PrepareCombineResult {
  combineHash: `0x${string}`;
  plaintext:   string;            // decimal
  transcript:  `0x${string}`;
  proof:       `0x${string}`;
  input:       `0x${string}`;
}

async function runGoFixture(args: string[]) {
  const repoRoot    = path.resolve(__dirname, '..', '..');
  const fixtureMain = path.join(repoRoot, 'cmd', 'sdk-test-fixture');
  return new Promise<{ status: number | null; stdout: string; stderr: string }>((resolve, reject) => {
    const child = spawn('go', ['run', fixtureMain, ...args], { cwd: repoRoot });
    let stdout = '';
    let stderr = '';
    child.stdout.on('data', (d) => { stdout += d.toString('utf8'); });
    child.stderr.on('data', (d) => { stderr += d.toString('utf8'); });
    const timer = setTimeout(() => { child.kill('SIGKILL'); reject(new Error('Go fixture timed out')); }, 600_000);
    child.on('error', (err) => { clearTimeout(timer); reject(err); });
    child.on('close', (code) => { clearTimeout(timer); resolve({ status: code, stdout, stderr }); });
  });
}

function lastJsonLine<T>(stdout: string): T | null {
  const line = stdout.split('\n').map((l) => l.trim()).filter((l) => l.startsWith('{')).at(-1);
  if (!line) return null;
  try { return JSON.parse(line) as T; } catch { return null; }
}

const ZERO_AID = ('0x' + '00'.repeat(32)) as `0x${string}`;

describe('SDK combineDecryption end-to-end (writer drives the on-chain combine)', () => {
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
      console.warn('[combine-e2e] fixture create failed — skipping. stderr:', out.stderr.slice(0, 500));
      return;
    }
    fixture = lastJsonLine<FixtureCreateResult>(out.stdout);
  });

  it('SDK writer.combineDecryption with aid=0x00..00 lands a plaintext on-chain', async () => {
    const { enabled, rpcUrl, addressesFile } = useHarness();
    if (!enabled || !fixture) return;

    // 1. SDK encrypts + submits the ciphertext.
    const pk = await client.getCollectivePublicKey(fixture.epochId);
    const plaintext = 137n;
    const eg = await buildElGamal();
    const ct = eg.encrypt(plaintext, [pk.x, pk.y]);
    const ciphertextIndex = 1;
    const submitTx = await writer.submitCiphertext(
      fixture.epochId, ciphertextIndex, ct.c1[0], ct.c1[1], ct.c2[0], ct.c2[1],
    );
    await writer.publicClient.waitForTransactionReceipt({ hash: submitTx });

    // 2. Go fixture builds the proof + submits the partial decryption,
    //    then hands back the combine bytes for the SDK to send.
    const prepOut = await runGoFixture([
      '--rpc-url', rpcUrl,
      '--addresses-file', addressesFile,
      '--action=prepare-combine',
      '--epoch-id', fixture.epochId,
      '--ciphertext-index', String(ciphertextIndex),
      '--share', fixture.share,
    ]);
    if (prepOut.status !== 0) {
      throw new Error(`prepare-combine failed: ${prepOut.stderr.slice(0, 1000)}`);
    }
    const payload = lastJsonLine<PrepareCombineResult>(prepOut.stdout);
    expect(payload).not.toBeNull();
    expect(payload!.plaintext).toBe(plaintext.toString());

    // 3. SDK writer drives the on-chain combine. This is the load-bearing
    //    assertion: any drift in the ABI shape (`aid` arg, ordering, types)
    //    would surface as a simulateContract revert here.
    const combineTx = await writer.combineDecryption(
      fixture.epochId,
      ZERO_AID,
      ciphertextIndex,
      payload!.combineHash,
      BigInt(payload!.plaintext),
      payload!.transcript,
      payload!.proof,
      payload!.input,
    );
    await writer.publicClient.waitForTransactionReceipt({ hash: combineTx });

    const recovered = await client.getPlaintext(fixture.epochId, ciphertextIndex);
    expect(recovered).toBe(plaintext);
  }, 900_000);
});
