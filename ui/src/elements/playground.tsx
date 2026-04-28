import { useCallback, useState } from 'react'
import { Box, Grid, GridItem, HStack, Stack, Text } from '@chakra-ui/react'
import { useAccount } from 'wagmi'
import type { Hex } from 'viem'
import { RoundStatus, type ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
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
import { PageHeader } from '~components/Layout/PageHeader'
import { useRound } from '~queries/rounds'
import { useBlockNumber } from '~queries/chain'
import { roundFailure } from '~lib/round-utils'

// Playground page. Editorial masthead, then a two-column layout: the
// numbered step cards on the left, a sticky activity-log panel on the
// right. The right rail collapses behind a disclosure when debug mode is
// off so casual readers don't see a wall of mono terminal output by
// default.
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

  // Surface a round-failure state up here so every downstream step can
  // freeze. Same React Query key as the WatchProgressStep, so this is a
  // free read from the cache.
  const round = useRound((roundId ?? undefined) as `0x${string}` | undefined)
  const { data: block } = useBlockNumber()
  const failure = round.data ? roundFailure(round.data.round, block ?? null) : null
  const aborted = round.data?.round.status === RoundStatus.Aborted
  const blocked = Boolean(failure || aborted)

  const stepWallet: StepStatus = isConnected ? 'done' : 'active'
  const stepCreate: StepStatus = !isConnected ? 'pending' : roundId ? 'done' : 'active'
  const stepWatch: StepStatus = !roundId
    ? 'pending'
    : blocked
      ? 'error'
      : collectivePubKey
        ? 'done'
        : 'active'
  const stepKey: StepStatus = !roundId
    ? 'pending'
    : blocked
      ? 'pending'
      : collectivePubKey
        ? 'done'
        : 'active'
  const stepEncrypt: StepStatus = !collectivePubKey || blocked ? 'pending' : ciphertext ? 'done' : 'active'
  const stepSubmit: StepStatus = !ciphertext || blocked ? 'pending' : submittedIndex ? 'done' : 'active'
  const stepVerify: StepStatus = !submittedIndex || blocked ? 'pending' : 'active'

  const onSubmitted = useCallback((idx: number, _hash: Hex) => {
    setSubmittedIndex(idx)
  }, [])

  return (
    <Stack gap={{ base: 8, md: 12 }}>
      <PageHeader
        title='Playground'
        subtitle='A guided walkthrough of the full DKG flow: create a round, wait for nodes to contribute, read the collective public key, encrypt a value, submit the ciphertext, and watch the committee threshold-decrypt it on-chain.'
      />

      <Grid templateColumns={{ base: '1fr', lg: '2fr 1fr' }} gap={{ base: 6, lg: 8 }} alignItems='start'>
        <GridItem>
          <Stack gap={5}>
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
          <Box
            borderWidth='1px'
            borderColor='border.subtle'
            borderRadius='lg'
            bg='surface'
            p={4}
            boxShadow='inset'
          >
            <HStack
              mb={3}
              fontFamily='mono'
              fontSize='2xs'
              color='ink.3'
              letterSpacing='0.08em'
              gap={2}
            >
              <Box w='6px' h='6px' borderRadius='full' bg='live.fg' />
              <Text textTransform='uppercase'>Activity log</Text>
            </HStack>
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
