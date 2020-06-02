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

	// Calculate the total number of levels based on
	// the max index provided.
	var levels int
	for {
		pow := math.Pow(2, float64(levels))
		if maxIdx < int(pow) {
                   break
		}
		levels++
	}
	levels--

	// Capture the positional data to use.
	data := generateData(levels)
