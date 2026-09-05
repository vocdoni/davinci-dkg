// Package finalizer turns an epoch's accepted contributions into usable keys
// with one on-chain operation (docs/pool-keys-v4.md §11): it reads every
// accepted contribution record at one block, recovers each dealer's compact
// calldata, rebuilds all MaxK pool keys and every committee member's share
// commitments, proves the batched finalization circuit and submits
// `finalizeEpoch`. Live therefore means every key and every share-commitment
// root is stored.
//
// Used by the davinci-dkg-node daemon (via its stagger rotation) and, without
// the proof, by committee members that rebuild their Merkle leaves from the
// contributions instead of trusting the finalization transaction's calldata.
package finalizer

import (
	"bytes"
	"context"
	"errors"
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
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/log"
	gtypes "github.com/vocdoni/davinci-dkg/solidity/golang-types"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
	"github.com/vocdoni/davinci-dkg/web3/txmanager"
)

// EpochPhase values as stored on-chain (DKGTypes.EpochPhase).
const (
	phaseKeyAssembly uint8 = 2
	phaseLive        uint8 = 3
)

// ErrAlreadyLive reports that the epoch went Live under us: another
// finalizer won the race. It is a benign outcome, not a failure to retry.
var ErrAlreadyLive = errors.New("epoch is already Live")

// ErrStale reports that the accepted set the statement was proven over no
// longer matches the chain (a reorg while proving): the proof must not be
// sent and the caller rebuilds from scratch.
var ErrStale = errors.New("finalization statement is stale")

