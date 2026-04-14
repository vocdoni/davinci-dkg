package finalize

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/native/twistededwards"
	ccommon "github.com/vocdoni/davinci-dkg/circuits/common"
)

// FinalizeCircuit proves the phase-5 statement from the paper over commitment
// points: accepted commitment vectors aggregate into Cbar(k), Cbar(0) is the
// public key, and D_i = sum_k i^k * Cbar(k). The wide transcript is compressed
// with a BRLC commitment to keep verifier public inputs constant-sized.
type FinalizeCircuit struct {
	RoundHash            frontend.Variable `gnark:",public"`
	Threshold            frontend.Variable `gnark:",public"`
	CommitteeSize        frontend.Variable `gnark:",public"`
	AcceptedCount        frontend.Variable `gnark:",public"`
	AggregateHash        frontend.Variable `gnark:",public"`
	CollectivePublicKey  frontend.Variable `gnark:",public"`
	ShareCommitmentHash  frontend.Variable `gnark:",public"`
	Challenge            frontend.Variable `gnark:",public"`
	TranscriptCommitment frontend.Variable `gnark:",public"`

	ParticipantIndexes      [MaxParticipants]frontend.Variable
	ContributionCommitments [MaxParticipants][MaxCoefficients]twistededwards.Point
	AggregateCommitments    [MaxCoefficients]twistededwards.Point
	ShareCommitments        [MaxParticipants]twistededwards.Point
}

func (c *FinalizeCircuit) Define(api frontend.API) error {
	coeffMask := ccommon.PrefixMask(api, c.Threshold, MaxCoefficients)
	participantMask := ccommon.PrefixMask(api, c.AcceptedCount, MaxParticipants)

	for k := range MaxCoefficients {
		for i := range MaxParticipants {
			if err := ccommon.AssertPointOnCurve(api, c.ContributionCommitments[i][k]); err != nil {
				return err
			}
		}
		if err := ccommon.AssertPointOnCurve(api, c.AggregateCommitments[k]); err != nil {
			return err
		}
		sum := ccommon.IdentityPoint()
		for i := range MaxParticipants {
			next := ccommon.AddPointIfEnabled(api, sum, c.ContributionCommitments[i][k], participantMask[i])
			sum.X = next.X
			sum.Y = next.Y
		}
		// Conditional equality: when coeffMask[k] == 1 the aggregate commitment
		// must equal the running sum; when coeffMask[k] == 0 the constraint is
		// trivially satisfied. This replaces two SelectPoint + AssertPointEqual
		// (~6 constraints) with two muls + two asserts (~4 constraints).
		diffX := api.Sub(c.AggregateCommitments[k].X, sum.X)
		diffY := api.Sub(c.AggregateCommitments[k].Y, sum.Y)
		api.AssertIsEqual(api.Mul(coeffMask[k], diffX), 0)
		api.AssertIsEqual(api.Mul(coeffMask[k], diffY), 0)
	}

	publicKeyHash, err := ccommon.MultiHash(api, c.RoundHash, c.AggregateCommitments[0].X, c.AggregateCommitments[0].Y)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.CollectivePublicKey, publicKeyHash)

	aggregateInputs := []frontend.Variable{c.RoundHash, c.Threshold, c.CommitteeSize, c.AcceptedCount}
	for k := range MaxCoefficients {
		aggregateInputs = append(
			aggregateInputs,
			api.Select(coeffMask[k], c.AggregateCommitments[k].X, 0),
			api.Select(coeffMask[k], c.AggregateCommitments[k].Y, 1),
		)
	}
	aggregateHash, err := ccommon.MultiHash(api, aggregateInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.AggregateHash, aggregateHash)

	// Precompute masked aggregate commitments once, so that the per-participant
	// CommitmentPolynomialValue call below can iterate without repeating the
	// per-coefficient mask Select. Unused coefficient slots are replaced with the
	// twisted-Edwards identity point (0, 1) so a subsequent unconditional Add is a
	// no-op for those slots.
	maskedAggregate := make([]twistededwards.Point, MaxCoefficients)
	for k := range MaxCoefficients {
		maskedAggregate[k] = twistededwards.Point{
			X: api.Mul(coeffMask[k], c.AggregateCommitments[k].X),
			Y: api.Select(coeffMask[k], c.AggregateCommitments[k].Y, 1),
		}
	}

	shareInputs := []frontend.Variable{c.RoundHash, c.Threshold, c.CommitteeSize, c.AcceptedCount}
	for i := range MaxParticipants {
		if err := ccommon.AssertPointOnCurve(api, c.ShareCommitments[i]); err != nil {
			return err
		}
		shareCommitment, err := ccommon.CommitmentPolynomialValue(
			api,
			maskedAggregate,
			nil, // mask already baked in
			c.ParticipantIndexes[i],
		)
		if err != nil {
			return err
		}
		// Conditional equality (see comment on the aggregate-vs-sum check above):
		// when participantMask[i] == 1 we must have ShareCommitments[i] ==
		// shareCommitment; otherwise the constraint is trivially satisfied.
		dShareX := api.Sub(c.ShareCommitments[i].X, shareCommitment.X)
		dShareY := api.Sub(c.ShareCommitments[i].Y, shareCommitment.Y)
		api.AssertIsEqual(api.Mul(participantMask[i], dShareX), 0)
		api.AssertIsEqual(api.Mul(participantMask[i], dShareY), 0)
		shareInputs = append(
			shareInputs,
			api.Select(participantMask[i], c.ParticipantIndexes[i], 0),
			api.Select(participantMask[i], c.ShareCommitments[i].X, 0),
			api.Select(participantMask[i], c.ShareCommitments[i].Y, 1),
		)
	}
	shareHash, err := ccommon.MultiHash(api, shareInputs...)
	if err != nil {
		return err
	}
	api.AssertIsEqual(c.ShareCommitmentHash, shareHash)

	transcript := make([]frontend.Variable, 0, 168)
	for i := range MaxParticipants {
		transcript = append(transcript, api.Select(participantMask[i], c.ParticipantIndexes[i], 0))
	}
	for i := range MaxParticipants {
		for k := range MaxCoefficients {
			transcript = append(
				transcript,
				api.Select(participantMask[i], c.ContributionCommitments[i][k].X, 0),
				api.Select(participantMask[i], c.ContributionCommitments[i][k].Y, 1),
			)
		}
	}
	for k := range MaxCoefficients {
		transcript = append(
			transcript,
			api.Select(coeffMask[k], c.AggregateCommitments[k].X, 0),
			api.Select(coeffMask[k], c.AggregateCommitments[k].Y, 1),
		)
	}
	for i := range MaxParticipants {
		transcript = append(
			transcript,
			api.Select(participantMask[i], c.ShareCommitments[i].X, 0),
			api.Select(participantMask[i], c.ShareCommitments[i].Y, 1),
		)
	}
	api.AssertIsEqual(c.TranscriptCommitment, ccommon.BRLC(api, c.Challenge, transcript))
	return nil
}
