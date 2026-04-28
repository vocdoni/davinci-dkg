import { Box, Heading, HStack, Stack, Text } from '@chakra-ui/react'
import type { ReactNode } from 'react'

interface Props {
  /** The page H1. Sans, semibold, tight. */
  title: string
  /** Optional one-line subtitle in dim ink. */
  subtitle?: ReactNode
  /** Optional right-aligned action (link, button, etc). */
  action?: ReactNode
}

// Minimal page header used by every route. No eyebrow, no rule, no
// uppercase mono section label — those experiments were too loud. Just
// a quiet H1 with an optional subtitle on the next line.
export function PageHeader({ title, subtitle, action }: Props) {
  return (
    <Box mb={{ base: 8, md: 10 }}>
      <HStack justify='space-between' align='start' gap={4} wrap='wrap'>
        <Stack gap={2} flex='1' minW={0}>
          <Heading
            as='h1'
            fontSize={{ base: '3xl', md: '4xl' }}
            fontWeight={500}
            color='ink.0'
            letterSpacing='-0.02em'
            lineHeight='1.15'
          >
            {title}
          </Heading>
          {subtitle && (
            <Text fontSize={{ base: 'sm', md: 'md' }} color='ink.3' lineHeight='1.55' maxW='62ch'>
              {subtitle}
            </Text>
          )}
        </Stack>
        {action && <Box flexShrink={0}>{action}</Box>}
      </HStack>
    </Box>
  )
}
