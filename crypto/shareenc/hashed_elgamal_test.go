package shareenc

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestEncryptDecryptShare(t *testing.T) {
	c := qt.New(t)

	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)

	recipient := types.NodeKey{
		Operator: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		PubX:     encodedKey.X,
		PubY:     encodedKey.Y,
	}

	ciphertext, err := EncryptShare("epoch-1", 1, 2, 0, big.NewInt(33), recipient)
	c.Assert(err, qt.IsNil)

	share, err := DecryptShare("epoch-1", 1, 2, 0, *ciphertext, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(big.NewInt(33)), qt.Equals, 0)
}

func TestDecryptShareRejectsWrongKey(t *testing.T) {
	c := qt.New(t)

	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)

	recipient := types.NodeKey{
		Operator: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		PubX:     encodedKey.X,
		PubY:     encodedKey.Y,
	}

	ciphertext, err := EncryptShare("epoch-1", 1, 2, 0, big.NewInt(33), recipient)
	c.Assert(err, qt.IsNil)

	share, err := DecryptShare("epoch-1", 1, 2, 0, *ciphertext, big.NewInt(19))
	// A wrong key unmasks an unrelated field element: usually not even a
	// canonical scalar, so decryption errors; when it happens to be one, it
	// is not the share.
	if err == nil {
		c.Assert(share.Cmp(big.NewInt(33)) == 0, qt.IsFalse)
	}
}

func TestEncryptShareRejectsNonCanonicalShare(t *testing.T) {
	c := qt.New(t)
	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)
	recipient := types.NodeKey{PubX: encodedKey.X, PubY: encodedKey.Y}
	_, err := EncryptShare("epoch-1", 1, 2, 0, group.ScalarField(), recipient)
	c.Assert(err, qt.IsNotNil)
	_, err = EncryptShare("epoch-1", 1, 2, 0, big.NewInt(-1), recipient)
	c.Assert(err, qt.IsNotNil)
}

func TestEncryptShareWithNonce(t *testing.T) {
	c := qt.New(t)

	privateKey := big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)

	recipient := types.NodeKey{
		Operator: common.HexToAddress("0x1000000000000000000000000000000000000001"),
		PubX:     encodedKey.X,
		PubY:     encodedKey.Y,
	}

	nonce := big.NewInt(23)
	ciphertext, err := EncryptShareWithNonce("epoch-1", 1, 2, 0, big.NewInt(33), recipient, nonce)
	c.Assert(err, qt.IsNil)
	c.Assert(ciphertext, qt.Not(qt.IsNil))

	expectedEphemeral := group.NewPoint()
	expectedEphemeral.ScalarBaseMult(nonce)
	expected := group.Encode(expectedEphemeral)
	c.Assert(ciphertext.Ephemeral.X.Cmp(expected.X), qt.Equals, 0)
	c.Assert(ciphertext.Ephemeral.Y.Cmp(expected.Y), qt.Equals, 0)

	share, err := DecryptShare("epoch-1", 1, 2, 0, *ciphertext, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(big.NewInt(33)), qt.Equals, 0)
}

// The same ECDH secret serves every pool key, so the key index must change
// the mask or two keys' shares would be a one-time pad reuse.
func TestKeyIndexChangesMask(t *testing.T) {
	privateKey := big.NewInt(23)
	pub := group.NewPoint()
	pub.ScalarBaseMult(privateKey)
	enc := group.Encode(pub)
	recipient := types.NodeKey{Operator: common.HexToAddress("0x1000000000000000000000000000000000000001"), PubX: enc.X, PubY: enc.Y}
	nonce := big.NewInt(77)
	a, err := EncryptShareWithNonce("epoch-1", 1, 2, 0, big.NewInt(33), recipient, nonce)
	if err != nil {
		t.Fatal(err)
	}
	b, err := EncryptShareWithNonce("epoch-1", 1, 2, 1, big.NewInt(33), recipient, nonce)
	if err != nil {
		t.Fatal(err)
	}
	if a.MaskedShare.Cmp(b.MaskedShare) == 0 {
		t.Fatal("key index does not change the mask")
	}
	if _, err := DecryptShare("epoch-1", 1, 2, 0, *b, privateKey); err == nil {
		got, _ := DecryptShare("epoch-1", 1, 2, 0, *b, privateKey)
		if got.Cmp(big.NewInt(33)) == 0 {
			t.Fatal("share decrypted under the wrong key index")
		}
	}
}

