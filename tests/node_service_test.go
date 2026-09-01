package tests

import (
	"context"
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
	"github.com/vocdoni/davinci-dkg/crypto/schnorr"
	"github.com/vocdoni/davinci-dkg/node"
	golangtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	"github.com/vocdoni/davinci-dkg/tests/helpers"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

// TestNodesServiceApplicationCiphertexts runs three real node instances
// against the harness chain and checks the whole production path: lottery
// claim, contribution, finalize, then partial decryption + combine for
// ciphertexts submitted under a mode-0 application, a mode-1 (organizer
// co-decryption) application and the legacy aid=0 path.
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
	tx, err := services.Manager.CreateEpoch(auth, 2, 3, 3, helpers.DefaultLotteryAlphaBps, golangtypes.DKGTypesDecryptionPolicy{})
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

	// ── application A: public derivation ──────────────────────────────────
	aidA := randomAid(c)
	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.AppManager.RegisterApplication(auth, epochID, aidA, golangtypes.DKGTypesAppPolicy{})
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	appA, err := services.AppManager.GetApplication(services.CallOpts(ctx), epochID, aidA)
	c.Assert(err, qt.IsNil)
	pkA, err := helpers.AddPoints(pkEp, helpers.ScalarBasePoint(appA.DerivationS))
	c.Assert(err, qt.IsNil)

	// ── application B: organizer co-decryption ────────────────────────────
	aidB := randomAid(c)
	skOrg, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 250))
	c.Assert(err, qt.IsNil)
	skOrg.Add(skOrg, big.NewInt(1))
	pkOrgX, pkOrgY, orgProof, err := schnorr.ProveOrganizerRegister(skOrg, epochID, aidB)
	c.Assert(err, qt.IsNil)
	auth, err = services.TxManager.NewTransactOpts(ctx)
	c.Assert(err, qt.IsNil)
	tx, err = services.AppManager.RegisterApplicationCoDec(auth, epochID, aidB, golangtypes.DKGTypesAppPolicy{},
		pkOrgX, pkOrgY, orgProof.Ax, orgProof.Ay, orgProof.Z)
	c.Assert(err, qt.IsNil)
	c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
	pkB, err := helpers.AddPoints(pkEp, types.CurvePoint{X: pkOrgX, Y: pkOrgY})
	c.Assert(err, qt.IsNil)

	// ── ciphertexts ───────────────────────────────────────────────────────
	submitter := &helpers.TestActor{Contracts: services.Contracts, Manager: services.Manager, Registry: services.Registry, TxManager: services.TxManager}
	type want struct {
		aid [32]byte
		idx uint16
		m   *big.Int
	}
	wants := []want{
		{aidA, 1, big.NewInt(42)},
		{aidA, 2, big.NewInt(1_000_000)},
		{aidB, 1, big.NewInt(7)},
		{[32]byte{}, 1, big.NewInt(5)},
	}
	pkFor := map[[32]byte]types.CurvePoint{aidA: pkA, aidB: pkB, {}: pkEp}
	for _, w := range wants {
		c1, c2, err := helpers.EncryptScalar(pkFor[w.aid], w.m)
		c.Assert(err, qt.IsNil)
		c.Assert(helpers.SubmitCiphertextAsApp(ctx, submitter, epochID, w.aid, w.idx, c1.X, c1.Y, c2.X, c2.Y), qt.IsNil)
		if w.aid == aidB {
			share, err := helpers.BuildOrganizerShareSubmission(ctx, epochID, aidB, w.idx, c1, skOrg)
			c.Assert(err, qt.IsNil)
			auth, err = services.TxManager.NewTransactOpts(ctx)
			c.Assert(err, qt.IsNil)
			tx, err = services.AppManager.SubmitOrganizerShare(auth, epochID, aidB, w.idx,
				c1.X, c1.Y, c2.X, c2.Y, share.Delta.X, share.Delta.Y, share.Proof, share.Input)
			c.Assert(err, qt.IsNil)
			c.Assert(services.TxManager.WaitTxByHash(tx.Hash(), helpers.DefaultTxTimeout), qt.IsNil)
		}
	}

	// ── the committee must decrypt every one of them ──────────────────────
	for _, w := range wants {
		w := w
		c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
			rec, err := services.Manager.GetCombinedDecryption(services.CallOpts(ctx), epochID, w.aid, w.idx)
			return err == nil && rec.Completed
		}), qt.IsNil, qt.Commentf("aid=%x idx=%d never combined", w.aid, w.idx))
		got, err := services.Manager.GetPlaintext(services.CallOpts(ctx), epochID, w.aid, w.idx)
		c.Assert(err, qt.IsNil)
		c.Assert(got.String(), qt.Equals, w.m.String())
	}
}
