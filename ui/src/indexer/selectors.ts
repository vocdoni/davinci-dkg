// Pure selectors over the entity store.
//
// Every page reads the store through these and nothing else: they are plain
// functions of `IndexerStore`, so they unit-test against the synthetic fixture
// and memoise on store identity (the indexer hands out a new top-level object
// on every change, so a `WeakMap` keyed by the store is an exact cache).

import {
  applicationKey,
  ciphertextKey,
  epochKey,
  operatorKey,
  txKey,
  type Address,
  type Aid,
  type AppModeName,
  type ApplicationEntity,
  type CiphertextEntity,
  type ContributionEntity,
  type EpochEntity,
  type EpochId,
  type EpochPhaseName,
  type Hex,
  type IndexedEvent,
  type IndexerStore,
  type OperatorEntity,
  type Point,
  type PoolKeyEntity,
  type TxMeta,
} from './types'

// ── memoisation ──────────────────────────────────────────────────────────────

function memoPerStore<T>(compute: (store: IndexerStore) => T): (store: IndexerStore) => T {
  const cache = new WeakMap<IndexerStore, T>()
  return (store) => {
    const hit = cache.get(store)
    if (hit !== undefined) return hit
    const value = compute(store)
    cache.set(store, value)
    return value
  }
}

// ── shared helpers ───────────────────────────────────────────────────────────

/** 2²⁵⁶ — the size of the lottery hash space. */
const HASH_SPACE = 1n << 256n

export function gasOf(store: IndexerStore, tx: Hex | null | undefined): number | null {
  if (!tx) return null
  return store.txMeta[txKey(tx)]?.gasUsed ?? null
}

export function senderOf(store: IndexerStore, tx: Hex | null | undefined): Address | null {
  if (!tx) return null
  return store.txMeta[txKey(tx)]?.from ?? null
}

export function txMetaOf(store: IndexerStore, tx: Hex | null | undefined): TxMeta | null {
  if (!tx) return null
  return store.txMeta[txKey(tx)] ?? null
}

export function getEpoch(store: IndexerStore, id: EpochId | string): EpochEntity | null {
  return store.epochs[epochKey(id)] ?? null
}

export function getOperator(store: IndexerStore, address: Address | string): OperatorEntity | null {
  return store.operators[operatorKey(address)] ?? null
}

export function getApplication(store: IndexerStore, epoch: EpochId, aid: Aid): ApplicationEntity | null {
  return store.applications[applicationKey(epoch, aid)] ?? null
}

export function getCiphertext(
  store: IndexerStore,
  epoch: EpochId,
  aid: Aid,
  index: number,
): CiphertextEntity | null {
  return store.ciphertexts[ciphertextKey(epoch, aid, index)] ?? null
}

/** Epochs in descending nonce order (newest first). */
export const epochsNewestFirst = memoPerStore((store: IndexerStore): EpochEntity[] =>
  store.epochOrder.map((key) => store.epochs[key]).reverse(),
)

export function newestEpoch(store: IndexerStore): EpochEntity | null {
  return epochsNewestFirst(store)[0] ?? null
}

/** Epochs whose key is currently usable. */
export function liveEpochs(store: IndexerStore): EpochEntity[] {
  return epochsNewestFirst(store).filter((epoch) => epoch.status === 'live')
}

/**
 * Addresses the registry knows about. The store also holds addresses that only
 * ever appear as a ciphertext submitter or an application organizer — real
 * facts worth keeping, but not rows in the operators table.
 */
export const registeredOperators = memoPerStore((store: IndexerStore): OperatorEntity[] =>
  store.operatorOrder
    .map((key) => store.operators[key])
    .filter((operator) => operator.status !== 'none' || operator.registeredAtBlock > 0),
)

// ── network stats ────────────────────────────────────────────────────────────

export interface NetworkStats {
  chainId: number
  chainName: string
  managerAddress: Address
  registryAddress: Address | null
  appManagerAddress: Address | null
  explorerUrl: string
  headBlock: number
  lastIndexedBlock: number
  deployBlock: number
  blockTimeSeconds: number
  epochDurationBlocks: number | null
  nextEpochStartBlock: number | null
  /** Blocks until `createEpoch` may fire again; null when unknown. */
  blocksToNextEpoch: number | null
  operatorsRegistered: number
  operatorsActive: number
  operatorsInactive: number
  inactivityWindow: number | null
  epochs: number
  epochsLive: number
  epochsAborted: number
  newestEpoch: EpochEntity | null
  /** Threshold and committee size of the newest epoch that has a policy. */
  thresholdInForce: number | null
  committeeSizeInForce: number | null
  applications: number
  ciphertexts: number
  ciphertextsDecrypted: number
  partials: number
  contributions: number
  claims: number
  events: number
}

