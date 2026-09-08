package common

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	edbn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/types"
)

func nativeMul(scalar *big.Int, base *edbn254.PointAffine) (*big.Int, *big.Int) {
	var p edbn254.PointAffine
	p.ScalarMultiplication(base, scalar)
	x, y := new(big.Int), new(big.Int)
	p.X.BigInt(x)
	p.Y.BigInt(y)
	return x, y
}

type canonicalFixedBaseCircuit struct {
	S    frontend.Variable
	X, Y frontend.Variable `gnark:",public"`
}

func (c *canonicalFixedBaseCircuit) Define(api frontend.API) error {
	p := FixedBaseMulBits(api, CanonicalScalarBits(api, c.S))
	api.AssertIsEqual(p.X, c.X)
	api.AssertIsEqual(p.Y, c.Y)
	return nil
}

type fixedBaseCircuit struct {
	S    frontend.Variable
	X, Y frontend.Variable `gnark:",public"`
}

func (c *fixedBaseCircuit) Define(api frontend.API) error {
	p := FixedBaseMul(api, c.S)
	api.AssertIsEqual(p.X, c.X)
	api.AssertIsEqual(p.Y, c.Y)
	return nil
}

func TestFixedBaseMulBitsMatchesNative(t *testing.T) {
	c := qt.New(t)
	params := edbn254.GetEdwardsCurve()
	rMinusOne := new(big.Int).Sub(&params.Order, big.NewInt(1))
	scalars := []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2), big.NewInt(3), rMinusOne}
	for range 3 {
		s, err := rand.Int(rand.Reader, &params.Order)
		c.Assert(err, qt.IsNil)
		scalars = append(scalars, s)
	}
	for _, s := range scalars {
		x, y := nativeMul(s, &params.Base)
		c.Assert(test.IsSolved(&canonicalFixedBaseCircuit{}, &canonicalFixedBaseCircuit{S: s, X: x, Y: y}, ecc.BN254.ScalarField()), qt.IsNil)
		c.Assert(test.IsSolved(&fixedBaseCircuit{}, &fixedBaseCircuit{S: s, X: x, Y: y}, ecc.BN254.ScalarField()), qt.IsNil)
		wrong := new(big.Int).Add(x, big.NewInt(1))
		c.Assert(test.IsSolved(&canonicalFixedBaseCircuit{}, &canonicalFixedBaseCircuit{S: s, X: wrong, Y: y}, ecc.BN254.ScalarField()), qt.IsNotNil)
	}
	// FixedBaseMul accepts any field element (all below 2^254); the window
	// sum equals [s mod r]·G (the native routine expects a reduced scalar).
	pMinusOne := new(big.Int).Sub(ecc.BN254.ScalarField(), big.NewInt(1))
	x, y := nativeMul(new(big.Int).Mod(pMinusOne, &params.Order), &params.Base)
	c.Assert(test.IsSolved(&fixedBaseCircuit{}, &fixedBaseCircuit{S: pMinusOne, X: x, Y: y}, ecc.BN254.ScalarField()), qt.IsNil)
}

func TestCanonicalScalarBitsRejectsNonCanonical(t *testing.T) {
	c := qt.New(t)
	params := edbn254.GetEdwardsCurve()
	for _, bad := range []*big.Int{
		new(big.Int).Set(&params.Order),
		new(big.Int).Add(&params.Order, big.NewInt(1)),
		new(big.Int).Mul(big.NewInt(7), &params.Order),
		new(big.Int).Sub(ecc.BN254.ScalarField(), big.NewInt(1)),
	} {
		x, y := nativeMul(bad, &params.Base)
		c.Assert(test.IsSolved(&canonicalFixedBaseCircuit{}, &canonicalFixedBaseCircuit{S: bad, X: x, Y: y}, ecc.BN254.ScalarField()),
			qt.IsNotNil, qt.Commentf("%s must be rejected", bad))
	}
}

type varMulBitsCircuit struct {
	P    twistededwards.Point
	S    frontend.Variable
	X, Y frontend.Variable `gnark:",public"`
}

func (c *varMulBitsCircuit) Define(api frontend.API) error {
	bs := api.ToBinary(c.S, 254)
	p := ScalarMulVarBits(api, c.P, bs)
	api.AssertIsEqual(p.X, c.X)
	api.AssertIsEqual(p.Y, c.Y)
	return nil
}

