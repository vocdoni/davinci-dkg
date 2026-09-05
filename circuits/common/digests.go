package common

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/vocdoni/davinci-dkg/types"
)

// The contribution digests are two-level: one Poseidon per pool key over its
// commitment vector, one per recipient row over that row's MaxK masked shares,
// then one sponge over the digests. A flat absorption would blow past the
// sponge's 256-input cap at MaxK·MaxN inputs. docs/pool-keys.md pins the
// formulas (unchanged by v4, which only compacts the calldata transcript);
// the finalize circuit recomputes CommitmentKeyDigest and CommitmentsHash to
// reproduce a dealer's on-chain commitments hash, so both sides must read them
// from here. Digests always absorb the padded vectors: inactive scalars are
// zero and inactive points are the identity (0, 1).

// CommitmentKeyDigest hashes one pool key's commitment vector. Callers pass
// coefficient slots already masked to the identity (0, 1) beyond the threshold.
func CommitmentKeyDigest(api frontend.API, commitments []twistededwards.Point) (frontend.Variable, error) {
	inputs := make([]frontend.Variable, 0, 2*len(commitments))
	for _, commitment := range commitments {
		inputs = append(inputs, commitment.X, commitment.Y)
	}
	return MultiHash(api, inputs...)
}

// CommitmentKeyDigestNative mirrors CommitmentKeyDigest for witness builders.
func CommitmentKeyDigestNative(commitments []types.CurvePoint) (*big.Int, error) {
	inputs := make([]*big.Int, 0, 2*len(commitments))
	for i, commitment := range commitments {
		if err := commitment.Validate(); err != nil {
			return nil, fmt.Errorf("commitment %d: %w", i, err)
		}
		inputs = append(inputs, commitment.X, commitment.Y)
	}
	return MultiHashNative(inputs...)
}

// CommitmentsHash binds a contributor's MaxK key digests to the epoch, its
// one-based index and the threshold. This is the value the contract stores as
// ContributionRecord.commitmentsHash.
func CommitmentsHash(
	api frontend.API,
	roundHash, contributorIndex, threshold frontend.Variable,
	keyDigests []frontend.Variable,
) (frontend.Variable, error) {
	inputs := make([]frontend.Variable, 0, 3+len(keyDigests))
	inputs = append(inputs, roundHash, contributorIndex, threshold)
	inputs = append(inputs, keyDigests...)
	return MultiHash(api, inputs...)
}

// CommitmentsHashNative mirrors CommitmentsHash for witness builders.
func CommitmentsHashNative(roundHash, contributorIndex, threshold *big.Int, keyDigests []*big.Int) (*big.Int, error) {
	inputs := make([]*big.Int, 0, 3+len(keyDigests))
	inputs = append(inputs, roundHash, contributorIndex, threshold)
	inputs = append(inputs, keyDigests...)
	return MultiHashNative(inputs...)
}

// EncryptedShareRowDigest hashes one recipient row: its index, its public key,
// the ephemeral shared by every pool key and the MaxK masked shares it gets.
func EncryptedShareRowDigest(
	api frontend.API,
	recipientIndex frontend.Variable,
	recipientKey, ephemeral twistededwards.Point,
	maskedShares []frontend.Variable,
) (frontend.Variable, error) {
	inputs := make([]frontend.Variable, 0, 5+len(maskedShares))
	inputs = append(inputs, recipientIndex, recipientKey.X, recipientKey.Y, ephemeral.X, ephemeral.Y)
	inputs = append(inputs, maskedShares...)
	return MultiHash(api, inputs...)
}

// EncryptedShareRowDigestNative mirrors EncryptedShareRowDigest natively.
func EncryptedShareRowDigestNative(
	recipientIndex *big.Int,
	recipientKey, ephemeral types.CurvePoint,
	maskedShares []*big.Int,
) (*big.Int, error) {
	if err := recipientKey.Validate(); err != nil {
		return nil, fmt.Errorf("recipient key: %w", err)
	}
	if err := ephemeral.Validate(); err != nil {
		return nil, fmt.Errorf("ephemeral: %w", err)
	}
	inputs := make([]*big.Int, 0, 5+len(maskedShares))
	inputs = append(inputs, recipientIndex, recipientKey.X, recipientKey.Y, ephemeral.X, ephemeral.Y)
	inputs = append(inputs, maskedShares...)
	return MultiHashNative(inputs...)
}

// EncryptedSharesHash binds the per-recipient row digests to the epoch, the
// contributor index and the committee size.
func EncryptedSharesHash(
	api frontend.API,
	roundHash, contributorIndex, committeeSize frontend.Variable,
	rowDigests []frontend.Variable,
) (frontend.Variable, error) {
	inputs := make([]frontend.Variable, 0, 3+len(rowDigests))
	inputs = append(inputs, roundHash, contributorIndex, committeeSize)
	inputs = append(inputs, rowDigests...)
	return MultiHash(api, inputs...)
}

// EncryptedSharesHashNative mirrors EncryptedSharesHash for witness builders.
func EncryptedSharesHashNative(
	roundHash, contributorIndex, committeeSize *big.Int,
	rowDigests []*big.Int,
) (*big.Int, error) {
	inputs := make([]*big.Int, 0, 3+len(rowDigests))
	inputs = append(inputs, roundHash, contributorIndex, committeeSize)
	inputs = append(inputs, rowDigests...)
	return MultiHashNative(inputs...)
}
