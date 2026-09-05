// Columns for the cross-epoch applications table. Sized so the table fits a
// 1440 px viewport without a horizontal scroll: every track is fixed except
// the decryption window, which takes the slack; the pool key and the two
// windows stack on two lines rather than widen.

import { Link } from 'react-router-dom'
import { Address, BlockCell, Hash, ProgressBar, type AnyColumnDef } from '~kit'
import type { ApplicationRow } from '~indexer/selectors'
import type { EpochId } from '~indexer/types'
import { shortHash } from '~lib/format'
import { paths } from '~routes/paths'
import { StopClick } from '../operators/cells'
import {
  CiphertextCountCell,
  DecryptionWindowCell,
  ModeBadge,
  PoolKeyCell,
  SecretCell,
  SubmissionCell,
  WindowCell,
} from './cells'
import { submissionPolicy, submissionPolicyLabel } from './policy'

export type EpochLabel = (id: EpochId) => number | null | undefined

export function applicationColumns(nonceOf: EpochLabel): AnyColumnDef<ApplicationRow>[] {
  return [
    {
      id: 'epoch',
      header: 'Epoch',
      accessorFn: (row) => nonceOf(row.epoch) ?? 0,
      meta: { width: '72px' },
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
      meta: { width: '136px' },
      cell: ({ row }) => (
        <Hash value={row.original.aid} copy={false} href={paths.application(row.original.epoch, row.original.aid)} />
      ),
    },
    {
      id: 'mode',
      header: 'Mode',
      accessorKey: 'mode',
      meta: {
        width: '128px',
        headerTooltip:
          'organizer-locked: PK_aid = P_j + PK_org, partials and combines refused until the organizer reveals sk_org · automatic: PK_aid = P_j, no organizer key',
      },
      cell: ({ row }) => <ModeBadge mode={row.original.mode} />,
    },
    {
      id: 'organizer',
      header: 'Organizer',
      accessorKey: 'creator',
      meta: { width: '112px', headerTooltip: 'Registered the application and, when organizer-locked, holds sk_org until the reveal' },
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
      id: 'submission',
      header: 'Submission',
      accessorFn: (row) => submissionPolicyLabel(submissionPolicy(row)),
      meta: {
        width: '124px',
        headerTooltip: 'Who may call submitCiphertext: the registrant only, an allow-list, or anyone (open)',
      },
      cell: ({ row }) => (
        <StopClick>
          <SubmissionCell row={row.original} />
        </StopClick>
      ),
    },
    {
      id: 'window',
      header: 'Window',
      enableSorting: false,
      meta: { width: '112px', headerTooltip: 'Blocks between which ciphertexts are accepted; any block when unbounded' },
      cell: ({ row }) => <WindowCell row={row.original} />,
    },
    {
      id: 'pool',
      header: 'Pool key',
      accessorFn: (row) => row.poolIndex ?? -1,
      meta: { width: '96px', headerTooltip: 'Index of the epoch pool key this application claimed at registration, and P_j.x' },
      cell: ({ row }) => <PoolKeyCell row={row.original} />,
    },
    {
      id: 'decryption',
      header: 'Decryption window',
      accessorFn: (row) => row.decryptNotAfter ?? -1,
      meta: {
        headerTooltip:
          'policy.decryptNotBefore → decryptNotAfter — partials and combines revert outside it (and, when organizer-locked, before the reveal)',
      },
      cell: ({ row }) => (
        <DecryptionWindowCell notBefore={row.original.decryptNotBefore} notAfter={row.original.decryptNotAfter} compact />
      ),
    },
    {
      id: 'ciphertexts',
      header: 'Ciphertexts',
      accessorKey: 'ciphertexts',
      meta: { numeric: true, width: '96px', headerTooltip: 'Submitted / policy.maxCiphertexts (∞ = uncapped)' },
      cell: ({ row }) => <CiphertextCountCell row={row.original} />,
    },
    {
      id: 'decrypted',
      header: 'Decrypted',
      accessorFn: (row) => (row.ciphertexts === 0 ? -1 : row.decrypted / row.ciphertexts),
      meta: { width: '92px', headerTooltip: 'Ciphertexts with a combined plaintext on chain' },
      cell: ({ row }) =>
        row.original.ciphertexts === 0 ? (
          <span className='text-ash'>—</span>
        ) : (
          <ProgressBar
            value={row.original.decrypted}
            total={row.original.ciphertexts}
            label={false}
            size='sm'
            className='w-14'
          />
        ),
    },
    {
      id: 'secret',
      header: 'Secret',
      accessorFn: (row) => (row.mode === 'automatic' ? 2 : row.unlocked ? 1 : 0),
      meta: { width: '92px', headerTooltip: 'Organizer secret: revealed, kept, or none for an automatic application' },
      cell: ({ row }) => <SecretCell row={row.original} />,
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
