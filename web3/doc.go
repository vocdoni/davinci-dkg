// Package web3 contains the Ethereum client bindings used by davinci-dkg.
//
// It wraps go-ethereum's ethclient with strongly typed helpers for the
// DKGManager and DKGRegistry contracts, decodes Round/Contribution/Decryption
// state into Go structs (see types), and exposes BRLC transcript encoders
// that mirror the on-chain Solidity format bit-for-bit.
//
// The package is intentionally thin: it performs no business logic and no
// state caching. Higher-level orchestration lives in cmd/davinci-dkg-node
// and cmd/dkg-runner.
package web3
