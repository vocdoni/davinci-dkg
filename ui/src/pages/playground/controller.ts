// Everything the stepper does, in one hook: the reducer, the organizer secret,
// the URL and session round-trip, and the three writes.
//
// The panels below it are dumb — they render `controller.state` and call
// `controller.actions.*`. Keeping the wiring here is what makes the machine in
// `machine.ts` testable in isolation and the panels testable without a chain.

import { useCallback, useEffect, useMemo, useReducer, useRef, useState } from 'react'
import { useSearchParams } from 'react-router-dom'
import type { Address, Hex } from 'viem'
import type { BabyJubPoint, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import {
  clearOrganizerSecret,
  loadOrganizerSecret,
  parseOrganizerSecret,
  saveOrganizerSecret,
} from '~lib/organizer-secret'
import {
  activeStep,
  furthestStep,
  initialState,
  reducer,
  resolveEpoch,
  resumeState,
  type ChainFacts,
  type LogEntry,
  type PlaygroundState,
  type StepId,
} from './machine'
import { clearSession, loadSession, saveSession } from './session'
import {
  applicationPublicKey,
  ciphertextWords,
  deserialiseCiphertext,
  encryptValue,
  organizerPublicKey,
  parsePlaintext,
  randomAid,
  randomOrganizerSecret,
  registrationProof,
  serialiseCiphertext,
  shareProof,
  type RegistrationProof,
  type ShareProof,
} from './organizer'
import type { DecryptionView, PlaygroundChain } from './types'

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as Address

export interface EpochOption {
  id: string
  nonce: number
  threshold: number
  committeeSize: number
  liveSinceBlock: number | null
  /** `PK_ep`, for the "which key am I encrypting under" column. */
  key: { x: bigint; y: bigint } | null
}

/** The exact words each write signed, for the "advanced" panels. */
export interface Transcripts {
  registration: RegistrationProof | null
  share: ShareProof | null
}

export interface PlaygroundController {
  state: PlaygroundState
  step: StepId
  furthest: StepId
  facts: ChainFacts
  /** The epoch actually in force: the pinned one, else the newest Live one. */
  epochId: string | null
  secret: bigint | null
  /** True while the generated secret has not been acknowledged as saved. */
  secretFresh: boolean
  ciphertext: ElGamalCiphertext | null
  decryption: DecryptionView | null
  words: Array<{ label: string; value: bigint }>
  transcripts: Transcripts
  /** `PK_org = sk_org·G`, TE form. */
  organizerKey: BabyJubPoint | null
  /** `PK_aid = PK_ep + PK_org`, TE form; null until an epoch key is known. */
  applicationKeyPoint: BabyJubPoint | null
  actions: {
    connect: () => void
    selectEpoch: (id: string) => void
    /** Accept the epoch on screen — the default counts only once confirmed. */
    confirmEpoch: () => void
    rollIdentity: () => void
    useSecret: (input: string) => string | null
    acknowledgeSecret: () => void
    setCap: (cap: number) => void
    setSubmitter: (submitter: string) => void
    setValue: (value: string) => void
    register: () => Promise<void>
    encrypt: () => Promise<void>
    submit: () => Promise<void>
    release: () => Promise<void>
    withhold: () => void
    goto: (step: StepId) => void
    toggleAdvanced: () => void
    reset: () => void
  }
}

export function usePlaygroundController(chain: PlaygroundChain, epochs: EpochOption[]): PlaygroundController {
  const [params, setParams] = useSearchParams()
  const [state, dispatch] = useReducer(reducer, undefined, initialState)
  const [secret, setSecret] = useState<bigint | null>(null)
  const [secretFresh, setSecretFresh] = useState(false)
  const [transcripts, setTranscripts] = useState<Transcripts>({ registration: null, share: null })
  const hydrated = useRef(false)

  const log = useCallback(
    (entry: Omit<LogEntry, 'id' | 'block'> & { block?: number | null }) => {
      dispatch({ type: 'log', entry: { block: chain.headBlock, ...entry } })
    },
    [chain.headBlock]
  )

  // ── resume from `?epoch=&aid=` plus whatever the tab still remembers ──────
  useEffect(() => {
    if (hydrated.current) return
    hydrated.current = true
    const epochId = params.get('epoch')
    const aid = params.get('aid')
    if (!epochId && !aid) return
    const stored = epochId && aid ? loadOrganizerSecret(epochId as Hex, aid as Hex) : null
    const session = epochId && aid ? loadSession(epochId, aid) : null
    const resumed = resumeState(initialState(), {
      epochId,
      aid,
      hasSecret: stored != null,
      // `aid` only reaches the URL after a registration lands, so a link that
      // carries one names an application that exists; the secret is what
      // decides whether this tab can go on doing anything with it.
      registered: stored != null,
      session,
    })
    if (stored) setSecret(stored)
    dispatch({ type: 'hydrate', patch: resumed })
    if (stored) {
      dispatch({
        type: 'log',
        entry: {
          block: null,
          step: 'register',
          tone: 'ok',
          message: `Resumed application ${aid} on epoch ${epochId} from this tab's session`,
        },
      })
    } else if (aid) {
      dispatch({
        type: 'log',
        entry: {
          block: null,
          step: 'register',
          tone: 'warn',
          message: `No organizer secret stored for ${aid} in this tab — the later steps need it`,
        },
      })
    }
  }, [params])

  // ── keep the URL and the session record in step with the state ───────────
  const epochId = resolveEpoch(state, epochs.map((e) => e.id))

  // The newest Live epoch is a *default*, not a guess: commit it to the state
  // as soon as one exists so the step machine, the URL and the panels all
  // agree on which epoch is in force.
  useEffect(() => {
    if (state.epochId || state.pinned || !epochId) return
    dispatch({ type: 'default-epoch', epochId })
  }, [epochId, state.epochId, state.pinned])

  useEffect(() => {
    if (!hydrated.current) return
    const next = new URLSearchParams(params)
    let changed = false
    const set = (key: string, value: string | null) => {
      const current = next.get(key)
      if (value && current !== value) {
        next.set(key, value)
        changed = true
      } else if (!value && current) {
        next.delete(key)
        changed = true
      }
    }
    set('epoch', epochId)
    set('aid', state.registered ? state.aid : null)
    if (changed) setParams(next, { replace: true })
  }, [epochId, state.aid, state.registered, params, setParams])

  useEffect(() => {
    if (!state.registered || !epochId || !state.aid) return
    saveSession(epochId, state.aid, state)
  }, [state, epochId])

  // ── wallet ───────────────────────────────────────────────────────────────
  useEffect(() => {
    if (chain.wallet.connected && !state.connected) {
      dispatch({ type: 'connected', address: chain.wallet.address ?? '' })
      dispatch({
        type: 'log',
        entry: {
          block: chain.headBlock,
          step: 'connect',
          tone: 'ok',
          message: `Connected as ${chain.wallet.label}`,
        },
      })
    } else if (!chain.wallet.connected && state.connected) {
      dispatch({ type: 'disconnected' })
    }
  }, [chain.wallet.connected, chain.wallet.address, chain.wallet.label, chain.headBlock, state.connected])

  // ── identity: a fresh aid + secret, generated together ───────────────────
  const rollIdentity = useCallback(() => {
    if (state.pinned) return
    if (epochId && state.aid) clearOrganizerSecret(epochId as Hex, state.aid as Hex)
    const aid = randomAid()
    const sk = randomOrganizerSecret()
    setSecret(sk)
    setSecretFresh(true)
    dispatch({ type: 'set-aid', aid })
    // Persisted immediately, not after the transaction: a crash between the
    // two would otherwise strand an application nobody can ever decrypt.
    if (epochId) saveOrganizerSecret(epochId as Hex, aid, sk)
  }, [epochId, state.aid, state.pinned])

  // Generate the pair as soon as an epoch is known, so the register step has
  // something to show without the visitor pressing anything first.
  useEffect(() => {
    if (state.pinned || state.aid || !epochId || !state.connected) return
    rollIdentity()
  }, [epochId, state.aid, state.pinned, state.connected, rollIdentity])

  const useSecretInput = useCallback(
    (input: string): string | null => {
      if (state.pinned) return 'The application is already registered'
      const parsed = parseOrganizerSecret(input)
      if (!parsed) return 'Not a scalar — expected a decimal or 0x-hex integer'
      const aid = state.aid ?? randomAid()
      setSecret(parsed)
      setSecretFresh(false)
      dispatch({ type: 'set-aid', aid })
      if (epochId) saveOrganizerSecret(epochId as Hex, aid as Hex, parsed)
      log({ step: 'register', tone: 'info', message: 'Using a pasted organizer secret' })
      return null
    },
    [epochId, state.aid, state.pinned, log]
  )

  // ── writes ───────────────────────────────────────────────────────────────
  const fail = useCallback(
    (step: StepId, err: unknown) => {
      const message = err instanceof Error ? err.message.split('\n')[0] : String(err)
      dispatch({ type: 'error', message })
      dispatch({ type: 'log', entry: { block: chain.headBlock, step, tone: 'danger', message } })
    },
    [chain.headBlock]
  )

  const register = useCallback(async () => {
    if (!epochId || !state.aid || !secret) return
    dispatch({ type: 'busy', slot: 'register' })
    try {
      // The witness is drawn here, not inside the writer, so the transcript
      // the "advanced" panel prints is the one the transaction carries.
      const proof = registrationProof(secret, epochId as Hex, state.aid as Hex)
      setTranscripts((t) => ({ ...t, registration: proof }))
      const tx = await chain.register({
        aid: state.aid as Hex,
        skOrg: secret,
        authorizedSubmitter: (state.submitter.trim() || ZERO_ADDRESS) as Address,
        maxCiphertexts: state.cap,
        nonce: proof.nonce,
      })
      dispatch({ type: 'registered', aid: state.aid, tx })
      log({
        step: 'register',
        tone: 'ok',
        tx: tx.hash,
        message: `Registered application ${state.aid} on epoch ${epochId}`,
      })
    } catch (err) {
      fail('register', err)
    }
  }, [chain, epochId, state.aid, state.submitter, state.cap, secret, log, fail])

  const encrypt = useCallback(async () => {
    if (!secret || !chain.epochKey) return
    const parsed = parsePlaintext(state.value)
    if ('error' in parsed) {
      dispatch({ type: 'error', message: parsed.error })
      return
    }
    dispatch({ type: 'error', message: null })
    try {
      const ct = await encryptValue(parsed.value, chain.epochKey, organizerPublicKey(secret))
      dispatch({ type: 'encrypted', ciphertext: serialiseCiphertext(ct) })
      log({ step: 'encrypt', tone: 'ok', message: `Encrypted ${parsed.value} under PK_aid = PK_ep + PK_org` })
    } catch (err) {
      fail('encrypt', err)
    }
  }, [chain.epochKey, secret, state.value, log, fail])

  const ciphertext = useMemo(
    () => (state.ciphertext ? deserialiseCiphertext(state.ciphertext) : null),
    [state.ciphertext]
  )

  const submit = useCallback(async () => {
    if (!state.aid || !ciphertext) return
    dispatch({ type: 'busy', slot: 'submit' })
    try {
      const { tx, ciphertextIndex } = await chain.submitCiphertext({ aid: state.aid as Hex, ciphertext })
      dispatch({ type: 'submitted', ciphertextIndex, tx })
      log({
        step: 'submit',
        tone: 'ok',
        tx: tx.hash,
        message: `Ciphertext accepted as index ${ciphertextIndex}`,
      })
    } catch (err) {
      fail('submit', err)
    }
  }, [chain, ciphertext, state.aid, log, fail])

  const release = useCallback(async () => {
    if (!state.aid || !ciphertext || state.ciphertextIndex == null || !secret) return
    dispatch({ type: 'busy', slot: 'share' })
    try {
      const proof = shareProof(epochId as Hex, state.aid as Hex, state.ciphertextIndex, secret, ciphertext)
      setTranscripts((t) => ({ ...t, share: proof }))
      const tx = await chain.releaseShare({
        aid: state.aid as Hex,
        ciphertext,
        ciphertextIndex: state.ciphertextIndex,
        skOrg: secret,
        nonce: proof.nonce,
      })
      dispatch({ type: 'share-released', tx })
      log({ step: 'share', tone: 'ok', tx: tx.hash, message: 'Organizer share Δ = sk_org·C1 published' })
    } catch (err) {
      fail('share', err)
    }
  }, [chain, ciphertext, epochId, state.aid, state.ciphertextIndex, secret, log, fail])

  const withhold = useCallback(() => {
    dispatch({ type: 'share-withheld' })
    log({
      step: 'share',
      tone: 'warn',
      message: 'Share withheld — the committee will still post partials, but nothing can be combined',
    })
  }, [log])

  const reset = useCallback(() => {
    if (epochId && state.aid) {
      clearOrganizerSecret(epochId as Hex, state.aid as Hex)
      clearSession(epochId, state.aid)
    }
    setSecret(null)
    setSecretFresh(false)
    setTranscripts({ registration: null, share: null })
    dispatch({ type: 'reset' })
    log({ step: 'connect', tone: 'info', message: 'Walkthrough reset' })
  }, [epochId, state.aid, log])

  // ── derived ──────────────────────────────────────────────────────────────
  const decryption = chain.decryption(state.ciphertextIndex)
  const facts: ChainFacts = { combined: decryption?.combined.done ?? false }
  const words = useMemo(() => (ciphertext ? ciphertextWords(ciphertext) : []), [ciphertext])
  const organizerKey = useMemo(() => (secret ? organizerPublicKey(secret) : null), [secret])
  const applicationKeyPoint = useMemo(
    () => (chain.epochKey && organizerKey ? applicationPublicKey(chain.epochKey, organizerKey) : null),
    [chain.epochKey, organizerKey]
  )

  return {
    state,
    step: activeStep(state, facts),
    furthest: furthestStep(state, facts),
    facts,
    epochId,
    secret,
    secretFresh,
    ciphertext,
    decryption,
    words,
    transcripts,
    organizerKey,
    applicationKeyPoint,
    actions: {
      connect: chain.wallet.connect,
      selectEpoch: (id: string) => dispatch({ type: 'select-epoch', epochId: id }),
      confirmEpoch: () => {
        if (epochId) dispatch({ type: 'select-epoch', epochId })
        dispatch({ type: 'goto', step: 'register' })
      },
      rollIdentity,
      useSecret: useSecretInput,
      acknowledgeSecret: () => setSecretFresh(false),
      setCap: (cap: number) => dispatch({ type: 'set-cap', cap }),
      setSubmitter: (submitter: string) => dispatch({ type: 'set-submitter', submitter }),
      setValue: (value: string) => dispatch({ type: 'set-value', value }),
      register,
      encrypt,
      submit,
      release,
      withhold,
      goto: (step: StepId) => dispatch({ type: 'goto', step }),
      toggleAdvanced: () => dispatch({ type: 'toggle-advanced' }),
      reset,
    },
  }
}
