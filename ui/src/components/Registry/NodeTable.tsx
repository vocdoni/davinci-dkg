import { Badge, Box, HStack, Stack, Table, Text } from '@chakra-ui/react'
import type { NodeKey } from '@vocdoni/davinci-dkg-sdk'
import { NodeStatus } from '@vocdoni/davinci-dkg-sdk'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { bigIntToHex, blocksToDuration } from '~lib/format'

interface Props {
  nodes: NodeKey[]
  currentBlock: bigint | null
}

function nodeStatusBadge(status: number) {
  if (status === NodeStatus.Active) return <Badge colorPalette='green'>Active</Badge>
  if (status === NodeStatus.Inactive) return <Badge colorPalette='gray'>Inactive</Badge>
  return <Badge colorPalette='red'>Unknown</Badge>
}

export function NodeTable({ nodes, currentBlock }: Props) {
  if (nodes.length === 0) {
    return (
      <Box p={6} textAlign='center' color='gray.500' fontSize='sm'>
        No nodes registered.
      </Box>
    )
  }

  return (
    <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' overflow='hidden'>
      <Table.Root size='sm'>
        <Table.Header>
          <Table.Row bg='gray.900'>
            <Table.ColumnHeader>Operator</Table.ColumnHeader>
            <Table.ColumnHeader>Status</Table.ColumnHeader>
            <Table.ColumnHeader>Last seen</Table.ColumnHeader>
            <Table.ColumnHeader>Public key</Table.ColumnHeader>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {nodes.map((node) => {
            const lastSeenBlocks = currentBlock != null ? Number(currentBlock - node.lastActiveBlock) : null
            return (
              <Table.Row key={node.operator} _hover={{ bg: 'gray.800' }}>
                <Table.Cell>
                  <HashCell value={node.operator} head={6} tail={6} />
                </Table.Cell>
                <Table.Cell>{nodeStatusBadge(node.status)}</Table.Cell>
                <Table.Cell>
                  <Stack gap={0}>
                    <Text fontSize='xs' color='gray.300'>
                      {lastSeenBlocks != null ? blocksToDuration(lastSeenBlocks) + ' ago' : '—'}
                    </Text>
                    <Text fontSize='2xs' color='gray.500' fontFamily='mono'>
                      block #{node.lastActiveBlock.toString()}
                    </Text>
                  </Stack>
                </Table.Cell>
                <Table.Cell>
                  <DetailDisclosure title='Show key coordinates'>
                    <Stack gap={0.5}>
                      <HStack gap={2}>
                        <Text fontSize='2xs' color='gray.500'>x:</Text>
                        <HashCell value={bigIntToHex(node.pubX)} head={6} tail={6} />
                      </HStack>
                      <HStack gap={2}>
                        <Text fontSize='2xs' color='gray.500'>y:</Text>
                        <HashCell value={bigIntToHex(node.pubY)} head={6} tail={6} />
                      </HStack>
                    </Stack>
                  </DetailDisclosure>
                </Table.Cell>
              </Table.Row>
            )
          })}
        </Table.Body>
      </Table.Root>
    </Box>
  )
}
