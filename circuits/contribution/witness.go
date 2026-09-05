package contribution

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	dkghash "github.com/vocdoni/davinci-dkg/crypto/hash"
	"github.com/vocdoni/davinci-dkg/crypto/shareenc"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/types"
)

var contributionTranscriptDomain = protocol.DomainContributionTranscriptV1

// PublicInputs is the native representation of the public contribution inputs.
// The per-key vectors (Commitments, Shares, EncryptedShares) are indexed by
// pool key first and by coefficient / recipient second.
type PublicInputs struct {
	RoundHash            *big.Int
	Threshold            *big.Int
	CommitteeSize        *big.Int
	ContributorIndex     *big.Int
	CommitmentHash       *big.Int
	ShareHash            *big.Int
	Challenge            *big.Int
	TranscriptCommitment *big.Int
	Commitments          [][]types.CurvePoint
	RecipientKeys        []types.CurvePoint
	Shares               [][]*big.Int
	EncryptedShares      [][]types.EncryptedShare
	RecipientIndexes     []*big.Int
}

// BuildWitness materializes the native assignment into a gnark witness and the
// corresponding public inputs.
func BuildWitness(a Assignment) (*ContributionCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	threshold := big.NewInt(int64(a.Threshold))
	committeeSize := big.NewInt(int64(a.CommitteeSize))
	contributorIndex := big.NewInt(int64(a.ContributorIndex))
	subgroupOrder := group.ScalarField()

	recipientIndexes, err := ccommon.PadBigInts(ccommon.Uint16sToBigInts(a.RecipientIndexes), MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	encryptionNonces, err := ccommon.PadBigInts(a.EncryptionNonces, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	recipientKeys := make([]types.CurvePoint, len(a.RecipientKeys))
	for i, key := range a.RecipientKeys {
		recipientKeys[i] = types.CurvePoint{X: key.PubX, Y: key.PubY}
	}
	paddedRecipientKeys, err := ccommon.PadPoints(recipientKeys, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}

	// Per pool key: the padded coefficient vector, its Feldman commitments
	// (identity beyond the threshold) and the share of every recipient.
	coefficients := make([][]*big.Int, MaxKeys)
	commitments := make([][]types.CurvePoint, MaxKeys)
	paddedCommitments := make([][]types.CurvePoint, MaxKeys)
	shares := newScalarGrid()
	for j := range MaxKeys {
		coefficients[j], err = ccommon.PadBigInts(a.Coefficients[j], MaxCoefficients)
		if err != nil {
			return nil, nil, fmt.Errorf("pad key %d coefficients: %w", j, err)
		}
		commitments[j] = make([]types.CurvePoint, len(a.Coefficients[j]))
		for m, coefficient := range a.Coefficients[j] {
			point := group.NewPoint()
			point.ScalarBaseMult(coefficient)
			commitments[j][m] = group.Encode(point)
		}
		paddedCommitments[j], err = ccommon.PadPoints(commitments[j], MaxCoefficients)
		if err != nil {
			return nil, nil, fmt.Errorf("pad key %d commitments: %w", j, err)
		}
		for i := range a.RecipientIndexes {
			share, evalErr := ccommon.EvaluatePolynomialNative(a.Coefficients[j], big.NewInt(int64(a.RecipientIndexes[i])))
			if evalErr != nil {
				return nil, nil, fmt.Errorf("evaluate key %d share %d: %w", j, i, evalErr)
			}
			shares[j][i] = share
		}
	}

	ephemerals, err := ccommon.PadPoints(nil, MaxRecipients)
	if err != nil {
		return nil, nil, err
	}
	maskedShares := newScalarGrid()
	shareMasks := newScalarGrid()
	maskQuotients := newScalarGrid()
	maskedShareCarries := newScalarGrid()
	encryptedShares := make([][]types.EncryptedShare, MaxKeys)
	for j := range MaxKeys {
		encryptedShares[j] = make([]types.EncryptedShare, 0, len(a.RecipientIndexes))
	}
	for i := range a.RecipientIndexes {
		recipientPoint, decodeErr := group.Decode(paddedRecipientKeys[i])
		if decodeErr != nil {
			return nil, nil, fmt.Errorf("decode recipient key %d: %w", i, decodeErr)
		}
		// One ECDH secret per recipient, reused by every pool key; the key
		// index is what keeps the MaxK masks derived from it independent.
		sharedPoint := group.NewPoint()
		sharedPoint.ScalarMult(recipientPoint, a.EncryptionNonces[i])
		shared := group.Encode(sharedPoint)
		for j := range MaxKeys {
			ciphertext, encryptErr := shareenc.EncryptShareWithNonceRoundHash(
				a.RoundHash,
				a.ContributorIndex,
				a.RecipientIndexes[i],
				uint8(j),
				shares[j][i],
				a.RecipientKeys[i],
				a.EncryptionNonces[i],
			)
			if encryptErr != nil {
				return nil, nil, fmt.Errorf("encrypt key %d share %d: %w", j, i, encryptErr)
			}
			shareMask := new(big.Int).Sub(ciphertext.MaskedShare, shares[j][i])
			shareMask.Mod(shareMask, subgroupOrder)
			rawMask, maskErr := rawShareMask(a.RoundHash, a.ContributorIndex, a.RecipientIndexes[i], uint8(j), shared)
			if maskErr != nil {
				return nil, nil, fmt.Errorf("derive key %d raw mask %d: %w", j, i, maskErr)
			}
			maskQuotient := new(big.Int).Sub(rawMask, shareMask)
			maskQuotient.Div(maskQuotient, subgroupOrder)
			carry := big.NewInt(0)
			if new(big.Int).Add(shares[j][i], shareMask).Cmp(subgroupOrder) >= 0 {
				carry.SetInt64(1)
			}

			ephemerals[i] = ciphertext.Ephemeral
			maskedShares[j][i] = ciphertext.MaskedShare
			shareMasks[j][i] = shareMask
			maskQuotients[j][i] = maskQuotient
			maskedShareCarries[j][i] = carry
			encryptedShares[j] = append(encryptedShares[j], types.EncryptedShare{
				Recipient:      a.RecipientKeys[i].Operator,
				RecipientIndex: a.RecipientIndexes[i],
				Ephemeral:      ciphertext.Ephemeral,
				Ciphertext:     ciphertext.MaskedShare,
			})
		}
	}

	keyDigests := make([]*big.Int, MaxKeys)
	for j := range MaxKeys {
		keyDigests[j], err = ccommon.CommitmentKeyDigestNative(paddedCommitments[j])
		if err != nil {
			return nil, nil, fmt.Errorf("digest key %d commitments: %w", j, err)
		}
	}
	commitmentHash, err := ccommon.CommitmentsHashNative(a.RoundHash, contributorIndex, threshold, keyDigests)
	if err != nil {
		return nil, nil, fmt.Errorf("hash commitment inputs: %w", err)
	}

	rowDigests := make([]*big.Int, MaxRecipients)
	for i := range MaxRecipients {
		rowShares := make([]*big.Int, MaxKeys)
		for j := range MaxKeys {
			rowShares[j] = maskedShares[j][i]
		}
		rowDigests[i], err = ccommon.EncryptedShareRowDigestNative(
			recipientIndexes[i],
			paddedRecipientKeys[i],
			ephemerals[i],
			rowShares,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("share row %d digest: %w", i, err)
		}
	}
	shareHash, err := ccommon.EncryptedSharesHashNative(a.RoundHash, contributorIndex, committeeSize, rowDigests)
	if err != nil {
		return nil, nil, fmt.Errorf("hash share inputs: %w", err)
	}

	transcriptValues := transcriptWords(paddedCommitments, recipientIndexes, paddedRecipientKeys, ephemerals, maskedShares)
	anchor, err := ccommon.ChallengeAnchor(transcriptValues, commitmentHash, shareHash)
	if err != nil {
		return nil, nil, fmt.Errorf("hash contribution challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, contributionTranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive contribution challenge: %w", err)
	}
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
	}
	for i := range MaxRecipients {
		witness.RecipientPubKeys[i] = ccommon.CircuitPoint(paddedRecipientKeys[i])
		witness.Ephemerals[i] = ccommon.CircuitPoint(ephemerals[i])
		witness.EncryptionNonces[i] = encryptionNonces[i]
		witness.RecipientIndexes[i] = recipientIndexes[i]
	}
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			witness.Coefficients[j][m] = coefficients[j][m]
			witness.Commitments[j][m] = ccommon.CircuitPoint(paddedCommitments[j][m])
		}
		for i := range MaxRecipients {
			witness.Shares[j][i] = shares[j][i]
			witness.MaskedShares[j][i] = maskedShares[j][i]
			witness.MaskQuotients[j][i] = maskQuotients[j][i]
			witness.ShareMasks[j][i] = shareMasks[j][i]
			witness.MaskedShareCarries[j][i] = maskedShareCarries[j][i]
		}
	}

	publicShares := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		publicShares[j] = shares[j][:len(a.RecipientIndexes)]
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
		Shares:               publicShares,
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
// Order must match the circuit field declaration order.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Threshold,
		p.CommitteeSize,
		p.ContributorIndex,
		p.CommitmentHash,
		p.ShareHash,
		p.Challenge,
		p.TranscriptCommitment,
	}
}

