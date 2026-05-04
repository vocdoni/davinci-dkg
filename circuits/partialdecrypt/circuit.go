package partialdecrypt

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// PartialDecryptCircuit proves a BabyJubJub Chaum-Pedersen relation for one
// partial decryption: Y = xG, Delta = xM, A1 = wG, A2 = wM.
//
// Per paper §4.4 lines 695–704, the Fiat-Shamir challenge
// transcript binds (eid, aid, ctIdx, role, i, G, C_1, D_i, δ_i, A_i, B_i).
// Role tags committee partial decryptions (1) versus organizer shares (2)
// per paper §6.3 line 1161; binding it prevents cross-protocol replay.
type PartialDecryptCircuit struct {
	RoundHash        frontend.Variable    `gnark:",public"` // semantically: eid
	Aid              frontend.Variable    `gnark:",public"` // application identifier
	CtIdx            frontend.Variable    `gnark:",public"` // per-application ciphertext index
	Role             frontend.Variable    `gnark:",public"` // 1 = COMMITTEE, 2 = ORGANIZER
	ParticipantIndex frontend.Variable    `gnark:",public"` // i (committee slot, 0 for organizer)
	Base             twistededwards.Point `gnark:",public"` // C_1
	PublicKey        twistededwards.Point `gnark:",public"` // D_i  (or PK_org for role=ORGANIZER)
	Delta            twistededwards.Point `gnark:",public"` // δ_i  (or Δ_org for role=ORGANIZER)
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

	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.Secret), publicKey)
	ccommon.AssertPointEqual(api, curve.ScalarMul(base, c.Secret), delta)
	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.Nonce), a1)
	ccommon.AssertPointEqual(api, curve.ScalarMul(base, c.Nonce), a2)

	// Role must be 1 (COMMITTEE) or 2 (ORGANIZER).
	// (role - 1) * (role - 2) == 0 ⟺ role ∈ {1, 2}.
	// The Solidity entry-point gates already enforce this per call site
	// (submitPartialDecryption requires role=1; submitOrganizerShare
	// requires role=2), so this is defence-in-depth: the circuit
	// statement itself is now closed under valid roles.
	api.AssertIsEqual(api.Mul(api.Sub(c.Role, 1), api.Sub(c.Role, 2)), 0)

	// Per paper §4.4 lines 695–704: bind the Fiat-Shamir
	// challenge to the full transcript
	//   (eid, aid, ctIdx, role, i, G, C_1, D_i, δ_i, A_i, B_i)
	// so a proof cannot be replayed across epochs, applications,
	// ciphertexts, participants, or roles.
	state, err := ccommon.HashFieldElements(
		api,
		ccommon.PartialDecryptDomain(),
		c.RoundHash, // eid
		c.Aid,
		c.CtIdx,
		c.Role,
		c.ParticipantIndex, // i
	)
	if err != nil {
		return err
	}
	challenge, err := ccommon.HashPointTuple(
		api,
		state,
		c.PublicKey, // D_i (or PK_org)
		c.Base,      // C_1
		c.Delta,     // δ_i
		c.A1,        // A_i
		c.A2,        // B_i
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
