import { Box, HStack, Text } from '@chakra-ui/react'
import { Global } from '@emotion/react'
import { LuCheck, LuCircle, LuCircleDot } from 'react-icons/lu'
import { phaseSequence, roundPhase, roundPhaseLabel, type RoundPhase } from '~lib/round-utils'
import type { Round } from '@vocdoni/davinci-dkg-sdk'

// A halo that softly grows + fades around the current phase dot, drawing
// the eye to where the round is right now without being distracting. Two
// rules: keep the animation slow (1.6s) so it doesn't look frantic, and
// pin opacity to 0 at the end of the cycle so the badge below shows
// through cleanly.
//
// Defined as a global @keyframes so we can drive the pulse with a plain
// CSS `animation` string anywhere in the component without per-instance
// keyframe declarations.
const pulseStyles = `
@keyframes dkgPhasePulse {
  0%   { transform: scale(0.85); opacity: 0.65; }
  70%  { transform: scale(1.7); opacity: 0; }
  100% { transform: scale(1.7); opacity: 0; }
}
`

// Horizontal four-step timeline (Registration → Contribution → Finalized →
// Completed). Aborted rounds short-circuit and render a single red badge.
export function PhaseTimeline({ round }: { round: Round }) {
  const current = roundPhase(round)

  if (current === 'aborted') {
    return (
      <HStack gap={2} color='red.300'>
        <LuCircleDot />
        <Text fontSize='sm' fontWeight='semibold'>
          Aborted
        </Text>
      </HStack>
    )
  }

  const reachedIdx = phaseSequence.indexOf(current)
  return (
    <>
      <Global styles={pulseStyles} />
      <HStack gap={0} align='center' wrap='wrap'>
        {phaseSequence.map((phase, i) => {
          const state: 'past' | 'current' | 'future' =
            reachedIdx < 0 ? 'future' : i < reachedIdx ? 'past' : i === reachedIdx ? 'current' : 'future'
          return (
            <HStack key={phase} gap={0}>
              <PhaseDot phase={phase} state={state} />
              {i < phaseSequence.length - 1 && (
                <Box w={{ base: 6, md: 12 }} h='1px' bg={i < reachedIdx ? 'cyan.700' : 'gray.700'} />
              )}
            </HStack>
          )
        })}
      </HStack>
    </>
  )
}

function PhaseDot({ phase, state }: { phase: RoundPhase; state: 'past' | 'current' | 'future' }) {
  const color = state === 'past' ? 'cyan.500' : state === 'current' ? 'cyan.300' : 'gray.600'
  const Icon = state === 'past' ? LuCheck : state === 'current' ? LuCircleDot : LuCircle
  return (
    <HStack gap={1.5} px={1}>
      <Box position='relative' color={color} display='inline-flex'>
        <Icon />
        {state === 'current' && (
          <Box
            position='absolute'
            inset={0}
            borderRadius='full'
            bg='cyan.400'
            css={{ animation: 'dkgPhasePulse 1.6s ease-out infinite' }}
            pointerEvents='none'
          />
        )}
      </Box>
      <Text fontSize='xs' color={state === 'future' ? 'gray.500' : 'gray.300'} fontWeight={state === 'current' ? 'semibold' : 'normal'}>
        {roundPhaseLabel(phase)}
      </Text>
    </HStack>
  )
}
