package helpers

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"sync"

	gnec "github.com/consensys/gnark-crypto/ecc"
	groth16backend "github.com/consensys/gnark/backend/groth16"
	groth16bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/frontend"
	"github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/vocdoni/davinci-dkg/circuits"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/circuits/contribution"
	"github.com/vocdoni/davinci-dkg/circuits/decryptcombine"
	"github.com/vocdoni/davinci-dkg/circuits/finalize"
	"github.com/vocdoni/davinci-dkg/circuits/partialdecrypt"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/types"
)

type ContributionSubmission struct {
	Proof               []byte
	Input               []byte
	Transcript          []byte
	CommitmentsHash     [32]byte
	EncryptedSharesHash [32]byte
	RoundHash           *big.Int
}

// FinalizeSubmission is everything `finalizeEpoch` needs, plus the pool keys
// and share commitments the callers turn into Merkle paths for the partial
// decryptions submitted against those keys.
//
// ParticipantIndexes lists the accepted contributors (the transcript's
// active rows). ShareCommitments[j] is member-indexed: entry i is D_{j,i+1}
// for every committee member i < committeeSize, contributing or not, because
// a member that only claimed a slot still received a share of every accepted
// polynomial. Shares[j] is the Merkle tree over exactly those leaves, the
// root the contract stores for key j.
type FinalizeSubmission struct {
	Proof              []byte
	Input              []byte
	Transcript         []byte
	TranscriptDigest   [32]byte
	RoundHash          *big.Int
	ParticipantIndexes []uint16
	PoolKeys           []types.CurvePoint
	ShareCommitments   [][]types.CurvePoint
	Shares             []ShareTree
}

// PoolKey is P_j, the committee key an application registered against key j
// encrypts under (plus PK_org when it is organizer-locked).
func (s *FinalizeSubmission) PoolKey(keyIndex uint8) types.CurvePoint {
	return s.PoolKeys[keyIndex]
}

// ShareTree is key j's share-commitment tree.
func (s *FinalizeSubmission) ShareTree(keyIndex uint8) ShareTree {
	return s.Shares[keyIndex]
}

type PartialDecryptionSubmission struct {
	Proof     []byte
	Input     []byte
	DeltaHash [32]byte
	Delta     types.CurvePoint
	RoundHash *big.Int
	// ShareProof is the MerkleDepth-long sibling path proving the member's
	// share commitment against the pool key's root.
	ShareProof [][32]byte
	// C1, C2 are the on-chain ciphertext coords the proof binds to,
	// captured here so SubmitPartialDecryption callers can pass them
	// straight through. Set by
	// BuildPartialDecryptionSubmissionFromBase from the caller's `base`.
	C1 types.CurvePoint
	C2 types.CurvePoint
}

type DecryptCombineOutput struct {
	Proof        []byte
	Input        []byte
	Transcript   []byte
	CombineHash  [32]byte
	Plaintext    *big.Int
	CiphertextC1 types.CurvePoint
	CiphertextC2 types.CurvePoint
	// OrganizerPK is the key the contract checks the transcript's w[4..5]
	// against: sk_org·G for an organizer-locked application, the identity
	// (0, 1) for an automatic one.
	OrganizerPK types.CurvePoint
}

var (
	contributionRuntimeOnce   sync.Once
	contributionRuntime       *circuits.CircuitRuntime
	contributionRuntimeErr    error
	finalizeRuntimeOnce       sync.Once
	finalizeRuntime           *circuits.CircuitRuntime
	finalizeRuntimeErr        error
	partialDecryptRuntimeOnce sync.Once
	partialDecryptRuntime     *circuits.CircuitRuntime
	partialDecryptRuntimeErr  error
	decryptCombineRuntimeOnce sync.Once
	decryptCombineRuntime     *circuits.CircuitRuntime
	decryptCombineRuntimeErr  error
)

