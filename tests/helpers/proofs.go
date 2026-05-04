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

// BuildPartialDecryptionSubmission for the legacy per-epoch path —
// defaults aid/role to (0, COMMITTEE) and lets the caller pick the
// ciphertextIndex so multi-ciphertext tests can produce proofs that
// match the on-chain (publicInputs[2]==ctIdx) check.
func BuildPartialDecryptionSubmission(
	ctx context.Context,
	epochID [12]byte,
	ciphertextIndex uint16,
	participantIndex uint16,
	base *big.Int,
	secret *big.Int,
	nonce *big.Int,
) (*PartialDecryptionSubmission, error) {
	basePoint := group.Generator()
	basePoint.ScalarBaseMult(base)
	// Legacy callers don't have C2; pass identity. The on-chain C1 binding
	// will still match because the test fixture uses the canonical TEST_CT
	// vectors with C1 = generator. C2 = identity makes the keccak match
	// unrealistic in real flows but works for the few unit-test paths
	// that still go through this entry point.
	identityC2 := types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	return BuildPartialDecryptionSubmissionFromBase(ctx, epochID, [32]byte{}, ciphertextIndex, participantIndex, group.Encode(basePoint), identityC2, secret, nonce)
}

// BuildPartialDecryptionSubmissionFromBase is the variant used when the caller
// already has the c1 ciphertext point (e.g. recovered from a CiphertextSubmitted
// event log) instead of the scalar k that produced it. The flow.test SDK e2e
// path goes through this entry point because the SDK encrypts with a random k
// that the test fixture never sees.
//
// `aid` and `ciphertextIndex` are bound into the Fiat-Shamir transcript
// via the witness builder so the on-chain submitPartialDecryption check
// (publicInputs[1]==aid, publicInputs[2]==ctIdx, publicInputs[3]==COMMITTEE)
// succeeds. Pass `[32]byte{}` aid for the legacy per-epoch path.
//
// `c2` is just stashed on the returned struct so the caller can pass it
// through to SubmitPartialDecryption. Callers that
// don't have c2 (legacy single-CT-test paths) can pass the identity
// point and use the FromBase variant whose API knows the full ct.
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
		Role:             big.NewInt(int64(protocol.RoleCommittee)),
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

func BuildDecryptCombineOutput(
	ctx context.Context,
	epochID [12]byte,
	ciphertextIndex uint16,
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

	return BuildDecryptCombineOutputFromCiphertext(ctx, epochID, ciphertextIndex, threshold,
		group.Encode(c1Point), group.Encode(c2Point),
		participantIndexes, partialDecryptions, plaintext)
}

// BuildDecryptCombineOutputFromCiphertext is the variant used when the caller
// already has c1, c2 as curve points (e.g. recovered from a SDK-submitted
// CiphertextSubmitted event log) and the plaintext was discovered out-of-band
// via brute-force discrete log on m·G = c2 - sum(λᵢ·Δᵢ).
// BuildDecryptCombineOutputFromCiphertext builds the combine proof for the
// legacy per-epoch path (aid=0, mode=0=PUBLIC_DERIVATION, S=0, no organizer
// share). The on-chain `combineDecryption` requires publicInputs[1..3] to
// equal aid / ctIdx / mode, so we set them explicitly even when zero.
//
// `ciphertextIndex` is bound into both the contract-side check
// (publicInputs[2]==ciphertextIndex) and the in-circuit transcript.
func BuildDecryptCombineOutputFromCiphertext(
	ctx context.Context,
	epochID [12]byte,
	ciphertextIndex uint16,
	threshold uint16,
	ciphertextC1 types.CurvePoint,
	ciphertextC2 types.CurvePoint,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	return BuildDecryptCombineOutputForApp(
		ctx, epochID, [32]byte{}, ciphertextIndex, 0 /* mode */, big.NewInt(0), /* S */
		identityPoint(), threshold, ciphertextC1, ciphertextC2,
		participantIndexes, partialDecryptions, plaintext,
	)
}

// identityPoint returns the BabyJubJub identity (0, 1) used as the default
// DeltaOrg for the legacy mode-0 combine path.
func identityPoint() types.CurvePoint {
	return types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
}

// BuildDecryptCombineOutputForApp is the per-application variant. Mode 0
// uses S+identity-DeltaOrg; mode 1 uses S=0+real DeltaOrg from the
// organizer's submitted share. The witness bindings here MUST match the
// contract's submitOrganizerShare / combineDecryption checks.
func BuildDecryptCombineOutputForApp(
	ctx context.Context,
	epochID [12]byte,
	aid [32]byte,
	ciphertextIndex uint16,
	mode uint8,
	s *big.Int,
	deltaOrg types.CurvePoint,
	threshold uint16,
	ciphertextC1 types.CurvePoint,
	ciphertextC2 types.CurvePoint,
	participantIndexes []uint16,
	partialDecryptions []types.CurvePoint,
	plaintext *big.Int,
) (*DecryptCombineOutput, error) {
	assignment := decryptcombine.Assignment{
		RoundHash:          RoundScalar(epochID),
		Aid:                new(big.Int).SetBytes(aid[:]),
		CtIdx:              new(big.Int).SetUint64(uint64(ciphertextIndex)),
		Mode:               new(big.Int).SetUint64(uint64(mode)),
		S:                  s,
		DeltaOrg:           deltaOrg,
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
