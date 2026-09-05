package common

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/test"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

type scalarMulVarCircuit struct {
	P, Q   twistededwards.Point
	Scalar frontend.Variable
}

func (c *scalarMulVarCircuit) Define(api frontend.API) error {
	AssertPointEqual(api, ScalarMulVar(api, c.P, c.Scalar), c.Q)
	return nil
}

// The hint-free gadget must agree with the native library for random
// full-width scalars (including ones at or above the subgroup order) and
// must refuse a wrong product.
func TestScalarMulVarMatchesNative(t *testing.T) {
	base := group.NewPoint()
	base.ScalarBaseMult(big.NewInt(12345))
	for _, s := range []*big.Int{big.NewInt(0), big.NewInt(1), new(big.Int).Sub(group.ScalarField(), big.NewInt(1)), group.ScalarField(), randomScalar(t)} {
		q := group.NewPoint()
		q.ScalarMult(base, s)
		p, qq := group.Encode(base), group.Encode(q)
		good := &scalarMulVarCircuit{P: twistededwards.Point{X: p.X, Y: p.Y}, Q: twistededwards.Point{X: qq.X, Y: qq.Y}, Scalar: s}
		if err := test.IsSolved(&scalarMulVarCircuit{}, good, ecc.BN254.ScalarField()); err != nil {
			t.Fatalf("scalar %s: %v", s, err)
		}
		bad := *good
		bad.Q = twistededwards.Point{X: p.X, Y: p.Y}
		if s.Cmp(big.NewInt(1)) != 0 && test.IsSolved(&scalarMulVarCircuit{}, &bad, ecc.BN254.ScalarField()) == nil {
			t.Fatalf("scalar %s: wrong product accepted", s)
		}
	}
}

func randomScalar(t *testing.T) *big.Int {
	s, err := rand.Int(rand.Reader, ecc.BN254.ScalarField())
	if err != nil {
		t.Fatal(err)
	}
	return s
}
