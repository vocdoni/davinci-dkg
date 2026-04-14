package decryptcombine

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// DecryptCombineCircuit reconstructs the combined decryption value from a
// threshold set of partial decryptions.
type DecryptCombineCircuit struct {
	RoundHash            frontend.Variable `gnark:",public"`
	Threshold            frontend.Variable `gnark:",public"`
	ShareCount           frontend.Variable `gnark:",public"`
	CombineHash          frontend.Variable `gnark:",public"`
	PlaintextHash        frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	CiphertextC1       twistededwards.Point
	CiphertextC2       twistededwards.Point
	Plaintext          frontend.Variable
	ParticipantIndexes [MaxShares]frontend.Variable
	PartialDecryptions [MaxShares]twistededwards.Point
	// LagrangeCoefficients are pre-computed natively in the BJJ scalar field (r_bjj)
	// and passed as private witnesses. Computing them in-circuit via api.Div would use
	// BN254.Fr arithmetic, giving wrong results for negative coefficients (e.g. -1 ≠ r_bjj-1).
	LagrangeCoefficients [MaxShares]frontend.Variable
}

func (c *DecryptCombineCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ccommon.BabyJubJubCurveID())
	if err != nil {
		return err
	}
	mask := ccommon.PrefixMask(api, c.ShareCount, MaxShares)
	for _, point := range []twistededwards.Point{c.CiphertextC1, c.CiphertextC2} {
		if err := ccommon.AssertPointOnCurve(api, point); err != nil {
			return err
		}
	}
	for i := range MaxShares {
		if err := ccommon.AssertPointOnCurve(api, c.PartialDecryptions[i]); err != nil {
			return err
		}
	}

	hashInputs := []frontend.Variable{
		c.RoundHash,
		c.Threshold,
		c.ShareCount,
		c.CiphertextC1.X,
		c.CiphertextC1.Y,
		c.CiphertextC2.X,
		c.CiphertextC2.Y,
	}
	for i := range MaxShares {
		hashInputs = append(
			hashInputs,
			api.Select(mask[i], c.ParticipantIndexes[i], 0),
			api.Select(mask[i], c.PartialDecryptions[i].X, 0),
			api.Select(mask[i], c.PartialDecryptions[i].Y, 1),
		)
	}
	combineHash, err := ccommon.MultiHash(api, hashInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.CombineHash, combineHash)
	api.AssertIsEqual(c.PlaintextHash, c.Plaintext)

	// The Lagrange interpolation accumulator. For inactive slots we mask the
	// scalar to 0 so curve.ScalarMul yields the identity, then unconditionally
	// add — saving the per-iteration result Select compared to the original.
	combined := ccommon.IdentityPoint()
	for i := range MaxShares {
		// Use pre-computed Lagrange coefficient (in r_bjj, the BJJ scalar field).
		// In-circuit api.Div uses BN254.Fr, giving wrong results for negative
		// coefficients because BN254.Fr-1 ≠ r_bjj-1 as BJJ scalars.
		lambda := api.Mul(mask[i], c.LagrangeCoefficients[i])
		scaled := curve.ScalarMul(c.PartialDecryptions[i], lambda)
		combined = curve.Add(combined, scaled)
	}
	messagePoint := ccommon.FixedBaseMul(api, c.Plaintext)
	expectedC2 := curve.Add(messagePoint, combined)
	ccommon.AssertPointEqual(api, expectedC2, c.CiphertextC2)
	transcript := make([]frontend.Variable, 0, 28)
	transcript = append(transcript, c.CiphertextC1.X, c.CiphertextC1.Y, c.CiphertextC2.X, c.CiphertextC2.Y)
	for i := range MaxShares {
		transcript = append(transcript, c.ParticipantIndexes[i])
	}
	for i := range MaxShares {
		transcript = append(transcript, c.PartialDecryptions[i].X, c.PartialDecryptions[i].Y)
	}
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
