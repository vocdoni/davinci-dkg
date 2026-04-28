import { Box, Grid, GridItem, Heading, HStack, SimpleGrid, Stack, Tabs, Text } from '@chakra-ui/react'
import { useParams } from 'react-router-dom'
import type { ReactNode } from 'react'
import { useRound, useRoundEvents } from '~queries/rounds'
import { useBlockNumber } from '~queries/chain'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { PageHeader } from '~components/Layout/PageHeader'
import { StatusBadge } from '~components/Round/StatusBadge'
import { PhaseTimeline } from '~components/Round/PhaseTimeline'
import { ParticipantList } from '~components/Round/ParticipantList'
import { EventLog } from '~components/Round/EventLog'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { RawJson } from '~components/Debug/RawJson'
import { roundSummary } from '~lib/round-utils'
import { blocksRemaining, blocksToDuration } from '~lib/format'

// Round detail page rendered as a typeset article on a single round.
//
//   Masthead   eyebrow with §-anchor + chain badge, large mono round id,
//              status pill + nonce, plain-English summary paragraph,
//              phase timeline.
//   §1 Policy  marginalia layout: small-caps mono labels in the gutter,
//              human-readable values in the body column. Raw policy JSON
//              hidden behind an appendix disclosure.
//   §2 Counts  five mono counters, each with a left-rule.
//   §3 Tabs    Participants / Activity, plain mono tab triggers underlined
//              with the gold rule when active.
export function RoundView() {
  const { id } = useParams<{ id: string }>()
  const roundId = (id as `0x${string}`) || undefined

  const round = useRound(roundId)
  const events = useRoundEvents(roundId)
  const { data: blockNumber } = useBlockNumber()

  return (
    <Stack gap={{ base: 10, md: 14 }}>
      <PageHeader
        title='Round'
        subtitle={
          roundId && (
            <HStack gap={2} display='inline-flex'>
              <HashCell value={roundId} head={8} tail={6} />
            </HStack>
          )
        }
      />

      <QueryDataLayout isLoading={round.isLoading} isError={round.isError} error={round.error}>
        {round.data && (
          <Stack gap={{ base: 10, md: 14 }}>
            <Box>
              <HStack gap={4} mb={4} wrap='wrap'>
                <StatusBadge round={round.data.round} />
                <Text fontFamily='mono' fontSize='2xs' color='ink.4' letterSpacing='0.06em'>
                  NONCE {round.data.round.nonce.toString()}
                </Text>
              </HStack>
              <Text
                fontSize={{ base: 'md', md: 'lg' }}
                color='ink.1'
                lineHeight='1.55'
                maxW='62ch'
                mb={8}
              >
                {roundSummary(round.data.round, blockNumber ?? null)}
              </Text>
              <Box
                p={{ base: 5, md: 6 }}
                borderWidth='1px'
                borderColor='border.subtle'
                borderRadius='lg'
                bg='surface'
                boxShadow='inset'
              >
                <PhaseTimeline round={round.data.round} />
              </Box>
            </Box>

            <RoundSection title='Round parameters'>
              <SimpleGrid columns={{ base: 1, md: 2 }} gap={{ base: 5, md: 6 }} columnGap={10}>
                <PolicyRow
                  label='Threshold'
                  value={`${round.data.round.policy.threshold} of ${round.data.round.policy.committeeSize}`}
                  hint={
                    round.data.round.policy.minValidContributions === round.data.round.policy.threshold
                      ? `${round.data.round.policy.threshold} contributions needed to finalize`
                      : `${round.data.round.policy.minValidContributions} needed to finalize (extra redundancy)`
                  }
                />
                <PolicyRow
                  label='Lottery α'
                  value={`${round.data.round.policy.lotteryAlphaBps / 100}%`}
                  hint='candidate-pool oversubscription'
                />
                <PolicyRow
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
                <PolicyRow
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
                <PolicyRow
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
                <PolicyRow
                  label='Disclosure'
                  value={round.data.round.policy.disclosureAllowed ? 'Allowed' : 'Disabled'}
                  hint='secret-key reveal'
                />
                <PolicyRow
                  label='Organizer'
                  value={<HashCell value={round.data.round.organizer} head={6} tail={4} />}
                />
                <PolicyRow
                  label='Seed'
                  value={
                    BigInt(round.data.round.seed) === 0n ? (
                      <Text color='ink.3' fontSize='sm'>
                        pending
                      </Text>
                    ) : (
                      <HashCell value={round.data.round.seed} head={6} tail={4} />
                    )
                  }
                  hint={`block #${round.data.round.seedBlock.toString()}`}
                />
              </SimpleGrid>
              <Box mt={6}>
                <DetailDisclosure title='Show raw policy and decryption-policy fields'>
                  <Stack gap={4}>
                    <Box>
                      <Text fontSize='2xs' color='ink.4' mb={2} letterSpacing='0.06em'>
                        POLICY
                      </Text>
                      <RawJson value={round.data.round.policy} />
                    </Box>
                    <Box>
                      <Text fontSize='2xs' color='ink.4' mb={2} letterSpacing='0.06em'>
                        DECRYPTION POLICY
                      </Text>
                      <RawJson value={round.data.round.decryptionPolicy} />
                    </Box>
                  </Stack>
                </DetailDisclosure>
              </Box>
            </RoundSection>

            <RoundSection title='Activity counters'>
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
                <Counter
                  label='Partial decryptions'
                  value={round.data.round.partialDecryptionCount.toString()}
                />
                <Counter label='Revealed shares' value={round.data.round.revealedShareCount.toString()} />
              </SimpleGrid>
            </RoundSection>

            <RoundSection title='Participants & activity'>
              <Tabs.Root defaultValue='participants' variant='line'>
                <Tabs.List borderColor='rule'>
                  <DkgTab value='participants'>Participants</DkgTab>
                  <DkgTab value='activity'>Activity</DkgTab>
                </Tabs.List>
                <Tabs.Content value='participants' pt={6}>
                  <ParticipantList participants={round.data.participants} />
                </Tabs.Content>
                <Tabs.Content value='activity' pt={6}>
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
            </RoundSection>
          </Stack>
        )}
      </QueryDataLayout>
    </Stack>
  )
}

