import { Box, Button, Container, Heading, Stack, Text } from '@chakra-ui/react'
import { useRouteError, isRouteErrorResponse, useNavigate } from 'react-router-dom'
import { ErrorReportButton } from '~components/Debug/ErrorReportButton'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { Routes } from '~router/routes'

// React Router error boundary. Distinguishes 404-style "not found" from
// rendering errors so the user gets a relevant message in either case.
export function ErrorElement() {
  const error = useRouteError()
  const navigate = useNavigate()

  let title = 'Something went wrong'
  let summary = 'An unexpected error occurred while rendering this view.'
  if (isRouteErrorResponse(error)) {
    if (error.status === 404) {
      title = 'Page not found'
      summary = "We don't have anything at that URL."
    } else {
      title = `Error ${error.status}`
      summary = error.statusText || summary
    }
  } else if (error instanceof Error) {
    summary = error.message
  }

  return (
    <Container maxW='2xl' py={20}>
      <Stack gap={5}>
        <Heading size='lg' color='red.300'>
          {title}
        </Heading>
        <Text color='gray.300'>{summary}</Text>
        <Box>
          <Button onClick={() => navigate(Routes.home)} colorPalette='cyan' variant='outline'>
            Back to overview
          </Button>
        </Box>
        <ErrorReportButton error={error} />
        <DetailDisclosure>
          <pre>{error instanceof Error ? error.stack : JSON.stringify(error, null, 2)}</pre>
        </DetailDisclosure>
      </Stack>
    </Container>
  )
}
