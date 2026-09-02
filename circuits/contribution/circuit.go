package contribution

import (
	ecc_tweds "github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// MaxCoefficients/MaxRecipients are aliases of the single shared constant
// `circuits/common.MaxN`. Edit `circuits/common/sizes.go` to change the bound.
const (
	MaxCoefficients = ccommon.MaxN
	MaxRecipients   = ccommon.MaxN
)

// ContributionCircuit proves the full DKG phase-4 statement from the paper:
// coefficient commitments, Feldman consistency, nonce points, and hashed share encryption.
type ContributionCircuit struct {
	RoundHash            frontend.Variable `gnark:",public"`
	Threshold            frontend.Variable `gnark:",public"`
	CommitteeSize        frontend.Variable `gnark:",public"`
	ContributorIndex     frontend.Variable `gnark:",public"`
	CommitmentHash       frontend.Variable `gnark:",public"`
	ShareHash            frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	Commitments      [MaxCoefficients]twistededwards.Point
	RecipientPubKeys [MaxRecipients]twistededwards.Point
	Ephemerals       [MaxRecipients]twistededwards.Point
	MaskedShares     [MaxRecipients]frontend.Variable

	Coefficients       [MaxCoefficients]frontend.Variable
	EncryptionNonces   [MaxRecipients]frontend.Variable
	RecipientIndexes   [MaxRecipients]frontend.Variable
	Shares             [MaxRecipients]frontend.Variable
	MaskQuotients      [MaxRecipients]frontend.Variable
	ShareMasks         [MaxRecipients]frontend.Variable
	MaskedShareCarries [MaxRecipients]frontend.Variable
}

func (c *ContributionCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ecc_tweds.BN254)
	if err != nil {
		return err
	}
	// Bound the public count inputs to their fixed array sizes.
	// PrefixMask returns all-active when count > size, so
	// without these the statement could prove a partial set while
	// claiming a larger one.
	api.AssertIsLessOrEqual(c.Threshold, MaxCoefficients)
	api.AssertIsLessOrEqual(c.CommitteeSize, MaxRecipients)
	api.AssertIsLessOrEqual(c.Threshold, c.CommitteeSize)
	// Honest contributor index is one-based in [1, CommitteeSize]. Asserting
	// the upper bound here closes a future-composition gap; the contract
	// already enforces non-zero.
	api.AssertIsLessOrEqual(c.ContributorIndex, c.CommitteeSize)

	coeffMask := ccommon.PrefixMask(api, c.Threshold, MaxCoefficients)
	recipientMask := ccommon.PrefixMask(api, c.CommitteeSize, MaxRecipients)

	// Pre-mask the coefficients and the commitment points once, so that the
	// per-recipient CommitmentPolynomialValue calls below
	// can iterate without repeating per-coefficient Select work. Inactive slots
	// are folded to 0 (scalars) and the curve identity (0, 1) (points), which
	// makes a subsequent unconditional Add a no-op.
	maskedCommitments := make([]twistededwards.Point, MaxCoefficients)
	commitmentInputs := []frontend.Variable{c.RoundHash, c.ContributorIndex, c.Threshold}
	subgroupOrderMinusOne := ccommon.SubgroupOrderMinusOne()
	for i := range MaxCoefficients {
		if err := ccommon.AssertPointOnCurve(api, c.Commitments[i]); err != nil {
			return err
		}
		// Range-check the coefficient witness to its canonical [0, r) form.
		// FixedBaseMul itself wraps mod r, so the constraint is defence in
		// depth against future composition.
		api.AssertIsLessOrEqual(c.Coefficients[i], subgroupOrderMinusOne)
		expectedCommitment := ccommon.FixedBaseMul(api, c.Coefficients[i])
		// Conditional equality: when coeffMask[i] == 1 the witness commitment
		// must equal the FixedBaseMul of the witness coefficient; otherwise the
		// constraint is trivially satisfied. Replaces 4 Selects + 2 Asserts
		// (~6 constraints) with 2 Muls + 2 Asserts (~4 constraints).
		dCommitX := api.Sub(c.Commitments[i].X, expectedCommitment.X)
		dCommitY := api.Sub(c.Commitments[i].Y, expectedCommitment.Y)
		api.AssertIsEqual(api.Mul(coeffMask[i], dCommitX), 0)
		api.AssertIsEqual(api.Mul(coeffMask[i], dCommitY), 0)

		// Pre-masked commitments reused across the recipient loop.
		maskedCommitments[i] = ccommon.MaskPoint(api, coeffMask[i], c.Commitments[i])

		commitmentInputs = append(
			commitmentInputs,
			maskedCommitments[i].X,
			maskedCommitments[i].Y,
		)
	}
	commitmentHash, err := ccommon.MultiHash(api, commitmentInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.CommitmentHash, commitmentHash)

	shareInputs := []frontend.Variable{c.RoundHash, c.ContributorIndex, c.CommitteeSize}
	maskedIndexes := make([]frontend.Variable, MaxRecipients)
	maskedKeys := make([]twistededwards.Point, MaxRecipients)
	maskedEphemerals := make([]twistededwards.Point, MaxRecipients)
	maskedShares := make([]frontend.Variable, MaxRecipients)
	for i := range MaxRecipients {
		// Shares[i] must lie in [0, r). Without this check
		// the prover can pick s' = honest_share + 7·r (still <p when
		// honest_share<δ) and have AddModSubgroupOrder publish a
		// MaskedShare that decrypts to honest_share+(r−δ)≠honest_share.
		// That breaks recipient-side Feldman and DoSes the epoch.
		api.AssertIsLessOrEqual(c.Shares[i], subgroupOrderMinusOne)
		activeShare := api.Select(recipientMask[i], c.Shares[i], 0)

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

		feldmanPoint, err := ccommon.CommitmentPolynomialValue(api, maskedCommitments, c.RecipientIndexes[i])
		if err != nil {
			return err
		}
		sharePoint := ccommon.FixedBaseMul(api, activeShare)
		// Conditional equality on the Feldman consistency check.
		dFeldX := api.Sub(sharePoint.X, feldmanPoint.X)
		dFeldY := api.Sub(sharePoint.Y, feldmanPoint.Y)
		api.AssertIsEqual(api.Mul(recipientMask[i], dFeldX), 0)
		api.AssertIsEqual(api.Mul(recipientMask[i], dFeldY), 0)

		expectedEphemeral := ccommon.FixedBaseMul(api, c.EncryptionNonces[i])
		// Conditional equality on the ephemeral consistency check.
		dEphX := api.Sub(c.Ephemerals[i].X, expectedEphemeral.X)
		dEphY := api.Sub(c.Ephemerals[i].Y, expectedEphemeral.Y)
		api.AssertIsEqual(api.Mul(recipientMask[i], dEphX), 0)
		api.AssertIsEqual(api.Mul(recipientMask[i], dEphY), 0)

		sharedSecret := curve.ScalarMul(c.RecipientPubKeys[i], c.EncryptionNonces[i])
		rawMask, err := ccommon.ShareMaskHash(
			api,
			c.RoundHash,
			c.ContributorIndex,
			c.RecipientIndexes[i],
			sharedSecret.X,
			sharedSecret.Y,
		)
		if err != nil {
			return err
		}
		activeRawMask := api.Select(recipientMask[i], rawMask, 0)
		activeMaskQuotient := api.Select(recipientMask[i], c.MaskQuotients[i], 0)
		activeReducedMask := api.Select(recipientMask[i], c.ShareMasks[i], 0)
		mask := ccommon.ReduceToSubgroupOrder(api, activeRawMask, activeMaskQuotient, activeReducedMask)
		activeMaskedShare := api.Select(recipientMask[i], c.MaskedShares[i], 0)
		expectedMaskedShare := ccommon.AddModSubgroupOrder(
			api,
			activeShare,
			mask,
			api.Select(recipientMask[i], c.MaskedShareCarries[i], 0),
			activeMaskedShare,
		)
		api.AssertIsEqual(
			activeMaskedShare,
			api.Select(recipientMask[i], expectedMaskedShare, 0),
		)

		// Every transcript word must be fixed by a digest before the BRLC
		// challenge exists (the contract derives ρ from the digests and the
		// calldata), so the share digest also absorbs the recipient keys and
		// all four vectors are masked to constants in inactive slots.
		maskedIndexes[i] = api.Select(recipientMask[i], c.RecipientIndexes[i], 0)
		maskedKeys[i] = ccommon.MaskPoint(api, recipientMask[i], c.RecipientPubKeys[i])
		maskedEphemerals[i] = ccommon.MaskPoint(api, recipientMask[i], c.Ephemerals[i])
		maskedShares[i] = activeMaskedShare
		shareInputs = append(
			shareInputs,
			maskedIndexes[i],
			maskedKeys[i].X, maskedKeys[i].Y,
			maskedEphemerals[i].X, maskedEphemerals[i].Y,
			maskedShares[i],
		)
	}
	shareHash, err := ccommon.MultiHash(api, shareInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.ShareHash, shareHash)
	transcript := make([]frontend.Variable, 0, 8*ccommon.MaxN)
	for i := range MaxCoefficients {
		transcript = append(transcript, maskedCommitments[i].X, maskedCommitments[i].Y)
	}
	transcript = append(transcript, maskedIndexes...)
	for i := range MaxRecipients {
		transcript = append(transcript, maskedKeys[i].X, maskedKeys[i].Y)
	}
	for i := range MaxRecipients {
		transcript = append(transcript, maskedEphemerals[i].X, maskedEphemerals[i].Y)
	}
	transcript = append(transcript, maskedShares...)
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
