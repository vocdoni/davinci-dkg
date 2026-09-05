package web3

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	qt "github.com/frankban/quicktest"
)

// batchRPCServer is a JSON-RPC server that answers eth_call with the call's
// own input (so a test can tell the answers apart) and records every HTTP
// request body it received: a batch arrives as one JSON array.
type batchRPCServer struct {
	mu       sync.Mutex
	requests []json.RawMessage
	blocks   []string
	srv      *httptest.Server
}

func newBatchRPCServer(t *testing.T) *batchRPCServer {
	t.Helper()
	b := &batchRPCServer{}
	b.srv = httptest.NewServer(http.HandlerFunc(b.serve))
	t.Cleanup(b.srv.Close)
	return b
}

func (b *batchRPCServer) serve(w http.ResponseWriter, r *http.Request) {
	var body json.RawMessage
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b.mu.Lock()
	b.requests = append(b.requests, body)
	b.mu.Unlock()

	var reqs []testRPCRequest
	batch := len(body) > 0 && body[0] == '['
	if batch {
		_ = json.Unmarshal(body, &reqs)
	} else {
		var one testRPCRequest
		_ = json.Unmarshal(body, &one)
		reqs = []testRPCRequest{one}
	}
	resps := make([]testRPCResponse, len(reqs))
	for i, req := range reqs {
		resps[i] = testRPCResponse{JSONRPC: "2.0", ID: req.ID}
		switch req.Method {
		case "eth_chainId":
			resps[i].Result = "0x7a69"
		case "eth_call":
			var call struct {
				To    common.Address `json:"to"`
				Input hexutil.Bytes  `json:"input"`
			}
			_ = json.Unmarshal(req.Params[0], &call)
			var block string
			_ = json.Unmarshal(req.Params[1], &block)
			b.mu.Lock()
			b.blocks = append(b.blocks, block)
			b.mu.Unlock()
			if len(call.Input) > 0 && call.Input[0] == 0xff {
				resps[i].Error = &testRPCError{Code: 3, Message: "execution reverted"}
				break
			}
			resps[i].Result = hexutil.Bytes(append([]byte{0xee}, call.Input...))
		default:
			resps[i].Error = &testRPCError{Code: -32601, Message: "not handled"}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	if batch {
		_ = json.NewEncoder(w).Encode(resps)
		return
	}
	_ = json.NewEncoder(w).Encode(resps[0])
}

// Several reads issued together travel as one JSON-RPC batch — one HTTP
// round trip — pinned to the same block tag, each call getting its own
// answer or its own error; the pool counts the batch and its methods.
func TestBatchCallSendsOneRoundTripPinnedToTheBlock(t *testing.T) {
	c := qt.New(t)
	server := newBatchRPCServer(t)
	pool, err := NewRPCPool([]string{server.srv.URL})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	backend := NewPooledBackend(pool)

	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	calls := []*Call{
		{To: to, Data: []byte{1, 2, 3}},
		{To: to, Data: []byte{0xff, 9}}, // this one reverts
		{To: to, Data: []byte{4, 5, 6}},
	}
	BatchCall(context.Background(), backend, big.NewInt(0x1234), calls)

	c.Assert(server.requests, qt.HasLen, 1, qt.Commentf("three calls must cost one HTTP request"))
	c.Assert(server.requests[0][0], qt.Equals, byte('['), qt.Commentf("the request must be a JSON-RPC batch"))
	c.Assert(server.blocks, qt.DeepEquals, []string{"0x1234", "0x1234", "0x1234"})
	c.Assert(calls[0].Err, qt.IsNil)
	c.Assert(calls[0].Output, qt.DeepEquals, []byte{0xee, 1, 2, 3})
	c.Assert(calls[1].Err, qt.IsNotNil, qt.Commentf("a reverting call fails alone"))
	c.Assert(calls[1].Output, qt.HasLen, 0)
	c.Assert(calls[2].Err, qt.IsNil)
	c.Assert(calls[2].Output, qt.DeepEquals, []byte{0xee, 4, 5, 6})

	stats := pool.Snapshot()
	c.Assert(stats["batch"], qt.Equals, uint64(1))
	c.Assert(stats["eth_call"], qt.Equals, uint64(3))

	// nil block reads the latest state.
	BatchCall(context.Background(), backend, nil, calls[:1])
	c.Assert(server.blocks[len(server.blocks)-1], qt.Equals, "latest")
	c.Assert(len(calls), qt.Equals, 3)
	BatchCall(context.Background(), backend, nil, nil)
	c.Assert(server.requests, qt.HasLen, 2, qt.Commentf("an empty batch sends nothing"))
}

// singleCaller is a bind.ContractCaller without batch support.
type singleCaller struct {
	calls []ethereum.CallMsg
	block []*big.Int
}

func (s *singleCaller) CodeAt(context.Context, common.Address, *big.Int) ([]byte, error) {
	return nil, errors.New("unsupported")
}

func (s *singleCaller) CallContract(_ context.Context, msg ethereum.CallMsg, block *big.Int) ([]byte, error) {
	s.calls = append(s.calls, msg)
	s.block = append(s.block, block)
	if msg.Data[0] == 0xff {
		return nil, errors.New("execution reverted")
	}
	return append([]byte{0xee}, msg.Data...), nil
}

// A caller that cannot batch gets the same calls one by one, at the same
// block, with the same per-call outcomes.
func TestBatchCallFallsBackToSingleCalls(t *testing.T) {
	c := qt.New(t)
	caller := &singleCaller{}
	to := common.HexToAddress("0x1")
	calls := []*Call{{To: to, Data: []byte{1}}, {To: to, Data: []byte{0xff}}, {To: to, Data: []byte{2}}}
	block := big.NewInt(77)
	BatchCall(context.Background(), caller, block, calls)

	c.Assert(caller.calls, qt.HasLen, 3)
	for i, msg := range caller.calls {
		c.Assert(*msg.To, qt.Equals, to)
		c.Assert(msg.Data, qt.DeepEquals, calls[i].Data)
		c.Assert(caller.block[i], qt.Equals, block)
	}
	c.Assert(calls[0].Output, qt.DeepEquals, []byte{0xee, 1})
	c.Assert(calls[1].Err, qt.IsNotNil)
	c.Assert(calls[2].Output, qt.DeepEquals, []byte{0xee, 2})
}

// Every request the pooled backend sends is counted under its JSON-RPC
// method, so a node can log what a poll cycle costs its provider.
func TestPooledBackendCountsRequestsPerMethod(t *testing.T) {
	c := qt.New(t)
	server := testRPCServer()
	defer server.Close()
	pool, err := NewRPCPool([]string{server.URL})
	c.Assert(err, qt.IsNil)
	defer pool.Close()
	backend := NewPooledBackend(pool)
	ctx := context.Background()

	_, _ = backend.FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
	_, _ = backend.FilterLogs(ctx, ethereum.FilterQuery{FromBlock: big.NewInt(1), ToBlock: big.NewInt(2)})
	_, _ = backend.BlockNumber(ctx)         // not served by the test server: still one request
	_, _ = backend.HeaderByNumber(ctx, nil) // idem
	to := common.HexToAddress("0x2000000000000000000000000000000000000002")
	_, _ = backend.CallContract(ctx, ethereum.CallMsg{To: &to, Data: managerABI.Methods["appManager"].ID}, nil)

	stats := pool.Snapshot()
	c.Assert(stats["eth_getLogs"], qt.Equals, uint64(2))
	c.Assert(stats["eth_blockNumber"], qt.Equals, uint64(1))
	c.Assert(stats["eth_getBlockByNumber"], qt.Equals, uint64(1))
	c.Assert(stats["eth_call"], qt.Equals, uint64(1))
	stats["eth_call"] = 99
	c.Assert(pool.Snapshot()["eth_call"], qt.Equals, uint64(1), qt.Commentf("Snapshot returns a copy"))
}
