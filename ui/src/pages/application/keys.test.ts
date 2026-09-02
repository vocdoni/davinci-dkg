import { describe, expect, it } from 'vitest'
import { Base8, addPoint, mulPointEscalar } from '@zk-kit/baby-jubjub'
import { applicationPublicKey, formatPointPair } from './keys'

const pkEp = mulPointEscalar(Base8, 12345n) as [bigint, bigint]
const pkOrg = mulPointEscalar(Base8, 67890n) as [bigint, bigint]

describe('applicationPublicKey', () => {
  it('is PK_ep + PK_org on the curve', () => {
    const key = applicationPublicKey({ x: pkEp[0], y: pkEp[1] }, { x: pkOrg[0], y: pkOrg[1] })
    const expected = addPoint(pkEp, pkOrg)
    expect(key).not.toBeNull()
    expect(key?.x).toBe(expected[0])
    expect(key?.y).toBe(expected[1])
  })

  it('equals (sk_ep + sk_org)·G — the property decryption relies on', () => {
    const key = applicationPublicKey({ x: pkEp[0], y: pkEp[1] }, { x: pkOrg[0], y: pkOrg[1] })
    const direct = mulPointEscalar(Base8, 12345n + 67890n)
    expect(key?.x).toBe(direct[0])
    expect(key?.y).toBe(direct[1])
  })

  it('is null until the epoch key exists', () => {
    expect(applicationPublicKey(null, { x: pkOrg[0], y: pkOrg[1] })).toBeNull()
    expect(applicationPublicKey({ x: pkEp[0], y: pkEp[1] }, undefined)).toBeNull()
  })
})

describe('formatPointPair', () => {
  it('renders both coordinates as 32-byte hex, x first', () => {
    expect(formatPointPair({ x: 1n, y: 2n })).toBe(`0x${'0'.repeat(63)}1,0x${'0'.repeat(63)}2`)
  })
})
