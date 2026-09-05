// Package finalizer drives the two on-chain steps that turn an epoch's
// accepted contributions into usable keys: the proof-less finalizeEpoch,
// which freezes the accepted set and makes the epoch Live, and
// activatePoolKey, which proves one of the MaxK dealt polynomials into a
// committee key from the contributions' on-chain calldata.
//
// Used by the davinci-dkg-node daemon (via its stagger rotations).
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
	"github.com/vocdoni/davinci-dkg/circuits"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/poolkey"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// Result carries the outputs of a successful pool-key activation.
// ParticipantIndexes lists the accepted contributors; ShareCommitments is
// member-indexed, one D_p per committee member p = i+1 (contributing or
// not), the leaves of the Merkle root the contract stored.
type Result struct {
	PoolKey            nodetypes.CurvePoint
	TranscriptDigest   common.Hash
	ParticipantIndexes []uint16
	ShareCommitments   []nodetypes.CurvePoint
	GasUsed            uint64
}

// FinalizeEpoch submits the proof-less finalizeEpoch transaction, which
// freezes the accepted contributor set and moves the epoch to Live. The
// caller is responsible for block.number >= policy.LiveNotBeforeBlock and
// for contributionCount >= minValidContributions — the contract gates on
// both.
func FinalizeEpoch(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	epochID [12]byte,
) (uint64, error) {
	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := m.FinalizeEpoch(auth, epochID)
	if err != nil {
		return 0, err
	}
	txm.RecordPending(tx)
	if err := txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		return 0, err
	}
	var gasUsed uint64
	if receipt, err := c.Client().TransactionReceipt(ctx, tx.Hash()); err == nil {
		gasUsed = receipt.GasUsed
		log.Infow("finalizeEpoch tx mined", "tx", tx.Hash().Hex(), "gas", receipt.GasUsed)
	}
	return gasUsed, nil
}

// CalldataCache memoises raw submitContribution calldata on disk so a
// rate-limited RPC is not rescanned for the same dealer's transaction on
// every attempt. Get returns ok=false on any read problem (the caller then
// falls back to the event log); a failed Put only costs a later refetch.
// Entries are written after the calldata has been validated and a contribution
// is immutable once accepted, so they never go stale.
type CalldataCache interface {
	Get(epochID [12]byte, dealer common.Address) ([]byte, bool)
	Put(epochID [12]byte, dealer common.Address, data []byte)
}

// ActivationAssignment reconstructs the activation statement of pool key
// `keyIndex` from the accepted contributions' on-chain calldata: the input
// poolkey.BuildWitness turns into the aggregate commitments, P_j and every
// committee member's share commitment D_p. No proof is involved. A non-nil
// `cache` is consulted before each dealer's event-log scan and filled with
// the calldata once it decodes.
//
// Activation is permissionless but needs every contributor's *whole*
// commitment set: reproducing the stored Poseidon commitmentsHash absorbs the
// digests of the keys that are not being activated.
func ActivationAssignment(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	epochID [12]byte,
	t, n uint16,
	committee []common.Address,
	keyIndex uint8,
	cache CalldataCache,
) (poolkey.Assignment, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	// Bound the event-log scan to blocks since this epoch was created. Using
	// the epoch's SeedBlock as a lower bound keeps the FilterLogs cheap even
	// on long-lived deployments. We back off by 1 block to be safe against
	// any off-by-one in the seed-block emission relative to the contribution
	// submission window (contributions can only land after registration
	// closes, which is after seedBlock, so this is conservative).
	epoch, err := c.GetEpoch(ctx, epochID)
	if err != nil {
		return poolkey.Assignment{}, fmt.Errorf("get epoch for log scan range: %w", err)
	}
	startBlock := uint64(0)
	if epoch.SeedBlock > 0 {
		startBlock = uint64(epoch.SeedBlock) - 1
	}

	acceptedIdxs := make([]uint16, 0, n)
	allCommitments := make([][][]nodetypes.CurvePoint, 0, n)

	for i, addr := range committee {
		rec, err := m.GetContribution(callOpts, epochID, addr)
		if err != nil {
			// A read failure is not a missing contribution: proving over a
			// subset would activate a key the contract rejects (or, with a
			// forged accepted count, the wrong key).
			return poolkey.Assignment{}, fmt.Errorf("get contribution of %s: %w", addr, err)
		}
		if !rec.Accepted {
			continue
		}
		data, cached := []byte(nil), false
		if cache != nil {
			data, cached = cache.Get(epochID, addr)
		}
		if data == nil {
			data, err = ContributionCalldata(ctx, c.Client(), m, epochID, addr, startBlock)
			if err != nil {
				return poolkey.Assignment{}, fmt.Errorf("contribution calldata for %s: %w", addr, err)
			}
		}
		sets, err := parseCommitmentPoints(data, t)
		if err != nil {
			return poolkey.Assignment{}, fmt.Errorf("parse commitment points for %s: %w", addr, err)
		}
		if cache != nil && !cached {
			cache.Put(epochID, addr, data)
		}
		acceptedIdxs = append(acceptedIdxs, uint16(i+1))
		allCommitments = append(allCommitments, sets)
	}
	// The contract activates over exactly the frozen accepted set; anything
	// else means the committee list or the chain view is stale.
	if len(acceptedIdxs) != int(epoch.ContributionCount) {
		return poolkey.Assignment{}, fmt.Errorf(
			"reconstructed %d accepted contributions, epoch has %d", len(acceptedIdxs), epoch.ContributionCount,
		)
	}
	if len(acceptedIdxs) < int(t) {
		return poolkey.Assignment{}, fmt.Errorf("only %d/%d accepted contributions", len(acceptedIdxs), t)
	}

	return poolkey.Assignment{
		RoundHash:          new(big.Int).SetBytes(epochID[:]),
		Threshold:          t,
		CommitteeSize:      n,
		KeyIndex:           keyIndex,
		ParticipantIndexes: acceptedIdxs,
		Commitments:        allCommitments,
	}, nil
}

