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
    const logs = await this.publicClient.getLogs({
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
      fromBlock: options?.fromBlock ?? 0n,
      toBlock: options?.toBlock ?? 'latest',
    });
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
    }>
  > {
    const logs = await this.publicClient.getLogs({
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
    });
    return logs.map((l) => ({
      aggregateCommitmentsHash: (l.args as any).aggregateCommitmentsHash as `0x${string}`,
      collectivePublicKeyHash: (l.args as any).collectivePublicKeyHash as `0x${string}`,
      shareCommitmentHash: (l.args as any).shareCommitmentHash as `0x${string}`,
      blockNumber: l.blockNumber ?? 0n,
    }));
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
    const logs = await this.publicClient.getLogs({
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
    });

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
    const logs = await this.publicClient.getContractEvents({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      fromBlock,
      toBlock: 'latest',
    });
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
