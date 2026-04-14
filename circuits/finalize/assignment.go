package finalize

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// MaxCoefficients/MaxParticipants are aliases of the single shared constant
// `circuits/common.MaxN`. Edit `circuits/common/sizes.go` to change the bound.
const (
	MaxCoefficients = ccommon.MaxN
	MaxParticipants = ccommon.MaxN
)

// Assignment is the native input model used to build a finalize witness.
type Assignment struct {
	RoundHash                *big.Int
	Threshold                uint16
	CommitteeSize            uint16
	ParticipantIndexes       []uint16
	ContributionCoefficients [][]*big.Int
}

// Validate checks that the assignment fits the current circuit bounds.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("round hash is required")
	}
	if a.Threshold == 0 || a.CommitteeSize == 0 {
		return fmt.Errorf("threshold and committee size are required")
	}
	if a.Threshold > a.CommitteeSize {
		return fmt.Errorf("threshold cannot exceed committee size")
	}
	if len(a.ParticipantIndexes) < int(a.Threshold) {
		return fmt.Errorf("participant count %d is below threshold %d", len(a.ParticipantIndexes), a.Threshold)
	}
	if len(a.ParticipantIndexes) > int(a.CommitteeSize) {
		return fmt.Errorf("participant count %d exceeds committee size %d", len(a.ParticipantIndexes), a.CommitteeSize)
	}
	if len(a.ParticipantIndexes) > MaxParticipants {
		return fmt.Errorf("participant count %d exceeds max %d", len(a.ParticipantIndexes), MaxParticipants)
	}
	if len(a.ContributionCoefficients) != len(a.ParticipantIndexes) {
		return fmt.Errorf(
			"contribution count %d does not match participant count %d",
			len(a.ContributionCoefficients),
			len(a.ParticipantIndexes),
		)
	}
	for i, index := range a.ParticipantIndexes {
		if index == 0 {
			return fmt.Errorf("participant index %d is zero", i)
		}
	}
	for i, contribution := range a.ContributionCoefficients {
		if len(contribution) != int(a.Threshold) {
			return fmt.Errorf(
				"contribution %d coefficient count %d does not match threshold %d",
				i,
				len(contribution),
				a.Threshold,
			)
		}
		for j, coefficient := range contribution {
			if coefficient == nil {
				return fmt.Errorf("contribution %d coefficient %d is nil", i, j)
			}
		}
	}
	return nil
}
