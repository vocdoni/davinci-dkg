package helpers

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestRecoverParticipantShares(t *testing.T) {
	c := qt.New(t)

	contributions := [][][]*big.Int{
		DealPoolCoefficients([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		DealPoolCoefficients([]*big.Int{big.NewInt(5), big.NewInt(2)}),
	}

	// Key 0 keeps the base polynomials verbatim: f1(x)=3+x, f2(x)=5+2x.
	shares, err := RecoverParticipantShares(contributions, 0, []uint16{1, 2})
	c.Assert(err, qt.IsNil)
	c.Assert(shares, qt.HasLen, 2)
	c.Assert(shares[0].Cmp(big.NewInt(11)), qt.Equals, 0)
	c.Assert(shares[1].Cmp(big.NewInt(14)), qt.Equals, 0)

	// Every other key is a distinct polynomial, so the shares differ.
	other, err := RecoverParticipantShares(contributions, 1, []uint16{1, 2})
	c.Assert(err, qt.IsNil)
	c.Assert(other[0].Cmp(shares[0]), qt.Not(qt.Equals), 0)

	_, err = RecoverParticipantShares(contributions, ccommon.MaxK, []uint16{1})
	c.Assert(err, qt.IsNotNil)
}

func TestDealPoolCoefficientsFillsThePool(t *testing.T) {
	c := qt.New(t)

	pool := DealPoolCoefficients([]*big.Int{big.NewInt(7), big.NewInt(9)})
	c.Assert(pool, qt.HasLen, ccommon.MaxK)
	for _, key := range pool {
		c.Assert(key, qt.HasLen, 2)
	}
	c.Assert(pool[0][0].Cmp(big.NewInt(7)), qt.Equals, 0)
	c.Assert(pool[0][1].Cmp(big.NewInt(9)), qt.Equals, 0)
}

func TestShareTreeProofMatchesTheRoot(t *testing.T) {
	c := qt.New(t)

	indexes := []uint16{1, 2, 3}
	shares := []*big.Int{big.NewInt(11), big.NewInt(14), big.NewInt(17)}
	tree, err := ShareTreeFromShares(indexes, shares)
	c.Assert(err, qt.IsNil)

	root := tree.Root()
	c.Assert(root, qt.Not(qt.Equals), [32]byte{})

	for i, index := range indexes {
		path, err := tree.Proof(index)
		c.Assert(err, qt.IsNil)
		c.Assert(path, qt.HasLen, ccommon.MerkleDepth)

		// Recompute the root the way DKGManager does: the leaf at index−1,
		// folded with its siblings bottom-up through the tagged node hash.
		node := tree[index-1]
		position := int(index) - 1
		for _, sibling := range path {
			if position%2 == 0 {
				node = ccommon.MerkleNode(node, sibling)
			} else {
				node = ccommon.MerkleNode(sibling, node)
			}
			position /= 2
		}
		c.Assert(node, qt.Equals, root, qt.Commentf("share %d", i))
	}

	// Slots beyond the committee are the fixed empty leaf, never a point.
	c.Assert(tree[len(indexes)], qt.Equals, ccommon.EmptyLeaf)

	_, err = tree.Proof(0)
	c.Assert(err, qt.IsNotNil)
}

func TestCommitteeShareTreeIsMemberIndexed(t *testing.T) {
	c := qt.New(t)

	commitments := []types.CurvePoint{ScalarBasePoint(big.NewInt(11)), ScalarBasePoint(big.NewInt(14))}
	committee, err := CommitteeShareTree(commitments)
	c.Assert(err, qt.IsNil)
	explicit, err := NewShareTree([]uint16{1, 2}, commitments)
	c.Assert(err, qt.IsNil)
	c.Assert(committee.Root(), qt.Equals, explicit.Root())

	leaf, err := ccommon.ShareCommitmentLeaf(commitments[1])
	c.Assert(err, qt.IsNil)
	c.Assert(committee[1], qt.Equals, leaf, qt.Commentf("member 2 sits at leaf 1"))
}

func TestParticipantShareIsPerPoolKey(t *testing.T) {
	c := qt.New(t)

	// The SDK fixture epoch: one contributor, one coefficient per key.
	round := &FinalizedRoundResult{
		ParticipantIndexes: []uint16{1},
		Contributions:      [][][]*big.Int{DealPoolCoefficients([]*big.Int{big.NewInt(FixtureShare)})},
	}

	key0, err := round.ParticipantShare(0, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(key0.Cmp(big.NewInt(FixtureShare)), qt.Equals, 0)

	// Key 1 is dealt from its own polynomial, so the same member holds a
	// different share of it; a partial built from key 0's share would not
	// match key 1's share root.
	key1, err := round.ParticipantShare(1, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(key1.Cmp(key0), qt.Not(qt.Equals), 0)

	tree0, err := ShareTreeFromShares(round.ParticipantIndexes, []*big.Int{key0})
	c.Assert(err, qt.IsNil)
	tree1, err := ShareTreeFromShares(round.ParticipantIndexes, []*big.Int{key1})
	c.Assert(err, qt.IsNil)
	c.Assert(tree0.Root(), qt.Not(qt.Equals), tree1.Root())

	_, err = round.ParticipantShare(ccommon.MaxK, 1)
	c.Assert(err, qt.IsNotNil)
}
