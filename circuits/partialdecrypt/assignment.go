package partialdecrypt

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a partial decrypt witness.
//
// Aid, CtIdx, and Role are the per-application transcript binding fields
// (paper §4.4 lines 695–704, PLAN §5.3). They default to zero / committee
// when callers don't supply them, preserving backward compatibility with
// pre-application call sites until those are migrated to per-app workers.
type Assignment struct {
	RoundHash        *big.Int // semantically: eid
	Aid              *big.Int // application identifier (0 for legacy callers)
	CtIdx            *big.Int // per-application ciphertext index (0 for legacy callers)
	Role             *big.Int // 1 = COMMITTEE (default), 2 = ORGANIZER
	ParticipantIndex uint16
	Base             types.CurvePoint
	Secret           *big.Int
	Nonce            *big.Int
}

// Validate checks that the assignment is complete.
func (a Assignment) Validate() error {
	if a.RoundHash == nil || a.Secret == nil || a.Nonce == nil {
		return fmt.Errorf("epoch hash, base, secret, and nonce are required")
	}
	if a.ParticipantIndex == 0 {
		return fmt.Errorf("participant index is required")
	}
	if err := a.Base.Validate(); err != nil {
		return fmt.Errorf("base point: %w", err)
	}
	return nil
}
