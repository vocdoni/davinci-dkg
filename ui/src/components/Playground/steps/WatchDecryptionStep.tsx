import { useEffect } from 'react'
import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import {
  pointFromTEtoRTE,
  type ElGamalCiphertext,
} from '@vocdoni/davinci-dkg-sdk'
import { LuUsers, LuCombine, LuEye } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { useEpoch } from '~queries/epochs'
import { useCiphertextPartials } from '~queries/epochs'
import { useCiphertextRecord, useCiphertextStatus } from '~queries/applications'

interface Props {
  status: StepStatus
  epochId: Hex | null
  aid: Hex | null
  ciphertextIndex: number | null
  /** The ciphertext this browser built, for the local match check. */
  ciphertext: ElGamalCiphertext | null
  /** The number this browser encrypted, for the local plaintext check. */
  expectedPlaintext: bigint | null
  /** True while the organizer is deliberately sitting on its share. */
  shareWithheld: boolean
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'chain') => void
}

/**
 * The last step: watch the committee answer, and check the result locally.
 *
 * Two things are shown that the contract's counters cannot give you.
 * `PartialDecryptionSubmitted` names its submitter, so the committee members
 * that answered *this* ciphertext are listed by address — and the count is
 * expected to stop around `t`, not `n`: members answer on a seed-derived
 * stagger so the ones later in the rotation usually never spend the gas. The
 * organizer-share row is the other half: partials keep arriving while it is
 * missing, and `combineDecryption` reverts until it lands.
 */
