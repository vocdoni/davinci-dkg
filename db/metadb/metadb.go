package metadb

import (
	"fmt"
	"testing"

	"github.com/vocdoni/davinci-dkg/db"
	"github.com/vocdoni/davinci-dkg/db/pebble"
)

func New(typ, path string) (db.Database, error) {
	switch typ {
	case db.TypePebble:
		return pebble.New(db.Options{Path: path})
	default:
		return nil, fmt.Errorf("unsupported db type %q", typ)
	}
}

func NewTest(tb testing.TB) db.Database {
	tb.Helper()

	database, err := New(db.TypePebble, tb.TempDir())
	if err != nil {
		tb.Fatal(err)
	}
	tb.Cleanup(func() {
		_ = database.Close()
	})
	return database
}
