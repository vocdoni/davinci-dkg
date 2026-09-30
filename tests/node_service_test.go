package tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
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
// claim, contribution, the proof-carrying finalize that stores every pool
// key, then partial decryption + combine for ciphertexts submitted under two
// applications. One is automatic — the committee owns it end to end — and
// one is organizer-locked: the contract refuses every partial of it until
// the organizer calls revealOrganizerSecret, so the nodes park its slots and
// nothing of it exists on chain before the reveal half way through the test.
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
	live, err := services.Contracts.GetEpoch(ctx, epochID)
	c.Assert(err, qt.IsNil)
	c.Assert(live.ClaimedCount, qt.Equals, uint16(3))
	c.Assert(live.ContributionCount, qt.Equals, uint16(3), qt.Commentf("every member deals"))

	// Live means every pool key is stored: the two registrations below claim
	// keys 0 and 1 without any further node work.
	next, err := services.Manager.GetPoolStatus(services.CallOpts(ctx), epochID)
	c.Assert(err, qt.IsNil)
	c.Assert(next, qt.Equals, uint8(0))

	// ── two applications: one automatic, one organizer-locked ─────────────
	self := selfActor()
	aidAuto := randomAid(c)
	aidLocked := randomAid(c)
	skOrg := randomOrganizerSecret(c)

	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, epochID, aidAuto, nil,
		golangtypes.DKGTypesAppPolicy{Mode: uint8(types.AppModeAutomatic)},
	), qt.IsNil)
	c.Assert(helpers.RegisterApplication(
		ctx, self, services.AppManager, epochID, aidLocked, skOrg, golangtypes.DKGTypesAppPolicy{},
	), qt.IsNil)

	pkFor := map[[32]byte]types.CurvePoint{}
	for _, aid := range [][32]byte{aidAuto, aidLocked} {
		rec, recErr := services.AppManager.GetApplication(services.CallOpts(ctx), epochID, aid)
		c.Assert(recErr, qt.IsNil)
		c.Assert(rec.Exists, qt.IsTrue)
		x, y, keyErr := services.Manager.GetPoolKey(services.CallOpts(ctx), epochID, rec.PoolIndex)
		c.Assert(keyErr, qt.IsNil)
		pkAid, keyErr := elgamal.ApplicationKey(
			types.CurvePoint{X: x, Y: y},
			types.CurvePoint{X: rec.OrganizerPK.X, Y: rec.OrganizerPK.Y},
		)
		c.Assert(keyErr, qt.IsNil)
		pkFor[aid] = pkAid
	}

	// ── ciphertexts ───────────────────────────────────────────────────────
	type want struct {
		aid [32]byte
		idx uint16
		m   *big.Int
	}
	wants := []want{
		{aidAuto, 1, big.NewInt(42)},
		{aidAuto, 2, big.NewInt(1_000_000)},
		{aidLocked, 1, big.NewInt(7)},
		{aidLocked, 2, big.NewInt(123_456)},
	}
	for _, w := range wants {
		c1, c2, encErr := elgamal.Encrypt(pkFor[w.aid], w.m)
		c.Assert(encErr, qt.IsNil)
		assigned, subErr := helpers.SubmitCiphertextAs(ctx, self, epochID, w.aid, c1, c2)
		c.Assert(subErr, qt.IsNil)
		c.Assert(assigned, qt.Equals, w.idx, qt.Commentf("indices are assigned sequentially per application"))
	}

	// ── the automatic application decrypts on its own ─────────────────────
	for _, w := range wants {
		if w.aid != aidAuto {
			continue
		}
		waitCombined(ctx, c, epochID, w.aid, w.idx, w.m)
	}

	// ── the locked one stays sealed until the secret is out ───────────────
	// The automatic combines above took the nodes through many ticks with
	// the locked ciphertexts pending; not one partial of them may be on
	// chain, let alone a combine — the contract reverts
	// OrganizerSecretNotRevealed and the nodes park the slots instead.
	for _, w := range wants {
		if w.aid != aidLocked {
			continue
		}
		for member := uint16(1); member <= 3; member++ {
			partial, partialErr := services.Manager.GetPartialDecryption(services.CallOpts(ctx), epochID, w.aid, member, w.idx)
			c.Assert(partialErr, qt.IsNil)
			c.Assert(partial.Accepted, qt.IsFalse,
				qt.Commentf("aid=%x idx=%d member %d posted a partial before the reveal", w.aid, w.idx, member))
		}
		rec, recErr := services.Manager.GetCombinedDecryption(services.CallOpts(ctx), epochID, w.aid, w.idx)
		c.Assert(recErr, qt.IsNil)
		c.Assert(rec.Completed, qt.IsFalse,
			qt.Commentf("aid=%x idx=%d combined without the organizer secret", w.aid, w.idx))
	}

	// The reveal wakes the parked slots; partials and combines follow.
	c.Assert(helpers.RevealOrganizerSecretAs(ctx, self, services.AppManager, epochID, aidLocked, skOrg), qt.IsNil)

	for _, w := range wants {
		if w.aid != aidLocked {
			continue
		}
		waitCombined(ctx, c, epochID, w.aid, w.idx, w.m)
	}
}

