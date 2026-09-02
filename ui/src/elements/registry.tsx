import { useMemo } from 'react'
import { Heading, SimpleGrid, Stack } from '@chakra-ui/react'
import { useRegistryNodes, useRegistryStats } from '~queries/registry'
import { useOperatorStats } from '~queries/operators'
import { useBlockNumber } from '~queries/chain'
import { StatCard } from '~components/ui/StatCard'
import { OperatorLeaderboard } from '~components/Registry/OperatorLeaderboard'
import { OperatorActivityChart } from '~components/Charts/OperatorActivityChart'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { PageHeader } from '~components/Layout/PageHeader'
import { blocksToDuration } from '~lib/format'
import { formatParticipation, summarizeOperatorStats } from '~lib/operator-stats'

export function Registry() {
  const stats = useRegistryStats()
  const nodes = useRegistryNodes()
  const operators = useOperatorStats()
  const { data: block } = useBlockNumber()

  const summary = useMemo(
    () => summarizeOperatorStats(operators.data ?? []),
    [operators.data],
  )

  return (
    <Stack gap={{ base: 10, md: 12 }}>
      <PageHeader
        title='Node registry'
        subtitle='Operators registered on the DKG registry contract. The lottery only picks committee members from active nodes; inactive ones are pruned automatically by the on-chain inactivity window. Everything below is counted from the contract event log since the deployment block.'
      />

      <QueryDataLayout isLoading={stats.isLoading} isError={stats.isError} error={stats.error}>
        {stats.data && (
          <SimpleGrid columns={{ base: 2, md: 4 }} gap={{ base: 3, md: 4 }}>
            <StatCard
              label='Active'
              value={stats.data.active.toString()}
              hint={`${stats.data.total.toString()} ever registered`}
              tone='live'
            />
            <StatCard
              label='Contributions'
              value={operators.data ? summary.contributions.toString() : '—'}
              hint={operators.data ? `${summary.claims} slots claimed` : 'counting…'}
            />
            <StatCard
              label='Partial decryptions'
              value={operators.data ? summary.partials.toString() : '—'}
              hint={
                operators.data
                  ? `${summary.finalizations} finalized · ${summary.combines} combined`
                  : 'counting…'
              }
            />
            <StatCard
              label='Participation'
              value={operators.data ? formatParticipation(summary.participation) : '—'}
              hint='contributions ÷ claims, all operators'
              tone='accent'
            />
          </SimpleGrid>
        )}
      </QueryDataLayout>

      <OperatorActivityChart rows={operators.data ?? []} loading={operators.isLoading} />

      <Stack gap={5}>
        <Heading
          as='h2'
          fontSize={{ base: 'lg', md: 'xl' }}
          fontWeight={500}
          color='ink.0'
          letterSpacing='-0.01em'
        >
          Leaderboard
        </Heading>
        <QueryDataLayout
          isLoading={operators.isLoading}
          isError={operators.isError}
          error={operators.error}
          isEmpty={operators.data?.length === 0}
          emptyMessage='No nodes have ever registered against this registry.'
        >
          {operators.data && (
            <OperatorLeaderboard
              rows={operators.data}
              nodes={nodes.data ?? []}
              currentBlock={block ?? null}
            />
          )}
        </QueryDataLayout>
        {stats.data && (
          <SimpleGrid columns={{ base: 1, md: 2 }} gap={3}>
            <StatCard
              label='Inactivity window'
              value={blocksToDuration(Number(stats.data.inactivity))}
              hint={`${stats.data.inactivity.toString()} blocks without a heartbeat before a node can be reaped`}
            />
            <StatCard
              label='Latest block'
              value={block ? `#${block.toString()}` : '—'}
              hint='chain head'
              tone='live'
            />
          </SimpleGrid>
        )}
      </Stack>
    </Stack>
  )
}
