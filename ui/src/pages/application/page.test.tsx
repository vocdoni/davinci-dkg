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

    // PK_aid = PK_ep + PK_org, rendered as its two coordinates.
    const epoch = store.epochs[app.epoch.toLowerCase()]
    const pkAid = applicationPublicKey(epoch.collectivePublicKey, app.organizerPK)
    expect(pkAid).not.toBeNull()
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

  it('disables the organizer tools in demo mode and says why', () => {
    const { source, app } = harness()
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${app.aid}`)

    expect(screen.getByRole('heading', { name: 'Release organizer share' })).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'Release organizer share' })).toBeDisabled()
    expect(screen.getByText('Demo mode: no chain to send to')).toBeInTheDocument()
    expect(screen.getByLabelText('Organizer secret')).toBeDisabled()
  })

  it('explains an application that does not exist', () => {
    const { source, app } = harness()
    const aid = `0x${'11'.repeat(32)}`
    renderPage(source, <ApplicationPage />, `/applications/${app.epoch}/${aid}`)

    expect(screen.getByText('Not found')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: 'Browse all applications' })).toHaveAttribute('href', '/applications')
  })
})
