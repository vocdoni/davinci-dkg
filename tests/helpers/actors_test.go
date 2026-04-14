package helpers

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestDefaultAnvilPrivateKeys(t *testing.T) {
	c := qt.New(t)

	c.Assert(len(DefaultAnvilPrivateKeys) >= 3, qt.IsTrue)
	c.Assert(DefaultAnvilPrivateKeys[0], qt.Equals, LocalAccountPrivKey)
}
