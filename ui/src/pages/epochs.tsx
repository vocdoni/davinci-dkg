import { useMemo, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import {
  Address,
  DataTable,
  EmptyState,
  Hash,
  Input,
  Panel,
  ProgressBar,
  SearchIcon,
  SectionHeader,
  Stack,
  StatCell,
  StatRow,
  type AnyColumnDef,
} from '~kit'
import { useEpochs, useIndexer, useNetworkStats, useTxMeta } from '~data/hooks'
import type { EpochRow } from '~indexer/selectors'
import { paths } from '~routes/paths'
import { PhaseBadge } from './epochs/PhaseBadge'
import { PhaseFilter, type PhaseFilterValue } from './epochs/PhaseFilter'
import { elapsedSince } from './epochs/cadence'

/** Above this the table windows its rows instead of rendering all of them. */
const VIRTUALIZE_ABOVE = 50

export function EpochsPage() {
  const navigate = useNavigate()
  const indexer = useIndexer()
  const stats = useNetworkStats()
  const [phase, setPhase] = useState<PhaseFilterValue>('all')
  const [query, setQuery] = useState('')

  const all = useEpochs()
  const rows = useEpochs({ phase, query: query.trim() })
  const loading = (indexer.scanning || indexer.status.phase === 'loading') && all.length === 0

  const counts = useMemo(() => {
    const out: Record<string, number> = {}
    for (const row of all) out[row.phase] = (out[row.phase] ?? 0) + 1
    return out
  }, [all])

  // The finalizer is the transaction sender, so ask the indexer to resolve the
  // receipts of the rows actually on screen before the rest of its queue.
  const finalizationTxs = useMemo(() => rows.slice(0, 60).map((row) => row.finalizationTx), [rows])
  useTxMeta(finalizationTxs)

  const columns = useMemo<AnyColumnDef<EpochRow>[]>(
    () => [
      {
        id: 'id',
        header: 'Epoch',
        accessorKey: 'id',
        meta: { width: '156px' },
        cell: ({ row }) => <Hash value={row.original.id} chars={6} copy={false} />,
      },
      {
        id: 'nonce',
        header: 'Nonce',
        accessorKey: 'nonce',
        meta: { numeric: true, width: '68px', headerTooltip: 'Sequence number inside the bytes12 epoch id' },
        cell: ({ row }) => `#${row.original.nonce}`,
      },
      {
        id: 'phase',
        header: 'Phase',
        accessorKey: 'phase',
        meta: { width: '104px' },
        cell: ({ row }) => <PhaseBadge phase={row.original.phase} size='sm' />,
      },
      {
        id: 'policy',
        header: 't of n',
        accessorFn: (row) => row.threshold,
        meta: { numeric: true, width: '76px', headerTooltip: 'Threshold t of committee size n' },
        cell: ({ row }) => `${row.original.threshold} / ${row.original.committeeSize}`,
      },
      {
        id: 'claims',
        header: 'Claims',
        accessorFn: (row) => row.claimProgress,
        meta: { width: '128px', headerTooltip: 'Lottery slots claimed of n' },
        cell: ({ row }) => (
          <ProgressBar
            value={row.original.claims}
            total={row.original.committeeSize}
            label={false}
            size='sm'
            className='w-28'
          />
        ),
      },
      {
        id: 'contributions',
        header: 'Contributions',
        accessorFn: (row) => row.contributionProgress,
        meta: { width: '146px', headerTooltip: 'Accepted contributions of n, with the m_min marker' },
        cell: ({ row }) => (
          <ProgressBar
            value={row.original.contributions}
            total={row.original.committeeSize}
            threshold={row.original.minValidContributions || undefined}
            label={false}
            size='sm'
            className='w-32'
          />
        ),
      },
      {
        id: 'ciphertexts',
        header: 'Ciphertexts',
        accessorFn: (row) => row.ciphertexts,
        meta: { numeric: true, width: '104px', headerTooltip: 'Decrypted of submitted' },
        cell: ({ row }) => (
          <span>
            {row.original.decrypted}
            <span className='text-ash'> / {row.original.ciphertexts}</span>
          </span>
        ),
      },
      {
        id: 'live',
        header: 'Live since',
        accessorFn: (row) => row.liveSinceBlock ?? 0,
        meta: { numeric: true, width: '112px', headerTooltip: 'Blocks since finalizeEpoch landed' },
        cell: ({ row }) => {
          const since = elapsedSince(row.original.liveSinceBlock, stats.headBlock, stats.blockTimeSeconds)
          if (!since) return <span className='text-ash'>—</span>
          return (
            <span title={`block ${row.original.liveSinceBlock?.toLocaleString()}`}>
              {since.text}
            </span>
          )
        },
      },
      {
        id: 'creator',
        header: 'Creator',
        accessorKey: 'creator',
        meta: { width: '132px', headerTooltip: 'Whoever won the permissionless createEpoch race' },
        cell: ({ row }) => <Address value={row.original.creator} copy={false} explorer={false} />,
      },
      {
        id: 'finalizer',
        header: 'Finalizer',
        accessorFn: (row) => row.finalizer ?? '',
        meta: { width: '132px', headerTooltip: 'EpochLive names nobody — resolved from the transaction sender' },
        cell: ({ row }) =>
          row.original.finalizer ? (
            <Address value={row.original.finalizer} copy={false} explorer={false} />
          ) : (
            <span className='text-ash' title={row.original.finalizationTx ? 'awaiting the receipt' : 'not finalized'}>
              —
            </span>
          ),
      },
    ],
    [stats.headBlock, stats.blockTimeSeconds]
  )

  const virtualized = rows.length > VIRTUALIZE_ABOVE

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Epochs'
        title='Epochs'
        description='Every epoch the manager has created, newest first. One epoch is one DKG run: a lottery, a committee, a pool of keys and the decryptions it served.'
      />

      <StatRow>
        <StatCell label='Created' value={stats.epochs.toLocaleString()} mono hint='all time' />
        <StatCell
          label='Live'
          value={stats.epochsLive}
          mono
          tone={stats.epochsLive > 0 ? 'accent' : 'default'}
          hint='keys usable right now'
        />
        <StatCell
          label='Aborted'
          value={stats.epochsAborted}
          mono
          tone={stats.epochsAborted > 0 ? 'warn' : 'default'}
          hint='committee never filled'
        />
        <StatCell
          label='Committee in force'
          value={stats.committeeSizeInForce != null ? `${stats.thresholdInForce ?? '?'} of ${stats.committeeSizeInForce}` : '—'}
          mono
          hint='threshold t of n'
        />
      </StatRow>

      <Panel
        label='Table'
        title={`${rows.length.toLocaleString()} epoch${rows.length === 1 ? '' : 's'}`}
        description={virtualized ? 'Windowed rows — scrolling and filtering stay instant at any history length.' : undefined}
        bodyClassName='p-0'
        actions={
          <Input
            value={query}
            onChange={(event) => setQuery(event.target.value)}
            placeholder='id, nonce or address'
            aria-label='Search epochs'
            size='sm'
            mono
            iconLeft={<SearchIcon size={13} />}
            wrapperClassName='w-44 sm:w-64'
          />
        }
      >
        <div className='border-b border-charcoal px-4 py-3'>
          <PhaseFilter value={phase} onChange={setPhase} counts={counts} total={all.length} />
        </div>
        <div className='overflow-x-auto scroll-slim'>
          <div className='min-w-[1160px]'>
            <DataTable
              data={rows}
              columns={columns}
              loading={loading}
              loadingRows={10}
              virtualized={virtualized}
              rowHeight={44}
              maxHeight={virtualized ? 620 : undefined}
              getRowId={(row) => row.id}
              onRowClick={(row) => navigate(paths.epoch(row.id))}
              empty={
                <EmptyState
                  title='No epoch matches this view'
                  description={
                    all.length === 0
                      ? 'The manager has not created an epoch in the scanned range yet.'
                      : 'Clear the search box or pick another phase.'
                  }
                />
              }
            />
          </div>
        </div>
      </Panel>
    </Stack>
  )
}
