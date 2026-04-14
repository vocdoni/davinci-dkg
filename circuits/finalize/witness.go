package finalize

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/types"
)

var finalizeTranscriptDomain = ethcrypto.Keccak256Hash([]byte("davinci-dkg:finalize:v1"))

// PublicInputs is the native representation of the finalize public inputs.
type PublicInputs struct {
	RoundHash               *big.Int
	Threshold               *big.Int
	CommitteeSize           *big.Int
	AcceptedCount           *big.Int
	AggregateHash           *big.Int
	CollectivePublicKey     *big.Int
	ShareCommitmentHash     *big.Int
	Challenge               *big.Int
	TranscriptCommitment    *big.Int
	ParticipantIndexes      []*big.Int
	ContributionCommitments [MaxParticipants][MaxCoefficients]types.CurvePoint
	AggregateCommitments    []types.CurvePoint
	ShareCommitments        []types.CurvePoint
}

// BuildWitness materializes the finalize native assignment.
func BuildWitness(a Assignment) (*FinalizeCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	modulus := ecc.BN254.ScalarField()
	threshold := big.NewInt(int64(a.Threshold))
	committeeSize := big.NewInt(int64(a.CommitteeSize))
	acceptedCount := big.NewInt(int64(len(a.ParticipantIndexes)))

	participantIndexes := ccommon.Uint16sToBigInts(a.ParticipantIndexes)
	participantIndexes, err := ccommon.PadBigInts(participantIndexes, MaxParticipants)
	if err != nil {
		return nil, nil, err
	}

	var contributionCommitments [MaxParticipants][MaxCoefficients]types.CurvePoint
	aggregateCommitments := make([]types.CurvePoint, MaxCoefficients)
	for i := range MaxParticipants {
		for k := range MaxCoefficients {
			contributionCommitments[i][k] = types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
		}
	}
	for k := range MaxCoefficients {
		aggregateCommitments[k] = types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	}
	for i := range len(a.ContributionCoefficients) {
		padded, err := ccommon.PadBigInts(a.ContributionCoefficients[i], MaxCoefficients)
		if err != nil {
			return nil, nil, fmt.Errorf("pad contribution %d: %w", i, err)
		}
		for k := range MaxCoefficients {
			commitmentPoint := group.NewPoint()
			commitmentPoint.ScalarBaseMult(padded[k])
			contributionCommitments[i][k] = group.Encode(commitmentPoint)
			if i == 0 {
				aggregateCommitments[k] = group.Encode(commitmentPoint)
				continue
			}
			acc, err := group.Decode(aggregateCommitments[k])
			if err != nil {
				return nil, nil, fmt.Errorf("decode aggregate commitment %d: %w", k, err)
			}
			acc.Add(acc, commitmentPoint)
			aggregateCommitments[k] = group.Encode(acc)
		}
	}
	aggregateInputs := []*big.Int{a.RoundHash, threshold, committeeSize, acceptedCount}
	for k := range MaxCoefficients {
		if k < int(a.Threshold) {
			aggregateInputs = append(aggregateInputs, aggregateCommitments[k].X, aggregateCommitments[k].Y)
			continue
		}
		aggregateInputs = append(aggregateInputs, big.NewInt(0), big.NewInt(1))
	}
	aggregateHash, err := ccommon.MultiHashNative(aggregateInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("hash aggregate coefficients: %w", err)
	}

	shareInputs := []*big.Int{a.RoundHash, threshold, committeeSize, acceptedCount}
	shareCommitments := make([]types.CurvePoint, 0, len(a.ParticipantIndexes))
	for i := range len(a.ParticipantIndexes) {
		sharePoint := group.NewPoint()
		sharePoint.SetZero()
		power := big.NewInt(1)
		x := big.NewInt(int64(a.ParticipantIndexes[i]))
		for k := 0; k < int(a.Threshold); k++ {
			commitmentPoint, err := group.Decode(aggregateCommitments[k])
			if err != nil {
				return nil, nil, fmt.Errorf("decode aggregate point %d: %w", k, err)
			}
			term := group.NewPoint()
			term.ScalarMult(commitmentPoint, power)
			sharePoint.Add(sharePoint, term)
			power.Mul(power, x)
			power.Mod(power, modulus)
		}
		encoded := group.Encode(sharePoint)
		shareCommitments = append(shareCommitments, encoded)
		shareInputs = append(shareInputs, big.NewInt(int64(a.ParticipantIndexes[i])), encoded.X, encoded.Y)
	}
	for i := len(a.ParticipantIndexes); i < MaxParticipants; i++ {
		shareInputs = append(shareInputs, big.NewInt(0), big.NewInt(0), big.NewInt(1))
	}
	shareHash, err := ccommon.MultiHashNative(shareInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("hash share commitments: %w", err)
	}

	publicKeyHash, err := ccommon.MultiHashNative(a.RoundHash, aggregateCommitments[0].X, aggregateCommitments[0].Y)
	if err != nil {
		return nil, nil, fmt.Errorf("hash collective public key: %w", err)
	}

	anchor, err := ccommon.HashPackedBigIntsNative(aggregateHash, publicKeyHash, shareHash)
	if err != nil {
		return nil, nil, fmt.Errorf("hash finalize challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, finalizeTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive finalize challenge: %w", err)
	}
	transcriptValues := make([]*big.Int, 0, 168)
	for i := range MaxParticipants {
		transcriptValues = append(transcriptValues, participantIndexes[i])
	}
	for i := range MaxParticipants {
		for k := range MaxCoefficients {
			point := contributionCommitments[i][k]
			transcriptValues = append(transcriptValues, point.X, point.Y)
		}
	}
	for k := range MaxCoefficients {
		point := aggregateCommitments[k]
		transcriptValues = append(transcriptValues, point.X, point.Y)
	}
	for i := range MaxParticipants {
		if i < len(shareCommitments) {
			transcriptValues = append(transcriptValues, shareCommitments[i].X, shareCommitments[i].Y)
			continue
		}
		transcriptValues = append(transcriptValues, big.NewInt(0), big.NewInt(1))
	}
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc finalize transcript: %w", err)
	}

	witness := &FinalizeCircuit{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            threshold,
		CommitteeSize:        committeeSize,
		AcceptedCount:        acceptedCount,
		AggregateHash:        aggregateHash,
		CollectivePublicKey:  publicKeyHash,
		ShareCommitmentHash:  shareHash,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
	}
	for i := range MaxParticipants {
		witness.ParticipantIndexes[i] = participantIndexes[i]
		for k := range MaxCoefficients {
			witness.ContributionCommitments[i][k] = ccommon.CircuitPoint(contributionCommitments[i][k])
		}
	}
	for k := range MaxCoefficients {
		witness.AggregateCommitments[k] = ccommon.CircuitPoint(aggregateCommitments[k])
	}
	for i := range MaxParticipants {
		if i < len(shareCommitments) {
			witness.ShareCommitments[i] = ccommon.CircuitPoint(shareCommitments[i])
			continue
		}
		witness.ShareCommitments[i] = ccommon.IdentityPoint()
	}

	publicInputs := &PublicInputs{
		RoundHash:               new(big.Int).Set(a.RoundHash),
		Threshold:               new(big.Int).Set(threshold),
		CommitteeSize:           new(big.Int).Set(committeeSize),
		AcceptedCount:           new(big.Int).Set(acceptedCount),
		AggregateHash:           new(big.Int).Set(aggregateHash),
		CollectivePublicKey:     new(big.Int).Set(publicKeyHash),
		ShareCommitmentHash:     new(big.Int).Set(shareHash),
		Challenge:               new(big.Int).Set(challenge),
		TranscriptCommitment:    new(big.Int).Set(transcriptCommitment),
		ParticipantIndexes:      participantIndexes,
		ContributionCommitments: contributionCommitments,
		AggregateCommitments:    aggregateCommitments,
		ShareCommitments:        shareCommitments,
	}
	return witness, publicInputs, nil
}

