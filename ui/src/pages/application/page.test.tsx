import { describe, expect, it } from 'vitest'
import { screen } from '@testing-library/react'
import type { ReactElement } from 'react'
import { Route, Routes } from 'react-router-dom'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import type { DataSource } from '~data/source'
import type { ApplicationEntity, IndexerStore } from '~indexer/types'
import { checksum } from '~kit'
import { renderWithProviders } from '../../test-utils'
import { ApplicationPage } from '../application'
import { applicationPublicKey, formatPointPair } from './keys'

const FIXTURE = {
  operators: 24,
  epochs: 4,
  committeeSize: 6,
  threshold: 3,
  minValidContributions: 4,
  applicationsPerEpoch: 1,
  ciphertextsPerApplication: 4,
  blockIntervalMs: 0,
}

function harness(): { source: DataSource; store: IndexerStore; app: ApplicationEntity } {
  const source = createDemoDataSource(FIXTURE)
  const store = source.getSnapshot().store
  const app = store.applications[store.applicationOrder[0]]
  return { source, store, app }
}

/** The one organizer-locked application whose organizer has not revealed. */
function lockedApp(store: IndexerStore): ApplicationEntity {
  const app = store.applicationOrder
    .map((key) => store.applications[key])
    .find((a) => a.mode === 'organizer-locked' && a.organizerSecret == null)
  if (!app) throw new Error('the fixture produced no unrevealed application')
  return app
}

/** Demo mode is the config here: the organizer panel must not reach for wagmi. */
function renderPage(source: DataSource, ui: ReactElement, route: string) {
  return renderWithProviders(
    <DataSourceProvider source={source}>
      <Routes>
        <Route path='/applications/:epoch/:aid' element={ui} />
      </Routes>
    </DataSourceProvider>,
    { route, config: { demo: true } }
  )
}

