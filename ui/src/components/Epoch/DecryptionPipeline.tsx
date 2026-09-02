import { Box, Grid, HStack, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import {
  useApplication,
  useCiphertextCount,
  useCiphertextStatus,
} from '~queries/applications'
import { useEpoch } from '~queries/epochs'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { RawJson } from '~components/Debug/RawJson'

// DecryptionPipeline — read-only view of one application's decryption
// progress. Shows:
//   - the application's organizer key
//   - committee partials submitted vs. threshold
//   - per-ciphertext: is the organizer share on chain, is it combined
//
// The organizer-share column is the load-bearing one: a ciphertext with
// enough committee partials still cannot be combined until the organizer
// releases Δ = sk_org·C1 (the contract reverts `OrganizerShareMissing()`), so
// without it "stuck" is indistinguishable from "committee is slow".
//
// This component is a status panel, not a participant. It does not submit any
// transactions; the heavy lifting happens in the Go committee node, the
// organizer's browser, and the on-chain `combineDecryption` call.
//
// It answers "what is the state of *this* aid"; the per-application pipeline
// above it on the epoch page answers "what is happening in this epoch" and
// carries the per-ciphertext partial counts.

interface Props {
  epochId: Hex
  aid: Hex
}

/** How many ciphertext rows to render — one pair of reads each. */
const MAX_ROWS = 10

export function DecryptionPipeline({ epochId, aid }: Props) {
  const epoch = useEpoch(epochId)
  const app = useApplication(epochId, aid)
  const count = useCiphertextCount(epochId, aid)

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
          No application registered for this <code>aid</code> yet. Ciphertexts can only be
          submitted under a registered application.
        </Text>
      </PanelShell>
    )
  }

  const ep = epoch.data?.epoch
  const threshold = ep?.policy.threshold ?? 0
  const committeeSize = ep?.policy.committeeSize ?? 0
  const partials = ep?.partialDecryptionCount ?? 0
  const ciphertexts = Number(count.data ?? 0)
  const rows = Array.from({ length: Math.min(ciphertexts, MAX_ROWS) }, (_, i) => i + 1)

  // Committee progress is shown as a fraction of the *threshold*, not the
  // committee size — once we have ≥ threshold partials for a single
  // ciphertext the combine can run. We don't have per-ciphertext partial
  // counts via the client today (only the epoch-wide total), so this is a
  // coarse indicator until the SDK exposes `getPartialDecryptionsByCt`.
  return (
    <PanelShell title='Decryption pipeline'>
      <Stack gap={6}>
        <SimpleGrid columns={{ base: 2, md: 4 }} gap={3}>
          <Stat label='Threshold' value={`${threshold} of ${committeeSize}`} />
          <Stat label='Ciphertexts (this aid)' value={ciphertexts.toString()} />
          <Stat label='Partials (epoch-wide)' value={partials.toString()} />
          <Stat label='Organizer' value='required' />
        </SimpleGrid>

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
            <Label>Submitter</Label>
            <HashCell value={app.data.policy.authorizedSubmitter} head={6} tail={4} />
            <Label>Created at</Label>
            <Text fontFamily='mono' fontSize='xs' color='ink.1'>
              block #{app.data.createdAtBlock.toString()}
            </Text>
            <Label>PK_org.x</Label>
            <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
              {app.data.organizerPK[0].toString()}
            </Text>
            <Label>PK_org.y</Label>
            <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
              {app.data.organizerPK[1].toString()}
            </Text>
          </Grid>
        </Stack>

        <Stack gap={2}>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            Ciphertexts
          </Text>
          {rows.length === 0 ? (
            <Text fontSize='sm' color='ink.3'>
              No ciphertexts submitted under this application yet.
            </Text>
          ) : (
            <Stack gap={0}>
              <Grid
                templateColumns='60px 1fr 1fr 1fr'
                gap={3}
                pb={2}
                borderBottomWidth='1px'
                borderColor='border.subtle'
              >
                <Label>#</Label>
                <Label>Organizer share</Label>
                <Label>Combined</Label>
                <Label>Plaintext</Label>
              </Grid>
              {rows.map((ix) => (
                <CiphertextRow key={ix} epochId={epochId} aid={aid} index={ix} />
              ))}
              {ciphertexts > MAX_ROWS && (
                <Text fontSize='2xs' color='ink.4' mt={2}>
                  Showing the first {MAX_ROWS} of {ciphertexts}.
                </Text>
              )}
            </Stack>
          )}
        </Stack>

        <DetailDisclosure title='Show raw application record'>
          <RawJson value={app.data} />
        </DetailDisclosure>
      </Stack>
    </PanelShell>
  )
}

// ─── small leaf components ──────────────────────────────────────────────────

function CiphertextRow({ epochId, aid, index }: { epochId: Hex; aid: Hex; index: number }) {
  const status = useCiphertextStatus(epochId, aid, index)
  const share = status.data?.organizerShare
  const combined = status.data?.combined
  return (
    <Grid
      templateColumns='60px 1fr 1fr 1fr'
      gap={3}
      py={2}
      borderBottomWidth='1px'
      borderColor='border.subtle'
      alignItems='center'
    >
      <Text className='dkg-tabular' fontFamily='mono' fontSize='xs' color='ink.1'>
        {index}
      </Text>
      <Pill
        ok={share === true}
        okLabel='released'
        pendingLabel={status.isLoading ? 'checking…' : 'awaiting organizer'}
      />
      <Pill
        ok={combined === true}
        okLabel='combined'
        pendingLabel={status.isLoading ? 'checking…' : 'awaiting committee'}
      />
      <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
        {combined ? status.data!.plaintext.toString() : '—'}
      </Text>
    </Grid>
  )
}

function Pill({ ok, okLabel, pendingLabel }: { ok: boolean; okLabel: string; pendingLabel: string }) {
  return (
    <HStack gap={2} align='center'>
      <Box
        w='8px'
        h='8px'
        borderRadius='full'
        bg={ok ? 'live.fg' : 'border.subtle'}
        borderWidth='1px'
        borderColor={ok ? 'live.fg' : 'border'}
      />
      <Text fontSize='xs' color={ok ? 'ink.1' : 'ink.3'}>
        {ok ? okLabel : pendingLabel}
      </Text>
    </HStack>
  )
}

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