export const networkStats = memoPerStore((store: IndexerStore): NetworkStats => {
  const epochs = store.epochOrder.map((key) => store.epochs[key])
  let epochsLive = 0
  let epochsAborted = 0
  let claims = 0
  let contributions = 0
  for (const epoch of epochs) {
    if (epoch.status === 'live') epochsLive += 1
    if (epoch.status === 'aborted') epochsAborted += 1
    claims += epoch.slots.length
    contributions += epoch.contributions.length
  }

  const registered = registeredOperators(store)
  let operatorsActive = 0
  let operatorsInactive = 0
  for (const operator of registered) {
    if (operator.status === 'active') operatorsActive += 1
    else if (operator.status === 'inactive') operatorsInactive += 1
  }

  let ciphertexts = 0
  let decrypted = 0
  let partials = 0
  for (const key of Object.keys(store.ciphertexts)) {
    const ct = store.ciphertexts[key]
    ciphertexts += 1
    partials += ct.partials.length
    if (ct.combined) decrypted += 1
  }

  const withPolicy = epochs.filter((epoch) => epoch.policy != null)
  const policyEpoch = withPolicy.length > 0 ? withPolicy[withPolicy.length - 1] : null
  const head = store.chain.headBlock
  const nextStart = store.chain.nextEpochStartBlock

  return {
    chainId: store.chain.chainId,
    chainName: store.chain.chainName,
    managerAddress: store.chain.managerAddress,
    registryAddress: store.chain.registryAddress,
    appManagerAddress: store.chain.appManagerAddress,
    explorerUrl: store.chain.explorerUrl,
    headBlock: head,
    lastIndexedBlock: store.lastIndexedBlock,
    deployBlock: store.chain.deployBlock,
    blockTimeSeconds: store.chain.blockTimeSeconds,
    epochDurationBlocks: store.chain.epochDurationBlocks,
    nextEpochStartBlock: nextStart,
    blocksToNextEpoch: nextStart != null ? Math.max(0, nextStart - head) : null,
    operatorsRegistered: store.chain.nodeCount ?? registered.length,
    operatorsActive: store.chain.activeCount ?? operatorsActive,
    operatorsInactive,
    inactivityWindow: store.chain.inactivityWindow,
    epochs: epochs.length,
    epochsLive,
    epochsAborted,
    newestEpoch: epochs.length > 0 ? epochs[epochs.length - 1] : null,
    thresholdInForce: policyEpoch?.policy?.threshold ?? null,
    committeeSizeInForce: policyEpoch?.policy?.committeeSize ?? null,
    applications: store.applicationOrder.length,
    ciphertexts,
    ciphertextsDecrypted: decrypted,
    partials,
    contributions,
    claims,
    events: store.events.length,
  }
})

// ── epoch list ───────────────────────────────────────────────────────────────

export interface EpochRow {
  id: EpochId
  nonce: number
  phase: EpochPhaseName
  threshold: number
  committeeSize: number
  minValidContributions: number
  claims: number
  contributions: number
  ciphertexts: number
  decrypted: number
  applications: number
  createdBlock: number
  startBlock: number
  liveSinceBlock: number | null
  endBlock: number | null
  creator: Address
  finalizer: Address | null
  finalizationTx: Hex | null
  finalizationGas: number | null
  /** Pool keys activated / claimed so far, out of `POOL_SIZE`. */
  poolActivated: number
  poolClaimed: number
  /** 0…1 over the committee. */
  claimProgress: number
  contributionProgress: number
}

export interface EpochFilter {
  phase?: EpochPhaseName | 'all'
  /** Substring match against the epoch id, nonce or creator. */
  query?: string
  limit?: number
}

function epochRow(store: IndexerStore, epoch: EpochEntity): EpochRow {
  const n = epoch.policy?.committeeSize ?? epoch.committee.length
  const claims = epoch.slots.length
  const contributions = epoch.contributions.length
  let ciphertexts = 0
  let decrypted = 0
  for (const appKey of epoch.applications) {
    const app = store.applications[appKey]
    for (const ctKey of app.ciphertexts) {
      ciphertexts += 1
      if (store.ciphertexts[ctKey].combined) decrypted += 1
    }
  }
  const duration = store.chain.epochDurationBlocks
  return {
    id: epoch.id,
    nonce: epoch.nonce,
    phase: epoch.status,
    threshold: epoch.policy?.threshold ?? 0,
    committeeSize: n,
    minValidContributions: epoch.policy?.minValidContributions ?? 0,
    claims,
    contributions,
    ciphertexts,
    decrypted,
    applications: epoch.applications.length,
    createdBlock: epoch.createdBlock,
    startBlock: epoch.startBlock,
    liveSinceBlock: epoch.finalization?.block ?? null,
    endBlock: duration != null ? epoch.startBlock + duration : null,
    creator: epoch.creator,
    finalizer: epoch.finalization?.by ?? senderOf(store, epoch.finalization?.tx),
    finalizationTx: epoch.finalization?.tx ?? null,
    finalizationGas: gasOf(store, epoch.finalization?.tx),
    poolActivated: epoch.poolKeys.filter((slot) => slot.key != null).length,
    poolClaimed: epoch.poolKeys.filter((slot) => slot.claimedBy != null).length,
    claimProgress: n > 0 ? Math.min(1, claims / n) : 0,
    contributionProgress: n > 0 ? Math.min(1, contributions / n) : 0,
  }
}

export const allEpochRows = memoPerStore((store: IndexerStore): EpochRow[] =>
  epochsNewestFirst(store).map((epoch) => epochRow(store, epoch)),
)

export function epochRows(store: IndexerStore, filter: EpochFilter = {}): EpochRow[] {
  let rows = allEpochRows(store)
  if (filter.phase && filter.phase !== 'all') rows = rows.filter((row) => row.phase === filter.phase)
  const query = filter.query?.trim().toLowerCase()
  if (query) {
    // Nonce is matched exactly and the creator by prefix: a substring match on
    // an address turns any single hex digit into "everything".
    rows = rows.filter(
      (row) =>
        String(row.nonce) === query ||
        row.id.toLowerCase().includes(query) ||
        row.creator.toLowerCase().startsWith(query) ||
        (row.finalizer?.toLowerCase().startsWith(query) ?? false),
    )
  }
  if (filter.limit != null) rows = rows.slice(0, filter.limit)
  return rows
}

// ── epoch detail ─────────────────────────────────────────────────────────────

