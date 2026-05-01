import { useQuery } from '@tanstack/react-query'
import { Box, HStack, Text, Tooltip } from '@chakra-ui/react'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useEstimatedBlock } from '~hooks/use-estimated-block'
import { useBlockTimeSeconds } from '~hooks/use-block-time'

// Refresh nextEpochStartBlock + epochDurationBlocks every 30s. Both move
// only on createEpoch, so this is plenty frequent — the wall-clock display
// is interpolated client-side from useEstimatedBlock + useBlockTimeSeconds.
const STALE_TIME_MS = 30_000

function fmtMs(ms: number): string {
  if (ms <= 0) return '0s'
  const totalS = Math.floor(ms / 1000)
  const m = Math.floor(totalS / 60)
  const s = totalS % 60
  if (m === 0) return `${s}s`
  return `${m}m ${s.toString().padStart(2, '0')}s`
}

/**
 * Header pill that shows time + blocks remaining until the next epoch's
 * `createEpoch` is allowed by the contract's cadence guard. When the
 * threshold is reached the pill flips to a "ready" indicator (the
 * dkg-node binaries auto-fire createEpoch with jitter, so the user
 * usually doesn't need to do anything; the pill just signals readiness).
 */
export function NextEpochCountdown({ compact = false }: { compact?: boolean }) {
  const client = useDkgClient()
  const blockTimeS = useBlockTimeSeconds()
  const liveBlock = useEstimatedBlock(blockTimeS)

  const { data } = useQuery({
    queryKey: ['nextEpochStartBlock', client?.dkg.managerAddress],
    enabled: !!client,
    refetchInterval: STALE_TIME_MS,
    staleTime: STALE_TIME_MS,
    queryFn: async () => {
      if (!client) return null
      const [next, dur] = await Promise.all([
        client.dkg.getNextEpochStartBlock(),
        client.dkg.getEpochDurationBlocks(),
      ])
      return { next, dur }
    },
  })

  if (!client || !data || liveBlock == null) return null
  const { next, dur } = data
  const remainingBlocks = next > liveBlock ? next - liveBlock : 0n
  const remainingMs = Number(remainingBlocks) * blockTimeS * 1000

  const ready = remainingBlocks === 0n
  const label = ready ? 'Next epoch ready' : `Next epoch in ${fmtMs(remainingMs)}`
  const tooltip = ready
    ? `Cadence threshold reached — any node may now fire createEpoch. ` +
      `Epoch length: ${dur.toString()} blocks ` +
      `(~${Math.round(Number(dur) * blockTimeS)}s @ ${blockTimeS.toFixed(1)}s/block).`
    : `Cadence anchor: block ${next.toString()}. ` +
      `Currently at ~${liveBlock.toString()}. ` +
      `${remainingBlocks.toString()} blocks remaining @ ${blockTimeS.toFixed(1)}s/block.`

  return (
    <Tooltip.Root openDelay={300}>
      <Tooltip.Trigger asChild>
        <HStack
          gap={2}
          px={3}
          py={1.5}
          borderRadius='md'
          borderWidth='1px'
          borderColor='rule'
          bg={ready ? 'live.bg' : 'transparent'}
          fontFamily='mono'
          fontSize='xs'
          color={ready ? 'live.fg' : 'ink.3'}
          cursor='help'
          whiteSpace='nowrap'
        >
          <Box w='5px' h='5px' borderRadius='full' bg={ready ? 'live.fg' : 'ink.4'} />
          {compact ? (
            <Text>{ready ? '⏵ ready' : fmtMs(remainingMs)}</Text>
          ) : (
            <Text>{label}</Text>
          )}
        </HStack>
      </Tooltip.Trigger>
      <Tooltip.Positioner>
        <Tooltip.Content maxW='320px' fontSize='xs'>
          {tooltip}
        </Tooltip.Content>
      </Tooltip.Positioner>
    </Tooltip.Root>
  )
}
