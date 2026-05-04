package helpers

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRoundIDHelpers(t *testing.T) {
	c := qt.New(t)

	epochID := RoundIDFromString("epoch-1")

	c.Assert(RoundIDToString(epochID), qt.Equals, "epoch-1")

	overflow := RoundIDFromString("epoch-identifier-overflow")
	c.Assert(RoundIDToString(overflow), qt.Equals, "epoch-identi")
}