export interface LotteryInfo {
  seed: Hex | null
  seedBlock: number
  seedResolvedBlock: number | null
  /** τ, the raw eligibility threshold. */
  threshold: bigint
  /** τ / 2²⁵⁶ — the share of the hash space that wins. */
  thresholdFraction: number
  alphaBps: number
  alpha: number
  /** Registry size snapshotted at `createEpoch`, when it can be recovered. */
  registrySnapshot: number | null
  /** min(1, α·n/R): the chance a given registered operator is admissible. */
  admissibleProbability: number | null
  claims: Array<{ slot: number; operator: Address; block: number; tx: Hex | null }>
}

export interface CommitteeRow {
  /** 0-based position in `selectedParticipants`. */
  slot: number
  /**
   * 1-based index the protocol uses for this member: `contributorIndex`,
   * `participantIndex` and `getShareCommitmentHash` are all `slot + 1`.
   */
  participantIndex: number
  operator: Address
  claimBlock: number | null
  claimTx: Hex | null
  contributed: boolean
  contributionBlock: number | null
  contributionTx: Hex | null
  contributionGas: number | null
  commitmentsHash: Hex | null
  encryptedSharesHash: Hex | null
  partials: number
}

export interface EpochWindows {
  createdBlock: number
  startBlock: number
  seedBlock: number
  committeeSelectionDeadline: number | null
  keyAssemblyDeadline: number | null
  liveNotBefore: number | null
  endBlock: number | null
}

export interface ApplicationRow {
  key: string
  epoch: EpochId
  aid: Aid
  creator: Address
  /** `PK_org`, TE form; the identity `(0, 1)` when automatic. */
  organizerPK: Point
  mode: AppModeName
  /** Pool key claimed at registration; null until seen. */
  poolIndex: number | null
  /** `P_j`, TE form; null until the key's activation has been indexed. */
  poolKey: Point | null
  /** Policy fields are null until the record has been read on chain. */
  openSubmission: boolean | null
  /** Allow-list; empty means the registrant only. */
  submitters: Address[] | null
  maxCiphertexts: number | null
  notBeforeBlock: number | null
  notAfterBlock: number | null
  /** Decryption window, unix seconds; 0 = unbounded on that side. */
  decryptNotBefore: number | null
  decryptNotAfter: number | null
  /**
   * Whether the committee can combine on its own: always for an automatic
   * application, and once `sk_org` is revealed for an organizer-locked one.
   */
  unlocked: boolean
  /** `sk_org` after the reveal; null before it and for automatic applications. */
  organizerSecret: bigint | null
  revealBlock: number | null
  revealTx: Hex | null
  createdBlock: number
  createdTx: Hex | null
  ciphertexts: number
  decrypted: number
}

export type PoolSlotState = 'inactive' | 'activated' | 'claimed'

/** One of the `POOL_SIZE` keys of an epoch, as the pool panel draws it. */
export interface PoolSlotRow extends PoolKeyEntity {
  state: PoolSlotState
  activatedBy: Address | null
  activatedGas: number | null
}

export interface EpochDetail {
  epoch: EpochEntity
  row: EpochRow
  lottery: LotteryInfo
  committee: CommitteeRow[]
  windows: EpochWindows
  applications: ApplicationRow[]
  finalization: {
    by: Address | null
    block: number
    tx: Hex | null
    gasUsed: number | null
    contributionCount: number
  } | null
  /** The pool, by key index. */
  pool: PoolSlotRow[]
  poolNext: number
  poolActivated: number
  poolClaimed: number
  events: IndexedEvent[]
  /** Contributions in submission order. */
  contributions: CommitteeRow[]
}

/** `P_j` of the key `app` claimed, once its activation is in the store. */
export function poolKeyOf(store: IndexerStore, app: ApplicationEntity): Point | null {
  if (app.poolIndex == null) return null
  return store.epochs[epochKey(app.epoch)]?.poolKeys[app.poolIndex]?.key ?? null
}

/** True once the committee alone can combine this application's ciphertexts. */
export function isUnlocked(app: ApplicationEntity): boolean {
  return app.mode === 'automatic' || app.organizerSecret != null
}

function applicationRow(store: IndexerStore, app: ApplicationEntity): ApplicationRow {
  let decrypted = 0
  for (const key of app.ciphertexts) {
    if (store.ciphertexts[key].combined) decrypted += 1
  }
  return {
    key: app.key,
    epoch: app.epoch,
    aid: app.aid,
    creator: app.creator,
    organizerPK: app.organizerPK,
    mode: app.mode,
    poolIndex: app.poolIndex,
    poolKey: poolKeyOf(store, app),
    openSubmission: app.policy?.openSubmission ?? null,
    submitters: app.policy?.submitters ?? null,
    maxCiphertexts: app.policy?.maxCiphertexts ?? null,
    notBeforeBlock: app.policy?.notBeforeBlock ?? null,
    notAfterBlock: app.policy?.notAfterBlock ?? null,
    decryptNotBefore: app.policy?.decryptNotBefore ?? null,
    decryptNotAfter: app.policy?.decryptNotAfter ?? null,
    unlocked: isUnlocked(app),
    organizerSecret: app.organizerSecret,
    revealBlock: app.organizerReveal?.block ?? null,
    revealTx: app.organizerReveal?.tx ?? null,
    createdBlock: app.createdBlock,
    createdTx: app.createdTx,
    ciphertexts: app.ciphertexts.length,
    decrypted,
  }
}

function poolSlotRow(store: IndexerStore, slot: PoolKeyEntity): PoolSlotRow {
  return {
    ...slot,
    state: slot.claimedBy != null ? 'claimed' : slot.key != null ? 'activated' : 'inactive',
    activatedBy: senderOf(store, slot.activatedTx),
    activatedGas: gasOf(store, slot.activatedTx),
  }
}

