import { Badge, Box, HStack, Stack, Text } from '@chakra-ui/react'
import type { Address } from 'viem'
import { HashCell } from '~components/ui/HashCell'

export function ParticipantList({ participants }: { participants: Address[] }) {
  if (participants.length === 0) {
    return (
      <Text fontSize='sm' color='gray.500'>
        No committee members selected yet.
      </Text>
    )
  }
  return (
    <Stack gap={2}>
      {participants.map((addr, i) => (
        <Box key={addr} borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' px={3} py={2}>
          <HStack gap={3}>
            <Badge colorPalette='gray' fontFamily='mono' fontSize='2xs'>
              #{i + 1}
            </Badge>
            <HashCell value={addr} head={6} tail={6} />
          </HStack>
        </Box>
      ))}
    </Stack>
  )
}
