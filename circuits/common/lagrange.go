package common

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// LagrangeCoefficientAtZero computes one Lagrange basis coefficient evaluated at 0.
func LagrangeCoefficientAtZero(api frontend.API, xs, mask []frontend.Variable, idx int) frontend.Variable {
	coeff := frontend.Variable(1)
	xi := xs[idx]
	for j := range len(xs) {
		if j == idx {
			continue
		}
		numerator := api.Select(mask[j], xs[j], 1)
		denominator := api.Select(mask[j], api.Sub(xs[j], xi), 1)
		coeff = api.Mul(coeff, api.Div(numerator, denominator))
	}
	return api.Select(mask[idx], coeff, 0)
}

// InterpolateAtZero reconstructs the constant term from masked shares in-circuit.
func InterpolateAtZero(api frontend.API, xs, ys, mask []frontend.Variable) frontend.Variable {
	result := frontend.Variable(0)
	for i := range len(xs) {
		coefficient := LagrangeCoefficientAtZero(api, xs, mask, i)
		term := api.Mul(coefficient, api.Select(mask[i], ys[i], 0))
		result = api.Add(result, term)
	}
	return result
}

// LagrangeCoefficientsAtZeroNative computes native Lagrange coefficients evaluated at 0.
// The coefficients are computed modulo the BJJ group order (r_bjj), which is the correct
// field for scalar multiplication on the BabyJubJub curve.
func LagrangeCoefficientsAtZeroNative(indexes []*big.Int) ([]*big.Int, error) {
	modulus := group.ScalarField()
	coeffs := make([]*big.Int, len(indexes))
	for i, xi := range indexes {
		if xi == nil {
			return nil, fmt.Errorf("index %d is nil", i)
		}
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)
		for j, xj := range indexes {
			if i == j {
				continue
			}
			if xj == nil {
				return nil, fmt.Errorf("index %d is nil", j)
			}
			numerator.Mul(numerator, xj)
			numerator.Mod(numerator, modulus)

			diff := new(big.Int).Sub(xj, xi)
			diff.Mod(diff, modulus)
			denominator.Mul(denominator, diff)
			denominator.Mod(denominator, modulus)
		}
		denominatorInv := new(big.Int).ModInverse(denominator, modulus)
		if denominatorInv == nil {
			return nil, fmt.Errorf("no inverse for denominator at index %d", i)
		}
		coeff := new(big.Int).Mul(numerator, denominatorInv)
		coeff.Mod(coeff, modulus)
		coeffs[i] = coeff
	}
	return coeffs, nil
}

// InterpolateAtZeroNative reconstructs the constant term from native shares.
func InterpolateAtZeroNative(indexes, values []*big.Int) (*big.Int, error) {
	if len(indexes) == 0 || len(indexes) != len(values) {
		return nil, fmt.Errorf("indexes and values must have the same non-zero length")
	}
	coeffs, err := LagrangeCoefficientsAtZeroNative(indexes)
	if err != nil {
		return nil, err
	}
	modulus := group.ScalarField()
	result := big.NewInt(0)
	for i, value := range values {
		if value == nil {
			return nil, fmt.Errorf("value %d is nil", i)
		}
		term := new(big.Int).Mul(coeffs[i], value)
		term.Mod(term, modulus)
		result.Add(result, term)
		result.Mod(result, modulus)
	}
	return result, nil
}

// InterpolatePointsAtZeroNative reconstructs the constant point from native partial decryptions.
func InterpolatePointsAtZeroNative(indexes []*big.Int, values []types.CurvePoint) (types.CurvePoint, error) {
	if len(indexes) == 0 || len(indexes) != len(values) {
		return types.CurvePoint{}, fmt.Errorf("indexes and values must have the same non-zero length")
	}
	coeffs, err := LagrangeCoefficientsAtZeroNative(indexes)
	if err != nil {
		return types.CurvePoint{}, err
	}
	sum := group.NewPoint()
	sum.SetZero()
	for i, value := range values {
		point, err := group.Decode(value)
		if err != nil {
			return types.CurvePoint{}, fmt.Errorf("decode point %d: %w", i, err)
		}
		term := group.NewPoint()
		term.ScalarMult(point, coeffs[i])
		sum.Add(sum, term)
	}
	return group.Encode(sum), nil
}
