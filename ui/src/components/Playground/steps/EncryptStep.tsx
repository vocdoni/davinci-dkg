import { useState } from 'react'
import { Box, Button, Field, HStack, Input, Stack, Text } from '@chakra-ui/react'
import { buildElGamal, type BabyJubPoint, type ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { LuFile, LuLock, LuPackage } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { bigIntToHex } from '~lib/format'

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

  const onEncrypt = async () => {
    if (!collectivePubKey) return
    setBusy(true)
    try {
      const m = BigInt(plaintext || '0')
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
          <HStack gap={3} align='end'>
            <Field.Root maxW='200px'>
              <Field.Label fontSize='xs'>Number to encrypt</Field.Label>
              <Input
                size='sm'
                fontFamily='mono'
                value={plaintext}
                onChange={(e) => {
                  setPlaintext(e.target.value)
                  setCt(null)
                }}
              />
              <Field.HelperText fontSize='2xs'>
                Keep it small (under ~1 000 000) so the demo can decode it quickly.
              </Field.HelperText>
            </Field.Root>
            <Button colorPalette='purple' size='sm' onClick={onEncrypt} loading={busy}>
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
