package types

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// PartialDecryption is one participant's decryption share plus proof material.
type PartialDecryption struct {
	RoundID          string
	Participant      common.Address
	ParticipantIndex uint16
	CiphertextIndex  uint16
	Delta            CurvePoint
	Proof            []byte
}

// Validate checks that the partial decryption payload is minimally complete.
func (d PartialDecryption) Validate() error {
	if d.RoundID == "" {
		return fmt.Errorf("round id is required")
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

// RevealedShare is the emergency disclosure path for one participant share.
type RevealedShare struct {
	RoundID          string
	Participant      common.Address
	ParticipantIndex uint16
	Share            *big.Int
}

// Validate checks that the reveal-share payload is minimally complete.
func (s RevealedShare) Validate() error {
	if s.RoundID == "" {
		return fmt.Errorf("round id is required")
	}
	if s.Participant == (common.Address{}) {
		return fmt.Errorf("participant is required")
	}
	if s.ParticipantIndex == 0 {
		return fmt.Errorf("participant index is required")
	}
	if s.Share == nil {
		return fmt.Errorf("share is required")
	}
	return nil
}
