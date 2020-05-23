/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package queue
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
	func New(cap int) (*Queue, error)
	// Enqueue inserts data into the queue if there
	// is available capacity.
	func (q *Queue) Enqueue(data *Data) error

	// Dequeue removes data into the queue if data exists.
	func (q *Queue) Dequeue() (*Data, error)
	// Operate accepts a function that takes data and calls
	// the specified function for every piece of data found.
	func (q *Queue) Operate(f func(d *Data) error) error
*/

package queue_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/queue"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestNew validates the New functionality.
func TestNew(t *testing.T) {
	t.Log("Given the need to test New functionality.")
	{
		t.Logf("\tTest 0:\tWhen creating a new queue with invalid capacity.")
		{
			var cap int
			_, err := queue.New(cap)
			if err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould not be able to create a queue for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%s\tTest 0:\tShould not be able to create a queue for %d items.", succeed, cap)

			cap = -1
			_, err = queue.New(cap)
