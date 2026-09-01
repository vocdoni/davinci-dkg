package node

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/log"
	nodetypes "github.com/vocdoni/davinci-dkg/types"
)

// logRangeBlocks bounds each eth_getLogs call; most public RPC providers cap
// the range at 10k blocks.
const logRangeBlocks = 10_000

// ctKey identifies one ciphertext slot: (epoch, application, index).
type ctKey struct {
	epoch [12]byte
	aid   [32]byte
	idx   uint16
}

func (k ctKey) String() string {
	return fmt.Sprintf("epoch=%x aid=%x idx=%d", k.epoch, k.aid[:4], k.idx)
}

// ciphertext is a decoded CiphertextSubmitted event payload.
type ciphertext struct {
	c1, c2 nodetypes.CurvePoint
}

// scanCiphertexts pulls every CiphertextSubmitted event since the last scan
// (or since head-lookback on first run) into the pending set.
func (n *Node) scanCiphertexts(ctx context.Context) error {
	head, err := n.contracts.Client().BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("head: %w", err)
	}
	from := n.lastCtScan + 1
	if n.lastCtScan == 0 {
		from = 0
		if head > n.lookback {
			from = head - n.lookback
		}
	}
	for start := from; start <= head; start += logRangeBlocks {
		end := min(start+logRangeBlocks-1, head)
		it, err := n.manager.FilterCiphertextSubmitted(&bind.FilterOpts{Context: ctx, Start: start, End: &end}, nil, nil, nil)
		if err != nil {
			return fmt.Errorf("filter CiphertextSubmitted [%d,%d]: %w", start, end, err)
		}
		for it.Next() {
			ev := it.Event
			key := ctKey{ev.EpochId, ev.Aid, ev.CiphertextIndex}
			if _, seen := n.pending[key]; seen {
				continue
			}
			n.pending[key] = &ciphertext{
				c1: nodetypes.CurvePoint{X: new(big.Int).Set(ev.C1x), Y: new(big.Int).Set(ev.C1y)},
				c2: nodetypes.CurvePoint{X: new(big.Int).Set(ev.C2x), Y: new(big.Int).Set(ev.C2y)},
			}
			log.Infow("ciphertext discovered", "ct", key.String(), "block", ev.Raw.BlockNumber)
		}
		err = it.Error()
		_ = it.Close()
		if err != nil {
			return fmt.Errorf("iterate CiphertextSubmitted: %w", err)
		}
		n.lastCtScan = end
	}
	return nil
}

// serviceCiphertexts advances every pending ciphertext and drops the ones
// that need no further work.
func (n *Node) serviceCiphertexts(ctx context.Context) {
	for key, ct := range n.pending {
		done, err := n.serviceCiphertext(ctx, key, ct)
		if err != nil {
			log.Warnw("ciphertext service failed", "ct", key.String(), "err", err)
		}
		if done {
			delete(n.pending, key)
			delete(n.partialDone, key)
		}
	}
}

// serviceCiphertext submits this node's partial decryption (if it sits on
// the committee) and then tries to combine. Returns done=true once the slot
// is combined on-chain or can never be.
func (n *Node) serviceCiphertext(ctx context.Context, key ctKey, ct *ciphertext) (bool, error) {
	if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
		return true, nil
	}
	epoch, err := n.contracts.GetEpoch(ctx, key.epoch)
	if err != nil {
		return false, fmt.Errorf("get epoch: %w", err)
	}
	if epoch.Status != epochLive {
		return epoch.Status == epochAborted, nil
	}
	selected, err := n.selected(ctx, key.epoch)
	if err != nil {
		return false, err
	}
	if idx := myIndex(selected, n.address); idx != 0 && !n.partialDone[key] {
		toxic, err := n.submitPartial(ctx, key, ct, idx, epoch, selected)
		if err != nil {
			return false, err
		}
		if toxic {
			return true, nil
		}
	}
	return n.tryCombine(ctx, key, ct, epoch)
}

// selected caches the committee of a Live epoch (it never changes).
func (n *Node) selected(ctx context.Context, epochID [12]byte) ([]common.Address, error) {
	if s, ok := n.selectedCache[epochID]; ok {
		return s, nil
	}
	s, err := n.contracts.SelectedParticipants(ctx, epochID)
	if err != nil {
		return nil, fmt.Errorf("selected participants: %w", err)
	}
	n.selectedCache[epochID] = s
	return s, nil
}

