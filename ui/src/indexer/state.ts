// Contract state the events do not carry: epoch structs, registry records,
// application records, share commitments and transaction metadata.
//
// Everything goes through `multicall` against the canonical Multicall3
// address (deployed on Sepolia and pre-deployed by Anvil). A chain without it
// simply falls back to bounded-concurrency `eth_call`s, so the explorer still
// works — it just costs more requests.

import { dkgAppManagerAbi, dkgManagerAbi, dkgRegistryAbi, fromRTEtoTE } from '@vocdoni/davinci-dkg-sdk'
import type { Abi, Address, PublicClient } from 'viem'
import type { Aid, ChainMeta, EpochId, EpochPolicy, Hex, TxMeta } from './types'
import type { ApplicationStateUpdate, EpochStateUpdate, OperatorStateUpdate } from './reduce'

/** Canonical Multicall3 deployment, identical on every supported chain. */
export const MULTICALL3_ADDRESS = '0xcA11bde05977b3631167028862bE2a173976CA11' as Address

type CallSpec = {
  address: Address
  abi: Abi
  functionName: string
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  args?: any[]
}

type CallResult = { ok: true; value: unknown } | { ok: false; error: unknown }

export interface StateReaderOptions {
  client: PublicClient
  managerAddress: Address
  registryAddress?: Address | null
  appManagerAddress?: Address | null
  /** Calls per multicall request. */
  batchSize?: number
  /** Parallel `eth_call`s when multicall is unavailable. */
  concurrency?: number
  multicallAddress?: Address
}

const MANAGER_ABI = dkgManagerAbi as unknown as Abi
const REGISTRY_ABI = dkgRegistryAbi as unknown as Abi
const APP_MANAGER_ABI = dkgAppManagerAbi as unknown as Abi

function num(v: unknown): number {
  return typeof v === 'bigint' ? Number(v) : Number(v ?? 0)
}

function big(v: unknown): bigint {
  return typeof v === 'bigint' ? v : BigInt(Number(v ?? 0))
}

export class StateReader {
  readonly client: PublicClient
  readonly managerAddress: Address
  registryAddress: Address | null
  appManagerAddress: Address | null
  private readonly batchSize: number
  private readonly concurrency: number
  private readonly multicallAddress: Address
  /** Flips to false the first time multicall is rejected by the chain. */
  private multicallSupported = true
  /** Requests issued (multicall batches + individual calls). */
  requests = 0

  constructor(options: StateReaderOptions) {
    this.client = options.client
    this.managerAddress = options.managerAddress.toLowerCase() as Address
    this.registryAddress = (options.registryAddress?.toLowerCase() as Address) ?? null
    this.appManagerAddress = (options.appManagerAddress?.toLowerCase() as Address) ?? null
    this.batchSize = options.batchSize ?? 64
    this.concurrency = options.concurrency ?? 8
    this.multicallAddress = options.multicallAddress ?? MULTICALL3_ADDRESS
  }

  get usesMulticall(): boolean {
    return this.multicallSupported
  }

  /** Run `calls`, never throwing per call: failures come back as `ok: false`. */
  async read(calls: CallSpec[]): Promise<CallResult[]> {
    if (calls.length === 0) return []
    const out: CallResult[] = []
    for (let i = 0; i < calls.length; i += this.batchSize) {
      const batch = calls.slice(i, i + this.batchSize)
      out.push(...(await this.readBatch(batch)))
    }
    return out
  }

  private async readBatch(batch: CallSpec[]): Promise<CallResult[]> {
    if (this.multicallSupported) {
      try {
        this.requests += 1
        const results = await this.client.multicall({
          // eslint-disable-next-line @typescript-eslint/no-explicit-any
          contracts: batch as any,
          allowFailure: true,
          multicallAddress: this.multicallAddress,
        })
        return results.map((r) =>
          r.status === 'success' ? { ok: true as const, value: r.result } : { ok: false as const, error: r.error },
        )
      } catch (err) {
        // Chain without Multicall3, or a provider that rejects the batch:
        // degrade once and stay degraded for the session.
        this.multicallSupported = false
        void err
      }
    }
    return this.readSequential(batch)
  }