export function epochDetail(store: IndexerStore, id: EpochId | string): EpochDetail | null {
  const epoch = getEpoch(store, id)
  if (!epoch) return null

  const n = epoch.policy?.committeeSize ?? epoch.committee.length
  const alpha = (epoch.policy?.lotteryAlphaBps ?? 0) / 10_000
  const registrySnapshot = epoch.registrySnapshot ?? recoverRegistrySnapshot(epoch, n)
  const lottery: LotteryInfo = {
    seed: epoch.seed,
    seedBlock: epoch.seedBlock,
    seedResolvedBlock: epoch.seedResolvedBlock,
    threshold: epoch.lotteryThreshold,
    thresholdFraction: fractionOfHashSpace(epoch.lotteryThreshold),
    alphaBps: epoch.policy?.lotteryAlphaBps ?? 0,
    alpha,
    registrySnapshot,
    admissibleProbability:
      registrySnapshot && registrySnapshot > 0 ? Math.min(1, (alpha * n) / registrySnapshot) : null,
    claims: epoch.slots.map((key) => {
      const slot = store.slots[key]
      return { slot: slot.slot, operator: slot.operator, block: slot.block, tx: slot.tx }
    }),
  }

  // Slots are 0-based; contributorIndex / participantIndex are 1-based
  // (`DKGManager` checks `epochParticipants[contributorIndex - 1] == sender`),
  // so every join below goes through `slot = index - 1`.
  const contributionBySlot = new Map<number, ContributionEntity>()
  const contributionByOperator = new Map<string, ContributionEntity>()
  for (const key of epoch.contributions) {
    const contribution = store.contributions[key]
    contributionBySlot.set(participantIndexToSlot(contribution.index), contribution)
    contributionByOperator.set(operatorKey(contribution.contributor), contribution)
  }

  const partialsBySlot = new Map<number, number>()
  for (const appKey of epoch.applications) {
    for (const ctKey of store.applications[appKey].ciphertexts) {
      for (const partial of store.ciphertexts[ctKey].partials) {
        const slot = participantIndexToSlot(partial.participantIndex)
        partialsBySlot.set(slot, (partialsBySlot.get(slot) ?? 0) + 1)
      }
    }
  }

  const size = Math.max(n, epoch.committee.length)
  const committee: CommitteeRow[] = []
  for (let i = 0; i < size; i++) {
    const operator = epoch.committee[i] ?? null
    const slot = store.slots[`${epoch.id}:${i}`] ?? null
    const contribution =
      contributionBySlot.get(i) ?? (operator ? contributionByOperator.get(operatorKey(operator)) : undefined)
    committee.push({
      slot: i,
      participantIndex: i + 1,
      operator: (operator ?? slot?.operator ?? '0x0000000000000000000000000000000000000000') as Address,
      claimBlock: slot?.block ?? null,
      claimTx: slot?.tx ?? null,
      contributed: contribution != null,
      contributionBlock: contribution?.block ?? null,
      contributionTx: contribution?.tx ?? null,
      contributionGas: gasOf(store, contribution?.tx),
      commitmentsHash: contribution?.commitmentsHash ?? null,
      encryptedSharesHash: contribution?.encryptedSharesHash ?? null,
      partials: partialsBySlot.get(i) ?? 0,
    })
  }

  const duration = store.chain.epochDurationBlocks
  const windows: EpochWindows = {
    createdBlock: epoch.createdBlock,
    startBlock: epoch.startBlock,
    seedBlock: epoch.seedBlock,
    committeeSelectionDeadline: epoch.policy?.committeeSelectionDeadlineBlock ?? null,
    keyAssemblyDeadline: epoch.policy?.keyAssemblyDeadlineBlock ?? null,
    liveNotBefore: epoch.policy?.liveNotBeforeBlock ?? null,
    endBlock: duration != null ? epoch.startBlock + duration : null,
  }

  const pool = epoch.poolKeys.map((slot) => poolSlotRow(store, slot))

  return {
    epoch,
    row: epochRow(store, epoch),
    lottery,
    committee,
    windows,
    applications: epoch.applications.map((key) => applicationRow(store, store.applications[key])),
    finalization: epoch.finalization
      ? {
          by: epoch.finalization.by ?? senderOf(store, epoch.finalization.tx),
          block: epoch.finalization.block,
          tx: epoch.finalization.tx,
          gasUsed: gasOf(store, epoch.finalization.tx),
          contributionCount: epoch.finalization.contributionCount,
        }
      : null,
    pool,
    poolNext: epoch.poolNext,
    poolActivated: pool.filter((slot) => slot.key != null).length,
    poolClaimed: pool.filter((slot) => slot.claimedBy != null).length,
    events: epoch.events.map((i) => store.events[i]),
    contributions: epoch.contributions
      .map((key) => committee[participantIndexToSlot(store.contributions[key].index)])
      .filter((row): row is CommitteeRow => row != null),
  }
}

/**
 * `slot = participantIndex - 1`. The lottery numbers committee slots from 0,
 * while every proof-carrying call numbers participants from 1.
 */
export function participantIndexToSlot(participantIndex: number): number {
  return Math.max(0, participantIndex - 1)
}

/** τ as a fraction of 2²⁵⁶, with enough precision for a percentage. */
export function fractionOfHashSpace(threshold: bigint): number {
  if (threshold <= 0n) return 0
  if (threshold >= HASH_SPACE) return 1
  const scale = 1_000_000_000n
  return Number((threshold * scale) / HASH_SPACE) / Number(scale)
}

