/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package min
	// Min returns the minimum integer in the slice.
	func Min(n []int) (int, error)
*/

package min_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/slices/min"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestMax(t *testing.T) {
	tt := []struct {
		name     string
		input    []int
		expected int
		success  bool
	}{
		{"empty", []int{}, 0, false},
		{"nil", nil, 0, false},
		{"one", []int{10}, 10, true},
		{"even", []int{20, 30, 10, 50}, 10, true},
		{"odd", []int{30, 50, 10}, 10, true},
	}
