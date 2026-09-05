import { useState } from 'react'
import { Link } from 'react-router-dom'
import {
  Address,
  Badge,
  Button,
  Callout,
  CopyButton,
  EmptyState,
  Hash,
  Input,
  ProgressBar,
  Select,
  TxCell,
} from '~kit'
import { blocksToDuration, shortHash } from '~lib/format'
import { paths } from '~routes/paths'
import { cn } from '~lib/cn'
import { POOL_SIZE, type AppModeName } from '~indexer/types'
import type { PlaygroundController } from './controller'
import type { EpochOption } from './controller'
import { NextButton, Record, StepPanel, Transcript, TxLine, Words } from './parts'
import type { PlaygroundChain } from './types'

export interface EpochStepData {
  live: EpochOption[]
  /** Newest epoch of any phase — what to show when nothing is Live. */
  newest: { id: string; nonce: number; phase: string; startBlock: number; endBlock: number | null } | null
  headBlock: number
  blockTimeSeconds: number
  /** Block at which `createEpoch` may fire again. */
  nextEpochStartBlock: number | null
  blocksToNextEpoch: number | null
}

export interface StepProps {
  controller: PlaygroundController
  chain: PlaygroundChain
  epochs: EpochStepData
}

const point = (label: string, p: readonly [bigint, bigint] | null) =>
  p ? [{ label: `${label}.x`, value: p[0] }, { label: `${label}.y`, value: p[1] }] : []

/**
 * A field element cut down to a glance — the same decimal the transcripts
 * print in full, so a coordinate shown inline and the one behind "advanced"
 * read as the same number.
 */
const shortWord = (value: bigint): string => {
  const text = value.toString()
  return text.length > 14 ? `${text.slice(0, 12)}…` : text
}

// ── 1. connect ───────────────────────────────────────────────────────────────

export function ConnectStep({ controller, chain }: StepProps) {
  const { wallet } = chain
  return (
    <StepPanel
      step='connect'
      intro={
        <>
          <p>
            You are about to play the <strong className='text-silver'>organizer</strong> of an application. An organizer
            never creates an epoch and never runs a prover: the operator set produces epochs on a fixed block cadence
            all by itself, and everything below is a handful of curve multiplications in this tab plus two or three
            cheap transactions.
          </p>
          <p className='mt-2'>
            {chain.kind === 'demo'
              ? 'Demo mode: the wallet, the transactions and the committee are simulated from the synthetic fixture. No chain is touched and nothing is signed.'
              : 'The wallet you connect pays for the transactions and, by default, is the only address allowed to submit ciphertexts.'}
          </p>
        </>
      }
      actions={
        wallet.connected ? (
          <NextButton onClick={() => controller.actions.goto('epoch')}>Choose an epoch →</NextButton>
        ) : (
          <Button variant='primary' onClick={controller.actions.connect}>
            {chain.kind === 'demo' ? 'Use the demo wallet' : 'Connect wallet'}
          </Button>
        )
      }
    >
      {wallet.connected ? (
        <Record
          items={[
            {
              label: 'connected as',
              value:
                chain.kind === 'demo' && wallet.address ? (
                  <span className='inline-flex items-center gap-2'>
                    <Address value={wallet.address} explorer={false} />
                    <Badge tone='warn' size='sm'>
                      simulated
                    </Badge>
                  </span>
                ) : wallet.address ? (
                  <Address value={wallet.address} />
                ) : (
                  '—'
                ),
            },
            { label: 'head block', value: chain.headBlock.toLocaleString(), mono: true },
          ]}
        />
      ) : null}
      {wallet.problem && wallet.connected ? (
        <Callout tone='warn' title='The wallet cannot write yet' className='mt-4'>
          {wallet.problem}
        </Callout>
      ) : null}
    </StepPanel>
  )
}

// ── 2. epoch ─────────────────────────────────────────────────────────────────

