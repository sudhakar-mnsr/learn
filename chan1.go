// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create an unbuffered channel.
	share := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.

	var wg sync.WaitGroup
	wg.Add(2)

	// Launch two goroutines.

	go func() {
		goroutine("Bill", share)
		wg.Done()
	}()
