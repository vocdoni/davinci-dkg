import { useState } from 'react'
import { Box, Button, HStack, Input, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import { randomOrganizerSecret, type BabyJubPoint } from '@vocdoni/davinci-dkg-sdk'
import { LuKeyRound, LuSignature, LuShieldCheck } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { useDkgClient } from '~hooks/use-dkg-client'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { parseOrganizerSecret, saveOrganizerSecret } from '~lib/organizer-secret'

// Ciphertexts always live under a registered application, so the playground
// has to register one before it can encrypt anything. Registration binds an
// organizer key: PK_aid = PK_ep + PK_org. The secret half stays in this
// browser — we show it once, offer a copy button, and keep it in session
// storage so the "release the organizer share" step can find it again.

interface Props {
  status: StepStatus
  epochId: Hex | null
  collectivePubKey: { x: bigint; y: bigint } | null
  onRegistered: (aid: Hex, skOrg: bigint, pkOrg: BabyJubPoint) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx' | 'crypto') => void
}

const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000' as const

function randomAid(): Hex {
  const rand = new Uint8Array(32)
  globalThis.crypto.getRandomValues(rand)
  // `aid` is a BN254 scalar-field public input of every decryption proof, so
  // it must stay below the field modulus: clear the top three bits.
  rand[0] &= 0x1f
  return ('0x' + Array.from(rand).map((b) => b.toString(16).padStart(2, '0')).join('')) as Hex
}

export function RegisterAppStep({ status, epochId, collectivePubKey, onRegistered, log }: Props) {
  const writer = useDkgWriter()
  const { dkg } = useDkgClient()
  const [aid, setAid] = useState<Hex>(randomAid())
  const [secret, setSecret] = useState<string>(() => randomOrganizerSecret().toString())
  // A small cap is the honest default for a walkthrough: it is enough to try a
  // second ciphertext under the same application and it makes the policy field
  // visible instead of leaving it at the silent "unlimited" zero.
  const [maxCiphertexts, setMaxCiphertexts] = useState('4')
  const [busy, setBusy] = useState(false)
  const [tx, setTx] = useState<Hex | null>(null)
  const [copied, setCopied] = useState(false)

  const sk = parseOrganizerSecret(secret)

  const onRegister = async () => {
    if (!writer || !epochId || sk === null) return
    setBusy(true)
    try {
      log('Sending registerApplication (PK_org + Schnorr proof of possession)…', 'tx')
      const hash = await writer.registerApplication(
        epochId,
        aid,
        {
          authorizedSubmitter: ZERO_ADDRESS, // resolves on chain to your address
          maxCiphertexts: parseMaxCiphertexts(maxCiphertexts),
          notBeforeBlock: 0n,
          notAfterBlock: 0n,
        },
        sk,
      )
      await writer.waitForTransaction(hash)
      saveOrganizerSecret(epochId, aid, sk)
      setTx(hash)
      log(`registerApplication tx: ${hash}`, 'tx')

      // Read PK_org back from the chain rather than re-deriving it, so the key
      // we encrypt under is provably the one the contract stored.
      const app = await dkg.getApplication(epochId, aid)
      onRegistered(aid, sk, app.organizerPK)
      log('Application registered; PK_aid = PK_ep + PK_org is now the encryption key.', 'crypto')
    } catch (err) {
      log(`registerApplication failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setBusy(false)
    }
  }

  const copy = async () => {
    try {
      await globalThis.navigator?.clipboard?.writeText(secret)
      setCopied(true)
    } catch {
      log('Could not reach the clipboard — select the secret and copy it manually.', 'error')
    }
  }

  return (
    <StepCard
      n={3}
      title='Register your application'
      status={status}
      description='From here on you are the organizer. Ciphertexts belong to an application, not to the epoch, and registering one mixes your own secret into the encryption key — so nobody can have the committee decrypt your data behind your back.'
    >
      {!collectivePubKey ? (
        <Text fontSize='sm' color='ink.4'>
          Waiting for the shared encryption key.
        </Text>
      ) : (
        <Stack gap={4}>
          <Stack gap={2}>
            <Text fontSize='xs' color='ink.4'>
              Application id (aid)
            </Text>
            <Input
              size='sm'
              fontFamily='mono'
              fontSize='xs'
              value={aid}
              spellCheck={false}
              disabled={Boolean(tx)}
              onChange={(e) => setAid(e.target.value as Hex)}
            />
            <Text fontSize='2xs' color='ink.4' lineHeight='1.5'>
              A random 32-byte id with its top three bits cleared: every decryption proof takes
              the aid as a public input on the BN254 scalar field, so the contract rejects
              anything at or above the field modulus. Pick your own if you already have one.
            </Text>
          </Stack>

          <Stack gap={2}>
            <Text fontSize='xs' color='ink.4'>
              Max ciphertexts
            </Text>
            <Input
              size='sm'
              maxW='140px'
              fontFamily='mono'
              fontSize='xs'
              inputMode='numeric'
              value={maxCiphertexts}
              spellCheck={false}
              disabled={Boolean(tx)}
              onChange={(e) => setMaxCiphertexts(e.target.value)}
            />
            <Text fontSize='2xs' color='ink.4' lineHeight='1.5'>
              How many ciphertexts this application may ever accept; <code>0</code> means
              unlimited. The contract counts them and assigns each one its index.
            </Text>
          </Stack>

          <Stack gap={2}>
            <Text fontSize='xs' color='ink.4'>
              Organizer secret (sk_org)
            </Text>
            <HStack gap={2} wrap='wrap'>
              <Input
                size='sm'
                fontFamily='mono'
                fontSize='xs'
                value={secret}
                spellCheck={false}
                disabled={Boolean(tx)}
                onChange={(e) => {
                  setSecret(e.target.value)
                  setCopied(false)
                }}
              />
              {!tx && (
                <Button
                  size='sm'
                  variant='outline'
                  onClick={() => {
                    setSecret(randomOrganizerSecret().toString())
                    setCopied(false)
                  }}
                >
                  Re-roll
                </Button>
              )}
              <Button size='sm' variant='ghost' onClick={copy}>
                {copied ? 'Copied' : 'Copy'}
              </Button>
            </HStack>
            {sk === null && (
              <Text fontSize='2xs' color='danger.fg'>
                Not a non-zero integer (decimal or 0x-hex).
              </Text>
            )}
          </Stack>

          <Box
            borderWidth='1px'
            borderColor='border.subtle'
            bg='surface.sunken'
            px={4}
            py={3}
            borderRadius='md'
          >
            <Text fontSize='xs' color='ink.2' lineHeight='1.6'>
              <Box as='span' color='ink.0' fontWeight={500}>
                Authorised submitter:
              </Box>{' '}
              the wallet you are connected with. The policy field goes on chain as the zero
              address, which the contract resolves to the registering address — there is no open
              submission, so only you can add ciphertexts to this application. Register from a
              different wallet than the one that will submit and the submissions revert; a
              production organizer sets this to whatever service actually posts the ciphertexts.
            </Text>
          </Box>

          <Box
            borderLeftWidth='2px'
            borderColor='danger.fg'
            bg='danger.bg'
            px={4}
            py={3}
            borderRadius='md'
          >
            <Text fontSize='sm' color='danger.fg' fontWeight={500} mb={1}>
              Copy this secret before you register.
            </Text>
            <Text fontSize='xs' color='ink.1' lineHeight='1.55'>
              It never leaves this browser — only <code>PK_org = sk_org·G</code> and a proof
              that you hold it go on chain. <strong>Lose it and every ciphertext under this
              application is undecryptable forever</strong>, threshold or not. This page keeps a
              copy in session storage for the next steps; it disappears when you close the tab.
            </Text>
          </Box>

          {tx ? (
            <Stack gap={2} fontSize='sm'>
              <Text color='live.fg'>Application registered.</Text>
              <HashCell value={tx} head={8} tail={6} />
            </Stack>
          ) : (
            <HStack gap={3}>
              <Button
                colorPalette='cyan'
                size='sm'
                onClick={onRegister}
                loading={busy}
                disabled={!writer || !epochId || sk === null}
              >
                Register application →
              </Button>
              {!writer && (
                <Text fontSize='xs' color='ink.3'>
                  Connect a wallet to sign the <code>registerApplication</code> transaction.
                </Text>
              )}
            </HStack>
          )}

          <DetailDisclosure title='Show what goes on chain'>
            <Stack gap={1} fontSize='2xs' fontFamily='mono' color='ink.3'>
              <Text>registerApplication(eid, aid, policy, PK_org.x, PK_org.y, A.x, A.y, z)</Text>
              <Text>domain = davinci-dkg:organizer-register:v1</Text>
              <Text>c = keccak256(domain ‖ eid ‖ aid ‖ PK_org ‖ A) mod q</Text>
              <Text>z = w + c·sk_org mod q &nbsp; (sk_org itself is never sent)</Text>
            </Stack>
          </DetailDisclosure>

          <HowItWorks
            body={
              <>
                The committee's key alone would let anyone who can talk to the committee open
                anything encrypted under it. So each application adds a second, private key of
                its own: the encryption key becomes the sum of the two. Decryption then needs
                the committee <em>and</em> the application's organizer — which is you, here.
                The chain only ever sees the public half plus a short proof that you really
                hold the secret one.
              </>
            }
            flow={[
              { icon: <LuKeyRound />, label: 'Generate sk_org' },
              { icon: <LuSignature />, label: 'Prove you hold it' },
              { icon: <LuShieldCheck />, label: 'PK_aid = PK_ep + PK_org' },
            ]}
          />
        </Stack>
      )}
    </StepCard>
  )
}

/**
 * Clamp the max-ciphertexts policy field to the contract's uint16, treating a
 * blank or unparseable value as 0 (unlimited) rather than failing the tx.
 */
function parseMaxCiphertexts(input: string): number {
  const n = Number(input.trim())
  if (!Number.isFinite(n) || n < 0) return 0
  return Math.min(65535, Math.floor(n))
}
