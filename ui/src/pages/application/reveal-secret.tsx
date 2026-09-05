// Organizer tools: reveal `sk_org` for an organizer-locked application.
//
// One transaction, once, and irreversible: `revealOrganizerSecret` stores the
// scalar after checking `sk_org·G == PK_org`, and from then on every ciphertext
// of the application — past and future — combines on `t` partials alone. The
// same check runs here first, so a mistyped secret never costs gas. The
// relayer does not have to be the organizer: the call is permissionless and
// the secret is what authenticates it.

import { useState } from 'react'
import { useAccount, usePublicClient, useWalletClient } from 'wagmi'
import { DKGWriter } from '@vocdoni/davinci-dkg-sdk'
import { Button, Callout, Input, TxCell } from '~kit'
import type { Address, Aid, EpochId, Hex, Point } from '~indexer/types'
import { loadOrganizerSecret, parseOrganizerSecret } from '~lib/organizer-secret'
import { matchesOrganizerKey } from './keys'

export interface RevealSecretProps {
  epoch: EpochId
  aid: Aid
  /** The registered `PK_org`, TE form, to check the secret against locally. */
  organizerPK: Point
  ciphertexts: number
  managerAddress: Address
  appManagerAddress: Address | null
  /** Demo mode has no chain to send to; the panel explains itself instead. */
  demo: boolean
}

interface FieldProps {
  secret: string
  onSecret: (value: string) => void
  disabled?: boolean
  error?: string
}

function SecretField({ secret, onSecret, disabled, error }: FieldProps) {
  return (
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
      error={error}
      onChange={(e) => onSecret(e.target.value)}
      hint='Checked against PK_org in this tab before anything is sent. Prefilled from this tab’s playground session when one exists.'
    />
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

function LiveRevealSecret({
  epoch,
  aid,
  organizerPK,
  ciphertexts,
  managerAddress,
  appManagerAddress,
}: Omit<RevealSecretProps, 'demo'>) {
  const { isConnected, address } = useAccount()
  const publicClient = usePublicClient()
  const { data: walletClient } = useWalletClient()

  const [secret, setSecret] = useState(() => loadOrganizerSecret(epoch, aid)?.toString() ?? '')
  const [state, setState] = useState<SendState>({ phase: 'idle' })

  async function reveal(): Promise<void> {
    const skOrg = parseOrganizerSecret(secret)
    if (skOrg == null) {
      setState({ phase: 'error', message: 'The secret must be a non-zero integer, decimal or 0x-hex.' })
      return
    }
    if (!matchesOrganizerKey(skOrg, organizerPK)) {
      setState({ phase: 'error', message: 'sk_org·G is not the registered PK_org — the contract would revert InvalidOrganizerSecret().' })
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
      const hash = await writer.revealOrganizerSecret(epoch, aid, skOrg)
      setState({ phase: 'sent', hash })
    } catch (err) {
      setState({ phase: 'error', message: reason(err) })
    }
  }

  return (
    <div className='flex flex-col gap-4'>
      <div className='max-w-lg'>
        <SecretField
          secret={secret}
          onSecret={(value) => {
            setSecret(value)
            setState({ phase: 'idle' })
          }}
          disabled={state.phase === 'sending' || state.phase === 'sent'}
        />
      </div>

      <div className='flex flex-wrap items-center gap-3'>
        <Button
          variant='primary'
          size='md'
          loading={state.phase === 'sending'}
          disabled={!isConnected || state.phase === 'sent'}
          onClick={() => {
            void reveal()
          }}
        >
          Reveal organizer secret
        </Button>
        <span className='text-[12px] text-ash'>
          once, irreversibly — unlocks {ciphertexts} ciphertext{ciphertexts === 1 ? '' : 's'} and every later one
        </span>
        {address ? <span className='font-mono text-[11px] text-ash'>relaying from {address}</span> : null}
      </div>

      {!isConnected ? (
        <Callout tone='info' title='Connect a wallet'>
          The secret is checked here, but someone has to pay for the transaction. Connect a wallet with the top-right
          button — it does not have to be the organizer’s: revealing is permissionless, the secret itself is what the
          contract verifies against PK_org.
        </Callout>
      ) : null}
      {state.phase === 'sent' ? (
        <Callout tone='ok' title='Secret revealed'>
          <span className='flex items-center gap-2'>
            Transaction <TxCell hash={state.hash} copy />
          </span>
          From here on the committee combines on its own: once t partials are in, anyone can call{' '}
          <code>combineDecryption</code> and the plaintext lands on chain.
        </Callout>
      ) : null}
      {state.phase === 'error' ? (
        <Callout tone='danger' title='The secret was not revealed'>
          {state.message}
        </Callout>
      ) : null}
    </div>
  )
}

function DemoRevealSecret({ epoch, aid }: Pick<RevealSecretProps, 'epoch' | 'aid'>) {
  const secret = loadOrganizerSecret(epoch, aid)?.toString() ?? ''
  return (
    <div className='flex flex-col gap-4'>
      <div className='max-w-lg'>
        <SecretField secret={secret} onSecret={() => {}} disabled />
      </div>
      <Button variant='primary' size='md' disabled>
        Reveal organizer secret
      </Button>
      <Callout tone='info' title='Demo mode: no chain to send to'>
        This explorer is running entirely off the synthetic fixture, so there is no <code>DKGAppManager</code> to
        receive <code>revealOrganizerSecret</code> and no wallet to sign it. Point the explorer at a deployment (drop
        the <code>?demo=1</code>) and connect a wallet to reveal a real secret.
      </Callout>
    </div>
  )
}

/** Reveal `sk_org` for this organizer-locked application. */
export function RevealSecretPanel(props: RevealSecretProps) {
  const { demo, ...rest } = props
  if (demo) return <DemoRevealSecret epoch={rest.epoch} aid={rest.aid} />
  return <LiveRevealSecret {...rest} />
}
