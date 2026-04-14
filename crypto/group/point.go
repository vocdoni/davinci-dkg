package group

import (
	"github.com/vocdoni/davinci-node/crypto/ecc"
	bjj "github.com/vocdoni/davinci-node/crypto/ecc/bjj_gnark"
)

// NewPoint returns the gnark-backed reduced Twisted Edwards BabyJubJub point.
func NewPoint() ecc.Point {
	return bjj.New()
}

// Generator returns the canonical generator point for the DKG group.
func Generator() ecc.Point {
	p := NewPoint()
	p.SetGenerator()
	return p
}
