package contribution

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a contribution witness.
type Assignment struct {
	RoundHash        *big.Int
	Threshold        uint16
	CommitteeSize    uint16
	ContributorIndex uint16
	Coefficients     []*big.Int
	RecipientIndexes []uint16
	RecipientKeys    []types.NodeKey
	EncryptionNonces []*big.Int
}

// Validate checks that the assignment fits the current circuit bounds.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("round hash is required")
	}
	if a.Threshold == 0 {
		return fmt.Errorf("threshold is required")
	}
	if a.CommitteeSize == 0 {
		return fmt.Errorf("committee size is required")
	}
	if a.ContributorIndex == 0 {
		return fmt.Errorf("contributor index is required")
	}
	if len(a.Coefficients) == 0 {
		return fmt.Errorf("coefficients are required")
	}
	if int(a.Threshold) != len(a.Coefficients) {
		return fmt.Errorf("threshold %d does not match coefficient count %d", a.Threshold, len(a.Coefficients))
	}
	if len(a.Coefficients) > MaxCoefficients {
		return fmt.Errorf("coefficient count %d exceeds max %d", len(a.Coefficients), MaxCoefficients)
	}
	if int(a.CommitteeSize) != len(a.RecipientIndexes) {
		return fmt.Errorf("committee size %d does not match recipient count %d", a.CommitteeSize, len(a.RecipientIndexes))
	}
	if len(a.RecipientIndexes) > MaxRecipients {
		return fmt.Errorf("recipient count %d exceeds max %d", len(a.RecipientIndexes), MaxRecipients)
	}
	if len(a.RecipientKeys) != 0 && len(a.RecipientKeys) != len(a.RecipientIndexes) {
		return fmt.Errorf("recipient key count %d does not match recipient count %d", len(a.RecipientKeys), len(a.RecipientIndexes))
	}
	if len(a.EncryptionNonces) != 0 && len(a.EncryptionNonces) != len(a.RecipientIndexes) {
		return fmt.Errorf(
			"encryption nonce count %d does not match recipient count %d",
			len(a.EncryptionNonces),
			len(a.RecipientIndexes),
		)
	}
	for i, coefficient := range a.Coefficients {
		if coefficient == nil {
			return fmt.Errorf("coefficient %d is nil", i)
		}
	}
	for i, index := range a.RecipientIndexes {
		if index == 0 {
			return fmt.Errorf("recipient index %d is zero", i)
		}
	}
	for i, key := range a.RecipientKeys {
		if key.PubX == nil || key.PubY == nil {
			return fmt.Errorf("recipient key %d is missing coordinates", i)
		}
	}
	for i, nonce := range a.EncryptionNonces {
		if nonce == nil {
			return fmt.Errorf("encryption nonce %d is nil", i)
		}
	}
	return nil
}
