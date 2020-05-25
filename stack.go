// Package stack asks the student to implement a stack in Go.
package stack

import "errors"

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
func Make(cap int) *Stack {
	return &Stack{
		data: make([]*Data, 0, cap),
	}
}

// Count returns the number of items in the stack.
func (s *Stack) Count() int {
	return len(s.data)
}

// Push adds data into the top of the stack.
func (s *Stack) Push(data *Data) {
	s.data = append(s.data, data)
}
