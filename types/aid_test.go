package types

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	qt "github.com/frankban/quicktest"
)

// bn254ScalarField is the modulus the contract bounds application ids by.
var bn254ScalarField, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617", 10)

func TestApplicationIDLayout(t *testing.T) {
	c := qt.New(t)
	registrant := common.HexToAddress("0x42fC000000000000000000000000000000Ab589F")
	aid, err := ApplicationID(registrant, big.NewInt(0x1234))
	c.Assert(err, qt.IsNil)
	c.Assert(common.Hash(aid).Hex(), qt.Equals,
		"0x000000000000000000001234"+"42fc000000000000000000000000000000ab589f")
	c.Assert(ApplicationRegistrant(aid), qt.Equals, registrant)
}

func TestApplicationIDStaysInScalarField(t *testing.T) {
	c := qt.New(t)
	maxSalt := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), AppIDSaltBits), big.NewInt(1))
	aid, err := ApplicationID(common.MaxAddress, maxSalt)
	c.Assert(err, qt.IsNil)
	c.Assert(new(big.Int).SetBytes(aid[:]).Cmp(bn254ScalarField) < 0, qt.IsTrue)

	_, err = ApplicationID(common.MaxAddress, new(big.Int).Add(maxSalt, big.NewInt(1)))
	c.Assert(err, qt.ErrorIs, ErrAppIDSalt)
	_, err = ApplicationID(common.MaxAddress, big.NewInt(-1))
	c.Assert(err, qt.ErrorIs, ErrAppIDSalt)
	_, err = ApplicationID(common.MaxAddress, nil)
	c.Assert(err, qt.ErrorIs, ErrAppIDSalt)
}

func TestRandomApplicationID(t *testing.T) {
	c := qt.New(t)
	registrant := common.HexToAddress("0x00000000000000000000000000000000000A11CE")
	a, err := RandomApplicationID(registrant)
	c.Assert(err, qt.IsNil)
	b, err := RandomApplicationID(registrant)
	c.Assert(err, qt.IsNil)
	c.Assert(a, qt.Not(qt.Equals), b)
	for _, aid := range [][32]byte{a, b} {
		c.Assert(ApplicationRegistrant(aid), qt.Equals, registrant)
		c.Assert(new(big.Int).SetBytes(aid[:]).BitLen() <= 160+AppIDSaltBits, qt.IsTrue)
	}
}
