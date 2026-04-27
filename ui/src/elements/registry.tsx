import { Box, Heading, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import { useRegistryNodes, useRegistryStats } from '~queries/registry'
import { useBlockNumber } from '~queries/chain'
import { StatCard } from '~components/ui/StatCard'
import { NodeTable } from '~components/Registry/NodeTable'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { blocksToDuration } from '~lib/format'

export function Registry() {
  const stats = useRegistryStats()
  const nodes = useRegistryNodes()
  const { data: block } = useBlockNumber()

  return (
    <Stack gap={8}>
      <Box>
        <Heading size='lg'>Node registry</Heading>
        <Text color='gray.400' fontSize='sm' mt={1}>
          Operators registered on the DKG registry contract. The lottery only ever picks committee
          members from active nodes; inactive ones are pruned automatically by the on-chain
          inactivity window.
        </Text>
      </Box>

      <QueryDataLayout isLoading={stats.isLoading} isError={stats.isError} error={stats.error}>
        {stats.data && (
          <SimpleGrid columns={{ base: 2, md: 4 }} gap={4}>
            <StatCard label='Active' value={stats.data.active.toString()} />
            <StatCard label='Total registered' value={stats.data.total.toString()} />
            <StatCard
              label='Inactivity window'
              value={blocksToDuration(Number(stats.data.inactivity))}
              hint={`${stats.data.inactivity.toString()} blocks`}
            />
            <StatCard label='Latest block' value={block ? `#${block.toString()}` : '—'} />
          </SimpleGrid>
        )}
      </QueryDataLayout>

      <Box>
        <Heading size='sm' mb={3}>
          Nodes
        </Heading>
        <QueryDataLayout
          isLoading={nodes.isLoading}
          isError={nodes.isError}
          error={nodes.error}
          isEmpty={nodes.data?.length === 0}
          emptyMessage='No nodes have ever registered against this registry.'
        >
          {nodes.data && <NodeTable nodes={nodes.data} currentBlock={block ?? null} />}
        </QueryDataLayout>
      </Box>
    </Stack>
  )
}
