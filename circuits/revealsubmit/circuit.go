package revealsubmit

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// RevealSubmitCircuit proves d_i * G = D_i for disclosure phase 2.
type RevealSubmitCircuit struct {
	RoundHash        frontend.Variable    `gnark:",public"`
	ParticipantIndex frontend.Variable    `gnark:",public"`
	ShareValue       frontend.Variable    `gnark:",public"`
	ShareCommitment  twistededwards.Point `gnark:",public"`
}

func (c *RevealSubmitCircuit) Define(api frontend.API) error {
	if err := ccommon.AssertPointOnCurve(api, c.ShareCommitment); err != nil {
		return err
	}
	ccommon.AssertPointEqual(api, ccommon.FixedBaseMul(api, c.ShareValue), c.ShareCommitment)
	return nil
}
