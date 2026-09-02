import { useState } from 'react'
import { Box, Button, Field, HStack, Input, Stack, Text } from '@chakra-ui/react'
import {
  applicationKey,
  encryptForApplication,
  type BabyJubPoint,
  type ElGamalCiphertext,
} from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import { LuFile, LuLock, LuPackage } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { bigIntToHex } from '~lib/format'

// Mirrors the Go committee's MaxDLogPlaintext (cmd/davinci-dkg-node/dlog.go).
// The committee can only recover plaintexts strictly below this — submitting
// anything larger guarantees the epoch will fail at the combine step. We
// reject it client-side so the user gets immediate, actionable feedback
// instead of waiting for the chain to finalize a doomed epoch.
const MAX_PLAINTEXT = 1n << 50n // 1,125,899,906,842,624

// The playground encrypts under the application key registered in the
// previous step: PK_aid = PK_ep + PK_org.

interface Props {
  status: StepStatus
  epochId: Hex | null
  collectivePubKey: { x: bigint; y: bigint } | null
  /** Organizer public key of the registered application, TE form. */
  pkOrg: BabyJubPoint | null
  onEncrypted: (plaintext: bigint, ct: ElGamalCiphertext) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'crypto') => void
}

export function EncryptStep({ status, epochId, collectivePubKey, pkOrg, onEncrypted, log }: Props) {
  const [plaintext, setPlaintext] = useState('42')
  const [ct, setCt] = useState<ElGamalCiphertext | null>(null)
  const [busy, setBusy] = useState(false)

  // Validate on every render so the button reflects current input without
  // a useEffect roundtrip. We accept only non-negative integers under the
  // committee's recoverable cap.
  const validation = validatePlaintext(plaintext)

  const onEncrypt = async () => {
    if (!collectivePubKey || !pkOrg || !epochId || validation.error) return
    setBusy(true)
    try {
      const m = validation.value!
      const pkEp: BabyJubPoint = [collectivePubKey.x, collectivePubKey.y]
      // PK_aid = PK_ep + PK_org. There is no proof of knowledge of the
      // randomness: the submitter of an aggregated tally cannot know it, and
      // cross-application replay is stopped by the organizer key instead.
      const ciphertext = await encryptForApplication(m, pkEp, pkOrg)
      setCt(ciphertext)
      onEncrypted(m, ciphertext)
      log(`Encrypted plaintext m=${m} under PK_aid for ${epochId} as ElGamal (c1, c2).`, 'crypto')
    } catch (err) {
      log(`Encrypt failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setBusy(false)
    }
  }

  const pkAid = collectivePubKey && pkOrg
    ? applicationKey([collectivePubKey.x, collectivePubKey.y], pkOrg)
    : null

  return (
    <StepCard
      n={4}
      title='Encrypt a value for the committee'
      status={status}
      description='Pick any small number and encrypt it with the application key. Opening it later takes the committee and you, together.'
    >
      {!collectivePubKey || !pkOrg ? (
        <Text fontSize='sm' color='ink.4'>
          Register an application first — ciphertexts are always encrypted under an
          application key.
        </Text>
      ) : (
        <Stack gap={4}>
          <HStack gap={3} align='end' wrap='wrap'>
            <Field.Root maxW='260px' invalid={!!validation.error}>
              <Field.Label fontSize='xs'>Number to encrypt</Field.Label>
              <Input
                size='sm'
                fontFamily='mono'
                inputMode='numeric'
                value={plaintext}
                onChange={(e) => {
                  setPlaintext(e.target.value)
                  setCt(null)
                }}
              />
              {validation.error ? (
                <Field.ErrorText fontSize='2xs'>{validation.error}</Field.ErrorText>
              ) : (
                <Field.HelperText fontSize='2xs'>
                  Any non-negative integer up to 2<sup>50</sup> ≈ 1.13 × 10<sup>15</sup>. The
                  committee's discrete-log recovery is capped there.
                </Field.HelperText>
              )}
            </Field.Root>
            <Button
              colorPalette='purple'
              size='sm'
              onClick={onEncrypt}
              loading={busy}
              disabled={!!validation.error || busy || !epochId || !pkOrg}
            >
              Encrypt →
            </Button>
          </HStack>
          {ct && (
            <Box>
              <Text fontSize='sm' color='live.fg'>
                Ciphertext ready. Submit it on-chain in the next step.
              </Text>
              <DetailDisclosure title='Show ciphertext components'>
                <Stack gap={2}>
                  <Box>
                    <Text fontSize='2xs' color='ink.4'>
                      c1 = k·G (random ephemeral, discloses nothing about the message)
                    </Text>
                    <HashCell value={bigIntToHex(ct.c1[0])} head={6} tail={6} />
                    <HashCell value={bigIntToHex(ct.c1[1])} head={6} tail={6} />
                  </Box>
                  <Box>
                    <Text fontSize='2xs' color='ink.4'>
                      c2 = m·G + k·PK_aid (the message, blinded by the application key)
                    </Text>
                    <HashCell value={bigIntToHex(ct.c2[0])} head={6} tail={6} />
                    <HashCell value={bigIntToHex(ct.c2[1])} head={6} tail={6} />
                  </Box>
                  {pkAid && (
                    <Box>
                      <Text fontSize='2xs' color='ink.4'>
                        PK_aid = PK_ep + PK_org (the key this ciphertext is bound to)
                      </Text>
                      <HashCell value={bigIntToHex(pkAid[0])} head={6} tail={6} />
                      <HashCell value={bigIntToHex(pkAid[1])} head={6} tail={6} />
                    </Box>
                  )}
                </Stack>
              </DetailDisclosure>
            </Box>
          )}
          <HowItWorks
            body={
              <>
                ElGamal encryption mixes your number with a fresh random value and the
                application key, producing two points on a curve. Each point looks like noise
                on its own. The application key is the committee's key plus your own, so
                recovering the number later needs both halves: even someone who could get the
                committee to decrypt on demand learns nothing without your share.
              </>
            }
            flow={[
              { icon: <LuFile />, label: 'Your number' },
              { icon: <LuLock />, label: 'Mixed with PK_aid' },
              { icon: <LuPackage />, label: 'Sealed ciphertext' },
            ]}
          />
        </Stack>
      )}
    </StepCard>
  )
}

interface PlaintextValidation {
  value?: bigint
  error?: string
}

// Accept a non-negative decimal integer strictly below MAX_PLAINTEXT.
// We deliberately don't fall back to "0" on empty input (the old code did)
// — silently encrypting zero hides the issue from the user.
function validatePlaintext(input: string): PlaintextValidation {
  const trimmed = input.trim()
  if (trimmed === '') return { error: 'Enter a number to encrypt.' }
  if (!/^\d+$/.test(trimmed)) return { error: 'Plaintext must be a non-negative integer.' }
  let value: bigint
  try {
    value = BigInt(trimmed)
  } catch {
    return { error: 'Plaintext must be a valid integer.' }
  }
  if (value >= MAX_PLAINTEXT) {
    return {
      error: `Plaintext must be below 2^50 (≈ 1.13 × 10^15). The committee's discrete-log recovery cannot decode larger values.`,
    }
  }
  return { value }
}
