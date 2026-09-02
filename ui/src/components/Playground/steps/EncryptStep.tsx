import { useState } from 'react'
import { Box, Button, Field, HStack, Input, Stack, Text } from '@chakra-ui/react'
import {
  encryptWithProof,
  type BabyJubPoint,
  type CiphertextPoK,
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

// The playground encrypts under the bare epoch key, i.e. application id 0.
// The proof of knowledge is bound to (epochId, aid), so the same pair must
// be used when the ciphertext is submitted.
const ZERO_AID = ('0x' + '00'.repeat(32)) as Hex

interface Props {
  status: StepStatus
  epochId: Hex | null
  collectivePubKey: { x: bigint; y: bigint } | null
  onEncrypted: (plaintext: bigint, ct: ElGamalCiphertext, pok: CiphertextPoK) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'crypto') => void
}

export function EncryptStep({ status, epochId, collectivePubKey, onEncrypted, log }: Props) {
  const [plaintext, setPlaintext] = useState('42')
  const [ct, setCt] = useState<ElGamalCiphertext | null>(null)
  const [pok, setPok] = useState<CiphertextPoK | null>(null)
  const [busy, setBusy] = useState(false)

  // Validate on every render so the button reflects current input without
  // a useEffect roundtrip. We accept only non-negative integers under the
  // committee's recoverable cap.
  const validation = validatePlaintext(plaintext)

  const onEncrypt = async () => {
    if (!collectivePubKey || !epochId || validation.error) return
    setBusy(true)
    try {
      const m = validation.value!
      const pubKey: BabyJubPoint = [collectivePubKey.x, collectivePubKey.y]
      // encryptWithProof draws the randomness r, encrypts, and proves
      // knowledge of r for exactly this (epochId, aid). Committee nodes only
      // decrypt ciphertexts whose proof verifies.
      const result = await encryptWithProof(epochId, ZERO_AID, m, pubKey)
      setCt(result.ciphertext)
      setPok(result.pok)
      onEncrypted(m, result.ciphertext, result.pok)
      log(`Encrypted plaintext m=${m} as ElGamal ciphertext (c1, c2) + proof of knowledge for ${epochId}.`, 'crypto')
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
                  setPok(null)
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
              disabled={!!validation.error || busy || !epochId}
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
                  {pok && (
                    <Box>
                      <Text fontSize='2xs' color='ink.4'>
                        Schnorr proof of knowledge of k, bound to this epoch: A = w·G, z = w + c·k
                      </Text>
                      <HashCell value={bigIntToHex(pok.ax)} head={6} tail={6} />
                      <HashCell value={bigIntToHex(pok.ay)} head={6} tail={6} />
                      <HashCell value={bigIntToHex(pok.z)} head={6} tail={6} />
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
                committee's shared key, producing two points on a curve. Each point looks like
                noise on its own — only the committee, working together later, can subtract the
                blinding away and recover your original number. Alongside, the SDK produces a
                tiny proof that you know that random value; the committee refuses to decrypt
                ciphertexts that lack it, so nobody can copy someone else's ciphertext and have
                it opened for them.
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
