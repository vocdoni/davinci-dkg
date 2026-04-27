import { Box, Heading, HStack, Spinner, Text } from '@chakra-ui/react'
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

const colorByStatus: Record<StepStatus, { border: string; bg: string; pill: string }> = {
  pending: { border: 'gray.800', bg: 'gray.950', pill: 'gray.700' },
  active: { border: 'cyan.700', bg: 'gray.900', pill: 'cyan.500' },
  done: { border: 'green.700', bg: 'gray.900', pill: 'green.500' },
  error: { border: 'red.700', bg: 'gray.900', pill: 'red.500' },
}

export function StepCard({ n, title, status, description, children }: Props) {
  const c = colorByStatus[status]
  return (
    <Box
      borderWidth='1px'
      borderColor={c.border}
      bg={c.bg}
      borderRadius='md'
      p={5}
      opacity={status === 'pending' ? 0.7 : 1}
      transition='opacity 0.2s, border-color 0.2s'
    >
      <HStack mb={description ? 1 : 4} gap={3} align='center'>
        <Box
          w={7}
          h={7}
          borderRadius='full'
          bg={c.pill}
          color='white'
          display='flex'
          alignItems='center'
          justifyContent='center'
          fontSize='xs'
          fontWeight='bold'
          flexShrink={0}
        >
          {status === 'done' ? '✓' : status === 'error' ? '✗' : n}
        </Box>
        <Heading size='sm' color={status === 'pending' ? 'gray.400' : 'white'}>
          {title}
        </Heading>
        {status === 'active' && <Spinner size='sm' color='cyan.300' ml='auto' />}
      </HStack>
      {description && (
        <Text fontSize='xs' color='gray.500' mb={4} pl={10}>
          {description}
        </Text>
      )}
      {children}
    </Box>
  )
}
