import { useMemo, useState } from 'react'
import {
  Box,
  Button,
  Grid,
  GridItem,
  HStack,
  Input,
  Stack,
  Text,
} from '@chakra-ui/react'
import { toHex, type Hex } from 'viem'
import {
  proveOrganizer,
  randomOrganizerSecret,
  type AppPolicy,
} from '@vocdoni/davinci-dkg-sdk'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import {
  clearOrganizerSecret,
  parseOrganizerSecret,
  saveOrganizerSecret,
} from '~lib/organizer-secret'

// AppRegistrationForm — bind a per-application key to a Live epoch.
//
//   PK_aid = PK_ep + PK_org,   PK_org = sk_org · G
//
// Every application works this way: opening a ciphertext needs both the
// committee threshold and the organizer's share Δ = sk_org·C1. There is no
// public-derivation variant and no mode selector — an application nobody holds
// a secret for could be read by anyone who can reach the committee.
//
// The form generates `sk_org` in the browser (or accepts a pasted one), shows
// it exactly once with a copy button, and keeps it in session storage so the
// playground can release the share later. It is never transmitted: only
// `PK_org` and a Schnorr proof of possession go on chain.

interface Props {
  /** Hex-encoded bytes12 epoch id (the URL param). */
  epochId: Hex
  /** Epoch collective public key in TE coords; `null` until finalized. */
  pkEp: { x: bigint; y: bigint } | null
  /** Optional initial `aid`. Defaults to a fresh random bytes32. */
  initialAid?: Hex
  /** Called with the on-chain tx hash on a successful registration. */
  onRegistered?: (txHash: Hex, aid: Hex, skOrg: bigint) => void
}

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as const

function defaultAid(): Hex {
  // Random aid so users can re-submit without colliding while testing.
  const rand = new Uint8Array(32)
  globalThis.crypto.getRandomValues(rand)
  // aid is a BN254 scalar-field public input of every decryption proof, so
  // it must stay below the field modulus (~2^253.6): clear the top 3 bits.
  rand[0] &= 0x1f
  return ('0x' + Array.from(rand).map((b) => b.toString(16).padStart(2, '0')).join('')) as Hex
}

