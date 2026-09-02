import { useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useIndexer, useNetworkStats, useOperators } from '~data/hooks'
import {
  Callout,
  Card,
  DataTable,
  EmptyState,
  Input,
  Panel,
  SearchIcon,
  SectionHeader,
  Select,
  Stack,
  StatCell,
  StatRow,
} from '~kit'
import { Donut, StackedBars, formatCompact } from '~kit/charts'
import { blocksToDuration } from '~lib/format'
import { paths } from '~routes/paths'
import { operatorColumns } from './operators/columns'
import { operatorStatusSlices, operatorWorkChart, WORK_SERIES } from './operators/charts'
import { filterOperators, STATUS_OPTIONS, type OperatorStatusFilter } from './operators/filter'
import { StoreSearchHits } from './operators/store-search'

/**
 * The registry, at registry scale. Everything on this page is derived from the
 * one `useOperators()` selector; the table is virtualised so 300 rows cost the
 * same as 30, and the two charts answer the questions the table cannot: who is
 * doing the work, and how much of the set is still alive.
 */
export function OperatorsPage() {
  const rows = useOperators()
  const stats = useNetworkStats()
  const { scanning, status: indexerStatus } = useIndexer()
  const navigate = useNavigate()

  const [query, setQuery] = useState('')
  const [status, setStatus] = useState<OperatorStatusFilter>('all')

  const filtered = useMemo(() => filterOperators(rows, { query, status }), [rows, query, status])
  const chart = useMemo(() => operatorWorkChart(rows), [rows])
  const slices = useMemo(() => operatorStatusSlices(rows), [rows])
  const columns = useMemo(() => operatorColumns({ blockTimeSeconds: stats.blockTimeSeconds }), [stats.blockTimeSeconds])

  const reapable = useMemo(() => rows.filter((row) => row.reapable).length, [rows])
  const inactive = useMemo(() => rows.filter((row) => row.status === 'inactive').length, [rows])
  const newest = stats.newestEpoch
  const loading = scanning && rows.length === 0
  // Only when there is nothing to show: an empty registry and a failed scan
  // look identical otherwise.
  const failure =
    rows.length === 0 && !scanning ? (indexerStatus.errors[indexerStatus.errors.length - 1] ?? null) : null

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Registry'
        title='Operators'
        description='Every registered operator, what it has done and how reliably it does it. Counters are the whole indexed history; participation is contributions over claimed slots.'
      />

      {failure ? (
        <Callout tone='danger' title='The indexer could not read the chain'>
          {failure.scope}: {failure.message}
        </Callout>
      ) : null}

      <StatRow>
        <StatCell
          label='Active / registered'
          value={`${stats.operatorsActive} / ${stats.operatorsRegistered}`}
          mono
          tone='accent'
          hint={`${inactive} inactive · ${reapable} past the inactivity window`}
          loading={loading}
        />
        <StatCell
          label='Inactivity window'
          value={stats.inactivityWindow != null ? formatCompact(stats.inactivityWindow) : '—'}
          mono
          hint={
            stats.inactivityWindow != null
              ? `blocks of silence before a reap (${blocksToDuration(stats.inactivityWindow, stats.blockTimeSeconds)})`
              : 'not read from the registry yet'
          }
          loading={loading}
        />
        <StatCell
          label='Newest committee'
          value={newest ? `${newest.committee.length || newest.policy?.committeeSize || 0}` : '—'}
          mono
          hint={
            newest
              ? `epoch #${newest.nonce} · t = ${newest.policy?.threshold ?? '?'} of n = ${newest.policy?.committeeSize ?? '?'}`
              : 'no epoch indexed yet'
          }
          loading={loading}
        />
        <StatCell
          label='Work indexed'
          value={formatCompact(stats.contributions)}
          mono
          hint={`contributions · ${formatCompact(stats.partials)} partials · ${formatCompact(stats.claims)} claims`}
          loading={loading}
        />
      </StatRow>

      <div className='grid gap-4 lg:grid-cols-3'>
        <Panel
          className='lg:col-span-2'
          label='Distribution'
          title='Work per operator'
          description={
            chart.grouped > 0
              ? `Top ${chart.data.length - 1} by contributions; the remaining ${chart.grouped} operators are summed into the last column.`
              : 'Every operator, by contributions.'
          }
        >
          <StackedBars
            data={chart.data}
            series={WORK_SERIES}
            height={240}
            loading={loading}
            onBarClick={(_, index) => {
              const address = chart.addresses[index]
              if (address) navigate(paths.operator(address))
            }}
          />
        </Panel>

        <Panel label='Liveness' title='Registry status'>
          <Donut
            slices={slices}
            loading={loading}
            centerValue={rows.length}
            centerLabel='registered'
            size={176}
            thickness={20}
          />
        </Panel>
      </div>

      <Card flush className='overflow-hidden'>
        <div className='flex flex-wrap items-end justify-between gap-4 border-b border-charcoal px-5 py-4'>
          <div className='min-w-0'>
            <div className='label-caps mb-1.5 text-emerald'>Registry</div>
            <h2 className='text-[15px] font-semibold text-ghost'>
              {filtered.length === rows.length
                ? `${rows.length} operators`
                : `${filtered.length} of ${rows.length} operators`}
            </h2>
          </div>
          <div className='flex flex-wrap items-end gap-3'>
            <Input
              size='sm'
              mono
              value={query}
              placeholder='0x… address prefix'
              aria-label='Search operators by address'
              iconLeft={<SearchIcon size={13} />}
              onChange={(e) => setQuery(e.target.value)}
              wrapperClassName='w-64'
            />
            <Select
              size='sm'
              aria-label='Filter by status'
              value={status}
              onChange={(e) => setStatus(e.target.value as OperatorStatusFilter)}
              options={STATUS_OPTIONS}
              wrapperClassName='w-48'
            />
          </div>
        </div>

        {query.trim() !== '' ? <StoreSearchHits query={query} skip={['operator']} className='m-5 mb-0' /> : null}

        <div className='overflow-x-auto scroll-slim'>
          <div className='min-w-[1250px]'>
            <DataTable
              data={filtered}
              columns={columns}
              loading={loading}
              virtualized
              rowHeight={44}
              maxHeight={620}
              getRowId={(row) => row.address}
              initialSorting={[{ id: 'contributions', desc: true }]}
              onRowClick={(row) => navigate(paths.operator(row.address))}
              empty={
                <EmptyState
                  title='No operator matches'
                  description='Clear the address filter or pick another status.'
                />
              }
            />
          </div>
        </div>
      </Card>
    </Stack>
  )
}
