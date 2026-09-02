import {
  getAbiItem,
  getContract,
  type PublicClient,
  type Address,
  type GetContractReturnType,
} from 'viem';
import { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';
import {
  type Epoch,
  type EpochBounds,
  type ContributionRecord,
  type PartialDecryptionRecord,
  type CombinedDecryptionRecord,
  type CiphertextPoK,
  type NodeKey,
  type DKGConfig,
  type EpochPhaseValue,
  type EpochEvent,
  type EpochEntry,
} from './types.js';
import { buildEpochId } from './utils.js';
import { fromRTEtoTE } from './crypto/babyjub-form.js';
import { verifyCiphertextPoK } from './schnorr.js';

type ManagerContract = GetContractReturnType<typeof dkgManagerAbi, PublicClient>;
type RegistryContract = GetContractReturnType<typeof dkgRegistryAbi, PublicClient>;
type AppManagerContract = GetContractReturnType<typeof dkgAppManagerAbi, PublicClient>;

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
  private _appManager: AppManagerContract | null;
  private _resolvedAppManagerAddress: Address | null;

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

    if (config.appManagerAddress) {
      this._resolvedAppManagerAddress = config.appManagerAddress;
      this._appManager = getContract({
        address: config.appManagerAddress,
        abi: dkgAppManagerAbi,
        client: this.publicClient,
      });
    } else {
      this._resolvedAppManagerAddress = null;
      this._appManager = null;
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

  /**
   * The DKGAppManager address. Throws if not resolved yet — call
   * `_getAppManagerAddress()` (async) when you don't already have it cached.
   */
  get appManagerAddress(): Address {
    if (!this._resolvedAppManagerAddress) {
      throw new Error('appManagerAddress not yet resolved; call an app-manager method first or provide it in config');
    }
    return this._resolvedAppManagerAddress;
  }

  /** Resolve and cache the app manager contract, fetching its address from the manager when needed. */
  protected async _getAppManager(): Promise<AppManagerContract> {
    if (this._appManager) return this._appManager;
    const addr = (await this._manager.read.appManager()) as Address;
    this._resolvedAppManagerAddress = addr;
    this._appManager = getContract({
      address: addr,
      abi: dkgAppManagerAbi,
      client: this.publicClient,
    });
    return this._appManager;
  }

  /** Resolve and return the app manager address, fetching from the manager when needed. */
  protected async _getAppManagerAddress(): Promise<Address> {
    if (this._resolvedAppManagerAddress) return this._resolvedAppManagerAddress;
    await this._getAppManager();
    return this._resolvedAppManagerAddress!;
  }

  // ── Epoch ID utilities ─────────────────────────────────────────────────────

  /**
   * Fetch the current EPOCH_PREFIX and epochNonce, then assemble a epoch ID.
   * Call this after `createEpoch` is mined to derive the new epoch ID
   * without needing the transaction receipt.
   *
   * @param nonce  The nonce at epoch-creation time (epochNonce() before the tx)
   */
  async buildEpochId(nonce: bigint): Promise<`0x${string}`> {
    const prefix = await this._manager.read.EPOCH_PREFIX();
    return buildEpochId(prefix, nonce);
  }

  /** Current epoch nonce (incremented by each createEpoch call). */
  async epochNonce(): Promise<bigint> {
    return this._manager.read.epochNonce();
  }

  /**
   * The contract's compile-time epoch length in blocks. Set per-deploy via
   * the constructor; read once and cache. Pair with
   * {@link getNextEpochStartBlock} and a block-time estimate to render the
   * "next epoch in" countdown in the UI.
   */
  async getEpochDurationBlocks(): Promise<bigint> {
    return this._manager.read.epochDurationBlocks() as Promise<bigint>;
  }

  /**
   * Earliest block at which the next `createEpoch` call may succeed.
   * Equals `lastEpochStartBlock + EPOCH_DURATION_BLOCKS`, or `block.number`
   * before any epoch has been created.
   */
  async getNextEpochStartBlock(): Promise<bigint> {
    return BigInt(await this._manager.read.nextEpochStartBlock() as bigint);
  }

  /** Block in which the most recent epoch was created. */
  async getLastEpochStartBlock(): Promise<bigint> {
    return BigInt(await this._manager.read.lastEpochStartBlock() as bigint);
  }

  /**
   * Deploy-time bounds `createEpoch` enforces (the manager's immutables
   * `MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS`). The
   * contract additionally always requires
   * `1 ≤ threshold ≤ minValidContributions ≤ committeeSize ≤ MaxN` and
   * `lotteryAlphaBps ≥ 10000`; anything outside reverts `InvalidPolicy()`.
   * Immutable per deployment, so cache freely.
   */
  async getEpochBounds(): Promise<EpochBounds> {
    const [minThreshold, minCommitteeSize, maxLotteryAlphaBps] = await Promise.all([
      this._manager.read.MIN_THRESHOLD(),
      this._manager.read.MIN_COMMITTEE_SIZE(),
      this._manager.read.MAX_LOTTERY_ALPHA_BPS(),
    ]);
    return {
      minThreshold: Number(minThreshold),
      minCommitteeSize: Number(minCommitteeSize),
      maxLotteryAlphaBps: Number(maxLotteryAlphaBps),
    };
  }

  // ── Epoch queries ──────────────────────────────────────────────────────────

  /** Fetch full epoch state. */
  async getEpoch(epochId: `0x${string}`): Promise<Epoch> {
    const r = await this._manager.read.getEpoch([epochId as `0x${string}` & { length: 26 }]);
    return r as unknown as Epoch;
  }

  /** Fetch the list of addresses that claimed a slot in this epoch. */
  async selectedParticipants(epochId: `0x${string}`): Promise<Address[]> {
    return this._manager.read.selectedParticipants([epochId as any]) as Promise<Address[]>;
  }

  /** Fetch the contribution record for a specific contributor. */
  async getContribution(
    epochId: `0x${string}`,
    contributor: Address,
  ): Promise<ContributionRecord> {
    const r = await this._manager.read.getContribution([epochId as any, contributor]);
    return r as unknown as ContributionRecord;
  }

  /**
   * Fetch a partial decryption record for a specific participant and
   * ciphertext index. The `delta` curve point is converted from on-chain
   * RTE to TE for consistency with `getCollectivePublicKey`.
   */
  async getPartialDecryption(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    participantIndex: number,
    ciphertextIndex: number,
  ): Promise<PartialDecryptionRecord> {
    const r = await this._manager.read.getPartialDecryption([
      epochId as any,
      aid as any,
      participantIndex,
      ciphertextIndex,
    ]) as unknown as PartialDecryptionRecord;
    return r;
  }

  /**
   * Fetch the cached on-chain `Application` record for `(epochId, aid)`.
   * Returns an `exists: false` record when the application has not been
   * registered. Callers should check `record.exists` before reading
   * `mode`-specific fields. The `organizerPK` curve point is converted
   * from on-chain RTE to TE for consistency with `getCollectivePublicKey`.
   */
  async getApplication(
    epochId: `0x${string}`,
    aid: `0x${string}`,
  ): Promise<import('./types.js').ApplicationRecord> {
    const am = await this._getAppManager();
    const r = (await am.read.getApplication([epochId as any, aid as any])) as any;
    const [pkX, pkY] = fromRTEtoTE(BigInt(r.organizerPK.x), BigInt(r.organizerPK.y));
    return {
      creator: r.creator,
      mode: Number(r.mode) as 0 | 1,
      derivationS: BigInt(r.derivationS),
      organizerPK: [pkX, pkY],
      policy: {
        authorizedSubmitter: r.policy.authorizedSubmitter,
        maxCiphertexts: Number(r.policy.maxCiphertexts),
        notBeforeBlock: BigInt(r.policy.notBeforeBlock),
        notAfterBlock: BigInt(r.policy.notAfterBlock),
      },
      createdAtBlock: BigInt(r.createdAtBlock),
      exists: Boolean(r.exists),
    };
  }

  /** Fetch the combined decryption record for a ciphertext index. */
  async getCombinedDecryption(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertextIndex: number,
  ): Promise<CombinedDecryptionRecord> {
    const r = await this._manager.read.getCombinedDecryption([epochId as any, aid as any, ciphertextIndex]);
    return r as unknown as CombinedDecryptionRecord;
  }

  /**
   * Fetch the recovered plaintext for (epochId, ciphertextIndex). Returns 0n
   * when the decryption has not been combined yet — consumers that need to
   * disambiguate "not combined" from "plaintext is literally zero" should also
   * check `getCombinedDecryption(...).completed`.
   */
  async getPlaintext(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertextIndex: number,
  ): Promise<bigint> {
    return this._manager.read.getPlaintext([epochId as any, aid as any, ciphertextIndex]) as Promise<bigint>;
  }

  /**
   * keccak256(abi.encode(c1x, c1y, c2x, c2y)) for the ciphertext stored at
   * (epochId, aid, ciphertextIndex). Returns 0x00..00 when no ciphertext has
   * been submitted at this slot. The raw coordinates are only in the
   * `CiphertextSubmitted` event log.
   */
  async getCiphertextHash(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertextIndex: number,
  ): Promise<`0x${string}`> {
    return this._manager.read.getCiphertextHash([epochId as any, aid as any, ciphertextIndex]);
  }

  /**
   * Number of ciphertexts accepted so far under `(epochId, aid)`. Indices are
   * assigned on-chain as 1, 2, … in submission order, so the last accepted
   * ciphertext has index `ciphertextCount`.
   */
  async ciphertextCount(epochId: `0x${string}`, aid: `0x${string}`): Promise<number> {
    return Number(await this._manager.read.ciphertextCount([epochId as any, aid as any]));
  }

  /** Fetch the share-commitment hash for a given participant index. */
  async getShareCommitmentHash(
    epochId: `0x${string}`,
    participantIndex: number,
  ): Promise<`0x${string}`> {
    return this._manager.read.getShareCommitmentHash([epochId as any, participantIndex]);
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
   * Fetch EpochCreated events in the given block range.
   * Returns up to `count` most-recent events when `fromBlock` is omitted.
   */
  async getEpochCreatedEvents(options?: {
    fromBlock?: bigint;
    toBlock?: bigint;
  }): Promise<
    Array<{
      epochId: `0x${string}`;
      organizer: Address;
      startBlock: bigint;
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
          name: 'EpochCreated',
          inputs: [
            { name: 'epochId', type: 'bytes12', indexed: true },
            { name: 'organizer', type: 'address', indexed: true },
            { name: 'startBlock', type: 'uint64', indexed: false },
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
      epochId: (l.args as any).epochId as `0x${string}`,
      organizer: (l.args as any).organizer as Address,
      startBlock: BigInt((l.args as any).startBlock ?? 0),
      seedBlock: BigInt((l.args as any).seedBlock ?? 0),
      lotteryThreshold: BigInt((l.args as any).lotteryThreshold ?? 0),
      blockNumber: l.blockNumber ?? 0n,
    }));
  }

  /**
   * Fetch all EpochLive events for a specific epoch.
   */
  async getEpochLiveEvents(epochId: `0x${string}`): Promise<
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
          name: 'EpochLive',
          inputs: [
            { name: 'epochId', type: 'bytes12', indexed: true },
            { name: 'aggregateCommitmentsHash', type: 'bytes32', indexed: false },
            { name: 'collectivePublicKeyHash', type: 'bytes32', indexed: false },
            { name: 'shareCommitmentHash', type: 'bytes32', indexed: false },
          ],
        } as const,
        args: { epochId: epochId as any },
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
   * Fetch all CiphertextSubmitted events for a specific epoch, optionally
   * narrowed to one application (`aid`) and/or one `ciphertextIndex`. Each
   * entry carries the raw (C1, C2) coordinates in on-chain (RTE) form plus
   * the submitter's proof of knowledge; the contract only stores a keccak
   * hash, so this log is the only source of the coordinates nodes need for
   * threshold decryption.
   *
   * `pokValid` is the SDK's own verdict on the proof (`verifyCiphertextPoK`).
   * Committee nodes run the same check and never decrypt a ciphertext whose
   * proof fails, so a `false` here means the plaintext will never appear.
   */
  async getCiphertextSubmittedEvents(
    epochId: `0x${string}`,
    opts?: { aid?: `0x${string}`; ciphertextIndex?: number },
  ): Promise<
    Array<{
      aid: `0x${string}`;
      ciphertextIndex: number;
      submitter: Address;
      c1: { x: bigint; y: bigint };
      c2: { x: bigint; y: bigint };
      pok: CiphertextPoK;
      pokValid: boolean;
      blockNumber: bigint;
      transactionHash: `0x${string}` | null;
    }>
  > {
    const args: Record<string, unknown> = { epochId: epochId as any };
    if (opts?.aid != null) args.aid = opts.aid;
    if (opts?.ciphertextIndex != null) args.ciphertextIndex = opts.ciphertextIndex;
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: this.managerAddress,
        event: getAbiItem({ abi: dkgManagerAbi, name: 'CiphertextSubmitted' }),
        args,
        fromBlock: 0n,
        toBlock: 'latest',
      },
      { fallbackWindow: 50_000n },
    );
    return logs.map((l) => {
      const a = l.args as any;
      const aid = a.aid as `0x${string}`;
      const pok: CiphertextPoK = { ax: a.pokAx as bigint, ay: a.pokAy as bigint, z: a.pokZ as bigint };
      return {
        aid,
        ciphertextIndex: Number(a.ciphertextIndex),
        submitter: a.submitter as Address,
        c1: { x: a.c1x as bigint, y: a.c1y as bigint },
        c2: { x: a.c2x as bigint, y: a.c2y as bigint },
        pok,
        pokValid: verifyCiphertextPoK(epochId, aid, a.c1x, a.c1y, a.c2x, a.c2y, pok),
        blockNumber: l.blockNumber ?? 0n,
        transactionHash: (l.transactionHash ?? null) as `0x${string}` | null,
      };
    });
  }

  /**
   * Fetch all DecryptionCombined events for a specific epoch (optionally
   * filtered by `aid` and/or `ciphertextIndex`). Each entry contains the
   * plaintext scalar.
   */
  async getDecryptionCombinedEvents(
    epochId: `0x${string}`,
    opts?: { aid?: `0x${string}`; ciphertextIndex?: number },
  ): Promise<
    Array<{
      aid: `0x${string}`;
      ciphertextIndex: number;
      combineHash: `0x${string}`;
      plaintext: bigint;
      blockNumber: bigint;
      transactionHash: `0x${string}` | null;
    }>
  > {
    const args: Record<string, unknown> = { epochId: epochId as any };
    if (opts?.aid != null) args.aid = opts.aid;
    if (opts?.ciphertextIndex != null) args.ciphertextIndex = opts.ciphertextIndex;
    const logs = await getLogsChunked(
      this.publicClient,
      {
        address: this.managerAddress,
        event: getAbiItem({ abi: dkgManagerAbi, name: 'DecryptionCombined' }),
        args,
        fromBlock: 0n,
        toBlock: 'latest',
      },
      { fallbackWindow: 50_000n },
    );
    return logs.map((l) => {
      const a = l.args as any;
      return {
        aid: a.aid as `0x${string}`,
        ciphertextIndex: Number(a.ciphertextIndex),
        combineHash: a.combineHash as `0x${string}`,
        plaintext: a.plaintext as bigint,
        blockNumber: l.blockNumber ?? 0n,
        transactionHash: (l.transactionHash ?? null) as `0x${string}` | null,
      };
    });
  }

  /**
   * Returns the collective public key accumulated on-chain for the given epoch.
   *
   * The contract accumulates this key as contributions are submitted — each
   * accepted contributor's commitment[0] point (a_{i,0}·G) is added to a
   * running sum.  Once the epoch is finalized the value equals the full
   * collective public key.  The key is available as soon as the first
   * contribution is accepted.
   *
   * Returns { x: 0n, y: 1n } (the BabyJubJub identity) if no contributions
   * have been accepted yet.
   *
   * IMPORTANT: do NOT encrypt and call `submitCiphertext` with the value
   * returned during the Contribution phase. The contract's `submitCiphertext`
   * requires `EpochPhase.Live` and will revert otherwise. Either:
   *   - use `flow.ts:waitForCollectivePublicKeyHash` then read this getter, or
   *   - check `getEpoch(epochId).status === EpochPhase.Live` first.
   * Pre-finalize reads of this value are intended for monitoring/observation
   * (e.g. displaying the in-progress accumulator), not for producing
   * ciphertexts that will actually be sent on-chain.
   *
   * The point is converted from gnark's RTE form (used on-chain by the
   * contract and the ZK circuits) to circomlib's TE form so the result
   * composes directly with @zk-kit/baby-jubjub / circomlibjs operations
   * (mulPointEscalar, addPoint, etc.) used by this SDK's ElGamal layer.
   * See `crypto/babyjub-form.ts` for the rationale and the conversion
   * formula (vendored from davinci-sdk for wire compatibility).
   */
  async getCollectivePublicKey(
    epochId: `0x${string}`,
  ): Promise<{ x: bigint; y: bigint }> {
    const result = await this.publicClient.readContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'getCollectivePublicKey',
      args: [epochId as any],
    }) as { x: bigint; y: bigint };
    // Identity (0, 1) is the same in both forms — no-op conversion is safe.
    const [x, y] = fromRTEtoTE(result.x, result.y);
    return { x, y };
  }

  /**
   * @deprecated Use {@link getCollectivePublicKey} instead.
   * Kept for backwards compatibility — now simply delegates to the on-chain
   * getter which accumulates the key as contributions are submitted.
   */
  async getCollectivePublicKeyFromContributions(
    epochId: `0x${string}`,
    _participants?: Address[],
  ): Promise<{ x: bigint; y: bigint }> {
    return this.getCollectivePublicKey(epochId);
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

  /** Fetch the EPOCH_PREFIX constant. Cached cheaply because it is immutable. */
  async roundPrefix(): Promise<number> {
    return this._manager.read.EPOCH_PREFIX();
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
   * Fetch all DKGManager events that reference a specific epoch.
   *
   * Events are returned in block order (ascending).  The caller can filter
   * by `eventName` to isolate e.g. only `ContributionSubmitted` events.
   */
  async getAllEpochEvents(epochId: `0x${string}`, fromBlock = 0n): Promise<EpochEvent[]> {
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
      .filter((l) => 'args' in l && (l.args as any).epochId === epochId)
      .map((l) => ({
        eventName: (l as any).eventName as string,
        args: (l.args ?? {}) as Record<string, unknown>,
        blockNumber: l.blockNumber ?? 0n,
        transactionHash: (l.transactionHash ?? '0x') as `0x${string}`,
      }));
  }

  /**
   * Fetch the most recent `limit` epochs in descending nonce order.
   *
   * Epochs with status 0 (None) are omitted (an id that was never created).
   *
   * @param limit  Maximum number of epochs to return (default: 20)
   */
  async getRecentEpochs(limit = 20): Promise<EpochEntry[]> {
    const [nonce, prefix] = await Promise.all([
      this.epochNonce(),
      this.roundPrefix(),
    ]);
    if (nonce === 0n) return [];

    const RING_BUFFER_SIZE = 64n;
    const minNonce = nonce > RING_BUFFER_SIZE ? nonce - RING_BUFFER_SIZE + 1n : 1n;

    const ids: `0x${string}`[] = [];
    for (let i = nonce; i >= minNonce && ids.length < limit; i--) {
      ids.push(buildEpochId(prefix, i));
      if (i === 1n) break;
    }

    const entries = await Promise.all(
      ids.map(async (id) => {
        try {
          const epoch = await this.getEpoch(id);
          return { id, epoch };
        } catch {
          return null;
        }
      }),
    );
    return entries.filter(
      (e): e is EpochEntry => e !== null && Number(e.epoch.status) !== 0,
    );
  }

  /** @deprecated Use {@link getRecentEpochs}. Kept for SDK 0.1.x compatibility. */
  async getRecentRounds(limit = 20): Promise<EpochEntry[]> {
    return this.getRecentEpochs(limit);
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
