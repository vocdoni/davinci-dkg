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
  onReleased: (txHash: Hex) => void
  log: (msg: string, level?: 'info' | 'success' | 'error' | 'tx' | 'crypto') => void
}

export function ReleaseShareStep({
  status,
  epochId,
  aid,
  ciphertextIndex,
  ciphertext,
  skOrg,
  onReleased,
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
      n={8}
      title='Release your organizer share'
      status={status}
      description='The committee has done its half. Now you contribute yours — without it, the ciphertext stays sealed no matter how many committee members cooperate.'
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
          <HStack gap={3}>
            <Button
              colorPalette='cyan'
              size='sm'
              onClick={onRelease}
              loading={busy}
              disabled={!writer || !epochId || !aid}
            >
              Release organizer share →
            </Button>
            {!writer && (
              <Text fontSize='xs' color='ink.3'>
                Connect a wallet to enable submission.
              </Text>
            )}
          </HStack>
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
