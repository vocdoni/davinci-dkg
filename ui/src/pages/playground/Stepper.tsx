import { Badge, Card, Panel, SectionHeader, Stack, Toggle } from '~kit'
import { stepStatus, type StepId } from './machine'
import { usePlaygroundController, type EpochOption } from './controller'
import { ActivityLog, StepRail } from './parts'
import {
  ConnectStep,
  EncryptStep,
  EpochStep,
  RegisterStep,
  ShareStep,
  SubmitStep,
  VerifyStep,
  WatchStep,
  type EpochStepData,
  type StepProps,
} from './steps'
import type { PlaygroundChain } from './types'

const PANELS: Record<StepId, (props: StepProps) => JSX.Element> = {
  connect: ConnectStep,
  epoch: EpochStep,
  register: RegisterStep,
  encrypt: EncryptStep,
  submit: SubmitStep,
  share: ShareStep,
  watch: WatchStep,
  verify: VerifyStep,
}

export interface StepperProps {
  chain: PlaygroundChain
  epochs: EpochStepData
  /** Same list, in the shape the controller resolves the pinned epoch from. */
  options: EpochOption[]
}

/**
 * The playground proper: a rail of eight steps, one panel at a time, and a log
 * of everything that happened. Both modes render this — the only difference is
 * which `PlaygroundChain` is behind it.
 */
export function Stepper({ chain, epochs, options }: StepperProps) {
  const controller = usePlaygroundController(chain, options)
  const Panel_ = PANELS[controller.step]
  const { state } = controller

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Organizer'
        title='Playground'
        description='Register an application against a live epoch, encrypt a value under its key, submit it, decide when the plaintext may exist, and check the result against what this browser built.'
        actions={
          <div className='flex items-center gap-4'>
            {chain.kind === 'demo' ? (
              <Badge tone='warn' dot title='Running off the synthetic fixture; nothing is broadcast'>
                demo
              </Badge>
            ) : (
              <Badge tone='ok' dot>
                live
              </Badge>
            )}
            <Toggle
              checked={state.advanced}
              onChange={controller.actions.toggleAdvanced}
              label='Advanced'
              hint='print transcripts'
            />
          </div>
        }
      />

      <div className='grid gap-6 lg:grid-cols-[260px_minmax(0,1fr)]'>
        <div className='flex flex-col gap-4'>
          <Card flush className='p-3'>
            <StepRail
              active={controller.step}
              status={(step) => stepStatus(step, state, controller.facts)}
              onSelect={controller.actions.goto}
            />
          </Card>
          {controller.epochId ? (
            <Card level='onyx' className='text-[11px] leading-relaxed text-ash'>
              <div className='label-caps mb-1.5 text-[10px] text-pewter'>working on</div>
              <div className='truncate font-mono text-[11px] text-silver' title={controller.epochId}>
                {controller.epochId}
              </div>
              {state.aid ? (
                <div className='mt-1 truncate font-mono text-[11px] text-silver' title={state.aid}>
                  {state.aid}
                </div>
              ) : null}
              {state.pinned ? <div className='mt-2 text-emerald'>epoch pinned by the registration</div> : null}
            </Card>
          ) : null}
        </div>

        <div className='flex min-w-0 flex-col gap-6'>
          <Panel_ controller={controller} chain={chain} epochs={epochs} />
          <Panel title='Activity log' description='Every action this walkthrough took, newest first.'>
            <ActivityLog entries={state.log} />
          </Panel>
        </div>
      </div>
    </Stack>
  )
}
