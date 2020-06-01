package main

// Tree represents all values in the tree.
type Tree struct {
	root *Node
}

// Insert adds a value into the tree.
func (t *Tree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{value: value}
