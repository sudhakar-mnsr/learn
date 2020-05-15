// Package hash implements a hash table.
package hash

import (
	"fmt"
	"hash/maphash"
)

const numBuckets = 256

// An entry where we store key and value in the hash.
type entry struct {
	key   string
	value int
}

// Hash is a simple Hash table implementation.
type Hash struct {
	buckets [][]entry
	hash    maphash.Hash
}

// New returns a new hash table.
func New() *Hash {
	return &Hash{
		buckets: make([][]entry, numBuckets),
	}
}
