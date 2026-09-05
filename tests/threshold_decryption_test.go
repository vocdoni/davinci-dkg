package tests

import (
	"context"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestThresholdDecryptionAutomaticApplication decrypts under an automatic
// application: no organizer half at all, so `t` partials plus the combine
// proof (with sk_org = 0 and PK_org = identity) are the whole story.
func TestThresholdDecryptionAutomaticApplication(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()
	aid := randomAid(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, aid, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)

	share := poolShare(c, res, 0, 1)
	activation := res.Activation(0)

	partial, err := helpers.BuildPartialDecryptionSubmission(
		ctx, res.EpochID, aid, 1, 1, big.NewInt(9), share, big.NewInt(5), activation.Shares,
	)
	c.Assert(err, qt.IsNil)

	combine, err := helpers.BuildDecryptCombineOutput(ctx, res.EpochID, aid, 1, 1, big.NewInt(9), nil,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, big.NewInt(3))
	c.Assert(err, qt.IsNil)
	c.Assert(combine.OrganizerPK.Y.Cmp(big.NewInt(1)), qt.Equals, 0, qt.Commentf("automatic apps carry the identity key"))

	// submitCiphertext must precede submitPartialDecryption: the
	// partial-decrypt verifier binds pi[4..5] to the on-chain C1.
	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aid, combine.CiphertextC1, combine.CiphertextC2)
	c.Assert(err, qt.IsNil)
	c.Assert(assignedIdx, qt.Equals, uint16(1))

	partial.C2 = combine.CiphertextC2
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, partial), qt.IsNil)
	c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aid, assignedIdx, combine), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, res.EpochID, aid, assignedIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Completed, qt.IsTrue)
	c.Assert(record.Plaintext.String(), qt.Equals, "3")
}

// TestThresholdDecryptionLockedApplicationNeedsTheReveal drives the same
// ciphertext under an organizer-locked application. Until `sk_org` is
// revealed the contract refuses every partial (`OrganizerSecretNotRevealed`):
// nothing of the committee's half exists on chain, so the organizer learns
// the result together with everyone else. After the reveal the partial
// lands, a combine carrying the identity organizer key is still refused, and
// the honest combine recovers the plaintext.
func TestThresholdDecryptionLockedApplicationNeedsTheReveal(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()
	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	share := poolShare(c, res, 0, 1)
	activation := res.Activation(0)
	const base = 9
	plaintext := big.NewInt(3)

	partial, err := helpers.BuildPartialDecryptionSubmission(
		ctx, res.EpochID, aid, 1, 1, big.NewInt(base), share, big.NewInt(5), activation.Shares,
	)
	c.Assert(err, qt.IsNil)

	combine, err := helpers.BuildDecryptCombineOutput(ctx, res.EpochID, aid, 1, 1, big.NewInt(base), skOrg,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, plaintext)
	c.Assert(err, qt.IsNil)

	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aid, combine.CiphertextC1, combine.CiphertextC2)
	c.Assert(err, qt.IsNil)
	partial.C2 = combine.CiphertextC2

	// Sealed: the partial is refused by the contract, not by node policy.
	err = helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, partial)
	ok, got := helpers.RevertsWith(err, "OrganizerSecretNotRevealed")
	c.Assert(ok, qt.IsTrue, qt.Commentf("partial before the reveal must revert OrganizerSecretNotRevealed, got %s", got))
	sealed, err := services.Manager.GetPartialDecryption(services.CallOpts(ctx), res.EpochID, aid, 1, assignedIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(sealed.Accepted, qt.IsFalse, qt.Commentf("no partial exists on chain before the reveal"))

	// And the combine has nothing to combine, whichever key it claims.
	c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aid, assignedIdx, combine), qt.IsNotNil,
		qt.Commentf("a combine before the reveal must revert"))

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.OrganizerSecret.Sign(), qt.Equals, 0, qt.Commentf("nothing is published before the reveal"))

	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, res.EpochID, aid, skOrg), qt.IsNil)
	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, partial), qt.IsNil,
		qt.Commentf("the very same partial lands once the secret is out"))

	// What the committee alone can prove: the same C1/C2 read as an
	// automatic ciphertext, i.e. plaintext m' = m + sk_org·r. The proof is
	// valid, the contract rejects it because the transcript's PK_org is the
	// identity and not the application's registered key.
	shifted := new(big.Int).Mul(skOrg, big.NewInt(base))
	shifted.Add(shifted, plaintext)
	shifted.Mod(shifted, group.ScalarField())
	without, err := helpers.BuildDecryptCombineOutputFromCiphertext(
		ctx, res.EpochID, aid, assignedIdx, 1, combine.CiphertextC1, combine.CiphertextC2,
		helpers.IdentityPoint(), nil, []uint16{1}, []types.CurvePoint{partial.Delta}, shifted,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aid, assignedIdx, without), qt.IsNotNil,
		qt.Commentf("a combine without the application's organizer key must revert"))

	c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aid, assignedIdx, combine), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, res.EpochID, aid, assignedIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Completed, qt.IsTrue)
	c.Assert(record.Plaintext.String(), qt.Equals, plaintext.String())
}

