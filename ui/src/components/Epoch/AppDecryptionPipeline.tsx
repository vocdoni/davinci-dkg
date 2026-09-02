import { Box, Grid, HStack, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { LuCheck } from 'react-icons/lu'
import { useEpochDecryptionProgress } from '~queries/epochs'
import { HashCell } from '~components/ui/HashCell'
import type { ApplicationProgress, CiphertextProgress } from '~lib/decryption-overview'

interface Props {
  epochId: Hex
  /** Decryption threshold `t` — how many partials a combine needs. */
  threshold: number
  /** Committee size `n`. */
  committeeSize: number
  /** Only Live epochs have applications; skip the scans otherwise. */
  enabled?: boolean
}

/**
 * Every application registered against this epoch, each ciphertext walked
 * through the pipeline that opens it:
 *
 *   submitted → partials t/n → organizer share → combined
 *
 * The counts are real: partials are counted per ciphertext from
 * `PartialDecryptionSubmitted` (which names its submitter), not from the
 * epoch-wide counter, and the organizer-share stage reads the app manager's
 * own log. Members answer on a seed-derived stagger, so a row usually stops at
 * `t` rather than `n`: sitting at exactly `t` partials means finished, not
 * short.
 */
export function AppDecryptionPipeline({ epochId, threshold, committeeSize, enabled = true }: Props) {
  const progress = useEpochDecryptionProgress(epochId, enabled)

  if (!enabled) return null

  if (progress.isLoading) {
    return (
      <Shell>
        <Box h='72px' borderRadius='md' bg='surface.sunken' css={{ animation: 'dkgSkeletonPulse 1.6s ease-in-out infinite' }} />
      </Shell>
    )
  }
  if (progress.isError) {
    return (
      <Shell>
        <Text fontSize='sm' color='danger.fg'>
          {progress.error instanceof Error ? progress.error.message : String(progress.error)}
        </Text>
      </Shell>
    )
  }
  const apps = progress.data ?? []
  if (apps.length === 0) {
    return (
      <Shell>
        <Text fontSize='sm' color='ink.3'>
          No application has registered against this epoch yet. Applications bring their own
          organizer key; the epoch key alone cannot open anything.
        </Text>
      </Shell>
    )
  }

  return (
    <Shell>
      <Stack gap={5}>
        {apps.map((app) => (
          <AppRow key={app.aid} app={app} threshold={threshold} committeeSize={committeeSize} />
        ))}
      </Stack>
    </Shell>
  )
}

function AppRow({
  app,
  threshold,
  committeeSize,
}: {
  app: ApplicationProgress
  threshold: number
  committeeSize: number
}) {
  return (
    <Stack gap={3} borderTopWidth='1px' borderColor='rule' pt={4}>
      <HStack gap={4} wrap='wrap' align='baseline'>
        <HashCell value={app.aid} head={8} tail={6} />
        <Text fontSize='2xs' color='ink.4' fontFamily='mono'>
          registered at block #{app.registeredAtBlock.toString()}
        </Text>
        <Text fontSize='2xs' color='ink.4'>
          {app.submitted} ciphertext{app.submitted === 1 ? '' : 's'} · {app.combined} combined
        </Text>
      </HStack>

      {app.ciphertexts.length === 0 ? (
        <Text fontSize='xs' color='ink.4'>
          Registered, but nothing submitted under it yet.
        </Text>
      ) : (
        <Stack gap={2}>
          {app.ciphertexts.map((ct) => (
            <CiphertextStrip
              key={ct.index}
              ct={ct}
              threshold={threshold}
              committeeSize={committeeSize}
            />
          ))}
        </Stack>
      )}
    </Stack>
  )
}

/** One ciphertext as a four-stage strip with the real counts on each stage. */
function CiphertextStrip({
  ct,
  threshold,
  committeeSize,
}: {
  ct: CiphertextProgress
  threshold: number
  committeeSize: number
}) {
  const stages = [
    { label: 'Submitted', detail: `block #${ct.blockNumber.toString()}`, done: true },
    {
      label: `Partials ${ct.partials}/${committeeSize}`,
      detail: `${threshold} needed`,
      done: ct.partials >= threshold && threshold > 0,
    },
    {
      label: 'Organizer share',
      detail: ct.organizerShare ? 'released' : 'withheld',
      done: ct.organizerShare,
    },
    {
      label: 'Combined',
      detail: ct.combined && ct.plaintext != null ? `m = ${ct.plaintext.toString()}` : 'pending',
      done: ct.combined,
    },
  ]

  return (
    <Grid
      templateColumns={{ base: '32px 1fr', md: '48px 1fr' }}
      gap={3}
      alignItems='center'
      py={2}
      px={3}
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='md'
      bg='surface.sunken'
    >
      <Text className='dkg-tabular' fontFamily='mono' fontSize='xs' color='ink.3'>
        #{ct.index}
      </Text>
      <HStack gap={0} wrap='wrap' rowGap={2}>
        {stages.map((s, i) => (
          <HStack key={s.label} gap={2} align='center' pr={i < stages.length - 1 ? 2 : 0}>
            <Box
              w='14px'
              h='14px'
              borderRadius='full'
              borderWidth='1px'
              borderColor={s.done ? 'live.fg' : 'border.strong'}
              bg={s.done ? 'live.fg' : 'transparent'}
              color='canvas'
              display='flex'
              alignItems='center'
              justifyContent='center'
              flexShrink={0}
            >
              {s.done && <LuCheck size={9} strokeWidth={3} />}
            </Box>
            <Stack gap={0} minW='fit-content'>
              <Text fontSize='2xs' color={s.done ? 'ink.1' : 'ink.3'} whiteSpace='nowrap'>
                {s.label}
              </Text>
              <Text fontFamily='mono' fontSize='2xs' color='ink.4' whiteSpace='nowrap'>
                {s.detail}
              </Text>
            </Stack>
            {i < stages.length - 1 && (
              <Box w={{ base: 3, md: 6 }} h='1px' bg='rule' ml={2} flexShrink={0} />
            )}
          </HStack>
        ))}
      </HStack>
    </Grid>
  )
}

function Shell({ children }: { children: React.ReactNode }) {
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
        Decryption pipeline per application
      </Text>
      {children}
    </Box>
  )
}
