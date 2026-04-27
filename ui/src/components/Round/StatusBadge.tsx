import { Badge } from '@chakra-ui/react'
import { roundPhase, roundPhaseColor, roundPhaseLabel } from '~lib/round-utils'
import type { Round } from '@vocdoni/davinci-dkg-sdk'

export function StatusBadge({ round }: { round: Round }) {
  const phase = roundPhase(round)
  return (
    <Badge colorPalette={roundPhaseColor(phase)} variant='subtle'>
      {roundPhaseLabel(phase)}
    </Badge>
  )
}
