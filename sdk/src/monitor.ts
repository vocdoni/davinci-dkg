import { type Address } from 'viem';
import { dkgManagerAbi } from './abi.js';
import { type DKGClient } from './client.js';
import { EpochPhase, type CiphertextPoK, type EpochPhaseValue, type PollOptions } from './types.js';
import { sleep } from './utils.js';
import { fromRTEtoTE } from './crypto/babyjub-form.js';
import { verifyCiphertextPoK } from './schnorr.js';

const DEFAULT_INTERVAL_MS = 2_000;
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
 * Watch for a epoch being finalized.
 * Calls `onFinalized` once when the EpochLive event fires.
 * Returns an unsubscribe function.
 */
export function watchEpochLive(
  client: DKGClient,
  epochId: `0x${string}`,
  onFinalized: (collectivePublicKeyHash: `0x${string}`) => void,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: [
      {
        type: 'event',
        name: 'EpochLive',
        inputs: [
          { name: 'epochId', type: 'bytes12', indexed: true },
          { name: 'aggregateCommitmentsHash', type: 'bytes32', indexed: false },
          { name: 'collectivePublicKeyHash', type: 'bytes32', indexed: false },
          { name: 'shareCommitmentHash', type: 'bytes32', indexed: false },
        ],
      },
    ] as const,
    eventName: 'EpochLive',
    args: { epochId: epochId as any },
    onLogs: (logs) => {
      for (const log of logs) {
        const { collectivePublicKeyHash } = log.args as any;
        if (collectivePublicKeyHash) onFinalized(collectivePublicKeyHash as `0x${string}`);
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
 * BabyJubJub coordinates and the submitter's proof of knowledge — the
 * contract stores only the keccak hash, so the event log is the only way to
 * recover the coordinates nodes need for threshold decryption.
 *
 * `c1`/`c2` are converted from on-chain RTE form to TE form for consistency
 * with the rest of this SDK (see `crypto/babyjub-form.ts`); `pok` stays in
 * RTE form as reported on-chain. `pokValid` is the SDK's verdict on the
 * proof (`verifyCiphertextPoK`, computed over the RTE words): committee
 * nodes run the same check and never decrypt a ciphertext that fails it.
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
    pok: CiphertextPoK;
    pokValid: boolean;
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
        const { aid, ciphertextIndex, submitter, c1x, c1y, c2x, c2y, pokAx, pokAy, pokZ } =
          log.args as any;
        if (typeof ciphertextIndex === 'number' && submitter && aid) {
          const pok: CiphertextPoK = { ax: pokAx as bigint, ay: pokAy as bigint, z: pokZ as bigint };
          const pokValid = verifyCiphertextPoK(
            epochId, aid as `0x${string}`, c1x as bigint, c1y as bigint, c2x as bigint, c2y as bigint, pok,
          );
          const [c1xT, c1yT] = fromRTEtoTE(c1x as bigint, c1y as bigint);
          const [c2xT, c2yT] = fromRTEtoTE(c2x as bigint, c2y as bigint);
          onCiphertext({
            aid: aid as `0x${string}`,
            ciphertextIndex,
            submitter: submitter as Address,
            c1: { x: c1xT, y: c1yT },
            c2: { x: c2xT, y: c2yT },
            pok,
            pokValid,
          });
        }
      }
    },
  });
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