describe('ApplicationPage', () => {
  it('shows the record, the derived PK_aid and the ciphertext pipeline', () => {
    const { source, store, app } = harness()
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)

    expect(screen.getByText('On-chain application record')).toBeInTheDocument()
    expect(screen.getAllByTitle(checksum(app.creator)).length).toBeGreaterThan(0)

    // PK_aid = P_j + PK_org, rendered as its two coordinates.
    const epoch = store.epochs[app.epoch.toLowerCase()]
    expect(app.poolIndex).toBe(0)
    const pkAid = applicationPublicKey(epoch.poolKeys[app.poolIndex!].key, app.organizerPK)
    expect(pkAid).not.toBeNull()
    expect(screen.getByText('key 0')).toBeInTheDocument()
    const [x, y] = formatPointPair(pkAid!).split(',')
    expect(screen.getAllByTitle(x).length).toBeGreaterThan(0)
    expect(screen.getAllByTitle(y).length).toBeGreaterThan(0)

    // One row per ciphertext, indices as the contract assigned them.
    expect(screen.getByRole('heading', { name: 'Ciphertexts' })).toBeInTheDocument()
    for (let i = 1; i <= app.ciphertexts.length; i++) {
      expect(screen.getAllByText(String(i)).length).toBeGreaterThan(0)
    }
  })

  it('draws the partial matrix and links back to the playground', () => {
    const { source, app } = harness()
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)

    expect(screen.getByLabelText('Partial decryption matrix')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: 'Resume in playground' })).toHaveAttribute(
      'href',
      `/playground?epoch=${app.epoch}&aid=${app.aid}`
    )
  })

  it('disables the reveal in demo mode and says why', () => {
    const { source, store } = harness()
    const app = lockedApp(store)
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)

    expect(screen.getByRole('heading', { name: 'Reveal organizer secret' })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'Reveal organizer secret' })).toBeDisabled()
    expect(screen.getByText('Demo mode: no chain to send to')).toBeInTheDocument()
    expect(screen.getByLabelText('Organizer secret')).toBeDisabled()
    // No partial exists before the reveal — the contract refuses them — so
    // every ciphertext waits on the secret.
    expect(screen.getAllByText('kept').length).toBeGreaterThan(0)
    expect(screen.getAllByText('awaiting-reveal')).toHaveLength(app.ciphertexts.length)
    expect(screen.getByText('partials and combines refused until the reveal')).toBeInTheDocument()
  })

  it('shows the policy and the reveal of an organizer-locked application', () => {
    const { source, app } = harness()
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)

    expect(app.mode).toBe('organizer-locked')
    expect(screen.getByText('organizer-locked')).toBeInTheDocument()
    expect(screen.getByText('Submission policy')).toBeInTheDocument()
    // The fixture allow-lists the one address that submitted the ciphertexts.
    expect(screen.getAllByTitle(checksum(app.policy!.submitters[0])).length).toBeGreaterThan(0)
    // The window opened in 2023 and never closes: rendered as dates, and open;
    // the block window is unbounded and reads "any block".
    expect(screen.getByText('Decryption window')).toBeInTheDocument()
    expect(screen.getAllByText('open').length).toBeGreaterThanOrEqual(1)
    expect(screen.getByText(/no deadline/)).toBeInTheDocument()
    expect(screen.getByText('any block')).toBeInTheDocument()
    expect(screen.queryByText('closed')).not.toBeInTheDocument()
    // This organizer has revealed: the secret is shown and the tool is gone.
    expect(app.organizerSecret).not.toBeNull()
    expect(screen.getAllByText('revealed').length).toBeGreaterThan(0)
    expect(screen.getByText(`sk_org = ${app.organizerSecret!.toString()}`)).toBeInTheDocument()
    expect(screen.getByRole('heading', { name: 'Organizer secret revealed' })).toBeInTheDocument()
    expect(screen.queryByRole('heading', { name: 'Reveal organizer secret' })).not.toBeInTheDocument()
  })

  it('marks a closed decryption window', () => {
    const source = createDemoDataSource({ ...FIXTURE, epochs: 5 })
    const store = source.getSnapshot().store
    // Organizer-locked applications on even epochs closed their window in 2025.
    const app = store.applicationOrder
      .map((key) => store.applications[key])
      .find((a) => a.mode === 'organizer-locked' && store.epochs[a.epoch].nonce % 2 === 0)!
    expect(app).toBeDefined()
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)
    expect(screen.getAllByText('closed').length).toBeGreaterThan(0)
  })

  it('shows an automatic application with no organizer key and no organizer step', () => {
    const source = createDemoDataSource({ ...FIXTURE, applicationsPerEpoch: 2 })
    const store = source.getSnapshot().store
    const app = store.applicationOrder.map((key) => store.applications[key]).find((a) => a.mode === 'automatic')!
    expect(app).toBeDefined()
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)

    expect(screen.getAllByText('automatic').length).toBeGreaterThan(0)
    expect(screen.getByText('key 1')).toBeInTheDocument()
    expect(screen.queryByText('PK_org x')).not.toBeInTheDocument()
    expect(screen.queryByText('Organizer secret')).not.toBeInTheDocument()
    // "open" submission, plus the open decryption window badge (hence getAll).
    expect(screen.getAllByText('open').length).toBeGreaterThanOrEqual(2)
    // A deadline in 2033 renders as a date, not as closed.
    expect(screen.queryByText('closed')).not.toBeInTheDocument()
    expect(screen.getByRole('heading', { name: 'No organizer step' })).toBeInTheDocument()
    expect(screen.queryByRole('heading', { name: 'Reveal organizer secret' })).not.toBeInTheDocument()
  })

  it('explains an application that does not exist', () => {
    const { source, app } = harness()
    const aid = `0x${'11'.repeat(32)}`
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${aid}`)

    expect(screen.getByText('Not found')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: 'Browse all applications' })).toHaveAttribute('href', '/applications')
  })
})
