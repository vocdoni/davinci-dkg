// Pure reducers: events (and the contract state events do not carry) folded
// into the entity store. No RPC, no React, no side effects — the live indexer
// and the synthetic fixture both go through here, which is what keeps the two
// store shapes identical.

import { parseEpochId } from '@vocdoni/davinci-dkg-sdk'
import { compareEvents } from './events'
import {
  applicationKey,
  ciphertextKey,
  contributionKey,
  epochKey,
  operatorKey,
  slotKey,
  txKey,
  STORE_VERSION,
  type Address,
  type Aid,
  type AppPolicyEntity,
  type ApplicationEntity,
  type ChainMeta,
  type CiphertextEntity,
  type EpochEntity,
  type EpochId,
  type EpochPhaseName,
  type EpochPolicy,
  type Hex,
  type IndexedEvent,
  type IndexerStore,
  type NodeStatusName,
  type OperatorEntity,
  type Point,
  type TxMeta,
} from './types'

export interface StoreSeed {
  chainId: number
  chainName?: string
  managerAddress: Address
  registryAddress?: Address | null
  appManagerAddress?: Address | null
  explorerUrl?: string
  deployBlock: number
  blockTimeSeconds?: number
  staggerBlocks?: number
}

export function createEmptyStore(seed: StoreSeed): IndexerStore {
  const chain: ChainMeta = {
    chainId: seed.chainId,
    chainName: seed.chainName ?? '',
    managerAddress: seed.managerAddress.toLowerCase() as Address,
    registryAddress: (seed.registryAddress?.toLowerCase() as Address) ?? null,
    appManagerAddress: (seed.appManagerAddress?.toLowerCase() as Address) ?? null,
    explorerUrl: seed.explorerUrl ?? '',
    deployBlock: seed.deployBlock,
    headBlock: seed.deployBlock,
    epochPrefix: null,
    epochDurationBlocks: null,
    committeeSelectionBlocks: null,
    keyAssemblyBlocks: null,
    nextEpochStartBlock: null,
    inactivityWindow: null,
    activeCount: null,
    nodeCount: null,
    blockTimeSeconds: seed.blockTimeSeconds ?? 12,
    staggerBlocks: seed.staggerBlocks ?? 3,
    stateBlock: 0,
  }
  return {
    version: STORE_VERSION,
    chain,
    lastIndexedBlock: Math.max(0, seed.deployBlock - 1),
    operators: {},
    operatorOrder: [],
    epochs: {},
    epochOrder: [],
    slots: {},
    contributions: {},
    applications: {},
    applicationOrder: [],
    ciphertexts: {},
    txMeta: {},
    events: [],
  }
}

/** New top-level identity so `useSyncExternalStore` sees a change. */
export function bumpStore(store: IndexerStore): IndexerStore {
  return { ...store }
}

// ── entity upserts ───────────────────────────────────────────────────────────

function nonceOf(id: EpochId): number {
  try {
    return Number(parseEpochId(id).nonce)
  } catch {
    return 0
  }
}

export function ensureOperator(store: IndexerStore, address: Address, block: number): OperatorEntity {
  const key = operatorKey(address)
  let op = store.operators[key]
  if (!op) {
    op = {
      address: key as Address,
      pubKey: null,
      status: 'none',
      registeredAtBlock: 0,
      lastActiveBlock: 0,
      firstSeenBlock: block,
      keyUpdates: 0,
      reaps: 0,
      reactivations: 0,
      events: [],
      stateBlock: 0,
    }
    store.operators[key] = op
    store.operatorOrder.push(key)
  }
  return op
}

