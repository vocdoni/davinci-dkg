package poolkey

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/types"
)

var poolKeyTranscriptDomain = protocol.DomainPoolKeyTranscriptV1

// PublicInputs is the native representation of the activation public inputs
// plus the transcript vectors the contract streams alongside the proof. Every
// vector is padded to the circuit bounds: inactive participant slots carry a
// zero index and hash and the identity point, coefficient slots beyond the
// threshold carry the identity, and ShareCommitments[i] is D_{i+1} for every
// committee member i < CommitteeSize (the identity beyond).
type PublicInputs struct {
	RoundHash            *big.Int
	Threshold            *big.Int
	CommitteeSize        *big.Int
	AcceptedCount        *big.Int
	KeyIndex             *big.Int
	TranscriptDigest     *big.Int
	Challenge            *big.Int
	TranscriptCommitment *big.Int
	ParticipantIndexes   []*big.Int
	ContributionHashes   []*big.Int
	AggregateCommitments []types.CurvePoint
	ShareCommitments     []types.CurvePoint
	// PoolKey is AggregateCommitments[0], the key activatePoolKey stores.
	PoolKey types.CurvePoint
}

// BuildWitness materializes the native assignment into a gnark witness and the
// corresponding public inputs.
func BuildWitness(a Assignment) (*PoolKeyCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	accepted := len(a.ParticipantIndexes)
	threshold := big.NewInt(int64(a.Threshold))
	committeeSize := big.NewInt(int64(a.CommitteeSize))
	acceptedCount := big.NewInt(int64(accepted))
	keyIndex := new(big.Int).SetUint64(uint64(a.KeyIndex))

	participantIndexes, err := ccommon.PadBigInts(ccommon.Uint16sToBigInts(a.ParticipantIndexes), MaxParticipants)
	if err != nil {
		return nil, nil, err
	}

	// Per contributor: the activated key's commitment vector and the MaxK key
	// digests its stored commitments hash absorbs. Inactive slots stay at the
	// identity / zero, which the circuit masks out of every constraint.
	keyCommitments := make([][]types.CurvePoint, MaxParticipants)
	keyDigests := make([][]*big.Int, MaxParticipants)
	contributionHashes := make([]*big.Int, MaxParticipants)
	for i := range MaxParticipants {
		if keyCommitments[i], err = ccommon.PadPoints(nil, MaxCoefficients); err != nil {
			return nil, nil, err
		}
		keyDigests[i] = make([]*big.Int, MaxKeys)
		for j := range MaxKeys {
			keyDigests[i][j] = big.NewInt(0)
		}
		contributionHashes[i] = big.NewInt(0)
	}
	for i := range accepted {
		for j := range MaxKeys {
			padded, padErr := ccommon.PadPoints(a.Commitments[i][j], MaxCoefficients)
			if padErr != nil {
				return nil, nil, fmt.Errorf("pad contributor %d key %d commitments: %w", i, j, padErr)
			}
			digest, digestErr := ccommon.CommitmentKeyDigestNative(padded)
			if digestErr != nil {
				return nil, nil, fmt.Errorf("digest contributor %d key %d: %w", i, j, digestErr)
			}
			keyDigests[i][j] = digest
			if j == int(a.KeyIndex) {
				keyCommitments[i] = padded
			}
		}
		contributionHashes[i], err = ccommon.CommitmentsHashNative(
			a.RoundHash,
			participantIndexes[i],
			threshold,
			keyDigests[i],
		)
		if err != nil {
			return nil, nil, fmt.Errorf("hash contributor %d commitments: %w", i, err)
		}
	}

	// Ā[m] = Σ_i A_i[m] over the accepted contributors. Slots beyond the
	// threshold are the identity on every row, so they aggregate to it too.
	aggregateCommitments := make([]types.CurvePoint, MaxCoefficients)
	for m := range MaxCoefficients {
		sum := group.NewPoint()
		sum.SetZero()
		for i := range accepted {
			point, decodeErr := group.Decode(keyCommitments[i][m])
			if decodeErr != nil {
				return nil, nil, fmt.Errorf("decode contributor %d commitment %d: %w", i, m, decodeErr)
			}
			sum.Add(sum, point)
		}
		aggregateCommitments[m] = group.Encode(sum)
	}

	// D_p for every committee member p, contributing or not: the Merkle leaf
	// of member p is slot p−1, and a member that skipped the contribution
	// phase still holds a share of every accepted polynomial.
	shareCommitments, err := ccommon.PadPoints(nil, MaxParticipants)
	if err != nil {
		return nil, nil, err
	}
	for i := range int(a.CommitteeSize) {
		shareCommitments[i], err = ccommon.CommitmentPolynomialValueNative(aggregateCommitments, big.NewInt(int64(i+1)))
		if err != nil {
			return nil, nil, fmt.Errorf("evaluate share commitment of member %d: %w", i+1, err)
		}
	}

	transcriptValues := transcriptWords(participantIndexes, contributionHashes, aggregateCommitments, shareCommitments)
	digest, err := TranscriptDigestNative(a.RoundHash, keyIndex, transcriptValues)
	if err != nil {
		return nil, nil, err
	}
	// The anchor commits to the witness transcript (through the Poseidon
	// digest the circuit recomputes) and to the calldata transcript (through
	// keccak) before ρ exists: keccak(digest ‖ keccak(transcript)).
	anchor, err := ccommon.ChallengeAnchor(transcriptValues, digest)
	if err != nil {
		return nil, nil, fmt.Errorf("hash pool key challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, poolKeyTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive pool key challenge: %w", err)
	}
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc pool key transcript: %w", err)
	}

	witness := &PoolKeyCircuit{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            threshold,
		CommitteeSize:        committeeSize,
		AcceptedCount:        acceptedCount,
		KeyIndex:             keyIndex,
		TranscriptDigest:     digest,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
	}
	for i := range MaxParticipants {
		witness.ParticipantIndexes[i] = participantIndexes[i]
		witness.ContributionHashes[i] = contributionHashes[i]
		witness.ShareCommitments[i] = ccommon.CircuitPoint(shareCommitments[i])
		for m := range MaxCoefficients {
			witness.KeyCommitments[i][m] = ccommon.CircuitPoint(keyCommitments[i][m])
		}
		for j := range MaxKeys {
			witness.OtherKeyDigests[i][j] = keyDigests[i][j]
		}
	}
	for m := range MaxCoefficients {
		witness.AggregateCommitments[m] = ccommon.CircuitPoint(aggregateCommitments[m])
	}

	publicInputs := &PublicInputs{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            new(big.Int).Set(threshold),
		CommitteeSize:        new(big.Int).Set(committeeSize),
		AcceptedCount:        new(big.Int).Set(acceptedCount),
		KeyIndex:             new(big.Int).Set(keyIndex),
		TranscriptDigest:     new(big.Int).Set(digest),
		Challenge:            new(big.Int).Set(challenge),
		TranscriptCommitment: new(big.Int).Set(transcriptCommitment),
		ParticipantIndexes:   participantIndexes,
		ContributionHashes:   contributionHashes,
		AggregateCommitments: aggregateCommitments,
		ShareCommitments:     shareCommitments,
		PoolKey:              aggregateCommitments[0],
	}
	return witness, publicInputs, nil
}

