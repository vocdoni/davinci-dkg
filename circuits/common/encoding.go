package common

import (
	"fmt"
	"math/big"
)

// PadBigInts right-pads a slice with zeros until it reaches size.
func PadBigInts(values []*big.Int, size int) ([]*big.Int, error) {
	if len(values) > size {
		return nil, fmt.Errorf("got %d values, max is %d", len(values), size)
	}
	out := make([]*big.Int, size)
	for i := range size {
		if i < len(values) && values[i] != nil {
			out[i] = new(big.Int).Set(values[i])
			continue
		}
		out[i] = big.NewInt(0)
	}
	return out, nil
}

// Uint16sToBigInts converts uint16 inputs to big.Int values.
func Uint16sToBigInts(values []uint16) []*big.Int {
	out := make([]*big.Int, len(values))
	for i, value := range values {
		out[i] = big.NewInt(int64(value))
	}
	return out
}
