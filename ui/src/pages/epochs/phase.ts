// Phase presentation, shared by the three pages of this stream (overview,
// epochs, epoch). It lives here rather than in `~kit` because it maps a
// *protocol* value onto the design system, and the kit knows nothing about the
// protocol. Promotion candidate once a second stream needs it.

import type { EpochPhaseName } from '~indexer/types'
import type { BadgeTone } from '~kit'
import type { PhaseColorKey } from '~kit/charts'

/** Short label for a badge or a chip. `title` carries the long form. */
export const PHASE_LABEL: Record<EpochPhaseName, string> = {
  none: 'none',
  'committee-selection': 'selection',
  'key-assembly': 'assembly',
  live: 'live',
  aborted: 'aborted',
  completed: 'closed',
}

/** What the short label stands for — the badge's tooltip. */
export const PHASE_TITLE: Record<EpochPhaseName, string> = {
  none: 'no epoch',
  'committee-selection': 'committee selection — operators are claiming slots',
  'key-assembly': 'key assembly — the committee is submitting contributions',
  live: 'live — every pool key is stored and can be claimed',
  aborted: 'aborted — the committee never filled',
  completed: 'closed — the service window has passed',
}

/**
 * Emerald is reserved for live/ok, so only `live` gets it; assembly is the one
 * amber "in progress" state and aborted the one red failure.
 */
export const PHASE_TONE: Record<EpochPhaseName, BadgeTone> = {
  none: 'neutral',
  'committee-selection': 'neutral',
  'key-assembly': 'warn',
  live: 'ok',
  aborted: 'danger',
  completed: 'neutral',
}

/** Chart palette key for a phase (`~kit/charts` PHASE_COLORS). */
export const PHASE_COLOR_KEY: Record<EpochPhaseName, PhaseColorKey> = {
  none: 'pending',
  'committee-selection': 'selection',
  'key-assembly': 'assembly',
  live: 'live',
  aborted: 'aborted',
  completed: 'closed',
}

/** Filter chips, in lifecycle order. `all` first. */
export const PHASE_FILTERS: Array<EpochPhaseName | 'all'> = [
  'all',
  'committee-selection',
  'key-assembly',
  'live',
  'aborted',
  'completed',
]

export function phaseLabel(phase: EpochPhaseName | 'all'): string {
  return phase === 'all' ? 'all' : PHASE_LABEL[phase]
}
