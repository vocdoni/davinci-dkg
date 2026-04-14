package db

import (
	"fmt"
	"io"
)

const (
	TypePebble = "pebble"
)

var ErrKeyNotFound = fmt.Errorf("key not found")

type Options struct {
	Path string
}

type Database interface {
	io.Closer
	Reader
	WriteTx() WriteTx
	Compact() error
}

type Reader interface {
	Get(key []byte) ([]byte, error)
	Iterate(prefix []byte, callback func(key, value []byte) bool) error
}

type WriteTx interface {
	Reader
	Set(key, value []byte) error
	Delete(key []byte) error
	Commit() error
	Discard()
}
