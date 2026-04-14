package revealsubmit

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a reveal-submit witness.
type Assignment struct {
	RoundHash        *big.Int
	ParticipantIndex uint16
	ShareValue       *big.Int
	ShareCommitment  types.CurvePoint
}

// Validate checks that the assignment is complete.
func (a Assignment) Validate() error {
	if a.RoundHash == nil || a.ShareValue == nil {
		return fmt.Errorf("round hash and share value are required")
	}
	if a.ParticipantIndex == 0 {
		return fmt.Errorf("participant index is required")
	}
	if err := a.ShareCommitment.Validate(); err != nil {
		return fmt.Errorf("share commitment: %w", err)
	}
	return nil
}
