package common

import (
	"fmt"
	"math/big"
	"sync"

	edbn254 "github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	ecc_tweds "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/std/math/bits"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// ScalarBits is the bit length of the BabyJubJub subgroup order r: 251.
// Every canonical scalar fits it, so a canonical decomposition needs no
// more bits than that.
var ScalarBits = subgroupOrder.BitLen()

// CanonicalScalarBits range-checks v < r and returns its ScalarBits
// little-endian bits, so a caller can feed the same decomposition to
// FixedBaseMulBits instead of decomposing v a second time. About 500
// constraints for the check and the bits together, against 672 for
// AssertIsLessOrEqual alone.
func CanonicalScalarBits(api frontend.API, v frontend.Variable) []frontend.Variable {
	bs := bits.ToBinary(api, v, bits.WithNbDigits(ScalarBits))
	AssertBitsLessOrEqualConst(api, bs, subgroupOrderMinusOne)
	return bs
}

// AssertBitsLessOrEqualConst asserts that the little-endian boolean bits bs
// encode a value at most bound. It scans from the top bit keeping `eq`, one
// while every higher bit equals the bound's: a set bit where the bound has
// zero is only allowed once `eq` is zero. One constraint per bit.
func AssertBitsLessOrEqualConst(api frontend.API, bs []frontend.Variable, bound *big.Int) {
	var eq frontend.Variable = 1
	for k := len(bs) - 1; k >= 0; k-- {
		if bound.Bit(k) == 1 {
			eq = api.Mul(eq, bs[k])
		} else {
			api.AssertIsEqual(api.Mul(eq, bs[k]), 0)
		}
	}
}

// fixedBaseWindows[w][k] = k·2^(2w)·G for k in 0..3: the constants of one
// 2-bit window of the fixed-base multiplication.
var (
	fixedBaseWindows     [127][4][2]*big.Int
	fixedBaseWindowsOnce sync.Once
)

func initFixedBaseWindows() {
	params := edbn254.GetEdwardsCurve()
	for w := range fixedBaseWindows {
		fixedBaseWindows[w][0] = [2]*big.Int{big.NewInt(0), big.NewInt(1)}
		for k := 1; k < 4; k++ {
			s := new(big.Int).Lsh(big.NewInt(int64(k)), uint(2*w))
			var p edbn254.PointAffine
			p.ScalarMultiplication(&params.Base, s)
			x, y := new(big.Int), new(big.Int)
			p.X.BigInt(x)
			p.Y.BigInt(y)
			fixedBaseWindows[w][k] = [2]*big.Int{x, y}
		}
	}
}

// FixedBaseMulBits returns [s]·G for the scalar whose little-endian boolean
// bits are bs (at most 254 of them). Each 2-bit window selects one of four
// constants with the bilinear form
//
//	c00 + b0·(c10 − c00) + b1·(c01 − c00) + b0·b1·(c11 − c10 − c01 + c00),
//
// one multiplication (b0·b1) shared by both coordinates, then adds it to
// the accumulator with the complete twisted-Edwards formula (adding the
// identity for a zero window is correct and needs no select). 1,393
// constraints for 254 bits against 2,342 for the 4-bit nested-Lookup2
// gadget it replaces. Callers must pass constrained booleans: ToBinary or
// CanonicalScalarBits.
func FixedBaseMulBits(api frontend.API, bs []frontend.Variable) twistededwards.Point {
	fixedBaseWindowsOnce.Do(initFixedBaseWindows)
	if len(bs) > 2*len(fixedBaseWindows) {
		panic(fmt.Sprintf("FixedBaseMulBits: %d bits exceed the %d-bit table", len(bs), 2*len(fixedBaseWindows)))
	}
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		panic(err)
	}
	var acc twistededwards.Point
	for w := 0; 2*w < len(bs); w++ {
		b0 := bs[2*w]
		var b1 frontend.Variable = 0
		if 2*w+1 < len(bs) {
			b1 = bs[2*w+1]
		}
		t := fixedBaseWindows[w]
		b01 := api.Mul(b0, b1)
		coord := func(c int) frontend.Variable {
			c00, c10, c01, c11 := t[0][c], t[1][c], t[2][c], t[3][c]
			d10 := new(big.Int).Sub(c10, c00)
			d01 := new(big.Int).Sub(c01, c00)
			d11 := new(big.Int).Sub(new(big.Int).Sub(new(big.Int).Add(c11, c00), c10), c01)
			return api.Add(c00, api.Mul(b0, d10), api.Mul(b1, d01), api.Mul(b01, d11))
		}
		window := twistededwards.Point{X: coord(0), Y: coord(1)}
		if w == 0 {
			acc = window
			continue
		}
		acc = curve.Add(acc, window)
	}
	return acc
}

// ScalarMulVarBits multiplies a variable point by the scalar whose
// little-endian boolean bits are bs, with a plain double-and-add: no hints,
// complete formulas, exactly [s]·P (see ScalarMulVar for why nothing
// cleverer is used). Callers that already decomposed the scalar, e.g. for
// FixedBaseMulBits, pass the same bits and save a decomposition.
func ScalarMulVarBits(api frontend.API, point twistededwards.Point, bs []frontend.Variable) twistededwards.Point {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		panic(err)
	}
	acc := IdentityPoint()
	for i := len(bs) - 1; i >= 0; i-- {
		acc = curve.Double(acc)
		sum := curve.Add(acc, point)
		acc = twistededwards.Point{
			X: api.Select(bs[i], sum.X, acc.X),
			Y: api.Select(bs[i], sum.Y, acc.Y),
		}
	}
	return acc
}

// AssertCofactorPreimage asserts that point lies in the prime-order subgroup
// by exhibiting preimage with [8]·preimage = point: the curve has order 8·r,
// so the image of multiplication by 8 is exactly ⟨G⟩. The preimage is
// checked on the curve and multiplied by three constrained doublings; no
// hint is trusted. 21 constraints. An honest prover uses
// CofactorPreimageNative.
func AssertCofactorPreimage(api frontend.API, preimage, point twistededwards.Point) error {
	if err := AssertPointOnCurve(api, preimage); err != nil {
		return err
	}
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return err
	}
	q := curve.Double(curve.Double(curve.Double(preimage)))
	api.AssertIsEqual(q.X, point.X)
	api.AssertIsEqual(q.Y, point.Y)
	return nil
}

// cofactorInverse is 8⁻¹ mod r.
var cofactorInverse = new(big.Int).ModInverse(big.NewInt(8), subgroupOrder)

// CofactorPreimageNative returns [8⁻¹ mod r]·point, the witness
// AssertCofactorPreimage expects for a prime-subgroup point.
func CofactorPreimageNative(point types.CurvePoint) (types.CurvePoint, error) {
	p, err := group.Decode(point)
	if err != nil {
		return types.CurvePoint{}, err
	}
	q := group.NewPoint()
	q.ScalarMult(p, cofactorInverse)
	return group.Encode(q), nil
}
