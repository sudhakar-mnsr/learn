/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package list
	// Node represents the data being stored.
	type Node struct {
		Data string
		next *Node
		prev *Node
	}

	// List represents a list of nodes.
	type List struct {
		Count int
		first *Node
		last  *Node
	}


	// Add places a new node at the end of the list.
	func (l *List) Add(data string) *Node
	// AddFront places a new node at the front of the list.
	func (l *List) AddFront(data string) *Node
	// Find traverses the list looking for the specified data.
	func (l *List) Find(data string) (*Node, error)
	// FindReverse traverses the list in the opposite direction
	// looking for the specified data.
	func (l *List) FindReverse(data string) (*Node, error)
