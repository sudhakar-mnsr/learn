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
