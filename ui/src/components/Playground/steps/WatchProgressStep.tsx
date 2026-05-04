import { useEffect, useRef, useState } from 'react'
import { Alert, Box, Button, HStack, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { EpochPhase, type Epoch } from '@vocdoni/davinci-dkg-sdk'
import { LuClipboardList, LuKey, LuLock } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useEpoch, useEpochEvents } from '~queries/epochs'
import { useBlockNumber } from '~queries/chain'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { StatusBadge } from '~components/Epoch/StatusBadge'
import { PhaseTimeline } from '~components/Epoch/PhaseTimeline'
import { Countdown } from '~components/Epoch/Countdown'
import { HowItWorks } from '../HowItWorks'
import { roundFailure, roundSummary } from '~lib/epoch-utils'

interface Props {
  status: StepStatus
  epochId: Hex | null
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx' | 'chain') => void
}

export function WatchProgressStep({ status, epochId, log }: Props) {
  const writer = useDkgWriter()
  const epoch = useEpoch((epochId ?? undefined) as `0x${string}` | undefined)
  const events = useEpochEvents((epochId ?? undefined) as `0x${string}` | undefined)
  const { data: block } = useBlockNumber()

  const [abortBusy, setAbortBusy] = useState(false)

  // Mirror epoch-status / event-count transitions into the activity log.
  // We track the last-seen values in refs so the effect doesn't re-fire on
  // every render — only when the underlying value actually changed.
  const lastStatus = useRef<number | null>(null)
  const lastEventCount = useRef(0)
  useEffect(() => {
    if (!epoch.data) return
    const s = Number(epoch.data.epoch.status)
    if (lastStatus.current !== s) {
      lastStatus.current = s
      const labels = ['None', 'Registration', 'Contribution', 'Finalized', 'Aborted', 'Completed']
      log(`Epoch status → ${labels[s] ?? s}`, s === 3 ? 'success' : s === 4 ? 'error' : 'chain')
    }
  }, [epoch.data, log])
  useEffect(() => {
    if (!events.data) return
    if (events.data.length > lastEventCount.current) {
      for (let i = lastEventCount.current; i < events.data.length; i++) {
        const ev = events.data[i]
        log(`Event: ${ev.eventName} @ block #${ev.blockNumber}`, 'chain')
      }
      lastEventCount.current = events.data.length
    }
  }, [events.data, log])

  const onAbort = async () => {
    if (!writer || !epochId) return
    setAbortBusy(true)
    try {
      log('Aborting epoch…', 'error')
      const hash = await writer.abortEpoch(epochId)
      log(`Abort tx submitted: ${hash}`, 'tx')
      await writer.waitForTransaction(hash)
      log('Epoch aborted.', 'error')
    } catch (err) {
      log(`Abort failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setAbortBusy(false)
    }
  }

  const failure = epoch.data ? roundFailure(epoch.data.epoch, block ?? null) : null
  const canAbort =
    epoch.data &&
    (epoch.data.epoch.status === EpochPhase.CommitteeSelection ||
      epoch.data.epoch.status === EpochPhase.KeyAssembly)

  // Pick the headline counter for the current phase. After finalize the
  // counters are no longer meaningful here — KeyAvailableStep takes over.
  const headlineCounter = epoch.data ? pickHeadlineCounter(epoch.data.epoch) : null

  const cardStatus: StepStatus = failure ? 'error' : status

  return (
    <StepCard
      n={3}
      title='Watch the committee build the key'
      status={cardStatus}
      description='The committee members claim their slots, then each one publishes a cryptographic contribution. When enough have arrived, the epoch finalizes.'
    >
      {!epochId ? (
        <Text fontSize='sm' color='ink.4'>
          Create a epoch above first.
        </Text>
      ) : !epoch.data ? (
        <Text fontSize='sm' color='ink.4'>
          Loading epoch…
        </Text>
      ) : (
        <Stack gap={5}>
          <HStack gap={3} wrap='wrap'>
            <StatusBadge epoch={epoch.data.epoch} />
            <Text fontSize='xs' color='ink.4'>
              nonce {epoch.data.epoch.nonce.toString()}
            </Text>
          </HStack>
          <Text fontSize='sm' color='ink.2'>
            {roundSummary(epoch.data.epoch, block ?? null)}
          </Text>
          <Box>
            <PhaseTimeline epoch={epoch.data.epoch} />
          </Box>

          {/* Big live counter for the active phase. */}
          {headlineCounter && !failure && (
            <BigCounter
              label={headlineCounter.label}
              have={headlineCounter.have}
              total={headlineCounter.total}
              need={headlineCounter.need}
              tone={headlineCounter.tone}
              caption={headlineCounter.caption}
            />
          )}

          {/* Live countdown to whichever deadline is currently relevant. */}
          {epoch.data.epoch.status === EpochPhase.CommitteeSelection && !failure && (
            <Countdown
              target={epoch.data.epoch.policy.committeeSelectionDeadlineBlock}
              label='until registration closes'
            />
          )}
          {epoch.data.epoch.status === EpochPhase.KeyAssembly && !failure && (
            <Stack gap={1}>
              <Countdown
                target={epoch.data.epoch.policy.keyAssemblyDeadlineBlock}
                label='until contributions close'
              />
              <Countdown
                target={epoch.data.epoch.policy.liveNotBeforeBlock}
                label='until finalize unlocks'
              />
            </Stack>
          )}

          {/* Failure banner — epoch window expired without enough nodes. */}
          {failure && (
            <Alert.Root status='error'>
              <Alert.Indicator />
              <Alert.Content>
                <Alert.Title>
                  {failure.kind === 'committee-selection'
                    ? 'Registration closed without a viable committee.'
                    : 'Contribution window closed without enough contributions.'}
                </Alert.Title>
                <Alert.Description fontSize='xs'>
                  {failure.kind === 'committee-selection' ? (
                    <>
                      Only <b>{failure.have}</b> of the {failure.total} committee slot(s) were
                      claimed before the deadline — at least <b>{failure.need}</b> are needed for
                      the epoch to be decryptable. The playground cannot continue with this epoch.
                      Abort it and try again with a longer window, or wait for more nodes to come
                      online.
                    </>
                  ) : (
                    <>
                      Only <b>{failure.have}</b> contribution(s) arrived in time —{' '}
                      <b>{failure.need}</b> are required to finalize. The playground cannot
                      continue. Abort and create a new epoch.
                    </>
                  )}
                </Alert.Description>
              </Alert.Content>
            </Alert.Root>
          )}

          {canAbort && (
            <Box>
              <Button size='xs' colorPalette='red' variant='outline' onClick={onAbort} loading={abortBusy}>
                Abort epoch
              </Button>
              <Text fontSize='2xs' color='ink.4' mt={1}>
                Organizer-only. Useful if you started a epoch you don't intend to complete.
              </Text>
            </Box>
          )}
          {epoch.data.epoch.status === EpochPhase.Aborted && (
            <Alert.Root status='error'>
              <Alert.Indicator />
              <Alert.Title>This epoch was aborted.</Alert.Title>
            </Alert.Root>
          )}

          <HowItWorks
            body={
              <>
                Each committee node generates a small piece of the eventual private key, posts a
                public commitment to its piece, and shares the piece secretly with the other
                members. No single node ever learns the full key — that's the whole point of a{' '}
                <em>distributed</em> key generation. Once enough valid contributions have landed,
                anyone can call <em>finalize</em> and the public key becomes usable.
              </>
            }
            flow={[
              { icon: <LuClipboardList />, label: 'Members claim slots' },
              { icon: <LuKey />, label: 'Each posts a key piece' },
              { icon: <LuLock />, label: 'Public key locked in' },
            ]}
          />
        </Stack>
      )}
    </StepCard>
  )
}

interface CounterSpec {
  label: string
  have: number
  total: number
  need: number
  tone: 'accent' | 'live'
  caption: string
}

function pickHeadlineCounter(epoch: Epoch): CounterSpec | null {
  const min = epoch.policy.minValidContributions
  const n = epoch.policy.committeeSize
  switch (epoch.status) {
    case EpochPhase.CommitteeSelection:
      return {
        label: 'Slots claimed',
        have: epoch.claimedCount,
        total: n,
        need: Math.min(min, n),
        tone: epoch.claimedCount >= n ? 'live' : 'accent',
        caption:
          epoch.claimedCount >= n
            ? 'Committee full — moving to contribution phase.'
            : `${epoch.claimedCount} of ${n} eligible nodes have joined this committee.`,
      }
    case EpochPhase.KeyAssembly:
      return {
        label: 'Contributions accepted',
        have: epoch.contributionCount,
        total: n,
        need: min,
        tone: epoch.contributionCount >= min ? 'live' : 'accent',
        caption:
          epoch.contributionCount >= min
            ? `Threshold reached. Awaiting finalize at block ${epoch.policy.liveNotBeforeBlock.toString()}.`
            : `${epoch.contributionCount} of ${min} required contributions received.`,
      }
    default:
      return null
  }
}

interface BigCounterProps {
  label: string
  have: number
  total: number
  need: number
  tone: 'accent' | 'live'
  caption: string
}

// Bold "X / N" status panel with progress bar and threshold marker.
// Reads as the centrepiece of the active phase, replacing the previous
// quiet inline summary.
function BigCounter({ label, have, total, need, tone, caption }: BigCounterProps) {
  const pct = total > 0 ? Math.min(100, (have / total) * 100) : 0
  const needPct = total > 0 ? Math.min(100, (need / total) * 100) : 0
  const barFill = tone === 'live' ? 'live.fg' : 'accent.fg'
  const valueColor = tone === 'live' ? 'live.fg' : 'ink.0'

  return (
    <Box
      borderWidth='1px'
      borderColor={tone === 'live' ? 'rgba(134, 239, 172, 0.40)' : 'border.subtle'}
      bg={tone === 'live' ? 'live.bg' : 'surface.sunken'}
      borderRadius='lg'
      p={{ base: 4, md: 5 }}
    >
      <Stack gap={3}>
        <HStack justify='space-between' align='baseline' wrap='wrap' gap={2}>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            {label}
          </Text>
          <Text fontFamily='mono' fontSize='2xs' color='ink.4' letterSpacing='0.06em'>
            need {need} / {total}
          </Text>
        </HStack>
        <HStack align='baseline' gap={2}>
          <Text
            className='dkg-tabular'
            fontFamily='mono'
            fontSize={{ base: '4xl', md: '5xl' }}
            fontWeight={600}
            color={valueColor}
            lineHeight='1'
            letterSpacing='-0.02em'
          >
            {have}
          </Text>
          <Text
            className='dkg-tabular'
            fontFamily='mono'
            fontSize={{ base: 'xl', md: '2xl' }}
            color='ink.3'
            fontWeight={400}
            lineHeight='1'
          >
            / {total}
          </Text>
        </HStack>
        <Box position='relative' h='8px' bg='surface.raised' borderRadius='full' overflow='visible'>
          <Box
            position='absolute'
            inset={0}
            w={`${pct}%`}
            bg={barFill}
            borderRadius='full'
            transition='width 0.3s ease'
          />
          {/* Threshold marker — vertical hairline at the "need" position. */}
          {need > 0 && need < total && (
            <Box
              position='absolute'
              top='-3px'
              bottom='-3px'
              left={`${needPct}%`}
              w='2px'
              bg='ink.1'
              opacity={0.55}
              transform='translateX(-1px)'
              borderRadius='full'
            />
          )}
        </Box>
        <Text fontSize='xs' color='ink.3' lineHeight='1.5'>
          {caption}
        </Text>
      </Stack>
    </Box>
  )
}

// Re-export so the playground page can also display the partial-decryption
// counter on VerifyDecryptionStep without re-implementing it.
export { BigCounter }
export type { BigCounterProps }
