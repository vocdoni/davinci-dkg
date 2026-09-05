import { useMemo, useState } from 'react'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { useApplications, useEpochs, useIndexer } from '~data/hooks'
import type { AppModeName, EpochId } from '~indexer/types'
import {
  Callout,
  Card,
  DataTable,
  EmptyState,
  Input,
  SearchIcon,
  SectionHeader,
  Select,
  Stack,
  StatCell,
  StatRow,
} from '~kit'
import { formatCompact } from '~kit/charts'
import { paths } from '~routes/paths'
import { applicationColumns } from './applications/columns'
import { filterApplications, summarizeApplications } from './applications/filter'
import { StoreSearchHits } from './operators/store-search'

/**
 * Every application ever registered, across every epoch. The epoch filter is
 * in the URL so the epoch page can link straight into its own applications.
 */
export function ApplicationsPage() {
  const rows = useApplications()
  const epochs = useEpochs()
  const { scanning, status } = useIndexer()
  const navigate = useNavigate()

  const [params, setParams] = useSearchParams()
  const epoch = params.get('epoch') ?? 'all'
  const [query, setQuery] = useState('')
  const [mode, setMode] = useState<AppModeName | 'all'>('all')

  const nonceOf = useMemo(() => {
    const map = new Map<string, number>()
    for (const row of epochs) map.set(row.id.toLowerCase(), row.nonce)
    return (id: EpochId) => map.get(id.toLowerCase()) ?? null
  }, [epochs])

  const filtered = useMemo(() => filterApplications(rows, { query, epoch, mode }), [rows, query, epoch, mode])
  const totals = useMemo(() => summarizeApplications(filtered), [filtered])
  const columns = useMemo(() => applicationColumns(nonceOf), [nonceOf])

  const epochOptions = useMemo(
    () => [
      { value: 'all', label: 'All epochs' },
      ...epochs.map((row) => ({ value: row.id, label: `Epoch #${row.nonce} · ${row.phase}` })),
    ],
    [epochs]
  )

  const loading = scanning && rows.length === 0
  // Only when there is nothing to show: "no applications yet" and "the scan
  // failed" look identical otherwise.
  const failure = rows.length === 0 && !scanning ? (status.errors[status.errors.length - 1] ?? null) : null

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Applications'
        title='Applications'
        description='Every application registered against any epoch, with the state of its decryption pipeline. Each application claims one of its epoch’s pool keys: an automatic one encrypts under that key alone, an organizer-locked one under P_j + PK_org and opens nothing until its organizer reveals sk_org.'
      />

      {failure ? (
        <Callout tone='danger' title='The indexer could not read the chain'>
          {failure.scope}: {failure.message}
        </Callout>
      ) : null}

      <StatRow>
        <StatCell label='Applications' value={totals.applications} mono loading={loading} />
        <StatCell label='Ciphertexts' value={formatCompact(totals.ciphertexts)} mono loading={loading} />
        <StatCell
          label='Decrypted'
          value={formatCompact(totals.decrypted)}
          mono
          tone='accent'
          hint={
            totals.ciphertexts > 0
              ? `${Math.round((totals.decrypted / totals.ciphertexts) * 100)}% of all ciphertexts`
              : undefined
          }
          loading={loading}
        />
        <StatCell
          label='Secrets kept'
          value={formatCompact(totals.locked)}
          mono
          tone={totals.locked > 0 ? 'warn' : 'default'}
          hint='locked applications, sk_org not revealed yet'
          loading={loading}
        />
      </StatRow>

      <Card flush className='overflow-hidden'>
        <div className='flex flex-wrap items-end justify-between gap-4 border-b border-charcoal px-5 py-4'>
          <div className='min-w-0'>
            <div className='label-caps mb-1.5 text-emerald'>Registry</div>
            <h2 className='text-[15px] font-semibold text-ghost'>
              {filtered.length === rows.length
                ? `${rows.length} applications`
                : `${filtered.length} of ${rows.length} applications`}
            </h2>
          </div>
          <div className='flex flex-wrap items-end gap-3'>
            <Input
              size='sm'
              mono
              value={query}
              placeholder='application id or organizer'
              aria-label='Search applications'
              iconLeft={<SearchIcon size={13} />}
              onChange={(e) => setQuery(e.target.value)}
              wrapperClassName='w-72'
            />
            <Select
              size='sm'
              aria-label='Filter by mode'
              value={mode}
              onChange={(e) => setMode(e.target.value as AppModeName | 'all')}
              options={[
                { value: 'all', label: 'All modes' },
                { value: 'organizer-locked', label: 'organizer-locked' },
                { value: 'automatic', label: 'automatic' },
              ]}
              wrapperClassName='w-44'
            />
            <Select
              size='sm'
              aria-label='Filter by epoch'
              value={epoch}
              onChange={(e) => {
                const value = e.target.value
                const next = new URLSearchParams(params)
                if (value === 'all') next.delete('epoch')
                else next.set('epoch', value)
                setParams(next, { replace: true })
              }}
              options={epochOptions}
              wrapperClassName='w-56'
            />
          </div>
        </div>

        {query.trim() !== '' ? <StoreSearchHits query={query} skip={['application']} className='m-5 mb-0' /> : null}

        <div className='overflow-x-auto scroll-slim'>
          <div className='min-w-[1300px]'>
            <DataTable
              data={filtered}
              columns={columns}
              loading={loading}
              virtualized={filtered.length > 60}
              rowHeight={44}
              maxHeight={620}
              getRowId={(row) => row.key}
              initialSorting={[{ id: 'created', desc: true }]}
              onRowClick={(row) => navigate(paths.application(row.epoch, row.aid))}
              empty={
                <EmptyState
                  title='No application matches'
                  description='Clear the search or pick another epoch. Applications appear as soon as an organizer registers one against a live epoch.'
                />
              }
            />
          </div>
        </div>
      </Card>
    </Stack>
  )
}
