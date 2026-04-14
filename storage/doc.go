// Package storage is the persistence layer for a davinci-dkg node.
//
// It exposes CRUD operations for Rounds, Contributions, PartialDecryptions,
// CombinedDecryptions and RevealedShares on top of a key/value database
// (see db/pebble for the default backend). All values are serialized with
// Go's gob encoding and keyed by the round's 12-byte identifier plus a
// short per-type prefix.
//
// The package is intentionally synchronous and does no caching: callers
// provide the context and handle concurrency. The service workers call it
// from a single goroutine per round, so there is no contention in practice.
package storage
