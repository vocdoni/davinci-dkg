// Deterministic synthetic network.
//
// Same store shape as the live indexer — it is literally built by pushing a
// generated event stream through the same reducers — so `?demo=1` exercises
// every page, table and chart at production scale with no RPC:
//
//   300 operators (some reaped, a few reactivated)
//   8 epochs, committee 64, t = 33, m_min = 40, α = 1.5
//     · one aborted (committee never filled)
//     · one still in KeyAssembly (the newest)
//   pool keys activated ahead of demand, 2 applications per Live epoch
//     (one organizer-locked, one automatic), 8 ciphertexts each
//   partials in waves of t; every organizer reveals its secret right after
//     registering except the newest Live epoch's, whose ciphertexts sit at
//     "awaiting reveal" with no partials at all — the contract refuses them
//     until the reveal
//   gas from BENCHMARKS.md, blocks on a 12 s cadence

import { buildEpochId } from '@vocdoni/davinci-dkg-sdk'
import {
  applyApplicationState,
  applyChainState,
  applyEpochState,
  applyEvents,
  applyOperatorState,
  applyTxMeta,
  createEmptyStore,
  statusFromPhase,
} from '../indexer/reduce'
import {
  POOL_SIZE,
  type Address,
  type Aid,
  type EpochId,
  type EpochPhaseName,
  type EventDataMap,
  type Hex,
  type IndexedEvent,
  type IndexedEventName,
  type IndexerStore,
  type Point,
  type TxMeta,
} from '../indexer/types'

export interface FixtureOptions {
  seed?: number
  operators?: number
  epochs?: number
  committeeSize?: number
  threshold?: number
  minValidContributions?: number
  lotteryAlphaBps?: number
  applicationsPerEpoch?: number
  ciphertextsPerApplication?: number
  epochDurationBlocks?: number
  committeeSelectionBlocks?: number
  keyAssemblyBlocks?: number
  finalizeGapBlocks?: number
  inactivityWindow?: number
  staggerBlocks?: number
  deployBlock?: number
  chainId?: number
  chainName?: string
  managerAddress?: Address
  registryAddress?: Address
  appManagerAddress?: Address
  explorerUrl?: string
  /** Epoch prefix used to build the bytes12 ids (matches the Sepolia deploy). */
  epochPrefix?: number
}

export const DEFAULT_FIXTURE: Required<
  Omit<FixtureOptions, 'managerAddress' | 'registryAddress' | 'appManagerAddress'>
> & {
  managerAddress: Address
  registryAddress: Address
  appManagerAddress: Address
} = {
  seed: 0xda7c1,
  operators: 300,
  epochs: 8,
  committeeSize: 64,
  threshold: 33,
  minValidContributions: 40,
  lotteryAlphaBps: 15_000,
  applicationsPerEpoch: 2,
  ciphertextsPerApplication: 8,
  epochDurationBlocks: 300,
  committeeSelectionBlocks: 25,
  keyAssemblyBlocks: 25,
  finalizeGapBlocks: 5,
  inactivityWindow: 50_400,
  staggerBlocks: 3,
  deployBlock: 11_619_019,
  chainId: 11155111,
  chainName: 'sepolia',
  explorerUrl: 'https://sepolia.etherscan.io',
  epochPrefix: 0x2f1105e9,
  managerAddress: '0x3f9b338706a31f26d49159478015c8aaeab908ad',
  registryAddress: '0x9a1f2ce4bd0e0b1f5f6d0a4c7e2b3d8f1c5a6e70',
  appManagerAddress: '0x5c3e7a9d1b8f0426ea5d9c3b7f102a4d6e8b9c11',
}

/**
 * Gas figures measured on the Sepolia deployment (see ../BENCHMARKS.md). The
 * proof-less `finalizeEpoch`, `activatePoolKey` and `revealOrganizerSecret`
 * are estimates until the pool-key deployment is re-measured: activation
 * costs what the old proof-carrying finalize did.
 */
