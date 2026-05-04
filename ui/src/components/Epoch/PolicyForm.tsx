import { Alert, Field, NumberInput, SimpleGrid, Stack, Text } from '@chakra-ui/react'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'

export interface PolicyFormState {
  threshold: string
  committeeSize: string
  minValidContributions: string
  lotteryAlphaBps: string
}

export const defaultPolicyForm: PolicyFormState = {
  threshold: '2',
  committeeSize: '3',
  minValidContributions: '2',
  lotteryAlphaBps: '15000',
}

/**
 * Hard cap on committee size, set by the ZK circuits' `MaxN` constant
 * (`circuits/common/sizes.go`). The Solidity contract pins the same value
 * (`MAX_N`) and rejects createEpoch with a higher number; the UI mirrors
 * the cap so users can't even *type* an invalid value.
 *
 * Bumping MaxN requires a coordinated release: new circuits, fresh trusted
 * setup, redeployed verifier contracts, redeployed DKGManager, then
 * updating this constant. Don't change this number in isolation.
 */
export const MAX_COMMITTEE_SIZE = 32

/**
 * Returns a human-readable error string when the form is in an invalid
 * state, or null when it's safe to submit. Catches the gotchas the
 * contract itself doesn't enforce — most importantly
 * `minValidContributions >= threshold`, without which an epoch can finalize
 * but produce a key nobody can decrypt.
 */
export function validatePolicyForm(v: PolicyFormState): string | null {
  const t = Number(v.threshold)
  const n = Number(v.committeeSize)
  const m = Number(v.minValidContributions)
  if (!Number.isFinite(t) || t < 1) return 'Threshold must be at least 1.'
  if (!Number.isFinite(n) || n < 1) return 'Committee size must be at least 1.'
  if (n > MAX_COMMITTEE_SIZE) {
    return `Committee size cannot exceed ${MAX_COMMITTEE_SIZE} — this is the circuit's hard limit (MaxN).`
  }
  if (t > n) return 'Threshold cannot exceed committee size.'
  if (!Number.isFinite(m) || m < 1) return 'Min valid contributions must be at least 1.'
  if (m > n) return 'Min valid contributions cannot exceed committee size.'
  if (m < t) {
    return `Min valid contributions (${m}) must be ≥ threshold (${t}). Otherwise the epoch can finalize but no one will be able to decrypt.`
  }
  return null
}

interface Props {
  value: PolicyFormState
  onChange: (next: PolicyFormState) => void
  disabled?: boolean
}

