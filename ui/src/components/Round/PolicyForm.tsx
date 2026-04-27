import { Alert, Box, chakra, Field, HStack, NumberInput, SimpleGrid, Stack, Switch, Text } from '@chakra-ui/react'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'

export interface PolicyFormState {
  threshold: string
  committeeSize: string
  minValidContributions: string
  lotteryAlphaBps: string
  seedDelay: string
  regDeadlineOffset: string
  contribDeadlineOffset: string
  finalizeDelayBlocks: string
  disclosureAllowed: boolean
}

export const defaultPolicyForm: PolicyFormState = {
  threshold: '2',
  committeeSize: '3',
  minValidContributions: '2',
  lotteryAlphaBps: '15000',
  seedDelay: '2',
  regDeadlineOffset: '10',
  contribDeadlineOffset: '20',
  finalizeDelayBlocks: '1',
  disclosureAllowed: false,
}

/**
 * Hard cap on committee size, set by the ZK circuits' `MaxN` constant
 * (`circuits/common/sizes.go`). The Solidity contract pins the same value
 * (`MAX_N`) and rejects createRound with a higher number; the UI mirrors
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
 * `minValidContributions >= threshold`, without which a round can finalize
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
    return `Min valid contributions (${m}) must be ≥ threshold (${t}). Otherwise the round can finalize but no one will be able to decrypt.`
  }
  return null
}

// Round-duration presets cover the "I just want to see it work" cases so
// the simple view doesn't expose three separate block-offset knobs. The
// Custom option falls through to the advanced section.
const durationPresets: { id: string; label: string; reg: number; contrib: number; finalize: number }[] = [
  { id: 'quick', label: 'Quick (~2 min)', reg: 5, contrib: 10, finalize: 1 },
  { id: 'default', label: 'Default (~4 min)', reg: 10, contrib: 20, finalize: 1 },
  { id: 'long', label: 'Long (~10 min)', reg: 25, contrib: 50, finalize: 2 },
]

function detectPreset(v: PolicyFormState): string {
  return (
    durationPresets.find(
      (p) =>
        p.reg.toString() === v.regDeadlineOffset &&
        p.contrib.toString() === v.contribDeadlineOffset &&
        p.finalize.toString() === v.finalizeDelayBlocks
    )?.id ?? 'custom'
  )
}

interface Props {
  value: PolicyFormState
  onChange: (next: PolicyFormState) => void
  disabled?: boolean
}

// Two-tier policy form: the four basics most users care about (committee
// size, threshold, minimum contributions, round duration) live up top;
// every other knob — lottery oversubscription, seed delay, individual
// block offsets, secret-key disclosure — sits behind an Advanced
// disclosure so first-time users aren't drowned in ZK protocol jargon.
export function PolicyForm({ value, onChange, disabled }: Props) {
  const set = <K extends keyof PolicyFormState>(key: K, v: PolicyFormState[K]) => onChange({ ...value, [key]: v })

  const presetId = detectPreset(value)
  const applyPreset = (id: string) => {
    if (id === 'custom') return
    const p = durationPresets.find((x) => x.id === id)
    if (!p) return
    onChange({
      ...value,
      regDeadlineOffset: p.reg.toString(),
      contribDeadlineOffset: p.contrib.toString(),
      finalizeDelayBlocks: p.finalize.toString(),
    })
  }

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
          help='Members needed to decrypt later. By default this is also the minimum number of contributions required to finalize the round (override under Advanced for extra redundancy).'
          value={value.threshold}
          onChange={setThreshold}
          disabled={disabled}
          min={1}
        />
      </SimpleGrid>

      {/* Soft inline warning when the user has broken the link in Advanced
          and ended up with min < threshold — a state the contract sadly
          accepts even though the resulting round can finalize without a
          decryptable key. */}
      {Number(value.minValidContributions) < Number(value.threshold) && (
        <Alert.Root status='warning' size='sm'>
          <Alert.Indicator />
          <Alert.Content>
            <Alert.Title>Min valid contributions is below threshold.</Alert.Title>
            <Alert.Description fontSize='xs'>
              The round will finalize once {value.minValidContributions} contribution(s) arrive, but
              decryption needs {value.threshold}. With these settings, the round can lock in a key
              that nobody can ever use. Raise it under Advanced.
            </Alert.Description>
          </Alert.Content>
        </Alert.Root>
      )}

      <Box>
        <Text fontSize='xs' color='gray.400' mb={2}>
          Round duration
        </Text>
        <HStack gap={2} wrap='wrap'>
          {durationPresets.map((p) => (
            <PresetChip
              key={p.id}
              isActive={presetId === p.id}
              onClick={() => applyPreset(p.id)}
              disabled={disabled}
            >
              {p.label}
            </PresetChip>
          ))}
          <PresetChip isActive={presetId === 'custom'} disabled>
            {presetId === 'custom' ? 'Custom (set below)' : 'Custom'}
          </PresetChip>
        </HStack>
        <Text fontSize='2xs' color='gray.500' mt={1.5}>
          Controls registration / contribution / finalize-delay block offsets. Adjust individually
          under Advanced.
        </Text>
      </Box>

      {/* ── Advanced ────────────────────────────────────────────────────── */}
      <DetailDisclosure title='Advanced configuration'>
        <Stack gap={4} p={1}>
          <Text fontSize='xs' color='gray.400'>
            Fine-grained protocol parameters. Defaults are sensible — touch these only if you have a
            reason.
          </Text>
          <SimpleGrid columns={{ base: 1, md: 2 }} gap={3}>
            <SmallNumberField
              label='Min valid contributions'
              help={
                linked
                  ? 'Auto-tracks threshold (raise it for extra redundancy — e.g. t=3, n=10, min=7 means up to 4 share holders can go offline post-finalize and decryption still works).'
                  : 'How many contributions must arrive before finalize is allowed. Must be ≥ threshold or the round can finalize without a decryptable key.'
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
            <SmallNumberField
              label='Seed delay (blocks)'
              help='Block delay until the lottery seed is revealed.'
              value={value.seedDelay}
              onChange={(v) => set('seedDelay', v)}
              disabled={disabled}
            />
            <SmallNumberField
              label='Registration window (blocks)'
              help='How long nodes have to claim slots.'
              value={value.regDeadlineOffset}
              onChange={(v) => set('regDeadlineOffset', v)}
              disabled={disabled}
            />
            <SmallNumberField
              label='Contribution window (blocks)'
              help='How long nodes have to contribute.'
              value={value.contribDeadlineOffset}
              onChange={(v) => set('contribDeadlineOffset', v)}
              disabled={disabled}
            />
            <SmallNumberField
              label='Finalize delay (blocks)'
              help='Wait between contribution close and finalize.'
              value={value.finalizeDelayBlocks}
              onChange={(v) => set('finalizeDelayBlocks', v)}
              disabled={disabled}
            />
          </SimpleGrid>
          <Field.Root disabled={disabled} display='flex' alignItems='center' gap={3}>
            <Switch.Root
              checked={value.disclosureAllowed}
              onCheckedChange={(d) => set('disclosureAllowed', d.checked)}
            >
              <Switch.HiddenInput />
              <Switch.Control />
            </Switch.Root>
            <Text fontSize='sm'>Allow secret-key disclosure after completion</Text>
          </Field.Root>
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
      <Field.Label fontSize='xs'>{label}</Field.Label>
      <NumberInput.Root
        size='sm'
        value={value}
        min={min}
        max={max}
        onValueChange={(d) => onChange(d.value)}
        disabled={disabled}
      >
        <NumberInput.Input fontFamily='mono' />
      </NumberInput.Root>
      <Field.HelperText fontSize='2xs'>{help}</Field.HelperText>
    </Field.Root>
  )
}

// chakra.button is the recommended Chakra v3 way to attach style props to
// a real <button>, avoiding the polymorphic-as-prop type juggling that
// Box's `as` prop demands. It's a plain HTML button under the hood.
const StyledChip = chakra('button', {
  base: {
    px: 3,
    py: 1.5,
    borderRadius: 'md',
    borderWidth: '1px',
    fontSize: 'xs',
    transition: 'border-color 0.15s',
  },
})

function PresetChip({
  isActive,
  disabled,
  onClick,
  children,
}: {
  isActive: boolean
  disabled?: boolean
  onClick?: () => void
  children: React.ReactNode
}) {
  return (
    <StyledChip
      type='button'
      borderColor={isActive ? 'cyan.500' : 'gray.700'}
      bg={isActive ? 'cyan.900' : 'transparent'}
      color={isActive ? 'cyan.200' : 'gray.300'}
      cursor={disabled ? 'not-allowed' : 'pointer'}
      opacity={disabled ? 0.5 : 1}
      onClick={disabled ? undefined : onClick}
      disabled={disabled}
      _hover={!disabled ? { borderColor: 'cyan.400' } : undefined}
    >
      {children}
    </StyledChip>
  )
}
