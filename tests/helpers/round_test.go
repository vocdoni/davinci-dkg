package helpers

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestRoundIDHelpers(t *testing.T) {
	c := qt.New(t)

	roundID := RoundIDFromString("round-1")

	c.Assert(RoundIDToString(roundID), qt.Equals, "round-1")

	overflow := RoundIDFromString("round-identifier-overflow")
	c.Assert(RoundIDToString(overflow), qt.Equals, "round-identi")
}
