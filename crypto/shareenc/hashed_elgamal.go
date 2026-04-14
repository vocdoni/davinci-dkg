package shareenc

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/types"
)

// Ciphertext is the native share-encryption payload.
type Ciphertext struct {
	Ephemeral   types.CurvePoint
	MaskedShare *big.Int
}

// EncryptShare masks one Shamir share for a recipient using the protocol transcript.
func EncryptShare(roundID string, contributorIndex, recipientIndex uint16, share *big.Int, recipient types.NodeKey) (*Ciphertext, error) {
	modulus := group.ScalarField()
	nonce, err := rand.Int(rand.Reader, modulus)
	if err != nil {
		return nil, fmt.Errorf("generate nonce: %w", err)
	}
	if nonce.Sign() == 0 {
		nonce = big.NewInt(1)
	}

	return EncryptShareWithNonce(roundID, contributorIndex, recipientIndex, share, recipient, nonce)
}

// EncryptShareWithNonce masks one Shamir share for a recipient using caller-provided randomness.
func EncryptShareWithNonce(
	roundID string,
	contributorIndex, recipientIndex uint16,
	share *big.Int,
	recipient types.NodeKey,
	nonce *big.Int,
) (*Ciphertext, error) {
	return encryptShareWithRoundValue(
		new(big.Int).SetBytes([]byte(roundID)),
		contributorIndex,
		recipientIndex,
		share,
		recipient,
		nonce,
	)
}

// EncryptShareWithNonceRoundHash masks one Shamir share using a numeric round hash.
func EncryptShareWithNonceRoundHash(
	roundHash *big.Int,
	contributorIndex, recipientIndex uint16,
	share *big.Int,
	recipient types.NodeKey,
	nonce *big.Int,
) (*Ciphertext, error) {
	return encryptShareWithRoundValue(roundHash, contributorIndex, recipientIndex, share, recipient, nonce)
}

func encryptShareWithRoundValue(
	roundValue *big.Int,
	contributorIndex, recipientIndex uint16,
	share *big.Int,
	recipient types.NodeKey,
	nonce *big.Int,
) (*Ciphertext, error) {
	if share == nil {
		return nil, fmt.Errorf("share is required")
	}
	if roundValue == nil {
		return nil, fmt.Errorf("round hash is required")
	}
	if nonce == nil {
		return nil, fmt.Errorf("nonce is required")
	}
	if recipient.PubX == nil || recipient.PubY == nil {
		return nil, fmt.Errorf("recipient public key coordinates are required")
	}
	if contributorIndex == 0 || recipientIndex == 0 {
		return nil, fmt.Errorf("participant indices are required")
	}

	modulus := group.ScalarField()
	nonce = new(big.Int).Mod(new(big.Int).Set(nonce), modulus)
	if nonce.Sign() == 0 {
		return nil, fmt.Errorf("nonce must be non-zero")
	}

	recipientPoint, err := group.Decode(types.CurvePoint{X: recipient.PubX, Y: recipient.PubY})
	if err != nil {
		return nil, fmt.Errorf("decode recipient key: %w", err)
	}

	ephemeral := group.NewPoint()
	ephemeral.ScalarBaseMult(nonce)

	shared := group.NewPoint()
	shared.ScalarMult(recipientPoint, nonce)

	mask, err := shareMask(roundValue, contributorIndex, recipientIndex, group.Encode(shared))
	if err != nil {
		return nil, err
	}

	maskedShare := new(big.Int).Add(share, mask)
	maskedShare.Mod(maskedShare, modulus)

	return &Ciphertext{
		Ephemeral:   group.Encode(ephemeral),
		MaskedShare: maskedShare,
	}, nil
}

// DecryptShare removes the masking term from one encrypted share.
func DecryptShare(roundID string, contributorIndex, recipientIndex uint16, ciphertext Ciphertext, privateKey *big.Int) (*big.Int, error) {
	return DecryptShareRoundHash(
		new(big.Int).SetBytes([]byte(roundID)),
		contributorIndex,
		recipientIndex,
		ciphertext,
		privateKey,
	)
}

// DecryptShareRoundHash removes the masking term from one encrypted share using a numeric round hash.
func DecryptShareRoundHash(
	roundHash *big.Int,
	contributorIndex, recipientIndex uint16,
	ciphertext Ciphertext,
	privateKey *big.Int,
) (*big.Int, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("private key is required")
	}
	if roundHash == nil {
		return nil, fmt.Errorf("round hash is required")
	}
	if ciphertext.MaskedShare == nil {
		return nil, fmt.Errorf("masked share is required")
	}

	ephemeral, err := group.Decode(ciphertext.Ephemeral)
	if err != nil {
		return nil, fmt.Errorf("decode ephemeral point: %w", err)
	}

	shared := group.NewPoint()
	shared.ScalarMult(ephemeral, privateKey)

	mask, err := shareMask(roundHash, contributorIndex, recipientIndex, group.Encode(shared))
	if err != nil {
		return nil, err
	}

	share := new(big.Int).Sub(ciphertext.MaskedShare, mask)
	share.Mod(share, group.ScalarField())
	return share, nil
}

func shareMask(roundHash *big.Int, contributorIndex, recipientIndex uint16, shared types.CurvePoint) (*big.Int, error) {
	meta, err := dkghash.HashFieldElements(
		dkghash.DomainValue(dkghash.DomainShareEncryption),
		roundHash,
		new(big.Int).SetUint64((uint64(contributorIndex)<<16)|uint64(recipientIndex)),
	)
	if err != nil {
		return nil, fmt.Errorf("hash metadata: %w", err)
	}

	mask, err := dkghash.HashFieldElements(meta, shared.X, shared.Y)
	if err != nil {
		return nil, fmt.Errorf("hash shared secret: %w", err)
	}
	mask.Mod(mask, group.ScalarField())
	return mask, nil
}
