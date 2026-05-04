package types

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestEpochPhaseString(t *testing.T) {
	c := qt.New(t)

	c.Assert(EpochPhaseUnknown.String(), qt.Equals, "unknown")
	c.Assert(EpochPhaseCommitteeSelection.String(), qt.Equals, "registration")
	c.Assert(EpochPhaseKeyAssembly.String(), qt.Equals, "contribution")
	c.Assert(EpochPhaseLive.String(), qt.Equals, "finalized")
	c.Assert(EpochPhaseDecryption.String(), qt.Equals, "decryption")
	c.Assert(EpochPhaseAborted.String(), qt.Equals, "aborted")
	c.Assert(EpochPhaseCompleted.String(), qt.Equals, "completed")
}

func TestEpochPolicyValidate(t *testing.T) {
	c := qt.New(t)

	c.Run("accepts coherent policy", func(c *qt.C) {
		policy := EpochPolicy{
			Threshold:                       3,
			CommitteeSize:                   5,
			MinValidContributions:           3,
			LotteryAlphaBps:                 20000,
			CommitteeSelectionDeadlineBlock: 10,
			KeyAssemblyDeadlineBlock:        20,
			LiveNotBeforeBlock:              21,
		}

		err := policy.Validate()

		c.Assert(err, qt.IsNil)
	})

	c.Run("rejects threshold larger than committee", func(c *qt.C) {
		policy := EpochPolicy{
			Threshold:                       6,
			CommitteeSize:                   5,
			MinValidContributions:           3,
			CommitteeSelectionDeadlineBlock: 10,
			KeyAssemblyDeadlineBlock:        20,
			LiveNotBeforeBlock:              21,
		}

		err := policy.Validate()

		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "threshold")
	})

	// The "rejects non monotonic block windows" test was removed when phase
	// deadlines moved from caller-supplied to contract-derived (from
	// EPOCH_DURATION_BLOCKS). The Go-side policy struct no longer validates
	// deadline ordering; the on-chain constructor enforces it via the
	// derived per-phase offsets.

	c.Run("rejects lottery alpha below 1.0", func(c *qt.C) {
		policy := EpochPolicy{
			Threshold:             3,
			CommitteeSize:         5,
			MinValidContributions: 3,
			LotteryAlphaBps:       9999,
		}
		err := policy.Validate()
		c.Assert(err, qt.Not(qt.IsNil))
		c.Assert(err.Error(), qt.Contains, "alpha")
	})
}