/**
 * The lottery threshold is `τ = α·n·2²⁵⁶ / R`, so the registry size snapshotted
 * at `createEpoch` can be read back out of it — useful because no event
 * carries `R`.
 */
export function recoverRegistrySnapshot(epoch: EpochEntity, committeeSize: number): number | null {
  const alphaBps = epoch.policy?.lotteryAlphaBps
  if (!alphaBps || epoch.lotteryThreshold <= 0n || committeeSize <= 0) return null
  const numerator = (BigInt(alphaBps) * BigInt(committeeSize) * HASH_SPACE) / 10_000n
  const r = numerator / epoch.lotteryThreshold
  const value = Number(r)
  return Number.isFinite(value) && value > 0 ? value : null
}

// ── operators ────────────────────────────────────────────────────────────────

export interface OperatorEpochHistory {
  epoch: EpochId
  nonce: number
  claimed: boolean
  slot: number | null
  contributed: boolean
  contributionBlock: number | null
  finalized: boolean
  partials: number
  combines: number
}

export interface OperatorAggregate {
  claims: number
  contributions: number
  partials: number
  finalizations: number
  combines: number
  ciphertextsSubmitted: number
  applicationsRegistered: number
  epochsServed: number
  perEpoch: Map<string, OperatorEpochHistory>
}

function emptyAggregate(): OperatorAggregate {
  return {
    claims: 0,
    contributions: 0,
    partials: 0,
    finalizations: 0,
    combines: 0,
    ciphertextsSubmitted: 0,
    applicationsRegistered: 0,
    epochsServed: 0,
    perEpoch: new Map(),
  }
}

function historyFor(
  store: IndexerStore,
  aggregate: OperatorAggregate,
  epochId: string,
): OperatorEpochHistory {
  let entry = aggregate.perEpoch.get(epochId)
  if (!entry) {
    entry = {
      epoch: epochId as EpochId,
      nonce: store.epochs[epochId]?.nonce ?? 0,
      claimed: false,
      slot: null,
      contributed: false,
      contributionBlock: null,
      finalized: false,
      partials: 0,
      combines: 0,
    }
    aggregate.perEpoch.set(epochId, entry)
  }
  return entry
}

/** One pass over every event: per-operator counters and per-epoch history. */
export const operatorAggregates = memoPerStore(
  (store: IndexerStore): Map<string, OperatorAggregate> => {
    const map = new Map<string, OperatorAggregate>()
    const get = (address: string): OperatorAggregate => {
      const key = operatorKey(address)
      let aggregate = map.get(key)
      if (!aggregate) {
        aggregate = emptyAggregate()
        map.set(key, aggregate)
      }
      return aggregate
    }
    for (const key of store.operatorOrder) get(key)

    for (const event of store.events) {
      switch (event.name) {
        case 'SlotClaimed': {
          const aggregate = get(event.data.claimer)
          aggregate.claims += 1
          const history = historyFor(store, aggregate, epochKey(event.data.epochId))
          history.claimed = true
          history.slot = event.data.slot
          break
        }
        case 'ContributionSubmitted': {
          const aggregate = get(event.data.contributor)
          aggregate.contributions += 1
          const history = historyFor(store, aggregate, epochKey(event.data.epochId))
          history.contributed = true
          history.contributionBlock = event.block
          break
        }
        case 'PartialDecryptionSubmitted': {
          const aggregate = get(event.data.participant)
          aggregate.partials += 1
          historyFor(store, aggregate, epochKey(event.data.epochId)).partials += 1
          break
        }
        case 'CiphertextSubmitted':
          get(event.data.submitter).ciphertextsSubmitted += 1
          break
        case 'ApplicationRegistered':
          get(event.data.creator).applicationsRegistered += 1
          break
        default:
          break
      }
    }

    // Finalizations and combines name nobody in their event: they are
    // attributed through the transaction sender, once it has been resolved.
    for (const epochId of store.epochOrder) {
      const epoch = store.epochs[epochId]
      const finalizer = epoch.finalization?.by ?? senderOf(store, epoch.finalization?.tx)
      if (finalizer) {
        const aggregate = get(finalizer)
        aggregate.finalizations += 1
        historyFor(store, aggregate, epochId).finalized = true
      }
    }
    for (const key of Object.keys(store.ciphertexts)) {
      const ct = store.ciphertexts[key]
      const by = ct.combined?.by ?? senderOf(store, ct.combined?.tx)
      if (by) {
        const aggregate = get(by)
        aggregate.combines += 1
        historyFor(store, aggregate, epochKey(ct.epoch)).combines += 1
      }
    }

    for (const aggregate of map.values()) {
      let served = 0
      for (const history of aggregate.perEpoch.values()) {
        if (history.claimed || history.contributed) served += 1
      }
      aggregate.epochsServed = served
    }
    return map
  },
)

export interface OperatorRow {
  address: Address
  pubKey: Point | null
  status: 'none' | 'active' | 'inactive'
  registeredAtBlock: number
  lastActiveBlock: number
  epochsServed: number
  claims: number
  contributions: number
  partials: number
  finalizations: number
  combines: number
  /** contributions / claims, or null when the operator never claimed ("—"). */
  participation: number | null
  /** Blocks since `lastActiveBlock`; null when never active. */
  idleBlocks: number | null
  /** True when `lastActive + INACTIVITY_WINDOW` has passed. */
  reapable: boolean
}

