package helpers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
	"github.com/vocdoni/davinci-dkg/web3"
)

const (
	roundStatusContribution uint8 = 2
	// roundStatusLive is the phase finalizeEpoch moves the epoch to, with
	// every pool key stored; applications can only register there.
	roundStatusLive uint8 = 3
)

// FinalizedRoundResult is a Live epoch plus everything a test needs to use
// its pool and to reproduce the members' shares.
type FinalizedRoundResult struct {
	EpochID            [12]byte
	Epoch              web3.EpochView
	RoundHash          *big.Int
	Participant        common.Address
	Threshold          uint16
	CommitteeSize      uint16
	ParticipantIndexes []uint16
	// Contributions is indexed by accepted contributor, then pool key, then
	// coefficient — the shape BuildFinalizeSubmission takes.
	Contributions [][][]*big.Int
	// Finalize is the finalization that made the epoch Live: every pool key
	// and every key's share-commitment tree.
	Finalize *FinalizeSubmission
}

// PoolKey is P_j of the round's pool.
func (r *FinalizedRoundResult) PoolKey(keyIndex uint8) types.CurvePoint {
	return r.Finalize.PoolKey(keyIndex)
}

// Shares is key j's share-commitment tree, the root the contract stores and
// the tree a partial decryption proves its leaf against.
func (r *FinalizedRoundResult) Shares(keyIndex uint8) ShareTree {
	return r.Finalize.ShareTree(keyIndex)
}

// ParticipantShare returns the member's share of pool key `keyIndex`,
// d = Σ_c f_{c,keyIndex}(participantIndex) over the accepted contributions.
// Every key of the pool is dealt from its own polynomial, so a share is only
// valid for the key it was evaluated on: a partial built from key 0's share
// fails the share-root check of an application registered against key 1.
func (r *FinalizedRoundResult) ParticipantShare(keyIndex uint8, participantIndex uint16) (*big.Int, error) {
	shares, err := RecoverParticipantShares(r.Contributions, keyIndex, []uint16{participantIndex})
	if err != nil {
		return nil, fmt.Errorf("recover participant %d share of pool key %d: %w", participantIndex, keyIndex, err)
	}
	return shares[0], nil
}

// DefaultLotteryAlphaBps is the over-subscription factor applied to integration
// test epoch policies when the caller leaves LotteryAlphaBps at zero. With
// BootstrappedNodeKeys active operators, α·n ≥ registered for every n ≥ 1, so
// every test actor is lottery-eligible and committees are deterministic.
const DefaultLotteryAlphaBps uint16 = 65535

// DefaultSeedDelay matches the on-chain `SEED_DELAY_BLOCKS` constant in
// `solidity/src/libraries/Sizes.sol`. Kept as a Go-side constant for tests
// that need to advance past the seed block before claiming slots.
const DefaultSeedDelay uint64 = 1

func CreateContributionRound(ctx context.Context, services *TestServices, policy types.EpochPolicy) ([12]byte, error) {
	var zero [12]byte

	if policy.LotteryAlphaBps == 0 {
		policy.LotteryAlphaBps = DefaultLotteryAlphaBps
	}
	if err := policy.Validate(); err != nil {
		return zero, err
	}

	epochID, err := CreateEpoch(ctx, services, policy)
	if err != nil {
		return zero, err
	}

	// Lottery flow: advance past seedBlock so blockhash is available, then claim.
	head, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return zero, fmt.Errorf("get block number: %w", err)
	}
	seedBlock := head + DefaultSeedDelay
	if head <= seedBlock {
		if err := MineBlocks(ctx, services, seedBlock-head+1); err != nil {
			return zero, err
		}
	}
	if err := ClaimSlot(ctx, services, epochID); err != nil {
		return zero, err
	}
	if _, err := WaitEpochPhase(ctx, services, epochID, roundStatusContribution); err != nil {
		return zero, err
	}

	return epochID, nil
}

