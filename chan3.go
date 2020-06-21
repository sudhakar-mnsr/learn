// Write a program that uses goroutines to generate up to 100 random numbers.
// Do not send values that are divisible by 2. Have the main goroutine receive
// values and add them to a slice.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(goroutines)

	// Iterate and launch each goroutine.
	for gr := 0; gr < goroutines; gr++ {
		// Create an anonymous function for each goroutine.
		go func() {

			// Ensure the waitgroup is decremented when this function returns.
			defer wg.Done()
