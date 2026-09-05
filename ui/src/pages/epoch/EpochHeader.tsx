import { useMemo } from 'react'
import { Address, BlockCell, Hash, KeyValue, Panel, SectionHeader, TxCell, type KeyValueItem } from '~kit'
import { BlockTimeline, type BlockWindow } from '~kit/charts'
import type { EpochDetail } from '~indexer/selectors'
import { blockOrNull, formatCountdown, phaseCountdown } from '~pages/epochs/cadence'
import { PhaseBadge } from '~pages/epochs/PhaseBadge'

const fmt = (block: number | null | undefined): string => {
  const value = blockOrNull(block)
  return value == null ? '—' : value.toLocaleString()
}

export interface EpochHeaderProps {
  detail: EpochDetail
  head: number
  blockTimeSeconds: number
  epochDurationBlocks: number | null
}

/**
 * Identity plus the lifecycle on the block axis. The four windows are absolute
 * block counts fixed at deploy time, so the picture answers "what is this epoch
 * waiting for?" without reading a single number.
 */
export function EpochHeader({ detail, head, blockTimeSeconds, epochDurationBlocks }: EpochHeaderProps) {
  const { epoch, row, windows } = detail
  const countdown = phaseCountdown(epoch, head, epochDurationBlocks)
  const current = blockOrNull(head)

  const timeline = useMemo<BlockWindow[]>(() => {
    const span = (label: string, from: number | null, to: number | null, tone: string, detailText: string) =>
      from != null && to != null && Number.isFinite(from) && Number.isFinite(to) && to > from
        ? [{ label, from, to, tone, detail: detailText }]
        : []
    return [
      ...span(
        'committee selection',
        windows.startBlock,
        windows.committeeSelectionDeadline,
        'selection',
        `${row.claims} of ${row.committeeSize} slots claimed · lottery seed at block ${windows.seedBlock.toLocaleString()}`
      ),
      ...span(
        'key assembly',
        windows.committeeSelectionDeadline,
        windows.keyAssemblyDeadline,
        'assembly',
        `${row.contributions} of ${row.committeeSize} contributions · m_min = ${row.minValidContributions}`
      ),
      ...span(
        'finalize gap',
        windows.keyAssemblyDeadline,
        windows.liveNotBefore,
        'finalize',
        'cooldown before finalizeEpoch may run'
      ),
      ...span(
        'live',
        windows.liveNotBefore,
        windows.endBlock,
        'live',
        `${row.ciphertexts} ciphertexts · ${row.applications} applications`
      ),
    ]
  }, [windows, row])

  const items: KeyValueItem[] = [
    { label: 'epoch id', value: <Hash value={epoch.id} full copy />, mono: true },
    { label: 'nonce', value: `#${epoch.nonce}`, mono: true },
    {
      label: 'creator',
      value: <Address value={epoch.creator} chars={6} />,
      hint: 'won the permissionless createEpoch race',
    },
    {
      label: 'created',
      value: (
        <span className='inline-flex items-center gap-2'>
          <BlockCell block={blockOrNull(epoch.createdBlock)} />
          {epoch.createdTx ? <TxCell hash={epoch.createdTx} /> : null}
        </span>
      ),
    },
    {
      label: 'policy',
      value: `t = ${row.threshold} · m_min = ${row.minValidContributions} · n = ${row.committeeSize}`,
      mono: true,
      hint: 'threshold, minimum valid contributions, committee size',
    },
    {
      label: 'deadlines',
      value: `selection ≤ ${fmt(windows.committeeSelectionDeadline)} · assembly ≤ ${fmt(windows.keyAssemblyDeadline)} · live ≥ ${fmt(windows.liveNotBefore)}`,
      mono: true,
      hint: windows.endBlock != null ? `service window ends at block ${windows.endBlock.toLocaleString()}` : undefined,
    },
  ]

  return (
    <>
      <SectionHeader
        size='page'
        label='Epoch'
        title={`Epoch #${epoch.nonce}`}
        description='One DKG run: a lottery, a committee, a pool of keys, and every decryption it served.'
        actions={
          <span className='flex items-center gap-3'>
            <PhaseBadge phase={epoch.status} />
            {countdown ? (
              <span className='font-mono text-[12px] text-ash'>
                {countdown.label} ·{' '}
                {formatCountdown(countdown, blockTimeSeconds, epoch.status === 'live' ? 'ended' : 'passed')}
              </span>
            ) : null}
          </span>
        }
      />

      <Panel
        label='Lifecycle'
        title='Windows on the block axis'
        description='Committee selection, key assembly and the finalize gap are fixed block budgets; everything left over is service time.'
      >
        <BlockTimeline
          windows={timeline}
          current={current}
          start={blockOrNull(windows.startBlock) ?? undefined}
          end={blockOrNull(windows.endBlock) ?? undefined}
          height={120}
        />
        <KeyValue items={items} columns={2} className='mt-5' />
      </Panel>
    </>
  )
}
