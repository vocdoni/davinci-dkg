import type { EpochPhaseName } from '~indexer/types'
import { Badge } from '~kit'
import { PHASE_LABEL, PHASE_TITLE, PHASE_TONE } from './phase'

export interface PhaseBadgeProps {
  phase: EpochPhaseName
  size?: 'sm' | 'md'
}

/** The one badge that says where an epoch is in its lifecycle. */
export function PhaseBadge({ phase, size = 'md' }: PhaseBadgeProps) {
  return (
    <Badge tone={PHASE_TONE[phase]} size={size} dot={phase === 'live'} title={PHASE_TITLE[phase]}>
      {PHASE_LABEL[phase]}
    </Badge>
  )
}
