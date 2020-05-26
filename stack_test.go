/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package stack
	// Data represents what is being stored on the stack.
	type Data struct {
		Name string
	}

	// Stack represents a stack of data.
	type Stack struct {
		data []*Data
	}
	// Make allows the creation of a stack with an initial
	// capacity for efficiency. Otherwise a stack can be
	// used in its zero value state.
	func Make(cap int) *Stack

	// Count returns the number of items in the stack.
	func (s *Stack) Count() int
	// Push adds data into the top of the stack.
	func (s *Stack) Push(data *Data)
	// Pop removes data from the top of the stack.
	func (s *Stack) Pop() (*Data, error)

	// Peek provides the data stored on the stack based
	// on the level from the bottom. A value of 0 would
	// return the top piece of data.
	func (s *Stack) Peek(level int) (*Data, error)
	// Operate accepts a function that takes data and calls
	// the specified function for every piece of data found.
	// It traverses from the top down through the stack.
	func (s *Stack) Operate(f func(data *Data) error) error
*/

package stack_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/stack"
)
