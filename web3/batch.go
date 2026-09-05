package web3

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/vocdoni/davinci-dkg/log"
)

// Call is one eth_call of a BatchCall: the target and calldata going in, the
// return data or the call's own error coming out.
type Call struct {
	To     common.Address
	Data   []byte
	Output []byte
	Err    error
}

// BatchCaller is a caller that can send several JSON-RPC requests in one
// batch, i.e. one HTTP round trip; *PooledBackend is one.
type BatchCaller interface {
	BatchCallContext(ctx context.Context, batch []rpc.BatchElem) error
}

// BatchCall runs every call as eth_call at `block` (nil = latest). Through a
// BatchCaller the calls travel in one JSON-RPC batch carrying the same block
// tag, so a batch pinned to a block reads one consistent state in one round
// trip instead of one per call; any other bind.ContractCaller (a test fake,
// a backend without batch support) gets them one by one through
// CallContract, and so does a batch the endpoint refused as a whole. Each
// call's own outcome — its return data, a revert, a malformed answer — lands
// in that Call; BatchCall itself never fails.
func BatchCall(ctx context.Context, caller bind.ContractCaller, block *big.Int, calls []*Call) {
	if len(calls) == 0 {
		return
	}
	if bc, ok := caller.(BatchCaller); ok {
		err := batchCall(ctx, bc, block, calls)
		if err == nil {
			return
		}
		log.Warnw("batched eth_call refused, sending the calls one by one", "calls", len(calls), "err", err)
	}
	for _, call := range calls {
		call.Output, call.Err = caller.CallContract(ctx, ethereum.CallMsg{To: &call.To, Data: call.Data}, block)
	}
}

func batchCall(ctx context.Context, bc BatchCaller, block *big.Int, calls []*Call) error {
	elems := make([]rpc.BatchElem, len(calls))
	outputs := make([]hexutil.Bytes, len(calls))
	for i, call := range calls {
		elems[i] = rpc.BatchElem{
			Method: "eth_call",
			Args:   []any{callArg(call), blockArg(block)},
			Result: &outputs[i],
		}
	}
	if err := bc.BatchCallContext(ctx, elems); err != nil {
		return err
	}
	for i, call := range calls {
		call.Output, call.Err = []byte(outputs[i]), elems[i].Error
	}
	return nil
}

// callArg is the eth_call transaction object ethclient sends for a read:
// only the target and the calldata (the sender defaults to the zero address,
// as bind's CallOpts do).
func callArg(call *Call) map[string]any {
	return map[string]any{"to": call.To, "input": hexutil.Bytes(call.Data)}
}

// blockArg is the eth_call block tag: "latest" for nil, else the hex height.
func blockArg(block *big.Int) string {
	if block == nil {
		return "latest"
	}
	return hexutil.EncodeBig(block)
}
