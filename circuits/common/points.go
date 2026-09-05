package common

import (
	"fmt"
	"math/big"

	ecc_tweds "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	format "github.com/vocdoni/gnark-crypto-primitives/ecc/format"
	"github.com/vocdoni/gnark-crypto-primitives/elgamal"
	circuitposeidon "github.com/vocdoni/gnark-crypto-primitives/hash/native/bn254/poseidon"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/types"
)

func BabyJubJubCurveID() ecc_tweds.ID {
	return ecc_tweds.BN254
}

func IdentityPoint() twistededwards.Point {
	return twistededwards.Point{X: big.NewInt(0), Y: big.NewInt(1)}
}

func CircuitPoint(point types.CurvePoint) twistededwards.Point {
	return twistededwards.Point{X: point.X, Y: point.Y}
}

func ToRTE(api frontend.API, point twistededwards.Point) twistededwards.Point {
	x, y := format.FromTEtoRTE(api, point.X, point.Y)
	return twistededwards.Point{X: x, Y: y}
}

func ToTE(api frontend.API, point twistededwards.Point) twistededwards.Point {
	x, y := format.FromRTEtoTE(api, point.X, point.Y)
	return twistededwards.Point{X: x, Y: y}
}

func CircuitPoints(points []types.CurvePoint, size int) ([]twistededwards.Point, error) {
	out := make([]twistededwards.Point, size)
	for i := range size {
		if i < len(points) {
			out[i] = CircuitPoint(points[i])
			continue
		}
		out[i] = IdentityPoint()
	}
	return out, nil
}

func FixedBaseMul(api frontend.API, scalar frontend.Variable) twistededwards.Point {
	return elgamal.FixedBaseScalarMulBN254(api, scalar)
}

func AssertPointEqual(api frontend.API, left, right twistededwards.Point) {
	api.AssertIsEqual(left.X, right.X)
	api.AssertIsEqual(left.Y, right.Y)
}

func SelectPoint(
	api frontend.API,
	enabled frontend.Variable,
	enabledPoint twistededwards.Point,
	disabledPoint twistededwards.Point,
) twistededwards.Point {
	return twistededwards.Point{
		X: api.Select(enabled, enabledPoint.X, disabledPoint.X),
		Y: api.Select(enabled, enabledPoint.Y, disabledPoint.Y),
	}
}

func AddPointIfEnabled(
	api frontend.API,
	acc twistededwards.Point,
	term twistededwards.Point,
	enabled frontend.Variable,
) twistededwards.Point {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		panic(err)
	}
	next := curve.Add(acc, term)
	return SelectPoint(api, enabled, next, acc)
}

func AssertPointOnCurve(api frontend.API, point twistededwards.Point) error {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return err
	}
	curve.AssertIsOnCurve(point)
	return nil
}

// ScalarMulVar multiplies a variable point by a full-width scalar with a
// plain double-and-add over the scalar's 254 bits: no hints, no
// decomposition trick, no commitments. gnark's hinted fake-GLV ScalarMul was
// unsound on cofactor curves before v0.15.0 (ePrint 2026/1776) and its fixed
// version commits, which changes the Solidity verifier and adds a pairing per
// proof; this gadget is the boring alternative for every secret or
// prover-chosen scalar that multiplies a variable-base point. ToBinary pins
// the canonical binary expansion, the unified twisted-Edwards formulas are
// complete, so the result is exactly [s]·P for any scalar s < 2^254.
func ScalarMulVar(api frontend.API, point twistededwards.Point, scalar frontend.Variable) twistededwards.Point {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		panic(err)
	}
	bits := api.ToBinary(scalar, 254)
	acc := IdentityPoint()
	for i := len(bits) - 1; i >= 0; i-- {
		acc = curve.Double(acc)
		sum := curve.Add(acc, point)
		acc = twistededwards.Point{
			X: api.Select(bits[i], sum.X, acc.X),
			Y: api.Select(bits[i], sum.Y, acc.Y),
		}
	}
	return acc
}

