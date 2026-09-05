import { type Address } from 'viem';
import { dkgManagerAbi } from './abi.js';
import { type DKGClient } from './client.js';
import { EpochPhase, type EpochPhaseValue, type PollOptions } from './types.js';
import { sleep } from './utils.js';
import { fromRTEtoTE } from './crypto/babyjub-form.js';

const DEFAULT_INTERVAL_MS = 2_000;
/** All-zero bytes32 — the "nothing stored" sentinel for a ciphertext hash. */
const ZERO_BYTES32 = ('0x' + '00'.repeat(32)) as `0x${string}`;
const DEFAULT_TIMEOUT_MS = 120_000;

/**
 * Poll until the given epoch reaches the target status (or beyond).
 *
 * @throws If the epoch is Aborted when waiting for a later status.
 * @throws If the timeout is exceeded.
 */
export async function waitForEpochPhase(
  client: DKGClient,
  epochId: `0x${string}`,
  targetStatus: EpochPhaseValue,
  options?: PollOptions,
): Promise<void> {
  const intervalMs = options?.intervalMs ?? DEFAULT_INTERVAL_MS;
  const timeoutMs = options?.timeoutMs ?? DEFAULT_TIMEOUT_MS;
  const deadline = Date.now() + timeoutMs;

  while (Date.now() < deadline) {
    const epoch = await client.getEpoch(epochId);
    if (epoch.status === EpochPhase.Aborted) {
      throw new Error(`Epoch ${epochId} was aborted`);
    }
    if (epoch.status >= targetStatus) return;
    await sleep(intervalMs);
  }
  throw new Error(
    `Timeout waiting for epoch ${epochId} to reach status ${targetStatus}`,
  );
}

/**
 * Poll until the combined decryption for a ciphertext is marked complete.
 *
 * @returns The completed CombinedDecryptionRecord.
 * @throws If the epoch is Aborted or the timeout is exceeded.
 */
export async function waitForDecryption(
  client: DKGClient,
  epochId: `0x${string}`,
  aid: `0x${string}`,
  ciphertextIndex: number,
  options?: PollOptions,
) {
  const intervalMs = options?.intervalMs ?? DEFAULT_INTERVAL_MS;
  const timeoutMs = options?.timeoutMs ?? DEFAULT_TIMEOUT_MS;
  const deadline = Date.now() + timeoutMs;

  while (Date.now() < deadline) {
    const record = await client.getCombinedDecryption(epochId, aid, ciphertextIndex);
    if (record.completed) return record;

    // Also check if the epoch was aborted so we fail fast.
    const epoch = await client.getEpoch(epochId);
    if (epoch.status === EpochPhase.Aborted) {
      throw new Error(`Epoch ${epochId} was aborted`);
    }

    await sleep(intervalMs);
  }
  throw new Error(
    `Timeout waiting for decryption of ciphertext ${ciphertextIndex} in epoch ${epochId}`,
  );
}

/**
 * Watch for new epochs created after `fromBlock`.
 * Calls `onRound` for each EpochCreated event.
 * Returns an unsubscribe function.
 */
export function watchNewEpochs(
  client: DKGClient,
  onEpoch: (epochId: `0x${string}`, organizer: Address) => void,
  fromBlock?: bigint,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: [
      {
        type: 'event',
        name: 'EpochCreated',
        inputs: [
          { name: 'epochId', type: 'bytes12', indexed: true },
          { name: 'organizer', type: 'address', indexed: true },
          { name: 'startBlock', type: 'uint64', indexed: false },
          { name: 'seedBlock', type: 'uint64', indexed: false },
          { name: 'lotteryThreshold', type: 'uint256', indexed: false },
        ],
      },
    ] as const,
    eventName: 'EpochCreated',
    fromBlock,
    onLogs: (logs) => {
      for (const log of logs) {
        const { epochId, organizer } = log.args as any;
        if (epochId && organizer) onEpoch(epochId as `0x${string}`, organizer as Address);
      }
    },
  });
}

/** @deprecated Use {@link watchNewEpochs}. Kept for SDK 0.1.x compatibility. */
export const watchNewRounds = watchNewEpochs;

/**
 * Watch for a epoch being finalized (proof-less `finalizeEpoch`, freezing the
 * accepted contributor set). Calls `onLive` once when the EpochLive event
 * fires, with the number of accepted contributions. The epoch's pool keys
 * are not activated yet at this point — see `waitForPoolKeyActivated`.
 * Returns an unsubscribe function.
 */
export function watchEpochLive(
  client: DKGClient,
  epochId: `0x${string}`,
  onLive: (contributionCount: number) => void,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: dkgManagerAbi,
    eventName: 'EpochLive',
    args: { epochId: epochId as any },
    onLogs: (logs) => {
      for (const log of logs) {
        const { contributionCount } = log.args as any;
        if (contributionCount !== undefined) onLive(Number(contributionCount));
      }
    },
  });
}

/**
 * Watch for the DecryptionCombined event for a specific ciphertext
 * (optionally narrowed to one application via `opts.aid`; the index is
 * per `(epochId, aid)`, so pass the aid whenever more than one application
 * is registered on the epoch). The callback receives the recovered
 * plaintext scalar as a bigint. Returns an unsubscribe function.
 */
