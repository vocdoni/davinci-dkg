import { Box, HStack, Stack, Table, Text } from '@chakra-ui/react'
import type { NodeKey } from '@vocdoni/davinci-dkg-sdk'
import { NodeStatus } from '@vocdoni/davinci-dkg-sdk'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { bigIntToHex, blocksToDuration } from '~lib/format'

interface Props {
  nodes: NodeKey[]
  currentBlock: bigint | null
}

// Editorial node status pill matching the StatusBadge vocabulary.
function NodeStatusPill({ status }: { status: number }) {
  const tone =
    status === NodeStatus.Active
      ? { fg: 'live.fg', bg: 'live.bg', border: 'rgba(134, 239, 172, 0.30)', dot: 'live.fg', label: 'Active' }
      : status === NodeStatus.Inactive
        ? { fg: 'ink.3', bg: 'surface.raised', border: 'border', dot: 'ink.4', label: 'Inactive' }
        : { fg: 'danger.fg', bg: 'danger.bg', border: 'danger.border', dot: 'danger.fg', label: 'Unknown' }
  return (
    <HStack
      as='span'
      display='inline-flex'
      gap={2}
      px={2.5}
      py='4px'
      borderRadius='full'
      borderWidth='1px'
      borderColor={tone.border}
      bg={tone.bg}
    >
      <Box w='6px' h='6px' borderRadius='full' bg={tone.dot} />
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color={tone.fg}
        letterSpacing='0.08em'
        textTransform='uppercase'
        lineHeight='1'
      >
        {tone.label}
      </Text>
    </HStack>
  )
}

export function NodeTable({ nodes, currentBlock }: Props) {
  if (nodes.length === 0) {
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
          No nodes registered.
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
            <Th>Operator</Th>
            <Th>Status</Th>
            <Th>Last seen</Th>
            <Th>Public key</Th>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {nodes.map((node, i) => {
            const lastSeenBlocks =
              currentBlock != null ? Number(currentBlock - node.lastActiveBlock) : null
            return (
              <Table.Row
                key={node.operator}
                borderTopWidth={i === 0 ? 0 : '1px'}
                borderColor='rule'
                transition='background 0.12s'
                _hover={{ bg: 'accent.bg' }}
              >
                <Table.Cell py={3.5}>
                  <HashCell value={node.operator} head={6} tail={6} />
                </Table.Cell>
                <Table.Cell py={3.5}>
                  <NodeStatusPill status={node.status} />
                </Table.Cell>
                <Table.Cell py={3.5}>
                  <Stack gap={0}>
                    <Text fontSize='xs' color='ink.1'>
                      {lastSeenBlocks != null ? blocksToDuration(lastSeenBlocks) + ' ago' : '—'}
                    </Text>
                    <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.4'>
                      block #{node.lastActiveBlock.toString()}
                    </Text>
                  </Stack>
                </Table.Cell>
                <Table.Cell py={3.5}>
                  <DetailDisclosure title='Show key coordinates'>
                    <Stack gap={1}>
                      <HStack gap={3}>
                        <Text fontFamily='mono' fontSize='2xs' color='accent.dim' minW='10px'>
                          x
                        </Text>
                        <HashCell value={bigIntToHex(node.pubX)} head={6} tail={6} />
                      </HStack>
                      <HStack gap={3}>
                        <Text fontFamily='mono' fontSize='2xs' color='accent.dim' minW='10px'>
                          y
                        </Text>
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