// PoolKeyStatement derives, without proving, what activatePoolKey committed
// to for `keyIndex`: P_j and the member-indexed share commitments
// (ShareCommitments[p-1] = D_p for p ≤ n, the identity beyond). Committee
// members use it to rebuild their Merkle leaves from the contributions'
// calldata instead of trusting the activation transaction's outer calldata,
// which an activation relayed through a contract does not even carry. A
// non-nil `cache` (see ActivationAssignment) spares the per-dealer event-log
// rescans.
func PoolKeyStatement(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	epochID [12]byte,
	t, n uint16,
	committee []common.Address,
	keyIndex uint8,
	cache CalldataCache,
) (*poolkey.PublicInputs, error) {
	asgn, err := ActivationAssignment(ctx, c, m, epochID, t, n, committee, keyIndex, cache)
	if err != nil {
		return nil, err
	}
	_, pi, err := poolkey.BuildWitness(asgn)
	if err != nil {
		return nil, fmt.Errorf("build pool key statement: %w", err)
	}
	return pi, nil
}

// BuildAndSubmitActivation reads the accepted contributions from on-chain
// calldata, proves that pool key `keyIndex` is the sum of their key-j
// commitments, and submits activatePoolKey. It returns the activated key and
// the per-member share commitments the contract folded into the Merkle root.
// The pool-key runtime is loaded from the pinned artifacts on every call;
// long-lived callers pass their own to ProveAndSubmitActivation.
func BuildAndSubmitActivation(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	epochID [12]byte,
	t, n uint16,
	committee []common.Address,
	keyIndex uint8,
) (*Result, error) {
	runtime, err := poolkey.Artifacts.LoadPinned(ctx, &poolkey.PoolKeyCircuit{})
	if err != nil {
		return nil, fmt.Errorf("load pool key circuit: %w", err)
	}
	return ProveAndSubmitActivation(ctx, c, m, txm, runtime, epochID, t, n, committee, keyIndex, nil)
}

