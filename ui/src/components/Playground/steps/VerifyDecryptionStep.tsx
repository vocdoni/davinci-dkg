import { useEffect } from 'react'
import { Alert, Box, Spinner, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { useQuery } from '@tanstack/react-query'
import { LuUsers, LuCombine, LuEye } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useEpoch } from '~queries/epochs'
import { Polling } from '~constants/polling'
import { QueryKeys } from '~queries/keys'
import { HowItWorks } from '../HowItWorks'

interface Props {
  status: StepStatus
  epochId: Hex | null
  ciphertextIndex: number | null
  expectedPlaintext: bigint | null
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'chain') => void
}

export function VerifyDecryptionStep({ status, epochId, ciphertextIndex, expectedPlaintext, log }: Props) {
  const { dkg } = useDkgClient()
  const epoch = useEpoch((epochId ?? undefined) as `0x${string}` | undefined)

  // Polls the contract's CombinedDecryption record for the configured
  // ciphertext index. Stops polling automatically once `completed` is true.
  const decryption = useQuery({
    queryKey: epochId && ciphertextIndex ? QueryKeys.decryption(epochId, ciphertextIndex) : ['decryption', 'idle'],
    queryFn: () => {
      if (!epochId || !ciphertextIndex) throw new Error('idle')
      const ZERO_AID = ('0x' + '00'.repeat(32)) as `0x${string}`
      return dkg.getCombinedDecryption(epochId, ZERO_AID, ciphertextIndex)
    },
    enabled: Boolean(epochId && ciphertextIndex),
    refetchInterval: (q) => (q.state.data?.completed ? false : Polling.decryption),
  })

  useEffect(() => {
    if (decryption.data?.completed) {
      log(`Plaintext recovered on-chain: ${decryption.data.plaintext.toString()}`, 'success')
    }
  }, [decryption.data?.completed, decryption.data?.plaintext, log])

  return (
    <StepCard
      n={7}
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
                Pieces collected: {epoch.data.epoch.partialDecryptionCount.toString()} of{' '}
                {epoch.data.epoch.policy.threshold} needed
              </Text>
            )}
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
                Waiting for the committee to finish combining their pieces…
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
              threshold scheme. Each member contributes a partial decryption; once enough have
              arrived, anyone can combine them on-chain and the original number is published.
              Until that combine step happens, the ciphertext stays opaque to the world.
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
