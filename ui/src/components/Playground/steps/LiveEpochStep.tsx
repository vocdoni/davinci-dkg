import { useEffect, useRef } from 'react'
import { Box, Grid, HStack, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { EpochPhase, type EpochEntry } from '@vocdoni/davinci-dkg-sdk'
import { LuTimer, LuShuffle, LuKeyRound } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { HowItWorks } from '../HowItWorks'
import { StatusBadge } from '~components/Epoch/StatusBadge'
import { Countdown } from '~components/Epoch/Countdown'
import { NextEpochCountdown } from '~components/Layout/NextEpochCountdown'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { useRecentEpochs, useEpoch } from '~queries/epochs'
import { useCollectivePublicKey } from '~queries/applications'
import { useBlockNumber } from '~queries/chain'
import { bigIntToHex, blocksToDuration } from '~lib/format'

interface Props {
  status: StepStatus
  /** Reports the epoch the rest of the walkthrough runs against. */
  onEpochSelected: (epochId: Hex | null, key: { x: bigint; y: bigint } | null) => void
  /** Once an application is registered the walkthrough stays on its epoch. */
  pinnedEpochId?: Hex | null
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'chain' | 'crypto') => void
}

/**
 * Picks the newest Live epoch and shows the key it produced.
 *
 * The playground used to create an epoch here. It no longer does: on a real
 * deployment the committee nodes create epochs themselves on a fixed cadence,
 * and an application never has a reason to — it registers against whichever
 * epoch is already Live. So this step is a read: find that epoch, show its
 * committee, threshold and key, and hand the id to the steps below.
 */
export function LiveEpochStep({ status, onEpochSelected, log, pinnedEpochId }: Props) {
  const recent = useRecentEpochs(10)
  const { data: block } = useBlockNumber()

  const entries = recent.data ?? []
  const live = entries.find((e) => e.epoch.status === EpochPhase.Live)
  const newest = entries[0]
  const selected = pinnedEpochId ?? live?.id
  const epoch = useEpoch(selected)
  const key = useCollectivePublicKey(selected)
  // What the card shows: the pinned epoch (registered application) or the newest Live one.
  const shown: EpochEntry | undefined = pinnedEpochId
    ? (entries.find((e) => e.id === pinnedEpochId) ?? (epoch.data ? { id: pinnedEpochId, epoch: epoch.data.epoch } : undefined))
    : live
  const supersededBy = pinnedEpochId && live && live.id !== pinnedEpochId ? live.id : null

  // Report (epoch, key) upwards once both have landed, and again whenever the
  // selection changes — e.g. the next epoch goes Live while the page is open.
  const reported = useRef<string | null>(null)
  useEffect(() => {
    if (!selected || !key.data) return
    const stamp = `${selected}:${key.data.x}`
    if (reported.current === stamp) return
    reported.current = stamp
    onEpochSelected(selected, key.data)
    log(`Using epoch ${selected} — collective public key read from the contract.`, 'crypto')
  }, [selected, key.data, onEpochSelected, log])

  useEffect(() => {
    if (recent.data && !live && !pinnedEpochId) onEpochSelected(null, null)
  }, [recent.data, live, pinnedEpochId, onEpochSelected])

  return (
    <StepCard
      n={2}
      title='Use the epoch that is live right now'
      status={status}
      description='The committee that will decrypt for you already exists. Operators produce a new epoch on a fixed block cadence and generate its shared key between themselves; you just read the result.'
    >
      {recent.isLoading ? (
        <Text fontSize='sm' color='ink.4'>
          Looking for the newest Live epoch…
        </Text>
      ) : !shown ? (
        <NoLiveEpoch newest={newest} />
      ) : (
        <Stack gap={5}>
          {supersededBy && (
            <Text fontSize='sm' color='ink.3'>
              Pinned to the epoch your application is registered on. A newer epoch (
              <HashCell value={supersededBy} head={8} tail={6} />) has gone live meanwhile; reload the page to start
              over on it.
            </Text>
          )}
          <HStack gap={4} wrap='wrap' align='center'>
            <StatusBadge epoch={shown.epoch} />
            <HashCell value={shown.id} head={8} tail={6} />
            <Text fontFamily='mono' fontSize='2xs' color='ink.4' letterSpacing='0.06em'>
              NONCE {shown.epoch.nonce.toString()}
            </Text>
          </HStack>

          <SimpleGrid columns={{ base: 2, md: 4 }} gap={3}>
            <Fact
              label='Threshold'
              value={`${shown.epoch.policy.threshold} of ${shown.epoch.policy.committeeSize}`}
              hint='members needed to decrypt'
            />
            <Fact
              label='Committee'
              value={`${shown.epoch.claimedCount}`}
              hint='slots claimed via lottery'
            />
            <Fact
              label='Live since'
              value={
                block
                  ? `${blocksToDuration(Number(block - shown.epoch.policy.liveNotBeforeBlock))} ago`
                  : '—'
              }
              hint={`block #${shown.epoch.policy.liveNotBeforeBlock.toString()}`}
            />
            <Fact
              label='Ciphertexts'
              value={shown.epoch.ciphertextCount.toString()}
              hint='submitted so far'
            />
          </SimpleGrid>

          {epoch.data && epoch.data.participants.length > 0 && (
            <Stack gap={2}>
              <Text
                fontFamily='mono'
                fontSize='2xs'
                color='ink.3'
                letterSpacing='0.08em'
                textTransform='uppercase'
              >
                Committee members
              </Text>
              <Grid templateColumns={{ base: '1fr', md: '1fr 1fr' }} gap={1.5}>
                {epoch.data.participants.map((addr, i) => (
                  <HStack key={addr} gap={3}>
                    <Text
                      className='dkg-tabular'
                      fontFamily='mono'
                      fontSize='2xs'
                      color='accent.dim'
                    >
                      {(i + 1).toString().padStart(2, '0')}
                    </Text>
                    <HashCell value={addr} head={6} tail={6} />
                  </HStack>
                ))}
              </Grid>
            </Stack>
          )}

          {key.data ? (
            <Stack gap={2}>
              <Text fontSize='sm' color='live.fg'>
                Shared encryption key is on chain — you can register an application against it.
              </Text>
              <DetailDisclosure title='Show key coordinates'>
                <Stack gap={1}>
                  <Text fontSize='2xs' color='ink.4'>
                    PK_ep on BabyJubJub (twisted Edwards, BN254 scalar field).
                  </Text>
                  <Text>x:</Text>
                  <HashCell value={bigIntToHex(key.data.x)} head={6} tail={6} />
                  <Text>y:</Text>
                  <HashCell value={bigIntToHex(key.data.y)} head={6} tail={6} />
                </Stack>
              </DetailDisclosure>
            </Stack>
          ) : (
            <Text fontSize='sm' color='ink.4'>
              Reading the collective public key…
            </Text>
          )}

          <HowItWorks
            body={
              <>
                Nothing in this step is yours to do. A lottery picked this committee out of the
                registered operators, each member published one piece of a key nobody holds in
                full, and one of them paid for the proof that finalized it. The result is a
                single public key that stays usable for the rest of the epoch — and a fresh
                committee takes over in the next one, without any handover.
              </>
            }
            flow={[
              { icon: <LuTimer />, label: 'Cadence opens an epoch' },
              { icon: <LuShuffle />, label: 'Lottery picks the committee' },
              { icon: <LuKeyRound />, label: 'Shared key goes live' },
            ]}
          />
        </Stack>
      )}
    </StepCard>
  )
}

