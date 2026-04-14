package group

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestEncodeDecodeRoundTrip(t *testing.T) {
	c := qt.New(t)

	generator := Generator()
	encoded := Encode(generator)
	decoded, err := Decode(encoded)

	c.Assert(err, qt.IsNil)
	c.Assert(decoded.Equal(generator), qt.IsTrue)
}

func TestScalarField(t *testing.T) {
	c := qt.New(t)

	c.Assert(ScalarField().Sign() > 0, qt.IsTrue)
}