export function ensureEpoch(store: IndexerStore, id: EpochId, block: number): EpochEntity {
  const key = epochKey(id)
  let epoch = store.epochs[key]
  if (!epoch) {
    epoch = {
      id: key as EpochId,
      nonce: nonceOf(key as EpochId),
      creator: '0x0000000000000000000000000000000000000000',
      createdBlock: block,
      createdTx: null,
      startBlock: block,
      seedBlock: 0,
      seed: null,
      seedResolvedBlock: null,
      lotteryThreshold: 0n,
      status: 'committee-selection',
      policy: null,
      registrySnapshot: null,
      committee: [],
      committeeFilledBlock: null,
      abortedBlock: null,
      slots: [],
      contributions: [],
      finalization: null,
      collectivePublicKey: null,
      shareCommitmentHashes: [],
      applications: [],
      counts: { claims: 0, contributions: 0, ciphertexts: 0, partials: 0, combines: 0, applications: 0 },
      events: [],
      stateBlock: 0,
    }
    store.epochs[key] = epoch
    store.epochOrder.push(key)
    store.epochOrder.sort((a, b) => store.epochs[a].nonce - store.epochs[b].nonce)
  }
  return epoch
}

function ensureApplication(
  store: IndexerStore,
  epoch: EpochId,
  aid: Aid,
  block: number,
): ApplicationEntity {
  const key = applicationKey(epoch, aid)
  let app = store.applications[key]
  if (!app) {
    app = {
      key,
      epoch: epochKey(epoch) as EpochId,
      aid: aid.toLowerCase() as Aid,
      creator: '0x0000000000000000000000000000000000000000',
      organizerPK: { x: 0n, y: 0n },
      policy: null,
      createdBlock: block,
      createdTx: null,
      ciphertexts: [],
      events: [],
      stateBlock: 0,
    }
    store.applications[key] = app
    store.applicationOrder.push(key)
    const ep = ensureEpoch(store, epoch, block)
    if (!ep.applications.includes(key)) {
      ep.applications.push(key)
      ep.counts.applications = ep.applications.length
    }
  }
  return app
}

function ensureCiphertext(
  store: IndexerStore,
  epoch: EpochId,
  aid: Aid,
  index: number,
  block: number,
): CiphertextEntity {
  const key = ciphertextKey(epoch, aid, index)
  let ct = store.ciphertexts[key]
  if (!ct) {
    ct = {
      key,
      epoch: epochKey(epoch) as EpochId,
      aid: aid.toLowerCase() as Aid,
      index,
      submitter: '0x0000000000000000000000000000000000000000',
      c1: { x: 0n, y: 0n },
      c2: { x: 0n, y: 0n },
      block,
      tx: null,
      partials: [],
      organizerShare: null,
      combined: null,
    }
    store.ciphertexts[key] = ct
    const app = ensureApplication(store, epoch, aid, block)
    if (!app.ciphertexts.includes(key)) {
      app.ciphertexts.push(key)
      app.ciphertexts.sort((a, b) => store.ciphertexts[a].index - store.ciphertexts[b].index)
    }
  }
  return ct
}

// ── event fold ───────────────────────────────────────────────────────────────

function attach(entityEvents: number[], index: number): void {
  entityEvents.push(index)
}

