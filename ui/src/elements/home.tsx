import { Box, Heading, HStack, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import { Link as RouterLink } from 'react-router-dom'
import { LuArrowRight } from 'react-icons/lu'
import type { ReactNode } from 'react'
import { useConfig } from '~providers/ConfigProvider'
import { useBlockNumber } from '~queries/chain'
import { useRecentRounds } from '~queries/rounds'
import { useRegistryStats, useRoundCount } from '~queries/registry'
import { StatCard } from '~components/ui/StatCard'
import { RoundList } from '~components/Round/RoundList'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { Routes } from '~router/routes'

// Landing page. Style: research project / academic.
//
// Sections in reading order:
//   1. Abstract  — one-paragraph protocol summary, no decorative chrome.
//   2. Phases    — the three-phase flow as a numbered table, not as
//                  marketing cards.
//   3. Properties — design choices, neutrally phrased.
//   4. Network   — live KPI snapshot.
//   5. Recent    — the most actionable thing on the page.
export function Home() {
  const config = useConfig()
  const { data: block } = useBlockNumber()
  const { data: roundNonce } = useRoundCount()
  const stats = useRegistryStats()
  const recent = useRecentRounds(5)

  return (
    <Stack gap={12} maxW='4xl'>
      {/* ── 1. Abstract ─────────────────────────────────────────────── */}
      <Box pb={6} borderBottomWidth='1px' borderColor='gray.800'>
        <Heading size='xl' color='gray.100' mb={2}>
          davinci-dkg
        </Heading>
        <Text color='gray.400' fontSize='sm' mb={4}>
          A non-interactive distributed key generation and threshold decryption protocol for
          EVM-compatible chains.
        </Text>
        <Text color='gray.300' fontSize='sm' lineHeight='1.7'>
          A committee of independent operators jointly generates a single ElGamal public key on
          BabyJubJub. No party ever holds the corresponding private key. Anyone can encrypt a
          message for the committee; decryption requires a configurable threshold of operators to
          publish partial decryptions, which any caller can then combine on-chain into the
          plaintext. Every protocol step — contributions, finalize, partial decryptions, combine —
          is gated by a Groth16 zk-SNARK that the smart contract verifies before accepting state
          changes. There is no dispute phase and no off-chain coordination.
        </Text>
        <HStack gap={5} mt={5} fontSize='sm'>
          <RouterLink to={Routes.playground}>
            <HStack gap={1.5} color='cyan.300' _hover={{ color: 'cyan.200' }}>
              <Text>Playground</Text>
              <LuArrowRight />
            </HStack>
          </RouterLink>
          <RouterLink to={Routes.runNode}>
            <HStack gap={1.5} color='cyan.300' _hover={{ color: 'cyan.200' }}>
              <Text>Run a node</Text>
              <LuArrowRight />
            </HStack>
          </RouterLink>
          <RouterLink to={Routes.sdk}>
            <HStack gap={1.5} color='cyan.300' _hover={{ color: 'cyan.200' }}>
              <Text>SDK reference</Text>
              <LuArrowRight />
            </HStack>
          </RouterLink>
        </HStack>
      </Box>

      {/* ── 2. Phases ───────────────────────────────────────────────── */}
      <Box>
        <Heading size='md' color='gray.100' mb={3}>
          Protocol phases
        </Heading>
        <Stack gap={0} borderWidth='1px' borderColor='gray.800' borderRadius='md' overflow='hidden'>
          <Phase
            n={1}
            title='Lottery and committee selection'
            body='Anyone opens a round on the manager contract. A future-block hash seeds a deterministic lottery; the registered nodes whose addresses pass the threshold may claim slots until the committee is full or the registration window closes.'
          />
          <Phase
            n={2}
            title='Distributed key generation'
            body='Each selected node publishes a Feldman / Pedersen contribution: a zk-SNARK that simultaneously commits to its polynomial coefficients, encrypts the per-recipient shares, and proves both are consistent. The contract aggregates the accepted contributions into a single collective public key.'
          />
          <Phase
            n={3}
            title='Threshold decryption'
            body='Once the round is finalized, anyone can submit an ElGamal ciphertext. The committee posts partial decryptions, each with a Chaum–Pedersen DLEQ proof; once t partials accumulate, any caller can combine them into the plaintext.'
            last
          />
        </Stack>
      </Box>

      {/* ── 3. Properties ───────────────────────────────────────────── */}
      <Box>
        <Heading size='md' color='gray.100' mb={3}>
          Properties
        </Heading>
        <Stack gap={3}>
          <Property
            label='No trusted dealer'
            body='The collective key is generated jointly. Corrupting fewer than t operators reveals nothing about the secret.'
          />
          <Property
            label='Non-interactive'
            body='Every step is a single transaction with a self-contained ZK proof. There is no complaint or dispute round; invalid contributions are simply rejected at submission time.'
          />
          <Property
            label='Verifiable on-chain'
            body='The contract verifies a Groth16 proof for every state-changing call. No oracle, no off-chain auditor; correctness is enforced by the EVM.'
          />
          <Property
            label='Threshold by parameter'
            body='The (t, n) tradeoff between availability and resilience is set per round. A round with min ≥ t can tolerate n − t operators going offline post-finalize without losing decryptability.'
          />
          <Property
            label='EVM-native deployment'
            body='A single Solidity contract orchestrates the protocol; operators run a Docker image and consumers use a viem-based TypeScript SDK. No custom L2, no custom networking layer.'
          />
        </Stack>
      </Box>

      {/* ── 4. Network ──────────────────────────────────────────────── */}
      <Box>
        <Heading size='md' color='gray.100' mb={1}>
          Live deployment
        </Heading>
        <Text fontSize='sm' color='gray.400' mb={3}>
          Snapshot of the connected chain — refreshed every few seconds.
        </Text>
        <SimpleGrid columns={{ base: 2, md: 4 }} gap={3}>
          <StatCard
            label='Rounds'
            value={roundNonce != null ? roundNonce.toString() : '—'}
            hint='Total ever created'
          />
          <StatCard
            label='Active nodes'
            value={stats.data ? stats.data.active.toString() : '—'}
            hint={stats.data ? `${stats.data.total.toString()} registered` : undefined}
          />
          <StatCard label='Latest block' value={block ? `#${block.toString()}` : '—'} hint={config.chainName} />
          <StatCard label='Chain id' value={config.chainId.toString()} hint='from /config.json' />
        </SimpleGrid>
      </Box>

      {/* ── 5. Recent rounds ────────────────────────────────────────── */}
      <Box>
        <HStack justify='space-between' mb={3}>
          <Heading size='md' color='gray.100'>
            Recent rounds
          </Heading>
          <RouterLink to={Routes.rounds}>
            <HStack gap={1} fontSize='sm' color='cyan.300'>
              <Text>View all</Text>
              <LuArrowRight />
            </HStack>
          </RouterLink>
        </HStack>
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
    </Stack>
  )
}

function Phase({ n, title, body, last }: { n: number; title: string; body: string; last?: boolean }) {
  return (
    <HStack
      gap={4}
      align='start'
      p={4}
      borderBottomWidth={last ? 0 : '1px'}
      borderColor='gray.800'
      bg='gray.900'
    >
      <Box
        flexShrink={0}
        w={7}
        h={7}
        borderRadius='full'
        borderWidth='1px'
        borderColor='cyan.700'
        color='cyan.300'
        display='flex'
        alignItems='center'
        justifyContent='center'
        fontSize='xs'
        fontFamily='mono'
        fontWeight='semibold'
      >
        {n}
      </Box>
      <Box>
        <Text fontSize='sm' fontWeight='semibold' color='gray.100' mb={1}>
          {title}
        </Text>
        <Text fontSize='sm' color='gray.400' lineHeight='1.6'>
          {body}
        </Text>
      </Box>
    </HStack>
  )
}

function Property({ label, body }: { label: string; body: ReactNode }) {
  return (
    <HStack gap={4} align='start'>
      <Text
        fontSize='xs'
        color='cyan.300'
        fontFamily='mono'
        textTransform='uppercase'
        letterSpacing='wider'
        minW='160px'
        pt='2px'
        flexShrink={0}
      >
        {label}
      </Text>
      <Text fontSize='sm' color='gray.300' lineHeight='1.6'>
        {body}
      </Text>
    </HStack>
  )
}