export function EpochStep({ controller, epochs }: StepProps) {
  const { state, epochId } = controller
  const locked = state.pinned

  if (epochs.live.length === 0) {
    const newest = epochs.newest
    const toNext = epochs.blocksToNextEpoch
    return (
      <StepPanel
        step='epoch'
        intro='An application can only be registered against an epoch whose pool keys are being activated — an epoch in the Live phase. There is none right now.'
      >
        {newest ? (
          <Record
            items={[
              { label: 'newest epoch', value: <Hash value={newest.id} href={paths.epoch(newest.id)} />, mono: true },
              { label: 'nonce', value: newest.nonce, mono: true },
              { label: 'phase', value: <Badge tone='warn'>{newest.phase}</Badge> },
              { label: 'started at block', value: newest.startBlock.toLocaleString(), mono: true },
              { label: 'head block', value: epochs.headBlock.toLocaleString(), mono: true },
              {
                label: 'next epoch may open',
                value: epochs.nextEpochStartBlock != null ? `block ${epochs.nextEpochStartBlock}` : 'unknown',
                mono: true,
                hint:
                  toNext != null
                    ? toNext > 0
                      ? `${toNext} blocks away · ${blocksToDuration(toNext, epochs.blockTimeSeconds)}`
                      : 'the cadence window is open now'
                    : undefined,
              },
            ]}
          />
        ) : (
          <EmptyState title='No epochs yet' description='Nothing has been created on this deployment.' />
        )}
        <Callout tone='info' title='Nothing to do — this resolves itself' className='mt-4'>
          Epochs are produced by the operator set, not by applications: every node races to call{' '}
          <code className='font-mono text-emerald'>createEpoch</code> the moment the cadence window opens, then the
          lottery picks a committee, the committee deals the pool and the epoch flips to Live. Come back when the
          countdown above runs out.
        </Callout>
      </StepPanel>
    )
  }

  return (
    <StepPanel
      step='epoch'
      intro={
        locked
          ? 'The application is registered against this epoch, so it is pinned: it claimed one of this epoch’s pool keys, and a newer epoch going Live changes nothing below.'
          : `The newest Live epoch is selected by default. Any Live epoch with a free pool key works — each epoch deals ${POOL_SIZE} keys, one per application, and its nodes keep a couple activated ahead of demand.`
      }
      actions={<NextButton onClick={controller.actions.confirmEpoch}>Register an application →</NextButton>}
    >
      {locked ? (
        <Callout tone='ok' title='Epoch pinned' className='mb-4'>
          Registration binds the application to this epoch. Later steps will keep using it even if a newer epoch goes
          Live in the meantime.
        </Callout>
      ) : null}
      <ul className='m-0 flex list-none flex-col gap-2 p-0'>
        {epochs.live.map((option) => {
          const selected = option.id === epochId
          return (
            <li key={option.id}>
              <button
                type='button'
                disabled={locked}
                onClick={() => controller.actions.selectEpoch(option.id)}
                className={cn(
                  'flex w-full flex-wrap items-center justify-between gap-x-6 gap-y-2 rounded-sm border px-4 py-3 text-left transition-colors',
                  'disabled:cursor-not-allowed',
                  selected
                    ? 'border-emerald/50 bg-emerald/5'
                    : 'border-charcoal hover:border-warm-gray hover:bg-onyx disabled:opacity-50'
                )}
              >
                <span className='flex min-w-0 items-center gap-3'>
                  <span
                    className={cn(
                      'h-2.5 w-2.5 shrink-0 rounded-full border',
                      selected ? 'border-emerald bg-emerald' : 'border-charcoal'
                    )}
                  />
                  <span className='min-w-0'>
                    <span className='block font-mono text-[12px] text-silver'>{shortHash(option.id, 10, 6)}</span>
                    <span className='block text-[11px] text-ash'>nonce {option.nonce}</span>
                  </span>
                </span>
                <span className='font-mono text-[11px] tnum text-pewter'>
                  t={option.threshold} of n={option.committeeSize}
                </span>
                <span className='font-mono text-[11px] tnum text-ash'>
                  live since {option.liveSinceBlock != null ? `#${option.liveSinceBlock}` : '—'}
                </span>
                <span
                  className={cn(
                    'min-w-0 font-mono text-[11px]',
                    option.poolActivated > option.poolClaimed ? 'text-ash' : 'text-amber'
                  )}
                  title={`${POOL_SIZE} keys per epoch: ${option.poolActivated} activated (${option.poolActivated - option.poolClaimed} still free), ${option.poolClaimed} claimed by applications, ${POOL_SIZE - option.poolActivated} not activated yet`}
                >
                  {option.poolActivated - option.poolClaimed} activated free · {option.poolClaimed} claimed ·{' '}
                  {POOL_SIZE - option.poolActivated} not activated
                </span>
              </button>
            </li>
          )
        })}
      </ul>
      {epochId ? (
        <p className='mt-3 text-[12px] text-ash'>
          <Link to={paths.epoch(epochId)} className='text-emerald hover:underline'>
            Open this epoch in the explorer
          </Link>{' '}
          to see its lottery, committee and pool.
        </p>
      ) : null}
    </StepPanel>
  )
}