export const GAS = {
  registerKey: 322_112,
  createEpoch: 150_279,
  claimSlotFirst: 175_520,
  claimSlot: 103_725,
  submitContribution: 462_523,
  finalizeEpoch: 91_204,
  activatePoolKey: 1_112_337,
  registerApplication: 407_793,
  revealOrganizerSecret: 61_870,
  submitCiphertextFirst: 96_001,
  submitCiphertext: 78_901,
  submitPartialDecryption: 381_604,
  submitPartialDecryptionMax: 398_704,
  combineDecryption: 430_432,
  reap: 46_500,
  reactivate: 52_100,
  heartbeat: 31_400,
} as const

/**
 * Decryption windows (unix seconds) of the fixture's applications, pinned to
 * fixed instants so a screenshot taken today matches one taken next month:
 * organizer-locked ones open in 2023 and, on even epochs, closed in 2025;
 * automatic ones close in 2033.
 */
export const FIXTURE_DECRYPT_OPEN = 1_700_000_000
export const FIXTURE_DECRYPT_CLOSED = 1_750_000_000
export const FIXTURE_DECRYPT_DEADLINE = 2_000_000_000

/** Pool keys the fixture activates per Live epoch: one per application plus two ahead. */
export const FIXTURE_ACTIVATE_AHEAD = 2

// ── deterministic bytes ──────────────────────────────────────────────────────

function fnv1a(text: string): number {
  let hash = 0x811c9dc5
  for (let i = 0; i < text.length; i++) {
    hash ^= text.charCodeAt(i)
    hash = Math.imul(hash, 0x01000193) >>> 0
  }
  return hash >>> 0
}

/** Order-independent pseudo-random bytes: same label ⇒ same value, always. */
function pseudoHex(label: string, bytes: number, seed: number): Hex {
  let state = (fnv1a(label) ^ seed) >>> 0
  let out = '0x'
  for (let i = 0; i < bytes; i++) {
    state = (Math.imul(state, 1664525) + 1013904223) >>> 0
    out += ((state >>> 24) & 0xff).toString(16).padStart(2, '0')
  }
  return out as Hex
}

function pseudoAddress(label: string, seed: number): Address {
  return pseudoHex(label, 20, seed).toLowerCase() as Address
}

function pseudoBig(label: string, seed: number, bytes = 31): bigint {
  return BigInt(pseudoHex(label, bytes, seed))
}

function pseudoPoint(label: string, seed: number): Point {
  return { x: pseudoBig(`${label}:x`, seed), y: pseudoBig(`${label}:y`, seed) }
}

/** An `aid` must be a BN254 scalar: clear the top three bits. */
function pseudoAid(label: string, seed: number): Aid {
  const raw = pseudoHex(label, 32, seed).slice(2)
  const first = parseInt(raw.slice(0, 2), 16) & 0x1f
  return `0x${first.toString(16).padStart(2, '0')}${raw.slice(2)}` as Aid
}

/** Deterministic integer in `[min, max]`. */
function pick(label: string, seed: number, min: number, max: number): number {
  if (max <= min) return min
  return min + (fnv1a(`${label}:${seed}`) % (max - min + 1))
}

// ── event stream builder ─────────────────────────────────────────────────────

interface Emitted {
  events: IndexedEvent[]
  txMeta: TxMeta[]
}

class Builder {
  readonly events: IndexedEvent[] = []
  readonly txMeta: TxMeta[] = []
  private logIndexByBlock = new Map<number, number>()

  constructor(private readonly seed: number) {}

  tx(label: string, block: number, gasUsed: number, from: Address): Hex {
    const hash = pseudoHex(`tx:${label}`, 32, this.seed)
    this.txMeta.push({ hash, from, gasUsed, blockNumber: block, status: 'success' })
    return hash
  }

  emit<K extends IndexedEventName>(
    name: K,
    block: number,
    tx: Hex | null,
    data: EventDataMap[K],
    refs: { epoch?: EpochId | null; aid?: Aid | null; actor?: Address | null } = {},
  ): void {
    const logIndex = this.logIndexByBlock.get(block) ?? 0
    this.logIndexByBlock.set(block, logIndex + 1)
    this.events.push({
      name,
      block,
      tx,
      logIndex,
      epoch: refs.epoch ?? null,
      aid: refs.aid ?? null,
      actor: refs.actor ?? null,
      data,
    } as IndexedEvent)
  }

