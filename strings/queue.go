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
