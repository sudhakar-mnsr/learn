/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package max
	// Max returns the maximum integer in the slice.
	func Max(n []int) (int, error)
*/


package max_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/slices/max"
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
