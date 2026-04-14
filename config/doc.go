// Package config holds build-time constants that tie the Go implementation
// to the compiled circuit artifacts and the deployed Solidity contracts.
//
// The most load-bearing value is the set of artifact hashes in
// circuit_artifacts.go: every circuit's proving key, verifying key and
// constraint system is referenced by SHA-256 digest, which the prover
// package verifies on load. Changing a circuit without regenerating these
// hashes causes a clear startup error rather than silently running against
// stale artifacts.
//
// Run `make circuits-update-hashes` after recompiling circuits to keep
// this file in sync.
package config
