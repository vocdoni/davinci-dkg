package txmanager

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	qt "github.com/frankban/quicktest"
)

const (
	opChain      = uint64(10)
	gnosisChain  = uint64(100)
	gwei         = int64(1_000_000_000)
	opFallback   = int64(1_000_000)
	opFeeCapWei  = int64(20_000_000) // 0.02 gwei
	unknownChain = testChainID
)

func newChainManager(t *testing.T, f *fakeRPC, chainID uint64) *Manager {
	t.Helper()
	client, err := ethclient.Dial(f.srv.URL)
	qt.Assert(t, err, qt.IsNil)
	t.Cleanup(client.Close)
	m, err := New(func() *ethclient.Client { return client }, chainID, testKey)
	qt.Assert(t, err, qt.IsNil)
	return m
}

func headerJSON(baseFee *big.Int) any {
	h := map[string]any{
		"parentHash":       common.Hash{},
		"sha3Uncles":       gethtypes.EmptyUncleHash,
		"miner":            common.Address{},
		"stateRoot":        common.Hash{},
		"transactionsRoot": common.Hash{},
		"receiptsRoot":     common.Hash{},
		"logsBloom":        hexutil.Bytes(make([]byte, 256)),
		"difficulty":       "0x0",
		"number":           "0x1",
		"gasLimit":         "0x1c9c380",
		"gasUsed":          "0x0",
		"timestamp":        "0x1",
		"extraData":        "0x",
	}
	if baseFee != nil {
		h["baseFeePerGas"] = hexutil.EncodeBig(baseFee)
	}
	return h
}

// serveFees makes the fake RPC answer the two fee reads; a nil tip fails
// eth_maxPriorityFeePerGas.
func serveFees(f *fakeRPC, tip, baseFee *big.Int) {
	f.handle("eth_maxPriorityFeePerGas", func([]json.RawMessage) (any, error) {
		if tip == nil {
			return nil, fmt.Errorf("method not supported")
		}
		return hexutil.EncodeBig(tip), nil
	})
	f.handle("eth_getBlockByNumber", func([]json.RawMessage) (any, error) { return headerJSON(baseFee), nil })
}

