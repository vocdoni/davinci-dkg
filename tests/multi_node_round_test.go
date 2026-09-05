package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestCommitteeRoundHappyPath drives a three-member epoch by hand: lottery
// claims, three contributions each dealing the whole pool, the batched
// finalization proof that stores every key, and a threshold decryption under an
// organizer-locked application whose secret is revealed before the first
// partial — the contract accepts none until then.
func TestCommitteeRoundHappyPath(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	actor1, err := services.Actor(1)
	c.Assert(err, qt.IsNil)
	actor2, err := services.Actor(2)
	c.Assert(err, qt.IsNil)

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.EpochPolicy{
		Threshold:                       2,
		CommitteeSize:                   3,
		MinValidContributions:           2,
		LotteryAlphaBps:                 helpers.DefaultLotteryAlphaBps,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}

	epochID, err := helpers.CreateEpoch(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	// Lottery flow: advance past seedBlock so blockhash is available, then have
	// each registered actor self-claim a slot. The committee fills first-come
	// first-served; there is no organizer-driven SelectParticipants step.
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, epochID), qt.IsNil)

	epoch, err := helpers.WaitEpochPhase(ctx, services, epochID, 2)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Policy.CommitteeSize, qt.Equals, uint16(3))

	contributions := [][][]*big.Int{
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(5), big.NewInt(2)}),
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(7), big.NewInt(4)}),
	}
	participantIndexes := []uint16{1, 2, 3}
	actors := []*helpers.TestActor{selfActor(), actor1, actor2}

	for i, actor := range actors {
		sub, subErr := helpers.BuildContributionSubmission(
			ctx, services, epochID, 2, 3, uint16(i+1), contributions[i], participantIndexes,
		)
		c.Assert(subErr, qt.IsNil)
		c.Assert(helpers.SubmitContributionAs(ctx, actor, epochID, uint16(i+1), sub), qt.IsNil)
	}

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)
	finalization, err := helpers.BuildFinalizeSubmission(ctx, epochID, 2, 3, participantIndexes, contributions)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.FinalizeEpochAs(ctx, selfActor(), epochID, finalization), qt.IsNil)

	epoch, err = helpers.WaitEpochPhase(ctx, services, epochID, 3)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.ContributionCount, qt.Equals, uint16(3))

	shareTree := finalization.ShareTree(0)
	root, err := services.Manager.GetPoolShareRoot(services.CallOpts(ctx), epochID, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(root, qt.Equals, shareTree.Root())

	recoveredShares, err := helpers.RecoverParticipantShares(contributions, 0, participantIndexes)
	c.Assert(err, qt.IsNil)
	c.Assert(recoveredShares, qt.HasLen, 3)

	// An organizer-locked application over the epoch's first pool key.
	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)
	self := selfActor()
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, epochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	partial0, err := helpers.BuildPartialDecryptionSubmission(
		ctx, epochID, aid, 1, 1, big.NewInt(9), recoveredShares[0], big.NewInt(11), shareTree,
	)
	c.Assert(err, qt.IsNil)
	partial1, err := helpers.BuildPartialDecryptionSubmission(
		ctx, epochID, aid, 1, 2, big.NewInt(9), recoveredShares[1], big.NewInt(13), shareTree,
	)
	c.Assert(err, qt.IsNil)

	combineOutput, err := helpers.BuildDecryptCombineOutput(
		ctx, epochID, aid, 1, 2, big.NewInt(9), skOrg,
		[]uint16{1, 2}, []types.CurvePoint{partial0.Delta, partial1.Delta}, big.NewInt(3),
	)
	c.Assert(err, qt.IsNil)

	// submitCiphertext must precede submitPartialDecryption so the
	// proof's pi[4..5] can be bound against the on-chain C1.
	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, epochID, aid, combineOutput.CiphertextC1, combineOutput.CiphertextC2)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))

	partial0.C2 = combineOutput.CiphertextC2
	partial1.C2 = combineOutput.CiphertextC2

	// Sealed application: the contract refuses the partial outright.
	err = helpers.SubmitPartialDecryptionAs(ctx, self, epochID, aid, 1, assignedIdx, partial0)
	ok, got := helpers.RevertsWith(err, "OrganizerSecretNotRevealed")
	c.Assert(ok, qt.IsTrue, qt.Commentf("partial before the reveal must revert OrganizerSecretNotRevealed, got %s", got))

	// The organizer publishes sk_org once; from then on the committee owns
	// the whole decryption.
	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, epochID, aid, skOrg), qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, epochID, aid, 1, assignedIdx, partial0), qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, actor1, epochID, aid, 2, assignedIdx, partial1), qt.IsNil)
	c.Assert(helpers.CombineDecryptionAs(ctx, self, epochID, aid, assignedIdx, combineOutput), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, epochID, aid, assignedIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Completed, qt.IsTrue)
	c.Assert(record.Plaintext.String(), qt.Equals, "3")
}
