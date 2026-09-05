package common

import (
	"fmt"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/vocdoni/davinci-dkg/types"
)

// Compile-time guard: the share-commitment tree has exactly MaxN leaves, so
// MerkleDepth must stay log2(MaxN). Either term underflows uint otherwise.
const _ = uint(1<<MerkleDepth-MaxN) + uint(MaxN-1<<MerkleDepth)

// The keccak Merkle tree over one pool key's per-member share commitments:
// activatePoolKey stores its root and submitPartialDecryption proves the
// member's leaf against it. Keccak (not Poseidon) because only the contract
// ever recomputes a path. See docs/pool-keys.md.
//
// Tagged encoding, so no leaf can be replayed as an internal node (or the
// other way round) and an absent member is distinguishable from any point:
//
//	leaf[p-1] = keccak256(0x00 ‖ D_p.x ‖ D_p.y)      for committee member p
//	leaf[i]   = EmptyLeaf                           for i ≥ committeeSize
//	node      = keccak256(0x01 ‖ left ‖ right)
//
// Coordinates are 32-byte big-endian words; the tree is a fixed MerkleDepth
// levels over MaxN leaves, indexed by participant index − 1.

const (
	merkleLeafTag byte = 0x00
	merkleNodeTag byte = 0x01
)

// EmptyLeaf fills every leaf beyond the committee. It is a fixed domain
// hash rather than bytes32(0) so an empty slot is never confused with a
// leaf whose hash happens to be zero-like, and the constant is what the
// contract must mirror.
var EmptyLeaf = [32]byte(ethcrypto.Keccak256([]byte("davinci-dkg:merkle-empty:v1")))

// ShareCommitmentLeaf encodes one member's share commitment as its leaf.
func ShareCommitmentLeaf(commitment types.CurvePoint) ([32]byte, error) {
	if err := commitment.Validate(); err != nil {
		return [32]byte{}, err
	}
	packed := make([]byte, 0, 65)
	packed = append(packed, merkleLeafTag)
	packed = append(packed, commitment.X.FillBytes(make([]byte, 32))...)
	packed = append(packed, commitment.Y.FillBytes(make([]byte, 32))...)
	return [32]byte(ethcrypto.Keccak256(packed)), nil
}

// MerkleNode hashes two children into their parent.
func MerkleNode(left, right [32]byte) [32]byte {
	packed := make([]byte, 0, 65)
	packed = append(packed, merkleNodeTag)
	packed = append(packed, left[:]...)
	packed = append(packed, right[:]...)
	return [32]byte(ethcrypto.Keccak256(packed))
}

// ShareCommitmentLeaves lays out one pool key's share commitments as the MaxN
// tree leaves: leaf[p-1] = ShareCommitmentLeaf(D_p) for every listed
// participant p, EmptyLeaf for every other slot. Callers pass the whole
// committee (indexes 1..committeeSize), because the activation transcript
// carries a D_p for every member, contributing or not.
func ShareCommitmentLeaves(
	participantIndexes []uint16,
	shareCommitments []types.CurvePoint,
) ([MaxN][32]byte, error) {
	var leaves [MaxN][32]byte
	for i := range leaves {
		leaves[i] = EmptyLeaf
	}
	if len(participantIndexes) != len(shareCommitments) {
		return leaves, fmt.Errorf(
			"got %d participant indexes and %d share commitments",
			len(participantIndexes),
			len(shareCommitments),
		)
	}
	seen := make(map[uint16]struct{}, len(participantIndexes))
	for i, index := range participantIndexes {
		if index == 0 || int(index) > MaxN {
			return leaves, fmt.Errorf("participant index %d out of range [1, %d]", index, MaxN)
		}
		if _, duplicate := seen[index]; duplicate {
			return leaves, fmt.Errorf("duplicate participant index %d", index)
		}
		seen[index] = struct{}{}
		leaf, err := ShareCommitmentLeaf(shareCommitments[i])
		if err != nil {
			return leaves, fmt.Errorf("share commitment %d: %w", i, err)
		}
		leaves[index-1] = leaf
	}
	return leaves, nil
}

// MerkleRoot folds the leaves into the root over MerkleDepth levels of
// MerkleNode.
func MerkleRoot(leaves [MaxN][32]byte) [32]byte {
	level := make([][32]byte, MaxN)
	copy(level, leaves[:])
	for range MerkleDepth {
		level = merkleLevel(level)
	}
	return level[0]
}

// MerklePath returns the siblings of leaf `index` bottom-up: the proof
// submitPartialDecryption takes alongside the partial.
func MerklePath(leaves [MaxN][32]byte, index int) ([MerkleDepth][32]byte, error) {
	var path [MerkleDepth][32]byte
	if index < 0 || index >= MaxN {
		return path, fmt.Errorf("leaf index %d out of range [0, %d)", index, MaxN)
	}
	level := make([][32]byte, MaxN)
	copy(level, leaves[:])
	for depth := range MerkleDepth {
		path[depth] = level[index^1]
		index >>= 1
		level = merkleLevel(level)
	}
	return path, nil
}

func merkleLevel(nodes [][32]byte) [][32]byte {
	next := make([][32]byte, len(nodes)/2)
	for i := range next {
		next[i] = MerkleNode(nodes[2*i], nodes[2*i+1])
	}
	return next
}
