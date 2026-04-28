import { Box, HStack, Link, Text } from '@chakra-ui/react'
import { LuExternalLink } from 'react-icons/lu'
import { useConfig } from '~providers/ConfigProvider'
import { useBlockNumber } from '~queries/chain'
import { explorerAddressUrl } from '~lib/explorer'
import { shortHash } from '~lib/format'

interface Props {
  /**
   * `compact` drops the contract segment and shows just chain + block.
   * Use it on viewports where horizontal space is at a premium (lg, but
   * not yet xl).
   */
  compact?: boolean
}

// One unified header pill that fuses the previous NetworkBadge +
// ContractBadge into a single shape:
//
//   ●  sepolia  ·  #1234  │  contract 0xd3ef…298a ↗
//
// The chain dot pulses on each fetch (liveness signal). The contract
// segment is a real anchor — clicking it jumps to the configured block
// explorer in a new tab. When `explorerUrl` is unset the segment renders
// as plain text.
export function ChainPill({ compact = false }: Props) {
  const { chainName, managerAddress, explorerUrl } = useConfig()
  const { data: block, isFetching } = useBlockNumber()
  const explorerHref = explorerAddressUrl(explorerUrl, managerAddress)

  return (
    <HStack
      gap={2.5}
      px={2.5}
      py={1.5}
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='full'
      bg='surface.sunken'
      transition='border-color 0.15s ease'
      _hover={{ borderColor: 'border' }}
    >
      {/* Liveness dot. The pulsing halo fades on every poll, which is the
          cheapest way to make "the chain is alive" obvious at a glance. */}
      <Box position='relative' w='6px' h='6px' flexShrink={0}>
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
        {chainName}
      </Text>

      <Box w='1px' h='10px' bg='border' />

      <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.3'>
        {block ? `#${block.toString()}` : '—'}
      </Text>

      {!compact && (
        <>
          <Box w='1px' h='10px' bg='border' />
          {/* Wrap only the contract segment in a Link so the chain/block
              part doesn't behave like an anchor when there's no explorer.
              The address speaks for itself — no "CONTRACT" label needed. */}
          {explorerHref ? (
            <Link
              href={explorerHref}
              target='_blank'
              rel='noopener noreferrer'
              aria-label={`DKG manager contract ${managerAddress} on block explorer`}
              _hover={{ textDecoration: 'none', color: 'accent.fg' }}
              transition='color 0.15s ease'
              color='ink.2'
            >
              <HStack gap={1.5}>
                <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='inherit'>
                  {shortHash(managerAddress, 4, 4)}
                </Text>
                <Box display='flex' alignItems='center' color='ink.4'>
                  <LuExternalLink size={10} />
                </Box>
              </HStack>
            </Link>
          ) : (
            <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.2'>
              {shortHash(managerAddress, 4, 4)}
            </Text>
          )}
        </>
      )}
    </HStack>
  )
}