// ── 3. register ──────────────────────────────────────────────────────────────

export function RegisterStep({ controller, chain }: StepProps) {
  const { state, secret, secretFresh, transcripts, organizerKey } = controller
  const [paste, setPaste] = useState('')
  const [pasteError, setPasteError] = useState<string | null>(null)
  const registered = state.registered
  const automatic = state.mode === 'automatic'

  return (
    <StepPanel
      step='register'
      intro={
        <>
          <p>
            Registration claims the epoch’s next activated pool key{' '}
            <code className='font-mono text-emerald'>P_j</code> — one key per application, {POOL_SIZE} per epoch — and
            the <strong className='text-silver'>mode</strong> decides whether an organizer key sits on top of it.
          </p>
          {automatic ? (
            <p className='mt-2'>
              <strong className='text-silver'>Automatic:</strong> no organizer key at all. The encryption key is{' '}
              <code className='font-mono text-emerald'>PK_aid = P_j</code>, the committee combines on its own the
              moment <code className='font-mono text-emerald'>t</code> partials are in, and confidentiality rests on the
              committee threshold plus the decryption window below.
            </p>
          ) : (
            <p className='mt-2'>
              <strong className='text-silver'>Organizer-locked:</strong> the registration publishes{' '}
              <code className='font-mono text-emerald'>PK_org = sk_org·G</code> with a Schnorr proof of possession, so
              the key is <code className='font-mono text-emerald'>PK_aid = P_j + PK_org</code> and nothing combines
              until you reveal <code className='font-mono text-emerald'>sk_org</code> — once, later, when you decide.
              Until then the secret never leaves this tab and is not recoverable.
            </p>
          )}
        </>
      }
      error={state.error}
      actions={
        registered ? (
          <NextButton onClick={() => controller.actions.goto('encrypt')}>Encrypt a value →</NextButton>
        ) : (
          <>
            <Button
              variant='primary'
              loading={state.busy === 'register'}
              disabled={(!automatic && !secret) || !state.aid || !chain.wallet.connected}
              onClick={() => void controller.actions.register()}
            >
              Register application
            </Button>
            <span className='text-[12px] text-ash'>one transaction, ~408k gas · claims the next free pool key</span>
          </>
        )
      }
    >
      <div className='flex flex-col gap-5'>
        <div>
          <div className='label-caps mb-1.5 text-[11px] text-pewter'>application id (aid)</div>
          <div className='flex flex-wrap items-center gap-2'>
            <code className='min-w-0 flex-1 truncate rounded-sm border border-charcoal bg-obsidian px-3 py-2 font-mono text-[12px] text-silver'>
              {state.aid ?? '—'}
            </code>
            {state.aid ? <CopyButton value={state.aid} label='Copy application id' /> : null}
            {!registered ? (
              <Button size='sm' onClick={controller.actions.rollIdentity}>
                Re-roll
              </Button>
            ) : null}
          </div>
          <p className='mt-1.5 text-[11px] text-ash'>
            32 random bytes with the top three bits cleared, so the id is a BN254 scalar — it is a public input of every
            decryption proof and the contract rejects anything larger.
          </p>
        </div>

        {automatic ? (
          <Callout tone='info' title='No organizer key'>
            An automatic application registers with the identity as its organizer key and no proof. There is no secret
            to keep and nothing for you to do after the ciphertexts are in — which is exactly the trade it makes.
          </Callout>
        ) : secret ? (
          <div>
            <div className='label-caps mb-1.5 text-[11px] text-pewter'>organizer secret (sk_org)</div>
            <div className='flex flex-wrap items-center gap-2'>
              <code className='min-w-0 flex-1 truncate rounded-sm border border-amber/30 bg-amber/[0.04] px-3 py-2 font-mono text-[12px] text-silver'>
                {secret.toString()}
              </code>
              <CopyButton value={secret.toString()} label='Copy organizer secret' />
            </div>
            {secretFresh ? (
              <Callout tone='danger' title='Copy this now — it is shown once' className='mt-3'>
                It is held in this tab’s <code className='font-mono'>sessionStorage</code> so the walkthrough survives a
                reload, and it dies with the tab. Losing it makes every ciphertext under this aid permanently
                undecryptable. A real organizer copies it into whatever they use for key custody.
                <div className='mt-3'>
                  <Button size='sm' onClick={controller.actions.acknowledgeSecret}>
                    I have saved it
                  </Button>
                </div>
              </Callout>
            ) : null}
          </div>
        ) : null}

        {!registered && !automatic ? (
          <details className='rounded-sm border border-charcoal px-4 py-3'>
            <summary className='cursor-pointer text-[12px] text-pewter'>Bring your own secret instead</summary>
            <div className='mt-3 flex flex-wrap items-end gap-2'>
              <Input
                mono
                wrapperClassName='min-w-0 flex-1'
                label='sk_org'
                placeholder='decimal or 0x-hex scalar'
                value={paste}
                onChange={(e) => setPaste(e.target.value)}
                error={pasteError}
              />
              <Button
                onClick={() => {
                  setPasteError(controller.actions.useSecret(paste))
                  setPaste('')
                }}
              >
                Use it
              </Button>
            </div>
          </details>
        ) : null}

        <div className='grid gap-4 sm:grid-cols-3'>
          <Select
            label='mode'
            disabled={registered}
            value={state.mode}
            onChange={(e) => controller.actions.setMode(e.target.value as AppModeName)}
            options={[
              { value: 'organizer-locked', label: 'organizer-locked (default)' },
              { value: 'automatic', label: 'automatic' },
            ]}
            hint={
              automatic
                ? 'policy.mode = Automatic — PK_aid = P_j; the committee decrypts on its own.'
                : 'policy.mode = OrganizerLocked — PK_aid = P_j + PK_org; you reveal sk_org once.'
            }
          />
          <Select
            label='ciphertext cap'
            disabled={registered}
            value={String(state.cap)}
            onChange={(e) => controller.actions.setCap(Number(e.target.value))}
            options={[
              { value: '1', label: '1 ciphertext' },
              { value: '4', label: '4 ciphertexts (default)' },
              { value: '16', label: '16 ciphertexts' },
              { value: '0', label: 'unlimited' },
            ]}
            hint='policy.maxCiphertexts — the contract counts and refuses the one past it.'
          />
          <Input
            mono
            label='authorised submitter'
            disabled={registered}
            placeholder={chain.wallet.address ?? '0x… (defaults to you)'}
            value={state.submitter}
            onChange={(e) => controller.actions.setSubmitter(e.target.value)}
            hint='Becomes the one-address allow-list (policy.submitters). Left empty only the registering wallet may submit; the SDK can also open submission to anyone.'
          />
        </div>

        <div className='grid gap-4 sm:grid-cols-2'>
          <Input
            type='datetime-local'
            label='decryption opens'
            disabled={registered}
            value={state.decryptNotBefore}
            onChange={(e) => controller.actions.setWindow(e.target.value, state.decryptNotAfter)}
            hint='policy.decryptNotBefore — partials and combines revert DecryptionNotOpen() before it. Empty: no floor.'
          />
          <Input
            type='datetime-local'
            label='decryption closes'
            disabled={registered}
            value={state.decryptNotAfter}
            onChange={(e) => controller.actions.setWindow(state.decryptNotBefore, e.target.value)}
            hint={
              chain.kind === 'demo'
                ? 'policy.decryptNotAfter — a hard stop, must lie in the future. Empty: never. The demo committee ignores the window.'
                : 'policy.decryptNotAfter — a hard stop, must lie in the future; after it nothing decrypts, ever. Empty: never.'
            }
          />
        </div>

        {registered ? (
          <div className='flex flex-col gap-3'>
            <Callout tone='ok' title='Application registered'>
              Pool key {state.poolIndex ?? '?'} is yours and the epoch is pinned from here on.{' '}
              {controller.epochId && state.aid ? (
                <Link
                  to={paths.application(controller.epochId, state.aid)}
                  className='text-emerald hover:underline'
                >
                  Open it in the explorer
                </Link>
              ) : null}
            </Callout>
            <TxLine tx={state.txs.register} label='registerApplication' />
          </div>
        ) : null}
      </div>

      {state.advanced ? (
        <Transcript
          title='registration transcript'
          note={
            automatic
              ? 'Automatic mode sends no key and no proof: the contract stores the identity (0, 1) as PK_org whatever the calldata says.'
              : transcripts.registration
                ? 'The words this transaction carried. e = keccak(DOMAIN_ORGANIZER_REGISTER_V1 ‖ eid ‖ aid ‖ PK_org ‖ A) mod q, z = w + e·sk_org.'
                : 'PK_org, derived from the secret above. The Schnorr witness A and the response z appear once the transaction is sent.'
          }
        >
          <Words
            items={[
              ...(automatic ? point('PK_org', [0n, 1n]) : point('PK_org', transcripts.registration?.pkOrg ?? organizerKey)),
              ...point('A', transcripts.registration?.a ?? null),
              ...(transcripts.registration ? [{ label: 'z', value: transcripts.registration.z }] : []),
            ]}
          />
        </Transcript>
      ) : null}
    </StepPanel>
  )
}

