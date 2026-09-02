import { describe, expect, it } from 'vitest'
import { screen } from '@testing-library/react'
import type { ReactElement } from 'react'
import { Route, Routes } from 'react-router-dom'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import type { DataSource } from '~data/source'
import type { IndexerStore } from '~indexer/types'
import { checksum } from '~kit'
import { renderWithProviders } from '../../test-utils'
import { OperatorPage } from '../operator'

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

function harness(): { source: DataSource; store: IndexerStore } {
  const source = createDemoDataSource(FIXTURE)
  return { source, store: source.getSnapshot().store }
}

function renderPage(source: DataSource, ui: ReactElement, route: string) {
  return renderWithProviders(
    <DataSourceProvider source={source}>
      <Routes>
        <Route path='/operators/:address' element={ui} />
      </Routes>
    </DataSourceProvider>,
    { route }
  )
}

/** A committee member of the first epoch that actually assembled one. */
function busyOperator(store: IndexerStore): string {
  for (const key of store.epochOrder) {
    const epoch = store.epochs[key]
    if (epoch.committee.length > 0) return epoch.committee[0]
  }
  throw new Error('fixture has no committee')
}

describe('OperatorPage', () => {
  it('shows the identity, the per-epoch history and the event log', () => {
    const { source, store } = harness()
    const address = busyOperator(store)
    renderPage(source, <OperatorPage />, `/operators/${address}`)

    expect(screen.getByText('Registry record')).toBeInTheDocument()
    // Addresses are always rendered checksummed, whatever the log casing was.
    expect(screen.getAllByTitle(checksum(address)).length).toBeGreaterThan(0)

    expect(screen.getByText('Per-epoch record')).toBeInTheDocument()
    const events = store.operators[address.toLowerCase()].events.length
    expect(events).toBeGreaterThan(0)
    expect(screen.getByText(`${events} events attributed to this address, newest first.`)).toBeInTheDocument()
  })

  it('links every epoch it served from the history table', () => {
    const { source, store } = harness()
    const address = busyOperator(store)
    renderPage(source, <OperatorPage />, `/operators/${address}`)

    const links = screen.getAllByRole('link').filter((link) => link.getAttribute('href')?.startsWith('/epochs/'))
    expect(links.length).toBeGreaterThan(0)
  })

  it('shows the participation trend and the counters', () => {
    const { source, store } = harness()
    const address = busyOperator(store)
    renderPage(source, <OperatorPage />, `/operators/${address}`)

    expect(screen.getByText('Participation over epochs')).toBeInTheDocument()
    expect(screen.getByLabelText('participation trend')).toBeInTheDocument()
    expect(screen.getByText(/^participation (\d+%|—)$/)).toBeInTheDocument()
    expect(screen.getByLabelText('partials per epoch')).toBeInTheDocument()
  })

  it('explains an unknown address instead of rendering an empty page', () => {
    const { source } = harness()
    renderPage(source, <OperatorPage />, '/operators/0x000000000000000000000000000000000000dead')

    expect(screen.getByText('No such operator')).toBeInTheDocument()
    expect(screen.getByText('0x000000000000000000000000000000000000dead')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: 'Browse the registry' })).toHaveAttribute('href', '/operators')
  })
})
