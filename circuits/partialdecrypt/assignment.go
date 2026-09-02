package partialdecrypt

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a partial decrypt witness.
//
// Aid and CtIdx are the per-application transcript binding fields. They
// default to zero when callers don't supply them, so pre-application call
// sites (unit tests, benchmarks) keep working unchanged.
type Assignment struct {
	RoundHash        *big.Int // semantically: eid
	Aid              *big.Int // application identifier
	CtIdx            *big.Int // per-application ciphertext index
	ParticipantIndex uint16   // i, one-based committee slot
	Base             types.CurvePoint
	Secret           *big.Int
	Nonce            *big.Int
}

// Validate checks that the assignment is complete. Committee slots are
// one-based, so a zero participant index is always a caller bug.
func (a Assignment) Validate() error {
	if a.RoundHash == nil || a.Secret == nil || a.Nonce == nil {
		return fmt.Errorf("epoch hash, base, secret, and nonce are required")
	}
	if a.ParticipantIndex == 0 {
		return fmt.Errorf("participant index must be non-zero")
	}
	if err := a.Base.Validate(); err != nil {
		return fmt.Errorf("base point: %w", err)
	}
	return nil
}
