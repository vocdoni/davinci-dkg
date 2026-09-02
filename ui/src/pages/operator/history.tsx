import { Link } from 'react-router-dom'
import { Badge, BlockCell, DataTable, EmptyState, TxCell, type AnyColumnDef } from '~kit'
import { formatCompact } from '~kit/charts'
import { paths } from '~routes/paths'
import type { OperatorHistoryRow } from './history-rows'

function Yes({ children, tone = 'ok' }: { children: string; tone?: 'ok' | 'neutral' }) {
  return (
    <Badge size='sm' tone={tone}>
      {children}
    </Badge>
  )
}

const NO = <span className='text-ash'>—</span>

const columns: AnyColumnDef<OperatorHistoryRow>[] = [
  {
    id: 'epoch',
    header: 'Epoch',
    accessorKey: 'nonce',
    meta: { width: '110px' },
    cell: ({ row }) => (
      <Link
        to={paths.epoch(row.original.epoch)}
        className='font-mono text-[12px] text-silver transition-colors hover:text-emerald'
      >
        #{row.original.nonce}
      </Link>
    ),
  },
  {
    id: 'claimed',
    header: 'Claimed',
    accessorFn: (row) => (row.claimed ? 1 : 0),
    meta: { width: '120px' },
    cell: ({ row }) => (row.original.claimed ? <Yes tone='neutral'>{`slot ${row.original.slot ?? '?'}`}</Yes> : NO),
  },
  {
    id: 'contributed',
    header: 'Contributed',
    accessorFn: (row) => (row.contributed ? 1 : 0),
    meta: { width: '110px' },
    cell: ({ row }) => (row.original.contributed ? <Yes>yes</Yes> : NO),
  },
  {
    id: 'contributionBlock',
    header: 'Block',
    accessorFn: (row) => row.contributionBlock ?? 0,
    meta: { numeric: true, width: '110px' },
    cell: ({ row }) => <BlockCell block={row.original.contributionBlock} />,
  },
  {
    id: 'contributionTx',
    header: 'Transaction',
    enableSorting: false,
    meta: { width: '130px' },
    cell: ({ row }) => (row.original.contributionTx ? <TxCell hash={row.original.contributionTx} /> : NO),
  },
  {
    id: 'gas',
    header: 'Gas',
    accessorFn: (row) => row.contributionGas ?? 0,
    meta: { numeric: true, width: '90px', headerTooltip: 'Gas used by the contribution transaction' },
    cell: ({ row }) =>
      row.original.contributionGas != null ? (
        formatCompact(row.original.contributionGas)
      ) : (
        <span className='text-ash'>—</span>
      ),
  },
  {
    id: 'finalized',
    header: 'Finalized',
    accessorFn: (row) => (row.finalized ? 1 : 0),
    meta: { width: '100px' },
    cell: ({ row }) => (row.original.finalized ? <Yes>yes</Yes> : NO),
  },
  { id: 'partials', header: 'Partials', accessorKey: 'partials', meta: { numeric: true, width: '90px' } },
  { id: 'combines', header: 'Combines', accessorKey: 'combines', meta: { numeric: true, width: '90px' } },
]

/** Per-epoch history: what this operator did in every epoch it touched. */
export function OperatorHistoryTable({ rows, loading }: { rows: OperatorHistoryRow[]; loading?: boolean }) {
  return (
    <DataTable
      data={rows}
      columns={columns}
      loading={loading}
      maxHeight={460}
      getRowId={(row) => row.epoch}
      initialSorting={[{ id: 'epoch', desc: true }]}
      empty={
        <EmptyState
          compact
          title='No epoch history'
          description='This operator has not claimed a slot or contributed in any indexed epoch.'
        />
      }
    />
  )
}
