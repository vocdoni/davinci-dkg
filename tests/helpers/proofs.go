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
	"github.com/vocdoni/davinci-dkg/circuits/revealshare"
	"github.com/vocdoni/davinci-dkg/circuits/revealsubmit"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

type ContributionSubmission struct {
	Proof               []byte
	Input               []byte
	Transcript          []byte
	CommitmentsHash     [32]byte
	EncryptedSharesHash [32]byte
	Commitment0X        *big.Int
	Commitment0Y        *big.Int
	RoundHash           *big.Int
}

type FinalizeRoundOutput struct {
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
}

type DecryptCombineOutput struct {
	Proof        []byte
	Input        []byte
	Transcript   []byte
	CombineHash  [32]byte
	Plaintext    *big.Int
	CiphertextC1 types.CurvePoint
	CiphertextC2 types.CurvePoint
}

type RevealShareOutput struct {
	Proof                   []byte
	Input                   []byte
	Transcript              []byte
	ShareHash               [32]byte
	DisclosureHash          [32]byte
	ReconstructedSecretHash [32]byte
	ReconstructedSecret     *big.Int
}

type RevealShareSubmission struct {
	Proof      []byte
	Input      []byte
	ShareHash  [32]byte
	ShareValue *big.Int
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
	revealShareRuntimeOnce    sync.Once
	revealShareRuntime        *circuits.CircuitRuntime
	revealShareRuntimeErr     error
	revealSubmitRuntimeOnce   sync.Once
	revealSubmitRuntime       *circuits.CircuitRuntime
	revealSubmitRuntimeErr    error
)

func BuildContributionSubmission(
	ctx context.Context,
	services *TestServices,
	roundID [12]byte,
	threshold uint16,
	committeeSize uint16,
	contributorIndex uint16,
	coefficients []*big.Int,
	recipientIndexes []uint16,
) (*ContributionSubmission, error) {
	roundHash := RoundScalar(roundID)
	recipientKeys, encryptionNonces, err := contributionRecipients(ctx, services, roundID, recipientIndexes)
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
	transcriptBytes, err := encodeSolidityWords(publicInputs.TranscriptScalars()...)
	if err != nil {
		return nil, err
	}

	return &ContributionSubmission{
		Proof:               proofBytes,
		Input:               inputBytes,
		Transcript:          transcriptBytes,
		CommitmentsHash:     common.BigToHash(publicInputs.CommitmentHash),
		EncryptedSharesHash: common.BigToHash(publicInputs.ShareHash),
		Commitment0X:        new(big.Int).Set(publicInputs.CommitmentX0),
		Commitment0Y:        new(big.Int).Set(publicInputs.CommitmentY0),
		RoundHash:           new(big.Int).Set(roundHash),
	}, nil
}

