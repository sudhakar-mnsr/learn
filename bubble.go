// Implementation of Bubble sort in Go.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func main() {
	numbers := generateList(1e2)
	fmt.Println("Before:", numbers)
	bubbleSort(numbers)
	fmt.Println("Sequential:", numbers)
