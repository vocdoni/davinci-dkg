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

	acceptedIdxs := make([]uint16, 0, n)
	allPoints := make([][]nodetypes.CurvePoint, 0, n)

	for i, addr := range committee {
		rec, err := m.GetContribution(callOpts, roundID, addr)
		if err != nil || !rec.Accepted {
			continue
		}
		pts, err := commitmentPointsFromCalldata(ctx, c, addr, t)
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

// commitmentPointsFromCalldata scans recent blocks for the submitContribution
// transaction from `contributor` and returns its t Feldman commitment points.
func commitmentPointsFromCalldata(
	ctx context.Context,
	c *web3.Contracts,
	contributor common.Address,
	t uint16,
) ([]nodetypes.CurvePoint, error) {
	client := c.Client()
	chainID := new(big.Int).SetUint64(c.ChainID)
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	start := uint64(0)
	if latest > 2000 {
		start = latest - 2000
	}
	signer := ethtypes.NewCancunSigner(chainID)
	for blk := start; blk <= latest; blk++ {
		block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blk))
		if err != nil {
			continue
		}
		for _, tx := range block.Transactions() {
			if tx.To() == nil || *tx.To() != c.Addresses.Manager {
				continue
			}
			from, err := ethtypes.Sender(signer, tx)
			if err != nil || from != contributor {
				continue
			}
			pts, err := parseCommitmentPoints(tx.Data(), t)
			if err != nil {
				continue
			}
			return pts, nil
		}
	}
	return nil, fmt.Errorf("submitContribution tx not found for %s", contributor)
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
