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
import { useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { StatusBadge } from '../components/StatusBadge';
import { HashCell } from '../components/HashCell';
import {
  useActiveNodeCount,
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
  const activeNodes = useActiveNodeCount();
  const tip = useChainTip();
  const recent = useRecentRounds(5);

  useEffect(() => {
    if (cfg.isError) console.error('[Home] Failed to load config:', cfg.error);
  }, [cfg.isError, cfg.error]);
  useEffect(() => {
    if (recent.isError) console.error('[Home] Failed to load rounds:', recent.error);
  }, [recent.isError, recent.error]);

  return (
    <VStack align="stretch" spacing={6}>
      <Heading size="lg">Overview</Heading>
      <SimpleGrid columns={{ base: 1, md: 4 }} spacing={4}>
        <StatCard
          label="Total rounds"
          value={nonce.data !== undefined ? nonce.data.toString() : '…'}
        />
        <StatCard
          label="Active / total nodes"
          value={
            activeNodes.data !== undefined && registryCount.data !== undefined
              ? `${activeNodes.data.toString()} / ${registryCount.data.toString()}`
              : '…'
          }
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
