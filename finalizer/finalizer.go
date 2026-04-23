// Package finalizer builds and submits the finalizeRound transaction for a
// DKG round. It collects accepted contributions from on-chain calldata,
// generates the finalize ZK proof off-chain, and broadcasts the tx.
//
// Used by both the dkg-runner orchestrator and the davinci-dkg-node daemon
// (the latter via its auto-finalize stagger).
package finalizer

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"time"

	gnec "github.com/consensys/gnark-crypto/ecc"
	groth16backend "github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
// the finalize ZK proof, and submits finalizeRound. It returns the recovered
// per-participant share commitments on success.
//
// The caller is responsible for ensuring block.number >=
// policy.FinalizeNotBeforeBlock before calling — the contract gate would
// otherwise revert with FinalizeTooEarly.
func BuildAndSubmit(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	roundID [12]byte,
	t, n uint16,
	committee []common.Address,
) (*Result, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	// Bound the event-log scan to blocks since this round was created. Using
	// the round's SeedBlock as a lower bound keeps the FilterLogs cheap even
	// on long-lived deployments. We back off by 1 block to be safe against
	// any off-by-one in the seed-block emission relative to the contribution
	// submission window (contributions can only land after registration
	// closes, which is after seedBlock, so this is conservative).
	round, err := c.GetRound(ctx, roundID)
	if err != nil {
		return nil, fmt.Errorf("get round for log scan range: %w", err)
	}
	startBlock := uint64(0)
	if round.SeedBlock > 0 {
		startBlock = uint64(round.SeedBlock) - 1
	}

	acceptedIdxs := make([]uint16, 0, n)
	allPoints := make([][]nodetypes.CurvePoint, 0, n)

	for i, addr := range committee {
		rec, err := m.GetContribution(callOpts, roundID, addr)
		if err != nil || !rec.Accepted {
			continue
		}
		pts, err := commitmentPointsFromCalldata(ctx, c, m, roundID, addr, startBlock, t)
		if err != nil {
			return nil, fmt.Errorf("commitment points for %s: %w", addr, err)
		}
		acceptedIdxs = append(acceptedIdxs, uint16(i+1))
		allPoints = append(allPoints, pts)
	}
	if len(acceptedIdxs) < int(t) {
		return nil, fmt.Errorf("only %d/%d accepted contributions", len(acceptedIdxs), t)
	}

	roundHash := new(big.Int).SetBytes(roundID[:])
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
	runtime, err := finalize.Artifacts.LoadOrSetupForCircuit(ctx, &finalize.FinalizeCircuit{})
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
	tx, err := m.FinalizeRound(auth, roundID,
		common.BigToHash(pi.AggregateHash),
		common.BigToHash(pi.CollectivePublicKey),
		common.BigToHash(pi.ShareCommitmentHash),
		transcriptBytes, proofBytes, inputBytes,
	)
	if err != nil {
		return nil, err
	}
	if err := txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		return nil, err
	}
	res := &Result{ShareCommitments: pi.ShareCommitments}
	if receipt, err := c.Client().TransactionReceipt(ctx, tx.Hash()); err == nil {
		res.GasUsed = receipt.GasUsed
		log.Infow("finalizeRound tx mined", "tx", tx.Hash().Hex(), "gas", receipt.GasUsed)
	}
	return res, nil
}

// commitmentPointsFromCalldata locates the submitContribution tx from
// `contributor` for the given round via the ContributionSubmitted event log
// (indexed by roundId + contributor), then fetches that single transaction
// and parses its t Feldman commitment points from the calldata transcript.
//
// This replaces an earlier implementation that scanned the last ~2000 blocks
// serially via BlockByNumber — on a public RPC that produced multi-minute
// stalls per finalize attempt and was the dominant source of finalize latency.
// The event-log path is O(1) RPC calls regardless of how long ago the
// contribution landed, so finalize fires within the stagger window.
func commitmentPointsFromCalldata(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	roundID [12]byte,
	contributor common.Address,
	startBlock uint64,
	t uint16,
) ([]nodetypes.CurvePoint, error) {
	client := c.Client()
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("read head: %w", err)
	}

	filterOpts := &bind.FilterOpts{
		Context: ctx,
		Start:   startBlock,
		End:     &latest,
	}
	it, err := m.FilterContributionSubmitted(
		filterOpts,
		[][12]byte{roundID},
		[]common.Address{contributor},
	)
	if err != nil {
		return nil, fmt.Errorf("filter ContributionSubmitted: %w", err)
	}
	defer it.Close()

	if !it.Next() {
		if err := it.Error(); err != nil {
			return nil, fmt.Errorf("iterate ContributionSubmitted: %w", err)
		}
		return nil, fmt.Errorf("no ContributionSubmitted event for %s in round %x (range %d..%d)",
			contributor, roundID, startBlock, latest)
	}

	txHash := it.Event.Raw.TxHash
	tx, _, err := client.TransactionByHash(ctx, txHash)
	if err != nil {
		return nil, fmt.Errorf("fetch contribution tx %s: %w", txHash.Hex(), err)
	}
	pts, err := parseCommitmentPoints(tx.Data(), t)
	if err != nil {
		return nil, fmt.Errorf("parse commitment points from tx %s: %w", txHash.Hex(), err)
	}
	return pts, nil
}

// parseCommitmentPoints extracts the first t Feldman commitment points from
// the submitContribution calldata transcript.
func parseCommitmentPoints(data []byte, t uint16) ([]nodetypes.CurvePoint, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("data too short")
	}
	payload := data[4:]
	// transcript is the 7th parameter (index 6, after commitment0X and commitment0Y),
	// offset is at head bytes 192..223.
	if len(payload) < 224 {
		return nil, fmt.Errorf("payload head too short")
	}
	tOffset := int64(new(big.Int).SetBytes(pad32(payload[192:224])).Uint64())
	if int(tOffset)+32 > len(payload) {
		return nil, fmt.Errorf("transcript offset out of bounds")
	}
	tLen := int64(new(big.Int).SetBytes(pad32(payload[tOffset : tOffset+32])).Uint64())
	tStart := tOffset + 32
	if tStart+tLen > int64(len(payload)) {
		return nil, fmt.Errorf("transcript out of bounds")
	}
	tr := payload[tStart : tStart+tLen]
	if len(tr) < contribution.MaxCoefficients*64 {
		return nil, fmt.Errorf("transcript too short")
	}
	pts := make([]nodetypes.CurvePoint, t)
	for k := uint16(0); k < t; k++ {
		x := new(big.Int).SetBytes(tr[k*64 : k*64+32])
		y := new(big.Int).SetBytes(tr[k*64+32 : k*64+64])
		pts[k] = nodetypes.CurvePoint{X: x, Y: y}
	}
	return pts, nil
}

func pad32(b []byte) []byte {
	if len(b) >= 32 {
		return b[:32]
	}
	out := make([]byte, 32)
	copy(out[32-len(b):], b)
	return out
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