  result(): Emitted {
    return { events: this.events, txMeta: this.txMeta }
  }
}

// ── the generator ────────────────────────────────────────────────────────────

export interface FixtureMeta {
  options: typeof DEFAULT_FIXTURE
  operators: Address[]
  epochIds: EpochId[]
  /** Phase each epoch ends up in. */
  phases: Record<string, EpochPhaseName>
  applications: Array<{ epoch: EpochId; aid: Aid }>
  headBlock: number
}

export interface Fixture {
  store: IndexerStore
  meta: FixtureMeta
}

export function buildFixture(options: FixtureOptions = {}): Fixture {
  // Undefined callers' keys must not override the defaults (e.g. `staggerBlocks: undefined`).
  const defined = Object.fromEntries(Object.entries(options).filter(([, v]) => v !== undefined)) as FixtureOptions
  const o = { ...DEFAULT_FIXTURE, ...defined }
  const seed = o.seed
  const store = createEmptyStore({
    chainId: o.chainId,
    chainName: o.chainName,
    managerAddress: o.managerAddress,
    registryAddress: o.registryAddress,
    appManagerAddress: o.appManagerAddress,
    explorerUrl: o.explorerUrl,
    deployBlock: o.deployBlock,
    blockTimeSeconds: 12,
    staggerBlocks: o.staggerBlocks,
  })

  const builder = new Builder(seed)
  const operators: Address[] = Array.from({ length: o.operators }, (_, i) => pseudoAddress(`op:${i}`, seed))
  const pubKeys = new Map<Address, Point>()
  const registeredAt = new Map<Address, number>()
  const lastActive = new Map<Address, number>()
  const status = new Map<Address, 'active' | 'inactive'>()

  // ── registry: 300 operators over the first ~600 blocks ─────────────────────
  operators.forEach((operator, i) => {
    const block = o.deployBlock + 1 + Math.floor(i * 2) + pick(`reg:${i}`, seed, 0, 1)
    const key = pseudoPoint(`key:${operator}`, seed)
    pubKeys.set(operator, key)
    registeredAt.set(operator, block)
    lastActive.set(operator, block)
    status.set(operator, 'active')
    builder.emit(
      'NodeRegistered',
      block,
      builder.tx(`register:${i}`, block, GAS.registerKey, operator),
      { operator, pubX: key.x, pubY: key.y },
      { actor: operator },
    )
  })

  const registrationEnd = o.deployBlock + 2 * o.operators + 8
  const firstEpochStart = registrationEnd + 40

  // A tenth of the registry never comes back: reaped a few blocks before the
  // first epoch, so they sit out every lottery.
  const reaped = operators.filter((_, i) => i % 10 === 7)
  reaped.forEach((operator, i) => {
    const block = registrationEnd + 4 + i
    builder.emit(
      'NodeReaped',
      block,
      builder.tx(`reap:${operator}`, block, GAS.reap, operators[0]),
      { operator, lastActiveBlock: lastActive.get(operator) ?? 0 },
      { actor: operator },
    )
    status.set(operator, 'inactive')
  })
  // A handful rejoin.
  const reactivated = reaped.filter((_, i) => i % 5 === 0)
  reactivated.forEach((operator, i) => {
    const block = registrationEnd + 20 + i
    builder.emit(
      'NodeReactivated',
      block,
      builder.tx(`reactivate:${operator}`, block, GAS.reactivate, operator),
      { operator },
      { actor: operator },
    )
    status.set(operator, 'active')
    registeredAt.set(operator, block)
    lastActive.set(operator, block)
  })

  const activeCount = operators.filter((operator) => status.get(operator) === 'active').length

  // ── epochs ────────────────────────────────────────────────────────────────
  const epochIds: EpochId[] = []
  const phases: Record<string, EpochPhaseName> = {}
  const applications: Array<{ epoch: EpochId; aid: Aid }> = []
  const committees = new Map<string, Address[]>()
  const abortedNonce = 3
  const keyAssemblyNonce = o.epochs
  // The newest epoch that goes Live: its organizer keeps its secret and a
  // couple of its ciphertexts stay ready, so every pipeline state is visible.
  let newestLiveNonce = 0
  for (let nonce = 1; nonce <= o.epochs; nonce++) {
    if (nonce !== abortedNonce && nonce !== keyAssemblyNonce) newestLiveNonce = nonce
  }

  let headBlock = firstEpochStart

  for (let e = 0; e < o.epochs; e++) {
    const nonce = e + 1
    const epochId = buildEpochId(o.epochPrefix, BigInt(nonce)).toLowerCase() as EpochId
    epochIds.push(epochId)
    const startBlock = firstEpochStart + e * o.epochDurationBlocks
    const seedBlock = startBlock + 1
    const aborted = nonce === abortedNonce
    const keyAssembly = nonce === keyAssemblyNonce
    const creator = operators[pick(`creator:${nonce}`, seed, 0, o.operators - 1)]

    // τ = α·n·2²⁵⁶ / R, exactly as DKGManager snapshots it.
    const tau = (BigInt(o.lotteryAlphaBps) * BigInt(o.committeeSize) * (1n << 256n)) / 10_000n / BigInt(activeCount)
    builder.emit(
      'EpochCreated',
      startBlock,
      builder.tx(`create:${nonce}`, startBlock, GAS.createEpoch, creator),
      {
        epochId,
        organizer: creator,
        startBlock,
        seedBlock,
        lotteryThreshold: tau,
      },
      { epoch: epochId, actor: creator },
    )

    // Committee: the lottery's winners, rotated per epoch so the operator
    // pages show different service records.
    const claimants: Address[] = []
    const eligible = operators.filter((operator) => status.get(operator) === 'active')
    const offset = pick(`offset:${nonce}`, seed, 0, eligible.length - 1)
    const wanted = aborted ? Math.floor(o.committeeSize * 0.64) : o.committeeSize
    for (let i = 0; i < wanted; i++) {
      claimants.push(eligible[(offset + i * 7) % eligible.length])
    }
    committees.set(epochId, claimants)

    claimants.forEach((operator, slot) => {
      const block = seedBlock + 1 + Math.floor(slot / 4)
      if (slot === 0) {
        builder.emit('SeedResolved', block, null, { epochId, seed: pseudoHex(`seed:${nonce}`, 32, seed) }, {
          epoch: epochId,
        })
      }
      const gas = slot === 0 ? GAS.claimSlotFirst : GAS.claimSlot
      builder.emit(
        'SlotClaimed',
        block,
        builder.tx(`claim:${nonce}:${slot}`, block, gas, operator),
        { epochId, claimer: operator, slot },
        { epoch: epochId, actor: operator },
      )
      lastActive.set(operator, block)
    })

    const selectionEnd = startBlock + o.committeeSelectionBlocks
    if (aborted) {
      const abortBlock = selectionEnd + 2
      builder.emit(
        'EpochAborted',
        abortBlock,
        builder.tx(`abort:${nonce}`, abortBlock, 48_200, operators[0]),
        { epochId },
        { epoch: epochId },
      )
      phases[epochId] = 'aborted'
      headBlock = Math.max(headBlock, abortBlock)
      continue
    }

    const filledBlock = seedBlock + 1 + Math.floor((claimants.length - 1) / 4)
    builder.emit('CommitteeFilled', filledBlock, null, { epochId }, { epoch: epochId })

    // Contributions: everyone but a few stragglers, and only a partial set for
    // the epoch that is still assembling its key.
    const contributors = keyAssembly
      ? claimants.slice(0, 27)
      : claimants.filter((_, i) => i % 16 !== 15 || nonce % 2 === 0)
    contributors.forEach((operator, i) => {
      const slot = claimants.indexOf(operator)
      const block = selectionEnd + 1 + Math.floor(i / 3)
      const tx = builder.tx(`contribute:${nonce}:${slot}`, block, GAS.submitContribution, operator)
      builder.emit(
        'ContributionSubmitted',
        block,
        tx,
        {
          epochId,
          contributor: operator,
          // 1-based on chain: DKGManager checks
          // `epochParticipants[contributorIndex - 1] == msg.sender`.
          contributorIndex: slot + 1,
          commitmentsHash: pseudoHex(`commit:${nonce}:${slot}`, 32, seed),
          encryptedSharesHash: pseudoHex(`shares:${nonce}:${slot}`, 32, seed),
        },
        { epoch: epochId, actor: operator },
      )
      // DKGManager.submitContribution calls registry.markActive.
      builder.emit('NodeMarkedActive', block, tx, { operator, atBlock: block }, { actor: operator })
      lastActive.set(operator, block)
    })

    if (keyAssembly) {
      phases[epochId] = 'key-assembly'
      headBlock = Math.max(headBlock, selectionEnd + 12)
      continue
    }

    const liveBlock = startBlock + o.committeeSelectionBlocks + o.keyAssemblyBlocks + o.finalizeGapBlocks
    const finalizer = claimants[pick(`finalizer:${nonce}`, seed, 0, claimants.length - 1)]
    // Proof-less: finalize only freezes the contributor set and opens the pool.
    builder.emit(
      'EpochLive',
      liveBlock,
      builder.tx(`finalize:${nonce}`, liveBlock, GAS.finalizeEpoch, finalizer),
      { epochId, contributionCount: contributors.length },
      { epoch: epochId },
    )
    phases[epochId] = 'live'

    // ── pool keys: one per application plus two ahead, one block apart ──────
    const activated = Math.min(POOL_SIZE, o.applicationsPerEpoch + FIXTURE_ACTIVATE_AHEAD)
    const activatorOffset = pick(`activator:${nonce}`, seed, 0, claimants.length - 1)
    for (let j = 0; j < activated; j++) {
      const block = liveBlock + 1 + j
      const activator = claimants[(activatorOffset + j) % claimants.length]
      builder.emit(
        'PoolKeyActivated',
        block,
        builder.tx(`activate:${nonce}:${j}`, block, GAS.activatePoolKey, activator),
        { epochId, keyIndex: j, key: pseudoPoint(`pool:${nonce}:${j}`, seed) },
        { epoch: epochId },
      )
    }

    // ── applications and their ciphertexts ──────────────────────────────────
    for (let a = 0; a < Math.min(o.applicationsPerEpoch, POOL_SIZE); a++) {
      const aid = pseudoAid(`aid:${nonce}:${a}`, seed)
      const organizer = operators[pick(`organizer:${nonce}:${a}`, seed, 0, o.operators - 1)]
      const submitter = pseudoAddress(`submitter:${nonce}:${a}`, seed)
      // Odd applications are automatic: no organizer key, anyone may submit,
      // the committee combines on its own. Even ones are organizer-locked with
      // a one-address allow-list and a decryption window.
      const automatic = a % 2 === 1
      const poolIndex = a
      const registerBlock = liveBlock + 2 + a * 2
      applications.push({ epoch: epochId, aid })
      const registerTx = builder.tx(`app:${nonce}:${a}`, registerBlock, GAS.registerApplication, organizer)
      // The manager logs the claim and the app manager the registration, in
      // the same transaction.
      builder.emit(
        'PoolKeyClaimed',
        registerBlock,
        registerTx,
        { epochId, aid, keyIndex: poolIndex },
        { epoch: epochId, aid },
      )
      builder.emit(
        'ApplicationRegistered',
        registerBlock,
        registerTx,
        {
          epochId,
          aid,
          creator: organizer,
          organizerPK: automatic ? { x: 0n, y: 1n } : pseudoPoint(`pkorg:${nonce}:${a}`, seed),
          mode: automatic ? 'automatic' : 'organizer-locked',
          poolIndex,
        },
        { epoch: epochId, aid, actor: organizer },
      )

      // The organizer of a locked application reveals sk_org right after
      // registering — the contract refuses every partial and combine before
      // the reveal, so nothing else could happen first. The newest Live
      // epoch's keeps it, so that epoch shows ciphertexts parked at
      // "awaiting reveal" with no partials.
      const reveals = !automatic && nonce !== newestLiveNonce
      const revealBlock = registerBlock + 3
      if (reveals) {
        builder.emit(
          'OrganizerSecretRevealed',
          revealBlock,
          builder.tx(`reveal:${nonce}:${a}`, revealBlock, GAS.revealOrganizerSecret, organizer),
          { epochId, aid, organizerSecret: pseudoBig(`skorg:${epochId}:${aid}`, seed) },
          { epoch: epochId, aid },
        )
      }
      // Block from which the committee may combine: at once for an automatic
      // application, after the reveal for a locked one, never if it is kept.
      const unlockedAt = automatic ? 0 : reveals ? revealBlock : null

      for (let c = 0; c < o.ciphertextsPerApplication; c++) {
        const index = c + 1
        const ctBlock = registerBlock + 6 + c * 4
        const gas = c === 0 ? GAS.submitCiphertextFirst : GAS.submitCiphertext
        builder.emit(
          'CiphertextSubmitted',
          ctBlock,
          builder.tx(`ct:${nonce}:${a}:${index}`, ctBlock, gas, submitter),
          {
            epochId,
            aid,
            ciphertextIndex: index,
            submitter,
            c1: pseudoPoint(`c1:${nonce}:${a}:${index}`, seed),
            c2: pseudoPoint(`c2:${nonce}:${a}:${index}`, seed),
          },
          { epoch: epochId, aid, actor: submitter },
        )

        // The seed-derived rotation the committee answers in; the combiner is
        // drawn from it too.
        const rotation = pick(`wave:${nonce}:${a}:${index}`, seed, 0, claimants.length - 1)
        // No partial exists before the reveal: the contract refuses them, and
        // the newest Live epoch's organizer has not revealed.
        if (unlockedAt != null) {
          // Wave 0: exactly t members answer within one stagger window.
          // Wave 1: a few late members on a third of the ciphertexts.
          const responders: number[] = []
          for (let i = 0; i < o.threshold; i++) responders.push((rotation + i) % claimants.length)
          const lateCount = index % 3 === 0 ? 5 : 0
          const late: number[] = []
          for (let i = 0; i < lateCount; i++) late.push((rotation + o.threshold + i) % claimants.length)

          responders.forEach((slot, i) => {
            const block = ctBlock + (i % o.staggerBlocks)
            const operator = claimants[slot]
            const gasUsed =
              GAS.submitPartialDecryption +
              ((i * 137) % (GAS.submitPartialDecryptionMax - GAS.submitPartialDecryption))
            builder.emit(
              'PartialDecryptionSubmitted',
              block,
              builder.tx(`partial:${nonce}:${a}:${index}:${slot}`, block, gasUsed, operator),
              {
                epochId,
                aid,
                participant: operator,
                participantIndex: slot + 1,
                ciphertextIndex: index,
                delta: pseudoPoint(`delta:${nonce}:${a}:${index}:${slot}`, seed),
              },
              { epoch: epochId, aid, actor: operator },
            )
            lastActive.set(operator, block)
          })
          late.forEach((slot, i) => {
            const block = ctBlock + o.staggerBlocks + (i % o.staggerBlocks)
            const operator = claimants[slot]
            builder.emit(
              'PartialDecryptionSubmitted',
              block,
              builder.tx(`partial:${nonce}:${a}:${index}:${slot}`, block, GAS.submitPartialDecryption, operator),
              {
                epochId,
                aid,
                participant: operator,
                participantIndex: slot + 1,
                ciphertextIndex: index,
                delta: pseudoPoint(`delta:${nonce}:${a}:${index}:${slot}`, seed),
              },
              { epoch: epochId, aid, actor: operator },
            )
          })
        }

        // The t-th partial lands inside the first stagger window.
        const thresholdBlock = ctBlock + Math.min(o.threshold - 1, o.staggerBlocks - 1)
        if (unlockedAt != null) {
          // Most ciphertexts get combined; the newest Live epoch keeps a few
          // in flight so the "ready" state is visible somewhere.
          const pending = nonce === newestLiveNonce && index > o.ciphertextsPerApplication - 2
          if (!pending) {
            const combineBlock = Math.max(thresholdBlock, unlockedAt) + 3
            const combiner = claimants[(rotation + 1) % claimants.length]
            builder.emit(
              'DecryptionCombined',
              combineBlock,
              builder.tx(`combine:${nonce}:${a}:${index}`, combineBlock, GAS.combineDecryption, combiner),
              {
                epochId,
                aid,
                ciphertextIndex: index,
                combineHash: pseudoHex(`combine:${nonce}:${a}:${index}`, 32, seed),
                plaintext: BigInt(pick(`plain:${nonce}:${a}:${index}`, seed, 1, 4096)),
              },
              { epoch: epochId, aid },
            )
            headBlock = Math.max(headBlock, combineBlock)
          }
        }
        headBlock = Math.max(headBlock, ctBlock + 12)
      }
    }
  }

  // ── fold everything in, then the state events do not carry ────────────────
  const { events, txMeta } = builder.result()
  applyEvents(store, events)
  applyTxMeta(store, txMeta)

  headBlock += 9
  store.lastIndexedBlock = headBlock

  for (let e = 0; e < o.epochs; e++) {
    const epochId = epochIds[e]
    const epoch = store.epochs[epochId]
    const startBlock = epoch.startBlock
    const phase = phases[epochId] ?? 'committee-selection'
    const committee = committees.get(epochId) ?? []
    applyEpochState(store, epochId, {
      status: statusFromPhase(phase),
      policy: {
        threshold: o.threshold,
        committeeSize: o.committeeSize,
        minValidContributions: o.minValidContributions,
        lotteryAlphaBps: o.lotteryAlphaBps,
        committeeSelectionDeadlineBlock: startBlock + o.committeeSelectionBlocks,
        keyAssemblyDeadlineBlock: startBlock + o.committeeSelectionBlocks + o.keyAssemblyBlocks,
        liveNotBeforeBlock:
          startBlock + o.committeeSelectionBlocks + o.keyAssemblyBlocks + o.finalizeGapBlocks,
      },
      committee,
      poolNext: phase === 'live' ? Math.min(o.applicationsPerEpoch, POOL_SIZE) : 0,
      claimedCount: committee.length,
      contributionCount: epoch.contributions.length,
      stateBlock: headBlock,
    })
  }

  for (const operator of operators) {
    applyOperatorState(store, operator, {
      pubKey: pubKeys.get(operator) ?? null,
      status: status.get(operator) === 'inactive' ? 2 : 1,
      lastActiveBlock: lastActive.get(operator) ?? 0,
      registeredAtBlock: registeredAt.get(operator) ?? 0,
      stateBlock: headBlock,
    })
  }

  for (const { epoch, aid } of applications) {
    const app = store.applications[`${epoch}:${aid}`]
    if (!app) continue
    const automatic = app.mode === 'automatic'
    const nonce = store.epochs[epoch]?.nonce ?? 0
    const submitter = (store.ciphertexts[app.ciphertexts[0]]?.submitter ?? app.creator) as Address
    applyApplicationState(store, epoch, aid, {
      mode: app.mode,
      poolIndex: app.poolIndex ?? undefined,
      policy: {
        openSubmission: automatic,
        submitters: automatic ? [] : [submitter],
        maxCiphertexts: o.ciphertextsPerApplication,
        notBeforeBlock: 0,
        notAfterBlock: 0,
        decryptNotBefore: automatic ? 0 : FIXTURE_DECRYPT_OPEN,
        decryptNotAfter: automatic ? FIXTURE_DECRYPT_DEADLINE : nonce % 2 === 0 ? FIXTURE_DECRYPT_CLOSED : 0,
      },
      stateBlock: headBlock,
    })
  }

  applyChainState(store, {
    headBlock,
    epochPrefix: o.epochPrefix,
    epochDurationBlocks: o.epochDurationBlocks,
    committeeSelectionBlocks: o.committeeSelectionBlocks,
    keyAssemblyBlocks: o.keyAssemblyBlocks,
    nextEpochStartBlock: firstEpochStart + (o.epochs - 1) * o.epochDurationBlocks + o.epochDurationBlocks,
    inactivityWindow: o.inactivityWindow,
    nodeCount: o.operators,
    activeCount,
    stateBlock: headBlock,
  })

  return {
    store,
    meta: { options: o, operators, epochIds, phases, applications, headBlock },
  }
}

/** The default network, built once per module instance. */
let cached: Fixture | null = null

export function syntheticFixture(): Fixture {
  if (!cached) cached = buildFixture()
  return cached
}

export function resetSyntheticFixture(): void {
  cached = null
}