// TranscriptScalars returns the ordered transcript compressed by the verifier
// path: TranscriptWords words, commitments and masked shares key-major.
//
// The padding helpers can fail when a manually-constructed PublicInputs holds
// more entries than the circuit bounds; BuildWitness validates its assignment
// first, but a hand-built value would otherwise emit a malformed transcript.
func (p PublicInputs) TranscriptScalars() ([]*big.Int, error) {
	if len(p.Commitments) != MaxKeys {
		return nil, fmt.Errorf("contribution transcript: got %d commitment sets, expected %d", len(p.Commitments), MaxKeys)
	}
	if len(p.EncryptedShares) != MaxKeys {
		return nil, fmt.Errorf(
			"contribution transcript: got %d encrypted share sets, expected %d",
			len(p.EncryptedShares),
			MaxKeys,
		)
	}
	commitments := make([][]types.CurvePoint, MaxKeys)
	maskedShares := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		var err error
		if commitments[j], err = ccommon.PadPoints(p.Commitments[j], MaxCoefficients); err != nil {
			return nil, fmt.Errorf("contribution transcript: pad key %d commitments: %w", j, err)
		}
		values := make([]*big.Int, len(p.EncryptedShares[j]))
		for i, share := range p.EncryptedShares[j] {
			values[i] = share.Ciphertext
		}
		if maskedShares[j], err = ccommon.PadBigInts(values, MaxRecipients); err != nil {
			return nil, fmt.Errorf("contribution transcript: pad key %d masked shares: %w", j, err)
		}
	}
	indexes, err := ccommon.PadBigInts(p.RecipientIndexes, MaxRecipients)
	if err != nil {
		return nil, fmt.Errorf("contribution transcript: pad recipient indexes: %w", err)
	}
	recipientKeys, err := ccommon.PadPoints(p.RecipientKeys, MaxRecipients)
	if err != nil {
		return nil, fmt.Errorf("contribution transcript: pad recipient keys: %w", err)
	}
	// The ephemeral is shared by every key, so key 0's rows carry all of them.
	ephemerals := make([]types.CurvePoint, len(p.EncryptedShares[0]))
	for i, share := range p.EncryptedShares[0] {
		ephemerals[i] = share.Ephemeral
	}
	paddedEphemerals, err := ccommon.PadPoints(ephemerals, MaxRecipients)
	if err != nil {
		return nil, fmt.Errorf("contribution transcript: pad ephemerals: %w", err)
	}
	return transcriptWords(commitments, indexes, recipientKeys, paddedEphemerals, maskedShares), nil
}