// Result carries the outputs of a successful finalization: every pool key
// and share-commitment root the contract stored, plus the member-indexed
// share commitments those roots were built from (ShareCommitments[j][i] is
// D_{j,i+1}, one per committee member whether it contributed or not).
type Result struct {
	TxHash             common.Hash
	TranscriptDigest   common.Hash
	ParticipantIndexes []uint16
	PoolKeys           []nodetypes.CurvePoint
	ShareRoots         [][32]byte
	ShareCommitments   [][]nodetypes.CurvePoint
	GasUsed            uint64
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

// snapshot is the on-chain state one finalization statement was built from,
// all read at the same block so a contribution landing between two reads
// cannot turn into a missing dealer.
type snapshot struct {
	manager    common.Address
	block      uint64
	epoch      gtypes.IDKGManagerEpoch
	committee  []common.Address
	assignment finalize.Assignment
}

// FinalizeStatement derives, without proving, what `finalizeEpoch` commits
// to: every pool key P_j and the member-indexed share commitments of every
// key (ShareCommitments[j][i] = D_{j,i+1} for i < n, the identity beyond).
// Committee members use it to rebuild their Merkle leaves from the
// contributions' calldata instead of trusting the finalization transaction's
// outer calldata, which a finalization relayed through a contract does not
// even carry. A non-nil `cache` spares the per-dealer event-log rescans.
func FinalizeStatement(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	epochID [12]byte,
	cache CalldataCache,
) (*finalize.PublicInputs, error) {
	snap, err := reconstruct(ctx, c, m, epochID, cache)
	if err != nil {
		return nil, err
	}
	_, pi, err := finalize.BuildWitness(snap.assignment)
	if err != nil {
		return nil, fmt.Errorf("build finalize statement: %w", err)
	}
	return pi, nil
}

// BuildAndSubmitFinalize reads the accepted contributions from on-chain
// calldata, proves the batched finalization and submits `finalizeEpoch`. The
// finalize runtime is loaded from the pinned artifacts on every call;
// long-lived callers pass their own to ProveAndSubmitFinalize.
func BuildAndSubmitFinalize(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	epochID [12]byte,
	cache CalldataCache,
) (*Result, error) {
	runtime, err := finalize.Artifacts.LoadPinned(ctx, &finalize.FinalizeCircuit{})
	if err != nil {
		return nil, fmt.Errorf("load finalize circuit: %w", err)
	}
	return ProveAndSubmitFinalize(ctx, c, m, txm, runtime, epochID, cache)
}

// ProveAndSubmitFinalize is BuildAndSubmitFinalize with a caller-supplied
// finalize runtime (the node loads all four circuits once at startup).
//
// The statement is reconstructed at one block, proven, and the epoch's phase
// and every accepted dealer's record are re-read right before the transaction
// is sent: a finalization that landed meanwhile is reported as ErrAlreadyLive,
// and any other divergence (a reorg changed the accepted set) is ErrStale,
// which the caller retries from scratch.
func ProveAndSubmitFinalize(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	txm *txmanager.Manager,
	runtime *circuits.CircuitRuntime,
	epochID [12]byte,
	cache CalldataCache,
) (*Result, error) {
	if runtime == nil {
		return nil, fmt.Errorf("finalize circuit runtime is required")
	}
	snap, err := reconstruct(ctx, c, m, epochID, cache)
	if err != nil {
		return nil, err
	}
	if snap.epoch.Status == phaseLive {
		return nil, ErrAlreadyLive
	}
	if snap.epoch.Status != phaseKeyAssembly {
		return nil, fmt.Errorf("epoch is in phase %d, not KeyAssembly", snap.epoch.Status)
	}
	if snap.block < snap.epoch.Policy.LiveNotBeforeBlock {
		return nil, fmt.Errorf("finalize gate opens at block %d, head is %d", snap.epoch.Policy.LiveNotBeforeBlock, snap.block)
	}
	if snap.epoch.ContributionCount < snap.epoch.Policy.MinValidContributions {
		return nil, fmt.Errorf("only %d/%d contributions accepted",
			snap.epoch.ContributionCount, snap.epoch.Policy.MinValidContributions)
	}

	witness, pi, err := finalize.BuildWitness(snap.assignment)
	if err != nil {
		return nil, fmt.Errorf("build finalize witness: %w", err)
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
	transcriptScalars, err := pi.TranscriptScalars()
	if err != nil {
		return nil, fmt.Errorf("finalize transcript scalars: %w", err)
	}
	transcriptBytes, err := encodeWords(transcriptScalars...)
	if err != nil {
		return nil, err
	}
	// The Poseidon digest of the transcript is a public input the contract
	// folds into the Fiat–Shamir anchor; it travels as its own argument so
	// the contract never has to hash the transcript with Poseidon itself.
	digest := common.BigToHash(pi.TranscriptDigest)

	// Proving takes a while: re-read the epoch and the accepted records at
	// the head before paying for the transaction. A statement built over a
	// different accepted set (a reorg during the proof) is stale and must
	// be rebuilt, not sent.
	if err := snap.recheck(ctx, c.Client(), epochID); err != nil {
		return nil, err
	}

	auth, err := txm.NewTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := m.FinalizeEpoch(auth, epochID, digest, transcriptBytes, proofBytes, inputBytes)
	if err != nil {
		return nil, err
	}
	txm.RecordPending(tx)
	if err := txm.WaitTxByHash(tx.Hash(), 120*time.Second); err != nil {
		return nil, err
	}
	roots, err := pi.ShareRoots()
	if err != nil {
		return nil, fmt.Errorf("share roots: %w", err)
	}
	n := int(snap.epoch.Policy.CommitteeSize)
	res := &Result{
		TxHash:             tx.Hash(),
		TranscriptDigest:   digest,
		ParticipantIndexes: snap.assignment.ParticipantIndexes,
		PoolKeys:           pi.PoolKeys,
		ShareRoots:         roots[:],
		ShareCommitments:   make([][]nodetypes.CurvePoint, finalize.MaxKeys),
	}
	for j := range finalize.MaxKeys {
		res.ShareCommitments[j] = pi.ShareCommitments[j][:n]
	}
	if receipt, err := c.Client().TransactionReceipt(ctx, tx.Hash()); err == nil {
		res.GasUsed = receipt.GasUsed
		log.Infow("finalizeEpoch tx mined", "tx", tx.Hash().Hex(), "gas", receipt.GasUsed)
	}
	return res, nil
}

// reconstruct reads the epoch, its committee and every contribution record
// at one block, recovers each accepted dealer's compact calldata (through the
// cache when possible), verifies the stored commitmentsHash against the
// decoded commitments and assembles the finalization assignment. Any read or
// decode failure is an error: a dealer silently dropped would prove a
// statement the contract rejects, or with a forged count, the wrong keys.
//
// The reads are pinned to the head observed first and sent as two JSON-RPC
// batches (epoch and committee, then every member's record) rather than
// 2 + n single calls; the pinned block is what makes the batch one snapshot.
func reconstruct(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	epochID [12]byte,
	cache CalldataCache,
) (*snapshot, error) {
	head, err := c.Client().BlockNumber(ctx)
	if err != nil {
		return nil, fmt.Errorf("read head: %w", err)
	}
	block := new(big.Int).SetUint64(head)
	epoch, committee, err := epochAndCommittee(ctx, c.Client(), c.Addresses.Manager, block, epochID)
	if err != nil {
		return nil, fmt.Errorf("at block %d: %w", head, err)
	}
	if epoch.Organizer == (common.Address{}) {
		return nil, fmt.Errorf("epoch %x does not exist", epochID)
	}
	// The accepted set is only frozen once contributions can no longer land:
	// past the finalize gate (which sits after the contribution deadline) or
	// once the epoch is Live.
	if epoch.Status != phaseLive && head < epoch.Policy.LiveNotBeforeBlock {
		return nil, fmt.Errorf("contributions are still open until block %d (head %d)", epoch.Policy.LiveNotBeforeBlock, head)
	}
	t, n := epoch.Policy.Threshold, epoch.Policy.CommitteeSize
	layout, err := contribution.NewLayout(int(t), int(n))
	if err != nil {
		return nil, fmt.Errorf("epoch policy: %w", err)
	}
	if len(committee) != int(n) {
		return nil, fmt.Errorf("committee has %d members, policy says %d", len(committee), n)
	}
	records, err := ContributionRecords(ctx, c.Client(), c.Addresses.Manager, block, epochID, committee)
	if err != nil {
		return nil, fmt.Errorf("at block %d: %w", head, err)
	}

	// Bound the event-log scan to blocks since this epoch was created. The
	// seed block precedes every claim, and contributions only land after
	// the committee filled, so backing off one block from it is conservative.
	startBlock := uint64(0)
	if epoch.SeedBlock > 0 {
		startBlock = epoch.SeedBlock - 1
	}
	roundHash := new(big.Int).SetBytes(epochID[:])

	asgn := finalize.Assignment{
		RoundHash:          roundHash,
		Threshold:          t,
		CommitteeSize:      n,
		ParticipantIndexes: make([]uint16, 0, n),
		Commitments:        make([][][]nodetypes.CurvePoint, 0, n),
		ContributionHashes: make([]*big.Int, 0, n),
	}
	for i, addr := range committee {
		rec := records[i]
		if !rec.Accepted {
			continue
		}
		index := uint16(i + 1)
		if rec.ContributorIndex != index {
			return nil, fmt.Errorf("contribution of %s is stored under index %d, committee position is %d",
				addr, rec.ContributorIndex, index)
		}
		storedHash := new(big.Int).SetBytes(rec.CommitmentsHash[:])
		tr, err := dealerTranscript(ctx, c, m, epochID, addr, index, startBlock, layout, roundHash, storedHash, cache)
		if err != nil {
			return nil, fmt.Errorf("contribution of %s (member %d): %w", addr, index, err)
		}
		asgn.ParticipantIndexes = append(asgn.ParticipantIndexes, index)
		asgn.Commitments = append(asgn.Commitments, tr.Commitments)
		asgn.ContributionHashes = append(asgn.ContributionHashes, storedHash)
	}
	// The contract finalizes over exactly the frozen accepted set; anything
	// else means the chain view is inconsistent.
	if len(asgn.ParticipantIndexes) != int(epoch.ContributionCount) {
		return nil, fmt.Errorf("reconstructed %d accepted contributions, epoch has %d",
			len(asgn.ParticipantIndexes), epoch.ContributionCount)
	}
	if len(asgn.ParticipantIndexes) < int(t) {
		return nil, fmt.Errorf("only %d/%d accepted contributions", len(asgn.ParticipantIndexes), t)
	}
	return &snapshot{manager: c.Addresses.Manager, block: head, epoch: epoch, committee: committee, assignment: asgn}, nil
}

// recheck re-reads the epoch and every committee member's contribution record
// at the head — one JSON-RPC batch through a web3.BatchCaller — and compares
// them with the snapshot the statement was built from. A finalization that
// landed meanwhile is ErrAlreadyLive; a different accepted count, a different
// set of accepted dealers or a dealer whose stored index or commitmentsHash
// changed is ErrStale, so a proof over a reorged accepted set is never sent.
func (s *snapshot) recheck(ctx context.Context, caller bind.ContractCaller, epochID [12]byte) error {
	parsed, err := managerABI()
	if err != nil {
		return err
	}
	epochCall, err := managerCall(parsed, s.manager, "getEpoch", epochID)
	if err != nil {
		return err
	}
	calls := make([]*web3.Call, 0, 1+len(s.committee))
	calls = append(calls, epochCall)
	for _, addr := range s.committee {
		call, err := managerCall(parsed, s.manager, "getContribution", epochID, addr)
		if err != nil {
			return err
		}
		calls = append(calls, call)
	}
	web3.BatchCall(ctx, caller, nil, calls)
	current, err := unpackCall[gtypes.IDKGManagerEpoch](parsed, "getEpoch", epochCall)
	if err != nil {
		return fmt.Errorf("re-read epoch before sending: %w", err)
	}
	if current.Status == phaseLive {
		return ErrAlreadyLive
	}
	if current.Status != phaseKeyAssembly {
		return fmt.Errorf("epoch left KeyAssembly (phase %d) while proving", current.Status)
	}
	if current.ContributionCount != s.epoch.ContributionCount {
		return fmt.Errorf("%w: accepted set changed while proving (%d → %d)",
			ErrStale, s.epoch.ContributionCount, current.ContributionCount)
	}
	accepted := 0
	for i, addr := range s.committee {
		rec, err := unpackCall[gtypes.DKGTypesContributionRecord](parsed, "getContribution", calls[i+1])
		if err != nil {
			return fmt.Errorf("re-read contribution of %s before sending: %w", addr, err)
		}
		if !rec.Accepted {
			continue
		}
		if accepted >= len(s.assignment.ParticipantIndexes) ||
			rec.ContributorIndex != s.assignment.ParticipantIndexes[accepted] ||
			new(big.Int).SetBytes(rec.CommitmentsHash[:]).Cmp(s.assignment.ContributionHashes[accepted]) != 0 {
			return fmt.Errorf("%w: contribution of %s (member %d) changed while proving", ErrStale, addr, i+1)
		}
		accepted++
	}
	if accepted != len(s.assignment.ParticipantIndexes) {
		return fmt.Errorf("%w: %d accepted contributions at the head, statement has %d",
			ErrStale, accepted, len(s.assignment.ParticipantIndexes))
	}
	return nil
}

// dealerTranscript recovers and decodes one accepted dealer's compact
// transcript, checking that its commitments reproduce the commitmentsHash
// the contract stored. The cache is consulted first; an entry that does not
// decode or verify is replaced by a fresh fetch from the event log, so a
// corrupted file can never strand the finalization.
func dealerTranscript(
	ctx context.Context,
	c *web3.Contracts,
	m *gtypes.DKGManager,
	epochID [12]byte,
	dealer common.Address,
	index uint16,
	startBlock uint64,
	layout contribution.Layout,
	roundHash, storedHash *big.Int,
	cache CalldataCache,
) (*contribution.Transcript, error) {
	decode := func(data []byte) (*contribution.Transcript, error) {
		tr, err := DecodeContribution(data, layout)
		if err != nil {
			return nil, err
		}
		got, err := ContributionHash(roundHash, index, layout.Threshold, tr.Commitments)
		if err != nil {
			return nil, err
		}
		if got.Cmp(storedHash) != 0 {
			return nil, fmt.Errorf("commitments hash %s does not match the stored %s", got, storedHash)
		}
		return tr, nil
	}
	if cache != nil {
		if data, ok := cache.Get(epochID, dealer); ok {
			tr, err := decode(data)
			if err == nil {
				return tr, nil
			}
			log.Warnw("cached contribution calldata unusable, refetching", "dealer", dealer, "err", err)
		}
	}
	data, err := ContributionCalldata(ctx, c.Client(), m, epochID, dealer, startBlock)
	if err != nil {
		return nil, fmt.Errorf("calldata: %w", err)
	}
	tr, err := decode(data)
	if err != nil {
		return nil, err
	}
	if cache != nil {
		cache.Put(epochID, dealer, data)
	}
	return tr, nil
}

// managerABI resolves the DKGManager ABI once, for the batched view calls.
var managerABI = sync.OnceValues(func() (*abi.ABI, error) {
	parsed, err := gtypes.DKGManagerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("parse DKGManager ABI: %w", err)
	}
	return parsed, nil
})

// managerCall packs one DKGManager view call for web3.BatchCall.
func managerCall(parsed *abi.ABI, manager common.Address, method string, args ...any) (*web3.Call, error) {
	data, err := parsed.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("pack %s: %w", method, err)
	}
	return &web3.Call{To: manager, Data: data}, nil
}

