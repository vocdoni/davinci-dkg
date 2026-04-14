// Package hash collects the Poseidon1 hashing helpers and the protocol
// domain separators used by davinci-dkg.
//
// Every hash computed off-chain must reproduce the in-circuit hash bit for
// bit, so domain separators are centralized here and consumed both by the
// Go crypto helpers and by the circuit witness builders. See the
// DomainShareEncryption, DomainPartialDecrypt and DomainRoundSelection
// constants for the active tags.
package hash
