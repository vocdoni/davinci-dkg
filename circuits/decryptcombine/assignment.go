package decryptcombine

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a decrypt-combine witness.
type Assignment struct {
	RoundHash          *big.Int
	Threshold          uint16
	CiphertextC1       types.CurvePoint
	CiphertextC2       types.CurvePoint
	ParticipantIndexes []uint16
	PartialDecryptions []types.CurvePoint
	Plaintext          *big.Int
}

// Validate checks that the assignment is complete.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("round hash is required")
	}
	if a.Threshold == 0 {
		return fmt.Errorf("threshold is required")
	}
	if err := a.CiphertextC1.Validate(); err != nil {
		return fmt.Errorf("ciphertext C1: %w", err)
	}
	if err := a.CiphertextC2.Validate(); err != nil {
		return fmt.Errorf("ciphertext C2: %w", err)
	}
	if len(a.ParticipantIndexes) == 0 || len(a.ParticipantIndexes) != len(a.PartialDecryptions) {
		return fmt.Errorf("participant indexes and partial decryptions must have the same non-zero length")
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
	for i, value := range a.PartialDecryptions {
		if err := value.Validate(); err != nil {
			return fmt.Errorf("partial decryption %d: %w", i, err)
		}
	}
	if a.Plaintext == nil {
		return fmt.Errorf("plaintext is required")
	}
	return nil
}
