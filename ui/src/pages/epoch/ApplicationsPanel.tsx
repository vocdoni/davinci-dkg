import { useMemo } from 'react'
import { useNavigate } from 'react-router-dom'
import { Address, BlockCell, DataTable, EmptyState, Hash, Panel, type AnyColumnDef } from '~kit'
import type { ApplicationRow } from '~indexer/selectors'
import { paths } from '~routes/paths'
import { blockOrNull } from '~pages/epochs/cadence'
import {
  CiphertextCountCell,
  DecryptionWindowCell,
  ModeBadge,
  PoolKeyCell,
  SecretCell,
  SubmissionCell,
  WindowCell,
} from '~pages/applications/cells'
import { submissionPolicy, submissionPolicyLabel } from '~pages/applications/policy'

/**
 * Applications registered against this epoch. Each claimed one pool key, so
 * each row is an independent encryption context: an organizer-locked one
 * (`P_j + PK_org`) gets no partial and no combine until its organizer reveals
 * — the contract refuses them — an automatic one (`P_j`) only needs the
 * committee. Sized like the cross-epoch table: fixed tracks, the decryption
 * window takes the slack, so it fits a 1440 px viewport.
 */
export function ApplicationsPanel({ applications }: { applications: ApplicationRow[] }) {
  const navigate = useNavigate()

  const columns = useMemo<AnyColumnDef<ApplicationRow>[]>(
    () => [
      {
        id: 'aid',
        header: 'Application id',
        accessorKey: 'aid',
        meta: { width: '144px', headerTooltip: 'bytes32 aid, bound into every decryption proof' },
        cell: ({ row }) => <Hash value={row.original.aid} chars={8} copy={false} />,
      },
      {
        id: 'mode',
        header: 'Mode',
        accessorKey: 'mode',
        meta: {
          width: '128px',
          headerTooltip:
            'organizer-locked: P_j + PK_org, partials and combines refused until the reveal · automatic: P_j, no organizer key',
        },
        cell: ({ row }) => <ModeBadge mode={row.original.mode} />,
      },
      {
        id: 'pool',
        header: 'Pool key',
        accessorFn: (row) => row.poolIndex ?? -1,
        meta: { width: '96px', headerTooltip: 'Which of the epoch’s pool keys the application claimed, and P_j.x' },
        cell: ({ row }) => <PoolKeyCell row={row.original} />,
      },
      {
        id: 'organizer',
        header: 'Organizer',
        accessorKey: 'creator',
        meta: {
          width: '112px',
          headerTooltip: 'Registered the application — with PK_org and a Schnorr proof of possession when organizer-locked',
        },
        cell: ({ row }) => <Address value={row.original.creator} copy={false} explorer={false} />,
      },
      {
        id: 'submission',
        header: 'Submission',
        accessorFn: (row) => submissionPolicyLabel(submissionPolicy(row)),
        meta: {
          width: '124px',
          headerTooltip: 'Who may call submitCiphertext: the registrant only, an allow-list, or anyone',
        },
        cell: ({ row }) => <SubmissionCell row={row.original} />,
      },
      {
        id: 'window',
        header: 'Window',
        accessorFn: (row) => row.notBeforeBlock ?? 0,
        meta: { width: '112px', headerTooltip: 'Blocks in which ciphertexts may be submitted; any block when unbounded' },
        cell: ({ row }) => <WindowCell row={row.original} />,
      },
      {
        id: 'decryption',
        header: 'Decryption window',
        accessorFn: (row) => row.decryptNotAfter ?? -1,
        meta: {
          headerTooltip:
            'decryptNotBefore → decryptNotAfter; partials and combines revert outside it (and, when organizer-locked, before the reveal)',
        },
        cell: ({ row }) => (
          <DecryptionWindowCell
            notBefore={row.original.decryptNotBefore}
            notAfter={row.original.decryptNotAfter}
            compact
          />
        ),
      },
      {
        id: 'ciphertexts',
        header: 'Ciphertexts',
        accessorKey: 'ciphertexts',
        meta: { numeric: true, width: '96px', headerTooltip: 'Submitted / maxCiphertexts allowed under this aid (∞ = uncapped)' },
        cell: ({ row }) => <CiphertextCountCell row={row.original} />,
      },
      {
        id: 'decrypted',
        header: 'Decrypted',
        accessorKey: 'decrypted',
        meta: { numeric: true, width: '92px', headerTooltip: 'Combined / submitted' },
        cell: ({ row }) => (
          <span className={row.original.decrypted === row.original.ciphertexts ? 'text-emerald' : undefined}>
            {row.original.decrypted}
            <span className='text-ash'> / {row.original.ciphertexts}</span>
          </span>
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
        cell: ({ row }) => <BlockCell block={blockOrNull(row.original.createdBlock)} />,
      },
    ],
    []
  )

  return (
    <Panel
      label='Applications'
      title={`${applications.length} registered`}
      description='One pool key per application. An organizer-locked application also carries PK_org: the contract refuses its partials and combines until sk_org is revealed, and losing sk_org before the reveal makes its ciphertexts permanently undecryptable. An automatic application has no organizer key at all.'
      bodyClassName='p-0'
    >
      <div className='overflow-x-auto scroll-slim'>
        <div className='min-w-[1240px]'>
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
