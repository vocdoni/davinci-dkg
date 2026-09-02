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
    expect(await screen.findByRole('heading', { name: /release or withhold/i })).toBeInTheDocument()

    await user.click(screen.getByRole('button', { name: /release the share/i }))
    expect(await screen.findByRole('heading', { name: /watch the decryption/i })).toBeInTheDocument()
    expect(screen.getByText(/partial decryptions/i)).toBeInTheDocument()
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
      share: 'released',
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
