package decryptcombine

import (
	"fmt"
	"math/big"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/dleq"
	"github.com/vocdoni/davinci-dkg/types"
)

var decryptCombineTranscriptDomain = ethcrypto.Keccak256Hash([]byte("davinci-dkg:decrypt-combine:v1"))

// PublicInputs is the native representation of the decrypt-combine public
// inputs plus the private words the contract re-checks through the
// calldata transcript.
type PublicInputs struct {
	RoundHash            *big.Int // semantically: eid
	Aid                  *big.Int
	CtIdx                *big.Int
	DeltaOrg             types.CurvePoint
	Threshold            *big.Int
	ShareCount           *big.Int
	CombineHash          *big.Int
	PlaintextHash        *big.Int
	Challenge            *big.Int
	TranscriptCommitment *big.Int
	CiphertextC1         types.CurvePoint
	CiphertextC2         types.CurvePoint
	OrganizerPK          types.CurvePoint
	OrganizerA1          types.CurvePoint
	OrganizerA2          types.CurvePoint
	OrganizerZ           *big.Int
	OrganizerE           *big.Int
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
	aid := new(big.Int)
	if a.Aid != nil {
		aid.Set(a.Aid)
	}
	ctIdx := new(big.Int)
	if a.CtIdx != nil {
		ctIdx.Set(a.CtIdx)
	}
	organizerZ := new(big.Int).Set(a.OrganizerProof.Response)
	// `e` is derived here, never taken from the caller: the contract
	// recomputes exactly this keccak over the calldata and pins the
	// transcript word to it, so a divergent `e` would make the combine
	// revert rather than produce a wrong plaintext.
	organizerE := dleq.OrganizerShareChallenge(
		epochIDBytes(a.RoundHash),
		aidBytes(aid),
		uint16(ctIdx.Uint64()),
		a.OrganizerPK,
		a.CiphertextC1,
		a.DeltaOrg,
		a.OrganizerProof.A1,
		a.OrganizerProof.A2,
	)

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

	// Poseidon digest order must match DecryptCombineCircuit.Define.
	hashInputs := []*big.Int{
		a.RoundHash, // eid
		aid,
		ctIdx,
		a.DeltaOrg.X,
		a.DeltaOrg.Y,
		threshold,
		shareCount,
		a.CiphertextC1.X,
		a.CiphertextC1.Y,
		a.CiphertextC2.X,
		a.CiphertextC2.Y,
		a.OrganizerPK.X,
		a.OrganizerPK.Y,
		a.OrganizerProof.A1.X,
		a.OrganizerProof.A1.Y,
		a.OrganizerProof.A2.X,
		a.OrganizerProof.A2.Y,
		organizerZ,
		organizerE,
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
	transcriptValues := transcriptWords(
		a.CiphertextC1, a.CiphertextC2,
		a.OrganizerPK, a.OrganizerProof.A1, a.OrganizerProof.A2,
		organizerZ, organizerE,
		participantIndexes, paddedPartialDecryptions,
	)
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
		DeltaOrg:             ccommon.CircuitPoint(a.DeltaOrg),
		Threshold:            threshold,
		ShareCount:           shareCount,
		CombineHash:          combineHash,
		PlaintextHash:        plaintextHash,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
		Plaintext:            new(big.Int).Set(a.Plaintext),
		CiphertextC1:         ccommon.CircuitPoint(a.CiphertextC1),
		CiphertextC2:         ccommon.CircuitPoint(a.CiphertextC2),
		OrganizerPK:          ccommon.CircuitPoint(a.OrganizerPK),
		OrganizerA1:          ccommon.CircuitPoint(a.OrganizerProof.A1),
		OrganizerA2:          ccommon.CircuitPoint(a.OrganizerProof.A2),
		OrganizerZ:           new(big.Int).Set(organizerZ),
		OrganizerE:           new(big.Int).Set(organizerE),
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
		DeltaOrg:             a.DeltaOrg,
		Threshold:            new(big.Int).Set(threshold),
		ShareCount:           new(big.Int).Set(shareCount),
		CombineHash:          new(big.Int).Set(combineHash),
		PlaintextHash:        new(big.Int).Set(plaintextHash),
		Challenge:            new(big.Int).Set(challenge),
		TranscriptCommitment: new(big.Int).Set(transcriptCommitment),
		CiphertextC1:         a.CiphertextC1,
		CiphertextC2:         a.CiphertextC2,
		OrganizerPK:          a.OrganizerPK,
		OrganizerA1:          a.OrganizerProof.A1,
		OrganizerA2:          a.OrganizerProof.A2,
		OrganizerZ:           new(big.Int).Set(organizerZ),
		OrganizerE:           new(big.Int).Set(organizerE),
		ParticipantIndexes:   participantIndexes,
		PartialDecryptions:   paddedPartialDecryptions,
	}
	return witness, publicInputs, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *DecryptCombineCircuit {
	return &DecryptCombineCircuit{
		RoundHash:            p.RoundHash,
		Aid:                  p.Aid,
		CtIdx:                p.CtIdx,
		DeltaOrg:             ccommon.CircuitPoint(p.DeltaOrg),
		Threshold:            p.Threshold,
		ShareCount:           p.ShareCount,
		CombineHash:          p.CombineHash,
		PlaintextHash:        p.PlaintextHash,
		Challenge:            p.Challenge,
		TranscriptCommitment: p.TranscriptCommitment,
	}
}

// Scalars returns the 11 ordered public inputs the on-chain verifier reads.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Aid,
		p.CtIdx,
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

// TranscriptScalars returns the `TranscriptWords` calldata transcript the
// contract re-hashes against the stored ciphertext, the application's
// organizer key and the stored organizer share before folding it with ρ.
func (p PublicInputs) TranscriptScalars() []*big.Int {
	return transcriptWords(
		p.CiphertextC1, p.CiphertextC2,
		p.OrganizerPK, p.OrganizerA1, p.OrganizerA2,
		p.OrganizerZ, p.OrganizerE,
		p.ParticipantIndexes, p.PartialDecryptions,
	)
}

// transcriptWords lays out the calldata transcript in the one canonical
// order shared by the circuit, the witness builder and the contract:
//
//	[0..3]              C1.x C1.y C2.x C2.y
//	[4..5]              PK_org.x PK_org.y
//	[6..7]              A1.x A1.y
//	[8..9]              A2.x A2.y
//	[10]                z
//	[11]                e
//	[12 .. 12+MaxShares)          participant indexes (0 in inactive slots)
//	[12+MaxShares .. 12+3·MaxShares)  partials as (x, y) pairs ((0,1) inactive)
func transcriptWords(
	c1, c2, pkOrg, a1, a2 types.CurvePoint,
	z, e *big.Int,
	indexes []*big.Int,
	partials []types.CurvePoint,
) []*big.Int {
	values := make([]*big.Int, 0, TranscriptWords)
	values = append(
		values,
		c1.X, c1.Y, c2.X, c2.Y,
		pkOrg.X, pkOrg.Y,
		a1.X, a1.Y,
		a2.X, a2.Y,
		z,
		e,
	)
	values = append(values, indexes...)
	for i := range partials {
		values = append(values, partials[i].X, partials[i].Y)
	}
	return values
}

// epochIDBytes renders the epoch id as the 12 bytes the keccak transcripts
// use. Assignment.Validate has already bounded it to 96 bits.
func epochIDBytes(roundHash *big.Int) [12]byte {
	var out [12]byte
	roundHash.FillBytes(out[:])
	return out
}

// aidBytes renders the application id as the 32 bytes the keccak
// transcripts use. Assignment.Validate has already bounded it to 256 bits.
func aidBytes(aid *big.Int) [32]byte {
	var out [32]byte
	aid.FillBytes(out[:])
	return out
}
