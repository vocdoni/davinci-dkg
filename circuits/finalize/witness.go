package finalize

import (
	"fmt"
	"math/big"

	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
	"github.com/vocdoni/davinci-dkg/internal/protocol"
	"github.com/vocdoni/davinci-dkg/types"
)

// TranscriptDomain is the BRLC Fiat–Shamir domain of the finalization
// transcript (docs/pool-keys-v4.md §2).
var TranscriptDomain = protocol.DomainFinalizeTranscriptV2

// Transcript word offsets (docs/pool-keys-v4.md §7). The layout is fixed, so
// these are plain functions of the circuit bounds.

// IndexesStart is the first participant-index word.
const IndexesStart = 0

// HashesStart is the first contribution-hash word.
const HashesStart = MaxParticipants

// KeysStart is the first word of key 0's block.
const KeysStart = 2 * MaxParticipants

// KeyOffset is the first word of key j's block, which holds P[j].x.
func KeyOffset(key int) int { return KeysStart + key*KeyWords }

// PoolKeyOffset is the word of P[key].x; .y follows.
func PoolKeyOffset(key int) int { return KeyOffset(key) }

// ShareCommitmentOffset is the word of D[key][member].x, member zero-based
// (committee position member+1); .y follows.
func ShareCommitmentOffset(key, member int) int { return KeyOffset(key) + 2 + 2*member }

// PublicInputs is the native representation of the finalization public inputs
// plus the transcript vectors the contract streams alongside the proof. Every
// vector is padded to the circuit bounds: dealer rows beyond AcceptedCount
// carry a zero index and hash, ShareCommitments[j][i] is D[j][i] for every
// committee member i < CommitteeSize and the identity beyond. Aggregate
// coefficients above the constant term are not part of the transcript; they
// are exposed for callers that verify recovered shares against them.
type PublicInputs struct {
	RoundHash            *big.Int
	Threshold            *big.Int
	CommitteeSize        *big.Int
	AcceptedCount        *big.Int
	TranscriptDigest     *big.Int
	Challenge            *big.Int
	TranscriptCommitment *big.Int
	ParticipantIndexes   []*big.Int
	ContributionHashes   []*big.Int
	PoolKeys             []types.CurvePoint
	ShareCommitments     [][]types.CurvePoint
	AggregateCommitments [][]types.CurvePoint
}

