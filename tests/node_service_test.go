package tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/elgamal"
	"github.com/vocdoni/davinci-dkg/node"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// TestNodesServiceApplicationCiphertexts runs three real node instances
// against the harness chain and checks the whole production path: lottery
// claim, contribution, finalize, then partial decryption + combine for
// ciphertexts submitted under two registered applications. The test plays the
// organizer: it registers each application with its own sk_org, encrypts
// under PK_aid = PK_ep + PK_org and releases the organizer share. The nodes
// must not combine anything before that share is on chain.
func TestNodesServiceApplicationCiphertexts(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Minute)
	defer cancel()

	// ── nodes ─────────────────────────────────────────────────────────────
	for _, idx := range []int{3, 4, 5} {
		cfg := &node.Config{
			Web3:                  node.Web3Config{RPC: []string{services.RPCURL}, GasMultiplier: 1.2},
			PrivKey:               helpers.DefaultAnvilPrivateKeys[idx],
			ManagerAddr:           services.Addresses.Manager.Hex(),
			PollInterval:          time.Second,
			AutoCreateEpochs:      false,
			DecryptLookbackBlocks: 5,
		}
		n, err := node.New(cfg)
		c.Assert(err, qt.IsNil)
		c.Assert(n.EnsureRegistered(ctx), qt.IsNil)
		go n.Run(ctx, cfg)
	}

	// ── epoch: n=3, t=2 ───────────────────────────────────────────────────
	// createEpoch is cadence-gated; earlier tests may have just created one.
	c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
		next, err := services.Manager.NextEpochStartBlock(services.CallOpts(ctx))
		if err != nil {
			return false
		}
		head, err := services.Contracts.Client().BlockNumber(ctx)
		return err == nil && head >= next
	}), qt.IsNil)
	auth, err := services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err := services.Manager.CreateEpoch(auth, 2, 3, 3, helpers.DefaultLotteryAlphaBps)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	nonce, err := services.Manager.EpochNonce(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	prefix, err := services.Manager.EPOCHPREFIX(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	epochID := web3.EpochID(prefix, nonce)

	c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
		e, err := services.Contracts.GetEpoch(ctx, epochID)
		return err == nil && e.Status == 3
	}), qt.IsNil, qt.Commentf("nodes must claim, contribute and finalize on their own"))

	pkRaw, err := services.Manager.GetCollectivePublicKey(services.CallOpts(ctx), epochID)
	c.Assert(err, qt.IsNil)
	pkEp := types.CurvePoint{X: pkRaw.X, Y: pkRaw.Y}

	// ── two applications, each with its own organizer key ─────────────────
	self := selfActor()
	aidA, skOrgA := randomAid(c), randomOrganizerSecret(c)
	aidB, skOrgB := randomAid(c), randomOrganizerSecret(c)
	skOrgFor := map[[32]byte]*big.Int{aidA: skOrgA, aidB: skOrgB}
	pkFor := map[[32]byte]types.CurvePoint{}
	for aid, skOrg := range skOrgFor {
		c.Assert(helpers.RegisterApplication(
			ctx, self, services.AppManager, epochID, aid, skOrg, golangtypes.DKGTypesAppPolicy{},
		), qt.IsNil)
		rec, err := services.AppManager.GetApplication(services.CallOpts(ctx), epochID, aid)
		c.Assert(err, qt.IsNil)
		c.Assert(rec.Exists, qt.IsTrue)
		pkAid, err := elgamal.ApplicationKey(pkEp, types.CurvePoint{X: rec.OrganizerPK.X, Y: rec.OrganizerPK.Y})
		c.Assert(err, qt.IsNil)
		pkFor[aid] = pkAid
	}

	// ── ciphertexts ───────────────────────────────────────────────────────
	type want struct {
		aid [32]byte
		idx uint16
		m   *big.Int
	}
	wants := []want{
		{aidA, 1, big.NewInt(42)},
		{aidA, 2, big.NewInt(1_000_000)},
		{aidB, 1, big.NewInt(7)},
	}
	for _, w := range wants {
		c1, c2, err := elgamal.Encrypt(pkFor[w.aid], w.m)
		c.Assert(err, qt.IsNil)
		assigned, err := helpers.SubmitCiphertextAs(ctx, self, epochID, w.aid, c1, c2)
		c.Assert(err, qt.IsNil)
		c.Assert(assigned, qt.Equals, w.idx, qt.Commentf("indices are assigned sequentially per application"))

		// The organizer half. Until this lands the committee can only
		// recover sk_ep·C1, so no node may combine.
		_, _, err = helpers.SubmitOrganizerShareAs(
			ctx, self, services.AppManager, epochID, w.aid, assigned, c1, c2, skOrgFor[w.aid],
		)
		c.Assert(err, qt.IsNil)
	}

	// ── the committee must decrypt every one of them ──────────────────────
	for _, w := range wants {
		c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
			rec, err := services.Manager.GetCombinedDecryption(services.CallOpts(ctx), epochID, w.aid, w.idx)
			return err == nil && rec.Completed
		}), qt.IsNil, qt.Commentf("aid=%x idx=%d never combined", w.aid, w.idx))
		got, err := services.Manager.GetPlaintext(services.CallOpts(ctx), epochID, w.aid, w.idx)
		c.Assert(err, qt.IsNil)
		c.Assert(got.String(), qt.Equals, w.m.String())
	}
}
