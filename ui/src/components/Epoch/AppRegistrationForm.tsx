import { useMemo, useState } from 'react'
import {
  Box,
  Button,
  Field,
  Grid,
  GridItem,
  HStack,
  Input,
  RadioGroup,
  Stack,
  Text,
} from '@chakra-ui/react'
import { keccak256, toHex, type Hex } from 'viem'
import {
  AppMode,
  computeS,
  proveOrganizer,
  type AppPolicy,
} from '@vocdoni/davinci-dkg-sdk'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'

// AppRegistrationForm — register a per-application correction key against a
// finalized epoch (paper §4.3).
//
//   Mode 0 (public derivation):    PK_aid = PK_ep + S·G,  S = H(eid||PK_ep||aid).
//   Mode 1 (organizer co-decryption): PK_aid = PK_ep + PK_org. Caller supplies
//                                     `sk_org` so we can derive PK_org and the
//                                     Schnorr proof of knowledge in-browser.
//
// The form is intentionally lean: a real organizer will compute `aid` from
// their own application's namespace (e.g. ballot id), and may want to sign
// `sk_org` from a hardware wallet — both are out of scope. This component
// is the canonical demo + reference integration for SDK consumers.

interface Props {
  /** Hex-encoded bytes12 epoch id (the URL param). */
  epochId: Hex
  /** Epoch collective public key in TE coords; `null` until finalized. */
  pkEp: { x: bigint; y: bigint } | null
  /** Optional initial `aid`. Defaults to a fresh random bytes32. */
  initialAid?: Hex
  /** Called with the on-chain tx hash on a successful registration. */
  onRegistered?: (txHash: Hex, aid: Hex) => void
}

type Mode = 0 | 1

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as const

function defaultAid(): Hex {
  // Random aid so users can re-submit without colliding while testing.
  const rand = new Uint8Array(32)
  globalThis.crypto.getRandomValues(rand)
  return ('0x' + Array.from(rand).map((b) => b.toString(16).padStart(2, '0')).join('')) as Hex
}

