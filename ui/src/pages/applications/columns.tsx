// Columns for the cross-epoch applications table.

import { Link } from 'react-router-dom'
import { Address, Badge, BlockCell, Hash, ProgressBar, type AnyColumnDef } from '~kit'
import type { ApplicationRow } from '~indexer/selectors'
import type { EpochId } from '~indexer/types'
import { shortHash } from '~lib/format'
import { paths } from '~routes/paths'
import { StopClick } from '../operators/cells'
import { WindowCell } from './cells'

export type EpochLabel = (id: EpochId) => number | null | undefined

export function applicationColumns(nonceOf: EpochLabel): AnyColumnDef<ApplicationRow>[] {
  return [
    {
      id: 'epoch',
      header: 'Epoch',
      accessorFn: (row) => nonceOf(row.epoch) ?? 0,
      meta: { width: '80px' },
      cell: ({ row }) => {
        const nonce = nonceOf(row.original.epoch)
        return (
          <StopClick>
            <Link
              to={paths.epoch(row.original.epoch)}
              title={row.original.epoch}
              className='font-mono text-[12px] text-silver transition-colors hover:text-emerald'
            >
              {nonce != null ? `#${nonce}` : shortHash(row.original.epoch, 4, 4)}
            </Link>
          </StopClick>
        )
      },
    },
    {
      id: 'aid',
      header: 'Application',
      accessorKey: 'aid',
      cell: ({ row }) => (
        <Hash
          value={row.original.aid}
          chars={8}
          copy={false}
          href={paths.application(row.original.epoch, row.original.aid)}
        />
      ),
    },
    {
      id: 'organizer',
      header: 'Organizer',
      accessorKey: 'creator',
      meta: { width: '140px', headerTooltip: 'Registered the application and holds sk_org' },
      cell: ({ row }) => (
        <StopClick>
          <Address
            value={row.original.creator}
            to={paths.operator(row.original.creator)}
            copy={false}
            explorer={false}
          />
        </StopClick>
      ),
    },
    {
      id: 'submitter',
      header: 'Submitter',
      accessorFn: (row) => row.authorizedSubmitter ?? '',
      meta: { width: '140px', headerTooltip: 'The only address allowed to submit ciphertexts for this application' },
      cell: ({ row }) =>
        row.original.authorizedSubmitter ? (
          <StopClick>
            <Address value={row.original.authorizedSubmitter} copy={false} explorer={false} />
          </StopClick>
        ) : (
          <span className='text-ash'>—</span>
        ),
    },
    {
      id: 'window',
      header: 'Window',
      enableSorting: false,
      meta: { width: '150px', headerTooltip: 'Blocks between which ciphertexts are accepted (0 = unbounded)' },
      cell: ({ row }) => <WindowCell row={row.original} />,
    },
    {
      id: 'cap',
      header: 'Cap',
      accessorFn: (row) => row.maxCiphertexts ?? 0,
      meta: { numeric: true, width: '64px', headerTooltip: 'Maximum ciphertext index; 0 means uncapped' },
      cell: ({ row }) =>
        row.original.maxCiphertexts ? row.original.maxCiphertexts : <span className='text-ash'>∞</span>,
    },
    {
      id: 'ciphertexts',
      header: 'Ciphertexts',
      accessorKey: 'ciphertexts',
      meta: { numeric: true, width: '92px' },
    },
    {
      id: 'decrypted',
      header: 'Decrypted',
      accessorFn: (row) => (row.ciphertexts === 0 ? -1 : row.decrypted / row.ciphertexts),
      meta: { width: '120px' },
      cell: ({ row }) =>
        row.original.ciphertexts === 0 ? (
          <span className='text-ash'>—</span>
        ) : (
          <ProgressBar
            value={row.original.decrypted}
            total={row.original.ciphertexts}
            label={false}
            size='sm'
            className='w-24'
          />
        ),
    },
    {
      id: 'shares',
      header: 'Shares',
      accessorFn: (row) => row.sharesPublished,
      meta: { width: '104px', headerTooltip: 'Ciphertexts whose organizer share is on chain' },
      cell: ({ row }) => {
        const { sharesPublished, ciphertexts } = row.original
        if (ciphertexts === 0) return <span className='text-ash'>—</span>
        const all = sharesPublished >= ciphertexts
        return (
          <Badge size='sm' tone={all ? 'ok' : sharesPublished === 0 ? 'warn' : 'neutral'}>
            {sharesPublished} / {ciphertexts}
          </Badge>
        )
      },
    },
    {
      id: 'created',
      header: 'Registered',
      accessorKey: 'createdBlock',
      meta: { numeric: true, width: '104px' },
      cell: ({ row }) => (
        <StopClick>
          <BlockCell block={row.original.createdBlock} />
        </StopClick>
      ),
    },
  ]
}
