import { createContext, useContext, useEffect, useState, type ReactNode } from 'react'
import { Center, Spinner, Text, VStack } from '@chakra-ui/react'
import type { RuntimeConfig } from '~types/index'
import { readRpcOverride } from '~lib/rpc-override'

// Loads /config.json once at boot and caches it for the lifetime of the
// page. /config.json is templated by the nginx entrypoint at container
// start, so it always reflects the deployment's actual chain + manager
// address — no rebuild needed to retarget.

interface ConfigContextValue {
  config: RuntimeConfig
}

const ConfigContext = createContext<ConfigContextValue | null>(null)

export function ConfigProvider({ children }: { children: ReactNode }) {
  const [config, setConfig] = useState<RuntimeConfig | null>(null)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    let cancelled = false
    fetch('/config.json', { cache: 'no-store' })
      .then((res) => {
        if (!res.ok) throw new Error(`/config.json HTTP ${res.status}`)
        return res.json() as Promise<RuntimeConfig>
      })
      .then((cfg) => {
        if (cancelled) return
        // Apply any user-supplied RPC override on top of the baked-in
        // /config.json so power users can repoint the explorer at a private
        // RPC without rebuilding the image.
        const override = readRpcOverride()
        setConfig(override ? { ...cfg, rpcUrl: override } : cfg)
      })
      .catch((err) => {
        if (!cancelled) setError(err instanceof Error ? err.message : String(err))
      })
    return () => {
      cancelled = true
    }
  }, [])

  if (error) {
    return (
      <Center minH='100vh' p={8}>
        <VStack gap={4} maxW='md'>
          <Text fontSize='lg' fontWeight='semibold' color='danger.fg'>
            UI failed to load its runtime config
          </Text>
          <Text fontSize='sm' color='ink.3' textAlign='center'>
            The container could not read <code>/config.json</code>. Verify the entrypoint env vars
            (<code>DAVINCI_DKG_RPC_URL</code>, <code>DAVINCI_DKG_MANAGER_ADDRESS</code>,
            <code>DAVINCI_DKG_CHAIN_ID</code>, <code>DAVINCI_DKG_CHAIN_NAME</code>) are set.
          </Text>
          <Text fontSize='xs' color='ink.4' fontFamily='mono'>
            {error}
          </Text>
        </VStack>
      </Center>
    )
  }

  if (!config) {
    return (
      <Center minH='100vh'>
        <Spinner size='lg' color='cyan.400' />
      </Center>
    )
  }

  return <ConfigContext.Provider value={{ config }}>{children}</ConfigContext.Provider>
}

export function useConfig(): RuntimeConfig {
  const ctx = useContext(ConfigContext)
  if (!ctx) throw new Error('useConfig must be used within <ConfigProvider>')
  return ctx.config
}
