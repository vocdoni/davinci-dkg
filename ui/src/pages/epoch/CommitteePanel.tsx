import { useMemo } from 'react'
import { Address, BlockCell, DataTable, EmptyState, Panel, ProgressBar, TxCell, type AnyColumnDef } from '~kit'
import type { CommitteeRow, EpochDetail } from '~indexer/selectors'
import { cn } from '~lib/cn'
import { paths } from '~routes/paths'
import { blockOrNull } from '~pages/epochs/cadence'

/** Above this the member table windows its rows. A 64-member committee does. */
const VIRTUALIZE_ABOVE = 50

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'

/**
 * Every committee member, in slot order. Slots are 0-based and the protocol's
 * participant index is `slot + 1` — both columns are shown, because every
 * on-chain call (`contributorIndex`, `participantIndex`, the Merkle leaf of a
 * partial's share proof) uses the 1-based one.
 */
export function CommitteePanel({ detail }: { detail: EpochDetail }) {
  const { committee, row, epoch } = detail
  const claimed = committee.filter((member) => member.claimBlock != null).length
  // The accepted-contribution count is the number of ContributionSubmitted
  // logs, not the number of rows the join marked green.
  const contributed = row.contributions

  const columns = useMemo<AnyColumnDef<CommitteeRow>[]>(
    () => [
      {
        id: 'slot',
        header: 'Slot',
        accessorKey: 'slot',
        meta: { numeric: true, width: '64px', headerTooltip: 'Committee slot, counted from 0' },
      },
      {
        id: 'index',
        header: 'Index',
        accessorKey: 'participantIndex',
        meta: { numeric: true, width: '64px', headerTooltip: 'Participant index used on chain: slot + 1' },
      },
      {
        id: 'operator',
        header: 'Operator',
        accessorKey: 'operator',
        meta: { width: '186px' },
        cell: ({ row: r }) =>
          r.original.operator === ZERO_ADDRESS ? (
            <span className='text-ash' title='this slot was never claimed'>
              unclaimed
            </span>
          ) : (
            <Address
              value={r.original.operator}
              copy={false}
              explorer={false}
              to={paths.operator(r.original.operator)}
            />
          ),
      },
      {
        id: 'claim',
        header: 'Claimed',
        accessorFn: (member) => member.claimBlock ?? 0,
        meta: { numeric: true, width: '116px' },
        cell: ({ row: r }) => <BlockCell block={blockOrNull(r.original.claimBlock)} />,
      },
      {
        id: 'contribution',
        header: 'Contributed',
        accessorFn: (member) => member.contributionBlock ?? 0,
        meta: { numeric: true, width: '124px' },
        cell: ({ row: r }) => {
          if (r.original.contributed) return <BlockCell block={blockOrNull(r.original.contributionBlock)} />
          if (r.original.claimBlock == null) return <span className='text-ash'>—</span>
          return (
            <span className='text-amber' title='no contribution in the key-assembly window'>
              missing
            </span>
          )
        },
      },
      {
        id: 'gas',
        header: 'Gas',
        accessorFn: (member) => member.contributionGas ?? 0,
        meta: { numeric: true, width: '100px', headerTooltip: 'Gas used by submitContribution (from the receipt)' },
        cell: ({ row: r }) =>
          r.original.contributionGas != null ? (
            r.original.contributionGas.toLocaleString()
          ) : (
            <span className='text-ash'>—</span>
          ),
      },
      {
        id: 'tx',
        header: 'Transaction',
        meta: { width: '128px' },
        cell: ({ row: r }) =>
          r.original.contributionTx ? (
            <TxCell hash={r.original.contributionTx} />
          ) : (
            <span className='text-ash'>—</span>
          ),
      },
      {
        id: 'partials',
        header: 'Partials',
        accessorKey: 'partials',
        meta: { numeric: true, width: '88px', headerTooltip: 'Partial decryptions this member has published' },
      },
    ],
    []
  )

  const reached = contributed >= row.minValidContributions
  const virtualized = committee.length > VIRTUALIZE_ABOVE

  return (
    <Panel
      label='Committee'
      title={`${committee.length} members`}
      description='Each member ran a Feldman VSS contribution with a Groth16 proof; a slot with no contribution is a member that did not show up.'
      bodyClassName='p-0'
      actions={
        <span className='font-mono text-[11px] text-ash'>
          t = {row.threshold} · m_min = {row.minValidContributions} · n = {row.committeeSize}
        </span>
      }
    >
      <div className='grid gap-6 border-b border-charcoal p-5 md:grid-cols-2'>
        <ProgressBar value={claimed} total={row.committeeSize} label='slots claimed' />
        <div className='min-w-0'>
          {/* The kit's built-in caption calls the marker "t"; here it is m_min,
              so the caption is written out rather than inherited. */}
          <div className='mb-1.5 flex items-baseline justify-between gap-2 text-[11px]'>
            <span className='text-pewter'>contributions</span>
            <span className={cn('font-mono tnum', reached ? 'text-emerald' : 'text-silver')}>
              {contributed}
              <span className='text-ash'> / {row.committeeSize}</span>
              <span className='text-ash'> · m_min={row.minValidContributions}</span>
            </span>
          </div>
          <ProgressBar
            value={contributed}
            total={row.committeeSize}
            threshold={row.minValidContributions || undefined}
            label={false}
            tone={reached ? 'accent' : 'warn'}
          />
        </div>
      </div>

      <div className='overflow-x-auto scroll-slim'>
        <div className='min-w-[1010px]'>
          <DataTable
            data={committee}
            columns={columns}
            virtualized={virtualized}
            rowHeight={40}
            maxHeight={virtualized ? 560 : undefined}
            getRowId={(member) => String(member.slot)}
            empty={
              <EmptyState
                compact
                title='No committee yet'
                description={
                  epoch.status === 'aborted'
                    ? 'The lottery never filled every slot.'
                    : 'Slots are still being claimed.'
                }
              />
            }
          />
        </div>
      </div>
    </Panel>
  )
}