// ScalarMulSmallScalar computes `scalar · point` for a scalar that the
// caller has range-checked to fit in `nbBits` bits. 2-bit windowed
// left-to-right double-and-add over the explicitly-decomposed scalar,
// with `nbBits ≤ MaxN.BitLen()*N` typically much smaller than the
// BN254.Fr width — that's where the savings come from versus
// `curve.ScalarMul`, which goes through the half-GCD trick over the full
// ~252-bit scalar field even for small inputs.
//
// Bits are obtained via `api.ToBinary(scalar, nbBits)` which uses gnark's
// internal bit-decomposition hint AND emits the range-check constraints,
// so passing an oversized scalar fails the proof rather than producing a
// silently-wrong result.
//
// 2-bit windowing matches gnark's own `scalarMulGeneric` structure: per
// pair of bits we precompute (P, 2P, 3P) once, then for each window
// double twice and add the looked-up multiple. Saves ~5 constraints per
// pair vs naive 1-bit-at-a-time double-and-add, after amortising the
// (2P, 3P) precomputation across all windows.
//
// For `nbBits = 0` returns the identity (caller should special-case this).
func ScalarMulSmallScalar(
	api frontend.API,
	point twistededwards.Point,
	scalar frontend.Variable,
	nbBits int,
) twistededwards.Point {
	if nbBits <= 0 {
		return IdentityPoint()
	}
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		panic(err)
	}
	bits := api.ToBinary(scalar, nbBits)
	// 1-bit fallback: degenerate to a single conditional select.
	if nbBits == 1 {
		identity := IdentityPoint()
		return twistededwards.Point{
			X: api.Select(bits[0], point.X, identity.X),
			Y: api.Select(bits[0], point.Y, identity.Y),
		}
	}
	// Precompute the small multiple table {1·P, 2·P, 3·P}. The 0·P entry
	// is the identity (0, 1) and is handled by Lookup2 directly.
	twoP := curve.Double(point)
	threeP := curve.Add(twoP, point)

	// Top window. If nbBits is odd, the top window is just 1 bit and we
	// initialise with `bit_top ? P : identity`. Otherwise (even nbBits)
	// we initialise with the full 2-bit lookup over the top two bits.
	var res twistededwards.Point
	var startIdx int
	if nbBits%2 == 1 {
		identity := IdentityPoint()
		res = twistededwards.Point{
			X: api.Select(bits[nbBits-1], point.X, identity.X),
			Y: api.Select(bits[nbBits-1], point.Y, identity.Y),
		}
		startIdx = nbBits - 2
	} else {
		// 2-bit Lookup2 with identity at index (0,0).
		res = twistededwards.Point{
			X: api.Lookup2(bits[nbBits-1], bits[nbBits-2], 0, twoP.X, point.X, threeP.X),
			Y: api.Lookup2(bits[nbBits-1], bits[nbBits-2], 1, twoP.Y, point.Y, threeP.Y),
		}
		startIdx = nbBits - 3
	}
	for i := startIdx; i >= 1; i -= 2 {
		res = curve.Double(curve.Double(res))
		// Conditional add of {0, P, 2P, 3P} based on (bits[i], bits[i-1]).
		// (0, 0) → identity, just keep res; otherwise add.
		tmp := twistededwards.Point{
			X: api.Lookup2(bits[i], bits[i-1], 0, twoP.X, point.X, threeP.X),
			Y: api.Lookup2(bits[i], bits[i-1], 1, twoP.Y, point.Y, threeP.Y),
		}
		added := curve.Add(res, tmp)
		// When (bits[i], bits[i-1]) = (0, 0) the lookup returns the
		// identity, so `added` already equals `res`. Skip the Select.
		res.X = added.X
		res.Y = added.Y
	}
	return res
}

// IndexBits bounds a one-based committee index (≤ MaxN = 32) for the short
// scalar multiplications below; ScalarMulSmallScalar range-checks x to it.
const IndexBits = 6