// submitPartial posts δ_i = d_i·C1 with its DLEQ proof. Returns toxic=true
// when the ciphertext is malformed and must never be decrypted.
func (n *Node) submitPartial(
	ctx context.Context,
	key ctKey,
	ct *ciphertext,
	idx uint16,
	epoch epochView,
	selected []common.Address,
) (bool, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	if rec, err := n.manager.GetPartialDecryption(callOpts, key.epoch, key.aid, idx, key.idx); err == nil && rec.Accepted {
		n.partialDone[key] = true
		return false, nil
	}

	// Refuse small-order / off-curve ciphertexts before touching the share:
	// δ_i = d_i·C1 for a cofactor point would leak d_i mod 8 on-chain. The
	// contract only checks canonical/on-curve/non-identity, so this is the
	// load-bearing subgroup check.
	if err := group.ValidateCiphertext(ct.c1, ct.c2); err != nil {
		log.Warnw("rejecting toxic ciphertext — refusing partial decryption", "ct", key.String(), "err", err)
		return true, nil
	}

	dShare, err := n.buildPrivateShare(ctx, key.epoch, idx, selected, epoch, callOpts)
	if err != nil {
		return false, fmt.Errorf("build private share: %w", err)
	}
	nonce, err := randomScalars(1)
	if err != nil {
		return false, err
	}
	witness, pi, err := partialdecrypt.BuildWitness(partialdecrypt.Assignment{
		RoundHash:        roundScalar(key.epoch),
		Aid:              new(big.Int).SetBytes(key.aid[:]),
		CtIdx:            new(big.Int).SetUint64(uint64(key.idx)),
		ParticipantIndex: idx,
		Base:             ct.c1,
		Secret:           dShare,
		Nonce:            nonce[0],
	})
	if err != nil {
		return false, fmt.Errorf("build partial decrypt witness: %w", err)
	}
	runtime, err := partialdecrypt.Artifacts.LoadOrSetupForCircuit(ctx, &partialdecrypt.PartialDecryptCircuit{})
	if err != nil {
		return false, fmt.Errorf("load partial decrypt circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return false, fmt.Errorf("prove partial decrypt: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return false, err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return false, err
	}
	dHash := ethcrypto.Keccak256Hash(
		common.LeftPadBytes(pi.Delta.X.Bytes(), 32),
		common.LeftPadBytes(pi.Delta.Y.Bytes(), 32),
	)

	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return false, err
	}
	tx, err := n.manager.SubmitPartialDecryption(auth, key.epoch, key.aid, idx, key.idx,
		ct.c1.X, ct.c1.Y, ct.c2.X, ct.c2.Y, dHash, proofBytes, inputBytes)
	if err == nil {
		n.txm.RecordPending(tx)
		err = n.txm.WaitTxByHash(tx.Hash(), 120*time.Second)
	}
	if err != nil {
		reason := decodeContractError(err)
		if strings.Contains(reason, "AlreadyPartiallyDecrypted") {
			n.partialDone[key] = true
			return false, nil
		}
		if isPermanentRevert(err) {
			// Retrying the same proof would fail the same way.
			n.partialDone[key] = true
		}
		return false, fmt.Errorf("submit partial decryption: %s", reason)
	}
	n.partialDone[key] = true
	log.Infow("partial decryption submitted", "ct", key.String(), "index", idx, "tx", tx.Hash().Hex())
	return false, nil
}

// acceptedPartials reads up to `threshold` distinct partials for the slot
// from the event log (the contract only stores their hashes).
func (n *Node) acceptedPartials(ctx context.Context, key ctKey, seedBlock uint64, threshold uint16) ([]uint16, []nodetypes.CurvePoint, error) {
	start := uint64(0)
	if seedBlock > 0 {
		start = seedBlock - 1
	}
	it, err := n.manager.FilterPartialDecryptionSubmitted(
		&bind.FilterOpts{Context: ctx, Start: start},
		[][12]byte{key.epoch}, [][32]byte{key.aid}, nil,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("filter PartialDecryptionSubmitted: %w", err)
	}
	defer func() { _ = it.Close() }()
	seen := map[uint16]bool{}
	var idxs []uint16
	var deltas []nodetypes.CurvePoint
	for it.Next() && len(idxs) < int(threshold) {
		e := it.Event
		if e.CiphertextIndex != key.idx || seen[e.ParticipantIndex] {
			continue
		}
		seen[e.ParticipantIndex] = true
		idxs = append(idxs, e.ParticipantIndex)
		deltas = append(deltas, nodetypes.CurvePoint{X: new(big.Int).Set(e.DeltaX), Y: new(big.Int).Set(e.DeltaY)})
	}
	return idxs, deltas, it.Error()
}

// combineCorrection resolves the per-application correction term:
// mode 0 → T = S·C1 with S the stored derivation tag; mode 1 → T = Δ_org
// from the organizer's submitted share. ready=false while the organizer
// share is still missing.
func (n *Node) combineCorrection(ctx context.Context, key ctKey) (mode uint8, s *big.Int, deltaOrg nodetypes.CurvePoint, ready bool, err error) {
	identity := nodetypes.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	if key.aid == ([32]byte{}) {
		return 0, big.NewInt(0), identity, true, nil
	}
	corr, err := n.appManager.GetCombineCorrection(&bind.CallOpts{Context: ctx}, key.epoch, key.aid, key.idx)
	if err != nil {
		if strings.Contains(decodeContractError(err), "InsufficientPartialDecryptions") {
			return 0, nil, identity, false, nil // organizer share not posted yet
		}
		return 0, nil, identity, false, fmt.Errorf("combine correction: %w", err)
	}
	return corr.Mode, corr.DerivationS, nodetypes.CurvePoint{X: corr.DeltaOrgX, Y: corr.DeltaOrgY}, true, nil
}

