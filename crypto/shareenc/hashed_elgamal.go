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

// EncryptShare masks one Shamir share for a recipient using the protocol
// transcript. keyIndex is the pool key the share belongs to (0..MaxK-1).
func EncryptShare(
	epochID string, contributorIndex, recipientIndex uint16, keyIndex uint8, share *big.Int, recipient types.NodeKey,
) (*Ciphertext, error) {
	modulus := group.ScalarField()
	nonce, err := rand.Int(rand.Reader, modulus)
	if err != nil {
		return nil, fmt.Errorf("generate nonce: %w", err)
	}
	if nonce.Sign() == 0 {
		nonce = big.NewInt(1)
	}

	return EncryptShareWithNonce(epochID, contributorIndex, recipientIndex, keyIndex, share, recipient, nonce)
}

// EncryptShareWithNonce masks one Shamir share for a recipient using caller-provided randomness.
func EncryptShareWithNonce(
	epochID string,
	contributorIndex, recipientIndex uint16,
	keyIndex uint8,
	share *big.Int,
	recipient types.NodeKey,
	nonce *big.Int,
) (*Ciphertext, error) {
	return encryptShareWithRoundValue(
		new(big.Int).SetBytes([]byte(epochID)),
		contributorIndex,
		recipientIndex,
		keyIndex,
		share,
		recipient,
		nonce,
	)
}

// EncryptShareWithNonceRoundHash masks one Shamir share using a numeric epoch hash.
func EncryptShareWithNonceRoundHash(
	roundHash *big.Int,
	contributorIndex, recipientIndex uint16,
	keyIndex uint8,
	share *big.Int,
	recipient types.NodeKey,
	nonce *big.Int,
) (*Ciphertext, error) {
	return encryptShareWithRoundValue(roundHash, contributorIndex, recipientIndex, keyIndex, share, recipient, nonce)
}

func encryptShareWithRoundValue(
	roundValue *big.Int,
	contributorIndex, recipientIndex uint16,
	keyIndex uint8,
	share *big.Int,
	recipient types.NodeKey,
	nonce *big.Int,
) (*Ciphertext, error) {
	if share == nil {
		return nil, fmt.Errorf("share is required")
	}
	if roundValue == nil {
		return nil, fmt.Errorf("epoch hash is required")
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
	if share == nil || share.Sign() < 0 || share.Cmp(group.ScalarField()) >= 0 {
		return nil, fmt.Errorf("share must be a canonical scalar in [0, r)")
	}

	recipientPoint, err := group.Decode(types.CurvePoint{X: recipient.PubX, Y: recipient.PubY})
	if err != nil {
		return nil, fmt.Errorf("decode recipient key: %w", err)
	}

	ephemeral := group.NewPoint()
	ephemeral.ScalarBaseMult(nonce)

	shared := group.NewPoint()
	shared.ScalarMult(recipientPoint, nonce)

	mask, err := shareMask(roundValue, contributorIndex, recipientIndex, keyIndex, group.Encode(shared))
	if err != nil {
		return nil, err
	}

	// One-time pad in the BN254 scalar field: the recipient subtracts the
	// same mask mod p and gets the canonical share back exactly.
	maskedShare := new(big.Int).Add(share, mask)
	maskedShare.Mod(maskedShare, group.BaseField())

	return &Ciphertext{
		Ephemeral:   group.Encode(ephemeral),
		MaskedShare: maskedShare,
	}, nil
}

// DecryptShare removes the masking term from one encrypted share.
func DecryptShare(
	epochID string, contributorIndex, recipientIndex uint16, keyIndex uint8, ciphertext Ciphertext, privateKey *big.Int,
) (*big.Int, error) {
	return DecryptShareRoundHash(
		new(big.Int).SetBytes([]byte(epochID)),
		contributorIndex,
		recipientIndex,
		keyIndex,
		ciphertext,
		privateKey,
	)
}

// DecryptShareRoundHash removes the masking term from one encrypted share using a numeric epoch hash.
func DecryptShareRoundHash(
	roundHash *big.Int,
	contributorIndex, recipientIndex uint16,
	keyIndex uint8,
	ciphertext Ciphertext,
	privateKey *big.Int,
) (*big.Int, error) {
	if privateKey == nil {
		return nil, fmt.Errorf("private key is required")
	}
	if roundHash == nil {
		return nil, fmt.Errorf("epoch hash is required")
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

	mask, err := shareMask(roundHash, contributorIndex, recipientIndex, keyIndex, group.Encode(shared))
	if err != nil {
		return nil, err
	}

	share := new(big.Int).Sub(ciphertext.MaskedShare, mask)
	share.Mod(share, group.BaseField())
	if share.Cmp(group.ScalarField()) >= 0 {
		return nil, fmt.Errorf("recovered share is not a canonical scalar")
	}
	return share, nil
}

// ShareMaskSeed is the per-recipient KDF seed all MaxK masks of one (dealer,
// recipient) pair expand from: Poseidon(domain, roundHash,
// contributorIndex·2^16 + recipientIndex, S.x, S.y) over the ECDH secret S.
// Mirrors circuits/common.ShareMaskSeed.
func ShareMaskSeed(roundHash *big.Int, contributorIndex, recipientIndex uint16, shared types.CurvePoint) (*big.Int, error) {
	return dkghash.HashFieldElements(
		dkghash.DomainValue(dkghash.DomainShareEncryption),
		roundHash,
		new(big.Int).SetUint64((uint64(contributorIndex)<<16)|uint64(recipientIndex)),
		shared.X,
		shared.Y,
	)
}

// ShareMask expands the seed for pool key keyIndex: Poseidon(seed, j). It is
// a uniform element of the BN254 scalar field and is added to the share in
// that field, so no reduction to the subgroup order is involved. Mirrors
// circuits/common.ShareMask.
func ShareMask(seed *big.Int, keyIndex uint8) (*big.Int, error) {
	return dkghash.HashFieldElements(seed, new(big.Int).SetUint64(uint64(keyIndex)))
}

func shareMask(
	roundHash *big.Int, contributorIndex, recipientIndex uint16, keyIndex uint8, shared types.CurvePoint,
) (*big.Int, error) {
	seed, err := ShareMaskSeed(roundHash, contributorIndex, recipientIndex, shared)
	if err != nil {
		return nil, fmt.Errorf("hash mask seed: %w", err)
	}
	mask, err := ShareMask(seed, keyIndex)
	if err != nil {
		return nil, fmt.Errorf("expand mask: %w", err)
	}
	return mask, nil
}
