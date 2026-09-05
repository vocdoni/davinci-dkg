package tests

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/ethereum/go-ethereum/common"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
)

// TestApplicationRegistration exercises `DKGAppManager.registerApplication`
// in organizer-locked mode: bring an epoch to Live with pool key 0
// activated, register an application with an organizer key and a Schnorr
// proof of possession, and assert the cached on-chain record — including the
// pool key it claimed. The test fails if the Solidity-side
// `_organizerSchnorrChallenge` disagrees with the Go-side prover by even one
// byte.
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

	pkX, pkY, _, err := schnorr.ProveOrganizerRegister(skOrg, res.EpochID, aid)
	c.Assert(err, qt.IsNil)

	policy := golangtypes.DKGTypesAppPolicy{
		Submitters:     []common.Address{services.TxManager.Address()},
		MaxCiphertexts: 0, // unlimited
	}
	c.Assert(helpers.RegisterApplication(ctx, selfActor(), services.AppManager, res.EpochID, aid, skOrg, policy), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Exists, qt.IsTrue)
	c.Assert(app.Creator.Hex(), qt.Equals, services.TxManager.Address().Hex())
	c.Assert(app.OrganizerPK.X.Cmp(pkX), qt.Equals, 0,
		qt.Commentf("on-chain PK_org.x must match Go-side prover"))
	c.Assert(app.OrganizerPK.Y.Cmp(pkY), qt.Equals, 0,
		qt.Commentf("on-chain PK_org.y must match Go-side prover"))
	c.Assert(app.Policy.Submitters, qt.DeepEquals, policy.Submitters)
	c.Assert(app.Policy.Mode, qt.Equals, uint8(types.AppModeOrganizerLocked))
	c.Assert(app.OrganizerSecret.Sign(), qt.Equals, 0, qt.Commentf("sk_org is only published by revealOrganizerSecret"))

	// The registration claimed the epoch's first (and only activated) key.
	c.Assert(app.PoolIndex, qt.Equals, uint8(0))
	poolIndex, err := services.Manager.GetAppPoolIndex(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(poolIndex, qt.Equals, uint8(0))

	status, err := services.Manager.GetPoolStatus(services.CallOpts(ctx), res.EpochID)
	c.Assert(err, qt.IsNil)
	c.Assert(status.NextIndex, qt.Equals, uint8(1), qt.Commentf("the pool cursor moved forward"))
	c.Assert(status.Activated, qt.Equals, uint8(1), qt.Commentf("only key 0 was activated"))

	poolX, poolY, err := services.Manager.GetPoolKey(services.CallOpts(ctx), res.EpochID, 0)
	c.Assert(err, qt.IsNil)
	activation := res.Activation(0)
	c.Assert(poolX.Cmp(activation.PoolKey.X), qt.Equals, 0)
	c.Assert(poolY.Cmp(activation.PoolKey.Y), qt.Equals, 0)

	root, err := services.Manager.GetPoolShareRoot(services.CallOpts(ctx), res.EpochID, 0)
	c.Assert(err, qt.IsNil)
	c.Assert(root, qt.Equals, activation.Shares.Root())

	// Identical re-registration must revert (aid is a one-shot binding).
	c.Assert(
		helpers.RegisterApplication(ctx, selfActor(), services.AppManager, res.EpochID, aid, skOrg, policy),
		qt.IsNotNil,
	)
}

// TestApplicationRegistrationAutomaticStoresIdentity registers an Automatic
// application: no organizer key at all, the contract stores the identity
// (0, 1) and the secret stays 0 forever.
func TestApplicationRegistrationAutomaticStoresIdentity(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)
	policy := golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)}

	c.Assert(helpers.RegisterApplication(ctx, selfActor(), services.AppManager, res.EpochID, aid, nil, policy), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Policy.Mode, qt.Equals, uint8(types.AppModeAutomatic))
	c.Assert(app.OrganizerPK.X.Sign(), qt.Equals, 0)
	c.Assert(app.OrganizerPK.Y.Cmp(big.NewInt(1)), qt.Equals, 0)
	c.Assert(app.OrganizerSecret.Sign(), qt.Equals, 0)
	c.Assert(app.PoolIndex, qt.Equals, uint8(0))
}

