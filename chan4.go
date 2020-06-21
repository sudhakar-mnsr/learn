// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {

	// Create the channel for sharing results.
	values := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	poolSize := runtime.GOMAXPROCS(0)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(poolSize)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < poolSize; i++ {
		go func(id int) {

			// Start an infinite loop.
			for {
