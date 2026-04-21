import {
  getContract,
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

  /**
   * Fetch the recovered plaintext for (roundId, ciphertextIndex). Returns 0n
   * when the decryption has not been combined yet — consumers that need to
   * disambiguate "not combined" from "plaintext is literally zero" should also
   * check `getCombinedDecryption(...).completed`.
   */
  async getPlaintext(
    roundId: `0x${string}`,
    ciphertextIndex: number,
  ): Promise<bigint> {
    return this._manager.read.getPlaintext([roundId as any, ciphertextIndex]) as Promise<bigint>;
  }

  /**
   * keccak256(abi.encode(c1x, c1y, c2x, c2y)) for the ciphertext stored at
   * (roundId, ciphertextIndex). Returns 0x00..00 when no ciphertext has been
   * submitted at this slot. The raw coordinates are only in the
   * `CiphertextSubmitted` event log.
   */
  async getCiphertextHash(
    roundId: `0x${string}`,
    ciphertextIndex: number,
  ): Promise<`0x${string}`> {
    return this._manager.read.getCiphertextHash([roundId as any, ciphertextIndex]);
  }

  /** Fetch the decryption policy set at round creation. */
  async getDecryptionPolicy(
    roundId: `0x${string}`,
  ): Promise<import('./types.js').DecryptionPolicy> {
    const r = await this._manager.read.getDecryptionPolicy([roundId as any]);
    return r as unknown as import('./types.js').DecryptionPolicy;
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
   * Fetch all CiphertextSubmitted events for a specific round. Each entry
   * carries the raw (C1, C2) coordinates; nodes and consumers need these to
   * perform threshold decryption since the contract only stores a keccak hash.
   */
  async getCiphertextSubmittedEvents(
    roundId: `0x${string}`,
    opts?: { ciphertextIndex?: number },
  ): Promise<
    Array<{
      ciphertextIndex: number;
      submitter: Address;
      c1: { x: bigint; y: bigint };
      c2: { x: bigint; y: bigint };
      blockNumber: bigint;
      transactionHash: `0x${string}` | null;
    }>
  > {
    const args: Record<string, unknown> = { roundId: roundId as any };
    if (opts?.ciphertextIndex != null) args.ciphertextIndex = opts.ciphertextIndex;
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: this.managerAddress,
        event: {
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
        } as const,
        args,
        fromBlock: 0n,
        toBlock: 'latest',
      },
      { fallbackWindow: 50_000n },
    );
    return logs.map((l) => {
      const a = l.args as any;
      return {
        ciphertextIndex: Number(a.ciphertextIndex),
        submitter: a.submitter as Address,
        c1: { x: a.c1x as bigint, y: a.c1y as bigint },
        c2: { x: a.c2x as bigint, y: a.c2y as bigint },
        blockNumber: l.blockNumber ?? 0n,
        transactionHash: (l.transactionHash ?? null) as `0x${string}` | null,
      };
    });
  }

  /**
   * Fetch all DecryptionCombined events for a specific round (optionally
   * filtered by `ciphertextIndex`). Each entry contains the plaintext scalar.
   */
  async getDecryptionCombinedEvents(
    roundId: `0x${string}`,
    opts?: { ciphertextIndex?: number },
  ): Promise<
    Array<{
      ciphertextIndex: number;
      combineHash: `0x${string}`;
      plaintext: bigint;
      blockNumber: bigint;
      transactionHash: `0x${string}` | null;
    }>
  > {
    const args: Record<string, unknown> = { roundId: roundId as any };
    if (opts?.ciphertextIndex != null) args.ciphertextIndex = opts.ciphertextIndex;
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: this.managerAddress,
        event: {
          type: 'event',
          name: 'DecryptionCombined',
          inputs: [
            { name: 'roundId', type: 'bytes12', indexed: true },
            { name: 'ciphertextIndex', type: 'uint16', indexed: true },
            { name: 'combineHash', type: 'bytes32', indexed: false },
            { name: 'plaintext', type: 'uint256', indexed: false },
          ],
        } as const,
        args,
        fromBlock: 0n,
        toBlock: 'latest',
      },
      { fallbackWindow: 50_000n },
    );
    return logs.map((l) => {
      const a = l.args as any;
      return {
        ciphertextIndex: Number(a.ciphertextIndex),
        combineHash: a.combineHash as `0x${string}`,
        plaintext: a.plaintext as bigint,
        blockNumber: l.blockNumber ?? 0n,
        transactionHash: (l.transactionHash ?? null) as `0x${string}` | null,
      };
    });
  }

  /**
   * Returns the collective public key accumulated on-chain for the given round.
   *
   * The contract accumulates this key as contributions are submitted — each
   * accepted contributor's commitment[0] point (a_{i,0}·G) is added to a
   * running sum.  Once the round is finalized the value equals the full
   * collective public key.  The key is available as soon as the first
   * contribution is accepted.
   *
   * Returns { x: 0n, y: 1n } (the BabyJubJub identity) if no contributions
   * have been accepted yet.
   */
  async getCollectivePublicKey(
    roundId: `0x${string}`,
  ): Promise<{ x: bigint; y: bigint }> {
    const result = await this.publicClient.readContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'getCollectivePublicKey',
      args: [roundId as any],
    }) as { x: bigint; y: bigint };
    return { x: result.x, y: result.y };
  }

  /**
   * @deprecated Use {@link getCollectivePublicKey} instead.
   * Kept for backwards compatibility — now simply delegates to the on-chain
   * getter which accumulates the key as contributions are submitted.
   */
  async getCollectivePublicKeyFromContributions(
    roundId: `0x${string}`,
    _participants?: Address[],
  ): Promise<{ x: bigint; y: bigint }> {
    return this.getCollectivePublicKey(roundId);
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
