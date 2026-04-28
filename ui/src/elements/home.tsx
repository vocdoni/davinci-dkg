import { Box, Heading, HStack, Link, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import { Link as RouterLink } from 'react-router-dom'
import type { ReactNode } from 'react'
import { useConfig } from '~providers/ConfigProvider'
import { useBlockNumber } from '~queries/chain'
import { useRecentRounds } from '~queries/rounds'
import { useRegistryStats, useRoundCount } from '~queries/registry'
import { StatCard } from '~components/ui/StatCard'
import { RoundList } from '~components/Round/RoundList'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { Routes } from '~router/routes'

const PAPER_URL = 'https://eprint.iacr.org/2026/552'
const DAVINCI_URL = 'https://davinci.vote'

export function Home() {
  const config = useConfig()
  const { data: block } = useBlockNumber()
  const { data: roundNonce } = useRoundCount()
  const stats = useRegistryStats()
  const recent = useRecentRounds(5)

  return (
    <Stack gap={{ base: 14, md: 20 }}>
      {/* Lede */}
      <Stack gap={5} maxW='640px'>
        <Stack gap={2}>
          <Heading
            as='h1'
            fontSize={{ base: '4xl', md: '5xl' }}
            fontWeight={500}
            color='ink.0'
            letterSpacing='-0.02em'
            lineHeight='1.05'
          >
            NI-DKG
          </Heading>
          <Text
            fontSize={{ base: 'md', md: 'lg' }}
            color='ink.2'
            letterSpacing='-0.005em'
          >
            Non-Interactive Distributed Key Generation
          </Text>
        </Stack>
        <Text fontSize={{ base: 'sm', md: 'md' }} color='ink.2' lineHeight='1.65'>
          A committee of independent operators jointly generates an{' '}
          <Term>ElGamal</Term> public key on{' '}
          <Term>BabyJubJub</Term> — a SNARK-friendly elliptic curve. The matching
          private key is never reconstructed; ciphertexts are decrypted by
          combining partial decryptions, and ElGamal's additive homomorphism
          allows aggregating ciphertexts before they are decrypted.
        </Text>
        <Text fontSize={{ base: 'sm', md: 'md' }} color='ink.2' lineHeight='1.65'>
          NI-DKG is a building block of the{' '}
          <Ext href={DAVINCI_URL}>davinci protocol</Ext>. Full construction and
          security proofs are described in the{' '}
          <Ext href={PAPER_URL}>NI-DKG paper</Ext>.
        </Text>
      </Stack>

      {/* Phases */}
      <Section title='Protocol'>
        <Stack gap={0} mt={1}>
          <Phase
            n={1}
            title='Permissionless committee'
            body={
              <>
                Anyone can <RouterQuick to={Routes.runNode}>run a node</RouterQuick>{' '}
                and join the registry. For each round, a verifiable on-chain
                lottery selects a committee from the registered operators — no
                gatekeeper, no allowlist.
              </>
            }
          />
          <Phase
            n={2}
            title='Verifiable key generation'
            body={
              <>
                Every contribution is a single transaction with an attached
                zk-SNARK that proves the operator's polynomial commitments and
                encrypted shares are consistent. The contract aggregates the
                accepted contributions into a collective public key.
              </>
            }
          />
          <Phase
            n={3}
            title='Threshold decryption'
            body={
              <>
                The private key is never disclosed. Instead, the committee
                decrypts <em>specific ciphertexts</em>: the result of a vote can
                be decrypted while the individual ballots remain hidden. Each
                partial decryption ships with a Chaum–Pedersen DLEQ proof; once
                the threshold is met, anyone can combine them into the
                plaintext.
              </>
            }
            last
          />
        </Stack>
      </Section>

      {/* Properties */}
      <Section title='Properties'>
        <Stack gap={5} mt={1}>
          <Property
            label='No trusted dealer'
            body='The collective key is generated jointly. Corrupting fewer than t operators reveals nothing.'
          />
          <Property
            label='Non-interactive'
            body='Each step is a single self-contained transaction. No complaint or dispute round; invalid contributions are rejected at submission time.'
          />
          <Property
            label='Verified on-chain'
            body='Every state-changing call is gated by a Groth16 verifier. Correctness is enforced by the EVM, not by an off-chain auditor.'
          />
        </Stack>
      </Section>

      {/* Live deployment */}
      <Section
        title='Live deployment'
        subtitle={`Snapshot of ${config.chainName}, refreshed every few seconds.`}
      >
        <SimpleGrid columns={{ base: 2, md: 4 }} gap={{ base: 3, md: 4 }} mt={1}>
          <StatCard
            label='Rounds'
            value={roundNonce != null ? roundNonce.toString() : '—'}
            hint='total ever created'
          />
          <StatCard
            label='Active nodes'
            value={stats.data ? stats.data.active.toString() : '—'}
            hint={stats.data ? `${stats.data.total.toString()} registered` : undefined}
          />
          <StatCard
            label='Latest block'
            value={block ? `#${block.toString()}` : '—'}
            hint={config.chainName}
            tone='live'
          />
          <StatCard label='Chain id' value={config.chainId.toString()} hint='from /config.json' />
        </SimpleGrid>
      </Section>

      {/* Recent rounds */}
      <Section title='Recent rounds' action={<RouterQuick to={Routes.rounds}>View all</RouterQuick>}>
        <Box mt={1}>
          <QueryDataLayout
            isLoading={recent.isLoading}
            isError={recent.isError}
            error={recent.error}
            isEmpty={recent.data?.length === 0}
            emptyMessage='No rounds have been created yet.'
          >
            {recent.data && <RoundList rounds={recent.data} />}
          </QueryDataLayout>
        </Box>
      </Section>
    </Stack>
  )
}

function Section({
  title,
  subtitle,
  action,
  children,
}: {
  title: string
  subtitle?: string
  action?: ReactNode
  children: ReactNode
}) {
  return (
    <Box as='section'>
      <HStack justify='space-between' align='baseline' mb={subtitle ? 1 : 5} wrap='wrap' gap={3}>
        <Heading
          as='h2'
          fontSize={{ base: 'lg', md: 'xl' }}
          fontWeight={500}
          color='ink.0'
          letterSpacing='-0.01em'
        >
          {title}
        </Heading>
        {action}
      </HStack>
      {subtitle && (
        <Text fontSize='sm' color='ink.3' mb={5} lineHeight='1.5'>
          {subtitle}
        </Text>
      )}
      {children}
    </Box>
  )
}

function Phase({ n, title, body, last }: { n: number; title: string; body: ReactNode; last?: boolean }) {
  return (
    <HStack
      align='start'
      gap={{ base: 5, md: 7 }}
      py={{ base: 5, md: 6 }}
      borderTopWidth='1px'
      borderTopColor='rule'
      borderBottomWidth={last ? '1px' : 0}
      borderBottomColor='rule'
    >
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize={{ base: 'xs', md: 'sm' }}
        color='accent.fg'
        flexShrink={0}
        pt='3px'
        minW='28px'
      >
        {n.toString().padStart(2, '0')}
      </Text>
      <Box>
        <Heading
          as='h3'
          fontSize={{ base: 'md', md: 'lg' }}
          fontWeight={500}
          color='ink.0'
          mb={2}
          letterSpacing='-0.01em'
        >
          {title}
        </Heading>
        <Text fontSize={{ base: 'sm', md: 'md' }} color='ink.2' lineHeight='1.6' maxW='62ch'>
          {body}
        </Text>
      </Box>
    </HStack>
  )
}

function Property({ label, body }: { label: string; body: string }) {
  return (
    <Box>
      <Text fontSize='xs' fontWeight={600} color='accent.fg' letterSpacing='0.01em' mb={1.5}>
        {label}
      </Text>
      <Text fontSize={{ base: 'sm', md: 'md' }} color='ink.2' lineHeight='1.6' maxW='62ch'>
        {body}
      </Text>
    </Box>
  )
}

function Term({ children }: { children: ReactNode }) {
  return (
    <Box as='span' color='ink.0' fontWeight={500}>
      {children}
    </Box>
  )
}

function Ext({ href, children }: { href: string; children: ReactNode }) {
  return (
    <Link
      href={href}
      target='_blank'
      rel='noopener noreferrer'
      color='accent.fg'
      borderBottomWidth='1px'
      borderColor='accent.border'
      _hover={{ color: 'accent.bright', borderColor: 'accent.fg' }}
      transition='color 0.15s, border-color 0.15s'
    >
      {children}
    </Link>
  )
}

function RouterQuick({ to, children }: { to: string; children: ReactNode }) {
  return (
    <RouterLink to={to}>
      <Text
        as='span'
        color='accent.fg'
        fontWeight={500}
        borderBottomWidth='1px'
        borderColor='accent.border'
        pb='1px'
        _hover={{ color: 'accent.bright', borderColor: 'accent.fg' }}
        transition='color 0.15s, border-color 0.15s'
      >
        {children}
      </Text>
    </RouterLink>
  )
}
