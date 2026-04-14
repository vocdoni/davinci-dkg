package dleq

import (
	"fmt"

	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

// Verify checks a DLEQ proof for the relation pubKey = x*G and target = x*base.
func Verify(proof Proof, pubKey, base, target types.CurvePoint) error {
	if proof.Response == nil {
		return fmt.Errorf("response is required")
	}

	pubKeyPoint, err := group.Decode(pubKey)
	if err != nil {
		return fmt.Errorf("decode public key: %w", err)
	}
	basePoint, err := group.Decode(base)
	if err != nil {
		return fmt.Errorf("decode base point: %w", err)
	}
	targetPoint, err := group.Decode(target)
	if err != nil {
		return fmt.Errorf("decode target point: %w", err)
	}
	a1Point, err := group.Decode(proof.A1)
	if err != nil {
		return fmt.Errorf("decode a1: %w", err)
	}
	a2Point, err := group.Decode(proof.A2)
	if err != nil {
		return fmt.Errorf("decode a2: %w", err)
	}

	challenge, err := challenge(pubKeyPoint, basePoint, targetPoint, a1Point, a2Point)
	if err != nil {
		return err
	}

	left1 := group.NewPoint()
	left1.ScalarBaseMult(proof.Response)

	right1 := group.NewPoint()
	right1Term := group.NewPoint()
	right1Term.ScalarMult(pubKeyPoint, challenge)
	right1.Add(a1Point, right1Term)

	left2 := group.NewPoint()
	left2.ScalarMult(basePoint, proof.Response)

	right2 := group.NewPoint()
	right2Term := group.NewPoint()
	right2Term.ScalarMult(targetPoint, challenge)
	right2.Add(a2Point, right2Term)

	if !left1.Equal(right1) || !left2.Equal(right2) {
		return fmt.Errorf("invalid dleq proof")
	}
	return nil
}
