import { Heading, SimpleGrid, Stack } from '@chakra-ui/react'
import { useRegistryNodes, useRegistryStats } from '~queries/registry'
import { useBlockNumber } from '~queries/chain'
import { StatCard } from '~components/ui/StatCard'
import { NodeTable } from '~components/Registry/NodeTable'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { PageHeader } from '~components/Layout/PageHeader'
import { blocksToDuration } from '~lib/format'

export function Registry() {
  const stats = useRegistryStats()
  const nodes = useRegistryNodes()
  const { data: block } = useBlockNumber()

  return (
    <Stack gap={{ base: 10, md: 12 }}>
      <PageHeader
        title='Node registry'
        subtitle='Operators registered on the DKG registry contract. The lottery only picks committee members from active nodes; inactive ones are pruned automatically by the on-chain inactivity window.'
      />

      <QueryDataLayout isLoading={stats.isLoading} isError={stats.isError} error={stats.error}>
        {stats.data && (
          <SimpleGrid columns={{ base: 2, md: 4 }} gap={{ base: 3, md: 4 }}>
            <StatCard label='Active' value={stats.data.active.toString()} tone='live' />
            <StatCard label='Total registered' value={stats.data.total.toString()} />
            <StatCard
              label='Inactivity window'
              value={blocksToDuration(Number(stats.data.inactivity))}
              hint={`${stats.data.inactivity.toString()} blocks`}
            />
            <StatCard label='Latest block' value={block ? `#${block.toString()}` : '—'} tone='live' />
          </SimpleGrid>
        )}
      </QueryDataLayout>

      <Stack gap={5}>
        <Heading
          as='h2'
          fontSize={{ base: 'lg', md: 'xl' }}
          fontWeight={500}
          color='ink.0'
          letterSpacing='-0.01em'
        >
          Roster
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
      </Stack>
    </Stack>
  )
}