func TestSuggestFeesOPFallbackTip(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	serveFees(f, nil, big.NewInt(5_000))
	m := newChainManager(t, f, opChain)

	tip, feeCap, err := m.suggestFees(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(tip.Int64(), qt.Equals, opFallback)
	c.Assert(feeCap.Int64(), qt.Equals, 2*5_000+opFallback)
}

func TestSuggestFeesOPNoBaseFee(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	serveFees(f, nil, nil)
	m := newChainManager(t, f, opChain)

	tip, feeCap, err := m.suggestFees(context.Background())
	c.Assert(err, qt.IsNil)
	// tip + 1 gwei would be 1.001 gwei; the cap wins.
	c.Assert(feeCap.Int64(), qt.Equals, opFeeCapWei)
	c.Assert(tip.Int64(), qt.Equals, opFallback)
}

func TestSuggestFeesOPClampsHugeValues(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	serveFees(f, big.NewInt(50*gwei), big.NewInt(100*gwei))
	m := newChainManager(t, f, opChain)

	tip, feeCap, err := m.suggestFees(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(feeCap.Int64(), qt.Equals, opFeeCapWei)
	c.Assert(tip.Cmp(feeCap) <= 0, qt.IsTrue, qt.Commentf("tip %v above feeCap %v", tip, feeCap))
	c.Assert(tip.Int64(), qt.Equals, opFeeCapWei)
}

func TestSuggestFeesKeepsOPNormalValues(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	serveFees(f, big.NewInt(opFallback), big.NewInt(5_000))
	m := newChainManager(t, f, opChain)

	tip, feeCap, err := m.suggestFees(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(tip.Int64(), qt.Equals, opFallback)
	c.Assert(feeCap.Int64(), qt.Equals, 2*5_000+opFallback)
}

// Gnosis and unknown chains keep the generic behaviour: 1 gwei fallback,
// no cap.
func TestSuggestFeesUncappedChainsUnchanged(t *testing.T) {
	for name, id := range map[string]uint64{"gnosis": gnosisChain, "unknown": unknownChain} {
		t.Run(name, func(t *testing.T) {
			c := qt.New(t)
			f := newFakeRPC(t)
			serveFees(f, nil, big.NewInt(7*gwei))
			m := newChainManager(t, f, id)
			tip, feeCap, err := m.suggestFees(context.Background())
			c.Assert(err, qt.IsNil)
			c.Assert(tip.Int64(), qt.Equals, gwei)
			c.Assert(feeCap.Int64(), qt.Equals, 14*gwei+gwei)

			serveFees(f, nil, nil)
			tip, feeCap, err = m.suggestFees(context.Background())
			c.Assert(err, qt.IsNil)
			c.Assert(tip.Int64(), qt.Equals, gwei)
			c.Assert(feeCap.Int64(), qt.Equals, 2*gwei)

			serveFees(f, big.NewInt(500*gwei), big.NewInt(1000*gwei))
			tip, feeCap, err = m.suggestFees(context.Background())
			c.Assert(err, qt.IsNil)
			c.Assert(tip.Int64(), qt.Equals, 500*gwei)
			c.Assert(feeCap.Int64(), qt.Equals, 2500*gwei)
		})
	}
}

// When fee estimation fails outright a capped chain must not leave it to
// bind (uncapped): the fallback is min(tip, cap) with feeCap = cap.
func TestFallbackFeesUnderCap(t *testing.T) {
	c := qt.New(t)
	m := newChainManager(t, newFakeRPC(t), opChain)
	tip, feeCap, ok := m.fallbackFees()
	c.Assert(ok, qt.IsTrue)
	c.Assert(tip.Int64(), qt.Equals, opFallback)
	c.Assert(feeCap.Int64(), qt.Equals, opFeeCapWei)

	m.config.Chain.FallbackTipCap = big.NewInt(gwei) // above the cap
	tip, feeCap, ok = m.fallbackFees()
	c.Assert(ok, qt.IsTrue)
	c.Assert(tip.Int64(), qt.Equals, opFeeCapWei)
	c.Assert(feeCap.Int64(), qt.Equals, opFeeCapWei)

	_, _, ok = newChainManager(t, newFakeRPC(t), gnosisChain).fallbackFees()
	c.Assert(ok, qt.IsFalse, qt.Commentf("uncapped chains keep bind's own estimate"))
}

func TestNewTransactOptsSetsCappedFees(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	serveFees(f, big.NewInt(50*gwei), big.NewInt(100*gwei))
	m := newChainManager(t, f, opChain)
	auth, err := m.NewTransactOpts(context.Background())
	c.Assert(err, qt.IsNil)
	c.Assert(auth.GasFeeCap.Int64(), qt.Equals, opFeeCapWei)
	c.Assert(auth.GasTipCap.Cmp(auth.GasFeeCap) <= 0, qt.IsTrue)
}

func capTx(t *testing.T, m *Manager, tip, feeCap int64) *gethtypes.Transaction {
	t.Helper()
	to := common.HexToAddress("0xdead")
	signed, err := gethtypes.SignTx(gethtypes.NewTx(&gethtypes.DynamicFeeTx{
		ChainID: m.chainID, Nonce: 5, GasTipCap: big.NewInt(tip), GasFeeCap: big.NewInt(feeCap),
		Gas: 21_000, To: &to,
	}), gethtypes.LatestSignerForChainID(m.chainID), m.key)
	qt.Assert(t, err, qt.IsNil)
	return signed
}

func TestBumpedCopyStopsAtTheCap(t *testing.T) {
	c := qt.New(t)
	m := newChainManager(t, newFakeRPC(t), opChain)

	// 50 % bumps compound until both values reach the cap, then stop.
	tx := capTx(t, m, 1_000_000, 10_000_000)
	var last *gethtypes.Transaction
	for i := 0; i < 20; i++ {
		next, err := m.bumpedCopy(tx)
		c.Assert(err, qt.IsNil)
		c.Assert(next.GasFeeCap().Int64() <= opFeeCapWei, qt.IsTrue)
		c.Assert(next.GasTipCap().Cmp(next.GasFeeCap()) <= 0, qt.IsTrue)
		if next.Hash() == tx.Hash() {
			last = next
			break
		}
		tx = next
	}
	c.Assert(last, qt.IsNotNil, qt.Commentf("bumps never reached a fixed point"))
	c.Assert(last.GasFeeCap().Int64(), qt.Equals, opFeeCapWei)
}

func TestBumpedCopyPartialBumpClamps(t *testing.T) {
	c := qt.New(t)
	m := newChainManager(t, newFakeRPC(t), opChain)
	tx := capTx(t, m, 1_000_000, 16_000_000)
	next, err := m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.GasFeeCap().Int64(), qt.Equals, opFeeCapWei) // 24M clamped
	c.Assert(next.GasTipCap().Int64(), qt.Equals, int64(1_500_000))
	c.Assert(next.Hash() != tx.Hash(), qt.IsTrue)
}

func TestBumpedCopyUncappedChainsUnchanged(t *testing.T) {
	for _, id := range []uint64{gnosisChain, unknownChain} {
		c := qt.New(t)
		m := newChainManager(t, newFakeRPC(t), id)
		tx := capTx(t, m, 10*gwei, 1000*gwei)
		for i := 0; i < 3; i++ {
			next, err := m.bumpedCopy(tx)
			c.Assert(err, qt.IsNil)
			c.Assert(next.GasFeeCap().Cmp(new(big.Int).Div(new(big.Int).Mul(tx.GasFeeCap(), big.NewInt(150)), big.NewInt(100))), qt.Equals, 0)
			c.Assert(next.GasTipCap().Cmp(new(big.Int).Div(new(big.Int).Mul(tx.GasTipCap(), big.NewInt(150)), big.NewInt(100))), qt.Equals, 0)
			tx = next
		}
	}
}

// At the cap the stuck transaction is re-broadcast as is: same hash, the
// attempt counted, the pending clock restarted, the warning logged once.
func TestRetryStuckAtCapRebroadcastsOriginal(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newChainManager(t, f, opChain)
	tx := capTx(t, m, opFeeCapWei, opFeeCapWei)
	m.RecordPending(tx)

	f.handle("eth_getTransactionCount", func(params []json.RawMessage) (any, error) { return "0x5", nil })
	var sent []common.Hash
	f.handle("eth_sendRawTransaction", func(params []json.RawMessage) (any, error) {
		var raw hexutil.Bytes
		if err := json.Unmarshal(params[0], &raw); err != nil {
			return nil, err
		}
		var d gethtypes.Transaction
		if err := d.UnmarshalBinary(raw); err != nil {
			return nil, err
		}
		sent = append(sent, d.Hash())
		return d.Hash(), nil
	})

	for i := 1; i <= 2; i++ {
		m.pending[5].submittedAt = time.Now().Add(-m.config.MaxPendingTime - time.Second)
		// Present in the mempool so only the age path runs.
		f.handle("eth_getTransactionCount", func(params []json.RawMessage) (any, error) {
			var tag string
			_ = json.Unmarshal(params[1], &tag)
			if tag == tagPending {
				return "0x6", nil
			}
			return "0x5", nil
		})
		c.Assert(m.retryStuck(context.Background()), qt.IsNil)
		c.Assert(sent, qt.HasLen, i)
		c.Assert(sent[i-1], qt.Equals, tx.Hash())
		c.Assert(m.pending[5].hash, qt.Equals, tx.Hash())
		c.Assert(m.pending[5].retries, qt.Equals, i)
		c.Assert(time.Since(m.pending[5].submittedAt) < time.Minute, qt.IsTrue, qt.Commentf("clock restarts"))
		c.Assert(m.pending[5].capLogged, qt.IsTrue)
	}
}

// geth only accepts a same-nonce replacement with tip and feeCap both at
// least 10% higher; a clamp that leaves less is treated as at the cap.
func TestBumpedCopyNeedsTenPercentOverTheOriginal(t *testing.T) {
	c := qt.New(t)
	m := newChainManager(t, newFakeRPC(t), opChain)

	// 19M -> 28.5M -> clamped 20M is +5.3%: geth would refuse it.
	tx := capTx(t, m, 1_000_000, 19_000_000)
	next, err := m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.Hash(), qt.Equals, tx.Hash())

	// 18M -> 20M is +11.1%: accepted.
	tx = capTx(t, m, 1_000_000, 18_000_000)
	next, err = m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.Hash() != tx.Hash(), qt.IsTrue)
	c.Assert(next.GasFeeCap().Int64(), qt.Equals, opFeeCapWei)

	// A cap-aligned pair: tip 18M of cap 20M clamps feeCap to 20M and needs
	// tip >= 19.8M, which the 50% bump of 18M (27M -> 20M) meets.
	tx = capTx(t, m, 18_000_000, 18_000_000)
	next, err = m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.Hash() != tx.Hash(), qt.IsTrue)

	// Just under 10% on the fee cap: 18.2M * 1.1 = 20.02M > cap, so at cap.
	tx = capTx(t, m, 1_000_000, 18_200_000)
	next, err = m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.Hash(), qt.Equals, tx.Hash())

	// Just at 10%: ceil(18,181,818 * 1.1) = 20,000,000 = the cap -> ok.
	tx = capTx(t, m, 1_000_000, 18_181_818)
	next, err = m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.Hash() != tx.Hash(), qt.IsTrue)
	tx = capTx(t, m, 1_000_000, 18_181_819) // ceil(20,000,000.9) = 20,000,001
	next, err = m.bumpedCopy(tx)
	c.Assert(err, qt.IsNil)
	c.Assert(next.Hash(), qt.Equals, tx.Hash())
}

