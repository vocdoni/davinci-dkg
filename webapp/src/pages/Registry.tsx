import {
  Box,
  Heading,
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
import { ErrorBanner } from '../components/ErrorBanner';
import { HashCell } from '../components/HashCell';
import { useRegistryNodes } from '../lib/hooks';

const NODE_STATUS = ['None', 'Active', 'Suspended'];

export function Registry() {
  const nodes = useRegistryNodes();
  return (
    <VStack align="stretch" spacing={4}>
      <Heading size="lg">Registry</Heading>
      <Text color="gray.400" fontSize="sm">
        Operators registered with the DKGRegistry contract and their on-chain BabyJubJub public
        keys.
      </Text>
      {nodes.isError && <ErrorBanner error={nodes.error} title="Failed to load registry" />}
      <Box bg="gray.800" borderRadius="md" borderWidth="1px" borderColor="gray.700">
        <TableContainer>
          <Table size="sm" variant="simple">
            <Thead>
              <Tr>
                <Th>Operator</Th>
                <Th>Status</Th>
                <Th>Pub.X</Th>
                <Th>Pub.Y</Th>
              </Tr>
            </Thead>
            <Tbody>
              {nodes.data?.map((n: any) => (
                <Tr key={n.operator}>
                  <Td>
                    <HashCell value={n.operator} />
                  </Td>
                  <Td>
                    <Tag
                      colorScheme={Number(n.status) === 1 ? 'green' : 'gray'}
                      fontFamily="mono"
                    >
                      {NODE_STATUS[Number(n.status)] ?? 'Unknown'}
                    </Tag>
                  </Td>
                  <Td>
                    <HashCell
                      value={`0x${BigInt(n.pubX).toString(16).padStart(64, '0')}`}
                    />
                  </Td>
                  <Td>
                    <HashCell
                      value={`0x${BigInt(n.pubY).toString(16).padStart(64, '0')}`}
                    />
                  </Td>
                </Tr>
              ))}
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
