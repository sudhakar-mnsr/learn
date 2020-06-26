// This sample program demonstrates how the logger package works.
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

// device allows us to mock a device we write logs to.
type device struct {
	mu      sync.RWMutex
	problem bool
}

// Write implements the io.Writer interface.
func (d *device) Write(p []byte) (n int, err error) {

	// Simulate disk problems.
	for d.isProblem() {

		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
}

// isProblem checks in a safe way if there is a problem.
func (d *device) isProblem() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.problem
}
