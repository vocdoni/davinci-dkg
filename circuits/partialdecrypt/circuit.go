package partialdecrypt

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// PartialDecryptCircuit proves a BabyJubJub Chaum-Pedersen relation for one
// committee partial decryption: D_i = d_i·G, δ_i = d_i·C_1, A_i = w·G,
// B_i = w·C_1.
//
// The Fiat-Shamir challenge transcript binds
// (eid, aid, ctIdx, i, D_i, C_1, δ_i, A_i, B_i), so a proof cannot be
// replayed across epochs, applications, ciphertexts or participants. The
// organizer share is a different object entirely — a keccak Chaum-Pedersen
// proof verified inside the combine circuit, not here.
type PartialDecryptCircuit struct {
	RoundHash        frontend.Variable    `gnark:",public"` // semantically: eid
	Aid              frontend.Variable    `gnark:",public"` // application identifier
	CtIdx            frontend.Variable    `gnark:",public"` // per-application ciphertext index
	ParticipantIndex frontend.Variable    `gnark:",public"` // i (committee slot, one-based)
	Base             twistededwards.Point `gnark:",public"` // C_1
	PublicKey        twistededwards.Point `gnark:",public"` // D_i
	Delta            twistededwards.Point `gnark:",public"` // δ_i
	A1               twistededwards.Point `gnark:",public"` // A_i = w·G
	A2               twistededwards.Point `gnark:",public"` // B_i = w·C_1
	Response         frontend.Variable    `gnark:",public"` // z_i = w + e·d_i

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

	// A canonical d_i: with a cofactor component in C1 the eight witnesses
	// d_i + m·r would all satisfy D_i = d_i·G yet give different δ_i.
	api.AssertIsLessOrEqual(c.Secret, ccommon.SubgroupOrderMinusOne())
	api.AssertIsLessOrEqual(c.Response, ccommon.SubgroupOrderMinusOne())
	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.Secret), publicKey)
	ccommon.AssertPointEqual(api, ccommon.ScalarMulVar(api, base, c.Secret), delta)
	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.Nonce), a1)
	ccommon.AssertPointEqual(api, ccommon.ScalarMulVar(api, base, c.Nonce), a2)

	// Bind the Fiat-Shamir challenge to the full transcript
	//   (eid, aid, ctIdx, i, D_i, C_1, δ_i, A_i, B_i)
	// so a proof cannot be replayed across epochs, applications,
	// ciphertexts or participants.
	state, err := ccommon.HashFieldElements(
		api,
		ccommon.PartialDecryptDomain(),
		c.RoundHash, // eid
		c.Aid,
		c.CtIdx,
		c.ParticipantIndex, // i
	)
	if err != nil {
		return err
	}
	challenge, err := ccommon.HashPointTuple(
		api,
		state,
		c.PublicKey, // D_i
		c.Base,      // C_1
		c.Delta,     // δ_i
		c.A1,        // A_i
		c.A2,        // B_i
	)
	if err != nil {
		return err
	}

	left1 := ccommon.FixedBaseMul(api, c.Response)
	right1 := curve.Add(a1, ccommon.ScalarMulVar(api, publicKey, challenge))
	ccommon.AssertPointEqual(api, left1, right1)

	left2 := ccommon.ScalarMulVar(api, base, c.Response)
	right2 := curve.Add(a2, ccommon.ScalarMulVar(api, delta, challenge))
	ccommon.AssertPointEqual(api, left2, right2)
	return nil
}
