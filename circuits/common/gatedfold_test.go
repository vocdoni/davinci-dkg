package common

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
)

// Tiny stand-in for the contribution fold: a few "points" and a few
// "scalars", each gated by the same public prefix count, so the active
// words are a strict prefix of every region and the gated fold must equal
// plain BRLC over exactly those words.
const gatedFoldSlots = 4

type gatedFoldCircuit struct {
	Challenge     frontend.Variable `gnark:",public"`
	Count         frontend.Variable `gnark:",public"`
	ExpectedWords frontend.Variable `gnark:",public"`
	Expected      frontend.Variable `gnark:",public"`

	Points  [gatedFoldSlots][2]frontend.Variable
	Scalars [gatedFoldSlots]frontend.Variable
}

func (c *gatedFoldCircuit) Define(api frontend.API) error {
	mask := PrefixMask(api, c.Count, gatedFoldSlots)
	fold := NewGatedFold(api, c.Challenge)
	for i := range gatedFoldSlots {
		fold.Absorb(mask[i], c.Points[i][0], c.Points[i][1])
	}
	for i := range gatedFoldSlots {
		fold.Absorb(mask[i], c.Scalars[i])
	}
	api.AssertIsEqual(fold.Count(), c.ExpectedWords)
	api.AssertIsEqual(fold.Commitment(), c.Expected)
	return nil
}

// gatedFoldFixture lays the slot values out in traversal order with their
// gates, exactly as the circuit above absorbs them.
func gatedFoldFixture(count int) (points [gatedFoldSlots][2]*big.Int, scalars [gatedFoldSlots]*big.Int, words []*big.Int, gates []bool) {
	for i := range gatedFoldSlots {
		points[i][0] = big.NewInt(int64(1000 + 10*i))
		points[i][1] = big.NewInt(int64(2000 + 10*i))
		scalars[i] = big.NewInt(int64(3000 + i))
	}
	for i := range gatedFoldSlots {
		words = append(words, points[i][0], points[i][1])
		gates = append(gates, i < count, i < count)
	}
	for i := range gatedFoldSlots {
		words = append(words, scalars[i])
		gates = append(gates, i < count)
	}
	return points, scalars, words, gates
}

func activeWords(words []*big.Int, gates []bool) []*big.Int {
	var active []*big.Int
	for q, word := range words {
		if gates[q] {
			active = append(active, word)
		}
	}
	return active
}

// The reference fold applies the per-word rules literally; skipping a word
// must leave the exponent where it was, so the result is the plain BRLC of
// the active words and nothing else.
func TestFoldReferenceSkipsInactiveWords(t *testing.T) {
	c := qt.New(t)
	rho := big.NewInt(7919)

	for count := 0; count <= gatedFoldSlots; count++ {
		_, _, words, gates := gatedFoldFixture(count)
		acc, n, err := GatedBRLCNative(rho, words, gates)
		c.Assert(err, qt.IsNil)
		c.Assert(n, qt.Equals, 3*count)

		active := activeWords(words, gates)
		want, err := BRLCNative(rho, active...)
		c.Assert(err, qt.IsNil)
		c.Assert(acc.Cmp(want), qt.Equals, 0, qt.Commentf("count %d", count))

		if count > 0 && count < gatedFoldSlots {
			// The padded fold (inactive words zeroed, exponent still
			// advancing) is a different commitment: that is the v3.1 layout
			// the compact transcript no longer streams.
			padded := make([]*big.Int, len(words))
			for q := range words {
				padded[q] = big.NewInt(0)
				if gates[q] {
					padded[q] = words[q]
				}
			}
			zeroed, err := BRLCNative(rho, padded...)
			c.Assert(err, qt.IsNil)
			c.Assert(acc.Cmp(zeroed), qt.Not(qt.Equals), 0)
		}
	}

	// A gate with a nil value is an error, not a skip.
	_, _, words, gates := gatedFoldFixture(2)
	words[0] = nil
	_, _, err := GatedBRLCNative(rho, words, gates)
	c.Assert(err, qt.Not(qt.IsNil))
	_, _, err = GatedBRLCNative(rho, words[:1], gates)
	c.Assert(err, qt.Not(qt.IsNil))
}

// The gadget reproduces the reference for every prefix length, including the
// all-inactive and all-active corners, and refuses a commitment or a word
// count that belongs to another prefix.
func TestFoldGadgetMatchesReference(t *testing.T) {
	c := qt.New(t)
	rho := big.NewInt(104729)
	field := ecc.BN254.ScalarField()

	for count := 0; count <= gatedFoldSlots; count++ {
		points, scalars, words, gates := gatedFoldFixture(count)
		expected, n, err := GatedBRLCNative(rho, words, gates)
		c.Assert(err, qt.IsNil)

		assignment := &gatedFoldCircuit{
			Challenge:     rho,
			Count:         big.NewInt(int64(count)),
			ExpectedWords: big.NewInt(int64(n)),
			Expected:      expected,
		}
		for i := range gatedFoldSlots {
			assignment.Points[i][0] = points[i][0]
			assignment.Points[i][1] = points[i][1]
			assignment.Scalars[i] = scalars[i]
		}
		c.Assert(test.IsSolved(&gatedFoldCircuit{}, assignment, field), qt.IsNil, qt.Commentf("count %d", count))

		if count == 0 || count == gatedFoldSlots {
			continue
		}
		// Padded-layout commitment for the same words: must not solve.
		wrong := *assignment
		padded := make([]*big.Int, len(words))
		for q := range words {
			padded[q] = big.NewInt(0)
			if gates[q] {
				padded[q] = words[q]
			}
		}
		wrong.Expected, err = BRLCNative(rho, padded...)
		c.Assert(err, qt.IsNil)
		c.Assert(test.IsSolved(&gatedFoldCircuit{}, &wrong, field), qt.Not(qt.IsNil))

		// Claiming one more active word than the count admits: must not solve.
		wrong = *assignment
		wrong.ExpectedWords = big.NewInt(int64(n + 1))
		c.Assert(test.IsSolved(&gatedFoldCircuit{}, &wrong, field), qt.Not(qt.IsNil))
	}
}
