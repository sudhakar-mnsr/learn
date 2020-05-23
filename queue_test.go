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

			if err == nil {
				t.Fatalf("\t%s\tTest 0:\tShould not be able to create a queue for %d items : %v", failed, cap, err)
			}
			t.Logf("\t%s\tTest 0:\tShould not be able to create a queue for %d items.", succeed, cap)
		}
	}
}

// TestEnqueue validates the Enqueue functionality.
func TestEnqueue(t *testing.T) {
	t.Log("Given the need to test Enqueue functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen enqueuing %d items", items)
		{
			q, err := queue.New(items)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a queue for %d items : %v", failed, items, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a queue for %d items.", succeed, items)

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData += name
				if err := q.Enqueue(&queue.Data{Name: name}); err != nil {
					t.Fatalf("\t%s\tTest 0:\tShould be able to enqueue item %d in the queue : %v", failed, i, err)
				}
			}

			if q.Count != items {
				t.Logf("\t%s\tTest 0:\tShould be able to enqueue %d items.", failed, items)
				t.Fatalf("\t\tTest 0:\tGot %d, Expected %d.", q.Count, items)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to enqueue %d items.", succeed, items)

			var data string
			f := func(d *queue.Data) error {
				data += d.Name
				return nil
			}

			if err := q.Operate(f); err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to operate on the queue : %v", failed, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to operate on the queue.", succeed)

			if data != orgData {
				t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d items in FIFO order.", failed, items)
				t.Fatalf("\t\tTest 0:\tGot %s, Expected %s.", data, orgData)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to traverse over %d items in FIFO order.", succeed, items)
		}
	}
}

// TestDequeue validates the Dequeue functionality.
func TestDequeue(t *testing.T) {
	t.Log("Given the need to test Dequeue functionality.")
	{
		const items = 5
		t.Logf("\tTest 0:\tWhen dequeuing %d items", items)
		{
			q, err := queue.New(items)
			if err != nil {
				t.Fatalf("\t%s\tTest 0:\tShould be able to create a queue for %d items : %v", failed, items, err)
			}
			t.Logf("\t%s\tTest 0:\tShould be able to create a queue for %d items.", succeed, items)

			var orgData string
			for i := 0; i < items; i++ {
				name := fmt.Sprintf("Name%d", i)
				orgData += name
				if err := q.Enqueue(&queue.Data{Name: name}); err != nil {
					t.Fatalf("\t%s\tTest 0:\tShould be able to enqueue item %d in the queue : %v", failed, i+1, err)
				}
			}