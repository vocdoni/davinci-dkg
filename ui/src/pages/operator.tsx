import { useMemo } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useIndexer, useOperator, useStore, useTxMeta } from '~data/hooks'
import { gasOf, formatParticipation } from '~indexer/selectors'
import { epochKey, type EpochId, type Hex } from '~indexer/types'
import {
  Address,
  BlockCell,
  Callout,
  Hash,
  KeyValue,
  Panel,
  SectionHeader,
  SkeletonText,
  Stack,
  StatCell,
  StatRow,
  TxCell,
  buttonClasses,
  type KeyValueItem,
} from '~kit'
import { Sparkline, CHART_COLORS } from '~kit/charts'
import { bigIntToHex, blocksToDuration } from '~lib/format'
import { paths } from '~routes/paths'
import { OperatorStatusBadge } from './operators/cells'
import { operatorHistoryRows, participationTrend, partialsTrend } from './operator/history-rows'
import { OperatorHistoryTable } from './operator/history'
import { OperatorEventLog } from './operator/events'

/**
 * One operator, end to end: who it is, what it did in every epoch, and the
 * transaction behind each of those facts.
 */
export function OperatorPage() {
  const { address = '' } = useParams()
  const detail = useOperator(address)
  const store = useStore()
  const { scanning } = useIndexer()

  const historyRows = useMemo(
    () => (detail ? operatorHistoryRows(detail.history, detail.events, (tx: Hex | null) => gasOf(store, tx)) : []),
    [detail, store]
  )
  // Ask the indexer for the receipts behind the contributions the table shows,
  // so the gas column fills in without a page-level log scan.
  useTxMeta(historyRows.map((row) => row.contributionTx))

  const nonceOf = useMemo(() => (id: EpochId) => store.epochs[epochKey(id)]?.nonce ?? null, [store])
  const participation = useMemo(() => participationTrend(historyRows), [historyRows])
  const partials = useMemo(() => partialsTrend(historyRows), [historyRows])

  if (!detail) {
    if (scanning) return <SkeletonText lines={8} className='max-w-2xl' />
    return (
      <Stack>
        <SectionHeader
          size='page'
          label='Operator'
          title='Unknown address'
          description='Nothing in the indexed range registers, claims or contributes under this address.'
        />
        <Callout tone='warn' title='No such operator'>
          <p className='font-mono text-[12px] break-all'>{address}</p>
          <p className='mt-2'>
            The registry may have been deployed after the indexed start block, or the address may simply never have been
            registered.{' '}
            <Link to={paths.operators()} className='text-emerald hover:underline'>
              Browse the registry
            </Link>
            .
          </p>
        </Callout>
      </Stack>
    )
  }

  const { row, operator } = detail
  const registration = detail.events.find((event) => event.name === 'NodeRegistered')
  const identity: KeyValueItem[] = [
    { label: 'Address', value: <Address value={row.address} full copy /> },
    { label: 'Status', value: <OperatorStatusBadge row={row} size='md' /> },
    {
      label: 'Registered',
      value: <BlockCell block={row.registeredAtBlock} />,
      hint: registration?.tx ? <TxCell hash={registration.tx} /> : undefined,
    },
    {
      label: 'Last active',
      value: <BlockCell block={row.lastActiveBlock > 0 ? row.lastActiveBlock : null} />,
      hint:
        row.idleBlocks != null
          ? `${row.idleBlocks} blocks ago (${blocksToDuration(row.idleBlocks, store.chain.blockTimeSeconds)})`
          : 'never marked active',
    },
    {
      label: 'Key x',
      value: operator.pubKey ? <Hash value={bigIntToHex(operator.pubKey.x)} chars={12} /> : '—',
    },
    {
      label: 'Key y',
      value: operator.pubKey ? <Hash value={bigIntToHex(operator.pubKey.y)} chars={12} /> : '—',
    },
    { label: 'Key updates', value: operator.keyUpdates, mono: true },
    {
      label: 'Reaps / reactivations',
      value: `${operator.reaps} / ${operator.reactivations}`,
      mono: true,
      hint: row.reapable ? 'past the inactivity window — reapable now' : undefined,
    },
  ]

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Operator'
        title={<Address value={row.address} full copy explorer className='text-[20px]' />}
        description='Identity, per-epoch history and every event this address produced, each with its block and transaction.'
        actions={
          <Link to={paths.operators()} className={buttonClasses('secondary', 'sm')}>
            All operators
          </Link>
        }
      />

      <StatRow>
        <StatCell label='Epochs served' value={row.epochsServed} mono hint={`${row.claims} slots claimed`} />
        <StatCell
          label='Contributions'
          value={row.contributions}
          mono
          tone='accent'
          hint={`participation ${formatParticipation(row.participation)}`}
        />
        <StatCell label='Partials' value={row.partials} mono hint='partial decryptions published' />
        <StatCell
          label='Finalize / combine'
          value={`${row.finalizations} / ${row.combines}`}
          mono
          hint='attributed through the transaction sender'
        />
      </StatRow>

      <div className='grid gap-4 lg:grid-cols-3'>
        <Panel className='lg:col-span-2' label='Identity' title='Registry record'>
          <KeyValue items={identity} columns={2} />
        </Panel>
        <Panel label='Trend' title='Participation over epochs'>
          <div className='flex flex-col gap-4'>
            <div>
              <div className='label-caps mb-1 text-[10px] text-pewter'>contributions / claims, cumulative</div>
              <div className='flex items-center gap-3'>
                <Sparkline values={participation} width={160} height={36} ariaLabel='participation trend' />
                <span className='font-mono text-lg tnum text-emerald'>{formatParticipation(row.participation)}</span>
              </div>
            </div>
            <div>
              <div className='label-caps mb-1 text-[10px] text-pewter'>partials per epoch</div>
              <div className='flex items-center gap-3'>
                <Sparkline
                  values={partials}
                  width={160}
                  height={36}
                  color={CHART_COLORS.teal}
                  ariaLabel='partials per epoch'
                />
                <span className='font-mono text-lg tnum text-silver'>{row.partials}</span>
              </div>
            </div>
            <p className='text-[11px] leading-relaxed text-ash'>
              A claimed slot is a promise to publish a contribution in the key-assembly window; participation is how
              often this operator kept it.
            </p>
          </div>
        </Panel>
      </div>

      <Panel
        label='History'
        title='Per-epoch record'
        description='Every epoch this operator claimed a slot in or contributed to, newest first.'
        bodyClassName='p-0'
      >
        <OperatorHistoryTable rows={historyRows} />
      </Panel>

      <Panel
        label='Log'
        title='Events'
        description={`${detail.events.length} events attributed to this address, newest first.`}
      >
        <OperatorEventLog events={detail.events} nonceOf={nonceOf} />
      </Panel>
    </Stack>
  )
}
