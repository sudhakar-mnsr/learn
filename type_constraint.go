package main

import (
	"fmt"
)

func main() {
	n := Numbers{10, 20, 30, 40, 50}
	fmt.Println(n)

	n = Double(n)
	fmt.Println(n)
}

// =============================================================================

type operateFunc(type V) func(v V) V

type SliceConstraint(type T) interface {
	type []T
}
