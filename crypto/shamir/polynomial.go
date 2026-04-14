package shamir

import (
	"fmt"
	"math/big"
)

// Polynomial is a Shamir polynomial over a finite field.
type Polynomial struct {
	Coefficients []*big.Int
	Modulus      *big.Int
}

// NewPolynomial constructs a polynomial with coefficients reduced modulo the field.
func NewPolynomial(coefficients []*big.Int, modulus *big.Int) (*Polynomial, error) {
	if modulus == nil || modulus.Sign() <= 0 {
		return nil, fmt.Errorf("modulus must be positive")
	}
	if len(coefficients) == 0 {
		return nil, fmt.Errorf("at least one coefficient is required")
	}
	normalized := make([]*big.Int, len(coefficients))
	for i, coeff := range coefficients {
		if coeff == nil {
			return nil, fmt.Errorf("coefficient %d is nil", i)
		}
		normalized[i] = new(big.Int).Mod(new(big.Int).Set(coeff), modulus)
	}
	return &Polynomial{
		Coefficients: normalized,
		Modulus:      new(big.Int).Set(modulus),
	}, nil
}

// Evaluate returns f(x) mod modulus.
func (p *Polynomial) Evaluate(x *big.Int) *big.Int {
	result := big.NewInt(0)
	xPower := big.NewInt(1)
	for _, coeff := range p.Coefficients {
		term := new(big.Int).Mul(coeff, xPower)
		term.Mod(term, p.Modulus)
		result.Add(result, term)
		result.Mod(result, p.Modulus)

		xPower.Mul(xPower, x)
		xPower.Mod(xPower, p.Modulus)
	}
	return result
}

// ValidateDegree checks that the polynomial degree is below the threshold.
func (p *Polynomial) ValidateDegree(threshold uint16) error {
	if threshold == 0 {
		return fmt.Errorf("threshold must be non-zero")
	}
	if len(p.Coefficients) > int(threshold) {
		return fmt.Errorf("polynomial degree exceeds threshold")
	}
	return nil
}