  private async readSequential(batch: CallSpec[]): Promise<CallResult[]> {
    const out: CallResult[] = new Array(batch.length)
    let cursor = 0
    const worker = async (): Promise<void> => {
      for (;;) {
        const i = cursor++
        if (i >= batch.length) return
        const call = batch[i]
        try {
          this.requests += 1
          const value = await this.client.readContract({
            address: call.address,
            abi: call.abi,
            functionName: call.functionName,
            // eslint-disable-next-line @typescript-eslint/no-explicit-any
            args: call.args as any,
          })
          out[i] = { ok: true, value }
        } catch (error) {
          out[i] = { ok: false, error }
        }
      }
    }
    await Promise.all(Array.from({ length: Math.min(this.concurrency, batch.length) }, worker))
    return out
  }

  /** Addresses and immutables — read once per session, then rarely. */
  async readChainMeta(stateBlock: number): Promise<Partial<ChainMeta>> {
    const managerCalls: CallSpec[] = [
      { address: this.managerAddress, abi: MANAGER_ABI, functionName: 'REGISTRY' },
      { address: this.managerAddress, abi: MANAGER_ABI, functionName: 'appManager' },
      { address: this.managerAddress, abi: MANAGER_ABI, functionName: 'EPOCH_PREFIX' },
      { address: this.managerAddress, abi: MANAGER_ABI, functionName: 'epochDurationBlocks' },
      { address: this.managerAddress, abi: MANAGER_ABI, functionName: 'nextEpochStartBlock' },
    ]
    const [registry, appManager, prefix, duration, nextStart] = await this.read(managerCalls)
    const meta: Partial<ChainMeta> = { stateBlock }
    if (registry.ok) {
      this.registryAddress = String(registry.value).toLowerCase() as Address
      meta.registryAddress = this.registryAddress
    }
    if (appManager.ok) {
      this.appManagerAddress = String(appManager.value).toLowerCase() as Address
      meta.appManagerAddress = this.appManagerAddress
    }
    if (prefix.ok) meta.epochPrefix = num(prefix.value)
    if (duration.ok) meta.epochDurationBlocks = num(duration.value)
    if (nextStart.ok) meta.nextEpochStartBlock = num(nextStart.value)

    if (this.registryAddress) {
      const [nodeCount, activeCount, window] = await this.read([
        { address: this.registryAddress, abi: REGISTRY_ABI, functionName: 'nodeCount' },
        { address: this.registryAddress, abi: REGISTRY_ABI, functionName: 'activeCount' },
        { address: this.registryAddress, abi: REGISTRY_ABI, functionName: 'INACTIVITY_WINDOW' },
      ])
      if (nodeCount.ok) meta.nodeCount = num(nodeCount.value)
      if (activeCount.ok) meta.activeCount = num(activeCount.value)
      if (window.ok) meta.inactivityWindow = num(window.value)
    }
    return meta
  }

  /** Registry counters only — the cheap part of `readChainMeta`, per poll. */
  async readRegistryCounters(stateBlock: number): Promise<Partial<ChainMeta>> {
    if (!this.registryAddress) return {}
    const [nodeCount, activeCount] = await this.read([
      { address: this.registryAddress, abi: REGISTRY_ABI, functionName: 'nodeCount' },
      { address: this.registryAddress, abi: REGISTRY_ABI, functionName: 'activeCount' },
    ])
    const meta: Partial<ChainMeta> = { stateBlock }
    if (nodeCount.ok) meta.nodeCount = num(nodeCount.value)
    if (activeCount.ok) meta.activeCount = num(activeCount.value)
    return meta
  }

