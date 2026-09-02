// The application key.
//
//   PK_aid = PK_ep + PK_org
//
// Both halves are stored in the circomlib TE form the SDK works in (the
// indexer converts the on-chain RTE words on the way in), so the sum is the
// SDK's own `applicationKey` with no conversion in between. Decryption needs
// both halves: the committee's threshold alone does not open a ciphertext.

import { applicationKey } from '@vocdoni/davinci-dkg-sdk'
import type { Point } from '~indexer/types'
import { bigIntToHex } from '~lib/format'

export function applicationPublicKey(pkEp: Point | null | undefined, pkOrg: Point | null | undefined): Point | null {
  if (!pkEp || !pkOrg) return null
  const [x, y] = applicationKey([pkEp.x, pkEp.y], [pkOrg.x, pkOrg.y])
  return { x, y }
}

/** `0x<x>,0x<y>` — the pair, in the order every tool in the repo takes it. */
export function formatPointPair(point: Point): string {
  return `${bigIntToHex(point.x)},${bigIntToHex(point.y)}`
}