function applyEvent(store: IndexerStore, ev: IndexedEvent): void {
  const index = store.events.length
  store.events.push(ev)
  if (ev.block > store.lastIndexedBlock) store.lastIndexedBlock = ev.block

  if (ev.actor) attach(ensureOperator(store, ev.actor, ev.block).events, index)
  if (ev.epoch) attach(ensureEpoch(store, ev.epoch, ev.block).events, index)
  if (ev.epoch && ev.aid) attach(ensureApplication(store, ev.epoch, ev.aid, ev.block).events, index)

  switch (ev.name) {
    case 'NodeRegistered': {
      const op = ensureOperator(store, ev.data.operator, ev.block)
      op.pubKey = { x: ev.data.pubX, y: ev.data.pubY }
      op.status = 'active'
      op.registeredAtBlock = ev.block
      op.lastActiveBlock = Math.max(op.lastActiveBlock, ev.block)
      break
    }
    case 'NodeUpdated': {
      const op = ensureOperator(store, ev.data.operator, ev.block)
      op.pubKey = { x: ev.data.pubX, y: ev.data.pubY }
      op.status = 'active'
      op.keyUpdates += 1
      op.lastActiveBlock = Math.max(op.lastActiveBlock, ev.block)
      break
    }
    case 'NodeMarkedActive': {
      const op = ensureOperator(store, ev.data.operator, ev.block)
      op.lastActiveBlock = Math.max(op.lastActiveBlock, ev.data.atBlock || ev.block)
      break
    }
    case 'NodeReaped': {
      const op = ensureOperator(store, ev.data.operator, ev.block)
      op.status = 'inactive'
      op.reaps += 1
      op.lastActiveBlock = Math.max(op.lastActiveBlock, ev.data.lastActiveBlock)
      break
    }
    case 'NodeReactivated': {
      const op = ensureOperator(store, ev.data.operator, ev.block)
      op.status = 'active'
      op.reactivations += 1
      op.registeredAtBlock = ev.block
      op.lastActiveBlock = Math.max(op.lastActiveBlock, ev.block)
      break
    }
    case 'ManagerSet':
      break
    case 'EpochCreated': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.creator = ev.data.organizer
      epoch.createdBlock = ev.block
      epoch.createdTx = ev.tx
      epoch.startBlock = ev.data.startBlock || ev.block
      epoch.seedBlock = ev.data.seedBlock
      epoch.lotteryThreshold = ev.data.lotteryThreshold
      if (epoch.status === 'none') epoch.status = 'committee-selection'
      break
    }
    case 'SeedResolved': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.seed = ev.data.seed
      epoch.seedResolvedBlock = ev.block
      break
    }
    case 'SlotClaimed': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      const key = slotKey(epoch.id, ev.data.slot)
      if (!store.slots[key]) {
        store.slots[key] = {
          key,
          epoch: epoch.id,
          slot: ev.data.slot,
          operator: ev.data.claimer,
          block: ev.block,
          tx: ev.tx,
        }
        epoch.slots.push(key)
        epoch.slots.sort((a, b) => store.slots[a].slot - store.slots[b].slot)
      }
      while (epoch.committee.length <= ev.data.slot) {
        epoch.committee.push('0x0000000000000000000000000000000000000000')
      }
      epoch.committee[ev.data.slot] = ev.data.claimer
      epoch.counts.claims = epoch.slots.length
      break
    }
    case 'CommitteeFilled': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.committeeFilledBlock = ev.block
      if (epoch.status === 'committee-selection') epoch.status = 'key-assembly'
      break
    }
    case 'ContributionSubmitted': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      const key = contributionKey(epoch.id, ev.data.contributorIndex)
      if (!store.contributions[key]) {
        store.contributions[key] = {
          key,
          epoch: epoch.id,
          index: ev.data.contributorIndex,
          contributor: ev.data.contributor,
          commitmentsHash: ev.data.commitmentsHash,
          encryptedSharesHash: ev.data.encryptedSharesHash,
          block: ev.block,
          tx: ev.tx,
        }
        epoch.contributions.push(key)
      }
      epoch.counts.contributions = epoch.contributions.length
      break
    }
    case 'EpochLive': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.finalization = {
        by: ev.tx ? (store.txMeta[txKey(ev.tx)]?.from ?? null) : null,
        block: ev.block,
        tx: ev.tx,
        aggregateCommitmentsHash: ev.data.aggregateCommitmentsHash,
        collectivePublicKeyHash: ev.data.collectivePublicKeyHash,
        shareCommitmentHash: ev.data.shareCommitmentHash,
      }
      if (epoch.status !== 'aborted') epoch.status = 'live'
      break
    }
    case 'EpochAborted': {
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.status = 'aborted'
      epoch.abortedBlock = ev.block
      break
    }
    case 'CiphertextSubmitted': {
      const ct = ensureCiphertext(store, ev.data.epochId, ev.data.aid, ev.data.ciphertextIndex, ev.block)
      ct.submitter = ev.data.submitter
      ct.c1 = ev.data.c1
      ct.c2 = ev.data.c2
      ct.block = ev.block
      ct.tx = ev.tx
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.counts.ciphertexts += 1
      break
    }
    case 'PartialDecryptionSubmitted': {
      const ct = ensureCiphertext(store, ev.data.epochId, ev.data.aid, ev.data.ciphertextIndex, ev.block)
      const already = ct.partials.some((p) => p.participantIndex === ev.data.participantIndex)
      if (!already) {
        ct.partials.push({
          participant: ev.data.participant,
          participantIndex: ev.data.participantIndex,
          block: ev.block,
          tx: ev.tx,
          delta: ev.data.delta,
        })
        ct.partials.sort((a, b) => a.block - b.block || a.participantIndex - b.participantIndex)
        const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
        epoch.counts.partials += 1
      }
      break
    }
    case 'OrganizerShareSubmitted': {
      const ct = ensureCiphertext(store, ev.data.epochId, ev.data.aid, ev.data.ciphertextIndex, ev.block)
      const overwrites = ct.organizerShare ? ct.organizerShare.overwrites + 1 : 0
      ct.organizerShare = {
        block: ev.block,
        tx: ev.tx,
        delta: ev.data.delta,
        a1: ev.data.a1,
        a2: ev.data.a2,
        z: ev.data.z,
        overwrites,
      }
      break
    }
    case 'DecryptionCombined': {
      const ct = ensureCiphertext(store, ev.data.epochId, ev.data.aid, ev.data.ciphertextIndex, ev.block)
      ct.combined = {
        by: ev.tx ? (store.txMeta[txKey(ev.tx)]?.from ?? null) : null,
        block: ev.block,
        tx: ev.tx,
        combineHash: ev.data.combineHash,
        plaintext: ev.data.plaintext,
      }
      const epoch = ensureEpoch(store, ev.data.epochId, ev.block)
      epoch.counts.combines += 1
      break
    }
    case 'ApplicationRegistered': {
      const app = ensureApplication(store, ev.data.epochId, ev.data.aid, ev.block)
      app.creator = ev.data.creator
      app.organizerPK = ev.data.organizerPK
      app.createdBlock = ev.block
      app.createdTx = ev.tx
      break
    }
  }
}

