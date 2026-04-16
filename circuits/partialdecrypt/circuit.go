package partialdecrypt

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// PartialDecryptCircuit proves a BabyJubJub Chaum-Pedersen relation for one
// partial decryption: Y = xG, Delta = xM, A1 = wG, A2 = wM.
type PartialDecryptCircuit struct {
	RoundHash        frontend.Variable    `gnark:",public"`
	ParticipantIndex frontend.Variable    `gnark:",public"`
	Base             twistededwards.Point `gnark:",public"`
	PublicKey        twistededwards.Point `gnark:",public"`
	Delta            twistededwards.Point `gnark:",public"`
	A1               twistededwards.Point `gnark:",public"`
	A2               twistededwards.Point `gnark:",public"`
	Response         frontend.Variable    `gnark:",public"`

	Secret frontend.Variable
	Nonce  frontend.Variable
}

func (c *PartialDecryptCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ccommon.BabyJubJubCurveID())
	if err != nil {
		return err
	}

	base := c.Base
	publicKey := c.PublicKey
	delta := c.Delta
	a1 := c.A1
	a2 := c.A2
	for _, point := range []twistededwards.Point{base, publicKey, delta, a1, a2} {
		if err := ccommon.AssertPointOnCurve(api, point); err != nil {
			return err
		}
	}

	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.Secret), publicKey)
	ccommon.AssertPointEqual(api, curve.ScalarMul(base, c.Secret), delta)
	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.Nonce), a1)
	ccommon.AssertPointEqual(api, curve.ScalarMul(base, c.Nonce), a2)

	// SECURITY (H-1 / paper §4.2): bind the challenge to (rid, participant)
	// in addition to the proof points so a transcript cannot be replayed
	// across rounds or impersonated for a different participant index.
	state, err := ccommon.HashFieldElements(
		api,
		ccommon.PartialDecryptDomain(),
		c.RoundHash,
		c.ParticipantIndex,
	)
	if err != nil {
		return err
	}
	challenge, err := ccommon.HashPointTuple(
		api,
		state,
		c.PublicKey,
		c.Base,
		c.Delta,
		c.A1,
		c.A2,
	)
	if err != nil {
		return err
	}

	left1 := ccommon.FixedBaseMul(api, c.Response)
	right1 := curve.Add(a1, curve.ScalarMul(publicKey, challenge))
	ccommon.AssertPointEqual(api, left1, right1)

	left2 := curve.ScalarMul(base, c.Response)
	right2 := curve.Add(a2, curve.ScalarMul(delta, challenge))
	ccommon.AssertPointEqual(api, left2, right2)
	return nil
}