// CreateFinalizedSingleParticipantRound drives a one-member epoch to Live with
// its whole pool stored. `coefficients` is the single fixture polynomial; it
// becomes key 0 of the dealt pool verbatim (see DealPoolCoefficients), so the
// share of participant 1 under key 0 stays sum(coefficients).
func CreateFinalizedSingleParticipantRound(
	ctx context.Context,
	services *TestServices,
	policy types.EpochPolicy,
	coefficients []*big.Int,
) (*FinalizedRoundResult, error) {
	epochID, err := CreateContributionRound(ctx, services, policy)
	if err != nil {
		return nil, err
	}

	pool := DealPoolCoefficients(coefficients)
	submission, err := BuildContributionSubmission(ctx, services, epochID, 1, 1, 1, pool, []uint16{1})
	if err != nil {
		return nil, err
	}
	self := SelfActor(services)
	if err := SubmitContributionAs(ctx, self, epochID, 1, submission); err != nil {
		return nil, err
	}

	// Wait until block.number >= liveNotBeforeBlock so the on-chain gate is open.
	if err := WaitForFinalizeGate(ctx, services, epochID); err != nil {
		return nil, err
	}
	finalization, err := BuildFinalizeSubmission(ctx, epochID, 1, 1, []uint16{1}, [][][]*big.Int{pool})
	if err != nil {
		return nil, fmt.Errorf("build finalization: %w", err)
	}
	if err := FinalizeEpochAs(ctx, self, epochID, finalization); err != nil {
		return nil, err
	}
	epoch, err := WaitEpochPhase(ctx, services, epochID, roundStatusLive)
	if err != nil {
		return nil, err
	}

	return &FinalizedRoundResult{
		EpochID:            epochID,
		Epoch:              epoch,
		RoundHash:          submission.RoundHash,
		Participant:        services.TxManager.Address(),
		Threshold:          1,
		CommitteeSize:      1,
		ParticipantIndexes: []uint16{1},
		Contributions:      [][][]*big.Int{pool},
		Finalize:           finalization,
	}, nil
}

// SelfActor wraps the harness' own signer as a TestActor. PrivKey is left
// empty on purpose: the harness key comes from the environment in external
// mode, so deriving a node key from it here would be a guess. Tests that
// need one use services.Actor(i).
func SelfActor(services *TestServices) *TestActor {
	return &TestActor{
		Contracts: services.Contracts,
		Manager:   services.Manager,
		Registry:  services.Registry,
		TxManager: services.TxManager,
	}
}

// CombineSingleParticipantDecryption drives the partial decryption and the
// combine for a ciphertext that is already on chain at
// (epochID, aid, ciphertextIndex), assuming the epoch was created by
// CreateFinalizedSingleParticipantRound (committee=1, threshold=1, single
// participant index 1 owned by services.TxManager).
//
// `share` is the member's share of the application's pool key — for a single
// coefficient list `coefficients` under key 0, that is f(1) = sum(coefficients).
// `organizerSecret` is the revealed sk_org of an organizer-locked
// application, or 0 for an automatic one.
//
// Used by the SDK end-to-end ciphertext test (sdk/tests/ciphertext-e2e).
func CombineSingleParticipantDecryption(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	share *big.Int,
	organizerSecret *big.Int,
) error {
	payload, err := PrepareSingleParticipantCombinePayload(ctx, services, epochID, aid, ciphertextIndex, share, organizerSecret)
	if err != nil {
		return err
	}
	auth, err := services.TxManager.NewTransactOpts(ctx)
	if err != nil {
		return err
	}
	tx, err := services.Manager.CombineDecryption(
		auth, epochID, aid, ciphertextIndex,
		payload.CombineHash, payload.Plaintext,
		payload.Transcript, payload.Proof, payload.Input,
	)
	if err != nil {
		return fmt.Errorf("combine decryption: %w", err)
	}
	return services.TxManager.WaitTxByHash(tx.Hash(), DefaultTxTimeout)
}

