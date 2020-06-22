// Answer for exercise 1 of Race Conditions.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// numbers maintains a set of random numbers.
var numbers []int

// mutex will help protect the slice.
var mutex sync.Mutex

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main is the entry point for the application.
func main() {
	// Number of goroutines to use.
	const grs = 3

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)
