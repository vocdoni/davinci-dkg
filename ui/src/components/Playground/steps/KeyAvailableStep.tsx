import { useEffect, useState } from 'react'
import { Box, Spinner, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { RoundStatus } from '@vocdoni/davinci-dkg-sdk'
import { LuPuzzle, LuCombine, LuKeyRound } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useRound } from '~queries/rounds'
import { useDkgClient } from '~hooks/use-dkg-client'
import { Countdown } from '~components/Round/Countdown'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { bigIntToHex } from '~lib/format'

interface Props {
  status: StepStatus
  roundId: Hex | null
  onKeyReady: (key: { x: bigint; y: bigint }) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'crypto') => void
}

// Reads the on-chain collective public key once the round reaches Finalized.
// Cf. UI_PLAN.md §4 / sdk/src/client.ts:565 — we deliberately do NOT surface
// this key during the Contribution phase (the on-chain accumulator carries
// it, but submitCiphertext requires Finalized status).
export function KeyAvailableStep({ status, roundId, onKeyReady, log }: Props) {
  const { dkg } = useDkgClient()
  const round = useRound((roundId ?? undefined) as `0x${string}` | undefined)
  const [key, setKey] = useState<{ x: bigint; y: bigint } | null>(null)
  const [loading, setLoading] = useState(false)

  const isFinalized =
    round.data &&
    (round.data.round.status === RoundStatus.Finalized || round.data.round.status === RoundStatus.Completed)

  useEffect(() => {
    if (!roundId || !isFinalized || key || loading) return
    setLoading(true)
    dkg
      .getCollectivePublicKey(roundId)
      .then((pk) => {
        setKey(pk)
        onKeyReady(pk)
        log('Collective public key read from on-chain accumulator.', 'crypto')
      })
      .catch((err) => {
        log(`Failed to read collective public key: ${err instanceof Error ? err.message : String(err)}`, 'error')
      })
      .finally(() => setLoading(false))
  }, [roundId, isFinalized, key, loading, dkg, onKeyReady, log])

  return (
    <StepCard
      n={4}
      title='The shared encryption key is ready'
      status={status}
      description='Once enough committee members have contributed, the round finalizes and a single public key — built from all their pieces — becomes available.'
    >
      {!roundId || !round.data ? (
        <Text fontSize='sm' color='ink.4'>
          Create a round and wait for it to be finalized.
        </Text>
      ) : !isFinalized ? (
        <Stack gap={3}>
          <Text fontSize='sm' color='ink.2'>
            Hang tight — the committee is still building the key. The encrypt step will unlock as
            soon as the round finalizes.
          </Text>
          <Countdown
            target={round.data.round.policy.finalizeNotBeforeBlock}
            label='until finalize unlocks'
          />
          <HowItWorks
            body={
              <>
                Why wait? The contract intentionally only exposes the public key after the
                committee has formally locked it in. If the page used a partial key, anything you
                encrypted with it would be useless — the on-chain decryption flow needs the full
                committee setup to succeed.
              </>
            }
            flow={[
              { icon: <LuPuzzle />, label: 'Collect key pieces' },
              { icon: <LuCombine />, label: 'Combine into one key' },
              { icon: <LuKeyRound />, label: 'Public key ready' },
            ]}
          />
        </Stack>
      ) : loading ? (
        <Stack gap={2} align='start'>
          <Spinner size='sm' color='accent.fg' />
          <Text fontSize='xs' color='ink.4'>
            Reading the public key…
          </Text>
        </Stack>
      ) : key ? (
        <Stack gap={3}>
          <Box>
            <Text fontSize='sm' color='ink.2'>
              The shared encryption key is live. You can now encrypt a value for the committee to
              decrypt — head to the next step.
            </Text>
          </Box>
          <DetailDisclosure title='Show key coordinates'>
            <Stack gap={1}>
              <Text fontSize='2xs' color='ink.4'>
                BabyJubJub (twisted Edwards / BN254 scalar field). x and y are 254-bit field elements.
              </Text>
              <Text>x:</Text>
              <HashCell value={bigIntToHex(key.x)} head={6} tail={6} />
              <Text>y:</Text>
              <HashCell value={bigIntToHex(key.y)} head={6} tail={6} />
            </Stack>
          </DetailDisclosure>
        </Stack>
      ) : null}
    </StepCard>
  )
}
