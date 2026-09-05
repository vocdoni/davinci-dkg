package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

const testEpochID = "epoch-1"

func TestContributionValidate(t *testing.T) {
	c := qt.New(t)

	contribution := Contribution{
		EpochID:          testEpochID,
		Contributor:      common.HexToAddress("0x1000000000000000000000000000000000000001"),
		ContributorIndex: 1,
		Commitments: []CurvePoint{
			{X: big.NewInt(11), Y: big.NewInt(12)},
		},
		EncryptedShares: []EncryptedShare{
			{
				Recipient:      common.HexToAddress("0x2000000000000000000000000000000000000002"),
				RecipientIndex: 2,
				Ephemeral:      CurvePoint{X: big.NewInt(21), Y: big.NewInt(22)},
				Ciphertext:     big.NewInt(33),
			},
		},
		Proof: []byte{0x01, 0x02},
	}

	c.Assert(contribution.Validate(), qt.IsNil)
}

func TestContributionValidateRejectsMissingEncryptedShares(t *testing.T) {
	c := qt.New(t)

	contribution := Contribution{
		EpochID:          testEpochID,
		Contributor:      common.HexToAddress("0x1000000000000000000000000000000000000001"),
		ContributorIndex: 1,
		Commitments: []CurvePoint{
			{X: big.NewInt(11), Y: big.NewInt(12)},
		},
		Proof: []byte{0x01},
	}

	err := contribution.Validate()
	c.Assert(err, qt.Not(qt.IsNil))
	c.Assert(err.Error(), qt.Contains, "encrypted share")
}
