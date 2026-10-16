package web3

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	qt "github.com/frankban/quicktest"
)

// brokenPendingServer mimics an endpoint (mainnet.optimism.io) whose "pending"
// state is broken: no code and a zero transaction count, while "latest" is
// right. Only addresses in code have code.
func brokenPendingServer(code map[common.Address]string, latestNonce string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req testRPCRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		resp := testRPCResponse{JSONRPC: "2.0", ID: req.ID}
		var tag string
		if len(req.Params) > 1 {
			_ = json.Unmarshal(req.Params[1], &tag)
		}
		switch req.Method {
		case "eth_getCode":
			var addr common.Address
			_ = json.Unmarshal(req.Params[0], &addr)
			resp.Result = "0x"
			if tag != "pending" && code[addr] != "" {
				resp.Result = code[addr]
			}
		case "eth_getTransactionCount":
			resp.Result = latestNonce
			if tag == "pending" {
				resp.Result = "0x0"
			}
		case "eth_estimateGas":
			resp.Result = "0x5208"
		case "eth_sendRawTransaction":
			resp.Result = common.Hash{1}.Hex()
		default:
			resp.Error = &testRPCError{Code: -32601, Message: "not handled"}
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
}

// With a broken pending state, a bound transact must neither fail with
// ErrNoCode nor sign with the zero nonce; a codeless address still fails.
func TestTransactSurvivesBrokenPendingState(t *testing.T) {
	c := qt.New(t)
	live := common.HexToAddress("0x2000000000000000000000000000000000000002")
	empty := common.HexToAddress("0x3000000000000000000000000000000000000003")
	server := brokenPendingServer(map[common.Address]string{live: "0x6001"}, "0x23") // 35
	defer server.Close()
	pool, err := NewRPCPool([]string{server.URL})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	backend := NewPooledBackend(pool)
	ctx := context.Background()

	nonce, err := backend.PendingNonceAt(ctx, common.Address{9})
	c.Assert(err, qt.IsNil)
	c.Assert(nonce, qt.Equals, uint64(35))
	code, err := backend.PendingCodeAt(ctx, live)
	c.Assert(err, qt.IsNil)
	c.Assert(code, qt.DeepEquals, []byte{0x60, 0x01})

	var signedNonce uint64
	opts := &bind.TransactOpts{
		From:     common.Address{9},
		Context:  ctx,
		GasPrice: big.NewInt(1),
		Signer: func(_ common.Address, tx *types.Transaction) (*types.Transaction, error) {
			signedNonce = tx.Nonce()
			return tx, nil
		},
	}
	bc := bind.NewBoundContract(live, *managerABI, backend, backend, backend)
	_, err = bc.RawTransact(opts, []byte{1, 2, 3, 4})
	c.Assert(errors.Is(err, bind.ErrNoCode), qt.IsFalse, qt.Commentf("%v", err))
	c.Assert(signedNonce, qt.Equals, uint64(35))

	bc = bind.NewBoundContract(empty, *managerABI, backend, backend, backend)
	_, err = bc.RawTransact(opts, []byte{1, 2, 3, 4})
	c.Assert(errors.Is(err, bind.ErrNoCode), qt.IsTrue, qt.Commentf("%v", err))
}
