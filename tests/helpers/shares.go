package helpers

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

func RecoverParticipantShares(contributions [][]*big.Int, participantIndexes []uint16) ([]*big.Int, error) {
	if len(contributions) == 0 {
		return nil, fmt.Errorf("contributions are required")
	}
	if len(participantIndexes) == 0 {
		return nil, fmt.Errorf("participant indexes are required")
	}

	modulus := group.ScalarField()
	recovered := make([]*big.Int, len(participantIndexes))
	for i, participantIndex := range participantIndexes {
		sum := big.NewInt(0)
		for j, coefficients := range contributions {
			share, err := ccommon.EvaluatePolynomialNative(coefficients, big.NewInt(int64(participantIndex)))
			if err != nil {
				return nil, fmt.Errorf("evaluate contribution %d for participant %d: %w", j, participantIndex, err)
			}
			sum.Add(sum, share)
			sum.Mod(sum, modulus)
		}
		recovered[i] = sum
	}

	return recovered, nil
}