// tryCombine interpolates threshold partials, recovers the plaintext by BSGS
// and posts the combine proof. Anyone may combine; the first tx wins.
func (n *Node) tryCombine(ctx context.Context, key ctKey, ct *ciphertext, epoch epochView) (bool, error) {
	threshold := epoch.Policy.Threshold
	idxs, deltas, err := n.acceptedPartials(ctx, key, epoch.SeedBlock, threshold)
	if err != nil {
		return false, err
	}
	if len(idxs) < int(threshold) {
		return false, nil
	}
	mode, s, deltaOrg, ready, err := n.combineCorrection(ctx, key)
	if err != nil || !ready {
		return false, err
	}

	// M·G = C2 − Σ λ_k·δ_k − T
	combinedEnc, err := ccommon.InterpolatePointsAtZeroNative(ccommon.Uint16sToBigInts(idxs), deltas)
	if err != nil {
		return false, fmt.Errorf("interpolate partials: %w", err)
	}
	combined, err := group.Decode(combinedEnc)
	if err != nil {
		return false, err
	}
	c1, err := group.Decode(ct.c1)
	if err != nil {
		return false, err
	}
	c2, err := group.Decode(ct.c2)
	if err != nil {
		return false, err
	}
	correction := group.NewPoint()
	if mode == 0 {
		correction.ScalarMult(c1, s)
	} else {
		correction, err = group.Decode(deltaOrg)
		if err != nil {
			return false, fmt.Errorf("decode organizer share: %w", err)
		}
	}
	negCombined := group.NewPoint()
	negCombined.Neg(combined)
	negCorrection := group.NewPoint()
	negCorrection.Neg(correction)
	mG := group.NewPoint()
	mG.Add(c2, negCombined)
	mG.Add(mG, negCorrection)

	plaintext, err := dlogBSGS(mG)
	if err != nil {
		// Plaintext ≥ MaxDLogPlaintext: retrying can never succeed.
		log.Errorw(err, fmt.Sprintf("combine: dlog failed for %s (plaintext must be < 2^50)", key.String()))
		return true, nil
	}

	witness, pi, err := decryptcombine.BuildWitness(decryptcombine.Assignment{
		RoundHash:          roundScalar(key.epoch),
		Aid:                new(big.Int).SetBytes(key.aid[:]),
		CtIdx:              new(big.Int).SetUint64(uint64(key.idx)),
		Mode:               new(big.Int).SetUint64(uint64(mode)),
		S:                  s,
		DeltaOrg:           deltaOrg,
		Threshold:          threshold,
		CiphertextC1:       ct.c1,
		CiphertextC2:       ct.c2,
		ParticipantIndexes: idxs,
		PartialDecryptions: deltas,
		Plaintext:          plaintext,
	})
	if err != nil {
		return false, fmt.Errorf("build combine witness: %w", err)
	}
	runtime, err := decryptcombine.Artifacts.LoadOrSetupForCircuit(ctx, &decryptcombine.DecryptCombineCircuit{})
	if err != nil {
		return false, fmt.Errorf("load combine circuit: %w", err)
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return false, fmt.Errorf("prove combine: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return false, err
	}
	inputBytes, err := encodePublicWitness(pi.PublicWitness())
	if err != nil {
		return false, err
	}
	transcriptBytes, err := encodeWords(pi.TranscriptScalars()...)
	if err != nil {
		return false, err
	}

	// Every committee member races to combine; re-check right before paying
	// for the tx so a winner that landed while we were proving saves us gas.
	if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
		return true, nil
	}
	auth, err := n.txm.NewTransactOpts(ctx)
	if err != nil {
		return false, err
	}
	tx, err := n.manager.CombineDecryption(auth, key.epoch, key.aid, key.idx,
		common.BigToHash(pi.CombineHash), pi.PlaintextHash, transcriptBytes, proofBytes, inputBytes)
	if err == nil {
		n.txm.RecordPending(tx)
		err = n.txm.WaitTxByHash(tx.Hash(), 120*time.Second)
	}
	if err != nil {
		reason := decodeContractError(err)
		if strings.Contains(reason, "AlreadyCombined") {
			return true, nil
		}
		// A mined-but-reverted tx is almost always a lost race; the
		// on-chain re-check at the top of the next attempt settles it.
		if rec, err := n.contracts.GetCombinedDecryption(ctx, key.epoch, key.aid, key.idx); err == nil && rec.Completed {
			return true, nil
		}
		if isPermanentRevert(err) {
			return true, errors.New("combine rejected on-chain, giving up: " + reason)
		}
		return false, fmt.Errorf("submit combine: %w", err)
	}
	log.Infow("decryption combined", "ct", key.String(), "plaintext", plaintext.String(), "tx", tx.Hash().Hex())
	return true, nil
}