// BuildContributionSubmission proves one contribution. `coefficients` holds
// the MaxK polynomials the contributor deals, key-major: use
// DealPoolCoefficients to expand a single fixture polynomial into a full set.
func BuildContributionSubmission(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	threshold uint16,
	committeeSize uint16,
	contributorIndex uint16,
	coefficients [][]*big.Int,
	recipientIndexes []uint16,
) (*ContributionSubmission, error) {
	roundHash := RoundScalar(epochID)
	recipientKeys, encryptionNonces, err := contributionRecipients(ctx, services, epochID, recipientIndexes)
	if err != nil {
		return nil, err
	}
	assignment := contribution.Assignment{
		RoundHash:        roundHash,
		Threshold:        threshold,
		CommitteeSize:    committeeSize,
		ContributorIndex: contributorIndex,
		Coefficients:     coefficients,
		RecipientIndexes: recipientIndexes,
		RecipientKeys:    recipientKeys,
		EncryptionNonces: encryptionNonces,
	}
	witness, publicInputs, err := contribution.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}

	runtime, err := loadContributionRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove contribution: %w", err)
	}

	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicAssignment(publicInputs.PublicWitness())
	if err != nil {
		return nil, err
	}
	transcriptScalars, err := publicInputs.TranscriptScalars()
	if err != nil {
		return nil, fmt.Errorf("contribution transcript scalars: %w", err)
	}
	transcriptBytes, err := encodeSolidityWords(transcriptScalars...)
	if err != nil {
		return nil, err
	}

	return &ContributionSubmission{
		Proof:               proofBytes,
		Input:               inputBytes,
		Transcript:          transcriptBytes,
		CommitmentsHash:     common.BigToHash(publicInputs.CommitmentHash),
		EncryptedSharesHash: common.BigToHash(publicInputs.ShareHash),
		RoundHash:           new(big.Int).Set(roundHash),
	}, nil
}

func contributionRecipients(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	recipientIndexes []uint16,
) ([]types.NodeKey, []*big.Int, error) {
	participants, err := services.Contracts.SelectedParticipants(ctx, epochID)
	if err != nil {
		return nil, nil, fmt.Errorf("selected participants: %w", err)
	}

	keys := make([]types.NodeKey, 0, len(recipientIndexes))
	nonces := make([]*big.Int, 0, len(recipientIndexes))
	for _, recipientIndex := range recipientIndexes {
		if recipientIndex == 0 || int(recipientIndex) > len(participants) {
			return nil, nil, fmt.Errorf("recipient index %d out of range", recipientIndex)
		}
		node, err := services.Contracts.GetNode(ctx, participants[recipientIndex-1])
		if err != nil {
			return nil, nil, fmt.Errorf("get node %d: %w", recipientIndex, err)
		}
		keys = append(keys, types.NodeKey{
			Operator: node.Operator,
			PubX:     node.PubX,
			PubY:     node.PubY,
		})
		nonces = append(nonces, big.NewInt(int64(1000+recipientIndex)))
	}
	return keys, nonces, nil
}

// BuildFinalizeSubmission proves `finalizeEpoch` for the epoch: every pool
// key and every committee member's share commitment of every key, from the
// accepted contributions. `contributions` is indexed by accepted contributor
// (aligned with participantIndexes), then pool key, then coefficient.
func BuildFinalizeSubmission(
	ctx context.Context,
	epochID [12]byte,
	threshold uint16,
	committeeSize uint16,
	participantIndexes []uint16,
	contributions [][][]*big.Int,
) (*FinalizeSubmission, error) {
	return buildFinalizeSubmission(ctx, epochID, threshold, committeeSize, participantIndexes, contributions, nil)
}

// BuildFinalizeSubmissionWithNonCanonicalWord is BuildFinalizeSubmission with
// transcript word `wordIndex` published in calldata as `w + p` (p the BN254
// scalar field), which the BRLC commitment cannot tell from `w`. The proof is
// otherwise honest for that calldata: the Fiat–Shamir challenge is derived
// from the non-canonical bytes exactly as the contract derives it, and the
// circuit's transcript commitment over the reduced words equals the
// contract's over the raw ones. Only BRLC's canonical-word check stands
// between such a transcript and storage — the contract reads the pool keys
// and the share commitments straight from calldata — so it must revert
// `TranscriptWordNotInField`. Test-only by construction.
func BuildFinalizeSubmissionWithNonCanonicalWord(
	ctx context.Context,
	epochID [12]byte,
	threshold uint16,
	committeeSize uint16,
	participantIndexes []uint16,
	contributions [][][]*big.Int,
	wordIndex int,
) (*FinalizeSubmission, error) {
	if wordIndex < 0 || wordIndex >= finalize.TranscriptWords {
		return nil, fmt.Errorf("transcript word %d out of range [0, %d)", wordIndex, finalize.TranscriptWords)
	}
	lift := func(words []*big.Int) {
		words[wordIndex] = new(big.Int).Add(words[wordIndex], gnec.BN254.ScalarField())
	}
	return buildFinalizeSubmission(ctx, epochID, threshold, committeeSize, participantIndexes, contributions, lift)
}

