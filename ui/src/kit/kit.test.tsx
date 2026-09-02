import { describe, expect, it } from 'vitest'
import { screen, within } from '@testing-library/react'
import { renderWithProviders } from '../test-utils'
import { Address, Badge, DataTable, Input, ProgressBar, type AnyColumnDef } from '~kit'

const LOWER = '0x3f9b338706a31f26d49159478015c8aaeab908ad'
const CHECKSUMMED = '0x3F9B338706a31f26D49159478015C8AAEAb908Ad'

describe('Address', () => {
  it('checksums and truncates, keeping the full value reachable', () => {
    renderWithProviders(<Address value={LOWER} />)
    const el = screen.getByTitle(CHECKSUMMED)
    expect(el).toHaveTextContent('0x3F9B…08Ad')
  })

  it('renders the full address when asked', () => {
    renderWithProviders(<Address value={LOWER} full copy={false} explorer={false} />)
    expect(screen.getByText(CHECKSUMMED)).toBeInTheDocument()
  })

  it('links to the configured block explorer', () => {
    renderWithProviders(<Address value={LOWER} />, { config: { explorerUrl: 'https://sepolia.etherscan.io' } })
    expect(screen.getByRole('link')).toHaveAttribute('href', `https://sepolia.etherscan.io/address/${CHECKSUMMED}`)
  })

  it('omits the explorer link when the deployment has no explorer', () => {
    renderWithProviders(<Address value={LOWER} />, { config: { explorerUrl: undefined } })
    expect(screen.queryByRole('link')).toBeNull()
  })
})

describe('ProgressBar', () => {
  it('exposes the t of n values to assistive tech', () => {
    renderWithProviders(<ProgressBar value={31} total={64} threshold={33} label='contributions' />)
    const bar = screen.getByRole('progressbar')
    expect(bar).toHaveAttribute('aria-valuenow', '31')
    expect(bar).toHaveAttribute('aria-valuemax', '64')
  })

  it('marks the threshold', () => {
    renderWithProviders(<ProgressBar value={10} total={64} threshold={33} />)
    expect(screen.getByTitle('threshold 33')).toBeInTheDocument()
  })
})

describe('Badge', () => {
  it('renders its label', () => {
    renderWithProviders(<Badge tone='ok'>live</Badge>)
    expect(screen.getByText('live')).toBeInTheDocument()
  })
})

interface Row {
  id: string
  count: number
}

const columns: AnyColumnDef<Row>[] = [
  { id: 'id', header: 'Id', accessorKey: 'id' },
  { id: 'count', header: 'Count', accessorKey: 'count', meta: { numeric: true } },
]

describe('DataTable', () => {
  const rows: Row[] = [
    { id: 'b', count: 2 },
    { id: 'a', count: 9 },
  ]

  it('renders headers and rows', () => {
    renderWithProviders(<DataTable data={rows} columns={columns} />)
    expect(screen.getByRole('columnheader', { name: /Id/ })).toBeInTheDocument()
    expect(screen.getByText('9')).toBeInTheDocument()
  })

  it('applies the initial sort', () => {
    renderWithProviders(<DataTable data={rows} columns={columns} initialSorting={[{ id: 'count', desc: true }]} />)
    const [firstRow] = screen.getAllByRole('row').slice(1)
    expect(within(firstRow as HTMLElement).getByText('9')).toBeInTheDocument()
  })

  it('shows an empty state instead of a bare header', () => {
    renderWithProviders(<DataTable data={[]} columns={columns} />)
    expect(screen.getByText('Nothing here yet')).toBeInTheDocument()
  })

  it('shows skeleton rows while loading', () => {
    const { container } = renderWithProviders(<DataTable data={[]} columns={columns} loading loadingRows={3} />)
    expect(container.querySelectorAll('.animate-skeleton').length).toBeGreaterThanOrEqual(3)
  })

  it('virtualises without rendering every row', () => {
    const many: Row[] = Array.from({ length: 300 }, (_, i) => ({ id: `row-${i}`, count: i }))
    renderWithProviders(<DataTable data={many} columns={columns} virtualized rowHeight={40} maxHeight={200} />)
    // jsdom reports a zero-height scroll element, so the virtualiser renders
    // only its overscan window — the point is that it is far below 300.
    expect(screen.queryAllByText(/^row-/).length).toBeLessThan(60)
  })
})

describe('Input', () => {
  it('associates a generated id with its label', () => {
    renderWithProviders(<Input label='Application id' />)
    expect(screen.getByLabelText('Application id')).toBeInstanceOf(HTMLInputElement)
  })

  it('keeps a caller-supplied id', () => {
    renderWithProviders(<Input id='aid' label='Application id' />)
    expect(screen.getByLabelText('Application id')).toHaveAttribute('id', 'aid')
  })
})
