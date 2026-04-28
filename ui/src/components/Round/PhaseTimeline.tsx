import { Box, Flex, Stack, Text } from '@chakra-ui/react'
import { LuCheck } from 'react-icons/lu'
import { phaseSequence, roundPhase, roundPhaseLabel, type RoundPhase } from '~lib/round-utils'
import type { Round } from '@vocdoni/davinci-dkg-sdk'

// Editorial four-step timeline.
//   ── past phases     filled gold dot, check glyph, name in dim bone
//   ── current phase   pulsing gold halo, name in bone, semibold
//   ── future phases   open ring, name in dim ink
//
// Connector rules between phases are 1px hairlines, half-tinted to gold for
// the segments behind the cursor. Aborted rounds short-circuit to a single
// coral row.
//
// The pulse keyframes live in the global theme (`dkgPhasePulse`) so this
// component doesn't need its own <Global> tag.
export function PhaseTimeline({ round }: { round: Round }) {
  const current = roundPhase(round)

  if (current === 'aborted') {
    return (
      <Flex
        align='center'
        gap={3}
        py={3}
        px={4}
        borderWidth='1px'
        borderColor='danger.border'
        bg='danger.bg'
        borderRadius='lg'
      >
        <Box w='8px' h='8px' borderRadius='full' bg='danger.fg' />
        <Text
          fontFamily='mono'
          fontSize='xs'
          color='danger.fg'
          letterSpacing='0.06em'
          textTransform='uppercase'
        >
          Round aborted
        </Text>
      </Flex>
    )
  }

  const reachedIdx = phaseSequence.indexOf(current)

  return (
    <Flex align='start' wrap='wrap' rowGap={4}>
      {phaseSequence.map((phase, i) => {
        const state: 'past' | 'current' | 'future' =
          reachedIdx < 0 ? 'future' : i < reachedIdx ? 'past' : i === reachedIdx ? 'current' : 'future'
        const isLast = i === phaseSequence.length - 1
        return (
          <Flex key={phase} align='start' flex={isLast ? '0 0 auto' : '1 1 auto'} minW={0}>
            <PhaseDot phase={phase} state={state} index={i + 1} />
            {!isLast && (
              <Box flex='1' minW={{ base: '20px', md: '40px' }} pt='14px' px={2}>
                <Box
                  h='1px'
                  bg={i < reachedIdx ? 'accent.fg' : 'rule'}
                  opacity={i < reachedIdx ? 0.5 : 1}
                />
              </Box>
            )}
          </Flex>
        )
      })}
    </Flex>
  )
}

function PhaseDot({
  phase,
  state,
  index,
}: {
  phase: RoundPhase
  state: 'past' | 'current' | 'future'
  index: number
}) {
  const dotBg =
    state === 'past' ? 'accent.dim' : state === 'current' ? 'accent.fg' : 'transparent'
  const dotBorder =
    state === 'past' ? 'accent.dim' : state === 'current' ? 'accent.fg' : 'border.strong'
  const labelColor = state === 'future' ? 'ink.4' : state === 'current' ? 'ink.0' : 'ink.2'
  const labelWeight = state === 'current' ? 600 : 400

  return (
    <Stack gap={1.5} align='center' minW='90px' flexShrink={0}>
      <Box position='relative' w='28px' h='28px'>
        {state === 'current' && (
          <Box
            position='absolute'
            inset={0}
            borderRadius='full'
            bg='accent.fg'
            opacity={0.4}
            css={{ animation: 'dkgPhasePulse 1.8s ease-out infinite' }}
            pointerEvents='none'
          />
        )}
        <Box
          position='relative'
          w='28px'
          h='28px'
          borderRadius='full'
          bg={dotBg}
          borderWidth='1px'
          borderColor={dotBorder}
          display='flex'
          alignItems='center'
          justifyContent='center'
          color={state === 'past' ? 'ink.0' : 'ink.0'}
          fontSize='10px'
          fontFamily='mono'
          fontWeight={500}
        >
          {state === 'past' ? <LuCheck size={12} /> : <Text className='dkg-tabular'>{index}</Text>}
        </Box>
      </Box>
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color={labelColor}
        fontWeight={labelWeight}
        letterSpacing='0.06em'
        textTransform='uppercase'
        textAlign='center'
        whiteSpace='nowrap'
      >
        {roundPhaseLabel(phase)}
      </Text>
    </Stack>
  )
}
