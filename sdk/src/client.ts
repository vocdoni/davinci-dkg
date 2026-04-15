import {
  getContract,
  decodeFunctionData,
  type PublicClient,
  type Address,
  type GetContractReturnType,
} from 'viem';
import { dkgManagerAbi, dkgRegistryAbi } from './abi.js';
import {
  type Round,
  type ContributionRecord,
  type PartialDecryptionRecord,
  type CombinedDecryptionRecord,
  type RevealedShareRecord,
  type NodeKey,
  type DKGConfig,
  type RoundStatusValue,
  type RoundEvent,
  type RoundEntry,
} from './types.js';
import { buildRoundId } from './utils.js';

type ManagerContract = GetContractReturnType<typeof dkgManagerAbi, PublicClient>;
type RegistryContract = GetContractReturnType<typeof dkgRegistryAbi, PublicClient>;

/** Default chunk size for chunked getLogs (blocks per request). */
const DEFAULT_LOG_CHUNK = 2000n;
/** Minimum chunk size before giving up on adaptive reduction. */
const MIN_LOG_CHUNK = 100n;
/** Default fallback window when fromBlock is unknown (0). */
const DEFAULT_FALLBACK_WINDOW = 50_000n;

/**
 * Returns true when the error message indicates that the requested block
 * range exceeds the provider's getLogs limit.
 */
function isRangeTooLargeError(err: unknown): boolean {
  const msg = err instanceof Error ? err.message : String(err);
  return (
    msg.includes('range') ||
    msg.includes('block range') ||
    msg.includes('10000') ||
    msg.includes('10,000') ||
    msg.includes('exceed') ||
    msg.includes('too large') ||
    msg.includes('too many blocks') ||
    msg.includes('eth_getLogs is limited')
  );
}

/**
 * Compute the effective fromBlock for a log query.
 *
 * When `fromBlock` is 0n (unknown deployment block), clamp it to
 * `latestBlock - fallbackWindow` so queries never scan from genesis.
 */
function effectiveFromBlock(fromBlock: bigint, latestBlock: bigint, fallbackWindow: bigint): bigint {
  if (fromBlock === 0n && fallbackWindow > 0n) {
    return latestBlock > fallbackWindow ? latestBlock - fallbackWindow : 0n;
  }
  return fromBlock;
}

/**
 * Fetch logs over a potentially large block range by splitting it into chunks.
 * Uses `any` for the opts parameter to avoid viem's complex getLogs union types.
 * The caller is responsible for passing valid getLogs parameters.
 *
 * When `fromBlock` is 0n and `fallbackWindow > 0`, the scan floor is clamped
 * to `latestBlock - fallbackWindow` to avoid scanning from genesis.
 */
async function getLogsChunked(
  client: PublicClient,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  opts: any,
  options: { chunkSize?: bigint; fallbackWindow?: bigint } = {},
): Promise<any[]> {
  const chunkSize = options.chunkSize ?? DEFAULT_LOG_CHUNK;
  const fallbackWindow = options.fallbackWindow ?? DEFAULT_FALLBACK_WINDOW;

  const latest = await client.getBlockNumber();
  const toBlock: bigint =
    opts.toBlock === 'latest' || opts.toBlock == null ? latest : BigInt(opts.toBlock as string | bigint);
  const rawFrom: bigint = opts.fromBlock != null ? BigInt(opts.fromBlock as string | bigint) : 0n;
  const fromBlock = effectiveFromBlock(rawFrom, latest, fallbackWindow);

  const all: any[] = [];
  let currentChunk = chunkSize;
  let cursor = fromBlock;

  while (cursor <= toBlock) {
    const end = cursor + currentChunk - 1n > toBlock ? toBlock : cursor + currentChunk - 1n;
    try {
      const chunk = await client.getLogs({ ...opts, fromBlock: cursor, toBlock: end });
      all.push(...chunk);
      cursor = end + 1n;
    } catch (err) {
      if (isRangeTooLargeError(err) && currentChunk > MIN_LOG_CHUNK) {
        currentChunk = currentChunk / 2n;
        continue;
      }
      throw err;
    }
  }
  return all;
}