// waitCombined blocks until the nodes combined (epochID, aid, idx) and
// asserts the recovered plaintext.
func waitCombined(ctx context.Context, c *qt.C, epochID [12]byte, aid [32]byte, idx uint16, want *big.Int) {
	c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
		rec, err := services.Manager.GetCombinedDecryption(services.CallOpts(ctx), epochID, aid, idx)
		return err == nil && rec.Completed
	}), qt.IsNil, qt.Commentf("aid=%x idx=%d never combined", aid, idx))
	got, err := services.Manager.GetPlaintext(services.CallOpts(ctx), epochID, aid, idx)
	c.Assert(err, qt.IsNil)
	c.Assert(got.String(), qt.Equals, want.String())
}

// TestNodeAbortsADeadEpochAndCreatesTheNext runs one real node with
// auto-create on against an epoch whose committee can never fill: two seats
// and one running operator. Before the cadence createEpoch only goes through
// behind a spent or Aborted newest epoch, so once the selection deadline has
// passed the node must abort the dead epoch itself and then create the next
// one, long before the cadence.
func TestNodeAbortsADeadEpochAndCreatesTheNext(t *testing.T) {
	if !helpers.IsIntegrationEnabled() {
		t.Skip("integration tests disabled")
	}
	c := qt.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Two seats keep every epoch in this test short of a committee, the ones
	// the node creates included, so the node never has to deal a share.
	cfg := &node.Config{
		Web3:                  node.Web3Config{RPC: []string{services.RPCURL}, GasMultiplier: 1.2},
		PrivKey:               helpers.DefaultAnvilPrivateKeys[3],
		ManagerAddr:           services.Addresses.Manager.Hex(),
		PollInterval:          time.Second,
		AutoCreateEpochs:      true,
		AutoCreateJitter:      time.Second,
		DecryptLookbackBlocks: 5,
		EpochPolicy: node.EpochPolicyConfig{
			Threshold: 1, CommitteeSize: 2, MinValidContributions: 1, LotteryAlphaBps: helpers.DefaultLotteryAlphaBps,
		},
	}
	n, err := node.New(cfg)
	c.Assert(err, qt.IsNil)
	c.Assert(n.EnsureRegistered(ctx), qt.IsNil)
	operator, err := services.ActorFromPrivateKey(cfg.PrivKey)
	c.Assert(err, qt.IsNil)
	nodeAddr := operator.Address()

	// The dead epoch is created at the cadence, before the node runs, so the
	// node has no cadence-driven createEpoch of its own to fire.
	dead, err := helpers.CreateEpoch(ctx, services, types.EpochPolicy{
		Threshold: 1, CommitteeSize: 2, MinValidContributions: 1, LotteryAlphaBps: helpers.DefaultLotteryAlphaBps,
	})
	c.Assert(err, qt.IsNil)
	epoch, err := services.Contracts.GetEpoch(ctx, dead)
	c.Assert(err, qt.IsNil)
	cadence, err := services.Manager.NextEpochStartBlock(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)

	nodeCtx, stop := context.WithCancel(ctx)
	done := make(chan struct{})
	go func() {
		n.Run(nodeCtx, cfg)
		close(done)
	}()
	defer func() {
		stop()
		<-done
		// Let a transaction sent just before the stop land here rather than
		// in the next test.
		_ = helpers.MineBlocks(ctx, services, 2)
		_ = helpers.WaitUntilCondition(ctx, time.Second, func() bool {
			client := services.Contracts.Pool().Current()
			pending, perr := client.PendingNonceAt(ctx, nodeAddr)
			mined, merr := client.NonceAt(ctx, nodeAddr, nil)
			return perr == nil && merr == nil && pending == mined
		})
	}()

	head, err := services.Contracts.Client().BlockNumber(ctx)
	c.Assert(err, qt.IsNil)
	if deadline := epoch.Policy.CommitteeSelectionDeadlineBlock; head <= deadline {
		c.Assert(helpers.MineBlocks(ctx, services, deadline-head+1), qt.IsNil)
	}

	c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
		e, err := services.Contracts.GetEpoch(ctx, dead)
		return err == nil && e.Status == 4 // Aborted
	}), qt.IsNil, qt.Commentf("the node never aborted the dead epoch"))
	it, err := services.Manager.FilterEpochAborted(&bind.FilterOpts{Start: epoch.StartBlock, Context: ctx}, [][12]byte{dead})
	c.Assert(err, qt.IsNil)
	c.Assert(it.Next(), qt.IsTrue)
	abortTx, _, err := services.Contracts.Client().TransactionByHash(ctx, it.Event.Raw.TxHash)
	c.Assert(err, qt.IsNil)
	sender, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(abortTx.ChainId()), abortTx)
	c.Assert(err, qt.IsNil)
	c.Assert(sender, qt.Equals, nodeAddr)
	c.Assert(it.Close(), qt.IsNil)

	prefix, err := services.Manager.EPOCHPREFIX(services.CallOpts(ctx))
	c.Assert(err, qt.IsNil)
	next := web3.EpochID(prefix, epoch.Nonce+1)
	var created web3.EpochView
	c.Assert(helpers.WaitUntilCondition(ctx, time.Second, func() bool {
		created, err = services.Contracts.GetEpoch(ctx, next)
		return err == nil && created.Status != 0
	}), qt.IsNil, qt.Commentf("the node never created the epoch after the aborted one"))
	c.Assert(created.Organizer, qt.Equals, nodeAddr)
	c.Assert(created.StartBlock < cadence, qt.IsTrue,
		qt.Commentf("epoch %d started at %d, the cadence was %d", epoch.Nonce+1, created.StartBlock, cadence))
}
