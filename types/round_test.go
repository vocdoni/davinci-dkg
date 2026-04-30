package types

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestEpochPhaseString(t *testing.T) {
	c := qt.New(t)

	c.Assert(EpochPhaseUnknown.String(), qt.Equals, "unknown")
	c.Assert(EpochPhaseRegistration.String(), qt.Equals, "registration")
	c.Assert(EpochPhaseContribution.String(), qt.Equals, "contribution")
	c.Assert(EpochPhaseFinalized.String(), qt.Equals, "finalized")
	c.Assert(EpochPhaseDecryption.String(), qt.Equals, "decryption")
	c.Assert(EpochPhaseAborted.String(), qt.Equals, "aborted")
	c.Assert(EpochPhaseCompleted.String(), qt.Equals, "completed")
}

func TestEpochPolicyValidate(t *testing.T) {
	c := qt.New(t)

	c.Run("accepts coherent policy", func(c *qt.C) {
		policy := EpochPolicy{
			Threshold:                 3,
			CommitteeSize:             5,
			MinValidContributions:     3,
			LotteryAlphaBps:           20000,
			SeedDelay:                 4,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		}

		err := policy.Validate()

		c.Assert(err, qt.IsNil)
	})

	c.Run("rejects threshold larger than committee", func(c *qt.C) {
		policy := EpochPolicy{
			Threshold:                 6,
			CommitteeSize:             5,
			MinValidContributions:     3,
			RegistrationDeadlineBlock: 10,
			ContributionDeadlineBlock: 20,
			FinalizeNotBeforeBlock:    21,
		}

		err := policy.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "threshold")
	})

	c.Run("rejects non monotonic block windows", func(c *qt.C) {
		policy := EpochPolicy{
			Threshold:                 3,
			CommitteeSize:             5,
			MinValidContributions:     3,
			LotteryAlphaBps:           20000,
			SeedDelay:                 4,
			RegistrationDeadlineBlock: 20,
			ContributionDeadlineBlock: 10,
			FinalizeNotBeforeBlock:    11,
		}

		err := policy.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "deadline")
	})
}
