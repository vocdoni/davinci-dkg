// Package web3 contains the Ethereum client bindings used by davinci-dkg.
//
// It wraps go-ethereum's ethclient with strongly typed helpers for the
// DKGManager and DKGRegistry contracts and decodes Epoch/Contribution/Decryption
// state into Go structs (see types).
//
// The package is intentionally thin: it performs no business logic and no
// state caching. Higher-level orchestration lives in the node package.
package web3
