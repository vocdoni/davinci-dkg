package types

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

func TestContributionJSONRoundTrip(t *testing.T) {
	c := qt.New(t)

	original := Contribution{
		RoundID:          "round-1",
		Contributor:      common.HexToAddress("0x1000000000000000000000000000000000000001"),
		ContributorIndex: 1,
		Commitments: []CurvePoint{
			{X: big.NewInt(11), Y: big.NewInt(22)},
			{X: big.NewInt(33), Y: big.NewInt(44)},
		},
		EncryptedShares: []EncryptedShare{
			{
				Recipient:      common.HexToAddress("0x2000000000000000000000000000000000000002"),
				RecipientIndex: 2,
				Ephemeral:      CurvePoint{X: big.NewInt(55), Y: big.NewInt(66)},
				Ciphertext:     big.NewInt(77),
			},
		},
		Proof:           []byte{1, 2, 3},
		PublicInputHash: common.HexToHash("0x1234"),
	}

	payload, err := json.Marshal(original)
	c.Assert(err, qt.IsNil)

	var decoded Contribution
	c.Assert(json.Unmarshal(payload, &decoded), qt.IsNil)

	c.Assert(decoded.RoundID, qt.Equals, original.RoundID)
	c.Assert(decoded.Contributor, qt.Equals, original.Contributor)
	c.Assert(decoded.ContributorIndex, qt.Equals, original.ContributorIndex)
	c.Assert(decoded.Commitments[0].X.Cmp(original.Commitments[0].X), qt.Equals, 0)
	c.Assert(decoded.Commitments[1].Y.Cmp(original.Commitments[1].Y), qt.Equals, 0)
	c.Assert(decoded.EncryptedShares[0].Recipient, qt.Equals, original.EncryptedShares[0].Recipient)
	c.Assert(decoded.EncryptedShares[0].Ephemeral.X.Cmp(original.EncryptedShares[0].Ephemeral.X), qt.Equals, 0)
	c.Assert(decoded.EncryptedShares[0].Ciphertext.Cmp(original.EncryptedShares[0].Ciphertext), qt.Equals, 0)
	c.Assert(decoded.PublicInputHash, qt.Equals, original.PublicInputHash)
}
