import { describe, expect, it } from 'vitest'
import {
  applyEpochState,
  applyEvents,
  applyOperatorState,
  applyTxMeta,
  createEmptyStore,
  phaseFromStatus,
} from './reduce'
import type { Address, EpochId, EventDataMap, Hex, IndexedEvent, IndexedEventName } from './types'

const MANAGER = '0x3f9b338706a31f26d49159478015c8aaeab908ad' as Address
const EPOCH = '0x2f1105e90000000000000007' as EpochId
const AID = '0x0f00000000000000000000000000000000000000000000000000000000000001' as Hex
const ALICE = '0x1111111111111111111111111111111111111111' as Address
const BOB = '0x2222222222222222222222222222222222222222' as Address
const CAROL = '0x3333333333333333333333333333333333333333' as Address

let logIndex = 0

function ev<K extends IndexedEventName>(
  name: K,
  block: number,
  data: EventDataMap[K],
  refs: { epoch?: EpochId; aid?: Hex; actor?: Address; tx?: Hex } = {},
): IndexedEvent {
  return {
    name,
    block,
    tx: refs.tx ?? (`0x${(logIndex + 1).toString(16).padStart(64, '0')}` as Hex),
    logIndex: logIndex++,
    epoch: refs.epoch ?? null,
    aid: refs.aid ?? null,
    actor: refs.actor ?? null,
    data,
  } as IndexedEvent
}

function store() {
  logIndex = 0
  return createEmptyStore({ chainId: 11155111, managerAddress: MANAGER, deployBlock: 100 })
}

