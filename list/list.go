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
