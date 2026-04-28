import { useEffect, useRef, useState } from 'react'
import { Alert, Box, Button, HStack, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { RoundStatus, type Round } from '@vocdoni/davinci-dkg-sdk'
import { LuClipboardList, LuKey, LuLock } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useRound, useRoundEvents } from '~queries/rounds'
import { useBlockNumber } from '~queries/chain'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { StatusBadge } from '~components/Round/StatusBadge'
import { PhaseTimeline } from '~components/Round/PhaseTimeline'
import { Countdown } from '~components/Round/Countdown'
import { HowItWorks } from '../HowItWorks'
import { roundFailure, roundSummary } from '~lib/round-utils'

interface Props {
  status: StepStatus
  roundId: Hex | null
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx' | 'chain') => void
}

export function WatchProgressStep({ status, roundId, log }: Props) {
  const writer = useDkgWriter()
  const round = useRound((roundId ?? undefined) as `0x${string}` | undefined)
  const events = useRoundEvents((roundId ?? undefined) as `0x${string}` | undefined)
  const { data: block } = useBlockNumber()

  const [abortBusy, setAbortBusy] = useState(false)

  // Mirror round-status / event-count transitions into the activity log.
  // We track the last-seen values in refs so the effect doesn't re-fire on
  // every render — only when the underlying value actually changed.
  const lastStatus = useRef<number | null>(null)
  const lastEventCount = useRef(0)
  useEffect(() => {
    if (!round.data) return
    const s = Number(round.data.round.status)
    if (lastStatus.current !== s) {
      lastStatus.current = s
      const labels = ['None', 'Registration', 'Contribution', 'Finalized', 'Aborted', 'Completed']
      log(`Round status → ${labels[s] ?? s}`, s === 3 ? 'success' : s === 4 ? 'error' : 'chain')
    }
  }, [round.data, log])
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
    if (!writer || !roundId) return
    setAbortBusy(true)
    try {
      log('Aborting round…', 'error')
      const hash = await writer.abortRound(roundId)
      log(`Abort tx submitted: ${hash}`, 'tx')
      await writer.waitForTransaction(hash)
      log('Round aborted.', 'error')
    } catch (err) {
      log(`Abort failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setAbortBusy(false)
    }
  }

  const failure = round.data ? roundFailure(round.data.round, block ?? null) : null
  const canAbort =
    round.data &&
    (round.data.round.status === RoundStatus.Registration ||
      round.data.round.status === RoundStatus.Contribution)

  // Pick the headline counter for the current phase. After finalize the
  // counters are no longer meaningful here — KeyAvailableStep takes over.
  const headlineCounter = round.data ? pickHeadlineCounter(round.data.round) : null

  const cardStatus: StepStatus = failure ? 'error' : status

  return (
    <StepCard
      n={3}
      title='Watch the committee build the key'
      status={cardStatus}
      description='The committee members claim their slots, then each one publishes a cryptographic contribution. When enough have arrived, the round finalizes.'
    >
      {!roundId ? (
        <Text fontSize='sm' color='ink.4'>
          Create a round above first.
        </Text>
      ) : !round.data ? (
        <Text fontSize='sm' color='ink.4'>
          Loading round…
        </Text>
      ) : (
        <Stack gap={5}>
          <HStack gap={3} wrap='wrap'>
            <StatusBadge round={round.data.round} />
            <Text fontSize='xs' color='ink.4'>
              nonce {round.data.round.nonce.toString()}
            </Text>
          </HStack>
          <Text fontSize='sm' color='ink.2'>
            {roundSummary(round.data.round, block ?? null)}
          </Text>
          <Box>
            <PhaseTimeline round={round.data.round} />
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
          {round.data.round.status === RoundStatus.Registration && !failure && (
            <Countdown
              target={round.data.round.policy.registrationDeadlineBlock}
              label='until registration closes'
            />
          )}
          {round.data.round.status === RoundStatus.Contribution && !failure && (
            <Stack gap={1}>
              <Countdown
                target={round.data.round.policy.contributionDeadlineBlock}
                label='until contributions close'
              />
              <Countdown
                target={round.data.round.policy.finalizeNotBeforeBlock}
                label='until finalize unlocks'
              />
            </Stack>
          )}

          {/* Failure banner — round window expired without enough nodes. */}
          {failure && (
            <Alert.Root status='error'>
              <Alert.Indicator />
              <Alert.Content>
                <Alert.Title>
                  {failure.kind === 'registration'
                    ? 'Registration closed without a viable committee.'
                    : 'Contribution window closed without enough contributions.'}
                </Alert.Title>
                <Alert.Description fontSize='xs'>
                  {failure.kind === 'registration' ? (
                    <>
                      Only <b>{failure.have}</b> of the {failure.total} committee slot(s) were
                      claimed before the deadline — at least <b>{failure.need}</b> are needed for
                      the round to be decryptable. The playground cannot continue with this round.
                      Abort it and try again with a longer window, or wait for more nodes to come
                      online.
                    </>
                  ) : (
                    <>
                      Only <b>{failure.have}</b> contribution(s) arrived in time —{' '}
                      <b>{failure.need}</b> are required to finalize. The playground cannot
                      continue. Abort and create a new round.
                    </>
                  )}
                </Alert.Description>
              </Alert.Content>
            </Alert.Root>
          )}

          {canAbort && (
            <Box>
              <Button size='xs' colorPalette='red' variant='outline' onClick={onAbort} loading={abortBusy}>
                Abort round
              </Button>
              <Text fontSize='2xs' color='ink.4' mt={1}>
                Organizer-only. Useful if you started a round you don't intend to complete.
              </Text>
            </Box>
          )}
          {round.data.round.status === RoundStatus.Aborted && (
            <Alert.Root status='error'>
              <Alert.Indicator />
              <Alert.Title>This round was aborted.</Alert.Title>
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

function pickHeadlineCounter(round: Round): CounterSpec | null {
  const min = round.policy.minValidContributions
  const n = round.policy.committeeSize
  switch (round.status) {
    case RoundStatus.Registration:
      return {
        label: 'Slots claimed',
        have: round.claimedCount,
        total: n,
        need: Math.min(min, n),
        tone: round.claimedCount >= n ? 'live' : 'accent',
        caption:
          round.claimedCount >= n
            ? 'Committee full — moving to contribution phase.'
            : `${round.claimedCount} of ${n} eligible nodes have joined this committee.`,
      }
    case RoundStatus.Contribution:
      return {
        label: 'Contributions accepted',
        have: round.contributionCount,
        total: n,
        need: min,
        tone: round.contributionCount >= min ? 'live' : 'accent',
        caption:
          round.contributionCount >= min
            ? `Threshold reached. Awaiting finalize at block ${round.policy.finalizeNotBeforeBlock.toString()}.`
            : `${round.contributionCount} of ${min} required contributions received.`,
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