func TestScalarMulVarBitsMatchesNative(t *testing.T) {
	c := qt.New(t)
	params := edbn254.GetEdwardsCurve()
	px, py := nativeMul(big.NewInt(12345), &params.Base)
	var base edbn254.PointAffine
	base.X.SetBigInt(px)
	base.Y.SetBigInt(py)
	for _, s := range []*big.Int{big.NewInt(0), big.NewInt(1), new(big.Int).Sub(&params.Order, big.NewInt(1))} {
		x, y := nativeMul(s, &base)
		w := &varMulBitsCircuit{P: twistededwards.Point{X: px, Y: py}, S: s, X: x, Y: y}
		c.Assert(test.IsSolved(&varMulBitsCircuit{}, w, ecc.BN254.ScalarField()), qt.IsNil)
	}
}

type preimageCircuit struct {
	Q, C twistededwards.Point
}

func (c *preimageCircuit) Define(api frontend.API) error {
	return AssertCofactorPreimage(api, c.Q, c.C)
}

func TestCofactorPreimageCertifiesTheSubgroup(t *testing.T) {
	c := qt.New(t)
	params := edbn254.GetEdwardsCurve()
	cx, cy := nativeMul(big.NewInt(987654321), &params.Base)
	point := types.CurvePoint{X: cx, Y: cy}
	pre, err := CofactorPreimageNative(point)
	c.Assert(err, qt.IsNil)
	ok := &preimageCircuit{Q: CircuitPoint(pre), C: twistededwards.Point{X: cx, Y: cy}}
	c.Assert(test.IsSolved(&preimageCircuit{}, ok, ecc.BN254.ScalarField()), qt.IsNil)

	// The identity certifies itself.
	id := &preimageCircuit{Q: IdentityPoint(), C: IdentityPoint()}
	c.Assert(test.IsSolved(&preimageCircuit{}, id, ecc.BN254.ScalarField()), qt.IsNil)

	// A wrong preimage, an off-curve preimage, and a point outside the
	// subgroup (the order-2 point (0, p−1), which has no preimage at all).
	wrong := &preimageCircuit{Q: twistededwards.Point{X: cx, Y: cy}, C: twistededwards.Point{X: cx, Y: cy}}
	c.Assert(test.IsSolved(&preimageCircuit{}, wrong, ecc.BN254.ScalarField()), qt.IsNotNil)
	off := &preimageCircuit{Q: twistededwards.Point{X: 1, Y: 1}, C: twistededwards.Point{X: cx, Y: cy}}
	c.Assert(test.IsSolved(&preimageCircuit{}, off, ecc.BN254.ScalarField()), qt.IsNotNil)
	minusOne := new(big.Int).Sub(ecc.BN254.ScalarField(), big.NewInt(1))
	torsion := twistededwards.Point{X: 0, Y: minusOne}
	for _, q := range []twistededwards.Point{IdentityPoint(), torsion, CircuitPoint(pre)} {
		c.Assert(test.IsSolved(&preimageCircuit{}, &preimageCircuit{Q: q, C: torsion}, ecc.BN254.ScalarField()), qt.IsNotNil)
	}
}

type constMulCircuit struct {
	P    twistededwards.Point
	K    int
	X, Y frontend.Variable `gnark:",public"`
}

func (c *constMulCircuit) Define(api frontend.API) error {
	p := ScalarMulSmallScalar(api, c.P, c.K, IndexBits)
	api.AssertIsEqual(p.X, c.X)
	api.AssertIsEqual(p.Y, c.Y)
	return nil
}

func TestScalarMulSmallScalarConstantChains(t *testing.T) {
	c := qt.New(t)
	params := edbn254.GetEdwardsCurve()
	px, py := nativeMul(big.NewInt(4242), &params.Base)
	var base edbn254.PointAffine
	base.X.SetBigInt(px)
	base.Y.SetBigInt(py)
	for k := 0; k <= MaxN; k++ {
		x, y := nativeMul(big.NewInt(int64(k)), &base)
		w := &constMulCircuit{P: twistededwards.Point{X: px, Y: py}, K: k, X: x, Y: y}
		c.Assert(test.IsSolved(&constMulCircuit{K: k}, w, ecc.BN254.ScalarField()), qt.IsNil, qt.Commentf("k=%d", k))
	}
}
