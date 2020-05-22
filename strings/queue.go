// Package queue implements of a circular queue.
package queue

import (
	"errors"
)

// Data represents what is being stored on the queue.
type Data struct {
	Name string
}

// Queue represents a list of data.
type Queue struct {
	Count int
	data  []*Data
	front int
	end   int
}

// New returns a queue with a set capacity.
func New(cap int) (*Queue, error) {
	if cap <= 0 {
		return nil, errors.New("invalid capacity")
	}
