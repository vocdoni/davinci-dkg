// The application key.
//
//   PK_aid = P_j            (automatic)
//   PK_aid = P_j + PK_org   (organizer-locked)
//
// Both halves are stored in the circomlib TE form the SDK works in (the
// indexer converts the on-chain RTE words on the way in), so the sum is the
// SDK's own `applicationKey` with no conversion in between.

import { Base8, mulPointEscalar } from '@zk-kit/baby-jubjub'
import { applicationKey } from '@vocdoni/davinci-dkg-sdk'
import type { Point } from '~indexer/types'
import { bigIntToHex } from '~lib/format'

/**
 * `PK_aid` from the claimed pool key and, for an organizer-locked
 * application, `PK_org`. Pass no organizer key (null or undefined) for an
 * automatic application. Null while the pool key is not on chain yet.
 */
export function applicationPublicKey(poolKey: Point | null | undefined, pkOrg?: Point | null): Point | null {
  if (!poolKey) return null
  const [x, y] = applicationKey([poolKey.x, poolKey.y], pkOrg ? [pkOrg.x, pkOrg.y] : undefined)
  return { x, y }
}

/** `0x<x>,0x<y>` — the pair, in the order every tool in the repo takes it. */
export function formatPointPair(point: Point): string {
  return `${bigIntToHex(point.x)},${bigIntToHex(point.y)}`
}

/** True when `skOrg·G` is the registered organizer key — the check the contract repeats on reveal. */
export function matchesOrganizerKey(skOrg: bigint, organizerPK: Point): boolean {
  const [x, y] = mulPointEscalar(Base8, skOrg)
  return x === organizerPK.x && y === organizerPK.y
}
