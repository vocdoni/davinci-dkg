import { useState } from 'react'
import { Alert, Box, Button, Heading, Stack, Text } from '@chakra-ui/react'
import { LuPenLine, LuUsers, LuShuffle } from 'react-icons/lu'
import { buildEpochId, type DecryptionPolicy, type EpochPolicy } from '@vocdoni/davinci-dkg-sdk'
import { StepCard, type StepStatus } from '../StepCard'
import { PolicyForm, defaultPolicyForm, validatePolicyForm, type PolicyFormState } from '~components/Epoch/PolicyForm'
import {
  DecryptionPolicyForm,
  defaultDecryptionPolicyForm,
  type DecryptionPolicyFormState,
} from '~components/Epoch/DecryptionPolicyForm'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { HowItWorks } from '../HowItWorks'
import { useDkgWriter } from '~hooks/use-dkg-writer'
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
  const [form, setForm] = useState<PolicyFormState>(defaultPolicyForm)
  const [dpForm, setDpForm] = useState<DecryptionPolicyFormState>(defaultDecryptionPolicyForm)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [txHash, setTxHash] = useState<Hex | null>(null)

  // Validate the form on every render so the submit button reflects the
  // current state without an extra useEffect roundtrip. Cheap pure call.
  const validationError = validatePolicyForm(form)

  const onCreate = async () => {
    if (!writer) return
    if (validationError) {
      setError(validationError)
      return
    }
    setBusy(true)
    setError(null)
    try {
      // Phase deadline blocks are derived ON-CHAIN from the contract's
      // immutable EPOCH_DURATION_BLOCKS plus per-phase BPS constants. The
      // form fields below are unused by the writer but stay on PolicyFormState
      // for backward UI-test compatibility — leave them at 0n.
      const policy: EpochPolicy = {
        threshold: Number(form.threshold),
        committeeSize: Number(form.committeeSize),
        minValidContributions: Number(form.minValidContributions),
        lotteryAlphaBps: Number(form.lotteryAlphaBps),
        registrationDeadlineBlock: 0n,
        contributionDeadlineBlock: 0n,
        finalizeNotBeforeBlock: 0n,
      }
      const dp: DecryptionPolicy = {
        ownerOnly: dpForm.ownerOnly,
        maxDecryptions: Number(dpForm.maxDecryptions || '0'),
        notBeforeBlock: BigInt(dpForm.notBeforeBlock || '0'),
        notBeforeTimestamp: BigInt(dpForm.notBeforeTimestamp || '0'),
        notAfterBlock: BigInt(dpForm.notAfterBlock || '0'),
        notAfterTimestamp: BigInt(dpForm.notAfterTimestamp || '0'),
      }

      const currentBlock = await writer.blockNumber()
      log(`Creating epoch at block #${currentBlock} (t=${policy.threshold} of n=${policy.committeeSize})`, 'info')
      const hash = await writer.createEpoch(policy, dp)
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
              <PolicyForm value={form} onChange={setForm} disabled={busy || !writer} />
            </Box>
            <DetailDisclosure title='Restrict who can publish ciphertexts (optional)'>
              <Stack gap={3} p={1}>
                <Text fontSize='xs' color='ink.3'>
                  By default, anyone can publish a ciphertext for this epoch to be decrypted. Set
                  these limits if you only want yourself (or a specific time window, or a maximum
                  count) to be allowed. Leave everything blank for an open epoch.
                </Text>
                <DecryptionPolicyForm value={dpForm} onChange={setDpForm} disabled={busy || !writer} />
              </Stack>
            </DetailDisclosure>
            <HowItWorks
              body={
                <>
                  This step writes the epoch's rules onto the contract. As soon as the epoch
                  exists, the registry runs a small lottery and picks {form.committeeSize}{' '}
                  committee members from the active nodes. Those nodes will spend the next minutes
                  collaboratively building one shared encryption key — without any single one of
                  them ever knowing the matching private key.
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