// TestThresholdDecryptionRespectsTheWindow asserts `decryptNotBefore` gates
// the partials, not just the combine: an application whose window has not
// opened leaks nothing, because `t` partials alone already decrypt an
// automatic ciphertext off chain.
func TestThresholdDecryptionRespectsTheWindow(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()
	aid := randomAid(c)

	now, err := helpers.ChainTimestamp(ctx, services)
	c.Assert(err, qt.IsNil)
	// Anvil's clock is the wall clock; keep the window short so the test
	// spends seconds, not minutes, waiting for it to open.
	const windowDelaySeconds = 12
	opensAt := now + windowDelaySeconds
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, aid, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic), DecryptNotBefore: opensAt},
	), qt.IsNil)

	share := poolShare(c, res, 0, 1)
	activation := res.Activation(0)

	partial, err := helpers.BuildPartialDecryptionSubmission(
		ctx, res.EpochID, aid, 1, 1, big.NewInt(9), share, big.NewInt(5), activation.Shares,
	)
	c.Assert(err, qt.IsNil)
	combine, err := helpers.BuildDecryptCombineOutput(ctx, res.EpochID, aid, 1, 1, big.NewInt(9), nil,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, big.NewInt(3))
	c.Assert(err, qt.IsNil)

	// The submission window is untouched by the decryption window.
	assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aid, combine.CiphertextC1, combine.CiphertextC2)
	c.Assert(err, qt.IsNil)
	partial.C2 = combine.CiphertextC2

	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, partial), qt.IsNotNil,
		qt.Commentf("submitPartialDecryption must revert with DecryptionNotOpen"))

	c.Assert(helpers.MineUntilTimestamp(ctx, services, opensAt), qt.IsNil)

	c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, assignedIdx, partial), qt.IsNil)
	c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aid, assignedIdx, combine), qt.IsNil)

	record, err := helpers.WaitCombinedDecryption(ctx, services, res.EpochID, aid, assignedIdx)
	c.Assert(err, qt.IsNil)
	c.Assert(record.Plaintext.String(), qt.Equals, "3")
}

// TestThresholdDecryptionSupportsMultipleCiphertextsPerRound runs several
// ciphertexts through one automatic application, checking the per-application
// index assignment along the way.
func TestThresholdDecryptionSupportsMultipleCiphertextsPerRound(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}

	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()
	aid := randomAid(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, aid, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)

	share := poolShare(c, res, 0, 1)
	activation := res.Activation(0)
	baseValues := []*big.Int{big.NewInt(9), big.NewInt(13)}
	plaintexts := []*big.Int{big.NewInt(3), big.NewInt(5)}

	for i := range baseValues {
		ciphertextIndex := uint16(i + 1)

		partial, err := helpers.BuildPartialDecryptionSubmission(
			ctx, res.EpochID, aid, ciphertextIndex, 1, baseValues[i], share, big.NewInt(int64(5+i)), activation.Shares,
		)
		c.Assert(err, qt.IsNil)

		combine, err := helpers.BuildDecryptCombineOutput(
			ctx, res.EpochID, aid, ciphertextIndex, 1, baseValues[i], nil,
			[]uint16{1}, []types.CurvePoint{partial.Delta}, plaintexts[i],
		)
		c.Assert(err, qt.IsNil)

		assignedIdx, err := helpers.SubmitCiphertextAs(ctx, self, res.EpochID, aid, combine.CiphertextC1, combine.CiphertextC2)
		c.Assert(err, qt.IsNil)
		c.Assert(assignedIdx, qt.Equals, ciphertextIndex)

		partial.C2 = combine.CiphertextC2
		c.Assert(helpers.SubmitPartialDecryptionAs(ctx, self, res.EpochID, aid, 1, ciphertextIndex, partial), qt.IsNil)
		c.Assert(helpers.CombineDecryptionAs(ctx, self, res.EpochID, aid, ciphertextIndex, combine), qt.IsNil)

		record, err := helpers.WaitCombinedDecryption(ctx, services, res.EpochID, aid, ciphertextIndex)
		c.Assert(err, qt.IsNil)
		c.Assert(record.Completed, qt.IsTrue)
		c.Assert(record.Plaintext.String(), qt.Equals, plaintexts[i].String())
	}
}

// poolShare is the member's share of one pool key of a finalized round.
func poolShare(c *qt.C, res *helpers.FinalizedRoundResult, keyIndex uint8, participantIndex uint16) *big.Int {
	shares, err := helpers.RecoverParticipantShares(res.Contributions, keyIndex, res.ParticipantIndexes)
	c.Assert(err, qt.IsNil)
	for i, index := range res.ParticipantIndexes {
		if index == participantIndex {
			return shares[i]
		}
	}
	c.Fatalf("participant %d is not part of the round", participantIndex)
	return nil
}
