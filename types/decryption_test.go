package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestPartialDecryptionValidate(t *testing.T) {
	c := qt.New(t)

	decryption := PartialDecryption{
		RoundID:          "round-1",
		Participant:      common.HexToAddress("0x3000000000000000000000000000000000000003"),
		ParticipantIndex: 2,
		CiphertextIndex:  1,
		Delta:            CurvePoint{X: big.NewInt(7), Y: big.NewInt(8)},
		Proof:            []byte{0x0A},
	}

	c.Assert(decryption.Validate(), qt.IsNil)
}

func TestRevealedShareValidateRejectsMissingShare(t *testing.T) {
	c := qt.New(t)

	share := RevealedShare{
		RoundID:          "round-1",
		Participant:      common.HexToAddress("0x3000000000000000000000000000000000000003"),
		ParticipantIndex: 2,
	}

	err := share.Validate()
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "share is required")
}
