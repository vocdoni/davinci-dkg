import { Link } from 'react-router-dom'
import { Address, Badge, BlockCell, Callout, Hash, KeyValue, Panel, TxCell, type BadgeTone, type KeyValueItem } from '~kit'
import type { EpochDetail, PoolSlotRow, PoolSlotState } from '~indexer/selectors'
import { POOL_SIZE } from '~indexer/types'
import { bigIntToHex } from '~lib/format'
import { cn } from '~lib/cn'
import { paths } from '~routes/paths'
import { FINALIZE_TRANSCRIPT_WORDS, blockOrNull } from '~pages/epochs/cadence'

const SLOT_TONE: Record<PoolSlotState, BadgeTone> = {
  pending: 'neutral',
  free: 'accent',
  claimed: 'ok',
}

const SLOT_LABEL: Record<PoolSlotState, string> = {
  pending: 'reading…',
  free: 'free',
  claimed: 'claimed',
}

/**
 * The epoch's pool of `POOL_SIZE` keys. One proof-carrying `finalizeEpoch`
 * proves and stores every key `P_j` together with the Merkle root of the
 * committee's share commitments, so the whole pool exists from the block the
 * epoch went Live; each key is then claimed by exactly one application at
 * registration. This panel is the epoch's key material.
 */
export function PoolPanel({ detail }: { detail: EpochDetail }) {
  const { finalization, pool, poolNext, poolKnown, poolClaimed, row } = detail

  if (!finalization) {
    return (
      <Panel label='Pool' title='Pool keys' description={`${POOL_SIZE} keys per epoch, all proven by one finalization.`}>
        <Callout tone='info' title='No pool yet'>
          The epoch has not been finalized. finalizeEpoch — one Groth16 proof over the accepted contributions — may
          run once the key-assembly window has closed, the finalize gap has passed, and at least m_min ={' '}
          {row.minValidContributions} contributions are on chain; it stores every key and share root at once.
        </Callout>
      </Panel>
    )
  }

  const items: KeyValueItem[] = [
    {
      label: 'finalizer',
      value: finalization.by ? (
        <Address value={finalization.by} chars={6} to={paths.operator(finalization.by)} />
      ) : (
        <span className='text-ash'>resolving…</span>
      ),
      hint: 'EpochLive names nobody — this is the transaction sender',
    },
    {
      label: 'finalized',
      value: (
        <span className='inline-flex items-center gap-2'>
          <BlockCell block={blockOrNull(finalization.block)} />
          {finalization.tx ? <TxCell hash={finalization.tx} /> : null}
        </span>
      ),
      hint: 'every key below exists from this block on',
    },
    {
      label: 'contributions proven',
      value: finalization.contributionCount,
      mono: true,
      hint: `accepted contributors the pool keys are dealt from · m_min = ${row.minValidContributions}`,
    },
    {
      label: 'finalize gas',
      value: finalization.gasUsed != null ? finalization.gasUsed.toLocaleString() : '—',
      mono: true,
      hint: `one verifier, ${POOL_SIZE} Merkle trees and ${2 * POOL_SIZE} key/root writes`,
    },
    {
      label: 'keys claimed',
      value: `${poolClaimed} / ${POOL_SIZE}`,
      mono: true,
      hint: poolNext >= POOL_SIZE ? 'pool exhausted — the next epoch may open early' : `next index ${poolNext}`,
    },
    {
      label: 'finalize transcript',
      value: `${FINALIZE_TRANSCRIPT_WORDS.toLocaleString()} words`,
      mono: true,
      hint: `2·MaxN + MaxK·(2 + 2·MaxN) — ${(FINALIZE_TRANSCRIPT_WORDS * 32).toLocaleString()} bytes of calldata, holding every P_j and every member's share commitment`,
    },
  ]

  return (
    <Panel
      label='Pool'
      title='Pool keys'
      description={
        poolKnown < POOL_SIZE
          ? `finalizeEpoch proved and stored all ${POOL_SIZE} keys at once; the explorer is still reading them from the manager.`
          : `finalizeEpoch proved and stored all ${POOL_SIZE} keys at once. Each key P_j below is claimed by one application at registration.`
      }
    >
      <KeyValue items={items} columns={2} />
      <ul className='m-0 mt-5 grid list-none gap-2 p-0 sm:grid-cols-2 xl:grid-cols-4' aria-label='Pool keys'>
        {pool.map((slot) => (
          <PoolSlot key={slot.index} slot={slot} epoch={detail.epoch.id} />
        ))}
      </ul>
    </Panel>
  )
}

function PoolSlot({ slot, epoch }: { slot: PoolSlotRow; epoch: string }) {
  const claimed = slot.state === 'claimed'
  return (
    <li
      aria-label={`pool key ${slot.index}`}
      className={cn(
        'flex flex-col gap-2 rounded-sm border px-3 py-2.5',
        claimed ? 'border-emerald/30 bg-emerald/5' : slot.state === 'free' ? 'border-charcoal' : 'border-charcoal/60'
      )}
    >
      <div className='flex items-center justify-between gap-2'>
        <span className='font-mono text-[12px] text-silver'>key {slot.index}</span>
        <Badge size='sm' tone={SLOT_TONE[slot.state]}>
          {SLOT_LABEL[slot.state]}
        </Badge>
      </div>
      <div className='font-mono text-[11px] text-ash'>
        {slot.key ? (
          <span className='inline-flex flex-col gap-0.5'>
            <span className='inline-flex items-center gap-1'>
              x <Hash value={bigIntToHex(slot.key.x)} chars={6} />
            </span>
            <span className='inline-flex items-center gap-1'>
              y <Hash value={bigIntToHex(slot.key.y)} chars={6} />
            </span>
          </span>
        ) : (
          <span>P_{slot.index} not read yet</span>
        )}
      </div>
      {slot.claimedBy ? (
        <div className='flex flex-wrap items-center gap-x-2 text-[11px] text-ash'>
          <span>claimed by</span>
          <Link to={paths.application(epoch, slot.claimedBy)} className='hover:text-emerald'>
            <Hash value={slot.claimedBy} chars={6} copy={false} />
          </Link>
          {slot.claimedBlock != null ? <BlockCell block={slot.claimedBlock} /> : null}
        </div>
      ) : null}
    </li>
  )
}