export function AppRegistrationForm({ epochId, pkEp, initialAid, onRegistered }: Props) {
  const writer = useDkgWriter()

  const [mode, setMode] = useState<Mode>(0)
  const [aid, setAid] = useState<Hex>(initialAid ?? defaultAid())
  const [organizerSecret, setOrganizerSecret] = useState<string>('')
  const [maxCiphertexts, setMaxCiphertexts] = useState<string>('0')
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [lastTx, setLastTx] = useState<Hex | null>(null)

  // S preview (mode 0). We re-derive on every keystroke; the operation is
  // a single keccak256, well under a millisecond.
  const sPreview = useMemo<bigint | null>(() => {
    if (mode !== 0 || !pkEp || pkEp.x === 0n) return null
    try {
      return computeS(epochId, pkEp.x, pkEp.y, aid)
    } catch {
      return null
    }
  }, [mode, pkEp, epochId, aid])

  // PK_org preview (mode 1). Derived from the organizer secret as a
  // visible sanity check before submission.
  const pkOrgPreview = useMemo<{ x: bigint; y: bigint } | null>(() => {
    if (mode !== 1) return null
    if (!organizerSecret) return null
    let sk: bigint
    try {
      sk = organizerSecret.startsWith('0x')
        ? BigInt(organizerSecret)
        : BigInt(organizerSecret)
    } catch {
      return null
    }
    if (sk === 0n) return null
    try {
      // Use the convenience prover with a fixed nonce purely for preview;
      // the real submission redraws a fresh nonce at submit-time.
      const { pkOrgX, pkOrgY } = proveOrganizer(sk, epochId, aid, 1n)
      return { x: pkOrgX, y: pkOrgY }
    } catch {
      return null
    }
  }, [mode, organizerSecret, epochId, aid])

  async function handleSubmit() {
    if (!writer) {
      setError('Connect a wallet first.')
      return
    }
    setSubmitting(true)
    setError(null)
    setLastTx(null)
    try {
      const policy: AppPolicy = {
        authorizedSubmitter: ZERO_ADDRESS,
        maxCiphertexts: Math.max(0, Math.min(65535, Number(maxCiphertexts) || 0)),
        notBeforeBlock: 0n,
        notAfterBlock: 0n,
      }
      let tx: Hex
      if (mode === AppMode.PublicDerivation) {
        tx = await writer.registerApplication(epochId, aid, policy)
      } else {
        if (!organizerSecret) throw new Error('Organizer secret is required for mode 1.')
        const sk = BigInt(organizerSecret)
        if (sk === 0n) throw new Error('Organizer secret must be non-zero.')
        const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, epochId, aid)
        tx = await writer.registerApplicationCoDec(
          epochId,
          aid,
          policy,
          pkOrgX,
          pkOrgY,
          proof.ax,
          proof.ay,
          proof.z,
        )
      }
      setLastTx(tx)
      onRegistered?.(tx, aid)
    } catch (e) {
      setError(e instanceof Error ? e.message : String(e))
    } finally {
      setSubmitting(false)
    }
  }

  return (
    <Stack
      gap={6}
      p={{ base: 5, md: 6 }}
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='lg'
      bg='surface'
      boxShadow='inset'
    >
      <Stack gap={2}>
        <Text
          fontFamily='mono'
          fontSize='2xs'
          color='ink.3'
          letterSpacing='0.08em'
          textTransform='uppercase'
        >
          Application registration
        </Text>
        <Text fontSize='sm' color='ink.1' lineHeight='1.55'>
          Bind a per-application key to this epoch. Mode 0 derives publicly from a hash of{' '}
          <code>(eid, PK_ep, aid)</code>. Mode 1 adds an organizer's private key to the epoch
          key; decryption then requires both the committee threshold and the organizer's share.
        </Text>
      </Stack>

      <ModeRadio mode={mode} setMode={setMode} />

      <FieldRow label='Application id (aid)'>
        <Input
          value={aid}
          onChange={(e) => setAid(e.target.value as Hex)}
          fontFamily='mono'
          fontSize='xs'
          spellCheck={false}
        />
        <Text fontSize='2xs' color='ink.4' mt={1}>
          bytes32. Stable identifier — you'll need the same value to submit ciphertexts.
        </Text>
      </FieldRow>

      <FieldRow label='Max ciphertexts'>
        <Input
          value={maxCiphertexts}
          onChange={(e) => setMaxCiphertexts(e.target.value)}
          fontFamily='mono'
          fontSize='xs'
          maxW='12ch'
        />
        <Text fontSize='2xs' color='ink.4' mt={1}>
          0 = unlimited.
        </Text>
      </FieldRow>

      {mode === 1 && (
        <FieldRow label='Organizer secret (sk_org)'>
          <Input
            value={organizerSecret}
            onChange={(e) => setOrganizerSecret(e.target.value)}
            placeholder='decimal or 0x-hex; non-zero, < L'
            fontFamily='mono'
            fontSize='xs'
            type='password'
          />
          <Text fontSize='2xs' color='ink.4' mt={1}>
            Stays in your browser. Only PK_org and the Schnorr proof are sent on-chain.
          </Text>
        </FieldRow>
      )}

      {mode === 0 && sPreview !== null && (
        <Preview label='S = H(eid ‖ PK_ep ‖ aid) mod L' value={sPreview} />
      )}
      {mode === 1 && pkOrgPreview && (
        <Stack gap={1}>
          <Text fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase'>
            PK_org preview (TE coords)
          </Text>
          <HStack gap={4}>
            <HashCell value={toHex(pkOrgPreview.x, { size: 32 })} head={6} tail={6} />
            <HashCell value={toHex(pkOrgPreview.y, { size: 32 })} head={6} tail={6} />
          </HStack>
        </Stack>
      )}

      {error && (
        <Box
          borderLeftWidth='2px'
          borderColor='danger.fg'
          bg='danger.bg'
          px={4}
          py={3}
          borderRadius='md'
        >
          <Text fontSize='sm' color='danger.fg'>
            {error}
          </Text>
        </Box>
      )}

      {lastTx && (
        <Box
          borderLeftWidth='2px'
          borderColor='live.fg'
          bg='live.bg'
          px={4}
          py={3}
          borderRadius='md'
        >
          <Text fontSize='xs' color='ink.1' mb={1}>
            Application registered. Transaction:
          </Text>
          <HashCell value={lastTx} head={8} tail={8} />
        </Box>
      )}

      <HStack gap={3}>
        <Button
          onClick={handleSubmit}
          disabled={submitting || !writer || (mode === 1 && !organizerSecret)}
          variant='solid'
          colorPalette='cyan'
        >
          {submitting ? 'Submitting…' : 'Register application'}
        </Button>
        {!writer && (
          <Text fontSize='xs' color='ink.3'>
            Connect a wallet to enable submission.
          </Text>
        )}
      </HStack>

      <DetailDisclosure title='Show derivation transcript'>
        <Stack gap={1} fontSize='2xs' fontFamily='mono' color='ink.3'>
          <Text>eid = {epochId}</Text>
          <Text>aid = {aid}</Text>
          <Text>mode = {mode === 0 ? 'PUBLIC_DERIVATION' : 'ORGANIZER_CODEC'}</Text>
          {sPreview !== null && <Text>S = {sPreview.toString()}</Text>}
        </Stack>
      </DetailDisclosure>
    </Stack>
  )
}

