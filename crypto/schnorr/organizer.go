package schnorr

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

// OrganizerProof is the on-chain artifact submitted by an application
// organizer at registerApplicationCoDec time.
type OrganizerProof struct {
	Ax *big.Int
	Ay *big.Int
	Z  *big.Int
}

// ProveOrganizerRegister builds a Schnorr proof of knowledge of `sk_org`
// such that `PK_org = sk_org · G` on BabyJubJub, bound to the (epochId, aid)
// the application is being registered for. Returns the derived public key
// and the proof.
//
// Transcript layout (paper §6.2 line 1138, mirrored by
// `_organizerSchnorrChallenge` in DKGManager.sol):
//
//	inner = Poseidon([domain, eid, PK_org.x, PK_org.y])  (T5)
//	c     = Poseidon([inner, aid_field, A.x, A.y])        (T5)
//
// `aid_field = uint256(aid) % bn254Q` to keep the input in the BN254
// scalar field.
func ProveOrganizerRegister(
	skOrg *big.Int,
	epochID [12]byte,
	aid [32]byte,
) (pkX, pkY *big.Int, proof OrganizerProof, err error) {
	curve := twistededwards.GetEdwardsCurve()
	G := curve.Base
	L := curve.Order

	// PK_org = sk_org · G
	var pub twistededwards.PointAffine
	pub.ScalarMultiplication(&G, skOrg)

	// Sample fresh nonce w ∈ [1, L-1].
	w, err := rand.Int(rand.Reader, &L)
	if err != nil {
		return nil, nil, OrganizerProof{}, fmt.Errorf("sample witness: %w", err)
	}
	if w.Sign() == 0 {
		w = big.NewInt(1)
	}

	// A = w · G
	var A twistededwards.PointAffine
	A.ScalarMultiplication(&G, w)

	pkX = pub.X.BigInt(new(big.Int))
	pkY = pub.Y.BigInt(new(big.Int))
	ax := A.X.BigInt(new(big.Int))
	ay := A.Y.BigInt(new(big.Int))

	c, err := organizerSchnorrChallenge(epochID, aid, pkX, pkY, ax, ay)
	if err != nil {
		return nil, nil, OrganizerProof{}, err
	}

	// z = w + c · sk_org  (mod L)
	z := new(big.Int).Mul(c, skOrg)
	z.Add(z, w)
	z.Mod(z, &L)

	return pkX, pkY, OrganizerProof{Ax: ax, Ay: ay, Z: z}, nil
}

func organizerSchnorrChallenge(
	epochID [12]byte, aid [32]byte,
	pkX, pkY, ax, ay *big.Int,
) (*big.Int, error) {
	domainField := new(big.Int).Mod(
		new(big.Int).SetBytes(protocol.DomainOrganizerRegisterV1.Bytes()),
		bn254Q,
	)
	eidField := new(big.Int).SetBytes(epochID[:])
	aidField := new(big.Int).Mod(new(big.Int).SetBytes(aid[:]), bn254Q)

	inner, err := poseidon.Hash([]*big.Int{domainField, eidField, pkX, pkY})
	if err != nil {
		return nil, fmt.Errorf("inner poseidon: %w", err)
	}
	c, err := poseidon.Hash([]*big.Int{inner, aidField, ax, ay})
	if err != nil {
		return nil, fmt.Errorf("outer poseidon: %w", err)
	}
	return c, nil
}
