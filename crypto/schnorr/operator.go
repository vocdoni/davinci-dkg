// Package schnorr implements the Schnorr proof of knowledge constructions
// used by the davinci-dkg registration paths. The transcript layout
// matches `solidity/src/DKGRegistry.sol::_operatorSchnorrChallenge` and
// the cross-impl byte equality is asserted by the cmd/operator-schnorr-vectors
// generator.
package schnorr

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254/twistededwards"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

// OperatorProof is the on-chain artifact submitted alongside `pubX, pubY`
// at `registerKey` / `updateKey`.
type OperatorProof struct {
	Ax *big.Int
	Ay *big.Int
	Z  *big.Int
}


// ProveOperatorRegister builds a Schnorr proof of knowledge of `privateKey`
// such that `pubKey = privateKey · G` on BabyJubJub, bound to the
// operator's blockchain address. The returned `(pubX, pubY)` is the
// derived public key; the proof is `(Ax, Ay, Z)`. The witness `w` is
// drawn from a cryptographically secure RNG.
func ProveOperatorRegister(privateKey *big.Int, operator common.Address) (
	pubX, pubY *big.Int, proof OperatorProof, err error,
) {
	curve := twistededwards.GetEdwardsCurve()
	G := curve.Base
	L := curve.Order

	// pubKey = privateKey · G
	var pub twistededwards.PointAffine
	pub.ScalarMultiplication(&G, privateKey)

	// Sample a fresh nonce w ∈ [1, L-1].
	w, err := rand.Int(rand.Reader, &L)
	if err != nil {
		return nil, nil, OperatorProof{}, fmt.Errorf("sample witness: %w", err)
	}
	if w.Sign() == 0 {
		w = big.NewInt(1) // adversarially-improbable but correct fallback
	}

	// A = w · G
	var A twistededwards.PointAffine
	A.ScalarMultiplication(&G, w)

	pubX = pub.X.BigInt(new(big.Int))
	pubY = pub.Y.BigInt(new(big.Int))
	ax := A.X.BigInt(new(big.Int))
	ay := A.Y.BigInt(new(big.Int))

	c, err := operatorSchnorrChallenge(operator, pubX, pubY, ax, ay)
	if err != nil {
		return nil, nil, OperatorProof{}, err
	}

	// z = w + c · privateKey  (mod L)
	z := new(big.Int).Mul(c, privateKey)
	z.Add(z, w)
	z.Mod(z, &L)

	return pubX, pubY, OperatorProof{Ax: ax, Ay: ay, Z: z}, nil
}

// operatorSchnorrChallenge mirrors `_operatorSchnorrChallenge` in
// DKGRegistry.sol:
//
//	c = keccak256(domain || op || pubX || pubY || ax || ay) mod L
func operatorSchnorrChallenge(op common.Address, pubX, pubY, ax, ay *big.Int) (*big.Int, error) {
	curve := twistededwards.GetEdwardsCurve()
	L := curve.Order
	buf := make([]byte, 0, 32+20+32*4)
	buf = append(buf, protocol.DomainOperatorRegisterV1.Bytes()...)
	buf = append(buf, op.Bytes()...)
	buf = append(buf, padTo32(pubX)...)
	buf = append(buf, padTo32(pubY)...)
	buf = append(buf, padTo32(ax)...)
	buf = append(buf, padTo32(ay)...)
	h := ethcrypto.Keccak256(buf)
	c := new(big.Int).SetBytes(h)
	c.Mod(c, &L)
	return c, nil
}

// padTo32 returns a 32-byte big-endian encoding of the integer.
func padTo32(v *big.Int) []byte {
	b := v.Bytes()
	if len(b) > 32 {
		return b[len(b)-32:]
	}
	out := make([]byte, 32)
	copy(out[32-len(b):], b)
	return out
}
