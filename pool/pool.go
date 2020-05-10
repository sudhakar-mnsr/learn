// Example provided with help from Fatih Arslan and Gabriel Aszalos.

// Package pool manages a user defined set of resources.
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool manages a set of resources that can be shared safely by
// multiple goroutines. The resource being managed must implement
// the io.Closer interface.
type Pool struct {
	mu        sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed is returned when an Acquire returns on a
// closed pool.
var ErrPoolClosed = errors.New("Pool has been closed")

// New creates a pool that manages resources. A pool requires a
// function that can allocate a new resource and the size of
// the pool.
func New(size uint, f func() (io.Closer, error)) (*Pool, error) {
	if size == 0 {
		return nil, errors.New("Size value too small")
	}

	return &Pool{
		factory:   f,
		resources: make(chan io.Closer, size),
	}, nil
}
