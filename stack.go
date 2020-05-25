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