function buildOperatorRow(
  store: IndexerStore,
  operator: OperatorEntity,
  aggregate: OperatorAggregate,
): OperatorRow {
  const head = store.chain.headBlock
  const window = store.chain.inactivityWindow
  const idle = operator.lastActiveBlock > 0 ? Math.max(0, head - operator.lastActiveBlock) : null
  return {
    address: operator.address,
    pubKey: operator.pubKey,
    status: operator.status,
    registeredAtBlock: operator.registeredAtBlock,
    lastActiveBlock: operator.lastActiveBlock,
    epochsServed: aggregate.epochsServed,
    claims: aggregate.claims,
    contributions: aggregate.contributions,
    partials: aggregate.partials,
    finalizations: aggregate.finalizations,
    combines: aggregate.combines,
    participation: aggregate.claims > 0 ? aggregate.contributions / aggregate.claims : null,
    idleBlocks: idle,
    reapable: operator.status === 'active' && window != null && idle != null ? idle > window : false,
  }
}

export const operatorRows = memoPerStore((store: IndexerStore): OperatorRow[] => {
  const aggregates = operatorAggregates(store)
  return registeredOperators(store).map((operator) =>
    buildOperatorRow(store, operator, aggregates.get(operator.address) ?? emptyAggregate()),
  )
})

/** `"—"` when the operator never claimed a slot, a percentage otherwise. */
export function formatParticipation(participation: number | null, digits = 0): string {
  if (participation == null) return '—'
  return `${(participation * 100).toFixed(digits)}%`
}

export interface OperatorDetail {
  operator: OperatorEntity
  row: OperatorRow
  history: OperatorEpochHistory[]
  events: IndexedEvent[]
}

export function operatorDetail(store: IndexerStore, address: Address | string): OperatorDetail | null {
  const operator = getOperator(store, address)
  if (!operator) return null
  const key = operatorKey(address)
  const aggregate = operatorAggregates(store).get(key) ?? emptyAggregate()
  const row = buildOperatorRow(store, operator, aggregate)
  const history = [...aggregate.perEpoch.values()].sort((a, b) => b.nonce - a.nonce)
  return {
    operator,
    row,
    history,
    events: operator.events.map((i) => store.events[i]),
  }
}

// ── applications ─────────────────────────────────────────────────────────────

export const applicationRows = memoPerStore((store: IndexerStore): ApplicationRow[] =>
  store.applicationOrder
    .map((key) => applicationRow(store, store.applications[key]))
    .sort((a, b) => b.createdBlock - a.createdBlock),
)

/**
 * Where a ciphertext is in the pipeline. `awaiting-reveal`: the application
 * is organizer-locked and `sk_org` is not out yet — the contract refuses every
 * partial and combine until it is, so nothing else can have happened;
 * `partials`: some in, below `t`; `ready`: `t` partials and nothing else
 * stands in the way of a combine.
 */
export type CiphertextState = 'submitted' | 'partials' | 'awaiting-reveal' | 'ready' | 'combined'

/**
 * Block from which the contract accepts partials for a ciphertext: its own
 * block, or — organizer-locked — the reveal block when that came later. Waves
 * are counted from here, since that is when the committee could start.
 */
function decryptionOpensAt(ct: CiphertextEntity, app: ApplicationEntity | undefined): number {
  const reveal = app?.mode === 'organizer-locked' ? app.organizerReveal?.block : null
  return reveal != null ? Math.max(ct.block, reveal) : ct.block
}

export interface PartialRow {
  /** 1-based, as it appears on chain. */
  participantIndex: number
  /** 0-based committee slot the participant occupies. */
  slot: number
  participant: Address
  block: number
  tx: Hex | null
  /** Decryption wave the partial landed in (see `partialMatrix`). */
  wave: number
}

export interface CiphertextRow {
  key: string
  epoch: EpochId
  aid: Aid
  index: number
  submitter: Address
  block: number
  tx: Hex | null
  c1: Point
  c2: Point
  threshold: number
  committeeSize: number
  partials: PartialRow[]
  partialCount: number
  combined: {
    done: boolean
    by: Address | null
    block: number | null
    tx: Hex | null
    gasUsed: number | null
    plaintext: bigint | null
  }
  state: CiphertextState
}

function ciphertextRow(store: IndexerStore, ct: CiphertextEntity): CiphertextRow {
  const epoch = store.epochs[epochKey(ct.epoch)]
  const threshold = epoch?.policy?.threshold ?? 0
  const committeeSize = epoch?.policy?.committeeSize ?? epoch?.committee.length ?? 0
  const stagger = store.chain.staggerBlocks || 1
  const app = store.applications[applicationKey(ct.epoch, ct.aid)]
  const opensAt = decryptionOpensAt(ct, app)
  const partials: PartialRow[] = ct.partials.map((partial) => ({
    participantIndex: partial.participantIndex,
    slot: participantIndexToSlot(partial.participantIndex),
    participant: partial.participant,
    block: partial.block,
    tx: partial.tx,
    wave: Math.max(0, Math.floor((partial.block - opensAt) / stagger)),
  }))
  const unlocked = app ? isUnlocked(app) : true
  const thresholdMet = threshold > 0 && partials.length >= threshold
  let state: CiphertextState = 'submitted'
  if (ct.combined) state = 'combined'
  // A locked application has no partials before the reveal — the contract
  // refuses them — so every one of its ciphertexts waits there until sk_org
  // is on chain, whatever the partial count says.
  else if (!unlocked) state = 'awaiting-reveal'
  else if (thresholdMet) state = 'ready'
  else if (partials.length > 0) state = 'partials'

  return {
    key: ct.key,
    epoch: ct.epoch,
    aid: ct.aid,
    index: ct.index,
    submitter: ct.submitter,
    block: ct.block,
    tx: ct.tx,
    c1: ct.c1,
    c2: ct.c2,
    threshold,
    committeeSize,
    partials,
    partialCount: partials.length,
    combined: {
      done: ct.combined != null,
      by: ct.combined?.by ?? senderOf(store, ct.combined?.tx),
      block: ct.combined?.block ?? null,
      tx: ct.combined?.tx ?? null,
      gasUsed: gasOf(store, ct.combined?.tx),
      plaintext: ct.combined?.plaintext ?? null,
    },
    state,
  }
}