func contributionRecipients(
	ctx context.Context,
	services *TestServices,
	roundID [12]byte,
	recipientIndexes []uint16,
) ([]types.NodeKey, []*big.Int, error) {
	participants, err := services.Contracts.SelectedParticipants(ctx, roundID)
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

func BuildFinalizeRoundOutput(
	ctx context.Context,
	roundID [12]byte,
	threshold uint16,
	committeeSize uint16,
	participantIndexes []uint16,
	contributionCoefficients [][]*big.Int,
) (*FinalizeRoundOutput, error) {
	roundHash := RoundScalar(roundID)
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

	return &FinalizeRoundOutput{
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

func BuildRevealShareSubmission(
	ctx context.Context,
	roundID [12]byte,
	participantIndex uint16,
	shareValue *big.Int,
	shareCommitment types.CurvePoint,
) (*RevealShareSubmission, error) {
	assignment := revealsubmit.Assignment{
		RoundHash:        RoundScalar(roundID),
		ParticipantIndex: participantIndex,
		ShareValue:       shareValue,
		ShareCommitment:  shareCommitment,
	}
	witness, publicInputs, err := revealsubmit.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}

	runtime, err := loadRevealSubmitRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove reveal submit: %w", err)
	}
	proofBytes, err := marshalSolidityProof(proof)
	if err != nil {
		return nil, err
	}
	inputBytes, err := encodePublicAssignment(publicInputs.PublicWitness())
	if err != nil {
		return nil, err
	}
	return &RevealShareSubmission{
		Proof:      proofBytes,
		Input:      inputBytes,
		ShareHash:  common.BigToHash(shareValue),
		ShareValue: new(big.Int).Set(shareValue),
	}, nil
}

func BuildPartialDecryptionSubmission(
	ctx context.Context,
	roundID [12]byte,
	participantIndex uint16,
	base *big.Int,
	secret *big.Int,
	nonce *big.Int,
) (*PartialDecryptionSubmission, error) {
	basePoint := group.Generator()
	basePoint.ScalarBaseMult(base)
	return BuildPartialDecryptionSubmissionFromBase(ctx, roundID, participantIndex, group.Encode(basePoint), secret, nonce)
}

// BuildPartialDecryptionSubmissionFromBase is the variant used when the caller
// already has the c1 ciphertext point (e.g. recovered from a CiphertextSubmitted
// event log) instead of the scalar k that produced it. The flow.test SDK e2e
// path goes through this entry point because the SDK encrypts with a random k
// that the test fixture never sees.
func BuildPartialDecryptionSubmissionFromBase(
	ctx context.Context,
	roundID [12]byte,
	participantIndex uint16,
	base types.CurvePoint,
	secret *big.Int,
	nonce *big.Int,
) (*PartialDecryptionSubmission, error) {
	roundHash := RoundScalar(roundID)
	assignment := partialdecrypt.Assignment{
		RoundHash:        roundHash,
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
	}, nil
}

func BuildDecryptCombineOutput(
	ctx context.Context,
	roundID [12]byte,
	threshold uint16,
	base *big.Int,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	c1Point := group.Generator()
	c1Point.ScalarBaseMult(base)
	messagePoint := group.Generator()
	messagePoint.ScalarBaseMult(plaintext)

	indexes := ccommon.Uint16sToBigInts(participantIndexes)
	combinedPoint, err := ccommon.InterpolatePointsAtZeroNative(indexes, partialDecryptions)
	if err != nil {
		return nil, fmt.Errorf("interpolate combined partial decryptions: %w", err)
	}
	combinedNative, err := group.Decode(combinedPoint)
	if err != nil {
		return nil, fmt.Errorf("decode combined point: %w", err)
	}
	c2Point := group.NewPoint()
	c2Point.Set(messagePoint)
	c2Point.Add(c2Point, combinedNative)

	return BuildDecryptCombineOutputFromCiphertext(ctx, roundID, threshold,
		group.Encode(c1Point), group.Encode(c2Point),
		participantIndexes, partialDecryptions, plaintext)
}

// BuildDecryptCombineOutputFromCiphertext is the variant used when the caller
// already has c1, c2 as curve points (e.g. recovered from a SDK-submitted
// CiphertextSubmitted event log) and the plaintext was discovered out-of-band
// via brute-force discrete log on m·G = c2 - sum(λᵢ·Δᵢ).
func BuildDecryptCombineOutputFromCiphertext(
	ctx context.Context,
	roundID [12]byte,
	threshold uint16,
	ciphertextC1 types.CurvePoint,
	ciphertextC2 types.CurvePoint,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	assignment := decryptcombine.Assignment{
		RoundHash:          RoundScalar(roundID),
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
	}, nil
}

func BuildRevealShareOutput(
	ctx context.Context,
	roundID [12]byte,
	threshold uint16,
	participantIndexes []uint16,
	revealedShares []*big.Int,
) (*RevealShareOutput, error) {
	assignment := revealshare.Assignment{
		RoundHash:          RoundScalar(roundID),
		Threshold:          threshold,
		ParticipantIndexes: participantIndexes,
		RevealedShares:     revealedShares,
	}
	witness, publicInputs, err := revealshare.BuildWitness(assignment)
	if err != nil {
		return nil, err
	}

	runtime, err := loadRevealShareRuntime(ctx)
	if err != nil {
		return nil, err
	}
	proof, err := runtime.ProveAndVerify(witness)
	if err != nil {
		return nil, fmt.Errorf("prove reveal share: %w", err)
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

	return &RevealShareOutput{
		Proof:                   proofBytes,
		Input:                   inputBytes,
		Transcript:              transcriptBytes,
		ShareHash:               common.BigToHash(revealedShares[0]),
		DisclosureHash:          common.BigToHash(publicInputs.DisclosureHash),
		ReconstructedSecretHash: common.BigToHash(publicInputs.ReconstructedSecretHash),
		ReconstructedSecret:     new(big.Int).Set(publicInputs.ReconstructedSecretHash),
	}, nil
}

func RoundScalar(roundID [12]byte) *big.Int {
	return new(big.Int).SetBytes(roundID[:])
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

func loadRevealShareRuntime(ctx context.Context) (*circuits.CircuitRuntime, error) {
	if err := ensureArtifactsBaseDir(); err != nil {
		return nil, err
	}
	revealShareRuntimeOnce.Do(func() {
		revealShareRuntime, revealShareRuntimeErr = revealshare.Artifacts.LoadOrSetupForCircuit(
			ctx,
			&revealshare.RevealShareCircuit{},
		)
	})
	return revealShareRuntime, revealShareRuntimeErr
}

func loadRevealSubmitRuntime(ctx context.Context) (*circuits.CircuitRuntime, error) {
	if err := ensureArtifactsBaseDir(); err != nil {
		return nil, err
	}
	revealSubmitRuntimeOnce.Do(func() {
		revealSubmitRuntime, revealSubmitRuntimeErr = revealsubmit.Artifacts.LoadOrSetupForCircuit(
			ctx,
			&revealsubmit.RevealSubmitCircuit{},
		)
	})
	return revealSubmitRuntime, revealSubmitRuntimeErr
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
