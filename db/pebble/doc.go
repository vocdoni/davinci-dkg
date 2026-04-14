// Package pebble is the default key/value database backend for davinci-dkg.
//
// It adapts CockroachDB's pebble LSM to the minimal db.Database interface
// used by the storage package. Pebble is chosen for its fast single-writer
// performance and small operational footprint — a DKG node's working set
// is tiny compared to a general-purpose database, and pebble fits it in
// one directory with no background services.
package pebble
