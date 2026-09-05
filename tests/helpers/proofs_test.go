package helpers

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/poolkey"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

// TestPoolKeyTranscriptDomainMatchesProtocol pins the BRLC domain the pool
// key circuit binds into its challenge to the protocol constant the vectors
// and BuildPoolKeyActivationWithNonCanonicalWord use: re-deriving the
// challenge from the witness builder's own transcript with the protocol
// domain must reproduce the public input.
func TestPoolKeyTranscriptDomainMatchesProtocol(t *testing.T) {
	c := qt.New(t)

	epochID := RoundIDFromString("domain-check")
	commitments, err := commitmentSets([][][]*big.Int{DealPoolCoefficients([]*big.Int{big.NewInt(7)})})
	c.Assert(err, qt.IsNil)
	_, pi, err := poolkey.BuildWitness(poolkey.Assignment{
		RoundHash:          RoundScalar(epochID),
		Threshold:          1,
		CommitteeSize:      1,
		KeyIndex:           0,
		ParticipantIndexes: []uint16{1},
		Commitments:        commitments,
	})
	c.Assert(err, qt.IsNil)

	words, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	anchor, err := ccommon.ChallengeAnchor(words, pi.TranscriptDigest)
	c.Assert(err, qt.IsNil)
	challenge, err := ccommon.DeriveChallengeNative(pi.RoundHash, protocol.DomainPoolKeyTranscriptV1, anchor)
	c.Assert(err, qt.IsNil)
	c.Assert(challenge.Cmp(pi.Challenge), qt.Equals, 0,
		qt.Commentf("circuits/poolkey derives its challenge under %q", protocol.DomainPoolKeyTranscriptV1Str))

	// The share commitments cover the whole committee, member p at slot p−1:
	// with a single coefficient D_1 = 1^0·Ā[0] = Ā[0], and the slot past the
	// committee holds the identity.
	c.Assert(pi.ShareCommitments[0].X.Cmp(pi.AggregateCommitments[0].X), qt.Equals, 0)
	c.Assert(pi.ShareCommitments[0].Y.Cmp(pi.AggregateCommitments[0].Y), qt.Equals, 0)
	c.Assert(pi.ShareCommitments[1].X.Sign(), qt.Equals, 0)
	c.Assert(pi.ShareCommitments[1].Y.Cmp(big.NewInt(1)), qt.Equals, 0)
}