// CombinePayload bundles the bytes a caller needs to invoke
// `combineDecryption` on-chain: the keccak hash binding the participant set
// + deltas, the recovered plaintext, and the verifier transcript / proof /
// public-input blobs. Returned by `PrepareSingleParticipantCombinePayload`
// so callers (including the SDK) can drive the on-chain combine themselves.
type CombinePayload struct {
	CombineHash [32]byte
	Plaintext   *big.Int
	Transcript  []byte
	Proof       []byte
	Input       []byte
}

// PrepareSingleParticipantCombinePayload submits the participant-1 partial
// decryption on chain (the gnark proof has to be built in Go) and returns the
// bytes a caller needs to drive `combineDecryption` themselves. Used by the
// SDK ciphertext-e2e and combine-e2e tests so the SDK writer is what actually
// issues the on-chain combine — the same code path production node operators
// take.
//
// There is no organizer share any more, and the contract refuses every
// partial of an organizer-locked application until `revealOrganizerSecret`
// (`OrganizerSecretNotRevealed`). The reveal is therefore a precondition of
// this helper: when the application is still sealed, `organizerSecret` is
// published first, before any partial is built; when it was already
// revealed, it must match the on-chain value.
func PrepareSingleParticipantCombinePayload(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	share *big.Int,
	organizerSecret *big.Int,
) (*CombinePayload, error) {
	if err := ensureOrganizerRevealed(ctx, services, epochID, aid, organizerSecret); err != nil {
		return nil, err
	}
	c1, c2, err := ciphertextFromLog(ctx, services, epochID, aid, ciphertextIndex)
	if err != nil {
		return nil, err
	}

	// A one-member committee has one leaf, so the tree is fully determined by
	// this share: D_1 = share·G. The contract checks that leaf against the
	// share root of the application's pool key, so make sure the caller
	// handed us the share of that key and not of another one — the revert
	// (`InvalidProofInput`) would not say which input was wrong.
	tree, err := ShareTreeFromShares([]uint16{1}, []*big.Int{share})
	if err != nil {
		return nil, err
	}
	if err := checkShareRoot(ctx, services, epochID, aid, tree); err != nil {
		return nil, err
	}

	// Build the partial decryption (delta = share·C1) using the gnark proof.
	// For a single-participant committee the partial decryption IS the combined
	// decryption (Lagrange interpolation at zero with a single share is the
	// share itself), but we still go through both txs so this exercises the
	// full on-chain decryption path the SDK consumers depend on.
	const partialNonce = 1
	partial, err := BuildPartialDecryptionSubmissionFromBase(
		ctx, epochID, aid, ciphertextIndex, 1, c1, c2, share, big.NewInt(partialNonce), tree,
	)
	if err != nil {
		return nil, fmt.Errorf("build partial decryption: %w", err)
	}

	// SubmitPartialDecryption is idempotent at the contract level only when
	// no record exists yet; re-running this helper after a successful
	// partial submission would revert. Callers driving combine via the SDK
	// run prepare-combine exactly once per ciphertext.
	if err := SubmitPartialDecryptionAs(ctx, SelfActor(services), epochID, aid, 1, ciphertextIndex, partial); err != nil {
		return nil, err
	}

	deltaOrg, err := OrganizerTerm(organizerSecret, c1)
	if err != nil {
		return nil, err
	}
	// Recover the plaintext by brute-force discrete log over a small window.
	// Fixture epochs always submit ciphertexts of small integers so this
	// terminates immediately; the helper bounds the search to 2^20 and is
	// deliberately separate from the production node's BSGS (node/dlog.go,
	// cap 2^50). Keeping a tiny cap here means tests don't pay the ~30 s
	// table-build cost the production cap implies.
	plaintext, err := bruteForceELGamalPlaintext(c2, partial.Delta, deltaOrg)
	if err != nil {
		return nil, fmt.Errorf("recover plaintext: %w", err)
	}

	combineOutput, err := BuildDecryptCombineOutputFromCiphertext(
		ctx, epochID, aid, ciphertextIndex, 1, c1, c2,
		OrganizerKey(organizerSecret), organizerSecret,
		[]uint16{1}, []types.CurvePoint{partial.Delta}, plaintext,
	)
	if err != nil {
		return nil, fmt.Errorf("build combine output: %w", err)
	}

	return &CombinePayload{
		CombineHash: combineOutput.CombineHash,
		Plaintext:   combineOutput.Plaintext,
		Transcript:  combineOutput.Transcript,
		Proof:       combineOutput.Proof,
		Input:       combineOutput.Input,
	}, nil
}

