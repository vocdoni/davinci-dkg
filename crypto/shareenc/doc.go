// Package shareenc implements the hashed-ElGamal share encryption scheme
// used by the contribution phase of the DKG.
//
// A share s_i(j) is published on-chain as the pair (R, σ) where R = r·G is
// the ephemeral public key and σ = s_i(j) + H(rid, i, j, r·pub_j) mod q is
// the masked share. The recipient recovers s_i(j) from their long-term
// secret sk_j using σ − H(rid, i, j, sk_j·R). The hash H is the same
// Poseidon1 instance used by the contribution circuit, see the hash
// package for the domain separator.
package shareenc
