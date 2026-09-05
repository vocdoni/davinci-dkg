// The organizer's client-side cryptography, wrapped so the step panels never
// touch curve arithmetic directly.
//
// Everything here runs in the browser in both modes: the secret is drawn
// locally, `PK_org = sk_org·G` and the Schnorr proof of possession are built
// locally, and the ElGamal ciphertext is built locally. Revealing the secret
// later needs no proof at all — the contract checks `sk_org·G = PK_org`
// itself. An organizer needs no circuit artifacts and no prover; all the
// expensive proving is on the committee's side.
//
// The Schnorr proof is drawn against a nonce this module supplies rather than
// the SDK's internal one, for one reason: the "advanced" toggle prints the
// exact words that go into the transaction, and it can only do that if it
// holds the witness the writer will use.

import { Base8, mulPointEscalar, subOrder } from '@zk-kit/baby-jubjub'
import {
  applicationKey,
  encrypt,
  pointFromTEtoRTE,
  proveOrganizer,
  randomOrganizerSecret,
  type BabyJubPoint,
  type ElGamalCiphertext,
} from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import type { SerialCiphertext } from './machine'

export { randomOrganizerSecret }

/** Largest plaintext the committee's BSGS recovery can invert (`node/dlog.go`). */
export const MAX_PLAINTEXT = 1n << 50n

/** A curve point as the store and the SDK each like it. */
export interface Point {
  x: bigint
  y: bigint
}

export function toPair(point: Point): BabyJubPoint {
  return [point.x, point.y]
}

/**
 * A fresh application id: 32 random bytes with the top three bits cleared, so
 * the value is a BN254 scalar. `aid` is a public input of every decryption
 * proof, and the contract rejects anything at or above the field modulus.
 */
export function randomAid(): Hex {
  const bytes = globalThis.crypto.getRandomValues(new Uint8Array(32))
  bytes[0] &= 0x1f
  // A zero aid is rejected on chain; the odds are nil but the check is free.
  if (bytes.every((b) => b === 0)) bytes[31] = 1
  let hex = '0x'
  for (const b of bytes) hex += b.toString(16).padStart(2, '0')
  return hex as Hex
}

/** True for a `0x`-prefixed 32-byte value below the BN254 scalar field. */
export function isValidAid(value: string): boolean {
  if (!/^0x[0-9a-fA-F]{64}$/.test(value)) return false
  const v = BigInt(value)
  return v > 0n && v >> 253n === 0n
}

/** A witness scalar in `[1, q)`, drawn from the same CSPRNG the SDK uses. */
function randomNonce(): bigint {
  let s = 0n
  while (s === 0n) {
    const bytes = globalThis.crypto.getRandomValues(new Uint8Array(32))
    let acc = 0n
    for (let i = 0; i < bytes.length; i++) acc += BigInt(bytes[i]) << BigInt(8 * i)
    s = acc % subOrder
  }
  return s
}

/** `PK_org = sk_org·G`, in the TE form the SDK and the indexer both use. */
export function organizerPublicKey(skOrg: bigint): BabyJubPoint {
  return mulPointEscalar(Base8, skOrg) as BabyJubPoint
}

/** `PK_aid = P_j` (automatic) or `P_j + PK_org` (organizer-locked). */
export function applicationPublicKey(poolKey: BabyJubPoint, pkOrg: BabyJubPoint | null): BabyJubPoint {
  return applicationKey(poolKey, pkOrg ?? undefined)
}

export interface RegistrationProof {
  /** Witness pinned so the displayed words are the submitted words. */
  nonce: bigint
  pkOrg: BabyJubPoint
  /** A = w·G, in the on-chain (RTE) form the contract hashes. */
  a: BabyJubPoint
  z: bigint
}

/** The Schnorr proof of possession `registerApplication` puts on chain. */
export function registrationProof(skOrg: bigint, epochId: Hex, aid: Hex): RegistrationProof {
  const nonce = randomNonce()
  const { pkOrgX, pkOrgY, proof } = proveOrganizer(skOrg, epochId, aid, nonce)
  return { nonce, pkOrg: [pkOrgX, pkOrgY], a: [proof.ax, proof.ay], z: proof.z }
}

/** Encrypt `value` under `PK_aid`. */
export async function encryptValue(value: bigint, pkAid: BabyJubPoint): Promise<ElGamalCiphertext> {
  return encrypt(value, pkAid)
}

/** Parse the plaintext field: a non-negative integer under the BSGS cap. */
export function parsePlaintext(input: string): { value: bigint } | { error: string } {
  const trimmed = input.trim()
  if (trimmed === '') return { error: 'Enter a value' }
  if (!/^\d+$/.test(trimmed)) return { error: 'Plaintexts are non-negative integers' }
  const value = BigInt(trimmed)
  if (value >= MAX_PLAINTEXT) return { error: 'Must be below 2^50 — the committee cannot invert more' }
  return { value }
}

// ── serialisation ────────────────────────────────────────────────────────────

export function serialiseCiphertext(ct: ElGamalCiphertext): SerialCiphertext {
  return {
    c1: [ct.c1[0].toString(), ct.c1[1].toString()],
    c2: [ct.c2[0].toString(), ct.c2[1].toString()],
  }
}

export function deserialiseCiphertext(ct: SerialCiphertext): ElGamalCiphertext {
  return {
    c1: [BigInt(ct.c1[0]), BigInt(ct.c1[1])],
    c2: [BigInt(ct.c2[0]), BigInt(ct.c2[1])],
  }
}

/** The four calldata words `submitCiphertext` actually sends, in RTE form. */
export function ciphertextWords(ct: ElGamalCiphertext): Array<{ label: string; value: bigint }> {
  const [c1x, c1y] = pointFromTEtoRTE(ct.c1)
  const [c2x, c2y] = pointFromTEtoRTE(ct.c2)
  return [
    { label: 'C1.x', value: c1x },
    { label: 'C1.y', value: c1y },
    { label: 'C2.x', value: c2x },
    { label: 'C2.y', value: c2y },
  ]
}

/** Structural equality of two ciphertexts, for the local verification step. */
export function sameCiphertext(a: ElGamalCiphertext | null, b: ElGamalCiphertext | null): boolean {
  if (!a || !b) return false
  return a.c1[0] === b.c1[0] && a.c1[1] === b.c1[1] && a.c2[0] === b.c2[0] && a.c2[1] === b.c2[1]
}
