import { useCallback, useState } from 'react'
import { Box, Grid, GridItem, Heading, Stack, Text } from '@chakra-ui/react'
import { useAccount } from 'wagmi'
import type { Hex } from 'viem'
import type { ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { ConnectStep } from '~components/Playground/steps/ConnectStep'
import { CreateRoundStep } from '~components/Playground/steps/CreateRoundStep'
import { WatchProgressStep } from '~components/Playground/steps/WatchProgressStep'
import { KeyAvailableStep } from '~components/Playground/steps/KeyAvailableStep'
import { EncryptStep } from '~components/Playground/steps/EncryptStep'
import { SubmitCiphertextStep } from '~components/Playground/steps/SubmitCiphertextStep'
import { VerifyDecryptionStep } from '~components/Playground/steps/VerifyDecryptionStep'
import { ActivityLog, type LogEntry, type LogLevel } from '~components/Playground/ActivityLog'
import type { StepStatus } from '~components/Playground/StepCard'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { useDebugMode } from '~hooks/use-debug-mode'

export function Playground() {
  const { isConnected } = useAccount()
  const { enabled: debug } = useDebugMode()

  const [roundId, setRoundId] = useState<Hex | null>(null)
  const [collectivePubKey, setCollectivePubKey] = useState<{ x: bigint; y: bigint } | null>(null)
  const [ciphertext, setCiphertext] = useState<ElGamalCiphertext | null>(null)
  const [plaintext, setPlaintext] = useState<bigint | null>(null)
  const [submittedIndex, setSubmittedIndex] = useState<number | null>(null)
  const [log, setLog] = useState<LogEntry[]>([])

  const addLog = useCallback((msg: string, level: LogLevel = 'info') => {
    setLog((prev) => [...prev, { ts: Date.now(), msg, level }])
  }, [])

  // Step status derivation. The pattern mirrors the legacy webapp: a step
  // is `done` when its primary output exists, `active` when the previous
  // step finished and it's now "open", `pending` otherwise.
  const stepWallet: StepStatus = isConnected ? 'done' : 'active'
  const stepCreate: StepStatus = !isConnected ? 'pending' : roundId ? 'done' : 'active'
  const stepWatch: StepStatus = !roundId ? 'pending' : collectivePubKey ? 'done' : 'active'
  const stepKey: StepStatus = !roundId ? 'pending' : collectivePubKey ? 'done' : 'active'
  const stepEncrypt: StepStatus = !collectivePubKey ? 'pending' : ciphertext ? 'done' : 'active'
  const stepSubmit: StepStatus = !ciphertext ? 'pending' : submittedIndex ? 'done' : 'active'
  const stepVerify: StepStatus = !submittedIndex ? 'pending' : 'active'

  const onSubmitted = useCallback((idx: number, _hash: Hex) => {
    setSubmittedIndex(idx)
  }, [])

  return (
    <Stack gap={6}>
      <Box>
        <Heading size='lg'>Playground</Heading>
        <Text color='gray.400' fontSize='sm' mt={1}>
          A guided walkthrough of the full DKG flow: create a round, wait for nodes to contribute,
          read the collective public key, encrypt a value, submit the ciphertext, and watch the
          committee threshold-decrypt it on-chain.
        </Text>
      </Box>
      <Grid templateColumns={{ base: '1fr', lg: '2fr 1fr' }} gap={6} alignItems='start'>
        <GridItem>
          <Stack gap={4}>
            <ConnectStep status={stepWallet} />
            <CreateRoundStep status={stepCreate} roundId={roundId} setRoundId={setRoundId} log={addLog} />
            <WatchProgressStep status={stepWatch} roundId={roundId} log={addLog} />
            <KeyAvailableStep status={stepKey} roundId={roundId} onKeyReady={setCollectivePubKey} log={addLog} />
            <EncryptStep
              status={stepEncrypt}
              collectivePubKey={collectivePubKey}
              onEncrypted={(m, ct) => {
                setPlaintext(m)
                setCiphertext(ct)
                setSubmittedIndex(null)
              }}
              log={addLog}
            />
            <SubmitCiphertextStep
              status={stepSubmit}
              roundId={roundId}
              ciphertext={ciphertext}
              onSubmitted={onSubmitted}
              log={addLog}
            />
            <VerifyDecryptionStep
              status={stepVerify}
              roundId={roundId}
              ciphertextIndex={submittedIndex}
              expectedPlaintext={plaintext}
              log={addLog}
            />
          </Stack>
        </GridItem>
        <GridItem position={{ base: 'static', lg: 'sticky' }} top={{ lg: 24 }}>
          <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={4}>
            <Heading size='xs' mb={3} color='gray.300'>
              Activity log
            </Heading>
            {debug ? (
              <ActivityLog entries={log} />
            ) : (
              <DetailDisclosure title={`Show activity log (${log.length} entries)`}>
                <ActivityLog entries={log} />
              </DetailDisclosure>
            )}
          </Box>
        </GridItem>
      </Grid>
    </Stack>
  )
}
