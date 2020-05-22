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
	q := Queue{
		front: 0,
		end:   0,
		data:  make([]*Data, cap),
	}
	return &q, nil
}

// Enqueue inserts data into the queue if there
// is available capacity.
func (q *Queue) Enqueue(data *Data) error {

	// If the front of the queue is right behind the end or
	// if the front is at the end of the capacity and the end
	// is at the beginning of the capacity, the queue is full.
	//  F  E  - Enqueue (Full) |  E        F - Enqueue (Full)
	// [A][B][C]               | [A][B][C]
	if q.front+1 == q.end ||
		q.front == len(q.data) && q.end == 0 {
		return errors.New("queue at capacity")
	}
