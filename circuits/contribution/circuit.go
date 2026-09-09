package contribution

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// MaxCoefficients is `circuits/common.MaxT`, MaxRecipients is
// `circuits/common.MaxN` and MaxKeys is `circuits/common.MaxK`, the number of
// pool keys every epoch deals. Edit `circuits/common/sizes.go` to change them.
const (
	MaxCoefficients = ccommon.MaxT
	MaxRecipients   = ccommon.MaxN
	MaxKeys         = ccommon.MaxK
)

// ContributionCircuit proves the DKG dealing statement for the whole pool at
// once: MaxK polynomials of degree below t, given by their coefficient
// commitments, Feldman consistency of every recipient's share with them, and
// hashed share encryption under one ECDH secret per recipient with the key
// index separating the MaxK masks derived from it.
//
// Only the constant term of each polynomial is a witness: its commitment
// C_{j,0} = a_{j,0}·G pins the pool-key contribution to the prime subgroup.
// The higher coefficients are not needed — any t of the n ≥ t consistent
// shares proven below interpolate the polynomial, so knowledge of the shares
// is knowledge of the coefficients — but their commitments must lie in the
// prime subgroup for that argument to speak about discrete logarithms (a
// cofactor component in C_1 and C_2 that cancels at every integer position
// would otherwise pass every Feldman check), so each one comes with a
// cofactor preimage: a point Q with 8·Q = C. See docs/pool-keys.md.
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

	ConstantTerms       [MaxKeys]frontend.Variable
	CommitmentPreimages [MaxKeys][MaxCoefficients - 1]twistededwards.Point
	EncryptionNonces    [MaxRecipients]frontend.Variable
	RecipientIndexes    [MaxRecipients]frontend.Variable
	Shares              [MaxKeys][MaxRecipients]frontend.Variable
}

