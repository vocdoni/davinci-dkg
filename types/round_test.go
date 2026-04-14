package types

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRoundPhaseString(t *testing.T) {
	c := qt.New(t)

	c.Assert(RoundPhaseUnknown.String(), qt.Equals, "unknown")
	c.Assert(RoundPhaseRegistration.String(), qt.Equals, "registration")
	c.Assert(RoundPhaseContribution.String(), qt.Equals, "contribution")
	c.Assert(RoundPhaseFinalized.String(), qt.Equals, "finalized")
	c.Assert(RoundPhaseDecryption.String(), qt.Equals, "decryption")
	c.Assert(RoundPhaseDisclosure.String(), qt.Equals, "disclosure")
	c.Assert(RoundPhaseAborted.String(), qt.Equals, "aborted")
	c.Assert(RoundPhaseCompleted.String(), qt.Equals, "completed")
}

func TestRoundPolicyValidate(t *testing.T) {
	c := qt.New(t)

	c.Run("accepts coherent policy", func(c *qt.C) {
		policy := RoundPolicy{
			Threshold:                 3,
			CommitteeSize:             5,
			MinValidContributions:     3,
			LotteryAlphaBps:           20000,
			SeedDelay:                 4,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			DisclosureAllowed:         true,
		}

		err := policy.Validate()

		c.Assert(err, qt.IsNil)
	})

	c.Run("rejects threshold larger than committee", func(c *qt.C) {
		policy := RoundPolicy{
			Threshold:                 6,
			CommitteeSize:             5,
			MinValidContributions:     3,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			DisclosureAllowed:         true,
		}

		err := policy.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "threshold")
	})

	c.Run("rejects non monotonic block windows", func(c *qt.C) {
		policy := RoundPolicy{
			Threshold:                 3,
			CommitteeSize:             5,
			MinValidContributions:     3,
			LotteryAlphaBps:           20000,
			SeedDelay:                 4,
			RegistrationDeadlineBlock: 20,
			ContributionDeadlineBlock: 10,
			DisclosureAllowed:         true,
		}

		err := policy.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "deadline")
	})
}