// BuildWitness materializes the native assignment into a gnark witness and the
// corresponding public inputs.
func BuildWitness(a Assignment) (*FinalizeCircuit, *PublicInputs, error) {
	if err := a.Validate(); err != nil {
		return nil, nil, err
	}

	accepted := len(a.ParticipantIndexes)
	threshold := big.NewInt(int64(a.Threshold))
	committeeSize := big.NewInt(int64(a.CommitteeSize))
	acceptedCount := big.NewInt(int64(accepted))

	participantIndexes, err := ccommon.PadBigInts(ccommon.Uint16sToBigInts(a.ParticipantIndexes), MaxParticipants)
	if err != nil {
		return nil, nil, err
	}

	// Per dealer: every key's padded commitment vector and the stored
	// commitments hash those vectors reproduce. Inactive rows stay at the
	// identity / zero, which the circuit masks out of every constraint.
	commitments := make([][][]types.CurvePoint, MaxParticipants)
	contributionHashes := make([]*big.Int, MaxParticipants)
	for d := range MaxParticipants {
		commitments[d] = make([][]types.CurvePoint, MaxKeys)
		for j := range MaxKeys {
			if commitments[d][j], err = ccommon.PadPoints(nil, MaxCoefficients); err != nil {
				return nil, nil, err
			}
		}
		contributionHashes[d] = big.NewInt(0)
	}
	for d := range accepted {
		keyDigests := make([]*big.Int, MaxKeys)
		for j := range MaxKeys {
			padded, padErr := ccommon.PadPoints(a.Commitments[d][j], MaxCoefficients)
			if padErr != nil {
				return nil, nil, fmt.Errorf("pad dealer %d key %d commitments: %w", d, j, padErr)
			}
			commitments[d][j] = padded
			keyDigests[j], err = ccommon.CommitmentKeyDigestNative(padded)
			if err != nil {
				return nil, nil, fmt.Errorf("digest dealer %d key %d: %w", d, j, err)
			}
		}
		contributionHashes[d], err = ccommon.CommitmentsHashNative(a.RoundHash, participantIndexes[d], threshold, keyDigests)
		if err != nil {
			return nil, nil, fmt.Errorf("hash dealer %d commitments: %w", d, err)
		}
		if len(a.ContributionHashes) != 0 && a.ContributionHashes[d].Cmp(contributionHashes[d]) != 0 {
			return nil, nil, fmt.Errorf(
				"dealer %d (member %d): recomputed commitments hash %s does not match the stored %s",
				d, a.ParticipantIndexes[d], contributionHashes[d], a.ContributionHashes[d],
			)
		}
	}

	// Per key: Ā[j][m] = Σ_d A[d][j][m] over the accepted dealers (identity
	// beyond the threshold on every row, so it aggregates to it too), then
	// D[j][i] for every committee member i+1 — the Merkle leaf of member i+1
	// is slot i — and the identity beyond the committee.
	aggregate := make([][]types.CurvePoint, MaxKeys)
	poolKeys := make([]types.CurvePoint, MaxKeys)
	shareCommitments := make([][]types.CurvePoint, MaxKeys)
	for j := range MaxKeys {
		aggregate[j] = make([]types.CurvePoint, MaxCoefficients)
		for m := range MaxCoefficients {
			sum := group.NewPoint()
			sum.SetZero()
			for d := range accepted {
				point, decodeErr := group.Decode(commitments[d][j][m])
				if decodeErr != nil {
					return nil, nil, fmt.Errorf("decode dealer %d key %d commitment %d: %w", d, j, m, decodeErr)
				}
				sum.Add(sum, point)
			}
			aggregate[j][m] = group.Encode(sum)
		}
		poolKeys[j] = aggregate[j][0]
		if shareCommitments[j], err = ccommon.PadPoints(nil, MaxParticipants); err != nil {
			return nil, nil, err
		}
		for i := range int(a.CommitteeSize) {
			shareCommitments[j][i], err = ccommon.CommitmentPolynomialValueNative(aggregate[j], big.NewInt(int64(i+1)))
			if err != nil {
				return nil, nil, fmt.Errorf("evaluate key %d share commitment of member %d: %w", j, i+1, err)
			}
		}
	}

	publicInputs := &PublicInputs{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            threshold,
		CommitteeSize:        committeeSize,
		AcceptedCount:        acceptedCount,
		ParticipantIndexes:   participantIndexes,
		ContributionHashes:   contributionHashes,
		PoolKeys:             poolKeys,
		ShareCommitments:     shareCommitments,
		AggregateCommitments: aggregate,
	}
	transcriptValues, err := publicInputs.TranscriptScalars()
	if err != nil {
		return nil, nil, err
	}
	digest, err := TranscriptDigestNative(a.RoundHash, threshold, committeeSize, acceptedCount, transcriptValues)
	if err != nil {
		return nil, nil, err
	}
	// The anchor commits to the witness transcript (through the Poseidon
	// digest the circuit recomputes) and to the calldata transcript (through
	// keccak) before ρ exists: keccak(digest ‖ keccak(transcript)).
	anchor, err := ccommon.ChallengeAnchor(transcriptValues, digest)
	if err != nil {
		return nil, nil, fmt.Errorf("hash finalize challenge anchor: %w", err)
	}
	challenge, err := ccommon.DeriveChallengeNative(a.RoundHash, TranscriptDomain, anchor)
	if err != nil {
		return nil, nil, fmt.Errorf("derive finalize challenge: %w", err)
	}
	transcriptCommitment, err := ccommon.BRLCNative(challenge, transcriptValues...)
	if err != nil {
		return nil, nil, fmt.Errorf("brlc finalize transcript: %w", err)
	}
	publicInputs.TranscriptDigest = digest
	publicInputs.Challenge = challenge
	publicInputs.TranscriptCommitment = transcriptCommitment

	witness := &FinalizeCircuit{
		RoundHash:            new(big.Int).Set(a.RoundHash),
		Threshold:            threshold,
		CommitteeSize:        committeeSize,
		AcceptedCount:        acceptedCount,
		TranscriptDigest:     digest,
		Challenge:            challenge,
		TranscriptCommitment: transcriptCommitment,
	}
	for d := range MaxParticipants {
		witness.ParticipantIndexes[d] = participantIndexes[d]
		witness.ContributionHashes[d] = contributionHashes[d]
		for j := range MaxKeys {
			for m := range MaxCoefficients {
				witness.Commitments[d][j][m] = ccommon.CircuitPoint(commitments[d][j][m])
			}
		}
	}
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			witness.AggregateCommitments[j][m] = ccommon.CircuitPoint(aggregate[j][m])
		}
		for i := range MaxParticipants {
			witness.ShareCommitments[j][i] = ccommon.CircuitPoint(shareCommitments[j][i])
		}
	}
	return witness, publicInputs, nil
}