function RoundSection({ title, children }: { title: string; children: ReactNode }) {
  return (
    <Box as='section'>
      <Heading
        as='h2'
        fontSize={{ base: 'lg', md: 'xl' }}
        fontWeight={500}
        color='ink.0'
        letterSpacing='-0.01em'
        mb={5}
      >
        {title}
      </Heading>
      {children}
    </Box>
  )
}

// ─── Policy row (marginalia variant) ─────────────────────────────────────
// Two-column. Label on the left in mono small-caps; value + hint on the
// right. The label sits in the margin of the body, like a textbook's side
// annotation.
function PolicyRow({ label, value, hint }: { label: string; value: ReactNode; hint?: string }) {
  return (
    <Grid templateColumns={{ base: '120px 1fr', md: '140px 1fr' }} gap={4} alignItems='baseline'>
      <GridItem>
        <Text
          fontFamily='mono'
          fontSize='2xs'
          color='ink.3'
          letterSpacing='0.08em'
          textTransform='uppercase'
        >
          {label}
        </Text>
      </GridItem>
      <GridItem>
        <Box fontSize={{ base: 'sm', md: 'md' }} color='ink.0' fontWeight={500}>
          {value}
        </Box>
        {hint && (
          <Text fontSize='xs' color='ink.3' mt={0.5}>
            {hint}
          </Text>
        )}
      </GridItem>
    </Grid>
  )
}

// ─── Counter ────────────────────────────────────────────────────────────
function Counter({ label, value }: { label: string; value: string }) {
  return (
    <Box
      position='relative'
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='lg'
      bg='surface'
      p={{ base: 3, md: 4 }}
      boxShadow='inset'
    >
      <Box
        position='absolute'
        left={0}
        top='25%'
        bottom='25%'
        w='2px'
        bg='accent.fg'
        opacity={0.4}
        borderRightRadius='full'
      />
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
        mb={1.5}
      >
        {label}
      </Text>
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize={{ base: 'lg', md: 'xl' }}
        fontWeight={500}
        color='ink.0'
        lineHeight='1.1'
      >
        {value}
      </Text>
    </Box>
  )
}

// ─── Refined tab trigger ─────────────────────────────────────────────────
function DkgTab({ value, children }: { value: string; children: ReactNode }) {
  return (
    <Tabs.Trigger
      value={value}
      fontFamily='mono'
      fontSize='2xs'
      letterSpacing='0.08em'
      textTransform='uppercase'
      color='ink.3'
      px={3}
      py={2.5}
      _selected={{
        color: 'ink.0',
        borderColor: 'accent.fg',
      }}
      _hover={{ color: 'ink.1' }}
    >
      {children}
    </Tabs.Trigger>
  )
}