// TestApplicationRegistrationDefaultsToRegistrantOnly asserts an empty
// policy keeps submission closed: no allow-list, no open submission, so only
// the registrant may submit.
func TestApplicationRegistrationDefaultsToRegistrantOnly(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	aid := randomAid(c)

	c.Assert(helpers.RegisterApplication(
		ctx, selfActor(), services.AppManager, res.EpochID, aid, randomOrganizerSecret(c), golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.Creator.Hex(), qt.Equals, services.TxManager.Address().Hex())
	c.Assert(app.Policy.Submitters, qt.HasLen, 0)
	c.Assert(app.Policy.OpenSubmission, qt.IsFalse)
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

	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	_, err = services.AppManager.RegisterApplication(
		auth, res.EpochID, aid, golangtypes.DKGTypesAppPolicy{}, pkX, pkY, proof.Ax, proof.Ay, tampered,
	)
	c.Assert(err, qt.IsNotNil, qt.Commentf("tampered Schnorr response must revert"))
}

// TestApplicationRegistrationRequiresAnActivatedKey asserts the pool gate:
// the fixture epoch only has key 0 activated, so the second registration
// reverts with PoolKeyNotActive until another key is proven.
func TestApplicationRegistrationRequiresAnActivatedKey(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()

	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, randomAid(c), nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)

	// Key 1 is not activated yet: the claim reverts.
	blocked := randomAid(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, blocked, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNotNil, qt.Commentf("registration must revert with PoolKeyNotActive"))

	// Activate it and the very same registration goes through.
	_, err := helpers.ActivateRoundPoolKey(ctx, services, res, 1)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, blocked, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, blocked)
	c.Assert(err, qt.IsNil)
	c.Assert(app.PoolIndex, qt.Equals, uint8(1))
}

// TestApplicationRegistrationExhaustsThePool activates every key of one
// epoch, registers MaxK applications and asserts the next one reverts with
// PoolExhausted. This is the epoch's hard capacity: after it, the only way
// to register is a new epoch.
func TestApplicationRegistrationExhaustsThePool(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), helpers.MaxTestTimeout(t))
	defer cancel()

	res := finalizedEpochForApps(ctx, c)
	self := selfActor()
	policy := golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)}

	for keyIndex := 1; keyIndex < ccommon.MaxK; keyIndex++ {
		_, err := helpers.ActivateRoundPoolKey(ctx, services, res, uint8(keyIndex))
		c.Assert(err, qt.IsNil)
	}
	status, err := services.Manager.GetPoolStatus(services.CallOpts(ctx), res.EpochID)
	c.Assert(err, qt.IsNil)
	c.Assert(status.Activated, qt.Equals, uint8(0xff), qt.Commentf("every key of the pool is activated"))

	for i := range ccommon.MaxK {
		aid := randomAid(c)
		c.Assert(helpers.RegisterApplication(ctx, self, services.AppManager, res.EpochID, aid, nil, policy), qt.IsNil)
		app, appErr := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
		c.Assert(appErr, qt.IsNil)
		c.Assert(app.PoolIndex, qt.Equals, uint8(i), qt.Commentf("keys are claimed in order"))
	}

	status, err = services.Manager.GetPoolStatus(services.CallOpts(ctx), res.EpochID)
	c.Assert(err, qt.IsNil)
	c.Assert(status.NextIndex, qt.Equals, uint8(ccommon.MaxK))

	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, randomAid(c), nil, policy,
	), qt.IsNotNil, qt.Commentf("the %d+1-th registration must revert with PoolExhausted", ccommon.MaxK))
}

// TestRevealOrganizerSecret covers the one-shot reveal: a secret that does
// not match the registered PK_org is refused, the right one lands, and a
// second reveal (or a reveal on an automatic application) reverts.
func TestRevealOrganizerSecret(t *testing.T) {
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

	wrong := new(big.Int).Add(skOrg, big.NewInt(1))
	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, res.EpochID, aid, wrong), qt.IsNotNil,
		qt.Commentf("a secret whose sk·G is not PK_org must revert"))
	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, res.EpochID, aid, big.NewInt(0)), qt.IsNotNil,
		qt.Commentf("zero is not a valid organizer secret"))

	// The reveal is permissionless: a second actor can publish it.
	stranger, err := services.Actor(1)
	c.Assert(err, qt.IsNil)
	c.Assert(helpers.RevealOrganizerSecretAs(ctx, stranger, services.AppManager, res.EpochID, aid, skOrg), qt.IsNil)

	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), res.EpochID, aid)
	c.Assert(err, qt.IsNil)
	c.Assert(app.OrganizerSecret.Cmp(skOrg), qt.Equals, 0)

	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, res.EpochID, aid, skOrg), qt.IsNotNil,
		qt.Commentf("a second reveal must revert with AlreadyRevealed"))

	// An automatic application has nothing to reveal.
	_, err = helpers.ActivateRoundPoolKey(ctx, services, res, 1)
	c.Assert(err, qt.IsNil)
	auto := randomAid(c)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, res.EpochID, auto, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)
	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, res.EpochID, auto, skOrg), qt.IsNotNil,
		qt.Commentf("automatic applications reject revealOrganizerSecret"))
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

	_, err := helpers.SubmitCiphertextAs(ctx, selfActor(), res.EpochID, randomAid(c), c1, c2)
	c.Assert(err, qt.IsNotNil, qt.Commentf("submitCiphertext for an unregistered aid must revert"))
	_, err = helpers.SubmitCiphertextAs(ctx, selfActor(), res.EpochID, [32]byte{}, c1, c2)
	c.Assert(err, qt.IsNotNil, qt.Commentf("submitCiphertext with aid = 0 must revert"))
}

// finalizedEpochForApps drives a single-participant epoch to Live with pool
// key 0 activated, so the application tests don't have to re-implement that
// flow. Built on top of the existing helper used by the SDK e2e tests.
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
	return helpers.SelfActor(services)
}
