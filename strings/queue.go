// Package queue implements of a circular queue.
package queue

import (
	"errors"
)

// Data represents what is being stored on the queue.
type Data struct {
	Name string
}

