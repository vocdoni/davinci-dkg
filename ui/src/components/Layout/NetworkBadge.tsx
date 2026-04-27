import { Badge, HStack, Text } from '@chakra-ui/react'
import { useConfig } from '~providers/ConfigProvider'
import { useBlockNumber } from '~queries/chain'

// Compact "chain · #block" indicator for the header. Polls the block number
// at the cadence defined by Polling.blockNumber so the user always has a
// rough sense that the RPC is alive.
export function NetworkBadge() {
  const config = useConfig()
  const { data: block } = useBlockNumber()
  return (
    <HStack gap={2}>
      <Badge colorPalette='cyan' variant='subtle' textTransform='lowercase'>
        {config.chainName}
      </Badge>
      <Text fontFamily='mono' fontSize='xs' color='gray.400'>
        {block ? `#${block.toString()}` : '—'}
      </Text>
    </HStack>
  )
}
