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
// threshold set of committee partial decryptions plus the organizer's
// share Δ_org = sk_org·C_1:
//
//	C_2 = m·G + Σ λ_k · δ_{x_k} + Δ_org
//
// Δ_org is not calldata any more: the organizer secret is a private witness
// and the circuit proves knowledge of it —
//
//	PK_org == sk_org·G   and   Δ_org == sk_org·C_1
//
// — for the PK_org the contract pinned to the application record. An
// organizer-locked application reveals sk_org once (`revealOrganizerSecret`)
// and the committee combines by itself from then on; an automatic
// application never had a key and uses the identity PK_org = (0, 1) with
// sk_org = 0, which satisfies both relations because 0·P is the identity
// for any P. See docs/pool-keys.md, "Combine".
//
// Public inputs, in declaration order, are the 9 words the verifier reads:
//
//	RoundHash, Aid, CtIdx, Threshold, ShareCount, CombineHash,
//	PlaintextHash, Challenge, TranscriptCommitment
type DecryptCombineCircuit struct {
	RoundHash            frontend.Variable `gnark:",public"` // semantically: eid
	Aid                  frontend.Variable `gnark:",public"` // application identifier
	CtIdx                frontend.Variable `gnark:",public"` // per-app ciphertext index
	Threshold            frontend.Variable `gnark:",public"`
	ShareCount           frontend.Variable `gnark:",public"`
	CombineHash          frontend.Variable `gnark:",public"`
	PlaintextHash        frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	CiphertextC1 twistededwards.Point
	CiphertextC2 twistededwards.Point
	Plaintext    frontend.Variable
	// OrganizerPK is folded into CombineHash and into the BRLC transcript,
	// so the contract binds it to the application's registered key.
	// OrganizerSecret is its discrete log and never leaves the witness.
	OrganizerPK        twistededwards.Point
	OrganizerSecret    frontend.Variable
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
	// OrganizerPK needs no on-curve assertion: the equality below makes it
	// the output of FixedBaseMul, which is on the curve by construction.
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

	// The organizer secret must be a canonical BJJ scalar. Without the bound
	// a prover could offer sk' = sk + r_bjj, which multiplies to the same two
	// points but is a second witness for the same statement.
	api.AssertIsLessOrEqual(c.OrganizerSecret, ccommon.SubgroupOrderMinusOne())

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
		c.Threshold,
		c.ShareCount,
		c.CiphertextC1.X,
		c.CiphertextC1.Y,
		c.CiphertextC2.X,
		c.CiphertextC2.Y,
		c.OrganizerPK.X,
		c.OrganizerPK.Y,
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

	// Knowledge of sk_org for the application's registered key, and the
	// share it determines. Both are needed: PK_org == sk_org·G alone would
	// let a combiner pick any Δ_org and shift the recovered plaintext, and
	// Δ_org == sk_org·C_1 alone would not tie the secret to the application.
	ccommon.AssertPointEqual(api, c.OrganizerPK, ccommon.FixedBaseMul(api, c.OrganizerSecret))
	deltaOrg := ccommon.ScalarMulVar(api, c.CiphertextC1, c.OrganizerSecret)

	// The Lagrange interpolation accumulator. For inactive slots we mask the
	// scalar to 0 so curve.ScalarMul yields the identity, then unconditionally
	// add — saving the per-iteration result Select compared to the original.
	combined := ccommon.IdentityPoint()
	lambdaG := make([]twistededwards.Point, MaxShares)
	for i := range MaxShares {
		lambda := api.Mul(mask[i], c.LagrangeCoefficients[i])
		scaled := ccommon.ScalarMulVar(api, c.PartialDecryptions[i], lambda)
		combined = curve.Add(combined, scaled)
		lambdaG[i] = ccommon.FixedBaseMul(api, lambda)
	}
	// Pin λ to the canonical Lagrange vector at 0 of the qualifying set.
	// The decryption identity alone is not enough: a prover who knows
	// log_G δ_k for a single k (the encryptor colluding with one share
	// holder) can open the ciphertext to any plaintext by shifting λ_k.
	// The canonical vector is the unique solution of the Vandermonde system
	//   Σ_k λ_k · x_k^j = [j = 0]   for j = 0 … s−1   (mod r_bjj),
	// s = ShareCount, which we check on points: FixedBaseMul reduces λ_k
	// modulo r_bjj and x_k ≤ MaxShares keeps every step a 6-bit scalar
	// multiplication. The system must run to ShareCount, not Threshold:
	// with s > t shares the first t equations leave an (s − t)-dimensional
	// family of solutions, so only the square system pins λ to the one
	// canonical vector and makes the witness unique.
	powerMask := ccommon.PrefixMask(api, c.ShareCount, MaxShares)
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
	// C_2 = m·G + Σ λ_k·δ_k + Δ_org.
	messagePoint := ccommon.FixedBaseMul(api, c.Plaintext)
	expectedC2 := curve.Add(curve.Add(messagePoint, combined), deltaOrg)
	ccommon.AssertPointEqual(api, expectedC2, c.CiphertextC2)
	// The transcript uses the same masked values as CombineHash so that no
	// witness word outside the digest can be tuned after ρ is known. Order
	// is fixed by the contract's checks: ciphertext, organizer key, then
	// the indexes and the partials.
	transcript := make([]frontend.Variable, 0, TranscriptWords)
	transcript = append(
		transcript,
		c.CiphertextC1.X, c.CiphertextC1.Y, c.CiphertextC2.X, c.CiphertextC2.Y,
		c.OrganizerPK.X, c.OrganizerPK.Y,
	)
	transcript = append(transcript, maskedIndexes...)
	for i := range MaxShares {
		transcript = append(transcript, maskedPartials[i].X, maskedPartials[i].Y)
	}
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
