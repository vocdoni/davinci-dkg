// Package types contains the cross-package data structures used by the
// davinci-dkg node: Round, Contribution, PartialDecryption, RevealedShare,
// NodeKey, ContractAddresses and the supporting enums.
//
// These types are the lingua franca between the chain bindings in web3,
// the persistence layer in storage, and the phase workers in service.
// Keeping them in a neutral package avoids import cycles.
package types
