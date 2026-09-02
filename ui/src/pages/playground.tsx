import { useMemo } from 'react'
import { useSearchParams } from 'react-router-dom'
import type { Hex } from 'viem'
import { useRuntimeConfig } from '~config/config-context'
import { Stepper } from './playground/Stepper'
import { useEpochStepData } from './playground/epochs'
import { useDemoChain } from './playground/useDemoChain'
import { useLiveChain } from './playground/useLiveChain'
import type { PlaygroundTarget } from './playground/types'

/**
 * `/playground` — the organizer walkthrough.
 *
 * The demo and live variants are separate components on purpose: only
 * `LivePlayground` mounts the wagmi hooks, so `?demo=1` never asks the browser
 * for a wallet, never opens an RPC socket, and can be screenshotted (and
 * rendered in a unit test) with nothing but the synthetic fixture behind it.
 *
 * The `(epoch, aid)` pair lives in the URL rather than in component state, so
 * the walkthrough is a link: `/playground?epoch=…&aid=…` reopens it at the
 * right step whenever this tab still holds the organizer secret.
 */
export function PlaygroundPage() {
  const { demo } = useRuntimeConfig()
  const [params] = useSearchParams()
  const target = useMemo<PlaygroundTarget>(
    () => ({
      epochId: (params.get('epoch') as Hex | null) ?? null,
      aid: (params.get('aid') as Hex | null) ?? null,
    }),
    [params]
  )
  return demo ? <DemoPlayground target={target} /> : <LivePlayground target={target} />
}

function DemoPlayground({ target }: { target: PlaygroundTarget }) {
  const { data, options } = useEpochStepData()
  const chain = useDemoChain(target)
  return <Stepper chain={chain} epochs={data} options={options} />
}

function LivePlayground({ target }: { target: PlaygroundTarget }) {
  const { data, options } = useEpochStepData()
  const chain = useLiveChain(target)
  return <Stepper chain={chain} epochs={data} options={options} />
}
