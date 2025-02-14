package store

import (
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
)

// KVStore encapsulates key-value store abstraction, in minimalistic interface.
//
// KVStore MUST be thread safe.
type KVStore interface {
	Get(key []byte) ([]byte, error)     // Get gets the value for a key.
	Set(key []byte, value []byte) error // Set updates the value for a key.
	Delete(key []byte) error            // Delete deletes a key.
	NewBatch() Batch
}

// Batch enables batching of transactions
type Batch interface {
	Set(key, value []byte) error // Accumulates KV entries in a transaction
	Delete(key []byte) error     // Deletes the given key
	Commit() error               // Commits the transaction
	Discard()                    // Discards the transaction
}

// NewInMemoryKVStore builds KVStore that works in-memory (without accessing disk).
func NewDefaultInMemoryKVStore() KVStore {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic(err)
	}
	return &BadgerKV{
		db: db,
	}
}

func NewDefaultKVStore(rootDir, dbPath, dbName string) KVStore {
	path := filepath.Join(rootify(rootDir, dbPath), dbName)
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}
	return &BadgerKV{
		db: db,
	}
}

// rootify works just like in cosmos-sdk
func rootify(rootDir, dbPath string) string {
	if filepath.IsAbs(dbPath) {
		return dbPath
	}
	return filepath.Join(rootDir, dbPath)
}
