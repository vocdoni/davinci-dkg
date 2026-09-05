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

// TranscriptDomain is the BRLC Fiat–Shamir domain of the compact contribution
// transcript (docs/pool-keys-v4.md §2).
var TranscriptDomain = protocol.DomainContributionTranscriptV2

// PublicInputs is the native representation of the public contribution inputs
// plus the unpadded vectors the compact transcript is built from. The per-key
// vectors (Commitments, Shares, EncryptedShares) are indexed by pool key first
// and by coefficient / recipient second; Commitments[j] has Threshold entries,
// the recipient vectors have CommitteeSize entries.
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
	n := int(a.CommitteeSize)

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

	// Digests absorb the padded vectors (identity / zero in inactive slots),
	// exactly as the circuit does; only the calldata transcript is compact.
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

	publicShares := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		publicShares[j] = shares[j][:n]
	}
	publicInputs := &PublicInputs{
		RoundHash:        new(big.Int).Set(a.RoundHash),
		Threshold:        new(big.Int).Set(threshold),
		CommitteeSize:    new(big.Int).Set(committeeSize),
		ContributorIndex: new(big.Int).Set(contributorIndex),
		CommitmentHash:   new(big.Int).Set(commitmentHash),
		ShareHash:        new(big.Int).Set(shareHash),
		Commitments:      commitments,
		RecipientKeys:    recipientKeys,
		Shares:           publicShares,
		EncryptedShares:  encryptedShares,
		RecipientIndexes: recipientIndexes[:n],
	}

	transcriptValues, err := publicInputs.TranscriptScalars()
	if err != nil {
		return nil, nil, err
	}
	anchor, err := ccommon.ChallengeAnchor(transcriptValues, commitmentHash, shareHash)
	if err != nil {
		return nil, nil, fmt.Errorf("hash contribution challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, TranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive contribution challenge: %w", err)
	}
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc contribution transcript: %w", err)
	}
	publicInputs.Challenge = new(big.Int).Set(challenge)
	publicInputs.TranscriptCommitment = new(big.Int).Set(transcriptCommitment)

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

// Layout returns the compact transcript layout of these inputs' (t, n).
func (p PublicInputs) Layout() (Layout, error) {
	if p.Threshold == nil || p.CommitteeSize == nil || !p.Threshold.IsInt64() || !p.CommitteeSize.IsInt64() {
		return Layout{}, fmt.Errorf("contribution transcript: threshold and committee size are required")
	}
	return NewLayout(int(p.Threshold.Int64()), int(p.CommitteeSize.Int64()))
}

// Transcript assembles the structured compact transcript from the unpadded
// public vectors. The ephemeral is shared by every key, so key 0's encrypted
// shares carry all of them.
func (p PublicInputs) Transcript() (Transcript, error) {
	if len(p.EncryptedShares) != MaxKeys {
		return Transcript{}, fmt.Errorf(
			"contribution transcript: got %d encrypted share sets, expected %d", len(p.EncryptedShares), MaxKeys,
		)
	}
	maskedShares := make([][]*big.Int, MaxKeys)
	for j := range MaxKeys {
		maskedShares[j] = make([]*big.Int, len(p.EncryptedShares[j]))
		for i, share := range p.EncryptedShares[j] {
			maskedShares[j][i] = share.Ciphertext
		}
	}
	ephemerals := make([]types.CurvePoint, len(p.EncryptedShares[0]))
	for i, share := range p.EncryptedShares[0] {
		ephemerals[i] = share.Ephemeral
	}
	return Transcript{
		Commitments:      p.Commitments,
		RecipientIndexes: p.RecipientIndexes,
		RecipientKeys:    p.RecipientKeys,
		Ephemerals:       ephemerals,
		MaskedShares:     maskedShares,
	}, nil
}

// TranscriptScalars returns the L_C-word compact transcript the contract
// streams into its BRLC check and hashes into the challenge anchor.
func (p PublicInputs) TranscriptScalars() ([]*big.Int, error) {
	layout, err := p.Layout()
	if err != nil {
		return nil, err
	}
	transcript, err := p.Transcript()
	if err != nil {
		return nil, err
	}
	return layout.Encode(transcript)
}

// BRLCCommitment compresses the compact transcript into one scalar commitment.
func (p PublicInputs) BRLCCommitment(challenge *big.Int) (*big.Int, error) {
	scalars, err := p.TranscriptScalars()
	if err != nil {
		return nil, err
	}
	return ccommon.BRLCNative(challenge, scalars...)
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
