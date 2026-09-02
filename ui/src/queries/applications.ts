import { useQuery } from '@tanstack/react-query'
import { useDkgClient } from '~hooks/use-dkg-client'
import { Polling } from '~constants/polling'

/** The all-zero bytes32 the app manager returns when no share is stored. */
export const ZERO_BYTES32 = ('0x' + '00'.repeat(32)) as `0x${string}`

export function useCollectivePublicKey(epochId: `0x${string}` | undefined) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['collectivePublicKey', epochId],
    queryFn: () => {
      if (!epochId) throw new Error('epochId required')
      return dkg.getCollectivePublicKey(epochId)
    },
    enabled: Boolean(epochId),
    refetchInterval: Polling.default,
  })
}

// Per-application reads. The cache key includes (epochId, aid) so that
// multiple AppRegistrationForms / DecryptionPipelines on the same page
// don't collide.

export function useApplication(
  epochId: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['application', epochId, aid],
    queryFn: () => {
      if (!epochId || !aid) throw new Error('epochId + aid required')
      return dkg.getApplication(epochId, aid)
    },
    enabled: Boolean(epochId && aid),
    refetchInterval: Polling.default,
  })
}

/**
 * Organizer-share status for one ciphertext. `getOrganizerShareHash` returns
 * `0x00…00` until the organizer releases `Δ = sk_org·C1`; the committee cannot
 * combine before that (`combineDecryption` reverts `OrganizerShareMissing()`),
 * so this read is what tells "waiting for the committee" apart from "waiting
 * for the organizer".
 */
export function useOrganizerShareHash(
  epochId: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
  ciphertextIndex: number | undefined,
) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['organizerShare', epochId, aid, ciphertextIndex],
    queryFn: () => {
      if (!epochId || !aid || ciphertextIndex == null) {
        throw new Error('epochId + aid + ciphertextIndex required')
      }
      return dkg.getOrganizerShareHash(epochId, aid, ciphertextIndex)
    },
    enabled: Boolean(epochId && aid && ciphertextIndex != null),
    refetchInterval: (q) => (q.state.data && q.state.data !== ZERO_BYTES32 ? false : Polling.decryption),
  })
}

/** Number of ciphertexts accepted under `(epochId, aid)`; indices run 1…count. */
export function useCiphertextCount(
  epochId: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['ciphertextCount', epochId, aid],
    queryFn: () => {
      if (!epochId || !aid) throw new Error('epochId + aid required')
      return dkg.ciphertextCount(epochId, aid)
    },
    enabled: Boolean(epochId && aid),
    refetchInterval: Polling.default,
  })
}

/**
 * Per-ciphertext decryption pipeline snapshot: is the organizer share stored,
 * has the committee combined, and what plaintext came out. One query per
 * ciphertext keeps the cache keys aligned with what the UI renders per row.
 */
export function useCiphertextStatus(
  epochId: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
  ciphertextIndex: number,
) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['ciphertextStatus', epochId, aid, ciphertextIndex],
    queryFn: async () => {
      if (!epochId || !aid) throw new Error('epochId + aid required')
      const [shareHash, combined] = await Promise.all([
        dkg.getOrganizerShareHash(epochId, aid, ciphertextIndex),
        dkg.getCombinedDecryption(epochId, aid, ciphertextIndex),
      ])
      return {
        organizerShare: shareHash !== ZERO_BYTES32,
        organizerShareHash: shareHash,
        combined: combined.completed,
        plaintext: combined.plaintext,
      }
    },
    // Indices are assigned from 1, so a 0 here means "no ciphertext yet" and
    // the read would be a wasted round trip.
    enabled: Boolean(epochId && aid) && ciphertextIndex > 0,
    refetchInterval: (q) => (q.state.data?.combined ? false : Polling.decryption),
  })
}

/**
 * The `CiphertextSubmitted` event for one ciphertext index — the only source of
 * the actual (C1, C2) coordinates, since the contract stores just their hash.
 *
 * The playground uses it to check locally that the ciphertext the chain
 * decrypted is byte-for-byte the one this browser built. Coordinates come back
 * in the on-chain (RTE) form.
 */
export function useCiphertextRecord(
  epochId: `0x${string}` | undefined,
  aid: `0x${string}` | undefined,
  ciphertextIndex: number | null | undefined,
) {
  const { dkg } = useDkgClient()
  return useQuery({
    queryKey: ['ciphertextRecord', epochId, aid, ciphertextIndex],
    queryFn: async () => {
      if (!epochId || !aid || ciphertextIndex == null) throw new Error('epochId + aid + index required')
      const events = await dkg.getCiphertextSubmittedEvents(epochId, { aid, ciphertextIndex })
      return events.at(-1) ?? null
    },
    enabled: Boolean(epochId && aid && ciphertextIndex != null),
    // The coordinates never change once mined; one successful read is enough.
    staleTime: Infinity,
  })
}
