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
			bucket[idx].value = value
			return
		}
	}

	// This key does not exist, so add this new value.
	h.buckets[idx] = append(bucket, entry{key, value})
}

// Retrieve extracts a value from the hash table based on the key.
func (h *Hash) Retrieve(key string) (int, error) {

	// For the specified key, identify what bucket in
	// the slice we need to store the key/value inside of.
	idx := h.hashKey(key)

	// Iterate over the entries for the specified bucket.
	for _, entry := range h.buckets[idx] {

		// Compare the keys and if there is a match return
		// the value associated with the key.
		if entry.key == key {
			return entry.value, nil
		}
	}

	// The key was not found so return the error.
	return 0, fmt.Errorf("%q not found", key)
}

// Delete deletes an entry from the hash table.
func (h *Hash) Delete(key string) error {

	// For the specified key, identify what bucket in
	// the slice we need to store the key/value inside of.
	bucketIdx := h.hashKey(key)

	// Extract a copy of the bucket from the hash table.
	bucket := h.buckets[bucketIdx]