/** No Live epoch: say where the newest one is and when it will be usable. */
function NoLiveEpoch({ newest }: { newest?: EpochEntry }) {
  if (!newest) {
    return (
      <Stack gap={3}>
        <Text fontSize='sm' color='ink.2'>
          No epoch exists on this deployment yet. Epochs are created by the committee nodes
          themselves once the cadence window opens — start a node, or wait for one to fire.
        </Text>
        <NextEpochCountdown />
      </Stack>
    )
  }

  const phase = newest.epoch.status
  const waiting =
    phase === EpochPhase.CommitteeSelection
      ? 'the lottery is still handing out committee slots'
      : phase === EpochPhase.KeyAssembly
        ? 'the committee is publishing its contributions'
        : phase === EpochPhase.Aborted
          ? 'the newest epoch was aborted — the next one opens on schedule'
          : 'the newest epoch is not usable yet'

  return (
    <Stack gap={4}>
      <HStack gap={4} wrap='wrap'>
        <StatusBadge epoch={newest.epoch} />
        <HashCell value={newest.id} head={8} tail={6} />
      </HStack>
      <Text fontSize='sm' color='ink.2' lineHeight='1.6' maxW='68ch'>
        No epoch is Live at the moment: {waiting}. You do not have to do anything — operators
        create and finalize epochs automatically, and this step unlocks by itself as soon as one
        goes Live.
      </Text>
      {phase !== EpochPhase.Aborted ? (
        <Countdown
          target={newest.epoch.policy.liveNotBeforeBlock}
          label='until this epoch goes Live'
        />
      ) : (
        <NextEpochCountdown />
      )}
    </Stack>
  )
}

function Fact({ label, value, hint }: { label: string; value: string; hint: string }) {
  return (
    <Box position='relative' pl={3}>
      <Box
        position='absolute'
        left={0}
        top='15%'
        bottom='15%'
        w='2px'
        bg='accent.fg'
        opacity={0.4}
        borderRightRadius='full'
      />
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
      >
        {label}
      </Text>
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize='md'
        fontWeight={500}
        color='ink.0'
        lineHeight='1.3'
      >
        {value}
      </Text>
      <Text fontSize='2xs' color='ink.4'>
        {hint}
      </Text>
    </Box>
  )
}
