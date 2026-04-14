// Package group contains BabyJubJub curve helpers shared between native Go
// code (the node services and circuit assignments) and the circuit witness
// builders.
//
// BabyJubJub is the twisted Edwards curve over the BN254 scalar field used
// by every davinci-dkg circuit. This package exposes scalar/point encoding
// helpers, fixed-base scalar multiplication, and the Poseidon-friendly
// point serialization used by the transcript builders.
package group
