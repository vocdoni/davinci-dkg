package common

import (
	"math/big"

	ecc_tweds "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	format "github.com/vocdoni/gnark-crypto-primitives/ecc/format"
	"github.com/vocdoni/gnark-crypto-primitives/elgamal"
	circuitposeidon "github.com/vocdoni/gnark-crypto-primitives/hash/native/bn254/poseidon"

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

// CommitmentPolynomialValue evaluates a commitment polynomial Σ_k cₖ·x^k.
//
// If `mask` is non-nil, slot k is included only when mask[k] == 1; this matches
// the legacy callsite. If `mask` is nil, callers are expected to have already
// folded the mask into the commitments (e.g. by replacing inactive slots with
// the curve identity point), which lets the inner loop skip the per-iteration
// Select on the running sum and saves ~2 constraints per coefficient per call.
//
// `xMaxBits` is the caller's `⌈log₂(MaxN)⌉` bound on the recipient /
// participant index — i.e. the per-step bit growth of `power_k = x^k`
// when x is range-checked to ≤ MaxN. For one-based indexes the boundary
// case is `x = MaxN = 2^xMaxBits` where `x^k = 2^(xMaxBits·k)` needs
// `xMaxBits·k + 1` bits to encode (the leading bit is set).  We pass
// that bound to ScalarMulSmallScalar; it via api.ToBinary uses gnark's
// internal bit-decomposition hint and emits the matching range-check.
// Saves ~1.3M constraints in the contribution circuit at MaxN = 32 vs
// the full-width scalar mul. `xMaxBits = 0` falls back to the original
// full-width path.
func CommitmentPolynomialValue(
	api frontend.API,
	commitments []twistededwards.Point,
	mask []frontend.Variable,
	x frontend.Variable,
	xMaxBits int,
) (twistededwards.Point, error) {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return twistededwards.Point{}, err
	}
	sum := IdentityPoint()
	power := frontend.Variable(1)
	for i, commitment := range commitments {
		var scaled twistededwards.Point
		switch {
		case i == 0:
			// power = 1, scaled = commitment[0]. No mul, no doublings.
			scaled = commitment
		case xMaxBits > 0:
			// power < x^i ≤ MaxN^i = 2^(xMaxBits·i). Worst case x = MaxN
			// gives power = 2^(xMaxBits·i), needing xMaxBits·i + 1 bits.
			scaled = ScalarMulSmallScalar(api, commitment, power, xMaxBits*i+1)
		default:
			scaled = curve.ScalarMul(commitment, power)
		}
		next := curve.Add(sum, scaled)
		if mask == nil {
			sum.X = next.X
			sum.Y = next.Y
		} else {
			active := frontend.Variable(1)
			if len(mask) > i {
				active = mask[i]
			}
			sum.X = api.Select(active, next.X, sum.X)
			sum.Y = api.Select(active, next.Y, sum.Y)
		}
		power = api.Mul(power, x)
	}
	return sum, nil
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
