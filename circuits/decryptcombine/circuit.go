package decryptcombine

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// DecryptCombineCircuit reconstructs the per-application plaintext from a
// threshold set of partial decryptions, with per-application correction
// term `T` selected by the `Mode` flag (paper §5.5 lines 1051–1077):
//
//	mode = 0 (public derivation): T = S · C_1 (computed in-circuit)
//	mode = 1 (organizer co-decryption): T = Δ_org (consumed as point)
//
//	M = C_2 - Σ λ_k · δ_{x_k} - T
type DecryptCombineCircuit struct {
	RoundHash            frontend.Variable `gnark:",public"` // semantically: eid
	Aid                  frontend.Variable `gnark:",public"` // application identifier
	CtIdx                frontend.Variable `gnark:",public"` // per-app ciphertext index
	Mode                 frontend.Variable `gnark:",public"` // 0 = public derivation, 1 = co-decryption
	S                    frontend.Variable `gnark:",public"` // derivation tag scalar; ignored if Mode=1
	DeltaOrg             twistededwards.Point `gnark:",public"` // organizer Δ_org; ignored if Mode=0
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
	// Per paper line 1092, soundness rests on the DLP-hardness argument: a forged
	// non-canonical λ-vector would force the prover to solve a discrete log to
	// produce a valid plaintext that the contract accepts.
	LagrangeCoefficients [MaxShares]frontend.Variable
}

func (c *DecryptCombineCircuit) Define(api frontend.API) error {
	curve, err := twistededwards.NewEdCurve(api, ccommon.BabyJubJubCurveID())
	if err != nil {
		return err
	}
	// Bound ShareCount to MaxShares so PrefixMask
	// truly masks to the circuit's fixed-size slot count. Without this,
	// a count > MaxShares leaves every slot active and the circuit /
	// contract disagree on the transcript extent.
	api.AssertIsLessOrEqual(c.ShareCount, MaxShares)
	mask := ccommon.PrefixMask(api, c.ShareCount, MaxShares)
	for _, point := range []twistededwards.Point{c.CiphertextC1, c.CiphertextC2, c.DeltaOrg} {
		if err := ccommon.AssertPointOnCurve(api, point); err != nil {
			return err
		}
	}
	for i := range MaxShares {
		if err := ccommon.AssertPointOnCurve(api, c.PartialDecryptions[i]); err != nil {
			return err
		}
	}

	// Mode flag must be 0 or 1: enforce mode·(mode-1) == 0.
	api.AssertIsEqual(api.Mul(c.Mode, api.Sub(c.Mode, 1)), 0)

	hashInputs := []frontend.Variable{
		c.RoundHash, // eid
		c.Aid,
		c.CtIdx,
		c.Mode,
		c.S,
		c.DeltaOrg.X,
		c.DeltaOrg.Y,
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
	// Plaintext is interpreted as a BJJ scalar (M = m·G).
	// Without a canonical-range bound the prover could pick m' ≡ m (mod r_bjj)
	// outside [0, r_bjj-1]. The recovered plaintext stored on-chain would then
	// disagree with what an honest verifier expects from the encrypted value.
	api.AssertIsLessOrEqual(c.Plaintext, ccommon.SubgroupOrderMinusOne())

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
	// Per-application correction term (paper §5.5 lines 1071–1077):
	//   mode = 0 (public derivation):  T = S · C_1 (computed in-circuit)
	//   mode = 1 (organizer co-decryption): T = Δ_org (consumed as point)
	// The S·C_1 in-circuit computation is the soundness fix described in
	// paper line 1088. Both candidates are computed unconditionally and
	// the mode flag selects which one binds the verifier; the cost is one
	// extra scalar multiplication on the unused branch.
	scC1 := curve.ScalarMul(c.CiphertextC1, c.S)
	correctionX := api.Select(c.Mode, c.DeltaOrg.X, scC1.X)
	correctionY := api.Select(c.Mode, c.DeltaOrg.Y, scC1.Y)
	correction := twistededwards.Point{X: correctionX, Y: correctionY}
	expectedC2 := curve.Add(curve.Add(messagePoint, combined), correction)
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
