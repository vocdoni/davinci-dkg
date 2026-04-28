import { Box, Center, Spinner, Stack, Text } from '@chakra-ui/react'
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
// uniform across the app — no per-page reinvention.
//
// Loading: a soft gold spinner, no text. The spinner colour matches the
// rest of the accent system, so loading reads as part of the page rather
// than as a generic UI overlay.
//
// Error: a typeset card. Header in display serif, message in italic body
// serif, copy-error-report button beneath. Designed to feel like an
// erratum sheet, not an alert.
//
// Empty: italic body serif, centered, dim. Distinguishable from "loading"
// because it has no spinner.
export function QueryDataLayout({ isLoading, isError, isEmpty, error, emptyMessage, children }: Props) {
  if (isLoading) {
    return (
      <Center py={16}>
        <Spinner size='lg' color='accent.fg' borderWidth='2px' />
      </Center>
    )
  }
  if (isError) {
    const msg = error instanceof Error ? error.message : String(error ?? 'Unknown error')
    return (
      <Center py={12}>
        <Box
          maxW='md'
          borderWidth='1px'
          borderColor='danger.border'
          bg='danger.bg'
          borderRadius='lg'
          p={6}
          textAlign='center'
        >
          <Stack gap={3}>
            <Text
             
              fontSize='lg'
              fontWeight={500}
              color='danger.fg'
              letterSpacing='-0.01em'
            >
              Couldn't load this view
            </Text>
            <Text fontSize='sm' color='ink.2'>
              {msg}
            </Text>
            <Center pt={1}>
              <ErrorReportButton error={error} />
            </Center>
          </Stack>
        </Box>
      </Center>
    )
  }
  if (isEmpty) {
    return (
      <Center py={16}>
        <Text fontSize='sm' color='ink.3'>
          {emptyMessage ?? 'Nothing to show.'}
        </Text>
      </Center>
    )
  }
  return <>{children}</>
}
