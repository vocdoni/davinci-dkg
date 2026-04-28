import { Box, HStack, Text } from '@chakra-ui/react'
import { roundPhase, roundPhaseLabel, type RoundPhase } from '~lib/round-utils'
import type { Round } from '@vocdoni/davinci-dkg-sdk'

// Editorial status badge: a tiny indicator dot + small mono uppercase
// label. Five tones, each with a paired (fg / bg / dot) mapping. The dot
// pulses for "current" / "in-progress" phases (registration, contribution)
// so a glance at the badge tells you whether the round is moving.
type Tone = {
  fg: string
  bg: string
  border: string
  dot: string
  pulse: boolean
}

const tonesByPhase: Record<RoundPhase, Tone> = {
  registration: { fg: 'amber.300', bg: 'warn.bg', border: 'rgba(240, 198, 116, 0.30)', dot: 'amber.300', pulse: true },
  contribution: { fg: 'accent.bright', bg: 'accent.bg', border: 'accent.border', dot: 'accent.fg', pulse: true },
  finalized: { fg: 'live.fg', bg: 'live.bg', border: 'rgba(134, 239, 172, 0.30)', dot: 'live.fg', pulse: false },
  completed: { fg: 'ink.2', bg: 'surface.raised', border: 'border', dot: 'ink.3', pulse: false },
  aborted: { fg: 'danger.fg', bg: 'danger.bg', border: 'danger.border', dot: 'danger.fg', pulse: false },
  unknown: { fg: 'ink.3', bg: 'surface.raised', border: 'border', dot: 'ink.4', pulse: false },
}

export function StatusBadge({ round }: { round: Round }) {
  const phase = roundPhase(round)
  const t = tonesByPhase[phase]
  return (
    <HStack
      as='span'
      display='inline-flex'
      gap={2}
      px={2.5}
      py='4px'
      borderRadius='full'
      borderWidth='1px'
      borderColor={t.border}
      bg={t.bg}
    >
      <Box position='relative' w='6px' h='6px'>
        <Box position='absolute' inset={0} borderRadius='full' bg={t.dot} />
        {t.pulse && (
          <Box
            position='absolute'
            inset={0}
            borderRadius='full'
            bg={t.dot}
            opacity={0.5}
            transformOrigin='center'
            css={{ animation: 'dkgPhasePulse 1.8s ease-out infinite' }}
            pointerEvents='none'
          />
        )}
      </Box>
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color={t.fg}
        letterSpacing='0.08em'
        textTransform='uppercase'
        lineHeight='1'
      >
        {roundPhaseLabel(phase)}
      </Text>
    </HStack>
  )
}
