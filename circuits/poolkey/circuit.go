// Package poolkey holds the per-key activation circuit that turns one of an
// epoch's MaxK dealt polynomials into a usable committee key. It replaces the
// old single-key finalize circuit: finalizing an epoch no longer needs a
// proof, and each pool key is activated on its own, permissionlessly, from
// the accepted contributions' public commitments. See docs/pool-keys.md.
package poolkey

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
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

// TranscriptWords is the fixed word count of the activation calldata
// transcript (6·MaxN, see docs/pool-keys.md):
//
//	[0, N)    participantIndexes   (0 when the accepted row is inactive)
//	[N, 2N)   contributionHashes   (0 when inactive)
//	[2N, 4N)  aggregateCommitments (x, y) ((0,1) for m >= t)
//	[4N, 6N)  shareCommitments D_p (x, y) at slot p−1 for every committee
//	          member p ≤ committeeSize, contributing or not; (0,1) beyond
const TranscriptWords = 6 * MaxParticipants

// PoolKeyCircuit proves, for one pool key j: every accepted contributor's
// on-chain commitments hash is reproduced from its commitments for key j plus
// the digests of its other keys; Ā[m] is the sum of those commitments over the
// accepted set; and D_p = Σ_m p^m · Ā[m] is member p's share commitment for
// every committee member. P_j is Ā[0], read off the transcript by the
// contract. The wide transcript is compressed into one BRLC commitment to keep
// the public inputs constant-sized, and its Poseidon digest is a public input
// so the Fiat–Shamir challenge commits to the witness words, not only to the
// calldata.
//
// Public inputs, in declaration order, are the 8 words the verifier reads:
//
//	RoundHash, Threshold, CommitteeSize, AcceptedCount, KeyIndex,
//	TranscriptDigest, Challenge, TranscriptCommitment
type PoolKeyCircuit struct {
	RoundHash     frontend.Variable `gnark:",public"`
	Threshold     frontend.Variable `gnark:",public"`
	CommitteeSize frontend.Variable `gnark:",public"`
	AcceptedCount frontend.Variable `gnark:",public"`
	KeyIndex      frontend.Variable `gnark:",public"`
	// TranscriptDigest = MultiHash(RoundHash, KeyIndex, transcript words…)
	// over the exact masked vector the BRLC commits to. The contract folds it
	// into the challenge anchor, so ρ is derived after the witness transcript
	// is fixed; without it every transcript word would be a free witness
	// bound to calldata only by one linear BRLC equation.
	TranscriptDigest     frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	ParticipantIndexes [MaxParticipants]frontend.Variable
	ContributionHashes [MaxParticipants]frontend.Variable
	// KeyCommitments holds only the activated key's commitment vector per
	// contributor; OtherKeyDigests carries all MaxK of that contributor's key
	// digests, with slot KeyIndex ignored and replaced by the recomputed one.
	KeyCommitments       [MaxParticipants][MaxCoefficients]twistededwards.Point
	OtherKeyDigests      [MaxParticipants][MaxKeys]frontend.Variable
	AggregateCommitments [MaxCoefficients]twistededwards.Point
	// ShareCommitments[i] is D_{i+1}, the share commitment of committee
	// member i+1, whether or not that member contributed: the Merkle leaf a
	// partial decryption later proves against lives at slot i.
	ShareCommitments [MaxParticipants]twistededwards.Point
}

