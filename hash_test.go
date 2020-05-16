
/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package hash
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
	func (h *Hash) Store(key string, value int)
	// Retrieve extracts a value from the hash table based on the key.
	func (h *Hash) Retrieve(key string) (int, error)
	// Delete deletes an entry from the hash table.
	func (h *Hash) Delete(key string) error
	// Len return the number of elements in the hash.
	func (h *Hash) Len() int
	// Do calls fn on each key/value. If fn return false stops the iteration.
	func (h *Hash) Do(fn func(key string, value int) bool)
*/

package hash_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/hash"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestHash(t *testing.T) {
	t.Log("Given the need to test hash functionality.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen checking basic hashing operations", testID)
		{
