import { useCallback, useRef, useState } from 'react'
import { Box, Grid, GridItem, HStack, Stack, Text } from '@chakra-ui/react'
import { useAccount } from 'wagmi'
import type { Hex } from 'viem'
import type { BabyJubPoint, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { ConnectStep } from '~components/Playground/steps/ConnectStep'
import { LiveEpochStep } from '~components/Playground/steps/LiveEpochStep'
import { RegisterAppStep } from '~components/Playground/steps/RegisterAppStep'
import { EncryptStep } from '~components/Playground/steps/EncryptStep'
import { SubmitCiphertextStep } from '~components/Playground/steps/SubmitCiphertextStep'
import { ReleaseShareStep } from '~components/Playground/steps/ReleaseShareStep'
import { WatchDecryptionStep } from '~components/Playground/steps/WatchDecryptionStep'
import { ActivityLog, type LogEntry, type LogLevel } from '~components/Playground/ActivityLog'
import type { StepStatus } from '~components/Playground/StepCard'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { useDebugMode } from '~hooks/use-debug-mode'
import { PageHeader } from '~components/Layout/PageHeader'

// Playground page. You play the organizer of an application: the committee and
// its key already exist (operators produce epochs on a cadence, and nothing an
// application does influences that), so the walkthrough starts where an
// integrator actually starts — registering an application against the epoch
// that is live right now.
//
// Layout: numbered step cards on the left, a sticky activity log on the right.
// The log collapses behind a disclosure when debug mode is off so a casual
// reader doesn't get a wall of terminal output by default.
export function Playground() {
  const { isConnected } = useAccount()
  const { enabled: debug } = useDebugMode()

  const [epochId, setEpochId] = useState<Hex | null>(null)
  const [collectivePubKey, setCollectivePubKey] = useState<{ x: bigint; y: bigint } | null>(null)
  const [aid, setAid] = useState<Hex | null>(null)
  const [skOrg, setSkOrg] = useState<bigint | null>(null)
  const [pkOrg, setPkOrg] = useState<BabyJubPoint | null>(null)
  const [ciphertext, setCiphertext] = useState<ElGamalCiphertext | null>(null)
  const [plaintext, setPlaintext] = useState<bigint | null>(null)
  const [submittedIndex, setSubmittedIndex] = useState<number | null>(null)
  const [share, setShare] = useState<'idle' | 'released' | 'withheld'>('idle')
  const [log, setLog] = useState<LogEntry[]>([])

  const addLog = useCallback((msg: string, level: LogLevel = 'info') => {
    setLog((prev) => [...prev, { ts: Date.now(), msg, level }])
  }, [])

  // Once an application is registered, every later step must use the epoch it
  // lives on, even if a newer epoch goes Live while the page is open.
  const aidRef = useRef<Hex | null>(null)
  const onEpochSelected = useCallback((id: Hex | null, key: { x: bigint; y: bigint } | null) => {
    if (aidRef.current) return
    setEpochId(id)
    setCollectivePubKey(key)
  }, [])

  const onRegistered = useCallback((newAid: Hex, secret: bigint, organizerPK: BabyJubPoint) => {
    aidRef.current = newAid
    setAid(newAid)
    setSkOrg(secret)
    setPkOrg(organizerPK)
    // A new application invalidates anything produced under the previous one.
    setCiphertext(null)
    setSubmittedIndex(null)
    setShare('idle')
  }, [])

  const onSubmitted = useCallback((idx: number) => {
    setSubmittedIndex(idx)
    setShare('idle')
  }, [])

  const hasKey = Boolean(epochId && collectivePubKey)
  const stepWallet: StepStatus = isConnected ? 'done' : 'active'
  const stepEpoch: StepStatus = !isConnected ? 'pending' : hasKey ? 'done' : 'active'
  const stepRegister: StepStatus = !hasKey ? 'pending' : aid ? 'done' : 'active'
  const stepEncrypt: StepStatus = !aid ? 'pending' : ciphertext ? 'done' : 'active'
  const stepSubmit: StepStatus = !ciphertext ? 'pending' : submittedIndex != null ? 'done' : 'active'
  const stepShare: StepStatus =
    submittedIndex == null ? 'pending' : share === 'released' ? 'done' : 'active'
  const stepWatch: StepStatus = submittedIndex == null || share === 'idle' ? 'pending' : 'active'

  return (
    <Stack gap={{ base: 8, md: 12 }}>
      <PageHeader
        title='Playground'
        subtitle='Play the organizer of an application, start to finish: pick up the epoch key the committee has already produced, register an application with your own organizer secret, encrypt a value under the combined key, publish the ciphertext, decide when to release your half of the decryption — and watch the committee open exactly what you sent.'
      />

      <Grid templateColumns={{ base: '1fr', lg: '2fr 1fr' }} gap={{ base: 6, lg: 8 }} alignItems='start'>
        <GridItem>
          <Stack gap={5}>
            <ConnectStep status={stepWallet} />
            <LiveEpochStep
              status={stepEpoch}
              onEpochSelected={onEpochSelected}
              log={addLog}
              pinnedEpochId={aid ? epochId : null}
            />
            <RegisterAppStep
              status={stepRegister}
              epochId={epochId}
              collectivePubKey={collectivePubKey}
              onRegistered={onRegistered}
              log={addLog}
            />
            <EncryptStep
              status={stepEncrypt}
              epochId={epochId}
              collectivePubKey={collectivePubKey}
              pkOrg={pkOrg}
              onEncrypted={(m, ct) => {
                setPlaintext(m)
                setCiphertext(ct)
                setSubmittedIndex(null)
                setShare('idle')
              }}
              log={addLog}
            />
            <SubmitCiphertextStep
              status={stepSubmit}
              epochId={epochId}
              aid={aid}
              ciphertext={ciphertext}
              onSubmitted={onSubmitted}
              log={addLog}
            />
            <ReleaseShareStep
              status={stepShare}
              epochId={epochId}
              aid={aid}
              ciphertextIndex={submittedIndex}
              ciphertext={ciphertext}
              skOrg={skOrg}
              withheld={share === 'withheld'}
              onReleased={() => setShare('released')}
              onWithhold={() => {
                setShare('withheld')
                addLog('Organizer share withheld — partials can land, the combine cannot.', 'info')
              }}
              log={addLog}
            />
            <WatchDecryptionStep
              status={stepWatch}
              epochId={epochId}
              aid={aid}
              ciphertextIndex={submittedIndex}
              ciphertext={ciphertext}
              expectedPlaintext={plaintext}
              shareWithheld={share === 'withheld'}
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
