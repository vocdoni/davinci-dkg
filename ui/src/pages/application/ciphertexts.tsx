// The per-ciphertext pipeline table: submitted → partials → organizer share →
// combined, with the transaction behind every step.

import {
  Address,
  Badge,
  BlockCell,
  CopyButton,
  DataTable,
  EmptyState,
  ProgressBar,
  Tooltip,
  TxCell,
  type AnyColumnDef,
} from '~kit'
import { formatCompact } from '~kit/charts'
import type { CiphertextRow, CiphertextState } from '~indexer/selectors'
import { paths } from '~routes/paths'

const STATE_TONE: Record<CiphertextState, 'ok' | 'warn' | 'neutral' | 'accent'> = {
  submitted: 'neutral',
  partials: 'neutral',
  'threshold-met': 'accent',
  'awaiting-share': 'warn',
  ready: 'accent',
  combined: 'ok',
}

const STATE_HELP: Record<CiphertextState, string> = {
  submitted: 'On chain, no partial decryption yet',
  partials: 'Some partials in, still below the threshold t',
  'threshold-met': 'Threshold reached',
  'awaiting-share': 't partials are in; the organizer share is the only thing missing',
  ready: 't partials and the organizer share are on chain — anyone may combine',
  combined: 'Decrypted; the plaintext is on chain',
}

function PartialsCell({ row }: { row: CiphertextRow }) {
  const indices = row.partials.map((partial) => partial.participantIndex).sort((a, b) => a - b)
  const shown = indices.slice(0, 48)
  return (
    <Tooltip
      content={
        indices.length === 0 ? (
          'No committee member has answered this ciphertext yet'
        ) : (
          <div>
            <div className='text-ghost'>
              {indices.length} of {row.committeeSize} members · t = {row.threshold}
            </div>
            <div className='mt-1 font-mono text-[10px] break-words text-ash'>
              {shown.map((i) => `#${i}`).join(' ')}
              {indices.length > shown.length ? ` +${indices.length - shown.length}` : ''}
            </div>
          </div>
        )
      }
    >
      <span className='block w-full'>
        <ProgressBar
          value={row.partialCount}
          total={row.committeeSize || row.threshold || 1}
          threshold={row.threshold || undefined}
          label={false}
          size='sm'
          className='w-28'
        />
      </span>
    </Tooltip>
  )
}

function PlaintextCell({ row }: { row: CiphertextRow }) {
  if (row.combined.plaintext == null) return <span className='text-ash'>—</span>
  const text = row.combined.plaintext.toString()
  const short = text.length > 14 ? `${text.slice(0, 12)}…` : text
  return (
    <span className='inline-flex min-w-0 items-center gap-1'>
      <Tooltip content={text}>
        <span className='truncate font-mono text-[12px] tnum text-ghost'>{short}</span>
      </Tooltip>
      <CopyButton value={text} label='Copy plaintext' />
    </span>
  )
}

const columns: AnyColumnDef<CiphertextRow>[] = [
  {
    id: 'index',
    header: '#',
    accessorKey: 'index',
    meta: { numeric: true, width: '56px', headerTooltip: 'Ciphertext index, assigned on chain (1-based)' },
  },
  {
    id: 'state',
    header: 'State',
    accessorKey: 'state',
    meta: { width: '128px' },
    cell: ({ row }) => (
      <Badge size='sm' tone={STATE_TONE[row.original.state]} title={STATE_HELP[row.original.state]}>
        {row.original.state}
      </Badge>
    ),
  },
  {
    id: 'submitted',
    header: 'Submitted',
    accessorKey: 'block',
    meta: { width: '150px' },
    cell: ({ row }) => (
      <span className='flex items-center gap-2'>
        <BlockCell block={row.original.block} />
        {row.original.tx ? <TxCell hash={row.original.tx} chars={4} /> : null}
      </span>
    ),
  },
  {
    id: 'submitter',
    header: 'Submitter',
    accessorKey: 'submitter',
    meta: { width: '132px' },
    cell: ({ row }) => (
      <Address
        value={row.original.submitter}
        to={paths.operator(row.original.submitter)}
        copy={false}
        explorer={false}
      />
    ),
  },
  {
    id: 'partials',
    header: 'Partials',
    accessorKey: 'partialCount',
    meta: {
      width: '150px',
      headerTooltip: 'Committee members that published a partial; the tick marks the threshold t',
    },
    cell: ({ row }) => (
      <span className='flex items-center gap-2'>
        <PartialsCell row={row.original} />
        <span className='font-mono text-[11px] tnum text-ash'>
          {row.original.partialCount}/{row.original.threshold || '?'}
        </span>
      </span>
    ),
  },
  {
    id: 'share',
    header: 'Organizer share',
    accessorFn: (row) => (row.share.present ? 1 : 0),
    meta: { width: '148px', headerTooltip: 'Δ = sk_org·C1; re-submission overwrites until the ciphertext is combined' },
    cell: ({ row }) => {
      const { share } = row.original
      if (!share.present) {
        return (
          <Badge size='sm' tone='warn'>
            withheld
          </Badge>
        )
      }
      return (
        <span className='flex items-center gap-2'>
          <Badge size='sm' tone='ok'>
            released
          </Badge>
          {share.overwrites > 0 ? (
            <span className='font-mono text-[11px] tnum text-ash' title={`${share.overwrites} overwrite(s)`}>
              ×{share.overwrites + 1}
            </span>
          ) : null}
          {share.block != null ? <BlockCell block={share.block} /> : null}
        </span>
      )
    },
  },
  {
    id: 'combined',
    header: 'Combined',
    accessorFn: (row) => row.combined.block ?? 0,
    meta: { width: '150px' },
    cell: ({ row }) =>
      row.original.combined.done ? (
        <span className='flex items-center gap-2'>
          <BlockCell block={row.original.combined.block} />
          {row.original.combined.tx ? <TxCell hash={row.original.combined.tx} chars={4} /> : null}
        </span>
      ) : (
        <span className='text-ash'>—</span>
      ),
  },
  {
    id: 'plaintext',
    header: 'Plaintext',
    enableSorting: false,
    cell: ({ row }) => <PlaintextCell row={row.original} />,
  },
  {
    id: 'combiner',
    header: 'Combiner',
    enableSorting: false,
    meta: { width: '140px', headerTooltip: 'From the transaction sender: DecryptionCombined names nobody' },
    cell: ({ row }) => {
      const by = row.original.combined.by
      if (!by) return <span className='text-ash'>—</span>
      return (
        <span className='flex items-center gap-2'>
          <Address value={by} to={paths.operator(by)} copy={false} explorer={false} />
          {row.original.combined.gasUsed != null ? (
            <span className='font-mono text-[11px] tnum text-ash'>{formatCompact(row.original.combined.gasUsed)}</span>
          ) : null}
        </span>
      )
    },
  },
]

export function CiphertextTable({ rows, loading }: { rows: CiphertextRow[]; loading?: boolean }) {
  return (
    <div className='overflow-x-auto scroll-slim'>
      <div className='min-w-[1200px]'>
        <DataTable
          data={rows}
          columns={columns}
          loading={loading}
          maxHeight={560}
          getRowId={(row) => row.key}
          initialSorting={[{ id: 'index', desc: false }]}
          empty={
            <EmptyState
              title='No ciphertexts yet'
              description='Nothing has been submitted under this application. Its authorized submitter posts them with submitCiphertext.'
            />
          }
        />
      </div>
    </div>
  )
}