// unpackCall decodes the single return value of a batched view call into
// the abigen type T, exactly as the generated binding does.
func unpackCall[T any](parsed *abi.ABI, method string, call *web3.Call) (T, error) {
	var zero T
	if call.Err != nil {
		return zero, fmt.Errorf("call %s: %w", method, call.Err)
	}
	out, err := parsed.Unpack(method, call.Output)
	if err != nil {
		return zero, fmt.Errorf("unpack %s: %w", method, err)
	}
	if len(out) != 1 {
		return zero, fmt.Errorf("unpack %s: %d return values, want 1", method, len(out))
	}
	value, ok := abi.ConvertType(out[0], new(T)).(*T)
	if !ok {
		return zero, fmt.Errorf("unpack %s: unexpected return type %T", method, out[0])
	}
	return *value, nil
}

// epochAndCommittee reads an epoch's record and its selected participants at
// `block` (nil = latest) in one batch.
func epochAndCommittee(
	ctx context.Context,
	caller bind.ContractCaller,
	manager common.Address,
	block *big.Int,
	epochID [12]byte,
) (gtypes.IDKGManagerEpoch, []common.Address, error) {
	parsed, err := managerABI()
	if err != nil {
		return gtypes.IDKGManagerEpoch{}, nil, err
	}
	epochCall, err := managerCall(parsed, manager, "getEpoch", epochID)
	if err != nil {
		return gtypes.IDKGManagerEpoch{}, nil, err
	}
	committeeCall, err := managerCall(parsed, manager, "selectedParticipants", epochID)
	if err != nil {
		return gtypes.IDKGManagerEpoch{}, nil, err
	}
	web3.BatchCall(ctx, caller, block, []*web3.Call{epochCall, committeeCall})
	epoch, err := unpackCall[gtypes.IDKGManagerEpoch](parsed, "getEpoch", epochCall)
	if err != nil {
		return gtypes.IDKGManagerEpoch{}, nil, fmt.Errorf("get epoch: %w", err)
	}
	committee, err := unpackCall[[]common.Address](parsed, "selectedParticipants", committeeCall)
	if err != nil {
		return gtypes.IDKGManagerEpoch{}, nil, fmt.Errorf("selected participants: %w", err)
	}
	return epoch, committee, nil
}

