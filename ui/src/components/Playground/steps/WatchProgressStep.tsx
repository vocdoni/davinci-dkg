import { useEffect, useRef, useState } from 'react'
import { Alert, Box, Button, HStack, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { RoundStatus } from '@vocdoni/davinci-dkg-sdk'
import { LuClipboardList, LuKey, LuLock } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useRound, useRoundEvents } from '~queries/rounds'
import { useBlockNumber } from '~queries/chain'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { StatusBadge } from '~components/Round/StatusBadge'
import { PhaseTimeline } from '~components/Round/PhaseTimeline'
import { Countdown } from '~components/Round/Countdown'
import { HowItWorks } from '../HowItWorks'
import { roundSummary } from '~lib/round-utils'

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

  const canAbort =
    round.data && (round.data.round.status === RoundStatus.Registration || round.data.round.status === RoundStatus.Contribution)

  return (
    <StepCard
      n={3}
      title='Watch the committee build the key'
      status={status}
      description='The committee members claim their slots, then each one publishes a cryptographic contribution. When enough have arrived, the round finalizes.'
    >
      {!roundId ? (
        <Text fontSize='sm' color='gray.500'>
          Create a round above first.
        </Text>
      ) : !round.data ? (
        <Text fontSize='sm' color='gray.500'>
          Loading round…
        </Text>
      ) : (
        <Stack gap={4}>
          <HStack gap={3}>
            <StatusBadge round={round.data.round} />
            <Text fontSize='xs' color='gray.500'>
              nonce {round.data.round.nonce.toString()}
            </Text>
          </HStack>
          <Text fontSize='sm' color='gray.300'>
            {roundSummary(round.data.round, block ?? null)}
          </Text>
          <Box>
            <PhaseTimeline round={round.data.round} />
          </Box>

          {/* Live countdown to whichever deadline is currently relevant. */}
          {round.data.round.status === RoundStatus.Registration && (
            <Countdown
              target={round.data.round.policy.registrationDeadlineBlock}
              label='until registration closes'
            />
          )}
          {round.data.round.status === RoundStatus.Contribution && (
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

          {canAbort && (
            <Box>
              <Button size='xs' colorPalette='red' variant='outline' onClick={onAbort} loading={abortBusy}>
                Abort round
              </Button>
              <Text fontSize='2xs' color='gray.500' mt={1}>
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
                members. No single node ever learns the full key — that's the whole point of a
                <em> distributed </em>key generation. Once enough valid contributions have
                landed, anyone can call <em>finalize</em> and the public key becomes usable.
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
