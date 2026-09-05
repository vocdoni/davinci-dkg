import { describe, expect, it } from 'vitest'
import { Base8, addPoint, mulPointEscalar } from '@zk-kit/baby-jubjub'
import { applicationPublicKey, formatPointPair } from './keys'

const poolKey = mulPointEscalar(Base8, 12345n) as [bigint, bigint]
const pkOrg = mulPointEscalar(Base8, 67890n) as [bigint, bigint]

describe('applicationPublicKey', () => {
  it('is P_j + PK_org on the curve for an organizer-locked application', () => {
    const key = applicationPublicKey({ x: poolKey[0], y: poolKey[1] }, { x: pkOrg[0], y: pkOrg[1] })
    const expected = addPoint(poolKey, pkOrg)
    expect(key).not.toBeNull()
    expect(key?.x).toBe(expected[0])
    expect(key?.y).toBe(expected[1])
  })

  it('equals (sk_pool + sk_org)·G — the property the combine proof relies on', () => {
    const key = applicationPublicKey({ x: poolKey[0], y: poolKey[1] }, { x: pkOrg[0], y: pkOrg[1] })
    const direct = mulPointEscalar(Base8, 12345n + 67890n)
    expect(key?.x).toBe(direct[0])
    expect(key?.y).toBe(direct[1])
  })

  it('is the pool key itself for an automatic application', () => {
    expect(applicationPublicKey({ x: poolKey[0], y: poolKey[1] }, null)).toEqual({ x: poolKey[0], y: poolKey[1] })
    expect(applicationPublicKey({ x: poolKey[0], y: poolKey[1] })).toEqual({ x: poolKey[0], y: poolKey[1] })
  })

  it('is null until the pool key has been read', () => {
    expect(applicationPublicKey(null, { x: pkOrg[0], y: pkOrg[1] })).toBeNull()
    expect(applicationPublicKey(undefined)).toBeNull()
  })
})

describe('formatPointPair', () => {
  it('renders both coordinates as 32-byte hex, x first', () => {
    expect(formatPointPair({ x: 1n, y: 2n })).toBe(`0x${'0'.repeat(63)}1,0x${'0'.repeat(63)}2`)
  })
})
