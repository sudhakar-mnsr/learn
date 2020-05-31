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

	numbers = generateList(1e2)
	fmt.Println("Before:", numbers)
	bubbleSortConcurrent(runtime.GOMAXPROCS(0), numbers)
	fmt.Println("Concurrent:", numbers)
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)
	for i := 0; i < totalNumbers; i++ {
