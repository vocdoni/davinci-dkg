import { Box, Heading, HStack, SimpleGrid, Stack, Tabs, Text } from '@chakra-ui/react'
import { useParams } from 'react-router-dom'
import { useRound, useRoundEvents } from '~queries/rounds'
import { useBlockNumber } from '~queries/chain'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { StatusBadge } from '~components/Round/StatusBadge'
import { PhaseTimeline } from '~components/Round/PhaseTimeline'
import { ParticipantList } from '~components/Round/ParticipantList'
import { EventLog } from '~components/Round/EventLog'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { RawJson } from '~components/Debug/RawJson'
import { roundSummary } from '~lib/round-utils'
import { blocksRemaining, blocksToDuration } from '~lib/format'

export function RoundView() {
  const { id } = useParams<{ id: string }>()
  const roundId = (id as `0x${string}`) || undefined

  const round = useRound(roundId)
  const events = useRoundEvents(roundId)
  const { data: blockNumber } = useBlockNumber()

  return (
    <Stack gap={8}>
      <Box>
        <HStack gap={3} mb={2}>
          <Heading size='lg'>Round</Heading>
          <HashCell value={roundId} head={8} tail={6} />
        </HStack>
        <Text color='gray.500' fontSize='sm'>
          Round detail with on-chain status, policy, participants, and activity.
        </Text>
      </Box>

      <QueryDataLayout isLoading={round.isLoading} isError={round.isError} error={round.error}>
        {round.data && (
          <Stack gap={8}>
            {/* Header card: status, summary, timeline */}
            <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={5}>
              <HStack justify='space-between' align='start' wrap='wrap' gap={3}>
                <Stack gap={2}>
                  <HStack gap={3}>
                    <StatusBadge round={round.data.round} />
                    <Text fontSize='2xs' color='gray.500'>
                      nonce {round.data.round.nonce.toString()}
                    </Text>
                  </HStack>
                  <Text fontSize='sm' color='gray.300'>
                    {roundSummary(round.data.round, blockNumber ?? null)}
                  </Text>
                </Stack>
              </HStack>
              <Box mt={5}>
                <PhaseTimeline round={round.data.round} />
              </Box>
            </Box>

            {/* Policy: human-readable, with raw fields disclosed */}
            <Stack gap={3}>
              <Heading size='sm'>Policy</Heading>
              <SimpleGrid columns={{ base: 1, md: 2, lg: 4 }} gap={3}>
                <PolicyKV
                  label='Threshold'
                  value={`${round.data.round.policy.threshold} of ${round.data.round.policy.committeeSize}`}
                  hint={
                    round.data.round.policy.minValidContributions === round.data.round.policy.threshold
                      ? `${round.data.round.policy.threshold} contributions needed to finalize`
                      : `${round.data.round.policy.minValidContributions} contributions needed to finalize (extra redundancy)`
                  }
                />
                <PolicyKV
                  label='Lottery α'
                  value={`${round.data.round.policy.lotteryAlphaBps / 100}%`}
                  hint='candidate-pool size'
                />
                <PolicyKV
                  label='Registration closes'
                  value={
                    blockNumber
                      ? blocksToDuration(
                          blocksRemaining(blockNumber, round.data.round.policy.registrationDeadlineBlock) ?? 0
                        )
                      : '—'
                  }
                  hint={`block #${round.data.round.policy.registrationDeadlineBlock.toString()}`}
                />
                <PolicyKV
                  label='Contribution closes'
                  value={
                    blockNumber
                      ? blocksToDuration(
                          blocksRemaining(blockNumber, round.data.round.policy.contributionDeadlineBlock) ?? 0
                        )
                      : '—'
                  }
                  hint={`block #${round.data.round.policy.contributionDeadlineBlock.toString()}`}
                />
                <PolicyKV
                  label='Finalize unlocks'
                  value={
                    blockNumber
                      ? blocksToDuration(
                          blocksRemaining(blockNumber, round.data.round.policy.finalizeNotBeforeBlock) ?? 0
                        )
                      : '—'
                  }
                  hint={`block #${round.data.round.policy.finalizeNotBeforeBlock.toString()}`}
                />
                <PolicyKV
                  label='Disclosure'
                  value={round.data.round.policy.disclosureAllowed ? 'Allowed' : 'Disabled'}
                  hint='secret-key reveal'
                />
                <PolicyKV
                  label='Organizer'
                  value={<HashCell value={round.data.round.organizer} head={4} tail={4} />}
                />
                <PolicyKV
                  label='Seed'
                  value={
                    BigInt(round.data.round.seed) === 0n
                      ? 'pending'
                      : <HashCell value={round.data.round.seed} head={4} tail={4} />
                  }
                  hint={`block #${round.data.round.seedBlock.toString()}`}
                />
              </SimpleGrid>
              <DetailDisclosure title='Show raw policy and decryption-policy fields'>
                <Stack gap={3}>
                  <Box>
                    <Text fontSize='2xs' color='gray.500' mb={1}>
                      policy
                    </Text>
                    <RawJson value={round.data.round.policy} />
                  </Box>
                  <Box>
                    <Text fontSize='2xs' color='gray.500' mb={1}>
                      decryptionPolicy
                    </Text>
                    <RawJson value={round.data.round.decryptionPolicy} />
                  </Box>
                </Stack>
              </DetailDisclosure>
            </Stack>

            {/* Counters */}
            <Stack gap={3}>
              <Heading size='sm'>Counters</Heading>
              <SimpleGrid columns={{ base: 2, md: 5 }} gap={3}>
                <Counter
                  label='Claimed'
                  value={`${round.data.round.claimedCount}/${round.data.round.policy.committeeSize}`}
                />
                <Counter
                  label='Contributions'
                  value={`${round.data.round.contributionCount}/${round.data.round.policy.committeeSize}`}
                />
                <Counter label='Ciphertexts' value={round.data.round.ciphertextCount.toString()} />
                <Counter label='Partial decryptions' value={round.data.round.partialDecryptionCount.toString()} />
                <Counter label='Revealed shares' value={round.data.round.revealedShareCount.toString()} />
              </SimpleGrid>
            </Stack>

            {/* Tabs: participants + activity */}
            <Tabs.Root defaultValue='participants' variant='line'>
              <Tabs.List>
                <Tabs.Trigger value='participants'>Participants</Tabs.Trigger>
                <Tabs.Trigger value='activity'>Activity</Tabs.Trigger>
              </Tabs.List>
              <Tabs.Content value='participants' pt={4}>
                <ParticipantList participants={round.data.participants} />
              </Tabs.Content>
              <Tabs.Content value='activity' pt={4}>
                <QueryDataLayout
                  isLoading={events.isLoading}
                  isError={events.isError}
                  error={events.error}
                  isEmpty={events.data?.length === 0}
                  emptyMessage='No on-chain events for this round yet.'
                >
                  {events.data && <EventLog events={events.data} />}
                </QueryDataLayout>
              </Tabs.Content>
            </Tabs.Root>
          </Stack>
        )}
      </QueryDataLayout>
    </Stack>
  )
}

function PolicyKV({ label, value, hint }: { label: string; value: React.ReactNode; hint?: string }) {
  return (
    <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={3}>
      <Text fontSize='2xs' color='gray.500' mb={1}>
        {label}
      </Text>
      <Box fontSize='sm' fontWeight='medium' color='gray.100'>
        {value}
      </Box>
      {hint && (
        <Text fontSize='2xs' color='gray.500' mt={1}>
          {hint}
        </Text>
      )}
    </Box>
  )
}

function Counter({ label, value }: { label: string; value: string }) {
  return (
    <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={3}>
      <Text fontSize='2xs' color='gray.500'>
        {label}
      </Text>
      <Text fontSize='lg' fontFamily='mono' fontWeight='semibold'>
        {value}
      </Text>
    </Box>
  )
}