// DigestParts are the three Poseidon levels of the transcript digest T
// (docs/pool-keys-v4.md §7), exposed so cross-implementation vectors can pin
// every intermediate value and not only the result.
type DigestParts struct {
	// Rows is R = H(0, I[0…N−1], h[0…N−1]).
	Rows *big.Int
	// Keys[j] is B_j = H(1, j, P[j].x, P[j].y, D[j][0].x, D[j][0].y, …).
	Keys []*big.Int
	// Digest is T = H(2, eid, t, n, a, K, L_F, R, B_0, …, B_(K−1)).
	Digest *big.Int
}

// TranscriptDigestParts computes R, every B_j and T over a TranscriptWords-long
// masked transcript in calldata order. The three calls absorb 1 + 2N, 4 + 2N
// and 8 + K inputs (65, 68 and 24 at the current bounds), under the native
// sponge's 256-input cap.
func TranscriptDigestParts(
	roundHash, threshold, committeeSize, acceptedCount *big.Int,
	transcript []*big.Int,
) (*DigestParts, error) {
	if roundHash == nil || threshold == nil || committeeSize == nil || acceptedCount == nil {
		return nil, fmt.Errorf("finalize transcript digest: epoch hash and counts are required")
	}
	if len(transcript) != TranscriptWords {
		return nil, fmt.Errorf("finalize transcript digest: got %d words, want %d", len(transcript), TranscriptWords)
	}
	for q, word := range transcript {
		if word == nil {
			return nil, fmt.Errorf("finalize transcript digest: word %d is nil", q)
		}
	}
	rowsInputs := make([]*big.Int, 0, 1+KeysStart)
	rowsInputs = append(rowsInputs, big.NewInt(digestTagRows))
	rowsInputs = append(rowsInputs, transcript[:KeysStart]...)
	rows, err := ccommon.MultiHashNative(rowsInputs...)
	if err != nil {
		return nil, fmt.Errorf("finalize transcript digest: rows: %w", err)
	}
	keys := make([]*big.Int, MaxKeys)
	for j := range MaxKeys {
		keyInputs := make([]*big.Int, 0, 2+KeyWords)
		keyInputs = append(keyInputs, big.NewInt(digestTagKey), big.NewInt(int64(j)))
		keyInputs = append(keyInputs, transcript[KeyOffset(j):KeyOffset(j+1)]...)
		if keys[j], err = ccommon.MultiHashNative(keyInputs...); err != nil {
			return nil, fmt.Errorf("finalize transcript digest: key %d: %w", j, err)
		}
	}
	outer := make([]*big.Int, 0, 8+MaxKeys)
	outer = append(outer,
		big.NewInt(digestTagOuter), roundHash, threshold, committeeSize, acceptedCount,
		big.NewInt(MaxKeys), big.NewInt(TranscriptWords), rows,
	)
	outer = append(outer, keys...)
	digest, err := ccommon.MultiHashNative(outer...)
	if err != nil {
		return nil, fmt.Errorf("finalize transcript digest: outer: %w", err)
	}
	return &DigestParts{Rows: rows, Keys: keys, Digest: digest}, nil
}

// TranscriptDigestNative mirrors the circuit's TranscriptDigest public input.
func TranscriptDigestNative(
	roundHash, threshold, committeeSize, acceptedCount *big.Int,
	transcript []*big.Int,
) (*big.Int, error) {
	parts, err := TranscriptDigestParts(roundHash, threshold, committeeSize, acceptedCount, transcript)
	if err != nil {
		return nil, err
	}
	return parts.Digest, nil
}

// PublicWitness converts native public inputs into the circuit public witness.
func (p PublicInputs) PublicWitness() *FinalizeCircuit {
	return &FinalizeCircuit{
		RoundHash:            p.RoundHash,
		Threshold:            p.Threshold,
		CommitteeSize:        p.CommitteeSize,
		AcceptedCount:        p.AcceptedCount,
		TranscriptDigest:     p.TranscriptDigest,
		Challenge:            p.Challenge,
		TranscriptCommitment: p.TranscriptCommitment,
	}
}

