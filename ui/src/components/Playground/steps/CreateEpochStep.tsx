import { useState } from 'react'
import { Alert, Box, Button, Heading, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import { LuPenLine, LuUsers, LuShuffle } from 'react-icons/lu'
import { buildEpochId, type CreateEpochParams, type EpochBounds } from '@vocdoni/davinci-dkg-sdk'
import { StepCard, type StepStatus } from '../StepCard'
import {
  MAX_COMMITTEE_SIZE,
  MIN_LOTTERY_ALPHA_BPS,
  PolicyForm,
  defaultPolicyForm,
  validatePolicyForm,
  type PolicyFormState,
} from '~components/Epoch/PolicyForm'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { HowItWorks } from '../HowItWorks'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { useEpochBounds } from '~queries/epochs'
import { HashCell } from '~components/ui/HashCell'
import type { Hex } from 'viem'

interface Props {
  status: StepStatus
  epochId: Hex | null
  setRoundId: (id: Hex) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx') => void
}

export function CreateEpochStep({ status, epochId, setRoundId, log }: Props) {
  const writer = useDkgWriter()
  const bounds = useEpochBounds()
  const [form, setForm] = useState<PolicyFormState>(defaultPolicyForm)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [txHash, setTxHash] = useState<Hex | null>(null)

  // Validate the form on every render so the submit button reflects the
  // current state without an extra useEffect roundtrip. Cheap pure call.
  // The deployment bounds arrive asynchronously; until then only the
  // contract-wide invariants are checked.
  const validationError = validatePolicyForm(form, bounds.data ?? null)

  const onCreate = async () => {
    if (!writer) return
    if (validationError) {
      setError(validationError)
      return
    }
    setBusy(true)
    setError(null)
    try {
      // Only these four fields exist on createEpoch. Phase deadline blocks
      // are derived ON-CHAIN from the contract's immutable
      // EPOCH_DURATION_BLOCKS plus per-phase BPS constants, and there is no
      // per-epoch decryption policy any more — ciphertext submission is
      // gated per application instead.
      const policy: CreateEpochParams = {
        threshold: Number(form.threshold),
        committeeSize: Number(form.committeeSize),
        minValidContributions: Number(form.minValidContributions),
        lotteryAlphaBps: Number(form.lotteryAlphaBps),
      }

      const currentBlock = await writer.blockNumber()
      log(`Creating epoch at block #${currentBlock} (t=${policy.threshold} of n=${policy.committeeSize})`, 'info')
      const hash = await writer.createEpoch(policy)
      setTxHash(hash)
      log(`createEpoch tx submitted: ${hash}`, 'tx')

      const receipt = await writer.waitForTransaction(hash)
      log(`Mined in block #${receipt.blockNumber} (gas ${receipt.gasUsed.toString()})`, 'tx')

      const [nonce, prefix] = await Promise.all([writer.epochNonce(), writer.roundPrefix()])
      const id = buildEpochId(prefix, nonce)
      setRoundId(id)
      log(`Epoch created: ${id}`, 'success')
    } catch (err) {
      const msg = err instanceof Error ? err.message : String(err)
      setError(msg)
      log(`createEpoch failed: ${msg}`, 'error')
    } finally {
      setBusy(false)
    }
  }

  return (
    <StepCard
      n={2}
      title='Create a DKG epoch'
      status={status}
      description='Pick how many committee members will share the key and how many of them are needed to decrypt later.'
    >
      <Stack gap={5}>
        {!epochId ? (
          <Stack gap={5}>
            <Box>
              <Heading size='xs' mb={3} color='ink.2'>
                Epoch configuration
              </Heading>
              <PolicyForm value={form} onChange={setForm} disabled={busy || !writer} bounds={bounds.data ?? null} />
            </Box>
            <DeploymentBounds bounds={bounds.data ?? null} loading={bounds.isLoading} />
            <HowItWorks
              body={
                <>
                  This step writes the epoch's rules onto the contract. As soon as the epoch
                  exists, the registry runs a small lottery and picks {form.committeeSize}{' '}
                  committee members from the active nodes. Those nodes will spend the next minutes
                  collaboratively building one shared encryption key — without any single one of
                  them ever knowing the matching private key. Anyone may create an epoch once the
                  cadence window allows it.
                </>
              }
              flow={[
                { icon: <LuPenLine />, label: 'Configure rules' },
                { icon: <LuShuffle />, label: 'Lottery picks committee' },
                { icon: <LuUsers />, label: '{n} nodes claim slots' },
              ]}
            />
            <Box>
              <Button
                colorPalette='cyan'
                size='sm'
                onClick={onCreate}
                loading={busy}
                disabled={!writer || busy || validationError !== null}
              >
                Create epoch →
              </Button>
              {!writer && (
                <Text fontSize='xs' color='ink.4' mt={2}>
                  Connect a wallet first.
                </Text>
              )}
              {writer && validationError && (
                <Text fontSize='xs' color='orange.300' mt={2}>
                  Fix the validation issue above before submitting.
                </Text>
              )}
            </Box>
          </Stack>
        ) : (
          <Stack gap={2} fontSize='sm'>
            <Text color='live.fg'>Epoch created successfully.</Text>
            <Box>
              <Text fontSize='xs' color='ink.4'>
                Epoch ID
              </Text>
              <HashCell value={epochId} head={8} tail={6} />
            </Box>
            {txHash && (
              <DetailDisclosure title='Show transaction hash'>
                <HashCell value={txHash} full />
              </DetailDisclosure>
            )}
          </Stack>
        )}
        {error && (
          <Alert.Root status='error'>
            <Alert.Indicator />
            <Alert.Title>{error}</Alert.Title>
          </Alert.Root>
        )}
      </Stack>
    </StepCard>
  )
}

// The limits `createEpoch` enforces on this deployment. Three of them are
// constructor immutables of DKGManager (read on-chain); the α floor and the
// committee cap are protocol-wide. Shown in place of the old per-epoch
// decryption-policy form, which no longer exists — ciphertext submission
// is gated per application via AppPolicy.
function DeploymentBounds({ bounds, loading }: { bounds: EpochBounds | null; loading: boolean }) {
  const rows: { label: string; value: string; hint: string }[] = [
    {
      label: 'Min threshold',
      value: bounds ? String(bounds.minThreshold) : loading ? '…' : '—',
      hint: 'MIN_THRESHOLD',
    },
    {
      label: 'Min committee',
      value: bounds ? String(bounds.minCommitteeSize) : loading ? '…' : '—',
      hint: 'MIN_COMMITTEE_SIZE',
    },
    {
      label: 'Max committee',
      value: String(MAX_COMMITTEE_SIZE),
      hint: 'circuit MaxN',
    },
    {
      label: 'Lottery α range',
      value: bounds
        ? `${MIN_LOTTERY_ALPHA_BPS}–${bounds.maxLotteryAlphaBps} bps`
        : loading
          ? '…'
          : `≥ ${MIN_LOTTERY_ALPHA_BPS} bps`,
      hint: 'MAX_LOTTERY_ALPHA_BPS',
    },
  ]
  return (
    <Box borderWidth='1px' borderColor='border.subtle' borderRadius='lg' bg='surface.sunken' p={4}>
      <Text fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase' mb={3}>
        Deployment bounds
      </Text>
      <SimpleGrid columns={{ base: 2, md: 4 }} gap={3}>
        {rows.map((r) => (
          <Box key={r.label}>
            <Text fontSize='2xs' color='ink.4'>
              {r.label}
            </Text>
            <Text className='dkg-tabular' fontFamily='mono' fontSize='sm' color='ink.0' fontWeight={500}>
              {r.value}
            </Text>
            <Text fontFamily='mono' fontSize='2xs' color='ink.4'>
              {r.hint}
            </Text>
          </Box>
        ))}
      </SimpleGrid>
      <Text fontSize='2xs' color='ink.4' mt={3} lineHeight='1.55' maxW='62ch'>
        Fixed when the <code>DKGManager</code> was deployed; <code>createEpoch</code> reverts with{' '}
        <code>InvalidPolicy()</code> outside them. The contract also requires threshold ≤ min valid
        contributions ≤ committee size.
      </Text>
    </Box>
  )
}
