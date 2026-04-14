package shamir

import (
	"fmt"
	"math/big"
)

// Share is one Shamir evaluation f(i).
type Share struct {
	Index uint16
	Value *big.Int
}

// InterpolateConstant reconstructs f(0) from a set of Shamir shares.
func InterpolateConstant(shares []Share, modulus *big.Int) (*big.Int, error) {
	if modulus == nil || modulus.Sign() <= 0 {
		return nil, fmt.Errorf("modulus must be positive")
	}
	if len(shares) == 0 {
		return nil, fmt.Errorf("at least one share is required")
	}

	seen := make(map[uint16]struct{}, len(shares))
	secret := big.NewInt(0)
	for i, share := range shares {
		if share.Index == 0 {
			return nil, fmt.Errorf("share %d: index is required", i)
		}
		if share.Value == nil {
			return nil, fmt.Errorf("share %d: value is required", i)
		}
		if _, ok := seen[share.Index]; ok {
			return nil, fmt.Errorf("share %d: duplicate index", i)
		}
		seen[share.Index] = struct{}{}

		coefficient, err := lagrangeCoefficientAtZero(share.Index, shares, modulus)
		if err != nil {
			return nil, err
		}

		term := new(big.Int).Mul(share.Value, coefficient)
		term.Mod(term, modulus)
		secret.Add(secret, term)
		secret.Mod(secret, modulus)
	}
	return secret, nil
}

func lagrangeCoefficientAtZero(index uint16, shares []Share, modulus *big.Int) (*big.Int, error) {
	numerator := big.NewInt(1)
	denominator := big.NewInt(1)
	xi := big.NewInt(int64(index))

	for _, other := range shares {
		if other.Index == index {
			continue
		}

		xj := big.NewInt(int64(other.Index))
		numerator.Mul(numerator, new(big.Int).Neg(xj))
		numerator.Mod(numerator, modulus)

		diff := new(big.Int).Sub(xi, xj)
		diff.Mod(diff, modulus)
		denominator.Mul(denominator, diff)
		denominator.Mod(denominator, modulus)
	}

	inverse := new(big.Int).ModInverse(denominator, modulus)
	if inverse == nil {
		return nil, fmt.Errorf("share index %d: denominator not invertible", index)
	}
	coefficient := new(big.Int).Mul(numerator, inverse)
	coefficient.Mod(coefficient, modulus)
	return coefficient, nil
}
