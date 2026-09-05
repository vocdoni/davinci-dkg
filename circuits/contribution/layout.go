package contribution

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/types"
)

// Layout is the compact contribution transcript of one epoch policy
// (docs/pool-keys-v4.md §3, §5). The transcript carries no padding: its
// length L_C = MaxK·(2t+n) + 5n and every offset are functions of the
// threshold t and the committee size n, which the contract binds to the
// epoch before it streams the words, and which recipients and finalizers
// take from authoritative epoch state — never from the calldata itself.
//
//	[0, 2Kt)          for j, for m < t: A[j][m].x, A[j][m].y
//	[2Kt, 2Kt+n)      recipient indexes, committee order (slot i holds i+1)
//	[2Kt+n, 2Kt+3n)   recipient public keys (x, y)
//	[2Kt+3n, 2Kt+5n)  ephemerals (x, y)
//	[2Kt+5n, L_C)     for j, for i < n: masked share ms[j][i]
//
// The committee region [2Kt, 2Kt+3n) is what DKGManager compares against its
// committee snapshot.
type Layout struct {
	Threshold     int
	CommitteeSize int
}

// NewLayout validates 1 ≤ t ≤ n ≤ MaxN and returns the layout.
func NewLayout(threshold, committeeSize int) (Layout, error) {
	if threshold < 1 {
		return Layout{}, fmt.Errorf("threshold must be at least 1, got %d", threshold)
	}
	if committeeSize > MaxRecipients {
		return Layout{}, fmt.Errorf("committee size %d exceeds max %d", committeeSize, MaxRecipients)
	}
	if threshold > committeeSize {
		return Layout{}, fmt.Errorf("threshold %d exceeds committee size %d", threshold, committeeSize)
	}
	return Layout{Threshold: threshold, CommitteeSize: committeeSize}, nil
}

// Words is L_C, the transcript length in 32-byte words.
func (l Layout) Words() int {
	return ccommon.CompactContributionWords(l.Threshold, l.CommitteeSize)
}

// Bytes is the exact calldata length the contract requires.
func (l Layout) Bytes() int { return 32 * l.Words() }

// CommitmentOffset is the word of A[key][coefficient].x; .y follows.
func (l Layout) CommitmentOffset(key, coefficient int) int {
	return 2 * (key*l.Threshold + coefficient)
}

// RecipientIndexesStart is 2Kt, the first recipient-index word.
func (l Layout) RecipientIndexesStart() int { return 2 * MaxKeys * l.Threshold }

// RecipientIndexOffset is the word holding recipient slot i's index.
func (l Layout) RecipientIndexOffset(recipient int) int {
	return l.RecipientIndexesStart() + recipient
}

// RecipientKeysStart is 2Kt+n, the first recipient public-key word.
func (l Layout) RecipientKeysStart() int { return l.RecipientIndexesStart() + l.CommitteeSize }

// RecipientKeyOffset is the word of recipient slot i's public key x; y follows.
func (l Layout) RecipientKeyOffset(recipient int) int {
	return l.RecipientKeysStart() + 2*recipient
}

// EphemeralsStart is 2Kt+3n, the first ephemeral word.
func (l Layout) EphemeralsStart() int { return l.RecipientIndexesStart() + 3*l.CommitteeSize }

// EphemeralOffset is the word of recipient slot i's ephemeral x; y follows.
func (l Layout) EphemeralOffset(recipient int) int {
	return l.EphemeralsStart() + 2*recipient
}

// MaskedSharesStart is 2Kt+5n, the first masked-share word.
func (l Layout) MaskedSharesStart() int { return l.RecipientIndexesStart() + 5*l.CommitteeSize }

// MaskedShareOffset is the word of ms[key][recipient].
func (l Layout) MaskedShareOffset(key, recipient int) int {
	return l.MaskedSharesStart() + key*l.CommitteeSize + recipient
}

// CommitteeRegion is the half-open word interval [2Kt, 2Kt+3n) the contract
// compares against its committee snapshot: the indexes followed by the
// public keys, exactly 3n words.
func (l Layout) CommitteeRegion() (start, end int) {
	return l.RecipientIndexesStart(), l.EphemeralsStart()
}

// Transcript is the structured content of one compact transcript: the values
// a recipient or finalizer reads back from calldata, with no padding.
type Transcript struct {
	// Commitments[j] holds the t coefficient commitments of key j.
	Commitments [][]types.CurvePoint
	// RecipientIndexes holds the n committee indexes; slot i is i+1.
	RecipientIndexes []*big.Int
	// RecipientKeys and Ephemerals hold one point per committee slot.
	RecipientKeys []types.CurvePoint
	Ephemerals    []types.CurvePoint
	// MaskedShares[j] holds key j's masked share for each committee slot.
	MaskedShares [][]*big.Int
}

