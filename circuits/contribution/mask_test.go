package contribution

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	ecc_tweds "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/types"
)

type shareMaskCircuit struct {
	RoundHash        frontend.Variable    `gnark:",public"`
	ContributorIndex frontend.Variable    `gnark:",public"`
	RecipientIndex   frontend.Variable    `gnark:",public"`
	KeyIndex         frontend.Variable    `gnark:",public"`
	RecipientPubKey  twistededwards.Point `gnark:",public"`
	ExpectedMask     frontend.Variable    `gnark:",public"`

	Nonce        frontend.Variable
	MaskQuotient frontend.Variable
}

type sharedSecretCircuit struct {
	RecipientPubKey twistededwards.Point `gnark:",public"`
	ExpectedShared  twistededwards.Point `gnark:",public"`

	Nonce frontend.Variable
}

type directShareMaskCircuit struct {
	RoundHash        frontend.Variable `gnark:",public"`
	ContributorIndex frontend.Variable `gnark:",public"`
	RecipientIndex   frontend.Variable `gnark:",public"`
	KeyIndex         frontend.Variable `gnark:",public"`
	SharedX          frontend.Variable `gnark:",public"`
	SharedY          frontend.Variable `gnark:",public"`
	ExpectedMask     frontend.Variable `gnark:",public"`

	MaskQuotient frontend.Variable
}

func (c *shareMaskCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return err
	}
	shared := curve.ScalarMul(c.RecipientPubKey, c.Nonce)
	rawMask, err := ccommon.ShareMaskHash(
		api,
		c.RoundHash,
		c.ContributorIndex,
		c.RecipientIndex,
		shared.X,
		shared.Y,
		c.KeyIndex,
	)
	if err != nil {
		return err
	}
	mask := ccommon.ReduceToSubgroupOrder(api, rawMask, c.MaskQuotient, c.ExpectedMask)
	api.AssertIsEqual(c.ExpectedMask, mask)
	return nil
}

func TestShareMaskMatchesNative(t *testing.T) {
	c := qt.New(t)

	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)
	nonce := big.NewInt(23)
	share := big.NewInt(33)
	roundHash := big.NewInt(12345)

	// Every pool key index must reproduce its own mask under the same ECDH
	// secret, so exercise the whole range rather than key 0 alone.
	for keyIndex := range MaxKeys {
		recipient := types.NodeKey{PubX: encodedKey.X, PubY: encodedKey.Y}
		ciphertext, err := shareenc.EncryptShareWithNonceRoundHash(
			roundHash, 1, 2, uint8(keyIndex), share, recipient, nonce,
		)
		c.Assert(err, qt.IsNil)

		expectedMask := new(big.Int).Sub(ciphertext.MaskedShare, share)
		expectedMask.Mod(expectedMask, group.ScalarField())
		quotient := maskQuotient(t, roundHash, 1, 2, uint8(keyIndex), encodedKey, nonce, expectedMask)

		witness := &shareMaskCircuit{
			RoundHash:        roundHash,
			ContributorIndex: big.NewInt(1),
			RecipientIndex:   big.NewInt(2),
			KeyIndex:         big.NewInt(int64(keyIndex)),
			RecipientPubKey:  ccommon.CircuitPoint(types.CurvePoint{X: encodedKey.X, Y: encodedKey.Y}),
			ExpectedMask:     expectedMask,
			Nonce:            nonce,
			MaskQuotient:     quotient,
		}
		assert := test.NewAssert(t)
		assert.SolvingSucceeded(&shareMaskCircuit{}, witness, test.WithCurves(ecc.BN254))
	}
}

func (c *directShareMaskCircuit) Define(api frontend.API) error {
	rawMask, err := ccommon.ShareMaskHash(
		api, c.RoundHash, c.ContributorIndex, c.RecipientIndex, c.SharedX, c.SharedY, c.KeyIndex,
	)
	if err != nil {
		return err
	}
	mask := ccommon.ReduceToSubgroupOrder(api, rawMask, c.MaskQuotient, c.ExpectedMask)
	api.AssertIsEqual(c.ExpectedMask, mask)
	return nil
}

func (c *sharedSecretCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return err
	}
	shared := curve.ScalarMul(c.RecipientPubKey, c.Nonce)
	ccommon.AssertPointEqual(api, shared, c.ExpectedShared)
	return nil
}

func TestSharedSecretMatchesNative(t *testing.T) {
	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)
	nonce := big.NewInt(23)

	sharedPoint := group.NewPoint()
	sharedPoint.ScalarMult(publicPoint, nonce)
	expectedShared := group.Encode(sharedPoint)

	witness := &sharedSecretCircuit{
		RecipientPubKey: ccommon.CircuitPoint(types.CurvePoint{X: encodedKey.X, Y: encodedKey.Y}),
		ExpectedShared:  ccommon.CircuitPoint(types.CurvePoint{X: expectedShared.X, Y: expectedShared.Y}),
		Nonce:           nonce,
	}
	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&sharedSecretCircuit{}, witness, test.WithCurves(ecc.BN254))
}

func TestDirectShareMaskMatchesNative(t *testing.T) {
	const keyIndex = MaxKeys - 1

	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	nonce := big.NewInt(23)
	share := big.NewInt(33)
	roundHash := big.NewInt(12345)

	sharedPoint := group.NewPoint()
	sharedPoint.ScalarMult(publicPoint, nonce)
	sharedEncoded := group.Encode(sharedPoint)

	recipient := types.NodeKey{PubX: group.Encode(publicPoint).X, PubY: group.Encode(publicPoint).Y}
	ciphertext, err := shareenc.EncryptShareWithNonceRoundHash(roundHash, 1, 2, keyIndex, share, recipient, nonce)
	qt.New(t).Assert(err, qt.IsNil)
	expectedMask := new(big.Int).Sub(ciphertext.MaskedShare, share)
	expectedMask.Mod(expectedMask, group.ScalarField())
	quotient := maskQuotient(t, roundHash, 1, 2, keyIndex, group.Encode(publicPoint), nonce, expectedMask)

	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&directShareMaskCircuit{}, &directShareMaskCircuit{
		RoundHash:        roundHash,
		ContributorIndex: big.NewInt(1),
		RecipientIndex:   big.NewInt(2),
		KeyIndex:         big.NewInt(keyIndex),
		SharedX:          sharedEncoded.X,
		SharedY:          sharedEncoded.Y,
		ExpectedMask:     expectedMask,
		MaskQuotient:     quotient,
	}, test.WithCurves(ecc.BN254))
}

// maskQuotient recomputes the subgroup-order quotient the circuit takes as a
// witness for one (contributor, recipient, key) mask.
func maskQuotient(
	t *testing.T,
	roundHash *big.Int,
	contributorIndex, recipientIndex uint16,
	keyIndex uint8,
	recipientKey types.CurvePoint,
	nonce, reducedMask *big.Int,
) *big.Int {
	t.Helper()

	recipientPoint, err := group.Decode(recipientKey)
	qt.New(t).Assert(err, qt.IsNil)
	sharedPoint := group.NewPoint()
	sharedPoint.ScalarMult(recipientPoint, nonce)
	rawMask, err := rawShareMask(roundHash, contributorIndex, recipientIndex, keyIndex, group.Encode(sharedPoint))
	qt.New(t).Assert(err, qt.IsNil)
	quotient := new(big.Int).Sub(rawMask, reducedMask)
	return quotient.Div(quotient, group.ScalarField())
}
