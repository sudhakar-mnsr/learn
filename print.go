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

		// Draw the hashes for this row.
		dataHashIdx := dataIdx
		for h := 0; h < data[i].draw; h++ {
			if values[dataHashIdx] != maxInt {
				fmt.Printf("/")
			} else {
				fmt.Printf(" ")
			}
			for sp := 0; sp < data[i].padding; sp++ {
				fmt.Print(" ")
			}
			if values[dataHashIdx+1] != maxInt {
				fmt.Printf("\\")
			} else {
				fmt.Printf(" ")
			}
			dataHashIdx += 2

			if data[i].gaps != 0 && data[i].gaps > h {
				for sp := 0; sp < data[i].gapPad; sp++ {
					fmt.Print(" ")
				}
			}
		}
		fmt.Print("\n")

		// Set the edge of the next row.
		for sp := 0; sp < data[i+1].edge; sp++ {
			fmt.Print(" ")
		}

		// Draw the numbers for this row.
		for n := 0; n < data[i+1].draw; n++ {
			if values[dataIdx] != maxInt {
				fmt.Printf("%02d", values[dataIdx])
			} else {
				fmt.Printf("  ")
			}

			for sp := 0; sp < data[i+1].padding; sp++ {
				fmt.Print(" ")
			}
			if values[dataIdx+1] != maxInt {
				fmt.Printf("%02d", values[dataIdx+1])
			} else {
				fmt.Printf("  ")
			}
			dataIdx += 2

			if data[i+1].gaps != 0 && data[i+1].gaps > n {
				for sp := 0; sp < data[i+1].gapPad; sp++ {
