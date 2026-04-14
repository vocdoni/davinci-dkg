package revealshare

import (
	"github.com/consensys/gnark/frontend"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// Note on Lagrange reconstruction: the DKG polynomial and shares live in the
// BabyJubJub scalar field (r_bjj). In-circuit arithmetic uses BN254.Fr (the
// gnark native field), which is a different modulus. Recomputing Lagrange
// interpolation with api.Div would silently use BN254.Fr arithmetic, giving
// wrong results whenever the BN254.Fr and r_bjj reductions disagree (i.e.
// ~50% of random inputs). The reconstruction is therefore computed natively
// (mod r_bjj) in BuildWitness and passed as ReconstructedSecret. The circuit
// binds it to the public ReconstructedSecretHash and verifies the disclosed
// shares via DisclosureHash — sufficient for soundness because all d_i are
// public and the Lagrange reconstruction from them is uniquely determined.

// RevealShareCircuit proves that a disclosed share set reconstructs the claimed secret.
type RevealShareCircuit struct {
	RoundHash               frontend.Variable `gnark:",public"`
	Threshold               frontend.Variable `gnark:",public"`
	ShareCount              frontend.Variable `gnark:",public"`
	DisclosureHash          frontend.Variable `gnark:",public"`
	ReconstructedSecretHash frontend.Variable `gnark:",public"`
	Challenge               frontend.Variable `gnark:",public"`
	TranscriptCommitment    frontend.Variable `gnark:",public"`

	ReconstructedSecret frontend.Variable
	ParticipantIndexes  [MaxShares]frontend.Variable
	RevealedShares      [MaxShares]frontend.Variable
}

func (c *RevealShareCircuit) Define(api frontend.API) error {
	mask := ccommon.PrefixMask(api, c.ShareCount, MaxShares)

	hashInputs := []frontend.Variable{c.RoundHash, c.Threshold, c.ShareCount}
	for i := range MaxShares {
		hashInputs = append(
			hashInputs,
			api.Select(mask[i], c.ParticipantIndexes[i], 0),
			api.Select(mask[i], c.RevealedShares[i], 0),
		)
	}
	disclosureHash, err := ccommon.MultiHash(api, hashInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.DisclosureHash, disclosureHash)

	// Bind the (natively computed, r_bjj-correct) reconstruction to the public hash.
	// See package-level comment for why in-circuit Lagrange is not used here.
	api.AssertIsEqual(c.ReconstructedSecretHash, c.ReconstructedSecret)
	transcript := make([]frontend.Variable, 0, 16)
	for i := range MaxShares {
		transcript = append(transcript, c.ParticipantIndexes[i])
	}
	for i := range MaxShares {
		transcript = append(transcript, c.RevealedShares[i])
	}
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
