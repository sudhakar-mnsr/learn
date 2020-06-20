// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffer channel so no send every blocks. Don't allocate more buffers than
// you need. Have the main goroutine display each random number is receives and
// then terminate the program.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	goroutines = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
