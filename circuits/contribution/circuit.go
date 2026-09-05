package contribution

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// MaxCoefficients/MaxRecipients are aliases of the single shared constant
// `circuits/common.MaxN`; MaxKeys is `circuits/common.MaxK`, the number of
// pool keys every epoch deals. Edit `circuits/common/sizes.go` to change them.
const (
	MaxCoefficients = ccommon.MaxN
	MaxRecipients   = ccommon.MaxN
	MaxKeys         = ccommon.MaxK
)

// ContributionCircuit proves the DKG dealing statement for the whole pool at
// once: MaxK independent polynomials, their coefficient commitments, Feldman
// consistency of every share, and hashed share encryption under one ECDH
// secret per recipient — the key index separates the MaxK masks derived from
// that single secret.
type ContributionCircuit struct {
	RoundHash            frontend.Variable `gnark:",public"`
	Threshold            frontend.Variable `gnark:",public"`
	CommitteeSize        frontend.Variable `gnark:",public"`
	ContributorIndex     frontend.Variable `gnark:",public"`
	CommitmentHash       frontend.Variable `gnark:",public"`
	ShareHash            frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	Commitments      [MaxKeys][MaxCoefficients]twistededwards.Point
	RecipientPubKeys [MaxRecipients]twistededwards.Point
	Ephemerals       [MaxRecipients]twistededwards.Point
	MaskedShares     [MaxKeys][MaxRecipients]frontend.Variable

	Coefficients       [MaxKeys][MaxCoefficients]frontend.Variable
	EncryptionNonces   [MaxRecipients]frontend.Variable
	RecipientIndexes   [MaxRecipients]frontend.Variable
	Shares             [MaxKeys][MaxRecipients]frontend.Variable
	MaskQuotients      [MaxKeys][MaxRecipients]frontend.Variable
	ShareMasks         [MaxKeys][MaxRecipients]frontend.Variable
	MaskedShareCarries [MaxKeys][MaxRecipients]frontend.Variable
}