// CommitmentPointsAssignment is like Assignment but takes pre-computed
// Feldman commitment points instead of raw polynomial scalars. This lets
// the testnet runner (which doesn't hold private keys) build the finalize
// proof directly from on-chain transcript data.
type CommitmentPointsAssignment struct {
	RoundHash          *big.Int
	Threshold          uint16
	CommitteeSize      uint16
	ParticipantIndexes []uint16
	// ContributionPoints[i][k] = C_i(k) = a_{i,k} * G (public Feldman commitment).
	ContributionPoints [][]types.CurvePoint
}

// BuildWitnessFromCommitmentPoints is the public-data variant of BuildWitness:
// it computes the same proof from pre-computed commitment points rather than
// raw polynomial coefficients.
func BuildWitnessFromCommitmentPoints(a CommitmentPointsAssignment) (*FinalizeCircuit, *PublicInputs, error) {
	// Convert to the coefficient-based assignment by using the points directly
	// (set coefficients[k] = 1 and override the commitment computation).
	// We do this by building the witness manually from the provided points.
	if a.RoundHash == nil {
		return nil, nil, fmt.Errorf("round hash is required")
	}
	if a.Threshold == 0 || a.CommitteeSize == 0 {
		return nil, nil, fmt.Errorf("threshold and committee size are required")
	}
	if len(a.ParticipantIndexes) < int(a.Threshold) {
		return nil, nil, fmt.Errorf("participant count below threshold")
	}
	if len(a.ContributionPoints) != len(a.ParticipantIndexes) {
		return nil, nil, fmt.Errorf("contribution point count mismatch")
	}

	modulus := ecc.BN254.ScalarField()
	threshold := big.NewInt(int64(a.Threshold))
	committeeSize := big.NewInt(int64(a.CommitteeSize))
	acceptedCount := big.NewInt(int64(len(a.ParticipantIndexes)))

	participantIndexes := ccommon.Uint16sToBigInts(a.ParticipantIndexes)
	participantIndexes, err := ccommon.PadBigInts(participantIndexes, MaxParticipants)
	if err != nil {
		return nil, nil, err
	}

	// Fill in contribution commitment points (identity for unused slots).
	var contributionCommitments [MaxParticipants][MaxCoefficients]types.CurvePoint
	for i := range MaxParticipants {
		for k := range MaxCoefficients {
			contributionCommitments[i][k] = types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
		}
	}
	for i, pts := range a.ContributionPoints {
		for k, pt := range pts {
			if k >= MaxCoefficients {
				break
			}
			if pt.X != nil && pt.Y != nil {
				contributionCommitments[i][k] = pt
			}
		}
	}

	// Aggregate: Cbar(k) = Σ_i C_i(k).
	aggregateCommitments := make([]types.CurvePoint, MaxCoefficients)
	for k := range MaxCoefficients {
		aggregateCommitments[k] = types.CurvePoint{X: big.NewInt(0), Y: big.NewInt(1)}
	}
	for i := range len(a.ParticipantIndexes) {
		for k := range MaxCoefficients {
			if i == 0 {
				aggregateCommitments[k] = contributionCommitments[i][k]
				continue
			}
			acc, err := group.Decode(aggregateCommitments[k])
			if err != nil {
				return nil, nil, err
			}
			term, err := group.Decode(contributionCommitments[i][k])
			if err != nil {
				return nil, nil, err
			}
			acc.Add(acc, term)
			aggregateCommitments[k] = group.Encode(acc)
		}
	}

	aggregateInputs := []*big.Int{a.RoundHash, threshold, committeeSize, acceptedCount}
	for k := range MaxCoefficients {
		if k < int(a.Threshold) {
			aggregateInputs = append(aggregateInputs, aggregateCommitments[k].X, aggregateCommitments[k].Y)
		} else {
			aggregateInputs = append(aggregateInputs, big.NewInt(0), big.NewInt(1))
		}
	}
	aggregateHash, err := ccommon.MultiHashNative(aggregateInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("aggregate hash: %w", err)
	}

	shareInputs := []*big.Int{a.RoundHash, threshold, committeeSize, acceptedCount}
	shareCommitments := make([]types.CurvePoint, 0, len(a.ParticipantIndexes))
	for i := range len(a.ParticipantIndexes) {
		sharePoint := group.NewPoint()
		sharePoint.SetZero()
		power := big.NewInt(1)
		x := big.NewInt(int64(a.ParticipantIndexes[i]))
		for k := 0; k < int(a.Threshold); k++ {
			cp, err := group.Decode(aggregateCommitments[k])
			if err != nil {
				return nil, nil, err
			}
			term := group.NewPoint()
			term.ScalarMult(cp, power)
			sharePoint.Add(sharePoint, term)
			power.Mul(power, x)
			power.Mod(power, modulus)
		}
		enc := group.Encode(sharePoint)
		shareCommitments = append(shareCommitments, enc)
		shareInputs = append(shareInputs, big.NewInt(int64(a.ParticipantIndexes[i])), enc.X, enc.Y)
	}
	for i := len(a.ParticipantIndexes); i < MaxParticipants; i++ {
		shareInputs = append(shareInputs, big.NewInt(0), big.NewInt(0), big.NewInt(1))
	}
	shareHash, err := ccommon.MultiHashNative(shareInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("share hash: %w", err)
	}
	publicKeyHash, err := ccommon.MultiHashNative(a.RoundHash, aggregateCommitments[0].X, aggregateCommitments[0].Y)
	if err != nil {
		return nil, nil, fmt.Errorf("pk hash: %w", err)
	}
	anchor, err := ccommon.HashPackedBigIntsNative(aggregateHash, publicKeyHash, shareHash)
	if err != nil {
		return nil, nil, err
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, finalizeTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, err
	}

	transcriptValues := make([]*big.Int, 0, 168)
	for i := range MaxParticipants {
		transcriptValues = append(transcriptValues, participantIndexes[i])
	}
	for i := range MaxParticipants {
		for k := range MaxCoefficients {
			p := contributionCommitments[i][k]
			transcriptValues = append(transcriptValues, p.X, p.Y)
		}
	}
	for k := range MaxCoefficients {
		p := aggregateCommitments[k]
		transcriptValues = append(transcriptValues, p.X, p.Y)
	}
	for i := range MaxParticipants {
		if i < len(shareCommitments) {
			transcriptValues = append(transcriptValues, shareCommitments[i].X, shareCommitments[i].Y)
		} else {
			transcriptValues = append(transcriptValues, big.NewInt(0), big.NewInt(1))
		}
	}
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, err
	}

	witness := &FinalizeCircuit{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            threshold,
		CommitteeSize:        committeeSize,
		AcceptedCount:        acceptedCount,
		AggregateHash:        aggregateHash,
		CollectivePublicKey:  publicKeyHash,
		ShareCommitmentHash:  shareHash,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
	}
	for i := range MaxParticipants {
		witness.ParticipantIndexes[i] = participantIndexes[i]
		for k := range MaxCoefficients {
			witness.ContributionCommitments[i][k] = ccommon.CircuitPoint(contributionCommitments[i][k])
		}
	}
	for k := range MaxCoefficients {
		witness.AggregateCommitments[k] = ccommon.CircuitPoint(aggregateCommitments[k])
	}
	for i := range MaxParticipants {
		if i < len(shareCommitments) {
			witness.ShareCommitments[i] = ccommon.CircuitPoint(shareCommitments[i])
		} else {
			witness.ShareCommitments[i] = ccommon.IdentityPoint()
		}
	}

	pi := &PublicInputs{
		RoundHash:               new(big.Int).Set(a.RoundHash),
		Threshold:               new(big.Int).Set(threshold),
		CommitteeSize:           new(big.Int).Set(committeeSize),
		AcceptedCount:           new(big.Int).Set(acceptedCount),
		AggregateHash:           new(big.Int).Set(aggregateHash),
		CollectivePublicKey:     new(big.Int).Set(publicKeyHash),
		ShareCommitmentHash:     new(big.Int).Set(shareHash),
		Challenge:               new(big.Int).Set(challenge),
		TranscriptCommitment:    new(big.Int).Set(transcriptCommitment),
		ParticipantIndexes:      participantIndexes,
		ContributionCommitments: contributionCommitments,
		AggregateCommitments:    aggregateCommitments,
		ShareCommitments:        shareCommitments,
	}
	return witness, pi, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *FinalizeCircuit {
	return &FinalizeCircuit{
		RoundHash:            p.RoundHash,
		Threshold:            p.Threshold,
		CommitteeSize:        p.CommitteeSize,
		AcceptedCount:        p.AcceptedCount,
		AggregateHash:        p.AggregateHash,
		CollectivePublicKey:  p.CollectivePublicKey,
		ShareCommitmentHash:  p.ShareCommitmentHash,
		Challenge:            p.Challenge,
		TranscriptCommitment: p.TranscriptCommitment,
	}
}

