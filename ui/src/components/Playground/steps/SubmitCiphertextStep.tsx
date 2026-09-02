import { useState } from 'react'
import { Box, Button, HStack, Stack, Text } from '@chakra-ui/react'
import type { CiphertextPoK, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import { LuPackage, LuUpload, LuRadio } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { HashCell } from '~components/ui/HashCell'
import { HowItWorks } from '../HowItWorks'

interface Props {
  status: StepStatus
  epochId: Hex | null
  ciphertext: ElGamalCiphertext | null
  /** Proof of knowledge of the ciphertext randomness, produced in the encrypt step. */
  pok: CiphertextPoK | null
  onSubmitted: (ciphertextIndex: number, txHash: Hex) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx') => void
}

// The playground submits under the bare epoch key (aid = 0). The contract
// assigns the ciphertext index itself (1, 2, … per epoch and aid); we read it
// back from the CiphertextSubmitted event in the receipt.
const ZERO_AID = ('0x' + '00'.repeat(32)) as Hex

export function SubmitCiphertextStep({ status, epochId, ciphertext, pok, onSubmitted, log }: Props) {
  const writer = useDkgWriter()
  const [busy, setBusy] = useState(false)
  const [tx, setTx] = useState<Hex | null>(null)
  const [index, setIndex] = useState<number | null>(null)

  const onSubmit = async () => {
    if (!writer || !epochId || !ciphertext || !pok) return
    setBusy(true)
    try {
      log('Sending submitCiphertext (ciphertext + proof of knowledge)…', 'tx')
      // submitCiphertext verifies the proof locally, waits for the receipt
      // and returns the on-chain-assigned index.
      const result = await writer.submitCiphertext(epochId, ZERO_AID, ciphertext, pok)
      setTx(result.hash)
      setIndex(result.ciphertextIndex)
      log(`submitCiphertext tx: ${result.hash}`, 'tx')
      log(
        `Mined in block #${result.receipt.blockNumber}; contract assigned ciphertext index ${result.ciphertextIndex}`,
        'tx'
      )
      onSubmitted(result.ciphertextIndex, result.hash)
    } catch (err) {
      log(`submitCiphertext failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setBusy(false)
    }
  }

  return (
    <StepCard
      n={6}
      title='Publish the ciphertext on-chain'
      status={status}
      description='The committee watches the chain for new ciphertexts. As soon as yours lands, they check your proof and start cooperating to decrypt it.'
    >
      <Stack gap={4}>
        {!ciphertext || !pok ? (
          <Text fontSize='sm' color='ink.4'>
            Encrypt something in the previous step first.
          </Text>
        ) : tx ? (
          <Stack gap={2} fontSize='sm'>
            <Text color='live.fg'>Ciphertext published. Waiting for committee decryption…</Text>
            <HStack gap={6} align='start' wrap='wrap'>
              <Box>
                <Text fontSize='xs' color='ink.4'>
                  tx
                </Text>
                <HashCell value={tx} head={8} tail={6} />
              </Box>
              {index != null && (
                <Box>
                  <Text fontSize='xs' color='ink.4'>
                    assigned index
                  </Text>
                  <Text className='dkg-tabular' fontFamily='mono' fontSize='sm' color='ink.0'>
                    #{index}
                  </Text>
                </Box>
              )}
            </HStack>
          </Stack>
        ) : (
          <Button colorPalette='cyan' size='sm' onClick={onSubmit} loading={busy} disabled={!writer || !epochId}>
            Publish ciphertext →
          </Button>
        )}
        <HowItWorks
          body={
            <>
              Your wallet sends a single transaction that stores the ciphertext on-chain together
              with a small Schnorr proof that you know the randomness used to encrypt it. The
              contract numbers ciphertexts itself, so you never pick an index. Each committee node
              sees the new ciphertext through the contract's event log, verifies the proof (a
              ciphertext without a valid proof is simply ignored — it stops anyone from replaying
              someone else's ciphertext to use the committee as a decryption oracle) and
              spontaneously starts the decryption work — no off-chain coordination needed.
            </>
          }
          flow={[
            { icon: <LuPackage />, label: 'Sealed ciphertext + proof' },
            { icon: <LuUpload />, label: 'You publish on-chain' },
            { icon: <LuRadio />, label: 'Committee verifies & picks it up' },
          ]}
        />
      </Stack>
    </StepCard>
  )
}
