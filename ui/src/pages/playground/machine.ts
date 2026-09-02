// The playground's state machine, as a pure reducer.
//
// Everything the stepper decides — which step is live, which are done, whether
// the epoch may still change — is a function of this state plus two facts read
// off the chain (`ChainFacts`). Nothing here touches React, the DOM, the
// wallet or storage, so the whole walkthrough can be replayed in a unit test.
//
// The one rule that is easy to get wrong, and therefore the one this file
// exists to enforce: **the epoch is pinned the moment an application is
// registered**. `PK_aid = PK_ep + PK_org` is bound to the epoch the
// registration landed in, so a newer epoch going Live must not silently
// re-point the encrypt / submit / release steps at a key nothing was
// registered under.

export type StepId = 'connect' | 'epoch' | 'register' | 'encrypt' | 'submit' | 'share' | 'watch' | 'verify'

export const STEPS: readonly StepId[] = [
  'connect',
  'epoch',
  'register',
  'encrypt',
  'submit',
  'share',
  'watch',
  'verify',
]

export const STEP_TITLES: Record<StepId, string> = {
  connect: 'Connect a wallet',
  epoch: 'Choose a live epoch',
  register: 'Register an application',
  encrypt: 'Encrypt a value',
  submit: 'Submit the ciphertext',
  share: 'Release or withhold the share',
  watch: 'Watch the decryption',
  verify: 'Verify locally',
}

export type StepStatus = 'done' | 'current' | 'todo'

/** Which transaction a `TxRecord` belongs to. */
export type TxSlot = 'register' | 'submit' | 'share'

export type ShareDecision = 'undecided' | 'released' | 'withheld'

/** A ciphertext in the shape session storage can hold (decimal strings). */
export interface SerialCiphertext {
  c1: [string, string]
  c2: [string, string]
}

export interface TxState {
  hash: string
  block: number | null
  gasUsed: number | null
  simulated: boolean
}

export interface LogEntry {
  /** Monotonic id so React keys stay stable. */
  id: number
  /** Chain (or simulator) head when the entry was recorded. */
  block: number | null
  step: StepId
  message: string
  tone: 'info' | 'ok' | 'warn' | 'danger'
  tx?: string
}

export interface PlaygroundState {
  connected: boolean
  epochId: string | null
  /**
   * True once the epoch is the user's choice rather than the default the app
   * pre-selected. The epoch step is not "done" until this flips, so nobody is
   * silently walked past the one decision that binds everything after it.
   */
  epochChosen: boolean
  aid: string | null
  /** True once an application exists: the epoch may no longer be changed. */
  pinned: boolean
  registered: boolean
  /** `maxCiphertexts` for the application policy. */
  cap: number
  /** Authorised submitter; empty means "the connected address". */
  submitter: string
  /** The plaintext the user typed, kept as text so the field stays editable. */
  value: string
  ciphertext: SerialCiphertext | null
  ciphertextIndex: number | null
  share: ShareDecision
  txs: Partial<Record<TxSlot, TxState>>
  /** Step the user navigated to; clamped to what is actually reachable. */
  current: StepId | null
  advanced: boolean
  busy: TxSlot | null
  error: string | null
  log: LogEntry[]
  nextLogId: number
}

/** The two chain-derived facts the step derivation needs. */
export interface ChainFacts {
  /** The ciphertext has been combined and its plaintext is on chain. */
  combined: boolean
}

export const NO_FACTS: ChainFacts = { combined: false }

export function initialState(): PlaygroundState {
  return {
    connected: false,
    epochId: null,
    epochChosen: false,
    aid: null,
    pinned: false,
    registered: false,
    cap: 4,
    submitter: '',
    value: '42',
    ciphertext: null,
    ciphertextIndex: null,
    share: 'undecided',
    txs: {},
    current: null,
    advanced: false,
    busy: null,
    error: null,
    log: [],
    nextLogId: 1,
  }
}