// BRLCCommitment compresses the contribution transcript into one scalar commitment.
func (p PublicInputs) BRLCCommitment(challenge *big.Int) (*big.Int, error) {
	scalars, err := p.TranscriptScalars()
	if err != nil {
		return nil, err
	}
	return ccommon.BRLCNative(challenge, scalars...)
}

// transcriptWords lays out the calldata transcript both the circuit and the
// contract stream: commitments key-major, then the recipient indexes, their
// public keys, the shared ephemerals, and the masked shares key-major. Every
// argument must already be padded to the circuit bounds.
func transcriptWords(
	commitments [][]types.CurvePoint,
	recipientIndexes []*big.Int,
	recipientKeys, ephemerals []types.CurvePoint,
	maskedShares [][]*big.Int,
) []*big.Int {
	words := make([]*big.Int, 0, TranscriptWords)
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			words = append(words, commitments[j][m].X, commitments[j][m].Y)
		}
	}
	words = append(words, recipientIndexes...)
	for i := range MaxRecipients {
		words = append(words, recipientKeys[i].X, recipientKeys[i].Y)
	}
	for i := range MaxRecipients {
		words = append(words, ephemerals[i].X, ephemerals[i].Y)
	}
	for j := range MaxKeys {
		words = append(words, maskedShares[j]...)
	}
	return words
}

// rawShareMask recomputes ccommon.ShareMaskHash natively, before the
// subgroup-order reduction: the circuit takes the quotient of that reduction
// as a witness and shareenc only ever returns the reduced mask.
func rawShareMask(
	roundHash *big.Int,
	contributorIndex, recipientIndex uint16,
	keyIndex uint8,
	shared types.CurvePoint,
) (*big.Int, error) {
	meta, err := dkghash.HashFieldElements(
		ccommon.ShareEncryptionDomain(),
		roundHash,
		new(big.Int).SetUint64((uint64(contributorIndex)<<16)|uint64(recipientIndex)),
		new(big.Int).SetUint64(uint64(keyIndex)),
	)
	if err != nil {
		return nil, fmt.Errorf("hash mask metadata: %w", err)
	}
	rawMask, err := dkghash.HashFieldElements(meta, shared.X, shared.Y)
	if err != nil {
		return nil, fmt.Errorf("hash shared secret: %w", err)
	}
	return rawMask, nil
}

// newScalarGrid allocates a zeroed MaxKeys × MaxRecipients scalar grid.
func newScalarGrid() [][]*big.Int {
	grid := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		grid[j] = make([]*big.Int, MaxRecipients)
		for i := range MaxRecipients {
			grid[j][i] = big.NewInt(0)
		}
	}
	return grid
}