// A rebroadcast of a tx already in the pool errors ("already known"); the
// at-cap entry must still wait MaxPendingTime before the next attempt.
func TestRetryStuckAtCapAlreadyKnownKeepsTheCadence(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newChainManager(t, f, opChain)
	tx := capTx(t, m, opFeeCapWei, opFeeCapWei)
	m.RecordPending(tx)
	old := time.Now().Add(-m.config.MaxPendingTime - time.Second)
	m.pending[5].submittedAt = old

	f.handle("eth_getTransactionCount", func(params []json.RawMessage) (any, error) {
		var tag string
		_ = json.Unmarshal(params[1], &tag)
		if tag == tagPending {
			return "0x6", nil
		}
		return "0x5", nil
	})
	f.handle("eth_sendRawTransaction", func([]json.RawMessage) (any, error) {
		return nil, fmt.Errorf("already known")
	})

	c.Assert(m.retryStuck(context.Background()), qt.IsNil)
	c.Assert(f.calls["eth_sendRawTransaction"], qt.Equals, 1)
	c.Assert(m.pending[5].retries, qt.Equals, 1)
	c.Assert(m.pending[5].hash, qt.Equals, tx.Hash())
	c.Assert(m.pending[5].submittedAt.After(old), qt.IsTrue)
	c.Assert(m.pending[5].capLogged, qt.IsTrue)

	c.Assert(m.retryStuck(context.Background()), qt.IsNil)
	c.Assert(f.calls["eth_sendRawTransaction"], qt.Equals, 1, qt.Commentf("no resend right after"))
	c.Assert(m.pending[5].retries, qt.Equals, 1)
}

