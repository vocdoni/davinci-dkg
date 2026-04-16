package common

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/shamir"
)

// PrefixMask returns a latch-style active mask for the first count elements.
func PrefixMask(api frontend.API, count frontend.Variable, size int) []frontend.Variable {
	mask := make([]frontend.Variable, size)
	isActive := api.Sub(1, api.IsZero(count))
	for i := range size {
		mask[i] = isActive
		isEnd := api.IsZero(api.Sub(count, i+1))
		isActive = api.Mul(isActive, api.Sub(1, isEnd))
	}
	return mask
}

// EvaluatePolynomial evaluates a scalar polynomial in-circuit using Horner form.
//
// If `mask` is non-nil, slot k is zeroed when mask[k] == 0. If `mask` is nil,
// callers are expected to have pre-masked the coefficients themselves, which
// saves one Select per coefficient per call — important when the same set of
// masked coefficients is reused across many evaluations.
func EvaluatePolynomial(api frontend.API, coeffs, mask []frontend.Variable, x frontend.Variable) frontend.Variable {
	result := frontend.Variable(0)
	for i := len(coeffs) - 1; i >= 0; i-- {
		coeff := coeffs[i]
		if mask != nil && len(mask) > i {
			coeff = api.Select(mask[i], coeffs[i], 0)
		}
		result = api.Add(api.Mul(result, x), coeff)
	}
	return result
}

// EvaluatePolynomialNative evaluates a Shamir polynomial over the BabyJubJub subgroup order.
// Shares must live in the same field as the BabyJubJub scalar field so that
// AddModSubgroupOrder in the contribution circuit can bound the carry to {0,1}.
func EvaluatePolynomialNative(coefficients []*big.Int, x *big.Int) (*big.Int, error) {
	if x == nil {
		return nil, fmt.Errorf("x is required")
	}
	poly, err := shamir.NewPolynomial(coefficients, group.ScalarField())
	if err != nil {
		return nil, err
	}
	return poly.Evaluate(x), nil
}
