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

// The interesting code here is how V is used to satisfy type T
// for the SliceConstraint interface.
func operate(type T SliceConstraint(V), V interface{})(slice T, fn operateFunc(V)) T {
	ret := make(T, len(slice))
	for i, v := range slice {
		ret[i] = fn(v)
	}
	return ret
}

// =============================================================================

type Numbers []int

func Double(n Numbers) Numbers {
	fn := func(v int) int {
		return 2 * v
	}
	// The compiler can infer from n that V in the declaration of the
	// operate function will represent an integer when constructing the
	// operate function.
	v := operate(n, fn)
	return v
}
