package revealshare

import (
	"fmt"
	"math/big"
)

// Assignment is the native input model used to build a reveal-share witness.
type Assignment struct {
	RoundHash          *big.Int
	Threshold          uint16
	ParticipantIndexes []uint16
	RevealedShares     []*big.Int
}

// Validate checks that the assignment is complete.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("round hash is required")
	}
	if a.Threshold == 0 {
		return fmt.Errorf("threshold is required")
	}
	if len(a.ParticipantIndexes) == 0 || len(a.ParticipantIndexes) != len(a.RevealedShares) {
		return fmt.Errorf("participant indexes and revealed shares must have the same non-zero length")
	}
	if len(a.ParticipantIndexes) != int(a.Threshold) {
		return fmt.Errorf("share count %d does not match threshold %d", len(a.ParticipantIndexes), a.Threshold)
	}
	if len(a.ParticipantIndexes) > MaxShares {
		return fmt.Errorf("share count %d exceeds max %d", len(a.ParticipantIndexes), MaxShares)
	}
	for i, index := range a.ParticipantIndexes {
		if index == 0 {
			return fmt.Errorf("participant index %d is zero", i)
		}
	}
	for i, value := range a.RevealedShares {
		if value == nil {
			return fmt.Errorf("revealed share %d is nil", i)
		}
	}
	return nil
}
