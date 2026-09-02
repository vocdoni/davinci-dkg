import type { EpochPhaseName } from '~indexer/types'
import { cn } from '~lib/cn'
import { PHASE_FILTERS, PHASE_TITLE, phaseLabel } from './phase'

export type PhaseFilterValue = EpochPhaseName | 'all'

export interface PhaseFilterProps {
  value: PhaseFilterValue
  onChange: (value: PhaseFilterValue) => void
  /** Row count per phase, so a chip that would empty the table says so. */
  counts: Record<string, number>
  total: number
}

/**
 * Filter chips. The kit has no chip primitive yet — this is a local one, and a
 * promotion candidate: the operators and applications tables want the same
 * control.
 */
export function PhaseFilter({ value, onChange, counts, total }: PhaseFilterProps) {
  return (
    <div className='flex flex-wrap items-center gap-2' role='group' aria-label='Filter by phase'>
      {PHASE_FILTERS.map((phase) => {
        const count = phase === 'all' ? total : (counts[phase] ?? 0)
        const active = value === phase
        return (
          <button
            key={phase}
            type='button'
            aria-pressed={active}
            title={phase === 'all' ? 'every epoch' : PHASE_TITLE[phase]}
            onClick={() => onChange(phase)}
            disabled={count === 0 && !active && phase !== 'all'}
            className={cn(
              'rounded-pill border px-3 py-1 text-[12px] transition-colors',
              active
                ? 'border-emerald/40 bg-emerald/10 text-emerald'
                : 'border-charcoal text-pewter hover:border-warm-gray hover:text-ghost',
              count === 0 && !active && phase !== 'all' && 'cursor-not-allowed opacity-40 hover:border-charcoal'
            )}
          >
            {phaseLabel(phase)}
            <span className='ml-1.5 font-mono tnum text-[11px] text-ash'>{count}</span>
          </button>
        )
      })}
    </div>
  )
}
