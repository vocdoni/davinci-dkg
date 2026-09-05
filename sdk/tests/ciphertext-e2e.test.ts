// SDK end-to-end ciphertext test.
//
// Drives the full ElGamal epoch-trip against a live Anvil + Live DKG epoch,
// exercising the pool-key and BabyJubJub form-conversion plumbing that the
// monitor / writer / client expose for SDK consumers:
//
//   1. Go fixture (`sdk-test-fixture --action=create --keys=2`) creates a Live
//      single-participant epoch (committee=1, threshold=1) — the proof-carrying
//      finalizeEpoch stores all MAX_K pool keys at once — and reports
//      participant 1's share of keys 0 and 1. Every key is dealt from its own
//      polynomial, so the shares differ per key.
//   2. SDK reads `getPoolStatus` / `getPoolKey(epochId, j)` → TE form, and
//      checks them against the fixture's P_0 and the finalization calldata.
//   3. SDK registers an application; it claims the next unclaimed key.
//      Organizer-locked: `PK_aid = P_j + PK_org`; automatic: `PK_aid = P_j`.
//      `client.getApplicationKey` computes it.
//   4. SDK calls `writer.submitCiphertext(...)` — internally converts c1/c2
//      from TE → RTE so the contract's `_isOnBabyJubJub` check accepts them,
//      and returns the on-chain-assigned ciphertext index.
//   5. For the locked application the organizer publishes `sk_org` once with
//      `writer.revealOrganizerSecret` (there is no per-ciphertext share).
//      Until then the contract refuses every partial and combine of the
//      application (`OrganizerSecretNotRevealed`), so the reveal has to land
//      before the fixture is asked to decrypt.
//   6. Go fixture (`sdk-test-fixture --action=decrypt --share=<shares[j]> ...`)
//      drives partial decryption + combine on-chain with the share of the key
//      `j` the application claimed: the contract checks the partial's share
//      commitment against that key's share root.
//   7. SDK reads `getPlaintext(epochId, aid, idx)` and asserts the recovered
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
  AppMode,
  DKGClient,
  DKGWriter,
  applicationKey,
  encrypt,
  randomOrganizerSecret,
  type AppPolicyInput,
  type BabyJubPoint,
  randomAid,
} from '../src/index.js';
import { fromRTEtoTE } from '../src/crypto/babyjub-form.js';
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
  /** Participant 1's share of pool key 0 (= shares[0]), decimal. */
  share: string;
  /** Participant 1's share of pool key j, decimal, one entry per key the fixture was asked for. */
  shares: string[];
  /** P_0 in the contract's RTE form, decimal coordinates. */
  poolKey: { x: string; y: string };
  /** `shares.length`: the keys the output describes; the pool itself is always whole. */
  activatedKeys: number;
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

