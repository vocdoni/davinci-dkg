import { describe, expect, it } from 'vitest'
import { screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { renderWithProviders } from '../../test-utils'
import { saveOrganizerSecret } from '~lib/organizer-secret'
import { PlaygroundPage } from '../playground'
import { saveSession } from './session'
import { initialState } from './machine'

/**
 * Demo mode is the only mode a unit test can drive: there is no wallet, no
 * RPC and no wagmi provider here. That it renders at all is therefore also the
 * assertion that the demo tree never reaches for one.
 */
function renderPlayground() {
  return renderWithProviders(<PlaygroundPage />, { config: { demo: true }, route: '/playground' })
}

describe('playground (demo)', () => {
  it('opens on the connect step and offers the simulated wallet', async () => {
    renderPlayground()
    expect(await screen.findByRole('heading', { name: /connect a wallet/i })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: /use the demo wallet/i })).toBeInTheDocument()
  })

  it('walks the whole flow with no chain behind it', async () => {
    const user = userEvent.setup()
    renderPlayground()

    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))

    // Connecting unlocks the epoch step, which defaults to the newest Live epoch.
    expect(await screen.findByRole('heading', { name: /choose a live epoch/i })).toBeInTheDocument()
    await user.click(screen.getByRole('button', { name: 'Register an application →' }))

    // The identity is generated for us, and the secret comes with its warning.
    const heading = await screen.findByRole('heading', { name: /register an application/i })
    expect(heading).toBeInTheDocument()
    expect(await screen.findByText(/copy this now — it is shown once/i)).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /^register application$/i }))
    expect(await screen.findByRole('heading', { name: /encrypt a value/i })).toBeInTheDocument()
    // The write is recorded in the log with its simulated transaction.
    expect(screen.getByText(/^Registered application/)).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /^encrypt$/i }))
    expect(await screen.findByRole('heading', { name: /submit the ciphertext/i })).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /submit ciphertext/i }))
    expect(await screen.findByRole('heading', { name: /reveal the organizer secret/i })).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /reveal the secret/i }))
    expect(await screen.findByRole('heading', { name: /watch the decryption/i })).toBeInTheDocument()
    expect(screen.getByText(/partial decryptions/i)).toBeInTheDocument()
    expect(screen.getByText(/^revealed · block/)).toBeInTheDocument()
    // The simulated committee only starts answering after the reveal, then the
    // combine lands and is written to the activity log.
    await waitFor(() => expect(screen.getByText(/combined on chain from \d+ partials/)).toBeInTheDocument(), {
      timeout: 15_000,
    })
  }, 60_000)

  it('claims a pool key and encrypts under the application key', async () => {
    const user = userEvent.setup()
    renderPlayground()

    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))
    // The epoch list says how much of each pool is left.
    expect((await screen.findAllByText(/\d+ activated free · \d+ claimed · \d+ not activated/)).length).toBeGreaterThan(0)
    await user.click(await screen.findByRole('button', { name: 'Register an application →' }))
    await user.click(await screen.findByRole('button', { name: /^register application$/i }))
    // The fixture's live epochs have two keys claimed, so ours is key 2. The
    // encrypt step names it next to PK_aid, and the log line records the claim.
    expect(await screen.findByText(/^Pool key 2 · PK_aid\.x/)).toBeInTheDocument()
    expect(screen.getByText(/— pool key 2$/)).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /^encrypt$/i }))
    expect(await screen.findByText(/^Encrypted 42 under PK_aid = P_2 \+ PK_org$/)).toBeInTheDocument()
  }, 60_000)

  it('skips the reveal step for an automatic application', async () => {
    const user = userEvent.setup()
    renderPlayground()

    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))
    await user.click(await screen.findByRole('button', { name: 'Register an application →' }))
    await screen.findByRole('heading', { name: /register an application/i })

    await user.selectOptions(screen.getByLabelText(/^mode$/i), 'automatic')
    // No organizer key: the secret block gives way to a note.
    expect(await screen.findByText('No organizer key')).toBeInTheDocument()
    expect(screen.queryByText(/copy this now — it is shown once/i)).not.toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /^register application$/i }))
    expect(await screen.findByRole('heading', { name: /encrypt a value/i })).toBeInTheDocument()
    expect(screen.getByText(/^Registered automatic application/)).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /^encrypt$/i }))
    expect(await screen.findByText(/^Encrypted 42 under PK_aid = P_2$/)).toBeInTheDocument()
    await user.click(await screen.findByRole('button', { name: /submit ciphertext/i }))
    // Straight to the watch step: there is no secret to reveal or keep, and
    // the rail says the reveal step was skipped.
    expect(await screen.findByRole('heading', { name: /watch the decryption/i })).toBeInTheDocument()
    expect(screen.getByText('not needed')).toBeInTheDocument()
    expect(screen.getByText('skipped')).toBeInTheDocument()
  }, 60_000)

  it('refuses a decryption window that is already closed', async () => {
    const user = userEvent.setup()
    renderPlayground()

    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))
    await user.click(await screen.findByRole('button', { name: 'Register an application →' }))
    await screen.findByRole('heading', { name: /register an application/i })
    await user.type(screen.getByLabelText(/decryption closes/i), '2001-01-01T00:00')
    await user.click(screen.getByRole('button', { name: /^register application$/i }))
    expect(await screen.findByText(/decryptNotAfter must be in the future/)).toBeInTheDocument()
    expect(screen.queryByText(/^Registered application/)).not.toBeInTheDocument()
  }, 60_000)

  it('pins the epoch in the URL once an application is registered', async () => {
    const user = userEvent.setup()
    renderPlayground()

    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))
    await user.click(await screen.findByRole('button', { name: 'Register an application →' }))
    await user.click(await screen.findByRole('button', { name: /^register application$/i }))

    await screen.findByText(/epoch pinned by the registration/i)
    // Going back to the epoch step must not offer a different epoch any more.
    await user.click(screen.getByRole('button', { name: /choose a live epoch/i }))
    expect(
      await screen.findByText(/registration binds the application to this epoch/i)
    ).toBeInTheDocument()
  }, 60_000)

  it('records every action in the activity log', async () => {
    const user = userEvent.setup()
    renderPlayground()
    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))
    await waitFor(() => expect(screen.getByText(/connected as demo wallet/i)).toBeInTheDocument())
  })

  it('prints the transcripts behind the advanced toggle', async () => {
    const user = userEvent.setup()
    renderPlayground()
    await user.click(await screen.findByRole('button', { name: /use the demo wallet/i }))
    await user.click(await screen.findByRole('button', { name: 'Register an application →' }))
    await user.click(screen.getByRole('switch'))
    expect(await screen.findByText(/registration transcript/i)).toBeInTheDocument()
    expect(screen.getByText('PK_org.x')).toBeInTheDocument()
  }, 60_000)

  it('resumes a walkthrough from a deep link plus the tab session', async () => {
    const user = userEvent.setup()
    const epochId = '0x0102030405060708090a0b0c' as const
    const aid = `0x1c${'ab'.repeat(31)}` as `0x${string}`
    // What a previous visit to this tab would have left behind.
    saveOrganizerSecret(epochId, aid, 12345678901234567890n)
    saveSession(epochId, aid, {
      ...initialState(),
      value: '77',
      ciphertext: { c1: ['1', '2'], c2: ['3', '4'] },
      ciphertextIndex: 2,
      reveal: 'revealed',
      poolIndex: 2,
    })

    renderWithProviders(<PlaygroundPage />, {
      config: { demo: true },
      route: `/playground?epoch=${epochId}&aid=${aid}`,
    })

    expect(await screen.findByText(/resumed application/i)).toBeInTheDocument()
    // One click to re-attach a wallet, then straight to where we left off.
    await user.click(screen.getByRole('button', { name: /use the demo wallet/i }))
    expect(await screen.findByRole('heading', { name: /watch the decryption/i })).toBeInTheDocument()
  }, 30_000)
})
