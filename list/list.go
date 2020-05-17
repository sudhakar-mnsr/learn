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
