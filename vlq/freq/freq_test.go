/*
	// This is the API you need to build for these tests. You will need to
	// change the import path in this test to point to your code.
	package freq
	// Sequential uses a sequential algorithm.
	func Sequential(text []string) map[rune]int

	// ConcurrentUnlimited uses a concurrent algorithm based on an
	// unlimited fan out pattern.
	func ConcurrentUnlimited(text []string) map[rune]int
	// ConcurrentBounded uses a concurrent algorithm based on a bounded
	// fan out and no channels.
	func ConcurrentBounded(text []string) map[rune]int

	// ConcurrentBoundedChannel uses a concurrent algorithm based on a bounded
	// fan out using a channel.
	func ConcurrentBoundedChannel(text []string) map[rune]int
*/

package freq_test

import (
	"testing"

	"github.com/ardanlabs/gotraining/topics/go/algorithms/fun/freq"
)

// go test -run none -bench . -benchtime 3s
