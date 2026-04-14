package common

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/bits"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	nativeposeidon "github.com/vocdoni/davinci-node/crypto/hash/poseidon"
	circuitposeidon "github.com/vocdoni/gnark-crypto-primitives/hash/native/bn254/poseidon"
)

var (
	recipientIndexShift   = big.NewInt(1 << 16)
	subgroupOrder         = group.ScalarField()
	subgroupOrderMinusOne = new(big.Int).Sub(new(big.Int).Set(subgroupOrder), big.NewInt(1))
)

// MultiHash wraps the gnark Poseidon multihash used by davinci-node circuits.
func MultiHash(api frontend.API, inputs ...frontend.Variable) (frontend.Variable, error) {
	return circuitposeidon.MultiHash(api, inputs...)
}

// MultiHashNative wraps the native Poseidon multihash used in witness builders.
func MultiHashNative(inputs ...*big.Int) (*big.Int, error) {
	return nativeposeidon.MultiPoseidon(inputs...)
}

// HashFieldElementsNative mirrors crypto/hash.HashFieldElements for witness builders.
func HashFieldElementsNative(inputs ...*big.Int) (*big.Int, error) {
	return dkghash.HashFieldElements(inputs...)
}

// HashFieldElements mirrors crypto/hash.HashFieldElements in-circuit.
// Uses the same Poseidon1 (gnark-crypto-primitives MultiHash) as davinci-node circuits.
func HashFieldElements(api frontend.API, inputs ...frontend.Variable) (frontend.Variable, error) {
	return circuitposeidon.MultiHash(api, inputs...)
}

// ShareMaskHash computes the raw hashed-ElGamal masking scalar before subgroup-order reduction.
func ShareMaskHash(
	api frontend.API,
	roundHash, contributorIndex, recipientIndex, sharedX, sharedY frontend.Variable,
) (frontend.Variable, error) {
	packedIndexes := api.Add(api.Mul(contributorIndex, recipientIndexShift), recipientIndex)
	meta, err := HashFieldElements(api, ShareEncryptionDomain(), roundHash, packedIndexes)
	if err != nil {
		return 0, err
	}
	return HashFieldElements(api, meta, sharedX, sharedY)
}

// ReduceToSubgroupOrder proves value = quotient*subgroupOrder + remainder with
// remainder in [0, subgroupOrder-1] and quotient in [0, 7].
func ReduceToSubgroupOrder(
	api frontend.API,
	value, quotient, remainder frontend.Variable,
) frontend.Variable {
	_ = bits.ToBinary(api, quotient, bits.WithNbDigits(3))
	api.AssertIsLessOrEqual(remainder, subgroupOrderMinusOne)
	api.AssertIsEqual(value, api.Add(remainder, api.Mul(quotient, subgroupOrder)))
	return remainder
}

// AddModSubgroupOrder proves left + right = carry*subgroupOrder + remainder
// with carry in {0,1} and remainder in [0, subgroupOrder-1].
func AddModSubgroupOrder(
	api frontend.API,
	left, right, carry, remainder frontend.Variable,
) frontend.Variable {
	api.AssertIsBoolean(carry)
	api.AssertIsLessOrEqual(remainder, subgroupOrderMinusOne)
	api.AssertIsEqual(api.Add(left, right), api.Add(remainder, api.Mul(carry, subgroupOrder)))
	return remainder
}
