import { useState } from 'react'
import { Box, Button, Field, HStack, Input, Stack, Text } from '@chakra-ui/react'
import { buildElGamal, type BabyJubPoint, type ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { LuFile, LuLock, LuPackage } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { bigIntToHex } from '~lib/format'

// Mirrors the Go committee's MaxDLogPlaintext (cmd/davinci-dkg-node/dlog.go).
// The committee can only recover plaintexts strictly below this — submitting
// anything larger guarantees the round will fail at the combine step. We
// reject it client-side so the user gets immediate, actionable feedback
// instead of waiting for the chain to finalize a doomed round.
const MAX_PLAINTEXT = 1n << 50n // 1,125,899,906,842,624

interface Props {
  status: StepStatus
  collectivePubKey: { x: bigint; y: bigint } | null
  onEncrypted: (plaintext: bigint, ct: ElGamalCiphertext) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'crypto') => void
}

export function EncryptStep({ status, collectivePubKey, onEncrypted, log }: Props) {
  const [plaintext, setPlaintext] = useState('42')
  const [ct, setCt] = useState<ElGamalCiphertext | null>(null)
  const [busy, setBusy] = useState(false)

  // Validate on every render so the button reflects current input without
  // a useEffect roundtrip. We accept only non-negative integers under the
  // committee's recoverable cap.
  const validation = validatePlaintext(plaintext)

  const onEncrypt = async () => {
    if (!collectivePubKey || validation.error) return
    setBusy(true)
    try {
      const m = validation.value!
      const eg = await buildElGamal()
      const pubKey: BabyJubPoint = [collectivePubKey.x, collectivePubKey.y]
      const result = eg.encrypt(m, pubKey)
      setCt(result)
      onEncrypted(m, result)
      log(`Encrypted plaintext m=${m} as ElGamal ciphertext (c1, c2).`, 'crypto')
    } catch (err) {
      log(`Encrypt failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setBusy(false)
    }
  }

  return (
    <StepCard
      n={5}
      title='Encrypt a value for the committee'
      status={status}
      description='Pick any small number and encrypt it with the shared key. The committee will need to cooperate to decrypt it later.'
    >
      {!collectivePubKey ? (
        <Text fontSize='sm' color='ink.4'>
          Waiting for the shared encryption key.
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
              disabled={!!validation.error || busy}
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
                      c2 = m·G + k·Q (the message, blinded by the shared key)
                    </Text>
                    <HashCell value={bigIntToHex(ct.c2[0])} head={6} tail={6} />
                    <HashCell value={bigIntToHex(ct.c2[1])} head={6} tail={6} />
                  </Box>
                </Stack>
              </DetailDisclosure>
            </Box>
          )}
          <HowItWorks
            body={
              <>
                ElGamal encryption mixes your number with a fresh random value and the
                committee's shared key, producing two points on a curve. Each point looks like
                noise on its own — only the committee, working together later, can subtract the
                blinding away and recover your original number.
              </>
            }
            flow={[
              { icon: <LuFile />, label: 'Your number' },
              { icon: <LuLock />, label: 'Mixed with shared key' },
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