export interface ApplicationDetail {
  application: ApplicationEntity
  row: ApplicationRow
  epoch: EpochEntity | null
  ciphertexts: CiphertextRow[]
  events: IndexedEvent[]
  summary: {
    total: number
    combined: number
    thresholdMet: number
    /** Ciphertexts parked at `awaiting-reveal`. */
    awaitingReveal: number
  }
}

export function applicationDetail(
  store: IndexerStore,
  epoch: EpochId | string,
  aid: Aid | string,
): ApplicationDetail | null {
  const app = store.applications[`${epochKey(epoch)}:${String(aid).toLowerCase()}`]
  if (!app) return null
  const ciphertexts = app.ciphertexts.map((key) => ciphertextRow(store, store.ciphertexts[key]))
  return {
    application: app,
    row: applicationRow(store, app),
    epoch: getEpoch(store, app.epoch),
    ciphertexts,
    events: app.events.map((i) => store.events[i]),
    summary: {
      total: ciphertexts.length,
      combined: ciphertexts.filter((row) => row.combined.done).length,
      thresholdMet: ciphertexts.filter((row) => row.threshold > 0 && row.partialCount >= row.threshold).length,
      awaitingReveal: ciphertexts.filter((row) => row.state === 'awaiting-reveal').length,
    },
  }
}

// ── activity ─────────────────────────────────────────────────────────────────

export interface ActivityBucket {
  epoch: EpochId
  nonce: number
  phase: EpochPhaseName
  startBlock: number
  claims: number
  contributions: number
  ciphertexts: number
  partials: number
  combines: number
  applications: number
}

/** Per-epoch activity for the newest `count` epochs, oldest first. */
export function activity(store: IndexerStore, count = 30): ActivityBucket[] {
  const epochs = epochsNewestFirst(store).slice(0, count).reverse()
  return epochs.map((epoch) => {
    let ciphertexts = 0
    let partials = 0
    let combines = 0
    for (const appKey of epoch.applications) {
      for (const ctKey of store.applications[appKey].ciphertexts) {
        const ct = store.ciphertexts[ctKey]
        ciphertexts += 1
        partials += ct.partials.length
        if (ct.combined) combines += 1
      }
    }
    return {
      epoch: epoch.id,
      nonce: epoch.nonce,
      phase: epoch.status,
      startBlock: epoch.startBlock,
      claims: epoch.slots.length,
      contributions: epoch.contributions.length,
      ciphertexts,
      partials,
      combines,
      applications: epoch.applications.length,
    }
  })
}

// ── partial matrix ───────────────────────────────────────────────────────────

export interface MatrixCell {
  /** 1-based participant index, as emitted. */
  participantIndex: number
  /** 0-based committee slot — the matrix row. */
  slot: number
  column: number
  aid: Aid
  ciphertextIndex: number
  block: number | null
  wave: number | null
  tx: Hex | null
}

export interface MatrixColumn {
  column: number
  aid: Aid
  ciphertextIndex: number
  submitBlock: number
  partials: number
  threshold: number
  combined: boolean
}

export interface PartialMatrix {
  epoch: EpochId
  threshold: number
  staggerBlocks: number
  /** Committee members in slot order. */
  rows: Array<{ slot: number; participantIndex: number; operator: Address; partials: number }>
  columns: MatrixColumn[]
  /** `cells[slot][column]`; null where no partial was published. */
  cells: Array<Array<MatrixCell | null>>
}

/**
 * Members × ciphertexts heat-map data. `wave` is the decryption round a
 * partial landed in, counted from 0: the node schedules its `i`-th attempt
 * `i · staggerBlocks` after decryption opened (the ciphertext, or the reveal
 * for an organizer-locked application — see `decryptionOpensAt`), so
 * `floor((block − opened) / staggerBlocks)` recovers it.
 */
export function partialMatrix(
  store: IndexerStore,
  epochId: EpochId | string,
  aid?: Aid | string,
): PartialMatrix | null {
  const epoch = getEpoch(store, epochId)
  if (!epoch) return null
  const stagger = store.chain.staggerBlocks || 1
  const size = epoch.policy?.committeeSize ?? epoch.committee.length
  const wanted = aid ? String(aid).toLowerCase() : null

  const columns: MatrixColumn[] = []
  const ciphertexts: CiphertextEntity[] = []
  for (const appKey of epoch.applications) {
    const app = store.applications[appKey]
    if (wanted && app.aid !== wanted) continue
    for (const ctKey of app.ciphertexts) ciphertexts.push(store.ciphertexts[ctKey])
  }
  ciphertexts.sort((a, b) => a.block - b.block || a.index - b.index)

  const cells: Array<Array<MatrixCell | null>> = Array.from({ length: size }, () =>
    new Array<MatrixCell | null>(ciphertexts.length).fill(null),
  )
  const rowTotals = new Array<number>(size).fill(0)

  ciphertexts.forEach((ct, column) => {
    columns.push({
      column,
      aid: ct.aid,
      ciphertextIndex: ct.index,
      submitBlock: ct.block,
      partials: ct.partials.length,
      threshold: epoch.policy?.threshold ?? 0,
      combined: ct.combined != null,
    })
    const opensAt = decryptionOpensAt(ct, store.applications[applicationKey(ct.epoch, ct.aid)])
    for (const partial of ct.partials) {
      const row = participantIndexToSlot(partial.participantIndex)
      if (row < 0 || row >= size) continue
      cells[row][column] = {
        participantIndex: partial.participantIndex,
        slot: row,
        column,
        aid: ct.aid,
        ciphertextIndex: ct.index,
        block: partial.block,
        wave: Math.max(0, Math.floor((partial.block - opensAt) / stagger)),
        tx: partial.tx,
      }
      rowTotals[row] += 1
    }
  })

  return {
    epoch: epoch.id,
    threshold: epoch.policy?.threshold ?? 0,
    staggerBlocks: stagger,
    rows: Array.from({ length: size }, (_, i) => ({
      slot: i,
      participantIndex: i + 1,
      operator: (epoch.committee[i] ?? '0x0000000000000000000000000000000000000000') as Address,
      partials: rowTotals[i],
    })),
    columns,
    cells,
  }
}

