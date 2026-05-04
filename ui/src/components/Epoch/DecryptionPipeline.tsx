import { Box, Grid, HStack, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { useApplication } from '~queries/applications'
import { useEpoch } from '~queries/epochs'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { RawJson } from '~components/Debug/RawJson'

// DecryptionPipeline — read-only view of one application's decryption
// progress. Shows:
//   - the application's mode (public / co-decryption)
//   - committee partials submitted vs. threshold
//   - organizer share present? (mode 1 only)
//   - per-ciphertext combine status
//
// This component is a status panel, not a participant. It does not submit
// any transactions; the heavy lifting happens in the Go committee node and
// the on-chain `combineDecryption` call.
//
// The pipeline reads the cached `Application` record via useApplication,
// and falls back to a placeholder when the aid is unregistered.

interface Props {
  epochId: Hex
  aid: Hex
}

export function DecryptionPipeline({ epochId, aid }: Props) {
  const epoch = useEpoch(epochId)
  const app = useApplication(epochId, aid)

  if (app.isLoading || epoch.isLoading) {
    return <PanelShell title='Decryption pipeline'>Loading…</PanelShell>
  }
  if (app.isError) {
    return (
      <PanelShell title='Decryption pipeline'>
        <Text color='danger.fg' fontSize='sm'>
          {app.error instanceof Error ? app.error.message : String(app.error)}
        </Text>
      </PanelShell>
    )
  }
  if (!app.data?.exists) {
    return (
      <PanelShell title='Decryption pipeline'>
        <Text color='ink.3' fontSize='sm'>
          No application registered for this <code>aid</code> yet.
        </Text>
      </PanelShell>
    )
  }

  const mode = app.data.mode
  const ep = epoch.data?.epoch
  const threshold = ep?.policy.threshold ?? 0
  const committeeSize = ep?.policy.committeeSize ?? 0
  const partials = ep?.partialDecryptionCount ?? 0
  const ciphertexts = ep?.ciphertextCount ?? 0

  // Committee progress is shown as a fraction of the *threshold*, not the
  // committee size — once we have ≥ threshold partials for a single
  // ciphertext the combine can run. We don't have per-ciphertext partial
  // counts via the client today (only the epoch-wide total), so this is a
  // coarse indicator until the SDK exposes `getPartialDecryptionsByCt`.
  return (
    <PanelShell title='Decryption pipeline'>
      <Stack gap={6}>
        <SimpleGrid columns={{ base: 2, md: 4 }} gap={3}>
          <Stat label='Mode' value={mode === 0 ? 'Public derivation' : 'Organizer co-dec'} />
          <Stat label='Threshold' value={`${threshold} of ${committeeSize}`} />
          <Stat label='Ciphertexts' value={ciphertexts.toString()} />
          <Stat label='Partials (epoch-wide)' value={partials.toString()} />
        </SimpleGrid>

        <Stages mode={mode} hasPartials={partials > 0} />

        <Stack gap={2}>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            Application identity
          </Text>
          <Grid templateColumns={{ base: '1fr', md: '120px 1fr' }} gap={2}>
            <Label>aid</Label>
            <HashCell value={aid} head={8} tail={8} />
            <Label>Creator</Label>
            <HashCell value={app.data.creator} head={6} tail={4} />
            <Label>Created at</Label>
            <Text fontFamily='mono' fontSize='xs' color='ink.1'>
              block #{app.data.createdAtBlock.toString()}
            </Text>
            {mode === 0 && (
              <>
                <Label>S</Label>
                <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
                  {app.data.derivationS.toString()}
                </Text>
              </>
            )}
            {mode === 1 && (
              <>
                <Label>PK_org.x</Label>
                <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
                  {app.data.organizerPK[0].toString()}
                </Text>
                <Label>PK_org.y</Label>
                <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
                  {app.data.organizerPK[1].toString()}
                </Text>
              </>
            )}
          </Grid>
        </Stack>

        <DetailDisclosure title='Show raw application record'>
          <RawJson value={app.data} />
        </DetailDisclosure>
      </Stack>
    </PanelShell>
  )
}

// ─── small leaf components ──────────────────────────────────────────────────

function PanelShell({ title, children }: { title: string; children: React.ReactNode }) {
  return (
    <Box
      p={{ base: 5, md: 6 }}
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='lg'
      bg='surface'
      boxShadow='inset'
    >
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
        mb={4}
      >
        {title}
      </Text>
      {children}
    </Box>
  )
}

function Stat({ label, value }: { label: string; value: string }) {
  return (
    <Box position='relative' py={1.5} pl={3}>
      <Box
        position='absolute'
        left={0}
        top='25%'
        bottom='25%'
        w='2px'
        bg='accent.fg'
        opacity={0.4}
        borderRightRadius='full'
      />
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
        mb={1}
      >
        {label}
      </Text>
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize='md'
        fontWeight={500}
        color='ink.0'
      >
        {value}
      </Text>
    </Box>
  )
}

function Label({ children }: { children: React.ReactNode }) {
  return (
    <Text
      fontFamily='mono'
      fontSize='2xs'
      color='ink.4'
      letterSpacing='0.06em'
      textTransform='uppercase'
      mt={1}
    >
      {children}
    </Text>
  )
}

// Stages renders the per-app decryption flow as a horizontal pipeline so
// the user can see at a glance which step is active. The semantics are:
//
//   • Mode 0 (public derivation):
//       Ciphertext → Committee partials → Combine → Plaintext
//
//   • Mode 1 (organizer co-decryption):
//       Ciphertext → { Committee partials, Organizer share } → Combine → Plaintext
//
// `hasPartials` is the only signal we have today (the SDK doesn't yet
// expose per-ciphertext share counts); future iterations will replace it
// with per-ciphertext progress as the SDK gains those readers.
function Stages({ mode, hasPartials }: { mode: 0 | 1; hasPartials: boolean }) {
  const stages =
    mode === 0
      ? ['Ciphertext', 'Committee partials', 'Combine', 'Plaintext']
      : ['Ciphertext', 'Committee partials', 'Organizer Δ_org', 'Combine', 'Plaintext']
  return (
    <HStack gap={0} align='stretch' wrap='wrap'>
      {stages.map((s, i) => (
        <HStack key={s} gap={2} align='center' minH='32px'>
          <Box
            w='10px'
            h='10px'
            borderRadius='full'
            bg={hasPartials && i <= 1 ? 'live.fg' : 'border.subtle'}
            borderWidth='1px'
            borderColor={hasPartials && i <= 1 ? 'live.fg' : 'border'}
          />
          <Text fontSize='xs' color='ink.1' fontWeight={500} mr={3}>
            {s}
          </Text>
          {i < stages.length - 1 && (
            <Box w={{ base: 4, md: 8 }} h='1px' bg='rule' mr={3} />
          )}
        </HStack>
      ))}
    </HStack>
  )
}
