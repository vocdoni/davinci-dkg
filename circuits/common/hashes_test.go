package common

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/types"
)

type hashFieldElementsCircuit struct {
	A        frontend.Variable `gnark:",public"`
	B        frontend.Variable `gnark:",public"`
	C        frontend.Variable `gnark:",public"`
	Expected frontend.Variable `gnark:",public"`
}

func (c *hashFieldElementsCircuit) Define(api frontend.API) error {
	h, err := HashFieldElements(api, c.A, c.B, c.C)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.Expected, h)
	return nil
}

func TestHashFieldElementsMatchesNative(t *testing.T) {
	c := qt.New(t)

	expected, err := dkghash.HashFieldElements(big.NewInt(7), big.NewInt(11), big.NewInt(13))
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&hashFieldElementsCircuit{}, &hashFieldElementsCircuit{
		A:        big.NewInt(7),
		B:        big.NewInt(11),
		C:        big.NewInt(13),
		Expected: expected,
	}, test.WithCurves(ecc.BN254))
}

type shareMaskCircuit struct {
	RoundHash        frontend.Variable `gnark:",public"`
	ContributorIndex frontend.Variable `gnark:",public"`
	RecipientIndex   frontend.Variable `gnark:",public"`
	SharedX          frontend.Variable `gnark:",public"`
	SharedY          frontend.Variable `gnark:",public"`
	KeyIndex         frontend.Variable `gnark:",public"`
	ExpectedSeed     frontend.Variable `gnark:",public"`
	ExpectedMask     frontend.Variable `gnark:",public"`
}

func (c *shareMaskCircuit) Define(api frontend.API) error {
	seed, err := ShareMaskSeed(api, c.RoundHash, c.ContributorIndex, c.RecipientIndex, c.SharedX, c.SharedY)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.ExpectedSeed, seed)
	mask, err := ShareMask(api, seed, c.KeyIndex)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.ExpectedMask, mask)
	return nil
}

// The in-circuit seed and mask must equal crypto/shareenc's, which is what
// the recipient uses to unmask its share.
func TestShareMaskMatchesNative(t *testing.T) {
	c := qt.New(t)
	roundHash := big.NewInt(12345)
	sharedX, _ := new(big.Int).SetString("10815461618510795226726276893454730046020450225029756020987856892208744569026", 10)
	sharedY, _ := new(big.Int).SetString("160151196236506387551997808635915570015226215386948734197202744433655535177", 10)
	shared := types.CurvePoint{X: sharedX, Y: sharedY}
	seed, err := shareenc.ShareMaskSeed(roundHash, 1, 2, shared)
	c.Assert(err, qt.IsNil)
	mask, err := shareenc.ShareMask(seed, MaxK-1)
	c.Assert(err, qt.IsNil)
	other, err := shareenc.ShareMask(seed, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(mask.Cmp(other), qt.Not(qt.Equals), 0, qt.Commentf("the key index separates the masks"))

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&shareMaskCircuit{}, &shareMaskCircuit{
		RoundHash:        roundHash,
		ContributorIndex: big.NewInt(1),
		RecipientIndex:   big.NewInt(2),
		SharedX:          sharedX,
		SharedY:          sharedY,
		KeyIndex:         big.NewInt(MaxK - 1),
		ExpectedSeed:     seed,
		ExpectedMask:     mask,
	}, test.WithCurves(ecc.BN254))
}