  /**
   * `getEpoch` + `selectedParticipants` (+ `getCollectivePublicKey` once the
   * epoch is Live) for each id, in one batch.
   */
  async readEpochs(ids: EpochId[], stateBlock: number): Promise<Map<string, EpochStateUpdate>> {
    const out = new Map<string, EpochStateUpdate>()
    if (ids.length === 0) return out
    const calls: CallSpec[] = []
    for (const id of ids) {
      calls.push({ address: this.managerAddress, abi: MANAGER_ABI, functionName: 'getEpoch', args: [id] })
      calls.push({
        address: this.managerAddress,
        abi: MANAGER_ABI,
        functionName: 'selectedParticipants',
        args: [id],
      })
      calls.push({
        address: this.managerAddress,
        abi: MANAGER_ABI,
        functionName: 'getCollectivePublicKey',
        args: [id],
      })
    }
    const results = await this.read(calls)
    ids.forEach((id, i) => {
      const epochResult = results[i * 3]
      const participantsResult = results[i * 3 + 1]
      const pkResult = results[i * 3 + 2]
      if (!epochResult?.ok) return
      const raw = epochResult.value as {
        policy: Record<string, unknown>
        status: number
        claimedCount: number
        contributionCount: number
        partialDecryptionCount: number
        ciphertextCount: number
      }
      const policy: EpochPolicy = {
        threshold: num(raw.policy?.threshold),
        committeeSize: num(raw.policy?.committeeSize),
        minValidContributions: num(raw.policy?.minValidContributions),
        lotteryAlphaBps: num(raw.policy?.lotteryAlphaBps),
        committeeSelectionDeadlineBlock: num(raw.policy?.committeeSelectionDeadlineBlock),
        keyAssemblyDeadlineBlock: num(raw.policy?.keyAssemblyDeadlineBlock),
        liveNotBeforeBlock: num(raw.policy?.liveNotBeforeBlock),
      }
      const update: EpochStateUpdate = {
        status: num(raw.status),
        policy,
        claimedCount: num(raw.claimedCount),
        contributionCount: num(raw.contributionCount),
        partialDecryptionCount: num(raw.partialDecryptionCount),
        ciphertextCount: num(raw.ciphertextCount),
        stateBlock,
      }
      if (participantsResult?.ok && Array.isArray(participantsResult.value)) {
        update.committee = (participantsResult.value as string[]).map((a) => a.toLowerCase() as Address)
      }
      if (pkResult?.ok) {
        const point = pkResult.value as { x: bigint; y: bigint }
        const [x, y] = fromRTEtoTE(big(point?.x), big(point?.y))
        update.collectivePublicKey = x === 0n && y === 1n ? null : { x, y }
      }
      out.set(id.toLowerCase(), update)
    })
    return out
  }

  /**
   * `D_i` for every committee member of every given epoch, in one batch.
   *
   * Participant indices are 1-based on chain (`epochParticipants[i - 1]` is
   * the slot that owns index `i`), so the returned arrays are indexed by slot:
   * `result[epochId][slot]` is `getShareCommitmentHash(epochId, slot + 1)`.
   */
  async readShareCommitments(
    epochs: Array<{ id: EpochId; committeeSize: number }>,
  ): Promise<Map<string, (Hex | null)[]>> {
    const out = new Map<string, (Hex | null)[]>()
    const wanted = epochs.filter((epoch) => epoch.committeeSize > 0)
    if (wanted.length === 0) return out
    const calls: CallSpec[] = []
    for (const epoch of wanted) {
      for (let slot = 0; slot < epoch.committeeSize; slot++) {
        calls.push({
          address: this.managerAddress,
          abi: MANAGER_ABI,
          functionName: 'getShareCommitmentHash',
          args: [epoch.id, slot + 1],
        })
      }
    }
    const results = await this.read(calls)
    let cursor = 0
    for (const epoch of wanted) {
      const hashes: (Hex | null)[] = []
      for (let slot = 0; slot < epoch.committeeSize; slot++) {
        const result = results[cursor++]
        hashes.push(result?.ok ? (result.value as Hex) : null)
      }
      out.set(epoch.id.toLowerCase(), hashes)
    }
    return out
  }

