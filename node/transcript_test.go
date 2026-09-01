package node

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
)

// The node recovers other members' encrypted shares from submitContribution
// calldata, so the decoder must track the contract's real ABI layout.
func TestDecodeContributionTranscriptFollowsSubmitContributionABI(t *testing.T) {
	c := qt.New(t)
	const n = ccommon.MaxN
	words := make([]*big.Int, 0, 8*n)
	for i := 0; i < 8*n; i++ {
		words = append(words, big.NewInt(int64(1000+i)))
	}
	// recipientIndexes live at words [2N..3N), ephemerals at [5N..7N),
	// maskedShares at [7N..8N).
	words[2*n] = big.NewInt(7)
	words[5*n], words[5*n+1] = big.NewInt(11), big.NewInt(13)
	words[7*n] = big.NewInt(17)
	transcript, err := encodeWords(words...)
	c.Assert(err, qt.IsNil)

	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	c.Assert(err, qt.IsNil)
	calldata, err := parsed.Pack("submitContribution",
		[12]byte{1}, uint16(1), [32]byte{2}, [32]byte{3}, transcript, []byte{0xaa}, []byte{0xbb})
	c.Assert(err, qt.IsNil)

	eph, masked, idxs, err := decodeContributionTranscript(calldata)
	c.Assert(err, qt.IsNil)
	c.Assert(idxs[0], qt.Equals, uint16(7))
	c.Assert(eph[0][0].Int64(), qt.Equals, int64(11))
	c.Assert(eph[0][1].Int64(), qt.Equals, int64(13))
	c.Assert(masked[0].Int64(), qt.Equals, int64(17))
}