// Encode lays a transcript out as L_C words, checking that every vector has
// the shape the layout expects.
func (l Layout) Encode(t Transcript) ([]*big.Int, error) {
	if len(t.Commitments) != MaxKeys {
		return nil, fmt.Errorf("compact transcript: got %d commitment sets, expected %d", len(t.Commitments), MaxKeys)
	}
	if len(t.MaskedShares) != MaxKeys {
		return nil, fmt.Errorf("compact transcript: got %d masked share sets, expected %d", len(t.MaskedShares), MaxKeys)
	}
	n := l.CommitteeSize
	if len(t.RecipientIndexes) != n || len(t.RecipientKeys) != n || len(t.Ephemerals) != n {
		return nil, fmt.Errorf(
			"compact transcript: got %d indexes, %d keys and %d ephemerals for a committee of %d",
			len(t.RecipientIndexes), len(t.RecipientKeys), len(t.Ephemerals), n,
		)
	}
	words := make([]*big.Int, 0, l.Words())
	for j := range MaxKeys {
		if len(t.Commitments[j]) != l.Threshold {
			return nil, fmt.Errorf(
				"compact transcript: key %d has %d commitments, expected the threshold %d",
				j, len(t.Commitments[j]), l.Threshold,
			)
		}
		for m, point := range t.Commitments[j] {
			if err := point.Validate(); err != nil {
				return nil, fmt.Errorf("compact transcript: key %d commitment %d: %w", j, m, err)
			}
			words = append(words, point.X, point.Y)
		}
	}
	for i, index := range t.RecipientIndexes {
		if index == nil {
			return nil, fmt.Errorf("compact transcript: recipient index %d is nil", i)
		}
		words = append(words, index)
	}
	for i, point := range t.RecipientKeys {
		if err := point.Validate(); err != nil {
			return nil, fmt.Errorf("compact transcript: recipient key %d: %w", i, err)
		}
		words = append(words, point.X, point.Y)
	}
	for i, point := range t.Ephemerals {
		if err := point.Validate(); err != nil {
			return nil, fmt.Errorf("compact transcript: ephemeral %d: %w", i, err)
		}
		words = append(words, point.X, point.Y)
	}
	for j := range MaxKeys {
		if len(t.MaskedShares[j]) != n {
			return nil, fmt.Errorf(
				"compact transcript: key %d has %d masked shares for a committee of %d", j, len(t.MaskedShares[j]), n,
			)
		}
		for i, share := range t.MaskedShares[j] {
			if share == nil {
				return nil, fmt.Errorf("compact transcript: key %d masked share %d is nil", j, i)
			}
			words = append(words, share)
		}
	}
	return words, nil
}

// Decode parses L_C words into their regions. It rejects — never reduces —
// a non-canonical word (one outside [0, p)), a wrong length and a recipient
// slot whose index is not its committee position, so a transcript has
// exactly one encoding and callers can index the result by committee slot.
func (l Layout) Decode(words []*big.Int) (*Transcript, error) {
	if len(words) != l.Words() {
		return nil, fmt.Errorf("compact transcript: got %d words, expected %d for t=%d n=%d",
			len(words), l.Words(), l.Threshold, l.CommitteeSize)
	}
	modulus := ecc.BN254.ScalarField()
	for q, word := range words {
		if word == nil || word.Sign() < 0 || word.Cmp(modulus) >= 0 {
			return nil, fmt.Errorf("compact transcript: word %d is not a canonical field element", q)
		}
	}
	n := l.CommitteeSize
	out := &Transcript{
		Commitments:      make([][]types.CurvePoint, MaxKeys),
		RecipientIndexes: make([]*big.Int, n),
		RecipientKeys:    make([]types.CurvePoint, n),
		Ephemerals:       make([]types.CurvePoint, n),
		MaskedShares:     make([][]*big.Int, MaxKeys),
	}
	for j := range MaxKeys {
		out.Commitments[j] = make([]types.CurvePoint, l.Threshold)
		for m := range l.Threshold {
			q := l.CommitmentOffset(j, m)
			out.Commitments[j][m] = types.CurvePoint{X: words[q], Y: words[q+1]}
		}
	}
	for i := range n {
		index := words[l.RecipientIndexOffset(i)]
		if !index.IsInt64() || index.Int64() != int64(i+1) {
			return nil, fmt.Errorf("compact transcript: recipient slot %d carries index %s, expected %d", i, index, i+1)
		}
		out.RecipientIndexes[i] = index
		q := l.RecipientKeyOffset(i)
		out.RecipientKeys[i] = types.CurvePoint{X: words[q], Y: words[q+1]}
		q = l.EphemeralOffset(i)
		out.Ephemerals[i] = types.CurvePoint{X: words[q], Y: words[q+1]}
	}
	for j := range MaxKeys {
		out.MaskedShares[j] = make([]*big.Int, n)
		for i := range n {
			out.MaskedShares[j][i] = words[l.MaskedShareOffset(j, i)]
		}
	}
	return out, nil
}

// DecodeBytes decodes the raw `bytes transcript` calldata argument: exactly
// 32·L_C bytes of big-endian words.
func (l Layout) DecodeBytes(data []byte) (*Transcript, error) {
	if len(data) != l.Bytes() {
		return nil, fmt.Errorf("compact transcript: got %d bytes, expected %d for t=%d n=%d",
			len(data), l.Bytes(), l.Threshold, l.CommitteeSize)
	}
	words := make([]*big.Int, l.Words())
	for q := range words {
		words[q] = new(big.Int).SetBytes(data[32*q : 32*(q+1)])
	}
	return l.Decode(words)
}
