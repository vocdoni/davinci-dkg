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
      title='Connect a wallet to act as the organizer'
      status={status}
      description='Three transactions in this walkthrough come from your wallet: registering your application, publishing the ciphertext, and releasing your organizer share. Reading the explorer needs no wallet at all.'
    >
      <Stack gap={4}>
        {!isConnected ? (
          <ConnectButton />
        ) : (
          <Box bg='canvas.deep' borderWidth='1px' borderColor='border.subtle' borderRadius='md' p={3}>
            <HStack gap={4} fontSize='sm' wrap='wrap'>
              <Text color='ink.3'>Connected:</Text>
              <HashCell value={address} head={6} tail={6} />
              <Text color='ink.4' fontSize='xs'>
                on {chain?.name ?? 'unknown chain'}
              </Text>
            </HStack>
          </Box>
        )}
        <HowItWorks
          body={
            <>
              Your wallet is the application's identity here: the address you connect with becomes
              the organizer that registers the application and, by default, the only address
              allowed to submit ciphertexts under it. It signs three transactions in all —
              register, submit, release — and pays their gas. You never create an epoch: the
              committee nodes do that on their own cadence. Nothing is sent on-chain in this
              step; it just hands the page an address it can sign from.
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
