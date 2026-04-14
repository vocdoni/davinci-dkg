package revealsubmit

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/types"
)

// PublicInputs is the native representation of the reveal-submit public inputs.
type PublicInputs struct {
	RoundHash        *big.Int
	ParticipantIndex *big.Int
	ShareValue       *big.Int
	ShareCommitment  types.CurvePoint
}

// BuildWitness materializes the reveal-submit native assignment.
func BuildWitness(a Assignment) (*RevealSubmitCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	participantIndex := big.NewInt(int64(a.ParticipantIndex))
	witness := &RevealSubmitCircuit{
		RoundHash:        new(big.Int).Set(a.RoundHash),
		ParticipantIndex: participantIndex,
		ShareValue:       new(big.Int).Set(a.ShareValue),
		ShareCommitment:  ccommon.CircuitPoint(a.ShareCommitment),
	}
	publicInputs := &PublicInputs{
		RoundHash:        new(big.Int).Set(a.RoundHash),
		ParticipantIndex: new(big.Int).Set(participantIndex),
		ShareValue:       new(big.Int).Set(a.ShareValue),
		ShareCommitment:  a.ShareCommitment,
	}
	return witness, publicInputs, nil
}

func (p PublicInputs) PublicWitness() *RevealSubmitCircuit {
	return &RevealSubmitCircuit{
		RoundHash:        p.RoundHash,
		ParticipantIndex: p.ParticipantIndex,
		ShareValue:       p.ShareValue,
		ShareCommitment:  ccommon.CircuitPoint(p.ShareCommitment),
	}
}

func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.ParticipantIndex,
		p.ShareValue,
		p.ShareCommitment.X,
		p.ShareCommitment.Y,
	}
}

func (p PublicInputs) Validate() error {
	if p.RoundHash == nil || p.ParticipantIndex == nil || p.ShareValue == nil {
		return fmt.Errorf("missing required public inputs")
	}
	return nil
}
