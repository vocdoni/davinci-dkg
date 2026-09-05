package contribution

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
)

// Assignment is the native input model used to build a contribution witness.
// Coefficients is indexed by pool key first and coefficient second: one
// degree-(Threshold-1) polynomial per key, all MaxKeys of them dealt in the
// same proof under one ephemeral per recipient.
type Assignment struct {
	RoundHash        *big.Int
	Threshold        uint16
	CommitteeSize    uint16
	ContributorIndex uint16
	Coefficients     [][]*big.Int
	RecipientIndexes []uint16
	RecipientKeys    []types.NodeKey
	EncryptionNonces []*big.Int
}

// Validate checks that the assignment fits the current circuit bounds.
func (a Assignment) Validate() error {
	if a.RoundHash == nil {
		return fmt.Errorf("epoch hash is required")
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
	// Every epoch deals exactly MaxKeys polynomials; unused keys only cost
	// calldata, so a short set is a bug rather than a cheaper contribution.
	if len(a.Coefficients) != MaxKeys {
		return fmt.Errorf("got %d coefficient sets, expected one per pool key (%d)", len(a.Coefficients), MaxKeys)
	}
	if int(a.Threshold) > MaxCoefficients {
		return fmt.Errorf("threshold %d exceeds max %d", a.Threshold, MaxCoefficients)
	}
	for j, coefficients := range a.Coefficients {
		if int(a.Threshold) != len(coefficients) {
			return fmt.Errorf("key %d: threshold %d does not match coefficient count %d", j, a.Threshold, len(coefficients))
		}
		for m, coefficient := range coefficients {
			if coefficient == nil {
				return fmt.Errorf("key %d coefficient %d is nil", j, m)
			}
		}
	}
	if int(a.CommitteeSize) != len(a.RecipientIndexes) {
		return fmt.Errorf("committee size %d does not match recipient count %d", a.CommitteeSize, len(a.RecipientIndexes))
	}
	if len(a.RecipientIndexes) > MaxRecipients {
		return fmt.Errorf("recipient count %d exceeds max %d", len(a.RecipientIndexes), MaxRecipients)
	}
	if len(a.RecipientKeys) != len(a.RecipientIndexes) {
		return fmt.Errorf("recipient key count %d does not match recipient count %d", len(a.RecipientKeys), len(a.RecipientIndexes))
	}
	if len(a.EncryptionNonces) != len(a.RecipientIndexes) {
		return fmt.Errorf(
			"encryption nonce count %d does not match recipient count %d",
			len(a.EncryptionNonces),
			len(a.RecipientIndexes),
		)
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
