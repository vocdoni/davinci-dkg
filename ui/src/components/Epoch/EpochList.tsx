import { Box, HStack, Stack, Table, Text } from '@chakra-ui/react'
import { Link as RouterLink, generatePath } from 'react-router-dom'
import type { EpochEntry } from '@vocdoni/davinci-dkg-sdk'
import { Routes } from '~router/routes'
import { StatusBadge } from './StatusBadge'
import { HashCell } from '~components/ui/HashCell'

interface Props {
  epochs: EpochEntry[]
  /** Optional cap; useful for the home page's "recent 5". */
  limit?: number
}

// Refined editorial table: surface card with hairline header rule, mono
// digits everywhere a number lives, the row-hover tint is gold rather than
// the usual cyan/blue. Activity counters use a "k v" mono pattern that
// reads as data, not as UI.
export function EpochList({ epochs, limit }: Props) {
  const items = limit ? epochs.slice(0, limit) : epochs
  if (items.length === 0) {
    return (
      <Box
        borderWidth='1px'
        borderColor='border.subtle'
        borderRadius='lg'
        bg='surface'
        p={{ base: 8, md: 12 }}
        textAlign='center'
      >
        <Text color='ink.3' fontSize='sm'>
          No epochs yet on this network.
        </Text>
      </Box>
    )
  }
  return (
    <Box
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='lg'
      bg='surface'
      overflow='hidden'
      boxShadow='inset'
    >
      <Table.Root size='sm' interactive>
        <Table.Header>
          <Table.Row bg='surface.sunken'>
            <Th>Epoch</Th>
            <Th>Status</Th>
            <Th>Organizer</Th>
            <Th>Threshold</Th>
            <Th>Activity</Th>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {items.map(({ id, epoch }) => (
            <Table.Row
              key={id}
              borderTopWidth='1px'
              borderColor='rule'
              transition='background 0.12s'
              _hover={{ bg: 'accent.bg' }}
            >
              <Table.Cell py={3.5}>
                <RouterLink to={generatePath(Routes.epoch, { id })}>
                  <Stack gap={0.5}>
                    <HashCell value={id} head={6} tail={4} />
                    <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.4'>
                      nonce {epoch.nonce.toString()}
                    </Text>
                  </Stack>
                </RouterLink>
              </Table.Cell>
              <Table.Cell py={3.5}>
                <StatusBadge epoch={epoch} />
              </Table.Cell>
              <Table.Cell py={3.5}>
                <HashCell value={epoch.organizer} />
              </Table.Cell>
              <Table.Cell py={3.5}>
                <Text className='dkg-tabular' fontFamily='mono' fontSize='xs' color='ink.1'>
                  {epoch.policy.threshold}
                  <Box as='span' color='ink.4' mx='0.4em'>
                    of
                  </Box>
                  {epoch.policy.committeeSize}
                </Text>
              </Table.Cell>
              <Table.Cell py={3.5}>
                <HStack gap={4} fontFamily='mono' fontSize='2xs' className='dkg-tabular'>
                  <ActivityKV k='claims' v={`${epoch.claimedCount}/${epoch.policy.committeeSize}`} />
                  <ActivityKV k='contribs' v={epoch.contributionCount.toString()} />
                  <ActivityKV k='cipher' v={epoch.ciphertextCount.toString()} />
                </HStack>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table.Root>
    </Box>
  )
}

function Th({ children }: { children: React.ReactNode }) {
  return (
    <Table.ColumnHeader
      fontFamily='mono'
      fontSize='2xs'
      fontWeight={500}
      color='ink.3'
      letterSpacing='0.08em'
      textTransform='uppercase'
      py={3}
      borderColor='rule'
    >
      {children}
    </Table.ColumnHeader>
  )
}

function ActivityKV({ k, v }: { k: string; v: string }) {
  return (
    <HStack gap={1}>
      <Text color='ink.4'>{k}</Text>
      <Text color='ink.1'>{v}</Text>
    </HStack>
  )
}
