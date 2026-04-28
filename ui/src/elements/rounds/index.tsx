import { useMemo, useState } from 'react'
import { Box, chakra, HStack, Stack } from '@chakra-ui/react'
import { useRecentRounds } from '~queries/rounds'
import { RoundList } from '~components/Round/RoundList'
import { QueryDataLayout } from '~components/Layout/QueryDataLayout'
import { PageHeader } from '~components/Layout/PageHeader'
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

const filterOrder: Filter[] = [
  'all',
  'registration',
  'contribution',
  'finalized',
  'completed',
  'aborted',
]

// Rounds list page. Editorial: eyebrow rule, large display heading, mono
// caption, then a row of filter chips before the table. Filter chips are
// pill-shaped with a hairline border and a tiny indicator dot when active —
// matches the StatusBadge vocabulary.
export function RoundsList() {
  // The contract caps the recent-rounds ring buffer at 64; querying more
  // than that just yields the same window.
  const recent = useRecentRounds(64)
  const [filter, setFilter] = useState<Filter>('all')

  const filtered = useMemo(() => {
    if (!recent.data) return []
    if (filter === 'all') return recent.data
    return recent.data.filter((r) => roundPhase(r.round) === filter)
  }, [recent.data, filter])

  const counts = useMemo(() => {
    const map: Partial<Record<Filter, number>> = { all: recent.data?.length ?? 0 }
    if (recent.data) {
      for (const r of recent.data) {
        const p = roundPhase(r.round)
        map[p] = (map[p] ?? 0) + 1
      }
    }
    return map
  }, [recent.data])

  return (
    <Stack gap={{ base: 8, md: 10 }}>
      <PageHeader
        title='Rounds'
        subtitle={`The most recent ${recent.data?.length ?? '…'} rounds from the contract's ring buffer. Older rounds remain valid on-chain but are not enumerated here.`}
      />

      <HStack gap={1.5} wrap='wrap'>
        {filterOrder.map((f) => (
          <FilterChip
            key={f}
            isActive={filter === f}
            onClick={() => setFilter(f)}
            label={filterLabels[f]}
            count={counts[f] ?? 0}
          />
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

const ChipBtn = chakra('button', {
  base: {
    display: 'inline-flex',
    alignItems: 'center',
    gap: 2,
    px: 3,
    py: 1.5,
    borderRadius: 'full',
    borderWidth: '1px',
    fontFamily: 'sans',
    fontSize: 'xs',
    cursor: 'pointer',
    transition: 'background 0.12s, border-color 0.12s, color 0.12s',
  },
})

function FilterChip({
  isActive,
  onClick,
  label,
  count,
}: {
  isActive: boolean
  onClick: () => void
  label: string
  count: number
}) {
  return (
    <ChipBtn
      type='button'
      onClick={onClick}
      borderColor={isActive ? 'accent.border' : 'border.subtle'}
      bg={isActive ? 'accent.bg.strong' : 'transparent'}
      color={isActive ? 'accent.bright' : 'ink.2'}
      _hover={{
        borderColor: isActive ? 'accent.border' : 'border',
        color: isActive ? 'accent.bright' : 'ink.0',
      }}
    >
      <Box>{label}</Box>
      <Box
        as='span'
        className='dkg-tabular'
        fontFamily='mono'
        fontSize='2xs'
        color={isActive ? 'accent.fg' : 'ink.4'}
      >
        {count}
      </Box>
    </ChipBtn>
  )
}
