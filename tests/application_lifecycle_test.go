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

// TestApplicationRegistration exercises `DKGAppManager.registerApplication`:
// bring an epoch to Live, register an application with an organizer key and
// a Schnorr proof of possession, and assert the cached on-chain record. The
// test fails if the Solidity-side `_organizerSchnorrChallenge` disagrees with
// the Go-side prover by even one byte.
//
// The test runs against the live anvil + node fixture; gated on
// RUN_INTEGRATION_TESTS=true so unit-test runs stay fast.
func TestApplicationRegistration(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)

	pkX, pkY, proof, err := schnorr.ProveOrganizerRegister(skOrg, res.EpochID, aid)
	c.Assert(err, qt.IsNil)

	policy := golangtypes.DKGTypesAppPolicy{
		AuthorizedSubmitter: services.TxManager.Address(),
		MaxCiphertexts:      0, // unlimited
		NotBeforeBlock:      0,
		NotAfterBlock:       0,
	}
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.AppManager.RegisterApplication(
		auth, res.EpochID, aid, policy, pkX, pkY, proof.Ax, proof.Ay, proof.Z,
	)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Exists, qt.IsTrue)
	c.Assert(app.Creator.Hex(), qt.Equals, services.TxManager.Address().Hex())
	c.Assert(app.OrganizerPK.X.Cmp(pkX), qt.Equals, 0,
		qt.Commentf("on-chain PK_org.x must match Go-side prover"))
	c.Assert(app.OrganizerPK.Y.Cmp(pkY), qt.Equals, 0,
		qt.Commentf("on-chain PK_org.y must match Go-side prover"))
	c.Assert(app.Policy.AuthorizedSubmitter.Hex(), qt.Equals, policy.AuthorizedSubmitter.Hex())

	// Identical re-registration must revert (aid is a one-shot binding).
	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	if _, err = services.AppManager.RegisterApplication(
		auth, res.EpochID, aid, policy, pkX, pkY, proof.Ax, proof.Ay, proof.Z,
	); err == nil {
		t.Fatalf("expected duplicate registerApplication to revert")
	}
}

// TestApplicationRegistrationResolvesZeroSubmitter asserts the contract stores
// the registering address when the policy leaves authorizedSubmitter zero:
// there is no open-submission mode.
func TestApplicationRegistrationResolvesZeroSubmitter(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)
	skOrg := randomOrganizerSecret(c)

	actor := selfActor()
	c.Assert(helpers.RegisterApplication(
		ctx, actor, services.AppManager, res.EpochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Policy.AuthorizedSubmitter.Hex(), qt.Equals, services.TxManager.Address().Hex())
}

// TestApplicationRegistrationRejectsTamperedProof flips the response scalar
// `z` and asserts that the on-chain Schnorr verifier rejects the proof. The
// test is the load-bearing assertion that the cross-impl transcript is wired
// correctly: a one-bit change in `z` breaks `z·G == A + c·PK` only if both
// sides agree on `c`.
func TestApplicationRegistrationRejectsTamperedProof(t *testing.T) {
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
	if _, err = services.AppManager.RegisterApplication(
		auth, res.EpochID, aid, policy,
		pkX, pkY, proof.Ax, proof.Ay, tampered,
	); err == nil {
		t.Fatalf("expected tampered Schnorr response to revert")
	}
}

// TestSubmitCiphertextRequiresRegisteredApplication asserts the contract
// refuses a ciphertext for an aid nobody registered — there is no epoch-key
// path any more.
func TestSubmitCiphertextRequiresRegisteredApplication(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	c1 := helpers.ScalarBasePoint(big.NewInt(9))
	c2 := helpers.ScalarBasePoint(big.NewInt(11))

	if _, err := helpers.SubmitCiphertextAs(ctx, selfActor(), res.EpochID, randomAid(c), c1, c2); err == nil {
		t.Fatalf("expected submitCiphertext for an unregistered aid to revert")
	}
	if _, err := helpers.SubmitCiphertextAs(ctx, selfActor(), res.EpochID, [32]byte{}, c1, c2); err == nil {
		t.Fatalf("expected submitCiphertext with aid = 0 to revert")
	}
}

// finalizedEpochForApps drives a single-participant epoch to Finalized so
// the application tests don't have to re-implement that flow. Built on
// top of the existing helper used by the SDK ciphertext-e2e test.
func finalizedEpochForApps(ctx context.Context, c *qt.C) *helpers.FinalizedRoundResult {
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
	res, err := helpers.CreateFinalizedSingleParticipantRound(
		ctx, services, policy, []*big.Int{big.NewInt(7)},
	)
	c.Assert(err, qt.IsNil)
	return res
}

// randomAid returns a fresh application id below the BN254 scalar field
// modulus (the contract rejects larger ids since proofs cannot bind them).
func randomAid(c *qt.C) [32]byte {
	var aid [32]byte
	_, err := rand.Read(aid[:])
	c.Assert(err, qt.IsNil)
	aid[0] &= 0x1f
	return aid
}

// randomOrganizerSecret returns a fresh non-zero organizer scalar.
func randomOrganizerSecret(c *qt.C) *big.Int {
	sk, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 250))
	c.Assert(err, qt.IsNil)
	return sk.Add(sk, big.NewInt(1))
}

// selfActor wraps the harness' own signer as a TestActor.
func selfActor() *helpers.TestActor {
	return &helpers.TestActor{
		Contracts: services.Contracts,
		Manager:   services.Manager,
		Registry:  services.Registry,
		TxManager: services.TxManager,
	}
}