export type PlaygroundAction =
  | { type: 'hydrate'; patch: Partial<PlaygroundState> }
  | { type: 'connected'; address: string }
  | { type: 'disconnected' }
  | { type: 'select-epoch'; epochId: string }
  | { type: 'default-epoch'; epochId: string }
  | { type: 'set-aid'; aid: string }
  | { type: 'set-cap'; cap: number }
  | { type: 'set-submitter'; submitter: string }
  | { type: 'set-value'; value: string }
  | { type: 'registered'; aid: string; tx: TxState }
  | { type: 'encrypted'; ciphertext: SerialCiphertext }
  | { type: 'submitted'; ciphertextIndex: number; tx: TxState }
  | { type: 'share-released'; tx: TxState }
  | { type: 'share-withheld' }
  | { type: 'tx-resolved'; slot: TxSlot; block: number | null; gasUsed: number | null }
  | { type: 'goto'; step: StepId }
  | { type: 'toggle-advanced' }
  | { type: 'busy'; slot: TxSlot | null }
  | { type: 'error'; message: string | null }
  | { type: 'log'; entry: Omit<LogEntry, 'id'> }
  | { type: 'reset' }

function withLog(state: PlaygroundState, entry: Omit<LogEntry, 'id'>): PlaygroundState {
  return {
    ...state,
    log: [...state.log, { ...entry, id: state.nextLogId }],
    nextLogId: state.nextLogId + 1,
  }
}

export function reducer(state: PlaygroundState, action: PlaygroundAction): PlaygroundState {
  switch (action.type) {
    case 'hydrate':
      return { ...state, ...action.patch }

    case 'connected':
      return { ...state, connected: true }

    case 'disconnected':
      return { ...state, connected: false }

    case 'select-epoch':
      // The pin. Once an application is registered the epoch is part of the
      // key the ciphertexts were built under and cannot be swapped out.
      if (state.pinned) return state
      if (state.epochId === action.epochId && state.epochChosen) return state
      return { ...state, epochId: action.epochId, epochChosen: true }

    case 'default-epoch':
      // Pre-selection only: it fills the radio in, it does not advance.
      if (state.pinned || state.epochChosen || state.epochId === action.epochId) return state
      return { ...state, epochId: action.epochId }

    case 'set-aid':
      return state.pinned ? state : { ...state, aid: action.aid }

    case 'set-cap':
      return state.pinned ? state : { ...state, cap: action.cap }

    case 'set-submitter':
      return state.pinned ? state : { ...state, submitter: action.submitter }

    case 'set-value':
      // Re-typing the value invalidates a ciphertext that was already built
      // but not yet submitted; once it is on chain the value is frozen.
      if (state.ciphertextIndex != null) return state
      return { ...state, value: action.value, ciphertext: null }

    case 'registered':
      return {
        ...state,
        aid: action.aid,
        registered: true,
        pinned: true,
        txs: { ...state.txs, register: action.tx },
        busy: null,
        error: null,
        current: null,
      }

    case 'encrypted':
      return { ...state, ciphertext: action.ciphertext, error: null, current: null }

    case 'submitted':
      return {
        ...state,
        ciphertextIndex: action.ciphertextIndex,
        txs: { ...state.txs, submit: action.tx },
        busy: null,
        error: null,
        current: null,
      }

    case 'share-released':
      return {
        ...state,
        share: 'released',
        txs: { ...state.txs, share: action.tx },
        busy: null,
        error: null,
        current: null,
      }

    case 'share-withheld':
      return { ...state, share: 'withheld', current: null }

    case 'tx-resolved': {
      const tx = state.txs[action.slot]
      if (!tx) return state
      if (tx.block === action.block && tx.gasUsed === action.gasUsed) return state
      return {
        ...state,
        txs: { ...state.txs, [action.slot]: { ...tx, block: action.block, gasUsed: action.gasUsed } },
      }
    }

    case 'goto':
      return { ...state, current: action.step }

    case 'toggle-advanced':
      return { ...state, advanced: !state.advanced }

    case 'busy':
      return { ...state, busy: action.slot, error: action.slot ? null : state.error }

    case 'error':
      return { ...state, busy: null, error: action.message }

    case 'log':
      return withLog(state, action.entry)

    case 'reset': {
      const fresh = initialState()
      return { ...fresh, connected: state.connected, advanced: state.advanced, nextLogId: state.nextLogId }
    }
  }
}

