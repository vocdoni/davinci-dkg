import { useMemo } from 'react'
import { useNavigate } from 'react-router-dom'
import { Address, BlockCell, DataTable, EmptyState, Hash, Panel, type AnyColumnDef } from '~kit'
import type { ApplicationRow } from '~indexer/selectors'
import { paths } from '~routes/paths'
import { blockOrNull } from '~pages/epochs/cadence'

/**
 * Applications registered against this epoch. `PK_aid = PK_ep + PK_org`, so
 * each row is an independent encryption context that needs both the committee
 * and its organizer to decrypt.
 */
export function ApplicationsPanel({ applications }: { applications: ApplicationRow[] }) {
  const navigate = useNavigate()

  const columns = useMemo<AnyColumnDef<ApplicationRow>[]>(
    () => [
      {
        id: 'aid',
        header: 'Application id',
        accessorKey: 'aid',
        meta: { width: '188px', headerTooltip: 'bytes32 aid, bound into every decryption proof' },
        cell: ({ row }) => <Hash value={row.original.aid} chars={8} copy={false} />,
      },
      {
        id: 'organizer',
        header: 'Organizer',
        accessorKey: 'creator',
        meta: { width: '150px', headerTooltip: 'Registered PK_org with a Schnorr proof of possession' },
        cell: ({ row }) => <Address value={row.original.creator} copy={false} explorer={false} />,
      },
      {
        id: 'submitter',
        header: 'Submitter',
        accessorFn: (row) => row.authorizedSubmitter ?? '',
        meta: { width: '150px', headerTooltip: 'The only address allowed to call submitCiphertext' },
        cell: ({ row }) =>
          row.original.authorizedSubmitter ? (
            <Address value={row.original.authorizedSubmitter} copy={false} explorer={false} />
          ) : (
            <span className='text-ash'>—</span>
          ),
      },
      {
        id: 'window',
        header: 'Window',
        accessorFn: (row) => row.notBeforeBlock ?? 0,
        meta: { numeric: true, width: '164px', headerTooltip: 'Blocks in which ciphertexts may be submitted' },
        cell: ({ row }) => {
          const { notBeforeBlock: from, notAfterBlock: to } = row.original
          if (from == null || to == null) return <span className='text-ash'>—</span>
          // Both zero is how the contract spells "no window at all".
          if (from === 0 && to === 0) return <span className='text-ash'>any block</span>
          return (
            <span>
              {from.toLocaleString()}
              <span className='text-ash'> → </span>
              {to === 0 ? '∞' : to.toLocaleString()}
            </span>
          )
        },
      },
      {
        id: 'cap',
        header: 'Cap',
        accessorFn: (row) => row.maxCiphertexts ?? 0,
        meta: { numeric: true, width: '76px', headerTooltip: 'maxCiphertexts allowed under this aid' },
        cell: ({ row }) => row.original.maxCiphertexts ?? <span className='text-ash'>—</span>,
      },
      {
        id: 'ciphertexts',
        header: 'Ciphertexts',
        accessorKey: 'ciphertexts',
        meta: { numeric: true, width: '106px' },
      },
      {
        id: 'decrypted',
        header: 'Decrypted',
        accessorKey: 'decrypted',
        meta: { numeric: true, width: '106px' },
        cell: ({ row }) => (
          <span className={row.original.decrypted === row.original.ciphertexts ? 'text-emerald' : undefined}>
            {row.original.decrypted}
            <span className='text-ash'> / {row.original.ciphertexts}</span>
          </span>
        ),
      },
      {
        id: 'shares',
        header: 'Shares',
        accessorKey: 'sharesPublished',
        meta: { numeric: true, width: '92px', headerTooltip: 'Organizer shares published of ciphertexts submitted' },
        cell: ({ row }) => (
          <span className={row.original.sharesPublished < row.original.ciphertexts ? 'text-amber' : undefined}>
            {row.original.sharesPublished}
          </span>
        ),
      },
      {
        id: 'created',
        header: 'Registered',
        accessorKey: 'createdBlock',
        meta: { numeric: true, width: '116px' },
        cell: ({ row }) => <BlockCell block={blockOrNull(row.original.createdBlock)} />,
      },
    ],
    []
  )

  return (
    <Panel
      label='Applications'
      title={`${applications.length} registered`}
      description='One organizer key per application. Losing sk_org makes its ciphertexts permanently undecryptable.'
      bodyClassName='p-0'
    >
      <div className='overflow-x-auto scroll-slim'>
        <div className='min-w-[1140px]'>
          <DataTable
            data={applications}
            columns={columns}
            maxHeight={420}
            getRowId={(row) => row.key}
            onRowClick={(row) => navigate(paths.application(row.epoch, row.aid))}
            empty={
              <EmptyState
                compact
                title='No application yet'
                description='An organizer registers one with registerApplication while the epoch is live.'
              />
            }
          />
        </div>
      </div>
    </Panel>
  )
}
