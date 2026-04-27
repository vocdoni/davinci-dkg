import { Box, HStack, Stack, Table, Text } from '@chakra-ui/react'
import { Link as RouterLink, generatePath } from 'react-router-dom'
import type { RoundEntry } from '@vocdoni/davinci-dkg-sdk'
import { Routes } from '~router/routes'
import { StatusBadge } from './StatusBadge'
import { HashCell } from '~components/ui/HashCell'

interface Props {
  rounds: RoundEntry[]
  /** Optional cap; useful for the home page's "recent 5". */
  limit?: number
}

export function RoundList({ rounds, limit }: Props) {
  const items = limit ? rounds.slice(0, limit) : rounds
  if (items.length === 0) {
    return (
      <Box p={6} textAlign='center' color='gray.500' fontSize='sm'>
        No rounds yet on this network.
      </Box>
    )
  }
  return (
    <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' overflow='hidden'>
      <Table.Root size='sm'>
        <Table.Header>
          <Table.Row bg='gray.900'>
            <Table.ColumnHeader>Round</Table.ColumnHeader>
            <Table.ColumnHeader>Status</Table.ColumnHeader>
            <Table.ColumnHeader>Organizer</Table.ColumnHeader>
            <Table.ColumnHeader>Threshold</Table.ColumnHeader>
            <Table.ColumnHeader>Activity</Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {items.map(({ id, round }) => (
            <Table.Row key={id} _hover={{ bg: 'gray.800' }}>
              <Table.Cell>
                <RouterLink to={generatePath(Routes.round, { id })}>
                  <Stack gap={0}>
                    <HashCell value={id} head={6} tail={4} />
                    <Text fontSize='2xs' color='gray.500'>
                      nonce {round.nonce.toString()}
                    </Text>
                  </Stack>
                </RouterLink>
              </Table.Cell>
              <Table.Cell>
                <StatusBadge round={round} />
              </Table.Cell>
              <Table.Cell>
                <HashCell value={round.organizer} />
              </Table.Cell>
              <Table.Cell>
                <Text fontSize='xs' fontFamily='mono'>
                  {round.policy.threshold} of {round.policy.committeeSize}
                </Text>
              </Table.Cell>
              <Table.Cell>
                <HStack gap={3} fontSize='2xs' color='gray.400'>
                  <Text>claims {round.claimedCount}/{round.policy.committeeSize}</Text>
                  <Text>contribs {round.contributionCount}</Text>
                  <Text>cipher {round.ciphertextCount}</Text>
                </HStack>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </Box>
  )
}