// ── 4. encrypt ───────────────────────────────────────────────────────────────

export function EncryptStep({ controller }: StepProps) {
  const { state, words, keys, keysError, ciphertext } = controller
  const automatic = state.mode === 'automatic'
  return (
    <StepPanel
      step='encrypt'
      intro={
        <>
          <p>
            ElGamal on BabyJubJub, entirely in this tab:{' '}
            <code className='font-mono text-emerald'>C1 = k·G</code>,{' '}
            <code className='font-mono text-emerald'>C2 = m·G + k·PK_aid</code> for a fresh random{' '}
            <code className='font-mono text-emerald'>k</code>, where{' '}
            <code className='font-mono text-emerald'>PK_aid</code> is what the SDK’s{' '}
            <code className='font-mono text-emerald'>getApplicationKey</code> returns for this application:{' '}
            {automatic ? (
              <>
                its pool key <code className='font-mono text-emerald'>P_j</code>
              </>
            ) : (
              <code className='font-mono text-emerald'>P_j + PK_org</code>
            )}
            .
          </p>
          <p className='mt-2'>
            The plaintext is a non-negative integer below 2<sup>50</sup> — the cap of the baby-step / giant-step
            inversion the committee runs to recover it.
          </p>
        </>
      }
      error={state.error}
      actions={
        <>
          <Button
            variant={ciphertext ? 'secondary' : 'primary'}
            disabled={!state.registered || state.ciphertextIndex != null}
            onClick={() => void controller.actions.encrypt()}
          >
            {ciphertext ? 'Re-encrypt' : 'Encrypt'}
          </Button>
          {ciphertext ? (
            <NextButton onClick={() => controller.actions.goto('submit')}>Submit it →</NextButton>
          ) : null}
        </>
      }
    >
      <div className='max-w-sm'>
        <Input
          mono
          label='plaintext'
          value={state.value}
          disabled={state.ciphertextIndex != null}
          onChange={(e) => controller.actions.setValue(e.target.value)}
          hint={state.ciphertextIndex != null ? 'Frozen: this value is already on chain.' : 'Integer, below 2^50.'}
        />
      </div>
      {keysError ? (
        <Callout tone='warn' title='The application key could not be read yet' className='mt-4'>
          {keysError} — retried on the next block.
        </Callout>
      ) : keys ? (
        <p className='mt-4 text-[12px] text-ash'>
          Pool key {keys.poolIndex} · PK_aid.x{' '}
          <span className='font-mono text-silver' title={keys.key[0].toString()}>
            {shortWord(keys.key[0])}
          </span>
        </p>
      ) : (
        <p className='mt-4 text-[12px] text-ash'>Reading the application key off the chain…</p>
      )}
      {ciphertext ? (
        <div className='mt-5'>
          <div className='label-caps mb-2 text-[11px] text-pewter'>ciphertext — the four calldata words</div>
          <Words items={words} />
        </div>
      ) : null}
      {state.advanced ? (
        <Transcript
          title='keys'
          note='TE (circomlib) form, the convention the SDK and the indexer both use. The words above are the same points in the gnark RTE form the contract stores.'
        >
          <Words
            items={[
              ...point(`P_${keys?.poolIndex ?? 'j'}`, keys?.poolKey ?? null),
              ...(keys?.organizerPK ? point('PK_org', keys.organizerPK) : []),
              ...point('PK_aid', keys?.key ?? null),
            ]}
            columns={1}
          />
        </Transcript>
      ) : null}
    </StepPanel>
  )
}

