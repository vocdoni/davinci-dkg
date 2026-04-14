import {
  Box,
  Heading,
  SimpleGrid,
  Stat,
  StatLabel,
  StatNumber,
  Text,
  VStack,
} from '@chakra-ui/react';
import { Link as RouterLink } from 'react-router-dom';
import { ErrorBanner } from '../components/ErrorBanner';
import { StatusBadge } from '../components/StatusBadge';
import { HashCell } from '../components/HashCell';
import {
  useChainTip,
  useConfig,
  useRecentRounds,
  useRegistry,
  useRoundNonce,
} from '../lib/hooks';

function StatCard({ label, value }: { label: string; value: React.ReactNode }) {
  return (
    <Box bg="gray.800" p={5} borderRadius="md" borderWidth="1px" borderColor="gray.700">
      <Stat>
        <StatLabel color="gray.400" fontSize="xs" textTransform="uppercase">
          {label}
        </StatLabel>
        <StatNumber fontFamily="mono" fontSize="2xl" color="cyan.200">
          {value}
        </StatNumber>
      </Stat>
    </Box>
  );
}

export function Home() {
  const cfg = useConfig();
  const nonce = useRoundNonce();
  const registryCount = useRegistry();
  const tip = useChainTip();
  const recent = useRecentRounds(5);

  return (
    <VStack align="stretch" spacing={6}>
      <Heading size="lg">Overview</Heading>
      {cfg.isError && <ErrorBanner error={cfg.error} title="Failed to load config" />}
      {recent.isError && <ErrorBanner error={recent.error} title="Failed to load rounds" />}
      <SimpleGrid columns={{ base: 1, md: 4 }} spacing={4}>
        <StatCard
          label="Total rounds"
          value={nonce.data !== undefined ? nonce.data.toString() : '…'}
        />
        <StatCard
          label="Registered nodes"
          value={registryCount.data !== undefined ? registryCount.data.toString() : '…'}
        />
        <StatCard
          label="Latest block"
          value={tip.data ? `#${tip.data.number.toString()}` : '…'}
        />
        <StatCard label="Chain ID" value={cfg.data ? cfg.data.chainId : '…'} />
      </SimpleGrid>

      <Box bg="gray.800" p={5} borderRadius="md" borderWidth="1px" borderColor="gray.700">
        <Heading size="md" mb={4}>
          Recent rounds
        </Heading>
        {recent.isLoading && <Text color="gray.400">Loading…</Text>}
        {recent.data && recent.data.length === 0 && (
          <Text color="gray.400">No rounds yet.</Text>
        )}
        <VStack align="stretch" spacing={2}>
          {recent.data?.map(({ id, round }) => (
            <Box
              key={id}
              as={RouterLink}
              to={`/rounds/${id}`}
              p={3}
              borderRadius="sm"
              borderWidth="1px"
              borderColor="gray.700"
              _hover={{ bg: 'gray.750', borderColor: 'cyan.700' }}
              display="flex"
              alignItems="center"
              gap={4}
            >
              <Box minW="220px">
                <HashCell value={id} head={10} tail={6} />
              </Box>
              <StatusBadge status={Number(round.status)} />
              <Text fontFamily="mono" fontSize="sm" color="gray.400">
                nonce {round.nonce.toString()}
              </Text>
              <Text fontFamily="mono" fontSize="sm" color="gray.400">
                committee {round.policy.committeeSize} · t {round.policy.threshold}
              </Text>
            </Box>
          ))}
        </VStack>
      </Box>
    </VStack>
  );
}
