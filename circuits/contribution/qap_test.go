package contribution

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// TestPublicInputsHaveDedicatedLeftRows checks the QAP certificate of
// ccommon.CertifyPublicInputs on the compiled contribution circuit; see the
// helper for why weak simulation-extractability needs it.
func TestPublicInputsHaveDedicatedLeftRows(t *testing.T) {
	c := qt.New(t)
	ccs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &ContributionCircuit{})
	c.Assert(err, qt.IsNil)
	c.Assert(ccs.GetNbPublicVariables(), qt.Equals, 9)
	missing, err := ccommon.MissingDedicatedPublicRows(ccs)
	c.Assert(err, qt.IsNil)
	c.Assert(missing, qt.HasLen, 0, qt.Commentf("public wires without a dedicated left row (0 is the constant wire)"))
}
