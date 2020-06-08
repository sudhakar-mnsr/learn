// Package freq provides support for find the frequency in which a rune
// is found in a collection of text documents.
package freq

import (
	"runtime"
	"sync"
)

// Sequential uses a sequential algorithm.
func Sequential(text []string) map[rune]int {
	m := make(map[rune]int)
	for _, words := range text {
		for _, r := range words {
			m[r]++
		}
	}
	return m
}
