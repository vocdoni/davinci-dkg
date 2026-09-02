import { Box, HStack, Stack, Table, Text } from '@chakra-ui/react'
import type { NodeKey } from '@vocdoni/davinci-dkg-sdk'
import { NodeStatus } from '@vocdoni/davinci-dkg-sdk'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { bigIntToHex, blocksToDuration } from '~lib/format'
import { formatParticipation, type OperatorStats } from '~lib/operator-stats'
import { ChartColors } from '~components/Charts/palette'

interface Props {
  rows: OperatorStats[]
  /** Registry records, for the BabyJubJub key of each operator. */
  nodes: NodeKey[]
  currentBlock: bigint | null
}

/**
 * The operator leaderboard: who claimed, who contributed, who decrypted, who
 * finalized, and how reliably. Sorted by the aggregation (contributions first);
 * the participation score is defined in `lib/operator-stats.ts` and explained
 * in the footnote below the table so the column header can stay one word.
 */
export function OperatorLeaderboard({ rows, nodes, currentBlock }: Props) {
  const keyOf = new Map(nodes.map((n) => [n.operator.toLowerCase(), n]))

  if (rows.length === 0) {
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
          No operators registered.
        </Text>
      </Box>
    )
  }

  return (
    <Stack gap={3}>
      <Box
        borderWidth='1px'
        borderColor='border.subtle'
        borderRadius='lg'
        bg='surface'
        overflowX='auto'
        boxShadow='inset'
      >
        <Table.Root size='sm' interactive>
          <Table.Header>
            <Table.Row bg='surface.sunken'>
              <Th>Operator</Th>
              <Th>Status</Th>
              <Th align='right'>Claims</Th>
              <Th align='right' dot={ChartColors.contributions}>
                Contributions
              </Th>
              <Th align='right' dot={ChartColors.partials}>
                Partials
              </Th>
              <Th align='right'>Finalized</Th>
              <Th align='right'>Combined</Th>
              <Th align='right'>Participation</Th>
              <Th>Last seen</Th>
              <Th>Public key</Th>
            </Table.Row>
          </Table.Header>
          <Table.Body>
            {rows.map((row, i) => {
              const node = keyOf.get(row.operator.toLowerCase())
              const lastSeenBlocks =
                currentBlock != null && row.lastActiveBlock != null
                  ? Number(currentBlock - row.lastActiveBlock)
                  : null
              return (
                <Table.Row
                  key={row.operator}
                  borderTopWidth={i === 0 ? 0 : '1px'}
                  borderColor='rule'
                  transition='background 0.12s'
                  _hover={{ bg: 'accent.bg' }}
                >
                  <Table.Cell py={3.5}>
                    <HashCell value={row.operator} head={6} tail={6} />
                  </Table.Cell>
                  <Table.Cell py={3.5}>
                    <StatusPill status={row.status} />
                  </Table.Cell>
                  <Num value={row.claims} />
                  <Num value={row.contributions} strong />
                  <Num value={row.partials} strong />
                  <Num value={row.finalizations} />
                  <Num value={row.combines} />
                  <Table.Cell py={3.5} textAlign='right'>
                    <Text
                      className='dkg-tabular'
                      fontFamily='mono'
                      fontSize='xs'
                      color={scoreColor(row.participation)}
                    >
                      {formatParticipation(row.participation)}
                    </Text>
                  </Table.Cell>
                  <Table.Cell py={3.5}>
                    <Stack gap={0}>
                      <Text fontSize='xs' color='ink.1'>
                        {lastSeenBlocks != null ? `${blocksToDuration(lastSeenBlocks)} ago` : '—'}
                      </Text>
                      <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.4'>
                        {row.lastActiveBlock != null
                          ? `block #${row.lastActiveBlock.toString()}`
                          : 'not in registry'}
                      </Text>
                    </Stack>
                  </Table.Cell>
                  <Table.Cell py={3.5}>
                    {node ? (
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
                          <Text fontFamily='mono' fontSize='2xs' color='ink.4'>
                            registered at block #{node.registeredAtBlock.toString()}
                          </Text>
                        </Stack>
                      </DetailDisclosure>
                    ) : (
                      <Text fontSize='2xs' color='ink.4'>
                        reaped
                      </Text>
                    )}
                  </Table.Cell>
                </Table.Row>
              )
            })}
          </Table.Body>
        </Table.Root>
      </Box>
      <Text fontSize='2xs' color='ink.4' lineHeight='1.6' maxW='80ch'>
        <strong>Participation</strong> = accepted contributions ÷ slots claimed, in percent.
        Winning a lottery slot commits an operator to publishing a contribution in the
        key-assembly window, so the ratio is how often it followed through. Operators that have
        never won a slot show <code>—</code> rather than 100%: no track record is not a perfect
        one. Counts cover every epoch since the manager's deployment block. Finalizations and
        combines are attributed to the transaction sender — those two events do not name a
        submitter. Partial counts differ between operators by design: members answer each
        ciphertext on a seed-derived stagger, so roughly <code>t</code> of the <code>n</code>
        end up spending the gas and which ones varies per ciphertext.
      </Text>
    </Stack>
  )
}

function scoreColor(score: number | null): string {
  if (score == null) return 'ink.4'
  if (score >= 90) return 'live.fg'
  if (score >= 60) return 'ink.1'
  return 'warn.fg'
}

function StatusPill({ status }: { status: number }) {
  const tone =
    status === NodeStatus.Active
      ? { fg: 'live.fg', bg: 'live.bg', border: 'rgba(134, 239, 172, 0.30)', label: 'Active' }
      : status === NodeStatus.Inactive
        ? { fg: 'ink.3', bg: 'surface.raised', border: 'border', label: 'Inactive' }
        : { fg: 'ink.4', bg: 'transparent', border: 'border.subtle', label: 'Unregistered' }
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
      <Box w='6px' h='6px' borderRadius='full' bg={tone.fg} />
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

function Num({ value, strong }: { value: number; strong?: boolean }) {
  return (
    <Table.Cell py={3.5} textAlign='right'>
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize='xs'
        color={value === 0 ? 'ink.4' : strong ? 'ink.0' : 'ink.1'}
      >
        {value}
      </Text>
    </Table.Cell>
  )
}

function Th({
  children,
  align,
  dot,
}: {
  children: React.ReactNode
  align?: 'right'
  dot?: string
}) {
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
      textAlign={align}
      whiteSpace='nowrap'
    >
      {dot ? (
        <HStack gap={1.5} justify={align === 'right' ? 'flex-end' : 'flex-start'}>
          {/* Matches the bar colour in the chart above, so the two read as one view. */}
          <Box w='7px' h='7px' borderRadius='2px' bg={dot} />
          <Box as='span'>{children}</Box>
        </HStack>
      ) : (
        children
      )}
    </Table.ColumnHeader>
  )
}
