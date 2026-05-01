package schnorr

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

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
// Transcript layout (mirrored by `_organizerSchnorrChallenge` in
// DKGAppManager.sol):
//
//	c = keccak256(domain || epochId || aid || PK_org || A) mod L
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
	curve := twistededwards.GetEdwardsCurve()
	L := curve.Order
	buf := make([]byte, 0, 32+12+32*5)
	buf = append(buf, protocol.DomainOrganizerRegisterV1.Bytes()...)
	buf = append(buf, epochID[:]...)
	buf = append(buf, aid[:]...)
	buf = append(buf, padTo32(pkX)...)
	buf = append(buf, padTo32(pkY)...)
	buf = append(buf, padTo32(ax)...)
	buf = append(buf, padTo32(ay)...)
	h := ethcrypto.Keccak256(buf)
	c := new(big.Int).SetBytes(h)
	c.Mod(c, &L)
	return c, nil
}

// fmt is only used for error wrapping below.
var _ = fmt.Errorf
