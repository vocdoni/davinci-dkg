// Package finalizer builds and submits the finalizeEpoch transaction for a
// DKG epoch. It collects accepted contributions from on-chain calldata,
// generates the finalize ZK proof off-chain, and broadcasts the tx.
//
// Used by the davinci-dkg-node daemon
// (the latter via its auto-finalize stagger).
package finalizer

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"reflect"
	"sync"
	"time"

	gnec "github.com/consensys/gnark-crypto/ecc"
	groth16backend "github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// Result carries the outputs of a successful finalize.
type Result struct {
	ShareCommitments []nodetypes.CurvePoint
	GasUsed          uint64
}

// BuildAndSubmit reads accepted contributions from on-chain calldata, builds
// the finalize ZK proof, and submits finalizeEpoch. It returns the recovered
// per-participant share commitments on success.
//
// The caller is responsible for ensuring block.number >=
// policy.LiveNotBeforeBlock before calling — the contract gate would
// otherwise revert with FinalizeTooEarly.
func BuildAndSubmit(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	epochID [12]byte,
	t, n uint16,
	committee []common.Address,
) (*Result, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	// Bound the event-log scan to blocks since this epoch was created. Using
	// the epoch's SeedBlock as a lower bound keeps the FilterLogs cheap even
	// on long-lived deployments. We back off by 1 block to be safe against
	// any off-by-one in the seed-block emission relative to the contribution
	// submission window (contributions can only land after registration
	// closes, which is after seedBlock, so this is conservative).
	epoch, err := c.GetEpoch(ctx, epochID)
	if err != nil {
		return nil, fmt.Errorf("get epoch for log scan range: %w", err)
	}
	startBlock := uint64(0)
	if epoch.SeedBlock > 0 {
		startBlock = uint64(epoch.SeedBlock) - 1
	}

	acceptedIdxs := make([]uint16, 0, n)
	allPoints := make([][]nodetypes.CurvePoint, 0, n)

	for i, addr := range committee {
		rec, err := m.GetContribution(callOpts, epochID, addr)
		if err != nil || !rec.Accepted {
			continue
		}
		data, err := ContributionCalldata(ctx, c.Client(), m, epochID, addr, startBlock)
		if err != nil {
			return nil, fmt.Errorf("contribution calldata for %s: %w", addr, err)
		}
		pts, err := parseCommitmentPoints(data, t)
		if err != nil {
			return nil, fmt.Errorf("parse commitment points for %s: %w", addr, err)
		}
		acceptedIdxs = append(acceptedIdxs, uint16(i+1))
		allPoints = append(allPoints, pts)
	}
	if len(acceptedIdxs) < int(t) {
		return nil, fmt.Errorf("only %d/%d accepted contributions", len(acceptedIdxs), t)
	}

	roundHash := new(big.Int).SetBytes(epochID[:])
	asgn := finalize.CommitmentPointsAssignment{
		RoundHash:          roundHash,
		Threshold:          t,
		CommitteeSize:      n,
		ParticipantIndexes: acceptedIdxs,
		ContributionPoints: allPoints,
	}
	witness, pi, err := finalize.BuildWitnessFromCommitmentPoints(asgn)
	if err != nil {
		return nil, fmt.Errorf("build finalize witness: %w", err)
	}
	runtime, err := finalize.Artifacts.LoadPinned(ctx, &finalize.FinalizeCircuit{})
	if err != nil {
		return nil, fmt.Errorf("load finalize circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove finalize: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return nil, err
	}
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return nil, err
	}

	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := m.FinalizeEpoch(
		auth, epochID,
		common.BigToHash(pi.AggregateHash),
		common.BigToHash(pi.CollectivePublicKey),
		common.BigToHash(pi.ShareCommitmentHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		return nil, err
	}
	txm.RecordPending(tx)
	if err := txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		return nil, err
	}
	res := &Result{ShareCommitments: pi.ShareCommitments}
	if receipt, err := c.Client().TransactionReceipt(ctx, tx.Hash()); err == nil {
		res.GasUsed = receipt.GasUsed
		log.Infow("finalizeEpoch tx mined", "tx", tx.Hash().Hex(), "gas", receipt.GasUsed)
	}
	return res, nil
}

// LogRangeBlocks bounds each eth_getLogs call; most public RPC providers cap
// the range at 10k blocks.
const LogRangeBlocks = 10_000

// chainReader is the slice of ethclient.Client that ContributionCalldata
// needs, so tests can drive it without an RPC endpoint.
type chainReader interface {
	BlockNumber(ctx context.Context) (uint64, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (*ethtypes.Transaction, bool, error)
}

// ContributionCalldata locates the submitContribution tx from `contributor`
// for the given epoch via the ContributionSubmitted event log (indexed by
// epochId + contributor) and returns that transaction's raw calldata. The
// encrypted shares and commitment points only live there.
//
// The log scan walks [startBlock, head] in LogRangeBlocks chunks and stops
// at the first chunk that holds the event, so it stays within provider
// limits and costs O(epoch age / LogRangeBlocks) RPC calls.
func ContributionCalldata(
	ctx context.Context,
	client chainReader,
	m *gtypes.DKGManager,
	epochID [12]byte,
	contributor common.Address,
	startBlock uint64,
) ([]byte, error) {
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("read head: %w", err)
	}
	for start := startBlock; start <= latest; start += LogRangeBlocks {
		end := min(start+LogRangeBlocks-1, latest)
		txHash, found, err := findContributionTx(ctx, m, epochID, contributor, start, end)
		if err != nil {
			return nil, err
		}
		if !found {
			continue
		}
		tx, _, err := client.TransactionByHash(ctx, txHash)
		if err != nil {
			return nil, fmt.Errorf("fetch contribution tx %s: %w", txHash.Hex(), err)
		}
		return tx.Data(), nil
	}
	return nil, fmt.Errorf("no ContributionSubmitted event for %s in epoch %x (range %d..%d)",
		contributor, epochID, startBlock, latest)
}

