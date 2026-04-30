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
	"github.com/iden3/go-iden3-crypto/poseidon"

	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

// OperatorProof is the on-chain artifact submitted alongside `pubX, pubY`
// at `registerKey` / `updateKey`.
type OperatorProof struct {
	Ax *big.Int
	Ay *big.Int
	Z  *big.Int
}

// bn254Q is the BN254 scalar field prime — matches `BabyJubJub.Q` in the
// Solidity library and is the modulus the on-chain `_operatorSchnorrChallenge`
// reduces the domain digest into.
var bn254Q, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

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

// operatorSchnorrChallenge implements the same two-pass Poseidon transcript
// as `_operatorSchnorrChallenge` in DKGRegistry.sol:
//
//	inner = T6(domainField, op, pubX, pubY, ax)
//	c     = T3(inner, ay)
//
// where `domainField = uint256(DOMAIN_OPERATOR_REGISTER_V1) % bn254Q`.
func operatorSchnorrChallenge(op common.Address, pubX, pubY, ax, ay *big.Int) (*big.Int, error) {
	domainField := new(big.Int).Mod(
		new(big.Int).SetBytes(protocol.DomainOperatorRegisterV1.Bytes()),
		bn254Q,
	)
	inner, err := poseidon.Hash([]*big.Int{
		domainField,
		new(big.Int).SetBytes(op.Bytes()),
		pubX, pubY, ax,
	})
	if err != nil {
		return nil, fmt.Errorf("inner poseidon: %w", err)
	}
	c, err := poseidon.Hash([]*big.Int{inner, ay})
	if err != nil {
		return nil, fmt.Errorf("outer poseidon: %w", err)
	}
	return c, nil
}