// ─── small leaf components ──────────────────────────────────────────────────

function ModeRadio({ mode, setMode }: { mode: Mode; setMode: (m: Mode) => void }) {
  return (
    <Field.Root>
      <Field.Label
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
      >
        Decryption mode
      </Field.Label>
      <RadioGroup.Root
        value={mode.toString()}
        onValueChange={(d) => setMode(Number(d.value) as Mode)}
      >
        <Stack gap={2} mt={2}>
          <RadioGroup.Item value='0'>
            <RadioGroup.ItemHiddenInput />
            <RadioGroup.ItemIndicator />
            <RadioGroup.ItemText>
              <Text fontWeight={500} color='ink.0'>Public derivation</Text>
              <Text fontSize='xs' color='ink.3'>Anyone with (eid, aid) can derive PK_aid; only the committee can decrypt.</Text>
            </RadioGroup.ItemText>
          </RadioGroup.Item>
          <RadioGroup.Item value='1'>
            <RadioGroup.ItemHiddenInput />
            <RadioGroup.ItemIndicator />
            <RadioGroup.ItemText>
              <Text fontWeight={500} color='ink.0'>Organizer co-decryption</Text>
              <Text fontSize='xs' color='ink.3'>PK_aid = PK_ep + PK_org; decryption requires both the committee threshold and the organizer's share.</Text>
            </RadioGroup.ItemText>
          </RadioGroup.Item>
        </Stack>
      </RadioGroup.Root>
    </Field.Root>
  )
}

function FieldRow({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <Grid templateColumns={{ base: '1fr', md: '160px 1fr' }} gap={3} alignItems='start'>
      <GridItem>
        <Text
          fontFamily='mono'
          fontSize='2xs'
          color='ink.3'
          letterSpacing='0.08em'
          textTransform='uppercase'
          mt={2}
        >
          {label}
        </Text>
      </GridItem>
      <GridItem>{children}</GridItem>
    </Grid>
  )
}

function Preview({ label, value }: { label: string; value: bigint }) {
  return (
    <Stack gap={1}>
      <Text
        fontFamily='mono'
        fontSize='2xs'
        color='ink.3'
        letterSpacing='0.08em'
        textTransform='uppercase'
      >
        {label}
      </Text>
      <HashCell value={toHex(value, { size: 32 })} head={8} tail={8} />
    </Stack>
  )
}

// Suppress unused-import warning if keccak256 ever gets needed for aid hashing.
void keccak256
