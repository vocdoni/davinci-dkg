package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

func TestContributionRejectsMalformedProof(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}

	epochID, err := helpers.CreateContributionRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	pool := helpers.DealPoolCoefficients([]*big.Int{big.NewInt(21)})
	submission, err := helpers.BuildContributionSubmission(ctx, services, epochID, 1, 1, 1, pool, []uint16{1})
	c.Assert(err, qt.IsNil)
	submission.Proof = submission.Proof[:len(submission.Proof)-32]

	c.Assert(helpers.SubmitContributionAs(ctx, selfActor(), epochID, 1, submission), qt.IsNotNil)
}

// TestFinalizeRejectsBeforeLiveNotBeforeBlock verifies the on-chain
// finalize gate. With contributions in place AND threshold met, finalizeEpoch
// must still revert until block.number reaches policy.liveNotBeforeBlock.
func TestFinalizeRejectsBeforeLiveNotBeforeBlock(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	// Wide gap so the gate is comfortably in the future when we attempt
	// finalize the first time.
	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 200,
	}
	epochID, err := helpers.CreateContributionRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	// Submit a single accepted contribution so the threshold is met.
	pool := helpers.DealPoolCoefficients([]*big.Int{big.NewInt(11)})
	submission, err := helpers.BuildContributionSubmission(ctx, services, epochID, 1, 1, 1, pool, []uint16{1})
	c.Assert(err, qt.IsNil)
	self := selfActor()
	c.Assert(helpers.SubmitContributionAs(ctx, self, epochID, 1, submission), qt.IsNil)

	// The proof is valid for the frozen contribution set; only the gate
	// stands in its way.
	finalization, err := helpers.BuildFinalizeSubmission(ctx, epochID, 1, 1, []uint16{1}, [][][]*big.Int{pool})
	c.Assert(err, qt.IsNil)
	err = helpers.FinalizeEpochAs(ctx, self, epochID, finalization)
	ok, got := helpers.RevertsWith(err, "InvalidPhase")
	c.Assert(ok, qt.IsTrue, qt.Commentf("finalize before liveNotBeforeBlock must revert InvalidPhase, got %s", got))

	// Roll past the gate and finalize successfully.
	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)
	c.Assert(helpers.FinalizeEpochAs(ctx, self, epochID, finalization), qt.IsNil)

	epoch, err := services.Contracts.GetEpoch(ctx, epochID)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Status, qt.Equals, uint8(3))

	// A second finalize is refused: the epoch is already Live.
	err = helpers.FinalizeEpochAs(ctx, self, epochID, finalization)
	ok, got = helpers.RevertsWith(err, "AlreadyLive")
	c.Assert(ok, qt.IsTrue, qt.Commentf("a second finalize must revert AlreadyLive, got %s", got))
}

// TestFinalizeRejectsMissingContribution asserts an epoch with no accepted
// contribution never goes Live, however long the gate has been open.
func TestFinalizeRejectsMissingContribution(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}

	epochID, err := helpers.CreateContributionRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)

	// No contribution means no statement to prove; the contract refuses the
	// call on the count before it looks at any argument.
	err = helpers.FinalizeEpochAs(ctx, selfActor(), epochID, &helpers.FinalizeSubmission{})
	ok, got := helpers.RevertsWith(err, "InsufficientContributions")
	c.Assert(ok, qt.IsTrue, qt.Commentf("finalize must revert with InsufficientContributions, got %s", got))
}

// TestFinalizeRejectsWrongContributionSet feeds finalizeEpoch a dealer row
// the epoch never accepted: the contract cross-checks every contribution hash
// in the transcript against storage before the verifier is ever called.
func TestFinalizeRejectsWrongContributionSet(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	epochID, pool := contributionRoundAtGate(ctx, c, []*big.Int{big.NewInt(7)})
	self := selfActor()

	// Same shape, different coefficients: the proof is internally sound but
	// its contribution hash is not the one on chain.
	forged := [][][]*big.Int{helpers.DealPoolCoefficients([]*big.Int{big.NewInt(1234)})}
	forgery, err := helpers.BuildFinalizeSubmission(ctx, epochID, 1, 1, []uint16{1}, forged)
	c.Assert(err, qt.IsNil)
	err = helpers.FinalizeEpochAs(ctx, self, epochID, forgery)
	ok, got := helpers.RevertsWith(err, "InvalidProofInput")
	c.Assert(ok, qt.IsTrue, qt.Commentf("finalization over a forged contribution set must revert InvalidProofInput, got %s", got))

	// The honest finalization of the same epoch still works afterwards.
	honest, err := helpers.BuildFinalizeSubmission(ctx, epochID, 1, 1, []uint16{1}, [][][]*big.Int{pool})
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.FinalizeEpochAs(ctx, self, epochID, honest), qt.IsNil)

	// And an epoch can only be finalized once.
	err = helpers.FinalizeEpochAs(ctx, self, epochID, honest)
	ok, got = helpers.RevertsWith(err, "AlreadyLive")
	c.Assert(ok, qt.IsTrue, qt.Commentf("re-finalization must revert AlreadyLive, got %s", got))
}