/**
 * Fold a batch of events into the store, in chronological order, skipping
 * anything already present. Returns the number of events actually applied.
 *
 * Duplicates only ever appear at the boundary block of an incremental poll
 * (the scan re-reads the last indexed block to survive a re-org), so the
 * de-duplication set is built from the tail of the log rather than all of it.
 */
export function applyEvents(store: IndexerStore, incoming: IndexedEvent[]): number {
  if (incoming.length === 0) return 0
  const sorted = [...incoming].sort(compareEvents)
  const minBlock = sorted[0].block
  const seen = new Set<string>()
  for (let i = store.events.length - 1; i >= 0; i--) {
    const e = store.events[i]
    if (e.block < minBlock) break
    seen.add(`${e.block}:${e.logIndex}`)
  }
  let applied = 0
  for (const ev of sorted) {
    const key = `${ev.block}:${ev.logIndex}`
    if (seen.has(key)) continue
    seen.add(key)
    applyEvent(store, ev)
    applied += 1
  }
  return applied
}

// ── contract state (what events do not carry) ────────────────────────────────

const PHASE_BY_STATUS: EpochPhaseName[] = [
  'none',
  'committee-selection',
  'key-assembly',
  'live',
  'aborted',
  'completed',
]

export function phaseFromStatus(status: number): EpochPhaseName {
  return PHASE_BY_STATUS[status] ?? 'none'
}

export function statusFromPhase(phase: EpochPhaseName): number {
  const i = PHASE_BY_STATUS.indexOf(phase)
  return i < 0 ? 0 : i
}

const NODE_STATUS: NodeStatusName[] = ['none', 'active', 'inactive']

export function nodeStatusFromCode(status: number): NodeStatusName {
  return NODE_STATUS[status] ?? 'none'
}

export interface EpochStateUpdate {
  status?: number
  policy?: EpochPolicy
  committee?: Address[]
  collectivePublicKey?: Point | null
  shareCommitmentHashes?: (Hex | null)[]
  claimedCount?: number
  contributionCount?: number
  partialDecryptionCount?: number
  ciphertextCount?: number
  stateBlock: number
}

