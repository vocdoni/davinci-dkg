package hash

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

// HashFieldElements hashes 2–16 field elements with the Poseidon1 (iden3)
// hash function over BN254. This is the same primitive used by davinci-node
// circuits and by gnark-crypto-primitives/hash/native/bn254/poseidon in-circuit.
func HashFieldElements(values ...*big.Int) (*big.Int, error) {
	n := len(values)
	if n < 2 || n > 16 {
		return nil, fmt.Errorf("poseidon1 expects 2–16 inputs, got %d", n)
	}
	for i, v := range values {
		if v == nil {
			return nil, fmt.Errorf("input %d is nil", i)
		}
	}
	return poseidon.Hash(values)
}

// DomainValue reduces an arbitrary byte slice into the BN254 scalar field.
func DomainValue(value []byte) *big.Int {
	return new(big.Int).Mod(new(big.Int).SetBytes(value), ecc.BN254.ScalarField())
}
