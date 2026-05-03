package storage

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/types"
)

// Storage is the in-memory store for DKG state.
type Storage struct {
	db            db.Database
	epochs        map[string]types.Epoch
	ready         map[string]map[common.Address]struct{}
	contributions map[string]map[common.Address]types.Contribution
	decryptions   map[string]map[partialMemKey]types.PartialDecryption
}

// partialMemKey is the in-memory composite key for partial decryptions. It
// matches the on-disk keying scheme so that committee + organizer shares for
// the same (epoch, ciphertextIndex) on different applications coexist
// without colliding.
type partialMemKey struct {
	AID             [32]byte
	Role            types.Role
	Participant     common.Address
	CiphertextIndex uint16
}

// New creates a new in-memory storage.
func New() *Storage {
	return NewWithDB(nil)
}

// NewWithDB creates a storage instance backed by the given database when non-nil.
func NewWithDB(database db.Database) *Storage {
	return &Storage{
		db:            database,
		epochs:        make(map[string]types.Epoch),
		ready:         make(map[string]map[common.Address]struct{}),
		contributions: make(map[string]map[common.Address]types.Contribution),
		decryptions:   make(map[string]map[partialMemKey]types.PartialDecryption),
	}
}
