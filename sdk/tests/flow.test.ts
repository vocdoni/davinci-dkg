// Full DKG flow test.
//
// This test relies on the sdk-test-fixture Go binary to create a finalized
// single-participant round (including ZK proof generation).  If the binary
// cannot be built or run (e.g. circuit artifacts are not compiled), the test
// is skipped gracefully.
//
// What is tested:
//   • DKGClient reads a Finalized round correctly
//   • getRoundFinalizedEvents returns the expected event
//   • waitForRoundStatus resolves immediately for an already-finalized round
//   • getContribution / getShareCommitmentHash return accepted records
//   • ElGamal encrypt/decrypt roundtrip using a synthetic key pair
//   • buildRoundId / parseRoundId roundtrip on the fixture round ID

import { describe, it, expect, beforeAll } from 'vitest';
import { inject } from 'vitest';
import { spawnSync } from 'node:child_process';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';
import {
  DKGClient,
  RoundStatus,
  buildElGamal,
  waitForRoundStatus,
  parseRoundId,
  buildRoundId,
} from '../src/index.js';
import { makePublicClient } from './helpers/accounts.js';

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

interface FixtureResult {
  roundId: `0x${string}`;
  collectivePublicKeyHash: `0x${string}`;
}

function runGoFixture(rpcUrl: string, addressesFile: string): FixtureResult | null {
  // __dirname = sdk/tests → go up to sdk/ then repo root
  const repoRoot    = path.resolve(__dirname, '..', '..');
  const fixtureMain = path.join(repoRoot, 'cmd', 'sdk-test-fixture');

  const result = spawnSync(
    'go',
    ['run', fixtureMain, '--rpc-url', rpcUrl, '--addresses-file', addressesFile],
    {
      cwd:      repoRoot,
      encoding: 'utf8',
      timeout:  600_000, // 10 min — ZK trusted setup can be slow
    },
  );

  if (result.status !== 0) {
    console.warn('[flow-test] Go fixture failed — skipping full flow tests.');
    console.warn('[flow-test] stderr:', result.stderr?.slice(0, 500));
    return null;
  }

  // gnark's circuit logger writes to stdout even when we redirect our own
  // logger to stderr. Find the last line that looks like a JSON object.
  const jsonLine = result.stdout
    .split('\n')
    .map((l) => l.trim())
    .filter((l) => l.startsWith('{'))
    .at(-1);

  if (!jsonLine) {
    console.warn('[flow-test] Could not find JSON in fixture output:', result.stdout.slice(0, 500));
    return null;
  }

  try {
    return JSON.parse(jsonLine) as FixtureResult;
  } catch {
    console.warn('[flow-test] Could not parse fixture JSON:', jsonLine);
    return null;
  }
}

describe('Full DKG flow (via Go fixture)', () => {
  let client:  DKGClient;
  let fixture: FixtureResult | null = null;

  beforeAll(async () => {
    const { enabled, rpcUrl, managerAddress, addressesFile } = useHarness();
    if (!enabled) return;

    client = new DKGClient({
      publicClient:    makePublicClient(rpcUrl),
      managerAddress,
    });

    console.log('[flow-test] Running Go fixture to create a finalized round…');
    fixture = runGoFixture(rpcUrl, addressesFile);

    if (fixture) {
      console.log(`[flow-test] Fixture round: ${fixture.roundId}`);
    }
  });

  it('fixture round is in Finalized status', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const round = await client.getRound(fixture.roundId);
    expect(round.status).toBe(RoundStatus.Finalized);
    expect(round.policy.threshold).toBe(1);
  });

  it('waitForRoundStatus resolves immediately for an already-finalized round', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    await waitForRoundStatus(client, fixture.roundId, RoundStatus.Finalized, {
      intervalMs: 500,
      timeoutMs:  10_000,
    });
  });

  it('getRoundFinalizedEvents returns the finalization event', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const events = await client.getRoundFinalizedEvents(fixture.roundId);
    expect(events.length).toBeGreaterThan(0);

    const ev = events[0];
    expect(ev.collectivePublicKeyHash).toBe(fixture.collectivePublicKeyHash);
    expect(ev.aggregateCommitmentsHash).toMatch(/^0x[0-9a-f]{64}$/i);
    expect(ev.shareCommitmentHash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('selectedParticipants returns one participant', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const participants = await client.selectedParticipants(fixture.roundId);
    expect(participants).toHaveLength(1);
    expect(participants[0]).toMatch(/^0x[0-9a-fA-F]{40}$/);
  });

  it('getContribution returns an accepted contribution record', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const participants = await client.selectedParticipants(fixture.roundId);
    const contrib      = await client.getContribution(fixture.roundId, participants[0]);
    expect(contrib.accepted).toBe(true);
    expect(contrib.commitmentsHash).toMatch(/^0x[0-9a-f]{64}$/i);
  });

  it('getShareCommitmentHash returns a non-zero hash for participant 1', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const h = await client.getShareCommitmentHash(fixture.roundId, 1);
    expect(h).not.toBe('0x' + '0'.repeat(64));
  });

  it('parseRoundId on fixture round ID roundtrips through buildRoundId', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    const { prefix, nonce } = parseRoundId(fixture.roundId);
    const rebuilt = buildRoundId(prefix, nonce);
    expect(rebuilt.toLowerCase()).toBe(fixture.roundId.toLowerCase());
  });

  it('ElGamal encrypt/decrypt with a synthetic key pair', async () => {
    const { enabled } = useHarness();
    if (!enabled || !fixture) return;

    // Test the cryptographic primitive with a locally-generated key pair.
    // In the full protocol, the collective public key is fetched via
    // client.getCollectivePublicKey(roundId).
    const eg = await buildElGamal();
    const { privKey, pubKey } = eg.generateKeyPair();

    const plaintext  = 137n;
    const ciphertext = eg.encrypt(plaintext, pubKey);
    const recovered  = eg.decrypt(ciphertext, privKey);

    expect(recovered).toBe(plaintext);
  });
});
