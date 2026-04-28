import { Box, Stack, Text } from '@chakra-ui/react'
import type { ReactNode } from 'react'

interface Props {
  label: string
  value: ReactNode
  hint?: ReactNode
  /** Visual tone. Default reads as a neutral counter; "live" tints the
   *  value in phosphor green for actively-ticking metrics like block #. */
  tone?: 'neutral' | 'live' | 'accent'
}

// Editorial KPI card. Small-caps mono label up top, large mono numeric value
// in the middle, italic serif hint at the bottom. The hairline border + the
// inset shadow give it just enough depth to feel like a card without
// shouting "I am a card!". Each card gets a faint accent rule on the left
// to mark it as a quantitative block (vs prose).
export function StatCard({ label, value, hint, tone = 'neutral' }: Props) {
  const valueColor = tone === 'live' ? 'live.fg' : tone === 'accent' ? 'accent.bright' : 'ink.0'
  return (
    <Box
      position='relative'
      borderWidth='1px'
      borderColor='border.subtle'
      bg='surface'
      borderRadius='lg'
      p={{ base: 4, md: 5 }}
      boxShadow='inset'
      transition='border-color 0.15s'
      _hover={{ borderColor: 'border' }}
    >
      {/* Left rule — 2px wide, half-height, accent or live tone. Makes the
          card read as a labeled measurement, not a generic tile. */}
      <Box
        position='absolute'
        left={0}
        top='25%'
        bottom='25%'
        w='2px'
        bg={tone === 'live' ? 'live.fg' : 'accent.fg'}
        opacity={0.6}
        borderRightRadius='full'
      />
      <Stack gap={2}>
        <Text
          fontFamily='mono'
          fontSize='2xs'
          color='ink.3'
          letterSpacing='0.08em'
          textTransform='uppercase'
        >
          {label}
        </Text>
        <Text
          className='dkg-tabular'
          fontFamily='mono'
          fontSize={{ base: 'xl', md: '2xl' }}
          fontWeight={500}
          color={valueColor}
          lineHeight='1.1'
        >
          {value}
        </Text>
        {hint && (
          <Text
           
           
            fontSize='xs'
            color='ink.3'
            lineHeight='1.4'
          >
            {hint}
          </Text>
        )}
      </Stack>
    </Box>
  )
}