// ── derivation ───────────────────────────────────────────────────────────────

/** The furthest step the current state has actually unlocked. */
export function furthestStep(state: PlaygroundState, facts: ChainFacts = NO_FACTS): StepId {
  if (!state.connected) return 'connect'
  if (!state.epochId || !state.epochChosen) return 'epoch'
  if (!state.registered) return 'register'
  if (!state.ciphertext) return 'encrypt'
  if (state.ciphertextIndex == null) return 'submit'
  if (state.share === 'undecided') return 'share'
  if (!facts.combined) return 'watch'
  return 'verify'
}

export function stepIndex(step: StepId): number {
  return STEPS.indexOf(step)
}

/**
 * The step actually being shown: whatever the user navigated to, clamped to
 * the furthest unlocked one. Going back to a finished step is always allowed;
 * skipping ahead never is.
 */
export function activeStep(state: PlaygroundState, facts: ChainFacts = NO_FACTS): StepId {
  const furthest = furthestStep(state, facts)
  if (!state.current) return furthest
  return stepIndex(state.current) <= stepIndex(furthest) ? state.current : furthest
}

export function stepStatus(step: StepId, state: PlaygroundState, facts: ChainFacts = NO_FACTS): StepStatus {
  const active = activeStep(state, facts)
  if (step === active) return 'current'
  return stepIndex(step) < stepIndex(furthestStep(state, facts)) ? 'done' : 'todo'
}

/**
 * Which epoch the walkthrough is operating on.
 *
 * Before an application exists the newest Live epoch is the default, so a
 * visitor who just landed does not have to choose anything; afterwards the
 * pinned id wins even when a newer epoch has gone Live in the meantime.
 */
export function resolveEpoch(state: PlaygroundState, liveEpochIds: readonly string[]): string | null {
  if (state.epochId) return state.epochId
  if (state.pinned) return null
  return liveEpochIds[0] ?? null
}

/** True when the epoch selector must be read-only. */
export function epochLocked(state: PlaygroundState): boolean {
  return state.pinned
}

// ── URL / session round-trip ─────────────────────────────────────────────────

export interface ResumeInput {
  /** `?epoch=` */
  epochId: string | null
  /** `?aid=` */
  aid: string | null
  /** A secret is in session storage for that `(epoch, aid)` pair. */
  hasSecret: boolean
  /** The application exists on chain. */
  registered: boolean
  /** Whatever else was stashed for this pair (value, ciphertext, index …). */
  session: Partial<PlaygroundState> | null
}

/**
 * Rebuild the state a `/playground?epoch=X&aid=Y` deep link should resume at.
 *
 * A link alone is not enough to continue: without the organizer secret the
 * later steps cannot compute anything, so the walkthrough stops at the
 * register step and says why.
 */
export function resumeState(base: PlaygroundState, input: ResumeInput): PlaygroundState {
  const state: PlaygroundState = { ...base }
  if (input.epochId) {
    // A deep link names the epoch explicitly, which is a choice.
    state.epochId = input.epochId
    state.epochChosen = true
  }
  if (!input.aid) return state
  state.aid = input.aid
  if (!input.registered || !input.hasSecret) return state
  state.registered = true
  state.pinned = true
  if (input.session) {
    const { value, ciphertext, ciphertextIndex, share, cap, submitter, txs } = input.session
    if (typeof value === 'string') state.value = value
    if (ciphertext) state.ciphertext = ciphertext
    if (typeof ciphertextIndex === 'number') state.ciphertextIndex = ciphertextIndex
    if (share) state.share = share
    if (typeof cap === 'number') state.cap = cap
    if (typeof submitter === 'string') state.submitter = submitter
    if (txs) state.txs = txs
  }
  return state
}
