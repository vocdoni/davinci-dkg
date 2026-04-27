import { Center, Spinner, Text, VStack } from '@chakra-ui/react'
import type { ReactNode } from 'react'
import { ErrorReportButton } from '~components/Debug/ErrorReportButton'

interface Props {
  isLoading: boolean
  isError: boolean
  isEmpty?: boolean
  error?: unknown
  emptyMessage?: string
  children: ReactNode
}

// One wrapper, three states. Every page that fetches over react-query
// renders its content inside this so the loading / error / empty UX is
// uniform — no per-page reinvention.
export function QueryDataLayout({ isLoading, isError, isEmpty, error, emptyMessage, children }: Props) {
  if (isLoading) {
    return (
      <Center py={16}>
        <Spinner size='lg' color='cyan.400' />
      </Center>
    )
  }
  if (isError) {
    const msg = error instanceof Error ? error.message : String(error ?? 'Unknown error')
    return (
      <Center py={16}>
        <VStack gap={3} maxW='md'>
          <Text fontSize='md' fontWeight='semibold' color='red.300'>
            Couldn't load this view
          </Text>
          <Text fontSize='sm' color='gray.400' textAlign='center'>
            {msg}
          </Text>
          <ErrorReportButton error={error} />
        </VStack>
      </Center>
    )
  }
  if (isEmpty) {
    return (
      <Center py={16}>
        <Text fontSize='sm' color='gray.500'>
          {emptyMessage ?? 'Nothing to show.'}
        </Text>
      </Center>
    )
  }
  return <>{children}</>
}