// ContributionRecords reads every committee member's contribution record at
// `block` (nil = latest): one JSON-RPC batch through a web3.BatchCaller, one
// eth_call per member otherwise. Entry i is member i+1's record, the zero
// record (Accepted false) for a member that never contributed, exactly as
// getContribution returns it. Used by the finalizer to pin the accepted set
// to one block and by committee members rebuilding their private share.
func ContributionRecords(
	ctx context.Context,
	caller bind.ContractCaller,
	manager common.Address,
	block *big.Int,
	epochID [12]byte,
	committee []common.Address,
) ([]gtypes.DKGTypesContributionRecord, error) {
	parsed, err := managerABI()
	if err != nil {
		return nil, err
	}
	calls := make([]*web3.Call, len(committee))
	for i, addr := range committee {
		if calls[i], err = managerCall(parsed, manager, "getContribution", epochID, addr); err != nil {
			return nil, err
		}
	}
	web3.BatchCall(ctx, caller, block, calls)
	records := make([]gtypes.DKGTypesContributionRecord, len(committee))
	for i, call := range calls {
		if records[i], err = unpackCall[gtypes.DKGTypesContributionRecord](parsed, "getContribution", call); err != nil {
			return nil, fmt.Errorf("get contribution of %s: %w", committee[i], err)
		}
	}
	return records, nil
}

