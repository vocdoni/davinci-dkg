// SDK end-to-end ciphertext test.
//
// Drives the full ElGamal round-trip against a live Anvil + finalized DKG
// round, exercising the BabyJubJub form-conversion plumbing that the
// monitor / writer / client expose for SDK consumers:
//
//   1. Go fixture (`sdk-test-fixture --action=create`) creates a finalized
//      single-participant round (committee=1, threshold=1, share=11).
//   2. SDK reads `getCollectivePublicKey(roundId)` → returned in TE form.
//   3. SDK encrypts a small plaintext with `buildElGamal().encrypt()`.
//   4. SDK calls `writer.submitCiphertext(...)` — internally converts c1/c2
//      from TE → RTE so the contract's `_isOnBabyJubJub` check accepts them.
//   5. Go fixture (`sdk-test-fixture --action=decrypt --share=11 ...`) drives
//      partial decryption + combine on-chain.
//   6. SDK reads `getPlaintext(roundId, idx)` and asserts the recovered
//      value equals the original plaintext.
//
// This test is what would have caught the InvalidCiphertext() production bug:
// without the BJJ form conversion (or with the broken on-chain accumulator
// pre-fix), step 4 reverts with InvalidCiphertext().
//
// If the Go fixture is unavailable or the chain isn't ready, the test is
// skipped gracefully — same pattern as flow.test.ts.

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { spawn } from 'node:child_process';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import {
  DKGClient,
  DKGWriter,
  buildElGamal,
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
  roundId: `0x${string}`;
  collectivePublicKeyHash: `0x${string}`;
  share: string; // decimal
}

interface FixtureDecryptResult {
  ok: true;
}

async function runGoFixture(args: string[]): Promise<{ status: number | null; stdout: string; stderr: string } | null> {
  const repoRoot    = path.resolve(__dirname, '..', '..');
  const fixtureMain = path.join(repoRoot, 'cmd', 'sdk-test-fixture');
  return new Promise((resolve, reject) => {
    const child = spawn('go', ['run', fixtureMain, ...args], { cwd: repoRoot });
    let stdout = '';
    let stderr = '';
    child.stdout.on('data', (d) => { stdout += d.toString('utf8'); });
    child.stderr.on('data', (d) => { stderr += d.toString('utf8'); });
    const killTimer = setTimeout(() => { child.kill('SIGKILL'); reject(new Error('Go fixture timed out after 10 min')); }, 600_000);
    child.on('error', (err) => { clearTimeout(killTimer); reject(err); });
    child.on('close', (code) => { clearTimeout(killTimer); resolve({ status: code, stdout, stderr }); });
  }).catch((err) => {
    console.warn('[ciphertext-e2e] Go fixture error:', err);
    return null;
  });
}

function lastJsonLine<T>(stdout: string): T | null {
  const line = stdout
    .split('\n')
    .map((l) => l.trim())
    .filter((l) => l.startsWith('{'))
    .at(-1);
  if (!line) return null;
  try { return JSON.parse(line) as T; } catch { return null; }
}

describe('SDK ciphertext end-to-end (encrypt → submit → combine → getPlaintext)', () => {
  let client:  DKGClient;
  let writer:  DKGWriter;
  let fixture: FixtureCreateResult | null = null;

  beforeAll(async () => {
    const { enabled, rpcUrl, managerAddress, addressesFile } = useHarness();
    if (!enabled) return;

    client = new DKGClient({ publicClient: makePublicClient(rpcUrl), managerAddress });
    writer = new DKGWriter({
      publicClient: makePublicClient(rpcUrl),
      // Use account #1 so we don't clash with the fixture's tx nonce on account #0.
      walletClient: makeWalletClient(rpcUrl, 1),
      managerAddress,
    });

    console.log('[ciphertext-e2e] Running Go fixture (create) to set up a finalized round…');
    const createOut = await runGoFixture(['--rpc-url', rpcUrl, '--addresses-file', addressesFile, '--action=create']);
    if (!createOut || createOut.status !== 0) {
      console.warn('[ciphertext-e2e] fixture create failed — skipping. stderr:', createOut?.stderr.slice(0, 500));
      return;
    }
    const parsed = lastJsonLine<FixtureCreateResult>(createOut.stdout);
    if (!parsed) {
      console.warn('[ciphertext-e2e] could not parse fixture create stdout');
      return;
    }
    fixture = parsed;
    console.log(`[ciphertext-e2e] Fixture round: ${fixture.roundId}, share=${fixture.share}`);
  });

  it('SDK encrypt → submitCiphertext → combine → getPlaintext recovers plaintext', async () => {
    const { enabled, rpcUrl, addressesFile } = useHarness();
    if (!enabled || !fixture) return;

    // 1. Read the on-chain collective public key (returned in TE form thanks
    //    to the SDK's RTE→TE conversion in client.getCollectivePublicKey).
    const pk = await client.getCollectivePublicKey(fixture.roundId);
    expect(pk.x).not.toBe(0n);
    // y == 1 with x == 0 would be the identity, i.e. no contributions accepted yet.
    expect(!(pk.x === 0n && pk.y === 1n)).toBe(true);

    // 2. Encrypt a small plaintext using SDK ElGamal (operates in TE form).
    const plaintext = 42n;
    const eg = await buildElGamal();
    const ciphertext = eg.encrypt(plaintext, [pk.x, pk.y]);

    // 3. Submit to chain. The writer converts TE→RTE internally before sending,
    //    so the contract's `_isOnBabyJubJub` (RTE) check passes.
    const ciphertextIndex = 1;
    const submitTx = await writer.submitCiphertext(
      fixture.roundId,
      ciphertextIndex,
      ciphertext.c1[0], ciphertext.c1[1],
      ciphertext.c2[0], ciphertext.c2[1],
    );
    expect(submitTx).toMatch(/^0x[0-9a-f]{64}$/i);
    await writer.publicClient.waitForTransactionReceipt({ hash: submitTx });

    // Sanity: the contract now stores a non-zero ciphertext hash for this index.
    const ctHash = await client.getCiphertextHash(fixture.roundId, ciphertextIndex);
    expect(ctHash).not.toBe('0x' + '0'.repeat(64));

    // 4. Drive the on-chain decryption flow via the Go fixture (it builds the
    //    Groth16 proofs we can't generate in TS).
    console.log('[ciphertext-e2e] Running Go fixture (decrypt) to drive partial decrypt + combine…');
    const decryptOut = await runGoFixture([
      '--rpc-url', rpcUrl,
      '--addresses-file', addressesFile,
      '--action=decrypt',
      '--round-id', fixture.roundId,
      '--ciphertext-index', String(ciphertextIndex),
      '--share', fixture.share,
    ]);
    if (!decryptOut || decryptOut.status !== 0) {
      throw new Error(`fixture decrypt failed: ${decryptOut?.stderr.slice(0, 1000) ?? 'no output'}`);
    }
    const decryptParsed = lastJsonLine<FixtureDecryptResult>(decryptOut.stdout);
    expect(decryptParsed?.ok).toBe(true);

    // 5. Read the recovered plaintext from chain — must match what we sent.
    const recovered = await client.getPlaintext(fixture.roundId, ciphertextIndex);
    expect(recovered).toBe(plaintext);
  }, 900_000);
});
