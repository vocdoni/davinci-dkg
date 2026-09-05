package node

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
)

// The node recovers other members' encrypted shares from submitContribution
// calldata, so the decoder must track the contract's real ABI layout.
func TestDecodeContributionTranscriptFollowsSubmitContributionABI(t *testing.T) {
	c := qt.New(t)
	const (
		n       = ccommon.MaxN
		k       = ccommon.MaxK
		idxOff  = 2 * k * n
		ephOff  = idxOff + 3*n
		maskOff = idxOff + 5*n
	)
	words := make([]*big.Int, 0, contribution.TranscriptWords)
	for i := 0; i < contribution.TranscriptWords; i++ {
		words = append(words, big.NewInt(int64(1000+i)))
	}
	words[idxOff] = big.NewInt(7)
	words[ephOff], words[ephOff+1] = big.NewInt(11), big.NewInt(13)
	// The masked share the last pool key sends recipient slot 0.
	words[maskOff+(k-1)*n] = big.NewInt(17)
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
	c.Assert(masked[k-1][0].Int64(), qt.Equals, int64(17))
}
