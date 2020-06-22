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
