package storage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// Storage is the in-memory store for DKG state.
type Storage struct {
	db            db.Database
	rounds        map[string]types.Round
	ready         map[string]map[common.Address]struct{}
	contributions map[string]map[common.Address]types.Contribution
	decryptions   map[string]map[common.Address]map[uint16]types.PartialDecryption
	disclosures   map[string]map[common.Address]types.RevealedShare
}

// New creates a new in-memory storage.
func New() *Storage {
	return NewWithDB(nil)
}

// NewWithDB creates a storage instance backed by the given database when non-nil.
func NewWithDB(database db.Database) *Storage {
	return &Storage{
		db:            database,
		rounds:        make(map[string]types.Round),
		ready:         make(map[string]map[common.Address]struct{}),
		contributions: make(map[string]map[common.Address]types.Contribution),
		decryptions:   make(map[string]map[common.Address]map[uint16]types.PartialDecryption),
		disclosures:   make(map[string]map[common.Address]types.RevealedShare),
	}
}
