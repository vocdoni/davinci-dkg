import { useEffect } from 'react'
import { Alert, Box, Spinner, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { useQuery } from '@tanstack/react-query'
import { LuUsers, LuCombine, LuEye } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useEpoch } from '~queries/epochs'
import { useOrganizerShareHash, ZERO_BYTES32 } from '~queries/applications'
import { Polling } from '~constants/polling'
import { QueryKeys } from '~queries/keys'
import { HowItWorks } from '../HowItWorks'

interface Props {
  status: StepStatus
  epochId: Hex | null
  aid: Hex | null
  ciphertextIndex: number | null
  expectedPlaintext: bigint | null
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'chain') => void
}

export function VerifyDecryptionStep({
  status,
  epochId,
  aid,
  ciphertextIndex,
  expectedPlaintext,
  log,
}: Props) {
  const { dkg } = useDkgClient()
  const epoch = useEpoch((epochId ?? undefined) as `0x${string}` | undefined)

  // Polls the contract's CombinedDecryption record for the configured
  // ciphertext index. Stops polling automatically once `completed` is true.
  const decryption = useQuery({
    queryKey: epochId && ciphertextIndex ? QueryKeys.decryption(epochId, ciphertextIndex) : ['decryption', 'idle'],
    queryFn: () => {
      if (!epochId || !aid || !ciphertextIndex) throw new Error('idle')
      return dkg.getCombinedDecryption(epochId, aid, ciphertextIndex)
    },
    enabled: Boolean(epochId && aid && ciphertextIndex),
    refetchInterval: (q) => (q.state.data?.completed ? false : Polling.decryption),
  })

  // The organizer share is the other precondition for combining; surfacing it
  // separately is what tells "the committee is slow" apart from "nobody has
  // released the share".
  const share = useOrganizerShareHash(
    (epochId ?? undefined) as `0x${string}` | undefined,
    (aid ?? undefined) as `0x${string}` | undefined,
    ciphertextIndex ?? undefined,
  )
  const shareReleased = Boolean(share.data && share.data !== ZERO_BYTES32)

  useEffect(() => {
    if (decryption.data?.completed) {
      log(`Plaintext recovered on-chain: ${decryption.data.plaintext.toString()}`, 'success')
    }
  }, [decryption.data?.completed, decryption.data?.plaintext, log])

  return (
    <StepCard
      n={9}
      title='Confirm the committee recovered your message'
      status={status}
      description='Each committee member contributes a piece of the decryption. Once enough pieces arrive, the original number is reconstructed on-chain.'
    >
      <Stack gap={4}>
        {!ciphertextIndex ? (
          <Text fontSize='sm' color='ink.4'>
            Publish a ciphertext to start the decryption flow.
          </Text>
        ) : (
          <Stack gap={3}>
            {epoch.data && (
              <Text fontSize='xs' color='ink.3'>
                Committee pieces collected: {epoch.data.epoch.partialDecryptionCount.toString()} of{' '}
                {epoch.data.epoch.policy.threshold} needed
              </Text>
            )}
            <Text fontSize='xs' color={shareReleased ? 'live.fg' : 'ink.3'}>
              Organizer share: {shareReleased ? 'released' : 'not on chain yet'}
            </Text>
            {decryption.isLoading && (
              <Stack gap={2} align='start'>
                <Spinner size='sm' color='accent.fg' />
                <Text fontSize='xs' color='ink.4'>
                  Checking the contract…
                </Text>
              </Stack>
            )}
            {decryption.data && !decryption.data.completed && (
              <Text fontSize='xs' color='ink.4'>
                {shareReleased
                  ? 'Waiting for the committee to finish combining their pieces…'
                  : 'Waiting for the organizer share — combining reverts until it is on chain.'}
              </Text>
            )}
            {decryption.data?.completed && (
              <Box>
                {expectedPlaintext != null && decryption.data.plaintext === expectedPlaintext ? (
                  <Alert.Root status='success'>
                    <Alert.Indicator />
                    <Alert.Title>
                      Recovered: {decryption.data.plaintext.toString()} — matches what you encrypted ✓
                    </Alert.Title>
                  </Alert.Root>
                ) : (
                  <Alert.Root status='warning'>
                    <Alert.Indicator />
                    <Alert.Title>
                      Recovered: {decryption.data.plaintext.toString()}
                      {expectedPlaintext != null && ` (expected ${expectedPlaintext.toString()})`}
                    </Alert.Title>
                  </Alert.Root>
                )}
              </Box>
            )}
          </Stack>
        )}
        <HowItWorks
          body={
            <>
              No single committee member can decrypt on their own — that's the whole point of a
              threshold scheme. Each member contributes a partial decryption, and the
              application's organizer contributes theirs. Once enough committee pieces and the
              organizer share are on-chain, anyone can combine them and the original number is
              published. Until that combine step happens, the ciphertext stays opaque to the
              world.
            </>
          }
          flow={[
            { icon: <LuUsers />, label: 'Members publish partials' },
            { icon: <LuCombine />, label: 'Pieces combined on-chain' },
            { icon: <LuEye />, label: 'Plaintext revealed' },
          ]}
        />
      </Stack>
    </StepCard>
  )
}
