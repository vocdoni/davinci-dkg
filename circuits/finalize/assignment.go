package finalize

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a finalization witness:
// the accepted contributions of one epoch, read back from their calldata.
// Commitments is indexed by accepted dealer, then pool key, then coefficient.
// ContributionHashes optionally carries each dealer's on-chain
// commitmentsHash; when present the builder checks the recomputed hash
// against it so a mis-decoded contribution fails early with a message
// instead of as an unsatisfiable witness.
type Assignment struct {
	RoundHash          *big.Int
	Threshold          uint16
	CommitteeSize      uint16
	ParticipantIndexes []uint16
	Commitments        [][][]types.CurvePoint
	ContributionHashes []*big.Int
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
	if len(a.ParticipantIndexes) < int(a.Threshold) {
		return fmt.Errorf("participant count %d is below threshold %d", len(a.ParticipantIndexes), a.Threshold)
	}
	if len(a.ParticipantIndexes) > int(a.CommitteeSize) {
		return fmt.Errorf("participant count %d exceeds committee size %d", len(a.ParticipantIndexes), a.CommitteeSize)
	}
	// Reject duplicate participant indexes so local tooling cannot produce an
	// accepted set the contract would refuse.
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
	if len(a.ContributionHashes) != 0 && len(a.ContributionHashes) != len(a.ParticipantIndexes) {
		return fmt.Errorf(
			"contribution hash count %d does not match participant count %d",
			len(a.ContributionHashes),
			len(a.ParticipantIndexes),
		)
	}
	for i, hash := range a.ContributionHashes {
		if hash == nil {
			return fmt.Errorf("contribution hash %d is nil", i)
		}
	}
	for i, dealer := range a.Commitments {
		if len(dealer) != MaxKeys {
			return fmt.Errorf("dealer %d deals %d keys, expected %d", i, len(dealer), MaxKeys)
		}
		for j, key := range dealer {
			if len(key) != int(a.Threshold) {
				return fmt.Errorf(
					"dealer %d key %d has %d commitments, expected the threshold %d",
					i,
					j,
					len(key),
					a.Threshold,
				)
			}
			for m, point := range key {
				if err := point.Validate(); err != nil {
					return fmt.Errorf("dealer %d key %d commitment %d: %w", i, j, m, err)
				}
			}
		}
	}
	return nil
}
