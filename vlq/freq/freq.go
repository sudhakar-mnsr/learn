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

// ConcurrentUnlimited uses a concurrent algorithm based on an
// unlimited fan out pattern.
func ConcurrentUnlimited(text []string) map[rune]int {
	ch := make(chan map[rune]int, len(text))
	for _, words := range text {
		go func(words string) {
			lm := make(map[rune]int)
			for _, r := range words {
				lm[r]++
			}
			ch <- lm
		}(words)
	}

	all := make(map[rune]int)

	for range text {
		lm := <-ch
		for r, c := range lm {
			all[r] += c
		}
	}

	return all
}

// ConcurrentBounded uses a concurrent algorithm based on a bounded
// fan out and no channels.
func ConcurrentBounded(text []string) map[rune]int {
	m := make(map[rune]int)
