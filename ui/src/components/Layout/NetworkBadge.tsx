import { Box, HStack, Text } from '@chakra-ui/react'
import { useConfig } from '~providers/ConfigProvider'
import { useBlockNumber } from '~queries/chain'

// Compact "● <chain> · #<block>" indicator. The pulsing phosphor dot is a
// liveness signal — if the block number stops ticking, the dot stays but the
// number falls behind, and the operator notices instantly.
//
// The `dkg-tabular` class is load-bearing here: without it the height of
// the digit row jumps every time a thinner glyph (1) replaces a wider one
// (5), which makes the header feel jittery on every poll.
export function NetworkBadge() {
  const config = useConfig()
  const { data: block, isFetching } = useBlockNumber()
  return (
    <HStack
      gap={2.5}
      px={2.5}
      py={1.5}
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='full'
      bg='surface.sunken'
    >
      <Box position='relative' w='6px' h='6px'>
        <Box
          position='absolute'
          inset={0}
          borderRadius='full'
          bg={block != null ? 'live.fg' : 'ink.4'}
        />
        {block != null && (
          <Box
            position='absolute'
            inset={0}
            borderRadius='full'
            bg='live.fg'
            opacity={isFetching ? 0.7 : 0}
            transform='scale(2.2)'
            transition='opacity 0.4s ease'
          />
        )}
      </Box>
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.2'
        letterSpacing='0.04em'
        textTransform='lowercase'
      >
        {config.chainName}
      </Text>
      <Box w='1px' h='10px' bg='border' />
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
      >
        {block ? `#${block.toString()}` : '—'}
      </Text>
    </HStack>
  )
}
