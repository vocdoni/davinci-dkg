import { useEffect, useState, type ReactNode } from 'react'
import { ConfigContext } from './config-context'
import { isDemoRequested, loadRuntimeConfig, type RuntimeConfig } from './runtime-config'

/**
 * Gates the app on `/config.json`. Nothing that touches a chain mounts until
 * the config resolves, so no hook has to cope with a half-known deployment.
 */
export function ConfigProvider({ children }: { children: ReactNode }) {
  const [config, setConfig] = useState<RuntimeConfig | null>(null)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    let cancelled = false
    loadRuntimeConfig(isDemoRequested())
      .then((cfg) => {
        if (!cancelled) setConfig(cfg)
      })
      .catch((err: unknown) => {
        if (!cancelled) setError(err instanceof Error ? err.message : String(err))
      })
    return () => {
      cancelled = true
    }
  }, [])

  if (error) {
    return (
      <div className='flex min-h-screen items-center justify-center p-8'>
        <div className='max-w-lg rounded-md border border-charcoal bg-carbon p-6'>
          <p className='label-caps text-red'>Config error</p>
          <h1 className='mt-2 text-lg font-semibold text-ghost'>The explorer could not read its runtime config</h1>
          <p className='mt-3 text-sm leading-relaxed text-ash'>
            <code className='text-pewter'>/config.json</code> is missing or malformed. It is templated from{' '}
            <code className='text-pewter'>RPC_URL</code>, <code className='text-pewter'>MANAGER_ADDRESS</code>,{' '}
            <code className='text-pewter'>CHAIN_ID</code>, <code className='text-pewter'>CHAIN_NAME</code>,{' '}
            <code className='text-pewter'>DEPLOY_BLOCK</code> and <code className='text-pewter'>EXPLORER_URL</code> by{' '}
            <code className='text-pewter'>scripts/render-ui-config.sh</code>. Append{' '}
            <code className='text-pewter'>?demo=1</code> to browse the synthetic network instead.
          </p>
          <p className='mt-4 font-mono text-xs text-red'>{error}</p>
        </div>
      </div>
    )
  }

  if (!config) {
    return (
      <div className='flex min-h-screen items-center justify-center'>
        <div className='h-6 w-6 animate-spin rounded-full border-2 border-charcoal border-t-emerald' />
      </div>
    )
  }

  return <ConfigContext.Provider value={config}>{children}</ConfigContext.Provider>
}
