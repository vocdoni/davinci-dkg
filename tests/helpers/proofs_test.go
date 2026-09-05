package helpers

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
)

// TestFinalizeTranscriptDomainMatchesProtocol pins the BRLC domain the
// finalization circuit binds into its challenge to the protocol constant the
// vectors and BuildFinalizeSubmissionWithNonCanonicalWord use: re-deriving
// the challenge from the witness builder's own transcript with the protocol
// domain must reproduce the public input.
func TestFinalizeTranscriptDomainMatchesProtocol(t *testing.T) {
	c := qt.New(t)

	epochID := RoundIDFromString("domain-check")
	commitments, err := commitmentSets([][][]*big.Int{DealPoolCoefficients([]*big.Int{big.NewInt(7)})})
	c.Assert(err, qt.IsNil)
	_, pi, err := finalize.BuildWitness(finalize.Assignment{
		RoundHash:          RoundScalar(epochID),
		Threshold:          1,
		CommitteeSize:      1,
		ParticipantIndexes: []uint16{1},
		Commitments:        commitments,
	})
	c.Assert(err, qt.IsNil)

	words, err := pi.TranscriptScalars()
	c.Assert(err, qt.IsNil)
	anchor, err := ccommon.ChallengeAnchor(words, pi.TranscriptDigest)
	c.Assert(err, qt.IsNil)
	challenge, err := ccommon.DeriveChallengeNative(pi.RoundHash, protocol.DomainFinalizeTranscriptV2, anchor)
	c.Assert(err, qt.IsNil)
	c.Assert(challenge.Cmp(pi.Challenge), qt.Equals, 0,
		qt.Commentf("circuits/finalize derives its challenge under %q", protocol.DomainFinalizeTranscriptV2Str))

	// Every key's share commitments cover the whole committee, member p at
	// slot p−1: with a single coefficient D_{j,1} = 1^0·Ā_j[0] = P_j, and the
	// slot past the committee holds the identity.
	for j := range ccommon.MaxK {
		c.Assert(pi.ShareCommitments[j][0].X.Cmp(pi.PoolKeys[j].X), qt.Equals, 0, qt.Commentf("key %d", j))
		c.Assert(pi.ShareCommitments[j][0].Y.Cmp(pi.PoolKeys[j].Y), qt.Equals, 0)
		c.Assert(pi.ShareCommitments[j][1].X.Sign(), qt.Equals, 0)
		c.Assert(pi.ShareCommitments[j][1].Y.Cmp(big.NewInt(1)), qt.Equals, 0)
	}
	// Key 0 keeps the base polynomial verbatim, so P_0 = 7·G.
	c.Assert(pi.PoolKeys[0].X.Cmp(ScalarBasePoint(big.NewInt(7)).X), qt.Equals, 0)

	// The share trees the harness builds are the ones the contract stores:
	// one root per key, over the member-indexed leaves.
	roots, err := pi.ShareRoots()
	c.Assert(err, qt.IsNil)
	tree, err := CommitteeShareTree(pi.ShareCommitments[3][:1])
	c.Assert(err, qt.IsNil)
	c.Assert(tree.Root(), qt.Equals, roots[3])
}