// ── event feed ───────────────────────────────────────────────────────────────

export interface FeedEntry {
  index: number
  name: IndexedEvent['name']
  block: number
  tx: Hex | null
  epoch: EpochId | null
  aid: Aid | null
  actor: Address | null
  event: IndexedEvent
}

/** The newest `count` events, newest first. */
export function eventFeed(store: IndexerStore, count = 20): FeedEntry[] {
  const out: FeedEntry[] = []
  for (let i = store.events.length - 1; i >= 0 && out.length < count; i--) {
    const event = store.events[i]
    out.push({
      index: i,
      name: event.name,
      block: event.block,
      tx: event.tx,
      epoch: event.epoch,
      aid: event.aid,
      actor: event.actor,
      event,
    })
  }
  return out
}

// ── search ───────────────────────────────────────────────────────────────────

export type SearchKind = 'epoch' | 'operator' | 'application' | 'transaction'

export interface SearchResult {
  kind: SearchKind
  /** Primary identifier, as it should be displayed. */
  id: string
  label: string
  detail: string
  /**
   * Where to send the user. In-app for epochs, operators and applications;
   * the configured block explorer for a transaction (the explorer has no
   * transaction page of its own), in which case `external` is true.
   */
  href: string
  external: boolean
}

const HEX = /^0x[0-9a-f]*$/

/** Block-explorer link for a transaction; empty when none is configured. */
export function txHref(store: IndexerStore, hash: string): string {
  const base = store.chain.explorerUrl.replace(/\/+$/, '')
  return base ? `${base}/tx/${hash}` : ''
}

/**
 * Resolve a free-text query against the store: epoch id or nonce, application
 * id, operator address, transaction hash, and prefixes of any of them.
 */
export function searchStore(store: IndexerStore, query: string, limit = 10): SearchResult[] {
  const q = query.trim().toLowerCase()
  if (q.length === 0) return []
  const results: SearchResult[] = []
  const push = (result: SearchResult): void => {
    if (results.length < limit && !results.some((r) => r.kind === result.kind && r.id === result.id)) {
      results.push(result)
    }
  }

  const epochResult = (epoch: EpochEntity): SearchResult => ({
    kind: 'epoch',
    id: epoch.id,
    label: `Epoch #${epoch.nonce}`,
    detail: `${epoch.id} · ${epoch.status}`,
    href: `/epochs/${epoch.id}`,
    external: false,
  })

  // Bare number → epoch nonce.
  if (/^\d+$/.test(q)) {
    const nonce = Number(q)
    for (const key of store.epochOrder) {
      const epoch = store.epochs[key]
      if (epoch.nonce === nonce) push(epochResult(epoch))
    }
  }

  if (HEX.test(q)) {
    for (const key of store.epochOrder) {
      if (key.startsWith(q)) push(epochResult(store.epochs[key]))
    }
    for (const key of store.operatorOrder) {
      if (!key.startsWith(q)) continue
      const operator = store.operators[key]
      push({
        kind: 'operator',
        id: operator.address,
        label: operator.address,
        detail: `operator · ${operator.status}`,
        href: `/operators/${operator.address}`,
        external: false,
      })
    }
    for (const key of store.applicationOrder) {
      const app = store.applications[key]
      if (!app.aid.startsWith(q)) continue
      push({
        kind: 'application',
        id: app.aid,
        label: `Application ${app.aid.slice(0, 10)}…`,
        detail: `epoch #${store.epochs[epochKey(app.epoch)]?.nonce ?? '?'} · ${app.ciphertexts.length} ciphertexts`,
        href: `/applications/${app.epoch}/${app.aid}`,
        external: false,
      })
    }
    // Transactions: the cache first (cheap), then the event log.
    for (const hash of Object.keys(store.txMeta)) {
      if (!hash.startsWith(q)) continue
      const meta = store.txMeta[hash]
      push({
        kind: 'transaction',
        id: hash,
        label: hash,
        detail: `block ${meta.blockNumber}`,
        href: txHref(store, hash),
        external: true,
      })
    }
    if (q.length >= 6) {
      for (let i = store.events.length - 1; i >= 0 && results.length < limit; i--) {
        const event = store.events[i]
        if (!event.tx || !event.tx.toLowerCase().startsWith(q)) continue
        push({
          kind: 'transaction',
          id: event.tx.toLowerCase(),
          label: event.tx,
          detail: `${event.name} · block ${event.block}`,
          href: txHref(store, event.tx.toLowerCase()),
          external: true,
        })
      }
    }
  }

  return results
}
