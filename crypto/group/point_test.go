package group

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestNewPoint(t *testing.T) {
	c := qt.New(t)

	point := NewPoint()
	x, y := point.Point()

	c.Assert(point.Order().Sign() > 0, qt.IsTrue)
	c.Assert(x, qt.Not(qt.IsNil))
	c.Assert(y, qt.Not(qt.IsNil))
}

func TestGeneratorIsStable(t *testing.T) {
	c := qt.New(t)

	g1 := Generator()
	g2 := Generator()

	c.Assert(g1.Equal(g2), qt.IsTrue)
}