export function WatchDecryptionStep({
  status,
  epochId,
  aid,
  ciphertextIndex,
  ciphertext,
  expectedPlaintext,
  shareWithheld,
  log,
}: Props) {
  const epoch = useEpoch((epochId ?? undefined) as `0x${string}` | undefined)
  const partials = useCiphertextPartials(
    (epochId ?? undefined) as `0x${string}` | undefined,
    (aid ?? undefined) as `0x${string}` | undefined,
    ciphertextIndex,
  )
  const ctStatus = useCiphertextStatus(
    (epochId ?? undefined) as `0x${string}` | undefined,
    (aid ?? undefined) as `0x${string}` | undefined,
    ciphertextIndex ?? 0,
  )
  const record = useCiphertextRecord(
    (epochId ?? undefined) as `0x${string}` | undefined,
    (aid ?? undefined) as `0x${string}` | undefined,
    ciphertextIndex,
  )

  const threshold = epoch.data?.epoch.policy.threshold ?? 0
  const committeeSize = epoch.data?.epoch.policy.committeeSize ?? 0
  const responders = partials.data ?? []
  const shareReleased = ctStatus.data?.organizerShare === true
  const combined = ctStatus.data?.combined === true
  const plaintext = combined ? (ctStatus.data?.plaintext ?? null) : null

  useEffect(() => {
    if (combined && plaintext != null) {
      log(`Plaintext recovered on-chain: ${plaintext.toString()}`, 'success')
    }
  }, [combined, plaintext, log])

  const ciphertextMatches = matchesOnChain(ciphertext, record.data)
  const plaintextMatches =
    plaintext != null && expectedPlaintext != null ? plaintext === expectedPlaintext : null

  return (
    <StepCard
      n={7}
      title='Watch the committee open it — then check the result yourself'
      status={status}
      description='Committee members publish their partial decryptions on their own. With your share released too, anyone can combine them and the number lands back on chain.'
    >
      <Stack gap={5}>
        {ciphertextIndex == null ? (
          <Text fontSize='sm' color='ink.4'>
            Publish a ciphertext first — there is nothing to decrypt yet.
          </Text>
        ) : (
          <Stack gap={5}>
            <PartialsMeter
              have={responders.length}
              threshold={threshold}
              committeeSize={committeeSize}
              loading={partials.isLoading}
            />

            {responders.length > 0 && (
              <Stack gap={1.5}>
                <Text
                  fontFamily='mono'
                  fontSize='2xs'
                  color='ink.3'
                  letterSpacing='0.08em'
                  textTransform='uppercase'
                >
                  Members that answered
                </Text>
                {responders.map((p) => (
                  <HStack key={`${p.participant}-${p.blockNumber}`} gap={3} wrap='wrap'>
                    <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='accent.dim'>
                      #{p.participantIndex}
                    </Text>
                    <HashCell value={p.participant} head={6} tail={6} />
                    <Text fontFamily='mono' fontSize='2xs' color='ink.4'>
                      block #{p.blockNumber.toString()}
                    </Text>
                  </HStack>
                ))}
              </Stack>
            )}

            <Stack gap={2}>
              <Row
                ok={shareReleased}
                label='Organizer share'
                okText='released — the committee can combine'
                pendingText={
                  shareWithheld
                    ? 'withheld by you — partials keep arriving, but nothing can be combined'
                    : 'not on chain yet'
                }
              />
              <Row
                ok={combined}
                label='Combined'
                okText={`plaintext published on chain${plaintext != null ? `: m = ${plaintext.toString()}` : ''}`}
                pendingText={
                  shareReleased
                    ? 'waiting for a member to submit the combine proof'
                    : 'blocked — combineDecryption reverts with OrganizerShareMissing()'
                }
              />
            </Stack>

            {(plaintextMatches != null || ciphertextMatches != null) && (
              <Box
                borderLeftWidth='2px'
                borderColor={plaintextMatches === false ? 'danger.fg' : 'live.fg'}
                bg={plaintextMatches === false ? 'danger.bg' : 'live.bg'}
                px={4}
                py={3}
                borderRadius='md'
              >
                <Text fontSize='2xs' color='ink.3' letterSpacing='0.06em' textTransform='uppercase' mb={2}>
                  Checked in this browser
                </Text>
                <Stack gap={1.5}>
                  {ciphertextMatches != null && (
                    <Check
                      ok={ciphertextMatches}
                      okText='The ciphertext stored on chain is byte-for-byte the one this page built.'
                      badText='The ciphertext on chain differs from the one this page built.'
                    />
                  )}
                  {plaintextMatches != null && (
                    <Check
                      ok={plaintextMatches}
                      okText={`The committee recovered ${plaintext!.toString()} — exactly the number you encrypted.`}
                      badText={`The committee recovered ${plaintext!.toString()}, but you encrypted ${expectedPlaintext!.toString()}.`}
                    />
                  )}
                </Stack>
                <Text fontSize='2xs' color='ink.3' mt={2} lineHeight='1.55'>
                  Both comparisons run locally against values read straight from the chain — the
                  page is not taking anyone's word for the result.
                </Text>
              </Box>
            )}

            <DetailDisclosure title='Show what the chain stores'>
              <Stack gap={1} fontSize='2xs' fontFamily='mono' color='ink.3'>
                <Text>partials: δ_i = d_i·C1 + Groth16 proof of a Chaum–Pedersen DLEQ</Text>
                <Text>organizer: Δ = sk_org·C1, stored as keccak256(Δ ‖ A1 ‖ A2 ‖ z)</Text>
                <Text>combine: proves Σ λ_k·δ_k interpolates and m·G + Σ λ_k·δ_k + Δ = C2</Text>
                <Text>result: m, readable via getPlaintext(eid, aid, ctIdx)</Text>
              </Stack>
            </DetailDisclosure>
          </Stack>
        )}

        <HowItWorks
          body={
            <>
              No committee member can decrypt alone — that is the point of a threshold scheme.
              Each of the <em>t</em> members whose turn comes up publishes a partial decryption
              with a proof that it used its real share, and you publish yours. Once both halves
              are on chain any member can combine them; the combine proof re-checks your share's
              proof as part of the same SNARK, so a bad share is rejected rather than corrupting
              the answer. Until the combine lands, the ciphertext stays opaque to everyone.
            </>
          }
          flow={[
            { icon: <LuUsers />, label: 'Members publish partials' },
            { icon: <LuCombine />, label: 'Pieces combined on-chain' },
            { icon: <LuEye />, label: 'Plaintext revealed' },
          ]}
        />
      </Stack>
    </StepCard>
  )
}

