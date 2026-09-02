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
	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/crypto/group"
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

type FinalizeEpochOutput struct {
	Proof                    []byte
	Input                    []byte
	Transcript               []byte
	AggregateCommitmentsHash [32]byte
	CollectivePublicKeyHash  [32]byte
	ShareCommitmentHash      [32]byte
	RoundHash                *big.Int
	ShareCommitments         []types.CurvePoint
}

type PartialDecryptionSubmission struct {
	Proof     []byte
	Input     []byte
	DeltaHash [32]byte
	Delta     types.CurvePoint
	RoundHash *big.Int
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
	// DeltaOrg and OrganizerProof are the exact organizer-share words the
	// transcript commits to. `submitOrganizerShare` must post these same
	// words before `combineDecryption`, or the contract's share-hash check
	// fails.
	DeltaOrg       types.CurvePoint
	OrganizerProof dleq.Proof
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

func BuildContributionSubmission(
	ctx context.Context,
	services *TestServices,
	epochID [12]byte,
	threshold uint16,
	committeeSize uint16,
	contributorIndex uint16,
	coefficients []*big.Int,
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

func BuildFinalizeEpochOutput(
	ctx context.Context,
	epochID [12]byte,
	threshold uint16,
	committeeSize uint16,
	participantIndexes []uint16,
	contributionCoefficients [][]*big.Int,
) (*FinalizeEpochOutput, error) {
	roundHash := RoundScalar(epochID)
	assignment := finalize.Assignment{
		RoundHash:                roundHash,
		Threshold:                threshold,
		CommitteeSize:            committeeSize,
		ParticipantIndexes:       participantIndexes,
		ContributionCoefficients: contributionCoefficients,
	}
	witness, publicInputs, err := finalize.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}

	runtime, err := loadFinalizeRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove finalize: %w", err)
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

	return &FinalizeEpochOutput{
		Proof:                    proofBytes,
		Input:                    inputBytes,
		Transcript:               transcriptBytes,
		AggregateCommitmentsHash: common.BigToHash(publicInputs.AggregateHash),
		CollectivePublicKeyHash:  common.BigToHash(publicInputs.CollectivePublicKey),
		ShareCommitmentHash:      common.BigToHash(publicInputs.ShareCommitmentHash),
		RoundHash:                new(big.Int).Set(roundHash),
		ShareCommitments:         publicInputs.ShareCommitments,
	}, nil
}

// BuildPartialDecryptionSubmission builds a partial decryption over
// C1 = base·G. `aid` and `ciphertextIndex` are bound into the Fiat-Shamir
// transcript so the on-chain checks (publicInputs[1]==aid,
// publicInputs[2]==ctIdx) succeed.
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
	secret *big.Int,
	nonce *big.Int,
) (*PartialDecryptionSubmission, error) {
	basePoint := group.Generator()
	basePoint.ScalarBaseMult(base)
	identityC2 := types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	return BuildPartialDecryptionSubmissionFromBase(
		ctx, epochID, aid, ciphertextIndex, participantIndex,
		group.Encode(basePoint), identityC2, secret, nonce,
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
	secret *big.Int,
	nonce *big.Int,
) (*PartialDecryptionSubmission, error) {
	roundHash := RoundScalar(epochID)
	assignment := partialdecrypt.Assignment{
		RoundHash:        roundHash,
		Aid:              new(big.Int).SetBytes(aid[:]),
		CtIdx:            new(big.Int).SetUint64(uint64(ciphertextIndex)),
		ParticipantIndex: participantIndex,
		Base:             base,
		Secret:           secret,
		Nonce:            nonce,
	}
	witness, publicInputs, err := partialdecrypt.BuildWitness(assignment)
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
		Delta:     publicInputs.Delta,
		RoundHash: new(big.Int).Set(roundHash),
		C1:        base,
		C2:        c2,
	}, nil
}

// BuildDecryptCombineOutput builds a combine proof for a synthetic ciphertext
// with C1 = base·G. It derives the organizer share Δ = skOrg·C1 itself and
// sets C2 = m·G + Σ λ_k δ_k + Δ so the circuit's plaintext equation holds.
//
// The caller must post the returned DeltaOrg/OrganizerProof via
// submitOrganizerShare before calling combineDecryption.
func BuildDecryptCombineOutput(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	threshold uint16,
	base *big.Int,
	skOrg *big.Int,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	c1Point := group.Generator()
	c1Point.ScalarBaseMult(base)
	c1 := group.Encode(c1Point)

	deltaOrg, organizerProof, err := dleq.ProveOrganizerShare(epochID, aid, ciphertextIndex, skOrg, c1)
	if err != nil {
		return nil, fmt.Errorf("prove organizer share: %w", err)
	}

	c2, err := combineCiphertextC2(participantIndexes, partialDecryptions, deltaOrg, plaintext)
	if err != nil {
		return nil, err
	}

	return BuildDecryptCombineOutputFromCiphertext(ctx, epochID, aid, ciphertextIndex, threshold,
		c1, c2, ScalarBasePoint(skOrg), deltaOrg, organizerProof,
		participantIndexes, partialDecryptions, plaintext)
}

// combineCiphertextC2 reconstructs C2 = m·G + Σ λ_k δ_k + Δ_org, the value a
// real encryption under PK_aid = PK_ep + PK_org would have produced.
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
		return types.CurvePoint{}, fmt.Errorf("decode organizer share: %w", err)
	}
	c2Point := group.NewPoint()
	c2Point.Set(messagePoint)
	c2Point.Add(c2Point, combinedNative)
	c2Point.Add(c2Point, deltaNative)
	return group.Encode(c2Point), nil
}

// BuildDecryptCombineOutputFromCiphertext is the variant used when the caller
// already has C1, C2 and the organizer share that was (or will be) posted on
// chain. The witness bindings here MUST match the contract's
// submitOrganizerShare / combineDecryption checks: the circuit recomputes the
// DLEQ challenge `e` from exactly these words and the contract pins it.
func BuildDecryptCombineOutputFromCiphertext(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	threshold uint16,
	ciphertextC1 types.CurvePoint,
	ciphertextC2 types.CurvePoint,
	organizerPK types.CurvePoint,
	deltaOrg types.CurvePoint,
	organizerProof dleq.Proof,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	assignment := decryptcombine.Assignment{
		RoundHash:          RoundScalar(epochID),
		Aid:                new(big.Int).SetBytes(aid[:]),
		CtIdx:              new(big.Int).SetUint64(uint64(ciphertextIndex)),
		DeltaOrg:           deltaOrg,
		OrganizerPK:        organizerPK,
		OrganizerProof:     organizerProof,
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
		Proof:          proofBytes,
		Input:          inputBytes,
		Transcript:     transcriptBytes,
		CombineHash:    common.BigToHash(publicInputs.CombineHash),
		Plaintext:      new(big.Int).Set(plaintext),
		CiphertextC1:   ciphertextC1,
		CiphertextC2:   ciphertextC2,
		DeltaOrg:       deltaOrg,
		OrganizerProof: organizerProof,
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
	for i := 0; i < rv.Len(); i++ {
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