/** Fold a `getEpoch` / `selectedParticipants` read into the store. */
export function applyEpochState(store: IndexerStore, id: EpochId, update: EpochStateUpdate): void {
  const epoch = ensureEpoch(store, id, update.stateBlock)
  if (update.status != null) epoch.status = phaseFromStatus(update.status)
  if (update.policy) epoch.policy = update.policy
  if (update.committee && update.committee.length > 0) {
    epoch.committee = update.committee.map((a) => a.toLowerCase() as Address)
  }
  if (update.collectivePublicKey !== undefined) epoch.collectivePublicKey = update.collectivePublicKey
  if (update.shareCommitmentHashes) epoch.shareCommitmentHashes = update.shareCommitmentHashes
  if (update.claimedCount != null) epoch.counts.claims = Math.max(epoch.counts.claims, update.claimedCount)
  if (update.contributionCount != null) {
    epoch.counts.contributions = Math.max(epoch.counts.contributions, update.contributionCount)
  }
  if (update.partialDecryptionCount != null) {
    epoch.counts.partials = Math.max(epoch.counts.partials, update.partialDecryptionCount)
  }
  if (update.ciphertextCount != null) {
    epoch.counts.ciphertexts = Math.max(epoch.counts.ciphertexts, update.ciphertextCount)
  }
  epoch.stateBlock = update.stateBlock
}

export interface OperatorStateUpdate {
  pubKey?: Point | null
  status?: number
  lastActiveBlock?: number
  registeredAtBlock?: number
  stateBlock: number
}

/** Fold a registry `getNode` read into the store. */
export function applyOperatorState(
  store: IndexerStore,
  address: Address,
  update: OperatorStateUpdate,
): void {
  const op = ensureOperator(store, address, update.stateBlock)
  if (update.pubKey !== undefined) op.pubKey = update.pubKey
  if (update.status != null) op.status = nodeStatusFromCode(update.status)
  if (update.lastActiveBlock != null) op.lastActiveBlock = update.lastActiveBlock
  if (update.registeredAtBlock != null && update.registeredAtBlock > 0) {
    op.registeredAtBlock = update.registeredAtBlock
  }
  op.stateBlock = update.stateBlock
}

export interface ApplicationStateUpdate {
  policy?: AppPolicyEntity
  organizerPK?: Point
  creator?: Address
  createdBlock?: number
  stateBlock: number
}

/** Fold an app-manager `getApplication` read into the store. */
export function applyApplicationState(
  store: IndexerStore,
  epoch: EpochId,
  aid: Aid,
  update: ApplicationStateUpdate,
): void {
  const app = ensureApplication(store, epoch, aid, update.stateBlock)
  if (update.policy) app.policy = update.policy
  if (update.organizerPK) app.organizerPK = update.organizerPK
  if (update.creator) app.creator = update.creator.toLowerCase() as Address
  if (update.createdBlock != null && update.createdBlock > 0) app.createdBlock = update.createdBlock
  app.stateBlock = update.stateBlock
}

/** Fold transaction metadata (sender, gas) into the store and its consumers. */
export function applyTxMeta(store: IndexerStore, metas: TxMeta[]): void {
  if (metas.length === 0) return
  const byHash = new Map<string, TxMeta>()
  for (const meta of metas) {
    const key = txKey(meta.hash)
    const stored: TxMeta = { ...meta, hash: key as Hex }
    store.txMeta[key] = stored
    byHash.set(key, stored)
  }
  // One pass over the consumers that need attribution: `EpochLive` and
  // `DecryptionCombined` name no submitter, so the transaction sender is the
  // only way to say who finalized or who combined.
  for (const epochId of store.epochOrder) {
    const fin = store.epochs[epochId].finalization
    if (fin?.tx) {
      const meta = byHash.get(txKey(fin.tx))
      if (meta) fin.by = meta.from
    }
  }
  for (const key of Object.keys(store.ciphertexts)) {
    const combined = store.ciphertexts[key].combined
    if (combined?.tx) {
      const meta = byHash.get(txKey(combined.tx))
      if (meta) combined.by = meta.from
    }
  }
}

export function applyChainState(store: IndexerStore, update: Partial<ChainMeta>): void {
  store.chain = { ...store.chain, ...update }
}
