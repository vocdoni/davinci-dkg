import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import { useAccount } from 'wagmi'
import { LuWallet, LuSignature, LuLink } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { ConnectButton } from '~components/Layout/ConnectButton'
import { HashCell } from '~components/ui/HashCell'
import { HowItWorks } from '../HowItWorks'

export function ConnectStep({ status }: { status: StepStatus }) {
  const { address, isConnected, chain } = useAccount()
  return (
    <StepCard
      n={1}
      title='Connect a wallet'
      status={status}
      description='You need a wallet to create a round and submit a ciphertext. Browsing the explorer works without one.'
    >
      <Stack gap={4}>
        {!isConnected ? (
          <ConnectButton />
        ) : (
          <Box bg='gray.950' borderWidth='1px' borderColor='gray.800' borderRadius='md' p={3}>
            <HStack gap={4} fontSize='sm' wrap='wrap'>
              <Text color='gray.400'>Connected:</Text>
              <HashCell value={address} head={6} tail={6} />
              <Text color='gray.500' fontSize='xs'>
                on {chain?.name ?? 'unknown chain'}
              </Text>
            </HStack>
          </Box>
        )}
        <HowItWorks
          body={
            <>
              Your wallet signs the two transactions later in this walkthrough — one to create
              the round, and one to publish the ciphertext for the committee to decrypt. Nothing
              is sent on-chain yet; the connect step just hands the page an address it can sign
              from.
            </>
          }
          flow={[
            { icon: <LuWallet />, label: 'Your wallet' },
            { icon: <LuSignature />, label: 'Signs transactions' },
            { icon: <LuLink />, label: 'Sent to chain' },
          ]}
        />
      </Stack>
    </StepCard>
  )
}
