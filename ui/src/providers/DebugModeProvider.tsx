import { createContext, useCallback, useContext, useEffect, useState, type ReactNode } from 'react'
import { readDebugMode, writeDebugMode } from '~lib/debug'

interface DebugModeContextValue {
  enabled: boolean
  setEnabled: (v: boolean) => void
  toggle: () => void
}

const DebugModeContext = createContext<DebugModeContextValue | null>(null)

export function DebugModeProvider({ children }: { children: ReactNode }) {
  const [enabled, setEnabledState] = useState<boolean>(() => readDebugMode())

  const setEnabled = useCallback((v: boolean) => {
    setEnabledState(v)
    writeDebugMode(v)
  }, [])

  const toggle = useCallback(() => {
    setEnabledState((prev) => {
      writeDebugMode(!prev)
      return !prev
    })
  }, [])

  // Cross-tab synchronisation: if the user toggles debug mode in one tab,
  // mirror it into the other tabs of the same origin so they don't show
  // stale state.
  useEffect(() => {
    const onStorage = (e: StorageEvent) => {
      if (e.key !== 'dkg-ui:debug') return
      setEnabledState(e.newValue === '1')
    }
    window.addEventListener('storage', onStorage)
    return () => window.removeEventListener('storage', onStorage)
  }, [])

  return <DebugModeContext.Provider value={{ enabled, setEnabled, toggle }}>{children}</DebugModeContext.Provider>
}

export function useDebugMode(): DebugModeContextValue {
  const ctx = useContext(DebugModeContext)
  if (!ctx) throw new Error('useDebugMode must be used within <DebugModeProvider>')
  return ctx
}
