// SDK end-to-end ciphertext test.
//
// Drives the full ElGamal epoch-trip against a live Anvil + finalized DKG
// epoch, exercising the BabyJubJub form-conversion plumbing that the
// monitor / writer / client expose for SDK consumers:
//
//   1. Go fixture (`sdk-test-fixture --action=create`) creates a finalized
//      single-participant epoch (committee=1, threshold=1, share=11).
//   2. SDK reads `getCollectivePublicKey(epochId)` → returned in TE form.
//   3. SDK encrypts a small plaintext with `encryptWithProof()`, which also
//      produces the Schnorr proof of knowledge of the ElGamal randomness.
//   4. SDK calls `writer.submitCiphertext(...)` — internally converts c1/c2
//      from TE → RTE so the contract's `_isOnBabyJubJub` check accepts them,
//      and returns the on-chain-assigned ciphertext index.
//   5. Go fixture (`sdk-test-fixture --action=decrypt --share=11 ...`) drives
//      partial decryption + combine on-chain.
//   6. SDK reads `getPlaintext(epochId, idx)` and asserts the recovered
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
  encryptWithProof,
  verifyCiphertextPoK,
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
  collectivePublicKeyHash: `0x${string}`;
  share: string; // decimal
}

interface FixtureDecryptResult {
  ok: true;
}

async function runGoFixture(args: string[]): Promise<{ status: number | null; stdout: string; stderr: string } | null> {
  const repoRoot    = path.resolve(__dirname, '..', '..');
  const fixtureMain = path.join(repoRoot, 'cmd', 'sdk-test-fixture');
  return new Promise<{ status: number | null; stdout: string; stderr: string } | null>((resolve, reject) => {
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

    console.log('[ciphertext-e2e] Running Go fixture (create) to set up a finalized epoch…');
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
    console.log(`[ciphertext-e2e] Fixture epoch: ${fixture.epochId}, share=${fixture.share}`);
  });

  it('SDK encrypt → submitCiphertext → combine → getPlaintext recovers plaintext', async () => {
    const { enabled, rpcUrl, addressesFile } = useHarness();
    if (!enabled || !fixture) return;

    // 1. Read the on-chain collective public key (returned in TE form thanks
    //    to the SDK's RTE→TE conversion in client.getCollectivePublicKey).
    const pk = await client.getCollectivePublicKey(fixture.epochId);
    expect(pk.x).not.toBe(0n);
    // y == 1 with x == 0 would be the identity, i.e. no contributions accepted yet.
    expect(!(pk.x === 0n && pk.y === 1n)).toBe(true);

    // 2. Encrypt a small plaintext using SDK ElGamal (operates in TE form) and
    //    prove knowledge of the randomness for (epochId, aid).
    const plaintext = 42n;
    const zeroAid = ('0x' + '0'.repeat(64)) as `0x${string}`;
    const { ciphertext, pok } = await encryptWithProof(fixture.epochId, zeroAid, plaintext, [pk.x, pk.y]);

    // 3. Submit to chain. The writer converts TE→RTE internally before sending,
    //    so the contract's `_isOnBabyJubJub` (RTE) check passes, and hands back
    //    the index the contract assigned.
    const countBefore = await client.ciphertextCount(fixture.epochId, zeroAid);
    const { hash: submitTx, ciphertextIndex } = await writer.submitCiphertext(
      fixture.epochId, zeroAid, ciphertext, pok,
    );
    expect(submitTx).toMatch(/^0x[0-9a-f]{64}$/i);
    expect(ciphertextIndex).toBe(countBefore + 1);
    expect(await client.ciphertextCount(fixture.epochId, zeroAid)).toBe(ciphertextIndex);

    // Sanity: the contract now stores a non-zero ciphertext hash for this index,
    // and the event carries a proof the committee will accept.
    const ctHash = await client.getCiphertextHash(fixture.epochId, zeroAid, ciphertextIndex);
    expect(ctHash).not.toBe('0x' + '0'.repeat(64));
    const events = await client.getCiphertextSubmittedEvents(fixture.epochId, { aid: zeroAid, ciphertextIndex });
    expect(events).toHaveLength(1);
    expect(events[0].pokValid).toBe(true);
    expect(verifyCiphertextPoK(
      fixture.epochId, zeroAid, events[0].c1.x, events[0].c1.y, events[0].c2.x, events[0].c2.y, events[0].pok,
    )).toBe(true);

    // 4. Drive the on-chain decryption flow via the Go fixture (it builds the
    //    Groth16 proofs we can't generate in TS).
    console.log('[ciphertext-e2e] Running Go fixture (decrypt) to drive partial decrypt + combine…');
    const decryptOut = await runGoFixture([
      '--rpc-url', rpcUrl,
      '--addresses-file', addressesFile,
      '--action=decrypt',
      '--epoch-id', fixture.epochId,
      '--ciphertext-index', String(ciphertextIndex),
      '--share', fixture.share,
    ]);
    if (!decryptOut || decryptOut.status !== 0) {
      throw new Error(`fixture decrypt failed: ${decryptOut?.stderr.slice(0, 1000) ?? 'no output'}`);
    }
    const decryptParsed = lastJsonLine<FixtureDecryptResult>(decryptOut.stdout);
    expect(decryptParsed?.ok).toBe(true);

    // 5. Read the recovered plaintext from chain — must match what we sent.
    const recovered = await client.getPlaintext(fixture.epochId, zeroAid, ciphertextIndex);
    expect(recovered).toBe(plaintext);
  }, 900_000);
});
