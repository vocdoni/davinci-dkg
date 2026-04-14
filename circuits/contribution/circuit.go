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
	coeffMask := ccommon.PrefixMask(api, c.Threshold, MaxCoefficients)
	recipientMask := ccommon.PrefixMask(api, c.CommitteeSize, MaxRecipients)

	// Pre-mask the coefficients and the commitment points once, so that the
	// per-recipient EvaluatePolynomial / CommitmentPolynomialValue calls below
	// can iterate without repeating per-coefficient Select work. Inactive slots
	// are folded to 0 (scalars) and the curve identity (0, 1) (points), which
	// makes a subsequent unconditional Add a no-op.
	maskedCoeffs := make([]frontend.Variable, MaxCoefficients)
	maskedCommitments := make([]twistededwards.Point, MaxCoefficients)
	commitmentInputs := []frontend.Variable{c.RoundHash, c.ContributorIndex, c.Threshold}
	for i := range MaxCoefficients {
		if err := ccommon.AssertPointOnCurve(api, c.Commitments[i]); err != nil {
			return err
		}
		expectedCommitment := ccommon.FixedBaseMul(api, c.Coefficients[i])
		// Conditional equality: when coeffMask[i] == 1 the witness commitment
		// must equal the FixedBaseMul of the witness coefficient; otherwise the
		// constraint is trivially satisfied. Replaces 4 Selects + 2 Asserts
		// (~6 constraints) with 2 Muls + 2 Asserts (~4 constraints).
		dCommitX := api.Sub(c.Commitments[i].X, expectedCommitment.X)
		dCommitY := api.Sub(c.Commitments[i].Y, expectedCommitment.Y)
		api.AssertIsEqual(api.Mul(coeffMask[i], dCommitX), 0)
		api.AssertIsEqual(api.Mul(coeffMask[i], dCommitY), 0)

		// Pre-masked artefacts reused across the recipient loop.
		maskedCoeffs[i] = api.Mul(coeffMask[i], c.Coefficients[i])
		maskedCommitments[i] = twistededwards.Point{
			X: api.Mul(coeffMask[i], c.Commitments[i].X),
			Y: api.Select(coeffMask[i], c.Commitments[i].Y, 1),
		}

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
	for i := range MaxRecipients {
		expected := ccommon.EvaluatePolynomial(api, maskedCoeffs, nil, c.RecipientIndexes[i])
		activeShare := api.Select(recipientMask[i], c.Shares[i], 0)
		api.AssertIsEqual(activeShare, api.Select(recipientMask[i], expected, 0))

		if err := ccommon.AssertPointOnCurve(api, c.RecipientPubKeys[i]); err != nil {
			return err
		}
		if err := ccommon.AssertPointOnCurve(api, c.Ephemerals[i]); err != nil {
			return err
		}

		feldmanPoint, err := ccommon.CommitmentPolynomialValue(api, maskedCommitments, nil, c.RecipientIndexes[i])
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

		shareInputs = append(shareInputs,
			api.Select(recipientMask[i], c.RecipientIndexes[i], 0),
			api.Select(recipientMask[i], c.Ephemerals[i].X, 0),
			api.Select(recipientMask[i], c.Ephemerals[i].Y, 1),
			api.Select(recipientMask[i], c.MaskedShares[i], 0),
		)
	}
	shareHash, err := ccommon.MultiHash(api, shareInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.ShareHash, shareHash)
	transcript := make([]frontend.Variable, 0, 64)
	for i := range MaxCoefficients {
		transcript = append(transcript, c.Commitments[i].X, c.Commitments[i].Y)
	}
	for i := range MaxRecipients {
		transcript = append(transcript, c.RecipientIndexes[i])
	}
	for i := range MaxRecipients {
		transcript = append(transcript, c.RecipientPubKeys[i].X, c.RecipientPubKeys[i].Y)
	}
	for i := range MaxRecipients {
		transcript = append(transcript, c.Ephemerals[i].X, c.Ephemerals[i].Y)
	}
	for i := range MaxRecipients {
		transcript = append(transcript, c.MaskedShares[i])
	}
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
