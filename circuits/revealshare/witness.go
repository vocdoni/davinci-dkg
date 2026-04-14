package revealshare

import (
	"fmt"
	"math/big"

	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

var revealShareTranscriptDomain = ethcrypto.Keccak256Hash([]byte("davinci-dkg:reveal-share:v1"))

// PublicInputs is the native representation of the reveal-share public inputs.
type PublicInputs struct {
	RoundHash               *big.Int
	Threshold               *big.Int
	ShareCount              *big.Int
	DisclosureHash          *big.Int
	ReconstructedSecretHash *big.Int
	Challenge               *big.Int
	TranscriptCommitment    *big.Int
	ParticipantIndexes      []*big.Int
	RevealedShares          []*big.Int
}

// BuildWitness materializes the reveal-share native assignment.
func BuildWitness(a Assignment) (*RevealShareCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	threshold := big.NewInt(int64(a.Threshold))
	shareCount := big.NewInt(int64(len(a.ParticipantIndexes)))
	participantIndexes := ccommon.Uint16sToBigInts(a.ParticipantIndexes)
	participantIndexes, err := ccommon.PadBigInts(participantIndexes, MaxShares)
	if err != nil {
		return nil, nil, err
	}
	revealedShares, err := ccommon.PadBigInts(a.RevealedShares, MaxShares)
	if err != nil {
		return nil, nil, err
	}

	reconstructedSecret, err := ccommon.InterpolateAtZeroNative(
		ccommon.Uint16sToBigInts(a.ParticipantIndexes),
		a.RevealedShares,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("reconstruct secret: %w", err)
	}

	hashInputs := []*big.Int{a.RoundHash, threshold, shareCount}
	for i := range len(a.ParticipantIndexes) {
		hashInputs = append(hashInputs, big.NewInt(int64(a.ParticipantIndexes[i])), a.RevealedShares[i])
	}
	for i := len(a.ParticipantIndexes); i < MaxShares; i++ {
		hashInputs = append(hashInputs, big.NewInt(0), big.NewInt(0))
	}
	disclosureHash, err := ccommon.MultiHashNative(hashInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("hash disclosed shares: %w", err)
	}
	reconstructedSecretHash := new(big.Int).Set(reconstructedSecret)
	anchor, err := ccommon.HashPackedBigIntsNative(disclosureHash, reconstructedSecretHash)
	if err != nil {
		return nil, nil, fmt.Errorf("hash reveal share challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, revealShareTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive reveal share challenge: %w", err)
	}
	transcriptValues := make([]*big.Int, 0, 16)
	transcriptValues = append(transcriptValues, participantIndexes...)
	transcriptValues = append(transcriptValues, revealedShares...)
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc reveal share transcript: %w", err)
	}

	witness := &RevealShareCircuit{
		RoundHash:               new(big.Int).Set(a.RoundHash),
		Threshold:               threshold,
		ShareCount:              shareCount,
		DisclosureHash:          disclosureHash,
		ReconstructedSecretHash: reconstructedSecretHash,
		Challenge:               challenge,
		TranscriptCommitment:    transcriptCommitment,
		ReconstructedSecret:     reconstructedSecret,
	}
	for i := range MaxShares {
		witness.ParticipantIndexes[i] = participantIndexes[i]
		witness.RevealedShares[i] = revealedShares[i]
	}

	publicInputs := &PublicInputs{
		RoundHash:               new(big.Int).Set(a.RoundHash),
		Threshold:               new(big.Int).Set(threshold),
		ShareCount:              new(big.Int).Set(shareCount),
		DisclosureHash:          new(big.Int).Set(disclosureHash),
		ReconstructedSecretHash: new(big.Int).Set(reconstructedSecretHash),
		Challenge:               new(big.Int).Set(challenge),
		TranscriptCommitment:    new(big.Int).Set(transcriptCommitment),
		ParticipantIndexes:      participantIndexes,
		RevealedShares:          revealedShares,
	}
	return witness, publicInputs, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *RevealShareCircuit {
	witness := &RevealShareCircuit{
		RoundHash:               p.RoundHash,
		Threshold:               p.Threshold,
		ShareCount:              p.ShareCount,
		DisclosureHash:          p.DisclosureHash,
		ReconstructedSecretHash: p.ReconstructedSecretHash,
		Challenge:               p.Challenge,
		TranscriptCommitment:    p.TranscriptCommitment,
	}
	return witness
}

// Scalars returns the ordered public scalars used by the verifier.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Threshold,
		p.ShareCount,
		p.DisclosureHash,
		p.ReconstructedSecretHash,
		p.Challenge,
		p.TranscriptCommitment,
	}
}

// TranscriptScalars returns the compressed transcript vector checked on chain.
func (p PublicInputs) TranscriptScalars() []*big.Int {
	values := make([]*big.Int, 0, 16)
	for i := range MaxShares {
		values = append(values, p.ParticipantIndexes[i])
	}
	for i := range MaxShares {
		values = append(values, p.RevealedShares[i])
	}
	return values
}