// pendingIsConfirmed serves a load-balanced OP RPC that has not seen the
// pending tx: eth_getTransactionCount answers the confirmed nonce for both
// tags, so a live nonce-5 transaction looks unknown to the mempool.
func pendingIsConfirmed(f *fakeRPC) {
	f.handle("eth_getTransactionCount", func([]json.RawMessage) (any, error) { return "0x5", nil })
}

// sendRecorder answers "already known" for the transactions the pool already
// holds and accepts the rest, recording every broadcast.
func sendRecorder(f *fakeRPC, held map[common.Hash]bool) *[]*gethtypes.Transaction {
	var sent []*gethtypes.Transaction
	f.handle("eth_sendRawTransaction", func(params []json.RawMessage) (any, error) {
		var raw hexutil.Bytes
		if err := json.Unmarshal(params[0], &raw); err != nil {
			return nil, err
		}
		var decoded gethtypes.Transaction
		if err := decoded.UnmarshalBinary(raw); err != nil {
			return nil, err
		}
		sent = append(sent, &decoded)
		if held[decoded.Hash()] {
			return nil, fmt.Errorf("already known")
		}
		held[decoded.Hash()] = true
		return decoded.Hash(), nil
	})
	return &sent
}

// A tx older than MaxPendingTime is bumped even when the RPC reports it
// unknown to the mempool (OP has no public mempool; a load-balanced RPC can
// answer the pending nonce with the confirmed one): rebroadcasting it as is
// would keep it stuck under the base fee forever.
func TestRetryStuckBumpsAnOldTxThatLooksUnknown(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newChainManager(t, f, opChain)
	tx := capTx(t, m, 1_000_000, 1_016_000)
	m.RecordPending(tx)
	m.pending[5].submittedAt = time.Now().Add(-m.config.MaxPendingTime - time.Second)
	pendingIsConfirmed(f)
	sent := sendRecorder(f, map[common.Hash]bool{tx.Hash(): true})

	c.Assert(m.retryStuck(context.Background()), qt.IsNil)
	c.Assert(*sent, qt.HasLen, 1)
	c.Assert((*sent)[0].Hash() != tx.Hash(), qt.IsTrue, qt.Commentf("bumped, not rebroadcast as is"))
	c.Assert((*sent)[0].GasFeeCap().Cmp(tx.GasFeeCap()) > 0, qt.IsTrue)
	c.Assert(m.pending[5].hash, qt.Equals, (*sent)[0].Hash())
}

