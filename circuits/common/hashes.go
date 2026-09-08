package common

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
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

// ShareMaskSeed derives the per-recipient KDF seed that all MaxK share masks
// of one (dealer, recipient) pair expand from: H(domain, roundHash,
// contributorIndex·2^16 + recipientIndex, S.x, S.y) over the ECDH secret S.
// Mirrors crypto/shareenc.ShareMaskSeed.
func ShareMaskSeed(
	api frontend.API,
	roundHash, contributorIndex, recipientIndex, sharedX, sharedY frontend.Variable,
) (frontend.Variable, error) {
	packedIndexes := api.Add(api.Mul(contributorIndex, recipientIndexShift), recipientIndex)
	return HashFieldElements(api, ShareEncryptionDomain(), roundHash, packedIndexes, sharedX, sharedY)
}

// ShareMask expands a recipient's seed for pool key keyIndex: H(seed, j).
// The mask is used as is, a uniform element of F_p added to the share in the
// native field. Mirrors crypto/shareenc.ShareMask.
func ShareMask(api frontend.API, seed, keyIndex frontend.Variable) (frontend.Variable, error) {
	return HashFieldElements(api, seed, keyIndex)
}