// ── 5. submit ────────────────────────────────────────────────────────────────

export function SubmitStep({ controller }: StepProps) {
  const { state, words } = controller
  const done = state.ciphertextIndex != null
  return (
    <StepPanel
      step='submit'
      intro={
        <>
          <p>
            Six calldata words and no proof at all. There is deliberately no proof of knowledge of the randomness{' '}
            <code className='font-mono text-emerald'>k</code>: the submitter of a homomorphically aggregated tally
            cannot know it. Replay across applications is stopped by the pool keys instead — every application has
            its own <code className='font-mono text-emerald'>P_j</code>, so a{' '}
            <code className='font-mono text-emerald'>C1</code> copied into another application and opened there only
            yields a value under a different key.
          </p>
          <p className='mt-2'>
            The contract checks the points are canonical, on-curve and non-identity, assigns the next index for this
            aid, and emits <code className='font-mono text-emerald'>CiphertextSubmitted</code>. The prime-order subgroup
            check is skipped on chain (~2M gas) and done by every committee node off chain before it computes a partial.
          </p>
        </>
      }
      error={state.error}
      actions={
        done ? (
          state.mode === 'automatic' ? (
            <NextButton onClick={() => controller.actions.goto('watch')}>Watch the decryption →</NextButton>
          ) : (
            <NextButton onClick={() => controller.actions.goto('reveal')}>Decide about the secret →</NextButton>
          )
        ) : (
          <>
            <Button
              variant='primary'
              loading={state.busy === 'submit'}
              disabled={!state.ciphertext}
              onClick={() => void controller.actions.submit()}
            >
              Submit ciphertext
            </Button>
            <span className='text-[12px] text-ash'>~96k gas for the first, ~79k after</span>
          </>
        )
      }
    >
      <Words items={words} />
      {done ? (
        <div className='mt-5 flex flex-col gap-3'>
          <Record
            items={[
              { label: 'assigned index', value: state.ciphertextIndex, mono: true },
              { label: 'application', value: state.aid ? <Hash value={state.aid} /> : '—' },
            ]}
          />
          <TxLine tx={state.txs.submit} label='submitCiphertext' />
        </div>
      ) : null}
    </StepPanel>
  )
}

