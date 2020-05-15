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

// Store adds a value in the hash table based on the key.
func (h *Hash) Store(key string, value int) {

	// For the specified key, identify what bucket in
	// the slice we need to store the key/value inside of.
	idx := h.hashKey(key)

	// Extract a copy of the bucket from the hash table.
	bucket := h.buckets[idx]
	// Iterate over the indexes for the specified bucket.
	for idx := range bucket {

		// Compare the keys and if there is a match replace the
		// existing entry value for the new value.
		if bucket[idx].key == key {
