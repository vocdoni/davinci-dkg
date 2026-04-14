package common

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
)

type hashFieldElementsCircuit struct {
	A        frontend.Variable `gnark:",public"`
	B        frontend.Variable `gnark:",public"`
	C        frontend.Variable `gnark:",public"`
	Expected frontend.Variable `gnark:",public"`
}

type shareMaskMetaCircuit struct {
	RoundHash        frontend.Variable `gnark:",public"`
	ContributorIndex frontend.Variable `gnark:",public"`
	RecipientIndex   frontend.Variable `gnark:",public"`
	Expected         frontend.Variable `gnark:",public"`
}

func (c *hashFieldElementsCircuit) Define(api frontend.API) error {
	got, err := HashFieldElements(api, c.A, c.B, c.C)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.Expected, got)
	return nil
}

func (c *shareMaskMetaCircuit) Define(api frontend.API) error {
	packedIndexes := api.Add(api.Mul(c.ContributorIndex, recipientIndexShift), c.RecipientIndex)
	meta, err := HashFieldElements(api, ShareEncryptionDomain(), c.RoundHash, packedIndexes)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.Expected, meta)
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

func TestShareMaskTranscriptTuplesMatchNative(t *testing.T) {
	c := qt.New(t)

	roundHash := big.NewInt(12345)
	packed := new(big.Int).SetUint64((1 << 16) | 2)
	meta, err := dkghash.HashFieldElements(ShareEncryptionDomain(), roundHash, packed)
	c.Assert(err, qt.IsNil)
	expected, err := dkghash.HashFieldElements(meta, big.NewInt(17), big.NewInt(19))
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&hashFieldElementsCircuit{}, &hashFieldElementsCircuit{
		A:        meta,
		B:        big.NewInt(17),
		C:        big.NewInt(19),
		Expected: expected,
	}, test.WithCurves(ecc.BN254))
}

func TestShareMaskSharedPointHashMatchesNative(t *testing.T) {
	c := qt.New(t)

	meta, err := dkghash.HashFieldElements(ShareEncryptionDomain(), big.NewInt(12345), new(big.Int).SetUint64((1<<16)|2))
	c.Assert(err, qt.IsNil)
	sharedX, _ := new(big.Int).SetString("10815461618510795226726276893454730046020450225029756020987856892208744569026", 10)
	sharedY, _ := new(big.Int).SetString("160151196236506387551997808635915570015226215386948734197202744433655535177", 10)
	expected, err := dkghash.HashFieldElements(meta, sharedX, sharedY)
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&hashFieldElementsCircuit{}, &hashFieldElementsCircuit{
		A:        meta,
		B:        sharedX,
		C:        sharedY,
		Expected: expected,
	}, test.WithCurves(ecc.BN254))
}

func TestShareMaskMetaMatchesNative(t *testing.T) {
	c := qt.New(t)

	roundHash := big.NewInt(12345)
	meta, err := dkghash.HashFieldElements(ShareEncryptionDomain(), roundHash, new(big.Int).SetUint64((1<<16)|2))
	c.Assert(err, qt.IsNil)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&shareMaskMetaCircuit{}, &shareMaskMetaCircuit{
		RoundHash:        roundHash,
		ContributorIndex: big.NewInt(1),
		RecipientIndex:   big.NewInt(2),
		Expected:         meta,
	}, test.WithCurves(ecc.BN254))
}