/** `have` of `threshold` partials, with the committee size as context. */
function PartialsMeter({
  have,
  threshold,
  committeeSize,
  loading,
}: {
  have: number
  threshold: number
  committeeSize: number
  loading: boolean
}) {
  const done = threshold > 0 && have >= threshold
  // The bar is scaled to the committee, with the threshold marked on it: the
  // question "are we there yet" is `have ≥ t`, but the honest denominator for a
  // count of members is `n`.
  const pct = committeeSize > 0 ? Math.min(100, (have / committeeSize) * 100) : 0
  const thresholdPct = committeeSize > 0 ? Math.min(100, (threshold / committeeSize) * 100) : 0
  return (
    <Box
      borderWidth='1px'
      borderColor={done ? 'rgba(134, 239, 172, 0.40)' : 'border.subtle'}
      bg={done ? 'live.bg' : 'surface.sunken'}
      borderRadius='lg'
      p={{ base: 4, md: 5 }}
    >
      <Stack gap={3}>
        <HStack justify='space-between' align='baseline' wrap='wrap' gap={2}>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            Partial decryptions
          </Text>
          <Text fontFamily='mono' fontSize='2xs' color='ink.4' letterSpacing='0.06em'>
            threshold {threshold} of {committeeSize}
          </Text>
        </HStack>
        <HStack align='baseline' gap={2}>
          <Text
            className='dkg-tabular'
            fontFamily='mono'
            fontSize={{ base: '3xl', md: '4xl' }}
            fontWeight={600}
            color={done ? 'live.fg' : 'ink.0'}
            lineHeight='1'
          >
            {loading ? '…' : have}
          </Text>
          <Text className='dkg-tabular' fontFamily='mono' fontSize='xl' color='ink.3'>
            / {committeeSize} members
          </Text>
        </HStack>
        <Box position='relative' h='8px' bg='surface.raised' borderRadius='full'>
          <Box
            position='absolute'
            inset={0}
            w={`${pct}%`}
            bg={done ? 'live.fg' : 'accent.fg'}
            borderRadius='full'
            transition='width 0.3s ease'
          />
          {threshold > 0 && threshold < committeeSize && (
            <Box
              position='absolute'
              top='-3px'
              bottom='-3px'
              left={`${thresholdPct}%`}
              w='2px'
              bg='ink.1'
              opacity={0.55}
              transform='translateX(-1px)'
              borderRadius='full'
              title={`threshold: ${threshold}`}
            />
          )}
        </Box>
        <Text fontSize='xs' color='ink.3' lineHeight='1.5'>
          {done
            ? 'Threshold reached — enough partials exist to combine. Members answer on a seed-derived stagger, so the ones further down the rotation usually never spend the gas; any that do are harmless.'
            : 'Committee members react to the ciphertext on their own; nothing here is waiting on you.'}
        </Text>
      </Stack>
    </Box>
  )
}

function Row({
  ok,
  label,
  okText,
  pendingText,
}: {
  ok: boolean
  label: string
  okText: string
  pendingText: string
}) {
  return (
    <HStack gap={3} align='start'>
      <Box
        mt='5px'
        w='8px'
        h='8px'
        borderRadius='full'
        flexShrink={0}
        bg={ok ? 'live.fg' : 'border.strong'}
      />
      <Text fontSize='sm' color={ok ? 'ink.1' : 'ink.3'}>
        <Box as='span' color='ink.0'>
          {label}:
        </Box>{' '}
        {ok ? okText : pendingText}
      </Text>
    </HStack>
  )
}

function Check({ ok, okText, badText }: { ok: boolean; okText: string; badText: string }) {
  return (
    <Text fontSize='sm' color={ok ? 'ink.1' : 'danger.fg'}>
      {ok ? '✓ ' : '✗ '}
      {ok ? okText : badText}
    </Text>
  )
}

/**
 * Compare the locally-built ciphertext with the one the contract logged. The
 * event carries the on-chain (RTE) words while the browser holds TE
 * coordinates, so the local pair is converted before comparing.
 */
function matchesOnChain(
  local: ElGamalCiphertext | null,
  onChain: { c1: { x: bigint; y: bigint }; c2: { x: bigint; y: bigint } } | null | undefined,
): boolean | null {
  if (!local || !onChain) return null
  const [c1x, c1y] = pointFromTEtoRTE(local.c1)
  const [c2x, c2y] = pointFromTEtoRTE(local.c2)
  return (
    c1x === onChain.c1.x && c1y === onChain.c1.y && c2x === onChain.c2.x && c2y === onChain.c2.y
  )
}
