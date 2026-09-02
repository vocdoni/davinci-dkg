import { useEffect, useState } from 'react'
import { Box, Button, HStack, Input, Stack, Text } from '@chakra-ui/react'
import { useQueryClient } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { fromRTEtoTE } from '@vocdoni/davinci-dkg-sdk'
import { useDkgClient } from '~hooks/use-dkg-client'
import { useDkgWriter } from '~hooks/use-dkg-writer'
import { loadOrganizerSecret, parseOrganizerSecret } from '~lib/organizer-secret'
import { HashCell } from '~components/ui/HashCell'

interface Props {
  epochId: Hex
  aid: Hex
}

/**
 * Release the organizer share for one ciphertext of an application, outside
 * the playground walkthrough: the organizer who comes back later, or whose
 * walkthrough was interrupted, only needs the application id, the ciphertext
 * index and the organizer secret. The ciphertext itself is read back from the
 * CiphertextSubmitted event, the DLEQ is computed in the browser.
 */
export function OrganizerSharePanel({ epochId, aid }: Props) {
  const { dkg } = useDkgClient()
  const writer = useDkgWriter()
  const queryClient = useQueryClient()
  const [index, setIndex] = useState('1')
  const [secretInput, setSecretInput] = useState('')
  const [busy, setBusy] = useState(false)
  const [status, setStatus] = useState<{ kind: 'ok' | 'error'; text: string; tx?: Hex } | null>(null)

  // Prefill from the secret the playground kept for this (epoch, aid), if any.
  useEffect(() => {
    const stored = loadOrganizerSecret(epochId, aid)
    setSecretInput(stored !== null ? stored.toString() : '')
    setStatus(null)
  }, [epochId, aid])

  const secret = parseOrganizerSecret(secretInput)
  const ctIdx = Number.parseInt(index, 10)
  const ready = Boolean(writer && dkg && secret !== null && Number.isInteger(ctIdx) && ctIdx >= 1)

  const onRelease = async () => {
    if (!writer || !dkg || secret === null) return
    setBusy(true)
    setStatus(null)
    try {
      const events = await dkg.getCiphertextSubmittedEvents(epochId, { aid, ciphertextIndex: ctIdx })
      const ev = events[0]
      if (!ev) throw new Error(`no ciphertext with index ${ctIdx} under this application`)
      const c1 = fromRTEtoTE(ev.c1.x, ev.c1.y)
      const c2 = fromRTEtoTE(ev.c2.x, ev.c2.y)
      const hash = await writer.submitOrganizerShare(epochId, aid, ctIdx, { c1, c2 }, secret)
      await writer.waitForTransaction(hash)
      setStatus({ kind: 'ok', text: `Organizer share released for ciphertext ${ctIdx}.`, tx: hash })
      void queryClient.invalidateQueries({ queryKey: ['ciphertextStatus', epochId, aid] })
    } catch (err) {
      setStatus({ kind: 'error', text: err instanceof Error ? err.message : String(err) })
    } finally {
      setBusy(false)
    }
  }

  return (
    <Box borderWidth='1px' borderColor='ink.8' borderRadius='md' p={4}>
      <Stack gap={3}>
        <Text fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase'>
          Release organizer share
        </Text>
        <Text fontSize='sm' color='ink.4'>
          For the organizer of this application: pick the ciphertext index, paste the organizer secret, and
          the share Δ = sk_org·C1 with its proof is computed here and posted on chain. The committee's
          partials are already there; this is the half that lets anyone combine the plaintext.
        </Text>
        <HStack gap={3} align='flex-end' wrap='wrap'>
          <Stack gap={1} minW='8rem'>
            <Text fontSize='xs' color='ink.4'>
              Ciphertext index
            </Text>
            <Input value={index} onChange={(e) => setIndex(e.target.value)} fontFamily='mono' fontSize='xs' />
          </Stack>
          <Stack gap={1} flex='1' minW='16rem'>
            <Text fontSize='xs' color='ink.4'>
              Organizer secret (sk_org)
            </Text>
            <Input
              value={secretInput}
              onChange={(e) => setSecretInput(e.target.value)}
              placeholder='decimal or 0x-hex; never leaves this browser'
              fontFamily='mono'
              fontSize='xs'
              spellCheck={false}
            />
          </Stack>
          <Button onClick={onRelease} disabled={!ready || busy} loading={busy}>
            Release share
          </Button>
        </HStack>
        {!writer && (
          <Text fontSize='xs' color='ink.4'>
            Connect the organizer's wallet to sign the submitOrganizerShare transaction.
          </Text>
        )}
        {status && (
          <Text fontSize='sm' color={status.kind === 'ok' ? 'green.400' : 'red.400'}>
            {status.text} {status.tx && <HashCell value={status.tx} head={10} tail={6} />}
          </Text>
        )}
      </Stack>
    </Box>
  )
}
