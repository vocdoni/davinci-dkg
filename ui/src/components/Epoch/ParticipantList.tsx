import { Box, Grid, GridItem, Stack, Text } from '@chakra-ui/react'
import type { Address } from 'viem'
import { HashCell } from '~components/ui/HashCell'

// Editorial roster: numbered list of committee members, each row a
// hairline-divided strip with a mono index column and the address.
export function ParticipantList({ participants }: { participants: Address[] }) {
  if (participants.length === 0) {
    return (
      <Box
        borderWidth='1px'
        borderColor='border.subtle'
        borderRadius='lg'
        bg='surface'
        p={{ base: 8, md: 12 }}
        textAlign='center'
      >
        <Text color='ink.3' fontSize='sm'>
          No committee members selected yet.
        </Text>
      </Box>
    )
  }
  return (
    <Stack gap={0}>
      {participants.map((addr, i) => (
        <Grid
          key={addr}
          templateColumns='52px 1fr'
          gap={4}
          py={3}
          alignItems='center'
          borderTopWidth={i === 0 ? '1px' : 0}
          borderBottomWidth='1px'
          borderColor='rule'
          transition='background 0.12s'
          _hover={{ bg: 'accent.bg' }}
        >
          <GridItem>
            <Text
              className='dkg-tabular'
              fontFamily='mono'
              fontSize='xs'
              color='accent.dim'
              letterSpacing='0.04em'
              textAlign='right'
            >
              {(i + 1).toString().padStart(2, '0')}
            </Text>
          </GridItem>
          <GridItem>
            <HashCell value={addr} head={8} tail={8} />
          </GridItem>
        </Grid>
      ))}
    </Stack>
  )
}
