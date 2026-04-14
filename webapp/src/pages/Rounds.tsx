import {
  Box,
  Heading,
  Table,
  TableContainer,
  Tbody,
  Td,
  Text,
  Th,
  Thead,
  Tr,
  VStack,
} from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { ErrorBanner } from '../components/ErrorBanner';
import { HashCell } from '../components/HashCell';
import { StatusBadge } from '../components/StatusBadge';
import { useRecentRounds } from '../lib/hooks';

export function Rounds() {
  const rounds = useRecentRounds(64);
  const navigate = useNavigate();
  return (
    <VStack align="stretch" spacing={4}>
      <Heading size="lg">Rounds</Heading>
      <Text color="gray.400" fontSize="sm">
        Showing the most recent rounds retained in the on-chain ring buffer (max 64).
      </Text>
      {rounds.isError && <ErrorBanner error={rounds.error} title="Failed to load rounds" />}
      <Box bg="gray.800" borderRadius="md" borderWidth="1px" borderColor="gray.700">
        <TableContainer>
          <Table size="sm" variant="simple">
            <Thead>
              <Tr>
                <Th>Round ID</Th>
                <Th>Status</Th>
                <Th>Nonce</Th>
                <Th>Organizer</Th>
                <Th isNumeric>Committee</Th>
                <Th isNumeric>Threshold</Th>
                <Th isNumeric>Claimed</Th>
                <Th isNumeric>Contribs</Th>
              </Tr>
            </Thead>
            <Tbody>
              {rounds.data?.map(({ id, round }) => (
                <Tr
                  key={id}
                  onClick={() => navigate(`/rounds/${id}`)}
                  _hover={{ bg: 'gray.750' }}
                  cursor="pointer"
                >
                  <Td>
                    <HashCell value={id} head={10} tail={6} />
                  </Td>
                  <Td>
                    <StatusBadge status={Number(round.status)} />
                  </Td>
                  <Td fontFamily="mono">{round.nonce.toString()}</Td>
                  <Td>
                    <HashCell value={round.organizer} />
                  </Td>
                  <Td isNumeric fontFamily="mono">
                    {round.policy.committeeSize}
                  </Td>
                  <Td isNumeric fontFamily="mono">
                    {round.policy.threshold}
                  </Td>
                  <Td isNumeric fontFamily="mono">
                    {round.claimedCount}
                  </Td>
                  <Td isNumeric fontFamily="mono">
                    {round.contributionCount}
                  </Td>
                </Tr>
              ))}
            </Tbody>
          </Table>
        </TableContainer>
        {rounds.isLoading && (
          <Text p={4} color="gray.400">
            Loading…
          </Text>
        )}
        {rounds.data && rounds.data.length === 0 && (
          <Text p={4} color="gray.400">
            No rounds yet.
          </Text>
        )}
      </Box>
    </VStack>
  );
}
