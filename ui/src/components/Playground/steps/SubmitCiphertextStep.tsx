import { useState } from 'react'
import { Box, Button, HStack, Stack, Text } from '@chakra-ui/react'
import type { ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import { LuPackage, LuUpload, LuRadio } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { HashCell } from '~components/ui/HashCell'
import { HowItWorks } from '../HowItWorks'

interface Props {
  status: StepStatus
  epochId: Hex | null
  /** Application id registered earlier in the flow. */
  aid: Hex | null
  ciphertext: ElGamalCiphertext | null
  onSubmitted: (ciphertextIndex: number, txHash: Hex) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx') => void
}

// The contract assigns the ciphertext index itself (1, 2, … per epoch and
// aid); we read it back from the CiphertextSubmitted event in the receipt.

export function SubmitCiphertextStep({ status, epochId, aid, ciphertext, onSubmitted, log }: Props) {
  const writer = useDkgWriter()
  const [busy, setBusy] = useState(false)
  const [tx, setTx] = useState<Hex | null>(null)
  const [index, setIndex] = useState<number | null>(null)

  const onSubmit = async () => {
    if (!writer || !epochId || !aid || !ciphertext) return
    setBusy(true)
    try {
      log('Sending submitCiphertext (c1, c2)…', 'tx')
      // submitCiphertext waits for the receipt and returns the index the
      // contract assigned.
      const result = await writer.submitCiphertext(epochId, aid, ciphertext)
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
      n={5}
      title='Publish the ciphertext on-chain'
      status={status}
      description='The committee watches the chain for new ciphertexts. As soon as yours lands, they start producing their half of the decryption.'
    >
      <Stack gap={4}>
        {!ciphertext || !aid ? (
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
          <HStack gap={3} wrap='wrap'>
            <Button
              colorPalette='cyan'
              size='sm'
              onClick={onSubmit}
              loading={busy}
              disabled={!writer || !epochId || !aid}
            >
              Publish ciphertext →
            </Button>
            {!writer && (
              <Text fontSize='xs' color='ink.3'>
                Connect a wallet to sign the <code>submitCiphertext</code> transaction — it must
                be the application's authorised submitter.
              </Text>
            )}
          </HStack>
        )}
        <HowItWorks
          body={
            <>
              Your wallet sends a single transaction carrying just the two ciphertext points.
              The contract checks they belong to your registered application and numbers them
              itself, so you never pick an index. Each committee node sees the new ciphertext
              through the contract's event log and spontaneously starts its share of the
              decryption — no off-chain coordination needed. Nothing they produce is useful on
              its own: the last piece is yours, in the next step.
            </>
          }
          flow={[
            { icon: <LuPackage />, label: 'Sealed ciphertext' },
            { icon: <LuUpload />, label: 'You publish on-chain' },
            { icon: <LuRadio />, label: 'Committee picks it up' },
          ]}
        />
      </Stack>
    </StepCard>
  )
}
