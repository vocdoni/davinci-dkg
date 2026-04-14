package feldman

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// Commitments returns Feldman coefficient commitments over the DKG generator.
func Commitments(coefficients []*big.Int) ([]types.CurvePoint, error) {
	if len(coefficients) == 0 {
		return nil, fmt.Errorf("at least one coefficient is required")
	}

	commitments := make([]types.CurvePoint, len(coefficients))
	for i, coefficient := range coefficients {
		if coefficient == nil {
			return nil, fmt.Errorf("coefficient %d is nil", i)
		}

		point := group.NewPoint()
		point.ScalarBaseMult(coefficient)
		commitments[i] = group.Encode(point)
	}
	return commitments, nil
}
