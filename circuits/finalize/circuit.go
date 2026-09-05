// Package finalize holds the batched finalization circuit (docs/pool-keys-v4.md
// §6, §7): one proof, verified by `finalizeEpoch`, that derives every pool key
// of an epoch and every committee member's share commitment for each key from
// the accepted contributions. It replaces the proof-less finalize plus the
// per-key `circuits/poolkey` activation of v3.1, so `Live` means every key and
// every share-commitment root is stored and usable.
package finalize

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	"github.com/consensys/gnark/std/math/bits"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// MaxCoefficients/MaxParticipants are aliases of the single shared constant
// `circuits/common.MaxN`; MaxKeys is `circuits/common.MaxK`, the pool size.
// Edit `circuits/common/sizes.go` to change them.
const (
	MaxCoefficients = ccommon.MaxN
	MaxParticipants = ccommon.MaxN
	MaxKeys         = ccommon.MaxK
)

// TranscriptWords is the fixed word count L_F of the finalization calldata
// transcript, 2·MaxN + MaxK·(2 + 2·MaxN) (docs/pool-keys-v4.md §7):
//
//	[0, N)                    participant indexes I[d]      (0 for rows d ≥ a)
//	[N, 2N)                   contribution hashes h[d]      (0 for rows d ≥ a)
//	[2N + j·(2+2N), …)        key j: P[j].x, P[j].y, then D[j][0].x, D[j][0].y, …, D[j][N−1].x, D[j][N−1].y
//	                          (D[j][i] = (0,1) for i ≥ n)
const TranscriptWords = ccommon.FinalizeTranscriptWords

// KeyWords is the transcript span of one key: P[j] then N share commitments.
const KeyWords = 2 + 2*MaxParticipants

// PublicInputWords is the number of public inputs the verifier reads.
const PublicInputWords = 7

// Digest tags (§7): domain-separate the three Poseidon levels of the
// transcript digest.
const (
	digestTagRows  = 0
	digestTagKey   = 1
	digestTagOuter = 2
)

// FinalizeCircuit proves, for every pool key j at once: each accepted
// dealer's on-chain commitmentsHash is reproduced from its MaxK commitment
// vectors; Ā[j][m] = Σ_d A[d][j][m] over the accepted dealers; P[j] = Ā[j][0];
// and D[j][i] = Σ_m (i+1)^m · Ā[j][m] is member i+1's share commitment for
// every committee member, contributing or not. The transcript the contract
// streams is compressed into one BRLC commitment and its hierarchical
// Poseidon digest is a public input, so the Fiat–Shamir challenge commits to
// the witness words and not only to the calldata.
//
// Public inputs, in declaration order, are the 7 words the verifier reads:
//
//	RoundHash, Threshold, CommitteeSize, AcceptedCount,
//	TranscriptDigest, Challenge, TranscriptCommitment
type FinalizeCircuit struct {
	RoundHash     frontend.Variable `gnark:",public"`
	Threshold     frontend.Variable `gnark:",public"`
	CommitteeSize frontend.Variable `gnark:",public"`
	AcceptedCount frontend.Variable `gnark:",public"`
	// TranscriptDigest = T of §7: H(2, eid, t, n, a, K, L_F, R, B_0…B_(K−1))
	// with R the digest of the dealer rows and B_j the digest of key j's
	// transcript block, all over the exact masked words the BRLC commits to.
	TranscriptDigest     frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	ParticipantIndexes [MaxParticipants]frontend.Variable
	ContributionHashes [MaxParticipants]frontend.Variable
	// Commitments[d][j][m] is dealer row d's commitment to coefficient m of
	// key j: every vector of every accepted contribution, because a dealer's
	// stored hash absorbs all MaxK key digests.
	Commitments          [MaxParticipants][MaxKeys][MaxCoefficients]twistededwards.Point
	AggregateCommitments [MaxKeys][MaxCoefficients]twistededwards.Point
	// ShareCommitments[j][i] is D[j][i], committee member i+1's share
	// commitment for key j, the Merkle leaf its partials prove against.
	ShareCommitments [MaxKeys][MaxParticipants]twistededwards.Point
}

