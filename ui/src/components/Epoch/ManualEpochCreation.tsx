import { useState } from 'react'
import {
  Alert,
  Box,
  Button,
  chakra,
  Collapsible,
  HStack,
  SimpleGrid,
  Stack,
  Text,
} from '@chakra-ui/react'
import { LuChevronRight } from 'react-icons/lu'
import { buildEpochId, type CreateEpochParams, type EpochBounds } from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import {
  MAX_COMMITTEE_SIZE,
  MIN_LOTTERY_ALPHA_BPS,
  PolicyForm,
  defaultPolicyForm,
  validatePolicyForm,
  type PolicyFormState,
} from './PolicyForm'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { useEpochBounds } from '~queries/epochs'
import { HashCell } from '~components/ui/HashCell'
import { NextEpochCountdown } from '~components/Layout/NextEpochCountdown'

const TriggerBtn = chakra('button', {
  base: {
    bg: 'transparent',
    border: 'none',
    p: 0,
    cursor: 'pointer',
    textAlign: 'left',
  },
})

/**
 * Manual `createEpoch`, folded away behind a disclosure.
 *
 * In a running deployment nobody needs this: the committee nodes fire
 * `createEpoch` themselves as soon as the cadence window opens (with jitter, so
 * they don't all pay for the same revert). It stays in the explorer for the
 * cases where that is not true — a local Anvil, a chain whose nodes are all
 * down, a deployment being bootstrapped — and because the parameters it takes
 * are the clearest way to show what an epoch actually is.
 */
export function ManualEpochCreation() {
  const [open, setOpen] = useState(false)
  const writer = useDkgWriter()
  const bounds = useEpochBounds()
  const [form, setForm] = useState<PolicyFormState>(defaultPolicyForm)
  const [busy, setBusy] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [created, setCreated] = useState<{ id: Hex; tx: Hex } | null>(null)

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
      // createEpoch takes exactly these four fields; the phase windows are
      // contract immutables, identical for every epoch on the deployment.
      const policy: CreateEpochParams = {
        threshold: Number(form.threshold),
        committeeSize: Number(form.committeeSize),
        minValidContributions: Number(form.minValidContributions),
        lotteryAlphaBps: Number(form.lotteryAlphaBps),
      }
      const hash = await writer.createEpoch(policy)
      await writer.waitForTransaction(hash)
      const [nonce, prefix] = await Promise.all([writer.epochNonce(), writer.roundPrefix()])
      setCreated({ id: buildEpochId(prefix, nonce), tx: hash })
    } catch (err) {
      setError(err instanceof Error ? err.message : String(err))
    } finally {
      setBusy(false)
    }
  }

  return (
    <Box borderTopWidth='1px' borderColor='rule' pt={5}>
      <TriggerBtn
        type='button'
        onClick={() => setOpen((v) => !v)}
        aria-expanded={open}
        _hover={{ '& .dkg-adv-label': { color: 'accent.fg' } }}
      >
        <HStack gap={2}>
          <Box
            color='ink.4'
            transform={open ? 'rotate(90deg)' : undefined}
            transition='transform 0.15s'
            display='flex'
          >
            <LuChevronRight />
          </Box>
          <Text
            className='dkg-adv-label'
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.06em'
            textTransform='uppercase'
            transition='color 0.15s'
          >
            Advanced: create an epoch manually
          </Text>
        </HStack>
      </TriggerBtn>

      <Collapsible.Root open={open}>
        <Collapsible.Content>
          <Stack
            gap={5}
            mt={4}
            p={{ base: 4, md: 5 }}
            borderWidth='1px'
            borderColor='border.subtle'
            borderRadius='lg'
            bg='surface'
            boxShadow='inset'
          >
            <Stack gap={2}>
              <Text fontSize='sm' color='ink.2' lineHeight='1.6' maxW='68ch'>
                <strong>You almost certainly do not need this.</strong> Epochs are created
                automatically: every committee node watches the cadence and races to call{' '}
                <code>createEpoch</code> the moment the contract allows it, so a new epoch appears
                on schedule whether or not anyone is looking. The call is permissionless but
                cadence-gated — fired early it simply reverts, and only the first of the racing
                calls lands.
              </Text>
              <Text fontSize='sm' color='ink.3' lineHeight='1.6' maxW='68ch'>
                Use this only where no node is running the cadence for you: a local Anvil, a
                fresh deployment, or a chain whose operators are all offline. Applications never
                create epochs — they register against whichever epoch is Live.
              </Text>
              <Box mt={1}>
                <NextEpochCountdown />
              </Box>
            </Stack>

            {created ? (
              <Stack gap={2} fontSize='sm'>
                <Text color='live.fg'>Epoch created.</Text>
                <HashCell value={created.id} head={8} tail={6} />
                <HashCell value={created.tx} head={8} tail={6} />
              </Stack>
            ) : (
              <>
                <PolicyForm
                  value={form}
                  onChange={setForm}
                  disabled={busy || !writer}
                  bounds={bounds.data ?? null}
                />
                <DeploymentBounds bounds={bounds.data ?? null} loading={bounds.isLoading} />
                <HStack gap={3} wrap='wrap'>
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
                    <Text fontSize='xs' color='ink.3'>
                      Connect a wallet to sign the <code>createEpoch</code> transaction — it costs
                      gas on this chain.
                    </Text>
                  )}
                  {writer && validationError && (
                    <Text fontSize='xs' color='warn.fg'>
                      {validationError}
                    </Text>
                  )}
                </HStack>
              </>
            )}

            {error && (
              <Alert.Root status='error'>
                <Alert.Indicator />
                <Alert.Title>{error}</Alert.Title>
              </Alert.Root>
            )}
          </Stack>
        </Collapsible.Content>
      </Collapsible.Root>
    </Box>
  )
}

/**
 * The limits `createEpoch` enforces on this deployment: three `DKGManager`
 * immutables plus the protocol-wide committee cap and α floor.
 */
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
    { label: 'Max committee', value: String(MAX_COMMITTEE_SIZE), hint: 'circuit MaxN' },
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
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
        mb={3}
      >
        Deployment bounds
      </Text>
      <SimpleGrid columns={{ base: 2, md: 4 }} gap={3}>
        {rows.map((r) => (
          <Box key={r.label}>
            <Text fontSize='2xs' color='ink.4'>
              {r.label}
            </Text>
            <Text
              className='dkg-tabular'
              fontFamily='mono'
              fontSize='sm'
              color='ink.0'
              fontWeight={500}
            >
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
