import { useMemo, useState } from 'react'
import { Box, Heading, HStack, Stack, Text } from '@chakra-ui/react'
import type { ButtonProps } from '@chakra-ui/react'
import { Button } from '@chakra-ui/react'
import { useRecentRounds } from '~queries/rounds'
import { RoundList } from '~components/Round/RoundList'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { roundPhase } from '~lib/round-utils'
import type { RoundPhase } from '~lib/round-utils'

type Filter = 'all' | RoundPhase

const filterLabels: Record<Filter, string> = {
  all: 'All',
  registration: 'Registration',
  contribution: 'Contribution',
  finalized: 'Finalized',
  completed: 'Completed',
  aborted: 'Aborted',
  unknown: 'Unknown',
}

const filterOrder: Filter[] = ['all', 'registration', 'contribution', 'finalized', 'completed', 'aborted']

export function RoundsList() {
  // The contract caps the recent-rounds ring buffer at 64; querying more
  // than that just yields the same window. Phase 2 keeps the simple
  // single-page view; pagination lands in Phase 4 once we have a real
  // dataset to test against.
  const recent = useRecentRounds(64)
  const [filter, setFilter] = useState<Filter>('all')

  const filtered = useMemo(() => {
    if (!recent.data) return []
    if (filter === 'all') return recent.data
    return recent.data.filter((r) => roundPhase(r.round) === filter)
  }, [recent.data, filter])

  return (
    <Stack gap={6}>
      <Box>
        <Heading size='lg'>Rounds</Heading>
        <Text color='gray.400' fontSize='sm' mt={1}>
          The most recent {recent.data?.length ?? '…'} rounds from the contract's ring buffer.
        </Text>
      </Box>

      <HStack gap={2} wrap='wrap'>
        {filterOrder.map((f) => (
          <FilterChip
            key={f}
            isActive={filter === f}
            onClick={() => setFilter(f)}
          >
            {filterLabels[f]}
          </FilterChip>
        ))}
      </HStack>

      <QueryDataLayout
        isLoading={recent.isLoading}
        isError={recent.isError}
        error={recent.error}
        isEmpty={filtered.length === 0}
        emptyMessage={
          recent.data?.length === 0
            ? 'No rounds have been created yet on this network.'
            : 'No rounds match this filter.'
        }
      >
        <RoundList rounds={filtered} />
      </QueryDataLayout>
    </Stack>
  )
}

function FilterChip({ isActive, ...props }: ButtonProps & { isActive: boolean }) {
  return (
    <Button
      size='xs'
      variant={isActive ? 'solid' : 'outline'}
      colorPalette={isActive ? 'cyan' : 'gray'}
      {...props}
    />
  )
}
