import { type Address } from 'viem';
import { type DKGClient } from './client.js';
import { RoundStatus, type RoundStatusValue, type PollOptions } from './types.js';
import { sleep } from './utils.js';

const DEFAULT_INTERVAL_MS = 2_000;
const DEFAULT_TIMEOUT_MS = 120_000;

/**
 * Poll until the given round reaches the target status (or beyond).
 *
 * @throws If the round is Aborted when waiting for a later status.
 * @throws If the timeout is exceeded.
 */
export async function waitForRoundStatus(
  client: DKGClient,
  roundId: `0x${string}`,
  targetStatus: RoundStatusValue,
  options?: PollOptions,
): Promise<void> {
  const intervalMs = options?.intervalMs ?? DEFAULT_INTERVAL_MS;
  const timeoutMs = options?.timeoutMs ?? DEFAULT_TIMEOUT_MS;
  const deadline = Date.now() + timeoutMs;

  while (Date.now() < deadline) {
    const round = await client.getRound(roundId);
    if (round.status === RoundStatus.Aborted) {
      throw new Error(`Round ${roundId} was aborted`);
    }
    if (round.status >= targetStatus) return;
    await sleep(intervalMs);
  }
  throw new Error(
    `Timeout waiting for round ${roundId} to reach status ${targetStatus}`,
  );
}

/**
 * Poll until the combined decryption for a ciphertext is marked complete.
 *
 * @returns The completed CombinedDecryptionRecord.
 * @throws If the round is Aborted or the timeout is exceeded.
 */
export async function waitForDecryption(
  client: DKGClient,
  roundId: `0x${string}`,
  ciphertextIndex: number,
  options?: PollOptions,
) {
  const intervalMs = options?.intervalMs ?? DEFAULT_INTERVAL_MS;
  const timeoutMs = options?.timeoutMs ?? DEFAULT_TIMEOUT_MS;
  const deadline = Date.now() + timeoutMs;

  while (Date.now() < deadline) {
    const record = await client.getCombinedDecryption(roundId, ciphertextIndex);
    if (record.completed) return record;

    // Also check if the round was aborted so we fail fast.
    const round = await client.getRound(roundId);
    if (round.status === RoundStatus.Aborted) {
      throw new Error(`Round ${roundId} was aborted`);
    }

    await sleep(intervalMs);
  }
  throw new Error(
    `Timeout waiting for decryption of ciphertext ${ciphertextIndex} in round ${roundId}`,
  );
}

/**
 * Watch for new rounds created after `fromBlock`.
 * Calls `onRound` for each RoundCreated event.
 * Returns an unsubscribe function.
 */
export function watchNewRounds(
  client: DKGClient,
  onRound: (roundId: `0x${string}`, organizer: Address) => void,
  fromBlock?: bigint,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: [
      {
        type: 'event',
        name: 'RoundCreated',
        inputs: [
          { name: 'roundId', type: 'bytes12', indexed: true },
          { name: 'organizer', type: 'address', indexed: true },
          { name: 'seedBlock', type: 'uint64', indexed: false },
          { name: 'lotteryThreshold', type: 'uint256', indexed: false },
        ],
      },
    ] as const,
    eventName: 'RoundCreated',
    fromBlock,
    onLogs: (logs) => {
      for (const log of logs) {
        const { roundId, organizer } = log.args as any;
        if (roundId && organizer) onRound(roundId as `0x${string}`, organizer as Address);
      }
    },
  });
}

/**
 * Watch for a round being finalized.
 * Calls `onFinalized` once when the RoundFinalized event fires.
 * Returns an unsubscribe function.
 */
export function watchRoundFinalized(
  client: DKGClient,
  roundId: `0x${string}`,
  onFinalized: (collectivePublicKeyHash: `0x${string}`) => void,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: [
      {
        type: 'event',
        name: 'RoundFinalized',
        inputs: [
          { name: 'roundId', type: 'bytes12', indexed: true },
          { name: 'aggregateCommitmentsHash', type: 'bytes32', indexed: false },
          { name: 'collectivePublicKeyHash', type: 'bytes32', indexed: false },
          { name: 'shareCommitmentHash', type: 'bytes32', indexed: false },
        ],
      },
    ] as const,
    eventName: 'RoundFinalized',
    args: { roundId: roundId as any },
    onLogs: (logs) => {
      for (const log of logs) {
        const { collectivePublicKeyHash } = log.args as any;
        if (collectivePublicKeyHash) onFinalized(collectivePublicKeyHash as `0x${string}`);
      }
    },
  });
}

/**
 * Watch for the DecryptionCombined event for a specific ciphertext.
 * The callback receives the recovered plaintext scalar as a bigint.
 * Returns an unsubscribe function.
 */
export function watchDecryptionCombined(
  client: DKGClient,
  roundId: `0x${string}`,
  ciphertextIndex: number,
  onCombined: (combineHash: `0x${string}`, plaintext: bigint) => void,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: [
      {
        type: 'event',
        name: 'DecryptionCombined',
        inputs: [
          { name: 'roundId', type: 'bytes12', indexed: true },
          { name: 'ciphertextIndex', type: 'uint16', indexed: true },
          { name: 'combineHash', type: 'bytes32', indexed: false },
          { name: 'plaintext', type: 'uint256', indexed: false },
        ],
      },
    ] as const,
    eventName: 'DecryptionCombined',
    args: { roundId: roundId as any, ciphertextIndex } as any,
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
 * Watch for CiphertextSubmitted events on a round. The callback receives the
 * ciphertext index, submitter, and the raw (C1, C2) BabyJubJub coordinates —
 * the contract stores only the keccak hash, so the event log is the only way
 * to recover the coordinates nodes need for threshold decryption.
 * Returns an unsubscribe function.
 */
export function watchCiphertextSubmitted(
  client: DKGClient,
  roundId: `0x${string}`,
  onCiphertext: (payload: {
    ciphertextIndex: number;
    submitter: Address;
    c1: { x: bigint; y: bigint };
    c2: { x: bigint; y: bigint };
  }) => void,
): () => void {
  return client.publicClient.watchContractEvent({
    address: client.managerAddress,
    abi: [
      {
        type: 'event',
        name: 'CiphertextSubmitted',
        inputs: [
          { name: 'roundId', type: 'bytes12', indexed: true },
          { name: 'ciphertextIndex', type: 'uint16', indexed: true },
          { name: 'submitter', type: 'address', indexed: true },
          { name: 'c1x', type: 'uint256', indexed: false },
          { name: 'c1y', type: 'uint256', indexed: false },
          { name: 'c2x', type: 'uint256', indexed: false },
          { name: 'c2y', type: 'uint256', indexed: false },
        ],
      },
    ] as const,
    eventName: 'CiphertextSubmitted',
    args: { roundId: roundId as any } as any,
    onLogs: (logs) => {
      for (const log of logs) {
        const { ciphertextIndex, submitter, c1x, c1y, c2x, c2y } = log.args as any;
        if (typeof ciphertextIndex === 'number' && submitter) {
          onCiphertext({
            ciphertextIndex,
            submitter: submitter as Address,
            c1: { x: c1x as bigint, y: c1y as bigint },
            c2: { x: c2x as bigint, y: c2y as bigint },
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
  roundNonce: bigint;
}> {
  const [blockNumber, totalNodes, activeNodes, roundNonce] = await Promise.all([
    client.blockNumber(),
    client.nodeCount(),
    client.activeCount(),
    client.roundNonce(),
  ]);
  return { blockNumber, totalNodes, activeNodes, roundNonce };
}
