package dleq

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

// OrganizerShareChallenge derives the Fiat-Shamir challenge of the
// organizer's Chaum-Pedersen proof that Δ = sk_org·C1 with PK_org = sk_org·G:
//
//	e = keccak256(
//	      DOMAIN_ORGANIZER_SHARE_V1   // 32 bytes
//	    ‖ eid                         // 12 bytes
//	    ‖ aid                         // 32 bytes
//	    ‖ uint256(ctIdx)              // 32 bytes
//	    ‖ PK_org.x ‖ PK_org.y         // 2 × 32
//	    ‖ C1.x ‖ C1.y                 // 2 × 32
//	    ‖ Δ.x ‖ Δ.y                   // 2 × 32
//	    ‖ A1.x ‖ A1.y ‖ A2.x ‖ A2.y   // 4 × 32
//	  ) mod q
//
// with q the BabyJubJub prime-order subgroup order. Coordinates are the
// canonical reduced values used in calldata, left-padded to 32 bytes —
// exactly `abi.encodePacked` of those fields in that order. The
// decrypt-combine circuit consumes `e` as a transcript word, the contract
// recomputes it from calldata and pins the word to it, so this function is
// the one place the encoding is written down on the Go side.
//
// A nil coordinate is treated as zero; callers that care about
// well-formedness validate the points first (see VerifyOrganizerShare).
func OrganizerShareChallenge(
	eid [12]byte,
	aid [32]byte,
	ctIdx uint16,
	pkOrg, c1, delta, a1, a2 types.CurvePoint,
) *big.Int {
	buf := make([]byte, 0, 32+12+32+32+10*32)
	buf = append(buf, protocol.DomainOrganizerShareV1.Bytes()...)
	buf = append(buf, eid[:]...)
	buf = append(buf, aid[:]...)
	buf = append(buf, common.LeftPadBytes(new(big.Int).SetUint64(uint64(ctIdx)).Bytes(), 32)...)
	for _, coordinate := range []*big.Int{
		pkOrg.X, pkOrg.Y,
		c1.X, c1.Y,
		delta.X, delta.Y,
		a1.X, a1.Y, a2.X, a2.Y,
	} {
		value := coordinate
		if value == nil {
			value = new(big.Int)
		}
		buf = append(buf, common.LeftPadBytes(value.Bytes(), 32)...)
	}
	e := new(big.Int).SetBytes(ethcrypto.Keccak256(buf))
	return e.Mod(e, group.ScalarField())
}

// ProveOrganizerShare computes the organizer's share Δ = sk_org·C1 together
// with the Chaum-Pedersen proof (A1 = w·G, A2 = w·C1, z = w + e·sk_org) that
// binds it to (eid, aid, ctIdx) and to PK_org = sk_org·G. The nonce `w` is
// drawn from crypto/rand on every call and never reused: reusing it across
// two challenges leaks sk_org.
func ProveOrganizerShare(
	eid [12]byte,
	aid [32]byte,
	ctIdx uint16,
	skOrg *big.Int,
	c1 types.CurvePoint,
) (types.CurvePoint, Proof, error) {
	order := group.ScalarField()
	if skOrg == nil {
		return types.CurvePoint{}, Proof{}, fmt.Errorf("organizer secret is required")
	}
	secret := new(big.Int).Mod(skOrg, order)
	if secret.Sign() == 0 {
		return types.CurvePoint{}, Proof{}, fmt.Errorf("organizer secret must not be zero modulo the subgroup order")
	}
	if err := group.ValidateEncryptionPoint("c1", c1.X, c1.Y); err != nil {
		return types.CurvePoint{}, Proof{}, err
	}
	c1Point, err := group.Decode(c1)
	if err != nil {
		return types.CurvePoint{}, Proof{}, fmt.Errorf("decode c1: %w", err)
	}

	w, err := rand.Int(rand.Reader, order)
	if err != nil {
		return types.CurvePoint{}, Proof{}, fmt.Errorf("sample organizer nonce: %w", err)
	}

	pkOrgPoint := group.NewPoint()
	pkOrgPoint.ScalarBaseMult(secret)
	deltaPoint := group.NewPoint()
	deltaPoint.ScalarMult(c1Point, secret)
	a1Point := group.NewPoint()
	a1Point.ScalarBaseMult(w)
	a2Point := group.NewPoint()
	a2Point.ScalarMult(c1Point, w)

	delta := group.Encode(deltaPoint)
	proof := Proof{A1: group.Encode(a1Point), A2: group.Encode(a2Point)}
	e := OrganizerShareChallenge(eid, aid, ctIdx, group.Encode(pkOrgPoint), c1, delta, proof.A1, proof.A2)
	proof.Response = new(big.Int).Mod(new(big.Int).Add(w, new(big.Int).Mul(e, secret)), order)
	return delta, proof, nil
}

// VerifyOrganizerShare checks z·G == A1 + e·PK_org and z·C1 == A2 + e·Δ for
// the challenge of OrganizerShareChallenge. Every point is required to be
// canonical, on the curve and inside the prime-order subgroup, Δ must not be
// the identity (Δ = O means the organizer contributed nothing) and z must be
// canonically reduced (z < q) so a share has exactly one valid encoding.
//
// PK_org must come from the on-chain application record, never from whoever
// submitted the share.
func VerifyOrganizerShare(
	eid [12]byte,
	aid [32]byte,
	ctIdx uint16,
	pkOrg, c1, delta types.CurvePoint,
	proof Proof,
) bool {
	order := group.ScalarField()
	if proof.Response == nil || proof.Response.Sign() < 0 || proof.Response.Cmp(order) >= 0 {
		return false
	}
	// Δ, C1 and PK_org must be usable group elements; A1 and A2 may be the
	// identity (a nonce of zero is legal, if useless), so they only get the
	// subgroup check.
	for _, point := range []struct {
		kind  string
		value types.CurvePoint
	}{{"pkOrg", pkOrg}, {"c1", c1}, {"delta", delta}} {
		if err := group.ValidateEncryptionPoint(point.kind, point.value.X, point.value.Y); err != nil {
			return false
		}
	}
	for _, point := range []types.CurvePoint{proof.A1, proof.A2} {
		if !group.IsInPrimeSubgroup(point.X, point.Y) {
			return false
		}
	}

	pkOrgPoint, err := group.Decode(pkOrg)
	if err != nil {
		return false
	}
	c1Point, err := group.Decode(c1)
	if err != nil {
		return false
	}
	deltaPoint, err := group.Decode(delta)
	if err != nil {
		return false
	}
	a1Point, err := group.Decode(proof.A1)
	if err != nil {
		return false
	}
	a2Point, err := group.Decode(proof.A2)
	if err != nil {
		return false
	}

	e := OrganizerShareChallenge(eid, aid, ctIdx, pkOrg, c1, delta, proof.A1, proof.A2)

	zG := group.NewPoint()
	zG.ScalarBaseMult(proof.Response)
	ePK := group.NewPoint()
	ePK.ScalarMult(pkOrgPoint, e)
	rhs1 := group.NewPoint()
	rhs1.Add(a1Point, ePK)
	if !zG.Equal(rhs1) {
		return false
	}

	zC1 := group.NewPoint()
	zC1.ScalarMult(c1Point, proof.Response)
	eDelta := group.NewPoint()
	eDelta.ScalarMult(deltaPoint, e)
	rhs2 := group.NewPoint()
	rhs2.Add(a2Point, eDelta)
	return zC1.Equal(rhs2)
}
