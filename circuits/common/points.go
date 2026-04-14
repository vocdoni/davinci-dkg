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

// CommitmentPolynomialValue evaluates a commitment polynomial Σ_k cₖ·x^k.
//
// If `mask` is non-nil, slot k is included only when mask[k] == 1; this matches
// the legacy callsite. If `mask` is nil, callers are expected to have already
// folded the mask into the commitments (e.g. by replacing inactive slots with
// the curve identity point), which lets the inner loop skip the per-iteration
// Select on the running sum and saves ~2 constraints per coefficient per call.
func CommitmentPolynomialValue(
	api frontend.API,
	commitments []twistededwards.Point,
	mask []frontend.Variable,
	x frontend.Variable,
) (twistededwards.Point, error) {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return twistededwards.Point{}, err
	}
	sum := IdentityPoint()
	power := frontend.Variable(1)
	for i, commitment := range commitments {
		scaled := curve.ScalarMul(commitment, power)
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
