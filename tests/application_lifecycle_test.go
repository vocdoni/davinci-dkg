package tests

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestApplicationRegistration_Mode0 exercises the mode-0 (public derivation)
// branch of `DKGManager.registerApplication`: bring an epoch to Finalized,
// register an application against `(epochId, aid)`, and assert that the
// cached on-chain record carries the contract-derived `S` and the calling
// account as creator. Mirrors paper §4.3.
//
// The test runs against the live anvil + 7-node fixture; gated on
// RUN_INTEGRATION_TESTS=true so unit-test runs stay fast.
func TestApplicationRegistration_Mode0(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)

	policy := golangtypes.DKGTypesAppPolicy{
		AuthorizedSubmitter: services.TxManager.Address(), // restrict to organizer
		MaxCiphertexts:      0,                            // unlimited
		NotBeforeBlock:      0,
		NotAfterBlock:       0,
	}
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.RegisterApplication(auth, res.EpochID, aid, policy)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	app, err := services.Manager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Exists, qt.IsTrue)
	c.Assert(app.Mode, qt.Equals, uint8(0))
	c.Assert(app.DerivationS.Sign() > 0, qt.IsTrue,
		qt.Commentf("S should be non-zero for any non-trivial (eid, aid)"))
	c.Assert(app.Creator.Hex(), qt.Equals, services.TxManager.Address().Hex())
	c.Assert(app.Policy.AuthorizedSubmitter.Hex(), qt.Equals,
		policy.AuthorizedSubmitter.Hex())

	// Identical re-registration must revert (aid is a one-shot binding).
	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	if _, err = services.Manager.RegisterApplication(auth, res.EpochID, aid, policy); err == nil {
		t.Fatalf("expected duplicate registerApplication to revert")
	}
}

// TestApplicationRegistration_Mode1 exercises the mode-1 (organizer
// co-decryption) branch. Builds a Schnorr proof of knowledge of `sk_org`
// via `crypto/schnorr` and verifies that the on-chain transcript matches
// (the test fails if the Solidity-side `_organizerSchnorrChallenge`
// disagrees with the Go-side prover by even one byte).
func TestApplicationRegistration_Mode1(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)

	skOrg, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<62))
	c.Assert(err, qt.IsNil)
	skOrg.Add(skOrg, big.NewInt(1)) // ensure non-zero

	pkX, pkY, proof, err := schnorr.ProveOrganizerRegister(skOrg, res.EpochID, aid)
	c.Assert(err, qt.IsNil)

	policy := golangtypes.DKGTypesAppPolicy{
		AuthorizedSubmitter: services.TxManager.Address(),
		MaxCiphertexts:      0,
		NotBeforeBlock:      0,
		NotAfterBlock:       0,
	}
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.RegisterApplicationCoDec(
		auth, res.EpochID, aid, policy,
		pkX, pkY, proof.Ax, proof.Ay, proof.Z,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	app, err := services.Manager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Exists, qt.IsTrue)
	c.Assert(app.Mode, qt.Equals, uint8(1))
	c.Assert(app.OrganizerPK.X.Cmp(pkX), qt.Equals, 0,
		qt.Commentf("on-chain PK_org.x must match Go-side prover"))
	c.Assert(app.OrganizerPK.Y.Cmp(pkY), qt.Equals, 0,
		qt.Commentf("on-chain PK_org.y must match Go-side prover"))
	c.Assert(app.DerivationS.Sign(), qt.Equals, 0,
		qt.Commentf("mode 1 must store S=0; correction comes from PK_org"))
}

// TestApplicationRegistration_Mode1_RejectsTamperedProof flips the response
// scalar `z` and asserts that the on-chain Schnorr verifier rejects the
// proof. The test is the load-bearing assertion that the cross-impl
// Poseidon transcript is wired correctly: a one-bit change in `z` breaks
// `z·G == A + c·PK` only if both sides agree on `c`.
func TestApplicationRegistration_Mode1_RejectsTamperedProof(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)

	skOrg := big.NewInt(1234567890)
	pkX, pkY, proof, err := schnorr.ProveOrganizerRegister(skOrg, res.EpochID, aid)
	c.Assert(err, qt.IsNil)

	tampered := new(big.Int).Add(proof.Z, big.NewInt(1))

	policy := golangtypes.DKGTypesAppPolicy{}
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	if _, err = services.Manager.RegisterApplicationCoDec(
		auth, res.EpochID, aid, policy,
		pkX, pkY, proof.Ax, proof.Ay, tampered,
	); err == nil {
		t.Fatalf("expected tampered Schnorr response to revert")
	}
}

// finalizedEpochForApps drives a single-participant epoch to Finalized so
// the application tests don't have to re-implement that flow. Built on
// top of the existing helper used by the SDK ciphertext-e2e test.
func finalizedEpochForApps(ctx context.Context, c *qt.C) *helpers.FinalizedRoundResult {
	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)

	policy := types.EpochPolicy{
		Threshold:                 1,
		CommitteeSize:             1,
		MinValidContributions:     1,
		LotteryAlphaBps:           helpers.DefaultLotteryAlphaBps,
		SeedDelay:                 helpers.DefaultSeedDelay,
		RegistrationDeadlineBlock: head + 25,
		ContributionDeadlineBlock: head + 50,
		FinalizeNotBeforeBlock:    head + 51,
	}
	res, err := helpers.CreateFinalizedSingleParticipantRound(
		ctx, services, policy, []*big.Int{big.NewInt(7)},
	)
	c.Assert(err, qt.IsNil)
	return res
}

func randomAid(c *qt.C) [32]byte {
	var aid [32]byte
	_, err := rand.Read(aid[:])
	c.Assert(err, qt.IsNil)
	return aid
}