// contributionRoundAtGate drives a single-member epoch through its one
// contribution (key 0 of the dealt pool is `coefficients`, see
// DealPoolCoefficients) and up to the finalize gate, leaving finalization to
// the caller.
func contributionRoundAtGate(ctx context.Context, c *qt.C, coefficients []*big.Int) ([12]byte, [][]*big.Int) {
	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	policy := types.EpochPolicy{
		Threshold:                       1,
		CommitteeSize:                   1,
		MinValidContributions:           1,
		LotteryAlphaBps:                 helpers.DefaultLotteryAlphaBps,
		CommitteeSelectionDeadlineBlock: head + 25,
		KeyAssemblyDeadlineBlock:        head + 50,
		LiveNotBeforeBlock:              head + 51,
	}
	epochID, err := helpers.CreateContributionRound(ctx, services, policy)
	c.Assert(err, qt.IsNil)
	pool := helpers.DealPoolCoefficients(coefficients)
	submission, err := helpers.BuildContributionSubmission(ctx, services, epochID, 1, 1, 1, pool, []uint16{1})
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.SubmitContributionAs(ctx, selfActor(), epochID, 1, submission), qt.IsNil)
	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)
	return epochID, pool
}

func TestPartialDecryptRejectsMalformedProof(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()

	// Register the application and put a real ciphertext on chain, so the
	// tampered proof reaches the verifier instead of tripping the earlier
	// "no such ciphertext" gate.
	aid := randomAid(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, aid, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)

	c1 := helpers.ScalarBasePoint(big.NewInt(3))
	c2 := helpers.ScalarBasePoint(big.NewInt(19))
	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aid, c1, c2)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))

	share := poolShare(c, res, 0, 1)
	partial, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, res.EpochID, aid, assignedIdx, 1, c1, c2, share, big.NewInt(4), res.Shares(0),
	)
	c.Assert(err, qt.IsNil)

	truncated := *partial
	truncated.Input = partial.Input[:len(partial.Input)-32]
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, &truncated), qt.IsNotNil)

	// A well-formed proof with a bogus Merkle path is refused too.
	badPath := *partial
	badPath.ShareProof = make([][32]byte, len(partial.ShareProof))
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, &badPath), qt.IsNotNil)

	// The untouched submission lands.
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, partial), qt.IsNil)
}

func TestRoundCanFinalizeWithMissingContributorWhenPolicyPermits(t *testing.T) {
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
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, epochID), qt.IsNil)

	contributions := [][][]*big.Int{
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(5), big.NewInt(2)}),
	}
	submission0, err := helpers.BuildContributionSubmission(ctx, services, epochID, 2, 3, 1, contributions[0], []uint16{1, 2, 3})
	c.Assert(err, qt.IsNil)
	submission1, err := helpers.BuildContributionSubmission(ctx, services, epochID, 2, 3, 2, contributions[1], []uint16{1, 2, 3})
	c.Assert(err, qt.IsNil)

	self := selfActor()
	c.Assert(helpers.SubmitContributionAs(ctx, self, epochID, 1, submission0), qt.IsNil)
	c.Assert(helpers.SubmitContributionAs(ctx, actor1, epochID, 2, submission1), qt.IsNil)

	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)
	// The pool is dealt over the two accepted contributors only.
	finalization, err := helpers.BuildFinalizeSubmission(ctx, epochID, 2, 3, []uint16{1, 2}, contributions)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.FinalizeEpochAs(ctx, self, epochID, finalization), qt.IsNil)

	epoch, err := services.Contracts.GetEpoch(ctx, epochID)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Status, qt.Equals, uint8(3))
	c.Assert(epoch.ContributionCount, qt.Equals, uint16(2))

	x, y, err := services.Manager.GetPoolKey(services.CallOpts(ctx), epochID, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(x.Cmp(finalization.PoolKey(0).X), qt.Equals, 0)
	c.Assert(y.Cmp(finalization.PoolKey(0).Y), qt.Equals, 0)
}