export function watchDecryptionCombined(
  client: DKGClient,
  epochId: `0x${string}`,
  ciphertextIndex: number,
  onCombined: (combineHash: `0x${string}`, plaintext: bigint) => void,
  opts?: { aid?: `0x${string}` },
): () => void {
  const args: Record<string, unknown> = { epochId, ciphertextIndex };
  if (opts?.aid != null) args.aid = opts.aid;
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: dkgManagerAbi,
    eventName: 'DecryptionCombined',
    args: args as any,
    onLogs: (logs) => {
      for (const log of logs) {
        const { combineHash, plaintext } = log.args as any;
        if (combineHash && typeof plaintext === 'bigint')
          onCombined(combineHash as `0x${string}`, plaintext);
      }
    },
  });
}

/**
 * Watch for CiphertextSubmitted events on an epoch (optionally one
 * application via `opts.aid`). The callback receives the application id,
 * the on-chain-assigned ciphertext index, the submitter, the (C1, C2)
 * BabyJubJub coordinates — the contract stores only the keccak hash, so the
 * event log is the only way to recover the coordinates the committee and the
 * organizer need for decryption.
 *
 * `c1`/`c2` are converted from on-chain RTE form to TE form for consistency
 * with the rest of this SDK (see `crypto/babyjub-form.ts`), so they can be
 * handed straight to the committee's partial-decryption computation.
 *
 * Returns an unsubscribe function.
 */
export function watchCiphertextSubmitted(
  client: DKGClient,
  epochId: `0x${string}`,
  onCiphertext: (payload: {
    aid: `0x${string}`;
    ciphertextIndex: number;
    submitter: Address;
    c1: { x: bigint; y: bigint };
    c2: { x: bigint; y: bigint };
  }) => void,
  opts?: { aid?: `0x${string}` },
): () => void {
  const args: Record<string, unknown> = { epochId };
  if (opts?.aid != null) args.aid = opts.aid;
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: dkgManagerAbi,
    eventName: 'CiphertextSubmitted',
    args: args as any,
    onLogs: (logs) => {
      for (const log of logs) {
        const { aid, ciphertextIndex, submitter, c1x, c1y, c2x, c2y } = log.args as any;
        if (typeof ciphertextIndex === 'number' && submitter && aid) {
          const [c1xT, c1yT] = fromRTEtoTE(c1x as bigint, c1y as bigint);
          const [c2xT, c2yT] = fromRTEtoTE(c2x as bigint, c2y as bigint);
          onCiphertext({
            aid: aid as `0x${string}`,
            ciphertextIndex,
            submitter: submitter as Address,
            c1: { x: c1xT, y: c1yT },
            c2: { x: c2xT, y: c2yT },
          });
        }
      }
    },
  });
}

/**
 * One-shot snapshot of a ciphertext's decryption pipeline.
 *
 * A ciphertext is combinable once the committee has posted `threshold`
 * partial decryptions and the decryption window is open
 * (`requireDecryptionOpen`) — there is no separate organizer-share gate in
 * either mode: an `OrganizerLocked` application's combine proof consumes the
 * organizer secret directly (see `revealOrganizerSecret`), and an `Automatic`
 * one uses the identity secret. This snapshot doesn't count partials (no
 * cheap on-chain counter exists); use `getPartialDecryptionEvents` for that.
 */
export async function decryptionProgress(
  client: DKGClient,
  epochId: `0x${string}`,
  aid: `0x${string}`,
  ciphertextIndex: number,
): Promise<{
  ciphertext: boolean;
  combined: boolean;
  plaintext: bigint;
}> {
  const [ctHash, record] = await Promise.all([
    client.getCiphertextHash(epochId, aid, ciphertextIndex),
    client.getCombinedDecryption(epochId, aid, ciphertextIndex),
  ]);
  return {
    ciphertext: ctHash !== ZERO_BYTES32,
    combined: record.completed,
    plaintext: record.plaintext,
  };
}

/**
 * Poll until pool key `keyIndex` is activated for the epoch (bit `keyIndex`
 * set in `client.getPoolStatus(epochId).activated`). Registering an
 * application under that key (`registerApplication` → `claimPoolKey`)
 * reverts `PoolKeyNotActive` until then.
 *
 * @throws If the epoch is Aborted or the timeout is exceeded.
 */
export async function waitForPoolKeyActivated(
  client: DKGClient,
  epochId: `0x${string}`,
  keyIndex: number,
  options?: PollOptions,
): Promise<void> {
  const intervalMs = options?.intervalMs ?? DEFAULT_INTERVAL_MS;
  const timeoutMs = options?.timeoutMs ?? DEFAULT_TIMEOUT_MS;
  const deadline = Date.now() + timeoutMs;

  while (Date.now() < deadline) {
    const { activated } = await client.getPoolStatus(epochId);
    if ((activated & (1 << keyIndex)) !== 0) return;

    const epoch = await client.getEpoch(epochId);
    if (epoch.status === EpochPhase.Aborted) {
      throw new Error(`Epoch ${epochId} was aborted`);
    }
    await sleep(intervalMs);
  }
  throw new Error(
    `Timeout waiting for pool key ${keyIndex} to activate in epoch ${epochId}`,
  );
}

/**
 * Return a human-readable summary of the current DKG network state.
 */
export async function networkSummary(client: DKGClient): Promise<{
  blockNumber: bigint;
  totalNodes: bigint;
  activeNodes: bigint;
  epochNonce: bigint;
}> {
  const [blockNumber, totalNodes, activeNodes, epochNonce] = await Promise.all([
    client.blockNumber(),
    client.nodeCount(),
    client.activeCount(),
    client.epochNonce(),
  ]);
  return { blockNumber, totalNodes, activeNodes, epochNonce };
}