  /** Registry records for the given operators. */
  async readOperators(
    operators: Address[],
    stateBlock: number,
  ): Promise<Map<string, OperatorStateUpdate>> {
    const out = new Map<string, OperatorStateUpdate>()
    if (!this.registryAddress || operators.length === 0) return out
    const registry = this.registryAddress
    const results = await this.read(
      operators.map((operator) => ({
        address: registry,
        abi: REGISTRY_ABI,
        functionName: 'getNode',
        args: [operator],
      })),
    )
    operators.forEach((operator, i) => {
      const result = results[i]
      if (!result?.ok) return
      const node = result.value as {
        pubX: bigint
        pubY: bigint
        status: number
        lastActiveBlock: bigint
        registeredAtBlock: bigint
      }
      out.set(operator.toLowerCase(), {
        pubKey: { x: big(node?.pubX), y: big(node?.pubY) },
        status: num(node?.status),
        lastActiveBlock: num(node?.lastActiveBlock),
        registeredAtBlock: num(node?.registeredAtBlock),
        stateBlock,
      })
    })
    return out
  }

  /** Application records for the given `(epoch, aid)` pairs. */
  async readApplications(
    pairs: Array<{ epoch: EpochId; aid: Aid }>,
    stateBlock: number,
  ): Promise<Map<string, ApplicationStateUpdate>> {
    const out = new Map<string, ApplicationStateUpdate>()
    if (!this.appManagerAddress || pairs.length === 0) return out
    const appManager = this.appManagerAddress
    const results = await this.read(
      pairs.map(({ epoch, aid }) => ({
        address: appManager,
        abi: APP_MANAGER_ABI,
        functionName: 'getApplication',
        args: [epoch, aid],
      })),
    )
    pairs.forEach(({ epoch, aid }, i) => {
      const result = results[i]
      if (!result?.ok) return
      const app = result.value as {
        creator: string
        organizerPK: { x: bigint; y: bigint }
        policy: {
          authorizedSubmitter: string
          maxCiphertexts: number
          notBeforeBlock: bigint
          notAfterBlock: bigint
        }
        createdAtBlock: bigint
        exists: boolean
      }
      if (!app?.exists) return
      const [x, y] = fromRTEtoTE(big(app.organizerPK?.x), big(app.organizerPK?.y))
      out.set(`${epoch.toLowerCase()}:${aid.toLowerCase()}`, {
        creator: app.creator.toLowerCase() as Address,
        organizerPK: { x, y },
        policy: {
          authorizedSubmitter: app.policy.authorizedSubmitter.toLowerCase() as Address,
          maxCiphertexts: num(app.policy.maxCiphertexts),
          notBeforeBlock: num(app.policy.notBeforeBlock),
          notAfterBlock: num(app.policy.notAfterBlock),
        },
        createdBlock: num(app.createdAtBlock),
        stateBlock,
      })
    })
    return out
  }

  /**
   * Sender and gas for a set of transactions. `EpochLive` and
   * `DecryptionCombined` name no submitter, so this is how a finalization or a
   * combine gets attributed; the UI also asks for it on any row a user opens.
   */
  async readTxMeta(hashes: Hex[]): Promise<TxMeta[]> {
    const unique = [...new Set(hashes.map((h) => h.toLowerCase()))] as Hex[]
    const out: TxMeta[] = []
    let cursor = 0
    const worker = async (): Promise<void> => {
      for (;;) {
        const i = cursor++
        if (i >= unique.length) return
        const hash = unique[i]
        try {
          this.requests += 1
          const receipt = await this.client.getTransactionReceipt({ hash })
          out.push({
            hash,
            from: (receipt.from?.toLowerCase() ?? null) as Address | null,
            gasUsed: receipt.gasUsed != null ? Number(receipt.gasUsed) : null,
            blockNumber: num(receipt.blockNumber),
            status: receipt.status === 'success' ? 'success' : 'reverted',
          })
        } catch {
          // A pruned or re-orged transaction is simply not attributable.
        }
      }
    }
    await Promise.all(Array.from({ length: Math.min(this.concurrency, unique.length) }, worker))
    return out
  }
}

export function createStateReader(options: StateReaderOptions): StateReader {
  return new StateReader(options)
}