// Scalars returns the 7 ordered public scalars used by the verifier.
// Order must match the circuit field declaration order.
func (p PublicInputs) Scalars() []*big.Int {
	return []*big.Int{
		p.RoundHash,
		p.Threshold,
		p.CommitteeSize,
		p.AcceptedCount,
		p.TranscriptDigest,
		p.Challenge,
		p.TranscriptCommitment,
	}
}

// TranscriptScalars returns the TranscriptWords-long finalization transcript
// the contract streams into its BRLC check: the dealer indexes, the dealer
// hashes, then per key P[j] and the N share commitments.
func (p PublicInputs) TranscriptScalars() ([]*big.Int, error) {
	indexes, err := ccommon.PadBigInts(p.ParticipantIndexes, MaxParticipants)
	if err != nil {
		return nil, fmt.Errorf("finalize transcript: pad participant indexes: %w", err)
	}
	hashes, err := ccommon.PadBigInts(p.ContributionHashes, MaxParticipants)
	if err != nil {
		return nil, fmt.Errorf("finalize transcript: pad contribution hashes: %w", err)
	}
	if len(p.PoolKeys) != MaxKeys || len(p.ShareCommitments) != MaxKeys {
		return nil, fmt.Errorf(
			"finalize transcript: got %d pool keys and %d share commitment sets, expected %d",
			len(p.PoolKeys), len(p.ShareCommitments), MaxKeys,
		)
	}
	words := make([]*big.Int, 0, TranscriptWords)
	words = append(words, indexes...)
	words = append(words, hashes...)
	for j := range MaxKeys {
		if err := p.PoolKeys[j].Validate(); err != nil {
			return nil, fmt.Errorf("finalize transcript: pool key %d: %w", j, err)
		}
		shares, err := ccommon.PadPoints(p.ShareCommitments[j], MaxParticipants)
		if err != nil {
			return nil, fmt.Errorf("finalize transcript: pad key %d share commitments: %w", j, err)
		}
		words = append(words, p.PoolKeys[j].X, p.PoolKeys[j].Y)
		for i := range MaxParticipants {
			words = append(words, shares[i].X, shares[i].Y)
		}
	}
	return words, nil
}

// BRLCCommitment compresses the finalization transcript into one scalar
// commitment, matching the circuit's TranscriptCommitment public input and
// the contract's on-chain check.
func (p PublicInputs) BRLCCommitment(challenge *big.Int) (*big.Int, error) {
	scalars, err := p.TranscriptScalars()
	if err != nil {
		return nil, err
	}
	return ccommon.BRLCNative(challenge, scalars...)
}

// ShareLeaves returns key j's MaxN Merkle leaves: member p's commitment at
// slot p−1 for p ≤ CommitteeSize, EmptyLeaf beyond — what `finalizeEpoch`
// hashes into poolShareRoots[eid][j] and a partial later proves against.
func (p PublicInputs) ShareLeaves(key int) ([MaxParticipants][32]byte, error) {
	if key < 0 || key >= MaxKeys {
		return [MaxParticipants][32]byte{}, fmt.Errorf("key %d outside the pool [0, %d)", key, MaxKeys)
	}
	if p.CommitteeSize == nil || !p.CommitteeSize.IsInt64() || p.CommitteeSize.Int64() > MaxParticipants {
		return [MaxParticipants][32]byte{}, fmt.Errorf("committee size is missing or out of range")
	}
	n := int(p.CommitteeSize.Int64())
	if len(p.ShareCommitments) != MaxKeys || len(p.ShareCommitments[key]) < n {
		return [MaxParticipants][32]byte{}, fmt.Errorf("key %d has fewer share commitments than the committee", key)
	}
	members := make([]uint16, n)
	for i := range n {
		members[i] = uint16(i + 1)
	}
	return ccommon.ShareCommitmentLeaves(members, p.ShareCommitments[key][:n])
}

// ShareRoots returns the MaxK Merkle roots the contract stores at
// finalization, one per key.
func (p PublicInputs) ShareRoots() ([MaxKeys][32]byte, error) {
	var roots [MaxKeys][32]byte
	for j := range MaxKeys {
		leaves, err := p.ShareLeaves(j)
		if err != nil {
			return roots, err
		}
		roots[j] = ccommon.MerkleRoot(leaves)
	}
	return roots, nil
}