describe('applyEvents', () => {
  it('builds operators from registry events', () => {
    const s = store()
    applyEvents(s, [
      ev('NodeRegistered', 101, { operator: ALICE, pubX: 7n, pubY: 9n }, { actor: ALICE }),
      ev('NodeRegistered', 102, { operator: BOB, pubX: 1n, pubY: 2n }, { actor: BOB }),
      ev('NodeMarkedActive', 150, { operator: ALICE, atBlock: 150 }, { actor: ALICE }),
      ev('NodeReaped', 200, { operator: BOB, lastActiveBlock: 102 }, { actor: BOB }),
      ev('NodeReactivated', 220, { operator: BOB }, { actor: BOB }),
      ev('NodeUpdated', 230, { operator: ALICE, pubX: 11n, pubY: 12n }, { actor: ALICE }),
    ])

    expect(s.operatorOrder).toEqual([ALICE, BOB])
    const alice = s.operators[ALICE]
    expect(alice.pubKey).toEqual({ x: 11n, y: 12n })
    expect(alice.status).toBe('active')
    expect(alice.lastActiveBlock).toBe(230)
    expect(alice.keyUpdates).toBe(1)
    const bob = s.operators[BOB]
    expect(bob.status).toBe('active')
    expect(bob.reaps).toBe(1)
    expect(bob.reactivations).toBe(1)
    expect(bob.registeredAtBlock).toBe(220)
    expect(s.lastIndexedBlock).toBe(230)
  })

  it('builds an epoch through its whole lifecycle', () => {
    const s = store()
    applyEvents(s, [
      ev(
        'EpochCreated',
        300,
        { epochId: EPOCH, organizer: ALICE, startBlock: 300, seedBlock: 301, lotteryThreshold: 1n << 250n },
        { epoch: EPOCH, actor: ALICE },
      ),
      ev('SeedResolved', 302, { epochId: EPOCH, seed: ('0x' + 'ab'.repeat(32)) as Hex }, { epoch: EPOCH }),
      ev('SlotClaimed', 302, { epochId: EPOCH, claimer: ALICE, slot: 0 }, { epoch: EPOCH, actor: ALICE }),
      ev('SlotClaimed', 303, { epochId: EPOCH, claimer: BOB, slot: 1 }, { epoch: EPOCH, actor: BOB }),
      ev('SlotClaimed', 303, { epochId: EPOCH, claimer: CAROL, slot: 2 }, { epoch: EPOCH, actor: CAROL }),
      ev('CommitteeFilled', 303, { epochId: EPOCH }, { epoch: EPOCH }),
      ev(
        'ContributionSubmitted',
        330,
        {
          epochId: EPOCH,
          contributor: BOB,
          contributorIndex: 1,
          commitmentsHash: ('0x' + '11'.repeat(32)) as Hex,
          encryptedSharesHash: ('0x' + '22'.repeat(32)) as Hex,
        },
        { epoch: EPOCH, actor: BOB },
      ),
      ev(
        'EpochLive',
        360,
        {
          epochId: EPOCH,
          aggregateCommitmentsHash: ('0x' + '33'.repeat(32)) as Hex,
          collectivePublicKeyHash: ('0x' + '44'.repeat(32)) as Hex,
          shareCommitmentHash: ('0x' + '55'.repeat(32)) as Hex,
        },
        { epoch: EPOCH, tx: ('0x' + 'fe'.repeat(32)) as Hex },
      ),
    ])

    const epoch = s.epochs[EPOCH]
    expect(epoch.nonce).toBe(7)
    expect(epoch.creator).toBe(ALICE)
    expect(epoch.seed).toBe('0x' + 'ab'.repeat(32))
    expect(epoch.committee).toEqual([ALICE, BOB, CAROL])
    expect(epoch.slots).toHaveLength(3)
    expect(epoch.counts.claims).toBe(3)
    expect(epoch.committeeFilledBlock).toBe(303)
    expect(epoch.contributions).toHaveLength(1)
    expect(s.contributions[`${EPOCH}:1`].contributor).toBe(BOB)
    expect(epoch.status).toBe('live')
    expect(epoch.finalization?.block).toBe(360)
    expect(epoch.finalization?.by).toBeNull()
  })

  it('threads a ciphertext through partials, share and combine', () => {
    const s = store()
    applyEvents(s, [
      ev(
        'ApplicationRegistered',
        400,
        { epochId: EPOCH, aid: AID, creator: ALICE, organizerPK: { x: 5n, y: 6n } },
        { epoch: EPOCH, aid: AID, actor: ALICE },
      ),
      ev(
        'CiphertextSubmitted',
        410,
        {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 1,
          submitter: CAROL,
          c1: { x: 1n, y: 2n },
          c2: { x: 3n, y: 4n },
        },
        { epoch: EPOCH, aid: AID, actor: CAROL },
      ),
      ev(
        'PartialDecryptionSubmitted',
        411,
        {
          epochId: EPOCH,
          aid: AID,
          participant: ALICE,
          participantIndex: 0,
          ciphertextIndex: 1,
          delta: { x: 9n, y: 8n },
        },
        { epoch: EPOCH, aid: AID, actor: ALICE },
      ),
      ev(
        'PartialDecryptionSubmitted',
        414,
        {
          epochId: EPOCH,
          aid: AID,
          participant: BOB,
          participantIndex: 1,
          ciphertextIndex: 1,
          delta: { x: 7n, y: 6n },
        },
        { epoch: EPOCH, aid: AID, actor: BOB },
      ),
      ev(
        'OrganizerShareSubmitted',
        415,
        {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 1,
          delta: { x: 1n, y: 1n },
          a1: { x: 2n, y: 2n },
          a2: { x: 3n, y: 3n },
          z: 42n,
        },
        { epoch: EPOCH, aid: AID },
      ),
      ev(
        'OrganizerShareSubmitted',
        417,
        {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 1,
          delta: { x: 10n, y: 10n },
          a1: { x: 2n, y: 2n },
          a2: { x: 3n, y: 3n },
          z: 43n,
        },
        { epoch: EPOCH, aid: AID },
      ),
      ev(
        'DecryptionCombined',
        420,
        {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 1,
          combineHash: ('0x' + '66'.repeat(32)) as Hex,
          plaintext: 1234n,
        },
        { epoch: EPOCH, aid: AID, tx: ('0x' + 'cc'.repeat(32)) as Hex },
      ),
    ])

    const appKey = `${EPOCH}:${AID}`
    const app = s.applications[appKey]
    expect(app.creator).toBe(ALICE)
    expect(app.ciphertexts).toHaveLength(1)
    const ct = s.ciphertexts[`${EPOCH}:${AID}:1`]
    expect(ct.submitter).toBe(CAROL)
    expect(ct.partials.map((p) => p.participantIndex)).toEqual([0, 1])
    expect(ct.organizerShare?.overwrites).toBe(1)
    expect(ct.organizerShare?.z).toBe(43n)
    expect(ct.combined?.plaintext).toBe(1234n)
    expect(s.epochs[EPOCH].counts).toMatchObject({ ciphertexts: 1, partials: 2, combines: 1, applications: 1 })
  })

  it('is idempotent across an overlapping re-scan', () => {
    const s = store()
    const events = [
      ev('NodeRegistered', 101, { operator: ALICE, pubX: 7n, pubY: 9n }, { actor: ALICE }),
      ev('SlotClaimed', 102, { epochId: EPOCH, claimer: ALICE, slot: 0 }, { epoch: EPOCH, actor: ALICE }),
    ]
    expect(applyEvents(s, events)).toBe(2)
    expect(applyEvents(s, events)).toBe(0)
    expect(s.events).toHaveLength(2)
    expect(s.epochs[EPOCH].slots).toHaveLength(1)
  })

  it('attributes finalizations and combines through the transaction sender', () => {
    const s = store()
    const finalizeTx = ('0x' + 'fa'.repeat(32)) as Hex
    const combineTx = ('0x' + 'cb'.repeat(32)) as Hex
    applyEvents(s, [
      ev(
        'EpochLive',
        360,
        {
          epochId: EPOCH,
          aggregateCommitmentsHash: ('0x' + '33'.repeat(32)) as Hex,
          collectivePublicKeyHash: ('0x' + '44'.repeat(32)) as Hex,
          shareCommitmentHash: ('0x' + '55'.repeat(32)) as Hex,
        },
        { epoch: EPOCH, tx: finalizeTx },
      ),
      ev(
        'CiphertextSubmitted',
        410,
        {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 1,
          submitter: CAROL,
          c1: { x: 1n, y: 2n },
          c2: { x: 3n, y: 4n },
        },
        { epoch: EPOCH, aid: AID, actor: CAROL },
      ),
      ev(
        'DecryptionCombined',
        420,
        {
          epochId: EPOCH,
          aid: AID,
          ciphertextIndex: 1,
          combineHash: ('0x' + '66'.repeat(32)) as Hex,
          plaintext: 7n,
        },
        { epoch: EPOCH, aid: AID, tx: combineTx },
      ),
    ])

    applyTxMeta(s, [
      { hash: finalizeTx, from: BOB, gasUsed: 1_112_337, blockNumber: 360, status: 'success' },
      { hash: combineTx, from: CAROL, gasUsed: 430_432, blockNumber: 420, status: 'success' },
    ])

    expect(s.epochs[EPOCH].finalization?.by).toBe(BOB)
    expect(s.ciphertexts[`${EPOCH}:${AID}:1`].combined?.by).toBe(CAROL)
    expect(s.txMeta[finalizeTx].gasUsed).toBe(1_112_337)
  })
})

