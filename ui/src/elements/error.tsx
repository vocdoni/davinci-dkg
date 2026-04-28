import { Box, Button, Container, Heading, HStack, Stack, Text } from '@chakra-ui/react'
import { useRouteError, isRouteErrorResponse, useNavigate } from 'react-router-dom'
import { ErrorReportButton } from '~components/Debug/ErrorReportButton'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { Routes } from '~router/routes'

// React Router error boundary. Editorial layout: eyebrow with the HTTP-ish
// status, display-serif headline, italic body summary, then actions.
export function ErrorElement() {
  const error = useRouteError()
  const navigate = useNavigate()

  let eyebrow = 'Error'
  let title = 'Something went wrong'
  let summary = 'An unexpected error occurred while rendering this view.'
  if (isRouteErrorResponse(error)) {
    if (error.status === 404) {
      eyebrow = '404 · Not found'
      title = 'Page not found'
      summary = "We don't have anything at that URL."
    } else {
      eyebrow = `${error.status} · ${error.statusText || 'Error'}`
      title = error.statusText || title
      summary = error.statusText || summary
    }
  } else if (error instanceof Error) {
    summary = error.message
  }

  return (
    <Container maxW='2xl' py={{ base: 16, md: 24 }}>
      <Stack gap={5}>
        <HStack
          fontFamily='mono'
          fontSize='2xs'
          color='danger.fg'
          letterSpacing='0.08em'
          gap={2}
        >
          <Box w='6px' h='6px' borderRadius='full' bg='danger.fg' />
          <Text textTransform='uppercase'>{eyebrow}</Text>
        </HStack>
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
        <Text fontSize='md' color='ink.2' lineHeight='1.55'>
          {summary}
        </Text>
        <HStack gap={3} pt={2}>
          <Button
            onClick={() => navigate(Routes.home)}
            bg='accent.fg'
            color='canvas.deep'
            fontFamily='sans'
            fontWeight={500}
            _hover={{ bg: 'accent.bright' }}
          >
            Back to overview
          </Button>
          <ErrorReportButton error={error} />
        </HStack>
        <DetailDisclosure title='Show stack trace'>
          <Box as='pre' whiteSpace='pre-wrap'>
            {error instanceof Error ? error.stack : JSON.stringify(error, null, 2)}
          </Box>
        </DetailDisclosure>
      </Stack>
    </Container>
  )
}