// TestNonContributingMemberPostsPartials covers decryption liveness of
// `n − t`, not `m − t`: a member that claimed a slot but never contributed
// still received a share of every accepted dealer's polynomial, its D_p is a
// leaf of the pool key's share root, and its partial is accepted and
// combines with a contributor's.
func TestNonContributingMemberPostsPartials(t *testing.T) {
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
	c.Assert(helpers.MineBlocks(ctx, services, helpers.DefaultSeedDelay+1), qt.IsNil)
	c.Assert(helpers.ClaimSlot(ctx, services, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor1, epochID), qt.IsNil)
	c.Assert(helpers.ClaimSlotAs(ctx, actor2, epochID), qt.IsNil)

	// Members 1 and 2 deal to the whole committee; member 3 (actor2) never
	// contributes but is a recipient of both.
	committee := []uint16{1, 2, 3}
	contributions := [][][]*big.Int{
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(3), big.NewInt(1)}),
		helpers.DealPoolCoefficients([]*big.Int{big.NewInt(5), big.NewInt(2)}),
	}
	self := selfActor()
	for i, actor := range []*helpers.TestActor{self, actor1} {
		sub, subErr := helpers.BuildContributionSubmission(ctx, services, epochID, 2, 3, uint16(i+1), contributions[i], committee)
		c.Assert(subErr, qt.IsNil)
		c.Assert(helpers.SubmitContributionAs(ctx, actor, epochID, uint16(i+1), sub), qt.IsNil)
	}
	c.Assert(helpers.WaitForFinalizeGate(ctx, services, epochID), qt.IsNil)

	// The finalization proves the aggregate over the two accepted
	// contributors and publishes a share commitment for all three members.
	contributors := []uint16{1, 2}
	finalization, err := helpers.BuildFinalizeSubmission(ctx, epochID, 2, 3, contributors, contributions)
	c.Assert(err, qt.IsNil)
	c.Assert(finalization.ShareCommitments[0], qt.HasLen, 3)
	c.Assert(helpers.FinalizeEpochAs(ctx, self, epochID, finalization), qt.IsNil)

	shares, err := helpers.RecoverParticipantShares(contributions, 0, committee)
	c.Assert(err, qt.IsNil)
	d3 := helpers.ScalarBasePoint(shares[2])
	c.Assert(d3.X.Cmp(finalization.ShareCommitments[0][2].X), qt.Equals, 0, qt.Commentf("D_3 = e_3·G is leaf 2"))
	c.Assert(d3.Y.Cmp(finalization.ShareCommitments[0][2].Y), qt.Equals, 0)

	aid := randomAid(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, epochID, aid, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)
	plaintext := big.NewInt(77)
	c1, c2, err := elgamal.Encrypt(finalization.PoolKey(0), plaintext)
	c.Assert(err, qt.IsNil)
	ctIdx, err := helpers.SubmitCiphertextAs(ctx, self, epochID, aid, c1, c2)
	c.Assert(err, qt.IsNil)

	// Member 3 posts first: a partial from a slot that never contributed.
	partial3, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, epochID, aid, ctIdx, 3, c1, c2, shares[2], big.NewInt(31), finalization.ShareTree(0),
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, actor2, epochID, aid, 3, ctIdx, partial3), qt.IsNil,
		qt.Commentf("a committee member that did not contribute may post partials"))
	partial1, err := helpers.BuildPartialDecryptionSubmissionFromBase(
		ctx, epochID, aid, ctIdx, 1, c1, c2, shares[0], big.NewInt(29), finalization.ShareTree(0),
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, epochID, aid, 1, ctIdx, partial1), qt.IsNil)

	combine, err := helpers.BuildDecryptCombineOutputFromCiphertext(
		ctx, epochID, aid, ctIdx, 2, c1, c2, helpers.IdentityPoint(), nil,
		[]uint16{1, 3}, []types.CurvePoint{partial1.Delta, partial3.Delta}, plaintext,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.CombineDecryptionAs(ctx, self, epochID, aid, ctIdx, combine), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, epochID, aid, ctIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Plaintext.String(), qt.Equals, plaintext.String())
}

