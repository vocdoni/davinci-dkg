import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import type { RoundEvent } from '@vocdoni/davinci-dkg-sdk'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { RawJson } from '~components/Debug/RawJson'

// Plain-English event summaries with the raw `args` blob hidden behind a
// disclosure for debug mode.
function eventSummary(ev: RoundEvent): string {
  switch (ev.eventName) {
    case 'RoundCreated':
      return 'Round created on-chain.'
    case 'SeedRevealed':
      return 'Lottery seed revealed.'
    case 'SlotClaimed':
      return 'Committee member claimed a slot.'
    case 'ContributionSubmitted':
      return 'Contribution submitted and accepted.'
    case 'RoundFinalized':
      return 'Round finalized — collective public key locked in.'
    case 'CiphertextSubmitted':
      return 'Ciphertext submitted for threshold decryption.'
    case 'PartialDecryptionSubmitted':
      return 'Committee member submitted a partial decryption.'
    case 'DecryptionCombined':
      return 'Plaintext recovered via Lagrange-interpolated combine.'
    case 'RoundAborted':
      return 'Round aborted.'
    case 'ShareRevealed':
      return 'Secret share revealed.'
    case 'KeyReconstructed':
      return 'Secret key reconstructed.'
    default:
      return ev.eventName
  }
}

export function EventLog({ events }: { events: RoundEvent[] }) {
  if (events.length === 0) {
    return (
      <Box p={6} textAlign='center' color='gray.500' fontSize='sm'>
        No events for this round yet.
      </Box>
    )
  }
  // Newest first feels more natural in an "activity" view.
  const sorted = [...events].sort((a, b) => Number(b.blockNumber - a.blockNumber))
  return (
    <Stack gap={2}>
      {sorted.map((ev, i) => (
        <Box
          key={`${ev.transactionHash}-${i}`}
          borderWidth='1px'
          borderColor='gray.800'
          borderRadius='md'
          bg='gray.900'
          p={3}
        >
          <HStack justify='space-between' align='start' wrap='wrap' gap={2}>
            <Stack gap={0.5}>
              <Text fontSize='sm'>{eventSummary(ev)}</Text>
              <Text fontSize='2xs' color='gray.500' fontFamily='mono'>
                {ev.eventName}
              </Text>
            </Stack>
            <HStack gap={3} fontSize='2xs' color='gray.500'>
              <Text>block #{ev.blockNumber.toString()}</Text>
              <HashCell value={ev.transactionHash} head={6} tail={4} />
            </HStack>
          </HStack>
          <DetailDisclosure title='Show event args'>
            <RawJson value={ev.args} />
          </DetailDisclosure>
        </Box>
      ))}
    </Stack>
  )
}