export function AppRegistrationForm({ epochId, pkEp, initialAid, onRegistered }: Props) {
  const writer = useDkgWriter()

  const [aid, setAid] = useState<Hex>(initialAid ?? defaultAid())
  const [organizerSecret, setOrganizerSecret] = useState<string>('')
  const [generated, setGenerated] = useState(false)
  const [maxCiphertexts, setMaxCiphertexts] = useState<string>('0')
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [lastTx, setLastTx] = useState<Hex | null>(null)

  const sk = useMemo(() => parseOrganizerSecret(organizerSecret), [organizerSecret])

  // PK_org / PK_aid preview, derived from the secret as a visible sanity check
  // before submission. A fixed nonce is fine here — the preview only reads the
  // public key out of the prover; the real submission redraws a fresh witness.
  const preview = useMemo(() => {
    if (sk === null) return null
    try {
      const { pkOrgX, pkOrgY } = proveOrganizer(sk, epochId, aid, 1n)
      return { pkOrgX, pkOrgY }
    } catch {
      return null
    }
  }, [sk, epochId, aid])

  function generate() {
    const fresh = randomOrganizerSecret()
    setOrganizerSecret(fresh.toString())
    setGenerated(true)
    setError(null)
    clearOrganizerSecret(epochId, aid)
  }

  async function copySecret() {
    try {
      await globalThis.navigator?.clipboard?.writeText(organizerSecret)
    } catch {
      setError('Could not reach the clipboard — select the value and copy it manually.')
    }
  }

  async function handleSubmit() {
    if (!writer) {
      setError('Connect a wallet first.')
      return
    }
    if (sk === null) {
      setError('Enter or generate a non-zero organizer secret.')
      return
    }
    setSubmitting(true)
    setError(null)
    setLastTx(null)
    try {
      const policy: AppPolicy = {
        authorizedSubmitter: ZERO_ADDRESS, // resolves on chain to the caller
        maxCiphertexts: Math.max(0, Math.min(65535, Number(maxCiphertexts) || 0)),
        notBeforeBlock: 0n,
        notAfterBlock: 0n,
      }
      const tx = await writer.registerApplication(epochId, aid, policy, sk)
      // Only persist once the transaction is actually accepted, so a failed
      // registration doesn't leave a secret for an application that isn't there.
      saveOrganizerSecret(epochId, aid, sk)
      setLastTx(tx)
      onRegistered?.(tx, aid, sk)
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
          Bind a per-application key to this epoch. The key is{' '}
          <code>PK_aid = PK_ep + PK_org</code>, so decryption needs both the committee
          threshold and your organizer share. Only <code>PK_org</code> and a proof that you
          hold its secret are sent on chain.
        </Text>
      </Stack>

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

      <FieldRow label='Organizer secret (sk_org)'>
        <Stack gap={2}>
          <HStack gap={2} align='start' wrap='wrap'>
            <Input
              value={organizerSecret}
              onChange={(e) => {
                setOrganizerSecret(e.target.value)
                setGenerated(false)
              }}
              placeholder='decimal or 0x-hex; non-zero, below the subgroup order'
              fontFamily='mono'
              fontSize='xs'
              spellCheck={false}
            />
            <Button size='sm' variant='outline' onClick={generate}>
              Generate
            </Button>
            {organizerSecret && (
              <Button size='sm' variant='ghost' onClick={copySecret}>
                Copy
              </Button>
            )}
          </HStack>
          {organizerSecret && sk === null && (
            <Text fontSize='2xs' color='danger.fg'>
              Not a non-zero integer (decimal or 0x-hex).
            </Text>
          )}
        </Stack>
      </FieldRow>

      {generated && sk !== null && (
        <Box
          borderLeftWidth='2px'
          borderColor='danger.fg'
          bg='danger.bg'
          px={4}
          py={3}
          borderRadius='md'
        >
          <Text fontSize='sm' color='danger.fg' fontWeight={500} mb={1}>
            Copy this secret now — it is shown once.
          </Text>
          <Text fontSize='xs' color='ink.1' lineHeight='1.55'>
            <code>sk_org</code> is this application's only decryption capability. It is not sent
            anywhere and cannot be recovered from the chain. <strong>If you lose it, every
            ciphertext ever submitted under this aid is permanently undecryptable</strong> — the
            committee threshold alone cannot open them. This page keeps it in session storage so
            the playground can release the share; that copy disappears when you close the tab.
          </Text>
        </Box>
      )}

      {preview && (
        <Stack gap={1}>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            PK_org preview (on-chain coords)
          </Text>
          <HStack gap={4} wrap='wrap'>
            <HashCell value={toHex(preview.pkOrgX, { size: 32 })} head={6} tail={6} />
            <HashCell value={toHex(preview.pkOrgY, { size: 32 })} head={6} tail={6} />
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
          disabled={submitting || !writer || sk === null}
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

      <DetailDisclosure title='Show registration transcript'>
        <Stack gap={1} fontSize='2xs' fontFamily='mono' color='ink.3'>
          <Text>domain = davinci-dkg:organizer-register:v1</Text>
          <Text>eid = {epochId}</Text>
          <Text>aid = {aid}</Text>
          <Text>c = keccak256(domain ‖ eid ‖ aid ‖ PK_org ‖ A) mod q</Text>
          <Text>z = w + c·sk_org mod q</Text>
          {pkEp && <Text>PK_ep = ({pkEp.x.toString()}, {pkEp.y.toString()})</Text>}
        </Stack>
      </DetailDisclosure>
    </Stack>
  )
}

// ─── small leaf components ──────────────────────────────────────────────────

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
