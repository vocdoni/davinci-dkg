package dleq

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/types"
)

// Proof is a non-interactive Chaum-Pedersen proof of equal discrete logarithms.
type Proof struct {
	A1       types.CurvePoint
	A2       types.CurvePoint
	Response *big.Int
}

// Prove builds a DLEQ proof for the relation pubKey = x*G and target = x*base
// bound to (roundHash, participantIndex). The same (roundHash, participantIndex)
// must be passed to Verify; otherwise the Fiat-Shamir transcript will not
// match and verification fails.
func Prove(secret, roundHash, participantIndex *big.Int, pubKey, base, target types.CurvePoint) (*Proof, error) {
	if secret == nil {
		return nil, fmt.Errorf("secret is required")
	}
	if roundHash == nil {
		return nil, fmt.Errorf("round hash is required")
	}
	if participantIndex == nil {
		return nil, fmt.Errorf("participant index is required")
	}

	pubKeyPoint, err := group.Decode(pubKey)
	if err != nil {
		return nil, fmt.Errorf("decode public key: %w", err)
	}
	basePoint, err := group.Decode(base)
	if err != nil {
		return nil, fmt.Errorf("decode base point: %w", err)
	}
	targetPoint, err := group.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("decode target point: %w", err)
	}

	modulus := group.ScalarField()
	witness, err := rand.Int(rand.Reader, modulus)
	if err != nil {
		return nil, fmt.Errorf("generate witness: %w", err)
	}
	if witness.Sign() == 0 {
		witness = big.NewInt(1)
	}

	a1 := group.NewPoint()
	a1.ScalarBaseMult(witness)

	a2 := group.NewPoint()
	a2.ScalarMult(basePoint, witness)

	c, err := challenge(roundHash, participantIndex, pubKeyPoint, basePoint, targetPoint, a1, a2)
	if err != nil {
		return nil, err
	}

	response := new(big.Int).Mul(c, secret)
	response.Add(response, witness)
	response.Mod(response, modulus)

	return &Proof{
		A1:       group.Encode(a1),
		A2:       group.Encode(a2),
		Response: response,
	}, nil
}

// challenge derives the Fiat-Shamir challenge:
// e = Hash(domain, rid, idx) -> Hash(., Y) -> Hash(., base) -> Hash(., target)
// -> Hash(., A1) -> Hash(., A2). The leading (rid, idx) binding mirrors the
// in-circuit derivation in circuits/partialdecrypt/circuit.go.
func challenge(
	roundHash, participantIndex *big.Int,
	pubKey, base, target, a1, a2 interface{ Point() (*big.Int, *big.Int) },
) (*big.Int, error) {
	state, err := dkghash.HashFieldElements(
		dkghash.DomainValue(dkghash.DomainPartialDecrypt),
		roundHash,
		participantIndex,
	)
	if err != nil {
		return nil, err
	}
	state, err = hashPointTuple(state, pubKey)
	if err != nil {
		return nil, err
	}
	state, err = hashPointTuple(state, base)
	if err != nil {
		return nil, err
	}
	state, err = hashPointTuple(state, target)
	if err != nil {
		return nil, err
	}
	state, err = hashPointTuple(state, a1)
	if err != nil {
		return nil, err
	}
	return hashPointTuple(state, a2)
}

func hashPointTuple(state *big.Int, point interface{ Point() (*big.Int, *big.Int) }) (*big.Int, error) {
	x, y := point.Point()
	return dkghash.HashFieldElements(state, x, y)
}
