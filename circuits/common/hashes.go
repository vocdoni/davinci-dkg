package common

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
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
	// reduceQuotientMax = (p-1)/r = 7. The wider envelope p < 8r leaves a
	// tiny gap δ := p - 7r at the top of the q=7 stratum where
	// `value = 7r + remainder (mod p)` admits a second non-canonical
	// decomposition. ReduceToSubgroupOrder closes that gap by also asserting
	// `q==7 ⇒ remainder < δ`.
	deltaMinusOne = new(big.Int).Sub(
		new(big.Int).Sub(ecc.BN254.ScalarField(), new(big.Int).Mul(big.NewInt(7), subgroupOrder)),
		big.NewInt(1),
	)
)

// SubgroupOrderMinusOne returns r-1 as a *big.Int. Used by callers that need
// to range-check witnesses against the BabyJubJub scalar field.
func SubgroupOrderMinusOne() *big.Int {
	return new(big.Int).Set(subgroupOrderMinusOne)
}

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
//
// Because p = 7·r + δ with 0 < δ < r, the inner equality `value = q·r + r'`
// is taken modulo p, and a single rawMask in [0, 8r − p) admits two
// q-decompositions. To force the *unique* canonical decomposition we add the
// auxiliary constraint `q == 7  ⇒  remainder < δ`, which together with
// `q ≤ 7` and `remainder < r` makes `q·r + remainder < p` over the integers.
func ReduceToSubgroupOrder(
	api frontend.API,
	value, quotient, remainder frontend.Variable,
) frontend.Variable {
	_ = bits.ToBinary(api, quotient, bits.WithNbDigits(3))
	api.AssertIsLessOrEqual(remainder, subgroupOrderMinusOne)
	api.AssertIsEqual(value, api.Add(remainder, api.Mul(quotient, subgroupOrder)))
	isSeven := api.IsZero(api.Sub(quotient, 7))
	api.AssertIsLessOrEqual(api.Mul(isSeven, remainder), deltaMinusOne)
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
