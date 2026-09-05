import type { ReactNode } from 'react'
import type { Hex } from 'viem'
import { Badge, BlockCell, Button, Card, CopyButton, EmptyState, KeyValue, TxCell, type KeyValueItem } from '~kit'
import { txMetaOf } from '~indexer/selectors'
import { useStore, useTxMeta } from '~data/hooks'
import { cn } from '~lib/cn'
import { STEPS, STEP_TITLES, type LogEntry, type StepId, type StepStatus, type TxState } from './machine'

/**
 * Hash, block and gas for one step's transaction.
 *
 * The receipt is captured when the write returns, but `useTxMeta` also asks
 * the indexer to resolve the same hash — so a walkthrough resumed in a new tab
 * (which has the state but not the receipts) still shows a block and a gas
 * figure once the scan reaches it.
 */
export function TxLine({ tx, label = 'transaction' }: { tx: TxState | null | undefined; label?: string }) {
  const store = useStore()
  useTxMeta([tx?.hash as Hex | undefined])
  if (!tx) return null
  const meta = txMetaOf(store, tx.hash as Hex)
  const block = meta?.blockNumber ?? tx.block
  const gas = meta?.gasUsed ?? tx.gasUsed
  return (
    <div className='flex flex-wrap items-center gap-x-4 gap-y-1 rounded-sm border border-charcoal bg-onyx/40 px-3 py-2'>
      <span className='label-caps text-[11px] text-pewter'>{label}</span>
      <TxCell hash={tx.hash} copy />
      <span className='font-mono text-[11px] tnum text-ash'>
        block <BlockCell block={block} className='text-[11px]' />
      </span>
      <span className='font-mono text-[11px] tnum text-ash'>gas {gas != null ? gas.toLocaleString() : '—'}</span>
      {tx.simulated ? (
        <Badge tone='warn' size='sm' title='Demo mode: this transaction was never broadcast'>
          simulated
        </Badge>
      ) : null}
    </div>
  )
}

/** A list of 256-bit words — the transcripts behind the "advanced" toggle. */
export function Words({ items, columns = 2 }: { items: Array<{ label: string; value: bigint }>; columns?: 1 | 2 }) {
  if (items.length === 0) return null
  return (
    <div className={cn('grid gap-x-6 gap-y-1', columns === 2 ? 'md:grid-cols-2' : 'grid-cols-1')}>
      {items.map((item) => (
        <div key={item.label} className='flex min-w-0 items-baseline gap-2 border-b border-charcoal py-1.5'>
          <span className='w-16 shrink-0 font-mono text-[11px] text-pewter'>{item.label}</span>
          <span className='min-w-0 flex-1 truncate font-mono text-[11px] tnum text-silver' title={item.value.toString()}>
            {item.value.toString()}
          </span>
          <CopyButton value={item.value.toString()} label={`Copy ${item.label}`} size={11} />
        </div>
      ))}
    </div>
  )
}

/** The collapsible transcript block each step hides behind "advanced". */
export function Transcript({ title, note, children }: { title: string; note?: ReactNode; children: ReactNode }) {
  return (
    <Card level='onyx' className='mt-4 border-dashed'>
      <div className='label-caps mb-2 text-[11px] text-pewter'>{title}</div>
      {note ? <p className='mb-3 text-[12px] leading-relaxed text-ash'>{note}</p> : null}
      {children}
    </Card>
  )
}

const LOG_TONES: Record<LogEntry['tone'], string> = {
  info: 'text-pewter',
  ok: 'text-emerald',
  warn: 'text-amber',
  danger: 'text-red',
}

/** Every action the walkthrough took, in order, with the block it took it at. */
export function ActivityLog({ entries }: { entries: LogEntry[] }) {
  if (entries.length === 0) {
    return <EmptyState compact title='Nothing yet' description='Every action you take is recorded here.' />
  }
  return (
    <ol className='m-0 flex list-none flex-col gap-2 p-0'>
      {[...entries].reverse().map((entry) => (
        <li key={entry.id} className='flex gap-3 border-b border-charcoal pb-2 last:border-b-0'>
          <span className='w-20 shrink-0 font-mono text-[11px] tnum text-ash'>
            {entry.block != null ? `#${entry.block}` : '—'}
          </span>
          <div className='min-w-0 flex-1'>
            <div className={cn('text-[12px] leading-relaxed', LOG_TONES[entry.tone])}>{entry.message}</div>
            <div className='mt-0.5 text-[10px] uppercase tracking-[0.1em] text-ash'>{entry.step}</div>
          </div>
          {entry.tx ? <TxCell hash={entry.tx} chars={4} className='shrink-0' /> : null}
        </li>
      ))}
    </ol>
  )
}

