package hash

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestHashFieldElementsIsDeterministic(t *testing.T) {
	c := qt.New(t)

	got1, err := HashFieldElements(big.NewInt(1), big.NewInt(2), big.NewInt(3))
	c.Assert(err, qt.IsNil)

	got2, err := HashFieldElements(big.NewInt(1), big.NewInt(2), big.NewInt(3))
	c.Assert(err, qt.IsNil)

	c.Assert(got1.Cmp(got2), qt.Equals, 0)
}

func TestDomainValue(t *testing.T) {
	c := qt.New(t)

	value := DomainValue([]byte("davinci-dkg"))

	c.Assert(value.Sign() >= 0, qt.IsTrue)
}
