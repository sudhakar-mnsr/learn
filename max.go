package max

import "fmt"

// Max returns the maximum integer in the slice.
func Max(n []int) (int, error) {
	// First check there are numbers in the collection.
	if len(n) == 0 {
		return 0, fmt.Errorf("slice %#v has no elements", n)
	}
