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

	ciphertext, err := EncryptShare("round-1", 1, 2, big.NewInt(33), recipient)
	c.Assert(err, qt.IsNil)

	share, err := DecryptShare("round-1", 1, 2, *ciphertext, privateKey)
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

	ciphertext, err := EncryptShare("round-1", 1, 2, big.NewInt(33), recipient)
	c.Assert(err, qt.IsNil)

	share, err := DecryptShare("round-1", 1, 2, *ciphertext, big.NewInt(19))
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(big.NewInt(33)) == 0, qt.IsFalse)
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
	ciphertext, err := EncryptShareWithNonce("round-1", 1, 2, big.NewInt(33), recipient, nonce)
	c.Assert(err, qt.IsNil)
	c.Assert(ciphertext, qt.Not(qt.IsNil))

	expectedEphemeral := group.NewPoint()
	expectedEphemeral.ScalarBaseMult(nonce)
	expected := group.Encode(expectedEphemeral)
	c.Assert(ciphertext.Ephemeral.X.Cmp(expected.X), qt.Equals, 0)
	c.Assert(ciphertext.Ephemeral.Y.Cmp(expected.Y), qt.Equals, 0)

	share, err := DecryptShare("round-1", 1, 2, *ciphertext, privateKey)
	c.Assert(err, qt.IsNil)
	c.Assert(share.Cmp(big.NewInt(33)), qt.Equals, 0)
}