func (c *ContributionCircuit) Define(api frontend.API) error {
	// One dedicated left-hand row per public input certifies the QAP hypothesis
	// of weak simulation-extractability (ccommon.CertifyPublicInputs).
	ccommon.CertifyPublicInputs(api, c.RoundHash, c.Threshold, c.CommitteeSize, c.ContributorIndex, c.CommitmentHash, c.ShareHash, c.Challenge, c.TranscriptCommitment)

	// Bound the public count inputs to their fixed array sizes.
	// PrefixMask returns all-active when count > size, so
	// without these the statement could prove a partial set while
	// claiming a larger one.
	// 1 ≤ t ≤ n ≤ MaxN and 1 ≤ contributorIndex ≤ n (docs/pool-keys.md
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

	// Commitments, masked to the identity beyond the threshold so the
	// per-recipient Horner evaluations below need no per-coefficient select.
	// Slot 0 is proven as a·G from the canonical constant term; every other
	// slot carries a cofactor preimage certifying its subgroup membership.
	var maskedCommitments [MaxKeys][MaxCoefficients]twistededwards.Point
	keyDigests := make([]frontend.Variable, MaxKeys)
	var err error
	for j := range MaxKeys {
		for m := range MaxCoefficients {
			if err := ccommon.AssertPointOnCurve(api, c.Commitments[j][m]); err != nil {
				return err
			}
			maskedCommitments[j][m] = ccommon.MaskPoint(api, coeffMask[m], c.Commitments[j][m])
			if m == 0 {
				// t ≥ 1, so slot 0 is always active.
				bs := ccommon.CanonicalScalarBits(api, c.ConstantTerms[j])
				ccommon.AssertPointEqual(api, c.Commitments[j][0], ccommon.FixedBaseMulBits(api, bs))
				continue
			}
			if err := ccommon.AssertCofactorPreimage(api, c.CommitmentPreimages[j][m-1], maskedCommitments[j][m]); err != nil {
				return err
			}
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
		// Slot i of the committee snapshot is member i+1: the contract checks
		// the transcript's index words against the snapshot, and the circuit
		// pins the witness index to the same constant, so Horner below can
		// evaluate at a constant and the transcript keeps its index words.
		position := i + 1
		api.AssertIsEqual(api.Mul(recipientMask[i], api.Sub(c.RecipientIndexes[i], position)), 0)

		// One decomposition of the nonce serves both R_i = r_i·G and the
		// ECDH secret r_i·PK_i. Ephemerals[i] needs no on-curve check: when
		// active it must equal the fixed-base product, on-curve by
		// construction; when inactive it is masked out of every digest and
		// the transcript.
		nonceBits := api.ToBinary(c.EncryptionNonces[i], 254)
		expectedEphemeral := ccommon.FixedBaseMulBits(api, nonceBits)
		api.AssertIsEqual(api.Mul(recipientMask[i], api.Sub(c.Ephemerals[i].X, expectedEphemeral.X)), 0)
		api.AssertIsEqual(api.Mul(recipientMask[i], api.Sub(c.Ephemerals[i].Y, expectedEphemeral.Y)), 0)
		// One ECDH secret per recipient, shared by all MaxK keys: the whole
		// point of dealing the pool in one proof. The key index enters the
		// mask expansion below so the MaxK masks stay independent — reusing
		// one mask would be a one-time-pad reuse across keys.
		sharedSecret := ccommon.ScalarMulVarBits(api, c.RecipientPubKeys[i], nonceBits)
		seed, err := ccommon.ShareMaskSeed(api, c.RoundHash, c.ContributorIndex, position, sharedSecret.X, sharedSecret.Y)
		if err != nil {
			return err
		}

		// Every transcript word must be fixed by a digest before the BRLC
		// challenge exists (the contract derives ρ from the digests and the
		// calldata), so the share digest also absorbs the recipient keys and
		// every vector is masked to constants in inactive slots.
		maskedIndexes[i] = api.Mul(recipientMask[i], position)
		maskedKeys[i] = ccommon.MaskPoint(api, recipientMask[i], c.RecipientPubKeys[i])
		maskedEphemerals[i] = ccommon.MaskPoint(api, recipientMask[i], c.Ephemerals[i])

		rowShares := make([]frontend.Variable, MaxKeys)
		for j := range MaxKeys {
			// Shares[j][i] must be canonical, s < r: otherwise s' = s + k·r
			// passes the Feldman check below (same point) and the recipient
			// recovers a value it cannot use as its share, a DoS on the
			// epoch. The same bits then drive the fixed-base product.
			shareBits := ccommon.CanonicalScalarBits(api, c.Shares[j][i])
			sharePoint := ccommon.FixedBaseMulBits(api, shareBits)
			feldmanPoint, err := ccommon.CommitmentPolynomialValue(api, maskedCommitments[j][:], position)
			if err != nil {
				return err
			}
			api.AssertIsEqual(api.Mul(recipientMask[i], api.Sub(sharePoint.X, feldmanPoint.X)), 0)
			api.AssertIsEqual(api.Mul(recipientMask[i], api.Sub(sharePoint.Y, feldmanPoint.Y)), 0)

			// Hashed ElGamal in the native field: the mask is a Poseidon
			// output, uniform in F_p, so masked = s + mask (mod p) is a
			// one-time pad and the recipient recovers exactly the canonical
			// s. Inactive slots publish 0.
			mask, err := ccommon.ShareMask(api, seed, j)
			if err != nil {
				return err
			}
			api.AssertIsEqual(
				api.Mul(recipientMask[i], api.Sub(c.MaskedShares[j][i], api.Add(c.Shares[j][i], mask))), 0,
			)
			maskedShares[j][i] = api.Mul(recipientMask[i], c.MaskedShares[j][i])
			rowShares[j] = maskedShares[j][i]
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

	// Compact BRLC (docs/pool-keys.md §4): the same fixed-size region order
	// as before, but every word is gated by the public counts — a
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