/**
 * The left rail: eight steps, their state, and back-navigation. A step is
 * clickable once the walkthrough has reached it (`furthest`); a skipped step —
 * the reveal of an automatic application — is labelled as such and opens its
 * explanation once it has been passed.
 */
export function StepRail({
  status,
  active,
  furthest,
  onSelect,
}: {
  status: (step: StepId) => StepStatus
  active: StepId
  furthest: StepId
  onSelect: (step: StepId) => void
}) {
  return (
    <ol className='m-0 flex list-none flex-col gap-1 p-0'>
      {STEPS.map((step, i) => {
        const state = status(step)
        const reachable = state !== 'todo' && STEPS.indexOf(step) <= STEPS.indexOf(furthest)
        return (
          <li key={step}>
            <button
              type='button'
              disabled={!reachable}
              onClick={() => onSelect(step)}
              className={cn(
                'flex w-full items-center gap-3 rounded-sm border px-3 py-2 text-left transition-colors',
                'disabled:cursor-not-allowed disabled:opacity-40',
                step === active
                  ? 'border-emerald/40 bg-emerald/5 text-ghost'
                  : 'border-transparent text-pewter hover:border-charcoal hover:bg-onyx'
              )}
            >
              <span
                className={cn(
                  'flex h-5 w-5 shrink-0 items-center justify-center rounded-full border font-mono text-[10px] tnum',
                  state === 'done'
                    ? 'border-emerald bg-emerald/15 text-emerald'
                    : step === active
                      ? 'border-emerald text-emerald'
                      : state === 'skipped'
                        ? 'border-dashed border-charcoal text-ash'
                        : 'border-charcoal text-ash'
                )}
              >
                {state === 'done' ? '✓' : state === 'skipped' ? '–' : i + 1}
              </span>
              <span className={cn('min-w-0 truncate text-[13px]', state === 'skipped' && 'text-ash')}>
                {STEP_TITLES[step]}
              </span>
              {state === 'skipped' ? (
                <span className='ml-auto shrink-0 text-[10px] uppercase tracking-[0.1em] text-ash'>skipped</span>
              ) : null}
            </button>
          </li>
        )
      })}
    </ol>
  )
}

/** Panel chrome shared by every step: title, blurb, body, primary action. */
export function StepPanel({
  step,
  title,
  intro,
  children,
  actions,
  error,
}: {
  step: StepId
  title?: string
  intro?: ReactNode
  children?: ReactNode
  actions?: ReactNode
  error?: string | null
}) {
  return (
    <Card>
      <div className='label-caps mb-2 text-emerald'>step {STEPS.indexOf(step) + 1} of {STEPS.length}</div>
      <h2 className='text-lg font-semibold tracking-tight text-ghost'>{title ?? STEP_TITLES[step]}</h2>
      {intro ? <div className='mt-2 max-w-2xl text-[13px] leading-relaxed text-ash'>{intro}</div> : null}
      {children ? <div className='mt-5'>{children}</div> : null}
      {error ? (
        <p className='mt-4 rounded-sm border border-red/30 bg-red/5 px-3 py-2 text-[12px] text-red'>{error}</p>
      ) : null}
      {actions ? <div className='mt-5 flex flex-wrap items-center gap-3'>{actions}</div> : null}
    </Card>
  )
}

/** Label → value rows for the small record blocks inside the steps. */
export function Record({ items }: { items: KeyValueItem[] }) {
  return <KeyValue items={items} />
}

export function NextButton({ onClick, children }: { onClick: () => void; children: ReactNode }) {
  return (
    <Button variant='ghost' onClick={onClick}>
      {children}
    </Button>
  )
}