// ── 6. reveal ────────────────────────────────────────────────────────────────

export function RevealStep({ controller }: StepProps) {
  const { state, secret, organizerKey } = controller
  const decided = state.reveal !== 'undecided'
  if (state.mode === 'automatic') {
    return (
      <StepPanel
        step='reveal'
        intro={
          <>
            <p>
              Nothing to do here. This application is <strong className='text-silver'>automatic</strong>: it has no
              organizer key, so once <code className='font-mono text-emerald'>t</code> partials are in — and the
              decryption window is open — a committee node combines them by itself.
            </p>
            <p className='mt-2'>
              There is no secret to keep and no way to stop the plaintext from appearing — that is the trade an
              automatic application makes. An organizer-locked application keeps this step for its organizer.
            </p>
          </>
        }
        actions={<NextButton onClick={() => controller.actions.goto('watch')}>Watch the decryption →</NextButton>}
      >
        <Callout tone='info' title='No organizer key'>
          The combine proof runs with the identity as <code className='font-mono'>PK_org</code> and a zero secret; the
          only thing gating this ciphertext is the committee threshold.
        </Callout>
      </StepPanel>
    )
  }
  return (
    <StepPanel
      step='reveal'
      intro={
        <>
          <p>
            Nothing happens to this ciphertext until you act. The contract refuses every partial decryption and every
            combine of an organizer-locked application with{' '}
            <code className='font-mono text-emerald'>OrganizerSecretNotRevealed()</code> until{' '}
            <code className='font-mono text-emerald'>sk_org</code> is on chain, so the committee is parked on this
            application rather than working on it — and there is nothing on chain to collect from before you decide.
          </p>
          <p className='mt-2'>
            The reveal is one transaction and it is final: from that block on the committee answers every ciphertext
            of this application — this one, and every later one — and{' '}
            <code className='font-mono text-emerald'>t</code> members inside the decryption window recover the
            plaintext. Keeping the secret is a legitimate choice and costs nothing: ciphertexts keep landing, nothing
            decrypts, and you can reveal the second a poll closes. No proof, no prover — the contract checks{' '}
            <code className='font-mono text-emerald'>sk_org·G = PK_org</code> itself.
          </p>
        </>
      }
      error={state.error}
      actions={
        decided ? (
          <NextButton onClick={() => controller.actions.goto('watch')}>Watch the decryption →</NextButton>
        ) : (
          <>
            <Button
              variant='primary'
              loading={state.busy === 'reveal'}
              onClick={() => void controller.actions.reveal()}
            >
              Reveal the secret
            </Button>
            <Button onClick={controller.actions.keep}>Keep it for now</Button>
            <span className='text-[12px] text-ash'>~62k gas</span>
          </>
        )
      }
    >
      {state.reveal === 'revealed' ? (
        <div className='flex flex-col gap-3'>
          <Callout tone='ok' title='Secret revealed'>
            <code className='font-mono'>sk_org</code> is on chain for good and the committee starts answering from
            this block: partials first, then the combine. Anyone may relay a reveal — the secret is what
            authenticates it — and a second call reverts <code className='font-mono'>AlreadyRevealed()</code>.
          </Callout>
          <TxLine tx={state.txs.reveal} label='revealOrganizerSecret' />
        </div>
      ) : null}
      {state.reveal === 'kept' ? (
        <Callout tone='warn' title='Secret kept'>
          Nothing is broken. The ciphertext sits at <em>awaiting reveal</em> — the contract refuses every partial and
          combine of this application — until you change your mind.
          <div className='mt-3'>
            <Button size='sm' variant='ghost' loading={state.busy === 'reveal'} onClick={() => void controller.actions.reveal()}>
              Reveal it now
            </Button>
          </div>
        </Callout>
      ) : null}
      {state.advanced ? (
        <Transcript
          title='reveal transcript'
          note='Two calldata words and no proof: the scalar itself, which the contract multiplies by G and compares to the PK_org it stored at registration.'
        >
          <Words
            items={[...(secret ? [{ label: 'sk_org', value: secret }] : []), ...point('PK_org', organizerKey)]}
            columns={1}
          />
        </Transcript>
      ) : null}
    </StepPanel>
  )
}