// CommitmentPolynomialValue evaluates a commitment polynomial Σ_k cₖ·x^k by
// Horner's rule, acc ← x·acc + cₖ from the top coefficient down. Every step
// is one IndexBits-bit scalar multiplication plus one addition, independent
// of k (≈ 75 constraints), versus a growing-width multiplication per term.
// Inactive coefficient slots must already hold the identity point; callers
// pre-mask them once and reuse the vector across evaluations.
func CommitmentPolynomialValue(
	api frontend.API,
	commitments []twistededwards.Point,
	x frontend.Variable,
) (twistededwards.Point, error) {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return twistededwards.Point{}, err
	}
	acc := IdentityPoint()
	for k := len(commitments) - 1; k >= 0; k-- {
		if k != len(commitments)-1 {
			acc = ScalarMulSmallScalar(api, acc, x, IndexBits)
		}
		acc = curve.Add(acc, commitments[k])
	}
	return acc, nil
}

// MaskPoint folds a boolean mask into a point: the point itself when the
// slot is active, the identity (0, 1) otherwise.
func MaskPoint(api frontend.API, mask frontend.Variable, p twistededwards.Point) twistededwards.Point {
	return twistededwards.Point{X: api.Mul(mask, p.X), Y: api.Select(mask, p.Y, 1)}
}

// HashPoint hashes (state, point.X, point.Y) with Poseidon1, matching the
// same primitive used by all other in-circuit hashing in this package.
func HashPoint(api frontend.API, state frontend.Variable, point twistededwards.Point) (frontend.Variable, error) {
	return circuitposeidon.MultiHash(api, state, point.X, point.Y)
}

func ShareEncryptionDomain() *big.Int {
	return hash.DomainValue(hash.DomainShareEncryption)
}

func PartialDecryptDomain() *big.Int {
	return hash.DomainValue(hash.DomainPartialDecrypt)
}

func HashPointTuple(api frontend.API, state frontend.Variable, points ...twistededwards.Point) (frontend.Variable, error) {
	current := state
	var err error
	for _, point := range points {
		current, err = HashPoint(api, current, point)
		if err != nil {
			return 0, err
		}
	}
	return current, nil
}

func HashPointTupleNative(state *big.Int, points ...types.CurvePoint) (*big.Int, error) {
	current := new(big.Int).Set(state)
	var err error
	for _, point := range points {
		current, err = hash.HashFieldElements(current, point.X, point.Y)
		if err != nil {
			return nil, err
		}
	}
	return current, nil
}

// IdentityCurvePoint is the twisted-Edwards identity (0, 1) in native form:
// the value every masked-out point slot carries in a witness or transcript.
func IdentityCurvePoint() types.CurvePoint {
	return types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
}

// PadPoints right-pads a native point vector with the identity until it
// reaches size, the point analogue of PadBigInts.
func PadPoints(points []types.CurvePoint, size int) ([]types.CurvePoint, error) {
	if len(points) > size {
		return nil, fmt.Errorf("got %d points, max is %d", len(points), size)
	}
	out := make([]types.CurvePoint, size)
	for i := range size {
		if i < len(points) && points[i].X != nil && points[i].Y != nil {
			out[i] = types.CurvePoint{X: new(big.Int).Set(points[i].X), Y: new(big.Int).Set(points[i].Y)}
			continue
		}
		out[i] = IdentityCurvePoint()
	}
	return out, nil
}

// CommitmentPolynomialValueNative mirrors CommitmentPolynomialValue for
// witness builders: Σ_m x^m · commitments[m] over the BabyJubJub prime-order
// subgroup. Identity slots contribute nothing, so callers pass the same
// masked vector the circuit uses. Powers are reduced modulo the subgroup
// order, which is a no-op at the current bounds (x ≤ MaxN, m < MaxN).
func CommitmentPolynomialValueNative(commitments []types.CurvePoint, x *big.Int) (types.CurvePoint, error) {
	if x == nil {
		return types.CurvePoint{}, fmt.Errorf("x is required")
	}
	modulus := group.ScalarField()
	acc := group.NewPoint()
	acc.SetZero()
	power := big.NewInt(1)
	for m, commitment := range commitments {
		point, err := group.Decode(commitment)
		if err != nil {
			return types.CurvePoint{}, fmt.Errorf("decode commitment %d: %w", m, err)
		}
		term := group.NewPoint()
		term.ScalarMult(point, power)
		acc.Add(acc, term)
		power.Mul(power, x)
		power.Mod(power, modulus)
	}
	return group.Encode(acc), nil
}
