// Package list implements of a doubly link list in Go.
package list

import (
	"fmt"
	"strings"
)

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
func (l *List) Add(data string) *Node {

	// When creating the new node, have the new node
	// point to the last node in the list.
	n := Node{
		Data: data,
		prev: l.last,
	}

	// Increment the count for the new node.
	l.Count++

	// If this is the first node, attach it.
	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	// Fix the fact that the last node does not point back to the NEW node.
	l.last.next = &n

	//              First                                       Last           l.last.next
	//                V                                           V                    V
	// nil <- Prev.[Node0].Next <-> Prev.[Node1].Next <-> Prev.[Node2].Next <-> Prev.[NEW].Next -> nil

	// Fix the fact the Last pointer is not pointing to the true end of the list.

	l.last = &n

	//              First                                       Last                 Last
	//                V                                           V  ----> MOVE ---->  V
	// nil <- Prev.[Node0].Next <-> Prev.[Node1].Next <-> Prev.[Node2].Next <-> Prev.[NEW].Next -> nil

	return &n
}

// AddFront places a new node at the front of the list.
func (l *List) AddFront(data string) *Node {

	// When creating the new node, have the new node
	// point to the first node in the list.
	n := Node{
		Data: data,
		next: l.first,
	}

	// Increment the count for the new node.
	l.Count++

	// If this is the first node, attach it.
	if l.first == nil && l.last == nil {
		l.first = &n
		l.last = &n
		return &n
	}

	// Fix the fact that the first node does not point back to the NEW node.
	l.first.prev = &n

	//      l.first.prev                First                                       Last
	//               V                    V                                           V
	// nil <- Prev.[NEW].Next <-> Prev.[Node2].Next <-> Prev.[Node1].Next <-> Prev.[Node0].Next -> nil

	// Fix the fact the First pointer is not pointing to the true beginning of the list.