func (c *PoolKeyCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ccommon.BabyJubJubCurveID())
	if err != nil {
		return err
	}
	// Bound the public count inputs to their fixed array sizes and to each
	// other so PrefixMask cannot be coerced into masking the wrong slot count.
	api.AssertIsLessOrEqual(c.Threshold, MaxCoefficients)
	api.AssertIsLessOrEqual(c.CommitteeSize, MaxParticipants)
	api.AssertIsLessOrEqual(c.AcceptedCount, c.CommitteeSize)
	api.AssertIsLessOrEqual(c.Threshold, c.AcceptedCount)

	coeffMask := ccommon.PrefixMask(api, c.Threshold, MaxCoefficients)
	participantMask := ccommon.PrefixMask(api, c.AcceptedCount, MaxParticipants)
	committeeMask := ccommon.PrefixMask(api, c.CommitteeSize, MaxParticipants)

	// One-hot selector over the pool: exactly one slot of a contributor's MaxK
	// digests is the key being activated and gets the recomputed value, the
	// rest come from the witness. Asserting the selector sums to one is what
	// enforces KeyIndex < MaxK.
	keySelector := make([]frontend.Variable, MaxKeys)
	selectorSum := frontend.Variable(0)
	for j := range MaxKeys {
		keySelector[j] = api.IsZero(api.Sub(c.KeyIndex, j))
		selectorSum = api.Add(selectorSum, keySelector[j])
	}
	api.AssertIsEqual(selectorSum, 1)

	// Row-masked commitments: an inactive participant contributes the identity
	// to the aggregate, so its row is a no-op in the sum below.
	var maskedRows [MaxParticipants][MaxCoefficients]twistededwards.Point
	for i := range MaxParticipants {
		// An active row must name a real (one-based) member: index 0 is the
		// inactive marker and must never be hashed as a contributor.
		api.AssertIsEqual(api.Mul(participantMask[i], api.IsZero(c.ParticipantIndexes[i])), 0)
		api.AssertIsLessOrEqual(c.ParticipantIndexes[i], MaxParticipants)

		digestPoints := make([]twistededwards.Point, MaxCoefficients)
		for m := range MaxCoefficients {
			if err := ccommon.AssertPointOnCurve(api, c.KeyCommitments[i][m]); err != nil {
				return err
			}
			// The contributor digested its commitments with the slots beyond
			// the threshold folded to the identity; reproduce that masking or
			// the recomputed hash would not match the stored one.
			digestPoints[m] = ccommon.MaskPoint(api, coeffMask[m], c.KeyCommitments[i][m])
			maskedRows[i][m] = ccommon.MaskPoint(api, participantMask[i], digestPoints[m])
		}
		keyDigest, err := ccommon.CommitmentKeyDigest(api, digestPoints)
		if err != nil {
			return err
		}
		digests := make([]frontend.Variable, MaxKeys)
		for j := range MaxKeys {
			digests[j] = api.Select(keySelector[j], keyDigest, c.OtherKeyDigests[i][j])
		}
		recomputed, err := ccommon.CommitmentsHash(api, c.RoundHash, c.ParticipantIndexes[i], c.Threshold, digests)
		if err != nil {
			return err
		}
		// Conditional equality: an inactive slot carries zeros everywhere and
		// is not required to reproduce any stored hash.
		api.AssertIsEqual(api.Mul(participantMask[i], api.Sub(recomputed, c.ContributionHashes[i])), 0)
	}

	for m := range MaxCoefficients {
		if err := ccommon.AssertPointOnCurve(api, c.AggregateCommitments[m]); err != nil {
			return err
		}
		sum := ccommon.IdentityPoint()
		for i := range MaxParticipants {
			sum = curve.Add(sum, maskedRows[i][m])
		}
		// Conditional equality: when coeffMask[m] == 1 the aggregate must equal
		// the running sum; otherwise the constraint is trivially satisfied.
		// Two muls + two asserts instead of two SelectPoint + AssertPointEqual.
		diffX := api.Sub(c.AggregateCommitments[m].X, sum.X)
		diffY := api.Sub(c.AggregateCommitments[m].Y, sum.Y)
		api.AssertIsEqual(api.Mul(coeffMask[m], diffX), 0)
		api.AssertIsEqual(api.Mul(coeffMask[m], diffY), 0)
	}

	// Pre-mask the aggregate once: unused coefficient slots become the identity
	// so the Horner chain in CommitmentPolynomialValue can add unconditionally.
	maskedAggregate := make([]twistededwards.Point, MaxCoefficients)
	for m := range MaxCoefficients {
		maskedAggregate[m] = ccommon.MaskPoint(api, coeffMask[m], c.AggregateCommitments[m])
	}

	// Share commitments for the whole committee, by member index rather than
	// by accepted row: slot i is D_{i+1}, which exists for every member the
	// aggregate polynomial can be evaluated at, contributing or not. The
	// evaluation point is a constant, so the Horner chain costs no scalar
	// decomposition. Masked by the committee prefix, independent of the
	// accepted rows.
	maskedShareCommitments := make([]twistededwards.Point, MaxParticipants)
	for i := range MaxParticipants {
		if err := ccommon.AssertPointOnCurve(api, c.ShareCommitments[i]); err != nil {
			return err
		}
		shareCommitment, err := ccommon.CommitmentPolynomialValue(api, maskedAggregate, i+1)
		if err != nil {
			return err
		}
		dShareX := api.Sub(c.ShareCommitments[i].X, shareCommitment.X)
		dShareY := api.Sub(c.ShareCommitments[i].Y, shareCommitment.Y)
		api.AssertIsEqual(api.Mul(committeeMask[i], dShareX), 0)
		api.AssertIsEqual(api.Mul(committeeMask[i], dShareY), 0)
		maskedShareCommitments[i] = ccommon.MaskPoint(api, committeeMask[i], c.ShareCommitments[i])
	}

	transcript := make([]frontend.Variable, 0, TranscriptWords)
	for i := range MaxParticipants {
		transcript = append(transcript, api.Select(participantMask[i], c.ParticipantIndexes[i], 0))
	}
	for i := range MaxParticipants {
		transcript = append(transcript, api.Select(participantMask[i], c.ContributionHashes[i], 0))
	}
	for m := range MaxCoefficients {
		transcript = append(transcript, maskedAggregate[m].X, maskedAggregate[m].Y)
	}
	for i := range MaxParticipants {
		transcript = append(transcript, maskedShareCommitments[i].X, maskedShareCommitments[i].Y)
	}
	// Digest over the same masked vector, prefixed by the epoch and the key
	// so one transcript cannot be replayed for another activation.
	digestInputs := make([]frontend.Variable, 0, 2+TranscriptWords)
	digestInputs = append(digestInputs, c.RoundHash, c.KeyIndex)
	digestInputs = append(digestInputs, transcript...)
	digest, err := ccommon.MultiHash(api, digestInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.TranscriptDigest, digest)
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
