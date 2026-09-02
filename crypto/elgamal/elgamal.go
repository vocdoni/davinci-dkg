// Package elgamal implements exponential ElGamal on BabyJubJub as used by
// applications encrypting to a DKG key: (C1, C2) = (r·G, m·G + r·PK).
package elgamal

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/types"
)

// Encrypt encrypts the scalar m under pk with fresh randomness.
func Encrypt(pk types.CurvePoint, m *big.Int) (c1, c2 types.CurvePoint, err error) {
	r, err := rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return c1, c2, err
	}
	return encryptWithRandomness(pk, m, r)
}

// ApplicationKey derives PK_aid from the epoch key: PK_ep + S·G in mode 0,
// PK_ep + PK_org in mode 1.
func ApplicationKey(pkEp types.CurvePoint, mode uint8, s *big.Int, pkOrg types.CurvePoint) (types.CurvePoint, error) {
	base, err := group.Decode(pkEp)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode epoch key: %w", err)
	}
	term := group.NewPoint()
	switch mode {
	case 0:
		term.ScalarBaseMult(s)
	case 1:
		term, err = group.Decode(pkOrg)
		if err != nil {
			return types.CurvePoint{}, fmt.Errorf("decode organizer key: %w", err)
		}
	default:
		return types.CurvePoint{}, fmt.Errorf("unknown mode %d", mode)
	}
	sum := group.NewPoint()
	sum.Add(base, term)
	return group.Encode(sum), nil
}

// PoK is a Schnorr proof of knowledge of the randomness r behind C1 = r·G,
// bound to (epoch, aid, C1, C2). Committee nodes verify it before releasing
// a partial decryption, so a C1 copied (or re-randomised, C1 + x·G) from
// another application's ciphertext is never decrypted as an oracle.
type PoK struct {
	A types.CurvePoint // w·G
	Z *big.Int         // w + c·r mod q
}

// EncryptWithProof encrypts m under pk and proves knowledge of the randomness.
func EncryptWithProof(epochID [12]byte, aid [32]byte, pk types.CurvePoint, m *big.Int) (c1, c2 types.CurvePoint, pok PoK, err error) {
	r, err := rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return c1, c2, pok, err
	}
	c1, c2, err = encryptWithRandomness(pk, m, r)
	if err != nil {
		return c1, c2, pok, err
	}
	pok, err = ProveKnowledge(epochID, aid, c1, c2, r)
	return c1, c2, pok, err
}

// ProveKnowledge produces the Schnorr proof for a ciphertext whose randomness
// r the caller already holds.
func ProveKnowledge(epochID [12]byte, aid [32]byte, c1, c2 types.CurvePoint, r *big.Int) (PoK, error) {
	var pok PoK
	w, err := rand.Int(rand.Reader, group.ScalarField())
	if err != nil {
		return pok, err
	}
	aPoint := group.NewPoint()
	aPoint.ScalarBaseMult(w)
	pok.A = group.Encode(aPoint)
	c := pokChallenge(epochID, aid, c1, c2, pok.A)
	pok.Z = new(big.Int).Mod(new(big.Int).Add(w, new(big.Int).Mul(c, r)), group.ScalarField())
	return pok, nil
}

// NoPoK is a placeholder proof for callers that drive decryption themselves
// (tests, benchmarks); committee nodes refuse ciphertexts carrying it.
func NoPoK() PoK {
	return PoK{A: types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}, Z: big.NewInt(0)}
}

// VerifyPoK checks z·G == A + c·C1 for the transcript (epoch, aid, C1, C2, A).
func VerifyPoK(epochID [12]byte, aid [32]byte, c1, c2 types.CurvePoint, pok PoK) bool {
	if pok.Z == nil || pok.Z.Sign() < 0 || pok.Z.Cmp(group.ScalarField()) >= 0 {
		return false
	}
	aPoint, err := group.Decode(pok.A)
	if err != nil {
		return false
	}
	c1Point, err := group.Decode(c1)
	if err != nil {
		return false
	}
	c := pokChallenge(epochID, aid, c1, c2, pok.A)
	left := group.NewPoint()
	left.ScalarBaseMult(pok.Z)
	cC1 := group.NewPoint()
	cC1.ScalarMult(c1Point, c)
	right := group.NewPoint()
	right.Add(aPoint, cC1)
	l, r := group.Encode(left), group.Encode(right)
	return l.X.Cmp(r.X) == 0 && l.Y.Cmp(r.Y) == 0
}

// pokChallenge = keccak(dom ‖ eid ‖ aid ‖ c1 ‖ c2 ‖ A) mod q, mirroring
// DKGProtocol.DOMAIN_CIPHERTEXT_POK_V1 and the SDK.
func pokChallenge(epochID [12]byte, aid [32]byte, c1, c2, a types.CurvePoint) *big.Int {
	buf := make([]byte, 0, 32+12+32+6*32)
	buf = append(buf, protocol.DomainCiphertextPoKV1.Bytes()...)
	buf = append(buf, epochID[:]...)
	buf = append(buf, aid[:]...)
	for _, v := range []*big.Int{c1.X, c1.Y, c2.X, c2.Y, a.X, a.Y} {
		buf = append(buf, common.LeftPadBytes(v.Bytes(), 32)...)
	}
	c := new(big.Int).SetBytes(ethcrypto.Keccak256(buf))
	return c.Mod(c, group.ScalarField())
}

func encryptWithRandomness(pk types.CurvePoint, m, r *big.Int) (c1, c2 types.CurvePoint, err error) {
	pkPoint, err := group.Decode(pk)
	if err != nil {
		return c1, c2, fmt.Errorf("decode pk: %w", err)
	}
	c1Point := group.NewPoint()
	c1Point.ScalarBaseMult(r)
	mG := group.NewPoint()
	mG.ScalarBaseMult(m)
	rPK := group.NewPoint()
	rPK.ScalarMult(pkPoint, r)
	c2Point := group.NewPoint()
	c2Point.Add(mG, rPK)
	return group.Encode(c1Point), group.Encode(c2Point), nil
}
