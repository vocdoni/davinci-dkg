// SDK combine end-to-end test.
//
// Validates that `writer.combineDecryption(...)` composes correctly with the on-chain
// `combineDecryption` ABI. The Go fixture builds the Groth16 proof + combine
// payload (since gnark proving in TS is not feasible), but the on-chain
// combine call itself goes through the SDK writer. Without this test, an
// ABI signature drift in writer.ts would only be caught by a downstream
// consumer hitting the chain.
//
// Flow:
//   1. Go fixture (`--action=create`) → finalized single-participant epoch.
//   2. SDK registers an application, encrypts a plaintext under PK_aid,
//      calls writer.submitCiphertext and takes the on-chain-assigned index,
//      then releases the organizer share.
//   3. Go fixture (`--action=prepare-combine`) → submits the partial
//      decryption on-chain and emits the combine payload bytes.
//   4. SDK calls writer.combineDecryption(epochId, aid, idx, …) with those
//      bytes.
//   5. SDK reads getPlaintext and asserts the recovered value matches.

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { spawn } from 'node:child_process';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import {
  DKGClient,
  DKGWriter,
  applicationKey,
  encrypt,
  randomOrganizerSecret,
  type AppPolicy,
  type BabyJubPoint,
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

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as const;

function randomAid(): `0x${string}` {
  const buf = new Uint8Array(32);
  globalThis.crypto.getRandomValues(buf);
  buf[0] &= 0x1f; // keep `aid` below the BN254 scalar field modulus
  return ('0x' + Array.from(buf).map((b) => b.toString(16).padStart(2, '0')).join('')) as `0x${string}`;
}

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

  it('SDK writer.combineDecryption lands a plaintext on-chain', async () => {
    const { enabled, rpcUrl, addressesFile } = useHarness();
    if (!enabled || !fixture) return;

    // 1. Register an application, encrypt under PK_aid, submit; the contract
    //    assigns the index. Then release the organizer share — combine
    //    reverts OrganizerShareMissing() without it.
    const aid = randomAid();
    const skOrg = randomOrganizerSecret();
    const policy: AppPolicy = {
      authorizedSubmitter: ZERO_ADDRESS,
      maxCiphertexts: 0,
      notBeforeBlock: 0n,
      notAfterBlock: 0n,
    };
    const regTx = await writer.registerApplication(fixture.epochId, aid, policy, skOrg);
    await writer.publicClient.waitForTransactionReceipt({ hash: regTx });

    const pk = await client.getCollectivePublicKey(fixture.epochId);
    const app = await client.getApplication(fixture.epochId, aid);
    const pkEp: BabyJubPoint = [pk.x, pk.y];
    const plaintext = 137n;
    const ciphertext = await encrypt(plaintext, applicationKey(pkEp, app.organizerPK));
    const { ciphertextIndex } = await writer.submitCiphertext(fixture.epochId, aid, ciphertext);
    expect(ciphertextIndex).toBeGreaterThanOrEqual(1);

    const shareTx = await writer.submitOrganizerShare(
      fixture.epochId, aid, ciphertextIndex, ciphertext, skOrg,
    );
    await writer.publicClient.waitForTransactionReceipt({ hash: shareTx });

    // 2. Go fixture builds the proof + submits the partial decryption,
    //    then hands back the combine bytes for the SDK to send.
    const prepOut = await runGoFixture([
      '--rpc-url', rpcUrl,
      '--addresses-file', addressesFile,
      '--action=prepare-combine',
      '--epoch-id', fixture.epochId,
      '--aid', aid,
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
      aid,
      ciphertextIndex,
      payload!.combineHash,
      BigInt(payload!.plaintext),
      payload!.transcript,
      payload!.proof,
      payload!.input,
    );
    await writer.publicClient.waitForTransactionReceipt({ hash: combineTx });

    const recovered = await client.getPlaintext(fixture.epochId, aid, ciphertextIndex);
    expect(recovered).toBe(plaintext);
  }, 900_000);
});
