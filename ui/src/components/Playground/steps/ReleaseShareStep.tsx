import { useState } from 'react'
import { Box, Button, HStack, Stack, Text } from '@chakra-ui/react'
import type { Hex } from 'viem'
import type { ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { LuKeyRound, LuSigma, LuLockOpen } from 'react-icons/lu'
import { StepCard, type StepStatus } from '../StepCard'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { HowItWorks } from '../HowItWorks'
import { HashCell } from '~components/ui/HashCell'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'
import { loadOrganizerSecret } from '~lib/organizer-secret'

// The organizer's half of the decryption. Δ = sk_org·C1, published with a
// Chaum-Pedersen DLEQ proving the same sk_org relates (G, PK_org) and (C1, Δ).
//
// The contract stores only keccak256(Δ ‖ A1 ‖ A2 ‖ z) and does not verify the
// DLEQ; the committee's combine SNARK does, taking the challenge from the
// transcript the contract recomputes. Until this lands, `combineDecryption`
// reverts `OrganizerShareMissing()` no matter how many partials are in.

interface Props {
  status: StepStatus
  epochId: Hex | null
  aid: Hex | null
  ciphertextIndex: number | null
  ciphertext: ElGamalCiphertext | null
  /** The secret from the registration step; falls back to session storage. */
  skOrg: bigint | null
  /** True once the user chose to sit on the share for now. */
  withheld: boolean
  onReleased: (txHash: Hex) => void
  /** Move on without releasing — the decryption step then shows the deadlock. */
  onWithhold: () => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx' | 'crypto') => void
}

export function ReleaseShareStep({
  status,
  epochId,
  aid,
  ciphertextIndex,
  ciphertext,
  skOrg,
  withheld,
  onReleased,
  onWithhold,
  log,
}: Props) {
  const writer = useDkgWriter()
  const [busy, setBusy] = useState(false)
  const [tx, setTx] = useState<Hex | null>(null)

  // Prefer the in-memory value; fall back to session storage so a page reload
  // between publishing and releasing doesn't strand the ciphertext.
  const secret = skOrg ?? (epochId && aid ? loadOrganizerSecret(epochId, aid) : null)

  const onRelease = async () => {
    if (!writer || !epochId || !aid || !ciphertext || ciphertextIndex == null || secret === null) return
    setBusy(true)
    try {
      log('Computing Δ = sk_org·C1 and its DLEQ, then sending submitOrganizerShare…', 'crypto')
      const hash = await writer.submitOrganizerShare(
        epochId,
        aid,
        ciphertextIndex,
        ciphertext,
        secret,
      )
      await writer.waitForTransaction(hash)
      setTx(hash)
      log(`submitOrganizerShare tx: ${hash}`, 'tx')
      onReleased(hash)
    } catch (err) {
      log(`submitOrganizerShare failed: ${err instanceof Error ? err.message : String(err)}`, 'error')
    } finally {
      setBusy(false)
    }
  }

  return (
    <StepCard
      n={6}
      title='Release your organizer share — or hold it back'
      status={status}
      description='The committee does its half on its own schedule. Yours is the half that decides when the ciphertext may be opened at all: until it lands, no amount of committee cooperation reveals anything.'
    >
      <Stack gap={4}>
        {ciphertextIndex == null || !ciphertext ? (
          <Text fontSize='sm' color='ink.4'>
            Publish a ciphertext in the previous step first.
          </Text>
        ) : secret === null ? (
          <Box
            borderLeftWidth='2px'
            borderColor='danger.fg'
            bg='danger.bg'
            px={4}
            py={3}
            borderRadius='md'
          >
            <Text fontSize='sm' color='danger.fg' fontWeight={500} mb={1}>
              Organizer secret not available.
            </Text>
            <Text fontSize='xs' color='ink.1' lineHeight='1.55'>
              This tab no longer holds <code>sk_org</code> for this application — session storage
              is cleared when the tab closes. Without it this ciphertext can never be decrypted;
              register a fresh application and start again.
            </Text>
          </Box>
        ) : tx ? (
          <Stack gap={2} fontSize='sm'>
            <Text color='live.fg'>
              Share released. The committee can now combine and publish the plaintext.
            </Text>
            <HashCell value={tx} head={8} tail={6} />
          </Stack>
        ) : (
          <Stack gap={3}>
            <HStack gap={3} wrap='wrap'>
              <Button
                colorPalette='cyan'
                size='sm'
                onClick={onRelease}
                loading={busy}
                disabled={!writer || !epochId || !aid}
              >
                Release organizer share →
              </Button>
              <Button size='sm' variant='outline' onClick={onWithhold} disabled={busy}>
                Withhold for now
              </Button>
              {!writer && (
                <Text fontSize='xs' color='ink.3'>
                  Connect a wallet to sign the <code>submitOrganizerShare</code> transaction.
                </Text>
              )}
            </HStack>
            <Box
              borderWidth='1px'
              borderColor={withheld ? 'accent.border' : 'border.subtle'}
              bg={withheld ? 'accent.bg' : 'surface.sunken'}
              px={4}
              py={3}
              borderRadius='md'
            >
              <Text fontSize='xs' color='ink.2' lineHeight='1.6'>
                {withheld ? (
                  <>
                    <Box as='span' color='accent.fg' fontWeight={500}>
                      Holding the share back.
                    </Box>{' '}
                    Watch the next step: committee members keep publishing partial decryptions —
                    they react to the ciphertext, not to you — but nothing can be combined.
                    <code> combineDecryption</code> reverts with{' '}
                    <code>OrganizerShareMissing()</code> until you come back here and release it.
                    This is exactly how a poll stays sealed until it closes.
                  </>
                ) : (
                  <>
                    You can wait. The committee publishes its partials as soon as the ciphertext
                    lands, but the plaintext cannot be assembled from them alone, so releasing
                    later — after a poll closes, at a deadline, never — is a normal thing for an
                    organizer to do. Choose <em>Withhold for now</em> to see that state in the
                    next step; you can release from here whenever you like.
                  </>
                )}
              </Text>
            </Box>
          </Stack>
        )}

        <DetailDisclosure title='Show the share transcript'>
          <Stack gap={1} fontSize='2xs' fontFamily='mono' color='ink.3'>
            <Text>Δ = sk_org·C1 &nbsp; A1 = w·G &nbsp; A2 = w·C1</Text>
            <Text>domain = davinci-dkg:organizer-share:v1</Text>
            <Text>
              e = keccak256(domain ‖ eid ‖ aid ‖ uint256(ctIdx) ‖ PK_org ‖ C1 ‖ Δ ‖ A1 ‖ A2) mod q
            </Text>
            <Text>z = w + e·sk_org mod q</Text>
            <Text>on chain: keccak256(Δ ‖ A1 ‖ A2 ‖ z); the DLEQ is checked in the combine proof</Text>
          </Stack>
        </DetailDisclosure>

        <HowItWorks
          body={
            <>
              You send one point — your secret applied to the ciphertext's random part — plus a
              short proof that it really came from the same secret you registered. You are not
              revealing <em>sk_org</em>, and the point on its own reveals nothing about the
              message. The committee's proof checks your proof as part of combining, so a
              malformed share is simply rejected rather than corrupting the result; you can send
              a corrected one until the ciphertext is opened.
            </>
          }
          flow={[
            { icon: <LuKeyRound />, label: 'Your secret' },
            { icon: <LuSigma />, label: 'Δ = sk_org·C1 + proof' },
            { icon: <LuLockOpen />, label: 'Committee can combine' },
          ]}
        />
      </Stack>
    </StepCard>
  )
}