describe('contract state', () => {
  it('folds getEpoch and getNode reads over the event-derived entities', () => {
    const s = store()
    applyEvents(s, [
      ev(
        'EpochCreated',
        300,
        { epochId: EPOCH, organizer: ALICE, startBlock: 300, seedBlock: 301, lotteryThreshold: 5n },
        { epoch: EPOCH, actor: ALICE },
      ),
    ])
    applyEpochState(s, EPOCH, {
      status: 3,
      policy: {
        threshold: 33,
        committeeSize: 64,
        minValidContributions: 40,
        lotteryAlphaBps: 15_000,
        committeeSelectionDeadlineBlock: 325,
        keyAssemblyDeadlineBlock: 350,
        liveNotBeforeBlock: 355,
      },
      committee: [ALICE, BOB],
      claimedCount: 64,
      stateBlock: 500,
    })
    applyOperatorState(s, ALICE, { status: 2, lastActiveBlock: 480, registeredAtBlock: 90, stateBlock: 500 })

    const epoch = s.epochs[EPOCH]
    expect(epoch.status).toBe('live')
    expect(phaseFromStatus(3)).toBe('live')
    expect(epoch.policy?.committeeSize).toBe(64)
    expect(epoch.committee).toEqual([ALICE, BOB])
    expect(epoch.counts.claims).toBe(64)
    expect(epoch.stateBlock).toBe(500)
    expect(s.operators[ALICE].status).toBe('inactive')
    expect(s.operators[ALICE].registeredAtBlock).toBe(90)
  })
})
