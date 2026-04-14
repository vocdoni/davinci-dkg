package partialdecrypt

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a partial decrypt witness.
type Assignment struct {
	RoundHash        *big.Int
	ParticipantIndex uint16
	Base             types.CurvePoint
	Secret           *big.Int
	Nonce            *big.Int
}

// Validate checks that the assignment is complete.
func (a Assignment) Validate() error {
	if a.RoundHash == nil || a.Secret == nil || a.Nonce == nil {
		return fmt.Errorf("round hash, base, secret, and nonce are required")
	}
	if a.ParticipantIndex == 0 {
		return fmt.Errorf("participant index is required")
	}
	if err := a.Base.Validate(); err != nil {
		return fmt.Errorf("base point: %w", err)
	}
	return nil
}