// Two-tier policy form. The basics most users care about (committee size,
// threshold) live up top. Min-valid-contributions and lottery
// oversubscription sit behind an Advanced disclosure for power users.
//
// Phase budgets (CommitteeSelection / KeyAssembly / FinalizeGap blocks)
// and total epoch duration are *contract immutables* set at deploy time —
// the UI doesn't expose them per-epoch because every caller would have to
// agree on the same numbers anyway, and a per-epoch override would just
// be discarded by the writer.
export function PolicyForm({ value, onChange, disabled }: Props) {
  const set = <K extends keyof PolicyFormState>(key: K, v: PolicyFormState[K]) => onChange({ ...value, [key]: v })

  // Auto-track min valid contributions to the threshold when the user
  // hasn't manually overridden it. This keeps the basic UX two-knobs
  // (committee size + threshold) while still letting power users break
  // the link in the Advanced section if they really want extra
  // redundancy. The "linked" state is detected by simple equality, so
  // any manual edit in Advanced opts out automatically.
  const linked = value.minValidContributions === value.threshold
  const setThreshold = (next: string) => {
    if (linked) onChange({ ...value, threshold: next, minValidContributions: next })
    else onChange({ ...value, threshold: next })
  }

  return (
    <Stack gap={5}>
      {/* ── Basics ──────────────────────────────────────────────────────── */}
      <SimpleGrid columns={{ base: 1, md: 2 }} gap={3}>
        <SmallNumberField
          label='Committee size'
          help={`How many nodes share the key. Capped at ${MAX_COMMITTEE_SIZE} (circuit limit).`}
          value={value.committeeSize}
          onChange={(v) => set('committeeSize', v)}
          disabled={disabled}
          min={1}
          max={MAX_COMMITTEE_SIZE}
        />
        <SmallNumberField
          label='Threshold'
          help='Members needed to decrypt later. By default this is also the minimum number of contributions required for the epoch to go Live (override under Advanced for extra redundancy).'
          value={value.threshold}
          onChange={setThreshold}
          disabled={disabled}
          min={1}
        />
      </SimpleGrid>

      {/* Soft inline warning when the user has broken the link in Advanced
          and ended up with min < threshold — a state the contract sadly
          accepts even though the resulting epoch can finalize without a
          decryptable key. */}
      {Number(value.minValidContributions) < Number(value.threshold) && (
        <Alert.Root status='warning' size='sm'>
          <Alert.Indicator />
          <Alert.Content>
            <Alert.Title>Min valid contributions is below threshold.</Alert.Title>
            <Alert.Description fontSize='xs'>
              The epoch will go Live once {value.minValidContributions} contribution(s) arrive,
              but decryption needs {value.threshold}. With these settings, the epoch can lock in
              a key that nobody can ever use. Raise it under Advanced.
            </Alert.Description>
          </Alert.Content>
        </Alert.Root>
      )}

      <Text fontSize='2xs' color='ink.4' lineHeight='1.55' maxW='62ch'>
        Phase budgets and total epoch duration are fixed by the deployed{' '}
        <code>DKGManager</code> contract — every caller sees the same schedule. The next epoch
        starts automatically once the cadence window elapses; nodes race to fire{' '}
        <code>createEpoch</code> with a small random jitter.
      </Text>

      {/* ── Advanced ────────────────────────────────────────────────────── */}
      <DetailDisclosure title='Advanced configuration'>
        <Stack gap={4} p={1}>
          <Text fontSize='xs' color='ink.3'>
            Fine-grained protocol parameters. Defaults are sensible — touch these only if you have
            a reason.
          </Text>
          <SimpleGrid columns={{ base: 1, md: 2 }} gap={3}>
            <SmallNumberField
              label='Min valid contributions'
              help={
                linked
                  ? 'Auto-tracks threshold (raise it for extra redundancy — e.g. t=3, n=10, min=7 means up to 4 share holders can go offline post-finalize and decryption still works).'
                  : 'How many contributions must arrive before the epoch can go Live. Must be ≥ threshold or the epoch can finalize without a decryptable key.'
              }
              value={value.minValidContributions}
              onChange={(v) => set('minValidContributions', v)}
              disabled={disabled}
              min={1}
            />
            <SmallNumberField
              label='Lottery α (bps)'
              help='Candidate-pool size = α × committee. 10 000 = 1.0×.'
              value={value.lotteryAlphaBps}
              onChange={(v) => set('lotteryAlphaBps', v)}
              disabled={disabled}
            />
          </SimpleGrid>
        </Stack>
      </DetailDisclosure>
    </Stack>
  )
}

function SmallNumberField({
  label,
  help,
  value,
  onChange,
  disabled,
  min,
  max,
}: {
  label: string
  help: string
  value: string
  onChange: (v: string) => void
  disabled?: boolean
  min?: number
  max?: number
}) {
  return (
    <Field.Root disabled={disabled}>
      <Field.Label
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.06em'
        textTransform='uppercase'
      >
        {label}
      </Field.Label>
      <NumberInput.Root
        size='sm'
        value={value}
        min={min}
        max={max}
        onValueChange={(d) => onChange(d.value)}
        disabled={disabled}
        mt={1}
      >
        <NumberInput.Input
          fontFamily='mono'
          bg='surface.sunken'
          borderColor='border.subtle'
          color='ink.0'
          _hover={{ borderColor: 'border' }}
          _focus={{ borderColor: 'accent.fg', boxShadow: 'none' }}
        />
      </NumberInput.Root>
      <Field.HelperText fontSize='2xs' color='ink.3' mt={1.5}>
        {help}
      </Field.HelperText>
    </Field.Root>
  )
}