// wrappingVector returns a recipient key pair and the first small nonce for
// which share + mask ≥ p under key index 0, so the ciphertext wraps modulo p.
func wrappingVector(c *qt.C, roundHash, share *big.Int) (privateKey, nonce *big.Int, recipient types.NodeKey, mask *big.Int) {
	privateKey = big.NewInt(17)
	publicPoint := group.NewPoint()
	publicPoint.ScalarBaseMult(privateKey)
	encodedKey := group.Encode(publicPoint)
	recipient = types.NodeKey{PubX: encodedKey.X, PubY: encodedKey.Y}
	p := group.BaseField()
	for n := int64(1); n < 1000; n++ {
		nonce = big.NewInt(n)
		ephemeral := group.NewPoint()
		ephemeral.ScalarBaseMult(nonce)
		shared := group.NewPoint()
		shared.ScalarMult(ephemeral, privateKey)
		seed, err := ShareMaskSeed(roundHash, 1, 2, group.Encode(shared))
		c.Assert(err, qt.IsNil)
		mask, err = ShareMask(seed, 0)
		c.Assert(err, qt.IsNil)
		if new(big.Int).Add(share, mask).Cmp(p) >= 0 {
			return privateKey, nonce, recipient, mask
		}
	}
	c.Fatal("no wrapping nonce below 1000")
	return nil, nil, types.NodeKey{}, nil
}

// TestShareMaskWrapsPastBaseField covers the native-field boundary: the
// masked share is reduced modulo p, so the ciphertext is smaller than the
// mask, and decryption still recovers the share exactly.
func TestShareMaskWrapsPastBaseField(t *testing.T) {
	c := qt.New(t)
	roundHash := big.NewInt(12345)
	share := new(big.Int).Sub(group.ScalarField(), big.NewInt(1))
	privateKey, nonce, recipient, mask := wrappingVector(c, roundHash, share)

	ciphertext, err := EncryptShareWithNonceRoundHash(roundHash, 1, 2, 0, share, recipient, nonce)
	c.Assert(err, qt.IsNil)
	c.Assert(ciphertext.MaskedShare.Cmp(mask) < 0, qt.IsTrue, qt.Commentf("the sum must have wrapped past p"))
	expected := new(big.Int).Add(share, mask)
	expected.Sub(expected, group.BaseField())
	c.Assert(ciphertext.MaskedShare.Cmp(expected), qt.Equals, 0)

	recovered, err := DecryptShareRoundHash(roundHash, 1, 2, 0, *ciphertext, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(recovered.Cmp(share), qt.Equals, 0)
}

// TestDecryptShareRejectsSubgroupOrder: a ciphertext that unmasks to exactly
// q (the first non-canonical value) must be rejected, as must one that
// unmasks to p − 1.
func TestDecryptShareRejectsSubgroupOrder(t *testing.T) {
	c := qt.New(t)
	roundHash := big.NewInt(12345)
	privateKey, nonce, _, mask := wrappingVector(c, roundHash, new(big.Int).Sub(group.ScalarField(), big.NewInt(1)))
	ephemeral := group.NewPoint()
	ephemeral.ScalarBaseMult(nonce)
	p := group.BaseField()
	for _, plaintext := range []*big.Int{group.ScalarField(), new(big.Int).Sub(p, big.NewInt(1))} {
		masked := new(big.Int).Add(plaintext, mask)
		masked.Mod(masked, p)
		ciphertext := Ciphertext{Ephemeral: group.Encode(ephemeral), MaskedShare: masked}
		_, err := DecryptShareRoundHash(roundHash, 1, 2, 0, ciphertext, privateKey)
		c.Assert(err, qt.IsNotNil, qt.Commentf("plaintext %s must be rejected", plaintext))
	}
	// The value just below q is canonical and decrypts.
	masked := new(big.Int).Add(new(big.Int).Sub(group.ScalarField(), big.NewInt(1)), mask)
	masked.Mod(masked, p)
	recovered, err := DecryptShareRoundHash(roundHash, 1, 2, 0, Ciphertext{Ephemeral: group.Encode(ephemeral), MaskedShare: masked}, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(recovered.Cmp(new(big.Int).Sub(group.ScalarField(), big.NewInt(1))), qt.Equals, 0)
}
