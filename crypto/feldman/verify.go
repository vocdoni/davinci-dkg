package feldman

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// VerifyShare checks one Shamir share against the Feldman commitment vector.
func VerifyShare(commitments []types.CurvePoint, index uint16, share *big.Int) error {
	if len(commitments) == 0 {
		return fmt.Errorf("at least one commitment is required")
	}
	if index == 0 {
		return fmt.Errorf("index is required")
	}
	if share == nil {
		return fmt.Errorf("share is required")
	}

	expected := group.NewPoint()
	expected.SetZero()

	xPower := big.NewInt(1)
	modulus := group.ScalarField()
	x := big.NewInt(int64(index))
	for i, encoded := range commitments {
		point, err := group.Decode(encoded)
		if err != nil {
			return fmt.Errorf("commitment %d: %w", i, err)
		}

		term := group.NewPoint()
		term.ScalarMult(point, xPower)
		expected.Add(expected, term)

		xPower.Mul(xPower, x)
		xPower.Mod(xPower, modulus)
	}

	actual := group.NewPoint()
	actual.ScalarBaseMult(share)
	if !actual.Equal(expected) {
		return fmt.Errorf("share does not match commitments")
	}
	return nil
}
