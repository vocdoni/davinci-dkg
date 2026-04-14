package web3

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// DeriveBRLCChallenge derives a BN254-scalar challenge from domain-separated inputs.
func DeriveBRLCChallenge(roundID [12]byte, domain string, anchor common.Hash) *big.Int {
	modulus := ecc.BN254.ScalarField()
	payload := append(append(roundID[:], []byte(domain)...), anchor.Bytes()...)
	sum := crypto.Keccak256(payload)
	return new(big.Int).Mod(new(big.Int).SetBytes(sum), modulus)
}

// BRLCCommit compresses a vector of scalars using the provided challenge.
func BRLCCommit(challenge *big.Int, values ...*big.Int) (*big.Int, error) {
	modulus := ecc.BN254.ScalarField()
	rho := new(big.Int).Mod(new(big.Int).Set(challenge), modulus)
	acc := big.NewInt(0)
	power := new(big.Int).Set(rho)

	for _, value := range values {
		term := new(big.Int).Mul(power, value)
		term.Mod(term, modulus)
		acc.Add(acc, term)
		acc.Mod(acc, modulus)
		power.Mul(power, rho)
		power.Mod(power, modulus)
	}

	return acc, nil
}
