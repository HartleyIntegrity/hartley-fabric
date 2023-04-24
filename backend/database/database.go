package database

import (
	"sync"
)

type Database struct {
	store map[string][]byte
	mu    sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		store: make(map[string][]byte),
	}
}

func (db *Database) Get(key string) ([]byte, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	value, ok := db.store[key]
	return value, ok
}

func (db *Database) Put(key string, value []byte) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.store[key] = value
}

func (db *Database) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.store, key)
}