// ensureOrganizerRevealed makes the reveal precede the first partial of an
// organizer-locked application. Automatic applications have nothing to
// reveal, and a non-zero `organizerSecret` for one is a caller mix-up.
func ensureOrganizerRevealed(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	aid [32]byte,
	organizerSecret *big.Int,
) error {
	app, err := services.AppManager.GetApplication(services.CallOpts(ctx), epochID, aid)
	if err != nil {
		return fmt.Errorf("get application: %w", err)
	}
	if !app.Exists {
		return fmt.Errorf("application %x is not registered in epoch %x", aid, epochID)
	}
	secretGiven := organizerSecret != nil && organizerSecret.Sign() != 0
	if app.Policy.Mode == uint8(types.AppModeAutomatic) {
		if secretGiven {
			return fmt.Errorf("application %x is automatic: it has no organizer secret", aid)
		}
		return nil
	}
	if app.OrganizerSecret.Sign() != 0 {
		if secretGiven && app.OrganizerSecret.Cmp(organizerSecret) != 0 {
			return fmt.Errorf("application %x revealed a different organizer secret than the one given", aid)
		}
		return nil
	}
	if !secretGiven {
		return fmt.Errorf("application %x is organizer-locked and sealed: "+
			"the organizer secret is required, partials are refused before the reveal", aid)
	}
	if err := RevealOrganizerSecretAs(ctx, SelfActor(services), services.AppManager, epochID, aid, organizerSecret); err != nil {
		return fmt.Errorf("reveal organizer secret before the first partial: %w", err)
	}
	return nil
}

// checkShareRoot verifies that `tree` reproduces the share root the contract
// stores for the pool key the application claimed, i.e. that the shares it
// was built from belong to that key.
func checkShareRoot(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	aid [32]byte,
	tree ShareTree,
) error {
	keyIndex, err := services.Manager.GetAppPoolIndex(services.CallOpts(ctx), epochID, aid)
	if err != nil {
		return fmt.Errorf("get application pool index: %w", err)
	}
	root, err := services.Manager.GetPoolShareRoot(services.CallOpts(ctx), epochID, keyIndex)
	if err != nil {
		return fmt.Errorf("get pool key %d share root: %w", keyIndex, err)
	}
	if tree.Root() != root {
		return fmt.Errorf("share tree root does not match the on-chain root of pool key %d: "+
			"the share is not the member's share of the application's pool key", keyIndex)
	}
	return nil
}

