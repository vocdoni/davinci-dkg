package decryptcombine

import (
	"fmt"
	"math/big"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/types"
)

var decryptCombineTranscriptDomain = ethcrypto.Keccak256Hash([]byte("davinci-dkg:decrypt-combine:v1"))

// PublicInputs is the native representation of the decrypt-combine public inputs.
type PublicInputs struct {
	RoundHash            *big.Int // semantically: eid
	Aid                  *big.Int
	CtIdx                *big.Int
	Mode                 *big.Int
	S                    *big.Int
	DeltaOrg             types.CurvePoint
	Threshold            *big.Int
	ShareCount           *big.Int
	CombineHash          *big.Int
	PlaintextHash        *big.Int
	Challenge            *big.Int
	TranscriptCommitment *big.Int
	CiphertextC1         types.CurvePoint
	CiphertextC2         types.CurvePoint
	ParticipantIndexes   []*big.Int
	PartialDecryptions   []types.CurvePoint
}

// BuildWitness materializes the decrypt-combine native assignment.
func BuildWitness(a Assignment) (*DecryptCombineCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	threshold := big.NewInt(int64(a.Threshold))
	shareCount := big.NewInt(int64(len(a.ParticipantIndexes)))
	// Default per-application correction fields to zero / identity / mode 0
	// when callers don't supply them — preserves the legacy combine path
	// (T = 0·C_1 = O, identity addend in the C_2 equation).
	aid := new(big.Int)
	if a.Aid != nil {
		aid.Set(a.Aid)
	}
	ctIdx := new(big.Int)
	if a.CtIdx != nil {
		ctIdx.Set(a.CtIdx)
	}
	mode := new(big.Int)
	if a.Mode != nil {
		mode.Set(a.Mode)
	}
	s := new(big.Int)
	if a.S != nil {
		s.Set(a.S)
	}
	deltaOrg := types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	if a.DeltaOrg.X != nil && a.DeltaOrg.Y != nil {
		deltaOrg = a.DeltaOrg
	}
	participantIndexes := ccommon.Uint16sToBigInts(a.ParticipantIndexes)
	participantIndexes, err := ccommon.PadBigInts(participantIndexes, MaxShares)
	if err != nil {
		return nil, nil, err
	}
	partials, err := ccommon.CircuitPoints(a.PartialDecryptions, MaxShares)
	if err != nil {
		return nil, nil, err
	}
	paddedPartialDecryptions := make([]types.CurvePoint, MaxShares)
	for i := range MaxShares {
		if i < len(a.PartialDecryptions) {
			paddedPartialDecryptions[i] = a.PartialDecryptions[i]
		} else {
			paddedPartialDecryptions[i] = types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
		}
	}

	hashInputs := []*big.Int{
		a.RoundHash, // eid
		aid,
		ctIdx,
		mode,
		s,
		deltaOrg.X,
		deltaOrg.Y,
		threshold,
		shareCount,
		a.CiphertextC1.X,
		a.CiphertextC1.Y,
		a.CiphertextC2.X,
		a.CiphertextC2.Y,
	}
	for i := range len(a.ParticipantIndexes) {
		hashInputs = append(
			hashInputs,
			big.NewInt(int64(a.ParticipantIndexes[i])),
			a.PartialDecryptions[i].X,
			a.PartialDecryptions[i].Y,
		)
	}
	for i := len(a.ParticipantIndexes); i < MaxShares; i++ {
		hashInputs = append(hashInputs, big.NewInt(0), big.NewInt(0), big.NewInt(1))
	}
	combineHash, err := ccommon.MultiHashNative(hashInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("hash partial decryptions: %w", err)
	}
	plaintextHash := new(big.Int).Set(a.Plaintext)
	transcriptValues := make([]*big.Int, 0, 4+3*MaxShares)
	transcriptValues = append(
		transcriptValues,
		a.CiphertextC1.X,
		a.CiphertextC1.Y,
		a.CiphertextC2.X,
		a.CiphertextC2.Y,
	)
	transcriptValues = append(transcriptValues, participantIndexes...)
	for i := range MaxShares {
		transcriptValues = append(
			transcriptValues,
			paddedPartialDecryptions[i].X,
			paddedPartialDecryptions[i].Y,
		)
	}
	anchor, err := ccommon.ChallengeAnchor(transcriptValues, combineHash, plaintextHash)
	if err != nil {
		return nil, nil, fmt.Errorf("hash decrypt combine challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, decryptCombineTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive decrypt combine challenge: %w", err)
	}
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc decrypt combine transcript: %w", err)
	}

	// Compute Lagrange coefficients in the BJJ scalar field (r_bjj).
	// These are passed as private witnesses; the point equality check at the end
	// of the circuit validates that they were used correctly.
	activeIndexes := ccommon.Uint16sToBigInts(a.ParticipantIndexes)
	lagrangeCoeffs, err := ccommon.LagrangeCoefficientsAtZeroNative(activeIndexes)
	if err != nil {
		return nil, nil, fmt.Errorf("compute lagrange coefficients: %w", err)
	}
	lagrangeCoeffs, err = ccommon.PadBigInts(lagrangeCoeffs, MaxShares)
	if err != nil {
		return nil, nil, err
	}

	witness := &DecryptCombineCircuit{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Aid:                  new(big.Int).Set(aid),
		CtIdx:                new(big.Int).Set(ctIdx),
		Mode:                 new(big.Int).Set(mode),
		S:                    new(big.Int).Set(s),
		DeltaOrg:             ccommon.CircuitPoint(deltaOrg),
		Threshold:            threshold,
		ShareCount:           shareCount,
		CombineHash:          combineHash,
		PlaintextHash:        plaintextHash,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
		Plaintext:            new(big.Int).Set(a.Plaintext),
		CiphertextC1:         ccommon.CircuitPoint(a.CiphertextC1),
		CiphertextC2:         ccommon.CircuitPoint(a.CiphertextC2),
	}
	for i := range MaxShares {
		witness.ParticipantIndexes[i] = participantIndexes[i]
		witness.PartialDecryptions[i] = partials[i]
		witness.LagrangeCoefficients[i] = lagrangeCoeffs[i]
	}

	publicInputs := &PublicInputs{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Aid:                  new(big.Int).Set(aid),
		CtIdx:                new(big.Int).Set(ctIdx),
		Mode:                 new(big.Int).Set(mode),
		S:                    new(big.Int).Set(s),
		DeltaOrg:             deltaOrg,
		Threshold:            new(big.Int).Set(threshold),
		ShareCount:           new(big.Int).Set(shareCount),
		CombineHash:          new(big.Int).Set(combineHash),
		PlaintextHash:        new(big.Int).Set(plaintextHash),
		Challenge:            new(big.Int).Set(challenge),
		TranscriptCommitment: new(big.Int).Set(transcriptCommitment),
		CiphertextC1:         a.CiphertextC1,
		CiphertextC2:         a.CiphertextC2,
		ParticipantIndexes:   participantIndexes,
		PartialDecryptions:   paddedPartialDecryptions,
	}
	return witness, publicInputs, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *DecryptCombineCircuit {
	witness := &DecryptCombineCircuit{
		RoundHash:            p.RoundHash,
		Aid:                  p.Aid,
		CtIdx:                p.CtIdx,
		Mode:                 p.Mode,
		S:                    p.S,
		DeltaOrg:             ccommon.CircuitPoint(p.DeltaOrg),
		Threshold:            p.Threshold,
		ShareCount:           p.ShareCount,
		CombineHash:          p.CombineHash,
		PlaintextHash:        p.PlaintextHash,
		Challenge:            p.Challenge,
		TranscriptCommitment: p.TranscriptCommitment,
	}
	return witness
}

// Scalars returns the ordered public scalars used by the verifier.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Aid,
		p.CtIdx,
		p.Mode,
		p.S,
		p.DeltaOrg.X,
		p.DeltaOrg.Y,
		p.Threshold,
		p.ShareCount,
		p.CombineHash,
		p.PlaintextHash,
		p.Challenge,
		p.TranscriptCommitment,
	}
}

// TranscriptScalars returns the compressed transcript vector checked on chain.
func (p PublicInputs) TranscriptScalars() []*big.Int {
	values := []*big.Int{
		p.CiphertextC1.X,
		p.CiphertextC1.Y,
		p.CiphertextC2.X,
		p.CiphertextC2.Y,
	}
	for i := range MaxShares {
		values = append(values, p.ParticipantIndexes[i])
	}
	for i := range MaxShares {
		values = append(values, p.PartialDecryptions[i].X, p.PartialDecryptions[i].Y)
	}
	return values
}
