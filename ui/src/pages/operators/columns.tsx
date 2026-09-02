// Column definitions for the operators table.
//
// Every counter is a plain number so TanStack can sort it; the cells only
// decide how it looks. Widths are fixed (the table is virtualised, and a
// virtualised row is a flex row, not a `<table>`), so the panel scrolls
// horizontally inside itself rather than pushing the page sideways.

import { Address, BlockCell, type AnyColumnDef } from '~kit'
import { formatParticipation, type OperatorRow } from '~indexer/selectors'
import { paths } from '~routes/paths'
import { LastActiveCell, OperatorKeyCell, OperatorStatusBadge, StopClick } from './cells'

export interface OperatorColumnOptions {
  /** Seconds per block, for the "blocks ago" tooltip. */
  blockTimeSeconds: number
}

export function operatorColumns({ blockTimeSeconds }: OperatorColumnOptions): AnyColumnDef<OperatorRow>[] {
  return [
    {
      id: 'address',
      header: 'Operator',
      accessorKey: 'address',
      cell: ({ row }) => (
        <Address value={row.original.address} to={paths.operator(row.original.address)} copy={false} explorer={false} />
      ),
    },
    {
      id: 'status',
      header: 'Status',
      accessorFn: (row) => (row.status === 'active' ? (row.reapable ? 1 : 0) : 2),
      meta: { width: '84px' },
      cell: ({ row }) => <OperatorStatusBadge row={row.original} />,
    },
    {
      id: 'registered',
      header: 'Registered',
      accessorKey: 'registeredAtBlock',
      meta: { numeric: true, width: '96px', headerTooltip: 'Block the operator (re-)entered the active set' },
      cell: ({ row }) => (
        <StopClick>
          <BlockCell block={row.original.registeredAtBlock} />
        </StopClick>
      ),
    },
    {
      id: 'idle',
      header: 'Last active',
      accessorFn: (row) => row.idleBlocks ?? Number.MAX_SAFE_INTEGER,
      meta: { numeric: true, width: '88px', headerTooltip: 'Blocks since the last liveness-bearing transaction' },
      cell: ({ row }) => <LastActiveCell row={row.original} blockTimeSeconds={blockTimeSeconds} />,
    },
    {
      id: 'epochsServed',
      header: 'Epochs',
      accessorKey: 'epochsServed',
      meta: { numeric: true, width: '66px', headerTooltip: 'Epochs in which the operator claimed or contributed' },
    },
    { id: 'claims', header: 'Claims', accessorKey: 'claims', meta: { numeric: true, width: '62px' } },
    {
      id: 'contributions',
      header: 'Contribs',
      accessorKey: 'contributions',
      meta: { numeric: true, width: '78px', headerTooltip: 'Accepted contributions to a committee transcript' },
    },
    { id: 'partials', header: 'Partials', accessorKey: 'partials', meta: { numeric: true, width: '72px' } },
    {
      id: 'finalizations',
      header: 'Finals',
      accessorKey: 'finalizations',
      meta: { numeric: true, width: '64px', headerTooltip: 'Epochs this operator finalized (from the tx sender)' },
    },
    {
      id: 'combines',
      header: 'Combines',
      accessorKey: 'combines',
      meta: { numeric: true, width: '78px', headerTooltip: 'Ciphertexts this operator combined (from the tx sender)' },
    },
    {
      id: 'participation',
      header: 'Participation',
      accessorFn: (row) => row.participation ?? -1,
      meta: {
        numeric: true,
        width: '92px',
        headerTooltip: 'Contributions / claims — a claimed slot is a promise to contribute. "—" when never claimed',
      },
      cell: ({ row }) =>
        row.original.participation == null ? (
          <span className='text-ash'>—</span>
        ) : (
          <span className={row.original.participation >= 1 ? 'text-emerald' : undefined}>
            {formatParticipation(row.original.participation)}
          </span>
        ),
    },
    {
      id: 'key',
      header: 'Key',
      enableSorting: false,
      meta: { width: '72px', align: 'right' },
      cell: ({ row }) => <OperatorKeyCell pubKey={row.original.pubKey} />,
    },
  ]
}
