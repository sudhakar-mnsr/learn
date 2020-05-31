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
		numbers[i] = rand.Intn(totalNumbers)
	}
	return numbers
}

func bubbleSort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func bubbleSortConcurrent(goroutines int, numbers []int) {
	totalNumbers := len(numbers)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
