package pebble

import (
	"bytes"
	"errors"
	"os"

	"github.com/cockroachdb/pebble"
	"github.com/vocdoni/davinci-dkg/db"
)

type Database struct {
	db *pebble.DB
}

type WriteTx struct {
	batch *pebble.Batch
}

var (
	_ db.Database = (*Database)(nil)
	_ db.WriteTx  = (*WriteTx)(nil)
)

func New(opts db.Options) (*Database, error) {
	if err := os.MkdirAll(opts.Path, os.ModePerm); err != nil {
		return nil, err
	}
	instance, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return nil, err
	}
	return &Database{db: instance}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) Get(key []byte) ([]byte, error) {
	value, closer, err := d.db.Get(key)
	if errors.Is(err, pebble.ErrNotFound) {
		return nil, db.ErrKeyNotFound
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = closer.Close() }()
	return bytes.Clone(value), nil
}

func (d *Database) Iterate(prefix []byte, callback func(key, value []byte) bool) error {
	iter, err := d.db.NewIter(&pebble.IterOptions{
		LowerBound: prefix,
		UpperBound: keyUpperBound(prefix),
	})
	if err != nil {
		return err
	}
	defer func() { _ = iter.Close() }()
	for iter.First(); iter.Valid(); iter.Next() {
		if !callback(bytes.Clone(iter.Key()), bytes.Clone(iter.Value())) {
			break
		}
	}
	return iter.Error()
}

func (d *Database) WriteTx() db.WriteTx {
	return &WriteTx{batch: d.db.NewIndexedBatch()}
}

func (d *Database) Compact() error {
	return d.db.Compact(nil, []byte{0xff}, true)
}

func (tx *WriteTx) Get(key []byte) ([]byte, error) {
	value, closer, err := tx.batch.Get(key)
	if errors.Is(err, pebble.ErrNotFound) {
		return nil, db.ErrKeyNotFound
	}
	if err != nil {
		return nil, err
	}
	defer func() { _ = closer.Close() }()
	return bytes.Clone(value), nil
}

func (tx *WriteTx) Iterate(prefix []byte, callback func(key, value []byte) bool) error {
	iter, err := tx.batch.NewIter(&pebble.IterOptions{
		LowerBound: prefix,
		UpperBound: keyUpperBound(prefix),
	})
	if err != nil {
		return err
	}
	defer func() { _ = iter.Close() }()
	for iter.First(); iter.Valid(); iter.Next() {
		if !callback(bytes.Clone(iter.Key()), bytes.Clone(iter.Value())) {
			break
		}
	}
	return iter.Error()
}

func (tx *WriteTx) Set(key, value []byte) error {
	return tx.batch.Set(key, value, nil)
}

func (tx *WriteTx) Delete(key []byte) error {
	return tx.batch.Delete(key, nil)
}

func (tx *WriteTx) Commit() error {
	if tx.batch == nil {
		return nil
	}
	err := tx.batch.Commit(nil)
	tx.batch = nil
	return err
}

func (tx *WriteTx) Discard() {
	if tx.batch == nil {
		return
	}
	_ = tx.batch.Close()
	tx.batch = nil
}

func keyUpperBound(prefix []byte) []byte {
	end := bytes.Clone(prefix)
	for i := len(end) - 1; i >= 0; i-- {
		end[i]++
		if end[i] != 0 {
			return end[:i+1]
		}
	}
	return nil
}
