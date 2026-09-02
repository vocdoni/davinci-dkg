import { Link } from 'react-router-dom'
import { useLocation } from 'react-router-dom'
import { ButtonLink, EmptyState, SectionHeader, Stack } from '~kit'
import { paths } from '~routes/paths'

export function NotFoundPage() {
  const { pathname } = useLocation()
  return (
    <Stack>
      <SectionHeader size='page' label='404' title='No such page' description={`Nothing is routed at ${pathname}.`} />
      <EmptyState
        title='Try the overview'
        description='Or search for an epoch id, application id, address or transaction in the bar above.'
        action={
          <Link to={paths.home()}>
            <ButtonLink href={paths.home()} variant='ghost' size='sm'>
              Go to overview
            </ButtonLink>
          </Link>
        }
      />
    </Stack>
  )
}
