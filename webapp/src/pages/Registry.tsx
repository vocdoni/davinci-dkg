import {
  Box,
  Heading,
  HStack,
  Stat,
  StatLabel,
  StatNumber,
  SimpleGrid,
  Table,
  TableContainer,
  Tag,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  VStack,
} from '@chakra-ui/react';
import { useEffect } from 'react';
import { nodeStatusColor, nodeStatusLabel } from '../lib/abi';
import { HashCell } from '../components/HashCell';
import {
  useActiveNodeCount,
  useChainTip,
  useInactivityWindow,
  useRegistry,
  useRegistryNodes,
} from '../lib/hooks';

function formatBlockDelta(
  lastActive: bigint | undefined,
  head: bigint | undefined,
): string {
  if (lastActive === undefined || head === undefined) return '—';
  if (head < lastActive) return 'now';
  const delta = head - lastActive;
  if (delta === 0n) return 'this block';
  return `${delta.toString()} blocks ago`;
}

function StatCard({ label, value }: { label: string; value: React.ReactNode }) {
  return (
    <Box bg="gray.800" p={4} borderRadius="md" borderWidth="1px" borderColor="gray.700">
      <Stat>
        <StatLabel color="gray.400" fontSize="xs" textTransform="uppercase">
          {label}
        </StatLabel>
        <StatNumber fontFamily="mono" fontSize="xl" color="cyan.200">
          {value}
        </StatNumber>
      </Stat>
    </Box>
  );
}

export function Registry() {
  const nodes = useRegistryNodes();
  const total = useRegistry();
  const active = useActiveNodeCount();
  const tip = useChainTip();
  const window = useInactivityWindow();

  useEffect(() => {
    if (nodes.isError) console.error('[Registry] Failed to load nodes:', nodes.error);
  }, [nodes.isError, nodes.error]);

  return (
    <VStack align="stretch" spacing={4}>
      <Heading size="lg">Registry</Heading>
      <Text color="gray.400" fontSize="sm">
        Operators registered with the DKGRegistry contract and their on-chain BabyJubJub public
        keys. The lottery draws only from nodes whose status is <strong>Active</strong>;
        stragglers whose <code>lastActiveBlock</code> falls behind by more than the inactivity
        window can be pruned by anyone via <code>reap(operator)</code>.
      </Text>

      <SimpleGrid columns={{ base: 1, sm: 2, md: 4 }} spacing={3}>
        <StatCard
          label="Active nodes"
          value={active.data !== undefined ? active.data.toString() : '…'}
        />
        <StatCard
          label="Total ever registered"
          value={total.data !== undefined ? total.data.toString() : '…'}
        />
        <StatCard
          label="Latest block"
          value={tip.data ? `#${tip.data.number.toString()}` : '…'}
        />
        <StatCard
          label="Inactivity window"
          value={
            window.data !== undefined ? `${window.data.toString()} blocks` : '…'
          }
        />
      </SimpleGrid>

      <Box bg="gray.800" borderRadius="md" borderWidth="1px" borderColor="gray.700">
        <TableContainer>
          <Table size="sm" variant="simple">
            <Thead>
              <Tr>
                <Th>Operator</Th>
                <Th>Status</Th>
                <Th>Last active</Th>
                <Th>Pub.X</Th>
                <Th>Pub.Y</Th>
              </Tr>
            </Thead>
            <Tbody>
              {nodes.data?.map((n: any) => {
                const status = Number(n.status);
                const lastActive =
                  n.lastActiveBlock !== undefined ? BigInt(n.lastActiveBlock) : undefined;
                return (
                  <Tr key={n.operator}>
                    <Td>
                      <HashCell value={n.operator} />
                    </Td>
                    <Td>
                      <Tag colorScheme={nodeStatusColor(status)} fontFamily="mono">
                        {nodeStatusLabel(status)}
                      </Tag>
                    </Td>
                    <Td fontFamily="mono" fontSize="xs" color="gray.400">
                      <HStack spacing={2}>
                        <Text>
                          {lastActive !== undefined ? `#${lastActive.toString()}` : '—'}
                        </Text>
                        <Text color="gray.500">
                          {formatBlockDelta(lastActive, tip.data?.number)}
                        </Text>
                      </HStack>
                    </Td>
                    <Td>
                      <HashCell value={`0x${BigInt(n.pubX).toString(16).padStart(64, '0')}`} />
                    </Td>
                    <Td>
                      <HashCell value={`0x${BigInt(n.pubY).toString(16).padStart(64, '0')}`} />
                    </Td>
                  </Tr>
                );
              })}
            </Tbody>
          </Table>
        </TableContainer>
        {nodes.isLoading && (
          <Text p={4} color="gray.400">
            Loading…
          </Text>
        )}
        {nodes.data && nodes.data.length === 0 && (
          <Text p={4} color="gray.400">
            No nodes registered.
          </Text>
        )}
      </Box>
    </VStack>
  );
}
