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
	// Remove traverses the list looking for the specified data
	// and if found, removes the node from the list.
	func (l *List) Remove(data string) (*Node, error)
	// Operate accepts a function that takes a node and calls
	// the specified function for every node found.
	func (l *List) Operate(f func(n *Node) error) error
	// OperateReverse accepts a function that takes a node and
	// calls the specified function for every node found.
	func (l *List) OperateReverse(f func(n *Node) error) error
	// AddSort adds a node based on lexical ordering.
	func (l *List) AddSort(data string) *Node
*/

package list_test

import (
	"fmt"
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/data/list"
)

const succeed = "\u2713"
const failed = "\u2717"

// TestAdd validates the Add functionality.
