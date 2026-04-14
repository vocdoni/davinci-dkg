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

// Prove builds a DLEQ proof for the relation pubKey = x*G and target = x*base.
func Prove(secret *big.Int, pubKey, base, target types.CurvePoint) (*Proof, error) {
	if secret == nil {
		return nil, fmt.Errorf("secret is required")
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

	challenge, err := challenge(pubKeyPoint, basePoint, targetPoint, a1, a2)
	if err != nil {
		return nil, err
	}

	response := new(big.Int).Mul(challenge, secret)
	response.Add(response, witness)
	response.Mod(response, modulus)

	return &Proof{
		A1:       group.Encode(a1),
		A2:       group.Encode(a2),
		Response: response,
	}, nil
}

func challenge(pubKey, base, target, a1, a2 interface{ Point() (*big.Int, *big.Int) }) (*big.Int, error) {
	state, err := hashPointTuple(dkghash.DomainValue(dkghash.DomainPartialDecrypt), pubKey)
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
