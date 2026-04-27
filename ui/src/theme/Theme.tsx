import { ChakraProvider } from '@chakra-ui/react'
import type { ReactNode } from 'react'
import { ColorModeProvider } from './color-mode'
import { system } from './system'

export function Theme({ children }: { children: ReactNode }) {
  return (
    <ColorModeProvider>
      <ChakraProvider value={system}>{children}</ChakraProvider>
    </ColorModeProvider>
  )
}