// ── 7. watch ─────────────────────────────────────────────────────────────────

export function WatchStep({ controller, chain }: StepProps) {
  const { decryption, state } = controller
  const simulated = chain.kind === 'demo'
  if (!decryption) {
    return (
      <StepPanel step='watch' intro='Waiting for the ciphertext to appear on chain.'>
        <EmptyState
          compact
          title='No decryption state yet'
          description='The indexer has not seen this ciphertext yet. It will show up on the next poll.'
        />
      </StepPanel>
    )
  }
  const { partials, threshold, committeeSize, reveal, combined } = decryption
  return (
    <StepPanel
      step='watch'
      intro={
        <>
          Each selected node publishes <code className='font-mono text-emerald'>δ_i = e_{'{j,i}'}·C1</code> — its
          share <code className='font-mono text-emerald'>e_{'{j,i}'}</code> of this application’s pool key{' '}
          <code className='font-mono text-emerald'>P_j</code> — with a Groth16 proof of its DLEQ and a Merkle proof
          of the share. Only the first <code className='font-mono text-emerald'>t</code> members of a seed-derived
          rotation answer, so the wave column below (counted from 0) is{' '}
          <code className='font-mono text-emerald'>⌊(partial block − opened block) / stagger⌋</code>, where the
          ciphertext opened{' '}
          {state.mode === 'automatic'
            ? 'the block it landed'
            : 'at the reveal — the contract refuses every partial before it'}
          . Once <code className='font-mono text-emerald'>t</code> partials are on chain, one node combines them
          and lands the plaintext.
        </>
      }
    >
      <div className='flex flex-col gap-5'>
        <ProgressBar
          value={partials.length}
          total={committeeSize || threshold}
          threshold={threshold}
          label='partial decryptions'
        />

        <div className='grid gap-3 sm:grid-cols-3'>
          <StatusTile
            label='threshold'
            ok={partials.length >= threshold && threshold > 0}
            text={`${partials.length} of ${threshold}`}
          />
          <StatusTile
            label='organizer secret'
            ok={reveal.done}
            text={
              !reveal.required
                ? 'not needed'
                : reveal.done
                  ? `revealed · block ${reveal.block ?? '—'}`
                  : state.reveal === 'kept'
                    ? 'kept · partials refused'
                    : 'pending · partials refused'
            }
          />
          <StatusTile
            label='combined'
            ok={combined.done}
            text={combined.done ? `block ${combined.block ?? '—'}` : 'waiting'}
          />
        </div>

        {partials.length > 0 ? (
          <div className='overflow-x-auto scroll-slim'>
            <table className='w-full min-w-[420px] border-collapse text-left'>
              <thead>
                <tr className='border-b border-charcoal'>
                  {['participant', 'index', 'wave', 'block', 'tx'].map((h) => (
                    <th key={h} className='label-caps py-2 pr-4 text-[10px] text-pewter'>
                      {h}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {partials.map((partial) => (
                  <tr key={`${partial.participantIndex}:${partial.block}`} className='border-b border-charcoal/60'>
                    <td className='py-1.5 pr-4'>
                      {partial.participant ? (
                        <Address value={partial.participant} explorer={false} copy={false} />
                      ) : (
                        <span className='text-ash'>—</span>
                      )}
                    </td>
                    <td className='py-1.5 pr-4 font-mono text-[12px] tnum text-silver'>{partial.participantIndex}</td>
                    <td className='py-1.5 pr-4 font-mono text-[12px] tnum text-pewter'>{partial.wave}</td>
                    <td className='py-1.5 pr-4 font-mono text-[12px] tnum text-ash'>{partial.block}</td>
                    <td className='py-1.5 pr-4'>
                      {partial.tx ? <TxCell hash={partial.tx} chars={4} /> : <span className='text-ash'>—</span>}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ) : reveal.required && !reveal.done ? (
          <EmptyState
            compact
            title='Awaiting reveal — partials refused'
            description='The contract rejects every partial decryption and combine of an organizer-locked application until sk_org is on chain (OrganizerSecretNotRevealed). Reveal the secret and the committee starts answering.'
          />
        ) : (
          <EmptyState compact title='No partials yet' description='The committee has not answered this ciphertext.' />
        )}

        <div className='flex flex-col gap-2'>
          {reveal.done && reveal.tx ? (
            <TxLine
              tx={{ hash: reveal.tx, block: reveal.block, gasUsed: null, simulated }}
              label='revealOrganizerSecret'
            />
          ) : null}
          {combined.done && combined.tx ? (
            <TxLine
              tx={{ hash: combined.tx, block: combined.block, gasUsed: null, simulated }}
              label='combineDecryption'
            />
          ) : null}
        </div>

        {combined.done ? (
          <Callout tone='ok' title='Plaintext recovered on chain'>
            <span className='font-mono text-ghost'>{combined.plaintext?.toString() ?? '—'}</span>
          </Callout>
        ) : null}
      </div>
    </StepPanel>
  )
}

function StatusTile({ label, ok, text }: { label: string; ok: boolean; text: string }) {
  return (
    <div className={cn('rounded-sm border px-3 py-2', ok ? 'border-emerald/30 bg-emerald/5' : 'border-charcoal')}>
      <div className='label-caps text-[10px] text-pewter'>{label}</div>
      <div className={cn('mt-1 font-mono text-[13px] tnum', ok ? 'text-emerald' : 'text-silver')}>{text}</div>
    </div>
  )
}

// ── 8. verify ────────────────────────────────────────────────────────────────

export function VerifyStep({ controller }: StepProps) {
  const { decryption, ciphertext, state, words } = controller
  const onChain = decryption?.onChain ?? null
  const ciphertextMatches =
    ciphertext != null &&
    onChain != null &&
    ciphertext.c1[0] === onChain.c1[0] &&
    ciphertext.c1[1] === onChain.c1[1] &&
    ciphertext.c2[0] === onChain.c2[0] &&
    ciphertext.c2[1] === onChain.c2[1]
  const recovered = decryption?.combined.plaintext ?? null
  const plaintextMatches = recovered != null && recovered.toString() === state.value.trim()

  return (
    <StepPanel
      step='verify'
      intro='Two checks, both done here rather than taken on trust: the ciphertext the chain stores is the one this tab built, and the plaintext the committee recovered is the number you typed.'
      actions={<Button onClick={controller.actions.reset}>Start over</Button>}
    >
      <div className='flex flex-col gap-4'>
        <CheckRow
          ok={ciphertextMatches}
          title='on-chain (C1, C2) equals the local ciphertext'
          detail={
            onChain
              ? 'Read back from CiphertextSubmitted and converted to TE, compared coordinate by coordinate.'
              : 'The chain has not shown the ciphertext back yet.'
          }
        />
        <CheckRow
          ok={plaintextMatches}
          title='recovered plaintext equals the encrypted value'
          detail={
            recovered != null
              ? `chain says ${recovered.toString()}, this tab encrypted ${state.value.trim()}`
              : 'No combined decryption yet.'
          }
        />
        <div>
          <div className='label-caps mb-2 text-[11px] text-pewter'>the words that were compared</div>
          <Words items={words} />
        </div>
      </div>
    </StepPanel>
  )
}

function CheckRow({ ok, title, detail }: { ok: boolean; title: string; detail: string }) {
  return (
    <div
      className={cn(
        'flex items-start gap-3 rounded-sm border px-4 py-3',
        ok ? 'border-emerald/30 bg-emerald/5' : 'border-charcoal'
      )}
    >
      <span
        className={cn(
          'mt-0.5 flex h-4 w-4 shrink-0 items-center justify-center rounded-full border font-mono text-[10px]',
          ok ? 'border-emerald text-emerald' : 'border-charcoal text-ash'
        )}
      >
        {ok ? '✓' : '·'}
      </span>
      <div className='min-w-0'>
        <div className={cn('text-[13px]', ok ? 'text-ghost' : 'text-silver')}>{title}</div>
        <div className='mt-0.5 text-[11px] text-ash'>{detail}</div>
      </div>
    </div>
  )
}