// ContributionHash recomputes the dealer's outer Poseidon commitmentsHash
// from its decoded commitment vectors, exactly as the contribution circuit
// did: every key digest absorbs the vector padded to MaxCoefficients. It is
// what a reader of contribution calldata compares with the dealer's stored
// ContributionRecord.CommitmentsHash before trusting the transcript.
func ContributionHash(roundHash *big.Int, index uint16, threshold int, commitments [][]nodetypes.CurvePoint) (*big.Int, error) {
	if len(commitments) != contribution.MaxKeys {
		return nil, fmt.Errorf("got %d commitment sets, expected %d", len(commitments), contribution.MaxKeys)
	}
	digests := make([]*big.Int, contribution.MaxKeys)
	for j := range contribution.MaxKeys {
		padded, err := ccommon.PadPoints(commitments[j], contribution.MaxCoefficients)
		if err != nil {
			return nil, fmt.Errorf("pad key %d commitments: %w", j, err)
		}
		if digests[j], err = ccommon.CommitmentKeyDigestNative(padded); err != nil {
			return nil, fmt.Errorf("digest key %d commitments: %w", j, err)
		}
	}
	hash, err := ccommon.CommitmentsHashNative(roundHash, big.NewInt(int64(index)), big.NewInt(int64(threshold)), digests)
	if err != nil {
		return nil, fmt.Errorf("commitments hash: %w", err)
	}
	return hash, nil
}