// buildFinalizeSubmission builds the witness, optionally rewrites the calldata
// transcript words (re-deriving the challenge and the commitment the way the
// contract will), proves, and lays out the member-indexed share trees.
func buildFinalizeSubmission(
	ctx context.Context,
	epochID [12]byte,
	threshold uint16,
	committeeSize uint16,
	participantIndexes []uint16,
	contributions [][][]*big.Int,
	rewrite func(words []*big.Int),
) (*FinalizeSubmission, error) {
	roundHash := RoundScalar(epochID)
	commitments, err := commitmentSets(contributions)
	if err != nil {
		return nil, err
	}
	assignment := finalize.Assignment{
		RoundHash:          roundHash,
		Threshold:          threshold,
		CommitteeSize:      committeeSize,
		ParticipantIndexes: participantIndexes,
		Commitments:        commitments,
	}
	witness, publicInputs, err := finalize.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}
	transcriptScalars, err := publicInputs.TranscriptScalars()
	if err != nil {
		return nil, fmt.Errorf("finalize transcript scalars: %w", err)
	}
	if rewrite != nil {
		rewrite(transcriptScalars)
		// The contract anchors ρ on keccak(digest ‖ keccak(calldata)), so a
		// different calldata encoding means a different challenge — and a
		// different (but still consistent) commitment in the witness.
		anchor, anchorErr := ccommon.ChallengeAnchor(transcriptScalars, publicInputs.TranscriptDigest)
		if anchorErr != nil {
			return nil, fmt.Errorf("finalize challenge anchor: %w", anchorErr)
		}
		challenge, challengeErr := ccommon.DeriveChallengeNative(roundHash, protocol.DomainFinalizeTranscriptV2, anchor)
		if challengeErr != nil {
			return nil, fmt.Errorf("finalize challenge: %w", challengeErr)
		}
		commitment, commitErr := publicInputs.BRLCCommitment(challenge)
		if commitErr != nil {
			return nil, fmt.Errorf("finalize transcript commitment: %w", commitErr)
		}
		witness.Challenge, witness.TranscriptCommitment = challenge, commitment
		publicInputs.Challenge, publicInputs.TranscriptCommitment = challenge, commitment
	}

	runtime, err := loadFinalizeRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove finalization: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicAssignment(publicInputs.PublicWitness())
	if err != nil {
		return nil, err
	}
	transcriptBytes, err := encodeSolidityWords(transcriptScalars...)
	if err != nil {
		return nil, err
	}

	// Leaves for the whole committee, members 1..n, exactly as the contract
	// rebuilds every key's tree from the transcript's share rows.
	shareCommitments := make([][]types.CurvePoint, ccommon.MaxK)
	trees := make([]ShareTree, ccommon.MaxK)
	for j := range ccommon.MaxK {
		shareCommitments[j] = publicInputs.ShareCommitments[j][:committeeSize]
		if trees[j], err = CommitteeShareTree(shareCommitments[j]); err != nil {
			return nil, fmt.Errorf("key %d share tree: %w", j, err)
		}
	}

	return &FinalizeSubmission{
		Proof:              proofBytes,
		Input:              inputBytes,
		Transcript:         transcriptBytes,
		TranscriptDigest:   common.BigToHash(publicInputs.TranscriptDigest),
		RoundHash:          new(big.Int).Set(roundHash),
		ParticipantIndexes: participantIndexes,
		PoolKeys:           publicInputs.PoolKeys,
		ShareCommitments:   shareCommitments,
		Shares:             trees,
	}, nil
}

// commitmentSets turns coefficient scalars into the commitment points the
// finalization circuit consumes: A_{c,j,m} = a_{c,j,m}·G.
func commitmentSets(contributions [][][]*big.Int) ([][][]types.CurvePoint, error) {
	sets := make([][][]types.CurvePoint, len(contributions))
	for c, keys := range contributions {
		sets[c] = make([][]types.CurvePoint, len(keys))
		for j, coefficients := range keys {
			sets[c][j] = make([]types.CurvePoint, len(coefficients))
			for m, coefficient := range coefficients {
				if coefficient == nil {
					return nil, fmt.Errorf("contributor %d key %d coefficient %d is nil", c, j, m)
				}
				sets[c][j][m] = ScalarBasePoint(coefficient)
			}
		}
	}
	return sets, nil
}