// Scalars returns the ordered public scalars used by the verifier.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Threshold,
		p.CommitteeSize,
		p.AcceptedCount,
		p.AggregateHash,
		p.CollectivePublicKey,
		p.ShareCommitmentHash,
		p.Challenge,
		p.TranscriptCommitment,
	}
}

func (p PublicInputs) TranscriptScalars() []*big.Int {
	values := make([]*big.Int, 0, 168)
	for i := range MaxParticipants {
		values = append(values, p.ParticipantIndexes[i])
	}
	for i := range MaxParticipants {
		for k := range MaxCoefficients {
			point := p.ContributionCommitments[i][k]
			values = append(values, point.X, point.Y)
		}
	}
	for k := range MaxCoefficients {
		point := p.AggregateCommitments[k]
		values = append(values, point.X, point.Y)
	}
	for i := range MaxParticipants {
		if i < len(p.ShareCommitments) {
			values = append(values, p.ShareCommitments[i].X, p.ShareCommitments[i].Y)
			continue
		}
		values = append(values, big.NewInt(0), big.NewInt(1))
	}
	return values
}

// BRLCCommitment compresses the public input vector into one scalar commitment.
func (p PublicInputs) BRLCCommitment(challenge *big.Int) (*big.Int, error) {
	return ccommon.BRLCNative(challenge, p.Scalars()...)
}
