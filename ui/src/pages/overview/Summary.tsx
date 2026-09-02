import type { ReactNode } from 'react'
import { Link } from 'react-router-dom'
import { Address, BlockCell, Card, StatCell, StatRow } from '~kit'
import type { NetworkStats } from '~indexer/selectors'
import { blocksToDuration } from '~lib/format'
import { paths } from '~routes/paths'
import { PhaseBadge } from '~pages/epochs/PhaseBadge'
import { blockOrNull, formatCountdown, nextEpochCountdown, phaseCountdown } from '~pages/epochs/cadence'

/**
 * Chain, manager, head block and the cadence clock — the four facts that say
 * *which* deployment this is and whether it is moving.
 */
export function HeaderStrip({ stats, loading }: { stats: NetworkStats; loading: boolean }) {
  const next = nextEpochCountdown(stats)
  const head = blockOrNull(stats.headBlock)
  const indexed = blockOrNull(stats.lastIndexedBlock)
  return (
    <Card className='p-4'>
      <div className='grid grid-cols-2 gap-x-6 gap-y-4 lg:grid-cols-4'>
        <Cell label='Chain'>
          <span className='font-mono text-[13px] text-ghost'>{stats.chainName}</span>
          <span className='ml-2 font-mono text-[11px] text-ash'>#{stats.chainId}</span>
        </Cell>
        <Cell label='Manager'>
          <Address value={stats.managerAddress} chars={6} />
        </Cell>
        <Cell label='Block'>
          {loading && head == null ? (
            <span className='font-mono text-[13px] text-ash'>syncing…</span>
          ) : (
            <BlockCell block={head} suffix={indexed != null && head != null && indexed < head ? '· indexing' : ''} />
          )}
        </Cell>
        <Cell label='Next epoch'>
          <span className='font-mono text-[13px] text-silver'>{formatCountdown(next, stats.blockTimeSeconds)}</span>
          {next ? (
            <span className='ml-2 font-mono text-[11px] text-ash'>
              at {next.targetBlock.toLocaleString()}
              {next.source === 'cadence' ? ' (cadence)' : ''}
            </span>
          ) : null}
        </Cell>
      </div>
    </Card>
  )
}

function Cell({ label, children }: { label: string; children: ReactNode }) {
  return (
    <div className='min-w-0'>
      <div className='label-caps mb-1.5 text-[10px] text-pewter'>{label}</div>
      <div className='truncate'>{children}</div>
    </div>
  )
}

/** The five numbers that answer "is the network healthy right now?". */
export function StatusCards({ stats, loading }: { stats: NetworkStats; loading: boolean }) {
  const newest = stats.newestEpoch
  const countdown = phaseCountdown(newest, stats.headBlock, stats.epochDurationBlocks)
  const decryptedPct =
    stats.ciphertexts > 0 ? `${Math.round((stats.ciphertextsDecrypted / stats.ciphertexts) * 100)}%` : '—'

  return (
    <StatRow className='lg:grid-cols-5'>
      <StatCell
        label='Newest epoch'
        loading={loading && !newest}
        value={
          newest ? (
            <Link to={paths.epoch(newest.id)} className='transition-colors hover:text-emerald'>
              #{newest.nonce}
            </Link>
          ) : (
            '—'
          )
        }
        mono
        aside={newest ? <PhaseBadge phase={newest.status} size='sm' /> : null}
        hint={
          countdown
            ? `${countdown.label} ${countdown.passed ? 'now' : `in ${countdown.blocks.toLocaleString()} blocks`}`
            : newest
              ? 'nothing pending'
              : 'no epoch created yet'
        }
      />
      <StatCell
        label='Live epochs'
        loading={loading}
        value={stats.epochsLive}
        mono
        tone={stats.epochsLive > 0 ? 'accent' : 'default'}
        hint={`of ${stats.epochs.toLocaleString()} created · ${stats.epochsAborted} aborted`}
      />
      <StatCell
        label='Operators'
        loading={loading}
        value={`${stats.operatorsActive} / ${stats.operatorsRegistered}`}
        mono
        hint={
          stats.inactivityWindow != null
            ? `active / registered · reaped after ${blocksToDuration(stats.inactivityWindow, stats.blockTimeSeconds)}`
            : 'active / registered'
        }
      />
      <StatCell
        label='Committee'
        loading={loading}
        value={
          stats.committeeSizeInForce != null ? `${stats.thresholdInForce ?? '?'} of ${stats.committeeSizeInForce}` : '—'
        }
        mono
        hint='threshold t of n in force'
      />
      <StatCell
        label='Decrypted'
        loading={loading}
        value={stats.ciphertextsDecrypted.toLocaleString()}
        mono
        hint={`of ${stats.ciphertexts.toLocaleString()} submitted · ${decryptedPct} all time`}
      />
    </StatRow>
  )
}