// BuildPartialDecryptionSubmission builds a partial decryption over
// C1 = base·G. `aid` and `ciphertextIndex` are bound into the Fiat-Shamir
// transcript so the on-chain checks (publicInputs[1]==aid,
// publicInputs[2]==ctIdx) succeed. `share` is the member's share of the
// application's pool key and `tree` the key's share-commitment tree, from
// which the returned ShareProof is derived.
//
// C2 is left as the identity: it is not a proof input, it only travels to
// `submitPartialDecryption` as calldata, and callers that need the real one
// use BuildPartialDecryptionSubmissionFromBase.
func BuildPartialDecryptionSubmission(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	participantIndex uint16,
	base *big.Int,
	share *big.Int,
	nonce *big.Int,
	tree ShareTree,
) (*PartialDecryptionSubmission, error) {
	basePoint := group.Generator()
	basePoint.ScalarBaseMult(base)
	identityC2 := types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	return BuildPartialDecryptionSubmissionFromBase(
		ctx, epochID, aid, ciphertextIndex, participantIndex,
		group.Encode(basePoint), identityC2, share, nonce, tree,
	)
}

// BuildPartialDecryptionSubmissionFromBase is the variant used when the caller
// already has the C1 ciphertext point (e.g. recovered from a
// CiphertextSubmitted event log) instead of the scalar k that produced it.
// The SDK e2e path goes through this entry point because the SDK encrypts
// with a random k that the test fixture never sees.
//
// `c2` is just stashed on the returned struct so the caller can pass it
// through to submitPartialDecryption.
func BuildPartialDecryptionSubmissionFromBase(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	participantIndex uint16,
	base types.CurvePoint,
	c2 types.CurvePoint,
	share *big.Int,
	nonce *big.Int,
	tree ShareTree,
) (*PartialDecryptionSubmission, error) {
	roundHash := RoundScalar(epochID)
	assignment := partialdecrypt.Assignment{
		RoundHash:        roundHash,
		Aid:              new(big.Int).SetBytes(aid[:]),
		CtIdx:            new(big.Int).SetUint64(uint64(ciphertextIndex)),
		ParticipantIndex: participantIndex,
		Base:             base,
		Secret:           share,
		Nonce:            nonce,
	}
	witness, publicInputs, err := partialdecrypt.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}
	shareProof, err := tree.Proof(participantIndex)
	if err != nil {
		return nil, err
	}

	runtime, err := loadPartialDecryptRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove partial decrypt: %w", err)
	}

	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicAssignment(publicInputs.PublicWitness())
	if err != nil {
		return nil, err
	}

	return &PartialDecryptionSubmission{
		Proof: proofBytes,
		Input: inputBytes,
		DeltaHash: ethcrypto.Keccak256Hash(
			common.LeftPadBytes(publicInputs.Delta.X.Bytes(), 32),
			common.LeftPadBytes(publicInputs.Delta.Y.Bytes(), 32),
		),
		Delta:      publicInputs.Delta,
		RoundHash:  new(big.Int).Set(roundHash),
		ShareProof: shareProof,
		C1:         base,
		C2:         c2,
	}, nil
}

// BuildDecryptCombineOutput builds a combine proof for a synthetic ciphertext
// with C1 = base·G. It derives the organizer term Δ = sk_org·C1 itself and
// sets C2 = m·G + Σ λ_k δ_k + Δ so the circuit's plaintext equation holds.
//
// `organizerSecret` is the revealed sk_org of an organizer-locked
// application, or 0 (paired with the identity PK) for an automatic one.
func BuildDecryptCombineOutput(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	threshold uint16,
	base *big.Int,
	organizerSecret *big.Int,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	c1Point := group.Generator()
	c1Point.ScalarBaseMult(base)
	c1 := group.Encode(c1Point)

	deltaOrg, err := OrganizerTerm(organizerSecret, c1)
	if err != nil {
		return nil, err
	}
	c2, err := combineCiphertextC2(participantIndexes, partialDecryptions, deltaOrg, plaintext)
	if err != nil {
		return nil, err
	}

	return BuildDecryptCombineOutputFromCiphertext(ctx, epochID, aid, ciphertextIndex, threshold,
		c1, c2, OrganizerKey(organizerSecret), organizerSecret,
		participantIndexes, partialDecryptions, plaintext)
}

// OrganizerKey is PK_org = sk_org·G, or the identity (0, 1) when there is no
// organizer half — the key an automatic registration stores.
func OrganizerKey(organizerSecret *big.Int) types.CurvePoint {
	if organizerSecret == nil || organizerSecret.Sign() == 0 {
		return IdentityPoint()
	}
	return ScalarBasePoint(organizerSecret)
}