/**
 * Read-only client for the DKG Manager and Registry contracts.
 *
 * Construct it once with a viem `PublicClient` and use its methods
 * to query on-chain state without needing a signer.
 */
export class DKGClient {
  readonly publicClient: PublicClient;
  readonly managerAddress: Address;

  private _manager: ManagerContract;
  private _registry: RegistryContract | null;
  private _resolvedRegistryAddress: Address | null;

  constructor(config: DKGConfig) {
    this.publicClient = config.publicClient;
    this.managerAddress = config.managerAddress;

    this._manager = getContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      client: this.publicClient,
    });

    if (config.registryAddress) {
      this._resolvedRegistryAddress = config.registryAddress;
      this._registry = getContract({
        address: config.registryAddress,
        abi: dkgRegistryAbi,
        client: this.publicClient,
      });
    } else {
      this._resolvedRegistryAddress = null;
      this._registry = null;
    }
  }

  /**
   * The DKGRegistry address. Throws if registry has not been resolved yet.
   * Use `_getRegistryAddress()` in async methods instead.
   */
  get registryAddress(): Address {
    if (!this._resolvedRegistryAddress) {
      throw new Error('registryAddress not yet resolved; call a registry method first or provide it in config');
    }
    return this._resolvedRegistryAddress;
  }

  /** Resolve and cache the registry contract, fetching its address from the manager when needed. */
  private async _getRegistry(): Promise<RegistryContract> {
    if (this._registry) return this._registry;
    const addr = await this._manager.read.REGISTRY();
    this._resolvedRegistryAddress = addr;
    this._registry = getContract({
      address: addr,
      abi: dkgRegistryAbi,
      client: this.publicClient,
    });
    return this._registry;
  }

  /** Resolve and return the registry address, fetching it from the manager when needed. */
  private async _getRegistryAddress(): Promise<Address> {
    if (this._resolvedRegistryAddress) return this._resolvedRegistryAddress;
    await this._getRegistry();
    return this._resolvedRegistryAddress!;
  }

  // ── Round ID utilities ─────────────────────────────────────────────────────

  /**
   * Fetch the current ROUND_PREFIX and roundNonce, then assemble a round ID.
   * Call this after `createRound` is mined to derive the new round ID
   * without needing the transaction receipt.
   *
   * @param nonce  The nonce at round-creation time (roundNonce() before the tx)
   */
  async buildRoundId(nonce: bigint): Promise<`0x${string}`> {
    const prefix = await this._manager.read.ROUND_PREFIX();
    return buildRoundId(prefix, nonce);
  }

  /** Current round nonce (incremented by each createRound call). */
  async roundNonce(): Promise<bigint> {
    return this._manager.read.roundNonce();
  }

  // ── Round queries ──────────────────────────────────────────────────────────

  /** Fetch full round state. */
  async getRound(roundId: `0x${string}`): Promise<Round> {
    const r = await this._manager.read.getRound([roundId as `0x${string}` & { length: 26 }]);
    return r as unknown as Round;
  }

  /** Fetch the list of addresses that claimed a slot in this round. */
  async selectedParticipants(roundId: `0x${string}`): Promise<Address[]> {
    return this._manager.read.selectedParticipants([roundId as any]) as Promise<Address[]>;
  }

  /** Fetch the contribution record for a specific contributor. */
  async getContribution(
    roundId: `0x${string}`,
    contributor: Address,
  ): Promise<ContributionRecord> {
    const r = await this._manager.read.getContribution([roundId as any, contributor]);
    return r as unknown as ContributionRecord;
  }

  /** Fetch a partial decryption record for a specific participant and ciphertext index. */
  async getPartialDecryption(
    roundId: `0x${string}`,
    participant: Address,
    ciphertextIndex: number,
  ): Promise<PartialDecryptionRecord> {
    const r = await this._manager.read.getPartialDecryption([
      roundId as any,
      participant,
      ciphertextIndex,
    ]);
    return r as unknown as PartialDecryptionRecord;
  }

  /** Fetch the combined decryption record for a ciphertext index. */
  async getCombinedDecryption(
    roundId: `0x${string}`,
    ciphertextIndex: number,
  ): Promise<CombinedDecryptionRecord> {
    const r = await this._manager.read.getCombinedDecryption([roundId as any, ciphertextIndex]);
    return r as unknown as CombinedDecryptionRecord;
  }

  /** Fetch the revealed share record for a participant. */
  async getRevealedShare(
    roundId: `0x${string}`,
    participant: Address,
  ): Promise<RevealedShareRecord> {
    const r = await this._manager.read.getRevealedShare([roundId as any, participant]);
    return r as unknown as RevealedShareRecord;
  }

  /** Fetch the share-commitment hash for a given participant index. */
  async getShareCommitmentHash(
    roundId: `0x${string}`,
    participantIndex: number,
  ): Promise<`0x${string}`> {
    return this._manager.read.getShareCommitmentHash([roundId as any, participantIndex]);
  }

  // ── Verifier key hashes ────────────────────────────────────────────────────

  async getContributionVerifierVKeyHash(): Promise<`0x${string}`> {
    return this._manager.read.getContributionVerifierVKeyHash();
  }

  async getPartialDecryptVerifierVKeyHash(): Promise<`0x${string}`> {
    return this._manager.read.getPartialDecryptVerifierVKeyHash();
  }

  async getFinalizeVerifierVKeyHash(): Promise<`0x${string}`> {
    return this._manager.read.getFinalizeVerifierVKeyHash();
  }

  async getDecryptCombineVerifierVKeyHash(): Promise<`0x${string}`> {
    return this._manager.read.getDecryptCombineVerifierVKeyHash();
  }

  async getRevealSubmitVerifierVKeyHash(): Promise<`0x${string}`> {
    return this._manager.read.getRevealSubmitVerifierVKeyHash();
  }

  async getRevealShareVerifierVKeyHash(): Promise<`0x${string}`> {
    return this._manager.read.getRevealShareVerifierVKeyHash();
  }

  // ── Registry queries ───────────────────────────────────────────────────────

  /** Fetch the NodeKey record for an operator address. */
  async getNode(operator: Address): Promise<NodeKey> {
    const registry = await this._getRegistry();
    const r = await registry.read.getNode([operator]);
    return r as unknown as NodeKey;
  }

  /** Total number of ever-registered nodes. */
  async nodeCount(): Promise<bigint> {
    const registry = await this._getRegistry();
    return registry.read.nodeCount();
  }

  /** Number of currently-active nodes. */
  async activeCount(): Promise<bigint> {
    const registry = await this._getRegistry();
    return registry.read.activeCount();
  }

  /** Whether the given operator is currently active. */
  async isActive(operator: Address): Promise<boolean> {
    const registry = await this._getRegistry();
    return registry.read.isActive([operator]);
  }

  /** The inactivity window in blocks after which a node can be reaped. */
  async inactivityWindow(): Promise<bigint> {
    const registry = await this._getRegistry();
    return registry.read.INACTIVITY_WINDOW();
  }

  // ── Chain utilities ────────────────────────────────────────────────────────

  /** Current block number. */
  async blockNumber(): Promise<bigint> {
    return this.publicClient.getBlockNumber();
  }

  // ── Event queries ──────────────────────────────────────────────────────────

  /**
   * Fetch RoundCreated events in the given block range.
   * Returns up to `count` most-recent events when `fromBlock` is omitted.
   */
  async getRoundCreatedEvents(options?: {
    fromBlock?: bigint;
    toBlock?: bigint;
  }): Promise<
    Array<{
      roundId: `0x${string}`;
      organizer: Address;
      seedBlock: bigint;
      lotteryThreshold: bigint;
      blockNumber: bigint;
    }>
  > {
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: this.managerAddress,
        event: {
          type: 'event',
          name: 'RoundCreated',
          inputs: [
            { name: 'roundId', type: 'bytes12', indexed: true },
            { name: 'organizer', type: 'address', indexed: true },
            { name: 'seedBlock', type: 'uint64', indexed: false },
            { name: 'lotteryThreshold', type: 'uint256', indexed: false },
          ],
        } as const,
        fromBlock: options?.fromBlock,
        toBlock: options?.toBlock,
      },
      { fallbackWindow: 50_000n },
    );
    return logs.map((l) => ({
      roundId: (l.args as any).roundId as `0x${string}`,
      organizer: (l.args as any).organizer as Address,
      seedBlock: BigInt((l.args as any).seedBlock ?? 0),
      lotteryThreshold: BigInt((l.args as any).lotteryThreshold ?? 0),
      blockNumber: l.blockNumber ?? 0n,
    }));
  }

  /**
   * Fetch all RoundFinalized events for a specific round.
   */
  async getRoundFinalizedEvents(roundId: `0x${string}`): Promise<
    Array<{
      aggregateCommitmentsHash: `0x${string}`;
      collectivePublicKeyHash: `0x${string}`;
      shareCommitmentHash: `0x${string}`;
      blockNumber: bigint;
      transactionHash: `0x${string}` | null;
    }>
  > {
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: this.managerAddress,
        event: {
          type: 'event',
          name: 'RoundFinalized',
          inputs: [
            { name: 'roundId', type: 'bytes12', indexed: true },
            { name: 'aggregateCommitmentsHash', type: 'bytes32', indexed: false },
            { name: 'collectivePublicKeyHash', type: 'bytes32', indexed: false },
            { name: 'shareCommitmentHash', type: 'bytes32', indexed: false },
          ],
        } as const,
        args: { roundId: roundId as any },
        fromBlock: 0n,
        toBlock: 'latest',
      },
      { fallbackWindow: 50_000n },
    );
    return logs.map((l) => ({
      aggregateCommitmentsHash: (l.args as any).aggregateCommitmentsHash as `0x${string}`,
      collectivePublicKeyHash: (l.args as any).collectivePublicKeyHash as `0x${string}`,
      shareCommitmentHash: (l.args as any).shareCommitmentHash as `0x${string}`,
      blockNumber: l.blockNumber ?? 0n,
      transactionHash: (l.transactionHash ?? null) as `0x${string}` | null,
    }));
  }

  /**
   * Extract the collective public key (x, y) for a finalized round.
   *
   * The coordinates are recovered from the `finalizeRound` transaction
   * calldata: the `transcript` parameter encodes the aggregate commitment
   * points as 32-byte words, and the first aggregate commitment
   * (AggregateCommitments[0]) is the collective public key.
   *
   * Transcript layout (MAX_N = 16 words):
   *   words [0..N)          participantIndexes
   *   words [N..N+2N²)      contributionCommitments (N*N points × 2 coords)
   *   words [N+2N²..+2N)    aggregateCommitments    (N points × 2 coords)
   *   words [N+2N²+2N..+2N) shareCommitments
   *
   * AggregateCommitments[0].x is at word index N + 2*N*N = 528 (N=16).
   */
  async getCollectivePublicKey(
    roundId: `0x${string}`,
    fromBlock?: bigint,
  ): Promise<{ x: bigint; y: bigint }> {
    const events = await this.getRoundFinalizedEvents(roundId);
    if (events.length === 0) {
      throw new Error(`No RoundFinalized event found for round ${roundId}`);
    }
    const txHash = events[events.length - 1].transactionHash;
    if (!txHash) throw new Error('RoundFinalized log has no transaction hash');

    const tx = await this.publicClient.getTransaction({ hash: txHash });

    // Decode calldata: finalizeRound(bytes12, bytes32, bytes32, bytes32, bytes, bytes, bytes)
    const decoded = decodeFunctionData({ abi: dkgManagerAbi, data: tx.input });
    if (decoded.functionName !== 'finalizeRound') {
      throw new Error(`Unexpected function: ${decoded.functionName}`);
    }
    // args: [roundId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input]
    const transcript = decoded.args[4] as `0x${string}`;

    // Each word is 32 bytes = 64 hex chars.  MAX_N = 16.
    const N = 16;
    const aggOffset = N + 2 * N * N; // = 528
    const hexBody = transcript.slice(2);   // strip '0x'
    const x = BigInt('0x' + hexBody.slice(aggOffset * 64, (aggOffset + 1) * 64));
    const y = BigInt('0x' + hexBody.slice((aggOffset + 1) * 64, (aggOffset + 2) * 64));
    return { x, y };
  }

  /**
   * Derive the collective public key directly from the `submitContribution`
   * calldata of accepted contributors — without waiting for `finalizeRound`.
   *
   * The collective public key is the sum of each contributor's zeroth Feldman
   * commitment point: sum_i( commitment_i[0] ) = sum_i( a_{i,0} * G ).
   *
   * This is mathematically identical to what `finalizeRound` computes, so the
   * result can be used for ElGamal encryption even while the round is still in
   * the Contribution phase (status 2).  No ZK-proof verification is performed;
   * callers should treat this as an optimistic/early value and cross-check
   * against `collectivePublicKeyHash` once the round is finalized.
   *
   * Transcript layout of `submitContribution` (N = 16 = MaxN):
   *   bytes [0..2N×32)    commitment points  (N points, X then Y, 32 bytes each)
   *   bytes [2N×32..3N×32) recipient indexes
   *   bytes [3N×32..5N×32) recipient pub keys
   *   bytes [5N×32..7N×32) ephemerals
   *   bytes [7N×32..8N×32) masked shares
   *
   * @param participants  Addresses of the accepted contributors.  If omitted the
   *                      method calls `selectedParticipants` on-chain to obtain them.
   */
  async getCollectivePublicKeyFromContributions(
    roundId: `0x${string}`,
    participants?: Address[],
  ): Promise<{ x: bigint; y: bigint }> {
    const { addPoint } = await import('@zk-kit/baby-jubjub');

    const addrs = participants ?? await this.selectedParticipants(roundId);
    if (addrs.length === 0) {
      throw new Error('getCollectivePublicKeyFromContributions: no participants');
    }

    // Fetch all ContributionSubmitted events for this round to get tx hashes.
    const events = await getLogsChunked(this.publicClient, {
      address: this.managerAddress,
      event: {
        type: 'event',
        name: 'ContributionSubmitted',
        inputs: [
          { name: 'roundId',           type: 'bytes12', indexed: true  },
          { name: 'contributor',       type: 'address',  indexed: true  },
          { name: 'contributorIndex',  type: 'uint16',   indexed: false },
          { name: 'commitmentsHash',   type: 'bytes32',  indexed: false },
          { name: 'encryptedSharesHash', type: 'bytes32', indexed: false },
        ],
      },
      args: { roundId: roundId as any },
    }, { fallbackWindow: 50_000n });

    // Index events by contributor address (lowercase) for O(1) lookup.
    const txByContributor = new Map<string, `0x${string}`>();
    for (const ev of events) {
      const contributor = (ev.args as any).contributor as string | undefined;
      if (contributor && ev.transactionHash) {
        txByContributor.set(contributor.toLowerCase(), ev.transactionHash);
      }
    }

    // For each accepted participant, fetch their tx, decode commitment[0], and add.
    // BabyJubJub identity = (0, 1) in twisted-Edwards affine coordinates.
    const N = 16; // MaxN
    let acc: [bigint, bigint] = [0n, 1n];

    for (const addr of addrs) {
      const txHash = txByContributor.get(addr.toLowerCase());
      if (!txHash) {
        throw new Error(`getCollectivePublicKeyFromContributions: no ContributionSubmitted event for ${addr}`);
      }

      const tx = await this.publicClient.getTransaction({ hash: txHash });
      const calldata = tx.input; // hex string with '0x' prefix

      // ABI-decode the transcript bytes from submitContribution calldata.
      // Head layout (32 bytes each): roundId | contributorIndex | commitmentsHash |
      //   encryptedSharesHash | transcript_offset | proof_offset | input_offset
      // transcript_offset is at bytes 128..159 of the payload (after 4-byte selector).
      const hex = calldata.slice(2); // strip '0x'
      if (hex.length < (4 + 160) * 2) {
        throw new Error(`calldata too short for ${addr}`);
      }
      // payload starts at char 8 (4 bytes selector × 2 hex chars/byte)
      const payload = hex.slice(8);
      const transcriptOffsetHex = payload.slice(128 * 2, 160 * 2);
      const transcriptOffset = Number(BigInt('0x' + transcriptOffsetHex));

      if ((transcriptOffset + 32) * 2 > payload.length) {
        throw new Error(`transcript offset out of range for ${addr}`);
      }
      const transcriptLen = Number(BigInt('0x' + payload.slice(transcriptOffset * 2, (transcriptOffset + 32) * 2)));
      const transcriptStart = transcriptOffset + 32;

      if ((transcriptStart + transcriptLen) * 2 > payload.length) {
        throw new Error(`transcript out of range for ${addr}`);
      }
      const tr = payload.slice(transcriptStart * 2, (transcriptStart + transcriptLen) * 2);

      // commitment[0] = first 64 bytes of transcript (X at [0..32), Y at [32..64))
      if (tr.length < N * 2 * 64) {
        throw new Error(`transcript too short for ${addr}: ${tr.length / 2} bytes`);
      }
      const cx = BigInt('0x' + tr.slice(0, 64));
      const cy = BigInt('0x' + tr.slice(64, 128));

      acc = addPoint(acc, [cx, cy]) as [bigint, bigint];
    }

    return { x: acc[0], y: acc[1] };
  }

  /**
   * Watch for any DKG Manager event and call the handler with the parsed log.
   * Returns an unsubscribe function.
   */
  watchManagerEvents(handler: (log: any) => void): () => void {
    return this.publicClient.watchContractEvent({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      onLogs: (logs) => logs.forEach(handler),
    });
  }

  /**
   * Watch for any DKG Registry event and call the handler with the parsed log.
   * Returns a Promise that resolves to an unsubscribe function once the
   * registry address has been resolved.
   */
  async watchRegistryEvents(handler: (log: any) => void): Promise<() => void> {
    const addr = await this._getRegistryAddress();
    return this.publicClient.watchContractEvent({
      address: addr,
      abi: dkgRegistryAbi,
      onLogs: (logs) => logs.forEach(handler),
    });
  }

  // ── Extended queries ───────────────────────────────────────────────────────

  /** Fetch the ROUND_PREFIX constant. Cached cheaply because it is immutable. */
  async roundPrefix(): Promise<number> {
    return this._manager.read.ROUND_PREFIX();
  }

  /**
   * Fetch all registered node records.
   *
   * Discovers operator addresses via NodeRegistered events then fetches
   * the current NodeKey for each one (which reflects any key updates).
   */
  async getRegistryNodes(fromBlock = 0n): Promise<NodeKey[]> {
    const registryAddr = await this._getRegistryAddress();
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: registryAddr,
        event: {
          type: 'event',
          name: 'NodeRegistered',
          inputs: [
            { name: 'operator', type: 'address', indexed: true },
            { name: 'pubX', type: 'uint256', indexed: false },
            { name: 'pubY', type: 'uint256', indexed: false },
          ],
        } as const,
        fromBlock,
        toBlock: 'latest',
      },
      { fallbackWindow: 50_000n },
    );

    // De-duplicate by lower-cased address; preserve insertion order.
    const seen = new Set<string>();
    const operators: Address[] = [];
    for (const l of logs) {
      const op = ((l.args as any).operator as string | undefined)?.toLowerCase();
      if (op && !seen.has(op)) {
        seen.add(op);
        operators.push(op as Address);
      }
    }

    const nodes = await Promise.all(
      operators.map(async (op) => {
        try {
          return await this.getNode(op);
        } catch {
          return null;
        }
      }),
    );
    return nodes.filter((n): n is NodeKey => n !== null);
  }

  /**
   * Fetch all DKGManager events that reference a specific round.
   *
   * Events are returned in block order (ascending).  The caller can filter
   * by `eventName` to isolate e.g. only `ContributionSubmitted` events.
   */
  async getAllRoundEvents(roundId: `0x${string}`, fromBlock = 0n): Promise<RoundEvent[]> {
    const latest = await this.publicClient.getBlockNumber();
    let start = fromBlock;
    // Apply fallback window when fromBlock is unknown (0).
    if (start === 0n) {
      start = latest > 50_000n ? latest - 50_000n : 0n;
    }
    let currentChunk = DEFAULT_LOG_CHUNK;
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const allLogs: any[] = [];
    let cursor = start;
    while (cursor <= latest) {
      const end = cursor + currentChunk - 1n > latest ? latest : cursor + currentChunk - 1n;
      try {
        const chunk = await this.publicClient.getContractEvents({
          address: this.managerAddress,
          abi: dkgManagerAbi,
          fromBlock: cursor,
          toBlock: end,
        });
        allLogs.push(...chunk);
        cursor = end + 1n;
      } catch (err) {
        if (isRangeTooLargeError(err) && currentChunk > MIN_LOG_CHUNK) {
          currentChunk = currentChunk / 2n;
          continue;
        }
        throw err;
      }
    }
    const logs = allLogs;
    return logs
      .filter((l) => 'args' in l && (l.args as any).roundId === roundId)
      .map((l) => ({
        eventName: (l as any).eventName as string,
        args: (l.args ?? {}) as Record<string, unknown>,
        blockNumber: l.blockNumber ?? 0n,
        transactionHash: (l.transactionHash ?? '0x') as `0x${string}`,
      }));
  }

  /**
   * Fetch the most recent `limit` rounds in descending nonce order.
   *
   * Rounds with status 0 (None) are omitted — they indicate an evicted slot
   * in the ring buffer.
   *
   * @param limit  Maximum number of rounds to return (default: 20)
   */
  async getRecentRounds(limit = 20): Promise<RoundEntry[]> {
    const [nonce, prefix] = await Promise.all([
      this.roundNonce(),
      this.roundPrefix(),
    ]);
    if (nonce === 0n) return [];

    const RING_BUFFER_SIZE = 64n;
    const minNonce = nonce > RING_BUFFER_SIZE ? nonce - RING_BUFFER_SIZE + 1n : 1n;

    const ids: `0x${string}`[] = [];
    for (let i = nonce; i >= minNonce && ids.length < limit; i--) {
      ids.push(buildRoundId(prefix, i));
      if (i === 1n) break;
    }

    const entries = await Promise.all(
      ids.map(async (id) => {
        try {
          const round = await this.getRound(id);
          return { id, round };
        } catch {
          return null;
        }
      }),
    );
    return entries.filter(
      (e): e is RoundEntry => e !== null && Number(e.round.status) !== 0,
    );
  }

  // ── Internal access for DKGWriter ──────────────────────────────────────────

  /** @internal Exposed for DKGWriter to reuse the same contract handle. */
  get _managerContract(): ManagerContract {
    return this._manager;
  }

  /** @internal Async accessor for DKGWriter to reuse the registry contract handle. */
  async _registryContract(): Promise<RegistryContract> {
    return this._getRegistry();
  }

  /** @internal Async accessor for DKGWriter to get the resolved registry address. */
  async _registryAddressResolved(): Promise<Address> {
    return this._getRegistryAddress();
  }
}
