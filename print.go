package main

import (
	"fmt"
	"math"
)

// PrettyPrint takes a Tree value and displays a pretty print
// version of the tree.
func PrettyPrint(t *Tree) {

	// Build an index map of positions for print layout.
	values := make(map[int]int)
	maxIdx := buildIndexMap(values, 0, 0, t.root)
