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
import type { AppModeName } from '~indexer/types'
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
  ciphertextWords,
  deserialiseCiphertext,
  encryptValue,
  organizerPublicKey,
  parsePlaintext,
  randomAid,
  randomOrganizerSecret,
  registrationProof,
  serialiseCiphertext,
  type RegistrationProof,
} from './organizer'
import { validateWindow } from './window'
import type { ApplicationKeys, DecryptionView, PlaygroundChain } from './types'

export interface EpochOption {
  id: string
  nonce: number
  threshold: number
  committeeSize: number
  liveSinceBlock: number | null
  /** How much of the pool is left, for the "is there a key for me" column. */
  poolClaimed: number
  poolFree: number
}

/** The exact words each write signed, for the "advanced" panels. */
export interface Transcripts {
  /** Null in automatic mode, which sends no Schnorr proof. */
  registration: RegistrationProof | null
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
  /** `PK_org = sk_org·G`, TE form; null for an automatic application. */
  organizerKey: BabyJubPoint | null
  /** The pool key and `PK_aid` of the registered application; null until read. */
  keys: ApplicationKeys | null
  /** Why `keys` could not be read, when it could not. */
  keysError: string | null
  actions: {
    connect: () => void
    selectEpoch: (id: string) => void
    /** Accept the epoch on screen — the default counts only once confirmed. */
    confirmEpoch: () => void
    rollIdentity: () => void
    useSecret: (input: string) => string | null
    acknowledgeSecret: () => void
    setMode: (mode: AppModeName) => void
    setCap: (cap: number) => void
    setSubmitter: (submitter: string) => void
    setWindow: (notBefore: string, notAfter: string) => void
    setValue: (value: string) => void
    register: () => Promise<void>
    encrypt: () => Promise<void>
    submit: () => Promise<void>
    reveal: () => Promise<void>
    keep: () => void
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
  const [transcripts, setTranscripts] = useState<Transcripts>({ registration: null })
  const [keys, setKeys] = useState<ApplicationKeys | null>(null)
  const [keysError, setKeysError] = useState<string | null>(null)
  const hydrated = useRef(false)
  /** `epoch:aid:index` of the combine already written to the log. */
  const loggedCombine = useRef<string | null>(null)

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
    // two would otherwise strand an application nobody can ever decrypt. An
    // automatic application never uses it, but keeping one per pair is what
    // lets a deep link resume either mode.
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
    if (!epochId || !state.aid) return
    const automatic = state.mode === 'automatic'
    if (!automatic && !secret) return
    const window = validateWindow(state.decryptNotBefore, state.decryptNotAfter)
    if ('error' in window) {
      dispatch({ type: 'error', message: window.error })
      return
    }
    dispatch({ type: 'busy', slot: 'register' })
    try {
      // The witness is drawn here, not inside the writer, so the transcript
      // the "advanced" panel prints is the one the transaction carries. An
      // automatic registration carries no key and no proof at all.
      const proof = automatic || !secret ? null : registrationProof(secret, epochId as Hex, state.aid as Hex)
      setTranscripts({ registration: proof })
      const submitter = state.submitter.trim()
      const { tx, poolIndex } = await chain.register({
        aid: state.aid as Hex,
        skOrg: automatic ? null : secret,
        mode: state.mode,
        submitters: submitter ? [submitter as Address] : [],
        maxCiphertexts: state.cap,
        decryptNotBefore: window.notBefore,
        decryptNotAfter: window.notAfter,
        nonce: proof?.nonce ?? 0n,
      })
      dispatch({ type: 'registered', aid: state.aid, tx, poolIndex })
      log({
        step: 'register',
        tone: 'ok',
        tx: tx.hash,
        message: automatic
          ? `Registered automatic application ${state.aid} on epoch ${epochId} — pool key ${poolIndex}, no organizer key`
          : `Registered application ${state.aid} on epoch ${epochId} — pool key ${poolIndex}`,
      })
    } catch (err) {
      fail('register', err)
    }
  }, [
    chain,
    epochId,
    state.aid,
    state.mode,
    state.submitter,
    state.cap,
    state.decryptNotBefore,
    state.decryptNotAfter,
    secret,
    log,
    fail,
  ])

  // ── the application key, read back once the application exists ──────────
  // `getApplicationKey` on a live chain, the simulator's key on the demo one.
  // Re-read whenever the target changes; a failed read is retried on the next
  // head block rather than logged on every tick.
  const applicationKeys = chain.applicationKeys
  useEffect(() => {
    if (!state.registered || !state.aid) {
      setKeys(null)
      setKeysError(null)
      return
    }
    if (keys) return
    let cancelled = false
    applicationKeys(state.aid as Hex)
      .then((read) => {
        if (cancelled) return
        setKeys(read)
        setKeysError(null)
      })
      .catch((err: unknown) => {
        if (cancelled) return
        setKeysError(err instanceof Error ? err.message.split('\n')[0] : String(err))
      })
    return () => {
      cancelled = true
    }
    // `chain.headBlock` is a deliberate retry trigger.
  }, [applicationKeys, state.registered, state.aid, keys, chain.headBlock])

  const encrypt = useCallback(async () => {
    if (!state.aid || !state.registered) return
    const parsed = parsePlaintext(state.value)
    if ('error' in parsed) {
      dispatch({ type: 'error', message: parsed.error })
      return
    }
    dispatch({ type: 'error', message: null })
    try {
      const read = keys ?? (await applicationKeys(state.aid as Hex))
      if (!keys) setKeys(read)
      const ct = await encryptValue(parsed.value, read.key)
      dispatch({ type: 'encrypted', ciphertext: serialiseCiphertext(ct) })
      log({
        step: 'encrypt',
        tone: 'ok',
        message:
          state.mode === 'automatic'
            ? `Encrypted ${parsed.value} under PK_aid = P_${read.poolIndex}`
            : `Encrypted ${parsed.value} under PK_aid = P_${read.poolIndex} + PK_org`,
      })
    } catch (err) {
      fail('encrypt', err)
    }
  }, [applicationKeys, keys, state.aid, state.registered, state.mode, state.value, log, fail])

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

  const reveal = useCallback(async () => {
    if (!state.aid || !secret) return
    dispatch({ type: 'busy', slot: 'reveal' })
    try {
      const tx = await chain.revealSecret({ aid: state.aid as Hex, skOrg: secret })
      dispatch({ type: 'secret-revealed', tx })
      log({
        step: 'reveal',
        tone: 'ok',
        tx: tx.hash,
        message: 'Organizer secret revealed — the committee starts answering and combines on its own from here',
      })
    } catch (err) {
      fail('reveal', err)
    }
  }, [chain, state.aid, secret, log, fail])

  const keep = useCallback(() => {
    dispatch({ type: 'secret-kept' })
    log({
      step: 'reveal',
      tone: 'warn',
      message: 'Secret kept — the contract refuses partials and combines for this application until the reveal',
    })
  }, [log])

  const reset = useCallback(() => {
    if (epochId && state.aid) {
      clearOrganizerSecret(epochId as Hex, state.aid as Hex)
      clearSession(epochId, state.aid)
    }
    setSecret(null)
    setSecretFresh(false)
    setTranscripts({ registration: null })
    setKeys(null)
    setKeysError(null)
    loggedCombine.current = null
    dispatch({ type: 'reset' })
    log({ step: 'connect', tone: 'info', message: 'Walkthrough reset' })
  }, [epochId, state.aid, log])

  // ── derived ──────────────────────────────────────────────────────────────
  const decryption = chain.decryption(state.ciphertextIndex)
  const facts: ChainFacts = { combined: decryption?.combined.done ?? false }

  // The combine is the one event of the walkthrough nobody in this tab sends,
  // so it goes into the log when the chain (or the simulator) shows it — once
  // per ciphertext, keyed so a resumed session does not repeat it.
  const combinedDone = decryption?.combined.done ?? false
  const combinedTx = decryption?.combined.tx ?? null
  const combinedBlock = decryption?.combined.block ?? null
  const combinedPlaintext = decryption?.combined.plaintext?.toString() ?? null
  const partialCount = decryption?.partials.length ?? 0
  useEffect(() => {
    if (!combinedDone || state.ciphertextIndex == null || !state.aid) return
    const key = `${epochId}:${state.aid}:${state.ciphertextIndex}`
    if (loggedCombine.current === key) return
    loggedCombine.current = key
    dispatch({
      type: 'log',
      entry: {
        block: combinedBlock,
        step: 'watch',
        tone: 'ok',
        tx: combinedTx ?? undefined,
        message: `Ciphertext ${state.ciphertextIndex} combined on chain from ${partialCount} partials${
          combinedPlaintext != null ? ` — plaintext ${combinedPlaintext}` : ''
        }`,
      },
    })
  }, [combinedDone, combinedTx, combinedBlock, combinedPlaintext, partialCount, epochId, state.aid, state.ciphertextIndex])
  const words = useMemo(() => (ciphertext ? ciphertextWords(ciphertext) : []), [ciphertext])
  const organizerKey = useMemo(
    () => (secret && state.mode !== 'automatic' ? organizerPublicKey(secret) : null),
    [secret, state.mode]
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
    keys,
    keysError,
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
      setMode: (mode: AppModeName) => dispatch({ type: 'set-mode', mode }),
      setCap: (cap: number) => dispatch({ type: 'set-cap', cap }),
      setSubmitter: (submitter: string) => dispatch({ type: 'set-submitter', submitter }),
      setWindow: (notBefore: string, notAfter: string) => dispatch({ type: 'set-window', notBefore, notAfter }),
      setValue: (value: string) => dispatch({ type: 'set-value', value }),
      register,
      encrypt,
      submit,
      reveal,
      keep,
      goto: (step: StepId) => dispatch({ type: 'goto', step }),
      toggleAdvanced: () => dispatch({ type: 'toggle-advanced' }),
      reset,
    },
  }
}
