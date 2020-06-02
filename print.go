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

	// Set the edge of the top of the tree.
	for sp := 0; sp < data[0].edge; sp++ {
		fmt.Print(" ")
	}
	fmt.Printf("%02d", values[0])
	fmt.Print("\n")

	dataIdx := 1
	for i := 1; i < len(data); i = i + 2 {

		// Set the edge of this row.
		for sp := 0; sp < data[i].edge; sp++ {
			fmt.Print(" ")
		}
