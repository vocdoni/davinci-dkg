import { Suspense, type ReactNode } from 'react'
import { Center, Spinner } from '@chakra-ui/react'

export function SuspenseLoader({ children }: { children: ReactNode }) {
  return (
    <Suspense
      fallback={
        <Center minH='40vh'>
          <Spinner size='lg' color='cyan.400' />
        </Center>
      }
    >
      {children}
    </Suspense>
  )
}