// OrganizerTerm is Δ = sk_org·C1, the point the combine adds back on top of
// the interpolated partials. It is the identity for an automatic application.
func OrganizerTerm(organizerSecret *big.Int, c1 types.CurvePoint) (types.CurvePoint, error) {
	if organizerSecret == nil || organizerSecret.Sign() == 0 {
		return IdentityPoint(), nil
	}
	base, err := group.Decode(c1)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode C1: %w", err)
	}
	delta := group.NewPoint()
	delta.ScalarMult(base, organizerSecret)
	return group.Encode(delta), nil
}

// IdentityPoint is the twisted-Edwards identity (0, 1).
func IdentityPoint() types.CurvePoint {
	return types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
}

// combineCiphertextC2 reconstructs C2 = m·G + Σ λ_k δ_k + Δ_org, the value a
// real encryption under PK_aid = P_j + PK_org would have produced.
func combineCiphertextC2(
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	deltaOrg types.CurvePoint,
	plaintext *big.Int,
) (types.CurvePoint, error) {
	messagePoint := group.Generator()
	messagePoint.ScalarBaseMult(plaintext)

	indexes := ccommon.Uint16sToBigInts(participantIndexes)
	combinedPoint, err := ccommon.InterpolatePointsAtZeroNative(indexes, partialDecryptions)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("interpolate combined partial decryptions: %w", err)
	}
	combinedNative, err := group.Decode(combinedPoint)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode combined point: %w", err)
	}
	deltaNative, err := group.Decode(deltaOrg)
	if err != nil {
		return types.CurvePoint{}, fmt.Errorf("decode organizer term: %w", err)
	}
	c2Point := group.NewPoint()
	c2Point.Set(messagePoint)
	c2Point.Add(c2Point, combinedNative)
	c2Point.Add(c2Point, deltaNative)
	return group.Encode(c2Point), nil
}

// BuildDecryptCombineOutputFromCiphertext is the variant used when the caller
// already has C1 and C2. The circuit proves knowledge of sk_org with
// PK_org = sk_org·G and Δ = sk_org·C1; the contract only checks that the
// transcript's PK_org words equal the application's registered key.
func BuildDecryptCombineOutputFromCiphertext(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	threshold uint16,
	ciphertextC1 types.CurvePoint,
	ciphertextC2 types.CurvePoint,
	organizerPK types.CurvePoint,
	organizerSecret *big.Int,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	secret := big.NewInt(0)
	if organizerSecret != nil {
		secret = new(big.Int).Set(organizerSecret)
	}
	assignment := decryptcombine.Assignment{
		RoundHash:          RoundScalar(epochID),
		Aid:                new(big.Int).SetBytes(aid[:]),
		CtIdx:              new(big.Int).SetUint64(uint64(ciphertextIndex)),
		OrganizerPK:        organizerPK,
		OrganizerSecret:    secret,
		Threshold:          threshold,
		CiphertextC1:       ciphertextC1,
		CiphertextC2:       ciphertextC2,
		ParticipantIndexes: participantIndexes,
		PartialDecryptions: partialDecryptions,
		Plaintext:          plaintext,
	}
	witness, publicInputs, err := decryptcombine.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}

	runtime, err := loadDecryptCombineRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove decrypt combine: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicAssignment(publicInputs.PublicWitness())
	if err != nil {
		return nil, err
	}
	transcriptBytes, err := encodeSolidityWords(publicInputs.TranscriptScalars()...)
	if err != nil {
		return nil, err
	}

	return &DecryptCombineOutput{
		Proof:        proofBytes,
		Input:        inputBytes,
		Transcript:   transcriptBytes,
		CombineHash:  common.BigToHash(publicInputs.CombineHash),
		Plaintext:    new(big.Int).Set(plaintext),
		CiphertextC1: ciphertextC1,
		CiphertextC2: ciphertextC2,
		OrganizerPK:  organizerPK,
	}, nil
}

func RoundScalar(epochID [12]byte) *big.Int {
	return new(big.Int).SetBytes(epochID[:])
}

func loadContributionRuntime(ctx context.Context) (*circuits.CircuitRuntime, error) {
	if err := ensureArtifactsBaseDir(); err != nil {
		return nil, err
	}
	contributionRuntimeOnce.Do(func() {
		contributionRuntime, contributionRuntimeErr = contribution.Artifacts.LoadOrSetupForCircuit(
			ctx,
			&contribution.ContributionCircuit{},
		)
	})
	return contributionRuntime, contributionRuntimeErr
}

