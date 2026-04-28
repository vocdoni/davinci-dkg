import { Box, Heading, HStack, Spinner, Text } from '@chakra-ui/react'
import { LuCheck, LuX } from 'react-icons/lu'
import type { ReactNode } from 'react'

export type StepStatus = 'pending' | 'active' | 'done' | 'error'

interface Props {
  n: number
  title: string
  status: StepStatus
  /** Optional one-sentence subtitle explaining what this step does. */
  description?: string
  children: ReactNode
}

// Editorial step card. The chrome stays quiet; the visual signal that "this
// step is active" is carried by a 2px gold rule running down the left edge,
// not by changing the card background. Pending steps fade slightly so the
// reader's eye lands on what's open.
//
// The number badge mirrors the home page's phase-number treatment: large
// mono digits when active/pending, a check (or ✗) when complete.
const toneByStatus: Record<
  StepStatus,
  { rule: string; numColor: string; titleColor: string; opacity: number }
> = {
  pending: { rule: 'transparent', numColor: 'ink.4', titleColor: 'ink.3', opacity: 0.6 },
  active: { rule: 'accent.fg', numColor: 'accent.fg', titleColor: 'ink.0', opacity: 1 },
  done: { rule: 'live.fg', numColor: 'live.fg', titleColor: 'ink.0', opacity: 1 },
  error: { rule: 'danger.fg', numColor: 'danger.fg', titleColor: 'ink.0', opacity: 1 },
}

export function StepCard({ n, title, status, description, children }: Props) {
  const tone = toneByStatus[status]
  return (
    <Box
      position='relative'
      borderWidth='1px'
      borderColor='border.subtle'
      bg='surface'
      borderRadius='lg'
      p={{ base: 5, md: 6 }}
      pl={{ base: 6, md: 7 }}
      opacity={tone.opacity}
      transition='opacity 0.2s ease'
      boxShadow='inset'
      _hover={{ borderColor: status === 'pending' ? 'border.subtle' : 'border' }}
    >
      {/* Left rule — the only "active state" colour. Replaces the previous
          full-card border tint, which made every step shout for attention. */}
      <Box
        position='absolute'
        left={0}
        top={5}
        bottom={5}
        w='2px'
        bg={tone.rule}
        borderRightRadius='full'
      />

      <HStack mb={description ? 1 : 4} gap={4} align='baseline'>
        <Box
          flexShrink={0}
          w={8}
          h={8}
          borderRadius='full'
          borderWidth='1px'
          borderColor={status === 'pending' ? 'border.strong' : tone.numColor}
          bg={status === 'done' || status === 'error' ? tone.numColor : 'transparent'}
          color={status === 'done' || status === 'error' ? 'canvas' : tone.numColor}
          display='flex'
          alignItems='center'
          justifyContent='center'
          fontSize='xs'
          fontFamily='mono'
          fontWeight={500}
        >
          {status === 'done' ? (
            <LuCheck size={14} strokeWidth={2.5} />
          ) : status === 'error' ? (
            <LuX size={14} strokeWidth={2.5} />
          ) : (
            <Text className='dkg-tabular'>{n}</Text>
          )}
        </Box>
        <Heading
          as='h3'
         
          fontSize={{ base: 'lg', md: 'xl' }}
          fontWeight={500}
          color={tone.titleColor}
          letterSpacing='-0.01em'
          flex='1'
        >
          {title}
        </Heading>
        {status === 'active' && (
          <Spinner size='sm' color='accent.fg' borderWidth='2px' />
        )}
      </HStack>
      {description && (
        <Text
         
          fontSize='sm'
         
          color='ink.3'
          mb={5}
          pl={12}
          maxW='62ch'
          lineHeight='1.55'
        >
          {description}
        </Text>
      )}
      {children}
    </Box>
  )
}
