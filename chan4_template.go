// Fix the race condition in this program.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// numbers maintains a set of random numbers.
var numbers []int

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}