// TestFinalizeRejectsNonCanonicalWord lifts one transcript word to `w + p`.
// The BRLC commitment cannot tell it from `w`, the proof is honest for that
// calldata, and the contract would otherwise store the raw word (a pool key
// itself, or a share-commitment leaf) — so the canonical check in
// BRLC.commitCalldata has to be what refuses it, with
// `TranscriptWordNotInField`. The untouched finalization of the same epoch
// then lands, which pins down that nothing else about the proof was wrong.
func TestFinalizeRejectsNonCanonicalWord(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	epochID, pool := contributionRoundAtGate(ctx, c, []*big.Int{big.NewInt(7)})
	self := selfActor()
	contributions := [][][]*big.Int{pool}

	lifted := map[string]int{
		"pool key P_1.x": finalize.PoolKeyOffset(1),
		"share commitment D_{0,1}.x (member 1 leaf)": finalize.ShareCommitmentOffset(0, 0),
	}
	for label, word := range lifted {
		finalization, err := helpers.BuildFinalizeSubmissionWithNonCanonicalWord(
			ctx, epochID, 1, 1, []uint16{1}, contributions, word,
		)
		c.Assert(err, qt.IsNil, qt.Commentf("%s: the proof itself is valid", label))
		err = helpers.FinalizeEpochAs(ctx, self, epochID, finalization)
		ok, got := helpers.RevertsWith(err, "TranscriptWordNotInField")
		c.Assert(ok, qt.IsTrue, qt.Commentf("%s: expected TranscriptWordNotInField, got %s", label, got))
	}

	honest, err := helpers.BuildFinalizeSubmission(ctx, epochID, 1, 1, []uint16{1}, contributions)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.FinalizeEpochAs(ctx, self, epochID, honest), qt.IsNil,
		qt.Commentf("the canonical encoding of the same statement finalizes"))
}

// TestCreateEpochEarlyAfterAbort covers the narrow early-creation rule: an
// epoch still selecting its committee keeps the cadence, but once it is
// provably dead and recorded as Aborted the next epoch may start at once.
func TestCreateEpochEarlyAfterAbort(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	// Nobody claims, so the committee never fills.
	policy := types.EpochPolicy{
		Threshold:             2,
		CommitteeSize:         2,
		MinValidContributions: 2,
		LotteryAlphaBps:       helpers.DefaultLotteryAlphaBps,
	}
	dead, err := helpers.CreateEpoch(ctx, services, policy)
	c.Assert(err, qt.IsNil)

	// While the newest epoch is in preparation the cadence holds.
	_, err = helpers.CreateEpochNow(ctx, services, policy)
	ok, got := helpers.RevertsWith(err, "InvalidPhase")
	c.Assert(ok, qt.IsTrue, qt.Commentf("early createEpoch behind a selecting epoch must revert InvalidPhase, got %s", got))

	// And an epoch that can still fill cannot be aborted.
	self := selfActor()
	err = helpers.AbortEpochAs(ctx, self, dead)
	ok, got = helpers.RevertsWith(err, "InvalidPhase")
	c.Assert(ok, qt.IsTrue, qt.Commentf("abortEpoch before the selection deadline must revert InvalidPhase, got %s", got))

	epoch, err := services.Contracts.GetEpoch(ctx, dead)
	c.Assert(err, qt.IsNil)
	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.MineBlocks(ctx, services, epoch.Policy.CommitteeSelectionDeadlineBlock-head+1), qt.IsNil)
	c.Assert(helpers.AbortEpochAs(ctx, self, dead), qt.IsNil)
	epoch, err = services.Contracts.GetEpoch(ctx, dead)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Status, qt.Equals, uint8(4)) // Aborted

	// Still well inside the cadence: the Aborted newest epoch is what lets
	// the next one through.
	next, err := services.Manager.NextEpochStartBlock(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	head, err = services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	c.Assert(head < next, qt.IsTrue, qt.Commentf("head %d must be before the cadence %d for the test to mean anything", head, next))

	early, err := helpers.CreateEpochNow(ctx, services, policy)
	c.Assert(err, qt.IsNil, qt.Commentf("createEpoch after an abort must not wait for the cadence"))
	epoch, err = services.Contracts.GetEpoch(ctx, early)
	c.Assert(err, qt.IsNil)
	c.Assert(epoch.Status, qt.Equals, uint8(1)) // CommitteeSelection
}