// ProveAndSubmitActivation is BuildAndSubmitActivation with a caller-supplied
// pool-key runtime (the node loads all four circuits once at startup) and an
// optional contribution-calldata cache (see ActivationAssignment).
func ProveAndSubmitActivation(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	runtime *circuits.CircuitRuntime,
	epochID [12]byte,
	t, n uint16,
	committee []common.Address,
	keyIndex uint8,
	cache CalldataCache,
) (*Result, error) {
	if runtime == nil {
		return nil, fmt.Errorf("pool key circuit runtime is required")
	}
	asgn, err := ActivationAssignment(ctx, c, m, epochID, t, n, committee, keyIndex, cache)
	if err != nil {
		return nil, err
	}
	witness, pi, err := poolkey.BuildWitness(asgn)
	if err != nil {
		return nil, fmt.Errorf("build pool key witness: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove pool key: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return nil, err
	}
	transcriptScalars, err := pi.TranscriptScalars()
	if err != nil {
		return nil, fmt.Errorf("pool key transcript scalars: %w", err)
	}
	transcriptBytes, err := encodeWords(transcriptScalars...)
	if err != nil {
		return nil, err
	}

	// The Poseidon digest of the transcript is a public input the contract
	// folds into the Fiat–Shamir anchor; it travels as its own argument so
	// the contract never has to hash 194 Poseidon inputs itself.
	digest := common.BigToHash(pi.TranscriptDigest)

	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := m.ActivatePoolKey(auth, epochID, keyIndex, digest, transcriptBytes, proofBytes, inputBytes)
	if err != nil {
		return nil, err
	}
	txm.RecordPending(tx)
	if err := txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		return nil, err
	}
	res := &Result{
		PoolKey:            pi.PoolKey,
		TranscriptDigest:   digest,
		ParticipantIndexes: asgn.ParticipantIndexes,
		ShareCommitments:   pi.ShareCommitments[:n],
	}
	if receipt, err := c.Client().TransactionReceipt(ctx, tx.Hash()); err == nil {
		res.GasUsed = receipt.GasUsed
		log.Infow("activatePoolKey tx mined", "tx", tx.Hash().Hex(), "key", keyIndex, "gas", receipt.GasUsed)
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
	return calldataFromEvent(ctx, client, startBlock, "ContributionSubmitted",
		func(ctx context.Context, start, end uint64) (common.Hash, bool, error) {
			return findContributionTx(ctx, m, epochID, contributor, start, end)
		})
}

// ActivationCalldata locates the activatePoolKey tx for (epoch, keyIndex)
// via the PoolKeyActivated event log and returns its raw calldata. The
// per-member share commitments a partial decryption must prove a Merkle path
// against only live there.
func ActivationCalldata(
	ctx context.Context,
	client chainReader,
	m *gtypes.DKGManager,
	epochID [12]byte,
	keyIndex uint8,
	startBlock uint64,
) ([]byte, error) {
	return calldataFromEvent(ctx, client, startBlock, "PoolKeyActivated",
		func(ctx context.Context, start, end uint64) (common.Hash, bool, error) {
			return findActivationTx(ctx, m, epochID, keyIndex, start, end)
		})
}

// calldataFromEvent walks [startBlock, head] in LogRangeBlocks chunks until
// `find` reports the event, then fetches that transaction's calldata.
func calldataFromEvent(
	ctx context.Context,
	client chainReader,
	startBlock uint64,
	event string,
	find func(ctx context.Context, start, end uint64) (common.Hash, bool, error),
) ([]byte, error) {
	latest, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("read head: %w", err)
	}
	for start := startBlock; start <= latest; start += LogRangeBlocks {
		end := min(start+LogRangeBlocks-1, latest)
		txHash, found, err := find(ctx, start, end)
		if err != nil {
			return nil, err
		}
		if !found {
			continue
		}
		tx, _, err := client.TransactionByHash(ctx, txHash)
		if err != nil {
			return nil, fmt.Errorf("fetch %s tx %s: %w", event, txHash.Hex(), err)
		}
		return tx.Data(), nil
	}
	return nil, fmt.Errorf("no %s event in range %d..%d", event, startBlock, latest)
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

// findActivationTx returns the hash of the tx that emitted
// PoolKeyActivated(epochID, keyIndex) within [start, end].
func findActivationTx(
	ctx context.Context,
	m *gtypes.DKGManager,
	epochID [12]byte,
	keyIndex uint8,
	start, end uint64,
) (common.Hash, bool, error) {
	it, err := m.FilterPoolKeyActivated(
		&bind.FilterOpts{Context: ctx, Start: start, End: &end},
		[][12]byte{epochID},
		[]uint8{keyIndex},
	)
	if err != nil {
		return common.Hash{}, false, fmt.Errorf("filter PoolKeyActivated [%d,%d]: %w", start, end, err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		if err := it.Error(); err != nil {
			return common.Hash{}, false, fmt.Errorf("iterate PoolKeyActivated: %w", err)
		}
		return common.Hash{}, false, nil
	}
	return it.Event.Raw.TxHash, true, nil
}

// managerMethod resolves a DKGManager ABI method once. Its selector and
// input layout are the only things the transcript decoders trust about the
// calldata they are handed.
func managerMethod(name string) func() (abi.Method, error) {
	return sync.OnceValues(func() (abi.Method, error) {
		parsed, err := gtypes.DKGManagerMetaData.GetAbi()
		if err != nil {
			return abi.Method{}, fmt.Errorf("parse DKGManager ABI: %w", err)
		}
		method, ok := parsed.Methods[name]
		if !ok {
			return abi.Method{}, fmt.Errorf("%s missing from DKGManager ABI", name)
		}
		return method, nil
	})
}

var (
	submitContributionMethod = managerMethod("submitContribution")
	activatePoolKeyMethod    = managerMethod("activatePoolKey")
)

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
// never a panic. The returned slice is the contribution.TranscriptWords
// layout documented in DKGManager.submitContribution.
func ContributionTranscript(data []byte) ([]byte, error) {
	return bytesArg(data, submitContributionMethod, "transcript", contribution.TranscriptWords*32)
}

// PoolKeyTranscript extracts the `transcript` argument from raw
// activatePoolKey calldata:
//
//	activatePoolKey(bytes12 epochId, uint8 keyIndex, bytes32 transcriptDigest,
//	    bytes transcript, bytes proof, bytes input)
//
// Same hostile-calldata contract as ContributionTranscript.
func PoolKeyTranscript(data []byte) ([]byte, error) {
	return bytesArg(data, activatePoolKeyMethod, "transcript", poolkey.TranscriptWords*32)
}

// poolKeyPublicInputWords is the size of activatePoolKey's `input`: the 8
// public scalars of the pool-key proof, 32 bytes each.
const poolKeyPublicInputWords = 8

// bytesArg unpacks the `bytes` argument named `name` of `method` from
// calldata and checks it is exactly `size` bytes long. Arguments are looked
// up by name, not position, so an ABI change reorders nothing here.
func bytesArg(data []byte, method func() (abi.Method, error), name string, size int) ([]byte, error) {
	mth, err := method()
	if err != nil {
		return nil, err
	}
	if len(data) < 4 || !bytes.Equal(data[:4], mth.ID) {
		return nil, fmt.Errorf("calldata is not a %s call", mth.Name)
	}
	args, err := mth.Inputs.Unpack(data[4:])
	if err != nil {
		return nil, fmt.Errorf("decode %s calldata: %w", mth.Name, err)
	}
	pos := -1
	for i, input := range mth.Inputs {
		if input.Name == name {
			pos = i
			break
		}
	}
	if pos < 0 || len(args) <= pos {
		return nil, fmt.Errorf("%s calldata has no %s argument", mth.Name, name)
	}
	value, ok := args[pos].([]byte)
	if !ok {
		return nil, fmt.Errorf("%s argument is %T, want []byte", name, args[pos])
	}
	if len(value) != size {
		return nil, fmt.Errorf("%s is %d bytes, want %d", name, len(value), size)
	}
	return value, nil
}

// PoolKeyShareCommitments decodes the committee's share commitments D_p from
// an activatePoolKey transcript. The region is member-indexed: slot i holds
// D_{i+1} for every committee member i < committeeSize, contributing or not
// (the identity beyond). committeeSize is read from the call's public
// inputs, which the contract checked against the epoch policy before the
// transaction could land. The result is the (indexes 1..n, points) pair
// ccommon.ShareCommitmentLeaves takes.
func PoolKeyShareCommitments(data []byte) ([]uint16, []nodetypes.CurvePoint, error) {
	tr, err := PoolKeyTranscript(data)
	if err != nil {
		return nil, nil, err
	}
	input, err := bytesArg(data, activatePoolKeyMethod, "input", poolKeyPublicInputWords*32)
	if err != nil {
		return nil, nil, err
	}
	const nn = poolkey.MaxParticipants
	committeeSize := new(big.Int).SetBytes(input[2*32 : 3*32])
	if committeeSize.Sign() == 0 || !committeeSize.IsUint64() || committeeSize.Uint64() > nn {
		return nil, nil, fmt.Errorf("committee size %s out of range [1, %d]", committeeSize, nn)
	}
	n := int(committeeSize.Uint64())
	word := func(i int) *big.Int { return new(big.Int).SetBytes(tr[i*32 : (i+1)*32]) }
	idxs := make([]uint16, n)
	points := make([]nodetypes.CurvePoint, n)
	for i := range n {
		idxs[i] = uint16(i + 1)
		points[i] = nodetypes.CurvePoint{X: word(4*nn + 2*i), Y: word(4*nn + 2*i + 1)}
	}
	return idxs, points, nil
}

// parseCommitmentPoints extracts every pool key's first t Feldman commitment
// points from the submitContribution calldata transcript. The result is
// indexed by key then coefficient, the shape poolkey.Assignment wants.
func parseCommitmentPoints(data []byte, t uint16) ([][]nodetypes.CurvePoint, error) {
	tr, err := ContributionTranscript(data)
	if err != nil {
		return nil, err
	}
	const nn = contribution.MaxCoefficients
	sets := make([][]nodetypes.CurvePoint, contribution.MaxKeys)
	for j := range contribution.MaxKeys {
		pts := make([]nodetypes.CurvePoint, t)
		for k := range int(t) {
			off := (j*nn + k) * 64
			pts[k] = nodetypes.CurvePoint{
				X: new(big.Int).SetBytes(tr[off : off+32]),
				Y: new(big.Int).SetBytes(tr[off+32 : off+64]),
			}
		}
		sets[j] = pts
	}
	return sets, nil
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
