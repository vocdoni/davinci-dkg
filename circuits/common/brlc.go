package common

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// BRLC computes a binding random linear combination commitment inside a circuit.
func BRLC(api frontend.API, challenge frontend.Variable, values []frontend.Variable) frontend.Variable {
	acc := frontend.Variable(0)
	power := challenge
	for i, value := range values {
		if i == 0 {
			acc = api.Mul(power, value)
			continue
		}
		power = api.Mul(power, challenge)
		acc = api.Add(acc, api.Mul(power, value))
	}
	return acc
}

// BRLCNative computes a binding random linear combination over the BN254 scalar field.
func BRLCNative(challenge *big.Int, values ...*big.Int) (*big.Int, error) {
	if challenge == nil {
		return nil, fmt.Errorf("challenge is required")
	}

	modulus := ecc.BN254.ScalarField()
	rho := new(big.Int).Mod(new(big.Int).Set(challenge), modulus)
	acc := big.NewInt(0)
	power := new(big.Int).Set(rho)

	for i, value := range values {
		if value == nil {
			return nil, fmt.Errorf("value %d is nil", i)
		}

		term := new(big.Int).Mul(power, value)
		term.Mod(term, modulus)
		acc.Add(acc, term)
		acc.Mod(acc, modulus)

		power.Mul(power, rho)
		power.Mod(power, modulus)
	}

	return acc, nil
}

// HashPackedBigIntsNative mirrors keccak256(abi.encodePacked(bytes32...)).
func HashPackedBigIntsNative(values ...*big.Int) (*big.Int, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("at least one value is required")
	}
	packed := make([]byte, 0, len(values)*32)
	for i, value := range values {
		if value == nil {
			return nil, fmt.Errorf("value %d is nil", i)
		}
		packed = append(packed, value.FillBytes(make([]byte, 32))...)
	}
	return new(big.Int).SetBytes(ethcrypto.Keccak256(packed)), nil
}

// ChallengeAnchor mirrors the contract's anchor for the BRLC challenge:
// keccak(digest_1 ‖ … ‖ digest_k ‖ keccak(transcript words)). Every value
// the prover chooses — the calldata transcript and the Poseidon digests
// that fix the witness — is hashed in before ρ exists, which is what makes
// the random-linear-combination commitment binding.
func ChallengeAnchor(transcript []*big.Int, digests ...*big.Int) (*big.Int, error) {
	transcriptHash, err := HashPackedBigIntsNative(transcript...)
	if err != nil {
		return nil, fmt.Errorf("hash transcript: %w", err)
	}
	return HashPackedBigIntsNative(append(append([]*big.Int{}, digests...), transcriptHash)...)
}

// DeriveChallengeNative mirrors the Solidity BRLC challenge derivation.
func DeriveChallengeNative(roundHash *big.Int, domain [32]byte, anchor *big.Int) (*big.Int, error) {
	if roundHash == nil {
		return nil, fmt.Errorf("epoch hash is required")
	}
	if anchor == nil {
		return nil, fmt.Errorf("anchor is required")
	}
	if roundHash.BitLen() > 96 {
		return nil, fmt.Errorf("roundHash must fit in 12 bytes (≤96 bits), got %d bits", roundHash.BitLen())
	}
	modulus := ecc.BN254.ScalarField()
	challengeBytes := ethcrypto.Keccak256(
		roundHash.FillBytes(make([]byte, 12)),
		domain[:],
		anchor.FillBytes(make([]byte, 32)),
	)
	challenge := new(big.Int).SetBytes(challengeBytes)
	challenge.Mod(challenge, modulus)
	return challenge, nil
}

// GatedFold is the compact-transcript variant of BRLC (docs/pool-keys-v4.md
// §4). The circuit still traverses its fixed-size witness arrays in region
// order, but every candidate word carries a gate b ∈ {0, 1} derived from the
// public counts; an inactive word neither contributes to the commitment nor
// advances the exponent, so the fold equals BRLC over exactly the words the
// contract streams from calldata, with no padding in between:
//
//	acc   ← acc + b·power·v
//	power ← power·(1 + b·(ρ−1))
//	count ← count + b
//
// starting from (acc, power, count) = (0, ρ, 0). The caller asserts count
// against the expected compact length and acc against the public transcript
// commitment. Gates must be booleans (PrefixMask yields them); the fold does
// not re-assert that, to keep the per-word cost at three multiplications.
type GatedFold struct {
	api       frontend.API
	challenge frontend.Variable
	acc       frontend.Variable
	power     frontend.Variable
	count     frontend.Variable
}

// NewGatedFold starts a gated fold under the given challenge ρ.
func NewGatedFold(api frontend.API, challenge frontend.Variable) *GatedFold {
	return &GatedFold{
		api:       api,
		challenge: challenge,
		acc:       0,
		power:     challenge,
		count:     0,
	}
}

// Absorb folds every value under one shared gate: the natural unit is a
// point (two coordinates) or a scalar, all of whose words are active or
// inactive together. The exponent step 1 + b·(ρ−1) is computed once per call.
func (f *GatedFold) Absorb(gate frontend.Variable, values ...frontend.Variable) {
	step := f.api.Add(1, f.api.Mul(gate, f.api.Sub(f.challenge, 1)))
	for _, value := range values {
		f.acc = f.api.Add(f.acc, f.api.Mul(gate, f.api.Mul(f.power, value)))
		f.power = f.api.Mul(f.power, step)
		f.count = f.api.Add(f.count, gate)
	}
}

// Commitment returns the folded commitment Σ ρ^(q+1)·w[q] over the active
// words w in traversal order.
func (f *GatedFold) Commitment() frontend.Variable { return f.acc }

// Count returns the number of active words absorbed so far.
func (f *GatedFold) Count() frontend.Variable { return f.count }

// GatedBRLCNative is the reference implementation of the gated fold over the
// BN254 scalar field, applying the exact per-word update rules of GatedFold.
// It exists so tests can pin the gadget to the rules and both to plain
// BRLCNative over the active words: skipping an inactive word must leave the
// exponent untouched.
func GatedBRLCNative(challenge *big.Int, values []*big.Int, gates []bool) (*big.Int, int, error) {
	if challenge == nil {
		return nil, 0, fmt.Errorf("challenge is required")
	}
	if len(values) != len(gates) {
		return nil, 0, fmt.Errorf("got %d values and %d gates", len(values), len(gates))
	}
	modulus := ecc.BN254.ScalarField()
	rho := new(big.Int).Mod(new(big.Int).Set(challenge), modulus)
	acc := big.NewInt(0)
	power := new(big.Int).Set(rho)
	count := 0
	for i, value := range values {
		if value == nil {
			return nil, 0, fmt.Errorf("value %d is nil", i)
		}
		if !gates[i] {
			continue
		}
		term := new(big.Int).Mul(power, value)
		acc.Add(acc, term.Mod(term, modulus))
		acc.Mod(acc, modulus)
		power.Mul(power, rho)
		power.Mod(power, modulus)
		count++
	}
	return acc, count, nil
}
