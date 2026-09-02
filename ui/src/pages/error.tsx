import { isRouteErrorResponse, useRouteError } from 'react-router-dom'
import { Callout, PageContainer, SectionHeader, Stack } from '~kit'

/** Route-level error boundary. Renders outside the shell, so it stays plain. */
export function RouteError() {
  const error = useRouteError()
  const message = isRouteErrorResponse(error)
    ? `${error.status} ${error.statusText}`
    : error instanceof Error
      ? error.message
      : 'Unknown error'
  const stack = error instanceof Error ? error.stack : undefined

  return (
    <PageContainer className='py-16'>
      <Stack>
        <SectionHeader size='page' label='Error' title='Something broke while rendering this page' />
        <Callout tone='danger' title={message}>
          {stack ? (
            <pre className='mt-2 max-h-64 overflow-auto rounded-sm border border-charcoal bg-obsidian p-3 text-[11px] leading-relaxed scroll-slim'>
              {stack}
            </pre>
          ) : null}
        </Callout>
      </Stack>
    </PageContainer>
  )
}
