import { useState } from 'react'
import { Box, Button, Stack, Text } from '@chakra-ui/react'
import type { ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import { LuPackage, LuUpload, LuRadio } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { HashCell } from '~components/ui/HashCell'
import { HowItWorks } from '../HowItWorks'

interface Props {
  status: StepStatus
  roundId: Hex | null
  ciphertext: ElGamalCiphertext | null
  onSubmitted: (ciphertextIndex: number, txHash: Hex) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx') => void
}

// Index 1 by default — the playground only ever submits one ciphertext per
// round. The contract enforces write-once per index.
const CIPHERTEXT_INDEX = 1

export function SubmitCiphertextStep({ status, roundId, ciphertext, onSubmitted, log }: Props) {
  const writer = useDkgWriter()
  const [busy, setBusy] = useState(false)
  const [tx, setTx] = useState<Hex | null>(null)

  const onSubmit = async () => {
    if (!writer || !roundId || !ciphertext) return
    setBusy(true)
    try {
      log('Sending submitCiphertext…', 'tx')
      const hash = await writer.submitCiphertext(
        roundId,
        CIPHERTEXT_INDEX,
        ciphertext.c1[0],
        ciphertext.c1[1],
        ciphertext.c2[0],
        ciphertext.c2[1]
      )
      setTx(hash)
      log(`submitCiphertext tx: ${hash}`, 'tx')
      const receipt = await writer.waitForTransaction(hash)
      log(`Mined in block #${receipt.blockNumber}`, 'tx')
      onSubmitted(CIPHERTEXT_INDEX, hash)
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
      description='The committee watches the chain for new ciphertexts. As soon as yours lands, they start cooperating to decrypt it.'
    >
      <Stack gap={4}>
        {!ciphertext ? (
          <Text fontSize='sm' color='gray.500'>
            Encrypt something in the previous step first.
          </Text>
        ) : tx ? (
          <Stack gap={2} fontSize='sm'>
            <Text color='green.300'>Ciphertext published. Waiting for committee decryption…</Text>
            <Box>
              <Text fontSize='xs' color='gray.500'>
                tx
              </Text>
              <HashCell value={tx} head={8} tail={6} />
            </Box>
          </Stack>
        ) : (
          <Button colorPalette='cyan' size='sm' onClick={onSubmit} loading={busy} disabled={!writer || !roundId}>
            Publish ciphertext →
          </Button>
        )}
        <HowItWorks
          body={
            <>
              Your wallet sends a single transaction that stores the ciphertext on-chain. Each
              committee node sees the new ciphertext through the contract's event log and
              spontaneously starts the decryption work — no off-chain coordination needed.
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