func (c *FinalizeCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ccommon.BabyJubJubCurveID())
	if err != nil {
		return err
	}
	// 1 ≤ t ≤ a ≤ n ≤ N. Bounding every count to its array size keeps
	// PrefixMask from masking the wrong slot count.
	api.AssertIsDifferent(c.Threshold, 0)
	api.AssertIsLessOrEqual(c.Threshold, MaxCoefficients)
	api.AssertIsLessOrEqual(c.CommitteeSize, MaxParticipants)
	api.AssertIsLessOrEqual(c.AcceptedCount, c.CommitteeSize)
	api.AssertIsLessOrEqual(c.Threshold, c.AcceptedCount)

	coeffMask := ccommon.PrefixMask(api, c.Threshold, MaxCoefficients)
	dealerMask := ccommon.PrefixMask(api, c.AcceptedCount, MaxParticipants)
	committeeMask := ccommon.PrefixMask(api, c.CommitteeSize, MaxParticipants)

	// Dealer rows. An active row names a real member in [1, n] and reproduces
	// its stored commitmentsHash from all MaxK recomputed key digests, once.
	// Its commitments, masked to the identity beyond the threshold and on
	// inactive rows, feed the aggregates below.
	maskedIndexes := make([]frontend.Variable, MaxParticipants)
	maskedHashes := make([]frontend.Variable, MaxParticipants)
	var maskedRows [MaxParticipants][MaxKeys][MaxCoefficients]twistededwards.Point
	for d := range MaxParticipants {
		active := dealerMask[d]
		maskedIndexes[d] = api.Mul(active, c.ParticipantIndexes[d])
		maskedHashes[d] = api.Mul(active, c.ContributionHashes[d])
		// Index 0 is the inactive marker and must never be hashed as a dealer.
		api.AssertIsEqual(api.Mul(active, api.IsZero(c.ParticipantIndexes[d])), 0)
		// maskedIndex ≤ n: both the index and n − index fit IndexBits bits.
		// n ≤ MaxN < 2^IndexBits, so a negative difference wraps to a field
		// element far outside the range and fails the decomposition.
		bits.ToBinary(api, maskedIndexes[d], bits.WithNbDigits(ccommon.IndexBits))
		bits.ToBinary(api, api.Sub(c.CommitteeSize, maskedIndexes[d]), bits.WithNbDigits(ccommon.IndexBits))

		keyDigests := make([]frontend.Variable, MaxKeys)
		for j := range MaxKeys {
			digestPoints := make([]twistededwards.Point, MaxCoefficients)
			for m := range MaxCoefficients {
				if err := ccommon.AssertPointOnCurve(api, c.Commitments[d][j][m]); err != nil {
					return err
				}
				// The dealer digested its commitments with the slots beyond
				// the threshold folded to the identity; reproduce that
				// masking or the recomputed hash would not match.
				digestPoints[m] = ccommon.MaskPoint(api, coeffMask[m], c.Commitments[d][j][m])
				maskedRows[d][j][m] = ccommon.MaskPoint(api, active, digestPoints[m])
			}
			keyDigests[j], err = ccommon.CommitmentKeyDigest(api, digestPoints)
			if err != nil {
				return err
			}
		}
		recomputed, err := ccommon.CommitmentsHash(api, c.RoundHash, c.ParticipantIndexes[d], c.Threshold, keyDigests)
		if err != nil {
			return err
		}
		// Conditional equality: an inactive row carries zeros and is not
		// required to reproduce any stored hash.
		api.AssertIsEqual(api.Mul(active, api.Sub(recomputed, c.ContributionHashes[d])), 0)
	}

	transcript := make([]frontend.Variable, 0, TranscriptWords)
	transcript = append(transcript, maskedIndexes...)
	transcript = append(transcript, maskedHashes...)
	rowsDigest, err := ccommon.MultiHash(api, append([]frontend.Variable{digestTagRows}, transcript...)...)
	if err != nil {
		return err
	}

	keyDigests := make([]frontend.Variable, MaxKeys)
	for j := range MaxKeys {
		// Ā[j][m] = Σ_d A[d][j][m] over the accepted rows, identity beyond
		// the threshold. The witness aggregate is pinned to the running sum
		// on active coefficients and masked for the Horner chain below.
		maskedAggregate := make([]twistededwards.Point, MaxCoefficients)
		for m := range MaxCoefficients {
			if err := ccommon.AssertPointOnCurve(api, c.AggregateCommitments[j][m]); err != nil {
				return err
			}
			sum := maskedRows[0][j][m]
			for d := 1; d < MaxParticipants; d++ {
				sum = curve.Add(sum, maskedRows[d][j][m])
			}
			api.AssertIsEqual(api.Mul(coeffMask[m], api.Sub(c.AggregateCommitments[j][m].X, sum.X)), 0)
			api.AssertIsEqual(api.Mul(coeffMask[m], api.Sub(c.AggregateCommitments[j][m].Y, sum.Y)), 0)
			maskedAggregate[m] = ccommon.MaskPoint(api, coeffMask[m], c.AggregateCommitments[j][m])
		}

		// Key block: P[j] = Ā[j][0] (t ≥ 1, so slot 0 is always active), then
		// D[j][i] for every committee member i+1 — contributing or not, since
		// every member received a share of every accepted polynomial — and
		// the identity for slots beyond the committee. The evaluation point
		// is a constant, so the Horner chain costs no scalar decomposition.
		keyWords := make([]frontend.Variable, 0, KeyWords)
		keyWords = append(keyWords, maskedAggregate[0].X, maskedAggregate[0].Y)
		for i := range MaxParticipants {
			if err := ccommon.AssertPointOnCurve(api, c.ShareCommitments[j][i]); err != nil {
				return err
			}
			expected, err := ccommon.CommitmentPolynomialValue(api, maskedAggregate, i+1)
			if err != nil {
				return err
			}
			api.AssertIsEqual(api.Mul(committeeMask[i], api.Sub(c.ShareCommitments[j][i].X, expected.X)), 0)
			api.AssertIsEqual(api.Mul(committeeMask[i], api.Sub(c.ShareCommitments[j][i].Y, expected.Y)), 0)
			masked := ccommon.MaskPoint(api, committeeMask[i], c.ShareCommitments[j][i])
			keyWords = append(keyWords, masked.X, masked.Y)
		}
		transcript = append(transcript, keyWords...)
		keyDigests[j], err = ccommon.MultiHash(api, append([]frontend.Variable{digestTagKey, j}, keyWords...)...)
		if err != nil {
			return err
		}
	}

	// T = H(2, eid, t, n, a, K, L_F, R, B_0, …, B_(K−1)): the digest the
	// contract folds into the challenge anchor, so ρ is derived after the
	// witness transcript is fixed.
	outer := []frontend.Variable{
		digestTagOuter, c.RoundHash, c.Threshold, c.CommitteeSize, c.AcceptedCount,
		MaxKeys, TranscriptWords, rowsDigest,
	}
	digest, err := ccommon.MultiHash(api, append(outer, keyDigests...)...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.TranscriptDigest, digest)
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
