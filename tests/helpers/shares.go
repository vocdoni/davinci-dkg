package helpers

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// ShareTree is one pool key's keccak Merkle tree over the per-member share
// commitments: `finalizeEpoch` stores its root and every
// `submitPartialDecryption` proves the member's leaf against it.
type ShareTree [ccommon.MaxN][32]byte

// NewShareTree lays out the leaves from committee members and their share
// commitments D_p (leaf p−1). The encoding — tagged leaves, tagged nodes, a
// fixed empty leaf — lives in circuits/common; nothing is re-derived here.
func NewShareTree(participantIndexes []uint16, shareCommitments []types.CurvePoint) (ShareTree, error) {
	leaves, err := ccommon.ShareCommitmentLeaves(participantIndexes, shareCommitments)
	if err != nil {
		return ShareTree{}, fmt.Errorf("share commitment leaves: %w", err)
	}
	return ShareTree(leaves), nil
}

// CommitteeShareTree is NewShareTree over the whole committee: entry i of
// `shareCommitments` is D_{i+1}, the layout a finalization transcript carries
// (one commitment per member, contributing or not).
func CommitteeShareTree(shareCommitments []types.CurvePoint) (ShareTree, error) {
	indexes := make([]uint16, len(shareCommitments))
	for i := range indexes {
		indexes[i] = uint16(i + 1)
	}
	return NewShareTree(indexes, shareCommitments)
}

// ShareTreeFromShares is the same tree built from the private shares: the
// commitment of member i is D_i = d_i·G.
func ShareTreeFromShares(participantIndexes []uint16, shares []*big.Int) (ShareTree, error) {
	if len(participantIndexes) != len(shares) {
		return ShareTree{}, fmt.Errorf("got %d participant indexes and %d shares", len(participantIndexes), len(shares))
	}
	commitments := make([]types.CurvePoint, len(shares))
	for i, share := range shares {
		if share == nil {
			return ShareTree{}, fmt.Errorf("share %d is nil", i)
		}
		commitments[i] = ScalarBasePoint(share)
	}
	return NewShareTree(participantIndexes, commitments)
}

// Root is the value `finalizeEpoch` stores for the key.
func (t ShareTree) Root() [32]byte {
	return ccommon.MerkleRoot(t)
}

// Proof returns the `bytes32[]` sibling path (bottom-up) that
// `submitPartialDecryption` takes for the given committee member.
func (t ShareTree) Proof(participantIndex uint16) ([][32]byte, error) {
	if participantIndex == 0 || int(participantIndex) > ccommon.MaxN {
		return nil, fmt.Errorf("participant index %d out of range [1, %d]", participantIndex, ccommon.MaxN)
	}
	path, err := ccommon.MerklePath(t, int(participantIndex)-1)
	if err != nil {
		return nil, fmt.Errorf("merkle path for participant %d: %w", participantIndex, err)
	}
	return path[:], nil
}

// RecoverParticipantShares rebuilds d_i = Σ_c f_c(i) for pool key `keyIndex`
// over the given contributions. `contributions` is indexed by contributor,
// then pool key, then coefficient — the same shape BuildFinalizeSubmission
// takes.
func RecoverParticipantShares(
	contributions [][][]*big.Int,
	keyIndex uint8,
	participantIndexes []uint16,
) ([]*big.Int, error) {
	if len(contributions) == 0 {
		return nil, fmt.Errorf("contributions are required")
	}
	if len(participantIndexes) == 0 {
		return nil, fmt.Errorf("participant indexes are required")
	}
	if int(keyIndex) >= ccommon.MaxK {
		return nil, fmt.Errorf("key index %d is outside the pool [0, %d)", keyIndex, ccommon.MaxK)
	}

	modulus := group.ScalarField()
	recovered := make([]*big.Int, len(participantIndexes))
	for i, participantIndex := range participantIndexes {
		sum := big.NewInt(0)
		for c, keys := range contributions {
			if len(keys) != ccommon.MaxK {
				return nil, fmt.Errorf("contributor %d deals %d keys, expected %d", c, len(keys), ccommon.MaxK)
			}
			share, err := ccommon.EvaluatePolynomialNative(keys[keyIndex], big.NewInt(int64(participantIndex)))
			if err != nil {
				return nil, fmt.Errorf("evaluate contribution %d for participant %d: %w", c, participantIndex, err)
			}
			sum.Add(sum, share)
			sum.Mod(sum, modulus)
		}
		recovered[i] = sum
	}

	return recovered, nil
}

// DealPoolCoefficients expands one polynomial into the MaxK polynomials every
// contribution now deals. Key 0 keeps `base` verbatim so fixtures that quote a
// single coefficient list (and the share it implies) stay valid; the other
// keys are deterministic, distinct offsets of it.
func DealPoolCoefficients(base []*big.Int) [][]*big.Int {
	keys := make([][]*big.Int, ccommon.MaxK)
	modulus := group.ScalarField()
	for j := range keys {
		keys[j] = make([]*big.Int, len(base))
		for m, coefficient := range base {
			value := new(big.Int).Add(coefficient, big.NewInt(int64(j*(m+1))))
			keys[j][m] = value.Mod(value, modulus)
		}
	}
	return keys
}
