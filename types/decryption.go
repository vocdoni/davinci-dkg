package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// Role tags partial decryptions with their producer per paper §6.3 line 1161.
// Mirrors `solidity/src/libraries/DKGTypes.sol::Role` and
// `internal/protocol/protocol.go::Role`.
type Role uint8

const (
	RoleNone      Role = 0
	RoleCommittee Role = 1
	RoleOrganizer Role = 2
)

// PartialDecryption is one participant's decryption share plus proof material.
//
// `AID` and `Role` are P9 additions for per-application decryption (paper §4.4
// lines 695–704, paper §6.3 line 1161). They default to zero / RoleCommittee
// for backward-compatible callers that operate at the per-epoch level only.
type PartialDecryption struct {
	EpochID          string
	AID              [32]byte // application identifier (zeros for legacy per-epoch path)
	Role             Role     // 1 = COMMITTEE (default), 2 = ORGANIZER
	Participant      common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
	Delta            CurvePoint
	Proof            []byte
}

// Validate checks that the partial decryption payload is minimally complete.
func (d PartialDecryption) Validate() error {
	if d.EpochID == "" {
		return fmt.Errorf("epoch id is required")
	}
	if d.Participant == (common.Address{}) {
		return fmt.Errorf("participant is required")
	}
	if d.ParticipantIndex == 0 {
		return fmt.Errorf("participant index is required")
	}
	if d.CiphertextIndex == 0 {
		return fmt.Errorf("ciphertext index is required")
	}
	if err := d.Delta.Validate(); err != nil {
		return fmt.Errorf("delta point: %w", err)
	}
	if len(d.Proof) == 0 {
		return fmt.Errorf("proof is required")
	}
	return nil
}
