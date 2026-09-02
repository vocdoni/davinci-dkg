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
            all by itself, and everything below is a handful of curve multiplications in this tab plus three cheap
            transactions.
          </p>
          <p className='mt-2'>
            {chain.kind === 'demo'
              ? 'Demo mode: the wallet, the transactions and the committee are simulated from the synthetic fixture. No chain is touched and nothing is signed.'
              : 'The wallet you connect pays for three transactions and, by default, becomes the application’s authorised submitter.'}
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
        intro='An application can only be registered against an epoch whose collective key is already assembled — an epoch in the Live phase. There is none right now.'
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
          lottery picks a committee, the committee assembles the key and the epoch flips to Live. Come back when the
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
          ? 'The application is registered against this epoch, so it is pinned: PK_aid = PK_ep + PK_org is bound to this epoch’s key, and a newer epoch going Live changes nothing below.'
          : 'The newest Live epoch is selected by default. Any Live epoch works — its committee is already assembled and its collective key PK_ep is on chain.'
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
                <span className='min-w-0 font-mono text-[11px] text-ash' title={option.key ? option.key.x.toString() : ''}>
                  PK_ep {option.key ? `${option.key.x.toString(16).slice(0, 10)}…` : '—'}
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
          to see its lottery, committee and key.
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

  return (
    <StepPanel
      step='register'
      intro={
        <>
          <p>
            Registration publishes <code className='font-mono text-emerald'>PK_org = sk_org·G</code> with a Schnorr
            proof that you hold <code className='font-mono text-emerald'>sk_org</code>. From then on the application’s
            encryption key is <code className='font-mono text-emerald'>PK_aid = PK_ep + PK_org</code>, so opening a
            ciphertext needs both the committee threshold and you.
          </p>
          <p className='mt-2'>
            The secret itself never leaves this tab. <strong className='text-silver'>It is also not recoverable</strong>{' '}
            — nothing on chain can reconstruct it, and without it every ciphertext of this application stays encrypted
            forever.
          </p>
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
              disabled={!secret || !state.aid || !chain.wallet.connected}
              onClick={() => void controller.actions.register()}
            >
              Register application
            </Button>
            <span className='text-[12px] text-ash'>one transaction, ~408k gas</span>
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

        {secret ? (
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

        {!registered ? (
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

        <div className='grid gap-4 sm:grid-cols-2'>
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
            hint='Only this address may submit ciphertexts. Left empty it resolves to the registering wallet — there is no open submission.'
          />
        </div>

        {registered ? (
          <div className='flex flex-col gap-3'>
            <Callout tone='ok' title='Application registered'>
              The epoch is pinned from here on.{' '}
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
            transcripts.registration
              ? 'The words this transaction carried. e = keccak(DOMAIN_ORGANIZER_REGISTER_V1 ‖ eid ‖ aid ‖ PK_org ‖ A) mod q, z = w + e·sk_org.'
              : 'PK_org, derived from the secret above. The Schnorr witness A and the response z appear once the transaction is sent.'
          }
        >
          <Words
            items={[
              ...point('PK_org', transcripts.registration?.pkOrg ?? organizerKey),
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

export function EncryptStep({ controller, chain }: StepProps) {
  const { state, words, applicationKeyPoint, ciphertext } = controller
  return (
    <StepPanel
      step='encrypt'
      intro={
        <>
          <p>
            ElGamal on BabyJubJub, entirely in this tab:{' '}
            <code className='font-mono text-emerald'>C1 = k·G</code>,{' '}
            <code className='font-mono text-emerald'>C2 = m·G + k·PK_aid</code> for a fresh random{' '}
            <code className='font-mono text-emerald'>k</code>.
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
            disabled={!chain.epochKey || state.ciphertextIndex != null}
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
      {!chain.epochKey ? (
        <Callout tone='warn' title='No collective key yet' className='mt-4'>
          The selected epoch has not published <code className='font-mono'>PK_ep</code>.
        </Callout>
      ) : null}
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
            items={[...point('PK_ep', chain.epochKey), ...point('PK_aid', applicationKeyPoint)]}
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
            cannot know it. Replay across applications is stopped by the organizer key instead — a{' '}
            <code className='font-mono text-emerald'>C1</code> copied into another application and opened there only
            yields <code className='font-mono text-emerald'>sk_ep·C1</code>.
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
          <NextButton onClick={() => controller.actions.goto('share')}>Decide about the share →</NextButton>
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

// ── 6. share ─────────────────────────────────────────────────────────────────

export function ShareStep({ controller }: StepProps) {
  const { state, transcripts } = controller
  const decided = state.share !== 'undecided'
  return (
    <StepPanel
      step='share'
      intro={
        <>
          <p>
            The committee will answer this ciphertext on its own — you do not ask it to. Its partials alone open
            nothing: the plaintext exists only once you publish{' '}
            <code className='font-mono text-emerald'>Δ = sk_org·C1</code> with a Chaum–Pedersen DLEQ proving the same
            secret relates <code className='font-mono text-emerald'>(G, PK_org)</code> and{' '}
            <code className='font-mono text-emerald'>(C1, Δ)</code>.
          </p>
          <p className='mt-2'>
            Withholding it is a legitimate choice and costs nothing: partials keep arriving,{' '}
            <code className='font-mono text-emerald'>combineDecryption</code> reverts{' '}
            <code className='font-mono text-emerald'>OrganizerShareMissing()</code>, and you can release the share the
            second a poll closes. The whole proof is a few curve multiplications and a keccak in this tab — no circuit
            artifacts, no prover.
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
              loading={state.busy === 'share'}
              onClick={() => void controller.actions.release()}
            >
              Release the share
            </Button>
            <Button onClick={controller.actions.withhold}>Withhold for now</Button>
            <span className='text-[12px] text-ash'>~88k gas</span>
          </>
        )
      }
    >
      {state.share === 'released' ? (
        <div className='flex flex-col gap-3'>
          <Callout tone='ok' title='Share published'>
            The contract stores only <code className='font-mono'>keccak256(Δ ‖ A1 ‖ A2 ‖ z)</code> and never verifies
            the DLEQ — the committee’s combine SNARK does, taking the challenge <code className='font-mono'>e</code>{' '}
            from the transcript the contract recomputed. Anyone may relay a share, and re-submission overwrites until
            the ciphertext is combined, so a malformed one cannot brick it.
          </Callout>
          <TxLine tx={state.txs.share} label='submitOrganizerShare' />
        </div>
      ) : null}
      {state.share === 'withheld' ? (
        <Callout tone='warn' title='Share withheld'>
          Nothing is broken. Partials will still land and the ciphertext will sit at{' '}
          <em>threshold met, awaiting share</em> until you change your mind.
          <div className='mt-3'>
            <Button size='sm' variant='ghost' loading={state.busy === 'share'} onClick={() => void controller.actions.release()}>
              Release it now
            </Button>
          </div>
        </Callout>
      ) : null}
      {state.advanced && transcripts.share ? (
        <Transcript
          title='organizer share transcript'
          note='e = keccak256(DOMAIN_ORGANIZER_SHARE_V1 ‖ eid ‖ aid ‖ uint256(ctIdx) ‖ PK_org ‖ C1 ‖ Δ ‖ A1 ‖ A2) mod q, z = w + e·sk_org mod q. All coordinates are the on-chain RTE words.'
        >
          <Words
            items={[
              ...point('Δ', transcripts.share.delta),
              ...point('A1', transcripts.share.a1),
              ...point('A2', transcripts.share.a2),
              { label: 'e', value: transcripts.share.e },
              { label: 'z', value: transcripts.share.z },
            ]}
          />
          <p className='mt-3 text-[11px] text-ash'>
            Local DLEQ check: {transcripts.share.valid ? 'verifies against PK_org.' : 'DOES NOT VERIFY.'}
          </p>
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
  const { partials, threshold, committeeSize, share, combined } = decryption
  return (
    <StepPanel
      step='watch'
      intro={
        <>
          Each selected node publishes <code className='font-mono text-emerald'>δ_i = d_i·C1</code> with a Groth16 proof
          of its own DLEQ. Only the first <code className='font-mono text-emerald'>t</code> members of a seed-derived
          rotation answer, so the wave column below is{' '}
          <code className='font-mono text-emerald'>⌊(partial block − ciphertext block) / stagger⌋</code>. Once{' '}
          <code className='font-mono text-emerald'>t</code> partials <em>and</em> the organizer share are on chain, one
          node combines them and lands the plaintext.
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
            label='organizer share'
            ok={share.present}
            text={share.present ? `block ${share.block ?? '—'}` : state.share === 'withheld' ? 'withheld' : 'pending'}
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
        ) : (
          <EmptyState compact title='No partials yet' description='The committee has not answered this ciphertext.' />
        )}

        <div className='flex flex-col gap-2'>
          {share.present && share.tx ? (
            <TxLine tx={{ hash: share.tx, block: share.block, gasUsed: null, simulated }} label='organizer share' />
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
