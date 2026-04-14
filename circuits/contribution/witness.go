package contribution

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/types"
)

var contributionTranscriptDomain = ethcrypto.Keccak256Hash([]byte("davinci-dkg:contribution:v1"))

// PublicInputs is the native representation of the public contribution inputs.
type PublicInputs struct {
	RoundHash            *big.Int
	Threshold            *big.Int
	CommitteeSize        *big.Int
	ContributorIndex     *big.Int
	CommitmentHash       *big.Int
	ShareHash            *big.Int
	Challenge            *big.Int
	TranscriptCommitment *big.Int
	Commitments          []types.CurvePoint
	RecipientKeys        []types.CurvePoint
	Shares               []*big.Int
	EncryptedShares      []types.EncryptedShare
	RecipientIndexes     []*big.Int
}

// BuildWitness materializes the native assignment into a gnark witness and the
// corresponding public inputs.
func BuildWitness(a Assignment) (*ContributionCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	coefficients, err := ccommon.PadBigInts(a.Coefficients, MaxCoefficients)
	if err != nil {
		return nil, nil, err
	}

	recipientIndexes := ccommon.Uint16sToBigInts(a.RecipientIndexes)
	recipientIndexes, err = ccommon.PadBigInts(recipientIndexes, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}

	shares := make([]*big.Int, 0, len(a.RecipientIndexes))
	for i := range a.CommitteeSize {
		share, err := ccommon.EvaluatePolynomialNative(a.Coefficients, big.NewInt(int64(a.RecipientIndexes[i])))
		if err != nil {
			return nil, nil, fmt.Errorf("evaluate share %d: %w", i, err)
		}
		shares = append(shares, share)
	}
	shares, err = ccommon.PadBigInts(shares, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}

	commitments := make([]types.CurvePoint, len(a.Coefficients))
	for i, coefficient := range a.Coefficients {
		point := group.NewPoint()
		point.ScalarBaseMult(coefficient)
		commitments[i] = group.Encode(point)
	}
	paddedCommitments, err := ccommon.CircuitPoints(commitments, MaxCoefficients)
	if err != nil {
		return nil, nil, err
	}

	recipientKeys := make([]types.CurvePoint, len(a.RecipientKeys))
	for i, key := range a.RecipientKeys {
		recipientKeys[i] = types.CurvePoint{X: key.PubX, Y: key.PubY}
	}
	paddedRecipientKeys, err := ccommon.CircuitPoints(recipientKeys, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}

	encryptionNonces, err := ccommon.PadBigInts(a.EncryptionNonces, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}

	encryptedShares := make([]types.EncryptedShare, 0, len(a.RecipientIndexes))
	shareMasks := make([]*big.Int, 0, len(a.RecipientIndexes))
	maskQuotients := make([]*big.Int, 0, len(a.RecipientIndexes))
	maskedShareCarries := make([]*big.Int, 0, len(a.RecipientIndexes))
	if len(a.RecipientKeys) == len(a.RecipientIndexes) && len(a.EncryptionNonces) == len(a.RecipientIndexes) {
		for i := range a.RecipientIndexes {
			ciphertext, encryptErr := shareenc.EncryptShareWithNonceRoundHash(
				a.RoundHash,
				a.ContributorIndex,
				a.RecipientIndexes[i],
				shares[i],
				a.RecipientKeys[i],
				a.EncryptionNonces[i],
			)
			if encryptErr != nil {
				return nil, nil, fmt.Errorf("encrypt share %d: %w", i, encryptErr)
			}
			shareMask := new(big.Int).Sub(ciphertext.MaskedShare, shares[i])
			shareMask.Mod(shareMask, group.ScalarField())
			recipientPoint, decodeErr := group.Decode(types.CurvePoint{X: a.RecipientKeys[i].PubX, Y: a.RecipientKeys[i].PubY})
			if decodeErr != nil {
				return nil, nil, fmt.Errorf("decode recipient key %d: %w", i, decodeErr)
			}
			sharedPoint := group.NewPoint()
			sharedPoint.ScalarMult(recipientPoint, a.EncryptionNonces[i])
			shared := group.Encode(sharedPoint)
			meta, hashErr := dkghash.HashFieldElements(
				ccommon.ShareEncryptionDomain(),
				a.RoundHash,
				new(big.Int).SetUint64((uint64(a.ContributorIndex)<<16)|uint64(a.RecipientIndexes[i])),
			)
			if hashErr != nil {
				return nil, nil, fmt.Errorf("derive meta %d: %w", i, hashErr)
			}
			rawMask, hashErr := dkghash.HashFieldElements(meta, shared.X, shared.Y)
			if hashErr != nil {
				return nil, nil, fmt.Errorf("derive raw mask %d: %w", i, hashErr)
			}
			maskQuotient := new(big.Int).Sub(rawMask, shareMask)
			maskQuotient.Div(maskQuotient, group.ScalarField())
			shareSum := new(big.Int).Add(shares[i], shareMask)
			maskedShareCarry := big.NewInt(0)
			if shareSum.Cmp(group.ScalarField()) >= 0 {
				maskedShareCarry.SetInt64(1)
			}
			encryptedShares = append(encryptedShares, types.EncryptedShare{
				Recipient:      a.RecipientKeys[i].Operator,
				RecipientIndex: a.RecipientIndexes[i],
				Ephemeral:      ciphertext.Ephemeral,
				Ciphertext:     ciphertext.MaskedShare,
			})
			shareMasks = append(shareMasks, shareMask)
			maskQuotients = append(maskQuotients, maskQuotient)
			maskedShareCarries = append(maskedShareCarries, maskedShareCarry)
		}
	}
	ephemerals := make([]types.CurvePoint, len(encryptedShares))
	maskedShares := make([]*big.Int, len(encryptedShares))
	for i, encryptedShare := range encryptedShares {
		ephemerals[i] = encryptedShare.Ephemeral
		maskedShares[i] = encryptedShare.Ciphertext
	}
	paddedEphemerals, err := ccommon.CircuitPoints(ephemerals, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	maskedShares, err = ccommon.PadBigInts(maskedShares, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	shareMasks, err = ccommon.PadBigInts(shareMasks, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	maskQuotients, err = ccommon.PadBigInts(maskQuotients, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	maskedShareCarries, err = ccommon.PadBigInts(maskedShareCarries, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}

	threshold := big.NewInt(int64(a.Threshold))
	committeeSize := big.NewInt(int64(a.CommitteeSize))
	contributorIndex := big.NewInt(int64(a.ContributorIndex))

	commitmentInputs := []*big.Int{a.RoundHash, contributorIndex, threshold}
	for _, commitment := range paddedCommitments {
		commitmentInputs = append(
			commitmentInputs,
			ephemeralCoordinate([]twistededwards.Point{commitment}, 0, true),
			ephemeralCoordinate([]twistededwards.Point{commitment}, 0, false),
		)
	}
	commitmentHash, err := ccommon.MultiHashNative(commitmentInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("hash commitment inputs: %w", err)
	}

	shareInputs := []*big.Int{a.RoundHash, contributorIndex, committeeSize}
	for i := range MaxRecipients {
		shareInputs = append(
			shareInputs,
			recipientIndexes[i],
			ephemeralCoordinate(paddedEphemerals, i, true),
			ephemeralCoordinate(paddedEphemerals, i, false),
			maskedShares[i],
		)
	}
	shareHash, err := ccommon.MultiHashNative(shareInputs...)
	if err != nil {
		return nil, nil, fmt.Errorf("hash share inputs: %w", err)
	}
	anchor, err := ccommon.HashPackedBigIntsNative(commitmentHash, shareHash)
	if err != nil {
		return nil, nil, fmt.Errorf("hash contribution challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, contributionTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive contribution challenge: %w", err)
	}
	transcriptValues := make([]*big.Int, 0, 64)
	for i := range MaxCoefficients {
		transcriptValues = append(
			transcriptValues,
			ephemeralCoordinate(paddedCommitments, i, true),
			ephemeralCoordinate(paddedCommitments, i, false),
		)
	}
	transcriptValues = append(transcriptValues, recipientIndexes...)
	for i := range MaxRecipients {
		transcriptValues = append(
			transcriptValues,
			ephemeralCoordinate(paddedRecipientKeys, i, true),
			ephemeralCoordinate(paddedRecipientKeys, i, false),
		)
	}
	for i := range MaxRecipients {
		transcriptValues = append(
			transcriptValues,
			ephemeralCoordinate(paddedEphemerals, i, true),
			ephemeralCoordinate(paddedEphemerals, i, false),
		)
	}
	transcriptValues = append(transcriptValues, maskedShares...)
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc contribution transcript: %w", err)
	}

	witness := &ContributionCircuit{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            threshold,
		CommitteeSize:        committeeSize,
		ContributorIndex:     contributorIndex,
		CommitmentHash:       commitmentHash,
		ShareHash:            shareHash,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
		MaskedShares:         toWitnessScalars(maskedShares),
	}
	for i := range MaxCoefficients {
		witness.Coefficients[i] = coefficients[i]
		witness.Commitments[i] = paddedCommitments[i]
	}
	for i := range MaxRecipients {
		witness.RecipientPubKeys[i] = paddedRecipientKeys[i]
		witness.Ephemerals[i] = paddedEphemerals[i]
		witness.EncryptionNonces[i] = encryptionNonces[i]
		witness.RecipientIndexes[i] = recipientIndexes[i]
		witness.Shares[i] = shares[i]
		witness.MaskQuotients[i] = maskQuotients[i]
		witness.ShareMasks[i] = shareMasks[i]
		witness.MaskedShareCarries[i] = maskedShareCarries[i]
	}

	publicInputs := &PublicInputs{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            new(big.Int).Set(threshold),
		CommitteeSize:        new(big.Int).Set(committeeSize),
		ContributorIndex:     new(big.Int).Set(contributorIndex),
		CommitmentHash:       new(big.Int).Set(commitmentHash),
		ShareHash:            new(big.Int).Set(shareHash),
		Challenge:            new(big.Int).Set(challenge),
		TranscriptCommitment: new(big.Int).Set(transcriptCommitment),
		Commitments:          commitments,
		RecipientKeys:        recipientKeys,
		Shares:               shares[:len(a.RecipientIndexes)],
		EncryptedShares:      encryptedShares,
		RecipientIndexes:     recipientIndexes,
	}
	return witness, publicInputs, nil
}

// PublicWitness converts the native public inputs into the circuit's public assignment.
func (p PublicInputs) PublicWitness() *ContributionCircuit {
	return &ContributionCircuit{
		RoundHash:            p.RoundHash,
		Threshold:            p.Threshold,
		CommitteeSize:        p.CommitteeSize,
		ContributorIndex:     p.ContributorIndex,
		CommitmentHash:       p.CommitmentHash,
		ShareHash:            p.ShareHash,
		Challenge:            p.Challenge,
		TranscriptCommitment: p.TranscriptCommitment,
	}
}

// Scalars returns the ordered public scalars used by the verifier.
func (p PublicInputs) Scalars() []*big.Int {
	scalars := []*big.Int{
		p.RoundHash,
		p.Threshold,
		p.CommitteeSize,
		p.ContributorIndex,
		p.CommitmentHash,
		p.ShareHash,
		p.Challenge,
		p.TranscriptCommitment,
	}
	return scalars
}

// TranscriptScalars returns the ordered transcript compressed by the verifier path.
func (p PublicInputs) TranscriptScalars() []*big.Int {
	transcript := make([]*big.Int, 0, 64)
	for _, commitment := range publicCommitmentPoints(p.Commitments) {
		transcript = append(
			transcript,
			ephemeralCoordinate([]twistededwards.Point{commitment}, 0, true),
			ephemeralCoordinate([]twistededwards.Point{commitment}, 0, false),
		)
	}
	indexes, _ := ccommon.PadBigInts(p.RecipientIndexes, MaxRecipients)
	transcript = append(transcript, indexes...)
	for _, recipient := range publicRecipientPoints(p.RecipientKeys) {
		transcript = append(
			transcript,
			ephemeralCoordinate([]twistededwards.Point{recipient}, 0, true),
			ephemeralCoordinate([]twistededwards.Point{recipient}, 0, false),
		)
	}
	for _, ephemeral := range publicEphemeralPoints(p.EncryptedShares) {
		transcript = append(
			transcript,
			ephemeralCoordinate([]twistededwards.Point{ephemeral}, 0, true),
			ephemeralCoordinate([]twistededwards.Point{ephemeral}, 0, false),
		)
	}
	for _, maskedShare := range publicMaskedShares(p.EncryptedShares) {
		value, ok := maskedShare.(*big.Int)
		if !ok {
			value = big.NewInt(0)
		}
		transcript = append(transcript, value)
	}
	return transcript
}

// BRLCCommitment compresses the contribution transcript into one scalar commitment.
func (p PublicInputs) BRLCCommitment(challenge *big.Int) (*big.Int, error) {
	return ccommon.BRLCNative(challenge, p.TranscriptScalars()...)
}

func toWitnessScalars(values []*big.Int) [MaxRecipients]frontend.Variable {
	var out [MaxRecipients]frontend.Variable
	for i := range MaxRecipients {
		out[i] = values[i]
	}
	return out
}

func publicRecipientPoints(points []types.CurvePoint) [MaxRecipients]twistededwards.Point {
	var out [MaxRecipients]twistededwards.Point
	for i := range MaxRecipients {
		out[i] = ccommon.IdentityPoint()
		if i < len(points) {
			out[i] = ccommon.CircuitPoint(points[i])
		}
	}
	return out
}

func ephemeralCoordinate(points []twistededwards.Point, index int, x bool) *big.Int {
	value, ok := points[index].X.(*big.Int)
	if !x {
		value, ok = points[index].Y.(*big.Int)
	}
	if ok {
		return value
	}
	if x {
		return big.NewInt(0)
	}
	return big.NewInt(1)
}

func publicCommitmentPoints(points []types.CurvePoint) [MaxCoefficients]twistededwards.Point {
	var out [MaxCoefficients]twistededwards.Point
	for i := range MaxCoefficients {
		out[i] = ccommon.IdentityPoint()
		if i < len(points) {
			out[i] = ccommon.CircuitPoint(points[i])
		}
	}
	return out
}

func publicEphemeralPoints(shares []types.EncryptedShare) [MaxRecipients]twistededwards.Point {
	var out [MaxRecipients]twistededwards.Point
	for i := range MaxRecipients {
		out[i] = ccommon.IdentityPoint()
		if i < len(shares) {
			out[i] = ccommon.CircuitPoint(shares[i].Ephemeral)
		}
	}
	return out
}

func publicMaskedShares(shares []types.EncryptedShare) [MaxRecipients]frontend.Variable {
	var out [MaxRecipients]frontend.Variable
	for i := range MaxRecipients {
		out[i] = big.NewInt(0)
		if i < len(shares) && shares[i].Ciphertext != nil {
			out[i] = shares[i].Ciphertext
		}
	}
	return out
}