// The whole lifecycle under that RPC: a monitor pass every MonitorInterval
// for well past MaxPendingTime. The tx is rebroadcast while young, never
// dropped (which would free its nonce while it can still be mined), and
// bumped once it is MaxPendingTime old.
func TestRetryStuckFalseUnknownIsBumpedNotDropped(t *testing.T) {
	c := qt.New(t)
	f := newFakeRPC(t)
	m := newChainManager(t, f, opChain)
	tx := capTx(t, m, 1_000_000, 1_016_000)
	m.RecordPending(tx)
	pendingIsConfirmed(f)
	sent := sendRecorder(f, map[common.Hash]bool{tx.Hash(): true})

	passes := int(m.config.MaxPendingTime/m.config.MonitorInterval) + 5
	bumpedAt := -1
	for i := 0; i < passes; i++ {
		// One MonitorInterval goes by between passes.
		m.mu.Lock()
		cur, ok := m.pending[5]
		if ok {
			cur.submittedAt = cur.submittedAt.Add(-m.config.MonitorInterval)
		}
		m.mu.Unlock()
		c.Assert(ok, qt.IsTrue, qt.Commentf("dropped at pass %d (retries %d)", i, m.config.MaxRetries))
		c.Assert(m.retryStuck(context.Background()), qt.IsNil)
		if bumpedAt < 0 && m.pending[5] != nil && m.pending[5].hash != tx.Hash() {
			bumpedAt = i
		}
	}
	c.Assert(bumpedAt >= 0, qt.IsTrue, qt.Commentf("never bumped in %d passes; sends %d", passes, len(*sent)))
	c.Assert(m.pending[5].signed.GasFeeCap().Cmp(tx.GasFeeCap()) > 0, qt.IsTrue)
	c.Assert(m.pending[5].retries < m.config.MaxRetries, qt.IsTrue)
}

// Only "the pool already holds it" answers are spared from MaxRetries.
func TestAlreadyKnown(t *testing.T) {
	c := qt.New(t)
	for msg, want := range map[string]bool{
		"already known":            true,
		"ALREADY KNOWN":            true,
		"known transaction: 0xabc": true,
		"AlreadyKnown":             true,
		"insufficient funds for gas * price + value":             false,
		"replacement transaction underpriced":                    false,
		"nonce too low":                                          false,
		"transaction underpriced: tip needed 1, tip permitted 0": false,
	} {
		c.Assert(alreadyKnown(fmt.Errorf("%s", msg)), qt.Equals, want, qt.Commentf("%q", msg))
	}
}