func loadFinalizeRuntime(ctx context.Context) (*circuits.CircuitRuntime, error) {
	if err := ensureArtifactsBaseDir(); err != nil {
		return nil, err
	}
	finalizeRuntimeOnce.Do(func() {
		finalizeRuntime, finalizeRuntimeErr = finalize.Artifacts.LoadOrSetupForCircuit(ctx, &finalize.FinalizeCircuit{})
	})
	return finalizeRuntime, finalizeRuntimeErr
}

func loadPartialDecryptRuntime(ctx context.Context) (*circuits.CircuitRuntime, error) {
	if err := ensureArtifactsBaseDir(); err != nil {
		return nil, err
	}
	partialDecryptRuntimeOnce.Do(func() {
		partialDecryptRuntime, partialDecryptRuntimeErr = partialdecrypt.Artifacts.LoadOrSetupForCircuit(
			ctx,
			&partialdecrypt.PartialDecryptCircuit{},
		)
	})
	return partialDecryptRuntime, partialDecryptRuntimeErr
}

func loadDecryptCombineRuntime(ctx context.Context) (*circuits.CircuitRuntime, error) {
	if err := ensureArtifactsBaseDir(); err != nil {
		return nil, err
	}
	decryptCombineRuntimeOnce.Do(func() {
		decryptCombineRuntime, decryptCombineRuntimeErr = decryptcombine.Artifacts.LoadOrSetupForCircuit(
			ctx,
			&decryptcombine.DecryptCombineCircuit{},
		)
	})
	return decryptCombineRuntime, decryptCombineRuntimeErr
}

func marshalSolidityProof(proof groth16backend.Proof) ([]byte, error) {
	bn254Proof, ok := proof.(*groth16bn254.Proof)
	if !ok {
		return nil, fmt.Errorf("unexpected proof type %T", proof)
	}
	return bn254Proof.MarshalSolidity(), nil
}

func encodeSolidityWords(values ...*big.Int) ([]byte, error) {
	encoded := make([]byte, 0, len(values)*32)
	for i, value := range values {
		if value == nil {
			return nil, fmt.Errorf("value %d is nil", i)
		}
		if value.Sign() < 0 {
			return nil, fmt.Errorf("value %d is negative", i)
		}
		word := common.LeftPadBytes(value.Bytes(), 32)
		encoded = append(encoded, word...)
	}
	return encoded, nil
}

func encodePublicAssignment(publicAssignment frontend.Circuit) ([]byte, error) {
	w, err := frontend.NewWitness(publicAssignment, gnec.BN254.ScalarField(), frontend.PublicOnly())
	if err != nil {
		return nil, fmt.Errorf("build public witness: %w", err)
	}
	rawValues, err := witnessVectorBigInts(w.Vector())
	if err != nil {
		return nil, fmt.Errorf("extract public witness vector: %w", err)
	}
	return encodeSolidityWords(rawValues...)
}

func witnessVectorBigInts(vector any) ([]*big.Int, error) {
	rv := reflect.ValueOf(vector)
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("unexpected witness vector type %T", vector)
	}
	values := make([]*big.Int, rv.Len())
	for i := range rv.Len() {
		method := rv.Index(i).Addr().MethodByName("BigInt")
		if !method.IsValid() {
			return nil, fmt.Errorf("witness element %d does not expose BigInt", i)
		}
		out := method.Call([]reflect.Value{reflect.ValueOf(new(big.Int))})
		if len(out) != 1 {
			return nil, fmt.Errorf("unexpected BigInt result arity for witness element %d", i)
		}
		value, ok := out[0].Interface().(*big.Int)
		if !ok {
			return nil, fmt.Errorf("unexpected BigInt result type for witness element %d: %T", i, out[0].Interface())
		}
		values[i] = new(big.Int).Set(value)
	}
	return values, nil
}

// ensureArtifactsBaseDir is a no-op retained so the existing call sites keep
// compiling without edits. The circuits package's init() already points
// BaseDir at $DAVINCI_DKG_ARTIFACTS_DIR (or ~/.davinci/artifacts), which is
// the same location maintained by go-test-circuits.yml's actions/cache step
// and by `make circuits`. Overriding to an in-repo `artifacts/` directory
// (the previous behavior) is unsafe because that directory is .gitignore'd
// and almost always stale; the override caused integration tests to fall
// back to a local trusted setup whose vkey did not match the on-chain
// verifier, producing ProofInvalid() reverts.
func ensureArtifactsBaseDir() error {
	return nil
}
