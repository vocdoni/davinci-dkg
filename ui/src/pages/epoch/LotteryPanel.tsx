import { useMemo } from 'react'
import {
  Address,
  BlockCell,
  Callout,
  DataTable,
  EmptyState,
  Hash,
  KeyValue,
  Panel,
  TxCell,
  type AnyColumnDef,
  type KeyValueItem,
} from '~kit'
import { Gauge } from '~kit/charts'
import type { EpochDetail, LotteryInfo } from '~indexer/selectors'
import { bigIntToHex } from '~lib/format'
import { paths } from '~routes/paths'
import { blockOrNull } from '~pages/epochs/cadence'

type Claim = LotteryInfo['claims'][number]

/**
 * The lottery in full: `keccak(seed ‖ operator) < τ` decides who may claim, so
 * τ as a share of the hash space *is* the admission rate. Anyone can replay it
 * from the seed and the addresses below.
 */
export function LotteryPanel({ detail }: { detail: EpochDetail }) {
  const { lottery, epoch, row } = detail
  const claims = useMemo(() => [...lottery.claims].sort((a, b) => a.slot - b.slot), [lottery.claims])

  const items: KeyValueItem[] = [
    {
      label: 'seed block',
      value: <BlockCell block={blockOrNull(lottery.seedBlock)} />,
      hint: 'blockhash(startBlock + SEED_DELAY_BLOCKS)',
    },
    {
      label: 'seed',
      value: lottery.seed ? <Hash value={lottery.seed} chars={10} /> : <span className='text-ash'>unresolved</span>,
      mono: true,
      hint:
        lottery.seedResolvedBlock != null
          ? `resolved at block ${lottery.seedResolvedBlock.toLocaleString()} by the first claim`
          : 'resolved on the first claimSlot call',
    },
    { label: 'τ', value: <Hash value={bigIntToHex(lottery.threshold)} chars={10} />, mono: true },
    {
      label: 'α',
      value: `${lottery.alpha} (${lottery.alphaBps} bps)`,
      mono: true,
      hint: 'oversubscription factor chosen by the epoch creator',
    },
    {
      label: 'N snapshotted',
      value: lottery.registrySnapshot != null ? Math.round(lottery.registrySnapshot).toLocaleString() : '—',
      mono: true,
      hint: 'registry activeCount() at createEpoch, recovered from τ',
    },
    {
      label: 'admissible',
      value: lottery.admissibleProbability != null ? `${(lottery.admissibleProbability * 100).toFixed(2)}%` : '—',
      mono: true,
      hint: 'min(1, α·n/N) — the chance a registered operator may claim',
    },
    {
      label: 'claims',
      value: `${row.claims} of ${row.committeeSize}`,
      mono: true,
      hint: 'first come, first served among the admissible',
    },
  ]

  const columns: AnyColumnDef<Claim>[] = [
    {
      id: 'slot',
      header: 'Slot',
      accessorKey: 'slot',
      meta: { numeric: true, width: '72px', headerTooltip: 'Committee slot, counted from 0' },
    },
    {
      id: 'index',
      header: 'Index',
      accessorFn: (claim) => claim.slot + 1,
      meta: { numeric: true, width: '72px', headerTooltip: 'Participant index the protocol uses: slot + 1' },
    },
    {
      id: 'operator',
      header: 'Operator',
      accessorKey: 'operator',
      meta: { width: '200px' },
      cell: ({ row: r }) => (
        <Address value={r.original.operator} copy={false} to={paths.operator(r.original.operator)} />
      ),
    },
    {
      id: 'block',
      header: 'Block',
      accessorKey: 'block',
      meta: { numeric: true, width: '120px' },
      cell: ({ row: r }) => <BlockCell block={blockOrNull(r.original.block)} />,
    },
    {
      id: 'tx',
      header: 'Transaction',
      meta: { width: '140px' },
      cell: ({ row: r }) => (r.original.tx ? <TxCell hash={r.original.tx} /> : <span className='text-ash'>—</span>),
    },
  ]

  return (
    <Panel
      label='Lottery'
      title='Committee selection'
      description='Trustless and replayable: no organizer can prefer an operator, and only operators registered before createEpoch may claim.'
    >
      {epoch.status === 'aborted' ? (
        <Callout tone='danger' title='Epoch aborted' className='mb-5'>
          The committee never filled: {row.claims} of {row.committeeSize} slots were claimed before the selection window
          closed
          {detail.windows.committeeSelectionDeadline != null
            ? ` at block ${detail.windows.committeeSelectionDeadline.toLocaleString()}`
            : ''}
          {epoch.abortedBlock != null ? `, and the abort was recorded at block ${epoch.abortedBlock.toLocaleString()}` : ''}
          . The next scheduled epoch opens automatically.
        </Callout>
      ) : null}

      <div className='grid gap-6 lg:grid-cols-[240px_1fr]'>
        <Gauge
          value={lottery.thresholdFraction}
          label='τ / 2²⁵⁶'
          caption={`α = ${lottery.alpha} · n = ${row.committeeSize}`}
          size={200}
        />
        <KeyValue items={items} />
      </div>

      <div className='mt-6'>
        <div className='label-caps mb-3 text-[10px] text-pewter'>Claims in slot order</div>
        <div className='overflow-x-auto scroll-slim'>
          <div className='min-w-[620px]'>
            <DataTable
              data={claims}
              columns={columns}
              maxHeight={420}
              getRowId={(claim) => String(claim.slot)}
              empty={
                <EmptyState
                  compact
                  title='No slot claimed yet'
                  description='The lottery seed resolves on the first claimSlot call.'
                />
              }
            />
          </div>
        </div>
      </div>
    </Panel>
  )
}
