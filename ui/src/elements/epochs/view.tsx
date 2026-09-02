import { useState } from 'react'
import { Box, Grid, GridItem, Heading, HStack, Input, SimpleGrid, Stack, Tabs, Text } from '@chakra-ui/react'
import { useParams } from 'react-router-dom'
import type { ReactNode } from 'react'
import type { Hex } from 'viem'
import { useEpoch, useEpochEvents } from '~queries/epochs'
import { useBlockNumber } from '~queries/chain'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { PageHeader } from '~components/Layout/PageHeader'
import { StatusBadge } from '~components/Epoch/StatusBadge'
import { PhaseTimeline } from '~components/Epoch/PhaseTimeline'
import { ParticipantList } from '~components/Epoch/ParticipantList'
import { EventLog } from '~components/Epoch/EventLog'
import { AppRegistrationForm } from '~components/Epoch/AppRegistrationForm'
import { DecryptionPipeline } from '~components/Epoch/DecryptionPipeline'
import { useCollectivePublicKey } from '~queries/applications'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { RawJson } from '~components/Debug/RawJson'
import { roundSummary } from '~lib/epoch-utils'
import { blocksRemaining, blocksToDuration } from '~lib/format'

// Epoch detail page rendered as a typeset article on a single epoch.
//
//   Masthead   eyebrow with §-anchor + chain badge, large mono epoch id,
//              status pill + nonce, plain-English summary paragraph,
//              phase timeline.
//   §1 Policy  marginalia layout: small-caps mono labels in the gutter,
//              human-readable values in the body column. Raw policy JSON
//              hidden behind an appendix disclosure.
//   §2 Counts  five mono counters, each with a left-rule.
//   §3 Tabs    Participants / Activity, plain mono tab triggers underlined
//              with the gold rule when active.
export function EpochView() {
  const { id } = useParams<{ id: string }>()
  const epochId = (id as `0x${string}`) || undefined

  const epoch = useEpoch(epochId)
  const events = useEpochEvents(epochId)
  const { data: blockNumber } = useBlockNumber()

  return (
    <Stack gap={{ base: 10, md: 14 }}>
      <PageHeader
        title='Epoch'
        subtitle={
          epochId && (
            <HStack gap={2} display='inline-flex'>
              <HashCell value={epochId} head={8} tail={6} />
            </HStack>
          )
        }
      />

      <QueryDataLayout isLoading={epoch.isLoading} isError={epoch.isError} error={epoch.error}>
        {epoch.data && (
          <Stack gap={{ base: 10, md: 14 }}>
            <Box>
              <HStack gap={4} mb={4} wrap='wrap'>
                <StatusBadge epoch={epoch.data.epoch} />
                <Text fontFamily='mono' fontSize='2xs' color='ink.4' letterSpacing='0.06em'>
                  NONCE {epoch.data.epoch.nonce.toString()}
                </Text>
              </HStack>
              <Text
                fontSize={{ base: 'md', md: 'lg' }}
                color='ink.1'
                lineHeight='1.55'
                maxW='62ch'
                mb={8}
              >
                {roundSummary(epoch.data.epoch, blockNumber ?? null)}
              </Text>
              <Box
                p={{ base: 5, md: 6 }}
                borderWidth='1px'
                borderColor='border.subtle'
                borderRadius='lg'
                bg='surface'
                boxShadow='inset'
              >
                <PhaseTimeline epoch={epoch.data.epoch} />
              </Box>
            </Box>

            <RoundSection title='Epoch parameters'>
              <SimpleGrid columns={{ base: 1, md: 2 }} gap={{ base: 5, md: 6 }} columnGap={10}>
                <PolicyRow
                  label='Threshold'
                  value={`${epoch.data.epoch.policy.threshold} of ${epoch.data.epoch.policy.committeeSize}`}
                  hint={
                    epoch.data.epoch.policy.minValidContributions === epoch.data.epoch.policy.threshold
                      ? `${epoch.data.epoch.policy.threshold} contributions needed to finalize`
                      : `${epoch.data.epoch.policy.minValidContributions} needed to finalize (extra redundancy)`
                  }
                />
                <PolicyRow
                  label='Lottery α'
                  value={`${epoch.data.epoch.policy.lotteryAlphaBps / 100}%`}
                  hint='candidate-pool oversubscription'
                />
                <PolicyRow
                  label='Committee Selection closes'
                  value={
                    blockNumber
                      ? blocksToDuration(
                          blocksRemaining(blockNumber, epoch.data.epoch.policy.committeeSelectionDeadlineBlock) ?? 0
                        )
                      : '—'
                  }
                  hint={`block #${epoch.data.epoch.policy.committeeSelectionDeadlineBlock.toString()}`}
                />
                <PolicyRow
                  label='Key Assembly closes'
                  value={
                    blockNumber
                      ? blocksToDuration(
                          blocksRemaining(blockNumber, epoch.data.epoch.policy.keyAssemblyDeadlineBlock) ?? 0
                        )
                      : '—'
                  }
                  hint={`block #${epoch.data.epoch.policy.keyAssemblyDeadlineBlock.toString()}`}
                />
                <PolicyRow
                  label='Goes Live at'
                  value={
                    blockNumber
                      ? blocksToDuration(
                          blocksRemaining(blockNumber, epoch.data.epoch.policy.liveNotBeforeBlock) ?? 0
                        )
                      : '—'
                  }
                  hint={`block #${epoch.data.epoch.policy.liveNotBeforeBlock.toString()}`}
                />
                <PolicyRow
                  label='Organizer'
                  value={<HashCell value={epoch.data.epoch.organizer} head={6} tail={4} />}
                />
                <PolicyRow
                  label='Seed'
                  value={
                    BigInt(epoch.data.epoch.seed) === 0n ? (
                      <Text color='ink.3' fontSize='sm'>
                        pending
                      </Text>
                    ) : (
                      <HashCell value={epoch.data.epoch.seed} head={6} tail={4} />
                    )
                  }
                  hint={`block #${epoch.data.epoch.seedBlock.toString()}`}
                />
              </SimpleGrid>
              <Box mt={6}>
                <DetailDisclosure title='Show raw policy fields'>
                  <Box>
                    <Text fontSize='2xs' color='ink.4' mb={2} letterSpacing='0.06em'>
                      POLICY
                    </Text>
                    <RawJson value={epoch.data.epoch.policy} />
                  </Box>
                </DetailDisclosure>
              </Box>
            </RoundSection>

            <RoundSection title='Activity counters'>
              <SimpleGrid columns={{ base: 2, md: 5 }} gap={3}>
                <Counter
                  label='Claimed'
                  value={`${epoch.data.epoch.claimedCount}/${epoch.data.epoch.policy.committeeSize}`}
                />
                <Counter
                  label='Contributions'
                  value={`${epoch.data.epoch.contributionCount}/${epoch.data.epoch.policy.committeeSize}`}
                />
                <Counter label='Ciphertexts' value={epoch.data.epoch.ciphertextCount.toString()} />
                <Counter
                  label='Partial decryptions'
                  value={epoch.data.epoch.partialDecryptionCount.toString()}
                />
              </SimpleGrid>
            </RoundSection>

            <ApplicationsSection epochId={epochId as Hex} />

            <RoundSection title='Participants & activity'>
              <Tabs.Root defaultValue='participants' variant='line'>
                <Tabs.List borderColor='rule'>
                  <DkgTab value='participants'>Participants</DkgTab>
                  <DkgTab value='activity'>Activity</DkgTab>
                </Tabs.List>
                <Tabs.Content value='participants' pt={6}>
                  <ParticipantList participants={epoch.data.participants} />
                </Tabs.Content>
                <Tabs.Content value='activity' pt={6}>
                  <QueryDataLayout
                    isLoading={events.isLoading}
                    isError={events.isError}
                    error={events.error}
                    isEmpty={events.data?.length === 0}
                    emptyMessage='No on-chain events for this epoch yet.'
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

// ─── Per-application section ─────────────────────────────────────────────
// Two stacked panels: a registration form (organizer-driven) and a status
// pipeline (read-only). The pipeline takes the aid the user has typed into
// the inspector, so a single page can both register and observe.
function ApplicationsSection({ epochId }: { epochId: Hex }) {
  const [aidUnderInspection, setAidUnderInspection] = useState<Hex>(
    ('0x' + '00'.repeat(32)) as Hex,
  )
  const pkEp = useCollectivePublicKey(epochId)
  return (
    <RoundSection title='Per-application keys'>
      <Stack gap={6}>
        <AppRegistrationForm
          epochId={epochId}
          pkEp={pkEp.data ?? null}
          onRegistered={(_tx, aid) => setAidUnderInspection(aid)}
        />
        <Stack gap={2}>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            Inspect application
          </Text>
          <Input
            value={aidUnderInspection}
            onChange={(e) => setAidUnderInspection(e.target.value as Hex)}
            fontFamily='mono'
            fontSize='xs'
            spellCheck={false}
          />
        </Stack>
        <DecryptionPipeline epochId={epochId} aid={aidUnderInspection} />
      </Stack>
    </RoundSection>
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
