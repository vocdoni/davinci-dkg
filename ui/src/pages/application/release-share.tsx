// Organizer tools: release Δ = sk_org·C1 for one ciphertext.
//
// The secret never leaves the browser. `submitOrganizerShare` proves knowledge
// of it with a Chaum–Pedersen DLEQ against the registered PK_org and sends
// only Δ and the proof; the contract stores the hash and the committee's
// combine SNARK is what actually checks the DLEQ. Relaying is permissionless —
// the wallet that pays for this transaction does not have to be the
// organizer's — and re-submission overwrites until the ciphertext is combined,
// so a malformed share cannot brick anything.

import { useState } from 'react'
import { useAccount, usePublicClient, useWalletClient } from 'wagmi'
import { DKGWriter, fromRTEtoTE } from '@vocdoni/davinci-dkg-sdk'
import { Button, Callout, Input, Select, TxCell } from '~kit'
import type { CiphertextRow } from '~indexer/selectors'
import type { Address, Aid, EpochId, Hex } from '~indexer/types'
import { loadOrganizerSecret, parseOrganizerSecret } from '~lib/organizer-secret'

export interface ReleaseShareProps {
  epoch: EpochId
  aid: Aid
  ciphertexts: CiphertextRow[]
  managerAddress: Address
  appManagerAddress: Address | null
  /** Demo mode has no chain to send to; the panel explains itself instead. */
  demo: boolean
}

interface FieldsProps {
  ciphertexts: CiphertextRow[]
  index: string
  onIndex: (value: string) => void
  secret: string
  onSecret: (value: string) => void
  disabled?: boolean
  secretError?: string
}

function ShareFields({ ciphertexts, index, onIndex, secret, onSecret, disabled, secretError }: FieldsProps) {
  const options = ciphertexts.map((row) => ({
    value: String(row.index),
    label: `#${row.index} · ${row.partialCount}/${row.threshold || '?'} partials${row.share.present ? ' · share on chain' : ''}${row.combined.done ? ' · combined' : ''}`,
  }))
  return (
    <div className='grid gap-4 sm:grid-cols-[220px_1fr]'>
      <Select
        label='Ciphertext'
        aria-label='Ciphertext index'
        value={index}
        disabled={disabled || options.length === 0}
        onChange={(e) => onIndex(e.target.value)}
        options={options.length > 0 ? options : [{ value: '', label: 'no ciphertexts yet' }]}
        hint='index assigned on chain'
      />
      <Input
        label='Organizer secret (sk_org)'
        aria-label='Organizer secret'
        mono
        type='password'
        autoComplete='off'
        spellCheck={false}
        placeholder='decimal or 0x…'
        value={secret}
        disabled={disabled}
        error={secretError}
        onChange={(e) => onSecret(e.target.value)}
        hint='Never sent anywhere: the share and its DLEQ are computed in this tab. Prefilled from this tab’s playground session when one exists.'
      />
    </div>
  )
}

function reason(err: unknown): string {
  if (typeof err === 'object' && err !== null) {
    const e = err as { shortMessage?: unknown; message?: unknown }
    if (typeof e.shortMessage === 'string' && e.shortMessage !== '') return e.shortMessage
    if (typeof e.message === 'string' && e.message !== '') return e.message
  }
  return String(err)
}

type SendState =
  | { phase: 'idle' }
  | { phase: 'sending' }
  | { phase: 'sent'; hash: Hex }
  | { phase: 'error'; message: string }

