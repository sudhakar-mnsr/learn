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
