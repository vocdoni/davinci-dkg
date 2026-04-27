import { ThemeProvider as NextThemesProvider, useTheme } from 'next-themes'
import type { ReactNode } from 'react'

// Single re-export point for color-mode handling. Every component that wants
// to know "are we in dark or light?" goes through useTheme() from here, so
// swapping the underlying provider later is a one-file change.

export function ColorModeProvider({ children }: { children: ReactNode }) {
  return (
    <NextThemesProvider attribute='class' defaultTheme='dark' enableSystem disableTransitionOnChange>
      {children}
    </NextThemesProvider>
  )
}

export { useTheme }