/** Drive partial decryption + combine on-chain via the Go fixture. */
async function fixtureDecrypt(
  epochId: `0x${string}`,
  aid: `0x${string}`,
  ciphertextIndex: number,
  share: string,
  skOrg?: bigint,
): Promise<void> {
  const { rpcUrl, addressesFile } = useHarness();
  console.log('[ciphertext-e2e] Running Go fixture (decrypt) to drive partial decrypt + combine…');
  const args = [
    '--rpc-url', rpcUrl,
    '--addresses-file', addressesFile,
    '--action=decrypt',
    '--epoch-id', epochId,
    '--aid', aid,
    '--ciphertext-index', String(ciphertextIndex),
    '--share', share,
  ];
  if (skOrg !== undefined) args.push('--org-secret', '0x' + skOrg.toString(16));
  const out = await runGoFixture(args);
  if (!out || out.status !== 0) {
    throw new Error(`fixture decrypt failed: ${out?.stderr.slice(0, 1000) ?? 'no output'}`);
  }
  expect(lastJsonLine<FixtureDecryptResult>(out.stdout)?.ok).toBe(true);
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

    console.log('[ciphertext-e2e] Running Go fixture (create) to set up a Live epoch…');
    // Shares of two keys: one per application registered below.
    const createOut = await runGoFixture([
      '--rpc-url', rpcUrl, '--addresses-file', addressesFile, '--action=create', '--keys=2',
    ]);
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
    console.log(
      `[ciphertext-e2e] Fixture epoch: ${fixture.epochId}, shares=${fixture.shares.join(',')}, keys=${fixture.activatedKeys}`,
    );
  });

  it('the pool is whole at Live and getPoolKey returns P_0 in TE form', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    expect(fixture.activatedKeys).toBe(2);
    // One share per reported key; `share` is the key-0 one.
    expect(fixture.shares).toHaveLength(2);
    expect(fixture.shares[0]).toBe(fixture.share);
    expect(fixture.shares[1]).not.toBe(fixture.shares[0]);
    // Nothing claimed yet; every key already exists.
    const status = await client.getPoolStatus(fixture.epochId);
    expect(status.nextIndex).toBe(0);

    // The fixture reports P_0 in the contract's RTE form; the client converts
    // to TE at the boundary.
    const expected = fromRTEtoTE(BigInt(fixture.poolKey.x), BigInt(fixture.poolKey.y));
    const p0 = await client.getPoolKey(fixture.epochId, 0);
    expect(p0).toEqual(expected);
    // y == 1 with x == 0 would be the identity, i.e. no contributions accepted.
    expect(!(p0[0] === 0n && p0[1] === 1n)).toBe(true);
    const p1 = await client.getPoolKey(fixture.epochId, 1);
    expect(p1).not.toEqual(p0);

    // The finalization calldata carries the raw on-chain words of every key.
    const finalize = await client.getFinalizeTranscript(fixture.epochId);
    expect(finalize).not.toBeNull();
    expect(finalize!.transcript.poolKeys[0].x).toBe(BigInt(fixture.poolKey.x));
    expect(finalize!.transcript.poolKeys[0].y).toBe(BigInt(fixture.poolKey.y));
  });

  it('organizer-locked: register → encrypt → submit → reveal sk_org → combine → getPlaintext', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    // 1. Register an organizer-locked application. sk_org never leaves this
    //    process; only PK_org and the proof of possession go on chain. It
    //    claims the next unclaimed key: 0 on a fresh epoch.
    const aid = randomAid();
    const skOrg = randomOrganizerSecret();
    // Organizer-locked, registrant-only: the defaults of every field.
    const policy: AppPolicyInput = { maxCiphertexts: 0 };
    const regTx = await writer.registerApplication(fixture.epochId, aid, policy, skOrg);
    await writer.publicClient.waitForTransactionReceipt({ hash: regTx });

    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.exists).toBe(true);
    expect(app.policy.mode).toBe(AppMode.OrganizerLocked);
    expect(app.poolIndex).toBe(0);
    expect(app.organizerSecret).toBe(0n);
    expect(await client.getAppPoolIndex(fixture.epochId, aid)).toBe(0);
    expect((await client.getPoolStatus(fixture.epochId)).nextIndex).toBe(1);

    // 2. PK_aid = P_0 + PK_org (both TE at this boundary), as computed by
    //    the client and by hand.
    const poolKey = await client.getPoolKey(fixture.epochId, app.poolIndex);
    const pkAid = await client.getApplicationKey(fixture.epochId, aid);
    expect(pkAid).toEqual(applicationKey(poolKey, app.organizerPK));
    expect(pkAid).not.toEqual(poolKey);

    // 3. Encrypt and submit. The writer converts TE→RTE internally before
    //    sending, so the contract's `_isOnBabyJubJub` (RTE) check passes, and
    //    hands back the index the contract assigned.
    const plaintext = 42n;
    const ciphertext = await encrypt(plaintext, pkAid as BabyJubPoint);
    const countBefore = await client.ciphertextCount(fixture.epochId, aid);
    const { hash: submitTx, ciphertextIndex } = await writer.submitCiphertext(
      fixture.epochId, aid, ciphertext,
    );
    expect(submitTx).toMatch(/^0x[0-9a-f]{64}$/i);
    expect(ciphertextIndex).toBe(countBefore + 1);
    expect(await client.ciphertextCount(fixture.epochId, aid)).toBe(ciphertextIndex);

    // Sanity: the contract now stores a non-zero ciphertext hash for this index
    // and the event replays the coordinates we sent.
    const ctHash = await client.getCiphertextHash(fixture.epochId, aid, ciphertextIndex);
    expect(ctHash).not.toBe('0x' + '0'.repeat(64));
    const events = await client.getCiphertextSubmittedEvents(fixture.epochId, { aid, ciphertextIndex });
    expect(events).toHaveLength(1);

    // 4. Reveal the organizer secret once. The contract checks sk·G == PK_org
    //    and from then on the committee combines by itself. Before it,
    //    requireDecryptionOpen reverts OrganizerSecretNotRevealed, which the
    //    client reports as "not open".
    expect(await client.getOrganizerSecretRevealedEvents(fixture.epochId, aid)).toHaveLength(0);
    expect(await client.isDecryptionOpen(fixture.epochId, aid)).toBe(false);
    const revealTx = await writer.revealOrganizerSecret(fixture.epochId, aid, skOrg);
    await writer.publicClient.waitForTransactionReceipt({ hash: revealTx });
    expect((await client.getApplication(fixture.epochId, aid)).organizerSecret).toBe(skOrg);
    const revealed = await client.getOrganizerSecretRevealedEvents(fixture.epochId, aid);
    expect(revealed).toHaveLength(1);
    expect(revealed[0].organizerSecret).toBe(skOrg);
    expect(await client.isDecryptionOpen(fixture.epochId, aid)).toBe(true);

    // 5. Drive the on-chain decryption flow via the Go fixture (it builds the
    //    Groth16 proofs we can't generate in TS) with the member's share of
    //    the key this application claimed.
    await fixtureDecrypt(fixture.epochId, aid, ciphertextIndex, fixture.shares[app.poolIndex], skOrg);

    // 6. Read the recovered plaintext from chain — must match what we sent.
    const recovered = await client.getPlaintext(fixture.epochId, aid, ciphertextIndex);
    expect(recovered).toBe(plaintext);
  }, 900_000);

  it('automatic: register → encryptAndSubmit under P_1 → combine → getPlaintext', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    // No organizer key at all: the committee threshold alone opens the
    // ciphertext, and the application key is the bare pool key.
    const aid = randomAid();
    const regTx = await writer.registerApplication(fixture.epochId, aid, { mode: AppMode.Automatic });
    await writer.publicClient.waitForTransactionReceipt({ hash: regTx });

    const app = await client.getApplication(fixture.epochId, aid);
    expect(app.policy.mode).toBe(AppMode.Automatic);
    expect(app.poolIndex).toBe(1);
    expect(app.organizerPK).toEqual([0n, 1n]);
    const poolKey = await client.getPoolKey(fixture.epochId, 1);
    expect(await client.getApplicationKey(fixture.epochId, aid)).toEqual(poolKey);

    const plaintext = 7n;
    const { ciphertextIndex, ciphertext } = await writer.encryptAndSubmit(fixture.epochId, aid, plaintext);
    expect(ciphertextIndex).toBe(1);
    expect(ciphertext.c1[0]).not.toBe(0n);

    // Nothing to reveal for an automatic application.
    await expect(writer.revealOrganizerSecret(fixture.epochId, aid, 1n)).rejects.toThrow();

    // The application sits on key 1, whose share differs from key 0's: the
    // partial's share commitment must match P_1's share root on chain.
    await fixtureDecrypt(fixture.epochId, aid, ciphertextIndex, fixture.shares[app.poolIndex]);
    expect(await client.getPlaintext(fixture.epochId, aid, ciphertextIndex)).toBe(plaintext);
  }, 900_000);
});
