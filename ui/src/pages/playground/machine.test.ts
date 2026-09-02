import { describe, expect, it } from 'vitest'
import {
  activeStep,
  furthestStep,
  initialState,
  reducer,
  resolveEpoch,
  resumeState,
  stepStatus,
  type PlaygroundState,
  type TxState,
} from './machine'

const EPOCH_A = '0x0102030405060708090a0b0c'
const EPOCH_B = '0x0102030405060708090a0b0d'
const AID = `0x1c${'ab'.repeat(31)}`

const tx = (hash: string): TxState => ({ hash, block: 100, gasUsed: 21_000, simulated: true })

/** Drive the reducer through a list of actions. */
function run(state: PlaygroundState, ...actions: Parameters<typeof reducer>[1][]): PlaygroundState {
  return actions.reduce(reducer, state)
}

const connected = () => run(initialState(), { type: 'connected', address: '0xabc' })

const registered = () =>
  run(
    connected(),
    { type: 'select-epoch', epochId: EPOCH_A },
    { type: 'registered', aid: AID, tx: tx('0xreg') }
  )

const submitted = () =>
  run(
    registered(),
    { type: 'encrypted', ciphertext: { c1: ['1', '2'], c2: ['3', '4'] } },
    { type: 'submitted', ciphertextIndex: 1, tx: tx('0xsub') }
  )

describe('step derivation', () => {
  it('starts at connect and unlocks one step at a time', () => {
    expect(furthestStep(initialState())).toBe('connect')
    expect(furthestStep(connected())).toBe('epoch')

    const chosen = run(connected(), { type: 'select-epoch', epochId: EPOCH_A })
    expect(furthestStep(chosen)).toBe('register')
    expect(furthestStep(registered())).toBe('encrypt')

    const encrypted = run(registered(), { type: 'encrypted', ciphertext: { c1: ['1', '2'], c2: ['3', '4'] } })
    expect(furthestStep(encrypted)).toBe('submit')
    expect(furthestStep(submitted())).toBe('share')
  })

  it('stops at watch until the chain says the ciphertext was combined', () => {
    const released = run(submitted(), { type: 'share-released', tx: tx('0xshare') })
    expect(furthestStep(released, { combined: false })).toBe('watch')
    expect(furthestStep(released, { combined: true })).toBe('verify')
  })

  it('treats a withheld share as a decision, not a dead end', () => {
    const withheld = run(submitted(), { type: 'share-withheld' })
    expect(furthestStep(withheld)).toBe('watch')
  })

  it('lets the user walk back but never skip ahead', () => {
    const state = run(submitted(), { type: 'goto', step: 'register' })
    expect(activeStep(state)).toBe('register')

    const skipped = run(submitted(), { type: 'goto', step: 'verify' })
    expect(activeStep(skipped)).toBe('share')
  })

  it('reports done / current / todo per step', () => {
    const state = submitted()
    expect(stepStatus('register', state)).toBe('done')
    expect(stepStatus('share', state)).toBe('current')
    expect(stepStatus('verify', state)).toBe('todo')
  })

  it('returns to the furthest step after each successful write', () => {
    const state = run(registered(), { type: 'goto', step: 'epoch' }, { type: 'encrypted', ciphertext: { c1: ['1', '2'], c2: ['3', '4'] } })
    expect(activeStep(state)).toBe('submit')
  })
})

describe('epoch pinning', () => {
  it('defaults to the newest live epoch before anything is registered', () => {
    expect(resolveEpoch(initialState(), [EPOCH_B, EPOCH_A])).toBe(EPOCH_B)
    expect(resolveEpoch(initialState(), [])).toBeNull()
  })

  it('honours an explicit selection', () => {
    const state = run(connected(), { type: 'select-epoch', epochId: EPOCH_A })
    expect(resolveEpoch(state, [EPOCH_B, EPOCH_A])).toBe(EPOCH_A)
  })

  it('pins the epoch on registration and ignores later selections', () => {
    const state = registered()
    expect(state.pinned).toBe(true)

    const attempted = reducer(state, { type: 'select-epoch', epochId: EPOCH_B })
    expect(attempted).toBe(state)
    expect(resolveEpoch(attempted, [EPOCH_B, EPOCH_A])).toBe(EPOCH_A)
  })

  it('freezes the identity and the policy once registered', () => {
    const state = registered()
    expect(reducer(state, { type: 'set-aid', aid: '0x00' })).toBe(state)
    expect(reducer(state, { type: 'set-cap', cap: 9 })).toBe(state)
    expect(reducer(state, { type: 'set-submitter', submitter: '0x00' })).toBe(state)
  })
})

describe('value and transactions', () => {
  it('invalidates an unsubmitted ciphertext when the value changes', () => {
    const encrypted = run(registered(), { type: 'encrypted', ciphertext: { c1: ['1', '2'], c2: ['3', '4'] } })
    const retyped = reducer(encrypted, { type: 'set-value', value: '7' })
    expect(retyped.value).toBe('7')
    expect(retyped.ciphertext).toBeNull()
  })

  it('freezes the value once the ciphertext is on chain', () => {
    const state = submitted()
    expect(reducer(state, { type: 'set-value', value: '7' })).toBe(state)
  })

  it('fills in the block and gas of a transaction once resolved', () => {
    const state = reducer(registered(), {
      type: 'tx-resolved',
      slot: 'register',
      block: 4242,
      gasUsed: 407_793,
    })
    expect(state.txs.register).toMatchObject({ block: 4242, gasUsed: 407_793 })
  })

  it('ignores a resolution for a transaction that was never sent', () => {
    const state = connected()
    expect(reducer(state, { type: 'tx-resolved', slot: 'share', block: 1, gasUsed: 1 })).toBe(state)
  })
})

describe('resume from a deep link', () => {
  const base = () => run(initialState(), { type: 'connected', address: '0xabc' })

  it('selects the epoch from the URL', () => {
    const state = resumeState(base(), { epochId: EPOCH_A, aid: null, hasSecret: false, registered: false, session: null })
    expect(state.epochId).toBe(EPOCH_A)
    expect(state.pinned).toBe(false)
  })

  it('stops at the register step when this tab has no secret', () => {
    const state = resumeState(base(), { epochId: EPOCH_A, aid: AID, hasSecret: false, registered: true, session: null })
    expect(state.aid).toBe(AID)
    expect(state.registered).toBe(false)
    expect(furthestStep(state)).toBe('register')
  })

  it('jumps to the right step when the secret and the session are both there', () => {
    const state = resumeState(base(), {
      epochId: EPOCH_A,
      aid: AID,
      hasSecret: true,
      registered: true,
      session: {
        value: '99',
        ciphertext: { c1: ['1', '2'], c2: ['3', '4'] },
        ciphertextIndex: 3,
        share: 'released',
        cap: 16,
        submitter: '0xfeed',
        txs: { register: tx('0xreg') },
      },
    })
    expect(state.pinned).toBe(true)
    expect(state.value).toBe('99')
    expect(state.ciphertextIndex).toBe(3)
    expect(furthestStep(state, { combined: false })).toBe('watch')
    expect(furthestStep(state, { combined: true })).toBe('verify')
  })
})

describe('reset', () => {
  it('clears the walkthrough but keeps the wallet and the advanced toggle', () => {
    const state = run(submitted(), { type: 'toggle-advanced' }, { type: 'reset' })
    expect(state.connected).toBe(true)
    expect(state.advanced).toBe(true)
    expect(state.aid).toBeNull()
    expect(state.pinned).toBe(false)
    expect(furthestStep(state)).toBe('epoch')
  })
})
