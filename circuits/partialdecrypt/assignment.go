package partialdecrypt

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a partial decrypt witness.
//
// Aid, CtIdx, and Role are the per-application transcript binding fields
// (paper §4.4 lines 695–704). They default to zero / committee
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

// Role constants — mirror solidity/src/libraries/DKGProtocol.sol and
// internal/protocol/protocol.go. Kept inline here so the witness
// validator can be role-aware without pulling in the cross-layer
// protocol package.
const (
	RoleCommittee = 1
	RoleOrganizer = 2
)

// Validate checks that the assignment is complete and that the
// participant-index policy matches the chosen role:
//
//   - Default / role = 1 (COMMITTEE): participantIndex MUST be non-zero
//     (committee slots are one-based, 1..n).
//   - Role = 2 (ORGANIZER): participantIndex MUST be exactly 0
//     (paper §6.3 line 1161 and solidity/src/DKGManager.sol's
//     submitOrganizerShare check).
func (a Assignment) Validate() error {
	if a.RoundHash == nil || a.Secret == nil || a.Nonce == nil {
		return fmt.Errorf("epoch hash, base, secret, and nonce are required")
	}
	role := int64(RoleCommittee)
	if a.Role != nil {
		role = a.Role.Int64()
	}
	switch role {
	case RoleCommittee:
		if a.ParticipantIndex == 0 {
			return fmt.Errorf("committee role requires non-zero participant index")
		}
	case RoleOrganizer:
		if a.ParticipantIndex != 0 {
			return fmt.Errorf("organizer role requires participant index 0, got %d", a.ParticipantIndex)
		}
	default:
		return fmt.Errorf("unknown role %d (expected %d=COMMITTEE or %d=ORGANIZER)",
			role, RoleCommittee, RoleOrganizer)
	}
	if err := a.Base.Validate(); err != nil {
		return fmt.Errorf("base point: %w", err)
	}
	return nil
}