// ciphertextFromLog recovers the ciphertext coordinates from the
// CiphertextSubmitted event (the contract only stores their keccak hash).
func ciphertextFromLog(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
) (c1, c2 types.CurvePoint, err error) {
	ciphertextHash, err := services.Manager.GetCiphertextHash(services.CallOpts(ctx), epochID, aid, ciphertextIndex)
	if err != nil {
		return c1, c2, fmt.Errorf("get ciphertext hash: %w", err)
	}
	var zero common.Hash
	if ciphertextHash == zero {
		return c1, c2, fmt.Errorf("ciphertext at (%x, %x, %d) not yet submitted", epochID, aid, ciphertextIndex)
	}

	// Scan from the epoch's seedBlock to limit the filter range.
	epoch, err := services.Contracts.GetEpoch(ctx, epochID)
	if err != nil {
		return c1, c2, fmt.Errorf("get epoch: %w", err)
	}
	startBlock := uint64(0)
	if epoch.SeedBlock > 0 {
		startBlock = epoch.SeedBlock - 1
	}
	latest, err := services.Contracts.Client().BlockNumber(ctx)
	if err != nil {
		return c1, c2, fmt.Errorf("read head: %w", err)
	}
	filterOpts := &bind.FilterOpts{Context: ctx, Start: startBlock, End: &latest}
	it, err := services.Manager.FilterCiphertextSubmitted(
		filterOpts, [][12]byte{epochID}, [][32]byte{aid}, []uint16{ciphertextIndex},
	)
	if err != nil {
		return c1, c2, fmt.Errorf("filter CiphertextSubmitted: %w", err)
	}
	defer func() { _ = it.Close() }()
	if !it.Next() {
		if err := it.Error(); err != nil {
			return c1, c2, fmt.Errorf("iterate CiphertextSubmitted: %w", err)
		}
		return c1, c2, fmt.Errorf("no CiphertextSubmitted event for (%x, %x, %d)", epochID, aid, ciphertextIndex)
	}
	c1 = types.CurvePoint{X: new(big.Int).Set(it.Event.C1x), Y: new(big.Int).Set(it.Event.C1y)}
	c2 = types.CurvePoint{X: new(big.Int).Set(it.Event.C2x), Y: new(big.Int).Set(it.Event.C2y)}
	return c1, c2, nil
}

// bruteForceELGamalPlaintext recovers m from c2 − δ − Δ_org = m·G by trying
// every m in [0, 2^20). Used by CombineSingleParticipantDecryption to
// discover the plaintext that the SDK encrypted (the SDK chose a random k, so
// this is the only way the fixture can learn what was sent without
// round-tripping through the original encryption). Δ_org is the identity for
// an automatic application.
//
// Production decryption uses node/dlog.go (BSGS, cap 2^50); this helper
// deliberately keeps the cheaper linear scan because every fixture submits
// values well under 2^20.
//
// Note on the loop: gnark's PointAffine zero value is the affine origin (0, 0),
// which is NOT a point on twisted Edwards (the identity is (0, 1)). We can't
// start from `group.NewPoint()` and add G repeatedly, because adding (0, 0) + G
// produces an invalid result. Instead we use the identity encoded as (0, 1)
// for the m=0 check and start the iteration from G itself for m=1+.
func bruteForceELGamalPlaintext(c2, delta, deltaOrg types.CurvePoint) (*big.Int, error) {
	c2Native, err := group.Decode(c2)
	if err != nil {
		return nil, fmt.Errorf("decode c2: %w", err)
	}
	target := group.NewPoint()
	target.Set(c2Native)
	for _, term := range []types.CurvePoint{delta, deltaOrg} {
		native, err := group.Decode(term)
		if err != nil {
			return nil, fmt.Errorf("decode subtrahend: %w", err)
		}
		neg := group.NewPoint()
		neg.Neg(native)
		target.Add(target, neg)
	}
	targetEnc := group.Encode(target)

	// m = 0 → target is the curve identity (0, 1)
	if targetEnc.X.Sign() == 0 && targetEnc.Y.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(0), nil
	}

	// m >= 1: candidate starts at G, iterate candidate += G.
	g := group.Generator()
	candidate := group.NewPoint()
	candidate.Set(g)
	for i := int64(1); i < 1<<20; i++ {
		candEnc := group.Encode(candidate)
		if candEnc.X.Cmp(targetEnc.X) == 0 && candEnc.Y.Cmp(targetEnc.Y) == 0 {
			return big.NewInt(i), nil
		}
		next := group.NewPoint()
		next.Add(candidate, g)
		candidate = next
	}
	return nil, fmt.Errorf("plaintext out of brute-force range (> 2^20)")
}
