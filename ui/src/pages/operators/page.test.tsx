import { describe, expect, it } from 'vitest'
import { screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import type { ReactElement } from 'react'
import { Route, Routes } from 'react-router-dom'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import type { DataSource } from '~data/source'
import type { IndexerStore } from '~indexer/types'
import { renderWithProviders } from '../../test-utils'
import { OperatorsPage } from '../operators'

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

function renderPage(source: DataSource, ui: ReactElement, route = '/operators') {
  return renderWithProviders(
    <DataSourceProvider source={source}>
      <Routes>
        <Route path='/operators' element={ui} />
      </Routes>
    </DataSourceProvider>,
    { route }
  )
}

describe('OperatorsPage', () => {
  it('shows the registry header cards and the full roster', () => {
    const { source, store } = harness()
    renderPage(source, <OperatorsPage />)

    expect(screen.getByText('24 operators')).toBeInTheDocument()
    expect(screen.getByText('Active / registered')).toBeInTheDocument()
    expect(screen.getByText('Inactivity window')).toBeInTheDocument()
    expect(screen.getByText('Newest committee')).toBeInTheDocument()

    const newest = store.epochs[store.epochOrder[store.epochOrder.length - 1]]
    expect(screen.getByText(`epoch #${newest.nonce}`, { exact: false })).toBeInTheDocument()
  })

  it('filters the table by address prefix', async () => {
    const user = userEvent.setup()
    const { source, store } = harness()
    renderPage(source, <OperatorsPage />)

    const address = store.operatorOrder[3]
    await user.type(screen.getByLabelText('Search operators by address'), address.slice(0, 10))
    expect(screen.getByText('1 of 24 operators')).toBeInTheDocument()
  })

  it('says so when nothing matches, rather than showing an empty table', async () => {
    const user = userEvent.setup()
    const { source } = harness()
    renderPage(source, <OperatorsPage />)

    await user.type(screen.getByLabelText('Search operators by address'), '0xdeadbeefcafe')
    expect(screen.getByText('0 of 24 operators')).toBeInTheDocument()
    expect(screen.getByText('No operator matches')).toBeInTheDocument()
  })

  it('offers the store search hit when the query is an epoch id, not an address', async () => {
    const user = userEvent.setup()
    const { source, store } = harness()
    renderPage(source, <OperatorsPage />)

    const epoch = store.epochs[store.epochOrder[0]]
    await user.type(screen.getByLabelText('Search operators by address'), epoch.id)
    expect(screen.getByText('elsewhere in the explorer')).toBeInTheDocument()
    const link = screen.getByRole('link', { name: new RegExp(`Epoch #${epoch.nonce}`) })
    expect(link).toHaveAttribute('href', `/epochs/${epoch.id}`)
  })

  it('filters by status', async () => {
    const user = userEvent.setup()
    const { source, store } = harness()
    renderPage(source, <OperatorsPage />)

    const inactive = Object.values(store.operators).filter((operator) => operator.status === 'inactive').length
    expect(inactive).toBeGreaterThan(0)
    await user.selectOptions(screen.getByLabelText('Filter by status'), 'inactive')
    expect(screen.getByText(`${inactive} of 24 operators`)).toBeInTheDocument()
  })
})
