package decryptcombine

import (
	"math/big"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
	"github.com/vocdoni/davinci-dkg/crypto/group"
)

// indexBits bounds a participant index (≤ MaxShares = 32) for the
// small-scalar multiplications of the Vandermonde check.
const indexBits = 6

// generator is G in the coordinates FixedBaseMul produces, so the
// Vandermonde identity below can compare against 1·G as constants.
var generator = func() twistededwards.Point {
	g := group.NewPoint()
	g.ScalarBaseMult(big.NewInt(1))
	enc := group.Encode(g)
	return twistededwards.Point{X: enc.X, Y: enc.Y}
}()

// DecryptCombineCircuit reconstructs the per-application plaintext from a
// threshold set of partial decryptions, with per-application correction
// term `T` selected by the `Mode` flag (paper §5.5 lines 1051–1077):
//
//	mode = 0 (public derivation): T = S · C_1 (computed in-circuit)
//	mode = 1 (organizer co-decryption): T = Δ_org (consumed as point)
//
//	M = C_2 - Σ λ_k · δ_{x_k} - T
type DecryptCombineCircuit struct {
	RoundHash            frontend.Variable    `gnark:",public"` // semantically: eid
	Aid                  frontend.Variable    `gnark:",public"` // application identifier
	CtIdx                frontend.Variable    `gnark:",public"` // per-app ciphertext index
	Mode                 frontend.Variable    `gnark:",public"` // 0 = public derivation, 1 = co-decryption
	S                    frontend.Variable    `gnark:",public"` // derivation tag scalar; ignored if Mode=1
	DeltaOrg             twistededwards.Point `gnark:",public"` // organizer Δ_org; ignored if Mode=0
	Threshold            frontend.Variable    `gnark:",public"`
	ShareCount           frontend.Variable    `gnark:",public"`
	CombineHash          frontend.Variable    `gnark:",public"`
	PlaintextHash        frontend.Variable    `gnark:",public"`
	Challenge            frontend.Variable    `gnark:",public"`
	TranscriptCommitment frontend.Variable    `gnark:",public"`

	CiphertextC1       twistededwards.Point
	CiphertextC2       twistededwards.Point
	Plaintext          frontend.Variable
	ParticipantIndexes [MaxShares]frontend.Variable
	PartialDecryptions [MaxShares]twistededwards.Point
	// LagrangeCoefficients are pre-computed natively in the BJJ scalar field
	// (r_bjj) and passed as private witnesses; api.Div would compute them in
	// BN254.Fr, which is the wrong field. Define pins them to the canonical
	// vector through the Vandermonde identity, so soundness does not depend
	// on who is proving (see the comment in Define).
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
	api.AssertIsLessOrEqual(c.Threshold, c.ShareCount)
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

	maskedIndexes := make([]frontend.Variable, MaxShares)
	maskedPartials := make([]twistededwards.Point, MaxShares)
	for i := range MaxShares {
		maskedIndexes[i] = api.Select(mask[i], c.ParticipantIndexes[i], 0)
		maskedPartials[i] = ccommon.MaskPoint(api, mask[i], c.PartialDecryptions[i])
	}
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
		hashInputs = append(hashInputs, maskedIndexes[i], maskedPartials[i].X, maskedPartials[i].Y)
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
	lambdaG := make([]twistededwards.Point, MaxShares)
	for i := range MaxShares {
		lambda := api.Mul(mask[i], c.LagrangeCoefficients[i])
		scaled := curve.ScalarMul(c.PartialDecryptions[i], lambda)
		combined = curve.Add(combined, scaled)
		lambdaG[i] = ccommon.FixedBaseMul(api, lambda)
	}
	// Pin λ to the canonical Lagrange vector at 0 of the qualifying set.
	// The decryption identity alone is not enough: a prover who knows
	// log_G δ_k for a single k (the encryptor colluding with one share
	// holder) can open the ciphertext to any plaintext by shifting λ_k.
	// The canonical vector is the unique solution of the Vandermonde system
	//   Σ_k λ_k · x_k^j = [j = 0]   for j = 0 … t−1   (mod r_bjj),
	// which we check on points: FixedBaseMul reduces λ_k modulo r_bjj and
	// x_k ≤ MaxShares keeps every step a 6-bit scalar multiplication.
	// Checking j < Threshold suffices because the shared polynomial has
	// degree < t, so any λ passing it interpolates F(0) exactly.
	powerMask := ccommon.PrefixMask(api, c.Threshold, MaxShares)
	for j := range MaxShares {
		sum := ccommon.IdentityPoint()
		for i := range MaxShares {
			sum = curve.Add(sum, lambdaG[i])
		}
		expected := ccommon.IdentityPoint()
		if j == 0 {
			expected = generator
		}
		api.AssertIsEqual(api.Mul(powerMask[j], api.Sub(sum.X, expected.X)), 0)
		api.AssertIsEqual(api.Mul(powerMask[j], api.Sub(sum.Y, expected.Y)), 0)
		if j+1 < MaxShares {
			for i := range MaxShares {
				lambdaG[i] = ccommon.ScalarMulSmallScalar(api, lambdaG[i], c.ParticipantIndexes[i], indexBits)
			}
		}
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
	// The transcript uses the same masked values as CombineHash so that no
	// witness word outside the digest can be tuned after ρ is known.
	transcript := make([]frontend.Variable, 0, 4+3*MaxShares)
	transcript = append(transcript, c.CiphertextC1.X, c.CiphertextC1.Y, c.CiphertextC2.X, c.CiphertextC2.Y)
	transcript = append(transcript, maskedIndexes...)
	for i := range MaxShares {
		transcript = append(transcript, maskedPartials[i].X, maskedPartials[i].Y)
	}
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
