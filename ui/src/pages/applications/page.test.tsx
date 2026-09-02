import { describe, expect, it } from 'vitest'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import type { ReactElement } from 'react'
import { Route, Routes } from 'react-router-dom'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import type { DataSource } from '~data/source'
import type { ApplicationEntity, IndexerStore } from '~indexer/types'
import { renderWithProviders } from '../../test-utils'
import { ApplicationsPage } from '../applications'

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

function harness(): { source: DataSource; store: IndexerStore; apps: ApplicationEntity[] } {
  const source = createDemoDataSource(FIXTURE)
  const store = source.getSnapshot().store
  return { source, store, apps: store.applicationOrder.map((key) => store.applications[key]) }
}

function renderPage(source: DataSource, ui: ReactElement, route = '/applications') {
  return renderWithProviders(
    <DataSourceProvider source={source}>
      <Routes>
        <Route path='/applications' element={ui} />
      </Routes>
    </DataSourceProvider>,
    { route }
  )
}

describe('ApplicationsPage', () => {
  it('lists every application across every epoch, with the pipeline totals', () => {
    const { source, apps } = harness()
    expect(apps.length).toBeGreaterThan(0)
    renderPage(source, <ApplicationsPage />)

    expect(screen.getByText(`${apps.length} applications`)).toBeInTheDocument()
    expect(screen.getByText('Pending shares')).toBeInTheDocument()
    expect(screen.getByText('ciphertexts still waiting on an organizer share')).toBeInTheDocument()

    const links = screen.getAllByRole('link').filter((a) => a.getAttribute('href')?.startsWith('/applications/'))
    expect(links.length).toBe(apps.length)
  })

  it('reads the epoch filter out of the URL', () => {
    const { source, apps } = harness()
    const epoch = apps[0].epoch
    const inEpoch = apps.filter((app) => app.epoch === epoch).length
    renderPage(source, <ApplicationsPage />, `/applications?epoch=${epoch}`)

    expect(screen.getByText(`${inEpoch} of ${apps.length} applications`)).toBeInTheDocument()
  })

  it('filters by application id prefix', async () => {
    const user = userEvent.setup()
    const { source, apps } = harness()
    renderPage(source, <ApplicationsPage />)

    await user.type(screen.getByLabelText('Search applications'), apps[0].aid.slice(0, 12))
    expect(screen.getByText(`1 of ${apps.length} applications`)).toBeInTheDocument()
  })

  it('explains an empty result and offers the store hit for a foreign query', async () => {
    const user = userEvent.setup()
    const { source, store, apps } = harness()
    renderPage(source, <ApplicationsPage />)

    const operator = store.operatorOrder[0]
    await user.type(screen.getByLabelText('Search applications'), operator)
    expect(screen.getByText(`0 of ${apps.length} applications`)).toBeInTheDocument()
    expect(screen.getByText('No application matches')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: new RegExp(operator.slice(0, 10), 'i') })).toHaveAttribute(
      'href',
      `/operators/${operator}`
    )
  })
})