// TranscriptDigestNative mirrors the circuit's TranscriptDigest public input:
// Poseidon MultiHash over (roundHash, keyIndex, transcript…), the
// TranscriptWords-long masked transcript in calldata order. 2 + 6·MaxN
// inputs, under the native sponge's 256-input cap.
func TranscriptDigestNative(roundHash, keyIndex *big.Int, transcript []*big.Int) (*big.Int, error) {
	if roundHash == nil || keyIndex == nil {
		return nil, fmt.Errorf("pool key transcript digest: epoch hash and key index are required")
	}
	if len(transcript) != TranscriptWords {
		return nil, fmt.Errorf("pool key transcript digest: got %d words, want %d", len(transcript), TranscriptWords)
	}
	inputs := make([]*big.Int, 0, 2+TranscriptWords)
	inputs = append(inputs, roundHash, keyIndex)
	inputs = append(inputs, transcript...)
	digest, err := ccommon.MultiHashNative(inputs...)
	if err != nil {
		return nil, fmt.Errorf("pool key transcript digest: %w", err)
	}
	return digest, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *PoolKeyCircuit {
	return &PoolKeyCircuit{
		RoundHash:            p.RoundHash,
		Threshold:            p.Threshold,
		CommitteeSize:        p.CommitteeSize,
		AcceptedCount:        p.AcceptedCount,
		KeyIndex:             p.KeyIndex,
		TranscriptDigest:     p.TranscriptDigest,
		Challenge:            p.Challenge,
		TranscriptCommitment: p.TranscriptCommitment,
	}
}

// Scalars returns the 8 ordered public scalars used by the verifier.
// Order must match the circuit field declaration order.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Threshold,
		p.CommitteeSize,
		p.AcceptedCount,
		p.KeyIndex,
		p.TranscriptDigest,
		p.Challenge,
		p.TranscriptCommitment,
	}
}

// TranscriptScalars returns the TranscriptWords-long activation transcript the
// contract streams into its BRLC check.
func (p PublicInputs) TranscriptScalars() ([]*big.Int, error) {
	indexes, err := ccommon.PadBigInts(p.ParticipantIndexes, MaxParticipants)
	if err != nil {
		return nil, fmt.Errorf("pool key transcript: pad participant indexes: %w", err)
	}
	hashes, err := ccommon.PadBigInts(p.ContributionHashes, MaxParticipants)
	if err != nil {
		return nil, fmt.Errorf("pool key transcript: pad contribution hashes: %w", err)
	}
	aggregate, err := ccommon.PadPoints(p.AggregateCommitments, MaxCoefficients)
	if err != nil {
		return nil, fmt.Errorf("pool key transcript: pad aggregate commitments: %w", err)
	}
	shares, err := ccommon.PadPoints(p.ShareCommitments, MaxParticipants)
	if err != nil {
		return nil, fmt.Errorf("pool key transcript: pad share commitments: %w", err)
	}
	return transcriptWords(indexes, hashes, aggregate, shares), nil
}

// BRLCCommitment compresses the activation transcript into one scalar
// commitment, matching the circuit's TranscriptCommitment public input and the
// contract's on-chain check.
func (p PublicInputs) BRLCCommitment(challenge *big.Int) (*big.Int, error) {
	scalars, err := p.TranscriptScalars()
	if err != nil {
		return nil, err
	}
	return ccommon.BRLCNative(challenge, scalars...)
}

// transcriptWords lays out the activation transcript: participant indexes,
// contribution hashes, aggregate commitments, share commitments. Every
// argument must already be padded to the circuit bounds.
func transcriptWords(
	participantIndexes, contributionHashes []*big.Int,
	aggregateCommitments, shareCommitments []types.CurvePoint,
) []*big.Int {
	words := make([]*big.Int, 0, TranscriptWords)
	words = append(words, participantIndexes...)
	words = append(words, contributionHashes...)
	for m := range MaxCoefficients {
		words = append(words, aggregateCommitments[m].X, aggregateCommitments[m].Y)
	}
	for i := range MaxParticipants {
		words = append(words, shareCommitments[i].X, shareCommitments[i].Y)
	}
	return words
}
