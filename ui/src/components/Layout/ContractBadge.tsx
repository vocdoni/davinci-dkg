import { Box, HStack, Link, Text } from '@chakra-ui/react'
import { LuExternalLink } from 'react-icons/lu'
import { useConfig } from '~providers/ConfigProvider'
import { explorerAddressUrl } from '~lib/explorer'
import { shortHash } from '~lib/format'

// Compact "manager 0xd3ef…298a ↗" pill linking to the block explorer.
// Mirrors NetworkBadge in shape so the two sit comfortably next to each
// other in the header. When no explorerUrl is configured the pill renders
// as plain text (no link) — the address is still informative.
export function ContractBadge() {
  const { managerAddress, explorerUrl } = useConfig()
  const href = explorerAddressUrl(explorerUrl, managerAddress)
  const short = shortHash(managerAddress, 6, 4)

  const inner = (
    <HStack gap={2} align='center'>
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.06em'
        textTransform='uppercase'
      >
        contract
      </Text>
      <Box w='1px' h='10px' bg='border' />
      <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.1'>
        {short}
      </Text>
      {href && (
        <Box color='ink.3' display='flex' alignItems='center'>
          <LuExternalLink size={10} />
        </Box>
      )}
    </HStack>
  )

  const pillProps = {
    px: 2.5,
    py: 1.5,
    borderWidth: '1px',
    borderColor: 'border.subtle',
    borderRadius: 'full',
    bg: 'surface.sunken',
  } as const

  if (!href) {
    return <Box {...pillProps}>{inner}</Box>
  }

  return (
    <Link
      href={href}
      target='_blank'
      rel='noopener noreferrer'
      aria-label={`DKG manager contract ${managerAddress} on block explorer`}
      _hover={{ borderColor: 'accent.border', textDecoration: 'none' }}
      transition='border-color 0.15s ease'
      {...pillProps}
    >
      {inner}
    </Link>
  )
}