func (c *ContributionCircuit) Define(api frontend.API) error {
	// Bound the public count inputs to their fixed array sizes.
	// PrefixMask returns all-active when count > size, so
	// without these the statement could prove a partial set while
	// claiming a larger one.
	// 1 ≤ t ≤ n ≤ MaxN and 1 ≤ contributorIndex ≤ n (docs/pool-keys-v4.md
	// §3). The compact transcript length is a function of t and n, so the
	// zero cases are excluded here and not left to the contract alone.
	api.AssertIsDifferent(c.Threshold, 0)
	api.AssertIsLessOrEqual(c.Threshold, MaxCoefficients)
	api.AssertIsLessOrEqual(c.CommitteeSize, MaxRecipients)
	api.AssertIsLessOrEqual(c.Threshold, c.CommitteeSize)
	api.AssertIsDifferent(c.ContributorIndex, 0)
	api.AssertIsLessOrEqual(c.ContributorIndex, c.CommitteeSize)

	coeffMask := ccommon.PrefixMask(api, c.Threshold, MaxCoefficients)
	recipientMask := ccommon.PrefixMask(api, c.CommitteeSize, MaxRecipients)
	subgroupOrderMinusOne := ccommon.SubgroupOrderMinusOne()

	// Pre-mask each key's coefficients and commitment points once, so that the
	// per-recipient CommitmentPolynomialValue calls below can iterate without
	// repeating per-coefficient Select work. Inactive slots are folded to 0
	// (scalars) and the curve identity (0, 1) (points), which makes a
	// subsequent unconditional Add a no-op.
	var maskedCommitments [MaxKeys][MaxCoefficients]twistededwards.Point
	keyDigests := make([]frontend.Variable, MaxKeys)
	var err error
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			if err := ccommon.AssertPointOnCurve(api, c.Commitments[j][m]); err != nil {
				return err
			}
			// Range-check the coefficient witness to its canonical [0, r) form.
			// FixedBaseMul itself wraps mod r, so the constraint is defence in
			// depth against future composition.
			api.AssertIsLessOrEqual(c.Coefficients[j][m], subgroupOrderMinusOne)
			expectedCommitment := ccommon.FixedBaseMul(api, c.Coefficients[j][m])
			// Conditional equality: when coeffMask[m] == 1 the witness commitment
			// must equal the FixedBaseMul of the witness coefficient; otherwise the
			// constraint is trivially satisfied. Replaces 4 Selects + 2 Asserts
			// (~6 constraints) with 2 Muls + 2 Asserts (~4 constraints).
			dCommitX := api.Sub(c.Commitments[j][m].X, expectedCommitment.X)
			dCommitY := api.Sub(c.Commitments[j][m].Y, expectedCommitment.Y)
			api.AssertIsEqual(api.Mul(coeffMask[m], dCommitX), 0)
			api.AssertIsEqual(api.Mul(coeffMask[m], dCommitY), 0)

			maskedCommitments[j][m] = ccommon.MaskPoint(api, coeffMask[m], c.Commitments[j][m])
		}
		keyDigests[j], err = ccommon.CommitmentKeyDigest(api, maskedCommitments[j][:])
		if err != nil {
			return err
		}
	}
	commitmentHash, err := ccommon.CommitmentsHash(api, c.RoundHash, c.ContributorIndex, c.Threshold, keyDigests)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.CommitmentHash, commitmentHash)

	maskedIndexes := make([]frontend.Variable, MaxRecipients)
	maskedKeys := make([]twistededwards.Point, MaxRecipients)
	maskedEphemerals := make([]twistededwards.Point, MaxRecipients)
	rowDigests := make([]frontend.Variable, MaxRecipients)
	var maskedShares [MaxKeys][MaxRecipients]frontend.Variable
	for i := range MaxRecipients {
		if err := ccommon.AssertPointOnCurve(api, c.RecipientPubKeys[i]); err != nil {
			return err
		}
		// Ephemerals[i] doesn't need an explicit on-curve check: when
		// recipientMask[i] == 1 the conditional equality below forces it
		// to equal `expectedEphemeral = FixedBaseMul(EncryptionNonces[i])`
		// which is on-curve by construction. When inactive the value is
		// masked out of the share-hash and transcript and the sharedSecret
		// scalar mul on RecipientPubKeys[i] is the only consumer left.

		// Range-check the recipient index to ≤ MaxRecipients (one-based;
		// the contract enforces non-zero). CommitmentPolynomialValue also
		// bounds it to IndexBits bits for its short scalar multiplications.
		api.AssertIsLessOrEqual(c.RecipientIndexes[i], MaxRecipients)

		expectedEphemeral := ccommon.FixedBaseMul(api, c.EncryptionNonces[i])
		// Conditional equality on the ephemeral consistency check.
		dEphX := api.Sub(c.Ephemerals[i].X, expectedEphemeral.X)
		dEphY := api.Sub(c.Ephemerals[i].Y, expectedEphemeral.Y)
		api.AssertIsEqual(api.Mul(recipientMask[i], dEphX), 0)
		api.AssertIsEqual(api.Mul(recipientMask[i], dEphY), 0)

		// One ECDH secret per recipient, shared by all MaxK keys: the whole
		// point of dealing the pool in one proof. The key index enters the
		// mask hash below so the MaxK masks derived from it stay independent
		// — reusing one mask would be a one-time-pad reuse across keys.
		sharedSecret := ccommon.ScalarMulVar(api, c.RecipientPubKeys[i], c.EncryptionNonces[i])

		// Every transcript word must be fixed by a digest before the BRLC
		// challenge exists (the contract derives ρ from the digests and the
		// calldata), so the share digest also absorbs the recipient keys and
		// every vector is masked to constants in inactive slots.
		maskedIndexes[i] = api.Select(recipientMask[i], c.RecipientIndexes[i], 0)
		maskedKeys[i] = ccommon.MaskPoint(api, recipientMask[i], c.RecipientPubKeys[i])
		maskedEphemerals[i] = ccommon.MaskPoint(api, recipientMask[i], c.Ephemerals[i])

		rowShares := make([]frontend.Variable, MaxKeys)
		for j := range MaxKeys {
			// Shares[j][i] must lie in [0, r). Without this check
			// the prover can pick s' = honest_share + 7·r (still <p when
			// honest_share<δ) and have AddModSubgroupOrder publish a
			// MaskedShare that decrypts to honest_share+(r−δ)≠honest_share.
			// That breaks recipient-side Feldman and DoSes the epoch.
			api.AssertIsLessOrEqual(c.Shares[j][i], subgroupOrderMinusOne)
			activeShare := api.Select(recipientMask[i], c.Shares[j][i], 0)

			feldmanPoint, err := ccommon.CommitmentPolynomialValue(api, maskedCommitments[j][:], c.RecipientIndexes[i])
			if err != nil {
				return err
			}
			sharePoint := ccommon.FixedBaseMul(api, activeShare)
			// Conditional equality on the Feldman consistency check.
			dFeldX := api.Sub(sharePoint.X, feldmanPoint.X)
			dFeldY := api.Sub(sharePoint.Y, feldmanPoint.Y)
			api.AssertIsEqual(api.Mul(recipientMask[i], dFeldX), 0)
			api.AssertIsEqual(api.Mul(recipientMask[i], dFeldY), 0)

			rawMask, err := ccommon.ShareMaskHash(
				api,
				c.RoundHash,
				c.ContributorIndex,
				c.RecipientIndexes[i],
				sharedSecret.X,
				sharedSecret.Y,
				j,
			)
			if err != nil {
				return err
			}
			mask := ccommon.ReduceToSubgroupOrder(
				api,
				api.Select(recipientMask[i], rawMask, 0),
				api.Select(recipientMask[i], c.MaskQuotients[j][i], 0),
				api.Select(recipientMask[i], c.ShareMasks[j][i], 0),
			)
			activeMaskedShare := api.Select(recipientMask[i], c.MaskedShares[j][i], 0)
			// AddModSubgroupOrder pins activeMaskedShare to share + mask mod r.
			ccommon.AddModSubgroupOrder(
				api,
				activeShare,
				mask,
				api.Select(recipientMask[i], c.MaskedShareCarries[j][i], 0),
				activeMaskedShare,
			)
			maskedShares[j][i] = activeMaskedShare
			rowShares[j] = activeMaskedShare
		}
		// One Poseidon per recipient row, one over the row digests: the flat
		// list would exceed the sponge's input cap at MaxK·MaxRecipients words.
		rowDigests[i], err = ccommon.EncryptedShareRowDigest(
			api,
			maskedIndexes[i],
			maskedKeys[i],
			maskedEphemerals[i],
			rowShares,
		)
		if err != nil {
			return err
		}
	}
	shareHash, err := ccommon.EncryptedSharesHash(api, c.RoundHash, c.ContributorIndex, c.CommitteeSize, rowDigests)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.ShareHash, shareHash)

	// Compact BRLC (docs/pool-keys-v4.md §4): the same fixed-size region
	// order as before, but every word is gated by the public counts — a
	// commitment coordinate by [m < t], everything else by [i < n] — so an
	// inactive slot neither contributes nor advances the exponent. The fold
	// therefore equals the contract's canonical BRLC over the L_C calldata
	// words, which carry no padding, and the word count pins the gates to
	// the public t and n.
	fold := ccommon.NewGatedFold(api, c.Challenge)
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			fold.Absorb(coeffMask[m], maskedCommitments[j][m].X, maskedCommitments[j][m].Y)
		}
	}
	for i := range MaxRecipients {
		fold.Absorb(recipientMask[i], maskedIndexes[i])
	}
	for i := range MaxRecipients {
		fold.Absorb(recipientMask[i], maskedKeys[i].X, maskedKeys[i].Y)
	}
	for i := range MaxRecipients {
		fold.Absorb(recipientMask[i], maskedEphemerals[i].X, maskedEphemerals[i].Y)
	}
	for j := range MaxKeys {
		for i := range MaxRecipients {
			fold.Absorb(recipientMask[i], maskedShares[j][i])
		}
	}
	// L_C = MaxK·(2t + n) + 5n, linear in the public inputs.
	expectedWords := api.Add(
		api.Mul(MaxKeys, api.Add(api.Mul(2, c.Threshold), c.CommitteeSize)),
		api.Mul(5, c.CommitteeSize),
	)
	api.AssertIsEqual(fold.Count(), expectedWords)
	api.AssertIsEqual(c.TranscriptCommitment, fold.Commitment())
	return nil
}
