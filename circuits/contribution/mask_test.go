package contribution

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/test"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/types"
)

// shareMaskCircuit derives the ECDH secret, the seed and one mask exactly as
// ContributionCircuit does and checks the masked share.
type shareMaskCircuit struct {
	RoundHash        frontend.Variable    `gnark:",public"`
	ContributorIndex frontend.Variable    `gnark:",public"`
	RecipientIndex   frontend.Variable    `gnark:",public"`
	KeyIndex         frontend.Variable    `gnark:",public"`
	RecipientPubKey  twistededwards.Point `gnark:",public"`
	MaskedShare      frontend.Variable    `gnark:",public"`

	Nonce frontend.Variable
	Share frontend.Variable
}

func (c *shareMaskCircuit) Define(api frontend.API) error {
	nonceBits := api.ToBinary(c.Nonce, 254)
	shared := ccommon.ScalarMulVarBits(api, c.RecipientPubKey, nonceBits)
	seed, err := ccommon.ShareMaskSeed(api, c.RoundHash, c.ContributorIndex, c.RecipientIndex, shared.X, shared.Y)
	if err != nil {
		return err
	}
	mask, err := ccommon.ShareMask(api, seed, c.KeyIndex)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.MaskedShare, api.Add(c.Share, mask))
	return nil
}

// The circuit's masked share equals crypto/shareenc's ciphertext for the
// same nonce, so the recipient's DecryptShare recovers the share.
func TestShareMaskMatchesNative(t *testing.T) {
	c := qt.New(t)
	const keyIndex = MaxKeys - 1
	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)
	recipient := types.NodeKey{PubX: encodedKey.X, PubY: encodedKey.Y}
	nonce := big.NewInt(23)
	roundHash := big.NewInt(12345)
	// A share close to r, so the native-field sum wraps past r without
	// wrapping past p.
	share := new(big.Int).Sub(group.ScalarField(), big.NewInt(5))

	ciphertext, err := shareenc.EncryptShareWithNonceRoundHash(roundHash, 1, 2, keyIndex, share, recipient, nonce)
	c.Assert(err, qt.IsNil)
	recovered, err := shareenc.DecryptShareRoundHash(roundHash, 1, 2, keyIndex, *ciphertext, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(recovered.Cmp(share), qt.Equals, 0)

	witness := &shareMaskCircuit{
		RoundHash:        roundHash,
		ContributorIndex: big.NewInt(1),
		RecipientIndex:   big.NewInt(2),
		KeyIndex:         big.NewInt(keyIndex),
		RecipientPubKey:  ccommon.CircuitPoint(encodedKey),
		MaskedShare:      ciphertext.MaskedShare,
		Nonce:            nonce,
		Share:            share,
	}
	assert := test.NewAssert(t)
	assert.SolvingSucceeded(&shareMaskCircuit{}, witness, test.WithCurves(ecc.BN254))

	// Another key index or another share does not match.
	wrongKey := *witness
	wrongKey.KeyIndex = big.NewInt(keyIndex - 1)
	assert.SolvingFailed(&shareMaskCircuit{}, &wrongKey, test.WithCurves(ecc.BN254))
	wrongShare := *witness
	wrongShare.Share = new(big.Int).Add(share, big.NewInt(1))
	assert.SolvingFailed(&shareMaskCircuit{}, &wrongShare, test.WithCurves(ecc.BN254))

	// A vector whose native-field sum wraps past p: the largest share, q−1,
	// with the first nonce whose mask exceeds p − (q−1). The ciphertext is
	// then smaller than the mask, and circuit and shareenc must still agree.
	share = new(big.Int).Sub(group.ScalarField(), big.NewInt(1))
	nonce = wrappingNonce(c, roundHash, privateKey, share)
	ciphertext, err = shareenc.EncryptShareWithNonceRoundHash(roundHash, 1, 2, keyIndex, share, recipient, nonce)
	c.Assert(err, qt.IsNil)
	recovered, err = shareenc.DecryptShareRoundHash(roundHash, 1, 2, keyIndex, *ciphertext, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(recovered.Cmp(share), qt.Equals, 0)
	witness.Nonce, witness.Share, witness.MaskedShare = nonce, share, ciphertext.MaskedShare
	assert.SolvingSucceeded(&shareMaskCircuit{}, witness, test.WithCurves(ecc.BN254))
}

// wrappingNonce returns the first small nonce for which share + mask ≥ p
// under key index MaxKeys−1, so the masked share wraps modulo p.
func wrappingNonce(c *qt.C, roundHash, privateKey, share *big.Int) *big.Int {
	p := group.BaseField()
	for n := int64(1); n < 1000; n++ {
		nonce := big.NewInt(n)
		ephemeral := group.NewPoint()
		ephemeral.ScalarBaseMult(nonce)
		shared := group.NewPoint()
		shared.ScalarMult(ephemeral, privateKey)
		seed, err := shareenc.ShareMaskSeed(roundHash, 1, 2, group.Encode(shared))
		c.Assert(err, qt.IsNil)
		mask, err := shareenc.ShareMask(seed, MaxKeys-1)
		c.Assert(err, qt.IsNil)
		if new(big.Int).Add(share, mask).Cmp(p) >= 0 {
			return nonce
		}
	}
	c.Fatal("no wrapping nonce below 1000")
	return nil
}