// LogRangeBlocks bounds each eth_getLogs call; most public RPC providers cap
// the range at 10k blocks.
const LogRangeBlocks = 10_000

// chainReader is the slice of the chain client that calldata recovery
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

// FinalizeCalldata locates the finalizeEpoch tx of the epoch via the
// EpochLive event log and returns its raw calldata. The per-member share
// commitments a partial decryption must prove a Merkle path against only
// live there (the contract stores their Merkle roots).
func FinalizeCalldata(
	ctx context.Context,
	client chainReader,
	m *gtypes.DKGManager,
	epochID [12]byte,
	startBlock uint64,
) ([]byte, error) {
	return calldataFromEvent(ctx, client, startBlock, "EpochLive",
		func(ctx context.Context, start, end uint64) (common.Hash, bool, error) {
			return findFinalizeTx(ctx, m, epochID, start, end)
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

// findFinalizeTx returns the hash of the tx that emitted EpochLive(epochID)
// within [start, end].
func findFinalizeTx(
	ctx context.Context,
	m *gtypes.DKGManager,
	epochID [12]byte,
	start, end uint64,
) (common.Hash, bool, error) {
	it, err := m.FilterEpochLive(&bind.FilterOpts{Context: ctx, Start: start, End: &end}, [][12]byte{epochID})
	if err != nil {
		return common.Hash{}, false, fmt.Errorf("filter EpochLive [%d,%d]: %w", start, end, err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		if err := it.Error(); err != nil {
			return common.Hash{}, false, fmt.Errorf("iterate EpochLive: %w", err)
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
	finalizeEpochMethod      = managerMethod("finalizeEpoch")
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
// never a panic. The transcript is compact (docs/pool-keys-v4.md §3): its
// length is a function of the epoch's (t, n), which the caller takes from
// authoritative epoch state and passes as `layout`, never from the calldata.
func ContributionTranscript(data []byte, layout contribution.Layout) ([]byte, error) {
	return bytesArg(data, submitContributionMethod, "transcript", layout.Bytes())
}

// DecodeContribution is ContributionTranscript followed by the layout's
// decoder: the dealer's commitment vectors, the committee region and the
// masked shares, indexed by committee slot (slot i is member i+1).
func DecodeContribution(data []byte, layout contribution.Layout) (*contribution.Transcript, error) {
	transcript, err := ContributionTranscript(data, layout)
	if err != nil {
		return nil, err
	}
	tr, err := layout.DecodeBytes(transcript)
	if err != nil {
		return nil, fmt.Errorf("decode contribution transcript: %w", err)
	}
	return tr, nil
}

// FinalizeTranscript extracts the `transcript` argument from raw
// finalizeEpoch calldata:
//
//	finalizeEpoch(bytes12 epochId, bytes32 transcriptDigest,
//	    bytes transcript, bytes proof, bytes input)
//
// Same hostile-calldata contract as ContributionTranscript. The transcript
// has the fixed layout of docs/pool-keys-v4.md §7.
func FinalizeTranscript(data []byte) ([]byte, error) {
	return bytesArg(data, finalizeEpochMethod, "transcript", finalize.TranscriptWords*32)
}

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

// finalizeCommitteeSize reads the committee size from a finalizeEpoch call's
// public inputs (position 2 of the 7), which the contract checked against
// the epoch policy before the transaction could land.
func finalizeCommitteeSize(data []byte) (int, error) {
	input, err := bytesArg(data, finalizeEpochMethod, "input", finalize.PublicInputWords*32)
	if err != nil {
		return 0, err
	}
	committeeSize := new(big.Int).SetBytes(input[2*32 : 3*32])
	if committeeSize.Sign() == 0 || !committeeSize.IsUint64() || committeeSize.Uint64() > finalize.MaxParticipants {
		return 0, fmt.Errorf("committee size %s out of range [1, %d]", committeeSize, finalize.MaxParticipants)
	}
	return int(committeeSize.Uint64()), nil
}

// FinalizePoolKey decodes pool key `keyIndex` from a finalizeEpoch
// transcript: the value the contract stored as poolKeys[eid][keyIndex].
func FinalizePoolKey(data []byte, keyIndex uint8) (nodetypes.CurvePoint, error) {
	if int(keyIndex) >= finalize.MaxKeys {
		return nodetypes.CurvePoint{}, fmt.Errorf("key %d outside the pool [0, %d)", keyIndex, finalize.MaxKeys)
	}
	tr, err := FinalizeTranscript(data)
	if err != nil {
		return nodetypes.CurvePoint{}, err
	}
	q := finalize.PoolKeyOffset(int(keyIndex))
	return nodetypes.CurvePoint{X: transcriptWord(tr, q), Y: transcriptWord(tr, q+1)}, nil
}

// FinalizeShareCommitments decodes the committee's share commitments of pool
// key `keyIndex` from a finalizeEpoch transcript. The region is
// member-indexed: slot i holds D_{keyIndex,i+1} for every committee member
// i < committeeSize, contributing or not (the identity beyond). The result is
// the (indexes 1..n, points) pair ccommon.ShareCommitmentLeaves takes.
func FinalizeShareCommitments(data []byte, keyIndex uint8) ([]uint16, []nodetypes.CurvePoint, error) {
	if int(keyIndex) >= finalize.MaxKeys {
		return nil, nil, fmt.Errorf("key %d outside the pool [0, %d)", keyIndex, finalize.MaxKeys)
	}
	tr, err := FinalizeTranscript(data)
	if err != nil {
		return nil, nil, err
	}
	n, err := finalizeCommitteeSize(data)
	if err != nil {
		return nil, nil, err
	}
	idxs := make([]uint16, n)
	points := make([]nodetypes.CurvePoint, n)
	for i := range n {
		idxs[i] = uint16(i + 1)
		q := finalize.ShareCommitmentOffset(int(keyIndex), i)
		points[i] = nodetypes.CurvePoint{X: transcriptWord(tr, q), Y: transcriptWord(tr, q+1)}
	}
	return idxs, points, nil
}

// transcriptWord reads 32-byte word q of a transcript whose length the
// caller has already checked.
func transcriptWord(transcript []byte, q int) *big.Int {
	return new(big.Int).SetBytes(transcript[q*32 : (q+1)*32])
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