function LiveReleaseShare({
  epoch,
  aid,
  ciphertexts,
  managerAddress,
  appManagerAddress,
}: Omit<ReleaseShareProps, 'demo'>) {
  const { isConnected, address } = useAccount()
  const publicClient = usePublicClient()
  const { data: walletClient } = useWalletClient()

  const [index, setIndex] = useState(() => String(ciphertexts[0]?.index ?? ''))
  const [secret, setSecret] = useState(() => loadOrganizerSecret(epoch, aid)?.toString() ?? '')
  const [state, setState] = useState<SendState>({ phase: 'idle' })

  const selected = ciphertexts.find((row) => String(row.index) === index) ?? null

  async function release(): Promise<void> {
    if (!selected) {
      setState({ phase: 'error', message: 'Pick a ciphertext first.' })
      return
    }
    const skOrg = parseOrganizerSecret(secret)
    if (skOrg == null) {
      setState({ phase: 'error', message: 'The secret must be a non-zero integer, decimal or 0x-hex.' })
      return
    }
    if (!walletClient || !publicClient) {
      setState({ phase: 'error', message: 'Connect a wallet before sending the transaction.' })
      return
    }
    setState({ phase: 'sending' })
    try {
      // The SDK resolves its own copy of viem, so its client types are
      // nominally different from the ones wagmi hands us even when the runtime
      // objects are identical. One cast at the boundary, nowhere else.
      const writer = new DKGWriter({
        publicClient,
        walletClient,
        managerAddress,
        ...(appManagerAddress ? { appManagerAddress } : {}),
      } as unknown as ConstructorParameters<typeof DKGWriter>[0])
      // The store keeps the raw on-chain (RTE) words; the SDK works in TE.
      const c1 = fromRTEtoTE(selected.c1.x, selected.c1.y)
      const c2 = fromRTEtoTE(selected.c2.x, selected.c2.y)
      const hash = await writer.submitOrganizerShare(epoch, aid, selected.index, { c1, c2 }, skOrg)
      setState({ phase: 'sent', hash })
    } catch (err) {
      setState({ phase: 'error', message: reason(err) })
    }
  }

  return (
    <div className='flex flex-col gap-4'>
      <ShareFields
        ciphertexts={ciphertexts}
        index={index}
        onIndex={(value) => {
          setIndex(value)
          setState({ phase: 'idle' })
        }}
        secret={secret}
        onSecret={(value) => {
          setSecret(value)
          setState({ phase: 'idle' })
        }}
        disabled={state.phase === 'sending'}
      />

      <div className='flex flex-wrap items-center gap-3'>
        <Button
          variant='primary'
          size='md'
          loading={state.phase === 'sending'}
          disabled={!isConnected || ciphertexts.length === 0}
          onClick={() => {
            void release()
          }}
        >
          Release organizer share
        </Button>
        {selected?.share.present ? (
          <span className='text-[12px] text-ash'>
            A share is already on chain for #{selected.index}; sending again overwrites it.
          </span>
        ) : null}
        {address ? <span className='font-mono text-[11px] text-ash'>relaying from {address}</span> : null}
      </div>

      {!isConnected ? (
        <Callout tone='info' title='Connect a wallet'>
          The share is computed here, but someone has to pay for the transaction. Connect a wallet with the top-right
          button — it does not have to be the organizer’s: relaying a share is permissionless, the DLEQ is what binds it
          to PK_org.
        </Callout>
      ) : null}
      {state.phase === 'sent' ? (
        <Callout tone='ok' title='Share submitted'>
          <span className='flex items-center gap-2'>
            Transaction <TxCell hash={state.hash} copy />
          </span>
          Once t partials are in, anyone can call <code>combineDecryption</code> and the plaintext lands on chain.
        </Callout>
      ) : null}
      {state.phase === 'error' ? (
        <Callout tone='danger' title='The share was not submitted'>
          {state.message}
        </Callout>
      ) : null}
    </div>
  )
}

function DemoReleaseShare({ epoch, aid, ciphertexts }: Pick<ReleaseShareProps, 'epoch' | 'aid' | 'ciphertexts'>) {
  const [index, setIndex] = useState(() => String(ciphertexts[0]?.index ?? ''))
  const secret = loadOrganizerSecret(epoch, aid)?.toString() ?? ''
  return (
    <div className='flex flex-col gap-4'>
      <ShareFields
        ciphertexts={ciphertexts}
        index={index}
        onIndex={setIndex}
        secret={secret}
        onSecret={() => {}}
        disabled
      />
      <Button variant='primary' size='md' disabled>
        Release organizer share
      </Button>
      <Callout tone='info' title='Demo mode: no chain to send to'>
        This explorer is running entirely off the synthetic fixture, so there is no <code>DKGAppManager</code> to
        receive <code>submitOrganizerShare</code> and no wallet to sign it. Point the explorer at a deployment (drop the{' '}
        <code>?demo=1</code>) and connect the organizer’s wallet to release a real share.
      </Callout>
    </div>
  )
}

/** Release the organizer share for one ciphertext of this application. */
export function ReleaseSharePanel(props: ReleaseShareProps) {
  const { demo, ...rest } = props
  if (demo) return <DemoReleaseShare epoch={rest.epoch} aid={rest.aid} ciphertexts={rest.ciphertexts} />
  return <LiveReleaseShare {...rest} />
}
