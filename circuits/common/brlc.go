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

// DeriveChallengeNative mirrors the Solidity BRLC challenge derivation.
func DeriveChallengeNative(roundHash *big.Int, domain [32]byte, anchor *big.Int) (*big.Int, error) {
	if roundHash == nil {
		return nil, fmt.Errorf("round hash is required")
	}
	if anchor == nil {
		return nil, fmt.Errorf("anchor is required")
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
