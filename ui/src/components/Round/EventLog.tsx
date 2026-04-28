import { Box, Grid, GridItem, HStack, Stack, Text } from '@chakra-ui/react'
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

// Editorial event log. Each event is a journal-entry-style row:
//
//   #BLK ───── plain-English summary
//              evname · txhash
//              [+ raw args]
//
// The block number column anchors each row, mono and tabular, so a long
// log reads as a chronicle of state changes. The hairline rules between
// rows form a typeset list, no per-row card.
export function EventLog({ events }: { events: RoundEvent[] }) {
  if (events.length === 0) {
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
          No events for this round yet.
        </Text>
      </Box>
    )
  }
  // Newest first feels more natural in an "activity" view.
  const sorted = [...events].sort((a, b) => Number(b.blockNumber - a.blockNumber))
  return (
    <Stack gap={0}>
      {sorted.map((ev, i) => (
        <Grid
          key={`${ev.transactionHash}-${i}`}
          templateColumns={{ base: '1fr', md: '120px 1fr' }}
          gap={{ base: 2, md: 6 }}
          py={4}
          borderTopWidth={i === 0 ? '1px' : 0}
          borderBottomWidth='1px'
          borderColor='rule'
        >
          <GridItem>
            <Text
              className='dkg-tabular'
              fontFamily='mono'
              fontSize='xs'
              color='ink.4'
              letterSpacing='0.04em'
              whiteSpace='nowrap'
            >
              #{ev.blockNumber.toString()}
            </Text>
          </GridItem>
          <GridItem>
            <Text fontSize='sm' color='ink.1' lineHeight='1.5' mb={1.5}>
              {eventSummary(ev)}
            </Text>
            <HStack gap={3} fontFamily='mono' fontSize='2xs' color='ink.4'>
              <Text color='accent.dim'>{ev.eventName}</Text>
              <Box w='1px' h='10px' bg='border' />
              <HashCell value={ev.transactionHash} head={6} tail={4} />
            </HStack>
            <DetailDisclosure title='Show event args'>
              <RawJson value={ev.args} />
            </DetailDisclosure>
          </GridItem>
        </Grid>
      ))}
    </Stack>
  )
}
