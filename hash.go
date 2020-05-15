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
