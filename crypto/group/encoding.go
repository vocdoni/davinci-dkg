package group

import (
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-node/crypto/ecc"
)

// ScalarField returns the subgroup order used by the DKG group.
func ScalarField() *big.Int {
	return NewPoint().Order()
}

// Encode converts a curve point into the explicit affine representation used by domain types.
func Encode(point ecc.Point) types.CurvePoint {
	x, y := point.Point()
	return types.CurvePoint{
		X: new(big.Int).Set(x),
		Y: new(big.Int).Set(y),
	}
}

// Decode converts an explicit affine point into the local curve wrapper type.
func Decode(encoded types.CurvePoint) (ecc.Point, error) {
	if err := encoded.Validate(); err != nil {
		return nil, err
	}
	point := NewPoint()
	decoded := point.SetPoint(encoded.X, encoded.Y)
	if decoded == nil {
		return nil, fmt.Errorf("failed to decode point")
	}
	return decoded, nil
}
