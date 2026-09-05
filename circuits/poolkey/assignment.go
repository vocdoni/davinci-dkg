package poolkey

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a pool-key activation
// witness. Commitments is indexed by accepted contributor, then pool key, then
// coefficient: activation is permissionless and reads every set from the
// contribution calldata, because reproducing a contributor's stored
// commitments hash needs the digests of the keys that are not being activated.
type Assignment struct {
	RoundHash          *big.Int
	Threshold          uint16
	CommitteeSize      uint16
	KeyIndex           uint8
	ParticipantIndexes []uint16
	Commitments        [][][]types.CurvePoint
}

// Validate checks that the assignment fits the current circuit bounds.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("epoch hash is required")
	}
	if a.Threshold == 0 || a.CommitteeSize == 0 {
		return fmt.Errorf("threshold and committee size are required")
	}
	if a.Threshold > a.CommitteeSize {
		return fmt.Errorf("threshold cannot exceed committee size")
	}
	if int(a.Threshold) > MaxCoefficients {
		return fmt.Errorf("threshold %d exceeds max %d", a.Threshold, MaxCoefficients)
	}
	if int(a.CommitteeSize) > MaxParticipants {
		return fmt.Errorf("committee size %d exceeds max %d", a.CommitteeSize, MaxParticipants)
	}
	if int(a.KeyIndex) >= MaxKeys {
		return fmt.Errorf("key index %d is outside the pool [0, %d)", a.KeyIndex, MaxKeys)
	}
	if len(a.ParticipantIndexes) < int(a.Threshold) {
		return fmt.Errorf("participant count %d is below threshold %d", len(a.ParticipantIndexes), a.Threshold)
	}
	if len(a.ParticipantIndexes) > int(a.CommitteeSize) {
		return fmt.Errorf("participant count %d exceeds committee size %d", len(a.ParticipantIndexes), a.CommitteeSize)
	}
	// Reject duplicate participant indexes so local tooling cannot produce an
	// accepted set the contract would refuse (and that a malicious prover could
	// otherwise use to activate a key over a set disjoint from the epoch's).
	seen := make(map[uint16]struct{}, len(a.ParticipantIndexes))
	for i, index := range a.ParticipantIndexes {
		if index == 0 {
			return fmt.Errorf("participant index %d is zero", i)
		}
		if index > a.CommitteeSize {
			return fmt.Errorf("participant index %d exceeds committee size %d", index, a.CommitteeSize)
		}
		if _, duplicate := seen[index]; duplicate {
			return fmt.Errorf("duplicate participant index %d", index)
		}
		seen[index] = struct{}{}
	}
	if len(a.Commitments) != len(a.ParticipantIndexes) {
		return fmt.Errorf(
			"commitment set count %d does not match participant count %d",
			len(a.Commitments),
			len(a.ParticipantIndexes),
		)
	}
	for i, contributor := range a.Commitments {
		if len(contributor) != MaxKeys {
			return fmt.Errorf("contributor %d deals %d keys, expected %d", i, len(contributor), MaxKeys)
		}
		for j, key := range contributor {
			if len(key) != int(a.Threshold) {
				return fmt.Errorf(
					"contributor %d key %d has %d commitments, expected the threshold %d",
					i,
					j,
					len(key),
					a.Threshold,
				)
			}
			for m, point := range key {
				if err := point.Validate(); err != nil {
					return fmt.Errorf("contributor %d key %d commitment %d: %w", i, j, m, err)
				}
			}
		}
	}
	return nil
}