// findContributionTx returns the hash of the tx that emitted
// ContributionSubmitted(epochID, contributor) within [start, end].
func findContributionTx(
	ctx context.Context,
	m *gtypes.DKGManager,
	epochID [12]byte,
	contributor common.Address,
	start, end uint64,
) (common.Hash, bool, error) {
	it, err := m.FilterContributionSubmitted(
		&bind.FilterOpts{Context: ctx, Start: start, End: &end},
		[][12]byte{epochID},
		[]common.Address{contributor},
	)
	if err != nil {
		return common.Hash{}, false, fmt.Errorf("filter ContributionSubmitted [%d,%d]: %w", start, end, err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		if err := it.Error(); err != nil {
			return common.Hash{}, false, fmt.Errorf("iterate ContributionSubmitted: %w", err)
		}
		return common.Hash{}, false, nil
	}
	return it.Event.Raw.TxHash, true, nil
}

// submitContributionMethod resolves the submitContribution ABI method once;
// its selector and input layout are the only things ContributionTranscript
// trusts about the calldata it is handed.
var submitContributionMethod = sync.OnceValues(func() (abi.Method, error) {
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	if err != nil {
		return abi.Method{}, fmt.Errorf("parse DKGManager ABI: %w", err)
	}
	method, ok := parsed.Methods["submitContribution"]
	if !ok {
		return abi.Method{}, fmt.Errorf("submitContribution missing from DKGManager ABI")
	}
	return method, nil
})

// ContributionTranscript extracts the `transcript` argument from raw
// submitContribution calldata:
//
//	submitContribution(bytes12 epochId, uint16 contributorIndex,
//	    bytes32 commitmentsHash, bytes32 encryptedSharesHash,
//	    bytes transcript, bytes proof, bytes input)
//
// The calldata is located through an event log and may be anything a
// transaction sender chose to put there, so the selector is checked against
// the ABI and the dynamic offsets are decoded by the ABI unpacker, which
// bounds-checks them without overflow. Any malformed input yields an error,
// never a panic. The returned slice is the 8·MaxN-word layout documented in
// DKGManager.submitContribution.
func ContributionTranscript(data []byte) ([]byte, error) {
	method, err := submitContributionMethod()
	if err != nil {
		return nil, err
	}
	if len(data) < 4 || !bytes.Equal(data[:4], method.ID) {
		return nil, fmt.Errorf("calldata is not a submitContribution call")
	}
	args, err := method.Inputs.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("decode submitContribution calldata: %w", err)
	}
	const transcriptArg = 4
	if len(args) <= transcriptArg {
		return nil, fmt.Errorf("submitContribution calldata has %d arguments", len(args))
	}
	tr, ok := args[transcriptArg].([]byte)
	if !ok {
		return nil, fmt.Errorf("transcript argument is %T, want []byte", args[transcriptArg])
	}
	if want := 8 * contribution.MaxRecipients * 32; len(tr) != want {
		return nil, fmt.Errorf("transcript is %d bytes, want %d", len(tr), want)
	}
	return tr, nil
}

// parseCommitmentPoints extracts the first t Feldman commitment points from
// the submitContribution calldata transcript.
func parseCommitmentPoints(data []byte, t uint16) ([]nodetypes.CurvePoint, error) {
	tr, err := ContributionTranscript(data)
	if err != nil {
		return nil, err
	}
	pts := make([]nodetypes.CurvePoint, t)
	for k := uint16(0); k < t; k++ {
		x := new(big.Int).SetBytes(tr[k*64 : k*64+32])
		y := new(big.Int).SetBytes(tr[k*64+32 : k*64+64])
		pts[k] = nodetypes.CurvePoint{X: x, Y: y}
	}
	return pts, nil
}

// ── proof / witness helpers (duplicated from cmd/* until consolidated) ─────

func marshalSolidityProof(proof groth16backend.Proof) ([]byte, error) {
	p, ok := proof.(*groth16bn254.Proof)
	if !ok {
		return nil, fmt.Errorf("unexpected proof type %T", proof)
	}
	return p.MarshalSolidity(), nil
}

func encodePublicWitness(pub frontend.Circuit) ([]byte, error) {
	w, err := frontend.NewWitness(pub, gnec.BN254.ScalarField(), frontend.PublicOnly())
	if err != nil {
		return nil, fmt.Errorf("build public witness: %w", err)
	}
	vals, err := witnessVector(w.Vector())
	if err != nil {
		return nil, err
	}
	return encodeWords(vals...)
}

func encodeWords(values ...*big.Int) ([]byte, error) {
	out := make([]byte, 0, len(values)*32)
	for i, v := range values {
		if v == nil {
			return nil, fmt.Errorf("value %d is nil", i)
		}
		out = append(out, common.LeftPadBytes(v.Bytes(), 32)...)
	}
	return out, nil
}

func witnessVector(vector any) ([]*big.Int, error) {
	rv := reflect.ValueOf(vector)
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("unexpected witness vector type %T", vector)
	}
	out := make([]*big.Int, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		mth := rv.Index(i).Addr().MethodByName("BigInt")
		if !mth.IsValid() {
			return nil, fmt.Errorf("element %d missing BigInt", i)
		}
		res := mth.Call([]reflect.Value{reflect.ValueOf(new(big.Int))})
		v, ok := res[0].Interface().(*big.Int)
		if !ok {
			return nil, fmt.Errorf("element %d BigInt bad type", i)
		}
		out[i] = new(big.Int).Set(v)
	}
	return out, nil
}
