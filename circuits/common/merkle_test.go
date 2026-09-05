package common

import (
	"encoding/hex"
	"math/big"
	"testing"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/types"
)

func testShareCommitments(count int) ([]uint16, []types.CurvePoint) {
	indexes := make([]uint16, count)
	points := make([]types.CurvePoint, count)
	for i := range count {
		indexes[i] = uint16(i + 1)
		points[i] = types.CurvePoint{X: big.NewInt(int64(100 + i)), Y: big.NewInt(int64(200 + i))}
	}
	return indexes, points
}

// climb recomputes the root the way the contract does: fold the leaf with its
// siblings bottom-up, choosing the side from the leaf index bits.
func climb(leaf [32]byte, index int, path [MerkleDepth][32]byte) [32]byte {
	node := leaf
	for depth := range MerkleDepth {
		if index&1 == 0 {
			node = MerkleNode(node, path[depth])
		} else {
			node = MerkleNode(path[depth], node)
		}
		index >>= 1
	}
	return node
}

func TestMerklePathVerifies(t *testing.T) {
	c := qt.New(t)

	// A partially-filled committee: slots beyond the committee stay empty.
	indexes, points := testShareCommitments(MaxN - 3)
	leaves, err := ShareCommitmentLeaves(indexes, points)
	c.Assert(err, qt.IsNil)
	for i := MaxN - 3; i < MaxN; i++ {
		c.Assert(leaves[i], qt.Equals, EmptyLeaf)
	}

	root := MerkleRoot(leaves)
	for leafIndex := range MaxN {
		path, err := MerklePath(leaves, leafIndex)
		c.Assert(err, qt.IsNil)
		c.Assert(climb(leaves[leafIndex], leafIndex, path), qt.Equals, root)
	}
}

func TestMerkleRootChangesWithLeaf(t *testing.T) {
	c := qt.New(t)

	indexes, points := testShareCommitments(4)
	leaves, err := ShareCommitmentLeaves(indexes, points)
	c.Assert(err, qt.IsNil)
	root := MerkleRoot(leaves)

	points[2].Y = big.NewInt(999)
	tampered, err := ShareCommitmentLeaves(indexes, points)
	c.Assert(err, qt.IsNil)
	c.Assert(MerkleRoot(tampered) == root, qt.IsFalse)
}

// The tagged encoding keeps the three node kinds apart: a leaf hash is never
// a valid internal node and an absent member is never a point.
func TestMerkleEncodingIsTagged(t *testing.T) {
	c := qt.New(t)

	x, y := big.NewInt(1), big.NewInt(2)
	leaf, err := ShareCommitmentLeaf(types.CurvePoint{X: x, Y: y})
	c.Assert(err, qt.IsNil)
	c.Assert(leaf[:], qt.DeepEquals, ethcrypto.Keccak256(
		[]byte{0x00}, x.FillBytes(make([]byte, 32)), y.FillBytes(make([]byte, 32)),
	))
	var left, right [32]byte
	left[31], right[31] = 1, 2
	node := MerkleNode(left, right)
	c.Assert(node[:], qt.DeepEquals, ethcrypto.Keccak256([]byte{0x01}, left[:], right[:]))
	c.Assert(EmptyLeaf[:], qt.DeepEquals, ethcrypto.Keccak256([]byte("davinci-dkg:merkle-empty:v1")))
	c.Assert(leaf == EmptyLeaf, qt.IsFalse)
}

// Cross-implementation vector: a 3-member committee with
// D_1 = (1, 2), D_2 = (3, 4), D_3 = (5, 6), the remaining MaxN − 3 leaves
// empty. The Solidity and TypeScript trees must produce the same root; the
// hex values are pinned here so a change in the encoding is a test failure
// on every side.
func TestMerkleVector(t *testing.T) {
	c := qt.New(t)

	indexes := []uint16{1, 2, 3}
	points := []types.CurvePoint{
		{X: big.NewInt(1), Y: big.NewInt(2)},
		{X: big.NewInt(3), Y: big.NewInt(4)},
		{X: big.NewInt(5), Y: big.NewInt(6)},
	}
	leaves, err := ShareCommitmentLeaves(indexes, points)
	c.Assert(err, qt.IsNil)

	c.Assert(hex.EncodeToString(EmptyLeaf[:]), qt.Equals,
		"11a12e535a08d28aa7434e11614f2eb9b34da3fcba5746a376ed981855fb01f0")
	c.Assert(hex.EncodeToString(leaves[0][:]), qt.Equals,
		"f08fd8860b9467156f3947e4270748a83c6966d43dd5cb94162189af575662d7")
	c.Assert(hex.EncodeToString(leaves[1][:]), qt.Equals,
		"bf911900795c617010eddba8019fbeb09566bd227a18c05b3151b18e7f28a1c4")
	c.Assert(hex.EncodeToString(leaves[2][:]), qt.Equals,
		"7242095be79d9496003c9bcd83e14cf940aebb29562d12d2b8505a0ec11d4914")
	root := MerkleRoot(leaves)
	c.Assert(hex.EncodeToString(root[:]), qt.Equals,
		"7312971cbb106416972bc40e2e5f311d5d190314dec121522726f8d8235b8bcc")
	// Siblings of leaf 0, bottom-up: the path a partial decryption by
	// participant 1 submits.
	path, err := MerklePath(leaves, 0)
	c.Assert(err, qt.IsNil)
	want := []string{
		"bf911900795c617010eddba8019fbeb09566bd227a18c05b3151b18e7f28a1c4",
		"b74c4ee2822448e5e7cd821528a0401d54f683d89b5abe4a5364a0fb2cb0af90",
		"792fdb07a7f30f6cea4a7d2cbfa86415d1715f91b67edcf7d3e0058afc80397b",
		"d7eff24cec648d254dcb4afb6aa0384710aa970cb4a191c125c1a7a325eea482",
		"3ab752de39c18f6086f960d696a2d35405819452ac9a3db9d1dd8b7e344ce744",
	}
	for depth, sibling := range path {
		c.Assert(hex.EncodeToString(sibling[:]), qt.Equals, want[depth], qt.Commentf("depth %d", depth))
	}
	c.Assert(climb(leaves[0], 0, path), qt.Equals, root)
}

func TestShareCommitmentLeavesRejectsBadInput(t *testing.T) {
	c := qt.New(t)

	_, points := testShareCommitments(2)
	_, err := ShareCommitmentLeaves([]uint16{1}, points)
	c.Assert(err, qt.Not(qt.IsNil))

	_, err = ShareCommitmentLeaves([]uint16{1, 1}, points)
	c.Assert(err, qt.Not(qt.IsNil))

	_, err = ShareCommitmentLeaves([]uint16{0, 2}, points)
	c.Assert(err, qt.Not(qt.IsNil))

	_, err = ShareCommitmentLeaves([]uint16{1, MaxN + 1}, points)
	c.Assert(err, qt.Not(qt.IsNil))
}

func TestMerklePathRejectsOutOfRangeIndex(t *testing.T) {
	c := qt.New(t)

	var leaves [MaxN][32]byte
	_, err := MerklePath(leaves, -1)
	c.Assert(err, qt.Not(qt.IsNil))
	_, err = MerklePath(leaves, MaxN)
	c.Assert(err, qt.Not(qt.IsNil))
}
